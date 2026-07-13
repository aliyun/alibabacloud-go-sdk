// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiRegistrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeAtiRegistrantRequest
	GetClientToken() *string
	SetRegistrantId(v string) *DescribeAtiRegistrantRequest
	GetRegistrantId() *string
}

type DescribeAtiRegistrantRequest struct {
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the real-name verified registrant.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2072277378616354816
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
}

func (s DescribeAtiRegistrantRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiRegistrantRequest) GoString() string {
	return s.String()
}

func (s *DescribeAtiRegistrantRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAtiRegistrantRequest) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *DescribeAtiRegistrantRequest) SetClientToken(v string) *DescribeAtiRegistrantRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAtiRegistrantRequest) SetRegistrantId(v string) *DescribeAtiRegistrantRequest {
	s.RegistrantId = &v
	return s
}

func (s *DescribeAtiRegistrantRequest) Validate() error {
	return dara.Validate(s)
}
