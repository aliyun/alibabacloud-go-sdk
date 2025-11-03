// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentlessLlmService(v int32) *DescribePropertyCountResponseBody
	GetAgentlessLlmService() *int32
	SetAgentlessScaAiComponent(v int32) *DescribePropertyCountResponseBody
	GetAgentlessScaAiComponent() *int32
	SetAutorun(v int32) *DescribePropertyCountResponseBody
	GetAutorun() *int32
	SetCron(v int32) *DescribePropertyCountResponseBody
	GetCron() *int32
	SetDatabase(v int32) *DescribePropertyCountResponseBody
	GetDatabase() *int32
	SetLkm(v int32) *DescribePropertyCountResponseBody
	GetLkm() *int32
	SetPort(v int32) *DescribePropertyCountResponseBody
	GetPort() *int32
	SetProcess(v int32) *DescribePropertyCountResponseBody
	GetProcess() *int32
	SetRequestId(v string) *DescribePropertyCountResponseBody
	GetRequestId() *string
	SetSca(v int32) *DescribePropertyCountResponseBody
	GetSca() *int32
	SetSoftware(v int32) *DescribePropertyCountResponseBody
	GetSoftware() *int32
	SetUser(v int32) *DescribePropertyCountResponseBody
	GetUser() *int32
	SetWeb(v int32) *DescribePropertyCountResponseBody
	GetWeb() *int32
	SetWebserver(v int32) *DescribePropertyCountResponseBody
	GetWebserver() *int32
}

type DescribePropertyCountResponseBody struct {
	// The number of AI services.
	//
	// example:
	//
	// 3
	AgentlessLlmService *int32 `json:"AgentlessLlmService,omitempty" xml:"AgentlessLlmService,omitempty"`
	// The number of AI tools.
	//
	// example:
	//
	// 13
	AgentlessScaAiComponent *int32 `json:"AgentlessScaAiComponent,omitempty" xml:"AgentlessScaAiComponent,omitempty"`
	// The number of startup items.
	//
	// example:
	//
	// 3
	Autorun *int32 `json:"Autorun,omitempty" xml:"Autorun,omitempty"`
	// The number of scheduled tasks.
	//
	// example:
	//
	// 123
	Cron *int32 `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The number of databases.
	//
	// example:
	//
	// 2
	Database *int32 `json:"Database,omitempty" xml:"Database,omitempty"`
	// The number of kernel modules.
	//
	// example:
	//
	// 4
	Lkm *int32 `json:"Lkm,omitempty" xml:"Lkm,omitempty"`
	// The number of ports.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The number of processes.
	//
	// example:
	//
	// 367
	Process *int32 `json:"Process,omitempty" xml:"Process,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of middleware assets.
	//
	// example:
	//
	// 112
	Sca *int32 `json:"Sca,omitempty" xml:"Sca,omitempty"`
	// The number of software assets.
	//
	// example:
	//
	// 111
	Software *int32 `json:"Software,omitempty" xml:"Software,omitempty"`
	// The number of accounts.
	//
	// example:
	//
	// 214
	User *int32 `json:"User,omitempty" xml:"User,omitempty"`
	// The number of websites.
	//
	// example:
	//
	// 65
	Web *int32 `json:"Web,omitempty" xml:"Web,omitempty"`
	// The number of web services.
	//
	// example:
	//
	// 8
	Webserver *int32 `json:"Webserver,omitempty" xml:"Webserver,omitempty"`
}

func (s DescribePropertyCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyCountResponseBody) GetAgentlessLlmService() *int32 {
	return s.AgentlessLlmService
}

func (s *DescribePropertyCountResponseBody) GetAgentlessScaAiComponent() *int32 {
	return s.AgentlessScaAiComponent
}

func (s *DescribePropertyCountResponseBody) GetAutorun() *int32 {
	return s.Autorun
}

func (s *DescribePropertyCountResponseBody) GetCron() *int32 {
	return s.Cron
}

func (s *DescribePropertyCountResponseBody) GetDatabase() *int32 {
	return s.Database
}

func (s *DescribePropertyCountResponseBody) GetLkm() *int32 {
	return s.Lkm
}

func (s *DescribePropertyCountResponseBody) GetPort() *int32 {
	return s.Port
}

func (s *DescribePropertyCountResponseBody) GetProcess() *int32 {
	return s.Process
}

func (s *DescribePropertyCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyCountResponseBody) GetSca() *int32 {
	return s.Sca
}

func (s *DescribePropertyCountResponseBody) GetSoftware() *int32 {
	return s.Software
}

func (s *DescribePropertyCountResponseBody) GetUser() *int32 {
	return s.User
}

func (s *DescribePropertyCountResponseBody) GetWeb() *int32 {
	return s.Web
}

func (s *DescribePropertyCountResponseBody) GetWebserver() *int32 {
	return s.Webserver
}

func (s *DescribePropertyCountResponseBody) SetAgentlessLlmService(v int32) *DescribePropertyCountResponseBody {
	s.AgentlessLlmService = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetAgentlessScaAiComponent(v int32) *DescribePropertyCountResponseBody {
	s.AgentlessScaAiComponent = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetAutorun(v int32) *DescribePropertyCountResponseBody {
	s.Autorun = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetCron(v int32) *DescribePropertyCountResponseBody {
	s.Cron = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetDatabase(v int32) *DescribePropertyCountResponseBody {
	s.Database = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetLkm(v int32) *DescribePropertyCountResponseBody {
	s.Lkm = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetPort(v int32) *DescribePropertyCountResponseBody {
	s.Port = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetProcess(v int32) *DescribePropertyCountResponseBody {
	s.Process = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetRequestId(v string) *DescribePropertyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetSca(v int32) *DescribePropertyCountResponseBody {
	s.Sca = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetSoftware(v int32) *DescribePropertyCountResponseBody {
	s.Software = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetUser(v int32) *DescribePropertyCountResponseBody {
	s.User = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetWeb(v int32) *DescribePropertyCountResponseBody {
	s.Web = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetWebserver(v int32) *DescribePropertyCountResponseBody {
	s.Webserver = &v
	return s
}

func (s *DescribePropertyCountResponseBody) Validate() error {
	return dara.Validate(s)
}
