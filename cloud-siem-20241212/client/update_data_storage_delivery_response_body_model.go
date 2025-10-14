// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataStorageDeliveryResponseBody
	GetRequestId() *string
}

type UpdateDataStorageDeliveryResponseBody struct {
	// example:
	//
	// 6D7FBF4A-5B95-5760-8B5A-BF8983D4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataStorageDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataStorageDeliveryResponseBody) SetRequestId(v string) *UpdateDataStorageDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataStorageDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
