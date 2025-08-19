// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckForConsolidatedBillingAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingAccountId(v string) *PrecheckForConsolidatedBillingAccountRequest
	GetBillingAccountId() *string
}

type PrecheckForConsolidatedBillingAccountRequest struct {
	// The ID of the management account or member to be used as a main financial account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111***089
	BillingAccountId *string `json:"BillingAccountId,omitempty" xml:"BillingAccountId,omitempty"`
}

func (s PrecheckForConsolidatedBillingAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s PrecheckForConsolidatedBillingAccountRequest) GoString() string {
	return s.String()
}

func (s *PrecheckForConsolidatedBillingAccountRequest) GetBillingAccountId() *string {
	return s.BillingAccountId
}

func (s *PrecheckForConsolidatedBillingAccountRequest) SetBillingAccountId(v string) *PrecheckForConsolidatedBillingAccountRequest {
	s.BillingAccountId = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountRequest) Validate() error {
	return dara.Validate(s)
}
