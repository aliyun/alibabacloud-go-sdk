// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeSourceId(v string) *CreateCodeSourceResponseBody
	GetCodeSourceId() *string
	SetRequestId(v string) *CreateCodeSourceResponseBody
	GetRequestId() *string
}

type CreateCodeSourceResponseBody struct {
	// The ID of the created code build.
	//
	// example:
	//
	// code-20********
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCodeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceResponseBody) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *CreateCodeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCodeSourceResponseBody) SetCodeSourceId(v string) *CreateCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *CreateCodeSourceResponseBody) SetRequestId(v string) *CreateCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCodeSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
