// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *ActivateZonesRequest
	GetBody() *string
	SetClientToken(v string) *ActivateZonesRequest
	GetClientToken() *string
}

type ActivateZonesRequest struct {
	// The request body.
	//
	// example:
	//
	// Empty
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies the client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ActivateZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateZonesRequest) GoString() string {
	return s.String()
}

func (s *ActivateZonesRequest) GetBody() *string {
	return s.Body
}

func (s *ActivateZonesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ActivateZonesRequest) SetBody(v string) *ActivateZonesRequest {
	s.Body = &v
	return s
}

func (s *ActivateZonesRequest) SetClientToken(v string) *ActivateZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *ActivateZonesRequest) Validate() error {
	return dara.Validate(s)
}
