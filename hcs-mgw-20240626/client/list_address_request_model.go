// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListAddressRequest
	GetCount() *int32
	SetMarker(v string) *ListAddressRequest
	GetMarker() *string
}

type ListAddressRequest struct {
	// Specifies the number of migration addresses to be returned.\\
	//
	// Valid values: 0 - 1000 (excluding 0).\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 100
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which the migration addresses are listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test_marker
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddressRequest) GoString() string {
	return s.String()
}

func (s *ListAddressRequest) GetCount() *int32 {
	return s.Count
}

func (s *ListAddressRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListAddressRequest) SetCount(v int32) *ListAddressRequest {
	s.Count = &v
	return s
}

func (s *ListAddressRequest) SetMarker(v string) *ListAddressRequest {
	s.Marker = &v
	return s
}

func (s *ListAddressRequest) Validate() error {
	return dara.Validate(s)
}
