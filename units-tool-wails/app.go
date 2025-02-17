package main

import (
	"log"
	"os/exec"
	"context"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetFerramentasVersion() string {
  cmd := exec.Command("./ferramentas.exe", "-v")
	cmd.Dir = "frontend/src"

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Erro ao obter a versão do ferramentas: ", err.Error())
		return "Ferramentas não encontrado."
	}

  return string(output)
}