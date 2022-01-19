// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateConsoleAccessWhiteListRequest struct {
	DstPort          *int32  `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	InstanceIdList   *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	InstanceInfoList *string `json:"InstanceInfoList,omitempty" xml:"InstanceInfoList,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LiveTime         *int32  `json:"LiveTime,omitempty" xml:"LiveTime,omitempty"`
	Note             *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ProductName      *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SourceCode       *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SrcIP            *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	WhiteListType    *int32  `json:"WhiteListType,omitempty" xml:"WhiteListType,omitempty"`
}

func (s CreateConsoleAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsoleAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *CreateConsoleAccessWhiteListRequest) SetDstPort(v int32) *CreateConsoleAccessWhiteListRequest {
	s.DstPort = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetInstanceIdList(v string) *CreateConsoleAccessWhiteListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetInstanceInfoList(v string) *CreateConsoleAccessWhiteListRequest {
	s.InstanceInfoList = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetLang(v string) *CreateConsoleAccessWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetLiveTime(v int32) *CreateConsoleAccessWhiteListRequest {
	s.LiveTime = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetNote(v string) *CreateConsoleAccessWhiteListRequest {
	s.Note = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetProductName(v string) *CreateConsoleAccessWhiteListRequest {
	s.ProductName = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetResourceOwnerId(v int64) *CreateConsoleAccessWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetSourceCode(v string) *CreateConsoleAccessWhiteListRequest {
	s.SourceCode = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetSourceIp(v string) *CreateConsoleAccessWhiteListRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetSrcIP(v string) *CreateConsoleAccessWhiteListRequest {
	s.SrcIP = &v
	return s
}

func (s *CreateConsoleAccessWhiteListRequest) SetWhiteListType(v int32) *CreateConsoleAccessWhiteListRequest {
	s.WhiteListType = &v
	return s
}

type CreateConsoleAccessWhiteListResponseBody struct {
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConsoleAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConsoleAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsoleAccessWhiteListResponseBody) SetModule(v string) *CreateConsoleAccessWhiteListResponseBody {
	s.Module = &v
	return s
}

func (s *CreateConsoleAccessWhiteListResponseBody) SetRequestId(v string) *CreateConsoleAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type CreateConsoleAccessWhiteListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConsoleAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConsoleAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConsoleAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *CreateConsoleAccessWhiteListResponse) SetHeaders(v map[string]*string) *CreateConsoleAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *CreateConsoleAccessWhiteListResponse) SetBody(v *CreateConsoleAccessWhiteListResponseBody) *CreateConsoleAccessWhiteListResponse {
	s.Body = v
	return s
}

type DeleteConsoleAccessWhiteListRequest struct {
	DisableWhitelist *string `json:"DisableWhitelist,omitempty" xml:"DisableWhitelist,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceCode       *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteConsoleAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsoleAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsoleAccessWhiteListRequest) SetDisableWhitelist(v string) *DeleteConsoleAccessWhiteListRequest {
	s.DisableWhitelist = &v
	return s
}

func (s *DeleteConsoleAccessWhiteListRequest) SetLang(v string) *DeleteConsoleAccessWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *DeleteConsoleAccessWhiteListRequest) SetSourceCode(v string) *DeleteConsoleAccessWhiteListRequest {
	s.SourceCode = &v
	return s
}

func (s *DeleteConsoleAccessWhiteListRequest) SetSourceIp(v string) *DeleteConsoleAccessWhiteListRequest {
	s.SourceIp = &v
	return s
}

type DeleteConsoleAccessWhiteListResponseBody struct {
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConsoleAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsoleAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsoleAccessWhiteListResponseBody) SetModule(v string) *DeleteConsoleAccessWhiteListResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteConsoleAccessWhiteListResponseBody) SetRequestId(v string) *DeleteConsoleAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConsoleAccessWhiteListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConsoleAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConsoleAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsoleAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsoleAccessWhiteListResponse) SetHeaders(v map[string]*string) *DeleteConsoleAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsoleAccessWhiteListResponse) SetBody(v *DeleteConsoleAccessWhiteListResponseBody) *DeleteConsoleAccessWhiteListResponse {
	s.Body = v
	return s
}

type DescribeAccessWhiteListSlbListRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAccessWhiteListSlbListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessWhiteListSlbListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessWhiteListSlbListRequest) SetLang(v string) *DescribeAccessWhiteListSlbListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAccessWhiteListSlbListRequest) SetSourceCode(v string) *DescribeAccessWhiteListSlbListRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeAccessWhiteListSlbListRequest) SetSourceIp(v string) *DescribeAccessWhiteListSlbListRequest {
	s.SourceIp = &v
	return s
}

type DescribeAccessWhiteListSlbListResponseBody struct {
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlbList    []*DescribeAccessWhiteListSlbListResponseBodySlbList `json:"SlbList,omitempty" xml:"SlbList,omitempty" type:"Repeated"`
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Module     *string                                              `json:"module,omitempty" xml:"module,omitempty"`
}

func (s DescribeAccessWhiteListSlbListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessWhiteListSlbListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessWhiteListSlbListResponseBody) SetRequestId(v string) *DescribeAccessWhiteListSlbListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessWhiteListSlbListResponseBody) SetSlbList(v []*DescribeAccessWhiteListSlbListResponseBodySlbList) *DescribeAccessWhiteListSlbListResponseBody {
	s.SlbList = v
	return s
}

func (s *DescribeAccessWhiteListSlbListResponseBody) SetTotalCount(v int32) *DescribeAccessWhiteListSlbListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessWhiteListSlbListResponseBody) SetModule(v string) *DescribeAccessWhiteListSlbListResponseBody {
	s.Module = &v
	return s
}

type DescribeAccessWhiteListSlbListResponseBodySlbList struct {
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ItemSign     *string `json:"ItemSign,omitempty" xml:"ItemSign,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeAccessWhiteListSlbListResponseBodySlbList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessWhiteListSlbListResponseBodySlbList) GoString() string {
	return s.String()
}

func (s *DescribeAccessWhiteListSlbListResponseBodySlbList) SetIP(v string) *DescribeAccessWhiteListSlbListResponseBodySlbList {
	s.IP = &v
	return s
}

func (s *DescribeAccessWhiteListSlbListResponseBodySlbList) SetInstanceId(v string) *DescribeAccessWhiteListSlbListResponseBodySlbList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccessWhiteListSlbListResponseBodySlbList) SetInstanceName(v string) *DescribeAccessWhiteListSlbListResponseBodySlbList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAccessWhiteListSlbListResponseBodySlbList) SetItemSign(v string) *DescribeAccessWhiteListSlbListResponseBodySlbList {
	s.ItemSign = &v
	return s
}

func (s *DescribeAccessWhiteListSlbListResponseBodySlbList) SetRegion(v string) *DescribeAccessWhiteListSlbListResponseBodySlbList {
	s.Region = &v
	return s
}

type DescribeAccessWhiteListSlbListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccessWhiteListSlbListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessWhiteListSlbListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessWhiteListSlbListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessWhiteListSlbListResponse) SetHeaders(v map[string]*string) *DescribeAccessWhiteListSlbListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessWhiteListSlbListResponse) SetBody(v *DescribeAccessWhiteListSlbListResponseBody) *DescribeAccessWhiteListSlbListResponse {
	s.Body = v
	return s
}

type DescribeAccessWhitelistEcsListRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAccessWhitelistEcsListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessWhitelistEcsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessWhitelistEcsListRequest) SetLang(v string) *DescribeAccessWhitelistEcsListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAccessWhitelistEcsListRequest) SetSourceCode(v string) *DescribeAccessWhitelistEcsListRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeAccessWhitelistEcsListRequest) SetSourceIp(v string) *DescribeAccessWhitelistEcsListRequest {
	s.SourceIp = &v
	return s
}

type DescribeAccessWhitelistEcsListResponseBody struct {
	EcsList    []*DescribeAccessWhitelistEcsListResponseBodyEcsList `json:"EcsList,omitempty" xml:"EcsList,omitempty" type:"Repeated"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Module     *string                                              `json:"module,omitempty" xml:"module,omitempty"`
}

func (s DescribeAccessWhitelistEcsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessWhitelistEcsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessWhitelistEcsListResponseBody) SetEcsList(v []*DescribeAccessWhitelistEcsListResponseBodyEcsList) *DescribeAccessWhitelistEcsListResponseBody {
	s.EcsList = v
	return s
}

func (s *DescribeAccessWhitelistEcsListResponseBody) SetRequestId(v string) *DescribeAccessWhitelistEcsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessWhitelistEcsListResponseBody) SetTotalCount(v int32) *DescribeAccessWhitelistEcsListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessWhitelistEcsListResponseBody) SetModule(v string) *DescribeAccessWhitelistEcsListResponseBody {
	s.Module = &v
	return s
}

type DescribeAccessWhitelistEcsListResponseBodyEcsList struct {
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DescribeAccessWhitelistEcsListResponseBodyEcsList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessWhitelistEcsListResponseBodyEcsList) GoString() string {
	return s.String()
}

func (s *DescribeAccessWhitelistEcsListResponseBodyEcsList) SetIP(v string) *DescribeAccessWhitelistEcsListResponseBodyEcsList {
	s.IP = &v
	return s
}

func (s *DescribeAccessWhitelistEcsListResponseBodyEcsList) SetInstanceId(v string) *DescribeAccessWhitelistEcsListResponseBodyEcsList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccessWhitelistEcsListResponseBodyEcsList) SetInstanceName(v string) *DescribeAccessWhitelistEcsListResponseBodyEcsList {
	s.InstanceName = &v
	return s
}

type DescribeAccessWhitelistEcsListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccessWhitelistEcsListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessWhitelistEcsListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessWhitelistEcsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessWhitelistEcsListResponse) SetHeaders(v map[string]*string) *DescribeAccessWhitelistEcsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessWhitelistEcsListResponse) SetBody(v *DescribeAccessWhitelistEcsListResponseBody) *DescribeAccessWhitelistEcsListResponse {
	s.Body = v
	return s
}

type DescribeAttackEventRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime      *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType  *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServerIpList *string `json:"ServerIpList,omitempty" xml:"ServerIpList,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime    *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAttackEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttackEventRequest) SetCurrentPage(v int32) *DescribeAttackEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAttackEventRequest) SetEndTime(v int32) *DescribeAttackEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAttackEventRequest) SetLang(v string) *DescribeAttackEventRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAttackEventRequest) SetPageSize(v int32) *DescribeAttackEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAttackEventRequest) SetProductType(v string) *DescribeAttackEventRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeAttackEventRequest) SetRegion(v string) *DescribeAttackEventRequest {
	s.Region = &v
	return s
}

func (s *DescribeAttackEventRequest) SetServerIpList(v string) *DescribeAttackEventRequest {
	s.ServerIpList = &v
	return s
}

func (s *DescribeAttackEventRequest) SetSourceIp(v string) *DescribeAttackEventRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAttackEventRequest) SetStartTime(v int32) *DescribeAttackEventRequest {
	s.StartTime = &v
	return s
}

type DescribeAttackEventResponseBody struct {
	EventList []*DescribeAttackEventResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	Module    *string                                     `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAttackEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttackEventResponseBody) SetEventList(v []*DescribeAttackEventResponseBodyEventList) *DescribeAttackEventResponseBody {
	s.EventList = v
	return s
}

func (s *DescribeAttackEventResponseBody) SetModule(v string) *DescribeAttackEventResponseBody {
	s.Module = &v
	return s
}

func (s *DescribeAttackEventResponseBody) SetRequestId(v string) *DescribeAttackEventResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAttackEventResponseBodyEventList struct {
	AttackType     *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateStamp *int32  `json:"GmtCreateStamp,omitempty" xml:"GmtCreateStamp,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VmIp           *string `json:"VmIp,omitempty" xml:"VmIp,omitempty"`
}

func (s DescribeAttackEventResponseBodyEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackEventResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *DescribeAttackEventResponseBodyEventList) SetAttackType(v string) *DescribeAttackEventResponseBodyEventList {
	s.AttackType = &v
	return s
}

func (s *DescribeAttackEventResponseBodyEventList) SetGmtCreate(v string) *DescribeAttackEventResponseBodyEventList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAttackEventResponseBodyEventList) SetGmtCreateStamp(v int32) *DescribeAttackEventResponseBodyEventList {
	s.GmtCreateStamp = &v
	return s
}

func (s *DescribeAttackEventResponseBodyEventList) SetGmtModified(v string) *DescribeAttackEventResponseBodyEventList {
	s.GmtModified = &v
	return s
}

func (s *DescribeAttackEventResponseBodyEventList) SetSourceIp(v string) *DescribeAttackEventResponseBodyEventList {
	s.SourceIp = &v
	return s
}

func (s *DescribeAttackEventResponseBodyEventList) SetUrl(v string) *DescribeAttackEventResponseBodyEventList {
	s.Url = &v
	return s
}

func (s *DescribeAttackEventResponseBodyEventList) SetVmIp(v string) *DescribeAttackEventResponseBodyEventList {
	s.VmIp = &v
	return s
}

type DescribeAttackEventResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAttackEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAttackEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttackEventResponse) SetHeaders(v map[string]*string) *DescribeAttackEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttackEventResponse) SetBody(v *DescribeAttackEventResponseBody) *DescribeAttackEventResponse {
	s.Body = v
	return s
}

type DescribeAttackedIpRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime      *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType  *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServerIpList *string `json:"ServerIpList,omitempty" xml:"ServerIpList,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime    *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAttackedIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackedIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttackedIpRequest) SetCurrentPage(v int32) *DescribeAttackedIpRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAttackedIpRequest) SetEndTime(v int32) *DescribeAttackedIpRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAttackedIpRequest) SetLang(v string) *DescribeAttackedIpRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAttackedIpRequest) SetPageSize(v int32) *DescribeAttackedIpRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAttackedIpRequest) SetProductType(v string) *DescribeAttackedIpRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeAttackedIpRequest) SetRegion(v string) *DescribeAttackedIpRequest {
	s.Region = &v
	return s
}

func (s *DescribeAttackedIpRequest) SetServerIpList(v string) *DescribeAttackedIpRequest {
	s.ServerIpList = &v
	return s
}

func (s *DescribeAttackedIpRequest) SetSourceIp(v string) *DescribeAttackedIpRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAttackedIpRequest) SetStartTime(v int32) *DescribeAttackedIpRequest {
	s.StartTime = &v
	return s
}

type DescribeAttackedIpResponseBody struct {
	IpList    []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	Module    *string   `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAttackedIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackedIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttackedIpResponseBody) SetIpList(v []*string) *DescribeAttackedIpResponseBody {
	s.IpList = v
	return s
}

func (s *DescribeAttackedIpResponseBody) SetModule(v string) *DescribeAttackedIpResponseBody {
	s.Module = &v
	return s
}

func (s *DescribeAttackedIpResponseBody) SetRequestId(v string) *DescribeAttackedIpResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAttackedIpResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAttackedIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAttackedIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackedIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttackedIpResponse) SetHeaders(v map[string]*string) *DescribeAttackedIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttackedIpResponse) SetBody(v *DescribeAttackedIpResponseBody) *DescribeAttackedIpResponse {
	s.Body = v
	return s
}

type DescribeConsoleAccessWhiteListRequest struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DstIP         *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceCode    *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SrcIP         *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WhiteListType *int32  `json:"WhiteListType,omitempty" xml:"WhiteListType,omitempty"`
	QueryProduct  *string `json:"queryProduct,omitempty" xml:"queryProduct,omitempty"`
}

func (s DescribeConsoleAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsoleAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeConsoleAccessWhiteListRequest) SetCurrentPage(v int32) *DescribeConsoleAccessWhiteListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetDstIP(v string) *DescribeConsoleAccessWhiteListRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetLang(v string) *DescribeConsoleAccessWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetPageSize(v int32) *DescribeConsoleAccessWhiteListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetSourceCode(v string) *DescribeConsoleAccessWhiteListRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetSourceIp(v string) *DescribeConsoleAccessWhiteListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetSrcIP(v string) *DescribeConsoleAccessWhiteListRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetStatus(v string) *DescribeConsoleAccessWhiteListRequest {
	s.Status = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetWhiteListType(v int32) *DescribeConsoleAccessWhiteListRequest {
	s.WhiteListType = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListRequest) SetQueryProduct(v string) *DescribeConsoleAccessWhiteListRequest {
	s.QueryProduct = &v
	return s
}

type DescribeConsoleAccessWhiteListResponseBody struct {
	DataList  []*DescribeConsoleAccessWhiteListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	PageInfo  *DescribeConsoleAccessWhiteListResponseBodyPageInfo   `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string                                               `json:"module,omitempty" xml:"module,omitempty"`
}

func (s DescribeConsoleAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsoleAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsoleAccessWhiteListResponseBody) SetDataList(v []*DescribeConsoleAccessWhiteListResponseBodyDataList) *DescribeConsoleAccessWhiteListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBody) SetPageInfo(v *DescribeConsoleAccessWhiteListResponseBodyPageInfo) *DescribeConsoleAccessWhiteListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBody) SetRequestId(v string) *DescribeConsoleAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBody) SetModule(v string) *DescribeConsoleAccessWhiteListResponseBody {
	s.Module = &v
	return s
}

type DescribeConsoleAccessWhiteListResponseBodyDataList struct {
	DstIp         *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	GmtCreate     *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtRealExpire *string `json:"GmtRealExpire,omitempty" xml:"GmtRealExpire,omitempty"`
	Id            *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	InsProduct    *string `json:"InsProduct,omitempty" xml:"InsProduct,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SrcIp         *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeConsoleAccessWhiteListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsoleAccessWhiteListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeConsoleAccessWhiteListResponseBodyDataList) SetDstIp(v string) *DescribeConsoleAccessWhiteListResponseBodyDataList {
	s.DstIp = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyDataList) SetGmtCreate(v string) *DescribeConsoleAccessWhiteListResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyDataList) SetGmtRealExpire(v string) *DescribeConsoleAccessWhiteListResponseBodyDataList {
	s.GmtRealExpire = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyDataList) SetId(v int32) *DescribeConsoleAccessWhiteListResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyDataList) SetInsProduct(v string) *DescribeConsoleAccessWhiteListResponseBodyDataList {
	s.InsProduct = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyDataList) SetRegionId(v string) *DescribeConsoleAccessWhiteListResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyDataList) SetSrcIp(v string) *DescribeConsoleAccessWhiteListResponseBodyDataList {
	s.SrcIp = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyDataList) SetStatus(v string) *DescribeConsoleAccessWhiteListResponseBodyDataList {
	s.Status = &v
	return s
}

type DescribeConsoleAccessWhiteListResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	PageSize    *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total       *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s DescribeConsoleAccessWhiteListResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsoleAccessWhiteListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeConsoleAccessWhiteListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeConsoleAccessWhiteListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyPageInfo) SetPageSize(v int32) *DescribeConsoleAccessWhiteListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponseBodyPageInfo) SetTotal(v int32) *DescribeConsoleAccessWhiteListResponseBodyPageInfo {
	s.Total = &v
	return s
}

type DescribeConsoleAccessWhiteListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConsoleAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConsoleAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsoleAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeConsoleAccessWhiteListResponse) SetHeaders(v map[string]*string) *DescribeConsoleAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeConsoleAccessWhiteListResponse) SetBody(v *DescribeConsoleAccessWhiteListResponseBody) *DescribeConsoleAccessWhiteListResponse {
	s.Body = v
	return s
}

type DescribeCountAttackEventRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime      *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType  *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServerIpList *string `json:"ServerIpList,omitempty" xml:"ServerIpList,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime    *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCountAttackEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCountAttackEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeCountAttackEventRequest) SetCurrentPage(v int32) *DescribeCountAttackEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCountAttackEventRequest) SetEndTime(v int32) *DescribeCountAttackEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCountAttackEventRequest) SetLang(v string) *DescribeCountAttackEventRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCountAttackEventRequest) SetPageSize(v int32) *DescribeCountAttackEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCountAttackEventRequest) SetProductType(v string) *DescribeCountAttackEventRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeCountAttackEventRequest) SetRegion(v string) *DescribeCountAttackEventRequest {
	s.Region = &v
	return s
}

func (s *DescribeCountAttackEventRequest) SetServerIpList(v string) *DescribeCountAttackEventRequest {
	s.ServerIpList = &v
	return s
}

func (s *DescribeCountAttackEventRequest) SetSourceIp(v string) *DescribeCountAttackEventRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCountAttackEventRequest) SetStartTime(v int32) *DescribeCountAttackEventRequest {
	s.StartTime = &v
	return s
}

type DescribeCountAttackEventResponseBody struct {
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCountAttackEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCountAttackEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCountAttackEventResponseBody) SetCount(v int64) *DescribeCountAttackEventResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCountAttackEventResponseBody) SetModule(v string) *DescribeCountAttackEventResponseBody {
	s.Module = &v
	return s
}

func (s *DescribeCountAttackEventResponseBody) SetRequestId(v string) *DescribeCountAttackEventResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCountAttackEventResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCountAttackEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCountAttackEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCountAttackEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeCountAttackEventResponse) SetHeaders(v map[string]*string) *DescribeCountAttackEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeCountAttackEventResponse) SetBody(v *DescribeCountAttackEventResponseBody) *DescribeCountAttackEventResponse {
	s.Body = v
	return s
}

type DescribePhoneInfoRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PhoneNum   *string `json:"phoneNum,omitempty" xml:"phoneNum,omitempty"`
	SourceCode *string `json:"sourceCode,omitempty" xml:"sourceCode,omitempty"`
}

func (s DescribePhoneInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneInfoRequest) SetLang(v string) *DescribePhoneInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribePhoneInfoRequest) SetSourceIp(v string) *DescribePhoneInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePhoneInfoRequest) SetPhoneNum(v string) *DescribePhoneInfoRequest {
	s.PhoneNum = &v
	return s
}

func (s *DescribePhoneInfoRequest) SetSourceCode(v string) *DescribePhoneInfoRequest {
	s.SourceCode = &v
	return s
}

type DescribePhoneInfoResponseBody struct {
	Module     *string `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DetectTime *string `json:"detectTime,omitempty" xml:"detectTime,omitempty"`
	PhoneNum   *int64  `json:"phoneNum,omitempty" xml:"phoneNum,omitempty"`
	RiskLevel  *int64  `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
}

func (s DescribePhoneInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneInfoResponseBody) SetModule(v string) *DescribePhoneInfoResponseBody {
	s.Module = &v
	return s
}

func (s *DescribePhoneInfoResponseBody) SetRequestId(v string) *DescribePhoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneInfoResponseBody) SetDetectTime(v string) *DescribePhoneInfoResponseBody {
	s.DetectTime = &v
	return s
}

func (s *DescribePhoneInfoResponseBody) SetPhoneNum(v int64) *DescribePhoneInfoResponseBody {
	s.PhoneNum = &v
	return s
}

func (s *DescribePhoneInfoResponseBody) SetRiskLevel(v int64) *DescribePhoneInfoResponseBody {
	s.RiskLevel = &v
	return s
}

type DescribePhoneInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePhoneInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneInfoResponse) SetHeaders(v map[string]*string) *DescribePhoneInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneInfoResponse) SetBody(v *DescribePhoneInfoResponseBody) *DescribePhoneInfoResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("jarvis-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConsoleAccessWhiteListWithOptions(request *CreateConsoleAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *CreateConsoleAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DstPort)) {
		query["DstPort"] = request.DstPort
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceInfoList)) {
		query["InstanceInfoList"] = request.InstanceInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LiveTime)) {
		query["LiveTime"] = request.LiveTime
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		query["SourceCode"] = request.SourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SrcIP)) {
		query["SrcIP"] = request.SrcIP
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteListType)) {
		query["WhiteListType"] = request.WhiteListType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConsoleAccessWhiteList"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConsoleAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConsoleAccessWhiteList(request *CreateConsoleAccessWhiteListRequest) (_result *CreateConsoleAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConsoleAccessWhiteListResponse{}
	_body, _err := client.CreateConsoleAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConsoleAccessWhiteListWithOptions(request *DeleteConsoleAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *DeleteConsoleAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisableWhitelist)) {
		query["DisableWhitelist"] = request.DisableWhitelist
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		query["SourceCode"] = request.SourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConsoleAccessWhiteList"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConsoleAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConsoleAccessWhiteList(request *DeleteConsoleAccessWhiteListRequest) (_result *DeleteConsoleAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConsoleAccessWhiteListResponse{}
	_body, _err := client.DeleteConsoleAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessWhiteListSlbListWithOptions(request *DescribeAccessWhiteListSlbListRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessWhiteListSlbListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		query["SourceCode"] = request.SourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessWhiteListSlbList"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessWhiteListSlbListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessWhiteListSlbList(request *DescribeAccessWhiteListSlbListRequest) (_result *DescribeAccessWhiteListSlbListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessWhiteListSlbListResponse{}
	_body, _err := client.DescribeAccessWhiteListSlbListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessWhitelistEcsListWithOptions(request *DescribeAccessWhitelistEcsListRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessWhitelistEcsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		query["SourceCode"] = request.SourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessWhitelistEcsList"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessWhitelistEcsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessWhitelistEcsList(request *DescribeAccessWhitelistEcsListRequest) (_result *DescribeAccessWhitelistEcsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessWhitelistEcsListResponse{}
	_body, _err := client.DescribeAccessWhitelistEcsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAttackEventWithOptions(request *DescribeAttackEventRequest, runtime *util.RuntimeOptions) (_result *DescribeAttackEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIpList)) {
		query["ServerIpList"] = request.ServerIpList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAttackEvent"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAttackEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAttackEvent(request *DescribeAttackEventRequest) (_result *DescribeAttackEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAttackEventResponse{}
	_body, _err := client.DescribeAttackEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAttackedIpWithOptions(request *DescribeAttackedIpRequest, runtime *util.RuntimeOptions) (_result *DescribeAttackedIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIpList)) {
		query["ServerIpList"] = request.ServerIpList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAttackedIp"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAttackedIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAttackedIp(request *DescribeAttackedIpRequest) (_result *DescribeAttackedIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAttackedIpResponse{}
	_body, _err := client.DescribeAttackedIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConsoleAccessWhiteListWithOptions(request *DescribeConsoleAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeConsoleAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DstIP)) {
		query["DstIP"] = request.DstIP
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		query["SourceCode"] = request.SourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SrcIP)) {
		query["SrcIP"] = request.SrcIP
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteListType)) {
		query["WhiteListType"] = request.WhiteListType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryProduct)) {
		query["queryProduct"] = request.QueryProduct
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConsoleAccessWhiteList"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConsoleAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConsoleAccessWhiteList(request *DescribeConsoleAccessWhiteListRequest) (_result *DescribeConsoleAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConsoleAccessWhiteListResponse{}
	_body, _err := client.DescribeConsoleAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCountAttackEventWithOptions(request *DescribeCountAttackEventRequest, runtime *util.RuntimeOptions) (_result *DescribeCountAttackEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIpList)) {
		query["ServerIpList"] = request.ServerIpList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCountAttackEvent"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCountAttackEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCountAttackEvent(request *DescribeCountAttackEventRequest) (_result *DescribeCountAttackEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCountAttackEventResponse{}
	_body, _err := client.DescribeCountAttackEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneInfoWithOptions(request *DescribePhoneInfoRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		query["phoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		query["sourceCode"] = request.SourceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneInfo"),
		Version:     tea.String("2018-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneInfo(request *DescribePhoneInfoRequest) (_result *DescribePhoneInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneInfoResponse{}
	_body, _err := client.DescribePhoneInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
