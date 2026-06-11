// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalMenuAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPortalMenuAuthorizationResponseBody
	GetRequestId() *string
	SetResult(v []*ListPortalMenuAuthorizationResponseBodyResult) *ListPortalMenuAuthorizationResponseBody
	GetResult() []*ListPortalMenuAuthorizationResponseBodyResult
	SetSuccess(v bool) *ListPortalMenuAuthorizationResponseBody
	GetSuccess() *bool
}

type ListPortalMenuAuthorizationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of authorization details for the BI portal menus.
	Result []*ListPortalMenuAuthorizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPortalMenuAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPortalMenuAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPortalMenuAuthorizationResponseBody) GetResult() []*ListPortalMenuAuthorizationResponseBodyResult {
	return s.Result
}

func (s *ListPortalMenuAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPortalMenuAuthorizationResponseBody) SetRequestId(v string) *ListPortalMenuAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBody) SetResult(v []*ListPortalMenuAuthorizationResponseBodyResult) *ListPortalMenuAuthorizationResponseBody {
	s.Result = v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBody) SetSuccess(v bool) *ListPortalMenuAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPortalMenuAuthorizationResponseBodyResult struct {
	// The ID of the leaf-node menu in the BI portal.
	//
	// example:
	//
	// 54kqgoa****
	MenuId *string `json:"MenuId,omitempty" xml:"MenuId,omitempty"`
	// The details of the authorization objects for the menu.
	Receivers []*ListPortalMenuAuthorizationResponseBodyResultReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// Indicates whether the menu is visible only to authorized users. Valid values:
	//
	// - true: The menu is visible only to authorized users.
	//
	// - false: The menu is visible to all users.
	//
	// example:
	//
	// true
	ShowOnlyWithAccess *bool `json:"ShowOnlyWithAccess,omitempty" xml:"ShowOnlyWithAccess,omitempty"`
}

func (s ListPortalMenuAuthorizationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListPortalMenuAuthorizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) GetMenuId() *string {
	return s.MenuId
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) GetReceivers() []*ListPortalMenuAuthorizationResponseBodyResultReceivers {
	return s.Receivers
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) GetShowOnlyWithAccess() *bool {
	return s.ShowOnlyWithAccess
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) SetMenuId(v string) *ListPortalMenuAuthorizationResponseBodyResult {
	s.MenuId = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) SetReceivers(v []*ListPortalMenuAuthorizationResponseBodyResultReceivers) *ListPortalMenuAuthorizationResponseBodyResult {
	s.Receivers = v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) SetShowOnlyWithAccess(v bool) *ListPortalMenuAuthorizationResponseBodyResult {
	s.ShowOnlyWithAccess = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) Validate() error {
	if s.Receivers != nil {
		for _, item := range s.Receivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPortalMenuAuthorizationResponseBodyResultReceivers struct {
	// The authorization type for the menu. Valid values:
	//
	// - 1: View
	//
	// - 11: Edit
	//
	// - 3: Export and view
	//
	// - 10: Manage data entry
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	AuthPointsValue *int32 `json:"AuthPointsValue,omitempty" xml:"AuthPointsValue,omitempty"`
	// The ID of the authorization object.
	//
	// > - If the authorization object is an organization, this ID is the organization ID.
	//
	// >
	//
	// > - If the authorization object is a workspace, this ID is the workspace ID.
	//
	// example:
	//
	// 121344444790****
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// The type of the authorization object. Valid values:
	//
	// - 0: User
	//
	// - 1: User group
	//
	// - 2: Organization
	//
	// - 3: Workspace
	//
	// example:
	//
	// 0
	ReceiverType *int32 `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
}

func (s ListPortalMenuAuthorizationResponseBodyResultReceivers) String() string {
	return dara.Prettify(s)
}

func (s ListPortalMenuAuthorizationResponseBodyResultReceivers) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) GetAuthPointsValue() *int32 {
	return s.AuthPointsValue
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) GetReceiverType() *int32 {
	return s.ReceiverType
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) SetAuthPointsValue(v int32) *ListPortalMenuAuthorizationResponseBodyResultReceivers {
	s.AuthPointsValue = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) SetReceiverId(v string) *ListPortalMenuAuthorizationResponseBodyResultReceivers {
	s.ReceiverId = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) SetReceiverType(v int32) *ListPortalMenuAuthorizationResponseBodyResultReceivers {
	s.ReceiverType = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) Validate() error {
	return dara.Validate(s)
}
