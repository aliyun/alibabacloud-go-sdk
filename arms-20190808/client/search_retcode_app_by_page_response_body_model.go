// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRetcodeAppByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchRetcodeAppByPageResponseBodyPageBean) *SearchRetcodeAppByPageResponseBody
	GetPageBean() *SearchRetcodeAppByPageResponseBodyPageBean
	SetRequestId(v string) *SearchRetcodeAppByPageResponseBody
	GetRequestId() *string
}

type SearchRetcodeAppByPageResponseBody struct {
	// The returned page information.
	PageBean *SearchRetcodeAppByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 626037F5-FDEB-45B0-804C-B3C92797A64E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBody) GetPageBean() *SearchRetcodeAppByPageResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchRetcodeAppByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchRetcodeAppByPageResponseBody) SetPageBean(v *SearchRetcodeAppByPageResponseBodyPageBean) *SearchRetcodeAppByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBody) SetRequestId(v string) *SearchRetcodeAppByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchRetcodeAppByPageResponseBodyPageBean struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The browser monitoring tasks that are returned.
	RetcodeApps []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) GetRetcodeApps() []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	return s.RetcodeApps
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetPageSize(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetRetcodeApps(v []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.RetcodeApps = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) Validate() error {
	if s.RetcodeApps != nil {
		for _, item := range s.RetcodeApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps struct {
	// The ID of the application. The parameter is an auto-increment parameter.
	//
	// example:
	//
	// 16064
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// a3
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 1545363321000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The alias of the application.
	//
	// example:
	//
	// c1
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The process identifier (PID) of the application.
	//
	// example:
	//
	// eb4zdose6v@9781be0f44d****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- `web`: web application
	//
	// 	- `weex`: Weex mobile app
	//
	// 	- `mini_dd`: DingTalk mini program
	//
	// 	- `mini_alipay`: Alipay mini program
	//
	// 	- `mini_wx`: WeChat mini program
	//
	// 	- `mini_common`: mini program on other platforms
	//
	// example:
	//
	// web
	RetcodeAppType *string `json:"RetcodeAppType,omitempty" xml:"RetcodeAppType,omitempty"`
	// The tag.
	Tags []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the monitoring task. Valid values:
	//
	// 	- `TRACE`: Application Monitoring
	//
	// 	- `RETCODE`: Browser Monitoring
	//
	// example:
	//
	// RETCODE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the task was updated.
	//
	// example:
	//
	// 1545363321000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 12341234
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetAppId() *int64 {
	return s.AppId
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetAppName() *string {
	return s.AppName
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetNickName() *string {
	return s.NickName
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetPid() *string {
	return s.Pid
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetRetcodeAppType() *string {
	return s.RetcodeAppType
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetTags() []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags {
	return s.Tags
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetType() *string {
	return s.Type
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetUserId() *string {
	return s.UserId
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppId(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppName(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppName = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetCreateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.CreateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetNickName(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.NickName = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetPid(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Pid = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetRegionId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.RegionId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetResourceGroupId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetRetcodeAppType(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.RetcodeAppType = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetTags(v []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Tags = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetType(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Type = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUpdateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUserId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UserId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags) GetKey() *string {
	return s.Key
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags) GetValue() *string {
	return s.Value
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags) SetKey(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags {
	s.Key = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags) SetValue(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags {
	s.Value = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeAppsTags) Validate() error {
	return dara.Validate(s)
}
