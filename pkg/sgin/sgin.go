package sgin

import (
	"github.com/gin-gonic/gin"
)

type Srv struct {
	engine      *gin.Engine
	groups      map[string]*gin.RouterGroup
	middlewares []gin.HandlerFunc
	addr        string
	mode        string
}

func New() *Srv {
	engine := gin.New()
	return &Srv{
		engine:      engine,
		groups:      make(map[string]*gin.RouterGroup),
		middlewares: nil,
		addr:        ":8080", // 默认端口
		mode:        "debug", // 默认模式
	}
}

func (s *Srv) WithAddr(addr string) *Srv {
	s.addr = addr
	return s
}

func (s *Srv) WithMode(mode string) *Srv {
	s.mode = mode
	gin.SetMode(mode)
	return s
}

func (s *Srv) Engine() *gin.Engine {
	return s.engine
}

func (s *Srv) Start() error {
	return s.engine.Run(s.addr)
}

func (s *Srv) AddHandlers(handlers ...func(engine *gin.Engine, groups map[string]*gin.RouterGroup)) *Srv {
	for _, handler := range handlers {
		handler(s.engine, s.groups)
	}
	return s
}

func (s *Srv) AddGroup(name string) *gin.RouterGroup {
	group := s.engine.Group(name)
	s.groups[name] = group
	return group
}

func (s *Srv) Use(mws ...gin.HandlerFunc) {
	s.middlewares = append(s.middlewares, mws...)
	s.engine.Use(mws...)
}

func (s *Srv) SetupRouter(setupFunc func(*gin.Engine)) *Srv {
	setupFunc(s.engine)
	return s
}
