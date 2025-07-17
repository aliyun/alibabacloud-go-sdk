// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRetcodeAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteRetcodeAppRequest
	GetAppId() *string
	SetAppName(v string) *DeleteRetcodeAppRequest
	GetAppName() *string
	SetPid(v string) *DeleteRetcodeAppRequest
	GetPid() *string
	SetRegionId(v string) *DeleteRetcodeAppRequest
	GetRegionId() *string
}

type DeleteRetcodeAppRequest struct {
	// The application ID.
	//
	// example:
	//
	// 1231
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The process identifier (PID) of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// aokcdqn3ly@741623b4e91****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRetcodeAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteRetcodeAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteRetcodeAppRequest) GetPid() *string {
	return s.Pid
}

func (s *DeleteRetcodeAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRetcodeAppRequest) SetAppId(v string) *DeleteRetcodeAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRetcodeAppRequest) SetAppName(v string) *DeleteRetcodeAppRequest {
	s.AppName = &v
	return s
}

func (s *DeleteRetcodeAppRequest) SetPid(v string) *DeleteRetcodeAppRequest {
	s.Pid = &v
	return s
}

func (s *DeleteRetcodeAppRequest) SetRegionId(v string) *DeleteRetcodeAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRetcodeAppRequest) Validate() error {
	return dara.Validate(s)
}
