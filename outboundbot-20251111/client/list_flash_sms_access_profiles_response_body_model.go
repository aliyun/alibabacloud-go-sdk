// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsAccessProfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFlashSmsAccessProfilesResponseBody
	GetCode() *string
	SetData(v *ListFlashSmsAccessProfilesResponseBodyData) *ListFlashSmsAccessProfilesResponseBody
	GetData() *ListFlashSmsAccessProfilesResponseBodyData
	SetHttpStatusCode(v int32) *ListFlashSmsAccessProfilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFlashSmsAccessProfilesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListFlashSmsAccessProfilesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListFlashSmsAccessProfilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFlashSmsAccessProfilesResponseBody
	GetSuccess() *bool
}

type ListFlashSmsAccessProfilesResponseBody struct {
	// 返回码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *ListFlashSmsAccessProfilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	//
	// example:
	//
	// Instance does not exist. Instance=outb001
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误信息中的变量值列表
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// 请求ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否调用成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFlashSmsAccessProfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsAccessProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlashSmsAccessProfilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFlashSmsAccessProfilesResponseBody) GetData() *ListFlashSmsAccessProfilesResponseBodyData {
	return s.Data
}

func (s *ListFlashSmsAccessProfilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFlashSmsAccessProfilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlashSmsAccessProfilesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListFlashSmsAccessProfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlashSmsAccessProfilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFlashSmsAccessProfilesResponseBody) SetCode(v string) *ListFlashSmsAccessProfilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBody) SetData(v *ListFlashSmsAccessProfilesResponseBodyData) *ListFlashSmsAccessProfilesResponseBody {
	s.Data = v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBody) SetHttpStatusCode(v int32) *ListFlashSmsAccessProfilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBody) SetMessage(v string) *ListFlashSmsAccessProfilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBody) SetParams(v []*string) *ListFlashSmsAccessProfilesResponseBody {
	s.Params = v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBody) SetRequestId(v string) *ListFlashSmsAccessProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBody) SetSuccess(v bool) *ListFlashSmsAccessProfilesResponseBody {
	s.Success = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFlashSmsAccessProfilesResponseBodyData struct {
	// 数据列表
	FlashSmsAccessProfiles []*ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles `json:"FlashSmsAccessProfiles,omitempty" xml:"FlashSmsAccessProfiles,omitempty" type:"Repeated"`
	// 页码，从1开始
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页记录数
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 符合条件的记录总数
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlashSmsAccessProfilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsAccessProfilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) GetFlashSmsAccessProfiles() []*ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles {
	return s.FlashSmsAccessProfiles
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) SetFlashSmsAccessProfiles(v []*ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) *ListFlashSmsAccessProfilesResponseBodyData {
	s.FlashSmsAccessProfiles = v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) SetPageNumber(v int32) *ListFlashSmsAccessProfilesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) SetPageSize(v int32) *ListFlashSmsAccessProfilesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) SetTotalCount(v int32) *ListFlashSmsAccessProfilesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyData) Validate() error {
	if s.FlashSmsAccessProfiles != nil {
		for _, item := range s.FlashSmsAccessProfiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles struct {
	// 接入配置
	//
	// example:
	//
	// {"apiId":"100235","apiKey":"3aRsPrTsDG3OPNq5","aesKey":"TQChVEAabhaNp2AB","capAppId":"300012117547"}
	AccessProfile *string `json:"AccessProfile,omitempty" xml:"AccessProfile,omitempty"`
	// 接入配置ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// 创建时间，毫秒级时间戳
	//
	// example:
	//
	// 1735660800000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 供应商ID
	//
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// 供应商名称
	//
	// example:
	//
	// 北京优音通信有限公司
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// 更新时间，毫秒级时间戳
	//
	// example:
	//
	// 1735660800000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) GoString() string {
	return s.String()
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) GetAccessProfile() *string {
	return s.AccessProfile
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) SetAccessProfile(v string) *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles {
	s.AccessProfile = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) SetAccessProfileId(v string) *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles {
	s.AccessProfileId = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) SetCreatedTime(v int64) *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles {
	s.CreatedTime = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) SetProviderId(v string) *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles {
	s.ProviderId = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) SetProviderName(v string) *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles {
	s.ProviderName = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) SetUpdatedTime(v int64) *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles {
	s.UpdatedTime = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponseBodyDataFlashSmsAccessProfiles) Validate() error {
	return dara.Validate(s)
}
