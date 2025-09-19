// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindContainerNetworkConnectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnects(v []*FindContainerNetworkConnectResponseBodyConnects) *FindContainerNetworkConnectResponseBody
	GetConnects() []*FindContainerNetworkConnectResponseBodyConnects
	SetPageInfo(v *FindContainerNetworkConnectResponseBodyPageInfo) *FindContainerNetworkConnectResponseBody
	GetPageInfo() *FindContainerNetworkConnectResponseBodyPageInfo
	SetRequestId(v string) *FindContainerNetworkConnectResponseBody
	GetRequestId() *string
}

type FindContainerNetworkConnectResponseBody struct {
	// The information about the network connections.
	Connects []*FindContainerNetworkConnectResponseBodyConnects `json:"Connects,omitempty" xml:"Connects,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *FindContainerNetworkConnectResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8686CE6E-9BFA-5436-A9D9-77B984AEE7F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindContainerNetworkConnectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectResponseBody) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectResponseBody) GetConnects() []*FindContainerNetworkConnectResponseBodyConnects {
	return s.Connects
}

func (s *FindContainerNetworkConnectResponseBody) GetPageInfo() *FindContainerNetworkConnectResponseBodyPageInfo {
	return s.PageInfo
}

func (s *FindContainerNetworkConnectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindContainerNetworkConnectResponseBody) SetConnects(v []*FindContainerNetworkConnectResponseBodyConnects) *FindContainerNetworkConnectResponseBody {
	s.Connects = v
	return s
}

func (s *FindContainerNetworkConnectResponseBody) SetPageInfo(v *FindContainerNetworkConnectResponseBodyPageInfo) *FindContainerNetworkConnectResponseBody {
	s.PageInfo = v
	return s
}

func (s *FindContainerNetworkConnectResponseBody) SetRequestId(v string) *FindContainerNetworkConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBody) Validate() error {
	return dara.Validate(s)
}

type FindContainerNetworkConnectResponseBodyConnects struct {
	// The information about the destination container.
	//
	// > This parameter is not supported.
	DstContainer *FindContainerNetworkConnectResponseBodyConnectsDstContainer `json:"DstContainer,omitempty" xml:"DstContainer,omitempty" type:"Struct"`
	// The destination IP address.
	//
	// example:
	//
	// 172.20.62.176
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 443
	DstPort *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The timestamp when the connection was first established.
	//
	// example:
	//
	// 2022-11-11 20:54:32
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// The ID of the network connection.
	//
	// example:
	//
	// 1458
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timestamp when the connection was last established.
	//
	// example:
	//
	// 2022-11-24 10:26:00
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The information about the source container.
	//
	// > This parameter is not supported.
	SrcContainer *FindContainerNetworkConnectResponseBodyConnectsSrcContainer `json:"SrcContainer,omitempty" xml:"SrcContainer,omitempty" type:"Struct"`
	// The source IP address.
	//
	// example:
	//
	// 35.233.62.116
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// The source port.
	//
	// example:
	//
	// 10240
	SrcPort *string `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
}

func (s FindContainerNetworkConnectResponseBodyConnects) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectResponseBodyConnects) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetDstContainer() *FindContainerNetworkConnectResponseBodyConnectsDstContainer {
	return s.DstContainer
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetDstIp() *string {
	return s.DstIp
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetDstPort() *string {
	return s.DstPort
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetId() *int64 {
	return s.Id
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetLastTime() *int64 {
	return s.LastTime
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetSrcContainer() *FindContainerNetworkConnectResponseBodyConnectsSrcContainer {
	return s.SrcContainer
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetSrcIp() *string {
	return s.SrcIp
}

func (s *FindContainerNetworkConnectResponseBodyConnects) GetSrcPort() *string {
	return s.SrcPort
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetDstContainer(v *FindContainerNetworkConnectResponseBodyConnectsDstContainer) *FindContainerNetworkConnectResponseBodyConnects {
	s.DstContainer = v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetDstIp(v string) *FindContainerNetworkConnectResponseBodyConnects {
	s.DstIp = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetDstPort(v string) *FindContainerNetworkConnectResponseBodyConnects {
	s.DstPort = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetFirstTime(v int64) *FindContainerNetworkConnectResponseBodyConnects {
	s.FirstTime = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetId(v int64) *FindContainerNetworkConnectResponseBodyConnects {
	s.Id = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetLastTime(v int64) *FindContainerNetworkConnectResponseBodyConnects {
	s.LastTime = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetSrcContainer(v *FindContainerNetworkConnectResponseBodyConnectsSrcContainer) *FindContainerNetworkConnectResponseBodyConnects {
	s.SrcContainer = v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetSrcIp(v string) *FindContainerNetworkConnectResponseBodyConnects {
	s.SrcIp = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) SetSrcPort(v string) *FindContainerNetworkConnectResponseBodyConnects {
	s.SrcPort = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnects) Validate() error {
	return dara.Validate(s)
}

type FindContainerNetworkConnectResponseBodyConnectsDstContainer struct {
	// The ID of the destination container.
	//
	// example:
	//
	// 48a6dxxx9d5a5866
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
}

func (s FindContainerNetworkConnectResponseBodyConnectsDstContainer) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectResponseBodyConnectsDstContainer) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectResponseBodyConnectsDstContainer) GetContainerId() *string {
	return s.ContainerId
}

func (s *FindContainerNetworkConnectResponseBodyConnectsDstContainer) SetContainerId(v string) *FindContainerNetworkConnectResponseBodyConnectsDstContainer {
	s.ContainerId = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnectsDstContainer) Validate() error {
	return dara.Validate(s)
}

type FindContainerNetworkConnectResponseBodyConnectsSrcContainer struct {
	// The ID of the source container.
	//
	// example:
	//
	// 48a6xxx5709d5a5866
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
}

func (s FindContainerNetworkConnectResponseBodyConnectsSrcContainer) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectResponseBodyConnectsSrcContainer) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectResponseBodyConnectsSrcContainer) GetContainerId() *string {
	return s.ContainerId
}

func (s *FindContainerNetworkConnectResponseBodyConnectsSrcContainer) SetContainerId(v string) *FindContainerNetworkConnectResponseBodyConnectsSrcContainer {
	s.ContainerId = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyConnectsSrcContainer) Validate() error {
	return dara.Validate(s)
}

type FindContainerNetworkConnectResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindContainerNetworkConnectResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) GetCount() *int64 {
	return s.Count
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) SetCount(v int64) *FindContainerNetworkConnectResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) SetCurrentPage(v int64) *FindContainerNetworkConnectResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) SetPageSize(v int64) *FindContainerNetworkConnectResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) SetTotalCount(v int64) *FindContainerNetworkConnectResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *FindContainerNetworkConnectResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
