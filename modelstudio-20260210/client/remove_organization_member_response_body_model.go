// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOrganizationMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveOrganizationMemberResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveOrganizationMemberResponseBody
	GetMessage() *string
	SetSuccess(v bool) *RemoveOrganizationMemberResponseBody
	GetSuccess() *bool
}

type RemoveOrganizationMemberResponseBody struct {
	// The error code. This parameter is empty if the operation is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is empty if the operation is successful.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveOrganizationMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveOrganizationMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveOrganizationMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveOrganizationMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveOrganizationMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveOrganizationMemberResponseBody) SetCode(v string) *RemoveOrganizationMemberResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveOrganizationMemberResponseBody) SetMessage(v string) *RemoveOrganizationMemberResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveOrganizationMemberResponseBody) SetSuccess(v bool) *RemoveOrganizationMemberResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveOrganizationMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
