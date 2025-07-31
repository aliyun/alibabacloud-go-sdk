// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAuthorizedAppsByGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfoList(v []*GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) *GetDataServiceAuthorizedAppsByGroupIdResponseBody
	GetAppInfoList() []*GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList
	SetCode(v string) *GetDataServiceAuthorizedAppsByGroupIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetDataServiceAuthorizedAppsByGroupIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceAuthorizedAppsByGroupIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceAuthorizedAppsByGroupIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceAuthorizedAppsByGroupIdResponseBody
	GetSuccess() *bool
}

type GetDataServiceAuthorizedAppsByGroupIdResponseBody struct {
	AppInfoList []*GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
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

func (s GetDataServiceAuthorizedAppsByGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAuthorizedAppsByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) GetAppInfoList() []*GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList {
	return s.AppInfoList
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) SetAppInfoList(v []*GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) *GetDataServiceAuthorizedAppsByGroupIdResponseBody {
	s.AppInfoList = v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) SetCode(v string) *GetDataServiceAuthorizedAppsByGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) SetHttpStatusCode(v int32) *GetDataServiceAuthorizedAppsByGroupIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) SetMessage(v string) *GetDataServiceAuthorizedAppsByGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) SetRequestId(v string) *GetDataServiceAuthorizedAppsByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) SetSuccess(v bool) *GetDataServiceAuthorizedAppsByGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList struct {
	// AppKey
	//
	// example:
	//
	// 202212
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// AppId
	//
	// example:
	//
	// 1021
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) GetId() *int32 {
	return s.Id
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) GetName() *string {
	return s.Name
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) SetAppKey(v int64) *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList {
	s.AppKey = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) SetId(v int32) *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList {
	s.Id = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) SetName(v string) *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList {
	s.Name = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponseBodyAppInfoList) Validate() error {
	return dara.Validate(s)
}
