// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVerifyLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMerchantBizId(v string) *CheckVerifyLogRequest
	GetMerchantBizId() *string
	SetTransactionId(v string) *CheckVerifyLogRequest
	GetTransactionId() *string
}

type CheckVerifyLogRequest struct {
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// hksb7ba1b*********015d694361bee4
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CheckVerifyLogRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckVerifyLogRequest) GoString() string {
	return s.String()
}

func (s *CheckVerifyLogRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *CheckVerifyLogRequest) GetTransactionId() *string {
	return s.TransactionId
}

func (s *CheckVerifyLogRequest) SetMerchantBizId(v string) *CheckVerifyLogRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CheckVerifyLogRequest) SetTransactionId(v string) *CheckVerifyLogRequest {
	s.TransactionId = &v
	return s
}

func (s *CheckVerifyLogRequest) Validate() error {
	return dara.Validate(s)
}
