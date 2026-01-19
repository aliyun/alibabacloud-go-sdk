// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationM2MClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableApplicationM2MClientResponseBody
	GetRequestId() *string
}

type DisableApplicationM2MClientResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationM2MClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationM2MClientResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationM2MClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationM2MClientResponseBody) SetRequestId(v string) *DisableApplicationM2MClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationM2MClientResponseBody) Validate() error {
	return dara.Validate(s)
}
