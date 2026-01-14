// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpSetResponseBody
	GetRequestId() *string
}

type UpdateIpSetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7D2F7E4E-B958-439C-9821-56D6213A31EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpSetResponseBody) SetRequestId(v string) *UpdateIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpSetResponseBody) Validate() error {
	return dara.Validate(s)
}
