// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTagResponseBody
	GetRequestId() *string
}

type CreateTagResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) Validate() error {
	return dara.Validate(s)
}
