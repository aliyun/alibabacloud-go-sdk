// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPluginsInstancesResponseBody
	GetCode() *string
	SetData(v []*ListPluginsInstancesResponseBodyData) *ListPluginsInstancesResponseBody
	GetData() []*ListPluginsInstancesResponseBodyData
	SetMessage(v string) *ListPluginsInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPluginsInstancesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListPluginsInstancesResponseBody
	GetTotal() *int64
}

type ListPluginsInstancesResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListPluginsInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 42
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPluginsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginsInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPluginsInstancesResponseBody) GetData() []*ListPluginsInstancesResponseBodyData {
	return s.Data
}

func (s *ListPluginsInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPluginsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginsInstancesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListPluginsInstancesResponseBody) SetCode(v string) *ListPluginsInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPluginsInstancesResponseBody) SetData(v []*ListPluginsInstancesResponseBodyData) *ListPluginsInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListPluginsInstancesResponseBody) SetMessage(v string) *ListPluginsInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPluginsInstancesResponseBody) SetRequestId(v string) *ListPluginsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginsInstancesResponseBody) SetTotal(v int64) *ListPluginsInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *ListPluginsInstancesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPluginsInstancesResponseBodyData struct {
	// example:
	//
	// i-bp118piqcio9tiwgh84b
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// example:
	//
	// allowed-repos-qmf8w
	InstanceName *string                                            `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	InstanceTag  []*ListPluginsInstancesResponseBodyDataInstanceTag `json:"instance_tag,omitempty" xml:"instance_tag,omitempty" type:"Repeated"`
	// example:
	//
	// Alibaba Cloud Linux  3.2104 LTS 64 bit
	OsName *string `json:"os_name,omitempty" xml:"os_name,omitempty"`
	// example:
	//
	// 1.1.1.1
	PrivateIp *string `json:"private_ip,omitempty" xml:"private_ip,omitempty"`
	// example:
	//
	// 1.1.1.1
	PublicIp *string `json:"public_ip,omitempty" xml:"public_ip,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// example:
	//
	// default resource group
	ResourceGroupName *string `json:"resource_group_name,omitempty" xml:"resource_group_name,omitempty"`
}

func (s ListPluginsInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPluginsInstancesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPluginsInstancesResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListPluginsInstancesResponseBodyData) GetInstanceTag() []*ListPluginsInstancesResponseBodyDataInstanceTag {
	return s.InstanceTag
}

func (s *ListPluginsInstancesResponseBodyData) GetOsName() *string {
	return s.OsName
}

func (s *ListPluginsInstancesResponseBodyData) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListPluginsInstancesResponseBodyData) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListPluginsInstancesResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *ListPluginsInstancesResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPluginsInstancesResponseBodyData) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListPluginsInstancesResponseBodyData) SetInstanceId(v string) *ListPluginsInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) SetInstanceName(v string) *ListPluginsInstancesResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) SetInstanceTag(v []*ListPluginsInstancesResponseBodyDataInstanceTag) *ListPluginsInstancesResponseBodyData {
	s.InstanceTag = v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) SetOsName(v string) *ListPluginsInstancesResponseBodyData {
	s.OsName = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) SetPrivateIp(v string) *ListPluginsInstancesResponseBodyData {
	s.PrivateIp = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) SetPublicIp(v string) *ListPluginsInstancesResponseBodyData {
	s.PublicIp = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) SetRegion(v string) *ListPluginsInstancesResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) SetResourceGroupId(v string) *ListPluginsInstancesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) SetResourceGroupName(v string) *ListPluginsInstancesResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyData) Validate() error {
	if s.InstanceTag != nil {
		for _, item := range s.InstanceTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPluginsInstancesResponseBodyDataInstanceTag struct {
	// example:
	//
	// test_tag_key
	TagKey *string `json:"tag_key,omitempty" xml:"tag_key,omitempty"`
	// example:
	//
	// test_tag_value
	TagValue *string `json:"tag_value,omitempty" xml:"tag_value,omitempty"`
}

func (s ListPluginsInstancesResponseBodyDataInstanceTag) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsInstancesResponseBodyDataInstanceTag) GoString() string {
	return s.String()
}

func (s *ListPluginsInstancesResponseBodyDataInstanceTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListPluginsInstancesResponseBodyDataInstanceTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListPluginsInstancesResponseBodyDataInstanceTag) SetTagKey(v string) *ListPluginsInstancesResponseBodyDataInstanceTag {
	s.TagKey = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyDataInstanceTag) SetTagValue(v string) *ListPluginsInstancesResponseBodyDataInstanceTag {
	s.TagValue = &v
	return s
}

func (s *ListPluginsInstancesResponseBodyDataInstanceTag) Validate() error {
	return dara.Validate(s)
}
