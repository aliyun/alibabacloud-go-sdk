// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRefundPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingInstanceIds(v string) *QueryRefundPriceRequest
	GetBillingInstanceIds() *string
	SetInstanceId(v string) *QueryRefundPriceRequest
	GetInstanceId() *string
}

type QueryRefundPriceRequest struct {
	// example:
	//
	// ng-1syeu23,ng-81293sa
	BillingInstanceIds *string `json:"billingInstanceIds,omitempty" xml:"billingInstanceIds,omitempty"`
	// example:
	//
	// c-82su12s0kl12
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s QueryRefundPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRefundPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryRefundPriceRequest) GetBillingInstanceIds() *string {
	return s.BillingInstanceIds
}

func (s *QueryRefundPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryRefundPriceRequest) SetBillingInstanceIds(v string) *QueryRefundPriceRequest {
	s.BillingInstanceIds = &v
	return s
}

func (s *QueryRefundPriceRequest) SetInstanceId(v string) *QueryRefundPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRefundPriceRequest) Validate() error {
	return dara.Validate(s)
}
