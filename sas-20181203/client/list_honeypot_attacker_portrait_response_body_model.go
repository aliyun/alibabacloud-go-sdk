// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAttackerPortraitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHoneypotAttackerPortraitResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListHoneypotAttackerPortraitResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListHoneypotAttackerPortraitResponseBodyList) *ListHoneypotAttackerPortraitResponseBody
	GetList() []*ListHoneypotAttackerPortraitResponseBodyList
	SetMessage(v string) *ListHoneypotAttackerPortraitResponseBody
	GetMessage() *string
	SetPageInfo(v *ListHoneypotAttackerPortraitResponseBodyPageInfo) *ListHoneypotAttackerPortraitResponseBody
	GetPageInfo() *ListHoneypotAttackerPortraitResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotAttackerPortraitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHoneypotAttackerPortraitResponseBody
	GetSuccess() *bool
}

type ListHoneypotAttackerPortraitResponseBody struct {
	// The status code that is returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
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
	// The details of the attacker profile.
	List []*ListHoneypotAttackerPortraitResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListHoneypotAttackerPortraitResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8A5A2DA6-67EA-5968-960F-6B20FD0C*****
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

func (s ListHoneypotAttackerPortraitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerPortraitResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerPortraitResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHoneypotAttackerPortraitResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHoneypotAttackerPortraitResponseBody) GetList() []*ListHoneypotAttackerPortraitResponseBodyList {
	return s.List
}

func (s *ListHoneypotAttackerPortraitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHoneypotAttackerPortraitResponseBody) GetPageInfo() *ListHoneypotAttackerPortraitResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotAttackerPortraitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotAttackerPortraitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHoneypotAttackerPortraitResponseBody) SetCode(v string) *ListHoneypotAttackerPortraitResponseBody {
	s.Code = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBody) SetHttpStatusCode(v int32) *ListHoneypotAttackerPortraitResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBody) SetList(v []*ListHoneypotAttackerPortraitResponseBodyList) *ListHoneypotAttackerPortraitResponseBody {
	s.List = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBody) SetMessage(v string) *ListHoneypotAttackerPortraitResponseBody {
	s.Message = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBody) SetPageInfo(v *ListHoneypotAttackerPortraitResponseBodyPageInfo) *ListHoneypotAttackerPortraitResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBody) SetRequestId(v string) *ListHoneypotAttackerPortraitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBody) SetSuccess(v bool) *ListHoneypotAttackerPortraitResponseBody {
	s.Success = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotAttackerPortraitResponseBodyList struct {
	// The number of attacks.
	//
	// example:
	//
	// 10
	AttackCount *int32 `json:"AttackCount,omitempty" xml:"AttackCount,omitempty"`
	// The information about the browsers of the attack source.
	Browser []*string `json:"Browser,omitempty" xml:"Browser,omitempty" type:"Repeated"`
	// The information about the hosts of the attack source.
	Host []*string `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
	// The timestamp at which the attack was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1679896965
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The network information about the attack source.
	Network *ListHoneypotAttackerPortraitResponseBodyListNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The attacker profile ID.
	//
	// example:
	//
	// cd48604a-1694-4f03-ade0-ec6994c3****
	PortraitId *string `json:"PortraitId,omitempty" xml:"PortraitId,omitempty"`
	// The social information about the attack source.
	Social []*string `json:"Social,omitempty" xml:"Social,omitempty" type:"Repeated"`
}

func (s ListHoneypotAttackerPortraitResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerPortraitResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) GetAttackCount() *int32 {
	return s.AttackCount
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) GetBrowser() []*string {
	return s.Browser
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) GetHost() []*string {
	return s.Host
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) GetLastTime() *int64 {
	return s.LastTime
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) GetNetwork() *ListHoneypotAttackerPortraitResponseBodyListNetwork {
	return s.Network
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) GetPortraitId() *string {
	return s.PortraitId
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) GetSocial() []*string {
	return s.Social
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) SetAttackCount(v int32) *ListHoneypotAttackerPortraitResponseBodyList {
	s.AttackCount = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) SetBrowser(v []*string) *ListHoneypotAttackerPortraitResponseBodyList {
	s.Browser = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) SetHost(v []*string) *ListHoneypotAttackerPortraitResponseBodyList {
	s.Host = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) SetLastTime(v int64) *ListHoneypotAttackerPortraitResponseBodyList {
	s.LastTime = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) SetNetwork(v *ListHoneypotAttackerPortraitResponseBodyListNetwork) *ListHoneypotAttackerPortraitResponseBodyList {
	s.Network = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) SetPortraitId(v string) *ListHoneypotAttackerPortraitResponseBodyList {
	s.PortraitId = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) SetSocial(v []*string) *ListHoneypotAttackerPortraitResponseBodyList {
	s.Social = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotAttackerPortraitResponseBodyListNetwork struct {
	// The public IP addresses.
	ExternalIp []*string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty" type:"Repeated"`
	// The private IP addresses.
	InternalIp []*string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty" type:"Repeated"`
	// The originating IP addresses.
	RealIp []*string `json:"RealIp,omitempty" xml:"RealIp,omitempty" type:"Repeated"`
}

func (s ListHoneypotAttackerPortraitResponseBodyListNetwork) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerPortraitResponseBodyListNetwork) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerPortraitResponseBodyListNetwork) GetExternalIp() []*string {
	return s.ExternalIp
}

func (s *ListHoneypotAttackerPortraitResponseBodyListNetwork) GetInternalIp() []*string {
	return s.InternalIp
}

func (s *ListHoneypotAttackerPortraitResponseBodyListNetwork) GetRealIp() []*string {
	return s.RealIp
}

func (s *ListHoneypotAttackerPortraitResponseBodyListNetwork) SetExternalIp(v []*string) *ListHoneypotAttackerPortraitResponseBodyListNetwork {
	s.ExternalIp = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyListNetwork) SetInternalIp(v []*string) *ListHoneypotAttackerPortraitResponseBodyListNetwork {
	s.InternalIp = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyListNetwork) SetRealIp(v []*string) *ListHoneypotAttackerPortraitResponseBodyListNetwork {
	s.RealIp = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyListNetwork) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotAttackerPortraitResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 11
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 2
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
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHoneypotAttackerPortraitResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerPortraitResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) SetCount(v int32) *ListHoneypotAttackerPortraitResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotAttackerPortraitResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotAttackerPortraitResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotAttackerPortraitResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
