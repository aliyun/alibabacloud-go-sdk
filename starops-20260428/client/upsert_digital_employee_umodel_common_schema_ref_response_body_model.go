// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody
	GetRequestId() *string
}

type UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody struct {
	// The unique identifier of the request.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-1234567890AB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody) SetRequestId(v string) *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody) Validate() error {
	return dara.Validate(s)
}
