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
	// example:
	//
	// 79284133-D4BA-56B3-954C-D538256F7EAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// stack-as1d4vld898ppnqbxxxxx
	StackId *string `json:"stackId,omitempty" xml:"stackId,omitempty"`
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
