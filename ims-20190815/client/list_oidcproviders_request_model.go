// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOIDCProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListOIDCProvidersRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListOIDCProvidersRequest
	GetMaxItems() *int32
}

type ListOIDCProvidersRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListOIDCProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOIDCProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListOIDCProvidersRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListOIDCProvidersRequest) SetMarker(v string) *ListOIDCProvidersRequest {
	s.Marker = &v
	return s
}

func (s *ListOIDCProvidersRequest) SetMaxItems(v int32) *ListOIDCProvidersRequest {
	s.MaxItems = &v
	return s
}

func (s *ListOIDCProvidersRequest) Validate() error {
	return dara.Validate(s)
}
