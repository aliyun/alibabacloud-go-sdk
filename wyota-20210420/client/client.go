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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ActivateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ActivateDeviceResponse) SetBody(v *ActivateDeviceResponseBody) *ActivateDeviceResponse {
	s.Body = v
	return s
}

type AddDeviceFromSNRequest struct {
	Alias         *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	LabelContents *string `json:"LabelContents,omitempty" xml:"LabelContents,omitempty"`
	SerialNo      *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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

func (s *AddDeviceFromSNRequest) SetLabelContents(v string) *AddDeviceFromSNRequest {
	s.LabelContents = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetSerialNo(v string) *AddDeviceFromSNRequest {
	s.SerialNo = &v
	return s
}

type AddDeviceFromSNResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDeviceFromSNResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddDeviceFromSNResponse) SetBody(v *AddDeviceFromSNResponseBody) *AddDeviceFromSNResponse {
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDevicesFromCSVResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddLabelsResponse) SetBody(v *AddLabelsResponseBody) *AddLabelsResponse {
	s.Body = v
	return s
}

type AttachEndUsersRequest struct {
	EndUserIds *string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty"`
	SerialNo   *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachEndUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachEndUsersResponse) SetBody(v *AttachEndUsersResponseBody) *AttachEndUsersResponse {
	s.Body = v
	return s
}

type AttachLabelRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachLabelResponse) SetBody(v *AttachLabelResponseBody) *AttachLabelResponse {
	s.Body = v
	return s
}

type AttachLabelsRequest struct {
	LabelIds *string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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

type AttachLabelsResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachLabelsResponse) SetBody(v *AttachLabelsResponseBody) *AttachLabelsResponse {
	s.Body = v
	return s
}

type CheckUuidValidRequest struct {
	BuildId  *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ChipId   *string `json:"ChipId,omitempty" xml:"ChipId,omitempty"`
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CheckUuidValidRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUuidValidRequest) GoString() string {
	return s.String()
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

func (s *CheckUuidValidRequest) SetSerialNo(v string) *CheckUuidValidRequest {
	s.SerialNo = &v
	return s
}

func (s *CheckUuidValidRequest) SetUuid(v string) *CheckUuidValidRequest {
	s.Uuid = &v
	return s
}

type CheckUuidValidResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckUuidValidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckUuidValidResponse) SetBody(v *CheckUuidValidResponseBody) *CheckUuidValidResponse {
	s.Body = v
	return s
}

type DeleteDevicesRequest struct {
	Force     *string `json:"Force,omitempty" xml:"Force,omitempty"`
	SerialNos *string `json:"SerialNos,omitempty" xml:"SerialNos,omitempty"`
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

type DeleteDevicesResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDevicesResponse) SetBody(v *DeleteDevicesResponseBody) *DeleteDevicesResponse {
	s.Body = v
	return s
}

type DeleteLabelRequest struct {
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteLabelResponse) SetBody(v *DeleteLabelResponseBody) *DeleteLabelResponse {
	s.Body = v
	return s
}

type DescribeDeviceSeatsRequest struct {
	PageNumber   *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

type DescribeDeviceSeatsShrinkRequest struct {
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SerialNoListShrink *string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty"`
	SiteId             *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TenantId           *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeDeviceSeatsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceSeatsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsShrinkRequest) SetPageNumber(v int32) *DescribeDeviceSeatsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeviceSeatsShrinkRequest) SetPageSize(v int32) *DescribeDeviceSeatsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceSeatsShrinkRequest) SetSerialNoListShrink(v string) *DescribeDeviceSeatsShrinkRequest {
	s.SerialNoListShrink = &v
	return s
}

func (s *DescribeDeviceSeatsShrinkRequest) SetSiteId(v string) *DescribeDeviceSeatsShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDeviceSeatsShrinkRequest) SetTenantId(v string) *DescribeDeviceSeatsShrinkRequest {
	s.TenantId = &v
	return s
}

type DescribeDeviceSeatsResponseBody struct {
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*DescribeDeviceSeatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDeviceSeatsResponse) SetBody(v *DescribeDeviceSeatsResponseBody) *DescribeDeviceSeatsResponse {
	s.Body = v
	return s
}

type DetachEndUsersRequest struct {
	EndUserIds *string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty"`
	SerialNo   *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachEndUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachEndUsersResponse) SetBody(v *DetachEndUsersResponseBody) *DetachEndUsersResponse {
	s.Body = v
	return s
}

type DetachLabelRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachLabelResponse) SetBody(v *DetachLabelResponseBody) *DetachLabelResponse {
	s.Body = v
	return s
}

type DetachLabelsRequest struct {
	LabelIds *string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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

type DetachLabelsResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachLabelsResponse) SetBody(v *DetachLabelsResponseBody) *DetachLabelsResponse {
	s.Body = v
	return s
}

type GetAppOtaLatestVersionRequest struct {
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	ClientType  *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientUid   *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	OsType      *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
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
	Code    *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetAppOtaLatestVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppOtaLatestVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAppOtaLatestVersionResponse) SetBody(v *GetAppOtaLatestVersionResponseBody) *GetAppOtaLatestVersionResponse {
	s.Body = v
	return s
}

type GetDeviceConfigsRequest struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
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

func (s *GetDeviceConfigsRequest) SetUserCustomId(v string) *GetDeviceConfigsRequest {
	s.UserCustomId = &v
	return s
}

type GetDeviceConfigsResponseBody struct {
	Code    *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetDeviceConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	SecureNetworkType *string `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	Uuid              *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetDeviceConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponseBodyData) SetSecureNetworkType(v string) *GetDeviceConfigsResponseBodyData {
	s.SecureNetworkType = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUuid(v string) *GetDeviceConfigsResponseBodyData {
	s.Uuid = &v
	return s
}

type GetDeviceConfigsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeviceConfigsResponse) SetBody(v *GetDeviceConfigsResponseBody) *GetDeviceConfigsResponse {
	s.Body = v
	return s
}

type GetDeviceOtaAutoStatusResponseBody struct {
	Code    *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetDeviceOtaAutoStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeviceOtaAutoStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaAutoStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetStatus(v int32) *GetDeviceOtaAutoStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetDeviceOtaAutoStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceOtaAutoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeviceOtaAutoStatusResponse) SetBody(v *GetDeviceOtaAutoStatusResponseBody) *GetDeviceOtaAutoStatusResponse {
	s.Body = v
	return s
}

type GetDeviceOtaInfoRequest struct {
	BaseVersion       *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	DeviceId          *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data    *GetDeviceOtaInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ForceUpgrade     *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	LocalDownloadUrl *string `json:"LocalDownloadUrl,omitempty" xml:"LocalDownloadUrl,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Model            *string `json:"Model,omitempty" xml:"Model,omitempty"`
	ReleaseNote      *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn    *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version          *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionCode      *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType      *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s GetDeviceOtaInfoResponseBodyDataVersion) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceOtaInfoResponseBodyDataVersion) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetCreator(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Creator = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.DownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetForceUpgrade(v int32) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.ForceUpgrade = &v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeviceOtaInfoResponse) SetBody(v *GetDeviceOtaInfoResponseBody) *GetDeviceOtaInfoResponse {
	s.Body = v
	return s
}

type GetDeviceOtaInfoTestRequest struct {
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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
	Code    *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetDeviceOtaInfoTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceOtaInfoTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeviceOtaInfoTestResponse) SetBody(v *GetDeviceOtaInfoTestResponseBody) *GetDeviceOtaInfoTestResponse {
	s.Body = v
	return s
}

type GetDeviceOtaTaskVersionInfoRequest struct {
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
	Code    *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetDeviceOtaTaskVersionInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceOtaTaskVersionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeviceOtaTaskVersionInfoResponse) SetBody(v *GetDeviceOtaTaskVersionInfoResponseBody) *GetDeviceOtaTaskVersionInfoResponse {
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
	Code    *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetFbOssConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFbOssConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFbOssConfigResponse) SetBody(v *GetFbOssConfigResponseBody) *GetFbOssConfigResponse {
	s.Body = v
	return s
}

type GetOssConfigResponseBody struct {
	Code    *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetOssConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessKeyId  *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	EndPoint     *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	OssPolicy    *string `json:"OssPolicy,omitempty" xml:"OssPolicy,omitempty"`
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
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

type GetOssConfigResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOssConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOssConfigResponse) SetBody(v *GetOssConfigResponseBody) *GetOssConfigResponse {
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
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceOtaTaskByTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDeviceOtaTaskByTenantResponse) SetBody(v *ListDeviceOtaTaskByTenantResponseBody) *ListDeviceOtaTaskByTenantResponse {
	s.Body = v
	return s
}

type ListDevicesRequest struct {
	Alias        *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BuildId      *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	Model        *string `json:"Model,omitempty" xml:"Model,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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

type ListDevicesResponseBody struct {
	Code       *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	ActiveTime        *string                                   `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	Alias             *string                                   `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AutoType          *string                                   `json:"AutoType,omitempty" xml:"AutoType,omitempty"`
	Bluetooth         *string                                   `json:"Bluetooth,omitempty" xml:"Bluetooth,omitempty"`
	BuildId           *string                                   `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientId          *string                                   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType        *string                                   `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	EndUserList       []*ListDevicesResponseBodyDataEndUserList `json:"EndUserList,omitempty" xml:"EndUserList,omitempty" type:"Repeated"`
	EtherMac          *string                                   `json:"EtherMac,omitempty" xml:"EtherMac,omitempty"`
	Id                *int64                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	IsActive          *string                                   `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	LabelList         []*ListDevicesResponseBodyDataLabelList   `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	Model             *string                                   `json:"Model,omitempty" xml:"Model,omitempty"`
	OrderId           *string                                   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	SecureNetworkType *string                                   `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	SerialNo          *string                                   `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Source            *string                                   `json:"Source,omitempty" xml:"Source,omitempty"`
	TenantId          *string                                   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Uuid              *string                                   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Wlan              *string                                   `json:"Wlan,omitempty" xml:"Wlan,omitempty"`
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

func (s *ListDevicesResponseBodyData) SetEndUserList(v []*ListDevicesResponseBodyDataEndUserList) *ListDevicesResponseBodyData {
	s.EndUserList = v
	return s
}

func (s *ListDevicesResponseBodyData) SetEtherMac(v string) *ListDevicesResponseBodyData {
	s.EtherMac = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetId(v int64) *ListDevicesResponseBodyData {
	s.Id = &v
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

func (s *ListDevicesResponseBodyData) SetModel(v string) *ListDevicesResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetOrderId(v string) *ListDevicesResponseBodyData {
	s.OrderId = &v
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

func (s *ListDevicesResponseBodyData) SetSource(v string) *ListDevicesResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetTenantId(v string) *ListDevicesResponseBodyData {
	s.TenantId = &v
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

type ListDevicesResponseBodyDataEndUserList struct {
	BindTime  *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SerialNo  *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDevicesResponseBodyDataEndUserList) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDataEndUserList) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataEndUserList) SetBindTime(v string) *ListDevicesResponseBodyDataEndUserList {
	s.BindTime = &v
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

type ListDevicesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDevicesResponse) SetBody(v *ListDevicesResponseBody) *ListDevicesResponse {
	s.Body = v
	return s
}

type ListFbIssueLabelsResponseBody struct {
	Code    *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ListFbIssueLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFbIssueLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code    *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ListFbIssueLabelsByLCResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFbIssueLabelsByLCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListLabelsResponse) SetBody(v *ListLabelsResponseBody) *ListLabelsResponse {
	s.Body = v
	return s
}

type ListTenantDeviceOtaInfoRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId     *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTenantDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTenantDeviceOtaInfoResponse) SetBody(v *ListTenantDeviceOtaInfoResponseBody) *ListTenantDeviceOtaInfoResponse {
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
	Code    *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ListUserFbAcIssuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserFbAcIssuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data       *ListUserFbIssuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserFbIssuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListUserFbIssuesResponse) SetBody(v *ListUserFbIssuesResponseBody) *ListUserFbIssuesResponse {
	s.Body = v
	return s
}

type ModifyDevicesSecureNetworkTypeRequest struct {
	AllDevices        *int64  `json:"AllDevices,omitempty" xml:"AllDevices,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDevicesSecureNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDevicesSecureNetworkTypeResponse) SetBody(v *ModifyDevicesSecureNetworkTypeResponseBody) *ModifyDevicesSecureNetworkTypeResponse {
	s.Body = v
	return s
}

type ModifySecureNetworkTypeRequest struct {
	SecureNetworkType *string `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	SerialNo          *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecureNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code    *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *RegisterDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportAppOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize  *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
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
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ReportUserFbAcIssueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportUserFbAcIssueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ErrorCode     *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg      *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType        *int32                              `json:"FbType,omitempty" xml:"FbType,omitempty"`
	FileList      []*ReportUserFbIssueRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	IssueLabel    *string                             `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	OccurTime     *int64                              `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	ReservedA     *string                             `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB     *string                             `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	Title         *string                             `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail     *string                             `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId        *string                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string                             `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WorkspaceId   *string                             `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

type ReportUserFbIssueRequestFileList struct {
	FileMd5   *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
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
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg       *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType         *int32  `json:"FbType,omitempty" xml:"FbType,omitempty"`
	FileListShrink *string `json:"FileList,omitempty" xml:"FileList,omitempty"`
	IssueLabel     *string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	OccurTime      *int64  `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	ReservedA      *string `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB      *string `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail      *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

type ReportUserFbIssueResponseBody struct {
	Code    *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ReportUserFbIssueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportUserFbIssueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReportUserFbIssueResponse) SetBody(v *ReportUserFbIssueResponseBody) *ReportUserFbIssueResponse {
	s.Body = v
	return s
}

type SetDeviceOtaAutoStatusRequest struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDeviceOtaAutoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceOtaAutoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaAutoStatusRequest) SetStatus(v string) *SetDeviceOtaAutoStatusRequest {
	s.Status = &v
	return s
}

type SetDeviceOtaAutoStatusResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDeviceOtaAutoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetDeviceOtaAutoStatusResponse) SetBody(v *SetDeviceOtaAutoStatusResponseBody) *SetDeviceOtaAutoStatusResponse {
	s.Body = v
	return s
}

type SetDeviceOtaTaskStatusRequest struct {
	OperationStatus *int32 `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	TaskId          *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDeviceOtaTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetDeviceOtaTaskStatusResponse) SetBody(v *SetDeviceOtaTaskStatusResponseBody) *SetDeviceOtaTaskStatusResponse {
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UnbindDeviceSeatsResponse) SetBody(v *UnbindDeviceSeatsResponseBody) *UnbindDeviceSeatsResponse {
	s.Body = v
	return s
}

type UpdateAliasRequest struct {
	Alias    *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
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

type UpdateAliasResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAliasResponse) SetBody(v *UpdateAliasResponseBody) *UpdateAliasResponse {
	s.Body = v
	return s
}

type UpdateDeviceBindedEndUserRequest struct {
	SerialNo         *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SourceEndUserIds *string `json:"SourceEndUserIds,omitempty" xml:"SourceEndUserIds,omitempty"`
	TargetEndUserIds *string `json:"TargetEndUserIds,omitempty" xml:"TargetEndUserIds,omitempty"`
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

func (s *UpdateDeviceBindedEndUserRequest) SetSourceEndUserIds(v string) *UpdateDeviceBindedEndUserRequest {
	s.SourceEndUserIds = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetTargetEndUserIds(v string) *UpdateDeviceBindedEndUserRequest {
	s.TargetEndUserIds = &v
	return s
}

type UpdateDeviceBindedEndUserResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceBindedEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateLabelResponse) SetBody(v *UpdateLabelResponseBody) *UpdateLabelResponse {
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

func (client *Client) AddDeviceFromSNWithOptions(request *AddDeviceFromSNRequest, runtime *util.RuntimeOptions) (_result *AddDeviceFromSNResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.LabelContents)) {
		body["LabelContents"] = request.LabelContents
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

func (client *Client) CheckUuidValidWithOptions(request *CheckUuidValidRequest, runtime *util.RuntimeOptions) (_result *CheckUuidValidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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

func (client *Client) DeleteDevicesWithOptions(request *DeleteDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		body["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNos)) {
		body["SerialNos"] = request.SerialNos
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
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

func (client *Client) DescribeDeviceSeatsWithOptions(tmpReq *DescribeDeviceSeatsRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceSeatsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeDeviceSeatsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SerialNoList)) {
		request.SerialNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SerialNoList, tea.String("SerialNoList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNoListShrink)) {
		body["SerialNoList"] = request.SerialNoListShrink
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

func (client *Client) GetDeviceConfigsWithOptions(request *GetDeviceConfigsRequest, runtime *util.RuntimeOptions) (_result *GetDeviceConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
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

func (client *Client) GetDeviceOtaAutoStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetDeviceOtaAutoStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
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

func (client *Client) GetDeviceOtaAutoStatus() (_result *GetDeviceOtaAutoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceOtaAutoStatusResponse{}
	_body, _err := client.GetDeviceOtaAutoStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceOtaInfoWithOptions(request *GetDeviceOtaInfoRequest, runtime *util.RuntimeOptions) (_result *GetDeviceOtaInfoResponse, _err error) {
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

func (client *Client) GetOssConfigWithOptions(runtime *util.RuntimeOptions) (_result *GetOssConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
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

func (client *Client) GetOssConfig() (_result *GetOssConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssConfigResponse{}
	_body, _err := client.GetOssConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
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

func (client *Client) SetDeviceOtaAutoStatusWithOptions(request *SetDeviceOtaAutoStatusRequest, runtime *util.RuntimeOptions) (_result *SetDeviceOtaAutoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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

func (client *Client) UpdateDeviceBindedEndUserWithOptions(request *UpdateDeviceBindedEndUserRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceBindedEndUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		body["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEndUserIds)) {
		body["SourceEndUserIds"] = request.SourceEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.TargetEndUserIds)) {
		body["TargetEndUserIds"] = request.TargetEndUserIds
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
