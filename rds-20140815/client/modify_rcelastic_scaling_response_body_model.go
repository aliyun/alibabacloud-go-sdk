// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCElasticScalingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyRCElasticScalingResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyRCElasticScalingResponseBody
	GetRequestId() *string
}

type ModifyRCElasticScalingResponseBody struct {
	// example:
	//
	// 264008926210869
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1F7B8B09-36F3-1165-BADB-13E376FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCElasticScalingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCElasticScalingResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCElasticScalingResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyRCElasticScalingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCElasticScalingResponseBody) SetOrderId(v string) *ModifyRCElasticScalingResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyRCElasticScalingResponseBody) SetRequestId(v string) *ModifyRCElasticScalingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCElasticScalingResponseBody) Validate() error {
	return dara.Validate(s)
}
