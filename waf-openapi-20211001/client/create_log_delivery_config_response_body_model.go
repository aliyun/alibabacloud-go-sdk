// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogDeliveryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryName(v string) *CreateLogDeliveryConfigResponseBody
	GetDeliveryName() *string
	SetRequestId(v string) *CreateLogDeliveryConfigResponseBody
	GetRequestId() *string
}

type CreateLogDeliveryConfigResponseBody struct {
	// The name of the log delivery configuration.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 26E46541-7AAB-5565-801D-F14DBDC5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLogDeliveryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLogDeliveryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogDeliveryConfigResponseBody) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *CreateLogDeliveryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLogDeliveryConfigResponseBody) SetDeliveryName(v string) *CreateLogDeliveryConfigResponseBody {
	s.DeliveryName = &v
	return s
}

func (s *CreateLogDeliveryConfigResponseBody) SetRequestId(v string) *CreateLogDeliveryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogDeliveryConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
