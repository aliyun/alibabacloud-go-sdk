// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListEdgeContainerAppRecordsRequest
	GetAppId() *string
	SetOrderKey(v string) *ListEdgeContainerAppRecordsRequest
	GetOrderKey() *string
	SetOrderType(v string) *ListEdgeContainerAppRecordsRequest
	GetOrderType() *string
	SetPageNumber(v int32) *ListEdgeContainerAppRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeContainerAppRecordsRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListEdgeContainerAppRecordsRequest
	GetSearchKey() *string
}

type ListEdgeContainerAppRecordsRequest struct {
	// This parameter is required.
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OrderKey   *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	OrderType  *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey  *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s ListEdgeContainerAppRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppRecordsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListEdgeContainerAppRecordsRequest) GetOrderKey() *string {
	return s.OrderKey
}

func (s *ListEdgeContainerAppRecordsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListEdgeContainerAppRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeContainerAppRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeContainerAppRecordsRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListEdgeContainerAppRecordsRequest) SetAppId(v string) *ListEdgeContainerAppRecordsRequest {
	s.AppId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetOrderKey(v string) *ListEdgeContainerAppRecordsRequest {
	s.OrderKey = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetOrderType(v string) *ListEdgeContainerAppRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetPageNumber(v int32) *ListEdgeContainerAppRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetPageSize(v int32) *ListEdgeContainerAppRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetSearchKey(v string) *ListEdgeContainerAppRecordsRequest {
	s.SearchKey = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) Validate() error {
	return dara.Validate(s)
}
