// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListPublishedAgentRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListPublishedAgentRequest
	GetPageSize() *int32
}

type ListPublishedAgentRequest struct {
	PageNo   *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPublishedAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListPublishedAgentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedAgentRequest) SetPageNo(v int32) *ListPublishedAgentRequest {
	s.PageNo = &v
	return s
}

func (s *ListPublishedAgentRequest) SetPageSize(v int32) *ListPublishedAgentRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishedAgentRequest) Validate() error {
	return dara.Validate(s)
}
