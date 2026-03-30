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
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.``
	//
	// When you call the operation for the first time, if the total number of returned entries exceeds the value of `MaxItems`, the entries are truncated. The system returns entries based on the value of `MaxItems` and does not return the excess entries. In this case, the value of the response parameter `IsTruncated` is `true`, and `Marker` is returned. In the next call, you can use the value of `Marker` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of the `IsTruncated` parameter becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page.
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
