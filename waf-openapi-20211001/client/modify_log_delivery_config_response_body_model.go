// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogDeliveryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryName(v string) *ModifyLogDeliveryConfigResponseBody
	GetDeliveryName() *string
	SetRequestId(v string) *ModifyLogDeliveryConfigResponseBody
	GetRequestId() *string
}

type ModifyLogDeliveryConfigResponseBody struct {
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
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogDeliveryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogDeliveryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogDeliveryConfigResponseBody) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *ModifyLogDeliveryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLogDeliveryConfigResponseBody) SetDeliveryName(v string) *ModifyLogDeliveryConfigResponseBody {
	s.DeliveryName = &v
	return s
}

func (s *ModifyLogDeliveryConfigResponseBody) SetRequestId(v string) *ModifyLogDeliveryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLogDeliveryConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
