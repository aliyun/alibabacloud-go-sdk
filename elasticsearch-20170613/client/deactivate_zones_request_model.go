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
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Used to ensure idempotency of the request. The client generates this parameter value and must guarantee its uniqueness across different requests, with a maximum length of 64 ASCII characters.
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
