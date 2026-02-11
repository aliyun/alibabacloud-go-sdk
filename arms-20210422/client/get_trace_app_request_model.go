// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *GetTraceAppRequest
	GetPid() *string
	SetRegionId(v string) *GetTraceAppRequest
	GetRegionId() *string
}

type GetTraceAppRequest struct {
	// This parameter is required.
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTraceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppRequest) GoString() string {
	return s.String()
}

func (s *GetTraceAppRequest) GetPid() *string {
	return s.Pid
}

func (s *GetTraceAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTraceAppRequest) SetPid(v string) *GetTraceAppRequest {
	s.Pid = &v
	return s
}

func (s *GetTraceAppRequest) SetRegionId(v string) *GetTraceAppRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceAppRequest) Validate() error {
	return dara.Validate(s)
}
