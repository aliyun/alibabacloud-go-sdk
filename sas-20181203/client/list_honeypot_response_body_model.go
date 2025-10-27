// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHoneypotResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListHoneypotResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListHoneypotResponseBodyList) *ListHoneypotResponseBody
	GetList() []*ListHoneypotResponseBodyList
	SetMessage(v string) *ListHoneypotResponseBody
	GetMessage() *string
	SetPageInfo(v *ListHoneypotResponseBodyPageInfo) *ListHoneypotResponseBody
	GetPageInfo() *ListHoneypotResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHoneypotResponseBody
	GetSuccess() *bool
}

type ListHoneypotResponseBody struct {
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
	// An array that consists of the information about the honeypots.
	List []*ListHoneypotResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListHoneypotResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C80AFF1F-CC20-502C-A4D4-F5433E529B69
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

func (s ListHoneypotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHoneypotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHoneypotResponseBody) GetList() []*ListHoneypotResponseBodyList {
	return s.List
}

func (s *ListHoneypotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHoneypotResponseBody) GetPageInfo() *ListHoneypotResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHoneypotResponseBody) SetCode(v string) *ListHoneypotResponseBody {
	s.Code = &v
	return s
}

func (s *ListHoneypotResponseBody) SetHttpStatusCode(v int32) *ListHoneypotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHoneypotResponseBody) SetList(v []*ListHoneypotResponseBodyList) *ListHoneypotResponseBody {
	s.List = v
	return s
}

func (s *ListHoneypotResponseBody) SetMessage(v string) *ListHoneypotResponseBody {
	s.Message = &v
	return s
}

func (s *ListHoneypotResponseBody) SetPageInfo(v *ListHoneypotResponseBodyPageInfo) *ListHoneypotResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotResponseBody) SetRequestId(v string) *ListHoneypotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotResponseBody) SetSuccess(v bool) *ListHoneypotResponseBody {
	s.Success = &v
	return s
}

func (s *ListHoneypotResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHoneypotResponseBodyList struct {
	// The name of the management node.
	//
	// example:
	//
	// 0804-pre
	ControlNodeName *string `json:"ControlNodeName,omitempty" xml:"ControlNodeName,omitempty"`
	// The ID of the honeypot.
	//
	// example:
	//
	// 76c2a1c72ef259777d96d55a7834e5f5d98f85666c49f76ad9caa447d8b7****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The display name of the honeypot image.
	//
	// example:
	//
	// MongoDB
	HoneypotImageDisplayName *string `json:"HoneypotImageDisplayName,omitempty" xml:"HoneypotImageDisplayName,omitempty"`
	// The ID of the honeypot image.
	//
	// example:
	//
	// sha256:eca5ced3757e46c24701e9ced4e652f2d730262d5685a4e001da22c4fb418fd4
	HoneypotImageId *string `json:"HoneypotImageId,omitempty" xml:"HoneypotImageId,omitempty"`
	// The name of the honeypot image.
	//
	// example:
	//
	// tcp_proxy
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The name of the honeypot.
	//
	// example:
	//
	// mx-rouyi
	HoneypotName *string `json:"HoneypotName,omitempty" xml:"HoneypotName,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// c94eff5b-ea48-4805-8b7f-e04d3509b117
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the custom configuration for the honeypot.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
	// An array that consists of the status information about the honeypot.
	State []*string `json:"State,omitempty" xml:"State,omitempty" type:"Repeated"`
}

func (s ListHoneypotResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListHoneypotResponseBodyList) GetControlNodeName() *string {
	return s.ControlNodeName
}

func (s *ListHoneypotResponseBodyList) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *ListHoneypotResponseBodyList) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *ListHoneypotResponseBodyList) GetHoneypotImageId() *string {
	return s.HoneypotImageId
}

func (s *ListHoneypotResponseBodyList) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *ListHoneypotResponseBodyList) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *ListHoneypotResponseBodyList) GetNodeId() *string {
	return s.NodeId
}

func (s *ListHoneypotResponseBodyList) GetPresetId() *string {
	return s.PresetId
}

func (s *ListHoneypotResponseBodyList) GetState() []*string {
	return s.State
}

func (s *ListHoneypotResponseBodyList) SetControlNodeName(v string) *ListHoneypotResponseBodyList {
	s.ControlNodeName = &v
	return s
}

func (s *ListHoneypotResponseBodyList) SetHoneypotId(v string) *ListHoneypotResponseBodyList {
	s.HoneypotId = &v
	return s
}

func (s *ListHoneypotResponseBodyList) SetHoneypotImageDisplayName(v string) *ListHoneypotResponseBodyList {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *ListHoneypotResponseBodyList) SetHoneypotImageId(v string) *ListHoneypotResponseBodyList {
	s.HoneypotImageId = &v
	return s
}

func (s *ListHoneypotResponseBodyList) SetHoneypotImageName(v string) *ListHoneypotResponseBodyList {
	s.HoneypotImageName = &v
	return s
}

func (s *ListHoneypotResponseBodyList) SetHoneypotName(v string) *ListHoneypotResponseBodyList {
	s.HoneypotName = &v
	return s
}

func (s *ListHoneypotResponseBodyList) SetNodeId(v string) *ListHoneypotResponseBodyList {
	s.NodeId = &v
	return s
}

func (s *ListHoneypotResponseBodyList) SetPresetId(v string) *ListHoneypotResponseBodyList {
	s.PresetId = &v
	return s
}

func (s *ListHoneypotResponseBodyList) SetState(v []*string) *ListHoneypotResponseBodyList {
	s.State = v
	return s
}

func (s *ListHoneypotResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
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
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHoneypotResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotResponseBodyPageInfo) SetCount(v int32) *ListHoneypotResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
