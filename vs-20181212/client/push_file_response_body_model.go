// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PushFileResponseBody
	GetRequestId() *string
}

type PushFileResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushFileResponseBody) GoString() string {
	return s.String()
}

func (s *PushFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushFileResponseBody) SetRequestId(v string) *PushFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushFileResponseBody) Validate() error {
	return dara.Validate(s)
}
