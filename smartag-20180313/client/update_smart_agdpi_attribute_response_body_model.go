// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGDpiAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSmartAGDpiAttributeResponseBody
	GetRequestId() *string
}

type UpdateSmartAGDpiAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmartAGDpiAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGDpiAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGDpiAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAGDpiAttributeResponseBody) SetRequestId(v string) *UpdateSmartAGDpiAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
