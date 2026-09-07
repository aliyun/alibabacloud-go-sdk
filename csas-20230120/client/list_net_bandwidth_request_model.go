// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListNetBandwidthRequest
	GetCurrentPage() *int32
	SetInstanceIds(v []*string) *ListNetBandwidthRequest
	GetInstanceIds() []*string
	SetNetType(v string) *ListNetBandwidthRequest
	GetNetType() *string
	SetPageSize(v int32) *ListNetBandwidthRequest
	GetPageSize() *int32
}

type ListNetBandwidthRequest struct {
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of instance IDs.
	//
	// if can be null:
	// true
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The network type. If this parameter is left empty, both VPC and Connector instances are queried.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNetBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ListNetBandwidthRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListNetBandwidthRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListNetBandwidthRequest) GetNetType() *string {
	return s.NetType
}

func (s *ListNetBandwidthRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetBandwidthRequest) SetCurrentPage(v int32) *ListNetBandwidthRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListNetBandwidthRequest) SetInstanceIds(v []*string) *ListNetBandwidthRequest {
	s.InstanceIds = v
	return s
}

func (s *ListNetBandwidthRequest) SetNetType(v string) *ListNetBandwidthRequest {
	s.NetType = &v
	return s
}

func (s *ListNetBandwidthRequest) SetPageSize(v int32) *ListNetBandwidthRequest {
	s.PageSize = &v
	return s
}

func (s *ListNetBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
