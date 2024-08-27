// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateDeviceRequest struct {
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ActivateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateDeviceRequest) GoString() string {
	return s.String()
}

func (s *ActivateDeviceRequest) SetUuid(v string) *ActivateDeviceRequest {
	s.Uuid = &v
	return s
}

type ActivateDeviceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateDeviceResponseBody) SetCode(v string) *ActivateDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateDeviceResponseBody) SetMessage(v string) *ActivateDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateDeviceResponseBody) SetRequestId(v string) *ActivateDeviceResponseBody {
	s.RequestId = &v
	return s
}

type ActivateDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateDeviceResponse) GoString() string {
	return s.String()
}

func (s *ActivateDeviceResponse) SetHeaders(v map[string]*string) *ActivateDeviceResponse {
	s.Headers = v
	return s
}

func (s *ActivateDeviceResponse) SetStatusCode(v int32) *ActivateDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateDeviceResponse) SetBody(v *ActivateDeviceResponseBody) *ActivateDeviceResponse {
	s.Body = v
	return s
}

type AddDeviceFromSNRequest struct {
	Alias             *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	CustomProperty    *string `json:"CustomProperty,omitempty" xml:"CustomProperty,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	LabelContents     *string `json:"LabelContents,omitempty" xml:"LabelContents,omitempty"`
	SecureNetworkType *string `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AddDeviceFromSNRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceFromSNRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceFromSNRequest) SetAlias(v string) *AddDeviceFromSNRequest {
	s.Alias = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetCustomProperty(v string) *AddDeviceFromSNRequest {
	s.CustomProperty = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetGroupId(v string) *AddDeviceFromSNRequest {
	s.GroupId = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetLabelContents(v string) *AddDeviceFromSNRequest {
	s.LabelContents = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetSecureNetworkType(v string) *AddDeviceFromSNRequest {
	s.SecureNetworkType = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetSerialNo(v string) *AddDeviceFromSNRequest {
	s.SerialNo = &v
	return s
}

type AddDeviceFromSNResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDeviceFromSNResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceFromSNResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceFromSNResponseBody) SetCode(v string) *AddDeviceFromSNResponseBody {
	s.Code = &v
	return s
}

func (s *AddDeviceFromSNResponseBody) SetMessage(v string) *AddDeviceFromSNResponseBody {
	s.Message = &v
	return s
}

func (s *AddDeviceFromSNResponseBody) SetRequestId(v string) *AddDeviceFromSNResponseBody {
	s.RequestId = &v
	return s
}

type AddDeviceFromSNResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDeviceFromSNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDeviceFromSNResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceFromSNResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceFromSNResponse) SetHeaders(v map[string]*string) *AddDeviceFromSNResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceFromSNResponse) SetStatusCode(v int32) *AddDeviceFromSNResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDeviceFromSNResponse) SetBody(v *AddDeviceFromSNResponseBody) *AddDeviceFromSNResponse {
	s.Body = v
	return s
}

type AddDeviceSeatsAndLabelsRequest struct {
	IsUnique  *bool     `json:"IsUnique,omitempty" xml:"IsUnique,omitempty"`
	Label     *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	LabelList []*string `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	SeatName  *string   `json:"SeatName,omitempty" xml:"SeatName,omitempty"`
	SerialNo  *string   `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	TenantId  *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId    *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddDeviceSeatsAndLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceSeatsAndLabelsRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceSeatsAndLabelsRequest) SetIsUnique(v bool) *AddDeviceSeatsAndLabelsRequest {
	s.IsUnique = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetLabel(v string) *AddDeviceSeatsAndLabelsRequest {
	s.Label = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetLabelList(v []*string) *AddDeviceSeatsAndLabelsRequest {
	s.LabelList = v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetSeatName(v string) *AddDeviceSeatsAndLabelsRequest {
	s.SeatName = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetSerialNo(v string) *AddDeviceSeatsAndLabelsRequest {
	s.SerialNo = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetTenantId(v string) *AddDeviceSeatsAndLabelsRequest {
	s.TenantId = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetZoneId(v string) *AddDeviceSeatsAndLabelsRequest {
	s.ZoneId = &v
	return s
}

type AddDeviceSeatsAndLabelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDeviceSeatsAndLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceSeatsAndLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceSeatsAndLabelsResponseBody) SetCode(v string) *AddDeviceSeatsAndLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponseBody) SetMessage(v string) *AddDeviceSeatsAndLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponseBody) SetRequestId(v string) *AddDeviceSeatsAndLabelsResponseBody {
	s.RequestId = &v
	return s
}

type AddDeviceSeatsAndLabelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDeviceSeatsAndLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDeviceSeatsAndLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceSeatsAndLabelsResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceSeatsAndLabelsResponse) SetHeaders(v map[string]*string) *AddDeviceSeatsAndLabelsResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponse) SetStatusCode(v int32) *AddDeviceSeatsAndLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponse) SetBody(v *AddDeviceSeatsAndLabelsResponseBody) *AddDeviceSeatsAndLabelsResponse {
	s.Body = v
	return s
}

type AddDevicesFromCSVRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileType *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	SeatCol  *int32  `json:"SeatCol,omitempty" xml:"SeatCol,omitempty"`
	SiteId   *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s AddDevicesFromCSVRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDevicesFromCSVRequest) GoString() string {
	return s.String()
}

func (s *AddDevicesFromCSVRequest) SetFileName(v string) *AddDevicesFromCSVRequest {
	s.FileName = &v
	return s
}

func (s *AddDevicesFromCSVRequest) SetFileType(v int32) *AddDevicesFromCSVRequest {
	s.FileType = &v
	return s
}

func (s *AddDevicesFromCSVRequest) SetSeatCol(v int32) *AddDevicesFromCSVRequest {
	s.SeatCol = &v
	return s
}

func (s *AddDevicesFromCSVRequest) SetSiteId(v string) *AddDevicesFromCSVRequest {
	s.SiteId = &v
	return s
}

func (s *AddDevicesFromCSVRequest) SetSiteName(v string) *AddDevicesFromCSVRequest {
	s.SiteName = &v
	return s
}

type AddDevicesFromCSVResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDevicesFromCSVResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDevicesFromCSVResponseBody) GoString() string {
	return s.String()
}

func (s *AddDevicesFromCSVResponseBody) SetCode(v string) *AddDevicesFromCSVResponseBody {
	s.Code = &v
	return s
}

func (s *AddDevicesFromCSVResponseBody) SetMessage(v string) *AddDevicesFromCSVResponseBody {
	s.Message = &v
	return s
}

func (s *AddDevicesFromCSVResponseBody) SetRequestId(v string) *AddDevicesFromCSVResponseBody {
	s.RequestId = &v
	return s
}

type AddDevicesFromCSVResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDevicesFromCSVResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDevicesFromCSVResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDevicesFromCSVResponse) GoString() string {
	return s.String()
}

func (s *AddDevicesFromCSVResponse) SetHeaders(v map[string]*string) *AddDevicesFromCSVResponse {
	s.Headers = v
	return s
}

func (s *AddDevicesFromCSVResponse) SetStatusCode(v int32) *AddDevicesFromCSVResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDevicesFromCSVResponse) SetBody(v *AddDevicesFromCSVResponseBody) *AddDevicesFromCSVResponse {
	s.Body = v
	return s
}

type AddLabelsRequest struct {
	LabelContents *string `json:"LabelContents,omitempty" xml:"LabelContents,omitempty"`
}

func (s AddLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLabelsRequest) GoString() string {
	return s.String()
}

func (s *AddLabelsRequest) SetLabelContents(v string) *AddLabelsRequest {
	s.LabelContents = &v
	return s
}

type AddLabelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AddLabelsResponseBody) SetCode(v string) *AddLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *AddLabelsResponseBody) SetMessage(v string) *AddLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *AddLabelsResponseBody) SetRequestId(v string) *AddLabelsResponseBody {
	s.RequestId = &v
	return s
}

type AddLabelsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLabelsResponse) GoString() string {
	return s.String()
}

func (s *AddLabelsResponse) SetHeaders(v map[string]*string) *AddLabelsResponse {
	s.Headers = v
	return s
}

func (s *AddLabelsResponse) SetStatusCode(v int32) *AddLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLabelsResponse) SetBody(v *AddLabelsResponseBody) *AddLabelsResponse {
	s.Body = v
	return s
}

type AddOrUpdateDeviceSeatsRequest struct {
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	UserCustomId *string `json:"UserCustomId,omitempty" xml:"UserCustomId,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddOrUpdateDeviceSeatsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateDeviceSeatsRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDeviceSeatsRequest) SetFileName(v string) *AddOrUpdateDeviceSeatsRequest {
	s.FileName = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsRequest) SetUserCustomId(v string) *AddOrUpdateDeviceSeatsRequest {
	s.UserCustomId = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsRequest) SetZoneId(v string) *AddOrUpdateDeviceSeatsRequest {
	s.ZoneId = &v
	return s
}

type AddOrUpdateDeviceSeatsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddOrUpdateDeviceSeatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateDeviceSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDeviceSeatsResponseBody) SetCode(v string) *AddOrUpdateDeviceSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponseBody) SetMessage(v string) *AddOrUpdateDeviceSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponseBody) SetRequestId(v string) *AddOrUpdateDeviceSeatsResponseBody {
	s.RequestId = &v
	return s
}

type AddOrUpdateDeviceSeatsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateDeviceSeatsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateDeviceSeatsResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDeviceSeatsResponse) SetHeaders(v map[string]*string) *AddOrUpdateDeviceSeatsResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponse) SetStatusCode(v int32) *AddOrUpdateDeviceSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponse) SetBody(v *AddOrUpdateDeviceSeatsResponseBody) *AddOrUpdateDeviceSeatsResponse {
	s.Body = v
	return s
}

type AddTerminalRequest struct {
	Alias           *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	SerialNumber    *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
}

func (s AddTerminalRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTerminalRequest) GoString() string {
	return s.String()
}

func (s *AddTerminalRequest) SetAlias(v string) *AddTerminalRequest {
	s.Alias = &v
	return s
}

func (s *AddTerminalRequest) SetSerialNumber(v string) *AddTerminalRequest {
	s.SerialNumber = &v
	return s
}

func (s *AddTerminalRequest) SetTerminalGroupId(v string) *AddTerminalRequest {
	s.TerminalGroupId = &v
	return s
}

type AddTerminalResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTerminalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *AddTerminalResponseBody) SetCode(v string) *AddTerminalResponseBody {
	s.Code = &v
	return s
}

func (s *AddTerminalResponseBody) SetHttpStatusCode(v int32) *AddTerminalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddTerminalResponseBody) SetMessage(v string) *AddTerminalResponseBody {
	s.Message = &v
	return s
}

func (s *AddTerminalResponseBody) SetRequestId(v string) *AddTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTerminalResponseBody) SetSuccess(v bool) *AddTerminalResponseBody {
	s.Success = &v
	return s
}

type AddTerminalResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTerminalResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTerminalResponse) GoString() string {
	return s.String()
}

func (s *AddTerminalResponse) SetHeaders(v map[string]*string) *AddTerminalResponse {
	s.Headers = v
	return s
}

func (s *AddTerminalResponse) SetStatusCode(v int32) *AddTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTerminalResponse) SetBody(v *AddTerminalResponseBody) *AddTerminalResponse {
	s.Body = v
	return s
}

type AttachEndUsersRequest struct {
	// This parameter is required.
	EndUserIds *string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AttachEndUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachEndUsersRequest) GoString() string {
	return s.String()
}

func (s *AttachEndUsersRequest) SetEndUserIds(v string) *AttachEndUsersRequest {
	s.EndUserIds = &v
	return s
}

func (s *AttachEndUsersRequest) SetSerialNo(v string) *AttachEndUsersRequest {
	s.SerialNo = &v
	return s
}

type AttachEndUsersResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEndUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachEndUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEndUsersResponseBody) SetCode(v string) *AttachEndUsersResponseBody {
	s.Code = &v
	return s
}

func (s *AttachEndUsersResponseBody) SetMessage(v string) *AttachEndUsersResponseBody {
	s.Message = &v
	return s
}

func (s *AttachEndUsersResponseBody) SetRequestId(v string) *AttachEndUsersResponseBody {
	s.RequestId = &v
	return s
}

type AttachEndUsersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEndUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachEndUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachEndUsersResponse) GoString() string {
	return s.String()
}

func (s *AttachEndUsersResponse) SetHeaders(v map[string]*string) *AttachEndUsersResponse {
	s.Headers = v
	return s
}

func (s *AttachEndUsersResponse) SetStatusCode(v int32) *AttachEndUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEndUsersResponse) SetBody(v *AttachEndUsersResponseBody) *AttachEndUsersResponse {
	s.Body = v
	return s
}

type AttachLabelRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AttachLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachLabelRequest) GoString() string {
	return s.String()
}

func (s *AttachLabelRequest) SetLabelContent(v string) *AttachLabelRequest {
	s.LabelContent = &v
	return s
}

func (s *AttachLabelRequest) SetLabelId(v string) *AttachLabelRequest {
	s.LabelId = &v
	return s
}

func (s *AttachLabelRequest) SetSerialNo(v string) *AttachLabelRequest {
	s.SerialNo = &v
	return s
}

type AttachLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachLabelResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLabelResponseBody) SetCode(v string) *AttachLabelResponseBody {
	s.Code = &v
	return s
}

func (s *AttachLabelResponseBody) SetMessage(v string) *AttachLabelResponseBody {
	s.Message = &v
	return s
}

func (s *AttachLabelResponseBody) SetRequestId(v string) *AttachLabelResponseBody {
	s.RequestId = &v
	return s
}

type AttachLabelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachLabelResponse) GoString() string {
	return s.String()
}

func (s *AttachLabelResponse) SetHeaders(v map[string]*string) *AttachLabelResponse {
	s.Headers = v
	return s
}

func (s *AttachLabelResponse) SetStatusCode(v int32) *AttachLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachLabelResponse) SetBody(v *AttachLabelResponseBody) *AttachLabelResponse {
	s.Body = v
	return s
}

type AttachLabelsRequest struct {
	LabelIds     *string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SerialNoList *string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty"`
}

func (s AttachLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachLabelsRequest) GoString() string {
	return s.String()
}

func (s *AttachLabelsRequest) SetLabelIds(v string) *AttachLabelsRequest {
	s.LabelIds = &v
	return s
}

func (s *AttachLabelsRequest) SetSerialNo(v string) *AttachLabelsRequest {
	s.SerialNo = &v
	return s
}

func (s *AttachLabelsRequest) SetSerialNoList(v string) *AttachLabelsRequest {
	s.SerialNoList = &v
	return s
}

type AttachLabelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLabelsResponseBody) SetCode(v string) *AttachLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *AttachLabelsResponseBody) SetMessage(v string) *AttachLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *AttachLabelsResponseBody) SetRequestId(v string) *AttachLabelsResponseBody {
	s.RequestId = &v
	return s
}

type AttachLabelsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachLabelsResponse) GoString() string {
	return s.String()
}

func (s *AttachLabelsResponse) SetHeaders(v map[string]*string) *AttachLabelsResponse {
	s.Headers = v
	return s
}

func (s *AttachLabelsResponse) SetStatusCode(v int32) *AttachLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachLabelsResponse) SetBody(v *AttachLabelsResponseBody) *AttachLabelsResponse {
	s.Body = v
	return s
}

type BindAccountLessLoginUserRequest struct {
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s BindAccountLessLoginUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAccountLessLoginUserRequest) GoString() string {
	return s.String()
}

func (s *BindAccountLessLoginUserRequest) SetEndUserId(v string) *BindAccountLessLoginUserRequest {
	s.EndUserId = &v
	return s
}

func (s *BindAccountLessLoginUserRequest) SetSerialNumber(v string) *BindAccountLessLoginUserRequest {
	s.SerialNumber = &v
	return s
}

func (s *BindAccountLessLoginUserRequest) SetUuid(v string) *BindAccountLessLoginUserRequest {
	s.Uuid = &v
	return s
}

type BindAccountLessLoginUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAccountLessLoginUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAccountLessLoginUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindAccountLessLoginUserResponseBody) SetCode(v string) *BindAccountLessLoginUserResponseBody {
	s.Code = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) SetHttpStatusCode(v int32) *BindAccountLessLoginUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) SetMessage(v string) *BindAccountLessLoginUserResponseBody {
	s.Message = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) SetRequestId(v string) *BindAccountLessLoginUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) SetSuccess(v bool) *BindAccountLessLoginUserResponseBody {
	s.Success = &v
	return s
}

type BindAccountLessLoginUserResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAccountLessLoginUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAccountLessLoginUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAccountLessLoginUserResponse) GoString() string {
	return s.String()
}

func (s *BindAccountLessLoginUserResponse) SetHeaders(v map[string]*string) *BindAccountLessLoginUserResponse {
	s.Headers = v
	return s
}

func (s *BindAccountLessLoginUserResponse) SetStatusCode(v int32) *BindAccountLessLoginUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAccountLessLoginUserResponse) SetBody(v *BindAccountLessLoginUserResponseBody) *BindAccountLessLoginUserResponse {
	s.Body = v
	return s
}

type CheckUuidValidRequest struct {
	Bluetooth *string `json:"Bluetooth,omitempty" xml:"Bluetooth,omitempty"`
	BuildId   *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// This parameter is required.
	ChipId   *string `json:"ChipId,omitempty" xml:"ChipId,omitempty"`
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	EtherMac *string `json:"EtherMac,omitempty" xml:"EtherMac,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	// This parameter is required.
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Wlan *string `json:"Wlan,omitempty" xml:"Wlan,omitempty"`
}

func (s CheckUuidValidRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUuidValidRequest) GoString() string {
	return s.String()
}

func (s *CheckUuidValidRequest) SetBluetooth(v string) *CheckUuidValidRequest {
	s.Bluetooth = &v
	return s
}

func (s *CheckUuidValidRequest) SetBuildId(v string) *CheckUuidValidRequest {
	s.BuildId = &v
	return s
}

func (s *CheckUuidValidRequest) SetChipId(v string) *CheckUuidValidRequest {
	s.ChipId = &v
	return s
}

func (s *CheckUuidValidRequest) SetClientId(v string) *CheckUuidValidRequest {
	s.ClientId = &v
	return s
}

func (s *CheckUuidValidRequest) SetCustomId(v string) *CheckUuidValidRequest {
	s.CustomId = &v
	return s
}

func (s *CheckUuidValidRequest) SetEtherMac(v string) *CheckUuidValidRequest {
	s.EtherMac = &v
	return s
}

func (s *CheckUuidValidRequest) SetSerialNo(v string) *CheckUuidValidRequest {
	s.SerialNo = &v
	return s
}

func (s *CheckUuidValidRequest) SetUuid(v string) *CheckUuidValidRequest {
	s.Uuid = &v
	return s
}

func (s *CheckUuidValidRequest) SetWlan(v string) *CheckUuidValidRequest {
	s.Wlan = &v
	return s
}

type CheckUuidValidResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckUuidValidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUuidValidResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUuidValidResponseBody) SetCode(v string) *CheckUuidValidResponseBody {
	s.Code = &v
	return s
}

func (s *CheckUuidValidResponseBody) SetMessage(v string) *CheckUuidValidResponseBody {
	s.Message = &v
	return s
}

func (s *CheckUuidValidResponseBody) SetRequestId(v string) *CheckUuidValidResponseBody {
	s.RequestId = &v
	return s
}

type CheckUuidValidResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUuidValidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUuidValidResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUuidValidResponse) GoString() string {
	return s.String()
}

func (s *CheckUuidValidResponse) SetHeaders(v map[string]*string) *CheckUuidValidResponse {
	s.Headers = v
	return s
}

func (s *CheckUuidValidResponse) SetStatusCode(v int32) *CheckUuidValidResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUuidValidResponse) SetBody(v *CheckUuidValidResponseBody) *CheckUuidValidResponse {
	s.Body = v
	return s
}

type CreateAppOtaTaskRequest struct {
	AppVersionUid *string   `json:"AppVersionUid,omitempty" xml:"AppVersionUid,omitempty"`
	Channel       *string   `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ClientIdList  []*string `json:"ClientIdList,omitempty" xml:"ClientIdList,omitempty" type:"Repeated"`
	ClientType    *int32    `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Creator       *string   `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description   *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ForceUpgrade  *int32    `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	Label         *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Project       *string   `json:"Project,omitempty" xml:"Project,omitempty"`
	Regions       []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	Status        *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskType      *int32    `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TenantId      *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateAppOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAppOtaTaskRequest) SetAppVersionUid(v string) *CreateAppOtaTaskRequest {
	s.AppVersionUid = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetChannel(v string) *CreateAppOtaTaskRequest {
	s.Channel = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetClientIdList(v []*string) *CreateAppOtaTaskRequest {
	s.ClientIdList = v
	return s
}

func (s *CreateAppOtaTaskRequest) SetClientType(v int32) *CreateAppOtaTaskRequest {
	s.ClientType = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetCreator(v string) *CreateAppOtaTaskRequest {
	s.Creator = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetDescription(v string) *CreateAppOtaTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetForceUpgrade(v int32) *CreateAppOtaTaskRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetLabel(v string) *CreateAppOtaTaskRequest {
	s.Label = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetName(v string) *CreateAppOtaTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetProject(v string) *CreateAppOtaTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetRegions(v []*string) *CreateAppOtaTaskRequest {
	s.Regions = v
	return s
}

func (s *CreateAppOtaTaskRequest) SetStatus(v int32) *CreateAppOtaTaskRequest {
	s.Status = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetTaskType(v int32) *CreateAppOtaTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetTenantId(v string) *CreateAppOtaTaskRequest {
	s.TenantId = &v
	return s
}

type CreateAppOtaTaskResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateAppOtaTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppOtaTaskResponseBody) SetCode(v string) *CreateAppOtaTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppOtaTaskResponseBody) SetData(v *CreateAppOtaTaskResponseBodyData) *CreateAppOtaTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppOtaTaskResponseBody) SetMessage(v string) *CreateAppOtaTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppOtaTaskResponseBody) SetRequestId(v string) *CreateAppOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppOtaTaskResponseBodyData struct {
	TaskUid *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s CreateAppOtaTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOtaTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppOtaTaskResponseBodyData) SetTaskUid(v string) *CreateAppOtaTaskResponseBodyData {
	s.TaskUid = &v
	return s
}

type CreateAppOtaTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAppOtaTaskResponse) SetHeaders(v map[string]*string) *CreateAppOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAppOtaTaskResponse) SetStatusCode(v int32) *CreateAppOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppOtaTaskResponse) SetBody(v *CreateAppOtaTaskResponseBody) *CreateAppOtaTaskResponse {
	s.Body = v
	return s
}

type CreateAppOtaVersionRequest struct {
	AppVersion       *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	Arch             *string `json:"Arch,omitempty" xml:"Arch,omitempty"`
	Channel          *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ClientType       *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Os               *string `json:"Os,omitempty" xml:"Os,omitempty"`
	OsType           *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OtaType          *int32  `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ReleaseNote      *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn    *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	ReleaseNoteJp    *string `json:"ReleaseNoteJp,omitempty" xml:"ReleaseNoteJp,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId       *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotRegionId *string `json:"SnapshotRegionId,omitempty" xml:"SnapshotRegionId,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionType      *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s CreateAppOtaVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOtaVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppOtaVersionRequest) SetAppVersion(v string) *CreateAppOtaVersionRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetArch(v string) *CreateAppOtaVersionRequest {
	s.Arch = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetChannel(v string) *CreateAppOtaVersionRequest {
	s.Channel = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetClientType(v int32) *CreateAppOtaVersionRequest {
	s.ClientType = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetCreator(v string) *CreateAppOtaVersionRequest {
	s.Creator = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetDownloadUrl(v string) *CreateAppOtaVersionRequest {
	s.DownloadUrl = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetMd5(v string) *CreateAppOtaVersionRequest {
	s.Md5 = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetOs(v string) *CreateAppOtaVersionRequest {
	s.Os = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetOsType(v string) *CreateAppOtaVersionRequest {
	s.OsType = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetOtaType(v int32) *CreateAppOtaVersionRequest {
	s.OtaType = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetProject(v string) *CreateAppOtaVersionRequest {
	s.Project = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetReleaseNote(v string) *CreateAppOtaVersionRequest {
	s.ReleaseNote = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetReleaseNoteEn(v string) *CreateAppOtaVersionRequest {
	s.ReleaseNoteEn = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetReleaseNoteJp(v string) *CreateAppOtaVersionRequest {
	s.ReleaseNoteJp = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetSize(v int64) *CreateAppOtaVersionRequest {
	s.Size = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetSnapshotId(v string) *CreateAppOtaVersionRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetSnapshotRegionId(v string) *CreateAppOtaVersionRequest {
	s.SnapshotRegionId = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetStatus(v int32) *CreateAppOtaVersionRequest {
	s.Status = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetVersionType(v string) *CreateAppOtaVersionRequest {
	s.VersionType = &v
	return s
}

type CreateAppOtaVersionResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateAppOtaVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppOtaVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOtaVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppOtaVersionResponseBody) SetCode(v string) *CreateAppOtaVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppOtaVersionResponseBody) SetData(v *CreateAppOtaVersionResponseBodyData) *CreateAppOtaVersionResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppOtaVersionResponseBody) SetMessage(v string) *CreateAppOtaVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppOtaVersionResponseBody) SetRequestId(v string) *CreateAppOtaVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppOtaVersionResponseBodyData struct {
	VersionUid *string `json:"VersionUid,omitempty" xml:"VersionUid,omitempty"`
}

func (s CreateAppOtaVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOtaVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppOtaVersionResponseBodyData) SetVersionUid(v string) *CreateAppOtaVersionResponseBodyData {
	s.VersionUid = &v
	return s
}

type CreateAppOtaVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppOtaVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppOtaVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOtaVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateAppOtaVersionResponse) SetHeaders(v map[string]*string) *CreateAppOtaVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateAppOtaVersionResponse) SetStatusCode(v int32) *CreateAppOtaVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppOtaVersionResponse) SetBody(v *CreateAppOtaVersionResponseBody) *CreateAppOtaVersionResponse {
	s.Body = v
	return s
}

type DeleteAppOtaVersionsRequest struct {
	VersionUidList []*string `json:"VersionUidList,omitempty" xml:"VersionUidList,omitempty" type:"Repeated"`
}

func (s DeleteAppOtaVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppOtaVersionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppOtaVersionsRequest) SetVersionUidList(v []*string) *DeleteAppOtaVersionsRequest {
	s.VersionUidList = v
	return s
}

type DeleteAppOtaVersionsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppOtaVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppOtaVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppOtaVersionsResponseBody) SetCode(v string) *DeleteAppOtaVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppOtaVersionsResponseBody) SetMessage(v string) *DeleteAppOtaVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppOtaVersionsResponseBody) SetRequestId(v string) *DeleteAppOtaVersionsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppOtaVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppOtaVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppOtaVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppOtaVersionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppOtaVersionsResponse) SetHeaders(v map[string]*string) *DeleteAppOtaVersionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppOtaVersionsResponse) SetStatusCode(v int32) *DeleteAppOtaVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppOtaVersionsResponse) SetBody(v *DeleteAppOtaVersionsResponseBody) *DeleteAppOtaVersionsResponse {
	s.Body = v
	return s
}

type DeleteDevicesRequest struct {
	// This parameter is required.
	Force     *string `json:"Force,omitempty" xml:"Force,omitempty"`
	SerialNos *string `json:"SerialNos,omitempty" xml:"SerialNos,omitempty"`
	Uuids     *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DeleteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevicesRequest) SetForce(v string) *DeleteDevicesRequest {
	s.Force = &v
	return s
}

func (s *DeleteDevicesRequest) SetSerialNos(v string) *DeleteDevicesRequest {
	s.SerialNos = &v
	return s
}

func (s *DeleteDevicesRequest) SetUuids(v string) *DeleteDevicesRequest {
	s.Uuids = &v
	return s
}

type DeleteDevicesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevicesResponseBody) SetCode(v string) *DeleteDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDevicesResponseBody) SetMessage(v string) *DeleteDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDevicesResponseBody) SetRequestId(v string) *DeleteDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDevicesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDevicesResponse) SetHeaders(v map[string]*string) *DeleteDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDevicesResponse) SetStatusCode(v int32) *DeleteDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDevicesResponse) SetBody(v *DeleteDevicesResponseBody) *DeleteDevicesResponse {
	s.Body = v
	return s
}

type DeleteLabelRequest struct {
	// This parameter is required.
	Force        *string `json:"Force,omitempty" xml:"Force,omitempty"`
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
}

func (s DeleteLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteLabelRequest) SetForce(v string) *DeleteLabelRequest {
	s.Force = &v
	return s
}

func (s *DeleteLabelRequest) SetLabelContent(v string) *DeleteLabelRequest {
	s.LabelContent = &v
	return s
}

func (s *DeleteLabelRequest) SetLabelId(v string) *DeleteLabelRequest {
	s.LabelId = &v
	return s
}

type DeleteLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLabelResponseBody) SetCode(v string) *DeleteLabelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLabelResponseBody) SetMessage(v string) *DeleteLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLabelResponseBody) SetRequestId(v string) *DeleteLabelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLabelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteLabelResponse) SetHeaders(v map[string]*string) *DeleteLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteLabelResponse) SetStatusCode(v int32) *DeleteLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLabelResponse) SetBody(v *DeleteLabelResponseBody) *DeleteLabelResponse {
	s.Body = v
	return s
}

type DescribeAppOtaVersionRequest struct {
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	Channel    *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ClientType *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Creator    *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionUid *string `json:"VersionUid,omitempty" xml:"VersionUid,omitempty"`
}

func (s DescribeAppOtaVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppOtaVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionRequest) SetAppVersion(v string) *DescribeAppOtaVersionRequest {
	s.AppVersion = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetChannel(v string) *DescribeAppOtaVersionRequest {
	s.Channel = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetClientType(v int32) *DescribeAppOtaVersionRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetCreator(v string) *DescribeAppOtaVersionRequest {
	s.Creator = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetProject(v string) *DescribeAppOtaVersionRequest {
	s.Project = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetStatus(v int32) *DescribeAppOtaVersionRequest {
	s.Status = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetVersionUid(v string) *DescribeAppOtaVersionRequest {
	s.VersionUid = &v
	return s
}

type DescribeAppOtaVersionResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeAppOtaVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppOtaVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppOtaVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionResponseBody) SetCode(v string) *DescribeAppOtaVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBody) SetData(v *DescribeAppOtaVersionResponseBodyData) *DescribeAppOtaVersionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAppOtaVersionResponseBody) SetMessage(v string) *DescribeAppOtaVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBody) SetRequestId(v string) *DescribeAppOtaVersionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAppOtaVersionResponseBodyData struct {
	AppOtaInfoDTOList []*DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList `json:"AppOtaInfoDTOList,omitempty" xml:"AppOtaInfoDTOList,omitempty" type:"Repeated"`
}

func (s DescribeAppOtaVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppOtaVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionResponseBodyData) SetAppOtaInfoDTOList(v []*DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) *DescribeAppOtaVersionResponseBodyData {
	s.AppOtaInfoDTOList = v
	return s
}

type DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList struct {
	AppVersion      *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	Channel         *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	DownloadUrl     *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FullDownloadUrl *string `json:"FullDownloadUrl,omitempty" xml:"FullDownloadUrl,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	OsType          *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ProtocolType    *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ReleaseNote     *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn   *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	SessionType     *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionCode     *int64  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType     *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionUid      *string `json:"VersionUid,omitempty" xml:"VersionUid,omitempty"`
}

func (s DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetAppVersion(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.AppVersion = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetChannel(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Channel = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetDownloadUrl(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetFullDownloadUrl(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.FullDownloadUrl = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetGmtCreate(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetMd5(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Md5 = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetOsType(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.OsType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetProject(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Project = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetProtocolType(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.ProtocolType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetReleaseNote(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetReleaseNoteEn(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetSessionType(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.SessionType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetSize(v int64) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Size = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetStatus(v int32) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Status = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetVersionCode(v int64) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.VersionCode = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetVersionType(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.VersionType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetVersionUid(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.VersionUid = &v
	return s
}

type DescribeAppOtaVersionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppOtaVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppOtaVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppOtaVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionResponse) SetHeaders(v map[string]*string) *DescribeAppOtaVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppOtaVersionResponse) SetStatusCode(v int32) *DescribeAppOtaVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppOtaVersionResponse) SetBody(v *DescribeAppOtaVersionResponseBody) *DescribeAppOtaVersionResponse {
	s.Body = v
	return s
}

type DescribeDeviceSeatsRequest struct {
	PageNumber   *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SerialNo     *string   `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SerialNoList []*string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty" type:"Repeated"`
	SiteId       *string   `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TenantId     *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeDeviceSeatsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceSeatsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsRequest) SetPageNumber(v int32) *DescribeDeviceSeatsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetPageSize(v int32) *DescribeDeviceSeatsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetSerialNo(v string) *DescribeDeviceSeatsRequest {
	s.SerialNo = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetSerialNoList(v []*string) *DescribeDeviceSeatsRequest {
	s.SerialNoList = v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetSiteId(v string) *DescribeDeviceSeatsRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetTenantId(v string) *DescribeDeviceSeatsRequest {
	s.TenantId = &v
	return s
}

type DescribeDeviceSeatsResponseBody struct {
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*DescribeDeviceSeatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeviceSeatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsResponseBody) SetCode(v string) *DescribeDeviceSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetData(v []*DescribeDeviceSeatsResponseBodyData) *DescribeDeviceSeatsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetMessage(v string) *DescribeDeviceSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetPageNumber(v int32) *DescribeDeviceSeatsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetPageSize(v int32) *DescribeDeviceSeatsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetRequestId(v string) *DescribeDeviceSeatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetTotalCount(v int64) *DescribeDeviceSeatsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDeviceSeatsResponseBodyData struct {
	SeatCol  *int32  `json:"SeatCol,omitempty" xml:"SeatCol,omitempty"`
	SeatName *string `json:"SeatName,omitempty" xml:"SeatName,omitempty"`
	SeatNo   *string `json:"SeatNo,omitempty" xml:"SeatNo,omitempty"`
	SeatRow  *int32  `json:"SeatRow,omitempty" xml:"SeatRow,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SiteId   *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s DescribeDeviceSeatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceSeatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSeatCol(v int32) *DescribeDeviceSeatsResponseBodyData {
	s.SeatCol = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSeatName(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SeatName = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSeatNo(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SeatNo = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSeatRow(v int32) *DescribeDeviceSeatsResponseBodyData {
	s.SeatRow = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSerialNo(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSiteId(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SiteId = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSiteName(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SiteName = &v
	return s
}

type DescribeDeviceSeatsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceSeatsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceSeatsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsResponse) SetHeaders(v map[string]*string) *DescribeDeviceSeatsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceSeatsResponse) SetStatusCode(v int32) *DescribeDeviceSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceSeatsResponse) SetBody(v *DescribeDeviceSeatsResponseBody) *DescribeDeviceSeatsResponse {
	s.Body = v
	return s
}

type DescribeDeviceVersionDetailRequest struct {
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DescribeDeviceVersionDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceVersionDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceVersionDetailRequest) SetModel(v string) *DescribeDeviceVersionDetailRequest {
	s.Model = &v
	return s
}

func (s *DescribeDeviceVersionDetailRequest) SetNetworkType(v string) *DescribeDeviceVersionDetailRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeDeviceVersionDetailRequest) SetRegion(v string) *DescribeDeviceVersionDetailRequest {
	s.Region = &v
	return s
}

func (s *DescribeDeviceVersionDetailRequest) SetVersionName(v string) *DescribeDeviceVersionDetailRequest {
	s.VersionName = &v
	return s
}

type DescribeDeviceVersionDetailResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeDeviceVersionDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceVersionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceVersionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceVersionDetailResponseBody) SetCode(v string) *DescribeDeviceVersionDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBody) SetData(v []*DescribeDeviceVersionDetailResponseBodyData) *DescribeDeviceVersionDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBody) SetMessage(v string) *DescribeDeviceVersionDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBody) SetRequestId(v string) *DescribeDeviceVersionDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeviceVersionDetailResponseBodyData struct {
	AndroidHorizontalMultiCnImageDownloadUrl *string `json:"AndroidHorizontalMultiCnImageDownloadUrl,omitempty" xml:"AndroidHorizontalMultiCnImageDownloadUrl,omitempty"`
	AndroidHorizontalMultiEnImageDownloadUrl *string `json:"AndroidHorizontalMultiEnImageDownloadUrl,omitempty" xml:"AndroidHorizontalMultiEnImageDownloadUrl,omitempty"`
	AndroidVerticalMultiCnImageDownloadUrl   *string `json:"AndroidVerticalMultiCnImageDownloadUrl,omitempty" xml:"AndroidVerticalMultiCnImageDownloadUrl,omitempty"`
	AndroidVerticalMultiEnImageDownloadUrl   *string `json:"AndroidVerticalMultiEnImageDownloadUrl,omitempty" xml:"AndroidVerticalMultiEnImageDownloadUrl,omitempty"`
	Channel                                  *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ClientType                               *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	CnImageDownloadUrl                       *string `json:"CnImageDownloadUrl,omitempty" xml:"CnImageDownloadUrl,omitempty"`
	Creator                                  *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl                              *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	EnImageDownloadUrl                       *string `json:"EnImageDownloadUrl,omitempty" xml:"EnImageDownloadUrl,omitempty"`
	Md5                                      *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Model                                    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	MultiCnImageDownloadUrl                  *string `json:"MultiCnImageDownloadUrl,omitempty" xml:"MultiCnImageDownloadUrl,omitempty"`
	MultiEnImageDownloadUrl                  *string `json:"MultiEnImageDownloadUrl,omitempty" xml:"MultiEnImageDownloadUrl,omitempty"`
	ReleaseNote                              *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn                            *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	Size                                     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version                                  *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionCode                              *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType                              *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DescribeDeviceVersionDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceVersionDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetAndroidHorizontalMultiCnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.AndroidHorizontalMultiCnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetAndroidHorizontalMultiEnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.AndroidHorizontalMultiEnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetAndroidVerticalMultiCnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.AndroidVerticalMultiCnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetAndroidVerticalMultiEnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.AndroidVerticalMultiEnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetChannel(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Channel = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetClientType(v int32) *DescribeDeviceVersionDetailResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetCnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.CnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetCreator(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Creator = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetEnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.EnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetMd5(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetModel(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Model = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetMultiCnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.MultiCnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetMultiEnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.MultiEnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetReleaseNote(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetReleaseNoteEn(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetSize(v int64) *DescribeDeviceVersionDetailResponseBodyData {
	s.Size = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetVersion(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetVersionCode(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetVersionType(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.VersionType = &v
	return s
}

type DescribeDeviceVersionDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceVersionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceVersionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceVersionDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceVersionDetailResponse) SetHeaders(v map[string]*string) *DescribeDeviceVersionDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceVersionDetailResponse) SetStatusCode(v int32) *DescribeDeviceVersionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponse) SetBody(v *DescribeDeviceVersionDetailResponseBody) *DescribeDeviceVersionDetailResponse {
	s.Body = v
	return s
}

type DescribeSnLabelCountsRequest struct {
	LabelList []*string `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	TenantId  *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId    *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName  *string   `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeSnLabelCountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnLabelCountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsRequest) SetLabelList(v []*string) *DescribeSnLabelCountsRequest {
	s.LabelList = v
	return s
}

func (s *DescribeSnLabelCountsRequest) SetTenantId(v string) *DescribeSnLabelCountsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeSnLabelCountsRequest) SetZoneId(v string) *DescribeSnLabelCountsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeSnLabelCountsRequest) SetZoneName(v string) *DescribeSnLabelCountsRequest {
	s.ZoneName = &v
	return s
}

type DescribeSnLabelCountsResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeSnLabelCountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSnLabelCountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnLabelCountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsResponseBody) SetCode(v string) *DescribeSnLabelCountsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBody) SetData(v *DescribeSnLabelCountsResponseBodyData) *DescribeSnLabelCountsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSnLabelCountsResponseBody) SetMessage(v string) *DescribeSnLabelCountsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBody) SetRequestId(v string) *DescribeSnLabelCountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSnLabelCountsResponseBodyData struct {
	LabelCountDTOList []*DescribeSnLabelCountsResponseBodyDataLabelCountDTOList `json:"LabelCountDTOList,omitempty" xml:"LabelCountDTOList,omitempty" type:"Repeated"`
	TotalCount        *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnLabelCountsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnLabelCountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsResponseBodyData) SetLabelCountDTOList(v []*DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) *DescribeSnLabelCountsResponseBodyData {
	s.LabelCountDTOList = v
	return s
}

func (s *DescribeSnLabelCountsResponseBodyData) SetTotalCount(v int32) *DescribeSnLabelCountsResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeSnLabelCountsResponseBodyDataLabelCountDTOList struct {
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) SetCount(v string) *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList {
	s.Count = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) SetLabel(v string) *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList {
	s.Label = &v
	return s
}

type DescribeSnLabelCountsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnLabelCountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnLabelCountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnLabelCountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsResponse) SetHeaders(v map[string]*string) *DescribeSnLabelCountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnLabelCountsResponse) SetStatusCode(v int32) *DescribeSnLabelCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnLabelCountsResponse) SetBody(v *DescribeSnLabelCountsResponseBody) *DescribeSnLabelCountsResponse {
	s.Body = v
	return s
}

type DescribeWorkZonesRequest struct {
	PageNumber   *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId     *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneIdList   []*string `json:"ZoneIdList,omitempty" xml:"ZoneIdList,omitempty" type:"Repeated"`
	ZoneNameList []*string `json:"ZoneNameList,omitempty" xml:"ZoneNameList,omitempty" type:"Repeated"`
}

func (s DescribeWorkZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesRequest) SetPageNumber(v int32) *DescribeWorkZonesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWorkZonesRequest) SetPageSize(v int32) *DescribeWorkZonesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWorkZonesRequest) SetTenantId(v string) *DescribeWorkZonesRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeWorkZonesRequest) SetZoneIdList(v []*string) *DescribeWorkZonesRequest {
	s.ZoneIdList = v
	return s
}

func (s *DescribeWorkZonesRequest) SetZoneNameList(v []*string) *DescribeWorkZonesRequest {
	s.ZoneNameList = v
	return s
}

type DescribeWorkZonesResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeWorkZonesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWorkZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesResponseBody) SetCode(v string) *DescribeWorkZonesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWorkZonesResponseBody) SetData(v *DescribeWorkZonesResponseBodyData) *DescribeWorkZonesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWorkZonesResponseBody) SetMessage(v string) *DescribeWorkZonesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWorkZonesResponseBody) SetRequestId(v string) *DescribeWorkZonesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWorkZonesResponseBodyData struct {
	TotalCount      *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WorkZoneDTOList []*DescribeWorkZonesResponseBodyDataWorkZoneDTOList `json:"WorkZoneDTOList,omitempty" xml:"WorkZoneDTOList,omitempty" type:"Repeated"`
}

func (s DescribeWorkZonesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesResponseBodyData) SetTotalCount(v int64) *DescribeWorkZonesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyData) SetWorkZoneDTOList(v []*DescribeWorkZonesResponseBodyDataWorkZoneDTOList) *DescribeWorkZonesResponseBodyData {
	s.WorkZoneDTOList = v
	return s
}

type DescribeWorkZonesResponseBodyDataWorkZoneDTOList struct {
	SeatCol  *int32  `json:"SeatCol,omitempty" xml:"SeatCol,omitempty"`
	SeatRow  *int32  `json:"SeatRow,omitempty" xml:"SeatRow,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeWorkZonesResponseBodyDataWorkZoneDTOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkZonesResponseBodyDataWorkZoneDTOList) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetSeatCol(v int32) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.SeatCol = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetSeatRow(v int32) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.SeatRow = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetTenantId(v string) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.TenantId = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetZoneId(v string) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.ZoneId = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetZoneName(v string) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.ZoneName = &v
	return s
}

type DescribeWorkZonesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWorkZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWorkZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesResponse) SetHeaders(v map[string]*string) *DescribeWorkZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkZonesResponse) SetStatusCode(v int32) *DescribeWorkZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWorkZonesResponse) SetBody(v *DescribeWorkZonesResponseBody) *DescribeWorkZonesResponse {
	s.Body = v
	return s
}

type DetachEndUsersRequest struct {
	// This parameter is required.
	EndUserIds *string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s DetachEndUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachEndUsersRequest) GoString() string {
	return s.String()
}

func (s *DetachEndUsersRequest) SetEndUserIds(v string) *DetachEndUsersRequest {
	s.EndUserIds = &v
	return s
}

func (s *DetachEndUsersRequest) SetSerialNo(v string) *DetachEndUsersRequest {
	s.SerialNo = &v
	return s
}

type DetachEndUsersResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachEndUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachEndUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DetachEndUsersResponseBody) SetCode(v string) *DetachEndUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DetachEndUsersResponseBody) SetMessage(v string) *DetachEndUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DetachEndUsersResponseBody) SetRequestId(v string) *DetachEndUsersResponseBody {
	s.RequestId = &v
	return s
}

type DetachEndUsersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachEndUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachEndUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachEndUsersResponse) GoString() string {
	return s.String()
}

func (s *DetachEndUsersResponse) SetHeaders(v map[string]*string) *DetachEndUsersResponse {
	s.Headers = v
	return s
}

func (s *DetachEndUsersResponse) SetStatusCode(v int32) *DetachEndUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachEndUsersResponse) SetBody(v *DetachEndUsersResponseBody) *DetachEndUsersResponse {
	s.Body = v
	return s
}

type DetachLabelRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s DetachLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachLabelRequest) GoString() string {
	return s.String()
}

func (s *DetachLabelRequest) SetLabelContent(v string) *DetachLabelRequest {
	s.LabelContent = &v
	return s
}

func (s *DetachLabelRequest) SetLabelId(v string) *DetachLabelRequest {
	s.LabelId = &v
	return s
}

func (s *DetachLabelRequest) SetSerialNo(v string) *DetachLabelRequest {
	s.SerialNo = &v
	return s
}

type DetachLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLabelResponseBody) SetCode(v string) *DetachLabelResponseBody {
	s.Code = &v
	return s
}

func (s *DetachLabelResponseBody) SetMessage(v string) *DetachLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DetachLabelResponseBody) SetRequestId(v string) *DetachLabelResponseBody {
	s.RequestId = &v
	return s
}

type DetachLabelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachLabelResponse) GoString() string {
	return s.String()
}

func (s *DetachLabelResponse) SetHeaders(v map[string]*string) *DetachLabelResponse {
	s.Headers = v
	return s
}

func (s *DetachLabelResponse) SetStatusCode(v int32) *DetachLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachLabelResponse) SetBody(v *DetachLabelResponseBody) *DetachLabelResponse {
	s.Body = v
	return s
}

type DetachLabelsRequest struct {
	LabelIds     *string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SerialNoList *string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty"`
}

func (s DetachLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachLabelsRequest) GoString() string {
	return s.String()
}

func (s *DetachLabelsRequest) SetLabelIds(v string) *DetachLabelsRequest {
	s.LabelIds = &v
	return s
}

func (s *DetachLabelsRequest) SetSerialNo(v string) *DetachLabelsRequest {
	s.SerialNo = &v
	return s
}

func (s *DetachLabelsRequest) SetSerialNoList(v string) *DetachLabelsRequest {
	s.SerialNoList = &v
	return s
}

type DetachLabelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLabelsResponseBody) SetCode(v string) *DetachLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *DetachLabelsResponseBody) SetMessage(v string) *DetachLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *DetachLabelsResponseBody) SetRequestId(v string) *DetachLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DetachLabelsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachLabelsResponse) GoString() string {
	return s.String()
}

func (s *DetachLabelsResponse) SetHeaders(v map[string]*string) *DetachLabelsResponse {
	s.Headers = v
	return s
}

func (s *DetachLabelsResponse) SetStatusCode(v int32) *DetachLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachLabelsResponse) SetBody(v *DetachLabelsResponseBody) *DetachLabelsResponse {
	s.Body = v
	return s
}

type GenerateOssUrlRequest struct {
	ObjectNameList []*string `json:"ObjectNameList,omitempty" xml:"ObjectNameList,omitempty" type:"Repeated"`
	SessionId      *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GenerateOssUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateOssUrlRequest) SetObjectNameList(v []*string) *GenerateOssUrlRequest {
	s.ObjectNameList = v
	return s
}

func (s *GenerateOssUrlRequest) SetSessionId(v string) *GenerateOssUrlRequest {
	s.SessionId = &v
	return s
}

type GenerateOssUrlResponseBody struct {
	Data      []*GenerateOssUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateOssUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOssUrlResponseBody) SetData(v []*GenerateOssUrlResponseBodyData) *GenerateOssUrlResponseBody {
	s.Data = v
	return s
}

func (s *GenerateOssUrlResponseBody) SetRequestId(v string) *GenerateOssUrlResponseBody {
	s.RequestId = &v
	return s
}

type GenerateOssUrlResponseBodyData struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ObjectName  *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s GenerateOssUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateOssUrlResponseBodyData) SetDownloadUrl(v string) *GenerateOssUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GenerateOssUrlResponseBodyData) SetObjectName(v string) *GenerateOssUrlResponseBodyData {
	s.ObjectName = &v
	return s
}

type GenerateOssUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateOssUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateOssUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateOssUrlResponse) SetHeaders(v map[string]*string) *GenerateOssUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateOssUrlResponse) SetStatusCode(v int32) *GenerateOssUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOssUrlResponse) SetBody(v *GenerateOssUrlResponseBody) *GenerateOssUrlResponse {
	s.Body = v
	return s
}

type GetAppOtaLatestVersionRequest struct {
	// This parameter is required.
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	ClientType  *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientUid   *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	// This parameter is required.
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetAppOtaLatestVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppOtaLatestVersionRequest) GoString() string {
	return s.String()
}

func (s *GetAppOtaLatestVersionRequest) SetBaseVersion(v string) *GetAppOtaLatestVersionRequest {
	s.BaseVersion = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) SetClientType(v int32) *GetAppOtaLatestVersionRequest {
	s.ClientType = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) SetClientUid(v string) *GetAppOtaLatestVersionRequest {
	s.ClientUid = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) SetOsType(v string) *GetAppOtaLatestVersionRequest {
	s.OsType = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) SetProject(v string) *GetAppOtaLatestVersionRequest {
	s.Project = &v
	return s
}

type GetAppOtaLatestVersionResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAppOtaLatestVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppOtaLatestVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppOtaLatestVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppOtaLatestVersionResponseBody) SetCode(v string) *GetAppOtaLatestVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBody) SetData(v *GetAppOtaLatestVersionResponseBodyData) *GetAppOtaLatestVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetAppOtaLatestVersionResponseBody) SetMessage(v string) *GetAppOtaLatestVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBody) SetRequestId(v string) *GetAppOtaLatestVersionResponseBody {
	s.RequestId = &v
	return s
}

type GetAppOtaLatestVersionResponseBodyData struct {
	AppVersion   *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	DownloadUrl  *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ForceUpgrade *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	Md5          *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	ReleaseNote  *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	TaskUid      *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
	VersionCode  *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s GetAppOtaLatestVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppOtaLatestVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetAppVersion(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetDownloadUrl(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetForceUpgrade(v int32) *GetAppOtaLatestVersionResponseBodyData {
	s.ForceUpgrade = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetMd5(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetReleaseNote(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.ReleaseNote = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetSize(v int64) *GetAppOtaLatestVersionResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetTaskUid(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.TaskUid = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetVersionCode(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetVersionType(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.VersionType = &v
	return s
}

type GetAppOtaLatestVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppOtaLatestVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppOtaLatestVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppOtaLatestVersionResponse) GoString() string {
	return s.String()
}

func (s *GetAppOtaLatestVersionResponse) SetHeaders(v map[string]*string) *GetAppOtaLatestVersionResponse {
	s.Headers = v
	return s
}

func (s *GetAppOtaLatestVersionResponse) SetStatusCode(v int32) *GetAppOtaLatestVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppOtaLatestVersionResponse) SetBody(v *GetAppOtaLatestVersionResponseBody) *GetAppOtaLatestVersionResponse {
	s.Body = v
	return s
}

type GetDeviceConfigsRequest struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	UrclVersion  *string `json:"UrclVersion,omitempty" xml:"UrclVersion,omitempty"`
	UserCustomId *string `json:"UserCustomId,omitempty" xml:"UserCustomId,omitempty"`
}

func (s GetDeviceConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigsRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsRequest) SetDeviceId(v string) *GetDeviceConfigsRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetNetworkType(v string) *GetDeviceConfigsRequest {
	s.NetworkType = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetRegion(v string) *GetDeviceConfigsRequest {
	s.Region = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetSerialNo(v string) *GetDeviceConfigsRequest {
	s.SerialNo = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetUrclVersion(v string) *GetDeviceConfigsRequest {
	s.UrclVersion = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetUserCustomId(v string) *GetDeviceConfigsRequest {
	s.UserCustomId = &v
	return s
}

type GetDeviceConfigsResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponseBody) SetCode(v string) *GetDeviceConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceConfigsResponseBody) SetData(v *GetDeviceConfigsResponseBodyData) *GetDeviceConfigsResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceConfigsResponseBody) SetMessage(v string) *GetDeviceConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceConfigsResponseBody) SetRequestId(v string) *GetDeviceConfigsResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceConfigsResponseBodyData struct {
	AutoLockScreenTime       *int32                                                 `json:"AutoLockScreenTime,omitempty" xml:"AutoLockScreenTime,omitempty"`
	AutoLogin                *int32                                                 `json:"AutoLogin,omitempty" xml:"AutoLogin,omitempty"`
	AutoUpdate               *int32                                                 `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	CustomIdleAction         *int32                                                 `json:"CustomIdleAction,omitempty" xml:"CustomIdleAction,omitempty"`
	CustomPowerOn            *int32                                                 `json:"CustomPowerOn,omitempty" xml:"CustomPowerOn,omitempty"`
	CustomResourcePackage    *GetDeviceConfigsResponseBodyDataCustomResourcePackage `json:"CustomResourcePackage,omitempty" xml:"CustomResourcePackage,omitempty" type:"Struct"`
	DefinePowerButton        *int32                                                 `json:"DefinePowerButton,omitempty" xml:"DefinePowerButton,omitempty"`
	DeviceLock               *int32                                                 `json:"DeviceLock,omitempty" xml:"DeviceLock,omitempty"`
	DisplayLayout            *string                                                `json:"DisplayLayout,omitempty" xml:"DisplayLayout,omitempty"`
	DisplayResolution        *string                                                `json:"DisplayResolution,omitempty" xml:"DisplayResolution,omitempty"`
	DisplayScaleRatio        *string                                                `json:"DisplayScaleRatio,omitempty" xml:"DisplayScaleRatio,omitempty"`
	EnableAdb                *int32                                                 `json:"EnableAdb,omitempty" xml:"EnableAdb,omitempty"`
	EnableAutoLockScreen     *int32                                                 `json:"EnableAutoLockScreen,omitempty" xml:"EnableAutoLockScreen,omitempty"`
	EnableBluetooth          *int32                                                 `json:"EnableBluetooth,omitempty" xml:"EnableBluetooth,omitempty"`
	EnableLockScreenPassword *int32                                                 `json:"EnableLockScreenPassword,omitempty" xml:"EnableLockScreenPassword,omitempty"`
	EnableModifyPassword     *int32                                                 `json:"EnableModifyPassword,omitempty" xml:"EnableModifyPassword,omitempty"`
	EnableScheduledPowerOff  *int32                                                 `json:"EnableScheduledPowerOff,omitempty" xml:"EnableScheduledPowerOff,omitempty"`
	EnableUnlockPassword     *int32                                                 `json:"EnableUnlockPassword,omitempty" xml:"EnableUnlockPassword,omitempty"`
	EnableWlan               *int32                                                 `json:"EnableWlan,omitempty" xml:"EnableWlan,omitempty"`
	IdleTime                 *int32                                                 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	LocalUsbPrint            *int32                                                 `json:"LocalUsbPrint,omitempty" xml:"LocalUsbPrint,omitempty"`
	LockPassword             *string                                                `json:"LockPassword,omitempty" xml:"LockPassword,omitempty"`
	ScheduledPowerOff        *string                                                `json:"ScheduledPowerOff,omitempty" xml:"ScheduledPowerOff,omitempty"`
	SecureNetworkType        *string                                                `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	SerialNo                 *string                                                `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SleepTime                *int32                                                 `json:"SleepTime,omitempty" xml:"SleepTime,omitempty"`
	Urcl                     *string                                                `json:"Urcl,omitempty" xml:"Urcl,omitempty"`
	UsbStorage               *int32                                                 `json:"UsbStorage,omitempty" xml:"UsbStorage,omitempty"`
	UserCustomId             *string                                                `json:"UserCustomId,omitempty" xml:"UserCustomId,omitempty"`
	Uuid                     *string                                                `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetDeviceConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponseBodyData) SetAutoLockScreenTime(v int32) *GetDeviceConfigsResponseBodyData {
	s.AutoLockScreenTime = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetAutoLogin(v int32) *GetDeviceConfigsResponseBodyData {
	s.AutoLogin = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetAutoUpdate(v int32) *GetDeviceConfigsResponseBodyData {
	s.AutoUpdate = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetCustomIdleAction(v int32) *GetDeviceConfigsResponseBodyData {
	s.CustomIdleAction = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetCustomPowerOn(v int32) *GetDeviceConfigsResponseBodyData {
	s.CustomPowerOn = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetCustomResourcePackage(v *GetDeviceConfigsResponseBodyDataCustomResourcePackage) *GetDeviceConfigsResponseBodyData {
	s.CustomResourcePackage = v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDefinePowerButton(v int32) *GetDeviceConfigsResponseBodyData {
	s.DefinePowerButton = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDeviceLock(v int32) *GetDeviceConfigsResponseBodyData {
	s.DeviceLock = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDisplayLayout(v string) *GetDeviceConfigsResponseBodyData {
	s.DisplayLayout = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDisplayResolution(v string) *GetDeviceConfigsResponseBodyData {
	s.DisplayResolution = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDisplayScaleRatio(v string) *GetDeviceConfigsResponseBodyData {
	s.DisplayScaleRatio = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableAdb(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableAdb = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableAutoLockScreen(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableAutoLockScreen = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableBluetooth(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableBluetooth = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableLockScreenPassword(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableLockScreenPassword = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableModifyPassword(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableModifyPassword = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableScheduledPowerOff(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableScheduledPowerOff = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableUnlockPassword(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableUnlockPassword = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableWlan(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableWlan = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetIdleTime(v int32) *GetDeviceConfigsResponseBodyData {
	s.IdleTime = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetLocalUsbPrint(v int32) *GetDeviceConfigsResponseBodyData {
	s.LocalUsbPrint = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetLockPassword(v string) *GetDeviceConfigsResponseBodyData {
	s.LockPassword = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetScheduledPowerOff(v string) *GetDeviceConfigsResponseBodyData {
	s.ScheduledPowerOff = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetSecureNetworkType(v string) *GetDeviceConfigsResponseBodyData {
	s.SecureNetworkType = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetSerialNo(v string) *GetDeviceConfigsResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetSleepTime(v int32) *GetDeviceConfigsResponseBodyData {
	s.SleepTime = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUrcl(v string) *GetDeviceConfigsResponseBodyData {
	s.Urcl = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUsbStorage(v int32) *GetDeviceConfigsResponseBodyData {
	s.UsbStorage = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUserCustomId(v string) *GetDeviceConfigsResponseBodyData {
	s.UserCustomId = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUuid(v string) *GetDeviceConfigsResponseBodyData {
	s.Uuid = &v
	return s
}

type GetDeviceConfigsResponseBodyDataCustomResourcePackage struct {
	ConfigAboutLogo     *string `json:"ConfigAboutLogo,omitempty" xml:"ConfigAboutLogo,omitempty"`
	DesktopWallpaper    *string `json:"DesktopWallpaper,omitempty" xml:"DesktopWallpaper,omitempty"`
	LoginPageBackground *string `json:"LoginPageBackground,omitempty" xml:"LoginPageBackground,omitempty"`
	LoginPageLogo       *string `json:"LoginPageLogo,omitempty" xml:"LoginPageLogo,omitempty"`
	PersonalCenterLogo  *string `json:"PersonalCenterLogo,omitempty" xml:"PersonalCenterLogo,omitempty"`
	StartLogo           *string `json:"StartLogo,omitempty" xml:"StartLogo,omitempty"`
	StartMenuLogo       *string `json:"StartMenuLogo,omitempty" xml:"StartMenuLogo,omitempty"`
	UpgradeLogo         *string `json:"UpgradeLogo,omitempty" xml:"UpgradeLogo,omitempty"`
}

func (s GetDeviceConfigsResponseBodyDataCustomResourcePackage) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigsResponseBodyDataCustomResourcePackage) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetConfigAboutLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.ConfigAboutLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetDesktopWallpaper(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.DesktopWallpaper = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetLoginPageBackground(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.LoginPageBackground = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetLoginPageLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.LoginPageLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetPersonalCenterLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.PersonalCenterLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetStartLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.StartLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetStartMenuLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.StartMenuLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetUpgradeLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.UpgradeLogo = &v
	return s
}

type GetDeviceConfigsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigsResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponse) SetHeaders(v map[string]*string) *GetDeviceConfigsResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceConfigsResponse) SetStatusCode(v int32) *GetDeviceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceConfigsResponse) SetBody(v *GetDeviceConfigsResponseBody) *GetDeviceConfigsResponse {
	s.Body = v
	return s
}

type GetDeviceOtaAutoStatusRequest struct {
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
}

func (s GetDeviceOtaAutoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaAutoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusRequest) SetClientType(v int32) *GetDeviceOtaAutoStatusRequest {
	s.ClientType = &v
	return s
}

type GetDeviceOtaAutoStatusResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceOtaAutoStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOtaAutoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaAutoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusResponseBody) SetCode(v string) *GetDeviceOtaAutoStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBody) SetData(v *GetDeviceOtaAutoStatusResponseBodyData) *GetDeviceOtaAutoStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBody) SetMessage(v string) *GetDeviceOtaAutoStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBody) SetRequestId(v string) *GetDeviceOtaAutoStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceOtaAutoStatusResponseBodyData struct {
	AutoUpdate             *int32  `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	AutoUpdateTimeSchedule *string `json:"AutoUpdateTimeSchedule,omitempty" xml:"AutoUpdateTimeSchedule,omitempty"`
	ForceUpgrade           *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	Status                 *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeviceOtaAutoStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaAutoStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetAutoUpdate(v int32) *GetDeviceOtaAutoStatusResponseBodyData {
	s.AutoUpdate = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetAutoUpdateTimeSchedule(v string) *GetDeviceOtaAutoStatusResponseBodyData {
	s.AutoUpdateTimeSchedule = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetForceUpgrade(v int32) *GetDeviceOtaAutoStatusResponseBodyData {
	s.ForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetStatus(v int32) *GetDeviceOtaAutoStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetDeviceOtaAutoStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOtaAutoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOtaAutoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaAutoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusResponse) SetHeaders(v map[string]*string) *GetDeviceOtaAutoStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOtaAutoStatusResponse) SetStatusCode(v int32) *GetDeviceOtaAutoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponse) SetBody(v *GetDeviceOtaAutoStatusResponseBody) *GetDeviceOtaAutoStatusResponse {
	s.Body = v
	return s
}

type GetDeviceOtaInfoRequest struct {
	// This parameter is required.
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	Channel     *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// This parameter is required.
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	Model             *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OsVersion         *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TargetVersionType *string `json:"TargetVersionType,omitempty" xml:"TargetVersionType,omitempty"`
	TenantId          *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDeviceOtaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoRequest) SetBaseVersion(v string) *GetDeviceOtaInfoRequest {
	s.BaseVersion = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetChannel(v string) *GetDeviceOtaInfoRequest {
	s.Channel = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetDeviceId(v string) *GetDeviceOtaInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetModel(v string) *GetDeviceOtaInfoRequest {
	s.Model = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetNetworkType(v string) *GetDeviceOtaInfoRequest {
	s.NetworkType = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetOsVersion(v string) *GetDeviceOtaInfoRequest {
	s.OsVersion = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetRegion(v string) *GetDeviceOtaInfoRequest {
	s.Region = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetRegionId(v string) *GetDeviceOtaInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetTargetVersionType(v string) *GetDeviceOtaInfoRequest {
	s.TargetVersionType = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetTenantId(v string) *GetDeviceOtaInfoRequest {
	s.TenantId = &v
	return s
}

type GetDeviceOtaInfoResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceOtaInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOtaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponseBody) SetCode(v string) *GetDeviceOtaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBody) SetData(v *GetDeviceOtaInfoResponseBodyData) *GetDeviceOtaInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceOtaInfoResponseBody) SetMessage(v string) *GetDeviceOtaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBody) SetRequestId(v string) *GetDeviceOtaInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceOtaInfoResponseBodyData struct {
	Version *GetDeviceOtaInfoResponseBodyDataVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
}

func (s GetDeviceOtaInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponseBodyData) SetVersion(v *GetDeviceOtaInfoResponseBodyDataVersion) *GetDeviceOtaInfoResponseBodyData {
	s.Version = v
	return s
}

type GetDeviceOtaInfoResponseBodyDataVersion struct {
	AndroidHorizontalMultiCnImageDownloadUrl *string `json:"AndroidHorizontalMultiCnImageDownloadUrl,omitempty" xml:"AndroidHorizontalMultiCnImageDownloadUrl,omitempty"`
	AndroidHorizontalMultiEnImageDownloadUrl *string `json:"AndroidHorizontalMultiEnImageDownloadUrl,omitempty" xml:"AndroidHorizontalMultiEnImageDownloadUrl,omitempty"`
	AndroidVerticalMultiCnImageDownloadUrl   *string `json:"AndroidVerticalMultiCnImageDownloadUrl,omitempty" xml:"AndroidVerticalMultiCnImageDownloadUrl,omitempty"`
	AndroidVerticalMultiEnImageDownloadUrl   *string `json:"AndroidVerticalMultiEnImageDownloadUrl,omitempty" xml:"AndroidVerticalMultiEnImageDownloadUrl,omitempty"`
	CnImageDownloadUrl                       *string `json:"CnImageDownloadUrl,omitempty" xml:"CnImageDownloadUrl,omitempty"`
	Creator                                  *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CustomForceUpgrade                       *bool   `json:"CustomForceUpgrade,omitempty" xml:"CustomForceUpgrade,omitempty"`
	DownloadUrl                              *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	EnImageDownloadUrl                       *string `json:"EnImageDownloadUrl,omitempty" xml:"EnImageDownloadUrl,omitempty"`
	ForceUpgrade                             *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	IsAppDownloadUrl                         *bool   `json:"IsAppDownloadUrl,omitempty" xml:"IsAppDownloadUrl,omitempty"`
	LocalDownloadUrl                         *string `json:"LocalDownloadUrl,omitempty" xml:"LocalDownloadUrl,omitempty"`
	Md5                                      *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Model                                    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	MultiCnImageDownloadUrl                  *string `json:"MultiCnImageDownloadUrl,omitempty" xml:"MultiCnImageDownloadUrl,omitempty"`
	MultiEnImageDownloadUrl                  *string `json:"MultiEnImageDownloadUrl,omitempty" xml:"MultiEnImageDownloadUrl,omitempty"`
	ReleaseNote                              *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn                            *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	Size                                     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version                                  *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionCode                              *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType                              *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s GetDeviceOtaInfoResponseBodyDataVersion) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoResponseBodyDataVersion) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetAndroidHorizontalMultiCnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.AndroidHorizontalMultiCnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetAndroidHorizontalMultiEnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.AndroidHorizontalMultiEnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetAndroidVerticalMultiCnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.AndroidVerticalMultiCnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetAndroidVerticalMultiEnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.AndroidVerticalMultiEnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetCnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.CnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetCreator(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Creator = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetCustomForceUpgrade(v bool) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.CustomForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.DownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetEnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.EnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetForceUpgrade(v int32) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.ForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetIsAppDownloadUrl(v bool) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.IsAppDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetLocalDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.LocalDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetMd5(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Md5 = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetModel(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Model = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetMultiCnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.MultiCnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetMultiEnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.MultiEnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetReleaseNote(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.ReleaseNote = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetReleaseNoteEn(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.ReleaseNoteEn = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetSize(v int64) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Size = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetVersion(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Version = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetVersionCode(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.VersionCode = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetVersionType(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.VersionType = &v
	return s
}

type GetDeviceOtaInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOtaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponse) SetHeaders(v map[string]*string) *GetDeviceOtaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOtaInfoResponse) SetStatusCode(v int32) *GetDeviceOtaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOtaInfoResponse) SetBody(v *GetDeviceOtaInfoResponseBody) *GetDeviceOtaInfoResponse {
	s.Body = v
	return s
}

type GetDeviceOtaInfoTestRequest struct {
	// This parameter is required.
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	// This parameter is required.
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDeviceOtaInfoTestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoTestRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestRequest) SetBaseVersion(v string) *GetDeviceOtaInfoTestRequest {
	s.BaseVersion = &v
	return s
}

func (s *GetDeviceOtaInfoTestRequest) SetDeviceId(v string) *GetDeviceOtaInfoTestRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceOtaInfoTestRequest) SetModel(v string) *GetDeviceOtaInfoTestRequest {
	s.Model = &v
	return s
}

func (s *GetDeviceOtaInfoTestRequest) SetTenantId(v string) *GetDeviceOtaInfoTestRequest {
	s.TenantId = &v
	return s
}

type GetDeviceOtaInfoTestResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceOtaInfoTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOtaInfoTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoTestResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestResponseBody) SetCode(v string) *GetDeviceOtaInfoTestResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBody) SetData(v *GetDeviceOtaInfoTestResponseBodyData) *GetDeviceOtaInfoTestResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBody) SetMessage(v string) *GetDeviceOtaInfoTestResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBody) SetRequestId(v string) *GetDeviceOtaInfoTestResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceOtaInfoTestResponseBodyData struct {
	Version *GetDeviceOtaInfoTestResponseBodyDataVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
}

func (s GetDeviceOtaInfoTestResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestResponseBodyData) SetVersion(v *GetDeviceOtaInfoTestResponseBodyDataVersion) *GetDeviceOtaInfoTestResponseBodyData {
	s.Version = v
	return s
}

type GetDeviceOtaInfoTestResponseBodyDataVersion struct {
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ForceUpgrade     *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	LocalDownloadUrl *string `json:"LocalDownloadUrl,omitempty" xml:"LocalDownloadUrl,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Model            *string `json:"Model,omitempty" xml:"Model,omitempty"`
	ReleaseNote      *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version          *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionCode      *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType      *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s GetDeviceOtaInfoTestResponseBodyDataVersion) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoTestResponseBodyDataVersion) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetCreator(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Creator = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetDownloadUrl(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.DownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetForceUpgrade(v int32) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.ForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetLocalDownloadUrl(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.LocalDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetMd5(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Md5 = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetModel(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Model = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetReleaseNote(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.ReleaseNote = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetSize(v int64) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Size = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetVersion(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Version = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetVersionCode(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.VersionCode = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetVersionType(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.VersionType = &v
	return s
}

type GetDeviceOtaInfoTestResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOtaInfoTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOtaInfoTestResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoTestResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestResponse) SetHeaders(v map[string]*string) *GetDeviceOtaInfoTestResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOtaInfoTestResponse) SetStatusCode(v int32) *GetDeviceOtaInfoTestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponse) SetBody(v *GetDeviceOtaInfoTestResponseBody) *GetDeviceOtaInfoTestResponse {
	s.Body = v
	return s
}

type GetDeviceOtaTaskVersionInfoRequest struct {
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDeviceOtaTaskVersionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaTaskVersionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaTaskVersionInfoRequest) SetTaskId(v string) *GetDeviceOtaTaskVersionInfoRequest {
	s.TaskId = &v
	return s
}

type GetDeviceOtaTaskVersionInfoResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceOtaTaskVersionInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOtaTaskVersionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaTaskVersionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) SetCode(v string) *GetDeviceOtaTaskVersionInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) SetData(v *GetDeviceOtaTaskVersionInfoResponseBodyData) *GetDeviceOtaTaskVersionInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) SetMessage(v string) *GetDeviceOtaTaskVersionInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) SetRequestId(v string) *GetDeviceOtaTaskVersionInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceOtaTaskVersionInfoResponseBodyData struct {
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetDeviceOtaTaskVersionInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaTaskVersionInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) SetReleaseNote(v string) *GetDeviceOtaTaskVersionInfoResponseBodyData {
	s.ReleaseNote = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) SetSize(v int64) *GetDeviceOtaTaskVersionInfoResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) SetVersion(v string) *GetDeviceOtaTaskVersionInfoResponseBodyData {
	s.Version = &v
	return s
}

type GetDeviceOtaTaskVersionInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOtaTaskVersionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOtaTaskVersionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaTaskVersionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaTaskVersionInfoResponse) SetHeaders(v map[string]*string) *GetDeviceOtaTaskVersionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponse) SetStatusCode(v int32) *GetDeviceOtaTaskVersionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponse) SetBody(v *GetDeviceOtaTaskVersionInfoResponseBody) *GetDeviceOtaTaskVersionInfoResponse {
	s.Body = v
	return s
}

type GetDeviceUpgradeStatusRequest struct {
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ClientUid  *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskUid    *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s GetDeviceUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusRequest) SetAppVersion(v string) *GetDeviceUpgradeStatusRequest {
	s.AppVersion = &v
	return s
}

func (s *GetDeviceUpgradeStatusRequest) SetClientUid(v string) *GetDeviceUpgradeStatusRequest {
	s.ClientUid = &v
	return s
}

func (s *GetDeviceUpgradeStatusRequest) SetProject(v string) *GetDeviceUpgradeStatusRequest {
	s.Project = &v
	return s
}

func (s *GetDeviceUpgradeStatusRequest) SetTaskUid(v string) *GetDeviceUpgradeStatusRequest {
	s.TaskUid = &v
	return s
}

type GetDeviceUpgradeStatusResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceUpgradeStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceUpgradeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusResponseBody) SetCode(v string) *GetDeviceUpgradeStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBody) SetData(v *GetDeviceUpgradeStatusResponseBodyData) *GetDeviceUpgradeStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBody) SetMessage(v string) *GetDeviceUpgradeStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBody) SetRequestId(v string) *GetDeviceUpgradeStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceUpgradeStatusResponseBodyData struct {
	AppOtaStatusDTOList []*GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList `json:"AppOtaStatusDTOList,omitempty" xml:"AppOtaStatusDTOList,omitempty" type:"Repeated"`
}

func (s GetDeviceUpgradeStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceUpgradeStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusResponseBodyData) SetAppOtaStatusDTOList(v []*GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) *GetDeviceUpgradeStatusResponseBodyData {
	s.AppOtaStatusDTOList = v
	return s
}

type GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList struct {
	BaseVersion   *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	ClientType    *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientUid     *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OsType        *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Project       *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	TaskUid       *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetBaseVersion(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.BaseVersion = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetClientType(v int32) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.ClientType = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetClientUid(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.ClientUid = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetNote(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.Note = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetOsType(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.OsType = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetProject(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.Project = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetStatus(v int32) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.Status = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetTargetVersion(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.TargetVersion = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetTaskUid(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.TaskUid = &v
	return s
}

type GetDeviceUpgradeStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusResponse) SetHeaders(v map[string]*string) *GetDeviceUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceUpgradeStatusResponse) SetStatusCode(v int32) *GetDeviceUpgradeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponse) SetBody(v *GetDeviceUpgradeStatusResponseBody) *GetDeviceUpgradeStatusResponse {
	s.Body = v
	return s
}

type GetExportDeviceInfoOssUrlRequest struct {
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s GetExportDeviceInfoOssUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExportDeviceInfoOssUrlRequest) GoString() string {
	return s.String()
}

func (s *GetExportDeviceInfoOssUrlRequest) SetTenantId(v string) *GetExportDeviceInfoOssUrlRequest {
	s.TenantId = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlRequest) SetZoneId(v string) *GetExportDeviceInfoOssUrlRequest {
	s.ZoneId = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlRequest) SetZoneName(v string) *GetExportDeviceInfoOssUrlRequest {
	s.ZoneName = &v
	return s
}

type GetExportDeviceInfoOssUrlResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetExportDeviceInfoOssUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExportDeviceInfoOssUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExportDeviceInfoOssUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetExportDeviceInfoOssUrlResponseBody) SetCode(v string) *GetExportDeviceInfoOssUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponseBody) SetData(v *GetExportDeviceInfoOssUrlResponseBodyData) *GetExportDeviceInfoOssUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponseBody) SetMessage(v string) *GetExportDeviceInfoOssUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponseBody) SetRequestId(v string) *GetExportDeviceInfoOssUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetExportDeviceInfoOssUrlResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetExportDeviceInfoOssUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetExportDeviceInfoOssUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExportDeviceInfoOssUrlResponseBodyData) SetUrl(v string) *GetExportDeviceInfoOssUrlResponseBodyData {
	s.Url = &v
	return s
}

type GetExportDeviceInfoOssUrlResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExportDeviceInfoOssUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExportDeviceInfoOssUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExportDeviceInfoOssUrlResponse) GoString() string {
	return s.String()
}

func (s *GetExportDeviceInfoOssUrlResponse) SetHeaders(v map[string]*string) *GetExportDeviceInfoOssUrlResponse {
	s.Headers = v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponse) SetStatusCode(v int32) *GetExportDeviceInfoOssUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponse) SetBody(v *GetExportDeviceInfoOssUrlResponseBody) *GetExportDeviceInfoOssUrlResponse {
	s.Body = v
	return s
}

type GetFbOssConfigRequest struct {
	DirPrefix       *string `json:"DirPrefix,omitempty" xml:"DirPrefix,omitempty"`
	IsDedicatedLine *int32  `json:"IsDedicatedLine,omitempty" xml:"IsDedicatedLine,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetFbOssConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFbOssConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFbOssConfigRequest) SetDirPrefix(v string) *GetFbOssConfigRequest {
	s.DirPrefix = &v
	return s
}

func (s *GetFbOssConfigRequest) SetIsDedicatedLine(v int32) *GetFbOssConfigRequest {
	s.IsDedicatedLine = &v
	return s
}

func (s *GetFbOssConfigRequest) SetRegion(v string) *GetFbOssConfigRequest {
	s.Region = &v
	return s
}

type GetFbOssConfigResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetFbOssConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFbOssConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFbOssConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetFbOssConfigResponseBody) SetCode(v string) *GetFbOssConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetFbOssConfigResponseBody) SetData(v *GetFbOssConfigResponseBodyData) *GetFbOssConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetFbOssConfigResponseBody) SetMessage(v string) *GetFbOssConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetFbOssConfigResponseBody) SetRequestId(v string) *GetFbOssConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetFbOssConfigResponseBodyData struct {
	AccessKeyId  *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	EndPoint     *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	OssPolicy    *string `json:"OssPolicy,omitempty" xml:"OssPolicy,omitempty"`
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetFbOssConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFbOssConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFbOssConfigResponseBodyData) SetAccessKeyId(v string) *GetFbOssConfigResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) SetEndPoint(v string) *GetFbOssConfigResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) SetOssPolicy(v string) *GetFbOssConfigResponseBodyData {
	s.OssPolicy = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) SetOssSignature(v string) *GetFbOssConfigResponseBodyData {
	s.OssSignature = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) SetSessionId(v string) *GetFbOssConfigResponseBodyData {
	s.SessionId = &v
	return s
}

type GetFbOssConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFbOssConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFbOssConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFbOssConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFbOssConfigResponse) SetHeaders(v map[string]*string) *GetFbOssConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFbOssConfigResponse) SetStatusCode(v int32) *GetFbOssConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFbOssConfigResponse) SetBody(v *GetFbOssConfigResponseBody) *GetFbOssConfigResponse {
	s.Body = v
	return s
}

type GetOssConfigRequest struct {
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetOssConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssConfigRequest) GoString() string {
	return s.String()
}

func (s *GetOssConfigRequest) SetType(v int32) *GetOssConfigRequest {
	s.Type = &v
	return s
}

type GetOssConfigResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOssConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssConfigResponseBody) SetCode(v string) *GetOssConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssConfigResponseBody) SetData(v *GetOssConfigResponseBodyData) *GetOssConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetOssConfigResponseBody) SetMessage(v string) *GetOssConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetOssConfigResponseBody) SetRequestId(v string) *GetOssConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetOssConfigResponseBodyData struct {
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	EndPoint      *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	OssPolicy     *string `json:"OssPolicy,omitempty" xml:"OssPolicy,omitempty"`
	OssSignature  *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetOssConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOssConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssConfigResponseBodyData) SetAccessKeyId(v string) *GetOssConfigResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetOssConfigResponseBodyData) SetEndPoint(v string) *GetOssConfigResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *GetOssConfigResponseBodyData) SetOssPolicy(v string) *GetOssConfigResponseBodyData {
	s.OssPolicy = &v
	return s
}

func (s *GetOssConfigResponseBodyData) SetOssSignature(v string) *GetOssConfigResponseBodyData {
	s.OssSignature = &v
	return s
}

func (s *GetOssConfigResponseBodyData) SetSecurityToken(v string) *GetOssConfigResponseBodyData {
	s.SecurityToken = &v
	return s
}

type GetOssConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssConfigResponse) GoString() string {
	return s.String()
}

func (s *GetOssConfigResponse) SetHeaders(v map[string]*string) *GetOssConfigResponse {
	s.Headers = v
	return s
}

func (s *GetOssConfigResponse) SetStatusCode(v int32) *GetOssConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssConfigResponse) SetBody(v *GetOssConfigResponseBody) *GetOssConfigResponse {
	s.Body = v
	return s
}

type GetVersionDownloadUrlRequest struct {
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetVersionDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVersionDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetVersionDownloadUrlRequest) SetVersionName(v string) *GetVersionDownloadUrlRequest {
	s.VersionName = &v
	return s
}

type GetVersionDownloadUrlResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetVersionDownloadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVersionDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVersionDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetVersionDownloadUrlResponseBody) SetCode(v string) *GetVersionDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetVersionDownloadUrlResponseBody) SetData(v *GetVersionDownloadUrlResponseBodyData) *GetVersionDownloadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetVersionDownloadUrlResponseBody) SetMessage(v string) *GetVersionDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetVersionDownloadUrlResponseBody) SetRequestId(v string) *GetVersionDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetVersionDownloadUrlResponseBodyData struct {
	FullDownloadUrl *string `json:"FullDownloadUrl,omitempty" xml:"FullDownloadUrl,omitempty"`
}

func (s GetVersionDownloadUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVersionDownloadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVersionDownloadUrlResponseBodyData) SetFullDownloadUrl(v string) *GetVersionDownloadUrlResponseBodyData {
	s.FullDownloadUrl = &v
	return s
}

type GetVersionDownloadUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVersionDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVersionDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVersionDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetVersionDownloadUrlResponse) SetHeaders(v map[string]*string) *GetVersionDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetVersionDownloadUrlResponse) SetStatusCode(v int32) *GetVersionDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVersionDownloadUrlResponse) SetBody(v *GetVersionDownloadUrlResponseBody) *GetVersionDownloadUrlResponse {
	s.Body = v
	return s
}

type ListDeviceOtaTaskByTenantRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceOtaTaskByTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantRequest) SetPageNumber(v int32) *ListDeviceOtaTaskByTenantRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantRequest) SetPageSize(v int32) *ListDeviceOtaTaskByTenantRequest {
	s.PageSize = &v
	return s
}

type ListDeviceOtaTaskByTenantResponseBody struct {
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *ListDeviceOtaTaskByTenantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeviceOtaTaskByTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetCode(v string) *ListDeviceOtaTaskByTenantResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetData(v *ListDeviceOtaTaskByTenantResponseBodyData) *ListDeviceOtaTaskByTenantResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetMessage(v string) *ListDeviceOtaTaskByTenantResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetPageNumber(v int32) *ListDeviceOtaTaskByTenantResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetPageSize(v int32) *ListDeviceOtaTaskByTenantResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetRequestId(v string) *ListDeviceOtaTaskByTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBody) SetTotalCount(v int64) *ListDeviceOtaTaskByTenantResponseBody {
	s.TotalCount = &v
	return s
}

type ListDeviceOtaTaskByTenantResponseBodyData struct {
	TenantDeviceOtaTasks []*ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks `json:"TenantDeviceOtaTasks,omitempty" xml:"TenantDeviceOtaTasks,omitempty" type:"Repeated"`
}

func (s ListDeviceOtaTaskByTenantResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantResponseBodyData) SetTenantDeviceOtaTasks(v []*ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) *ListDeviceOtaTaskByTenantResponseBodyData {
	s.TenantDeviceOtaTasks = v
	return s
}

type ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks struct {
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OperationStatus *int32  `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	PublishTime     *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UpgradeCount    *int64  `json:"UpgradeCount,omitempty" xml:"UpgradeCount,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetModel(v string) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.Model = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetOperationStatus(v int32) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.OperationStatus = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetPublishTime(v string) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.PublishTime = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetStatus(v int32) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.Status = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetTaskId(v int32) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.TaskId = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetUpgradeCount(v int64) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.UpgradeCount = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks) SetVersion(v string) *ListDeviceOtaTaskByTenantResponseBodyDataTenantDeviceOtaTasks {
	s.Version = &v
	return s
}

type ListDeviceOtaTaskByTenantResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceOtaTaskByTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceOtaTaskByTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantResponse) SetHeaders(v map[string]*string) *ListDeviceOtaTaskByTenantResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponse) SetStatusCode(v int32) *ListDeviceOtaTaskByTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponse) SetBody(v *ListDeviceOtaTaskByTenantResponseBody) *ListDeviceOtaTaskByTenantResponse {
	s.Body = v
	return s
}

type ListDeviceSeatsRequest struct {
	Label        *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	SeatNo       *string   `json:"SeatNo,omitempty" xml:"SeatNo,omitempty"`
	SerialNoList []*string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty" type:"Repeated"`
	TenantId     *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId       *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListDeviceSeatsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceSeatsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsRequest) SetLabel(v string) *ListDeviceSeatsRequest {
	s.Label = &v
	return s
}

func (s *ListDeviceSeatsRequest) SetSeatNo(v string) *ListDeviceSeatsRequest {
	s.SeatNo = &v
	return s
}

func (s *ListDeviceSeatsRequest) SetSerialNoList(v []*string) *ListDeviceSeatsRequest {
	s.SerialNoList = v
	return s
}

func (s *ListDeviceSeatsRequest) SetTenantId(v string) *ListDeviceSeatsRequest {
	s.TenantId = &v
	return s
}

func (s *ListDeviceSeatsRequest) SetZoneId(v string) *ListDeviceSeatsRequest {
	s.ZoneId = &v
	return s
}

type ListDeviceSeatsResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDeviceSeatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceSeatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsResponseBody) SetCode(v string) *ListDeviceSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceSeatsResponseBody) SetData(v *ListDeviceSeatsResponseBodyData) *ListDeviceSeatsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceSeatsResponseBody) SetMessage(v string) *ListDeviceSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceSeatsResponseBody) SetRequestId(v string) *ListDeviceSeatsResponseBody {
	s.RequestId = &v
	return s
}

type ListDeviceSeatsResponseBodyData struct {
	DeviceSeatDTOList []*ListDeviceSeatsResponseBodyDataDeviceSeatDTOList `json:"DeviceSeatDTOList,omitempty" xml:"DeviceSeatDTOList,omitempty" type:"Repeated"`
	TotalCount        *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeviceSeatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceSeatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsResponseBodyData) SetDeviceSeatDTOList(v []*ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) *ListDeviceSeatsResponseBodyData {
	s.DeviceSeatDTOList = v
	return s
}

func (s *ListDeviceSeatsResponseBodyData) SetTotalCount(v int64) *ListDeviceSeatsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListDeviceSeatsResponseBodyDataDeviceSeatDTOList struct {
	Label    *string `json:"Label,omitempty" xml:"Label,omitempty"`
	SeatName *string `json:"SeatName,omitempty" xml:"SeatName,omitempty"`
	SeatNo   *string `json:"SeatNo,omitempty" xml:"SeatNo,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SiteId   *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetLabel(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.Label = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSeatName(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SeatName = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSeatNo(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SeatNo = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSerialNo(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SerialNo = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSiteId(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SiteId = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSiteName(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SiteName = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetZoneId(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.ZoneId = &v
	return s
}

type ListDeviceSeatsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceSeatsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceSeatsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsResponse) SetHeaders(v map[string]*string) *ListDeviceSeatsResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceSeatsResponse) SetStatusCode(v int32) *ListDeviceSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceSeatsResponse) SetBody(v *ListDeviceSeatsResponseBody) *ListDeviceSeatsResponse {
	s.Body = v
	return s
}

type ListDevicesRequest struct {
	Alias          *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BuildId        *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientType     *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	DeviceGroupId  *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceIpV4     *string `json:"DeviceIpV4,omitempty" xml:"DeviceIpV4,omitempty"`
	DeviceName     *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceOS       *string `json:"DeviceOS,omitempty" xml:"DeviceOS,omitempty"`
	DevicePlatform *string `json:"DevicePlatform,omitempty" xml:"DevicePlatform,omitempty"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LabelContent   *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId        *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LocationInfo   *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	Model          *string `json:"Model,omitempty" xml:"Model,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SerialNo       *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	UserType       *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	Uuid           *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) SetAlias(v string) *ListDevicesRequest {
	s.Alias = &v
	return s
}

func (s *ListDevicesRequest) SetBuildId(v string) *ListDevicesRequest {
	s.BuildId = &v
	return s
}

func (s *ListDevicesRequest) SetClientType(v int32) *ListDevicesRequest {
	s.ClientType = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceGroupId(v string) *ListDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceIpV4(v string) *ListDevicesRequest {
	s.DeviceIpV4 = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceName(v string) *ListDevicesRequest {
	s.DeviceName = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceOS(v string) *ListDevicesRequest {
	s.DeviceOS = &v
	return s
}

func (s *ListDevicesRequest) SetDevicePlatform(v string) *ListDevicesRequest {
	s.DevicePlatform = &v
	return s
}

func (s *ListDevicesRequest) SetEndUserId(v string) *ListDevicesRequest {
	s.EndUserId = &v
	return s
}

func (s *ListDevicesRequest) SetLabelContent(v string) *ListDevicesRequest {
	s.LabelContent = &v
	return s
}

func (s *ListDevicesRequest) SetLabelId(v string) *ListDevicesRequest {
	s.LabelId = &v
	return s
}

func (s *ListDevicesRequest) SetLocationInfo(v string) *ListDevicesRequest {
	s.LocationInfo = &v
	return s
}

func (s *ListDevicesRequest) SetModel(v string) *ListDevicesRequest {
	s.Model = &v
	return s
}

func (s *ListDevicesRequest) SetPageNumber(v int32) *ListDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesRequest) SetPageSize(v int32) *ListDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevicesRequest) SetSerialNo(v string) *ListDevicesRequest {
	s.SerialNo = &v
	return s
}

func (s *ListDevicesRequest) SetUserType(v string) *ListDevicesRequest {
	s.UserType = &v
	return s
}

func (s *ListDevicesRequest) SetUuid(v string) *ListDevicesRequest {
	s.Uuid = &v
	return s
}

type ListDevicesResponseBody struct {
	Code       *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) SetCode(v string) *ListDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDevicesResponseBody) SetData(v []*ListDevicesResponseBodyData) *ListDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListDevicesResponseBody) SetMessage(v string) *ListDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevicesResponseBody) SetPageNumber(v int32) *ListDevicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesResponseBody) SetPageSize(v int32) *ListDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponseBody) SetTotalCount(v int64) *ListDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListDevicesResponseBodyData struct {
	ActiveTime                 *string                                           `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	Alias                      *string                                           `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AutoLockScreenTime         *int32                                            `json:"AutoLockScreenTime,omitempty" xml:"AutoLockScreenTime,omitempty"`
	AutoLogin                  *int32                                            `json:"AutoLogin,omitempty" xml:"AutoLogin,omitempty"`
	AutoType                   *string                                           `json:"AutoType,omitempty" xml:"AutoType,omitempty"`
	Bluetooth                  *string                                           `json:"Bluetooth,omitempty" xml:"Bluetooth,omitempty"`
	BuildId                    *string                                           `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientId                   *string                                           `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType                 *string                                           `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ConnectConfigs             []*ListDevicesResponseBodyDataConnectConfigs      `json:"ConnectConfigs,omitempty" xml:"ConnectConfigs,omitempty" type:"Repeated"`
	CustomIdleAction           *int32                                            `json:"CustomIdleAction,omitempty" xml:"CustomIdleAction,omitempty"`
	CustomPowerOn              *int32                                            `json:"CustomPowerOn,omitempty" xml:"CustomPowerOn,omitempty"`
	CustomProperty             *string                                           `json:"CustomProperty,omitempty" xml:"CustomProperty,omitempty"`
	CustomResourcePackage      *ListDevicesResponseBodyDataCustomResourcePackage `json:"CustomResourcePackage,omitempty" xml:"CustomResourcePackage,omitempty" type:"Struct"`
	DefinePowerButton          *int32                                            `json:"DefinePowerButton,omitempty" xml:"DefinePowerButton,omitempty"`
	DeviceIpV4                 *string                                           `json:"DeviceIpV4,omitempty" xml:"DeviceIpV4,omitempty"`
	DeviceLock                 *int32                                            `json:"DeviceLock,omitempty" xml:"DeviceLock,omitempty"`
	DeviceMqttConnectionStatus *int32                                            `json:"DeviceMqttConnectionStatus,omitempty" xml:"DeviceMqttConnectionStatus,omitempty"`
	DeviceName                 *string                                           `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceOS                   *string                                           `json:"DeviceOS,omitempty" xml:"DeviceOS,omitempty"`
	DevicePlatform             *string                                           `json:"DevicePlatform,omitempty" xml:"DevicePlatform,omitempty"`
	DisplayLayout              *string                                           `json:"DisplayLayout,omitempty" xml:"DisplayLayout,omitempty"`
	DisplayResolution          *string                                           `json:"DisplayResolution,omitempty" xml:"DisplayResolution,omitempty"`
	DisplayScaleRatio          *string                                           `json:"DisplayScaleRatio,omitempty" xml:"DisplayScaleRatio,omitempty"`
	EnableAdb                  *int32                                            `json:"EnableAdb,omitempty" xml:"EnableAdb,omitempty"`
	EnableAutoLockScreen       *int32                                            `json:"EnableAutoLockScreen,omitempty" xml:"EnableAutoLockScreen,omitempty"`
	EnableBluetooth            *int32                                            `json:"EnableBluetooth,omitempty" xml:"EnableBluetooth,omitempty"`
	EnableLockScreenPassword   *int32                                            `json:"EnableLockScreenPassword,omitempty" xml:"EnableLockScreenPassword,omitempty"`
	EnableModifyPassword       *int32                                            `json:"EnableModifyPassword,omitempty" xml:"EnableModifyPassword,omitempty"`
	EnableScheduledPowerOff    *int32                                            `json:"EnableScheduledPowerOff,omitempty" xml:"EnableScheduledPowerOff,omitempty"`
	EnableUnlockPassword       *int32                                            `json:"EnableUnlockPassword,omitempty" xml:"EnableUnlockPassword,omitempty"`
	EnableWlan                 *int32                                            `json:"EnableWlan,omitempty" xml:"EnableWlan,omitempty"`
	EndUserList                []*ListDevicesResponseBodyDataEndUserList         `json:"EndUserList,omitempty" xml:"EndUserList,omitempty" type:"Repeated"`
	EtherMac                   *string                                           `json:"EtherMac,omitempty" xml:"EtherMac,omitempty"`
	GmtModified                *string                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtSync                    *string                                           `json:"GmtSync,omitempty" xml:"GmtSync,omitempty"`
	Id                         *int64                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	IdleTime                   *int32                                            `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	IsActive                   *string                                           `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	LabelList                  []*ListDevicesResponseBodyDataLabelList           `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	LastLoginUser              *string                                           `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	LocalUsbPrint              *int32                                            `json:"LocalUsbPrint,omitempty" xml:"LocalUsbPrint,omitempty"`
	LocationInfo               *string                                           `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	LockPassword               *string                                           `json:"LockPassword,omitempty" xml:"LockPassword,omitempty"`
	Model                      *string                                           `json:"Model,omitempty" xml:"Model,omitempty"`
	OrderId                    *string                                           `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PeripheralConfig           *ListDevicesResponseBodyDataPeripheralConfig      `json:"PeripheralConfig,omitempty" xml:"PeripheralConfig,omitempty" type:"Struct"`
	ScheduledPowerOff          *string                                           `json:"ScheduledPowerOff,omitempty" xml:"ScheduledPowerOff,omitempty"`
	SecureNetworkType          *string                                           `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	SerialNo                   *string                                           `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SleepTime                  *int32                                            `json:"SleepTime,omitempty" xml:"SleepTime,omitempty"`
	Source                     *string                                           `json:"Source,omitempty" xml:"Source,omitempty"`
	TenantId                   *string                                           `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UsbStorage                 *int32                                            `json:"UsbStorage,omitempty" xml:"UsbStorage,omitempty"`
	Uuid                       *string                                           `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Wlan                       *string                                           `json:"Wlan,omitempty" xml:"Wlan,omitempty"`
}

func (s ListDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyData) SetActiveTime(v string) *ListDevicesResponseBodyData {
	s.ActiveTime = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetAlias(v string) *ListDevicesResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetAutoLockScreenTime(v int32) *ListDevicesResponseBodyData {
	s.AutoLockScreenTime = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetAutoLogin(v int32) *ListDevicesResponseBodyData {
	s.AutoLogin = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetAutoType(v string) *ListDevicesResponseBodyData {
	s.AutoType = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetBluetooth(v string) *ListDevicesResponseBodyData {
	s.Bluetooth = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetBuildId(v string) *ListDevicesResponseBodyData {
	s.BuildId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetClientId(v string) *ListDevicesResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetClientType(v string) *ListDevicesResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetConnectConfigs(v []*ListDevicesResponseBodyDataConnectConfigs) *ListDevicesResponseBodyData {
	s.ConnectConfigs = v
	return s
}

func (s *ListDevicesResponseBodyData) SetCustomIdleAction(v int32) *ListDevicesResponseBodyData {
	s.CustomIdleAction = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetCustomPowerOn(v int32) *ListDevicesResponseBodyData {
	s.CustomPowerOn = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetCustomProperty(v string) *ListDevicesResponseBodyData {
	s.CustomProperty = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetCustomResourcePackage(v *ListDevicesResponseBodyDataCustomResourcePackage) *ListDevicesResponseBodyData {
	s.CustomResourcePackage = v
	return s
}

func (s *ListDevicesResponseBodyData) SetDefinePowerButton(v int32) *ListDevicesResponseBodyData {
	s.DefinePowerButton = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceIpV4(v string) *ListDevicesResponseBodyData {
	s.DeviceIpV4 = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceLock(v int32) *ListDevicesResponseBodyData {
	s.DeviceLock = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceMqttConnectionStatus(v int32) *ListDevicesResponseBodyData {
	s.DeviceMqttConnectionStatus = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceName(v string) *ListDevicesResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceOS(v string) *ListDevicesResponseBodyData {
	s.DeviceOS = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDevicePlatform(v string) *ListDevicesResponseBodyData {
	s.DevicePlatform = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDisplayLayout(v string) *ListDevicesResponseBodyData {
	s.DisplayLayout = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDisplayResolution(v string) *ListDevicesResponseBodyData {
	s.DisplayResolution = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDisplayScaleRatio(v string) *ListDevicesResponseBodyData {
	s.DisplayScaleRatio = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableAdb(v int32) *ListDevicesResponseBodyData {
	s.EnableAdb = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableAutoLockScreen(v int32) *ListDevicesResponseBodyData {
	s.EnableAutoLockScreen = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableBluetooth(v int32) *ListDevicesResponseBodyData {
	s.EnableBluetooth = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableLockScreenPassword(v int32) *ListDevicesResponseBodyData {
	s.EnableLockScreenPassword = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableModifyPassword(v int32) *ListDevicesResponseBodyData {
	s.EnableModifyPassword = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableScheduledPowerOff(v int32) *ListDevicesResponseBodyData {
	s.EnableScheduledPowerOff = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableUnlockPassword(v int32) *ListDevicesResponseBodyData {
	s.EnableUnlockPassword = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableWlan(v int32) *ListDevicesResponseBodyData {
	s.EnableWlan = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEndUserList(v []*ListDevicesResponseBodyDataEndUserList) *ListDevicesResponseBodyData {
	s.EndUserList = v
	return s
}

func (s *ListDevicesResponseBodyData) SetEtherMac(v string) *ListDevicesResponseBodyData {
	s.EtherMac = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetGmtModified(v string) *ListDevicesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetGmtSync(v string) *ListDevicesResponseBodyData {
	s.GmtSync = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetId(v int64) *ListDevicesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetIdleTime(v int32) *ListDevicesResponseBodyData {
	s.IdleTime = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetIsActive(v string) *ListDevicesResponseBodyData {
	s.IsActive = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetLabelList(v []*ListDevicesResponseBodyDataLabelList) *ListDevicesResponseBodyData {
	s.LabelList = v
	return s
}

func (s *ListDevicesResponseBodyData) SetLastLoginUser(v string) *ListDevicesResponseBodyData {
	s.LastLoginUser = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetLocalUsbPrint(v int32) *ListDevicesResponseBodyData {
	s.LocalUsbPrint = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetLocationInfo(v string) *ListDevicesResponseBodyData {
	s.LocationInfo = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetLockPassword(v string) *ListDevicesResponseBodyData {
	s.LockPassword = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetModel(v string) *ListDevicesResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetOrderId(v string) *ListDevicesResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetPeripheralConfig(v *ListDevicesResponseBodyDataPeripheralConfig) *ListDevicesResponseBodyData {
	s.PeripheralConfig = v
	return s
}

func (s *ListDevicesResponseBodyData) SetScheduledPowerOff(v string) *ListDevicesResponseBodyData {
	s.ScheduledPowerOff = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetSecureNetworkType(v string) *ListDevicesResponseBodyData {
	s.SecureNetworkType = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetSerialNo(v string) *ListDevicesResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetSleepTime(v int32) *ListDevicesResponseBodyData {
	s.SleepTime = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetSource(v string) *ListDevicesResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetTenantId(v string) *ListDevicesResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetUsbStorage(v int32) *ListDevicesResponseBodyData {
	s.UsbStorage = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetUuid(v string) *ListDevicesResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetWlan(v string) *ListDevicesResponseBodyData {
	s.Wlan = &v
	return s
}

type ListDevicesResponseBodyDataConnectConfigs struct {
	ConnectScript  *string `json:"ConnectScript,omitempty" xml:"ConnectScript,omitempty"`
	PeripheralPid  *string `json:"PeripheralPid,omitempty" xml:"PeripheralPid,omitempty"`
	PeripheralVid  *string `json:"PeripheralVid,omitempty" xml:"PeripheralVid,omitempty"`
	RedirectPolicy *int32  `json:"RedirectPolicy,omitempty" xml:"RedirectPolicy,omitempty"`
}

func (s ListDevicesResponseBodyDataConnectConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDataConnectConfigs) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataConnectConfigs) SetConnectScript(v string) *ListDevicesResponseBodyDataConnectConfigs {
	s.ConnectScript = &v
	return s
}

func (s *ListDevicesResponseBodyDataConnectConfigs) SetPeripheralPid(v string) *ListDevicesResponseBodyDataConnectConfigs {
	s.PeripheralPid = &v
	return s
}

func (s *ListDevicesResponseBodyDataConnectConfigs) SetPeripheralVid(v string) *ListDevicesResponseBodyDataConnectConfigs {
	s.PeripheralVid = &v
	return s
}

func (s *ListDevicesResponseBodyDataConnectConfigs) SetRedirectPolicy(v int32) *ListDevicesResponseBodyDataConnectConfigs {
	s.RedirectPolicy = &v
	return s
}

type ListDevicesResponseBodyDataCustomResourcePackage struct {
	ConfigAboutLogo     *string `json:"ConfigAboutLogo,omitempty" xml:"ConfigAboutLogo,omitempty"`
	DesktopWallpaper    *string `json:"DesktopWallpaper,omitempty" xml:"DesktopWallpaper,omitempty"`
	LoginPageBackground *string `json:"LoginPageBackground,omitempty" xml:"LoginPageBackground,omitempty"`
	LoginPageLogo       *string `json:"LoginPageLogo,omitempty" xml:"LoginPageLogo,omitempty"`
	PersonalCenterLogo  *string `json:"PersonalCenterLogo,omitempty" xml:"PersonalCenterLogo,omitempty"`
	StartLogo           *string `json:"StartLogo,omitempty" xml:"StartLogo,omitempty"`
	StartMenuLogo       *string `json:"StartMenuLogo,omitempty" xml:"StartMenuLogo,omitempty"`
	UpgradeLogo         *string `json:"UpgradeLogo,omitempty" xml:"UpgradeLogo,omitempty"`
}

func (s ListDevicesResponseBodyDataCustomResourcePackage) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDataCustomResourcePackage) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetConfigAboutLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.ConfigAboutLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetDesktopWallpaper(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.DesktopWallpaper = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetLoginPageBackground(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.LoginPageBackground = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetLoginPageLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.LoginPageLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetPersonalCenterLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.PersonalCenterLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetStartLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.StartLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetStartMenuLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.StartMenuLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetUpgradeLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.UpgradeLogo = &v
	return s
}

type ListDevicesResponseBodyDataEndUserList struct {
	AdDomain    *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	BindTime    *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId   *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SerialNo    *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserType    *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListDevicesResponseBodyDataEndUserList) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDataEndUserList) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataEndUserList) SetAdDomain(v string) *ListDevicesResponseBodyDataEndUserList {
	s.AdDomain = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetBindTime(v string) *ListDevicesResponseBodyDataEndUserList {
	s.BindTime = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetDirectoryId(v string) *ListDevicesResponseBodyDataEndUserList {
	s.DirectoryId = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetEndUserId(v string) *ListDevicesResponseBodyDataEndUserList {
	s.EndUserId = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetId(v int64) *ListDevicesResponseBodyDataEndUserList {
	s.Id = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetSerialNo(v string) *ListDevicesResponseBodyDataEndUserList {
	s.SerialNo = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetTenantId(v string) *ListDevicesResponseBodyDataEndUserList {
	s.TenantId = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetUserType(v string) *ListDevicesResponseBodyDataEndUserList {
	s.UserType = &v
	return s
}

type ListDevicesResponseBodyDataLabelList struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LabelId     *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDevicesResponseBodyDataLabelList) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDataLabelList) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataLabelList) SetContent(v string) *ListDevicesResponseBodyDataLabelList {
	s.Content = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) SetGmtCreate(v int64) *ListDevicesResponseBodyDataLabelList {
	s.GmtCreate = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) SetGmtModified(v int64) *ListDevicesResponseBodyDataLabelList {
	s.GmtModified = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) SetLabelId(v string) *ListDevicesResponseBodyDataLabelList {
	s.LabelId = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) SetTenantId(v string) *ListDevicesResponseBodyDataLabelList {
	s.TenantId = &v
	return s
}

type ListDevicesResponseBodyDataPeripheralConfig struct {
	DefaultPolicy        *int32 `json:"DefaultPolicy,omitempty" xml:"DefaultPolicy,omitempty"`
	PolicyStrategy       *int32 `json:"PolicyStrategy,omitempty" xml:"PolicyStrategy,omitempty"`
	UsbAndInternalCamera *int32 `json:"UsbAndInternalCamera,omitempty" xml:"UsbAndInternalCamera,omitempty"`
	UsbPrinter           *int32 `json:"UsbPrinter,omitempty" xml:"UsbPrinter,omitempty"`
	UsbStorage           *int32 `json:"UsbStorage,omitempty" xml:"UsbStorage,omitempty"`
}

func (s ListDevicesResponseBodyDataPeripheralConfig) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDataPeripheralConfig) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetDefaultPolicy(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.DefaultPolicy = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetPolicyStrategy(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.PolicyStrategy = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetUsbAndInternalCamera(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.UsbAndInternalCamera = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetUsbPrinter(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.UsbPrinter = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetUsbStorage(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.UsbStorage = &v
	return s
}

type ListDevicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListDevicesResponse) SetHeaders(v map[string]*string) *ListDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListDevicesResponse) SetStatusCode(v int32) *ListDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevicesResponse) SetBody(v *ListDevicesResponseBody) *ListDevicesResponse {
	s.Body = v
	return s
}

type ListFbIssueLabelsResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListFbIssueLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFbIssueLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFbIssueLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsResponseBody) SetCode(v string) *ListFbIssueLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFbIssueLabelsResponseBody) SetData(v *ListFbIssueLabelsResponseBodyData) *ListFbIssueLabelsResponseBody {
	s.Data = v
	return s
}

func (s *ListFbIssueLabelsResponseBody) SetMessage(v string) *ListFbIssueLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFbIssueLabelsResponseBody) SetRequestId(v string) *ListFbIssueLabelsResponseBody {
	s.RequestId = &v
	return s
}

type ListFbIssueLabelsResponseBodyData struct {
	IssueLabel []*string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty" type:"Repeated"`
}

func (s ListFbIssueLabelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFbIssueLabelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsResponseBodyData) SetIssueLabel(v []*string) *ListFbIssueLabelsResponseBodyData {
	s.IssueLabel = v
	return s
}

type ListFbIssueLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFbIssueLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFbIssueLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFbIssueLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsResponse) SetHeaders(v map[string]*string) *ListFbIssueLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListFbIssueLabelsResponse) SetStatusCode(v int32) *ListFbIssueLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFbIssueLabelsResponse) SetBody(v *ListFbIssueLabelsResponseBody) *ListFbIssueLabelsResponse {
	s.Body = v
	return s
}

type ListFbIssueLabelsByLCRequest struct {
	Caller       *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	LanguageType *string `json:"LanguageType,omitempty" xml:"LanguageType,omitempty"`
}

func (s ListFbIssueLabelsByLCRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFbIssueLabelsByLCRequest) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsByLCRequest) SetCaller(v string) *ListFbIssueLabelsByLCRequest {
	s.Caller = &v
	return s
}

func (s *ListFbIssueLabelsByLCRequest) SetLanguageType(v string) *ListFbIssueLabelsByLCRequest {
	s.LanguageType = &v
	return s
}

type ListFbIssueLabelsByLCResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListFbIssueLabelsByLCResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFbIssueLabelsByLCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFbIssueLabelsByLCResponseBody) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsByLCResponseBody) SetCode(v string) *ListFbIssueLabelsByLCResponseBody {
	s.Code = &v
	return s
}

func (s *ListFbIssueLabelsByLCResponseBody) SetData(v *ListFbIssueLabelsByLCResponseBodyData) *ListFbIssueLabelsByLCResponseBody {
	s.Data = v
	return s
}

func (s *ListFbIssueLabelsByLCResponseBody) SetMessage(v string) *ListFbIssueLabelsByLCResponseBody {
	s.Message = &v
	return s
}

func (s *ListFbIssueLabelsByLCResponseBody) SetRequestId(v string) *ListFbIssueLabelsByLCResponseBody {
	s.RequestId = &v
	return s
}

type ListFbIssueLabelsByLCResponseBodyData struct {
	IssueLabel []*string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty" type:"Repeated"`
}

func (s ListFbIssueLabelsByLCResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFbIssueLabelsByLCResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsByLCResponseBodyData) SetIssueLabel(v []*string) *ListFbIssueLabelsByLCResponseBodyData {
	s.IssueLabel = v
	return s
}

type ListFbIssueLabelsByLCResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFbIssueLabelsByLCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFbIssueLabelsByLCResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFbIssueLabelsByLCResponse) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsByLCResponse) SetHeaders(v map[string]*string) *ListFbIssueLabelsByLCResponse {
	s.Headers = v
	return s
}

func (s *ListFbIssueLabelsByLCResponse) SetStatusCode(v int32) *ListFbIssueLabelsByLCResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFbIssueLabelsByLCResponse) SetBody(v *ListFbIssueLabelsByLCResponseBody) *ListFbIssueLabelsByLCResponse {
	s.Body = v
	return s
}

type ListLabelsRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListLabelsRequest) SetLabelContent(v string) *ListLabelsRequest {
	s.LabelContent = &v
	return s
}

func (s *ListLabelsRequest) SetLabelId(v string) *ListLabelsRequest {
	s.LabelId = &v
	return s
}

func (s *ListLabelsRequest) SetMaxResults(v int32) *ListLabelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLabelsRequest) SetNextToken(v string) *ListLabelsRequest {
	s.NextToken = &v
	return s
}

type ListLabelsResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLabelsResponseBody) SetCode(v string) *ListLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLabelsResponseBody) SetData(v []*ListLabelsResponseBodyData) *ListLabelsResponseBody {
	s.Data = v
	return s
}

func (s *ListLabelsResponseBody) SetMessage(v string) *ListLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLabelsResponseBody) SetNextToken(v string) *ListLabelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLabelsResponseBody) SetRequestId(v string) *ListLabelsResponseBody {
	s.RequestId = &v
	return s
}

type ListLabelsResponseBodyData struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId     *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListLabelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLabelsResponseBodyData) SetContent(v string) *ListLabelsResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetGmtCreate(v int64) *ListLabelsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetGmtModified(v int64) *ListLabelsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetId(v int64) *ListLabelsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetLabelId(v string) *ListLabelsResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetTenantId(v string) *ListLabelsResponseBodyData {
	s.TenantId = &v
	return s
}

type ListLabelsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListLabelsResponse) SetHeaders(v map[string]*string) *ListLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListLabelsResponse) SetStatusCode(v int32) *ListLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLabelsResponse) SetBody(v *ListLabelsResponseBody) *ListLabelsResponse {
	s.Body = v
	return s
}

type ListTenantDeviceOtaInfoRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTenantDeviceOtaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTenantDeviceOtaInfoRequest) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoRequest) SetPageNumber(v int32) *ListTenantDeviceOtaInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTenantDeviceOtaInfoRequest) SetPageSize(v int32) *ListTenantDeviceOtaInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListTenantDeviceOtaInfoRequest) SetTaskId(v int32) *ListTenantDeviceOtaInfoRequest {
	s.TaskId = &v
	return s
}

type ListTenantDeviceOtaInfoResponseBody struct {
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *ListTenantDeviceOtaInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTenantDeviceOtaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTenantDeviceOtaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetCode(v string) *ListTenantDeviceOtaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetData(v *ListTenantDeviceOtaInfoResponseBodyData) *ListTenantDeviceOtaInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetMessage(v string) *ListTenantDeviceOtaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetPageNumber(v int32) *ListTenantDeviceOtaInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetPageSize(v int32) *ListTenantDeviceOtaInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetRequestId(v string) *ListTenantDeviceOtaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetTotalCount(v int64) *ListTenantDeviceOtaInfoResponseBody {
	s.TotalCount = &v
	return s
}

type ListTenantDeviceOtaInfoResponseBodyData struct {
	TenantDeviceOtaInfos []*ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos `json:"TenantDeviceOtaInfos,omitempty" xml:"TenantDeviceOtaInfos,omitempty" type:"Repeated"`
}

func (s ListTenantDeviceOtaInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTenantDeviceOtaInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoResponseBodyData) SetTenantDeviceOtaInfos(v []*ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) *ListTenantDeviceOtaInfoResponseBodyData {
	s.TenantDeviceOtaInfos = v
	return s
}

type ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos struct {
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Model          *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) String() string {
	return tea.Prettify(s)
}

func (s ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) SetCurrentVersion(v string) *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos {
	s.CurrentVersion = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) SetDeviceId(v string) *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos {
	s.DeviceId = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) SetModel(v string) *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos {
	s.Model = &v
	return s
}

type ListTenantDeviceOtaInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantDeviceOtaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTenantDeviceOtaInfoResponse) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoResponse) SetHeaders(v map[string]*string) *ListTenantDeviceOtaInfoResponse {
	s.Headers = v
	return s
}

func (s *ListTenantDeviceOtaInfoResponse) SetStatusCode(v int32) *ListTenantDeviceOtaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponse) SetBody(v *ListTenantDeviceOtaInfoResponseBody) *ListTenantDeviceOtaInfoResponse {
	s.Body = v
	return s
}

type ListTerminalRequest struct {
	Alias           *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BuildId         *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientType      *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	InManage        *bool   `json:"InManage,omitempty" xml:"InManage,omitempty"`
	Ipv4            *string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	LocationInfo    *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SearchKeyword   *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	SerialNumber    *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTerminalRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalRequest) GoString() string {
	return s.String()
}

func (s *ListTerminalRequest) SetAlias(v string) *ListTerminalRequest {
	s.Alias = &v
	return s
}

func (s *ListTerminalRequest) SetBuildId(v string) *ListTerminalRequest {
	s.BuildId = &v
	return s
}

func (s *ListTerminalRequest) SetClientType(v int32) *ListTerminalRequest {
	s.ClientType = &v
	return s
}

func (s *ListTerminalRequest) SetInManage(v bool) *ListTerminalRequest {
	s.InManage = &v
	return s
}

func (s *ListTerminalRequest) SetIpv4(v string) *ListTerminalRequest {
	s.Ipv4 = &v
	return s
}

func (s *ListTerminalRequest) SetLocationInfo(v string) *ListTerminalRequest {
	s.LocationInfo = &v
	return s
}

func (s *ListTerminalRequest) SetMaxResults(v int32) *ListTerminalRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTerminalRequest) SetModel(v string) *ListTerminalRequest {
	s.Model = &v
	return s
}

func (s *ListTerminalRequest) SetNextToken(v string) *ListTerminalRequest {
	s.NextToken = &v
	return s
}

func (s *ListTerminalRequest) SetSearchKeyword(v string) *ListTerminalRequest {
	s.SearchKeyword = &v
	return s
}

func (s *ListTerminalRequest) SetSerialNumber(v string) *ListTerminalRequest {
	s.SerialNumber = &v
	return s
}

func (s *ListTerminalRequest) SetTerminalGroupId(v string) *ListTerminalRequest {
	s.TerminalGroupId = &v
	return s
}

func (s *ListTerminalRequest) SetUuid(v string) *ListTerminalRequest {
	s.Uuid = &v
	return s
}

type ListTerminalResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListTerminalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTerminalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *ListTerminalResponseBody) SetCode(v string) *ListTerminalResponseBody {
	s.Code = &v
	return s
}

func (s *ListTerminalResponseBody) SetData(v []*ListTerminalResponseBodyData) *ListTerminalResponseBody {
	s.Data = v
	return s
}

func (s *ListTerminalResponseBody) SetHttpStatusCode(v int32) *ListTerminalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTerminalResponseBody) SetMessage(v string) *ListTerminalResponseBody {
	s.Message = &v
	return s
}

func (s *ListTerminalResponseBody) SetNextToken(v string) *ListTerminalResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTerminalResponseBody) SetRequestId(v string) *ListTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTerminalResponseBody) SetSuccess(v bool) *ListTerminalResponseBody {
	s.Success = &v
	return s
}

func (s *ListTerminalResponseBody) SetTotalCount(v int32) *ListTerminalResponseBody {
	s.TotalCount = &v
	return s
}

type ListTerminalResponseBodyData struct {
	Alias         *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BindUserCount *int32  `json:"BindUserCount,omitempty" xml:"BindUserCount,omitempty"`
	BindUserId    *string `json:"BindUserId,omitempty" xml:"BindUserId,omitempty"`
	BuildId       *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientType    *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	InManage      *bool   `json:"InManage,omitempty" xml:"InManage,omitempty"`
	Ipv4          *string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	// Deprecated
	LastLoginUser   *string `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	LocationInfo    *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	LockSettings    *bool   `json:"LockSettings,omitempty" xml:"LockSettings,omitempty"`
	LoginUser       *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OnlineStatus    *bool   `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	SerialNumber    *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTerminalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTerminalResponseBodyData) SetAlias(v string) *ListTerminalResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetBindUserCount(v int32) *ListTerminalResponseBodyData {
	s.BindUserCount = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetBindUserId(v string) *ListTerminalResponseBodyData {
	s.BindUserId = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetBuildId(v string) *ListTerminalResponseBodyData {
	s.BuildId = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetClientType(v int32) *ListTerminalResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetDesktopId(v string) *ListTerminalResponseBodyData {
	s.DesktopId = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetInManage(v bool) *ListTerminalResponseBodyData {
	s.InManage = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetIpv4(v string) *ListTerminalResponseBodyData {
	s.Ipv4 = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetLastLoginUser(v string) *ListTerminalResponseBodyData {
	s.LastLoginUser = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetLocationInfo(v string) *ListTerminalResponseBodyData {
	s.LocationInfo = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetLockSettings(v bool) *ListTerminalResponseBodyData {
	s.LockSettings = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetLoginUser(v string) *ListTerminalResponseBodyData {
	s.LoginUser = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetModel(v string) *ListTerminalResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetOnlineStatus(v bool) *ListTerminalResponseBodyData {
	s.OnlineStatus = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetSerialNumber(v string) *ListTerminalResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetTerminalGroupId(v string) *ListTerminalResponseBodyData {
	s.TerminalGroupId = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetUuid(v string) *ListTerminalResponseBodyData {
	s.Uuid = &v
	return s
}

type ListTerminalResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTerminalResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalResponse) GoString() string {
	return s.String()
}

func (s *ListTerminalResponse) SetHeaders(v map[string]*string) *ListTerminalResponse {
	s.Headers = v
	return s
}

func (s *ListTerminalResponse) SetStatusCode(v int32) *ListTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTerminalResponse) SetBody(v *ListTerminalResponseBody) *ListTerminalResponse {
	s.Body = v
	return s
}

type ListTerminalsRequest struct {
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdEdsXbwG2ZlbWCzN4wTTg6wQvfp7u1BJl4bxCAby41POSaYAlCvfULQpkAnb0ff****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DemoDevice
	SearchKeyword *string   `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	SerialNumbers []*string `json:"SerialNumbers,omitempty" xml:"SerialNumbers,omitempty" type:"Repeated"`
	// example:
	//
	// tg-default
	TerminalGroupId *string   `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	Uuids           []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s ListTerminalsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalsRequest) GoString() string {
	return s.String()
}

func (s *ListTerminalsRequest) SetMaxResults(v int32) *ListTerminalsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTerminalsRequest) SetNextToken(v string) *ListTerminalsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTerminalsRequest) SetSearchKeyword(v string) *ListTerminalsRequest {
	s.SearchKeyword = &v
	return s
}

func (s *ListTerminalsRequest) SetSerialNumbers(v []*string) *ListTerminalsRequest {
	s.SerialNumbers = v
	return s
}

func (s *ListTerminalsRequest) SetTerminalGroupId(v string) *ListTerminalsRequest {
	s.TerminalGroupId = &v
	return s
}

func (s *ListTerminalsRequest) SetUuids(v []*string) *ListTerminalsRequest {
	s.Uuids = v
	return s
}

type ListTerminalsResponseBody struct {
	// example:
	//
	// TERMINAL_NOT_FOUND
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListTerminalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// terminal not found
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAdEdsXbwG2ZlbWCzN4wTTg6wQvfp7u1BJl4bxCAby41POSaYAlCvfULQpkAnb0ff****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// C5DCE54A-B266-522E-A6ED-468AF45F5AAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTerminalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTerminalsResponseBody) SetCode(v string) *ListTerminalsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTerminalsResponseBody) SetData(v []*ListTerminalsResponseBodyData) *ListTerminalsResponseBody {
	s.Data = v
	return s
}

func (s *ListTerminalsResponseBody) SetHttpStatusCode(v int32) *ListTerminalsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTerminalsResponseBody) SetMessage(v string) *ListTerminalsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTerminalsResponseBody) SetNextToken(v string) *ListTerminalsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTerminalsResponseBody) SetRequestId(v string) *ListTerminalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTerminalsResponseBody) SetSuccess(v bool) *ListTerminalsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTerminalsResponseBody) SetTotalCount(v int32) *ListTerminalsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTerminalsResponseBodyData struct {
	// example:
	//
	// DemoDevice
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// example:
	//
	// 7.0.2-RS-20240805.044924
	BuildId *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// ecd-drqmaogzbmbdf****
	CurrentConnectDesktop *string `json:"CurrentConnectDesktop,omitempty" xml:"CurrentConnectDesktop,omitempty"`
	// example:
	//
	// alice
	CurrentLoginUser *string `json:"CurrentLoginUser,omitempty" xml:"CurrentLoginUser,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	Ipv4         *string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	LocationInfo *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	ManageTime   *string `json:"ManageTime,omitempty" xml:"ManageTime,omitempty"`
	// example:
	//
	// US01
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// alice
	PasswordFreeLoginUser *string `json:"PasswordFreeLoginUser,omitempty" xml:"PasswordFreeLoginUser,omitempty"`
	// example:
	//
	// ODN49YQCPQYC****
	SerialNumber                 *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SetPasswordFreeLoginUserTime *string `json:"SetPasswordFreeLoginUserTime,omitempty" xml:"SetPasswordFreeLoginUserTime,omitempty"`
	// example:
	//
	// tg-default
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	// example:
	//
	// 04873D3898B51A7DF2455C1E1DC9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTerminalsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTerminalsResponseBodyData) SetAlias(v string) *ListTerminalsResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetBuildId(v string) *ListTerminalsResponseBodyData {
	s.BuildId = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetClientType(v int32) *ListTerminalsResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetCurrentConnectDesktop(v string) *ListTerminalsResponseBodyData {
	s.CurrentConnectDesktop = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetCurrentLoginUser(v string) *ListTerminalsResponseBodyData {
	s.CurrentLoginUser = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetIpv4(v string) *ListTerminalsResponseBodyData {
	s.Ipv4 = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetLocationInfo(v string) *ListTerminalsResponseBodyData {
	s.LocationInfo = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetManageTime(v string) *ListTerminalsResponseBodyData {
	s.ManageTime = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetModel(v string) *ListTerminalsResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetOnline(v bool) *ListTerminalsResponseBodyData {
	s.Online = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetPasswordFreeLoginUser(v string) *ListTerminalsResponseBodyData {
	s.PasswordFreeLoginUser = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetSerialNumber(v string) *ListTerminalsResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetSetPasswordFreeLoginUserTime(v string) *ListTerminalsResponseBodyData {
	s.SetPasswordFreeLoginUserTime = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetTerminalGroupId(v string) *ListTerminalsResponseBodyData {
	s.TerminalGroupId = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetUuid(v string) *ListTerminalsResponseBodyData {
	s.Uuid = &v
	return s
}

type ListTerminalsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTerminalsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalsResponse) GoString() string {
	return s.String()
}

func (s *ListTerminalsResponse) SetHeaders(v map[string]*string) *ListTerminalsResponse {
	s.Headers = v
	return s
}

func (s *ListTerminalsResponse) SetStatusCode(v int32) *ListTerminalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTerminalsResponse) SetBody(v *ListTerminalsResponseBody) *ListTerminalsResponse {
	s.Body = v
	return s
}

type ListTrustDevicesRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	UserCustomId *string `json:"UserCustomId,omitempty" xml:"UserCustomId,omitempty"`
}

func (s ListTrustDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrustDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListTrustDevicesRequest) SetLabelContent(v string) *ListTrustDevicesRequest {
	s.LabelContent = &v
	return s
}

func (s *ListTrustDevicesRequest) SetLabelId(v string) *ListTrustDevicesRequest {
	s.LabelId = &v
	return s
}

func (s *ListTrustDevicesRequest) SetSerialNo(v string) *ListTrustDevicesRequest {
	s.SerialNo = &v
	return s
}

func (s *ListTrustDevicesRequest) SetUserCustomId(v string) *ListTrustDevicesRequest {
	s.UserCustomId = &v
	return s
}

type ListTrustDevicesResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListTrustDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTrustDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrustDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrustDevicesResponseBody) SetCode(v string) *ListTrustDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTrustDevicesResponseBody) SetData(v []*ListTrustDevicesResponseBodyData) *ListTrustDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListTrustDevicesResponseBody) SetMessage(v string) *ListTrustDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTrustDevicesResponseBody) SetRequestId(v string) *ListTrustDevicesResponseBody {
	s.RequestId = &v
	return s
}

type ListTrustDevicesResponseBodyData struct {
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTrustDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTrustDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTrustDevicesResponseBodyData) SetModel(v string) *ListTrustDevicesResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListTrustDevicesResponseBodyData) SetSerialNo(v string) *ListTrustDevicesResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *ListTrustDevicesResponseBodyData) SetTenantId(v string) *ListTrustDevicesResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ListTrustDevicesResponseBodyData) SetUuid(v string) *ListTrustDevicesResponseBodyData {
	s.Uuid = &v
	return s
}

type ListTrustDevicesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrustDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrustDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrustDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListTrustDevicesResponse) SetHeaders(v map[string]*string) *ListTrustDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListTrustDevicesResponse) SetStatusCode(v int32) *ListTrustDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrustDevicesResponse) SetBody(v *ListTrustDevicesResponseBody) *ListTrustDevicesResponse {
	s.Body = v
	return s
}

type ListUserFbAcIssuesRequest struct {
	Account       *string `json:"Account,omitempty" xml:"Account,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IssueId       *string `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	Label         *string `json:"Label,omitempty" xml:"Label,omitempty"`
	ReservedA     *string `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB     *string `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	UserEmail     *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
}

func (s ListUserFbAcIssuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbAcIssuesRequest) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesRequest) SetAccount(v string) *ListUserFbAcIssuesRequest {
	s.Account = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetClientVersion(v string) *ListUserFbAcIssuesRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetErrorMessage(v string) *ListUserFbAcIssuesRequest {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetInstanceId(v string) *ListUserFbAcIssuesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetIssueId(v string) *ListUserFbAcIssuesRequest {
	s.IssueId = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetLabel(v string) *ListUserFbAcIssuesRequest {
	s.Label = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetReservedA(v string) *ListUserFbAcIssuesRequest {
	s.ReservedA = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetReservedB(v string) *ListUserFbAcIssuesRequest {
	s.ReservedB = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetUserEmail(v string) *ListUserFbAcIssuesRequest {
	s.UserEmail = &v
	return s
}

type ListUserFbAcIssuesResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListUserFbAcIssuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserFbAcIssuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbAcIssuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponseBody) SetCode(v string) *ListUserFbAcIssuesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBody) SetData(v *ListUserFbAcIssuesResponseBodyData) *ListUserFbAcIssuesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserFbAcIssuesResponseBody) SetMessage(v string) *ListUserFbAcIssuesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBody) SetRequestId(v string) *ListUserFbAcIssuesResponseBody {
	s.RequestId = &v
	return s
}

type ListUserFbAcIssuesResponseBodyData struct {
	IssueDataList []*ListUserFbAcIssuesResponseBodyDataIssueDataList `json:"IssueDataList,omitempty" xml:"IssueDataList,omitempty" type:"Repeated"`
	TotalCount    *int64                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserFbAcIssuesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbAcIssuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponseBodyData) SetIssueDataList(v []*ListUserFbAcIssuesResponseBodyDataIssueDataList) *ListUserFbAcIssuesResponseBodyData {
	s.IssueDataList = v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyData) SetTotalCount(v int64) *ListUserFbAcIssuesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListUserFbAcIssuesResponseBodyDataIssueDataList struct {
	Account       *string                                                    `json:"Account,omitempty" xml:"Account,omitempty"`
	ClientVersion *string                                                    `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ErrorMessage  *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FileList      []*ListUserFbAcIssuesResponseBodyDataIssueDataListFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	GmtCreated    *string                                                    `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified   *string                                                    `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceId    *string                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IssueId       *int64                                                     `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	Label         *string                                                    `json:"Label,omitempty" xml:"Label,omitempty"`
	ReservedA     *string                                                    `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB     *string                                                    `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	UserEmail     *string                                                    `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
}

func (s ListUserFbAcIssuesResponseBodyDataIssueDataList) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbAcIssuesResponseBodyDataIssueDataList) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetAccount(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.Account = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetClientVersion(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.ClientVersion = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetErrorMessage(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetFileList(v []*ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.FileList = v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetGmtCreated(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.GmtCreated = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetGmtModified(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.GmtModified = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetInstanceId(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.InstanceId = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetIssueId(v int64) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.IssueId = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetLabel(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.Label = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetReservedA(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.ReservedA = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetReservedB(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.ReservedB = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetUserEmail(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.UserEmail = &v
	return s
}

type ListUserFbAcIssuesResponseBodyDataIssueDataListFileList struct {
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize  *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) SetFileName(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	s.FileName = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) SetFileSize(v int32) *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	s.FileSize = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) SetFileType(v int32) *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	s.FileType = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) SetSessionId(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	s.SessionId = &v
	return s
}

type ListUserFbAcIssuesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserFbAcIssuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserFbAcIssuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbAcIssuesResponse) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponse) SetHeaders(v map[string]*string) *ListUserFbAcIssuesResponse {
	s.Headers = v
	return s
}

func (s *ListUserFbAcIssuesResponse) SetStatusCode(v int32) *ListUserFbAcIssuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserFbAcIssuesResponse) SetBody(v *ListUserFbAcIssuesResponseBody) *ListUserFbAcIssuesResponse {
	s.Body = v
	return s
}

type ListUserFbIssuesRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientId    *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientModel *string `json:"ClientModel,omitempty" xml:"ClientModel,omitempty"`
	ClientSn    *string `json:"ClientSn,omitempty" xml:"ClientSn,omitempty"`
	CustomerId  *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	ErrorCode   *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg    *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType      *int32  `json:"FbType,omitempty" xml:"FbType,omitempty"`
	IssueId     *int32  `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	IssueLabel  *string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail   *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WasRead     *int32  `json:"WasRead,omitempty" xml:"WasRead,omitempty"`
}

func (s ListUserFbIssuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbIssuesRequest) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesRequest) SetAppId(v string) *ListUserFbIssuesRequest {
	s.AppId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetClientId(v string) *ListUserFbIssuesRequest {
	s.ClientId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetClientModel(v string) *ListUserFbIssuesRequest {
	s.ClientModel = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetClientSn(v string) *ListUserFbIssuesRequest {
	s.ClientSn = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetCustomerId(v string) *ListUserFbIssuesRequest {
	s.CustomerId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetDescription(v string) *ListUserFbIssuesRequest {
	s.Description = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetDesktopId(v string) *ListUserFbIssuesRequest {
	s.DesktopId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetErrorCode(v string) *ListUserFbIssuesRequest {
	s.ErrorCode = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetErrorMsg(v string) *ListUserFbIssuesRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetFbType(v int32) *ListUserFbIssuesRequest {
	s.FbType = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetIssueId(v int32) *ListUserFbIssuesRequest {
	s.IssueId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetIssueLabel(v string) *ListUserFbIssuesRequest {
	s.IssueLabel = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetPageNumber(v int32) *ListUserFbIssuesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetPageSize(v int32) *ListUserFbIssuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetStatus(v int32) *ListUserFbIssuesRequest {
	s.Status = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetTitle(v string) *ListUserFbIssuesRequest {
	s.Title = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetUserEmail(v string) *ListUserFbIssuesRequest {
	s.UserEmail = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetUserId(v string) *ListUserFbIssuesRequest {
	s.UserId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetWasRead(v int32) *ListUserFbIssuesRequest {
	s.WasRead = &v
	return s
}

type ListUserFbIssuesResponseBody struct {
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *ListUserFbIssuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserFbIssuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbIssuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponseBody) SetCode(v string) *ListUserFbIssuesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetData(v *ListUserFbIssuesResponseBodyData) *ListUserFbIssuesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetMessage(v string) *ListUserFbIssuesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetPageNumber(v int32) *ListUserFbIssuesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetPageSize(v int32) *ListUserFbIssuesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetRequestId(v string) *ListUserFbIssuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetTotalCount(v int32) *ListUserFbIssuesResponseBody {
	s.TotalCount = &v
	return s
}

type ListUserFbIssuesResponseBodyData struct {
	Count             *string                                              `json:"Count,omitempty" xml:"Count,omitempty"`
	FeedbackIssueData []*ListUserFbIssuesResponseBodyDataFeedbackIssueData `json:"FeedbackIssueData,omitempty" xml:"FeedbackIssueData,omitempty" type:"Repeated"`
}

func (s ListUserFbIssuesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbIssuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponseBodyData) SetCount(v string) *ListUserFbIssuesResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyData) SetFeedbackIssueData(v []*ListUserFbIssuesResponseBodyDataFeedbackIssueData) *ListUserFbIssuesResponseBodyData {
	s.FeedbackIssueData = v
	return s
}

type ListUserFbIssuesResponseBodyDataFeedbackIssueData struct {
	AppId       *string                                                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientId    *string                                                      `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientModel *string                                                      `json:"ClientModel,omitempty" xml:"ClientModel,omitempty"`
	ClientSn    *string                                                      `json:"ClientSn,omitempty" xml:"ClientSn,omitempty"`
	CustomerId  *string                                                      `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Description *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId   *string                                                      `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	ErrorCode   *string                                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg    *string                                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType      *int32                                                       `json:"FbType,omitempty" xml:"FbType,omitempty"`
	FileList    []*ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	GmtCreated  *string                                                      `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	IssueId     *int32                                                       `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	IssueLabel  *string                                                      `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	Status      *int32                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Title       *string                                                      `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail   *string                                                      `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId      *string                                                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserFbIssuesResponseBodyDataFeedbackIssueData) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbIssuesResponseBodyDataFeedbackIssueData) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetAppId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.AppId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetClientId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ClientId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetClientModel(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ClientModel = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetClientSn(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ClientSn = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetCustomerId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.CustomerId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetDescription(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.Description = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetDesktopId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.DesktopId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetErrorCode(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ErrorCode = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetErrorMsg(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ErrorMsg = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetFbType(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.FbType = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetFileList(v []*ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.FileList = v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetGmtCreated(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.GmtCreated = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetIssueId(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.IssueId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetIssueLabel(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.IssueLabel = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetStatus(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.Status = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetTitle(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.Title = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetUserEmail(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.UserEmail = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetUserId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.UserId = &v
	return s
}

type ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList struct {
	FileMd5  *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	OssUrl   *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetFileMd5(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.FileMd5 = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetFileName(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.FileName = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetFileSize(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.FileSize = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetFileType(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.FileType = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetOssUrl(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.OssUrl = &v
	return s
}

type ListUserFbIssuesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserFbIssuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserFbIssuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserFbIssuesResponse) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponse) SetHeaders(v map[string]*string) *ListUserFbIssuesResponse {
	s.Headers = v
	return s
}

func (s *ListUserFbIssuesResponse) SetStatusCode(v int32) *ListUserFbIssuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserFbIssuesResponse) SetBody(v *ListUserFbIssuesResponseBody) *ListUserFbIssuesResponse {
	s.Body = v
	return s
}

type ModifyDevicesSecureNetworkTypeRequest struct {
	AllDevices *int64 `json:"AllDevices,omitempty" xml:"AllDevices,omitempty"`
	// This parameter is required.
	SecureNetworkType *string `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	SerialNos         *string `json:"SerialNos,omitempty" xml:"SerialNos,omitempty"`
}

func (s ModifyDevicesSecureNetworkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDevicesSecureNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDevicesSecureNetworkTypeRequest) SetAllDevices(v int64) *ModifyDevicesSecureNetworkTypeRequest {
	s.AllDevices = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeRequest) SetSecureNetworkType(v string) *ModifyDevicesSecureNetworkTypeRequest {
	s.SecureNetworkType = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeRequest) SetSerialNos(v string) *ModifyDevicesSecureNetworkTypeRequest {
	s.SerialNos = &v
	return s
}

type ModifyDevicesSecureNetworkTypeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDevicesSecureNetworkTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDevicesSecureNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) SetCode(v string) *ModifyDevicesSecureNetworkTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) SetMessage(v string) *ModifyDevicesSecureNetworkTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) SetRequestId(v string) *ModifyDevicesSecureNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDevicesSecureNetworkTypeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDevicesSecureNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDevicesSecureNetworkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDevicesSecureNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDevicesSecureNetworkTypeResponse) SetHeaders(v map[string]*string) *ModifyDevicesSecureNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponse) SetStatusCode(v int32) *ModifyDevicesSecureNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponse) SetBody(v *ModifyDevicesSecureNetworkTypeResponseBody) *ModifyDevicesSecureNetworkTypeResponse {
	s.Body = v
	return s
}

type ModifySecureNetworkTypeRequest struct {
	SecureNetworkType *string `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s ModifySecureNetworkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecureNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifySecureNetworkTypeRequest) SetSecureNetworkType(v string) *ModifySecureNetworkTypeRequest {
	s.SecureNetworkType = &v
	return s
}

func (s *ModifySecureNetworkTypeRequest) SetSerialNo(v string) *ModifySecureNetworkTypeRequest {
	s.SerialNo = &v
	return s
}

type ModifySecureNetworkTypeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecureNetworkTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecureNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecureNetworkTypeResponseBody) SetCode(v string) *ModifySecureNetworkTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySecureNetworkTypeResponseBody) SetMessage(v string) *ModifySecureNetworkTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySecureNetworkTypeResponseBody) SetRequestId(v string) *ModifySecureNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecureNetworkTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecureNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecureNetworkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecureNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifySecureNetworkTypeResponse) SetHeaders(v map[string]*string) *ModifySecureNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifySecureNetworkTypeResponse) SetStatusCode(v int32) *ModifySecureNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecureNetworkTypeResponse) SetBody(v *ModifySecureNetworkTypeResponseBody) *ModifySecureNetworkTypeResponse {
	s.Body = v
	return s
}

type RegisterDeviceRequest struct {
	Bluetooth  *string `json:"Bluetooth,omitempty" xml:"Bluetooth,omitempty"`
	BuildId    *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ChipId     *string `json:"ChipId,omitempty" xml:"ChipId,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Cpu        *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CustomId   *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	EtherMac   *string `json:"EtherMac,omitempty" xml:"EtherMac,omitempty"`
	Memory     *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Model      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	SerialNo   *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Storage    *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Wlan       *string `json:"Wlan,omitempty" xml:"Wlan,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) SetBluetooth(v string) *RegisterDeviceRequest {
	s.Bluetooth = &v
	return s
}

func (s *RegisterDeviceRequest) SetBuildId(v string) *RegisterDeviceRequest {
	s.BuildId = &v
	return s
}

func (s *RegisterDeviceRequest) SetChipId(v string) *RegisterDeviceRequest {
	s.ChipId = &v
	return s
}

func (s *RegisterDeviceRequest) SetClientId(v string) *RegisterDeviceRequest {
	s.ClientId = &v
	return s
}

func (s *RegisterDeviceRequest) SetClientType(v int32) *RegisterDeviceRequest {
	s.ClientType = &v
	return s
}

func (s *RegisterDeviceRequest) SetCpu(v string) *RegisterDeviceRequest {
	s.Cpu = &v
	return s
}

func (s *RegisterDeviceRequest) SetCustomId(v string) *RegisterDeviceRequest {
	s.CustomId = &v
	return s
}

func (s *RegisterDeviceRequest) SetEtherMac(v string) *RegisterDeviceRequest {
	s.EtherMac = &v
	return s
}

func (s *RegisterDeviceRequest) SetMemory(v string) *RegisterDeviceRequest {
	s.Memory = &v
	return s
}

func (s *RegisterDeviceRequest) SetModel(v string) *RegisterDeviceRequest {
	s.Model = &v
	return s
}

func (s *RegisterDeviceRequest) SetSerialNo(v string) *RegisterDeviceRequest {
	s.SerialNo = &v
	return s
}

func (s *RegisterDeviceRequest) SetStorage(v string) *RegisterDeviceRequest {
	s.Storage = &v
	return s
}

func (s *RegisterDeviceRequest) SetToken(v string) *RegisterDeviceRequest {
	s.Token = &v
	return s
}

func (s *RegisterDeviceRequest) SetWlan(v string) *RegisterDeviceRequest {
	s.Wlan = &v
	return s
}

type RegisterDeviceResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RegisterDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) SetCode(v string) *RegisterDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetData(v *RegisterDeviceResponseBodyData) *RegisterDeviceResponseBody {
	s.Data = v
	return s
}

func (s *RegisterDeviceResponseBody) SetMessage(v string) *RegisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetRequestId(v string) *RegisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

type RegisterDeviceResponseBodyData struct {
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RegisterDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBodyData) SetUuid(v string) *RegisterDeviceResponseBodyData {
	s.Uuid = &v
	return s
}

type RegisterDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) SetHeaders(v map[string]*string) *RegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceResponse) SetStatusCode(v int32) *RegisterDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
	s.Body = v
	return s
}

type ReportAppOtaInfoRequest struct {
	BaseVersion   *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	ClientType    *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientUid     *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OsType        *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Project       *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	TaskUid       *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s ReportAppOtaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportAppOtaInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportAppOtaInfoRequest) SetBaseVersion(v string) *ReportAppOtaInfoRequest {
	s.BaseVersion = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetClientType(v int32) *ReportAppOtaInfoRequest {
	s.ClientType = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetClientUid(v string) *ReportAppOtaInfoRequest {
	s.ClientUid = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetNote(v string) *ReportAppOtaInfoRequest {
	s.Note = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetOsType(v string) *ReportAppOtaInfoRequest {
	s.OsType = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetProject(v string) *ReportAppOtaInfoRequest {
	s.Project = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetStatus(v int32) *ReportAppOtaInfoRequest {
	s.Status = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetTargetVersion(v string) *ReportAppOtaInfoRequest {
	s.TargetVersion = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetTaskUid(v string) *ReportAppOtaInfoRequest {
	s.TaskUid = &v
	return s
}

type ReportAppOtaInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportAppOtaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportAppOtaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportAppOtaInfoResponseBody) SetCode(v string) *ReportAppOtaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ReportAppOtaInfoResponseBody) SetMessage(v string) *ReportAppOtaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ReportAppOtaInfoResponseBody) SetRequestId(v string) *ReportAppOtaInfoResponseBody {
	s.RequestId = &v
	return s
}

type ReportAppOtaInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportAppOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportAppOtaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportAppOtaInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportAppOtaInfoResponse) SetHeaders(v map[string]*string) *ReportAppOtaInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportAppOtaInfoResponse) SetStatusCode(v int32) *ReportAppOtaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportAppOtaInfoResponse) SetBody(v *ReportAppOtaInfoResponseBody) *ReportAppOtaInfoResponse {
	s.Body = v
	return s
}

type ReportDeviceOtaInfoRequest struct {
	BaseVersion   *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Model         *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s ReportDeviceOtaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceOtaInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportDeviceOtaInfoRequest) SetBaseVersion(v string) *ReportDeviceOtaInfoRequest {
	s.BaseVersion = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetDeviceId(v string) *ReportDeviceOtaInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetModel(v string) *ReportDeviceOtaInfoRequest {
	s.Model = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetNote(v string) *ReportDeviceOtaInfoRequest {
	s.Note = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetStatus(v int32) *ReportDeviceOtaInfoRequest {
	s.Status = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetTargetVersion(v string) *ReportDeviceOtaInfoRequest {
	s.TargetVersion = &v
	return s
}

type ReportDeviceOtaInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportDeviceOtaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceOtaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportDeviceOtaInfoResponseBody) SetCode(v string) *ReportDeviceOtaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ReportDeviceOtaInfoResponseBody) SetMessage(v string) *ReportDeviceOtaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ReportDeviceOtaInfoResponseBody) SetRequestId(v string) *ReportDeviceOtaInfoResponseBody {
	s.RequestId = &v
	return s
}

type ReportDeviceOtaInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportDeviceOtaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceOtaInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportDeviceOtaInfoResponse) SetHeaders(v map[string]*string) *ReportDeviceOtaInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportDeviceOtaInfoResponse) SetStatusCode(v int32) *ReportDeviceOtaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportDeviceOtaInfoResponse) SetBody(v *ReportDeviceOtaInfoResponseBody) *ReportDeviceOtaInfoResponse {
	s.Body = v
	return s
}

type ReportUserFbAcIssueRequest struct {
	Account       *string                               `json:"Account,omitempty" xml:"Account,omitempty"`
	ClientVersion *string                               `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ErrorMsg      *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FileList      []*ReportUserFbAcIssueRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	InstanceId    *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Labels        *string                               `json:"Labels,omitempty" xml:"Labels,omitempty"`
	ReservedA     *string                               `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB     *string                               `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	UserEmail     *string                               `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
}

func (s ReportUserFbAcIssueRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbAcIssueRequest) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueRequest) SetAccount(v string) *ReportUserFbAcIssueRequest {
	s.Account = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetClientVersion(v string) *ReportUserFbAcIssueRequest {
	s.ClientVersion = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetErrorMsg(v string) *ReportUserFbAcIssueRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetFileList(v []*ReportUserFbAcIssueRequestFileList) *ReportUserFbAcIssueRequest {
	s.FileList = v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetInstanceId(v string) *ReportUserFbAcIssueRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetLabels(v string) *ReportUserFbAcIssueRequest {
	s.Labels = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetReservedA(v string) *ReportUserFbAcIssueRequest {
	s.ReservedA = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetReservedB(v string) *ReportUserFbAcIssueRequest {
	s.ReservedB = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetUserEmail(v string) *ReportUserFbAcIssueRequest {
	s.UserEmail = &v
	return s
}

type ReportUserFbAcIssueRequestFileList struct {
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// This parameter is required.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ReportUserFbAcIssueRequestFileList) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbAcIssueRequestFileList) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueRequestFileList) SetFileName(v string) *ReportUserFbAcIssueRequestFileList {
	s.FileName = &v
	return s
}

func (s *ReportUserFbAcIssueRequestFileList) SetFileSize(v int32) *ReportUserFbAcIssueRequestFileList {
	s.FileSize = &v
	return s
}

func (s *ReportUserFbAcIssueRequestFileList) SetFileType(v int32) *ReportUserFbAcIssueRequestFileList {
	s.FileType = &v
	return s
}

func (s *ReportUserFbAcIssueRequestFileList) SetSessionId(v string) *ReportUserFbAcIssueRequestFileList {
	s.SessionId = &v
	return s
}

type ReportUserFbAcIssueShrinkRequest struct {
	Account        *string `json:"Account,omitempty" xml:"Account,omitempty"`
	ClientVersion  *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ErrorMsg       *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FileListShrink *string `json:"FileList,omitempty" xml:"FileList,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	ReservedA      *string `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB      *string `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	UserEmail      *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
}

func (s ReportUserFbAcIssueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbAcIssueShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueShrinkRequest) SetAccount(v string) *ReportUserFbAcIssueShrinkRequest {
	s.Account = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetClientVersion(v string) *ReportUserFbAcIssueShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetErrorMsg(v string) *ReportUserFbAcIssueShrinkRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetFileListShrink(v string) *ReportUserFbAcIssueShrinkRequest {
	s.FileListShrink = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetInstanceId(v string) *ReportUserFbAcIssueShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetLabels(v string) *ReportUserFbAcIssueShrinkRequest {
	s.Labels = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetReservedA(v string) *ReportUserFbAcIssueShrinkRequest {
	s.ReservedA = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetReservedB(v string) *ReportUserFbAcIssueShrinkRequest {
	s.ReservedB = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetUserEmail(v string) *ReportUserFbAcIssueShrinkRequest {
	s.UserEmail = &v
	return s
}

type ReportUserFbAcIssueResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReportUserFbAcIssueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportUserFbAcIssueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbAcIssueResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueResponseBody) SetCode(v string) *ReportUserFbAcIssueResponseBody {
	s.Code = &v
	return s
}

func (s *ReportUserFbAcIssueResponseBody) SetData(v *ReportUserFbAcIssueResponseBodyData) *ReportUserFbAcIssueResponseBody {
	s.Data = v
	return s
}

func (s *ReportUserFbAcIssueResponseBody) SetMessage(v string) *ReportUserFbAcIssueResponseBody {
	s.Message = &v
	return s
}

func (s *ReportUserFbAcIssueResponseBody) SetRequestId(v string) *ReportUserFbAcIssueResponseBody {
	s.RequestId = &v
	return s
}

type ReportUserFbAcIssueResponseBodyData struct {
	IssueId *int64 `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
}

func (s ReportUserFbAcIssueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbAcIssueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueResponseBodyData) SetIssueId(v int64) *ReportUserFbAcIssueResponseBodyData {
	s.IssueId = &v
	return s
}

type ReportUserFbAcIssueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportUserFbAcIssueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportUserFbAcIssueResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbAcIssueResponse) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueResponse) SetHeaders(v map[string]*string) *ReportUserFbAcIssueResponse {
	s.Headers = v
	return s
}

func (s *ReportUserFbAcIssueResponse) SetStatusCode(v int32) *ReportUserFbAcIssueResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportUserFbAcIssueResponse) SetBody(v *ReportUserFbAcIssueResponseBody) *ReportUserFbAcIssueResponse {
	s.Body = v
	return s
}

type ReportUserFbIssueRequest struct {
	AppId         *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientId      *string                             `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientModel   *string                             `json:"ClientModel,omitempty" xml:"ClientModel,omitempty"`
	ClientOsName  *string                             `json:"ClientOsName,omitempty" xml:"ClientOsName,omitempty"`
	ClientSn      *string                             `json:"ClientSn,omitempty" xml:"ClientSn,omitempty"`
	ClientVersion *string                             `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CustomerId    *string                             `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Description   *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId     *string                             `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopType   *int32                              `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	ErrorCode     *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg      *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType        *int32                              `json:"FbType,omitempty" xml:"FbType,omitempty"`
	FileList      []*ReportUserFbIssueRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	IssueLabel    *string                             `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	OccurTime     *int64                              `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	ReservedA     *string                             `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB     *string                             `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	TelNo         *string                             `json:"TelNo,omitempty" xml:"TelNo,omitempty"`
	Title         *string                             `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail     *string                             `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId        *string                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string                             `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WorkspaceId   *string                             `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WyId          *string                             `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ReportUserFbIssueRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbIssueRequest) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueRequest) SetAppId(v string) *ReportUserFbIssueRequest {
	s.AppId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientId(v string) *ReportUserFbIssueRequest {
	s.ClientId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientModel(v string) *ReportUserFbIssueRequest {
	s.ClientModel = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientOsName(v string) *ReportUserFbIssueRequest {
	s.ClientOsName = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientSn(v string) *ReportUserFbIssueRequest {
	s.ClientSn = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientVersion(v string) *ReportUserFbIssueRequest {
	s.ClientVersion = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetCustomerId(v string) *ReportUserFbIssueRequest {
	s.CustomerId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetDescription(v string) *ReportUserFbIssueRequest {
	s.Description = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetDesktopId(v string) *ReportUserFbIssueRequest {
	s.DesktopId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetDesktopType(v int32) *ReportUserFbIssueRequest {
	s.DesktopType = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetErrorCode(v string) *ReportUserFbIssueRequest {
	s.ErrorCode = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetErrorMsg(v string) *ReportUserFbIssueRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetFbType(v int32) *ReportUserFbIssueRequest {
	s.FbType = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetFileList(v []*ReportUserFbIssueRequestFileList) *ReportUserFbIssueRequest {
	s.FileList = v
	return s
}

func (s *ReportUserFbIssueRequest) SetIssueLabel(v string) *ReportUserFbIssueRequest {
	s.IssueLabel = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetOccurTime(v int64) *ReportUserFbIssueRequest {
	s.OccurTime = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetReservedA(v string) *ReportUserFbIssueRequest {
	s.ReservedA = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetReservedB(v string) *ReportUserFbIssueRequest {
	s.ReservedB = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetTelNo(v string) *ReportUserFbIssueRequest {
	s.TelNo = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetTitle(v string) *ReportUserFbIssueRequest {
	s.Title = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetUserEmail(v string) *ReportUserFbIssueRequest {
	s.UserEmail = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetUserId(v string) *ReportUserFbIssueRequest {
	s.UserId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetUserName(v string) *ReportUserFbIssueRequest {
	s.UserName = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetWorkspaceId(v string) *ReportUserFbIssueRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetWyId(v string) *ReportUserFbIssueRequest {
	s.WyId = &v
	return s
}

type ReportUserFbIssueRequestFileList struct {
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// This parameter is required.
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize  *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	OssUrl    *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ReportUserFbIssueRequestFileList) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbIssueRequestFileList) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueRequestFileList) SetFileMd5(v string) *ReportUserFbIssueRequestFileList {
	s.FileMd5 = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetFileName(v string) *ReportUserFbIssueRequestFileList {
	s.FileName = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetFileSize(v int32) *ReportUserFbIssueRequestFileList {
	s.FileSize = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetFileType(v int32) *ReportUserFbIssueRequestFileList {
	s.FileType = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetOssUrl(v string) *ReportUserFbIssueRequestFileList {
	s.OssUrl = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetSessionId(v string) *ReportUserFbIssueRequestFileList {
	s.SessionId = &v
	return s
}

type ReportUserFbIssueShrinkRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientId       *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientModel    *string `json:"ClientModel,omitempty" xml:"ClientModel,omitempty"`
	ClientOsName   *string `json:"ClientOsName,omitempty" xml:"ClientOsName,omitempty"`
	ClientSn       *string `json:"ClientSn,omitempty" xml:"ClientSn,omitempty"`
	ClientVersion  *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CustomerId     *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopType    *int32  `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg       *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType         *int32  `json:"FbType,omitempty" xml:"FbType,omitempty"`
	FileListShrink *string `json:"FileList,omitempty" xml:"FileList,omitempty"`
	IssueLabel     *string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	OccurTime      *int64  `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	ReservedA      *string `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB      *string `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	TelNo          *string `json:"TelNo,omitempty" xml:"TelNo,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail      *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WyId           *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ReportUserFbIssueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbIssueShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueShrinkRequest) SetAppId(v string) *ReportUserFbIssueShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientId(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientModel(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientModel = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientOsName(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientOsName = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientSn(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientSn = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientVersion(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetCustomerId(v string) *ReportUserFbIssueShrinkRequest {
	s.CustomerId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetDescription(v string) *ReportUserFbIssueShrinkRequest {
	s.Description = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetDesktopId(v string) *ReportUserFbIssueShrinkRequest {
	s.DesktopId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetDesktopType(v int32) *ReportUserFbIssueShrinkRequest {
	s.DesktopType = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetErrorCode(v string) *ReportUserFbIssueShrinkRequest {
	s.ErrorCode = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetErrorMsg(v string) *ReportUserFbIssueShrinkRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetFbType(v int32) *ReportUserFbIssueShrinkRequest {
	s.FbType = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetFileListShrink(v string) *ReportUserFbIssueShrinkRequest {
	s.FileListShrink = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetIssueLabel(v string) *ReportUserFbIssueShrinkRequest {
	s.IssueLabel = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetOccurTime(v int64) *ReportUserFbIssueShrinkRequest {
	s.OccurTime = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetReservedA(v string) *ReportUserFbIssueShrinkRequest {
	s.ReservedA = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetReservedB(v string) *ReportUserFbIssueShrinkRequest {
	s.ReservedB = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetTelNo(v string) *ReportUserFbIssueShrinkRequest {
	s.TelNo = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetTitle(v string) *ReportUserFbIssueShrinkRequest {
	s.Title = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetUserEmail(v string) *ReportUserFbIssueShrinkRequest {
	s.UserEmail = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetUserId(v string) *ReportUserFbIssueShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetUserName(v string) *ReportUserFbIssueShrinkRequest {
	s.UserName = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetWorkspaceId(v string) *ReportUserFbIssueShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetWyId(v string) *ReportUserFbIssueShrinkRequest {
	s.WyId = &v
	return s
}

type ReportUserFbIssueResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReportUserFbIssueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportUserFbIssueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbIssueResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueResponseBody) SetCode(v string) *ReportUserFbIssueResponseBody {
	s.Code = &v
	return s
}

func (s *ReportUserFbIssueResponseBody) SetData(v *ReportUserFbIssueResponseBodyData) *ReportUserFbIssueResponseBody {
	s.Data = v
	return s
}

func (s *ReportUserFbIssueResponseBody) SetMessage(v string) *ReportUserFbIssueResponseBody {
	s.Message = &v
	return s
}

func (s *ReportUserFbIssueResponseBody) SetRequestId(v string) *ReportUserFbIssueResponseBody {
	s.RequestId = &v
	return s
}

type ReportUserFbIssueResponseBodyData struct {
	IssueId *int32 `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
}

func (s ReportUserFbIssueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbIssueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueResponseBodyData) SetIssueId(v int32) *ReportUserFbIssueResponseBodyData {
	s.IssueId = &v
	return s
}

type ReportUserFbIssueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportUserFbIssueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportUserFbIssueResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportUserFbIssueResponse) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueResponse) SetHeaders(v map[string]*string) *ReportUserFbIssueResponse {
	s.Headers = v
	return s
}

func (s *ReportUserFbIssueResponse) SetStatusCode(v int32) *ReportUserFbIssueResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportUserFbIssueResponse) SetBody(v *ReportUserFbIssueResponseBody) *ReportUserFbIssueResponse {
	s.Body = v
	return s
}

type SendOpsMessageToTerminalsRequest struct {
	Delay      *bool     `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Msg        *string   `json:"Msg,omitempty" xml:"Msg,omitempty"`
	OpsAction  *string   `json:"OpsAction,omitempty" xml:"OpsAction,omitempty"`
	Uuids      []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
	WaitForAck *bool     `json:"WaitForAck,omitempty" xml:"WaitForAck,omitempty"`
}

func (s SendOpsMessageToTerminalsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendOpsMessageToTerminalsRequest) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalsRequest) SetDelay(v bool) *SendOpsMessageToTerminalsRequest {
	s.Delay = &v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) SetMsg(v string) *SendOpsMessageToTerminalsRequest {
	s.Msg = &v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) SetOpsAction(v string) *SendOpsMessageToTerminalsRequest {
	s.OpsAction = &v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) SetUuids(v []*string) *SendOpsMessageToTerminalsRequest {
	s.Uuids = v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) SetWaitForAck(v bool) *SendOpsMessageToTerminalsRequest {
	s.WaitForAck = &v
	return s
}

type SendOpsMessageToTerminalsResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendOpsMessageToTerminalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendOpsMessageToTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalsResponseBody) SetCode(v string) *SendOpsMessageToTerminalsResponseBody {
	s.Code = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetHttpStatusCode(v int32) *SendOpsMessageToTerminalsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetMessage(v string) *SendOpsMessageToTerminalsResponseBody {
	s.Message = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetRequestId(v string) *SendOpsMessageToTerminalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetSuccess(v bool) *SendOpsMessageToTerminalsResponseBody {
	s.Success = &v
	return s
}

type SendOpsMessageToTerminalsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendOpsMessageToTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendOpsMessageToTerminalsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendOpsMessageToTerminalsResponse) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalsResponse) SetHeaders(v map[string]*string) *SendOpsMessageToTerminalsResponse {
	s.Headers = v
	return s
}

func (s *SendOpsMessageToTerminalsResponse) SetStatusCode(v int32) *SendOpsMessageToTerminalsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponse) SetBody(v *SendOpsMessageToTerminalsResponseBody) *SendOpsMessageToTerminalsResponse {
	s.Body = v
	return s
}

type SetDeviceOtaAutoStatusRequest struct {
	AutoUpdate             *int32  `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	AutoUpdateTimeSchedule *string `json:"AutoUpdateTimeSchedule,omitempty" xml:"AutoUpdateTimeSchedule,omitempty"`
	ClientType             *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ForceUpgrade           *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDeviceOtaAutoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceOtaAutoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaAutoStatusRequest) SetAutoUpdate(v int32) *SetDeviceOtaAutoStatusRequest {
	s.AutoUpdate = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) SetAutoUpdateTimeSchedule(v string) *SetDeviceOtaAutoStatusRequest {
	s.AutoUpdateTimeSchedule = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) SetClientType(v int32) *SetDeviceOtaAutoStatusRequest {
	s.ClientType = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) SetForceUpgrade(v int32) *SetDeviceOtaAutoStatusRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) SetStatus(v string) *SetDeviceOtaAutoStatusRequest {
	s.Status = &v
	return s
}

type SetDeviceOtaAutoStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDeviceOtaAutoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceOtaAutoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaAutoStatusResponseBody) SetCode(v string) *SetDeviceOtaAutoStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeviceOtaAutoStatusResponseBody) SetMessage(v string) *SetDeviceOtaAutoStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SetDeviceOtaAutoStatusResponseBody) SetRequestId(v string) *SetDeviceOtaAutoStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetDeviceOtaAutoStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeviceOtaAutoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeviceOtaAutoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceOtaAutoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaAutoStatusResponse) SetHeaders(v map[string]*string) *SetDeviceOtaAutoStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDeviceOtaAutoStatusResponse) SetStatusCode(v int32) *SetDeviceOtaAutoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeviceOtaAutoStatusResponse) SetBody(v *SetDeviceOtaAutoStatusResponseBody) *SetDeviceOtaAutoStatusResponse {
	s.Body = v
	return s
}

type SetDeviceOtaTaskStatusRequest struct {
	// This parameter is required.
	OperationStatus *int32 `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// This parameter is required.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetDeviceOtaTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceOtaTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaTaskStatusRequest) SetOperationStatus(v int32) *SetDeviceOtaTaskStatusRequest {
	s.OperationStatus = &v
	return s
}

func (s *SetDeviceOtaTaskStatusRequest) SetTaskId(v int32) *SetDeviceOtaTaskStatusRequest {
	s.TaskId = &v
	return s
}

type SetDeviceOtaTaskStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDeviceOtaTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceOtaTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaTaskStatusResponseBody) SetCode(v string) *SetDeviceOtaTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeviceOtaTaskStatusResponseBody) SetMessage(v string) *SetDeviceOtaTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SetDeviceOtaTaskStatusResponseBody) SetRequestId(v string) *SetDeviceOtaTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetDeviceOtaTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeviceOtaTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeviceOtaTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceOtaTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaTaskStatusResponse) SetHeaders(v map[string]*string) *SetDeviceOtaTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDeviceOtaTaskStatusResponse) SetStatusCode(v int32) *SetDeviceOtaTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeviceOtaTaskStatusResponse) SetBody(v *SetDeviceOtaTaskStatusResponseBody) *SetDeviceOtaTaskStatusResponse {
	s.Body = v
	return s
}

type UnbindAccountLessLoginUserRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UnbindAccountLessLoginUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountLessLoginUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindAccountLessLoginUserRequest) SetSerialNumber(v string) *UnbindAccountLessLoginUserRequest {
	s.SerialNumber = &v
	return s
}

func (s *UnbindAccountLessLoginUserRequest) SetUuid(v string) *UnbindAccountLessLoginUserRequest {
	s.Uuid = &v
	return s
}

type UnbindAccountLessLoginUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindAccountLessLoginUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountLessLoginUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAccountLessLoginUserResponseBody) SetCode(v string) *UnbindAccountLessLoginUserResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) SetHttpStatusCode(v int32) *UnbindAccountLessLoginUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) SetMessage(v string) *UnbindAccountLessLoginUserResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) SetRequestId(v string) *UnbindAccountLessLoginUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) SetSuccess(v bool) *UnbindAccountLessLoginUserResponseBody {
	s.Success = &v
	return s
}

type UnbindAccountLessLoginUserResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAccountLessLoginUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAccountLessLoginUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountLessLoginUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindAccountLessLoginUserResponse) SetHeaders(v map[string]*string) *UnbindAccountLessLoginUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindAccountLessLoginUserResponse) SetStatusCode(v int32) *UnbindAccountLessLoginUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponse) SetBody(v *UnbindAccountLessLoginUserResponseBody) *UnbindAccountLessLoginUserResponse {
	s.Body = v
	return s
}

type UnbindDeviceSeatsRequest struct {
	SerialNoList []*string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty" type:"Repeated"`
}

func (s UnbindDeviceSeatsRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceSeatsRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceSeatsRequest) SetSerialNoList(v []*string) *UnbindDeviceSeatsRequest {
	s.SerialNoList = v
	return s
}

type UnbindDeviceSeatsShrinkRequest struct {
	SerialNoListShrink *string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty"`
}

func (s UnbindDeviceSeatsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceSeatsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceSeatsShrinkRequest) SetSerialNoListShrink(v string) *UnbindDeviceSeatsShrinkRequest {
	s.SerialNoListShrink = &v
	return s
}

type UnbindDeviceSeatsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDeviceSeatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDeviceSeatsResponseBody) SetCode(v string) *UnbindDeviceSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindDeviceSeatsResponseBody) SetMessage(v string) *UnbindDeviceSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindDeviceSeatsResponseBody) SetRequestId(v string) *UnbindDeviceSeatsResponseBody {
	s.RequestId = &v
	return s
}

type UnbindDeviceSeatsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDeviceSeatsResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceSeatsResponse) GoString() string {
	return s.String()
}

func (s *UnbindDeviceSeatsResponse) SetHeaders(v map[string]*string) *UnbindDeviceSeatsResponse {
	s.Headers = v
	return s
}

func (s *UnbindDeviceSeatsResponse) SetStatusCode(v int32) *UnbindDeviceSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDeviceSeatsResponse) SetBody(v *UnbindDeviceSeatsResponseBody) *UnbindDeviceSeatsResponse {
	s.Body = v
	return s
}

type UpdateAliasRequest struct {
	Alias    *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UpdateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) SetAlias(v string) *UpdateAliasRequest {
	s.Alias = &v
	return s
}

func (s *UpdateAliasRequest) SetSerialNo(v string) *UpdateAliasRequest {
	s.SerialNo = &v
	return s
}

func (s *UpdateAliasRequest) SetUuid(v string) *UpdateAliasRequest {
	s.Uuid = &v
	return s
}

type UpdateAliasResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponseBody) SetCode(v string) *UpdateAliasResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAliasResponseBody) SetMessage(v string) *UpdateAliasResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAliasResponseBody) SetRequestId(v string) *UpdateAliasResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponse) SetHeaders(v map[string]*string) *UpdateAliasResponse {
	s.Headers = v
	return s
}

func (s *UpdateAliasResponse) SetStatusCode(v int32) *UpdateAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAliasResponse) SetBody(v *UpdateAliasResponseBody) *UpdateAliasResponse {
	s.Body = v
	return s
}

type UpdateDeviceBindedEndUserRequest struct {
	// This parameter is required.
	SerialNo         *string                                             `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SourceAdEndUsers []*UpdateDeviceBindedEndUserRequestSourceAdEndUsers `json:"SourceAdEndUsers,omitempty" xml:"SourceAdEndUsers,omitempty" type:"Repeated"`
	SourceEndUserIds *string                                             `json:"SourceEndUserIds,omitempty" xml:"SourceEndUserIds,omitempty"`
	TargetAdEndUsers []*UpdateDeviceBindedEndUserRequestTargetAdEndUsers `json:"TargetAdEndUsers,omitempty" xml:"TargetAdEndUsers,omitempty" type:"Repeated"`
	TargetEndUserIds *string                                             `json:"TargetEndUserIds,omitempty" xml:"TargetEndUserIds,omitempty"`
	UserType         *string                                             `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s UpdateDeviceBindedEndUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceBindedEndUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserRequest) SetSerialNo(v string) *UpdateDeviceBindedEndUserRequest {
	s.SerialNo = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetSourceAdEndUsers(v []*UpdateDeviceBindedEndUserRequestSourceAdEndUsers) *UpdateDeviceBindedEndUserRequest {
	s.SourceAdEndUsers = v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetSourceEndUserIds(v string) *UpdateDeviceBindedEndUserRequest {
	s.SourceEndUserIds = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetTargetAdEndUsers(v []*UpdateDeviceBindedEndUserRequestTargetAdEndUsers) *UpdateDeviceBindedEndUserRequest {
	s.TargetAdEndUsers = v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetTargetEndUserIds(v string) *UpdateDeviceBindedEndUserRequest {
	s.TargetEndUserIds = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetUserType(v string) *UpdateDeviceBindedEndUserRequest {
	s.UserType = &v
	return s
}

type UpdateDeviceBindedEndUserRequestSourceAdEndUsers struct {
	AdDomain    *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId   *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
}

func (s UpdateDeviceBindedEndUserRequestSourceAdEndUsers) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceBindedEndUserRequestSourceAdEndUsers) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) SetAdDomain(v string) *UpdateDeviceBindedEndUserRequestSourceAdEndUsers {
	s.AdDomain = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) SetDirectoryId(v string) *UpdateDeviceBindedEndUserRequestSourceAdEndUsers {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) SetEndUserId(v string) *UpdateDeviceBindedEndUserRequestSourceAdEndUsers {
	s.EndUserId = &v
	return s
}

type UpdateDeviceBindedEndUserRequestTargetAdEndUsers struct {
	AdDomain    *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId   *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
}

func (s UpdateDeviceBindedEndUserRequestTargetAdEndUsers) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceBindedEndUserRequestTargetAdEndUsers) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) SetAdDomain(v string) *UpdateDeviceBindedEndUserRequestTargetAdEndUsers {
	s.AdDomain = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) SetDirectoryId(v string) *UpdateDeviceBindedEndUserRequestTargetAdEndUsers {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) SetEndUserId(v string) *UpdateDeviceBindedEndUserRequestTargetAdEndUsers {
	s.EndUserId = &v
	return s
}

type UpdateDeviceBindedEndUserResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeviceBindedEndUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceBindedEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserResponseBody) SetCode(v string) *UpdateDeviceBindedEndUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDeviceBindedEndUserResponseBody) SetMessage(v string) *UpdateDeviceBindedEndUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDeviceBindedEndUserResponseBody) SetRequestId(v string) *UpdateDeviceBindedEndUserResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDeviceBindedEndUserResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeviceBindedEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeviceBindedEndUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceBindedEndUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserResponse) SetHeaders(v map[string]*string) *UpdateDeviceBindedEndUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceBindedEndUserResponse) SetStatusCode(v int32) *UpdateDeviceBindedEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeviceBindedEndUserResponse) SetBody(v *UpdateDeviceBindedEndUserResponseBody) *UpdateDeviceBindedEndUserResponse {
	s.Body = v
	return s
}

type UpdateLabelRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
}

func (s UpdateLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelRequest) GoString() string {
	return s.String()
}

func (s *UpdateLabelRequest) SetLabelContent(v string) *UpdateLabelRequest {
	s.LabelContent = &v
	return s
}

func (s *UpdateLabelRequest) SetLabelId(v string) *UpdateLabelRequest {
	s.LabelId = &v
	return s
}

type UpdateLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLabelResponseBody) SetCode(v string) *UpdateLabelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLabelResponseBody) SetMessage(v string) *UpdateLabelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLabelResponseBody) SetRequestId(v string) *UpdateLabelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLabelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelResponse) GoString() string {
	return s.String()
}

func (s *UpdateLabelResponse) SetHeaders(v map[string]*string) *UpdateLabelResponse {
	s.Headers = v
	return s
}

func (s *UpdateLabelResponse) SetStatusCode(v int32) *UpdateLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLabelResponse) SetBody(v *UpdateLabelResponseBody) *UpdateLabelResponse {
	s.Body = v
	return s
}

type UpdateTerminalPolicyRequest struct {
	DisplayLayout           *string `json:"DisplayLayout,omitempty" xml:"DisplayLayout,omitempty"`
	DisplayResolution       *string `json:"DisplayResolution,omitempty" xml:"DisplayResolution,omitempty"`
	DisplayScaleRatio       *string `json:"DisplayScaleRatio,omitempty" xml:"DisplayScaleRatio,omitempty"`
	EnableAutoLockScreen    *int32  `json:"EnableAutoLockScreen,omitempty" xml:"EnableAutoLockScreen,omitempty"`
	EnableAutoLogin         *int32  `json:"EnableAutoLogin,omitempty" xml:"EnableAutoLogin,omitempty"`
	EnableBluetooth         *int32  `json:"EnableBluetooth,omitempty" xml:"EnableBluetooth,omitempty"`
	EnableModifyPassword    *int32  `json:"EnableModifyPassword,omitempty" xml:"EnableModifyPassword,omitempty"`
	EnableScheduledShutdown *int32  `json:"EnableScheduledShutdown,omitempty" xml:"EnableScheduledShutdown,omitempty"`
	EnableSwitchPersonal    *int32  `json:"EnableSwitchPersonal,omitempty" xml:"EnableSwitchPersonal,omitempty"`
	EnableWlan              *int32  `json:"EnableWlan,omitempty" xml:"EnableWlan,omitempty"`
	IdleTimeout             *int32  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	IdleTimeoutAction       *int32  `json:"IdleTimeoutAction,omitempty" xml:"IdleTimeoutAction,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PowerButtonDefine       *int32  `json:"PowerButtonDefine,omitempty" xml:"PowerButtonDefine,omitempty"`
	PowerButtonDefineForAs  *int32  `json:"PowerButtonDefineForAs,omitempty" xml:"PowerButtonDefineForAs,omitempty"`
	PowerButtonDefineForNs  *int32  `json:"PowerButtonDefineForNs,omitempty" xml:"PowerButtonDefineForNs,omitempty"`
	PowerOnBehavior         *int32  `json:"PowerOnBehavior,omitempty" xml:"PowerOnBehavior,omitempty"`
	ScheduledShutdown       *string `json:"ScheduledShutdown,omitempty" xml:"ScheduledShutdown,omitempty"`
	SettingLock             *int32  `json:"SettingLock,omitempty" xml:"SettingLock,omitempty"`
	TerminalPolicyId        *string `json:"TerminalPolicyId,omitempty" xml:"TerminalPolicyId,omitempty"`
}

func (s UpdateTerminalPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTerminalPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateTerminalPolicyRequest) SetDisplayLayout(v string) *UpdateTerminalPolicyRequest {
	s.DisplayLayout = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetDisplayResolution(v string) *UpdateTerminalPolicyRequest {
	s.DisplayResolution = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetDisplayScaleRatio(v string) *UpdateTerminalPolicyRequest {
	s.DisplayScaleRatio = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableAutoLockScreen(v int32) *UpdateTerminalPolicyRequest {
	s.EnableAutoLockScreen = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableAutoLogin(v int32) *UpdateTerminalPolicyRequest {
	s.EnableAutoLogin = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableBluetooth(v int32) *UpdateTerminalPolicyRequest {
	s.EnableBluetooth = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableModifyPassword(v int32) *UpdateTerminalPolicyRequest {
	s.EnableModifyPassword = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableScheduledShutdown(v int32) *UpdateTerminalPolicyRequest {
	s.EnableScheduledShutdown = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableSwitchPersonal(v int32) *UpdateTerminalPolicyRequest {
	s.EnableSwitchPersonal = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableWlan(v int32) *UpdateTerminalPolicyRequest {
	s.EnableWlan = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetIdleTimeout(v int32) *UpdateTerminalPolicyRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetIdleTimeoutAction(v int32) *UpdateTerminalPolicyRequest {
	s.IdleTimeoutAction = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetName(v string) *UpdateTerminalPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetPowerButtonDefine(v int32) *UpdateTerminalPolicyRequest {
	s.PowerButtonDefine = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetPowerButtonDefineForAs(v int32) *UpdateTerminalPolicyRequest {
	s.PowerButtonDefineForAs = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetPowerButtonDefineForNs(v int32) *UpdateTerminalPolicyRequest {
	s.PowerButtonDefineForNs = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetPowerOnBehavior(v int32) *UpdateTerminalPolicyRequest {
	s.PowerOnBehavior = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetScheduledShutdown(v string) *UpdateTerminalPolicyRequest {
	s.ScheduledShutdown = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetSettingLock(v int32) *UpdateTerminalPolicyRequest {
	s.SettingLock = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetTerminalPolicyId(v string) *UpdateTerminalPolicyRequest {
	s.TerminalPolicyId = &v
	return s
}

type UpdateTerminalPolicyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTerminalPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTerminalPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTerminalPolicyResponseBody) SetCode(v string) *UpdateTerminalPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) SetHttpStatusCode(v int32) *UpdateTerminalPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) SetMessage(v string) *UpdateTerminalPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) SetRequestId(v string) *UpdateTerminalPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) SetSuccess(v bool) *UpdateTerminalPolicyResponseBody {
	s.Success = &v
	return s
}

type UpdateTerminalPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTerminalPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTerminalPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTerminalPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateTerminalPolicyResponse) SetHeaders(v map[string]*string) *UpdateTerminalPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateTerminalPolicyResponse) SetStatusCode(v int32) *UpdateTerminalPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTerminalPolicyResponse) SetBody(v *UpdateTerminalPolicyResponseBody) *UpdateTerminalPolicyResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("wyota"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 设备激活
//
// @param request - ActivateDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateDeviceResponse
func (client *Client) ActivateDeviceWithOptions(request *ActivateDeviceRequest, runtime *util.RuntimeOptions) (_result *ActivateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateDevice"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备激活
//
// @param request - ActivateDeviceRequest
//
// @return ActivateDeviceResponse
func (client *Client) ActivateDevice(request *ActivateDeviceRequest) (_result *ActivateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateDeviceResponse{}
	_body, _err := client.ActivateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过序列号添加设备
//
// @param request - AddDeviceFromSNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDeviceFromSNResponse
func (client *Client) AddDeviceFromSNWithOptions(request *AddDeviceFromSNRequest, runtime *util.RuntimeOptions) (_result *AddDeviceFromSNResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.CustomProperty)) {
		body["CustomProperty"] = request.CustomProperty
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelContents)) {
		body["LabelContents"] = request.LabelContents
	}

	if !tea.BoolValue(util.IsUnset(request.SecureNetworkType)) {
		body["SecureNetworkType"] = request.SecureNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDeviceFromSN"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDeviceFromSNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过序列号添加设备
//
// @param request - AddDeviceFromSNRequest
//
// @return AddDeviceFromSNResponse
func (client *Client) AddDeviceFromSN(request *AddDeviceFromSNRequest) (_result *AddDeviceFromSNResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDeviceFromSNResponse{}
	_body, _err := client.AddDeviceFromSNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增设备座位和标签
//
// @param request - AddDeviceSeatsAndLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDeviceSeatsAndLabelsResponse
func (client *Client) AddDeviceSeatsAndLabelsWithOptions(request *AddDeviceSeatsAndLabelsRequest, runtime *util.RuntimeOptions) (_result *AddDeviceSeatsAndLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsUnique)) {
		body["IsUnique"] = request.IsUnique
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.LabelList)) {
		body["LabelList"] = request.LabelList
	}

	if !tea.BoolValue(util.IsUnset(request.SeatName)) {
		body["SeatName"] = request.SeatName
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDeviceSeatsAndLabels"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDeviceSeatsAndLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增设备座位和标签
//
// @param request - AddDeviceSeatsAndLabelsRequest
//
// @return AddDeviceSeatsAndLabelsResponse
func (client *Client) AddDeviceSeatsAndLabels(request *AddDeviceSeatsAndLabelsRequest) (_result *AddDeviceSeatsAndLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDeviceSeatsAndLabelsResponse{}
	_body, _err := client.AddDeviceSeatsAndLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过CSV文件添加设备
//
// @param request - AddDevicesFromCSVRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDevicesFromCSVResponse
func (client *Client) AddDevicesFromCSVWithOptions(request *AddDevicesFromCSVRequest, runtime *util.RuntimeOptions) (_result *AddDevicesFromCSVResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.SeatCol)) {
		body["SeatCol"] = request.SeatCol
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteName)) {
		body["SiteName"] = request.SiteName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDevicesFromCSV"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDevicesFromCSVResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过CSV文件添加设备
//
// @param request - AddDevicesFromCSVRequest
//
// @return AddDevicesFromCSVResponse
func (client *Client) AddDevicesFromCSV(request *AddDevicesFromCSVRequest) (_result *AddDevicesFromCSVResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDevicesFromCSVResponse{}
	_body, _err := client.AddDevicesFromCSVWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加标签
//
// @param request - AddLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLabelsResponse
func (client *Client) AddLabelsWithOptions(request *AddLabelsRequest, runtime *util.RuntimeOptions) (_result *AddLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelContents)) {
		body["LabelContents"] = request.LabelContents
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddLabels"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加标签
//
// @param request - AddLabelsRequest
//
// @return AddLabelsResponse
func (client *Client) AddLabels(request *AddLabelsRequest) (_result *AddLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLabelsResponse{}
	_body, _err := client.AddLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或更新设备工位
//
// @param request - AddOrUpdateDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateDeviceSeatsResponse
func (client *Client) AddOrUpdateDeviceSeatsWithOptions(request *AddOrUpdateDeviceSeatsRequest, runtime *util.RuntimeOptions) (_result *AddOrUpdateDeviceSeatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.UserCustomId)) {
		body["UserCustomId"] = request.UserCustomId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddOrUpdateDeviceSeats"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOrUpdateDeviceSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或更新设备工位
//
// @param request - AddOrUpdateDeviceSeatsRequest
//
// @return AddOrUpdateDeviceSeatsResponse
func (client *Client) AddOrUpdateDeviceSeats(request *AddOrUpdateDeviceSeatsRequest) (_result *AddOrUpdateDeviceSeatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddOrUpdateDeviceSeatsResponse{}
	_body, _err := client.AddOrUpdateDeviceSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTerminalResponse
func (client *Client) AddTerminalWithOptions(request *AddTerminalRequest, runtime *util.RuntimeOptions) (_result *AddTerminalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalGroupId)) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTerminal"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalRequest
//
// @return AddTerminalResponse
func (client *Client) AddTerminal(request *AddTerminalRequest) (_result *AddTerminalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTerminalResponse{}
	_body, _err := client.AddTerminalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备绑定终端用户
//
// @param request - AttachEndUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachEndUsersResponse
func (client *Client) AttachEndUsersWithOptions(request *AttachEndUsersRequest, runtime *util.RuntimeOptions) (_result *AttachEndUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachEndUsers"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachEndUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备绑定终端用户
//
// @param request - AttachEndUsersRequest
//
// @return AttachEndUsersResponse
func (client *Client) AttachEndUsers(request *AttachEndUsersRequest) (_result *AttachEndUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachEndUsersResponse{}
	_body, _err := client.AttachEndUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备绑定标签
//
// @param request - AttachLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachLabelResponse
func (client *Client) AttachLabelWithOptions(request *AttachLabelRequest, runtime *util.RuntimeOptions) (_result *AttachLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelContent)) {
		body["LabelContent"] = request.LabelContent
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachLabel"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备绑定标签
//
// @param request - AttachLabelRequest
//
// @return AttachLabelResponse
func (client *Client) AttachLabel(request *AttachLabelRequest) (_result *AttachLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachLabelResponse{}
	_body, _err := client.AttachLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量绑定标签
//
// @param request - AttachLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachLabelsResponse
func (client *Client) AttachLabelsWithOptions(request *AttachLabelsRequest, runtime *util.RuntimeOptions) (_result *AttachLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelIds)) {
		body["LabelIds"] = request.LabelIds
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNoList)) {
		body["SerialNoList"] = request.SerialNoList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachLabels"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量绑定标签
//
// @param request - AttachLabelsRequest
//
// @return AttachLabelsResponse
func (client *Client) AttachLabels(request *AttachLabelsRequest) (_result *AttachLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachLabelsResponse{}
	_body, _err := client.AttachLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindAccountLessLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAccountLessLoginUserResponse
func (client *Client) BindAccountLessLoginUserWithOptions(request *BindAccountLessLoginUserRequest, runtime *util.RuntimeOptions) (_result *BindAccountLessLoginUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAccountLessLoginUser"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAccountLessLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindAccountLessLoginUserRequest
//
// @return BindAccountLessLoginUserResponse
func (client *Client) BindAccountLessLoginUser(request *BindAccountLessLoginUserRequest) (_result *BindAccountLessLoginUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAccountLessLoginUserResponse{}
	_body, _err := client.BindAccountLessLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查uuid有效性
//
// @param request - CheckUuidValidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUuidValidResponse
func (client *Client) CheckUuidValidWithOptions(request *CheckUuidValidRequest, runtime *util.RuntimeOptions) (_result *CheckUuidValidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bluetooth)) {
		body["Bluetooth"] = request.Bluetooth
	}

	if !tea.BoolValue(util.IsUnset(request.BuildId)) {
		body["BuildId"] = request.BuildId
	}

	if !tea.BoolValue(util.IsUnset(request.ChipId)) {
		body["ChipId"] = request.ChipId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomId)) {
		body["CustomId"] = request.CustomId
	}

	if !tea.BoolValue(util.IsUnset(request.EtherMac)) {
		body["EtherMac"] = request.EtherMac
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.Wlan)) {
		body["Wlan"] = request.Wlan
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUuidValid"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUuidValidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查uuid有效性
//
// @param request - CheckUuidValidRequest
//
// @return CheckUuidValidResponse
func (client *Client) CheckUuidValid(request *CheckUuidValidRequest) (_result *CheckUuidValidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUuidValidResponse{}
	_body, _err := client.CheckUuidValidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateAppOtaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppOtaTaskResponse
func (client *Client) CreateAppOtaTaskWithOptions(request *CreateAppOtaTaskRequest, runtime *util.RuntimeOptions) (_result *CreateAppOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionUid)) {
		query["AppVersionUid"] = request.AppVersionUid
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIdList)) {
		query["ClientIdList"] = request.ClientIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ForceUpgrade)) {
		query["ForceUpgrade"] = request.ForceUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.Regions)) {
		query["Regions"] = request.Regions
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppOtaTask"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppOtaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateAppOtaTaskRequest
//
// @return CreateAppOtaTaskResponse
func (client *Client) CreateAppOtaTask(request *CreateAppOtaTaskRequest) (_result *CreateAppOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppOtaTaskResponse{}
	_body, _err := client.CreateAppOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建版本
//
// @param request - CreateAppOtaVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppOtaVersionResponse
func (client *Client) CreateAppOtaVersionWithOptions(request *CreateAppOtaVersionRequest, runtime *util.RuntimeOptions) (_result *CreateAppOtaVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Arch)) {
		query["Arch"] = request.Arch
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadUrl)) {
		query["DownloadUrl"] = request.DownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Md5)) {
		query["Md5"] = request.Md5
	}

	if !tea.BoolValue(util.IsUnset(request.Os)) {
		query["Os"] = request.Os
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OtaType)) {
		query["OtaType"] = request.OtaType
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNote)) {
		query["ReleaseNote"] = request.ReleaseNote
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNoteEn)) {
		query["ReleaseNoteEn"] = request.ReleaseNoteEn
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNoteJp)) {
		query["ReleaseNoteJp"] = request.ReleaseNoteJp
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotRegionId)) {
		query["SnapshotRegionId"] = request.SnapshotRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppOtaVersion"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppOtaVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建版本
//
// @param request - CreateAppOtaVersionRequest
//
// @return CreateAppOtaVersionResponse
func (client *Client) CreateAppOtaVersion(request *CreateAppOtaVersionRequest) (_result *CreateAppOtaVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppOtaVersionResponse{}
	_body, _err := client.CreateAppOtaVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除版本
//
// @param request - DeleteAppOtaVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppOtaVersionsResponse
func (client *Client) DeleteAppOtaVersionsWithOptions(request *DeleteAppOtaVersionsRequest, runtime *util.RuntimeOptions) (_result *DeleteAppOtaVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionUidList)) {
		query["VersionUidList"] = request.VersionUidList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppOtaVersions"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppOtaVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除版本
//
// @param request - DeleteAppOtaVersionsRequest
//
// @return DeleteAppOtaVersionsResponse
func (client *Client) DeleteAppOtaVersions(request *DeleteAppOtaVersionsRequest) (_result *DeleteAppOtaVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppOtaVersionsResponse{}
	_body, _err := client.DeleteAppOtaVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除设备
//
// @param request - DeleteDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDevicesResponse
func (client *Client) DeleteDevicesWithOptions(request *DeleteDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uuids)) {
		query["Uuids"] = request.Uuids
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		body["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNos)) {
		body["SerialNos"] = request.SerialNos
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevices"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除设备
//
// @param request - DeleteDevicesRequest
//
// @return DeleteDevicesResponse
func (client *Client) DeleteDevices(request *DeleteDevicesRequest) (_result *DeleteDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDevicesResponse{}
	_body, _err := client.DeleteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - DeleteLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLabelResponse
func (client *Client) DeleteLabelWithOptions(request *DeleteLabelRequest, runtime *util.RuntimeOptions) (_result *DeleteLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		body["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.LabelContent)) {
		body["LabelContent"] = request.LabelContent
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLabel"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - DeleteLabelRequest
//
// @return DeleteLabelResponse
func (client *Client) DeleteLabel(request *DeleteLabelRequest) (_result *DeleteLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLabelResponse{}
	_body, _err := client.DeleteLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询版本
//
// @param request - DescribeAppOtaVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppOtaVersionResponse
func (client *Client) DescribeAppOtaVersionWithOptions(request *DescribeAppOtaVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeAppOtaVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VersionUid)) {
		query["VersionUid"] = request.VersionUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppOtaVersion"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppOtaVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询版本
//
// @param request - DescribeAppOtaVersionRequest
//
// @return DescribeAppOtaVersionResponse
func (client *Client) DescribeAppOtaVersion(request *DescribeAppOtaVersionRequest) (_result *DescribeAppOtaVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppOtaVersionResponse{}
	_body, _err := client.DescribeAppOtaVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备座位
//
// @param request - DescribeDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceSeatsResponse
func (client *Client) DescribeDeviceSeatsWithOptions(request *DescribeDeviceSeatsRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceSeatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNoList)) {
		body["SerialNoList"] = request.SerialNoList
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceSeats"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备座位
//
// @param request - DescribeDeviceSeatsRequest
//
// @return DescribeDeviceSeatsResponse
func (client *Client) DescribeDeviceSeats(request *DescribeDeviceSeatsRequest) (_result *DescribeDeviceSeatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceSeatsResponse{}
	_body, _err := client.DescribeDeviceSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询版本信息
//
// @param request - DescribeDeviceVersionDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceVersionDetailResponse
func (client *Client) DescribeDeviceVersionDetailWithOptions(request *DescribeDeviceVersionDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceVersionDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		body["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceVersionDetail"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceVersionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询版本信息
//
// @param request - DescribeDeviceVersionDetailRequest
//
// @return DescribeDeviceVersionDetailResponse
func (client *Client) DescribeDeviceVersionDetail(request *DescribeDeviceVersionDetailRequest) (_result *DescribeDeviceVersionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceVersionDetailResponse{}
	_body, _err := client.DescribeDeviceVersionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备标签数量
//
// @param request - DescribeSnLabelCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnLabelCountsResponse
func (client *Client) DescribeSnLabelCountsWithOptions(request *DescribeSnLabelCountsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnLabelCountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelList)) {
		body["LabelList"] = request.LabelList
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneName)) {
		body["ZoneName"] = request.ZoneName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSnLabelCounts"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSnLabelCountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备标签数量
//
// @param request - DescribeSnLabelCountsRequest
//
// @return DescribeSnLabelCountsResponse
func (client *Client) DescribeSnLabelCounts(request *DescribeSnLabelCountsRequest) (_result *DescribeSnLabelCountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnLabelCountsResponse{}
	_body, _err := client.DescribeSnLabelCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工作区域
//
// @param request - DescribeWorkZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkZonesResponse
func (client *Client) DescribeWorkZonesWithOptions(request *DescribeWorkZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeWorkZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneIdList)) {
		body["ZoneIdList"] = request.ZoneIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneNameList)) {
		body["ZoneNameList"] = request.ZoneNameList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWorkZones"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWorkZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作区域
//
// @param request - DescribeWorkZonesRequest
//
// @return DescribeWorkZonesResponse
func (client *Client) DescribeWorkZones(request *DescribeWorkZonesRequest) (_result *DescribeWorkZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWorkZonesResponse{}
	_body, _err := client.DescribeWorkZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备解绑终端用户
//
// @param request - DetachEndUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachEndUsersResponse
func (client *Client) DetachEndUsersWithOptions(request *DetachEndUsersRequest, runtime *util.RuntimeOptions) (_result *DetachEndUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachEndUsers"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachEndUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备解绑终端用户
//
// @param request - DetachEndUsersRequest
//
// @return DetachEndUsersResponse
func (client *Client) DetachEndUsers(request *DetachEndUsersRequest) (_result *DetachEndUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachEndUsersResponse{}
	_body, _err := client.DetachEndUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备绑定标签
//
// @param request - DetachLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachLabelResponse
func (client *Client) DetachLabelWithOptions(request *DetachLabelRequest, runtime *util.RuntimeOptions) (_result *DetachLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelContent)) {
		body["LabelContent"] = request.LabelContent
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachLabel"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备绑定标签
//
// @param request - DetachLabelRequest
//
// @return DetachLabelResponse
func (client *Client) DetachLabel(request *DetachLabelRequest) (_result *DetachLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachLabelResponse{}
	_body, _err := client.DetachLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量解绑标签
//
// @param request - DetachLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachLabelsResponse
func (client *Client) DetachLabelsWithOptions(request *DetachLabelsRequest, runtime *util.RuntimeOptions) (_result *DetachLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelIds)) {
		body["LabelIds"] = request.LabelIds
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNoList)) {
		body["SerialNoList"] = request.SerialNoList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachLabels"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量解绑标签
//
// @param request - DetachLabelsRequest
//
// @return DetachLabelsResponse
func (client *Client) DetachLabels(request *DetachLabelsRequest) (_result *DetachLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachLabelsResponse{}
	_body, _err := client.DetachLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户上传的文件
//
// @param request - GenerateOssUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateOssUrlResponse
func (client *Client) GenerateOssUrlWithOptions(request *GenerateOssUrlRequest, runtime *util.RuntimeOptions) (_result *GenerateOssUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectNameList)) {
		body["ObjectNameList"] = request.ObjectNameList
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateOssUrl"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateOssUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户上传的文件
//
// @param request - GenerateOssUrlRequest
//
// @return GenerateOssUrlResponse
func (client *Client) GenerateOssUrl(request *GenerateOssUrlRequest) (_result *GenerateOssUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateOssUrlResponse{}
	_body, _err := client.GenerateOssUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用最新版本
//
// @param request - GetAppOtaLatestVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppOtaLatestVersionResponse
func (client *Client) GetAppOtaLatestVersionWithOptions(request *GetAppOtaLatestVersionRequest, runtime *util.RuntimeOptions) (_result *GetAppOtaLatestVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseVersion)) {
		query["BaseVersion"] = request.BaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientUid)) {
		query["ClientUid"] = request.ClientUid
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppOtaLatestVersion"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppOtaLatestVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用最新版本
//
// @param request - GetAppOtaLatestVersionRequest
//
// @return GetAppOtaLatestVersionResponse
func (client *Client) GetAppOtaLatestVersion(request *GetAppOtaLatestVersionRequest) (_result *GetAppOtaLatestVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppOtaLatestVersionResponse{}
	_body, _err := client.GetAppOtaLatestVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备配置
//
// @param request - GetDeviceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceConfigsResponse
func (client *Client) GetDeviceConfigsWithOptions(request *GetDeviceConfigsRequest, runtime *util.RuntimeOptions) (_result *GetDeviceConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.UrclVersion)) {
		body["UrclVersion"] = request.UrclVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UserCustomId)) {
		body["UserCustomId"] = request.UserCustomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceConfigs"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备配置
//
// @param request - GetDeviceConfigsRequest
//
// @return GetDeviceConfigsResponse
func (client *Client) GetDeviceConfigs(request *GetDeviceConfigsRequest) (_result *GetDeviceConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceConfigsResponse{}
	_body, _err := client.GetDeviceConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取是否开启自动升级状态
//
// @param request - GetDeviceOtaAutoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOtaAutoStatusResponse
func (client *Client) GetDeviceOtaAutoStatusWithOptions(request *GetDeviceOtaAutoStatusRequest, runtime *util.RuntimeOptions) (_result *GetDeviceOtaAutoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		body["ClientType"] = request.ClientType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceOtaAutoStatus"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceOtaAutoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取是否开启自动升级状态
//
// @param request - GetDeviceOtaAutoStatusRequest
//
// @return GetDeviceOtaAutoStatusResponse
func (client *Client) GetDeviceOtaAutoStatus(request *GetDeviceOtaAutoStatusRequest) (_result *GetDeviceOtaAutoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceOtaAutoStatusResponse{}
	_body, _err := client.GetDeviceOtaAutoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备升级信息
//
// @param request - GetDeviceOtaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOtaInfoResponse
func (client *Client) GetDeviceOtaInfoWithOptions(request *GetDeviceOtaInfoRequest, runtime *util.RuntimeOptions) (_result *GetDeviceOtaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseVersion)) {
		body["BaseVersion"] = request.BaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.OsVersion)) {
		body["OsVersion"] = request.OsVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersionType)) {
		body["TargetVersionType"] = request.TargetVersionType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceOtaInfo"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceOtaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备升级信息
//
// @param request - GetDeviceOtaInfoRequest
//
// @return GetDeviceOtaInfoResponse
func (client *Client) GetDeviceOtaInfo(request *GetDeviceOtaInfoRequest) (_result *GetDeviceOtaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceOtaInfoResponse{}
	_body, _err := client.GetDeviceOtaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备升级信息
//
// @param request - GetDeviceOtaInfoTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOtaInfoTestResponse
func (client *Client) GetDeviceOtaInfoTestWithOptions(request *GetDeviceOtaInfoTestRequest, runtime *util.RuntimeOptions) (_result *GetDeviceOtaInfoTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseVersion)) {
		body["BaseVersion"] = request.BaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceOtaInfoTest"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceOtaInfoTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备升级信息
//
// @param request - GetDeviceOtaInfoTestRequest
//
// @return GetDeviceOtaInfoTestResponse
func (client *Client) GetDeviceOtaInfoTest(request *GetDeviceOtaInfoTestRequest) (_result *GetDeviceOtaInfoTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceOtaInfoTestResponse{}
	_body, _err := client.GetDeviceOtaInfoTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户任务版本信息
//
// @param request - GetDeviceOtaTaskVersionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOtaTaskVersionInfoResponse
func (client *Client) GetDeviceOtaTaskVersionInfoWithOptions(request *GetDeviceOtaTaskVersionInfoRequest, runtime *util.RuntimeOptions) (_result *GetDeviceOtaTaskVersionInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceOtaTaskVersionInfo"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceOtaTaskVersionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户任务版本信息
//
// @param request - GetDeviceOtaTaskVersionInfoRequest
//
// @return GetDeviceOtaTaskVersionInfoResponse
func (client *Client) GetDeviceOtaTaskVersionInfo(request *GetDeviceOtaTaskVersionInfoRequest) (_result *GetDeviceOtaTaskVersionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceOtaTaskVersionInfoResponse{}
	_body, _err := client.GetDeviceOtaTaskVersionInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得设备升级详情
//
// @param request - GetDeviceUpgradeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceUpgradeStatusResponse
func (client *Client) GetDeviceUpgradeStatusWithOptions(request *GetDeviceUpgradeStatusRequest, runtime *util.RuntimeOptions) (_result *GetDeviceUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientUid)) {
		query["ClientUid"] = request.ClientUid
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUid)) {
		query["TaskUid"] = request.TaskUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceUpgradeStatus"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceUpgradeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得设备升级详情
//
// @param request - GetDeviceUpgradeStatusRequest
//
// @return GetDeviceUpgradeStatusResponse
func (client *Client) GetDeviceUpgradeStatus(request *GetDeviceUpgradeStatusRequest) (_result *GetDeviceUpgradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceUpgradeStatusResponse{}
	_body, _err := client.GetDeviceUpgradeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出设备工位列表
//
// @param request - GetExportDeviceInfoOssUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExportDeviceInfoOssUrlResponse
func (client *Client) GetExportDeviceInfoOssUrlWithOptions(request *GetExportDeviceInfoOssUrlRequest, runtime *util.RuntimeOptions) (_result *GetExportDeviceInfoOssUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneName)) {
		body["ZoneName"] = request.ZoneName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExportDeviceInfoOssUrl"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExportDeviceInfoOssUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出设备工位列表
//
// @param request - GetExportDeviceInfoOssUrlRequest
//
// @return GetExportDeviceInfoOssUrlResponse
func (client *Client) GetExportDeviceInfoOssUrl(request *GetExportDeviceInfoOssUrlRequest) (_result *GetExportDeviceInfoOssUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExportDeviceInfoOssUrlResponse{}
	_body, _err := client.GetExportDeviceInfoOssUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询OSS配置信息
//
// @param request - GetFbOssConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFbOssConfigResponse
func (client *Client) GetFbOssConfigWithOptions(request *GetFbOssConfigRequest, runtime *util.RuntimeOptions) (_result *GetFbOssConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirPrefix)) {
		body["DirPrefix"] = request.DirPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.IsDedicatedLine)) {
		body["IsDedicatedLine"] = request.IsDedicatedLine
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFbOssConfig"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFbOssConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询OSS配置信息
//
// @param request - GetFbOssConfigRequest
//
// @return GetFbOssConfigResponse
func (client *Client) GetFbOssConfig(request *GetFbOssConfigRequest) (_result *GetFbOssConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFbOssConfigResponse{}
	_body, _err := client.GetFbOssConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取OSS配置
//
// @param request - GetOssConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssConfigResponse
func (client *Client) GetOssConfigWithOptions(request *GetOssConfigRequest, runtime *util.RuntimeOptions) (_result *GetOssConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOssConfig"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取OSS配置
//
// @param request - GetOssConfigRequest
//
// @return GetOssConfigResponse
func (client *Client) GetOssConfig(request *GetOssConfigRequest) (_result *GetOssConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssConfigResponse{}
	_body, _err := client.GetOssConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取版本下载地址
//
// @param request - GetVersionDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVersionDownloadUrlResponse
func (client *Client) GetVersionDownloadUrlWithOptions(request *GetVersionDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetVersionDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVersionDownloadUrl"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVersionDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取版本下载地址
//
// @param request - GetVersionDownloadUrlRequest
//
// @return GetVersionDownloadUrlResponse
func (client *Client) GetVersionDownloadUrl(request *GetVersionDownloadUrlRequest) (_result *GetVersionDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVersionDownloadUrlResponse{}
	_body, _err := client.GetVersionDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户ota任务
//
// @param request - ListDeviceOtaTaskByTenantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceOtaTaskByTenantResponse
func (client *Client) ListDeviceOtaTaskByTenantWithOptions(request *ListDeviceOtaTaskByTenantRequest, runtime *util.RuntimeOptions) (_result *ListDeviceOtaTaskByTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceOtaTaskByTenant"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceOtaTaskByTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户ota任务
//
// @param request - ListDeviceOtaTaskByTenantRequest
//
// @return ListDeviceOtaTaskByTenantResponse
func (client *Client) ListDeviceOtaTaskByTenant(request *ListDeviceOtaTaskByTenantRequest) (_result *ListDeviceOtaTaskByTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceOtaTaskByTenantResponse{}
	_body, _err := client.ListDeviceOtaTaskByTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备座位列表
//
// @param request - ListDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceSeatsResponse
func (client *Client) ListDeviceSeatsWithOptions(request *ListDeviceSeatsRequest, runtime *util.RuntimeOptions) (_result *ListDeviceSeatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.SeatNo)) {
		body["SeatNo"] = request.SeatNo
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNoList)) {
		body["SerialNoList"] = request.SerialNoList
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceSeats"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备座位列表
//
// @param request - ListDeviceSeatsRequest
//
// @return ListDeviceSeatsResponse
func (client *Client) ListDeviceSeats(request *ListDeviceSeatsRequest) (_result *ListDeviceSeatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceSeatsResponse{}
	_body, _err := client.ListDeviceSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备列表
//
// @param request - ListDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDevicesResponse
func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIpV4)) {
		query["DeviceIpV4"] = request.DeviceIpV4
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceOS)) {
		query["DeviceOS"] = request.DeviceOS
	}

	if !tea.BoolValue(util.IsUnset(request.DevicePlatform)) {
		query["DevicePlatform"] = request.DevicePlatform
	}

	if !tea.BoolValue(util.IsUnset(request.LocationInfo)) {
		query["LocationInfo"] = request.LocationInfo
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.BuildId)) {
		body["BuildId"] = request.BuildId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		body["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelContent)) {
		body["LabelContent"] = request.LabelContent
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevices"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备列表
//
// @param request - ListDevicesRequest
//
// @return ListDevicesResponse
func (client *Client) ListDevices(request *ListDevicesRequest) (_result *ListDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevicesResponse{}
	_body, _err := client.ListDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户问题标签
//
// @param request - ListFbIssueLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFbIssueLabelsResponse
func (client *Client) ListFbIssueLabelsWithOptions(runtime *util.RuntimeOptions) (_result *ListFbIssueLabelsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListFbIssueLabels"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFbIssueLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户问题标签
//
// @return ListFbIssueLabelsResponse
func (client *Client) ListFbIssueLabels() (_result *ListFbIssueLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFbIssueLabelsResponse{}
	_body, _err := client.ListFbIssueLabelsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据语言类型和调用方查询标签列表
//
// @param request - ListFbIssueLabelsByLCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFbIssueLabelsByLCResponse
func (client *Client) ListFbIssueLabelsByLCWithOptions(request *ListFbIssueLabelsByLCRequest, runtime *util.RuntimeOptions) (_result *ListFbIssueLabelsByLCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		body["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageType)) {
		body["LanguageType"] = request.LanguageType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFbIssueLabelsByLC"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFbIssueLabelsByLCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据语言类型和调用方查询标签列表
//
// @param request - ListFbIssueLabelsByLCRequest
//
// @return ListFbIssueLabelsByLCResponse
func (client *Client) ListFbIssueLabelsByLC(request *ListFbIssueLabelsByLCRequest) (_result *ListFbIssueLabelsByLCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFbIssueLabelsByLCResponse{}
	_body, _err := client.ListFbIssueLabelsByLCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标签列表
//
// @param request - ListLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLabelsResponse
func (client *Client) ListLabelsWithOptions(request *ListLabelsRequest, runtime *util.RuntimeOptions) (_result *ListLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelContent)) {
		body["LabelContent"] = request.LabelContent
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLabels"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标签列表
//
// @param request - ListLabelsRequest
//
// @return ListLabelsResponse
func (client *Client) ListLabels(request *ListLabelsRequest) (_result *ListLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLabelsResponse{}
	_body, _err := client.ListLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户升级设备信息
//
// @param request - ListTenantDeviceOtaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantDeviceOtaInfoResponse
func (client *Client) ListTenantDeviceOtaInfoWithOptions(request *ListTenantDeviceOtaInfoRequest, runtime *util.RuntimeOptions) (_result *ListTenantDeviceOtaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTenantDeviceOtaInfo"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTenantDeviceOtaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户升级设备信息
//
// @param request - ListTenantDeviceOtaInfoRequest
//
// @return ListTenantDeviceOtaInfoResponse
func (client *Client) ListTenantDeviceOtaInfo(request *ListTenantDeviceOtaInfoRequest) (_result *ListTenantDeviceOtaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTenantDeviceOtaInfoResponse{}
	_body, _err := client.ListTenantDeviceOtaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询终端列表
//
// @param request - ListTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerminalResponse
func (client *Client) ListTerminalWithOptions(request *ListTerminalRequest, runtime *util.RuntimeOptions) (_result *ListTerminalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.BuildId)) {
		body["BuildId"] = request.BuildId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		body["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.InManage)) {
		body["InManage"] = request.InManage
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv4)) {
		body["Ipv4"] = request.Ipv4
	}

	if !tea.BoolValue(util.IsUnset(request.LocationInfo)) {
		body["LocationInfo"] = request.LocationInfo
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyword)) {
		body["SearchKeyword"] = request.SearchKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalGroupId)) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTerminal"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询终端列表
//
// @param request - ListTerminalRequest
//
// @return ListTerminalResponse
func (client *Client) ListTerminal(request *ListTerminalRequest) (_result *ListTerminalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTerminalResponse{}
	_body, _err := client.ListTerminalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询终端基本信息
//
// @param request - ListTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerminalsResponse
func (client *Client) ListTerminalsWithOptions(request *ListTerminalsRequest, runtime *util.RuntimeOptions) (_result *ListTerminalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNumbers)) {
		query["SerialNumbers"] = request.SerialNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.Uuids)) {
		query["Uuids"] = request.Uuids
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyword)) {
		body["SearchKeyword"] = request.SearchKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalGroupId)) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTerminals"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询终端基本信息
//
// @param request - ListTerminalsRequest
//
// @return ListTerminalsResponse
func (client *Client) ListTerminals(request *ListTerminalsRequest) (_result *ListTerminalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTerminalsResponse{}
	_body, _err := client.ListTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询可信设备列表
//
// @param request - ListTrustDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrustDevicesResponse
func (client *Client) ListTrustDevicesWithOptions(request *ListTrustDevicesRequest, runtime *util.RuntimeOptions) (_result *ListTrustDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelContent)) {
		body["LabelContent"] = request.LabelContent
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.UserCustomId)) {
		body["UserCustomId"] = request.UserCustomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrustDevices"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrustDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询可信设备列表
//
// @param request - ListTrustDevicesRequest
//
// @return ListTrustDevicesResponse
func (client *Client) ListTrustDevices(request *ListTrustDevicesRequest) (_result *ListTrustDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTrustDevicesResponse{}
	_body, _err := client.ListTrustDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询问题反馈列表
//
// @param request - ListUserFbAcIssuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserFbAcIssuesResponse
func (client *Client) ListUserFbAcIssuesWithOptions(request *ListUserFbAcIssuesRequest, runtime *util.RuntimeOptions) (_result *ListUserFbAcIssuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorMessage)) {
		body["ErrorMessage"] = request.ErrorMessage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IssueId)) {
		body["IssueId"] = request.IssueId
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedA)) {
		body["ReservedA"] = request.ReservedA
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedB)) {
		body["ReservedB"] = request.ReservedB
	}

	if !tea.BoolValue(util.IsUnset(request.UserEmail)) {
		body["UserEmail"] = request.UserEmail
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserFbAcIssues"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserFbAcIssuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询问题反馈列表
//
// @param request - ListUserFbAcIssuesRequest
//
// @return ListUserFbAcIssuesResponse
func (client *Client) ListUserFbAcIssues(request *ListUserFbAcIssuesRequest) (_result *ListUserFbAcIssuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserFbAcIssuesResponse{}
	_body, _err := client.ListUserFbAcIssuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户反馈问题列表
//
// @param request - ListUserFbIssuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserFbIssuesResponse
func (client *Client) ListUserFbIssuesWithOptions(request *ListUserFbIssuesRequest, runtime *util.RuntimeOptions) (_result *ListUserFbIssuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientModel)) {
		body["ClientModel"] = request.ClientModel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSn)) {
		body["ClientSn"] = request.ClientSn
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		body["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		body["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorCode)) {
		body["ErrorCode"] = request.ErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorMsg)) {
		body["ErrorMsg"] = request.ErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.FbType)) {
		body["FbType"] = request.FbType
	}

	if !tea.BoolValue(util.IsUnset(request.IssueId)) {
		body["IssueId"] = request.IssueId
	}

	if !tea.BoolValue(util.IsUnset(request.IssueLabel)) {
		body["IssueLabel"] = request.IssueLabel
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserEmail)) {
		body["UserEmail"] = request.UserEmail
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WasRead)) {
		body["WasRead"] = request.WasRead
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserFbIssues"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserFbIssuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户反馈问题列表
//
// @param request - ListUserFbIssuesRequest
//
// @return ListUserFbIssuesResponse
func (client *Client) ListUserFbIssues(request *ListUserFbIssuesRequest) (_result *ListUserFbIssuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserFbIssuesResponse{}
	_body, _err := client.ListUserFbIssuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改设备安全入网类型
//
// @param request - ModifyDevicesSecureNetworkTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDevicesSecureNetworkTypeResponse
func (client *Client) ModifyDevicesSecureNetworkTypeWithOptions(request *ModifyDevicesSecureNetworkTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyDevicesSecureNetworkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllDevices)) {
		body["AllDevices"] = request.AllDevices
	}

	if !tea.BoolValue(util.IsUnset(request.SecureNetworkType)) {
		body["SecureNetworkType"] = request.SecureNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNos)) {
		body["SerialNos"] = request.SerialNos
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDevicesSecureNetworkType"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDevicesSecureNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改设备安全入网类型
//
// @param request - ModifyDevicesSecureNetworkTypeRequest
//
// @return ModifyDevicesSecureNetworkTypeResponse
func (client *Client) ModifyDevicesSecureNetworkType(request *ModifyDevicesSecureNetworkTypeRequest) (_result *ModifyDevicesSecureNetworkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDevicesSecureNetworkTypeResponse{}
	_body, _err := client.ModifyDevicesSecureNetworkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 匿名api修改安全入网配置
//
// @param request - ModifySecureNetworkTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecureNetworkTypeResponse
func (client *Client) ModifySecureNetworkTypeWithOptions(request *ModifySecureNetworkTypeRequest, runtime *util.RuntimeOptions) (_result *ModifySecureNetworkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecureNetworkType)) {
		body["SecureNetworkType"] = request.SecureNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecureNetworkType"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecureNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 匿名api修改安全入网配置
//
// @param request - ModifySecureNetworkTypeRequest
//
// @return ModifySecureNetworkTypeResponse
func (client *Client) ModifySecureNetworkType(request *ModifySecureNetworkTypeRequest) (_result *ModifySecureNetworkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecureNetworkTypeResponse{}
	_body, _err := client.ModifySecureNetworkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备注册
//
// @param request - RegisterDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDeviceResponse
func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, runtime *util.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bluetooth)) {
		body["Bluetooth"] = request.Bluetooth
	}

	if !tea.BoolValue(util.IsUnset(request.BuildId)) {
		body["BuildId"] = request.BuildId
	}

	if !tea.BoolValue(util.IsUnset(request.ChipId)) {
		body["ChipId"] = request.ChipId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		body["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.CustomId)) {
		body["CustomId"] = request.CustomId
	}

	if !tea.BoolValue(util.IsUnset(request.EtherMac)) {
		body["EtherMac"] = request.EtherMac
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.Storage)) {
		body["Storage"] = request.Storage
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.Wlan)) {
		body["Wlan"] = request.Wlan
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterDevice"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备注册
//
// @param request - RegisterDeviceRequest
//
// @return RegisterDeviceResponse
func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上报升级信息
//
// @param request - ReportAppOtaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportAppOtaInfoResponse
func (client *Client) ReportAppOtaInfoWithOptions(request *ReportAppOtaInfoRequest, runtime *util.RuntimeOptions) (_result *ReportAppOtaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseVersion)) {
		query["BaseVersion"] = request.BaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientUid)) {
		query["ClientUid"] = request.ClientUid
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersion)) {
		query["TargetVersion"] = request.TargetVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUid)) {
		query["TaskUid"] = request.TaskUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportAppOtaInfo"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportAppOtaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上报升级信息
//
// @param request - ReportAppOtaInfoRequest
//
// @return ReportAppOtaInfoResponse
func (client *Client) ReportAppOtaInfo(request *ReportAppOtaInfoRequest) (_result *ReportAppOtaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportAppOtaInfoResponse{}
	_body, _err := client.ReportAppOtaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级信息上报
//
// @param request - ReportDeviceOtaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportDeviceOtaInfoResponse
func (client *Client) ReportDeviceOtaInfoWithOptions(request *ReportDeviceOtaInfoRequest, runtime *util.RuntimeOptions) (_result *ReportDeviceOtaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseVersion)) {
		body["BaseVersion"] = request.BaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersion)) {
		body["TargetVersion"] = request.TargetVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportDeviceOtaInfo"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportDeviceOtaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级信息上报
//
// @param request - ReportDeviceOtaInfoRequest
//
// @return ReportDeviceOtaInfoResponse
func (client *Client) ReportDeviceOtaInfo(request *ReportDeviceOtaInfoRequest) (_result *ReportDeviceOtaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportDeviceOtaInfoResponse{}
	_body, _err := client.ReportDeviceOtaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用中心用户问题反馈
//
// @param tmpReq - ReportUserFbAcIssueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportUserFbAcIssueResponse
func (client *Client) ReportUserFbAcIssueWithOptions(tmpReq *ReportUserFbAcIssueRequest, runtime *util.RuntimeOptions) (_result *ReportUserFbAcIssueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReportUserFbAcIssueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FileList)) {
		request.FileListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileList, tea.String("FileList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorMsg)) {
		body["ErrorMsg"] = request.ErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.FileListShrink)) {
		body["FileList"] = request.FileListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedA)) {
		body["ReservedA"] = request.ReservedA
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedB)) {
		body["ReservedB"] = request.ReservedB
	}

	if !tea.BoolValue(util.IsUnset(request.UserEmail)) {
		body["UserEmail"] = request.UserEmail
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportUserFbAcIssue"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportUserFbAcIssueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用中心用户问题反馈
//
// @param request - ReportUserFbAcIssueRequest
//
// @return ReportUserFbAcIssueResponse
func (client *Client) ReportUserFbAcIssue(request *ReportUserFbAcIssueRequest) (_result *ReportUserFbAcIssueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportUserFbAcIssueResponse{}
	_body, _err := client.ReportUserFbAcIssueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上报用户反馈问题
//
// @param tmpReq - ReportUserFbIssueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportUserFbIssueResponse
func (client *Client) ReportUserFbIssueWithOptions(tmpReq *ReportUserFbIssueRequest, runtime *util.RuntimeOptions) (_result *ReportUserFbIssueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReportUserFbIssueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FileList)) {
		request.FileListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileList, tea.String("FileList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientModel)) {
		body["ClientModel"] = request.ClientModel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOsName)) {
		body["ClientOsName"] = request.ClientOsName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSn)) {
		body["ClientSn"] = request.ClientSn
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		body["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		body["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopType)) {
		body["DesktopType"] = request.DesktopType
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorCode)) {
		body["ErrorCode"] = request.ErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorMsg)) {
		body["ErrorMsg"] = request.ErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.FbType)) {
		body["FbType"] = request.FbType
	}

	if !tea.BoolValue(util.IsUnset(request.FileListShrink)) {
		body["FileList"] = request.FileListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IssueLabel)) {
		body["IssueLabel"] = request.IssueLabel
	}

	if !tea.BoolValue(util.IsUnset(request.OccurTime)) {
		body["OccurTime"] = request.OccurTime
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedA)) {
		body["ReservedA"] = request.ReservedA
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedB)) {
		body["ReservedB"] = request.ReservedB
	}

	if !tea.BoolValue(util.IsUnset(request.TelNo)) {
		body["TelNo"] = request.TelNo
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserEmail)) {
		body["UserEmail"] = request.UserEmail
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WyId)) {
		body["WyId"] = request.WyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportUserFbIssue"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportUserFbIssueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上报用户反馈问题
//
// @param request - ReportUserFbIssueRequest
//
// @return ReportUserFbIssueResponse
func (client *Client) ReportUserFbIssue(request *ReportUserFbIssueRequest) (_result *ReportUserFbIssueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportUserFbIssueResponse{}
	_body, _err := client.ReportUserFbIssueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 向终端发送运维命令
//
// @param request - SendOpsMessageToTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendOpsMessageToTerminalsResponse
func (client *Client) SendOpsMessageToTerminalsWithOptions(request *SendOpsMessageToTerminalsRequest, runtime *util.RuntimeOptions) (_result *SendOpsMessageToTerminalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Delay)) {
		query["Delay"] = request.Delay
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Msg)) {
		body["Msg"] = request.Msg
	}

	if !tea.BoolValue(util.IsUnset(request.OpsAction)) {
		body["OpsAction"] = request.OpsAction
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uuids)) {
		bodyFlat["Uuids"] = request.Uuids
	}

	if !tea.BoolValue(util.IsUnset(request.WaitForAck)) {
		body["WaitForAck"] = request.WaitForAck
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendOpsMessageToTerminals"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendOpsMessageToTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向终端发送运维命令
//
// @param request - SendOpsMessageToTerminalsRequest
//
// @return SendOpsMessageToTerminalsResponse
func (client *Client) SendOpsMessageToTerminals(request *SendOpsMessageToTerminalsRequest) (_result *SendOpsMessageToTerminalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendOpsMessageToTerminalsResponse{}
	_body, _err := client.SendOpsMessageToTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置租户ota自动开启/关闭
//
// @param request - SetDeviceOtaAutoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeviceOtaAutoStatusResponse
func (client *Client) SetDeviceOtaAutoStatusWithOptions(request *SetDeviceOtaAutoStatusRequest, runtime *util.RuntimeOptions) (_result *SetDeviceOtaAutoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoUpdate)) {
		body["AutoUpdate"] = request.AutoUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUpdateTimeSchedule)) {
		body["AutoUpdateTimeSchedule"] = request.AutoUpdateTimeSchedule
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		body["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ForceUpgrade)) {
		body["ForceUpgrade"] = request.ForceUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDeviceOtaAutoStatus"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDeviceOtaAutoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置租户ota自动开启/关闭
//
// @param request - SetDeviceOtaAutoStatusRequest
//
// @return SetDeviceOtaAutoStatusResponse
func (client *Client) SetDeviceOtaAutoStatus(request *SetDeviceOtaAutoStatusRequest) (_result *SetDeviceOtaAutoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDeviceOtaAutoStatusResponse{}
	_body, _err := client.SetDeviceOtaAutoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 租户设置设备ota任务的状态
//
// @param request - SetDeviceOtaTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeviceOtaTaskStatusResponse
func (client *Client) SetDeviceOtaTaskStatusWithOptions(request *SetDeviceOtaTaskStatusRequest, runtime *util.RuntimeOptions) (_result *SetDeviceOtaTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationStatus)) {
		body["OperationStatus"] = request.OperationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDeviceOtaTaskStatus"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDeviceOtaTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 租户设置设备ota任务的状态
//
// @param request - SetDeviceOtaTaskStatusRequest
//
// @return SetDeviceOtaTaskStatusResponse
func (client *Client) SetDeviceOtaTaskStatus(request *SetDeviceOtaTaskStatusRequest) (_result *SetDeviceOtaTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDeviceOtaTaskStatusResponse{}
	_body, _err := client.SetDeviceOtaTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑免账号登录用户
//
// @param request - UnbindAccountLessLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAccountLessLoginUserResponse
func (client *Client) UnbindAccountLessLoginUserWithOptions(request *UnbindAccountLessLoginUserRequest, runtime *util.RuntimeOptions) (_result *UnbindAccountLessLoginUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindAccountLessLoginUser"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindAccountLessLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑免账号登录用户
//
// @param request - UnbindAccountLessLoginUserRequest
//
// @return UnbindAccountLessLoginUserResponse
func (client *Client) UnbindAccountLessLoginUser(request *UnbindAccountLessLoginUserRequest) (_result *UnbindAccountLessLoginUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindAccountLessLoginUserResponse{}
	_body, _err := client.UnbindAccountLessLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑设备座位
//
// @param tmpReq - UnbindDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDeviceSeatsResponse
func (client *Client) UnbindDeviceSeatsWithOptions(tmpReq *UnbindDeviceSeatsRequest, runtime *util.RuntimeOptions) (_result *UnbindDeviceSeatsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UnbindDeviceSeatsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SerialNoList)) {
		request.SerialNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SerialNoList, tea.String("SerialNoList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNoListShrink)) {
		body["SerialNoList"] = request.SerialNoListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindDeviceSeats"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindDeviceSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑设备座位
//
// @param request - UnbindDeviceSeatsRequest
//
// @return UnbindDeviceSeatsResponse
func (client *Client) UnbindDeviceSeats(request *UnbindDeviceSeatsRequest) (_result *UnbindDeviceSeatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindDeviceSeatsResponse{}
	_body, _err := client.UnbindDeviceSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新设备别名
//
// @param request - UpdateAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliasResponse
func (client *Client) UpdateAliasWithOptions(request *UpdateAliasRequest, runtime *util.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlias"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新设备别名
//
// @param request - UpdateAliasRequest
//
// @return UpdateAliasResponse
func (client *Client) UpdateAlias(request *UpdateAliasRequest) (_result *UpdateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAliasResponse{}
	_body, _err := client.UpdateAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量更新设备绑定的终端用户
//
// @param request - UpdateDeviceBindedEndUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeviceBindedEndUserResponse
func (client *Client) UpdateDeviceBindedEndUserWithOptions(request *UpdateDeviceBindedEndUserRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceBindedEndUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.SourceAdEndUsers)) {
		body["SourceAdEndUsers"] = request.SourceAdEndUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEndUserIds)) {
		body["SourceEndUserIds"] = request.SourceEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.TargetAdEndUsers)) {
		body["TargetAdEndUsers"] = request.TargetAdEndUsers
	}

	if !tea.BoolValue(util.IsUnset(request.TargetEndUserIds)) {
		body["TargetEndUserIds"] = request.TargetEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeviceBindedEndUser"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeviceBindedEndUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新设备绑定的终端用户
//
// @param request - UpdateDeviceBindedEndUserRequest
//
// @return UpdateDeviceBindedEndUserResponse
func (client *Client) UpdateDeviceBindedEndUser(request *UpdateDeviceBindedEndUserRequest) (_result *UpdateDeviceBindedEndUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceBindedEndUserResponse{}
	_body, _err := client.UpdateDeviceBindedEndUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改标签
//
// @param request - UpdateLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLabelResponse
func (client *Client) UpdateLabelWithOptions(request *UpdateLabelRequest, runtime *util.RuntimeOptions) (_result *UpdateLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelContent)) {
		body["LabelContent"] = request.LabelContent
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLabel"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改标签
//
// @param request - UpdateLabelRequest
//
// @return UpdateLabelResponse
func (client *Client) UpdateLabel(request *UpdateLabelRequest) (_result *UpdateLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLabelResponse{}
	_body, _err := client.UpdateLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改终端策略
//
// @param request - UpdateTerminalPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTerminalPolicyResponse
func (client *Client) UpdateTerminalPolicyWithOptions(request *UpdateTerminalPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateTerminalPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayLayout)) {
		body["DisplayLayout"] = request.DisplayLayout
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayResolution)) {
		body["DisplayResolution"] = request.DisplayResolution
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayScaleRatio)) {
		body["DisplayScaleRatio"] = request.DisplayScaleRatio
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAutoLockScreen)) {
		body["EnableAutoLockScreen"] = request.EnableAutoLockScreen
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAutoLogin)) {
		body["EnableAutoLogin"] = request.EnableAutoLogin
	}

	if !tea.BoolValue(util.IsUnset(request.EnableBluetooth)) {
		body["EnableBluetooth"] = request.EnableBluetooth
	}

	if !tea.BoolValue(util.IsUnset(request.EnableModifyPassword)) {
		body["EnableModifyPassword"] = request.EnableModifyPassword
	}

	if !tea.BoolValue(util.IsUnset(request.EnableScheduledShutdown)) {
		body["EnableScheduledShutdown"] = request.EnableScheduledShutdown
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSwitchPersonal)) {
		body["EnableSwitchPersonal"] = request.EnableSwitchPersonal
	}

	if !tea.BoolValue(util.IsUnset(request.EnableWlan)) {
		body["EnableWlan"] = request.EnableWlan
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		body["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeoutAction)) {
		body["IdleTimeoutAction"] = request.IdleTimeoutAction
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PowerButtonDefine)) {
		body["PowerButtonDefine"] = request.PowerButtonDefine
	}

	if !tea.BoolValue(util.IsUnset(request.PowerButtonDefineForAs)) {
		body["PowerButtonDefineForAs"] = request.PowerButtonDefineForAs
	}

	if !tea.BoolValue(util.IsUnset(request.PowerButtonDefineForNs)) {
		body["PowerButtonDefineForNs"] = request.PowerButtonDefineForNs
	}

	if !tea.BoolValue(util.IsUnset(request.PowerOnBehavior)) {
		body["PowerOnBehavior"] = request.PowerOnBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledShutdown)) {
		body["ScheduledShutdown"] = request.ScheduledShutdown
	}

	if !tea.BoolValue(util.IsUnset(request.SettingLock)) {
		body["SettingLock"] = request.SettingLock
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalPolicyId)) {
		body["TerminalPolicyId"] = request.TerminalPolicyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTerminalPolicy"),
		Version:     tea.String("2021-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTerminalPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改终端策略
//
// @param request - UpdateTerminalPolicyRequest
//
// @return UpdateTerminalPolicyResponse
func (client *Client) UpdateTerminalPolicy(request *UpdateTerminalPolicyRequest) (_result *UpdateTerminalPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTerminalPolicyResponse{}
	_body, _err := client.UpdateTerminalPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
