// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertUmodelCommonSchemaRefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpsertUmodelCommonSchemaRefResponseBody
	GetRequestId() *string
}

type UpsertUmodelCommonSchemaRefResponseBody struct {
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpsertUmodelCommonSchemaRefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertUmodelCommonSchemaRefResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertUmodelCommonSchemaRefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertUmodelCommonSchemaRefResponseBody) SetRequestId(v string) *UpsertUmodelCommonSchemaRefResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertUmodelCommonSchemaRefResponseBody) Validate() error {
	return dara.Validate(s)
}
