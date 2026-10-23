// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *DeactivateZonesRequest
	GetBody() *string
	SetClientToken(v string) *DeactivateZonesRequest
	GetClientToken() *string
}

type DeactivateZonesRequest struct {
	// The IDs of the zones to be taken offline.
	//
	// example:
	//
	// ["cn-hangzhou-i","cn-hangzhou-f"]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeactivateZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactivateZonesRequest) GoString() string {
	return s.String()
}

func (s *DeactivateZonesRequest) GetBody() *string {
	return s.Body
}

func (s *DeactivateZonesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeactivateZonesRequest) SetBody(v string) *DeactivateZonesRequest {
	s.Body = &v
	return s
}

func (s *DeactivateZonesRequest) SetClientToken(v string) *DeactivateZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeactivateZonesRequest) Validate() error {
	return dara.Validate(s)
}
