// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProxiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListModelProxiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelProxiesRequest
	GetPageSize() *int32
	SetProxyMode(v string) *ListModelProxiesRequest
	GetProxyMode() *string
	SetStatus(v string) *ListModelProxiesRequest
	GetStatus() *string
}

type ListModelProxiesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// proxyMode
	//
	// example:
	//
	// proxyMode
	ProxyMode *string `json:"proxyMode,omitempty" xml:"proxyMode,omitempty"`
	// example:
	//
	// CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListModelProxiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelProxiesRequest) GoString() string {
	return s.String()
}

func (s *ListModelProxiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelProxiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelProxiesRequest) GetProxyMode() *string {
	return s.ProxyMode
}

func (s *ListModelProxiesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListModelProxiesRequest) SetPageNumber(v int32) *ListModelProxiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelProxiesRequest) SetPageSize(v int32) *ListModelProxiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelProxiesRequest) SetProxyMode(v string) *ListModelProxiesRequest {
	s.ProxyMode = &v
	return s
}

func (s *ListModelProxiesRequest) SetStatus(v string) *ListModelProxiesRequest {
	s.Status = &v
	return s
}

func (s *ListModelProxiesRequest) Validate() error {
	return dara.Validate(s)
}
