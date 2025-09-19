// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpStatusCode(v int32) *GetModuleConfigResponseBody
	GetHttpStatusCode() *int32
	SetModuleConfigList(v []*GetModuleConfigResponseBodyModuleConfigList) *GetModuleConfigResponseBody
	GetModuleConfigList() []*GetModuleConfigResponseBodyModuleConfigList
	SetPageInfo(v *GetModuleConfigResponseBodyPageInfo) *GetModuleConfigResponseBody
	GetPageInfo() *GetModuleConfigResponseBodyPageInfo
	SetRequestId(v string) *GetModuleConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModuleConfigResponseBody
	GetSuccess() *bool
}

type GetModuleConfigResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// An array that consists of the configurations of the module.
	ModuleConfigList []*GetModuleConfigResponseBodyModuleConfigList `json:"ModuleConfigList,omitempty" xml:"ModuleConfigList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *GetModuleConfigResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 291B49F9-1685-4005-9D34-606B6F78****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModuleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetModuleConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetModuleConfigResponseBody) GetModuleConfigList() []*GetModuleConfigResponseBodyModuleConfigList {
	return s.ModuleConfigList
}

func (s *GetModuleConfigResponseBody) GetPageInfo() *GetModuleConfigResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetModuleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModuleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModuleConfigResponseBody) SetHttpStatusCode(v int32) *GetModuleConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModuleConfigResponseBody) SetModuleConfigList(v []*GetModuleConfigResponseBodyModuleConfigList) *GetModuleConfigResponseBody {
	s.ModuleConfigList = v
	return s
}

func (s *GetModuleConfigResponseBody) SetPageInfo(v *GetModuleConfigResponseBodyPageInfo) *GetModuleConfigResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetModuleConfigResponseBody) SetRequestId(v string) *GetModuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModuleConfigResponseBody) SetSuccess(v bool) *GetModuleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetModuleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetModuleConfigResponseBodyModuleConfigList struct {
	// The name of the configuration.
	//
	// example:
	//
	// timescan
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// An array that consists of the configuration items.
	Items []*GetModuleConfigResponseBodyModuleConfigListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The name of the module.
	//
	// example:
	//
	// alihids
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetModuleConfigResponseBodyModuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigResponseBodyModuleConfigList) GoString() string {
	return s.String()
}

func (s *GetModuleConfigResponseBodyModuleConfigList) GetConfigName() *string {
	return s.ConfigName
}

func (s *GetModuleConfigResponseBodyModuleConfigList) GetItems() []*GetModuleConfigResponseBodyModuleConfigListItems {
	return s.Items
}

func (s *GetModuleConfigResponseBodyModuleConfigList) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetModuleConfigResponseBodyModuleConfigList) SetConfigName(v string) *GetModuleConfigResponseBodyModuleConfigList {
	s.ConfigName = &v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigList) SetItems(v []*GetModuleConfigResponseBodyModuleConfigListItems) *GetModuleConfigResponseBodyModuleConfigList {
	s.Items = v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigList) SetModuleName(v string) *GetModuleConfigResponseBodyModuleConfigList {
	s.ModuleName = &v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigList) Validate() error {
	return dara.Validate(s)
}

type GetModuleConfigResponseBodyModuleConfigListItems struct {
	// The ID of the server group to which the server belongs.
	//
	// example:
	//
	// 5562414
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// i-uf6435dn4t59b9av****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the server.
	//
	// example:
	//
	// inStanceName****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 2.2.X.X
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

func (s GetModuleConfigResponseBodyModuleConfigListItems) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigResponseBodyModuleConfigListItems) GoString() string {
	return s.String()
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) GetGroupId() *int32 {
	return s.GroupId
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) GetIp() *string {
	return s.Ip
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) GetRegion() *string {
	return s.Region
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) GetUuid() *string {
	return s.Uuid
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) SetGroupId(v int32) *GetModuleConfigResponseBodyModuleConfigListItems {
	s.GroupId = &v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) SetInstanceId(v string) *GetModuleConfigResponseBodyModuleConfigListItems {
	s.InstanceId = &v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) SetInstanceName(v string) *GetModuleConfigResponseBodyModuleConfigListItems {
	s.InstanceName = &v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) SetIp(v string) *GetModuleConfigResponseBodyModuleConfigListItems {
	s.Ip = &v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) SetRegion(v string) *GetModuleConfigResponseBodyModuleConfigListItems {
	s.Region = &v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) SetUuid(v string) *GetModuleConfigResponseBodyModuleConfigListItems {
	s.Uuid = &v
	return s
}

func (s *GetModuleConfigResponseBodyModuleConfigListItems) Validate() error {
	return dara.Validate(s)
}

type GetModuleConfigResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 100
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetModuleConfigResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetModuleConfigResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *GetModuleConfigResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetModuleConfigResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetModuleConfigResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetModuleConfigResponseBodyPageInfo) SetCount(v int32) *GetModuleConfigResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *GetModuleConfigResponseBodyPageInfo) SetCurrentPage(v int32) *GetModuleConfigResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetModuleConfigResponseBodyPageInfo) SetPageSize(v int32) *GetModuleConfigResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetModuleConfigResponseBodyPageInfo) SetTotalCount(v int32) *GetModuleConfigResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *GetModuleConfigResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
