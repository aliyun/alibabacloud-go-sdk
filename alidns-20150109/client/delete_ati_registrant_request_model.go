// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAtiRegistrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteAtiRegistrantRequest
	GetClientToken() *string
	SetRegistrantId(v string) *DeleteAtiRegistrantRequest
	GetRegistrantId() *string
}

type DeleteAtiRegistrantRequest struct {
	// The client token that is used to ensure the idempotency of the request.
	//
	// Generate a parameter value from your client to ensure that the value is unique among different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the verified registrant.
	//
	// example:
	//
	// 2072277378616354816
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
}

func (s DeleteAtiRegistrantRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAtiRegistrantRequest) GoString() string {
	return s.String()
}

func (s *DeleteAtiRegistrantRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAtiRegistrantRequest) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *DeleteAtiRegistrantRequest) SetClientToken(v string) *DeleteAtiRegistrantRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAtiRegistrantRequest) SetRegistrantId(v string) *DeleteAtiRegistrantRequest {
	s.RegistrantId = &v
	return s
}

func (s *DeleteAtiRegistrantRequest) Validate() error {
	return dara.Validate(s)
}
