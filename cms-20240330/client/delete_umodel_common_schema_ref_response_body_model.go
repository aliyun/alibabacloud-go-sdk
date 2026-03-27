// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelCommonSchemaRefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUmodelCommonSchemaRefResponseBody
	GetRequestId() *string
}

type DeleteUmodelCommonSchemaRefResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteUmodelCommonSchemaRefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelCommonSchemaRefResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUmodelCommonSchemaRefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUmodelCommonSchemaRefResponseBody) SetRequestId(v string) *DeleteUmodelCommonSchemaRefResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUmodelCommonSchemaRefResponseBody) Validate() error {
	return dara.Validate(s)
}
