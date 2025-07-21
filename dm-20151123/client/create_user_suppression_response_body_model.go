// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSuppressionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserSuppressionResponseBody
	GetRequestId() *string
	SetSuppressionId(v string) *CreateUserSuppressionResponseBody
	GetSuppressionId() *string
}

type CreateUserSuppressionResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 1A846D66-5EC7-551B-9687-5BF1963DCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Invalid address number
	//
	// example:
	//
	// 59511
	SuppressionId *string `json:"SuppressionId,omitempty" xml:"SuppressionId,omitempty"`
}

func (s CreateUserSuppressionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSuppressionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserSuppressionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserSuppressionResponseBody) GetSuppressionId() *string {
	return s.SuppressionId
}

func (s *CreateUserSuppressionResponseBody) SetRequestId(v string) *CreateUserSuppressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserSuppressionResponseBody) SetSuppressionId(v string) *CreateUserSuppressionResponseBody {
	s.SuppressionId = &v
	return s
}

func (s *CreateUserSuppressionResponseBody) Validate() error {
	return dara.Validate(s)
}
