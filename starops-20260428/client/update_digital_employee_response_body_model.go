// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDigitalEmployeeResponseBody
	GetRequestId() *string
}

type UpdateDigitalEmployeeResponseBody struct {
	// example:
	//
	// 0CEC5375-C554-562B-A65F-***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDigitalEmployeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDigitalEmployeeResponseBody) SetRequestId(v string) *UpdateDigitalEmployeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDigitalEmployeeResponseBody) Validate() error {
	return dara.Validate(s)
}
