// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWuyingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartWuyingServerResponseBody
	GetRequestId() *string
}

type StartWuyingServerResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartWuyingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartWuyingServerResponseBody) GoString() string {
	return s.String()
}

func (s *StartWuyingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartWuyingServerResponseBody) SetRequestId(v string) *StartWuyingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartWuyingServerResponseBody) Validate() error {
	return dara.Validate(s)
}
