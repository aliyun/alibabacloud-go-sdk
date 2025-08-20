// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateStackResponseBody
	GetRequestId() *string
	SetStackId(v string) *CreateStackResponseBody
	GetStackId() *string
}

type CreateStackResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack ID.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CreateStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStackResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *CreateStackResponseBody) SetRequestId(v string) *CreateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStackResponseBody) SetStackId(v string) *CreateStackResponseBody {
	s.StackId = &v
	return s
}

func (s *CreateStackResponseBody) Validate() error {
	return dara.Validate(s)
}
