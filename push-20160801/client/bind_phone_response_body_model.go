// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindPhoneResponseBody
	GetRequestId() *string
}

type BindPhoneResponseBody struct {
	// example:
	//
	// 0D1126F0-F8FF-513D-BAFA-F140447BDED4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindPhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *BindPhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindPhoneResponseBody) SetRequestId(v string) *BindPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPhoneResponseBody) Validate() error {
	return dara.Validate(s)
}
