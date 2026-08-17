// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaInstanceCrlAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaIdentifier(v string) *GetCaInstanceCrlAddressRequest
	GetCaIdentifier() *string
	SetUuid(v string) *GetCaInstanceCrlAddressRequest
	GetUuid() *string
}

type GetCaInstanceCrlAddressRequest struct {
	// The CA certificate identifier.
	//
	// example:
	//
	// 1f0167b4-ee84-XXX-49bc4d39fa68
	CaIdentifier *string `json:"CaIdentifier,omitempty" xml:"CaIdentifier,omitempty"`
	// The zone ID of the China CAS instance.
	//
	// example:
	//
	// 1f047318-0815-XXX-f7ceb76b5c0a
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetCaInstanceCrlAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCaInstanceCrlAddressRequest) GoString() string {
	return s.String()
}

func (s *GetCaInstanceCrlAddressRequest) GetCaIdentifier() *string {
	return s.CaIdentifier
}

func (s *GetCaInstanceCrlAddressRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetCaInstanceCrlAddressRequest) SetCaIdentifier(v string) *GetCaInstanceCrlAddressRequest {
	s.CaIdentifier = &v
	return s
}

func (s *GetCaInstanceCrlAddressRequest) SetUuid(v string) *GetCaInstanceCrlAddressRequest {
	s.Uuid = &v
	return s
}

func (s *GetCaInstanceCrlAddressRequest) Validate() error {
	return dara.Validate(s)
}
