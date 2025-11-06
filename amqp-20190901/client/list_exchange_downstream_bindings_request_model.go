// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeDownstreamBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListExchangeDownstreamBindingsRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *ListExchangeDownstreamBindingsRequest
	GetCurrentPage() *int32
	SetExchangeName(v string) *ListExchangeDownstreamBindingsRequest
	GetExchangeName() *string
	SetInstanceId(v string) *ListExchangeDownstreamBindingsRequest
	GetInstanceId() *string
	SetPageSize(v int32) *ListExchangeDownstreamBindingsRequest
	GetPageSize() *int32
	SetVhostName(v string) *ListExchangeDownstreamBindingsRequest
	GetVhostName() *string
}

type ListExchangeDownstreamBindingsRequest struct {
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

func (s ListExchangeDownstreamBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeDownstreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListExchangeDownstreamBindingsRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListExchangeDownstreamBindingsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListExchangeDownstreamBindingsRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *ListExchangeDownstreamBindingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExchangeDownstreamBindingsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExchangeDownstreamBindingsRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *ListExchangeDownstreamBindingsRequest) SetConsoleSessionId(v string) *ListExchangeDownstreamBindingsRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListExchangeDownstreamBindingsRequest) SetCurrentPage(v int32) *ListExchangeDownstreamBindingsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListExchangeDownstreamBindingsRequest) SetExchangeName(v string) *ListExchangeDownstreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListExchangeDownstreamBindingsRequest) SetInstanceId(v string) *ListExchangeDownstreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangeDownstreamBindingsRequest) SetPageSize(v int32) *ListExchangeDownstreamBindingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExchangeDownstreamBindingsRequest) SetVhostName(v string) *ListExchangeDownstreamBindingsRequest {
	s.VhostName = &v
	return s
}

func (s *ListExchangeDownstreamBindingsRequest) Validate() error {
	return dara.Validate(s)
}
