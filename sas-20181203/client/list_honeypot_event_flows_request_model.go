// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotEventFlowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListHoneypotEventFlowsRequest
	GetCurrentPage() *int32
	SetDealed(v string) *ListHoneypotEventFlowsRequest
	GetDealed() *string
	SetLang(v string) *ListHoneypotEventFlowsRequest
	GetLang() *string
	SetPageSize(v int32) *ListHoneypotEventFlowsRequest
	GetPageSize() *int32
	SetRequestId(v string) *ListHoneypotEventFlowsRequest
	GetRequestId() *string
	SetSecurityEventId(v int64) *ListHoneypotEventFlowsRequest
	GetSecurityEventId() *int64
}

type ListHoneypotEventFlowsRequest struct {
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Dealed          *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
}

func (s ListHoneypotEventFlowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventFlowsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotEventFlowsRequest) GetDealed() *string {
	return s.Dealed
}

func (s *ListHoneypotEventFlowsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListHoneypotEventFlowsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotEventFlowsRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotEventFlowsRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *ListHoneypotEventFlowsRequest) SetCurrentPage(v int32) *ListHoneypotEventFlowsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotEventFlowsRequest) SetDealed(v string) *ListHoneypotEventFlowsRequest {
	s.Dealed = &v
	return s
}

func (s *ListHoneypotEventFlowsRequest) SetLang(v string) *ListHoneypotEventFlowsRequest {
	s.Lang = &v
	return s
}

func (s *ListHoneypotEventFlowsRequest) SetPageSize(v int32) *ListHoneypotEventFlowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotEventFlowsRequest) SetRequestId(v string) *ListHoneypotEventFlowsRequest {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotEventFlowsRequest) SetSecurityEventId(v int64) *ListHoneypotEventFlowsRequest {
	s.SecurityEventId = &v
	return s
}

func (s *ListHoneypotEventFlowsRequest) Validate() error {
	return dara.Validate(s)
}
