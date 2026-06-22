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
	// The page number of the current page in a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the event has been handled. Valid values:
	//
	// - **y**: Handled.
	//
	// - **n**: Unhandled.
	//
	// - **a**: All.
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return per page in a paged query. Default value: 100. If the PageSize parameter is left empty, 100 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Invalid parameter.
	//
	// example:
	//
	// 9F4E6157-9600-5588-86B9-38F09067****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the alert event.
	//
	// >You can call the [ListHoneypotEvents](~~ListHoneypotEvents~~) operation to obtain this parameter.
	//
	// example:
	//
	// 7455818
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
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
