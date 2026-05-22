// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIPv6ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIPv6ResponseBody
	GetRequestId() *string
}

type UpdateIPv6ResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIPv6ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIPv6ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIPv6ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIPv6ResponseBody) SetRequestId(v string) *UpdateIPv6ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIPv6ResponseBody) Validate() error {
	return dara.Validate(s)
}
