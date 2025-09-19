// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpaClusterStrategyNewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOpaClusterStrategyNewResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListOpaClusterStrategyNewResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListOpaClusterStrategyNewResponseBodyList) *ListOpaClusterStrategyNewResponseBody
	GetList() []*ListOpaClusterStrategyNewResponseBodyList
	SetMessage(v string) *ListOpaClusterStrategyNewResponseBody
	GetMessage() *string
	SetPageInfo(v *ListOpaClusterStrategyNewResponseBodyPageInfo) *ListOpaClusterStrategyNewResponseBody
	GetPageInfo() *ListOpaClusterStrategyNewResponseBodyPageInfo
	SetRequestId(v string) *ListOpaClusterStrategyNewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOpaClusterStrategyNewResponseBody
	GetSuccess() *bool
}

type ListOpaClusterStrategyNewResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The rules.
	List []*ListOpaClusterStrategyNewResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The message that shows the export task result. The value is fixed as **success**, which indicates that the export task is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListOpaClusterStrategyNewResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7DFD947C-9172-5129-B783-DD14C55191D2
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

func (s ListOpaClusterStrategyNewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOpaClusterStrategyNewResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpaClusterStrategyNewResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOpaClusterStrategyNewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOpaClusterStrategyNewResponseBody) GetList() []*ListOpaClusterStrategyNewResponseBodyList {
	return s.List
}

func (s *ListOpaClusterStrategyNewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOpaClusterStrategyNewResponseBody) GetPageInfo() *ListOpaClusterStrategyNewResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListOpaClusterStrategyNewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOpaClusterStrategyNewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOpaClusterStrategyNewResponseBody) SetCode(v string) *ListOpaClusterStrategyNewResponseBody {
	s.Code = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBody) SetHttpStatusCode(v int32) *ListOpaClusterStrategyNewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBody) SetList(v []*ListOpaClusterStrategyNewResponseBodyList) *ListOpaClusterStrategyNewResponseBody {
	s.List = v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBody) SetMessage(v string) *ListOpaClusterStrategyNewResponseBody {
	s.Message = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBody) SetPageInfo(v *ListOpaClusterStrategyNewResponseBodyPageInfo) *ListOpaClusterStrategyNewResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBody) SetRequestId(v string) *ListOpaClusterStrategyNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBody) SetSuccess(v bool) *ListOpaClusterStrategyNewResponseBody {
	s.Success = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOpaClusterStrategyNewResponseBodyList struct {
	// The action of the rule. Valid values:
	//
	// 	- **1**: trigger alerts
	//
	// 	- **2**: block
	//
	// 	- **3**: allow
	//
	// example:
	//
	// 1
	Action *int32 `json:"Action,omitempty" xml:"Action,omitempty"`
	// The number of clusters on which the rule takes effect.
	//
	// example:
	//
	// 1
	ClusterCount *int32 `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// The clusters on which the rule takes effect.
	ClusterIdList []*string `json:"ClusterIdList,omitempty" xml:"ClusterIdList,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// Config the Event Audit policys
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image names.
	ImageName []*string `json:"ImageName,omitempty" xml:"ImageName,omitempty" type:"Repeated"`
	// The tags that are added to the container.
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
	// Indicates whether the rule supports malicious Internet images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	MaliciousImage *bool `json:"MaliciousImage,omitempty" xml:"MaliciousImage,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 1
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// auto-strategy-vohuiq
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// Indicates whether the rule supports unscanned images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	UnScanedImage *bool `json:"UnScanedImage,omitempty" xml:"UnScanedImage,omitempty"`
}

func (s ListOpaClusterStrategyNewResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListOpaClusterStrategyNewResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetAction() *int32 {
	return s.Action
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetClusterCount() *int32 {
	return s.ClusterCount
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetClusterIdList() []*string {
	return s.ClusterIdList
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetDescription() *string {
	return s.Description
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetImageName() []*string {
	return s.ImageName
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetLabel() []*string {
	return s.Label
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetMaliciousImage() *bool {
	return s.MaliciousImage
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListOpaClusterStrategyNewResponseBodyList) GetUnScanedImage() *bool {
	return s.UnScanedImage
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetAction(v int32) *ListOpaClusterStrategyNewResponseBodyList {
	s.Action = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetClusterCount(v int32) *ListOpaClusterStrategyNewResponseBodyList {
	s.ClusterCount = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetClusterIdList(v []*string) *ListOpaClusterStrategyNewResponseBodyList {
	s.ClusterIdList = v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetDescription(v string) *ListOpaClusterStrategyNewResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetImageName(v []*string) *ListOpaClusterStrategyNewResponseBodyList {
	s.ImageName = v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetLabel(v []*string) *ListOpaClusterStrategyNewResponseBodyList {
	s.Label = v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetMaliciousImage(v bool) *ListOpaClusterStrategyNewResponseBodyList {
	s.MaliciousImage = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetStrategyId(v int64) *ListOpaClusterStrategyNewResponseBodyList {
	s.StrategyId = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetStrategyName(v string) *ListOpaClusterStrategyNewResponseBodyList {
	s.StrategyName = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) SetUnScanedImage(v bool) *ListOpaClusterStrategyNewResponseBodyList {
	s.UnScanedImage = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListOpaClusterStrategyNewResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 45
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpaClusterStrategyNewResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOpaClusterStrategyNewResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) SetCount(v int32) *ListOpaClusterStrategyNewResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) SetCurrentPage(v int32) *ListOpaClusterStrategyNewResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) SetPageSize(v int32) *ListOpaClusterStrategyNewResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) SetTotalCount(v int32) *ListOpaClusterStrategyNewResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
