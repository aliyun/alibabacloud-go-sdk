// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalUserId(v string) *GetUsageRequest
	GetExternalUserId() *string
	SetRedemptionOrderNo(v string) *GetUsageRequest
	GetRedemptionOrderNo() *string
}

type GetUsageRequest struct {
	// example:
	//
	// 1001
	ExternalUserId *string `json:"externalUserId,omitempty" xml:"externalUserId,omitempty"`
	// example:
	//
	// R123456789
	RedemptionOrderNo *string `json:"redemptionOrderNo,omitempty" xml:"redemptionOrderNo,omitempty"`
}

func (s GetUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUsageRequest) GoString() string {
	return s.String()
}

func (s *GetUsageRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *GetUsageRequest) GetRedemptionOrderNo() *string {
	return s.RedemptionOrderNo
}

func (s *GetUsageRequest) SetExternalUserId(v string) *GetUsageRequest {
	s.ExternalUserId = &v
	return s
}

func (s *GetUsageRequest) SetRedemptionOrderNo(v string) *GetUsageRequest {
	s.RedemptionOrderNo = &v
	return s
}

func (s *GetUsageRequest) Validate() error {
	return dara.Validate(s)
}
