// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAtiRegistrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RevokeAtiRegistrantRequest
	GetClientToken() *string
	SetRegistrantId(v string) *RevokeAtiRegistrantRequest
	GetRegistrantId() *string
}

type RevokeAtiRegistrantRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// If you do not specify this parameter, the system automatically uses the value of RequestId as the value of ClientToken. The value of RequestId may be different for each API request.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the real-name registrant.
	//
	// example:
	//
	// 2072277378616354816
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
}

func (s RevokeAtiRegistrantRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeAtiRegistrantRequest) GoString() string {
	return s.String()
}

func (s *RevokeAtiRegistrantRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RevokeAtiRegistrantRequest) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *RevokeAtiRegistrantRequest) SetClientToken(v string) *RevokeAtiRegistrantRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeAtiRegistrantRequest) SetRegistrantId(v string) *RevokeAtiRegistrantRequest {
	s.RegistrantId = &v
	return s
}

func (s *RevokeAtiRegistrantRequest) Validate() error {
	return dara.Validate(s)
}
