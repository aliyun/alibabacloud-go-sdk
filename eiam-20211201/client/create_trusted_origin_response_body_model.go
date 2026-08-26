// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrustedOriginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrustedOriginResponseBody
	GetRequestId() *string
	SetTrustedOriginId(v string) *CreateTrustedOriginResponseBody
	GetTrustedOriginId() *string
}

type CreateTrustedOriginResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trusted origin ID.
	//
	// example:
	//
	// to_example
	TrustedOriginId *string `json:"TrustedOriginId,omitempty" xml:"TrustedOriginId,omitempty"`
}

func (s CreateTrustedOriginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrustedOriginResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrustedOriginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrustedOriginResponseBody) GetTrustedOriginId() *string {
	return s.TrustedOriginId
}

func (s *CreateTrustedOriginResponseBody) SetRequestId(v string) *CreateTrustedOriginResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrustedOriginResponseBody) SetTrustedOriginId(v string) *CreateTrustedOriginResponseBody {
	s.TrustedOriginId = &v
	return s
}

func (s *CreateTrustedOriginResponseBody) Validate() error {
	return dara.Validate(s)
}
