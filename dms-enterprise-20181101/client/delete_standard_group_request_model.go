// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DeleteStandardGroupRequest
	GetGroupId() *int64
	SetTid(v int64) *DeleteStandardGroupRequest
	GetTid() *int64
}

type DeleteStandardGroupRequest struct {
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

func (s DeleteStandardGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteStandardGroupRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteStandardGroupRequest) SetGroupId(v int64) *DeleteStandardGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteStandardGroupRequest) SetTid(v int64) *DeleteStandardGroupRequest {
	s.Tid = &v
	return s
}

func (s *DeleteStandardGroupRequest) Validate() error {
	return dara.Validate(s)
}
