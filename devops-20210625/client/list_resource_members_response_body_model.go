// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListResourceMembersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListResourceMembersResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListResourceMembersResponseBody
	GetRequestId() *string
	SetResourceMembers(v []*ListResourceMembersResponseBodyResourceMembers) *ListResourceMembersResponseBody
	GetResourceMembers() []*ListResourceMembersResponseBodyResourceMembers
	SetSuccess(v bool) *ListResourceMembersResponseBody
	GetSuccess() *bool
}

type ListResourceMembersResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId       *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceMembers []*ListResourceMembersResponseBodyResourceMembers `json:"resourceMembers,omitempty" xml:"resourceMembers,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListResourceMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceMembersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListResourceMembersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListResourceMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceMembersResponseBody) GetResourceMembers() []*ListResourceMembersResponseBodyResourceMembers {
	return s.ResourceMembers
}

func (s *ListResourceMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourceMembersResponseBody) SetErrorCode(v string) *ListResourceMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListResourceMembersResponseBody) SetErrorMessage(v string) *ListResourceMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListResourceMembersResponseBody) SetRequestId(v string) *ListResourceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceMembersResponseBody) SetResourceMembers(v []*ListResourceMembersResponseBodyResourceMembers) *ListResourceMembersResponseBody {
	s.ResourceMembers = v
	return s
}

func (s *ListResourceMembersResponseBody) SetSuccess(v bool) *ListResourceMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourceMembersResponseBody) Validate() error {
	if s.ResourceMembers != nil {
		for _, item := range s.ResourceMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceMembersResponseBodyResourceMembers struct {
	// example:
	//
	// 22212212
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// admin
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// example:
	//
	// 张三
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListResourceMembersResponseBodyResourceMembers) String() string {
	return dara.Prettify(s)
}

func (s ListResourceMembersResponseBodyResourceMembers) GoString() string {
	return s.String()
}

func (s *ListResourceMembersResponseBodyResourceMembers) GetAccountId() *string {
	return s.AccountId
}

func (s *ListResourceMembersResponseBodyResourceMembers) GetRoleName() *string {
	return s.RoleName
}

func (s *ListResourceMembersResponseBodyResourceMembers) GetUsername() *string {
	return s.Username
}

func (s *ListResourceMembersResponseBodyResourceMembers) SetAccountId(v string) *ListResourceMembersResponseBodyResourceMembers {
	s.AccountId = &v
	return s
}

func (s *ListResourceMembersResponseBodyResourceMembers) SetRoleName(v string) *ListResourceMembersResponseBodyResourceMembers {
	s.RoleName = &v
	return s
}

func (s *ListResourceMembersResponseBodyResourceMembers) SetUsername(v string) *ListResourceMembersResponseBodyResourceMembers {
	s.Username = &v
	return s
}

func (s *ListResourceMembersResponseBodyResourceMembers) Validate() error {
	return dara.Validate(s)
}
