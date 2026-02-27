// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceAppResponseBody
	GetCode() *string
	SetData(v *GetDataServiceAppResponseBodyData) *GetDataServiceAppResponseBody
	GetData() *GetDataServiceAppResponseBodyData
	SetHttpStatusCode(v int32) *GetDataServiceAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceAppResponseBody
	GetSuccess() *bool
}

type GetDataServiceAppResponseBody struct {
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataServiceAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetDataServiceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceAppResponseBody) GetData() *GetDataServiceAppResponseBodyData {
	return s.Data
}

func (s *GetDataServiceAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceAppResponseBody) SetCode(v string) *GetDataServiceAppResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceAppResponseBody) SetData(v *GetDataServiceAppResponseBodyData) *GetDataServiceAppResponseBody {
	s.Data = v
	return s
}

func (s *GetDataServiceAppResponseBody) SetHttpStatusCode(v int32) *GetDataServiceAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceAppResponseBody) SetMessage(v string) *GetDataServiceAppResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceAppResponseBody) SetRequestId(v string) *GetDataServiceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceAppResponseBody) SetSuccess(v bool) *GetDataServiceAppResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceAppResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataServiceAppResponseBodyData struct {
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
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 默认应用
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 默认应用
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// example:
	//
	// true
	IpWhitelist       *string                                       `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	IpWhitelistStatus *bool                                         `json:"IpWhitelistStatus,omitempty" xml:"IpWhitelistStatus,omitempty"`
	OwnerList         []*GetDataServiceAppResponseBodyDataOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// example:
	//
	// 数据分析
	Scenarios *string `json:"Scenarios,omitempty" xml:"Scenarios,omitempty"`
}

func (s GetDataServiceAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppResponseBodyData) GetAppGroup() *string {
	return s.AppGroup
}

func (s *GetDataServiceAppResponseBodyData) GetAppId() *int32 {
	return s.AppId
}

func (s *GetDataServiceAppResponseBodyData) GetAppKey() *string {
	return s.AppKey
}

func (s *GetDataServiceAppResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetDataServiceAppResponseBodyData) GetAppSecret() *string {
	return s.AppSecret
}

func (s *GetDataServiceAppResponseBodyData) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *GetDataServiceAppResponseBodyData) GetIpWhitelistStatus() *bool {
	return s.IpWhitelistStatus
}

func (s *GetDataServiceAppResponseBodyData) GetOwnerList() []*GetDataServiceAppResponseBodyDataOwnerList {
	return s.OwnerList
}

func (s *GetDataServiceAppResponseBodyData) GetScenarios() *string {
	return s.Scenarios
}

func (s *GetDataServiceAppResponseBodyData) SetAppGroup(v string) *GetDataServiceAppResponseBodyData {
	s.AppGroup = &v
	return s
}

func (s *GetDataServiceAppResponseBodyData) SetAppId(v int32) *GetDataServiceAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetDataServiceAppResponseBodyData) SetAppKey(v string) *GetDataServiceAppResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *GetDataServiceAppResponseBodyData) SetAppName(v string) *GetDataServiceAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetDataServiceAppResponseBodyData) SetAppSecret(v string) *GetDataServiceAppResponseBodyData {
	s.AppSecret = &v
	return s
}

func (s *GetDataServiceAppResponseBodyData) SetIpWhitelist(v string) *GetDataServiceAppResponseBodyData {
	s.IpWhitelist = &v
	return s
}

func (s *GetDataServiceAppResponseBodyData) SetIpWhitelistStatus(v bool) *GetDataServiceAppResponseBodyData {
	s.IpWhitelistStatus = &v
	return s
}

func (s *GetDataServiceAppResponseBodyData) SetOwnerList(v []*GetDataServiceAppResponseBodyDataOwnerList) *GetDataServiceAppResponseBodyData {
	s.OwnerList = v
	return s
}

func (s *GetDataServiceAppResponseBodyData) SetScenarios(v string) *GetDataServiceAppResponseBodyData {
	s.Scenarios = &v
	return s
}

func (s *GetDataServiceAppResponseBodyData) Validate() error {
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

type GetDataServiceAppResponseBodyDataOwnerList struct {
	// example:
	//
	// 12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDataServiceAppResponseBodyDataOwnerList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppResponseBodyDataOwnerList) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppResponseBodyDataOwnerList) GetId() *string {
	return s.Id
}

func (s *GetDataServiceAppResponseBodyDataOwnerList) GetName() *string {
	return s.Name
}

func (s *GetDataServiceAppResponseBodyDataOwnerList) SetId(v string) *GetDataServiceAppResponseBodyDataOwnerList {
	s.Id = &v
	return s
}

func (s *GetDataServiceAppResponseBodyDataOwnerList) SetName(v string) *GetDataServiceAppResponseBodyDataOwnerList {
	s.Name = &v
	return s
}

func (s *GetDataServiceAppResponseBodyDataOwnerList) Validate() error {
	return dara.Validate(s)
}
