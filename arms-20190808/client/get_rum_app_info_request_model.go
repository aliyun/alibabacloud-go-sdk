// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroup(v string) *GetRumAppInfoRequest
	GetAppGroup() *string
	SetPid(v string) *GetRumAppInfoRequest
	GetPid() *string
	SetRegionId(v string) *GetRumAppInfoRequest
	GetRegionId() *string
}

type GetRumAppInfoRequest struct {
	// The group to which the application belongs.
	//
	// example:
	//
	// default
	AppGroup *string `json:"AppGroup,omitempty" xml:"AppGroup,omitempty"`
	// The process ID (PID) of the application.
	//
	// This parameter is required.
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
}

func (s GetRumAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoRequest) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoRequest) GetAppGroup() *string {
	return s.AppGroup
}

func (s *GetRumAppInfoRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRumAppInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumAppInfoRequest) SetAppGroup(v string) *GetRumAppInfoRequest {
	s.AppGroup = &v
	return s
}

func (s *GetRumAppInfoRequest) SetPid(v string) *GetRumAppInfoRequest {
	s.Pid = &v
	return s
}

func (s *GetRumAppInfoRequest) SetRegionId(v string) *GetRumAppInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetRumAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
