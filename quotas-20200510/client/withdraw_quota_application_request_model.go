// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawQuotaApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *WithdrawQuotaApplicationRequest
	GetApplicationId() *string
	SetReason(v string) *WithdrawQuotaApplicationRequest
	GetReason() *string
}

type WithdrawQuotaApplicationRequest struct {
	// The ID of the quota application.
	//
	// This parameter is required.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d42****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The reason for withdrawing the quota increase request.
	//
	// example:
	//
	// DesireValue is wrong
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s WithdrawQuotaApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s WithdrawQuotaApplicationRequest) GoString() string {
	return s.String()
}

func (s *WithdrawQuotaApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *WithdrawQuotaApplicationRequest) GetReason() *string {
	return s.Reason
}

func (s *WithdrawQuotaApplicationRequest) SetApplicationId(v string) *WithdrawQuotaApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *WithdrawQuotaApplicationRequest) SetReason(v string) *WithdrawQuotaApplicationRequest {
	s.Reason = &v
	return s
}

func (s *WithdrawQuotaApplicationRequest) Validate() error {
	return dara.Validate(s)
}
