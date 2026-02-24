// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegratedServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorDeliveryDataType(v string) *GetIntegratedServiceStatusResponseBody
	GetAggregatorDeliveryDataType() *string
	SetData(v bool) *GetIntegratedServiceStatusResponseBody
	GetData() *bool
	SetIntegratedTypes(v string) *GetIntegratedServiceStatusResponseBody
	GetIntegratedTypes() *string
	SetRequestId(v string) *GetIntegratedServiceStatusResponseBody
	GetRequestId() *string
}

type GetIntegratedServiceStatusResponseBody struct {
	AggregatorDeliveryDataType *string `json:"AggregatorDeliveryDataType,omitempty" xml:"AggregatorDeliveryDataType,omitempty"`
	Data                       *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	IntegratedTypes            *string `json:"IntegratedTypes,omitempty" xml:"IntegratedTypes,omitempty"`
	RequestId                  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIntegratedServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntegratedServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegratedServiceStatusResponseBody) GetAggregatorDeliveryDataType() *string {
	return s.AggregatorDeliveryDataType
}

func (s *GetIntegratedServiceStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *GetIntegratedServiceStatusResponseBody) GetIntegratedTypes() *string {
	return s.IntegratedTypes
}

func (s *GetIntegratedServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntegratedServiceStatusResponseBody) SetAggregatorDeliveryDataType(v string) *GetIntegratedServiceStatusResponseBody {
	s.AggregatorDeliveryDataType = &v
	return s
}

func (s *GetIntegratedServiceStatusResponseBody) SetData(v bool) *GetIntegratedServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetIntegratedServiceStatusResponseBody) SetIntegratedTypes(v string) *GetIntegratedServiceStatusResponseBody {
	s.IntegratedTypes = &v
	return s
}

func (s *GetIntegratedServiceStatusResponseBody) SetRequestId(v string) *GetIntegratedServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntegratedServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
