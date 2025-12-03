// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodAddComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MultimodAddComponentsResponseBody
	GetRequestId() *string
}

type MultimodAddComponentsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultimodAddComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultimodAddComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *MultimodAddComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultimodAddComponentsResponseBody) SetRequestId(v string) *MultimodAddComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultimodAddComponentsResponseBody) Validate() error {
	return dara.Validate(s)
}
