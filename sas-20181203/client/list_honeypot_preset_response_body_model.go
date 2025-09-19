// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotPresetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHoneypotPresetResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListHoneypotPresetResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListHoneypotPresetResponseBodyList) *ListHoneypotPresetResponseBody
	GetList() []*ListHoneypotPresetResponseBodyList
	SetMessage(v string) *ListHoneypotPresetResponseBody
	GetMessage() *string
	SetPageInfo(v *ListHoneypotPresetResponseBodyPageInfo) *ListHoneypotPresetResponseBody
	GetPageInfo() *ListHoneypotPresetResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotPresetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHoneypotPresetResponseBody
	GetSuccess() *bool
}

type ListHoneypotPresetResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// An array that consists of the honeypot templates.
	List []*ListHoneypotPresetResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListHoneypotPresetResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0C656B33-0D6B-5953-A26A-D766BD75B44A
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

func (s ListHoneypotPresetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotPresetResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotPresetResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHoneypotPresetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHoneypotPresetResponseBody) GetList() []*ListHoneypotPresetResponseBodyList {
	return s.List
}

func (s *ListHoneypotPresetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHoneypotPresetResponseBody) GetPageInfo() *ListHoneypotPresetResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotPresetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotPresetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHoneypotPresetResponseBody) SetCode(v string) *ListHoneypotPresetResponseBody {
	s.Code = &v
	return s
}

func (s *ListHoneypotPresetResponseBody) SetHttpStatusCode(v int32) *ListHoneypotPresetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHoneypotPresetResponseBody) SetList(v []*ListHoneypotPresetResponseBodyList) *ListHoneypotPresetResponseBody {
	s.List = v
	return s
}

func (s *ListHoneypotPresetResponseBody) SetMessage(v string) *ListHoneypotPresetResponseBody {
	s.Message = &v
	return s
}

func (s *ListHoneypotPresetResponseBody) SetPageInfo(v *ListHoneypotPresetResponseBodyPageInfo) *ListHoneypotPresetResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotPresetResponseBody) SetRequestId(v string) *ListHoneypotPresetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotPresetResponseBody) SetSuccess(v bool) *ListHoneypotPresetResponseBody {
	s.Success = &v
	return s
}

func (s *ListHoneypotPresetResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotPresetResponseBodyList struct {
	// The name of the management node.
	//
	// example:
	//
	// node1
	ControlNodeName *string `json:"ControlNodeName,omitempty" xml:"ControlNodeName,omitempty"`
	// The display name of the honeypot image.
	//
	// example:
	//
	// Metabase
	HoneypotImageDisplayName *string `json:"HoneypotImageDisplayName,omitempty" xml:"HoneypotImageDisplayName,omitempty"`
	// The name of the honeypot image.
	//
	// example:
	//
	// metabase
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The ID of the honeypot template.
	//
	// example:
	//
	// 3cc04a47-7229-418c-8101-f10a2887****
	HoneypotPresetId *string `json:"HoneypotPresetId,omitempty" xml:"HoneypotPresetId,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// c94eff5b-ea48-4805-8b7f-e04d3509****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The custom name of the honeypot template.
	//
	// example:
	//
	// WebMin-online
	PresetName *string `json:"PresetName,omitempty" xml:"PresetName,omitempty"`
	// The type of the honeypot template. Valid values:
	//
	// 	- **TEMP**: automatically generated template
	//
	// 	- **CUSTOM**: custom template
	//
	// 	- **DEFAULT**: default template
	//
	// example:
	//
	// CUSTOM
	PresetType *string `json:"PresetType,omitempty" xml:"PresetType,omitempty"`
}

func (s ListHoneypotPresetResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotPresetResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListHoneypotPresetResponseBodyList) GetControlNodeName() *string {
	return s.ControlNodeName
}

func (s *ListHoneypotPresetResponseBodyList) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *ListHoneypotPresetResponseBodyList) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *ListHoneypotPresetResponseBodyList) GetHoneypotPresetId() *string {
	return s.HoneypotPresetId
}

func (s *ListHoneypotPresetResponseBodyList) GetNodeId() *string {
	return s.NodeId
}

func (s *ListHoneypotPresetResponseBodyList) GetPresetName() *string {
	return s.PresetName
}

func (s *ListHoneypotPresetResponseBodyList) GetPresetType() *string {
	return s.PresetType
}

func (s *ListHoneypotPresetResponseBodyList) SetControlNodeName(v string) *ListHoneypotPresetResponseBodyList {
	s.ControlNodeName = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyList) SetHoneypotImageDisplayName(v string) *ListHoneypotPresetResponseBodyList {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyList) SetHoneypotImageName(v string) *ListHoneypotPresetResponseBodyList {
	s.HoneypotImageName = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyList) SetHoneypotPresetId(v string) *ListHoneypotPresetResponseBodyList {
	s.HoneypotPresetId = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyList) SetNodeId(v string) *ListHoneypotPresetResponseBodyList {
	s.NodeId = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyList) SetPresetName(v string) *ListHoneypotPresetResponseBodyList {
	s.PresetName = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyList) SetPresetType(v string) *ListHoneypotPresetResponseBodyList {
	s.PresetType = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotPresetResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHoneypotPresetResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotPresetResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotPresetResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotPresetResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotPresetResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotPresetResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotPresetResponseBodyPageInfo) SetCount(v int32) *ListHoneypotPresetResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotPresetResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotPresetResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotPresetResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotPresetResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
