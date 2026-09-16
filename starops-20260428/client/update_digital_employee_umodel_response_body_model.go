// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeUmodelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDigitalEmployeeUmodelResponseBody
	GetRequestId() *string
}

type UpdateDigitalEmployeeUmodelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-1234567890AB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDigitalEmployeeUmodelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeUmodelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDigitalEmployeeUmodelResponseBody) SetRequestId(v string) *UpdateDigitalEmployeeUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDigitalEmployeeUmodelResponseBody) Validate() error {
	return dara.Validate(s)
}
