// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecursionZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteRecursionZoneRequest
	GetClientToken() *string
	SetZoneId(v string) *DeleteRecursionZoneRequest
	GetZoneId() *string
}

type DeleteRecursionZoneRequest struct {
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 169783221000012
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteRecursionZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecursionZoneRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecursionZoneRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRecursionZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteRecursionZoneRequest) SetClientToken(v string) *DeleteRecursionZoneRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRecursionZoneRequest) SetZoneId(v string) *DeleteRecursionZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteRecursionZoneRequest) Validate() error {
	return dara.Validate(s)
}
