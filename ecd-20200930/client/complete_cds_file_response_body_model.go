// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteCdsFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompleteCdsFileResponseBody
	GetRequestId() *string
}

type CompleteCdsFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05C2791F-41A7-5E7C-B5E4-1401FD0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompleteCdsFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteCdsFileResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteCdsFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompleteCdsFileResponseBody) SetRequestId(v string) *CompleteCdsFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteCdsFileResponseBody) Validate() error {
	return dara.Validate(s)
}
