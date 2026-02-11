// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTraceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteTraceAppRequest
	GetAppId() *string
	SetPid(v string) *DeleteTraceAppRequest
	GetPid() *string
	SetRegionId(v string) *DeleteTraceAppRequest
	GetRegionId() *string
	SetType(v string) *DeleteTraceAppRequest
	GetType() *string
}

type DeleteTraceAppRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteTraceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTraceAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteTraceAppRequest) GetPid() *string {
	return s.Pid
}

func (s *DeleteTraceAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTraceAppRequest) GetType() *string {
	return s.Type
}

func (s *DeleteTraceAppRequest) SetAppId(v string) *DeleteTraceAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetPid(v string) *DeleteTraceAppRequest {
	s.Pid = &v
	return s
}

func (s *DeleteTraceAppRequest) SetRegionId(v string) *DeleteTraceAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetType(v string) *DeleteTraceAppRequest {
	s.Type = &v
	return s
}

func (s *DeleteTraceAppRequest) Validate() error {
	return dara.Validate(s)
}
