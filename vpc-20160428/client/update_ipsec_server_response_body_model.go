// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpsecServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpsecServerResponseBody
	GetRequestId() *string
}

type UpdateIpsecServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B61C08E5-403A-46A2-96C1-F7B1216DB10C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpsecServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpsecServerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpsecServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpsecServerResponseBody) SetRequestId(v string) *UpdateIpsecServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpsecServerResponseBody) Validate() error {
	return dara.Validate(s)
}
