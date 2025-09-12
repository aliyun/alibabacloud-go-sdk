// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseLindormInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseLindormInstanceResponseBody
	GetRequestId() *string
}

type ReleaseLindormInstanceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// F67BFFF3-F5C2-45B5-9C28-6E4A1E51****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseLindormInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseLindormInstanceResponseBody) SetRequestId(v string) *ReleaseLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseLindormInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
