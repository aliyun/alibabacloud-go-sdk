// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEndUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachEndUserResponseBody
	GetRequestId() *string
}

type AttachEndUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEndUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEndUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachEndUserResponseBody) SetRequestId(v string) *AttachEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachEndUserResponseBody) Validate() error {
	return dara.Validate(s)
}
