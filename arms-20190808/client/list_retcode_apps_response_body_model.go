// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRetcodeAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRetcodeAppsResponseBody
	GetRequestId() *string
	SetRetcodeApps(v []*ListRetcodeAppsResponseBodyRetcodeApps) *ListRetcodeAppsResponseBody
	GetRetcodeApps() []*ListRetcodeAppsResponseBodyRetcodeApps
}

type ListRetcodeAppsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 99A663CB-8D7B-4B0D-A006-03C8EE38E7BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of applications monitored by Browser Monitoring.
	RetcodeApps []*ListRetcodeAppsResponseBodyRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
}

func (s ListRetcodeAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRetcodeAppsResponseBody) GetRetcodeApps() []*ListRetcodeAppsResponseBodyRetcodeApps {
	return s.RetcodeApps
}

func (s *ListRetcodeAppsResponseBody) SetRequestId(v string) *ListRetcodeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRetcodeAppsResponseBody) SetRetcodeApps(v []*ListRetcodeAppsResponseBodyRetcodeApps) *ListRetcodeAppsResponseBody {
	s.RetcodeApps = v
	return s
}

func (s *ListRetcodeAppsResponseBody) Validate() error {
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

type ListRetcodeAppsResponseBodyRetcodeApps struct {
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
	// A1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The alias of the application monitored by Browser Monitoring.
	//
	// example:
	//
	// B1
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The process identifier (PID) of the application.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
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
	// The tags of the task.
	Tags []*ListRetcodeAppsResponseBodyRetcodeAppsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetAppId() *int64 {
	return s.AppId
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetAppName() *string {
	return s.AppName
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetNickName() *string {
	return s.NickName
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetPid() *string {
	return s.Pid
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetRetcodeAppType() *string {
	return s.RetcodeAppType
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetTags() []*ListRetcodeAppsResponseBodyRetcodeAppsTags {
	return s.Tags
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppId(v int64) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppId = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppName(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppName = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetNickName(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.NickName = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetPid(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.Pid = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetResourceGroupId(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetRetcodeAppType(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.RetcodeAppType = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetTags(v []*ListRetcodeAppsResponseBodyRetcodeAppsTags) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.Tags = v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) Validate() error {
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

type ListRetcodeAppsResponseBodyRetcodeAppsTags struct {
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

func (s ListRetcodeAppsResponseBodyRetcodeAppsTags) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsResponseBodyRetcodeAppsTags) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBodyRetcodeAppsTags) GetKey() *string {
	return s.Key
}

func (s *ListRetcodeAppsResponseBodyRetcodeAppsTags) GetValue() *string {
	return s.Value
}

func (s *ListRetcodeAppsResponseBodyRetcodeAppsTags) SetKey(v string) *ListRetcodeAppsResponseBodyRetcodeAppsTags {
	s.Key = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeAppsTags) SetValue(v string) *ListRetcodeAppsResponseBodyRetcodeAppsTags {
	s.Value = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeAppsTags) Validate() error {
	return dara.Validate(s)
}
