// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRetcodeAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateRetcodeAppRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateRetcodeAppRequest
	GetResourceGroupId() *string
	SetRetcodeAppName(v string) *CreateRetcodeAppRequest
	GetRetcodeAppName() *string
	SetRetcodeAppType(v string) *CreateRetcodeAppRequest
	GetRetcodeAppType() *string
	SetTags(v []*CreateRetcodeAppRequestTags) *CreateRetcodeAppRequest
	GetTags() []*CreateRetcodeAppRequestTags
}

type CreateRetcodeAppRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can obtain the resource group ID in the **Resource Management*	- console.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// SdkTest
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// mini_dd
	RetcodeAppType *string `json:"RetcodeAppType,omitempty" xml:"RetcodeAppType,omitempty"`
	// The tags that you want to add to the task.
	Tags []*CreateRetcodeAppRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateRetcodeAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRetcodeAppRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRetcodeAppRequest) GetRetcodeAppName() *string {
	return s.RetcodeAppName
}

func (s *CreateRetcodeAppRequest) GetRetcodeAppType() *string {
	return s.RetcodeAppType
}

func (s *CreateRetcodeAppRequest) GetTags() []*CreateRetcodeAppRequestTags {
	return s.Tags
}

func (s *CreateRetcodeAppRequest) SetRegionId(v string) *CreateRetcodeAppRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetResourceGroupId(v string) *CreateRetcodeAppRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppName(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppName = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppType(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppType = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetTags(v []*CreateRetcodeAppRequestTags) *CreateRetcodeAppRequest {
	s.Tags = v
	return s
}

func (s *CreateRetcodeAppRequest) Validate() error {
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

type CreateRetcodeAppRequestTags struct {
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

func (s CreateRetcodeAppRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppRequestTags) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateRetcodeAppRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateRetcodeAppRequestTags) SetKey(v string) *CreateRetcodeAppRequestTags {
	s.Key = &v
	return s
}

func (s *CreateRetcodeAppRequestTags) SetValue(v string) *CreateRetcodeAppRequestTags {
	s.Value = &v
	return s
}

func (s *CreateRetcodeAppRequestTags) Validate() error {
	return dara.Validate(s)
}
