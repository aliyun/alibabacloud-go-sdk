// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeUpstreamBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListExchangeUpstreamBindingsRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *ListExchangeUpstreamBindingsRequest
	GetCurrentPage() *int32
	SetExchangeName(v string) *ListExchangeUpstreamBindingsRequest
	GetExchangeName() *string
	SetInstanceId(v string) *ListExchangeUpstreamBindingsRequest
	GetInstanceId() *string
	SetPageSize(v int32) *ListExchangeUpstreamBindingsRequest
	GetPageSize() *int32
	SetVhostName(v string) *ListExchangeUpstreamBindingsRequest
	GetVhostName() *string
}

type ListExchangeUpstreamBindingsRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s ListExchangeUpstreamBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpstreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListExchangeUpstreamBindingsRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListExchangeUpstreamBindingsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListExchangeUpstreamBindingsRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *ListExchangeUpstreamBindingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExchangeUpstreamBindingsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExchangeUpstreamBindingsRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *ListExchangeUpstreamBindingsRequest) SetConsoleSessionId(v string) *ListExchangeUpstreamBindingsRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListExchangeUpstreamBindingsRequest) SetCurrentPage(v int32) *ListExchangeUpstreamBindingsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListExchangeUpstreamBindingsRequest) SetExchangeName(v string) *ListExchangeUpstreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListExchangeUpstreamBindingsRequest) SetInstanceId(v string) *ListExchangeUpstreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangeUpstreamBindingsRequest) SetPageSize(v int32) *ListExchangeUpstreamBindingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExchangeUpstreamBindingsRequest) SetVhostName(v string) *ListExchangeUpstreamBindingsRequest {
	s.VhostName = &v
	return s
}

func (s *ListExchangeUpstreamBindingsRequest) Validate() error {
	return dara.Validate(s)
}
