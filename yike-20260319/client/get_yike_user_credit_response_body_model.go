// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeUserCreditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreditTotal(v string) *GetYikeUserCreditResponseBody
	GetCreditTotal() *string
	SetCreditUsage(v string) *GetYikeUserCreditResponseBody
	GetCreditUsage() *string
	SetRequestId(v string) *GetYikeUserCreditResponseBody
	GetRequestId() *string
}

type GetYikeUserCreditResponseBody struct {
	// example:
	//
	// 10417.0
	CreditTotal *string `json:"CreditTotal,omitempty" xml:"CreditTotal,omitempty"`
	// example:
	//
	// 165.0
	CreditUsage *string `json:"CreditUsage,omitempty" xml:"CreditUsage,omitempty"`
	// RequestId
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetYikeUserCreditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeUserCreditResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeUserCreditResponseBody) GetCreditTotal() *string {
	return s.CreditTotal
}

func (s *GetYikeUserCreditResponseBody) GetCreditUsage() *string {
	return s.CreditUsage
}

func (s *GetYikeUserCreditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeUserCreditResponseBody) SetCreditTotal(v string) *GetYikeUserCreditResponseBody {
	s.CreditTotal = &v
	return s
}

func (s *GetYikeUserCreditResponseBody) SetCreditUsage(v string) *GetYikeUserCreditResponseBody {
	s.CreditUsage = &v
	return s
}

func (s *GetYikeUserCreditResponseBody) SetRequestId(v string) *GetYikeUserCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeUserCreditResponseBody) Validate() error {
	return dara.Validate(s)
}
