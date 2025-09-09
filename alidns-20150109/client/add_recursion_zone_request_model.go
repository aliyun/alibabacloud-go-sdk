// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecursionZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddRecursionZoneRequest
	GetClientToken() *string
	SetProxyPattern(v string) *AddRecursionZoneRequest
	GetProxyPattern() *string
	SetZoneName(v string) *AddRecursionZoneRequest
	GetZoneName() *string
}

type AddRecursionZoneRequest struct {
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// record
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s AddRecursionZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecursionZoneRequest) GoString() string {
	return s.String()
}

func (s *AddRecursionZoneRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddRecursionZoneRequest) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *AddRecursionZoneRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *AddRecursionZoneRequest) SetClientToken(v string) *AddRecursionZoneRequest {
	s.ClientToken = &v
	return s
}

func (s *AddRecursionZoneRequest) SetProxyPattern(v string) *AddRecursionZoneRequest {
	s.ProxyPattern = &v
	return s
}

func (s *AddRecursionZoneRequest) SetZoneName(v string) *AddRecursionZoneRequest {
	s.ZoneName = &v
	return s
}

func (s *AddRecursionZoneRequest) Validate() error {
	return dara.Validate(s)
}
