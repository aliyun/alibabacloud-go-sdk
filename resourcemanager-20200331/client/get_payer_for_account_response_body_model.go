// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPayerForAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPayerAccountId(v string) *GetPayerForAccountResponseBody
	GetPayerAccountId() *string
	SetPayerAccountName(v string) *GetPayerForAccountResponseBody
	GetPayerAccountName() *string
	SetRequestId(v string) *GetPayerForAccountResponseBody
	GetRequestId() *string
}

type GetPayerForAccountResponseBody struct {
	// The ID of the settlement account.
	//
	// example:
	//
	// 172841235500****
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	// The name of the settlement account.
	//
	// example:
	//
	// Alice
	PayerAccountName *string `json:"PayerAccountName,omitempty" xml:"PayerAccountName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPayerForAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPayerForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountResponseBody) GetPayerAccountId() *string {
	return s.PayerAccountId
}

func (s *GetPayerForAccountResponseBody) GetPayerAccountName() *string {
	return s.PayerAccountName
}

func (s *GetPayerForAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPayerForAccountResponseBody) SetPayerAccountId(v string) *GetPayerForAccountResponseBody {
	s.PayerAccountId = &v
	return s
}

func (s *GetPayerForAccountResponseBody) SetPayerAccountName(v string) *GetPayerForAccountResponseBody {
	s.PayerAccountName = &v
	return s
}

func (s *GetPayerForAccountResponseBody) SetRequestId(v string) *GetPayerForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPayerForAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
