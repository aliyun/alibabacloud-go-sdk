// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *GetStandardGroupRequest
	GetGroupId() *int64
	SetTid(v int64) *GetStandardGroupRequest
	GetTid() *int64
}

type GetStandardGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 242***
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 23****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStandardGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardGroupRequest) GoString() string {
	return s.String()
}

func (s *GetStandardGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetStandardGroupRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetStandardGroupRequest) SetGroupId(v int64) *GetStandardGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetStandardGroupRequest) SetTid(v int64) *GetStandardGroupRequest {
	s.Tid = &v
	return s
}

func (s *GetStandardGroupRequest) Validate() error {
	return dara.Validate(s)
}
