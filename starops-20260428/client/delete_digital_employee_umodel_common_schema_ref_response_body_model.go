// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody
	GetRequestId() *string
}

type DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody struct {
	// The unique identifier of the request.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-1234567890AB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody) SetRequestId(v string) *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody) Validate() error {
	return dara.Validate(s)
}
