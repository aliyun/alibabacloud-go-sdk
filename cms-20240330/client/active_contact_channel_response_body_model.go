// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveContactChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActiveContactChannelResponseBody
	GetRequestId() *string
}

type ActiveContactChannelResponseBody struct {
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ActiveContactChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveContactChannelResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveContactChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveContactChannelResponseBody) SetRequestId(v string) *ActiveContactChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveContactChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
