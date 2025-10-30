// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogDeliveryStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryName(v string) *ModifyResourceLogDeliveryStatusResponseBody
	GetDeliveryName() *string
	SetDeliveryType(v string) *ModifyResourceLogDeliveryStatusResponseBody
	GetDeliveryType() *string
	SetRequestId(v string) *ModifyResourceLogDeliveryStatusResponseBody
	GetRequestId() *string
	SetStatus(v bool) *ModifyResourceLogDeliveryStatusResponseBody
	GetStatus() *bool
}

type ModifyResourceLogDeliveryStatusResponseBody struct {
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyResourceLogDeliveryStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogDeliveryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) SetDeliveryName(v string) *ModifyResourceLogDeliveryStatusResponseBody {
	s.DeliveryName = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) SetDeliveryType(v string) *ModifyResourceLogDeliveryStatusResponseBody {
	s.DeliveryType = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) SetRequestId(v string) *ModifyResourceLogDeliveryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) SetStatus(v bool) *ModifyResourceLogDeliveryStatusResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
