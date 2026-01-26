// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeAppByPidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRetcodeAppByPidResponseBody
	GetRequestId() *string
	SetRetcodeApp(v *GetRetcodeAppByPidResponseBodyRetcodeApp) *GetRetcodeAppByPidResponseBody
	GetRetcodeApp() *GetRetcodeAppByPidResponseBodyRetcodeApp
}

type GetRetcodeAppByPidResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2983BEF7-4A0D-47A2-94A2-8E9C5E63****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned application data.
	RetcodeApp *GetRetcodeAppByPidResponseBodyRetcodeApp `json:"RetcodeApp,omitempty" xml:"RetcodeApp,omitempty" type:"Struct"`
}

func (s GetRetcodeAppByPidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeAppByPidResponseBody) GoString() string {
	return s.String()
}

func (s *GetRetcodeAppByPidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRetcodeAppByPidResponseBody) GetRetcodeApp() *GetRetcodeAppByPidResponseBodyRetcodeApp {
	return s.RetcodeApp
}

func (s *GetRetcodeAppByPidResponseBody) SetRequestId(v string) *GetRetcodeAppByPidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRetcodeAppByPidResponseBody) SetRetcodeApp(v *GetRetcodeAppByPidResponseBodyRetcodeApp) *GetRetcodeAppByPidResponseBody {
	s.RetcodeApp = v
	return s
}

func (s *GetRetcodeAppByPidResponseBody) Validate() error {
	if s.RetcodeApp != nil {
		if err := s.RetcodeApp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRetcodeAppByPidResponseBodyRetcodeApp struct {
	// The ID of the application. The parameter is an auto-increment parameter.
	//
	// example:
	//
	// 2787XXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application that is monitored by Browser Monitoring.
	//
	// example:
	//
	// testRetcodeAppXXXX
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The process identifier (PID) of the application.
	//
	// example:
	//
	// b590lhguqs@9781be0f44dXXXX
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the resource group. You can obtain the resource group ID in the **Resource Management*	- console.
	//
	// example:
	//
	// rg-acfmxidtzXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the application that is monitored by Browser Monitoring. Valid values:
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
	// The tags that are attached to the instance.
	Tags []*GetRetcodeAppByPidResponseBodyRetcodeAppTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetRetcodeAppByPidResponseBodyRetcodeApp) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeAppByPidResponseBodyRetcodeApp) GoString() string {
	return s.String()
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) GetAppId() *string {
	return s.AppId
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) GetAppName() *string {
	return s.AppName
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) GetPid() *string {
	return s.Pid
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) GetRetcodeAppType() *string {
	return s.RetcodeAppType
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) GetTags() []*GetRetcodeAppByPidResponseBodyRetcodeAppTags {
	return s.Tags
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) SetAppId(v string) *GetRetcodeAppByPidResponseBodyRetcodeApp {
	s.AppId = &v
	return s
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) SetAppName(v string) *GetRetcodeAppByPidResponseBodyRetcodeApp {
	s.AppName = &v
	return s
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) SetPid(v string) *GetRetcodeAppByPidResponseBodyRetcodeApp {
	s.Pid = &v
	return s
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) SetResourceGroupId(v string) *GetRetcodeAppByPidResponseBodyRetcodeApp {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) SetRetcodeAppType(v string) *GetRetcodeAppByPidResponseBodyRetcodeApp {
	s.RetcodeAppType = &v
	return s
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) SetTags(v []*GetRetcodeAppByPidResponseBodyRetcodeAppTags) *GetRetcodeAppByPidResponseBodyRetcodeApp {
	s.Tags = v
	return s
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeApp) Validate() error {
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

type GetRetcodeAppByPidResponseBodyRetcodeAppTags struct {
	// The key of the tag.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRetcodeAppByPidResponseBodyRetcodeAppTags) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeAppByPidResponseBodyRetcodeAppTags) GoString() string {
	return s.String()
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeAppTags) GetKey() *string {
	return s.Key
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeAppTags) GetValue() *string {
	return s.Value
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeAppTags) SetKey(v string) *GetRetcodeAppByPidResponseBodyRetcodeAppTags {
	s.Key = &v
	return s
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeAppTags) SetValue(v string) *GetRetcodeAppByPidResponseBodyRetcodeAppTags {
	s.Value = &v
	return s
}

func (s *GetRetcodeAppByPidResponseBodyRetcodeAppTags) Validate() error {
	return dara.Validate(s)
}
