package app

type App struct{}

type Builder struct{}

func NewAppBuilder() Builder {
	return Builder{}
}

func (Builder) Build() App {
	return App{}
}

func (a *App) StartListenAndServe() error {
	return nil
}
