// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChangeSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSetId(v string) *CreateChangeSetResponseBody
	GetChangeSetId() *string
	SetRequestId(v string) *CreateChangeSetResponseBody
	GetRequestId() *string
	SetStackId(v string) *CreateChangeSetResponseBody
	GetStackId() *string
}

type CreateChangeSetResponseBody struct {
	// The ID of the change set.
	//
	// example:
	//
	// e85abe0c-6528-43fb-ae93-fdf8de22****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CreateChangeSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChangeSetResponseBody) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *CreateChangeSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChangeSetResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *CreateChangeSetResponseBody) SetChangeSetId(v string) *CreateChangeSetResponseBody {
	s.ChangeSetId = &v
	return s
}

func (s *CreateChangeSetResponseBody) SetRequestId(v string) *CreateChangeSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChangeSetResponseBody) SetStackId(v string) *CreateChangeSetResponseBody {
	s.StackId = &v
	return s
}

func (s *CreateChangeSetResponseBody) Validate() error {
	return dara.Validate(s)
}
