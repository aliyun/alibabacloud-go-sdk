// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListListenersRequest
	GetAcceleratorId() *string
	SetPageNumber(v int32) *ListListenersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListListenersRequest
	GetPageSize() *int32
	SetProtocol(v string) *ListListenersRequest
	GetProtocol() *string
	SetRegionId(v string) *ListListenersRequest
	GetRegionId() *string
}

type ListListenersRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListListenersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListListenersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListListenersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListListenersRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ListListenersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListListenersRequest) SetAcceleratorId(v string) *ListListenersRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListListenersRequest) SetPageNumber(v int32) *ListListenersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListListenersRequest) SetPageSize(v int32) *ListListenersRequest {
	s.PageSize = &v
	return s
}

func (s *ListListenersRequest) SetProtocol(v string) *ListListenersRequest {
	s.Protocol = &v
	return s
}

func (s *ListListenersRequest) SetRegionId(v string) *ListListenersRequest {
	s.RegionId = &v
	return s
}

func (s *ListListenersRequest) Validate() error {
	return dara.Validate(s)
}
