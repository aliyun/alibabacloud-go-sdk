// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityList(v *ListAuthorityResponseBodyAuthorityList) *ListAuthorityResponseBody
	GetAuthorityList() *ListAuthorityResponseBodyAuthorityList
	SetCode(v int32) *ListAuthorityResponseBody
	GetCode() *int32
	SetMessage(v string) *ListAuthorityResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAuthorityResponseBody
	GetRequestId() *string
}

type ListAuthorityResponseBody struct {
	// The permissions.
	AuthorityList *ListAuthorityResponseBodyAuthorityList `json:"AuthorityList,omitempty" xml:"AuthorityList,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57609587-DFA2-41EC-****-*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBody) GetAuthorityList() *ListAuthorityResponseBodyAuthorityList {
	return s.AuthorityList
}

func (s *ListAuthorityResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAuthorityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorityResponseBody) SetAuthorityList(v *ListAuthorityResponseBodyAuthorityList) *ListAuthorityResponseBody {
	s.AuthorityList = v
	return s
}

func (s *ListAuthorityResponseBody) SetCode(v int32) *ListAuthorityResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuthorityResponseBody) SetMessage(v string) *ListAuthorityResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuthorityResponseBody) SetRequestId(v string) *ListAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorityResponseBody) Validate() error {
	if s.AuthorityList != nil {
		if err := s.AuthorityList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorityResponseBodyAuthorityList struct {
	Authority []*ListAuthorityResponseBodyAuthorityListAuthority `json:"Authority,omitempty" xml:"Authority,omitempty" type:"Repeated"`
}

func (s ListAuthorityResponseBodyAuthorityList) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityResponseBodyAuthorityList) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBodyAuthorityList) GetAuthority() []*ListAuthorityResponseBodyAuthorityListAuthority {
	return s.Authority
}

func (s *ListAuthorityResponseBodyAuthorityList) SetAuthority(v []*ListAuthorityResponseBodyAuthorityListAuthority) *ListAuthorityResponseBodyAuthorityList {
	s.Authority = v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityList) Validate() error {
	if s.Authority != nil {
		for _, item := range s.Authority {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorityResponseBodyAuthorityListAuthority struct {
	// The set of permissions.
	ActionList *ListAuthorityResponseBodyAuthorityListAuthorityActionList `json:"ActionList,omitempty" xml:"ActionList,omitempty" type:"Struct"`
	// The description of the permission group.
	//
	// example:
	//
	// Operations on applications
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the permission group.
	//
	// example:
	//
	// 1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the permission group.
	//
	// example:
	//
	// Application management
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAuthorityResponseBodyAuthorityListAuthority) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityResponseBodyAuthorityListAuthority) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) GetActionList() *ListAuthorityResponseBodyAuthorityListAuthorityActionList {
	return s.ActionList
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) GetDescription() *string {
	return s.Description
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) GetGroupId() *string {
	return s.GroupId
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) GetName() *string {
	return s.Name
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) SetActionList(v *ListAuthorityResponseBodyAuthorityListAuthorityActionList) *ListAuthorityResponseBodyAuthorityListAuthority {
	s.ActionList = v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) SetDescription(v string) *ListAuthorityResponseBodyAuthorityListAuthority {
	s.Description = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) SetGroupId(v string) *ListAuthorityResponseBodyAuthorityListAuthority {
	s.GroupId = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) SetName(v string) *ListAuthorityResponseBodyAuthorityListAuthority {
	s.Name = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) Validate() error {
	if s.ActionList != nil {
		if err := s.ActionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorityResponseBodyAuthorityListAuthorityActionList struct {
	Action []*ListAuthorityResponseBodyAuthorityListAuthorityActionListAction `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
}

func (s ListAuthorityResponseBodyAuthorityListAuthorityActionList) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityResponseBodyAuthorityListAuthorityActionList) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionList) GetAction() []*ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	return s.Action
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionList) SetAction(v []*ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) *ListAuthorityResponseBodyAuthorityListAuthorityActionList {
	s.Action = v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionList) Validate() error {
	if s.Action != nil {
		for _, item := range s.Action {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorityResponseBodyAuthorityListAuthorityActionListAction struct {
	// The code of the permission.
	//
	// example:
	//
	// 1
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the permission.
	//
	// example:
	//
	// Create an application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the permission group.
	//
	// example:
	//
	// 1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// Create an application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) GetCode() *string {
	return s.Code
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) GetDescription() *string {
	return s.Description
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) GetGroupId() *string {
	return s.GroupId
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) GetName() *string {
	return s.Name
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) SetCode(v string) *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	s.Code = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) SetDescription(v string) *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	s.Description = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) SetGroupId(v string) *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	s.GroupId = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) SetName(v string) *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	s.Name = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) Validate() error {
	return dara.Validate(s)
}
