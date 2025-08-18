// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateListResponseBody
	GetRequestId() *string
}

type UpdateListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateListResponseBody) SetRequestId(v string) *UpdateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateListResponseBody) Validate() error {
	return dara.Validate(s)
}
