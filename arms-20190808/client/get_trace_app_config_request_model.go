// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *GetTraceAppConfigRequest
	GetPid() *string
}

type GetTraceAppConfigRequest struct {
	// The process ID (PID) of the application.
	//
	// You can use one of the following methods to obtain the PID:
	//
	// 	- API: Call the [ListTraceApps](https://help.aliyun.com/document_detail/2588008.html) operation
	//
	// 	- Console: Log on to the Application Real-Time Monitoring Service (ARMS) console. In the left-side navigation pane, choose **Application Monitoring*	- > **Application List**. On the Application List page, click the name of your application. The URL in the address bar contains the PID of the application. The PID is indicated in the pid=xxx format. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is xxx%4074xxx, you must replace %40 with an at sign (@) to obtain xxx@74xxx.
	//
	// >
	//
	// 	- To obtain the PID in the console, your application must be monitored by Application Monitoring rather than Managed Service for OpenTelemetry.
	//
	// 	- The GetTraceAppConfig operation can query only the custom settings of applications that are monitored by Application Monitoring.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2n80plglh@745eddxxx
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetTraceAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppConfigRequest) GoString() string {
	return s.String()
}

func (s *GetTraceAppConfigRequest) GetPid() *string {
	return s.Pid
}

func (s *GetTraceAppConfigRequest) SetPid(v string) *GetTraceAppConfigRequest {
	s.Pid = &v
	return s
}

func (s *GetTraceAppConfigRequest) Validate() error {
	return dara.Validate(s)
}
