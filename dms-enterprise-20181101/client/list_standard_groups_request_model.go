// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStandardGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *ListStandardGroupsRequest
	GetTid() *int64
}

type ListStandardGroupsRequest struct {
	// The ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListStandardGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStandardGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListStandardGroupsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListStandardGroupsRequest) SetTid(v int64) *ListStandardGroupsRequest {
	s.Tid = &v
	return s
}

func (s *ListStandardGroupsRequest) Validate() error {
	return dara.Validate(s)
}
