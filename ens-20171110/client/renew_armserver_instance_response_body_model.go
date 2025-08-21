// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewARMServerInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenewARMServerInstanceResponseBody
	GetRequestId() *string
}

type RenewARMServerInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 370E61E0-6E6E-50FE-9259-EE706C55ABF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewARMServerInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewARMServerInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewARMServerInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewARMServerInstanceResponseBody) SetRequestId(v string) *RenewARMServerInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewARMServerInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
