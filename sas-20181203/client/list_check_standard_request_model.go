// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckStandardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *ListCheckStandardRequest
	GetInstanceIds() []*string
	SetInstanceSubTypes(v []*string) *ListCheckStandardRequest
	GetInstanceSubTypes() []*string
	SetInstanceTypes(v []*string) *ListCheckStandardRequest
	GetInstanceTypes() []*string
	SetLang(v string) *ListCheckStandardRequest
	GetLang() *string
	SetTaskSources(v []*string) *ListCheckStandardRequest
	GetTaskSources() []*string
	SetVendors(v []*string) *ListCheckStandardRequest
	GetVendors() []*string
}

type ListCheckStandardRequest struct {
	// The instance IDs of the cloud services to which the check items belong.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The subtypes of cloud services.
	InstanceSubTypes []*string `json:"InstanceSubTypes,omitempty" xml:"InstanceSubTypes,omitempty" type:"Repeated"`
	// The asset types of cloud services.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// List of task sources.
	TaskSources []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
	// The cloud service providers. Valid values:
	//
	// 	- **ALIYUN**: Alibaba Cloud.
	//
	// 	- **TENCENT**: Tencent Cloud.
	//
	// 	- **HUAWEICLOUD**: Huawei Cloud.
	//
	// 	- **MICROSOFT**: Microsoft Azure.
	//
	// 	- **AWS**: Amazon Web Services (AWS).
	Vendors []*string `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s ListCheckStandardRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckStandardRequest) GoString() string {
	return s.String()
}

func (s *ListCheckStandardRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListCheckStandardRequest) GetInstanceSubTypes() []*string {
	return s.InstanceSubTypes
}

func (s *ListCheckStandardRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ListCheckStandardRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckStandardRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *ListCheckStandardRequest) GetVendors() []*string {
	return s.Vendors
}

func (s *ListCheckStandardRequest) SetInstanceIds(v []*string) *ListCheckStandardRequest {
	s.InstanceIds = v
	return s
}

func (s *ListCheckStandardRequest) SetInstanceSubTypes(v []*string) *ListCheckStandardRequest {
	s.InstanceSubTypes = v
	return s
}

func (s *ListCheckStandardRequest) SetInstanceTypes(v []*string) *ListCheckStandardRequest {
	s.InstanceTypes = v
	return s
}

func (s *ListCheckStandardRequest) SetLang(v string) *ListCheckStandardRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckStandardRequest) SetTaskSources(v []*string) *ListCheckStandardRequest {
	s.TaskSources = v
	return s
}

func (s *ListCheckStandardRequest) SetVendors(v []*string) *ListCheckStandardRequest {
	s.Vendors = v
	return s
}

func (s *ListCheckStandardRequest) Validate() error {
	return dara.Validate(s)
}
