// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataServiceAppsResponseBody
	GetCode() *string
	SetData(v *ListDataServiceAppsResponseBodyData) *ListDataServiceAppsResponseBody
	GetData() *ListDataServiceAppsResponseBodyData
	SetHttpStatusCode(v int32) *ListDataServiceAppsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataServiceAppsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDataServiceAppsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceAppsResponseBody
	GetSuccess() *bool
}

type ListDataServiceAppsResponseBody struct {
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListDataServiceAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListDataServiceAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceAppsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataServiceAppsResponseBody) GetData() *ListDataServiceAppsResponseBodyData {
	return s.Data
}

func (s *ListDataServiceAppsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceAppsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataServiceAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceAppsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceAppsResponseBody) SetCode(v string) *ListDataServiceAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataServiceAppsResponseBody) SetData(v *ListDataServiceAppsResponseBodyData) *ListDataServiceAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataServiceAppsResponseBody) SetHttpStatusCode(v int32) *ListDataServiceAppsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceAppsResponseBody) SetMessage(v string) *ListDataServiceAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataServiceAppsResponseBody) SetRequestId(v string) *ListDataServiceAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceAppsResponseBody) SetSuccess(v bool) *ListDataServiceAppsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceAppsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceAppsResponseBodyData struct {
	AppList []*ListDataServiceAppsResponseBodyDataAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceAppsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataServiceAppsResponseBodyData) GetAppList() []*ListDataServiceAppsResponseBodyDataAppList {
	return s.AppList
}

func (s *ListDataServiceAppsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataServiceAppsResponseBodyData) SetAppList(v []*ListDataServiceAppsResponseBodyDataAppList) *ListDataServiceAppsResponseBodyData {
	s.AppList = v
	return s
}

func (s *ListDataServiceAppsResponseBodyData) SetTotalCount(v int64) *ListDataServiceAppsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceAppsResponseBodyData) Validate() error {
	if s.AppList != nil {
		for _, item := range s.AppList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServiceAppsResponseBodyDataAppList struct {
	// example:
	//
	// 默认分组
	AppGroup *string `json:"AppGroup,omitempty" xml:"AppGroup,omitempty"`
	// example:
	//
	// 12345
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 默认应用
	AppName   *string                                                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	IsMember  *bool                                                  `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	OwnerList []*ListDataServiceAppsResponseBodyDataAppListOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
}

func (s ListDataServiceAppsResponseBodyDataAppList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAppsResponseBodyDataAppList) GoString() string {
	return s.String()
}

func (s *ListDataServiceAppsResponseBodyDataAppList) GetAppGroup() *string {
	return s.AppGroup
}

func (s *ListDataServiceAppsResponseBodyDataAppList) GetAppId() *int32 {
	return s.AppId
}

func (s *ListDataServiceAppsResponseBodyDataAppList) GetAppName() *string {
	return s.AppName
}

func (s *ListDataServiceAppsResponseBodyDataAppList) GetIsMember() *bool {
	return s.IsMember
}

func (s *ListDataServiceAppsResponseBodyDataAppList) GetOwnerList() []*ListDataServiceAppsResponseBodyDataAppListOwnerList {
	return s.OwnerList
}

func (s *ListDataServiceAppsResponseBodyDataAppList) SetAppGroup(v string) *ListDataServiceAppsResponseBodyDataAppList {
	s.AppGroup = &v
	return s
}

func (s *ListDataServiceAppsResponseBodyDataAppList) SetAppId(v int32) *ListDataServiceAppsResponseBodyDataAppList {
	s.AppId = &v
	return s
}

func (s *ListDataServiceAppsResponseBodyDataAppList) SetAppName(v string) *ListDataServiceAppsResponseBodyDataAppList {
	s.AppName = &v
	return s
}

func (s *ListDataServiceAppsResponseBodyDataAppList) SetIsMember(v bool) *ListDataServiceAppsResponseBodyDataAppList {
	s.IsMember = &v
	return s
}

func (s *ListDataServiceAppsResponseBodyDataAppList) SetOwnerList(v []*ListDataServiceAppsResponseBodyDataAppListOwnerList) *ListDataServiceAppsResponseBodyDataAppList {
	s.OwnerList = v
	return s
}

func (s *ListDataServiceAppsResponseBodyDataAppList) Validate() error {
	if s.OwnerList != nil {
		for _, item := range s.OwnerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServiceAppsResponseBodyDataAppListOwnerList struct {
	// example:
	//
	// 12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDataServiceAppsResponseBodyDataAppListOwnerList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAppsResponseBodyDataAppListOwnerList) GoString() string {
	return s.String()
}

func (s *ListDataServiceAppsResponseBodyDataAppListOwnerList) GetId() *string {
	return s.Id
}

func (s *ListDataServiceAppsResponseBodyDataAppListOwnerList) GetName() *string {
	return s.Name
}

func (s *ListDataServiceAppsResponseBodyDataAppListOwnerList) SetId(v string) *ListDataServiceAppsResponseBodyDataAppListOwnerList {
	s.Id = &v
	return s
}

func (s *ListDataServiceAppsResponseBodyDataAppListOwnerList) SetName(v string) *ListDataServiceAppsResponseBodyDataAppListOwnerList {
	s.Name = &v
	return s
}

func (s *ListDataServiceAppsResponseBodyDataAppListOwnerList) Validate() error {
	return dara.Validate(s)
}
