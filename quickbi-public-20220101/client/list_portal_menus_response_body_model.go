// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalMenusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPortalMenusResponseBody
	GetRequestId() *string
	SetResult(v string) *ListPortalMenusResponseBody
	GetResult() *string
	SetSuccess(v bool) *ListPortalMenusResponseBody
	GetSuccess() *bool
}

type ListPortalMenusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 75912036-5527-4B7E-9265-B481D6651AC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A JSON string that levels the details of the portal menu list. Valid values:
	//
	// 	- menuType: the type of the menu.
	//
	//     	- 0: dashboard
	//
	//     	- 1: outer chain
	//
	//     	- 2: workbook
	//
	//     	- 4: directory folder
	//
	//     	- 5: form filling
	//
	//     	- 6: self-service data retrieval
	//
	// 	- menuId: menu ID
	//
	// 	- uri: ID or URL of the resource associated with the menu
	//
	// 	- showOnlyWithAccess: Authorized Only Visible
	//
	// 	- menuName: menu display name
	//
	// 	- dependentPermisson: whether the report resource associated with the menu has permissions
	//
	// 	- children: submenu
	//
	// example:
	//
	// [{"children":[{"children":[{"children":[{"menuId":"54kqgoa\\*\\*\\*\\*","menuName":"Report menu","menuType":0,"showOnlyWithAccess":true,"dependentPermisson":false,"uri":"e5da4a3f-d7f9-4262-a39e-a840043c\\*\\*\\*\\*"},{\\*\\*\\*\\	- "menu1nId":"pName" 135 "Directory menu","menuType":4,"showOnlyWithAccess":false,"dependentPermisson":true}],"menuId":"23a7d5d8-e55a-4737-b6a1-3c585505\\*\\*\\*\\*","menuName":"pop level -3 menu","menuType":4,"showOnlyWithAccess":true,"dependentPermisson":true}],"menuId":"80764 f3c-affd-45a1-aaa1-bb039d8a\\*\\*\\*\\*","menuName":"pop menu","menuType":4,"showOnlyWithAccess":false,"dependentPermisson":true}],"menuId":"277 f968a-22 ff-4ce6-83f0-a82950f4\\*\\*\\*\\*","menuName":"pop menu","menuType":4,"showOnlyWithAccess":false,"dependentPermisson":true}]
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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

func (s ListPortalMenusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPortalMenusResponseBody) GoString() string {
	return s.String()
}

func (s *ListPortalMenusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPortalMenusResponseBody) GetResult() *string {
	return s.Result
}

func (s *ListPortalMenusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPortalMenusResponseBody) SetRequestId(v string) *ListPortalMenusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPortalMenusResponseBody) SetResult(v string) *ListPortalMenusResponseBody {
	s.Result = &v
	return s
}

func (s *ListPortalMenusResponseBody) SetSuccess(v bool) *ListPortalMenusResponseBody {
	s.Success = &v
	return s
}

func (s *ListPortalMenusResponseBody) Validate() error {
	return dara.Validate(s)
}
