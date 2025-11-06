// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListExchangeRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *ListExchangeRequest
	GetCurrentPage() *int32
	SetExchangeNamePrefix(v string) *ListExchangeRequest
	GetExchangeNamePrefix() *string
	SetInstanceId(v string) *ListExchangeRequest
	GetInstanceId() *string
	SetPageSize(v int32) *ListExchangeRequest
	GetPageSize() *int32
	SetVhostName(v string) *ListExchangeRequest
	GetVhostName() *string
}

type ListExchangeRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage        *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ExchangeNamePrefix *string `json:"ExchangeNamePrefix,omitempty" xml:"ExchangeNamePrefix,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s ListExchangeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeRequest) GoString() string {
	return s.String()
}

func (s *ListExchangeRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListExchangeRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListExchangeRequest) GetExchangeNamePrefix() *string {
	return s.ExchangeNamePrefix
}

func (s *ListExchangeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExchangeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExchangeRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *ListExchangeRequest) SetConsoleSessionId(v string) *ListExchangeRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListExchangeRequest) SetCurrentPage(v int32) *ListExchangeRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListExchangeRequest) SetExchangeNamePrefix(v string) *ListExchangeRequest {
	s.ExchangeNamePrefix = &v
	return s
}

func (s *ListExchangeRequest) SetInstanceId(v string) *ListExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangeRequest) SetPageSize(v int32) *ListExchangeRequest {
	s.PageSize = &v
	return s
}

func (s *ListExchangeRequest) SetVhostName(v string) *ListExchangeRequest {
	s.VhostName = &v
	return s
}

func (s *ListExchangeRequest) Validate() error {
	return dara.Validate(s)
}
