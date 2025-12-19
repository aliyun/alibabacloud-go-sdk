// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteResourceTokenAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompleteResourceTokenAuthResponseBody
	GetRequestId() *string
}

type CompleteResourceTokenAuthResponseBody struct {
	// example:
	//
	// 6A3344F3-BFA1-5BCB-B5F5-17572C0D5586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompleteResourceTokenAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteResourceTokenAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteResourceTokenAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompleteResourceTokenAuthResponseBody) SetRequestId(v string) *CompleteResourceTokenAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteResourceTokenAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
