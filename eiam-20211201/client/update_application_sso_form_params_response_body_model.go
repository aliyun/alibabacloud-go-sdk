// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationSsoFormParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationSsoFormParamsResponseBody
	GetRequestId() *string
}

type UpdateApplicationSsoFormParamsResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationSsoFormParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationSsoFormParamsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationSsoFormParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationSsoFormParamsResponseBody) SetRequestId(v string) *UpdateApplicationSsoFormParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationSsoFormParamsResponseBody) Validate() error {
	return dara.Validate(s)
}
