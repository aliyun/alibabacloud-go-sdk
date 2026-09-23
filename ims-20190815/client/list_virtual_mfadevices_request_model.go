// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualMFADevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListVirtualMFADevicesRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListVirtualMFADevicesRequest
	GetMaxItems() *int32
}

type ListVirtualMFADevicesRequest struct {
	// The token for querying the next page of results. You do not need to specify `Marker` for the first API call.
	//
	// When you call the API for the first time, if the total number of entries exceeds the `MaxItems` limit, the data is truncated and only `MaxItems` entries are returned. In this case, the `IsTruncated` response parameter is `true` and a `Marker` is returned. You can use the `Marker` returned from the previous call to continue calling the API with the same request parameters to query the truncated data. You can repeat this process until `IsTruncated` is `false`, which indicates that all data has been retrieved.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListVirtualMFADevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualMFADevicesRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListVirtualMFADevicesRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListVirtualMFADevicesRequest) SetMarker(v string) *ListVirtualMFADevicesRequest {
	s.Marker = &v
	return s
}

func (s *ListVirtualMFADevicesRequest) SetMaxItems(v int32) *ListVirtualMFADevicesRequest {
	s.MaxItems = &v
	return s
}

func (s *ListVirtualMFADevicesRequest) Validate() error {
	return dara.Validate(s)
}
