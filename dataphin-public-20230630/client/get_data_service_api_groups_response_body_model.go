// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiGroupList(v []*GetDataServiceApiGroupsResponseBodyApiGroupList) *GetDataServiceApiGroupsResponseBody
	GetApiGroupList() []*GetDataServiceApiGroupsResponseBodyApiGroupList
	SetCode(v string) *GetDataServiceApiGroupsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetDataServiceApiGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceApiGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceApiGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceApiGroupsResponseBody
	GetSuccess() *bool
}

type GetDataServiceApiGroupsResponseBody struct {
	ApiGroupList []*GetDataServiceApiGroupsResponseBodyApiGroupList `json:"ApiGroupList,omitempty" xml:"ApiGroupList,omitempty" type:"Repeated"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceApiGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiGroupsResponseBody) GetApiGroupList() []*GetDataServiceApiGroupsResponseBodyApiGroupList {
	return s.ApiGroupList
}

func (s *GetDataServiceApiGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceApiGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceApiGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceApiGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceApiGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceApiGroupsResponseBody) SetApiGroupList(v []*GetDataServiceApiGroupsResponseBodyApiGroupList) *GetDataServiceApiGroupsResponseBody {
	s.ApiGroupList = v
	return s
}

func (s *GetDataServiceApiGroupsResponseBody) SetCode(v string) *GetDataServiceApiGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceApiGroupsResponseBody) SetHttpStatusCode(v int32) *GetDataServiceApiGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceApiGroupsResponseBody) SetMessage(v string) *GetDataServiceApiGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceApiGroupsResponseBody) SetRequestId(v string) *GetDataServiceApiGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceApiGroupsResponseBody) SetSuccess(v bool) *GetDataServiceApiGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceApiGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiGroupsResponseBodyApiGroupList struct {
	// example:
	//
	// 101231
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDataServiceApiGroupsResponseBodyApiGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiGroupsResponseBodyApiGroupList) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiGroupsResponseBodyApiGroupList) GetId() *int32 {
	return s.Id
}

func (s *GetDataServiceApiGroupsResponseBodyApiGroupList) GetName() *string {
	return s.Name
}

func (s *GetDataServiceApiGroupsResponseBodyApiGroupList) SetId(v int32) *GetDataServiceApiGroupsResponseBodyApiGroupList {
	s.Id = &v
	return s
}

func (s *GetDataServiceApiGroupsResponseBodyApiGroupList) SetName(v string) *GetDataServiceApiGroupsResponseBodyApiGroupList {
	s.Name = &v
	return s
}

func (s *GetDataServiceApiGroupsResponseBodyApiGroupList) Validate() error {
	return dara.Validate(s)
}
