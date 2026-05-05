package app

import (
	ausf_context "github.com/f0lkert/ausf/internal/context"
	"github.com/f0lkert/ausf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *ausf_context.AUSFContext
	Config() *factory.Config
}
