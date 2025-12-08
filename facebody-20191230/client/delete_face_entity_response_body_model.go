// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFaceEntityResponseBody
	GetRequestId() *string
}

type DeleteFaceEntityResponseBody struct {
	// example:
	//
	// DA7CAFEB-0A37-4496-9CDF-E3DBB6309CB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceEntityResponseBody) SetRequestId(v string) *DeleteFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
