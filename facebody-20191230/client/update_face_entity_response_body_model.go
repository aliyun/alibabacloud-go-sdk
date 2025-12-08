// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaceEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFaceEntityResponseBody
	GetRequestId() *string
}

type UpdateFaceEntityResponseBody struct {
	// example:
	//
	// DA7CAFEB-0A37-4496-9CDF-E3DBB6309CB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFaceEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFaceEntityResponseBody) SetRequestId(v string) *UpdateFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFaceEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
