package main

import (
	"context"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	tokenBuff, err := os.ReadFile("./token.txt")
	if err != nil {
		panic(err)
	}

	b, err := bot.New(string(tokenBuff), opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, helpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/toadmin", bot.MatchTypeExact, toadminHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/nextLesson", bot.MatchTypeExact, nextLessonHandler)
	// Для дня недели
	b.RegisterHandler(bot.HandlerTypeMessageText, "/scheduleOn", bot.MatchTypeExact, scheduleOnHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/scheduleTomorrow", bot.MatchTypeExact, scheduleTomorrowHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/scheduleToday", bot.MatchTypeExact, scheduleTodayHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/comment", bot.MatchTypeExact, commentHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/whereStudents", bot.MatchTypeExact, whereStudentsHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/whereTeacher", bot.MatchTypeExact, whereTeacherHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/whenExam", bot.MatchTypeExact, whenExamHandler)

	b.Start(ctx)
}

func whenExamHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func whereTeacherHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func whereStudentsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func commentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func scheduleTodayHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func scheduleTomorrowHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func scheduleOnHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func nextLessonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func toadminHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Команда ещё не реализована",
	})
}

func helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	commands := []string{
		"/help",
		"/toadmin",
		"/nextLesson",
		"/scheduleOn",
		"/scheduleTomorrow",
		"/scheduleToday",
		"/comment",
		"/whereStudents",
		"/whereTeacher",
		"/whenExam",
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Доступные команды:\n" + strings.Join(commands, ", "),
	})
}

func helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Привет, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*",
		ParseMode: models.ParseModeMarkdown,
	})
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Используйте /help для списка команд",
	})
}
