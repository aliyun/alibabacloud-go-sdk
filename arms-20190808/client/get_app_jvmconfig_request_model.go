// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppJVMConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetAppJVMConfigRequest
	GetEndTime() *int64
	SetPid(v string) *GetAppJVMConfigRequest
	GetPid() *string
	SetRegionId(v string) *GetAppJVMConfigRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetAppJVMConfigRequest
	GetStartTime() *int64
}

type GetAppJVMConfigRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1480607940000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDof the application.
	//
	//
	//
	// Log on to the **ARMS console**. In the left-side navigation pane, choose **Application Monitoring*	- > **Applications**. On the **Applications*	- page, click the name of an application. The URL in the address bar contains the process ID (PID) of the application. The PID is indicated in the pid=xxx format. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is eb4zdose6v%409781be0f44d\\*\\*\\*\\*, you must replace %40 with an at sign (@) to obtain eb4zdose6v@9781be0f44d\\*\\*\\*\\*.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1480521600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAppJVMConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppJVMConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAppJVMConfigRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAppJVMConfigRequest) GetPid() *string {
	return s.Pid
}

func (s *GetAppJVMConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAppJVMConfigRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAppJVMConfigRequest) SetEndTime(v int64) *GetAppJVMConfigRequest {
	s.EndTime = &v
	return s
}

func (s *GetAppJVMConfigRequest) SetPid(v string) *GetAppJVMConfigRequest {
	s.Pid = &v
	return s
}

func (s *GetAppJVMConfigRequest) SetRegionId(v string) *GetAppJVMConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetAppJVMConfigRequest) SetStartTime(v int64) *GetAppJVMConfigRequest {
	s.StartTime = &v
	return s
}

func (s *GetAppJVMConfigRequest) Validate() error {
	return dara.Validate(s)
}
