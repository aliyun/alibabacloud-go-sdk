// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSecureMobilePhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindSecureMobilePhoneResponseBody
	GetRequestId() *string
}

type BindSecureMobilePhoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0217AFEB-5318-56D4-B167-1933D83EDF3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindSecureMobilePhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindSecureMobilePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindSecureMobilePhoneResponseBody) SetRequestId(v string) *BindSecureMobilePhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSecureMobilePhoneResponseBody) Validate() error {
	return dara.Validate(s)
}
