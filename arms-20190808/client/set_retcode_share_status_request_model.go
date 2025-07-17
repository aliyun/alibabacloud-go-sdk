// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRetcodeShareStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SetRetcodeShareStatusRequest
	GetAppName() *string
	SetPid(v string) *SetRetcodeShareStatusRequest
	GetPid() *string
	SetStatus(v bool) *SetRetcodeShareStatusRequest
	GetStatus() *bool
}

type SetRetcodeShareStatusRequest struct {
	// The name of the application that is monitored by Browser Monitoring.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The process identifier (PID) of the application.
	//
	// Log on to the **ARMS console**. In the left-side navigation pane, choose **Browser Monitoring*	- > **Browser Monitoring**. On the Browser Monitoring page, click the name of an application. The URL in the address bar contains the process ID (PID) of the application. The PID is indicated in the `pid=xxx` format. The PID is usually percent encoded as `xxx%40xxx`. You must modify this value to remove the percent encoding. For example, if the PID in the URL is `eb4zdose6v%409781be0f44d****`, you must replace `%40` with @ to obtain `eb4zdose6v@9781be0f44d****`.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// Specifies whether to turn on or turn off the logon-free sharing switch. Valid values:
	//
	// 	- `true`: Turn on the switch.
	//
	// 	- `false`: Turn off the switch.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetRetcodeShareStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRetcodeShareStatusRequest) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetRetcodeShareStatusRequest) GetPid() *string {
	return s.Pid
}

func (s *SetRetcodeShareStatusRequest) GetStatus() *bool {
	return s.Status
}

func (s *SetRetcodeShareStatusRequest) SetAppName(v string) *SetRetcodeShareStatusRequest {
	s.AppName = &v
	return s
}

func (s *SetRetcodeShareStatusRequest) SetPid(v string) *SetRetcodeShareStatusRequest {
	s.Pid = &v
	return s
}

func (s *SetRetcodeShareStatusRequest) SetStatus(v bool) *SetRetcodeShareStatusRequest {
	s.Status = &v
	return s
}

func (s *SetRetcodeShareStatusRequest) Validate() error {
	return dara.Validate(s)
}
