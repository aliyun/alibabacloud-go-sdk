// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeBootRequest struct {
	PropertyUuid *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
}

func (s DescribeBootRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBootRequest) GoString() string {
	return s.String()
}

func (s *DescribeBootRequest) SetPropertyUuid(v string) *DescribeBootRequest {
	s.PropertyUuid = &v
	return s
}

type DescribeBootResponseBody struct {
	GmtCreate            *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WhiteListAffiliation *int32  `json:"WhiteListAffiliation,omitempty" xml:"WhiteListAffiliation,omitempty"`
	WhiteListDetail      *string `json:"WhiteListDetail,omitempty" xml:"WhiteListDetail,omitempty"`
}

func (s DescribeBootResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBootResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBootResponseBody) SetGmtCreate(v string) *DescribeBootResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBootResponseBody) SetGmtModified(v string) *DescribeBootResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeBootResponseBody) SetId(v int64) *DescribeBootResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeBootResponseBody) SetRequestId(v string) *DescribeBootResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBootResponseBody) SetWhiteListAffiliation(v int32) *DescribeBootResponseBody {
	s.WhiteListAffiliation = &v
	return s
}

func (s *DescribeBootResponseBody) SetWhiteListDetail(v string) *DescribeBootResponseBody {
	s.WhiteListDetail = &v
	return s
}

type DescribeBootResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBootResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBootResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBootResponse) GoString() string {
	return s.String()
}

func (s *DescribeBootResponse) SetHeaders(v map[string]*string) *DescribeBootResponse {
	s.Headers = v
	return s
}

func (s *DescribeBootResponse) SetStatusCode(v int32) *DescribeBootResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBootResponse) SetBody(v *DescribeBootResponseBody) *DescribeBootResponse {
	s.Body = v
	return s
}

type DescribeEventsRequest struct {
	CreateEndDate     *int64  `json:"CreateEndDate,omitempty" xml:"CreateEndDate,omitempty"`
	CreateStartDate   *int64  `json:"CreateStartDate,omitempty" xml:"CreateStartDate,omitempty"`
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EventAffiliation  *int32  `json:"EventAffiliation,omitempty" xml:"EventAffiliation,omitempty"`
	EventLevel        *int32  `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventStatus       *int32  `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	EventType         *int32  `json:"EventType,omitempty" xml:"EventType,omitempty"`
	HandleEndDate     *int64  `json:"HandleEndDate,omitempty" xml:"HandleEndDate,omitempty"`
	HandleStartDate   *int64  `json:"HandleStartDate,omitempty" xml:"HandleStartDate,omitempty"`
	HandleType        *int32  `json:"HandleType,omitempty" xml:"HandleType,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PropertyName      *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	PropertyPrivateIp *string `json:"PropertyPrivateIp,omitempty" xml:"PropertyPrivateIp,omitempty"`
	PropertyPublicIp  *string `json:"PropertyPublicIp,omitempty" xml:"PropertyPublicIp,omitempty"`
	PropertyUuid      *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
	Suspect           *string `json:"Suspect,omitempty" xml:"Suspect,omitempty"`
	SuspectSig        *string `json:"SuspectSig,omitempty" xml:"SuspectSig,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) SetCreateEndDate(v int64) *DescribeEventsRequest {
	s.CreateEndDate = &v
	return s
}

func (s *DescribeEventsRequest) SetCreateStartDate(v int64) *DescribeEventsRequest {
	s.CreateStartDate = &v
	return s
}

func (s *DescribeEventsRequest) SetCurrentPage(v int32) *DescribeEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventsRequest) SetEventAffiliation(v int32) *DescribeEventsRequest {
	s.EventAffiliation = &v
	return s
}

func (s *DescribeEventsRequest) SetEventLevel(v int32) *DescribeEventsRequest {
	s.EventLevel = &v
	return s
}

func (s *DescribeEventsRequest) SetEventStatus(v int32) *DescribeEventsRequest {
	s.EventStatus = &v
	return s
}

func (s *DescribeEventsRequest) SetEventType(v int32) *DescribeEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeEventsRequest) SetHandleEndDate(v int64) *DescribeEventsRequest {
	s.HandleEndDate = &v
	return s
}

func (s *DescribeEventsRequest) SetHandleStartDate(v int64) *DescribeEventsRequest {
	s.HandleStartDate = &v
	return s
}

func (s *DescribeEventsRequest) SetHandleType(v int32) *DescribeEventsRequest {
	s.HandleType = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int32) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetPropertyName(v string) *DescribeEventsRequest {
	s.PropertyName = &v
	return s
}

func (s *DescribeEventsRequest) SetPropertyPrivateIp(v string) *DescribeEventsRequest {
	s.PropertyPrivateIp = &v
	return s
}

func (s *DescribeEventsRequest) SetPropertyPublicIp(v string) *DescribeEventsRequest {
	s.PropertyPublicIp = &v
	return s
}

func (s *DescribeEventsRequest) SetPropertyUuid(v string) *DescribeEventsRequest {
	s.PropertyUuid = &v
	return s
}

func (s *DescribeEventsRequest) SetSuspect(v string) *DescribeEventsRequest {
	s.Suspect = &v
	return s
}

func (s *DescribeEventsRequest) SetSuspectSig(v string) *DescribeEventsRequest {
	s.SuspectSig = &v
	return s
}

type DescribeEventsResponseBody struct {
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribeEventsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) SetCurrentPage(v int32) *DescribeEventsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventsResponseBody) SetItems(v []*DescribeEventsResponseBodyItems) *DescribeEventsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeEventsResponseBody) SetPageSize(v int32) *DescribeEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalCount(v int64) *DescribeEventsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEventsResponseBodyItems struct {
	Accumulation      *string `json:"Accumulation,omitempty" xml:"Accumulation,omitempty"`
	EventAffiliation  *string `json:"EventAffiliation,omitempty" xml:"EventAffiliation,omitempty"`
	EventLevel        *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventStatus       *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	EventType         *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	EventUuid         *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtHandle         *int64  `json:"GmtHandle,omitempty" xml:"GmtHandle,omitempty"`
	GmtModified       *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HandleType        *string `json:"HandleType,omitempty" xml:"HandleType,omitempty"`
	PropertyName      *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	PropertyPrivateIp *string `json:"PropertyPrivateIp,omitempty" xml:"PropertyPrivateIp,omitempty"`
	PropertyPublicIp  *string `json:"PropertyPublicIp,omitempty" xml:"PropertyPublicIp,omitempty"`
	PropertyUuid      *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
	Suspect           *string `json:"Suspect,omitempty" xml:"Suspect,omitempty"`
	SuspectSig        *string `json:"SuspectSig,omitempty" xml:"SuspectSig,omitempty"`
	SuspectWhiteList  *string `json:"SuspectWhiteList,omitempty" xml:"SuspectWhiteList,omitempty"`
}

func (s DescribeEventsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyItems) SetAccumulation(v string) *DescribeEventsResponseBodyItems {
	s.Accumulation = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetEventAffiliation(v string) *DescribeEventsResponseBodyItems {
	s.EventAffiliation = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetEventLevel(v string) *DescribeEventsResponseBodyItems {
	s.EventLevel = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetEventStatus(v string) *DescribeEventsResponseBodyItems {
	s.EventStatus = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetEventType(v string) *DescribeEventsResponseBodyItems {
	s.EventType = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetEventUuid(v string) *DescribeEventsResponseBodyItems {
	s.EventUuid = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetGmtCreate(v int64) *DescribeEventsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetGmtHandle(v int64) *DescribeEventsResponseBodyItems {
	s.GmtHandle = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetGmtModified(v int64) *DescribeEventsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetHandleType(v string) *DescribeEventsResponseBodyItems {
	s.HandleType = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetPropertyName(v string) *DescribeEventsResponseBodyItems {
	s.PropertyName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetPropertyPrivateIp(v string) *DescribeEventsResponseBodyItems {
	s.PropertyPrivateIp = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetPropertyPublicIp(v string) *DescribeEventsResponseBodyItems {
	s.PropertyPublicIp = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetPropertyUuid(v string) *DescribeEventsResponseBodyItems {
	s.PropertyUuid = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSuspect(v string) *DescribeEventsResponseBodyItems {
	s.Suspect = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSuspectSig(v string) *DescribeEventsResponseBodyItems {
	s.SuspectSig = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSuspectWhiteList(v string) *DescribeEventsResponseBodyItems {
	s.SuspectWhiteList = &v
	return s
}

type DescribeEventsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponse) SetHeaders(v map[string]*string) *DescribeEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsResponse) SetStatusCode(v int32) *DescribeEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventsResponse) SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	PropertyUuid *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetPropertyUuid(v string) *DescribeInstanceRequest {
	s.PropertyUuid = &v
	return s
}

type DescribeInstanceResponseBody struct {
	Extensions           *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	GmtCreate            *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtRecentReport      *int64  `json:"GmtRecentReport,omitempty" xml:"GmtRecentReport,omitempty"`
	OnlineStatus         *int32  `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	ProgramExceptionNum  *int32  `json:"ProgramExceptionNum,omitempty" xml:"ProgramExceptionNum,omitempty"`
	ProgramTrustDetail   *string `json:"ProgramTrustDetail,omitempty" xml:"ProgramTrustDetail,omitempty"`
	ProgramTrustStatus   *int32  `json:"ProgramTrustStatus,omitempty" xml:"ProgramTrustStatus,omitempty"`
	ProgramWhiteListId   *int32  `json:"ProgramWhiteListId,omitempty" xml:"ProgramWhiteListId,omitempty"`
	ProgramWhiteListName *string `json:"ProgramWhiteListName,omitempty" xml:"ProgramWhiteListName,omitempty"`
	PropertyAffiliation  *int32  `json:"PropertyAffiliation,omitempty" xml:"PropertyAffiliation,omitempty"`
	PropertyName         *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	PropertyPrivateIp    *string `json:"PropertyPrivateIp,omitempty" xml:"PropertyPrivateIp,omitempty"`
	PropertyPublicIp     *string `json:"PropertyPublicIp,omitempty" xml:"PropertyPublicIp,omitempty"`
	PropertyUuid         *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemExceptionNum   *int32  `json:"SystemExceptionNum,omitempty" xml:"SystemExceptionNum,omitempty"`
	SystemTrustDetail    *string `json:"SystemTrustDetail,omitempty" xml:"SystemTrustDetail,omitempty"`
	SystemTrustStatus    *int32  `json:"SystemTrustStatus,omitempty" xml:"SystemTrustStatus,omitempty"`
	SystemWhiteListId    *int32  `json:"SystemWhiteListId,omitempty" xml:"SystemWhiteListId,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetExtensions(v string) *DescribeInstanceResponseBody {
	s.Extensions = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetGmtCreate(v int64) *DescribeInstanceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetGmtModified(v int64) *DescribeInstanceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetGmtRecentReport(v int64) *DescribeInstanceResponseBody {
	s.GmtRecentReport = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetOnlineStatus(v int32) *DescribeInstanceResponseBody {
	s.OnlineStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProgramExceptionNum(v int32) *DescribeInstanceResponseBody {
	s.ProgramExceptionNum = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProgramTrustDetail(v string) *DescribeInstanceResponseBody {
	s.ProgramTrustDetail = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProgramTrustStatus(v int32) *DescribeInstanceResponseBody {
	s.ProgramTrustStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProgramWhiteListId(v int32) *DescribeInstanceResponseBody {
	s.ProgramWhiteListId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProgramWhiteListName(v string) *DescribeInstanceResponseBody {
	s.ProgramWhiteListName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPropertyAffiliation(v int32) *DescribeInstanceResponseBody {
	s.PropertyAffiliation = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPropertyName(v string) *DescribeInstanceResponseBody {
	s.PropertyName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPropertyPrivateIp(v string) *DescribeInstanceResponseBody {
	s.PropertyPrivateIp = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPropertyPublicIp(v string) *DescribeInstanceResponseBody {
	s.PropertyPublicIp = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPropertyUuid(v string) *DescribeInstanceResponseBody {
	s.PropertyUuid = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetSystemExceptionNum(v int32) *DescribeInstanceResponseBody {
	s.SystemExceptionNum = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetSystemTrustDetail(v string) *DescribeInstanceResponseBody {
	s.SystemTrustDetail = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetSystemTrustStatus(v int32) *DescribeInstanceResponseBody {
	s.SystemTrustStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetSystemWhiteListId(v int32) *DescribeInstanceResponseBody {
	s.SystemWhiteListId = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type GenerateAikcertRequest struct {
	AikName      *string `json:"AikName,omitempty" xml:"AikName,omitempty"`
	AikPub       *string `json:"AikPub,omitempty" xml:"AikPub,omitempty"`
	EkCert       *string `json:"EkCert,omitempty" xml:"EkCert,omitempty"`
	EkPub        *string `json:"EkPub,omitempty" xml:"EkPub,omitempty"`
	NonceDigest  *string `json:"NonceDigest,omitempty" xml:"NonceDigest,omitempty"`
	PropertyUuid *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
}

func (s GenerateAikcertRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAikcertRequest) GoString() string {
	return s.String()
}

func (s *GenerateAikcertRequest) SetAikName(v string) *GenerateAikcertRequest {
	s.AikName = &v
	return s
}

func (s *GenerateAikcertRequest) SetAikPub(v string) *GenerateAikcertRequest {
	s.AikPub = &v
	return s
}

func (s *GenerateAikcertRequest) SetEkCert(v string) *GenerateAikcertRequest {
	s.EkCert = &v
	return s
}

func (s *GenerateAikcertRequest) SetEkPub(v string) *GenerateAikcertRequest {
	s.EkPub = &v
	return s
}

func (s *GenerateAikcertRequest) SetNonceDigest(v string) *GenerateAikcertRequest {
	s.NonceDigest = &v
	return s
}

func (s *GenerateAikcertRequest) SetPropertyUuid(v string) *GenerateAikcertRequest {
	s.PropertyUuid = &v
	return s
}

type GenerateAikcertResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GenerateAikcertResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GenerateAikcertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAikcertResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAikcertResponseBody) SetRequestId(v string) *GenerateAikcertResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAikcertResponseBody) SetData(v *GenerateAikcertResponseBodyData) *GenerateAikcertResponseBody {
	s.Data = v
	return s
}

type GenerateAikcertResponseBodyData struct {
	DataCredentialBlob *string `json:"dataCredentialBlob,omitempty" xml:"dataCredentialBlob,omitempty"`
	KeyCredentialBlob  *string `json:"keyCredentialBlob,omitempty" xml:"keyCredentialBlob,omitempty"`
}

func (s GenerateAikcertResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateAikcertResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateAikcertResponseBodyData) SetDataCredentialBlob(v string) *GenerateAikcertResponseBodyData {
	s.DataCredentialBlob = &v
	return s
}

func (s *GenerateAikcertResponseBodyData) SetKeyCredentialBlob(v string) *GenerateAikcertResponseBodyData {
	s.KeyCredentialBlob = &v
	return s
}

type GenerateAikcertResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateAikcertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateAikcertResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateAikcertResponse) GoString() string {
	return s.String()
}

func (s *GenerateAikcertResponse) SetHeaders(v map[string]*string) *GenerateAikcertResponse {
	s.Headers = v
	return s
}

func (s *GenerateAikcertResponse) SetStatusCode(v int32) *GenerateAikcertResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAikcertResponse) SetBody(v *GenerateAikcertResponseBody) *GenerateAikcertResponse {
	s.Body = v
	return s
}

type GenerateNonceRequest struct {
	AikName      *string `json:"AikName,omitempty" xml:"AikName,omitempty"`
	EkCert       *string `json:"EkCert,omitempty" xml:"EkCert,omitempty"`
	EkPub        *string `json:"EkPub,omitempty" xml:"EkPub,omitempty"`
	PropertyUuid *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
}

func (s GenerateNonceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateNonceRequest) GoString() string {
	return s.String()
}

func (s *GenerateNonceRequest) SetAikName(v string) *GenerateNonceRequest {
	s.AikName = &v
	return s
}

func (s *GenerateNonceRequest) SetEkCert(v string) *GenerateNonceRequest {
	s.EkCert = &v
	return s
}

func (s *GenerateNonceRequest) SetEkPub(v string) *GenerateNonceRequest {
	s.EkPub = &v
	return s
}

func (s *GenerateNonceRequest) SetPropertyUuid(v string) *GenerateNonceRequest {
	s.PropertyUuid = &v
	return s
}

type GenerateNonceResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GenerateNonceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GenerateNonceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateNonceResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateNonceResponseBody) SetRequestId(v string) *GenerateNonceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateNonceResponseBody) SetData(v *GenerateNonceResponseBodyData) *GenerateNonceResponseBody {
	s.Data = v
	return s
}

type GenerateNonceResponseBodyData struct {
	CredentialBlob *string `json:"CredentialBlob,omitempty" xml:"CredentialBlob,omitempty"`
}

func (s GenerateNonceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateNonceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateNonceResponseBodyData) SetCredentialBlob(v string) *GenerateNonceResponseBodyData {
	s.CredentialBlob = &v
	return s
}

type GenerateNonceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateNonceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateNonceResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateNonceResponse) GoString() string {
	return s.String()
}

func (s *GenerateNonceResponse) SetHeaders(v map[string]*string) *GenerateNonceResponse {
	s.Headers = v
	return s
}

func (s *GenerateNonceResponse) SetStatusCode(v int32) *GenerateNonceResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateNonceResponse) SetBody(v *GenerateNonceResponseBody) *GenerateNonceResponse {
	s.Body = v
	return s
}

type IgnoreEventsRequest struct {
	EventUuids  *string `json:"EventUuids,omitempty" xml:"EventUuids,omitempty"`
	HandleStyle *int32  `json:"HandleStyle,omitempty" xml:"HandleStyle,omitempty"`
}

func (s IgnoreEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s IgnoreEventsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreEventsRequest) SetEventUuids(v string) *IgnoreEventsRequest {
	s.EventUuids = &v
	return s
}

func (s *IgnoreEventsRequest) SetHandleStyle(v int32) *IgnoreEventsRequest {
	s.HandleStyle = &v
	return s
}

type IgnoreEventsResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IgnoreEventsResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreEventsResponseBody) SetData(v string) *IgnoreEventsResponseBody {
	s.Data = &v
	return s
}

func (s *IgnoreEventsResponseBody) SetRequestId(v string) *IgnoreEventsResponseBody {
	s.RequestId = &v
	return s
}

type IgnoreEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IgnoreEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IgnoreEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s IgnoreEventsResponse) GoString() string {
	return s.String()
}

func (s *IgnoreEventsResponse) SetHeaders(v map[string]*string) *IgnoreEventsResponse {
	s.Headers = v
	return s
}

func (s *IgnoreEventsResponse) SetStatusCode(v int32) *IgnoreEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *IgnoreEventsResponse) SetBody(v *IgnoreEventsResponseBody) *IgnoreEventsResponse {
	s.Body = v
	return s
}

type ProduceAikcertRequest struct {
	AikName          *string `json:"AikName,omitempty" xml:"AikName,omitempty"`
	CertRequest      *string `json:"CertRequest,omitempty" xml:"CertRequest,omitempty"`
	EkCert           *string `json:"EkCert,omitempty" xml:"EkCert,omitempty"`
	EkPubKey         *string `json:"EkPubKey,omitempty" xml:"EkPubKey,omitempty"`
	IncludeCertChain *bool   `json:"IncludeCertChain,omitempty" xml:"IncludeCertChain,omitempty"`
}

func (s ProduceAikcertRequest) String() string {
	return tea.Prettify(s)
}

func (s ProduceAikcertRequest) GoString() string {
	return s.String()
}

func (s *ProduceAikcertRequest) SetAikName(v string) *ProduceAikcertRequest {
	s.AikName = &v
	return s
}

func (s *ProduceAikcertRequest) SetCertRequest(v string) *ProduceAikcertRequest {
	s.CertRequest = &v
	return s
}

func (s *ProduceAikcertRequest) SetEkCert(v string) *ProduceAikcertRequest {
	s.EkCert = &v
	return s
}

func (s *ProduceAikcertRequest) SetEkPubKey(v string) *ProduceAikcertRequest {
	s.EkPubKey = &v
	return s
}

func (s *ProduceAikcertRequest) SetIncludeCertChain(v bool) *ProduceAikcertRequest {
	s.IncludeCertChain = &v
	return s
}

type ProduceAikcertResponseBody struct {
	Data      *ProduceAikcertResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProduceAikcertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProduceAikcertResponseBody) GoString() string {
	return s.String()
}

func (s *ProduceAikcertResponseBody) SetData(v *ProduceAikcertResponseBodyData) *ProduceAikcertResponseBody {
	s.Data = v
	return s
}

func (s *ProduceAikcertResponseBody) SetRequestId(v string) *ProduceAikcertResponseBody {
	s.RequestId = &v
	return s
}

type ProduceAikcertResponseBodyData struct {
	DataCredentialBlob *string `json:"DataCredentialBlob,omitempty" xml:"DataCredentialBlob,omitempty"`
	KeyCredentialBlob  *string `json:"KeyCredentialBlob,omitempty" xml:"KeyCredentialBlob,omitempty"`
}

func (s ProduceAikcertResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ProduceAikcertResponseBodyData) GoString() string {
	return s.String()
}

func (s *ProduceAikcertResponseBodyData) SetDataCredentialBlob(v string) *ProduceAikcertResponseBodyData {
	s.DataCredentialBlob = &v
	return s
}

func (s *ProduceAikcertResponseBodyData) SetKeyCredentialBlob(v string) *ProduceAikcertResponseBodyData {
	s.KeyCredentialBlob = &v
	return s
}

type ProduceAikcertResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ProduceAikcertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProduceAikcertResponse) String() string {
	return tea.Prettify(s)
}

func (s ProduceAikcertResponse) GoString() string {
	return s.String()
}

func (s *ProduceAikcertResponse) SetHeaders(v map[string]*string) *ProduceAikcertResponse {
	s.Headers = v
	return s
}

func (s *ProduceAikcertResponse) SetStatusCode(v int32) *ProduceAikcertResponse {
	s.StatusCode = &v
	return s
}

func (s *ProduceAikcertResponse) SetBody(v *ProduceAikcertResponseBody) *ProduceAikcertResponse {
	s.Body = v
	return s
}

type PutMessageRequest struct {
	FileData     *string `json:"FileData,omitempty" xml:"FileData,omitempty"`
	PropertyUuid *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
}

func (s PutMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMessageRequest) GoString() string {
	return s.String()
}

func (s *PutMessageRequest) SetFileData(v string) *PutMessageRequest {
	s.FileData = &v
	return s
}

func (s *PutMessageRequest) SetPropertyUuid(v string) *PutMessageRequest {
	s.PropertyUuid = &v
	return s
}

type PutMessageResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *PutMessageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s PutMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PutMessageResponseBody) SetRequestId(v string) *PutMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMessageResponseBody) SetData(v *PutMessageResponseBodyData) *PutMessageResponseBody {
	s.Data = v
	return s
}

type PutMessageResponseBodyData struct {
	KmoduleVerificationResult *PutMessageResponseBodyDataKmoduleVerificationResult `json:"kmoduleVerificationResult,omitempty" xml:"kmoduleVerificationResult,omitempty" type:"Struct"`
	NextClientIMAIndex        *int64                                               `json:"nextClientIMAIndex,omitempty" xml:"nextClientIMAIndex,omitempty"`
	PolicyProcResult          *PutMessageResponseBodyDataPolicyProcResult          `json:"policyProcResult,omitempty" xml:"policyProcResult,omitempty" type:"Struct"`
	ProgramVerificationResult *PutMessageResponseBodyDataProgramVerificationResult `json:"programVerificationResult,omitempty" xml:"programVerificationResult,omitempty" type:"Struct"`
	SystemVerificationResult  *PutMessageResponseBodyDataSystemVerificationResult  `json:"systemVerificationResult,omitempty" xml:"systemVerificationResult,omitempty" type:"Struct"`
}

func (s PutMessageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PutMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *PutMessageResponseBodyData) SetKmoduleVerificationResult(v *PutMessageResponseBodyDataKmoduleVerificationResult) *PutMessageResponseBodyData {
	s.KmoduleVerificationResult = v
	return s
}

func (s *PutMessageResponseBodyData) SetNextClientIMAIndex(v int64) *PutMessageResponseBodyData {
	s.NextClientIMAIndex = &v
	return s
}

func (s *PutMessageResponseBodyData) SetPolicyProcResult(v *PutMessageResponseBodyDataPolicyProcResult) *PutMessageResponseBodyData {
	s.PolicyProcResult = v
	return s
}

func (s *PutMessageResponseBodyData) SetProgramVerificationResult(v *PutMessageResponseBodyDataProgramVerificationResult) *PutMessageResponseBodyData {
	s.ProgramVerificationResult = v
	return s
}

func (s *PutMessageResponseBodyData) SetSystemVerificationResult(v *PutMessageResponseBodyDataSystemVerificationResult) *PutMessageResponseBodyData {
	s.SystemVerificationResult = v
	return s
}

type PutMessageResponseBodyDataKmoduleVerificationResult struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PutMessageResponseBodyDataKmoduleVerificationResult) String() string {
	return tea.Prettify(s)
}

func (s PutMessageResponseBodyDataKmoduleVerificationResult) GoString() string {
	return s.String()
}

func (s *PutMessageResponseBodyDataKmoduleVerificationResult) SetCode(v string) *PutMessageResponseBodyDataKmoduleVerificationResult {
	s.Code = &v
	return s
}

func (s *PutMessageResponseBodyDataKmoduleVerificationResult) SetStatus(v int32) *PutMessageResponseBodyDataKmoduleVerificationResult {
	s.Status = &v
	return s
}

type PutMessageResponseBodyDataPolicyProcResult struct {
	Code     *string `json:"code,omitempty" xml:"code,omitempty"`
	ProcData *string `json:"procData,omitempty" xml:"procData,omitempty"`
}

func (s PutMessageResponseBodyDataPolicyProcResult) String() string {
	return tea.Prettify(s)
}

func (s PutMessageResponseBodyDataPolicyProcResult) GoString() string {
	return s.String()
}

func (s *PutMessageResponseBodyDataPolicyProcResult) SetCode(v string) *PutMessageResponseBodyDataPolicyProcResult {
	s.Code = &v
	return s
}

func (s *PutMessageResponseBodyDataPolicyProcResult) SetProcData(v string) *PutMessageResponseBodyDataPolicyProcResult {
	s.ProcData = &v
	return s
}

type PutMessageResponseBodyDataProgramVerificationResult struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PutMessageResponseBodyDataProgramVerificationResult) String() string {
	return tea.Prettify(s)
}

func (s PutMessageResponseBodyDataProgramVerificationResult) GoString() string {
	return s.String()
}

func (s *PutMessageResponseBodyDataProgramVerificationResult) SetCode(v string) *PutMessageResponseBodyDataProgramVerificationResult {
	s.Code = &v
	return s
}

func (s *PutMessageResponseBodyDataProgramVerificationResult) SetStatus(v int32) *PutMessageResponseBodyDataProgramVerificationResult {
	s.Status = &v
	return s
}

type PutMessageResponseBodyDataSystemVerificationResult struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PutMessageResponseBodyDataSystemVerificationResult) String() string {
	return tea.Prettify(s)
}

func (s PutMessageResponseBodyDataSystemVerificationResult) GoString() string {
	return s.String()
}

func (s *PutMessageResponseBodyDataSystemVerificationResult) SetCode(v string) *PutMessageResponseBodyDataSystemVerificationResult {
	s.Code = &v
	return s
}

func (s *PutMessageResponseBodyDataSystemVerificationResult) SetStatus(v int32) *PutMessageResponseBodyDataSystemVerificationResult {
	s.Status = &v
	return s
}

type PutMessageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMessageResponse) GoString() string {
	return s.String()
}

func (s *PutMessageResponse) SetHeaders(v map[string]*string) *PutMessageResponse {
	s.Headers = v
	return s
}

func (s *PutMessageResponse) SetStatusCode(v int32) *PutMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMessageResponse) SetBody(v *PutMessageResponseBody) *PutMessageResponse {
	s.Body = v
	return s
}

type RegisterMessageRequest struct {
	Extensions          *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PropertyAffiliation *int32  `json:"PropertyAffiliation,omitempty" xml:"PropertyAffiliation,omitempty"`
	PropertyName        *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	PropertyPrivateIp   *string `json:"PropertyPrivateIp,omitempty" xml:"PropertyPrivateIp,omitempty"`
	PropertyPublicIp    *string `json:"PropertyPublicIp,omitempty" xml:"PropertyPublicIp,omitempty"`
	PropertyUuid        *string `json:"PropertyUuid,omitempty" xml:"PropertyUuid,omitempty"`
}

func (s RegisterMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterMessageRequest) GoString() string {
	return s.String()
}

func (s *RegisterMessageRequest) SetExtensions(v string) *RegisterMessageRequest {
	s.Extensions = &v
	return s
}

func (s *RegisterMessageRequest) SetInstanceId(v string) *RegisterMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *RegisterMessageRequest) SetInstanceType(v string) *RegisterMessageRequest {
	s.InstanceType = &v
	return s
}

func (s *RegisterMessageRequest) SetPropertyAffiliation(v int32) *RegisterMessageRequest {
	s.PropertyAffiliation = &v
	return s
}

func (s *RegisterMessageRequest) SetPropertyName(v string) *RegisterMessageRequest {
	s.PropertyName = &v
	return s
}

func (s *RegisterMessageRequest) SetPropertyPrivateIp(v string) *RegisterMessageRequest {
	s.PropertyPrivateIp = &v
	return s
}

func (s *RegisterMessageRequest) SetPropertyPublicIp(v string) *RegisterMessageRequest {
	s.PropertyPublicIp = &v
	return s
}

func (s *RegisterMessageRequest) SetPropertyUuid(v string) *RegisterMessageRequest {
	s.PropertyUuid = &v
	return s
}

type RegisterMessageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterMessageResponseBody) SetRequestId(v string) *RegisterMessageResponseBody {
	s.RequestId = &v
	return s
}

type RegisterMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterMessageResponse) GoString() string {
	return s.String()
}

func (s *RegisterMessageResponse) SetHeaders(v map[string]*string) *RegisterMessageResponse {
	s.Headers = v
	return s
}

func (s *RegisterMessageResponse) SetStatusCode(v int32) *RegisterMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterMessageResponse) SetBody(v *RegisterMessageResponseBody) *RegisterMessageResponse {
	s.Body = v
	return s
}

type TrustEventsRequest struct {
	EventUuids  *string `json:"EventUuids,omitempty" xml:"EventUuids,omitempty"`
	HandleStyle *int32  `json:"HandleStyle,omitempty" xml:"HandleStyle,omitempty"`
}

func (s TrustEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s TrustEventsRequest) GoString() string {
	return s.String()
}

func (s *TrustEventsRequest) SetEventUuids(v string) *TrustEventsRequest {
	s.EventUuids = &v
	return s
}

func (s *TrustEventsRequest) SetHandleStyle(v int32) *TrustEventsRequest {
	s.HandleStyle = &v
	return s
}

type TrustEventsResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TrustEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrustEventsResponseBody) GoString() string {
	return s.String()
}

func (s *TrustEventsResponseBody) SetData(v string) *TrustEventsResponseBody {
	s.Data = &v
	return s
}

func (s *TrustEventsResponseBody) SetRequestId(v string) *TrustEventsResponseBody {
	s.RequestId = &v
	return s
}

type TrustEventsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TrustEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TrustEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s TrustEventsResponse) GoString() string {
	return s.String()
}

func (s *TrustEventsResponse) SetHeaders(v map[string]*string) *TrustEventsResponse {
	s.Headers = v
	return s
}

func (s *TrustEventsResponse) SetStatusCode(v int32) *TrustEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *TrustEventsResponse) SetBody(v *TrustEventsResponseBody) *TrustEventsResponse {
	s.Body = v
	return s
}

type UnregisterMessageRequest struct {
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s UnregisterMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s UnregisterMessageRequest) GoString() string {
	return s.String()
}

func (s *UnregisterMessageRequest) SetProperties(v string) *UnregisterMessageRequest {
	s.Properties = &v
	return s
}

type UnregisterMessageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnregisterMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnregisterMessageResponseBody) GoString() string {
	return s.String()
}

func (s *UnregisterMessageResponseBody) SetRequestId(v string) *UnregisterMessageResponseBody {
	s.RequestId = &v
	return s
}

type UnregisterMessageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnregisterMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnregisterMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s UnregisterMessageResponse) GoString() string {
	return s.String()
}

func (s *UnregisterMessageResponse) SetHeaders(v map[string]*string) *UnregisterMessageResponse {
	s.Headers = v
	return s
}

func (s *UnregisterMessageResponse) SetStatusCode(v int32) *UnregisterMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *UnregisterMessageResponse) SetBody(v *UnregisterMessageResponseBody) *UnregisterMessageResponse {
	s.Body = v
	return s
}

type VerifyMessageRequest struct {
	FileData *string `json:"FileData,omitempty" xml:"FileData,omitempty"`
}

func (s VerifyMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyMessageRequest) GoString() string {
	return s.String()
}

func (s *VerifyMessageRequest) SetFileData(v string) *VerifyMessageRequest {
	s.FileData = &v
	return s
}

type VerifyMessageResponseBody struct {
	Data      *VerifyMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyMessageResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMessageResponseBody) SetData(v *VerifyMessageResponseBodyData) *VerifyMessageResponseBody {
	s.Data = v
	return s
}

func (s *VerifyMessageResponseBody) SetRequestId(v string) *VerifyMessageResponseBody {
	s.RequestId = &v
	return s
}

type VerifyMessageResponseBodyData struct {
	ImaVerificationResult *VerifyMessageResponseBodyDataImaVerificationResult `json:"ImaVerificationResult,omitempty" xml:"ImaVerificationResult,omitempty" type:"Struct"`
	PcrVerificationResult *VerifyMessageResponseBodyDataPcrVerificationResult `json:"PcrVerificationResult,omitempty" xml:"PcrVerificationResult,omitempty" type:"Struct"`
}

func (s VerifyMessageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VerifyMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyMessageResponseBodyData) SetImaVerificationResult(v *VerifyMessageResponseBodyDataImaVerificationResult) *VerifyMessageResponseBodyData {
	s.ImaVerificationResult = v
	return s
}

func (s *VerifyMessageResponseBodyData) SetPcrVerificationResult(v *VerifyMessageResponseBodyDataPcrVerificationResult) *VerifyMessageResponseBodyData {
	s.PcrVerificationResult = v
	return s
}

type VerifyMessageResponseBodyDataImaVerificationResult struct {
	Code        *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Status      *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	TcbObsolete []*string `json:"TcbObsolete,omitempty" xml:"TcbObsolete,omitempty" type:"Repeated"`
	Untrusted   []*string `json:"Untrusted,omitempty" xml:"Untrusted,omitempty" type:"Repeated"`
}

func (s VerifyMessageResponseBodyDataImaVerificationResult) String() string {
	return tea.Prettify(s)
}

func (s VerifyMessageResponseBodyDataImaVerificationResult) GoString() string {
	return s.String()
}

func (s *VerifyMessageResponseBodyDataImaVerificationResult) SetCode(v string) *VerifyMessageResponseBodyDataImaVerificationResult {
	s.Code = &v
	return s
}

func (s *VerifyMessageResponseBodyDataImaVerificationResult) SetStatus(v int32) *VerifyMessageResponseBodyDataImaVerificationResult {
	s.Status = &v
	return s
}

func (s *VerifyMessageResponseBodyDataImaVerificationResult) SetTcbObsolete(v []*string) *VerifyMessageResponseBodyDataImaVerificationResult {
	s.TcbObsolete = v
	return s
}

func (s *VerifyMessageResponseBodyDataImaVerificationResult) SetUntrusted(v []*string) *VerifyMessageResponseBodyDataImaVerificationResult {
	s.Untrusted = v
	return s
}

type VerifyMessageResponseBodyDataPcrVerificationResult struct {
	Code        *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Status      *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	TcbObsolete []*string `json:"TcbObsolete,omitempty" xml:"TcbObsolete,omitempty" type:"Repeated"`
	Untrusted   []*string `json:"Untrusted,omitempty" xml:"Untrusted,omitempty" type:"Repeated"`
}

func (s VerifyMessageResponseBodyDataPcrVerificationResult) String() string {
	return tea.Prettify(s)
}

func (s VerifyMessageResponseBodyDataPcrVerificationResult) GoString() string {
	return s.String()
}

func (s *VerifyMessageResponseBodyDataPcrVerificationResult) SetCode(v string) *VerifyMessageResponseBodyDataPcrVerificationResult {
	s.Code = &v
	return s
}

func (s *VerifyMessageResponseBodyDataPcrVerificationResult) SetStatus(v int32) *VerifyMessageResponseBodyDataPcrVerificationResult {
	s.Status = &v
	return s
}

func (s *VerifyMessageResponseBodyDataPcrVerificationResult) SetTcbObsolete(v []*string) *VerifyMessageResponseBodyDataPcrVerificationResult {
	s.TcbObsolete = v
	return s
}

func (s *VerifyMessageResponseBodyDataPcrVerificationResult) SetUntrusted(v []*string) *VerifyMessageResponseBodyDataPcrVerificationResult {
	s.Untrusted = v
	return s
}

type VerifyMessageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyMessageResponse) GoString() string {
	return s.String()
}

func (s *VerifyMessageResponse) SetHeaders(v map[string]*string) *VerifyMessageResponse {
	s.Headers = v
	return s
}

func (s *VerifyMessageResponse) SetStatusCode(v int32) *VerifyMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyMessageResponse) SetBody(v *VerifyMessageResponseBody) *VerifyMessageResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("trusted-server"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeBootWithOptions(request *DescribeBootRequest, runtime *util.RuntimeOptions) (_result *DescribeBootResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyUuid)) {
		query["PropertyUuid"] = request.PropertyUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBoot"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBootResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBoot(request *DescribeBootRequest) (_result *DescribeBootResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBootResponse{}
	_body, _err := client.DescribeBootWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateEndDate)) {
		query["CreateEndDate"] = request.CreateEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.CreateStartDate)) {
		query["CreateStartDate"] = request.CreateStartDate
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EventAffiliation)) {
		query["EventAffiliation"] = request.EventAffiliation
	}

	if !tea.BoolValue(util.IsUnset(request.EventLevel)) {
		query["EventLevel"] = request.EventLevel
	}

	if !tea.BoolValue(util.IsUnset(request.EventStatus)) {
		query["EventStatus"] = request.EventStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.HandleEndDate)) {
		query["HandleEndDate"] = request.HandleEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.HandleStartDate)) {
		query["HandleStartDate"] = request.HandleStartDate
	}

	if !tea.BoolValue(util.IsUnset(request.HandleType)) {
		query["HandleType"] = request.HandleType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyName)) {
		query["PropertyName"] = request.PropertyName
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyPrivateIp)) {
		query["PropertyPrivateIp"] = request.PropertyPrivateIp
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyPublicIp)) {
		query["PropertyPublicIp"] = request.PropertyPublicIp
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyUuid)) {
		query["PropertyUuid"] = request.PropertyUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Suspect)) {
		query["Suspect"] = request.Suspect
	}

	if !tea.BoolValue(util.IsUnset(request.SuspectSig)) {
		query["SuspectSig"] = request.SuspectSig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEvents"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyUuid)) {
		query["PropertyUuid"] = request.PropertyUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateAikcertWithOptions(request *GenerateAikcertRequest, runtime *util.RuntimeOptions) (_result *GenerateAikcertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AikName)) {
		query["AikName"] = request.AikName
	}

	if !tea.BoolValue(util.IsUnset(request.AikPub)) {
		query["AikPub"] = request.AikPub
	}

	if !tea.BoolValue(util.IsUnset(request.EkCert)) {
		query["EkCert"] = request.EkCert
	}

	if !tea.BoolValue(util.IsUnset(request.EkPub)) {
		query["EkPub"] = request.EkPub
	}

	if !tea.BoolValue(util.IsUnset(request.NonceDigest)) {
		query["NonceDigest"] = request.NonceDigest
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyUuid)) {
		query["PropertyUuid"] = request.PropertyUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateAikcert"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateAikcertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateAikcert(request *GenerateAikcertRequest) (_result *GenerateAikcertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateAikcertResponse{}
	_body, _err := client.GenerateAikcertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateNonceWithOptions(request *GenerateNonceRequest, runtime *util.RuntimeOptions) (_result *GenerateNonceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AikName)) {
		query["AikName"] = request.AikName
	}

	if !tea.BoolValue(util.IsUnset(request.EkCert)) {
		query["EkCert"] = request.EkCert
	}

	if !tea.BoolValue(util.IsUnset(request.EkPub)) {
		query["EkPub"] = request.EkPub
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyUuid)) {
		query["PropertyUuid"] = request.PropertyUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateNonce"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateNonceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateNonce(request *GenerateNonceRequest) (_result *GenerateNonceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateNonceResponse{}
	_body, _err := client.GenerateNonceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IgnoreEventsWithOptions(request *IgnoreEventsRequest, runtime *util.RuntimeOptions) (_result *IgnoreEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventUuids)) {
		query["EventUuids"] = request.EventUuids
	}

	if !tea.BoolValue(util.IsUnset(request.HandleStyle)) {
		query["HandleStyle"] = request.HandleStyle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IgnoreEvents"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IgnoreEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IgnoreEvents(request *IgnoreEventsRequest) (_result *IgnoreEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IgnoreEventsResponse{}
	_body, _err := client.IgnoreEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProduceAikcertWithOptions(request *ProduceAikcertRequest, runtime *util.RuntimeOptions) (_result *ProduceAikcertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AikName)) {
		query["AikName"] = request.AikName
	}

	if !tea.BoolValue(util.IsUnset(request.CertRequest)) {
		query["CertRequest"] = request.CertRequest
	}

	if !tea.BoolValue(util.IsUnset(request.EkCert)) {
		query["EkCert"] = request.EkCert
	}

	if !tea.BoolValue(util.IsUnset(request.EkPubKey)) {
		query["EkPubKey"] = request.EkPubKey
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeCertChain)) {
		query["IncludeCertChain"] = request.IncludeCertChain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ProduceAikcert"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ProduceAikcertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProduceAikcert(request *ProduceAikcertRequest) (_result *ProduceAikcertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProduceAikcertResponse{}
	_body, _err := client.ProduceAikcertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMessageWithOptions(request *PutMessageRequest, runtime *util.RuntimeOptions) (_result *PutMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileData)) {
		query["FileData"] = request.FileData
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyUuid)) {
		query["PropertyUuid"] = request.PropertyUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutMessage"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMessage(request *PutMessageRequest) (_result *PutMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMessageResponse{}
	_body, _err := client.PutMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterMessageWithOptions(request *RegisterMessageRequest, runtime *util.RuntimeOptions) (_result *RegisterMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extensions)) {
		query["Extensions"] = request.Extensions
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyAffiliation)) {
		query["PropertyAffiliation"] = request.PropertyAffiliation
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyName)) {
		query["PropertyName"] = request.PropertyName
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyPrivateIp)) {
		query["PropertyPrivateIp"] = request.PropertyPrivateIp
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyPublicIp)) {
		query["PropertyPublicIp"] = request.PropertyPublicIp
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyUuid)) {
		query["PropertyUuid"] = request.PropertyUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterMessage"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterMessage(request *RegisterMessageRequest) (_result *RegisterMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterMessageResponse{}
	_body, _err := client.RegisterMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TrustEventsWithOptions(request *TrustEventsRequest, runtime *util.RuntimeOptions) (_result *TrustEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventUuids)) {
		query["EventUuids"] = request.EventUuids
	}

	if !tea.BoolValue(util.IsUnset(request.HandleStyle)) {
		query["HandleStyle"] = request.HandleStyle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TrustEvents"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TrustEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TrustEvents(request *TrustEventsRequest) (_result *TrustEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TrustEventsResponse{}
	_body, _err := client.TrustEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnregisterMessageWithOptions(request *UnregisterMessageRequest, runtime *util.RuntimeOptions) (_result *UnregisterMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		query["Properties"] = request.Properties
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnregisterMessage"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnregisterMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnregisterMessage(request *UnregisterMessageRequest) (_result *UnregisterMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnregisterMessageResponse{}
	_body, _err := client.UnregisterMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyMessageWithOptions(request *VerifyMessageRequest, runtime *util.RuntimeOptions) (_result *VerifyMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileData)) {
		body["FileData"] = request.FileData
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyMessage"),
		Version:     tea.String("2020-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyMessage(request *VerifyMessageRequest) (_result *VerifyMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyMessageResponse{}
	_body, _err := client.VerifyMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
