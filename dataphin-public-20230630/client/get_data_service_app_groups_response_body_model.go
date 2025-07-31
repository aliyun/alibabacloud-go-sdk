// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroupList(v []*GetDataServiceAppGroupsResponseBodyAppGroupList) *GetDataServiceAppGroupsResponseBody
	GetAppGroupList() []*GetDataServiceAppGroupsResponseBodyAppGroupList
	SetCode(v string) *GetDataServiceAppGroupsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetDataServiceAppGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceAppGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceAppGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceAppGroupsResponseBody
	GetSuccess() *bool
}

type GetDataServiceAppGroupsResponseBody struct {
	AppGroupList []*GetDataServiceAppGroupsResponseBodyAppGroupList `json:"AppGroupList,omitempty" xml:"AppGroupList,omitempty" type:"Repeated"`
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

func (s GetDataServiceAppGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppGroupsResponseBody) GetAppGroupList() []*GetDataServiceAppGroupsResponseBodyAppGroupList {
	return s.AppGroupList
}

func (s *GetDataServiceAppGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceAppGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceAppGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceAppGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceAppGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceAppGroupsResponseBody) SetAppGroupList(v []*GetDataServiceAppGroupsResponseBodyAppGroupList) *GetDataServiceAppGroupsResponseBody {
	s.AppGroupList = v
	return s
}

func (s *GetDataServiceAppGroupsResponseBody) SetCode(v string) *GetDataServiceAppGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceAppGroupsResponseBody) SetHttpStatusCode(v int32) *GetDataServiceAppGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceAppGroupsResponseBody) SetMessage(v string) *GetDataServiceAppGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceAppGroupsResponseBody) SetRequestId(v string) *GetDataServiceAppGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceAppGroupsResponseBody) SetSuccess(v bool) *GetDataServiceAppGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceAppGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceAppGroupsResponseBodyAppGroupList struct {
	// example:
	//
	// 1022
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDataServiceAppGroupsResponseBodyAppGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppGroupsResponseBodyAppGroupList) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppGroupsResponseBodyAppGroupList) GetId() *int32 {
	return s.Id
}

func (s *GetDataServiceAppGroupsResponseBodyAppGroupList) GetName() *string {
	return s.Name
}

func (s *GetDataServiceAppGroupsResponseBodyAppGroupList) SetId(v int32) *GetDataServiceAppGroupsResponseBodyAppGroupList {
	s.Id = &v
	return s
}

func (s *GetDataServiceAppGroupsResponseBodyAppGroupList) SetName(v string) *GetDataServiceAppGroupsResponseBodyAppGroupList {
	s.Name = &v
	return s
}

func (s *GetDataServiceAppGroupsResponseBodyAppGroupList) Validate() error {
	return dara.Validate(s)
}
