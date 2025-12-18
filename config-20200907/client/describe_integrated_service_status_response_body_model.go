// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntegratedServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorDeliveryDataType(v string) *DescribeIntegratedServiceStatusResponseBody
	GetAggregatorDeliveryDataType() *string
	SetData(v bool) *DescribeIntegratedServiceStatusResponseBody
	GetData() *bool
	SetRequestId(v string) *DescribeIntegratedServiceStatusResponseBody
	GetRequestId() *string
}

type DescribeIntegratedServiceStatusResponseBody struct {
	// example:
	//
	// NonCompliantNotification
	AggregatorDeliveryDataType *string `json:"AggregatorDeliveryDataType,omitempty" xml:"AggregatorDeliveryDataType,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 2E265A38-84D9-5083-A333-B33A2B46D139
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIntegratedServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntegratedServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntegratedServiceStatusResponseBody) GetAggregatorDeliveryDataType() *string {
	return s.AggregatorDeliveryDataType
}

func (s *DescribeIntegratedServiceStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeIntegratedServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIntegratedServiceStatusResponseBody) SetAggregatorDeliveryDataType(v string) *DescribeIntegratedServiceStatusResponseBody {
	s.AggregatorDeliveryDataType = &v
	return s
}

func (s *DescribeIntegratedServiceStatusResponseBody) SetData(v bool) *DescribeIntegratedServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeIntegratedServiceStatusResponseBody) SetRequestId(v string) *DescribeIntegratedServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntegratedServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
