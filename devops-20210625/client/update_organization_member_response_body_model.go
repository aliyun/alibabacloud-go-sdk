// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateOrganizationMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateOrganizationMemberResponseBody
	GetErrorMessage() *string
	SetMember(v *UpdateOrganizationMemberResponseBodyMember) *UpdateOrganizationMemberResponseBody
	GetMember() *UpdateOrganizationMemberResponseBodyMember
	SetRequestId(v string) *UpdateOrganizationMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOrganizationMemberResponseBody
	GetSuccess() *bool
}

type UpdateOrganizationMemberResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string                                     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Member       *UpdateOrganizationMemberResponseBodyMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateOrganizationMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateOrganizationMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateOrganizationMemberResponseBody) GetMember() *UpdateOrganizationMemberResponseBodyMember {
	return s.Member
}

func (s *UpdateOrganizationMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOrganizationMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOrganizationMemberResponseBody) SetErrorCode(v string) *UpdateOrganizationMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) SetErrorMessage(v string) *UpdateOrganizationMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) SetMember(v *UpdateOrganizationMemberResponseBodyMember) *UpdateOrganizationMemberResponseBody {
	s.Member = v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) SetRequestId(v string) *UpdateOrganizationMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) SetSuccess(v bool) *UpdateOrganizationMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) Validate() error {
	if s.Member != nil {
		if err := s.Member.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateOrganizationMemberResponseBodyMember struct {
	// example:
	//
	// 292035769476261xxx
	AccountId              *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	OrganizationMemberName *string `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
}

func (s UpdateOrganizationMemberResponseBodyMember) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationMemberResponseBodyMember) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationMemberResponseBodyMember) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateOrganizationMemberResponseBodyMember) GetOrganizationMemberName() *string {
	return s.OrganizationMemberName
}

func (s *UpdateOrganizationMemberResponseBodyMember) SetAccountId(v string) *UpdateOrganizationMemberResponseBodyMember {
	s.AccountId = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBodyMember) SetOrganizationMemberName(v string) *UpdateOrganizationMemberResponseBodyMember {
	s.OrganizationMemberName = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBodyMember) Validate() error {
	return dara.Validate(s)
}
