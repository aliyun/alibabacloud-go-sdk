// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreditPackageId(v string) *DescribeCreditPackageRequest
	GetCreditPackageId() *string
	SetCreditPackageStatus(v string) *DescribeCreditPackageRequest
	GetCreditPackageStatus() *string
}

type DescribeCreditPackageRequest struct {
	// The ID of the credit package.
	//
	// example:
	//
	// crp-xagydbhfkah****
	CreditPackageId *string `json:"CreditPackageId,omitempty" xml:"CreditPackageId,omitempty"`
	// The status of the credit package.
	//
	// example:
	//
	// ACTIVE
	CreditPackageStatus *string `json:"CreditPackageStatus,omitempty" xml:"CreditPackageStatus,omitempty"`
}

func (s DescribeCreditPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackageRequest) GetCreditPackageId() *string {
	return s.CreditPackageId
}

func (s *DescribeCreditPackageRequest) GetCreditPackageStatus() *string {
	return s.CreditPackageStatus
}

func (s *DescribeCreditPackageRequest) SetCreditPackageId(v string) *DescribeCreditPackageRequest {
	s.CreditPackageId = &v
	return s
}

func (s *DescribeCreditPackageRequest) SetCreditPackageStatus(v string) *DescribeCreditPackageRequest {
	s.CreditPackageStatus = &v
	return s
}

func (s *DescribeCreditPackageRequest) Validate() error {
	return dara.Validate(s)
}
