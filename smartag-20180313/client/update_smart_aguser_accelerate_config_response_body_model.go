// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGUserAccelerateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSmartAGUserAccelerateConfigResponseBody
	GetRequestId() *string
}

type UpdateSmartAGUserAccelerateConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 465B5F94-C7CF-4D54-803D-A7BEAC8545D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmartAGUserAccelerateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGUserAccelerateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGUserAccelerateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAGUserAccelerateConfigResponseBody) SetRequestId(v string) *UpdateSmartAGUserAccelerateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
