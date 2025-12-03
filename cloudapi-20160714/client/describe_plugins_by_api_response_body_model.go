// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsByApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePluginsByApiResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginsByApiResponseBody
	GetPageSize() *int32
	SetPlugins(v *DescribePluginsByApiResponseBodyPlugins) *DescribePluginsByApiResponseBody
	GetPlugins() *DescribePluginsByApiResponseBodyPlugins
	SetRequestId(v string) *DescribePluginsByApiResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePluginsByApiResponseBody
	GetTotalCount() *int32
}

type DescribePluginsByApiResponseBody struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The plug-in information. The information is an array that consists of PluginAttribute data.
	Plugins *DescribePluginsByApiResponseBodyPlugins `json:"Plugins,omitempty" xml:"Plugins,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 46373DC4-19F1-4DC8-8C31-1107289BB5E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePluginsByApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginsByApiResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginsByApiResponseBody) GetPlugins() *DescribePluginsByApiResponseBodyPlugins {
	return s.Plugins
}

func (s *DescribePluginsByApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePluginsByApiResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePluginsByApiResponseBody) SetPageNumber(v int32) *DescribePluginsByApiResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsByApiResponseBody) SetPageSize(v int32) *DescribePluginsByApiResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsByApiResponseBody) SetPlugins(v *DescribePluginsByApiResponseBodyPlugins) *DescribePluginsByApiResponseBody {
	s.Plugins = v
	return s
}

func (s *DescribePluginsByApiResponseBody) SetRequestId(v string) *DescribePluginsByApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginsByApiResponseBody) SetTotalCount(v int32) *DescribePluginsByApiResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePluginsByApiResponseBody) Validate() error {
	if s.Plugins != nil {
		if err := s.Plugins.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePluginsByApiResponseBodyPlugins struct {
	PluginAttribute []*DescribePluginsByApiResponseBodyPluginsPluginAttribute `json:"PluginAttribute,omitempty" xml:"PluginAttribute,omitempty" type:"Repeated"`
}

func (s DescribePluginsByApiResponseBodyPlugins) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByApiResponseBodyPlugins) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiResponseBodyPlugins) GetPluginAttribute() []*DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	return s.PluginAttribute
}

func (s *DescribePluginsByApiResponseBodyPlugins) SetPluginAttribute(v []*DescribePluginsByApiResponseBodyPluginsPluginAttribute) *DescribePluginsByApiResponseBodyPlugins {
	s.PluginAttribute = v
	return s
}

func (s *DescribePluginsByApiResponseBodyPlugins) Validate() error {
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

type DescribePluginsByApiResponseBodyPluginsPluginAttribute struct {
	// The time when the plug-in was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-11T09:29:58Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The plug-in description.
	//
	// example:
	//
	// Traffic throttling
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the plug-in was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-11T09:29:58Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The definition statement of the plug-in.
	//
	// example:
	//
	// {\\"unit\\":\\"MINUTE\\",\\"apiDefault\\":20}
	PluginData *string `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	// The plug-in ID.
	//
	// example:
	//
	// 9a3f1a5279434f2ba74ccd91c295af9f
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The plug-in name.
	//
	// example:
	//
	// firstPlugin
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The plug-in type.
	//
	// example:
	//
	// trafficControl
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The region where the plug-in resides.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePluginsByApiResponseBodyPluginsPluginAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByApiResponseBodyPluginsPluginAttribute) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) GetPluginData() *string {
	return s.PluginData
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) GetPluginId() *string {
	return s.PluginId
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) GetPluginName() *string {
	return s.PluginName
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) GetPluginType() *string {
	return s.PluginType
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetCreatedTime(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetDescription(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.Description = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetModifiedTime(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetPluginData(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.PluginData = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetPluginId(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.PluginId = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetPluginName(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.PluginName = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetPluginType(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.PluginType = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetRegionId(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) Validate() error {
	return dara.Validate(s)
}
