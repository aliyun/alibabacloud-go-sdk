// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsByGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePluginsByGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginsByGroupResponseBody
	GetPageSize() *int32
	SetPlugins(v *DescribePluginsByGroupResponseBodyPlugins) *DescribePluginsByGroupResponseBody
	GetPlugins() *DescribePluginsByGroupResponseBodyPlugins
	SetRequestId(v string) *DescribePluginsByGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePluginsByGroupResponseBody
	GetTotalCount() *int32
}

type DescribePluginsByGroupResponseBody struct {
	// Pagination parameter: current page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Pagination parameter: number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Returns information about the plugins
	Plugins *DescribePluginsByGroupResponseBodyPlugins `json:"Plugins,omitempty" xml:"Plugins,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 5F5574BA-F22B-563D-841E-C817964F517F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of results returned
	//
	// example:
	//
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePluginsByGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginsByGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginsByGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginsByGroupResponseBody) GetPlugins() *DescribePluginsByGroupResponseBodyPlugins {
	return s.Plugins
}

func (s *DescribePluginsByGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePluginsByGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePluginsByGroupResponseBody) SetPageNumber(v int32) *DescribePluginsByGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsByGroupResponseBody) SetPageSize(v int32) *DescribePluginsByGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsByGroupResponseBody) SetPlugins(v *DescribePluginsByGroupResponseBodyPlugins) *DescribePluginsByGroupResponseBody {
	s.Plugins = v
	return s
}

func (s *DescribePluginsByGroupResponseBody) SetRequestId(v string) *DescribePluginsByGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginsByGroupResponseBody) SetTotalCount(v int32) *DescribePluginsByGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePluginsByGroupResponseBody) Validate() error {
	if s.Plugins != nil {
		if err := s.Plugins.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePluginsByGroupResponseBodyPlugins struct {
	PluginAttribute []*DescribePluginsByGroupResponseBodyPluginsPluginAttribute `json:"PluginAttribute,omitempty" xml:"PluginAttribute,omitempty" type:"Repeated"`
}

func (s DescribePluginsByGroupResponseBodyPlugins) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByGroupResponseBodyPlugins) GoString() string {
	return s.String()
}

func (s *DescribePluginsByGroupResponseBodyPlugins) GetPluginAttribute() []*DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	return s.PluginAttribute
}

func (s *DescribePluginsByGroupResponseBodyPlugins) SetPluginAttribute(v []*DescribePluginsByGroupResponseBodyPluginsPluginAttribute) *DescribePluginsByGroupResponseBodyPlugins {
	s.PluginAttribute = v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPlugins) Validate() error {
	if s.PluginAttribute != nil {
		for _, item := range s.PluginAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePluginsByGroupResponseBodyPluginsPluginAttribute struct {
	// Creation time, in GMT
	//
	// example:
	//
	// 2024-12-20T02:05:57Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Plugin description
	//
	// example:
	//
	// traffic controll
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Last modified time, in GMT
	//
	// example:
	//
	// 2022-03-15T02:30:18Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Plugin definition statement
	//
	// example:
	//
	// {\\"unit\\":\\"MINUTE\\",\\"apiDefault\\":20}
	PluginData *string `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	// Plugin ID
	//
	// example:
	//
	// 5e204eeb4aa94c919a49f471ad3fc716
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// Plugin name
	//
	// example:
	//
	// firstPlugin
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// Plugin type
	//
	// example:
	//
	// trafficControl
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// 插件所在Region
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePluginsByGroupResponseBodyPluginsPluginAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GoString() string {
	return s.String()
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GetPluginData() *string {
	return s.PluginData
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GetPluginId() *string {
	return s.PluginId
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GetPluginName() *string {
	return s.PluginName
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GetPluginType() *string {
	return s.PluginType
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) SetCreatedTime(v string) *DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) SetDescription(v string) *DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	s.Description = &v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) SetModifiedTime(v string) *DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) SetPluginData(v string) *DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	s.PluginData = &v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) SetPluginId(v string) *DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	s.PluginId = &v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) SetPluginName(v string) *DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	s.PluginName = &v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) SetPluginType(v string) *DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	s.PluginType = &v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) SetRegionId(v string) *DescribePluginsByGroupResponseBodyPluginsPluginAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribePluginsByGroupResponseBodyPluginsPluginAttribute) Validate() error {
	return dara.Validate(s)
}
