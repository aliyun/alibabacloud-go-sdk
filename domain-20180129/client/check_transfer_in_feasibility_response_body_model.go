// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTransferInFeasibilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanTransfer(v bool) *CheckTransferInFeasibilityResponseBody
	GetCanTransfer() *bool
	SetCode(v string) *CheckTransferInFeasibilityResponseBody
	GetCode() *string
	SetMessage(v string) *CheckTransferInFeasibilityResponseBody
	GetMessage() *string
	SetProductId(v string) *CheckTransferInFeasibilityResponseBody
	GetProductId() *string
	SetRequestId(v string) *CheckTransferInFeasibilityResponseBody
	GetRequestId() *string
}

type CheckTransferInFeasibilityResponseBody struct {
	// example:
	//
	// false
	CanTransfer *bool `json:"CanTransfer,omitempty" xml:"CanTransfer,omitempty"`
	// example:
	//
	// CheckTransferResult.DomainTransferProhibited
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// This domain name is in transfer prohibited status, so it cannot be transferred. You can contact your original registrar to change its status.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2a
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// FC0D6B89-2353-4D64-BD80-6606A7DBD7C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckTransferInFeasibilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckTransferInFeasibilityResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTransferInFeasibilityResponseBody) GetCanTransfer() *bool {
	return s.CanTransfer
}

func (s *CheckTransferInFeasibilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckTransferInFeasibilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckTransferInFeasibilityResponseBody) GetProductId() *string {
	return s.ProductId
}

func (s *CheckTransferInFeasibilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckTransferInFeasibilityResponseBody) SetCanTransfer(v bool) *CheckTransferInFeasibilityResponseBody {
	s.CanTransfer = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) SetCode(v string) *CheckTransferInFeasibilityResponseBody {
	s.Code = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) SetMessage(v string) *CheckTransferInFeasibilityResponseBody {
	s.Message = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) SetProductId(v string) *CheckTransferInFeasibilityResponseBody {
	s.ProductId = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) SetRequestId(v string) *CheckTransferInFeasibilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) Validate() error {
	return dara.Validate(s)
}
