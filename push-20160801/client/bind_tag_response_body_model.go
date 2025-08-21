// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindTagResponseBody
	GetRequestId() *string
}

type BindTagResponseBody struct {
	// example:
	//
	// 82FD0A09-5BB8-40FB-8221-9A11FE92D620
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindTagResponseBody) GoString() string {
	return s.String()
}

func (s *BindTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindTagResponseBody) SetRequestId(v string) *BindTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindTagResponseBody) Validate() error {
	return dara.Validate(s)
}
