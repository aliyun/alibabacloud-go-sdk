// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyUserScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetPolicyUserScopeResponseBody
	GetRequestId() *string
	SetUserScope(v *GetPolicyUserScopeResponseBodyUserScope) *GetPolicyUserScopeResponseBody
	GetUserScope() *GetPolicyUserScopeResponseBodyUserScope
}

type GetPolicyUserScopeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The users to whom the control policy applies.
	UserScope *GetPolicyUserScopeResponseBodyUserScope `json:"UserScope,omitempty" xml:"UserScope,omitempty" type:"Struct"`
}

func (s GetPolicyUserScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyUserScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyUserScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyUserScopeResponseBody) GetUserScope() *GetPolicyUserScopeResponseBodyUserScope {
	return s.UserScope
}

func (s *GetPolicyUserScopeResponseBody) SetRequestId(v string) *GetPolicyUserScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyUserScopeResponseBody) SetUserScope(v *GetPolicyUserScopeResponseBodyUserScope) *GetPolicyUserScopeResponseBody {
	s.UserScope = v
	return s
}

func (s *GetPolicyUserScopeResponseBody) Validate() error {
	if s.UserScope != nil {
		if err := s.UserScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPolicyUserScopeResponseBodyUserScope struct {
	// The scope of users to whom the control policy applies.
	//
	// 	- If **All*	- is returned for this parameter, the control policy applies to all users.
	//
	// 	- If no value is returned for this parameter, the control policy applies to the assets specified in the return values of UserGroupIds and UserIds.
	//
	// example:
	//
	// All
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The user groups to which the control policy applies.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The users to whom the control policy applies.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetPolicyUserScopeResponseBodyUserScope) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyUserScopeResponseBodyUserScope) GoString() string {
	return s.String()
}

func (s *GetPolicyUserScopeResponseBodyUserScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetPolicyUserScopeResponseBodyUserScope) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *GetPolicyUserScopeResponseBodyUserScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetPolicyUserScopeResponseBodyUserScope) SetScopeType(v string) *GetPolicyUserScopeResponseBodyUserScope {
	s.ScopeType = &v
	return s
}

func (s *GetPolicyUserScopeResponseBodyUserScope) SetUserGroupIds(v []*string) *GetPolicyUserScopeResponseBodyUserScope {
	s.UserGroupIds = v
	return s
}

func (s *GetPolicyUserScopeResponseBodyUserScope) SetUserIds(v []*string) *GetPolicyUserScopeResponseBodyUserScope {
	s.UserIds = v
	return s
}

func (s *GetPolicyUserScopeResponseBodyUserScope) Validate() error {
	return dara.Validate(s)
}
