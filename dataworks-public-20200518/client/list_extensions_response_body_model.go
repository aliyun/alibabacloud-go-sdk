// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListExtensionsResponseBodyPagingInfo) *ListExtensionsResponseBody
	GetPagingInfo() *ListExtensionsResponseBodyPagingInfo
	SetRequestId(v string) *ListExtensionsResponseBody
	GetRequestId() *string
}

type ListExtensionsResponseBody struct {
	// The pagination information.
	PagingInfo *ListExtensionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponseBody) GetPagingInfo() *ListExtensionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExtensionsResponseBody) SetPagingInfo(v *ListExtensionsResponseBodyPagingInfo) *ListExtensionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListExtensionsResponseBody) SetRequestId(v string) *ListExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExtensionsResponseBodyPagingInfo struct {
	// The list of extensions.
	Extensions []*ListExtensionsResponseBodyPagingInfoExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExtensionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListExtensionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponseBodyPagingInfo) GetExtensions() []*ListExtensionsResponseBodyPagingInfoExtensions {
	return s.Extensions
}

func (s *ListExtensionsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExtensionsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExtensionsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListExtensionsResponseBodyPagingInfo) SetExtensions(v []*ListExtensionsResponseBodyPagingInfoExtensions) *ListExtensionsResponseBodyPagingInfo {
	s.Extensions = v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListExtensionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfo) SetPageSize(v int32) *ListExtensionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListExtensionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListExtensionsResponseBodyPagingInfoExtensions struct {
	// The list of extension point events.
	BindEventList []*ListExtensionsResponseBodyPagingInfoExtensionsBindEventList `json:"BindEventList,omitempty" xml:"BindEventList,omitempty" type:"Repeated"`
	// The unique code of the extension.
	//
	// example:
	//
	// Extension Code
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The description of the extension.
	//
	// example:
	//
	// This is a description
	ExtensionDesc *string `json:"ExtensionDesc,omitempty" xml:"ExtensionDesc,omitempty"`
	// The name of the extension.
	//
	// example:
	//
	// Extension name
	ExtensionName *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 2003****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The state of the extension. Valid values: 0: Testing 1: Publishing 3: Disabled 4: Processing 5: Approved 6: Approve Failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListExtensionsResponseBodyPagingInfoExtensions) String() string {
	return dara.Prettify(s)
}

func (s ListExtensionsResponseBodyPagingInfoExtensions) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) GetBindEventList() []*ListExtensionsResponseBodyPagingInfoExtensionsBindEventList {
	return s.BindEventList
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) GetExtensionCode() *string {
	return s.ExtensionCode
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) GetExtensionDesc() *string {
	return s.ExtensionDesc
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) GetExtensionName() *string {
	return s.ExtensionName
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) GetOwner() *string {
	return s.Owner
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) GetStatus() *int32 {
	return s.Status
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) SetBindEventList(v []*ListExtensionsResponseBodyPagingInfoExtensionsBindEventList) *ListExtensionsResponseBodyPagingInfoExtensions {
	s.BindEventList = v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) SetExtensionCode(v string) *ListExtensionsResponseBodyPagingInfoExtensions {
	s.ExtensionCode = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) SetExtensionDesc(v string) *ListExtensionsResponseBodyPagingInfoExtensions {
	s.ExtensionDesc = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) SetExtensionName(v string) *ListExtensionsResponseBodyPagingInfoExtensions {
	s.ExtensionName = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) SetOwner(v string) *ListExtensionsResponseBodyPagingInfoExtensions {
	s.Owner = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) SetStatus(v int32) *ListExtensionsResponseBodyPagingInfoExtensions {
	s.Status = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfoExtensions) Validate() error {
	return dara.Validate(s)
}

type ListExtensionsResponseBodyPagingInfoExtensionsBindEventList struct {
	// The code of the event.
	//
	// example:
	//
	// commit-file
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// File submission pre-event
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
}

func (s ListExtensionsResponseBodyPagingInfoExtensionsBindEventList) String() string {
	return dara.Prettify(s)
}

func (s ListExtensionsResponseBodyPagingInfoExtensionsBindEventList) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponseBodyPagingInfoExtensionsBindEventList) GetEventCode() *string {
	return s.EventCode
}

func (s *ListExtensionsResponseBodyPagingInfoExtensionsBindEventList) GetEventName() *string {
	return s.EventName
}

func (s *ListExtensionsResponseBodyPagingInfoExtensionsBindEventList) SetEventCode(v string) *ListExtensionsResponseBodyPagingInfoExtensionsBindEventList {
	s.EventCode = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfoExtensionsBindEventList) SetEventName(v string) *ListExtensionsResponseBodyPagingInfoExtensionsBindEventList {
	s.EventName = &v
	return s
}

func (s *ListExtensionsResponseBodyPagingInfoExtensionsBindEventList) Validate() error {
	return dara.Validate(s)
}
