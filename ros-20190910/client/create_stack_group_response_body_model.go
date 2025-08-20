// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateStackGroupResponseBody
	GetRequestId() *string
	SetStackGroupId(v string) *CreateStackGroupResponseBody
	GetStackGroupId() *string
}

type CreateStackGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the stack group.
	//
	// example:
	//
	// 2c036e78-9e82-428e-afd6-177f5d04****
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
}

func (s CreateStackGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStackGroupResponseBody) GetStackGroupId() *string {
	return s.StackGroupId
}

func (s *CreateStackGroupResponseBody) SetRequestId(v string) *CreateStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStackGroupResponseBody) SetStackGroupId(v string) *CreateStackGroupResponseBody {
	s.StackGroupId = &v
	return s
}

func (s *CreateStackGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
