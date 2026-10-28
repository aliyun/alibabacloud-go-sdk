// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendVerifyActivateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResendVerifyActivateResponseBody
	GetRequestId() *string
}

type ResendVerifyActivateResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ResendVerifyActivateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResendVerifyActivateResponseBody) GoString() string {
	return s.String()
}

func (s *ResendVerifyActivateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResendVerifyActivateResponseBody) SetRequestId(v string) *ResendVerifyActivateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendVerifyActivateResponseBody) Validate() error {
	return dara.Validate(s)
}
