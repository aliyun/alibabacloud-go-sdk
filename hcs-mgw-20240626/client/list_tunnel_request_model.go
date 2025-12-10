// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTunnelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListTunnelRequest
	GetCount() *int32
	SetMarker(v string) *ListTunnelRequest
	GetMarker() *string
}

type ListTunnelRequest struct {
	// Specifies the number of tunnels to be returned.\\
	//
	// Valid values: 0 - 1000.\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which tunnels are listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 1
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListTunnelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTunnelRequest) GoString() string {
	return s.String()
}

func (s *ListTunnelRequest) GetCount() *int32 {
	return s.Count
}

func (s *ListTunnelRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListTunnelRequest) SetCount(v int32) *ListTunnelRequest {
	s.Count = &v
	return s
}

func (s *ListTunnelRequest) SetMarker(v string) *ListTunnelRequest {
	s.Marker = &v
	return s
}

func (s *ListTunnelRequest) Validate() error {
	return dara.Validate(s)
}
