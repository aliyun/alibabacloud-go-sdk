// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudVendorRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCloudVendorRegionsResponseBody
	GetCode() *string
	SetCount(v int32) *ListCloudVendorRegionsResponseBody
	GetCount() *int32
	SetData(v []*ListCloudVendorRegionsResponseBodyData) *ListCloudVendorRegionsResponseBody
	GetData() []*ListCloudVendorRegionsResponseBodyData
	SetHttpStatusCode(v int32) *ListCloudVendorRegionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCloudVendorRegionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCloudVendorRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCloudVendorRegionsResponseBody
	GetSuccess() *bool
}

type ListCloudVendorRegionsResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The regions that the service provider supports.
	Data []*ListCloudVendorRegionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C87EC6AD-4590-5546-9DF6-B8956579D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudVendorRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudVendorRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudVendorRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCloudVendorRegionsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListCloudVendorRegionsResponseBody) GetData() []*ListCloudVendorRegionsResponseBodyData {
	return s.Data
}

func (s *ListCloudVendorRegionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCloudVendorRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCloudVendorRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudVendorRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCloudVendorRegionsResponseBody) SetCode(v string) *ListCloudVendorRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBody) SetCount(v int32) *ListCloudVendorRegionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBody) SetData(v []*ListCloudVendorRegionsResponseBodyData) *ListCloudVendorRegionsResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudVendorRegionsResponseBody) SetHttpStatusCode(v int32) *ListCloudVendorRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBody) SetMessage(v string) *ListCloudVendorRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBody) SetRequestId(v string) *ListCloudVendorRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBody) SetSuccess(v bool) *ListCloudVendorRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCloudVendorRegionsResponseBodyData struct {
	// The area to which the region belongs. The valid values vary based on the value of the Vendor parameter.
	//
	// 	- Valid values if **Vendor*	- is set to Tencent:
	//
	// 	- **cn**: China
	//
	// 	- **southeast_asia**: Southeast Asia Pacific
	//
	// 	- **northeast_asia**: Northeast Asia Pacific
	//
	// 	- **southern_asia**: South Asia Pacific
	//
	// 	- **north_America**: North America
	//
	// 	- **south_America**: South America
	//
	// 	- **western_America**: Western United States
	//
	// 	- **eastern_America**: Eastern United States
	//
	// 	- **european**: Europe
	//
	// 	- Valid values if **Vendor*	- is set to HUAWEICLOUD:
	//
	// 	- **cn**: China
	//
	// 	- **africa**: Africa
	//
	// 	- **latin_america**: Latin America
	//
	// 	- **asia**: Asia Pacific
	//
	// 	- Valid values if **Vendor*	- is set to Azure:
	//
	// 	- **middle_east**: Middle East
	//
	// 	- **south_america**: South America
	//
	// 	- **canada**: Canada
	//
	// 	- **asia-pacific**: Asia Pacific
	//
	// 	- **europe**: Europe
	//
	// 	- **africa**: Africa
	//
	// 	- **us**: United States
	//
	// 	- **other**: other regions
	//
	// 	- Valid values if **Vendor*	- is set to AWS:
	//
	// 	- **cn**: China
	//
	// 	- **us**: United States
	//
	// 	- **eu**: Europe
	//
	// 	- **asia**: Asia Pacific
	//
	// 	- **south_america**: South America
	//
	// 	- **me**: Middle East
	//
	// 	- **ca**: Canada
	//
	// 	- **af**: Africa
	//
	// example:
	//
	// cn
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// Indicates whether the region is configured as a synchronization region on another site. Valid values:
	//
	// 	- **0**: The region is not configured as a synchronization region on another site.
	//
	// 	- **1**: The region is configured as a synchronization region on another site.
	//
	// example:
	//
	// 1
	Disable *int32 `json:"Disable,omitempty" xml:"Disable,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether the region is configured as a synchronization region on this site. Valid values:
	//
	// 	- **0**: The region is not configured as a synchronization region on this site.
	//
	// 	- **1**: The region is configured as a synchronization region on this site.
	//
	// example:
	//
	// 0
	Selected *int32 `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s ListCloudVendorRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCloudVendorRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudVendorRegionsResponseBodyData) GetArea() *string {
	return s.Area
}

func (s *ListCloudVendorRegionsResponseBodyData) GetDisable() *int32 {
	return s.Disable
}

func (s *ListCloudVendorRegionsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCloudVendorRegionsResponseBodyData) GetSelected() *int32 {
	return s.Selected
}

func (s *ListCloudVendorRegionsResponseBodyData) SetArea(v string) *ListCloudVendorRegionsResponseBodyData {
	s.Area = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBodyData) SetDisable(v int32) *ListCloudVendorRegionsResponseBodyData {
	s.Disable = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBodyData) SetRegionId(v string) *ListCloudVendorRegionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBodyData) SetSelected(v int32) *ListCloudVendorRegionsResponseBodyData {
	s.Selected = &v
	return s
}

func (s *ListCloudVendorRegionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
