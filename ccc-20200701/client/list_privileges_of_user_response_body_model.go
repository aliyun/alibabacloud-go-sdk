// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivilegesOfUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPrivilegesOfUserResponseBody
	GetCode() *string
	SetData(v []*ListPrivilegesOfUserResponseBodyData) *ListPrivilegesOfUserResponseBody
	GetData() []*ListPrivilegesOfUserResponseBodyData
	SetHttpStatusCode(v int32) *ListPrivilegesOfUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPrivilegesOfUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrivilegesOfUserResponseBody
	GetRequestId() *string
}

type ListPrivilegesOfUserResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListPrivilegesOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivilegesOfUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivilegesOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivilegesOfUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPrivilegesOfUserResponseBody) GetData() []*ListPrivilegesOfUserResponseBodyData {
	return s.Data
}

func (s *ListPrivilegesOfUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPrivilegesOfUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrivilegesOfUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivilegesOfUserResponseBody) SetCode(v string) *ListPrivilegesOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) SetData(v []*ListPrivilegesOfUserResponseBodyData) *ListPrivilegesOfUserResponseBody {
	s.Data = v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) SetHttpStatusCode(v int32) *ListPrivilegesOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) SetMessage(v string) *ListPrivilegesOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) SetRequestId(v string) *ListPrivilegesOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrivilegesOfUserResponseBodyData struct {
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Workbench:Call
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SELF_ONLY
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListPrivilegesOfUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrivilegesOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrivilegesOfUserResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPrivilegesOfUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListPrivilegesOfUserResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *ListPrivilegesOfUserResponseBodyData) SetInstanceId(v string) *ListPrivilegesOfUserResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBodyData) SetName(v string) *ListPrivilegesOfUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBodyData) SetScope(v string) *ListPrivilegesOfUserResponseBodyData {
	s.Scope = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
