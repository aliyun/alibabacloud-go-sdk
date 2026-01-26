// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomHostnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomHostnameResponseBody
	GetRequestId() *string
}

type UpdateCustomHostnameResponseBody struct {
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomHostnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomHostnameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomHostnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomHostnameResponseBody) SetRequestId(v string) *UpdateCustomHostnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomHostnameResponseBody) Validate() error {
	return dara.Validate(s)
}
