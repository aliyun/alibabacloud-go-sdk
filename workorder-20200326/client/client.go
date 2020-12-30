// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CloseTicketRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s CloseTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketRequest) GoString() string {
	return s.String()
}

func (s *CloseTicketRequest) SetLanguage(v string) *CloseTicketRequest {
	s.Language = &v
	return s
}

func (s *CloseTicketRequest) SetTicketId(v string) *CloseTicketRequest {
	s.TicketId = &v
	return s
}

type CloseTicketResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTicketResponseBody) SetMessage(v string) *CloseTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CloseTicketResponseBody) SetRequestId(v string) *CloseTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTicketResponseBody) SetCode(v int32) *CloseTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CloseTicketResponseBody) SetSuccess(v bool) *CloseTicketResponseBody {
	s.Success = &v
	return s
}

type CloseTicketResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketResponse) GoString() string {
	return s.String()
}

func (s *CloseTicketResponse) SetHeaders(v map[string]*string) *CloseTicketResponse {
	s.Headers = v
	return s
}

func (s *CloseTicketResponse) SetBody(v *CloseTicketResponseBody) *CloseTicketResponse {
	s.Body = v
	return s
}

type CreateTicketRequest struct {
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SecretContent   *string `json:"SecretContent,omitempty" xml:"SecretContent,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Phone           *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	NotifyTimeRange *string `json:"NotifyTimeRange,omitempty" xml:"NotifyTimeRange,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetLanguage(v string) *CreateTicketRequest {
	s.Language = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketRequest) SetContent(v string) *CreateTicketRequest {
	s.Content = &v
	return s
}

func (s *CreateTicketRequest) SetSecretContent(v string) *CreateTicketRequest {
	s.SecretContent = &v
	return s
}

func (s *CreateTicketRequest) SetProductCode(v string) *CreateTicketRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateTicketRequest) SetCategory(v string) *CreateTicketRequest {
	s.Category = &v
	return s
}

func (s *CreateTicketRequest) SetPhone(v string) *CreateTicketRequest {
	s.Phone = &v
	return s
}

func (s *CreateTicketRequest) SetEmail(v string) *CreateTicketRequest {
	s.Email = &v
	return s
}

func (s *CreateTicketRequest) SetNotifyTimeRange(v string) *CreateTicketRequest {
	s.NotifyTimeRange = &v
	return s
}

type CreateTicketResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetMessage(v string) *CreateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetData(v string) *CreateTicketResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTicketResponseBody) SetCode(v string) *CreateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTicketResponseBody) SetSuccess(v bool) *CreateTicketResponseBody {
	s.Success = &v
	return s
}

type CreateTicketResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type ListCategoriesRequest struct {
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListCategoriesRequest) SetLanguage(v string) *ListCategoriesRequest {
	s.Language = &v
	return s
}

func (s *ListCategoriesRequest) SetProductCode(v string) *ListCategoriesRequest {
	s.ProductCode = &v
	return s
}

type ListCategoriesResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListCategoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBody) SetMessage(v string) *ListCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoriesResponseBody) SetRequestId(v string) *ListCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoriesResponseBody) SetData(v *ListCategoriesResponseBodyData) *ListCategoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListCategoriesResponseBody) SetCode(v int32) *ListCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCategoriesResponseBody) SetSuccess(v bool) *ListCategoriesResponseBody {
	s.Success = &v
	return s
}

type ListCategoriesResponseBodyData struct {
	List []*ListCategoriesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListCategoriesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBodyData) SetList(v []*ListCategoriesResponseBodyDataList) *ListCategoriesResponseBodyData {
	s.List = v
	return s
}

type ListCategoriesResponseBodyDataList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListCategoriesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBodyDataList) SetName(v string) *ListCategoriesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListCategoriesResponseBodyDataList) SetId(v int32) *ListCategoriesResponseBodyDataList {
	s.Id = &v
	return s
}

type ListCategoriesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponse) SetHeaders(v map[string]*string) *ListCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListCategoriesResponse) SetBody(v *ListCategoriesResponseBody) *ListCategoriesResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetLanguage(v string) *ListProductsRequest {
	s.Language = &v
	return s
}

type ListProductsResponseBody struct {
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListProductsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetMessage(v string) *ListProductsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetData(v *ListProductsResponseBodyData) *ListProductsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductsResponseBody) SetCode(v int32) *ListProductsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductsResponseBody) SetSuccess(v bool) *ListProductsResponseBody {
	s.Success = &v
	return s
}

type ListProductsResponseBodyData struct {
	HotConsultation  []*ListProductsResponseBodyDataHotConsultation  `json:"HotConsultation,omitempty" xml:"HotConsultation,omitempty" type:"Repeated"`
	ConsultationMore []*ListProductsResponseBodyDataConsultationMore `json:"ConsultationMore,omitempty" xml:"ConsultationMore,omitempty" type:"Repeated"`
	HotTech          []*ListProductsResponseBodyDataHotTech          `json:"HotTech,omitempty" xml:"HotTech,omitempty" type:"Repeated"`
	TechMore         []*ListProductsResponseBodyDataTechMore         `json:"TechMore,omitempty" xml:"TechMore,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyData) SetHotConsultation(v []*ListProductsResponseBodyDataHotConsultation) *ListProductsResponseBodyData {
	s.HotConsultation = v
	return s
}

func (s *ListProductsResponseBodyData) SetConsultationMore(v []*ListProductsResponseBodyDataConsultationMore) *ListProductsResponseBodyData {
	s.ConsultationMore = v
	return s
}

func (s *ListProductsResponseBodyData) SetHotTech(v []*ListProductsResponseBodyDataHotTech) *ListProductsResponseBodyData {
	s.HotTech = v
	return s
}

func (s *ListProductsResponseBodyData) SetTechMore(v []*ListProductsResponseBodyDataTechMore) *ListProductsResponseBodyData {
	s.TechMore = v
	return s
}

type ListProductsResponseBodyDataHotConsultation struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyDataHotConsultation) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyDataHotConsultation) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataHotConsultation) SetDescription(v string) *ListProductsResponseBodyDataHotConsultation {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataHotConsultation) SetName(v string) *ListProductsResponseBodyDataHotConsultation {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataHotConsultation) SetProductCode(v string) *ListProductsResponseBodyDataHotConsultation {
	s.ProductCode = &v
	return s
}

type ListProductsResponseBodyDataConsultationMore struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyDataConsultationMore) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyDataConsultationMore) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataConsultationMore) SetDescription(v string) *ListProductsResponseBodyDataConsultationMore {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataConsultationMore) SetName(v string) *ListProductsResponseBodyDataConsultationMore {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataConsultationMore) SetProductCode(v string) *ListProductsResponseBodyDataConsultationMore {
	s.ProductCode = &v
	return s
}

type ListProductsResponseBodyDataHotTech struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyDataHotTech) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyDataHotTech) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataHotTech) SetDescription(v string) *ListProductsResponseBodyDataHotTech {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataHotTech) SetName(v string) *ListProductsResponseBodyDataHotTech {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataHotTech) SetProductCode(v string) *ListProductsResponseBodyDataHotTech {
	s.ProductCode = &v
	return s
}

type ListProductsResponseBodyDataTechMore struct {
	GroupName   *string                                            `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProductList []*ListProductsResponseBodyDataTechMoreProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyDataTechMore) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyDataTechMore) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataTechMore) SetGroupName(v string) *ListProductsResponseBodyDataTechMore {
	s.GroupName = &v
	return s
}

func (s *ListProductsResponseBodyDataTechMore) SetProductList(v []*ListProductsResponseBodyDataTechMoreProductList) *ListProductsResponseBodyDataTechMore {
	s.ProductList = v
	return s
}

type ListProductsResponseBodyDataTechMoreProductList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyDataTechMoreProductList) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyDataTechMoreProductList) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataTechMoreProductList) SetDescription(v string) *ListProductsResponseBodyDataTechMoreProductList {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataTechMoreProductList) SetName(v string) *ListProductsResponseBodyDataTechMoreProductList {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataTechMoreProductList) SetProductCode(v string) *ListProductsResponseBodyDataTechMoreProductList {
	s.ProductCode = &v
	return s
}

type ListProductsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProductsResponse) SetHeaders(v map[string]*string) *ListProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type ListTicketNotesRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ListTicketNotesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesRequest) GoString() string {
	return s.String()
}

func (s *ListTicketNotesRequest) SetLanguage(v string) *ListTicketNotesRequest {
	s.Language = &v
	return s
}

func (s *ListTicketNotesRequest) SetTicketId(v string) *ListTicketNotesRequest {
	s.TicketId = &v
	return s
}

type ListTicketNotesResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListTicketNotesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTicketNotesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBody) SetMessage(v string) *ListTicketNotesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetRequestId(v string) *ListTicketNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetData(v *ListTicketNotesResponseBodyData) *ListTicketNotesResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketNotesResponseBody) SetCode(v int32) *ListTicketNotesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetSuccess(v bool) *ListTicketNotesResponseBody {
	s.Success = &v
	return s
}

type ListTicketNotesResponseBodyData struct {
	List []*ListTicketNotesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListTicketNotesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyData) SetList(v []*ListTicketNotesResponseBodyDataList) *ListTicketNotesResponseBodyData {
	s.List = v
	return s
}

type ListTicketNotesResponseBodyDataList struct {
	GmtCreated   *int32  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	NoteId       *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	FromOfficial *bool   `json:"FromOfficial,omitempty" xml:"FromOfficial,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ListTicketNotesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyDataList) SetGmtCreated(v int32) *ListTicketNotesResponseBodyDataList {
	s.GmtCreated = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataList) SetNoteId(v string) *ListTicketNotesResponseBodyDataList {
	s.NoteId = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataList) SetFromOfficial(v bool) *ListTicketNotesResponseBodyDataList {
	s.FromOfficial = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataList) SetContent(v string) *ListTicketNotesResponseBodyDataList {
	s.Content = &v
	return s
}

type ListTicketNotesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTicketNotesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTicketNotesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponse) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponse) SetHeaders(v map[string]*string) *ListTicketNotesResponse {
	s.Headers = v
	return s
}

func (s *ListTicketNotesResponse) SetBody(v *ListTicketNotesResponseBody) *ListTicketNotesResponse {
	s.Body = v
	return s
}

type ListTicketsRequest struct {
	Ids               *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	CreatedAfterTime  *int64  `json:"CreatedAfterTime,omitempty" xml:"CreatedAfterTime,omitempty"`
	CreatedBeforeTime *int64  `json:"CreatedBeforeTime,omitempty" xml:"CreatedBeforeTime,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TicketStatus      *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
	PageStart         *int32  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	SubUserId         *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	Language          *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsRequest) SetIds(v string) *ListTicketsRequest {
	s.Ids = &v
	return s
}

func (s *ListTicketsRequest) SetCreatedAfterTime(v int64) *ListTicketsRequest {
	s.CreatedAfterTime = &v
	return s
}

func (s *ListTicketsRequest) SetCreatedBeforeTime(v int64) *ListTicketsRequest {
	s.CreatedBeforeTime = &v
	return s
}

func (s *ListTicketsRequest) SetProductCode(v string) *ListTicketsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListTicketsRequest) SetPageSize(v int32) *ListTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketsRequest) SetTicketStatus(v string) *ListTicketsRequest {
	s.TicketStatus = &v
	return s
}

func (s *ListTicketsRequest) SetPageStart(v int32) *ListTicketsRequest {
	s.PageStart = &v
	return s
}

func (s *ListTicketsRequest) SetSubUserId(v string) *ListTicketsRequest {
	s.SubUserId = &v
	return s
}

func (s *ListTicketsRequest) SetLanguage(v string) *ListTicketsRequest {
	s.Language = &v
	return s
}

type ListTicketsResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListTicketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBody) SetMessage(v string) *ListTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketsResponseBody) SetRequestId(v string) *ListTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketsResponseBody) SetData(v *ListTicketsResponseBodyData) *ListTicketsResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketsResponseBody) SetCode(v int32) *ListTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketsResponseBody) SetSuccess(v bool) *ListTicketsResponseBody {
	s.Success = &v
	return s
}

type ListTicketsResponseBodyData struct {
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*ListTicketsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTicketsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyData) SetCurrentPage(v int32) *ListTicketsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListTicketsResponseBodyData) SetList(v []*ListTicketsResponseBodyDataList) *ListTicketsResponseBodyData {
	s.List = v
	return s
}

func (s *ListTicketsResponseBodyData) SetPageSize(v int32) *ListTicketsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTicketsResponseBodyData) SetTotal(v int32) *ListTicketsResponseBodyData {
	s.Total = &v
	return s
}

type ListTicketsResponseBodyDataList struct {
	TicketStatus *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	CreatorId    *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	AddTime      *int32  `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListTicketsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyDataList) SetTicketStatus(v string) *ListTicketsResponseBodyDataList {
	s.TicketStatus = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetTitle(v string) *ListTicketsResponseBodyDataList {
	s.Title = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCreatorId(v string) *ListTicketsResponseBodyDataList {
	s.CreatorId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetAddTime(v int32) *ListTicketsResponseBodyDataList {
	s.AddTime = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetId(v string) *ListTicketsResponseBodyDataList {
	s.Id = &v
	return s
}

type ListTicketsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponse) GoString() string {
	return s.String()
}

func (s *ListTicketsResponse) SetHeaders(v map[string]*string) *ListTicketsResponse {
	s.Headers = v
	return s
}

func (s *ListTicketsResponse) SetBody(v *ListTicketsResponseBody) *ListTicketsResponse {
	s.Body = v
	return s
}

type ReplyTicketRequest struct {
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	TicketId      *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SecretContent *string `json:"SecretContent,omitempty" xml:"SecretContent,omitempty"`
}

func (s ReplyTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplyTicketRequest) GoString() string {
	return s.String()
}

func (s *ReplyTicketRequest) SetLanguage(v string) *ReplyTicketRequest {
	s.Language = &v
	return s
}

func (s *ReplyTicketRequest) SetTicketId(v string) *ReplyTicketRequest {
	s.TicketId = &v
	return s
}

func (s *ReplyTicketRequest) SetContent(v string) *ReplyTicketRequest {
	s.Content = &v
	return s
}

func (s *ReplyTicketRequest) SetSecretContent(v string) *ReplyTicketRequest {
	s.SecretContent = &v
	return s
}

type ReplyTicketResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReplyTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplyTicketResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyTicketResponseBody) SetMessage(v string) *ReplyTicketResponseBody {
	s.Message = &v
	return s
}

func (s *ReplyTicketResponseBody) SetRequestId(v string) *ReplyTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplyTicketResponseBody) SetCode(v int32) *ReplyTicketResponseBody {
	s.Code = &v
	return s
}

func (s *ReplyTicketResponseBody) SetSuccess(v bool) *ReplyTicketResponseBody {
	s.Success = &v
	return s
}

type ReplyTicketResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplyTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplyTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplyTicketResponse) GoString() string {
	return s.String()
}

func (s *ReplyTicketResponse) SetHeaders(v map[string]*string) *ReplyTicketResponse {
	s.Headers = v
	return s
}

func (s *ReplyTicketResponse) SetBody(v *ReplyTicketResponseBody) *ReplyTicketResponse {
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
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1": tea.String("workorder.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1": tea.String("workorder.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("workorder"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CloseTicketWithOptions(request *CloseTicketRequest, runtime *util.RuntimeOptions) (_result *CloseTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloseTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloseTicket"), tea.String("2020-03-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseTicket(request *CloseTicketRequest) (_result *CloseTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseTicketResponse{}
	_body, _err := client.CloseTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTicket"), tea.String("2020-03-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCategoriesWithOptions(request *ListCategoriesRequest, runtime *util.RuntimeOptions) (_result *ListCategoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCategoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCategories"), tea.String("2020-03-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCategories(request *ListCategoriesRequest) (_result *ListCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCategoriesResponse{}
	_body, _err := client.ListCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductsWithOptions(request *ListProductsRequest, runtime *util.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProducts"), tea.String("2020-03-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTicketNotesWithOptions(request *ListTicketNotesRequest, runtime *util.RuntimeOptions) (_result *ListTicketNotesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTicketNotesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTicketNotes"), tea.String("2020-03-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTicketNotes(request *ListTicketNotesRequest) (_result *ListTicketNotesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTicketNotesResponse{}
	_body, _err := client.ListTicketNotesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTicketsWithOptions(request *ListTicketsRequest, runtime *util.RuntimeOptions) (_result *ListTicketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTicketsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTickets"), tea.String("2020-03-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTickets(request *ListTicketsRequest) (_result *ListTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTicketsResponse{}
	_body, _err := client.ListTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplyTicketWithOptions(request *ReplyTicketRequest, runtime *util.RuntimeOptions) (_result *ReplyTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReplyTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReplyTicket"), tea.String("2020-03-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplyTicket(request *ReplyTicketRequest) (_result *ReplyTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplyTicketResponse{}
	_body, _err := client.ReplyTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
