// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSAMLProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListSAMLProvidersRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListSAMLProvidersRequest
	GetMaxItems() *int32
}

type ListSAMLProvidersRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListSAMLProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSAMLProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListSAMLProvidersRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListSAMLProvidersRequest) SetMarker(v string) *ListSAMLProvidersRequest {
	s.Marker = &v
	return s
}

func (s *ListSAMLProvidersRequest) SetMaxItems(v int32) *ListSAMLProvidersRequest {
	s.MaxItems = &v
	return s
}

func (s *ListSAMLProvidersRequest) Validate() error {
	return dara.Validate(s)
}
