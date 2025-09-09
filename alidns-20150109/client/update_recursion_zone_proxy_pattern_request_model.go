// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneProxyPatternRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionZoneProxyPatternRequest
	GetClientToken() *string
	SetProxyPattern(v string) *UpdateRecursionZoneProxyPatternRequest
	GetProxyPattern() *string
	SetZoneId(v string) *UpdateRecursionZoneProxyPatternRequest
	GetZoneId() *string
}

type UpdateRecursionZoneProxyPatternRequest struct {
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// record
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// example:
	//
	// 173671468000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateRecursionZoneProxyPatternRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneProxyPatternRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneProxyPatternRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionZoneProxyPatternRequest) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *UpdateRecursionZoneProxyPatternRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateRecursionZoneProxyPatternRequest) SetClientToken(v string) *UpdateRecursionZoneProxyPatternRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionZoneProxyPatternRequest) SetProxyPattern(v string) *UpdateRecursionZoneProxyPatternRequest {
	s.ProxyPattern = &v
	return s
}

func (s *UpdateRecursionZoneProxyPatternRequest) SetZoneId(v string) *UpdateRecursionZoneProxyPatternRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateRecursionZoneProxyPatternRequest) Validate() error {
	return dara.Validate(s)
}
