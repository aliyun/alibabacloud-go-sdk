// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFundAccountCreditAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v interface{}) *SetFundAccountCreditAmountResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *SetFundAccountCreditAmountResponseBody
	GetRequestId() *string
}

type SetFundAccountCreditAmountResponseBody struct {
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFundAccountCreditAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetFundAccountCreditAmountResponseBody) GoString() string {
	return s.String()
}

func (s *SetFundAccountCreditAmountResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *SetFundAccountCreditAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetFundAccountCreditAmountResponseBody) SetMetadata(v interface{}) *SetFundAccountCreditAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *SetFundAccountCreditAmountResponseBody) SetRequestId(v string) *SetFundAccountCreditAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetFundAccountCreditAmountResponseBody) Validate() error {
	return dara.Validate(s)
}
