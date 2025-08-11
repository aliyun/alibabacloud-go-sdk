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
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of authorization details of the portal menu.
	Result []*ListPortalMenuAuthorizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
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
	return dara.Validate(s)
}

type ListPortalMenuAuthorizationResponseBodyResult struct {
	// The menu ID of the BI portal leaf node.
	//
	// example:
	//
	// 54kqgoa****
	MenuId *string `json:"MenuId,omitempty" xml:"MenuId,omitempty"`
	// The details of the object to which the menu is authorized.
	Receivers []*ListPortalMenuAuthorizationResponseBodyResultReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// Whether only authorization is visible. Valid values:
	//
	// 	- true: Only the authorization is visible.
	//
	// 	- false: Both are visible.
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
	return dara.Validate(s)
}

type ListPortalMenuAuthorizationResponseBodyResultReceivers struct {
	// if can be null:
	// true
	AuthPointsValue *int32 `json:"AuthPointsValue,omitempty" xml:"AuthPointsValue,omitempty"`
	// The ID of the authorization object.
	//
	// example:
	//
	// 121344444790****
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// The type of the authorization object. Valid values:
	//
	// 	- 0: user
	//
	// 	- 1: user group
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
