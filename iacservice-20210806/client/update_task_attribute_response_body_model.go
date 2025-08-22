// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTaskAttributeResponseBody
	GetRequestId() *string
}

type UpdateTaskAttributeResponseBody struct {
	// example:
	//
	// 17793D91-A26F-520D-A948-3452A45D15B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateTaskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskAttributeResponseBody) SetRequestId(v string) *UpdateTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
