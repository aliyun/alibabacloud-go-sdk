// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserGroupsResponseBody
	GetCode() *string
	SetItems(v []interface{}) *ListUserGroupsResponseBody
	GetItems() []interface{}
	SetMessage(v string) *ListUserGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUserGroupsResponseBody
	GetRequestId() *string
}

type ListUserGroupsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of top-level user groups.
	Items []interface{} `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUserGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserGroupsResponseBody) GetItems() []interface{} {
	return s.Items
}

func (s *ListUserGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsResponseBody) SetCode(v string) *ListUserGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetItems(v []interface{}) *ListUserGroupsResponseBody {
	s.Items = v
	return s
}

func (s *ListUserGroupsResponseBody) SetMessage(v string) *ListUserGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
