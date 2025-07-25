// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpamPoolResponseBody
	GetRequestId() *string
}

type UpdateIpamPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DED57B9-7654-5B6D-93F7-BCA5839FEE38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpamPoolResponseBody) SetRequestId(v string) *UpdateIpamPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpamPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
