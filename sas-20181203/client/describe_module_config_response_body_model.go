// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModuleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeModuleConfigResponseBody
	GetCount() *int32
	SetHttpStatusCode(v int32) *DescribeModuleConfigResponseBody
	GetHttpStatusCode() *int32
	SetModuleConfigList(v []*DescribeModuleConfigResponseBodyModuleConfigList) *DescribeModuleConfigResponseBody
	GetModuleConfigList() []*DescribeModuleConfigResponseBodyModuleConfigList
	SetRequestId(v string) *DescribeModuleConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeModuleConfigResponseBody
	GetSuccess() *bool
}

type DescribeModuleConfigResponseBody struct {
	// The number of configurations for the module.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// An array that consists of the configurations of the module.
	ModuleConfigList []*DescribeModuleConfigResponseBodyModuleConfigList `json:"ModuleConfigList,omitempty" xml:"ModuleConfigList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeModuleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeModuleConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeModuleConfigResponseBody) GetModuleConfigList() []*DescribeModuleConfigResponseBodyModuleConfigList {
	return s.ModuleConfigList
}

func (s *DescribeModuleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModuleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeModuleConfigResponseBody) SetCount(v int32) *DescribeModuleConfigResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeModuleConfigResponseBody) SetHttpStatusCode(v int32) *DescribeModuleConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeModuleConfigResponseBody) SetModuleConfigList(v []*DescribeModuleConfigResponseBodyModuleConfigList) *DescribeModuleConfigResponseBody {
	s.ModuleConfigList = v
	return s
}

func (s *DescribeModuleConfigResponseBody) SetRequestId(v string) *DescribeModuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModuleConfigResponseBody) SetSuccess(v bool) *DescribeModuleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeModuleConfigResponseBody) Validate() error {
	if s.ModuleConfigList != nil {
		for _, item := range s.ModuleConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModuleConfigResponseBodyModuleConfigList struct {
	// The name of the configuration.
	//
	// example:
	//
	// timescan
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// An array that consists of the configuration items.
	Items []*DescribeModuleConfigResponseBodyModuleConfigListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The name of the module.
	//
	// example:
	//
	// alihids
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DescribeModuleConfigResponseBodyModuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleConfigResponseBodyModuleConfigList) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) GetConfigName() *string {
	return s.ConfigName
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) GetItems() []*DescribeModuleConfigResponseBodyModuleConfigListItems {
	return s.Items
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) SetConfigName(v string) *DescribeModuleConfigResponseBodyModuleConfigList {
	s.ConfigName = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) SetItems(v []*DescribeModuleConfigResponseBodyModuleConfigListItems) *DescribeModuleConfigResponseBodyModuleConfigList {
	s.Items = v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) SetModuleName(v string) *DescribeModuleConfigResponseBodyModuleConfigList {
	s.ModuleName = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModuleConfigResponseBodyModuleConfigListItems struct {
	// The ID of the server group.
	//
	// example:
	//
	// 173
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-uf6435dn4t59b9av****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// inStanceName****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 31.13.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The region in which the server resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// a47e3713-ed22-4015-93a3-d88ebe6****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeModuleConfigResponseBodyModuleConfigListItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleConfigResponseBodyModuleConfigListItems) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) GetIp() *string {
	return s.Ip
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetGroupId(v int32) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.GroupId = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetInstanceId(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetInstanceName(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetIp(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.Ip = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetRegion(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.Region = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetUuid(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.Uuid = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) Validate() error {
	return dara.Validate(s)
}
