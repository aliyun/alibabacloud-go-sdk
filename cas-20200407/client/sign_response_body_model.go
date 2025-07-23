// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SignResponseBody
	GetRequestId() *string
	SetSignature(v string) *SignResponseBody
	GetSignature() *string
}

type SignResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1ed33293-2e48-6b14-861e-538e28e408eb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature.
	//
	// example:
	//
	// eyaC0w3ROK5b3QcHmUtAhMY/sQjKu2t3uBfnf6J/gn7JfZtyxwcCUjzXbw5jmqJQRbj1te670Bshg9kUdanKhtHFhJjU5jX+ZMMBr6pH0gqQDJxR0K0yHXRc0Q5OQoUZ6BfpbI4Wt4jJvJSdCstz1vSg12CfEHS8Kd5qfhItK7Y=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s SignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SignResponseBody) GoString() string {
	return s.String()
}

func (s *SignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SignResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *SignResponseBody) SetRequestId(v string) *SignResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignResponseBody) SetSignature(v string) *SignResponseBody {
	s.Signature = &v
	return s
}

func (s *SignResponseBody) Validate() error {
	return dara.Validate(s)
}
