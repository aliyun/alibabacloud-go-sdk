// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRetcodeShareStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *SetRetcodeShareStatusRequest
	GetPid() *string
	SetRegionId(v string) *SetRetcodeShareStatusRequest
	GetRegionId() *string
	SetStatus(v bool) *SetRetcodeShareStatusRequest
	GetStatus() *bool
}

type SetRetcodeShareStatusRequest struct {
	// This parameter is required.
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetRetcodeShareStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRetcodeShareStatusRequest) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusRequest) GetPid() *string {
	return s.Pid
}

func (s *SetRetcodeShareStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetRetcodeShareStatusRequest) GetStatus() *bool {
	return s.Status
}

func (s *SetRetcodeShareStatusRequest) SetPid(v string) *SetRetcodeShareStatusRequest {
	s.Pid = &v
	return s
}

func (s *SetRetcodeShareStatusRequest) SetRegionId(v string) *SetRetcodeShareStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetRetcodeShareStatusRequest) SetStatus(v bool) *SetRetcodeShareStatusRequest {
	s.Status = &v
	return s
}

func (s *SetRetcodeShareStatusRequest) Validate() error {
	return dara.Validate(s)
}
