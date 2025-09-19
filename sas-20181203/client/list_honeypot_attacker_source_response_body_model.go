// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAttackerSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHoneypotAttackerSourceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListHoneypotAttackerSourceResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListHoneypotAttackerSourceResponseBodyList) *ListHoneypotAttackerSourceResponseBody
	GetList() []*ListHoneypotAttackerSourceResponseBodyList
	SetMessage(v string) *ListHoneypotAttackerSourceResponseBody
	GetMessage() *string
	SetPageInfo(v *ListHoneypotAttackerSourceResponseBodyPageInfo) *ListHoneypotAttackerSourceResponseBody
	GetPageInfo() *ListHoneypotAttackerSourceResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotAttackerSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHoneypotAttackerSourceResponseBody
	GetSuccess() *bool
}

type ListHoneypotAttackerSourceResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The source IP addresses of the attack.
	List []*ListHoneypotAttackerSourceResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListHoneypotAttackerSourceResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9915DC4D-B4DA-5140-8138-FD80636*****
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

func (s ListHoneypotAttackerSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHoneypotAttackerSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHoneypotAttackerSourceResponseBody) GetList() []*ListHoneypotAttackerSourceResponseBodyList {
	return s.List
}

func (s *ListHoneypotAttackerSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHoneypotAttackerSourceResponseBody) GetPageInfo() *ListHoneypotAttackerSourceResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotAttackerSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotAttackerSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHoneypotAttackerSourceResponseBody) SetCode(v string) *ListHoneypotAttackerSourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBody) SetHttpStatusCode(v int32) *ListHoneypotAttackerSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBody) SetList(v []*ListHoneypotAttackerSourceResponseBodyList) *ListHoneypotAttackerSourceResponseBody {
	s.List = v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBody) SetMessage(v string) *ListHoneypotAttackerSourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBody) SetPageInfo(v *ListHoneypotAttackerSourceResponseBodyPageInfo) *ListHoneypotAttackerSourceResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBody) SetRequestId(v string) *ListHoneypotAttackerSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBody) SetSuccess(v bool) *ListHoneypotAttackerSourceResponseBody {
	s.Success = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotAttackerSourceResponseBodyList struct {
	// The total number of attack events.
	//
	// example:
	//
	// 30
	EventCount *int32 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The most recent honeypot that was attacked.
	//
	// example:
	//
	// vpc tcp honeypot
	LastTargetHoneypot *string `json:"LastTargetHoneypot,omitempty" xml:"LastTargetHoneypot,omitempty"`
	// The most recent IP address that was attacked.
	//
	// example:
	//
	// 144.23.66.***
	LastTargetIp *string `json:"LastTargetIp,omitempty" xml:"LastTargetIp,omitempty"`
	// The last time when the attack event occurred.
	//
	// example:
	//
	// 1693446913000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **2**: low
	//
	// 	- **3**: medium
	//
	// 	- **4**: high
	//
	// example:
	//
	// 2
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The source IP address of the attack.
	//
	// example:
	//
	// 101.102.61.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
}

func (s ListHoneypotAttackerSourceResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerSourceResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerSourceResponseBodyList) GetEventCount() *int32 {
	return s.EventCount
}

func (s *ListHoneypotAttackerSourceResponseBodyList) GetLastTargetHoneypot() *string {
	return s.LastTargetHoneypot
}

func (s *ListHoneypotAttackerSourceResponseBodyList) GetLastTargetIp() *string {
	return s.LastTargetIp
}

func (s *ListHoneypotAttackerSourceResponseBodyList) GetLastTime() *int64 {
	return s.LastTime
}

func (s *ListHoneypotAttackerSourceResponseBodyList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListHoneypotAttackerSourceResponseBodyList) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListHoneypotAttackerSourceResponseBodyList) SetEventCount(v int32) *ListHoneypotAttackerSourceResponseBodyList {
	s.EventCount = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyList) SetLastTargetHoneypot(v string) *ListHoneypotAttackerSourceResponseBodyList {
	s.LastTargetHoneypot = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyList) SetLastTargetIp(v string) *ListHoneypotAttackerSourceResponseBodyList {
	s.LastTargetIp = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyList) SetLastTime(v int64) *ListHoneypotAttackerSourceResponseBodyList {
	s.LastTime = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyList) SetRiskLevel(v string) *ListHoneypotAttackerSourceResponseBodyList {
	s.RiskLevel = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyList) SetSrcIp(v string) *ListHoneypotAttackerSourceResponseBodyList {
	s.SrcIp = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotAttackerSourceResponseBodyPageInfo struct {
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
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHoneypotAttackerSourceResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerSourceResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) SetCount(v int32) *ListHoneypotAttackerSourceResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotAttackerSourceResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotAttackerSourceResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotAttackerSourceResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
