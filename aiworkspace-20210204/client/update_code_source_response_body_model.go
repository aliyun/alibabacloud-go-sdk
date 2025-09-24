// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCodeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeSourceId(v string) *UpdateCodeSourceResponseBody
	GetCodeSourceId() *string
	SetRequestId(v string) *UpdateCodeSourceResponseBody
	GetRequestId() *string
}

type UpdateCodeSourceResponseBody struct {
	// The ID of the code build.
	//
	// example:
	//
	// code-20********
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCodeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCodeSourceResponseBody) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *UpdateCodeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCodeSourceResponseBody) SetCodeSourceId(v string) *UpdateCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *UpdateCodeSourceResponseBody) SetRequestId(v string) *UpdateCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCodeSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
