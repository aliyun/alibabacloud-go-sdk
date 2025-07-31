// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppAuthorizedUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceAppAuthorizedUsersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetDataServiceAppAuthorizedUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceAppAuthorizedUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceAppAuthorizedUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceAppAuthorizedUsersResponseBody
	GetSuccess() *bool
	SetUserList(v []*GetDataServiceAppAuthorizedUsersResponseBodyUserList) *GetDataServiceAppAuthorizedUsersResponseBody
	GetUserList() []*GetDataServiceAppAuthorizedUsersResponseBodyUserList
}

type GetDataServiceAppAuthorizedUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	UserList  []*GetDataServiceAppAuthorizedUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetDataServiceAppAuthorizedUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppAuthorizedUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) GetUserList() []*GetDataServiceAppAuthorizedUsersResponseBodyUserList {
	return s.UserList
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) SetCode(v string) *GetDataServiceAppAuthorizedUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) SetHttpStatusCode(v int32) *GetDataServiceAppAuthorizedUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) SetMessage(v string) *GetDataServiceAppAuthorizedUsersResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) SetRequestId(v string) *GetDataServiceAppAuthorizedUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) SetSuccess(v bool) *GetDataServiceAppAuthorizedUsersResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) SetUserList(v []*GetDataServiceAppAuthorizedUsersResponseBodyUserList) *GetDataServiceAppAuthorizedUsersResponseBody {
	s.UserList = v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceAppAuthorizedUsersResponseBodyUserList struct {
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 30002101
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataServiceAppAuthorizedUsersResponseBodyUserList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppAuthorizedUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppAuthorizedUsersResponseBodyUserList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetDataServiceAppAuthorizedUsersResponseBodyUserList) GetId() *string {
	return s.Id
}

func (s *GetDataServiceAppAuthorizedUsersResponseBodyUserList) SetDisplayName(v string) *GetDataServiceAppAuthorizedUsersResponseBodyUserList {
	s.DisplayName = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponseBodyUserList) SetId(v string) *GetDataServiceAppAuthorizedUsersResponseBodyUserList {
	s.Id = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponseBodyUserList) Validate() error {
	return dara.Validate(s)
}
