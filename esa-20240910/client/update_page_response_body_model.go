// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePageResponseBody
	GetRequestId() *string
}

type UpdatePageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePageResponseBody) SetRequestId(v string) *UpdatePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePageResponseBody) Validate() error {
	return dara.Validate(s)
}
