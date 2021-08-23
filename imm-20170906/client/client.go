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

type CompareImageFacesRequest struct {
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageUriA *string `json:"ImageUriA,omitempty" xml:"ImageUriA,omitempty"`
	ImageUriB *string `json:"ImageUriB,omitempty" xml:"ImageUriB,omitempty"`
	FaceIdA   *string `json:"FaceIdA,omitempty" xml:"FaceIdA,omitempty"`
	FaceIdB   *string `json:"FaceIdB,omitempty" xml:"FaceIdB,omitempty"`
}

func (s CompareImageFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareImageFacesRequest) SetProject(v string) *CompareImageFacesRequest {
	s.Project = &v
	return s
}

func (s *CompareImageFacesRequest) SetSetId(v string) *CompareImageFacesRequest {
	s.SetId = &v
	return s
}

func (s *CompareImageFacesRequest) SetImageUriA(v string) *CompareImageFacesRequest {
	s.ImageUriA = &v
	return s
}

func (s *CompareImageFacesRequest) SetImageUriB(v string) *CompareImageFacesRequest {
	s.ImageUriB = &v
	return s
}

func (s *CompareImageFacesRequest) SetFaceIdA(v string) *CompareImageFacesRequest {
	s.FaceIdA = &v
	return s
}

func (s *CompareImageFacesRequest) SetFaceIdB(v string) *CompareImageFacesRequest {
	s.FaceIdB = &v
	return s
}

type CompareImageFacesResponseBody struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Similarity *float32                            `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
	SetId      *string                             `json:"SetId,omitempty" xml:"SetId,omitempty"`
	FaceA      *CompareImageFacesResponseBodyFaceA `json:"FaceA,omitempty" xml:"FaceA,omitempty" type:"Struct"`
	FaceB      *CompareImageFacesResponseBodyFaceB `json:"FaceB,omitempty" xml:"FaceB,omitempty" type:"Struct"`
}

func (s CompareImageFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBody) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBody) SetRequestId(v string) *CompareImageFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareImageFacesResponseBody) SetSimilarity(v float32) *CompareImageFacesResponseBody {
	s.Similarity = &v
	return s
}

func (s *CompareImageFacesResponseBody) SetSetId(v string) *CompareImageFacesResponseBody {
	s.SetId = &v
	return s
}

func (s *CompareImageFacesResponseBody) SetFaceA(v *CompareImageFacesResponseBodyFaceA) *CompareImageFacesResponseBody {
	s.FaceA = v
	return s
}

func (s *CompareImageFacesResponseBody) SetFaceB(v *CompareImageFacesResponseBodyFaceB) *CompareImageFacesResponseBody {
	s.FaceB = v
	return s
}

type CompareImageFacesResponseBodyFaceA struct {
	FaceId         *string                                           `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceAttributes *CompareImageFacesResponseBodyFaceAFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
}

func (s CompareImageFacesResponseBodyFaceA) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceA) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceA) SetFaceId(v string) *CompareImageFacesResponseBodyFaceA {
	s.FaceId = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceA) SetFaceAttributes(v *CompareImageFacesResponseBodyFaceAFaceAttributes) *CompareImageFacesResponseBodyFaceA {
	s.FaceAttributes = v
	return s
}

type CompareImageFacesResponseBodyFaceAFaceAttributes struct {
	FaceBoundary *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
}

func (s CompareImageFacesResponseBodyFaceAFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceAFaceAttributes) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributes) SetFaceBoundary(v *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) *CompareImageFacesResponseBodyFaceAFaceAttributes {
	s.FaceBoundary = v
	return s
}

type CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) SetLeft(v int32) *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) SetTop(v int32) *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) SetWidth(v int32) *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) SetHeight(v int32) *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

type CompareImageFacesResponseBodyFaceB struct {
	FaceId         *string                                           `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceAttributes *CompareImageFacesResponseBodyFaceBFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
}

func (s CompareImageFacesResponseBodyFaceB) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceB) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceB) SetFaceId(v string) *CompareImageFacesResponseBodyFaceB {
	s.FaceId = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceB) SetFaceAttributes(v *CompareImageFacesResponseBodyFaceBFaceAttributes) *CompareImageFacesResponseBodyFaceB {
	s.FaceAttributes = v
	return s
}

type CompareImageFacesResponseBodyFaceBFaceAttributes struct {
	FaceBoundary *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
}

func (s CompareImageFacesResponseBodyFaceBFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceBFaceAttributes) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributes) SetFaceBoundary(v *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) *CompareImageFacesResponseBodyFaceBFaceAttributes {
	s.FaceBoundary = v
	return s
}

type CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) SetLeft(v int32) *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) SetTop(v int32) *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) SetWidth(v int32) *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) SetHeight(v int32) *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

type CompareImageFacesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CompareImageFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompareImageFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponse) SetHeaders(v map[string]*string) *CompareImageFacesResponse {
	s.Headers = v
	return s
}

func (s *CompareImageFacesResponse) SetBody(v *CompareImageFacesResponseBody) *CompareImageFacesResponse {
	s.Body = v
	return s
}

type ConvertOfficeFormatRequest struct {
	Project        *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcUri         *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	TgtType        *string `json:"TgtType,omitempty" xml:"TgtType,omitempty"`
	TgtUri         *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	SrcType        *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StartPage      *int64  `json:"StartPage,omitempty" xml:"StartPage,omitempty"`
	EndPage        *int64  `json:"EndPage,omitempty" xml:"EndPage,omitempty"`
	MaxSheetRow    *int64  `json:"MaxSheetRow,omitempty" xml:"MaxSheetRow,omitempty"`
	MaxSheetCol    *int64  `json:"MaxSheetCol,omitempty" xml:"MaxSheetCol,omitempty"`
	MaxSheetCount  *int64  `json:"MaxSheetCount,omitempty" xml:"MaxSheetCount,omitempty"`
	SheetOnePage   *bool   `json:"SheetOnePage,omitempty" xml:"SheetOnePage,omitempty"`
	ModelId        *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	TgtFilePrefix  *string `json:"TgtFilePrefix,omitempty" xml:"TgtFilePrefix,omitempty"`
	TgtFileSuffix  *string `json:"TgtFileSuffix,omitempty" xml:"TgtFileSuffix,omitempty"`
	TgtFilePages   *string `json:"TgtFilePages,omitempty" xml:"TgtFilePages,omitempty"`
	FitToPagesTall *bool   `json:"FitToPagesTall,omitempty" xml:"FitToPagesTall,omitempty"`
	FitToPagesWide *bool   `json:"FitToPagesWide,omitempty" xml:"FitToPagesWide,omitempty"`
	PdfVector      *bool   `json:"PdfVector,omitempty" xml:"PdfVector,omitempty"`
	Hidecomments   *bool   `json:"Hidecomments,omitempty" xml:"Hidecomments,omitempty"`
}

func (s ConvertOfficeFormatRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertOfficeFormatRequest) GoString() string {
	return s.String()
}

func (s *ConvertOfficeFormatRequest) SetProject(v string) *ConvertOfficeFormatRequest {
	s.Project = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetSrcUri(v string) *ConvertOfficeFormatRequest {
	s.SrcUri = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtType(v string) *ConvertOfficeFormatRequest {
	s.TgtType = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtUri(v string) *ConvertOfficeFormatRequest {
	s.TgtUri = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetSrcType(v string) *ConvertOfficeFormatRequest {
	s.SrcType = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetStartPage(v int64) *ConvertOfficeFormatRequest {
	s.StartPage = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetEndPage(v int64) *ConvertOfficeFormatRequest {
	s.EndPage = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetMaxSheetRow(v int64) *ConvertOfficeFormatRequest {
	s.MaxSheetRow = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetMaxSheetCol(v int64) *ConvertOfficeFormatRequest {
	s.MaxSheetCol = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetMaxSheetCount(v int64) *ConvertOfficeFormatRequest {
	s.MaxSheetCount = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetSheetOnePage(v bool) *ConvertOfficeFormatRequest {
	s.SheetOnePage = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetModelId(v string) *ConvertOfficeFormatRequest {
	s.ModelId = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetPassword(v string) *ConvertOfficeFormatRequest {
	s.Password = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtFilePrefix(v string) *ConvertOfficeFormatRequest {
	s.TgtFilePrefix = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtFileSuffix(v string) *ConvertOfficeFormatRequest {
	s.TgtFileSuffix = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtFilePages(v string) *ConvertOfficeFormatRequest {
	s.TgtFilePages = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetFitToPagesTall(v bool) *ConvertOfficeFormatRequest {
	s.FitToPagesTall = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetFitToPagesWide(v bool) *ConvertOfficeFormatRequest {
	s.FitToPagesWide = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetPdfVector(v bool) *ConvertOfficeFormatRequest {
	s.PdfVector = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetHidecomments(v bool) *ConvertOfficeFormatRequest {
	s.Hidecomments = &v
	return s
}

type ConvertOfficeFormatResponseBody struct {
	PageCount *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertOfficeFormatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertOfficeFormatResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertOfficeFormatResponseBody) SetPageCount(v int32) *ConvertOfficeFormatResponseBody {
	s.PageCount = &v
	return s
}

func (s *ConvertOfficeFormatResponseBody) SetRequestId(v string) *ConvertOfficeFormatResponseBody {
	s.RequestId = &v
	return s
}

type ConvertOfficeFormatResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConvertOfficeFormatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConvertOfficeFormatResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertOfficeFormatResponse) GoString() string {
	return s.String()
}

func (s *ConvertOfficeFormatResponse) SetHeaders(v map[string]*string) *ConvertOfficeFormatResponse {
	s.Headers = v
	return s
}

func (s *ConvertOfficeFormatResponse) SetBody(v *ConvertOfficeFormatResponseBody) *ConvertOfficeFormatResponse {
	s.Body = v
	return s
}

type CreateGrabFrameTaskRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	VideoUri        *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	TargetList      *string `json:"TargetList,omitempty" xml:"TargetList,omitempty"`
	CustomMessage   *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
}

func (s CreateGrabFrameTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGrabFrameTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateGrabFrameTaskRequest) SetProject(v string) *CreateGrabFrameTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetVideoUri(v string) *CreateGrabFrameTaskRequest {
	s.VideoUri = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetNotifyTopicName(v string) *CreateGrabFrameTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetNotifyEndpoint(v string) *CreateGrabFrameTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetTargetList(v string) *CreateGrabFrameTaskRequest {
	s.TargetList = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetCustomMessage(v string) *CreateGrabFrameTaskRequest {
	s.CustomMessage = &v
	return s
}

type CreateGrabFrameTaskResponseBody struct {
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGrabFrameTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGrabFrameTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGrabFrameTaskResponseBody) SetTaskType(v string) *CreateGrabFrameTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateGrabFrameTaskResponseBody) SetRequestId(v string) *CreateGrabFrameTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGrabFrameTaskResponseBody) SetTaskId(v string) *CreateGrabFrameTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateGrabFrameTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGrabFrameTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGrabFrameTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGrabFrameTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateGrabFrameTaskResponse) SetHeaders(v map[string]*string) *CreateGrabFrameTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateGrabFrameTaskResponse) SetBody(v *CreateGrabFrameTaskResponseBody) *CreateGrabFrameTaskResponse {
	s.Body = v
	return s
}

type CreateGroupFacesJobRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
}

func (s CreateGroupFacesJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupFacesJobRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFacesJobRequest) SetProject(v string) *CreateGroupFacesJobRequest {
	s.Project = &v
	return s
}

func (s *CreateGroupFacesJobRequest) SetSetId(v string) *CreateGroupFacesJobRequest {
	s.SetId = &v
	return s
}

func (s *CreateGroupFacesJobRequest) SetNotifyTopicName(v string) *CreateGroupFacesJobRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateGroupFacesJobRequest) SetNotifyEndpoint(v string) *CreateGroupFacesJobRequest {
	s.NotifyEndpoint = &v
	return s
}

type CreateGroupFacesJobResponseBody struct {
	JobType   *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateGroupFacesJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupFacesJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupFacesJobResponseBody) SetJobType(v string) *CreateGroupFacesJobResponseBody {
	s.JobType = &v
	return s
}

func (s *CreateGroupFacesJobResponseBody) SetRequestId(v string) *CreateGroupFacesJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupFacesJobResponseBody) SetSetId(v string) *CreateGroupFacesJobResponseBody {
	s.SetId = &v
	return s
}

func (s *CreateGroupFacesJobResponseBody) SetJobId(v string) *CreateGroupFacesJobResponseBody {
	s.JobId = &v
	return s
}

type CreateGroupFacesJobResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupFacesJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupFacesJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupFacesJobResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupFacesJobResponse) SetHeaders(v map[string]*string) *CreateGroupFacesJobResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupFacesJobResponse) SetBody(v *CreateGroupFacesJobResponseBody) *CreateGroupFacesJobResponse {
	s.Body = v
	return s
}

type CreateImageProcessTaskRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri        *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	TargetList      *string `json:"TargetList,omitempty" xml:"TargetList,omitempty"`
}

func (s CreateImageProcessTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageProcessTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageProcessTaskRequest) SetProject(v string) *CreateImageProcessTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateImageProcessTaskRequest) SetImageUri(v string) *CreateImageProcessTaskRequest {
	s.ImageUri = &v
	return s
}

func (s *CreateImageProcessTaskRequest) SetNotifyTopicName(v string) *CreateImageProcessTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateImageProcessTaskRequest) SetNotifyEndpoint(v string) *CreateImageProcessTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateImageProcessTaskRequest) SetTargetList(v string) *CreateImageProcessTaskRequest {
	s.TargetList = &v
	return s
}

type CreateImageProcessTaskResponseBody struct {
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageProcessTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageProcessTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageProcessTaskResponseBody) SetTaskType(v string) *CreateImageProcessTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateImageProcessTaskResponseBody) SetRequestId(v string) *CreateImageProcessTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageProcessTaskResponseBody) SetTaskId(v string) *CreateImageProcessTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateImageProcessTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageProcessTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageProcessTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageProcessTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageProcessTaskResponse) SetHeaders(v map[string]*string) *CreateImageProcessTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageProcessTaskResponse) SetBody(v *CreateImageProcessTaskResponseBody) *CreateImageProcessTaskResponse {
	s.Body = v
	return s
}

type CreateMediaComplexTaskRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Parameters      *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
}

func (s CreateMediaComplexTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaComplexTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaComplexTaskRequest) SetProject(v string) *CreateMediaComplexTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateMediaComplexTaskRequest) SetParameters(v string) *CreateMediaComplexTaskRequest {
	s.Parameters = &v
	return s
}

func (s *CreateMediaComplexTaskRequest) SetNotifyTopicName(v string) *CreateMediaComplexTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateMediaComplexTaskRequest) SetNotifyEndpoint(v string) *CreateMediaComplexTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

type CreateMediaComplexTaskResponseBody struct {
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateMediaComplexTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaComplexTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaComplexTaskResponseBody) SetTaskType(v string) *CreateMediaComplexTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateMediaComplexTaskResponseBody) SetRequestId(v string) *CreateMediaComplexTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaComplexTaskResponseBody) SetTaskId(v string) *CreateMediaComplexTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateMediaComplexTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMediaComplexTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMediaComplexTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaComplexTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaComplexTaskResponse) SetHeaders(v map[string]*string) *CreateMediaComplexTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaComplexTaskResponse) SetBody(v *CreateMediaComplexTaskResponseBody) *CreateMediaComplexTaskResponse {
	s.Body = v
	return s
}

type CreateMergeFaceGroupsJobRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	GroupIdFrom     *string `json:"GroupIdFrom,omitempty" xml:"GroupIdFrom,omitempty"`
	GroupIdTo       *string `json:"GroupIdTo,omitempty" xml:"GroupIdTo,omitempty"`
	CustomMessage   *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
}

func (s CreateMergeFaceGroupsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeFaceGroupsJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMergeFaceGroupsJobRequest) SetProject(v string) *CreateMergeFaceGroupsJobRequest {
	s.Project = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetSetId(v string) *CreateMergeFaceGroupsJobRequest {
	s.SetId = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetNotifyTopicName(v string) *CreateMergeFaceGroupsJobRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetNotifyEndpoint(v string) *CreateMergeFaceGroupsJobRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetGroupIdFrom(v string) *CreateMergeFaceGroupsJobRequest {
	s.GroupIdFrom = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetGroupIdTo(v string) *CreateMergeFaceGroupsJobRequest {
	s.GroupIdTo = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetCustomMessage(v string) *CreateMergeFaceGroupsJobRequest {
	s.CustomMessage = &v
	return s
}

type CreateMergeFaceGroupsJobResponseBody struct {
	GroupIdFrom *string `json:"GroupIdFrom,omitempty" xml:"GroupIdFrom,omitempty"`
	JobType     *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId       *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	GroupIdTo   *string `json:"GroupIdTo,omitempty" xml:"GroupIdTo,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateMergeFaceGroupsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeFaceGroupsJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetGroupIdFrom(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.GroupIdFrom = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetJobType(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.JobType = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetRequestId(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetSetId(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.SetId = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetGroupIdTo(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.GroupIdTo = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetJobId(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.JobId = &v
	return s
}

type CreateMergeFaceGroupsJobResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMergeFaceGroupsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMergeFaceGroupsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeFaceGroupsJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMergeFaceGroupsJobResponse) SetHeaders(v map[string]*string) *CreateMergeFaceGroupsJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMergeFaceGroupsJobResponse) SetBody(v *CreateMergeFaceGroupsJobResponseBody) *CreateMergeFaceGroupsJobResponse {
	s.Body = v
	return s
}

type CreateOfficeConversionTaskRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcUri          *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	TgtType         *string `json:"TgtType,omitempty" xml:"TgtType,omitempty"`
	TgtUri          *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	SrcType         *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StartPage       *int64  `json:"StartPage,omitempty" xml:"StartPage,omitempty"`
	EndPage         *int64  `json:"EndPage,omitempty" xml:"EndPage,omitempty"`
	MaxSheetRow     *int64  `json:"MaxSheetRow,omitempty" xml:"MaxSheetRow,omitempty"`
	MaxSheetCol     *int64  `json:"MaxSheetCol,omitempty" xml:"MaxSheetCol,omitempty"`
	MaxSheetCount   *int64  `json:"MaxSheetCount,omitempty" xml:"MaxSheetCount,omitempty"`
	SheetOnePage    *bool   `json:"SheetOnePage,omitempty" xml:"SheetOnePage,omitempty"`
	ModelId         *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	TgtFilePrefix   *string `json:"TgtFilePrefix,omitempty" xml:"TgtFilePrefix,omitempty"`
	TgtFileSuffix   *string `json:"TgtFileSuffix,omitempty" xml:"TgtFileSuffix,omitempty"`
	TgtFilePages    *string `json:"TgtFilePages,omitempty" xml:"TgtFilePages,omitempty"`
	FitToPagesTall  *bool   `json:"FitToPagesTall,omitempty" xml:"FitToPagesTall,omitempty"`
	FitToPagesWide  *bool   `json:"FitToPagesWide,omitempty" xml:"FitToPagesWide,omitempty"`
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	PdfVector       *bool   `json:"PdfVector,omitempty" xml:"PdfVector,omitempty"`
	Hidecomments    *bool   `json:"Hidecomments,omitempty" xml:"Hidecomments,omitempty"`
	DisplayDpi      *int32  `json:"DisplayDpi,omitempty" xml:"DisplayDpi,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskRequest) SetProject(v string) *CreateOfficeConversionTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSrcUri(v string) *CreateOfficeConversionTaskRequest {
	s.SrcUri = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtType(v string) *CreateOfficeConversionTaskRequest {
	s.TgtType = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtUri(v string) *CreateOfficeConversionTaskRequest {
	s.TgtUri = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetNotifyTopicName(v string) *CreateOfficeConversionTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetNotifyEndpoint(v string) *CreateOfficeConversionTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSrcType(v string) *CreateOfficeConversionTaskRequest {
	s.SrcType = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetStartPage(v int64) *CreateOfficeConversionTaskRequest {
	s.StartPage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetEndPage(v int64) *CreateOfficeConversionTaskRequest {
	s.EndPage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetMaxSheetRow(v int64) *CreateOfficeConversionTaskRequest {
	s.MaxSheetRow = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetMaxSheetCol(v int64) *CreateOfficeConversionTaskRequest {
	s.MaxSheetCol = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetMaxSheetCount(v int64) *CreateOfficeConversionTaskRequest {
	s.MaxSheetCount = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSheetOnePage(v bool) *CreateOfficeConversionTaskRequest {
	s.SheetOnePage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetModelId(v string) *CreateOfficeConversionTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetPassword(v string) *CreateOfficeConversionTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtFilePrefix(v string) *CreateOfficeConversionTaskRequest {
	s.TgtFilePrefix = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtFileSuffix(v string) *CreateOfficeConversionTaskRequest {
	s.TgtFileSuffix = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtFilePages(v string) *CreateOfficeConversionTaskRequest {
	s.TgtFilePages = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetFitToPagesTall(v bool) *CreateOfficeConversionTaskRequest {
	s.FitToPagesTall = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetFitToPagesWide(v bool) *CreateOfficeConversionTaskRequest {
	s.FitToPagesWide = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetIdempotentToken(v string) *CreateOfficeConversionTaskRequest {
	s.IdempotentToken = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetPdfVector(v bool) *CreateOfficeConversionTaskRequest {
	s.PdfVector = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetHidecomments(v bool) *CreateOfficeConversionTaskRequest {
	s.Hidecomments = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetDisplayDpi(v int32) *CreateOfficeConversionTaskRequest {
	s.DisplayDpi = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetUserData(v string) *CreateOfficeConversionTaskRequest {
	s.UserData = &v
	return s
}

type CreateOfficeConversionTaskResponseBody struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Percent    *int32  `json:"Percent,omitempty" xml:"Percent,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TgtLoc     *string `json:"TgtLoc,omitempty" xml:"TgtLoc,omitempty"`
}

func (s CreateOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskResponseBody) SetStatus(v string) *CreateOfficeConversionTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetTaskId(v string) *CreateOfficeConversionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetRequestId(v string) *CreateOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetPercent(v int32) *CreateOfficeConversionTaskResponseBody {
	s.Percent = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetCreateTime(v string) *CreateOfficeConversionTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetTgtLoc(v string) *CreateOfficeConversionTaskResponseBody {
	s.TgtLoc = &v
	return s
}

type CreateOfficeConversionTaskResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *CreateOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOfficeConversionTaskResponse) SetBody(v *CreateOfficeConversionTaskResponseBody) *CreateOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type CreateSetRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
}

func (s CreateSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSetRequest) GoString() string {
	return s.String()
}

func (s *CreateSetRequest) SetProject(v string) *CreateSetRequest {
	s.Project = &v
	return s
}

func (s *CreateSetRequest) SetSetId(v string) *CreateSetRequest {
	s.SetId = &v
	return s
}

func (s *CreateSetRequest) SetSetName(v string) *CreateSetRequest {
	s.SetName = &v
	return s
}

type CreateSetResponseBody struct {
	VideoCount  *int32  `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VideoLength *int32  `json:"VideoLength,omitempty" xml:"VideoLength,omitempty"`
	SetId       *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageCount  *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	FaceCount   *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	SetName     *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s CreateSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSetResponseBody) SetVideoCount(v int32) *CreateSetResponseBody {
	s.VideoCount = &v
	return s
}

func (s *CreateSetResponseBody) SetRequestId(v string) *CreateSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSetResponseBody) SetCreateTime(v string) *CreateSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateSetResponseBody) SetVideoLength(v int32) *CreateSetResponseBody {
	s.VideoLength = &v
	return s
}

func (s *CreateSetResponseBody) SetSetId(v string) *CreateSetResponseBody {
	s.SetId = &v
	return s
}

func (s *CreateSetResponseBody) SetImageCount(v int32) *CreateSetResponseBody {
	s.ImageCount = &v
	return s
}

func (s *CreateSetResponseBody) SetFaceCount(v int32) *CreateSetResponseBody {
	s.FaceCount = &v
	return s
}

func (s *CreateSetResponseBody) SetSetName(v string) *CreateSetResponseBody {
	s.SetName = &v
	return s
}

func (s *CreateSetResponseBody) SetModifyTime(v string) *CreateSetResponseBody {
	s.ModifyTime = &v
	return s
}

type CreateSetResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSetResponse) GoString() string {
	return s.String()
}

func (s *CreateSetResponse) SetHeaders(v map[string]*string) *CreateSetResponse {
	s.Headers = v
	return s
}

func (s *CreateSetResponse) SetBody(v *CreateSetResponseBody) *CreateSetResponse {
	s.Body = v
	return s
}

type CreateVideoAbstractTaskRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	VideoUri        *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	TargetVideoUri  *string `json:"TargetVideoUri,omitempty" xml:"TargetVideoUri,omitempty"`
	TargetClipsUri  *string `json:"TargetClipsUri,omitempty" xml:"TargetClipsUri,omitempty"`
	AbstractLength  *int32  `json:"AbstractLength,omitempty" xml:"AbstractLength,omitempty"`
}

func (s CreateVideoAbstractTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoAbstractTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoAbstractTaskRequest) SetProject(v string) *CreateVideoAbstractTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateVideoAbstractTaskRequest) SetVideoUri(v string) *CreateVideoAbstractTaskRequest {
	s.VideoUri = &v
	return s
}

func (s *CreateVideoAbstractTaskRequest) SetNotifyTopicName(v string) *CreateVideoAbstractTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateVideoAbstractTaskRequest) SetNotifyEndpoint(v string) *CreateVideoAbstractTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateVideoAbstractTaskRequest) SetTargetVideoUri(v string) *CreateVideoAbstractTaskRequest {
	s.TargetVideoUri = &v
	return s
}

func (s *CreateVideoAbstractTaskRequest) SetTargetClipsUri(v string) *CreateVideoAbstractTaskRequest {
	s.TargetClipsUri = &v
	return s
}

func (s *CreateVideoAbstractTaskRequest) SetAbstractLength(v int32) *CreateVideoAbstractTaskRequest {
	s.AbstractLength = &v
	return s
}

type CreateVideoAbstractTaskResponseBody struct {
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVideoAbstractTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoAbstractTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoAbstractTaskResponseBody) SetTaskType(v string) *CreateVideoAbstractTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateVideoAbstractTaskResponseBody) SetRequestId(v string) *CreateVideoAbstractTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoAbstractTaskResponseBody) SetTaskId(v string) *CreateVideoAbstractTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateVideoAbstractTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVideoAbstractTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoAbstractTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoAbstractTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoAbstractTaskResponse) SetHeaders(v map[string]*string) *CreateVideoAbstractTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoAbstractTaskResponse) SetBody(v *CreateVideoAbstractTaskResponseBody) *CreateVideoAbstractTaskResponse {
	s.Body = v
	return s
}

type CreateVideoAnalyseTaskRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	VideoUri        *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	TgtUri          *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
}

func (s CreateVideoAnalyseTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoAnalyseTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoAnalyseTaskRequest) SetProject(v string) *CreateVideoAnalyseTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateVideoAnalyseTaskRequest) SetVideoUri(v string) *CreateVideoAnalyseTaskRequest {
	s.VideoUri = &v
	return s
}

func (s *CreateVideoAnalyseTaskRequest) SetTgtUri(v string) *CreateVideoAnalyseTaskRequest {
	s.TgtUri = &v
	return s
}

func (s *CreateVideoAnalyseTaskRequest) SetNotifyTopicName(v string) *CreateVideoAnalyseTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateVideoAnalyseTaskRequest) SetNotifyEndpoint(v string) *CreateVideoAnalyseTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

type CreateVideoAnalyseTaskResponseBody struct {
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVideoAnalyseTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoAnalyseTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoAnalyseTaskResponseBody) SetTaskType(v string) *CreateVideoAnalyseTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateVideoAnalyseTaskResponseBody) SetRequestId(v string) *CreateVideoAnalyseTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoAnalyseTaskResponseBody) SetTaskId(v string) *CreateVideoAnalyseTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateVideoAnalyseTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVideoAnalyseTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoAnalyseTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoAnalyseTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoAnalyseTaskResponse) SetHeaders(v map[string]*string) *CreateVideoAnalyseTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoAnalyseTaskResponse) SetBody(v *CreateVideoAnalyseTaskResponseBody) *CreateVideoAnalyseTaskResponse {
	s.Body = v
	return s
}

type CreateVideoCompressTaskRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	VideoUri        *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	TargetList      *string `json:"TargetList,omitempty" xml:"TargetList,omitempty"`
	CustomMessage   *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
	TargetContainer *string `json:"TargetContainer,omitempty" xml:"TargetContainer,omitempty"`
	TargetSegment   *string `json:"TargetSegment,omitempty" xml:"TargetSegment,omitempty"`
}

func (s CreateVideoCompressTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoCompressTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoCompressTaskRequest) SetProject(v string) *CreateVideoCompressTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetVideoUri(v string) *CreateVideoCompressTaskRequest {
	s.VideoUri = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetNotifyTopicName(v string) *CreateVideoCompressTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetNotifyEndpoint(v string) *CreateVideoCompressTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetTargetList(v string) *CreateVideoCompressTaskRequest {
	s.TargetList = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetCustomMessage(v string) *CreateVideoCompressTaskRequest {
	s.CustomMessage = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetTargetContainer(v string) *CreateVideoCompressTaskRequest {
	s.TargetContainer = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetTargetSegment(v string) *CreateVideoCompressTaskRequest {
	s.TargetSegment = &v
	return s
}

type CreateVideoCompressTaskResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateVideoCompressTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoCompressTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoCompressTaskResponseBody) SetTaskId(v string) *CreateVideoCompressTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVideoCompressTaskResponseBody) SetRequestId(v string) *CreateVideoCompressTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoCompressTaskResponseBody) SetTaskType(v string) *CreateVideoCompressTaskResponseBody {
	s.TaskType = &v
	return s
}

type CreateVideoCompressTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVideoCompressTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoCompressTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoCompressTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoCompressTaskResponse) SetHeaders(v map[string]*string) *CreateVideoCompressTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoCompressTaskResponse) SetBody(v *CreateVideoCompressTaskResponseBody) *CreateVideoCompressTaskResponse {
	s.Body = v
	return s
}

type CreateVideoProduceTaskRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Images          *string `json:"Images,omitempty" xml:"Images,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	TargetUri       *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
	CustomMessage   *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
	Music           *string `json:"Music,omitempty" xml:"Music,omitempty"`
	Width           *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	Height          *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateVideoProduceTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoProduceTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoProduceTaskRequest) SetProject(v string) *CreateVideoProduceTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetImages(v string) *CreateVideoProduceTaskRequest {
	s.Images = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetNotifyTopicName(v string) *CreateVideoProduceTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetNotifyEndpoint(v string) *CreateVideoProduceTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetTargetUri(v string) *CreateVideoProduceTaskRequest {
	s.TargetUri = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetCustomMessage(v string) *CreateVideoProduceTaskRequest {
	s.CustomMessage = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetMusic(v string) *CreateVideoProduceTaskRequest {
	s.Music = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetWidth(v int32) *CreateVideoProduceTaskRequest {
	s.Width = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetHeight(v int32) *CreateVideoProduceTaskRequest {
	s.Height = &v
	return s
}

func (s *CreateVideoProduceTaskRequest) SetTemplateName(v string) *CreateVideoProduceTaskRequest {
	s.TemplateName = &v
	return s
}

type CreateVideoProduceTaskResponseBody struct {
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVideoProduceTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoProduceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoProduceTaskResponseBody) SetTaskType(v string) *CreateVideoProduceTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateVideoProduceTaskResponseBody) SetRequestId(v string) *CreateVideoProduceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoProduceTaskResponseBody) SetTaskId(v string) *CreateVideoProduceTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateVideoProduceTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVideoProduceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoProduceTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoProduceTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoProduceTaskResponse) SetHeaders(v map[string]*string) *CreateVideoProduceTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoProduceTaskResponse) SetBody(v *CreateVideoProduceTaskResponseBody) *CreateVideoProduceTaskResponse {
	s.Body = v
	return s
}

type DecodeBlindWatermarkRequest struct {
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri         *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	OriginalImageUri *string `json:"OriginalImageUri,omitempty" xml:"OriginalImageUri,omitempty"`
	TargetUri        *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
	ImageQuality     *int32  `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty"`
	Model            *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s DecodeBlindWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s DecodeBlindWatermarkRequest) GoString() string {
	return s.String()
}

func (s *DecodeBlindWatermarkRequest) SetProject(v string) *DecodeBlindWatermarkRequest {
	s.Project = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetImageUri(v string) *DecodeBlindWatermarkRequest {
	s.ImageUri = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetOriginalImageUri(v string) *DecodeBlindWatermarkRequest {
	s.OriginalImageUri = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetTargetUri(v string) *DecodeBlindWatermarkRequest {
	s.TargetUri = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetImageQuality(v int32) *DecodeBlindWatermarkRequest {
	s.ImageQuality = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetModel(v string) *DecodeBlindWatermarkRequest {
	s.Model = &v
	return s
}

type DecodeBlindWatermarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TargetUri *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DecodeBlindWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecodeBlindWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *DecodeBlindWatermarkResponseBody) SetRequestId(v string) *DecodeBlindWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecodeBlindWatermarkResponseBody) SetTargetUri(v string) *DecodeBlindWatermarkResponseBody {
	s.TargetUri = &v
	return s
}

func (s *DecodeBlindWatermarkResponseBody) SetContent(v string) *DecodeBlindWatermarkResponseBody {
	s.Content = &v
	return s
}

type DecodeBlindWatermarkResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DecodeBlindWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DecodeBlindWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s DecodeBlindWatermarkResponse) GoString() string {
	return s.String()
}

func (s *DecodeBlindWatermarkResponse) SetHeaders(v map[string]*string) *DecodeBlindWatermarkResponse {
	s.Headers = v
	return s
}

func (s *DecodeBlindWatermarkResponse) SetBody(v *DecodeBlindWatermarkResponseBody) *DecodeBlindWatermarkResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetProject(v string) *DeleteImageRequest {
	s.Project = &v
	return s
}

func (s *DeleteImageRequest) SetSetId(v string) *DeleteImageRequest {
	s.SetId = &v
	return s
}

func (s *DeleteImageRequest) SetImageUri(v string) *DeleteImageRequest {
	s.ImageUri = &v
	return s
}

type DeleteImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageUri  *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetSetId(v string) *DeleteImageResponseBody {
	s.SetId = &v
	return s
}

func (s *DeleteImageResponseBody) SetImageUri(v string) *DeleteImageResponseBody {
	s.ImageUri = &v
	return s
}

type DeleteImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DeleteImageJobRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteImageJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageJobRequest) SetProject(v string) *DeleteImageJobRequest {
	s.Project = &v
	return s
}

func (s *DeleteImageJobRequest) SetJobType(v string) *DeleteImageJobRequest {
	s.JobType = &v
	return s
}

func (s *DeleteImageJobRequest) SetJobId(v string) *DeleteImageJobRequest {
	s.JobId = &v
	return s
}

type DeleteImageJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageJobResponseBody) SetRequestId(v string) *DeleteImageJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageJobResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageJobResponse) SetHeaders(v map[string]*string) *DeleteImageJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageJobResponse) SetBody(v *DeleteImageJobResponseBody) *DeleteImageJobResponse {
	s.Body = v
	return s
}

type DeleteOfficeConversionTaskRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteOfficeConversionTaskRequest) SetProject(v string) *DeleteOfficeConversionTaskRequest {
	s.Project = &v
	return s
}

func (s *DeleteOfficeConversionTaskRequest) SetTaskId(v string) *DeleteOfficeConversionTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteOfficeConversionTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOfficeConversionTaskResponseBody) SetRequestId(v string) *DeleteOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOfficeConversionTaskResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *DeleteOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteOfficeConversionTaskResponse) SetBody(v *DeleteOfficeConversionTaskResponseBody) *DeleteOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProject(v string) *DeleteProjectRequest {
	s.Project = &v
	return s
}

type DeleteProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteSetRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s DeleteSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteSetRequest) SetProject(v string) *DeleteSetRequest {
	s.Project = &v
	return s
}

func (s *DeleteSetRequest) SetSetId(v string) *DeleteSetRequest {
	s.SetId = &v
	return s
}

type DeleteSetResponseBody struct {
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSetResponseBody) SetSetId(v string) *DeleteSetResponseBody {
	s.SetId = &v
	return s
}

func (s *DeleteSetResponseBody) SetRequestId(v string) *DeleteSetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSetResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteSetResponse) SetHeaders(v map[string]*string) *DeleteSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteSetResponse) SetBody(v *DeleteSetResponseBody) *DeleteSetResponse {
	s.Body = v
	return s
}

type DeleteVideoRequest struct {
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri  *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	Resources *bool   `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s DeleteVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoRequest) SetProject(v string) *DeleteVideoRequest {
	s.Project = &v
	return s
}

func (s *DeleteVideoRequest) SetSetId(v string) *DeleteVideoRequest {
	s.SetId = &v
	return s
}

func (s *DeleteVideoRequest) SetVideoUri(v string) *DeleteVideoRequest {
	s.VideoUri = &v
	return s
}

func (s *DeleteVideoRequest) SetResources(v bool) *DeleteVideoRequest {
	s.Resources = &v
	return s
}

type DeleteVideoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoUri  *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s DeleteVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoResponseBody) SetRequestId(v string) *DeleteVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVideoResponseBody) SetVideoUri(v string) *DeleteVideoResponseBody {
	s.VideoUri = &v
	return s
}

func (s *DeleteVideoResponseBody) SetSetId(v string) *DeleteVideoResponseBody {
	s.SetId = &v
	return s
}

type DeleteVideoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoResponse) SetHeaders(v map[string]*string) *DeleteVideoResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideoResponse) SetBody(v *DeleteVideoResponseBody) *DeleteVideoResponse {
	s.Body = v
	return s
}

type DeleteVideoTaskRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoTaskRequest) SetProject(v string) *DeleteVideoTaskRequest {
	s.Project = &v
	return s
}

func (s *DeleteVideoTaskRequest) SetTaskType(v string) *DeleteVideoTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DeleteVideoTaskRequest) SetTaskId(v string) *DeleteVideoTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteVideoTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoTaskResponseBody) SetRequestId(v string) *DeleteVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVideoTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoTaskResponse) SetHeaders(v map[string]*string) *DeleteVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideoTaskResponse) SetBody(v *DeleteVideoTaskResponseBody) *DeleteVideoTaskResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectTypes []*string `json:"ProjectTypes,omitempty" xml:"ProjectTypes,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetProjectTypes(v []*string) *DescribeRegionsResponseBodyRegionsRegion {
	s.ProjectTypes = v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DetectImageBodiesRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s DetectImageBodiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesRequest) SetProject(v string) *DetectImageBodiesRequest {
	s.Project = &v
	return s
}

func (s *DetectImageBodiesRequest) SetImageUri(v string) *DetectImageBodiesRequest {
	s.ImageUri = &v
	return s
}

type DetectImageBodiesResponseBody struct {
	ImageUri  *string                                `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Bodies    []*DetectImageBodiesResponseBodyBodies `json:"Bodies,omitempty" xml:"Bodies,omitempty" type:"Repeated"`
}

func (s DetectImageBodiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponseBody) SetImageUri(v string) *DetectImageBodiesResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageBodiesResponseBody) SetRequestId(v string) *DetectImageBodiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageBodiesResponseBody) SetBodies(v []*DetectImageBodiesResponseBodyBodies) *DetectImageBodiesResponseBody {
	s.Bodies = v
	return s
}

type DetectImageBodiesResponseBodyBodies struct {
	BodyConfidence *float32                                         `json:"BodyConfidence,omitempty" xml:"BodyConfidence,omitempty"`
	BodyBoundary   *DetectImageBodiesResponseBodyBodiesBodyBoundary `json:"BodyBoundary,omitempty" xml:"BodyBoundary,omitempty" type:"Struct"`
}

func (s DetectImageBodiesResponseBodyBodies) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponseBodyBodies) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponseBodyBodies) SetBodyConfidence(v float32) *DetectImageBodiesResponseBodyBodies {
	s.BodyConfidence = &v
	return s
}

func (s *DetectImageBodiesResponseBodyBodies) SetBodyBoundary(v *DetectImageBodiesResponseBodyBodiesBodyBoundary) *DetectImageBodiesResponseBodyBodies {
	s.BodyBoundary = v
	return s
}

type DetectImageBodiesResponseBodyBodiesBodyBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DetectImageBodiesResponseBodyBodiesBodyBoundary) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponseBodyBodiesBodyBoundary) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponseBodyBodiesBodyBoundary) SetLeft(v int32) *DetectImageBodiesResponseBodyBodiesBodyBoundary {
	s.Left = &v
	return s
}

func (s *DetectImageBodiesResponseBodyBodiesBodyBoundary) SetTop(v int32) *DetectImageBodiesResponseBodyBodiesBodyBoundary {
	s.Top = &v
	return s
}

func (s *DetectImageBodiesResponseBodyBodiesBodyBoundary) SetWidth(v int32) *DetectImageBodiesResponseBodyBodiesBodyBoundary {
	s.Width = &v
	return s
}

func (s *DetectImageBodiesResponseBodyBodiesBodyBoundary) SetHeight(v int32) *DetectImageBodiesResponseBodyBodiesBodyBoundary {
	s.Height = &v
	return s
}

type DetectImageBodiesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectImageBodiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageBodiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponse) SetHeaders(v map[string]*string) *DetectImageBodiesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageBodiesResponse) SetBody(v *DetectImageBodiesResponseBody) *DetectImageBodiesResponse {
	s.Body = v
	return s
}

type DetectImageFacesRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s DetectImageFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageFacesRequest) SetProject(v string) *DetectImageFacesRequest {
	s.Project = &v
	return s
}

func (s *DetectImageFacesRequest) SetImageUri(v string) *DetectImageFacesRequest {
	s.ImageUri = &v
	return s
}

type DetectImageFacesResponseBody struct {
	ImageUri  *string                              `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Faces     []*DetectImageFacesResponseBodyFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
}

func (s DetectImageFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBody) SetImageUri(v string) *DetectImageFacesResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageFacesResponseBody) SetRequestId(v string) *DetectImageFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageFacesResponseBody) SetFaces(v []*DetectImageFacesResponseBodyFaces) *DetectImageFacesResponseBody {
	s.Faces = v
	return s
}

type DetectImageFacesResponseBodyFaces struct {
	EmotionConfidence    *float32                                         `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	Attractive           *float32                                         `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	AttractiveConfidence *float32                                         `json:"AttractiveConfidence,omitempty" xml:"AttractiveConfidence,omitempty"`
	Gender               *string                                          `json:"Gender,omitempty" xml:"Gender,omitempty"`
	AgeConfidence        *float32                                         `json:"AgeConfidence,omitempty" xml:"AgeConfidence,omitempty"`
	GenderConfidence     *float32                                         `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	FaceId               *string                                          `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceQuality          *float32                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Emotion              *string                                          `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	Age                  *int32                                           `json:"Age,omitempty" xml:"Age,omitempty"`
	FaceConfidence       *float32                                         `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	FaceAttributes       *DetectImageFacesResponseBodyFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	EmotionDetails       *DetectImageFacesResponseBodyFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
}

func (s DetectImageFacesResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFaces) SetEmotionConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetAttractive(v float32) *DetectImageFacesResponseBodyFaces {
	s.Attractive = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetAttractiveConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.AttractiveConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetGender(v string) *DetectImageFacesResponseBodyFaces {
	s.Gender = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetAgeConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.AgeConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetGenderConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.GenderConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetFaceId(v string) *DetectImageFacesResponseBodyFaces {
	s.FaceId = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetFaceQuality(v float32) *DetectImageFacesResponseBodyFaces {
	s.FaceQuality = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetEmotion(v string) *DetectImageFacesResponseBodyFaces {
	s.Emotion = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetAge(v int32) *DetectImageFacesResponseBodyFaces {
	s.Age = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetFaceConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.FaceConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetFaceAttributes(v *DetectImageFacesResponseBodyFacesFaceAttributes) *DetectImageFacesResponseBodyFaces {
	s.FaceAttributes = v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetEmotionDetails(v *DetectImageFacesResponseBodyFacesEmotionDetails) *DetectImageFacesResponseBodyFaces {
	s.EmotionDetails = v
	return s
}

type DetectImageFacesResponseBodyFacesFaceAttributes struct {
	GlassesConfidence *float32                                                     `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	Glasses           *string                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Mask              *string                                                      `json:"Mask,omitempty" xml:"Mask,omitempty"`
	BeardConfidence   *float32                                                     `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	MaskConfidence    *float32                                                     `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	Beard             *string                                                      `json:"Beard,omitempty" xml:"Beard,omitempty"`
	FaceBoundary      *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	HeadPose          *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
}

func (s DetectImageFacesResponseBodyFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetGlassesConfidence(v float32) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetGlasses(v string) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetMask(v string) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetBeardConfidence(v float32) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetMaskConfidence(v float32) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetBeard(v string) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.Beard = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetFaceBoundary(v *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetHeadPose(v *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.HeadPose = v
	return s
}

type DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) SetLeft(v int32) *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) SetTop(v int32) *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) SetWidth(v int32) *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) SetHeight(v int32) *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

type DetectImageFacesResponseBodyFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) SetPitch(v float32) *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) SetRoll(v float32) *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) SetYaw(v float32) *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type DetectImageFacesResponseBodyFacesEmotionDetails struct {
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
}

func (s DetectImageFacesResponseBodyFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetHAPPY(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetCALM(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetSURPRISED(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetDISGUSTED(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetANGRY(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetSAD(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetSCARED(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.SCARED = &v
	return s
}

type DetectImageFacesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectImageFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponse) SetHeaders(v map[string]*string) *DetectImageFacesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageFacesResponse) SetBody(v *DetectImageFacesResponseBody) *DetectImageFacesResponse {
	s.Body = v
	return s
}

type DetectImageLogosRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s DetectImageLogosRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLogosRequest) GoString() string {
	return s.String()
}

func (s *DetectImageLogosRequest) SetProject(v string) *DetectImageLogosRequest {
	s.Project = &v
	return s
}

func (s *DetectImageLogosRequest) SetImageUri(v string) *DetectImageLogosRequest {
	s.ImageUri = &v
	return s
}

type DetectImageLogosResponseBody struct {
	ImageUri  *string                              `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Logos     []*DetectImageLogosResponseBodyLogos `json:"Logos,omitempty" xml:"Logos,omitempty" type:"Repeated"`
}

func (s DetectImageLogosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLogosResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageLogosResponseBody) SetImageUri(v string) *DetectImageLogosResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageLogosResponseBody) SetRequestId(v string) *DetectImageLogosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageLogosResponseBody) SetLogos(v []*DetectImageLogosResponseBodyLogos) *DetectImageLogosResponseBody {
	s.Logos = v
	return s
}

type DetectImageLogosResponseBodyLogos struct {
	LogoConfidence *float32                                       `json:"LogoConfidence,omitempty" xml:"LogoConfidence,omitempty"`
	LogoName       *string                                        `json:"LogoName,omitempty" xml:"LogoName,omitempty"`
	LogoBoundary   *DetectImageLogosResponseBodyLogosLogoBoundary `json:"LogoBoundary,omitempty" xml:"LogoBoundary,omitempty" type:"Struct"`
}

func (s DetectImageLogosResponseBodyLogos) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLogosResponseBodyLogos) GoString() string {
	return s.String()
}

func (s *DetectImageLogosResponseBodyLogos) SetLogoConfidence(v float32) *DetectImageLogosResponseBodyLogos {
	s.LogoConfidence = &v
	return s
}

func (s *DetectImageLogosResponseBodyLogos) SetLogoName(v string) *DetectImageLogosResponseBodyLogos {
	s.LogoName = &v
	return s
}

func (s *DetectImageLogosResponseBodyLogos) SetLogoBoundary(v *DetectImageLogosResponseBodyLogosLogoBoundary) *DetectImageLogosResponseBodyLogos {
	s.LogoBoundary = v
	return s
}

type DetectImageLogosResponseBodyLogosLogoBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DetectImageLogosResponseBodyLogosLogoBoundary) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLogosResponseBodyLogosLogoBoundary) GoString() string {
	return s.String()
}

func (s *DetectImageLogosResponseBodyLogosLogoBoundary) SetLeft(v int32) *DetectImageLogosResponseBodyLogosLogoBoundary {
	s.Left = &v
	return s
}

func (s *DetectImageLogosResponseBodyLogosLogoBoundary) SetTop(v int32) *DetectImageLogosResponseBodyLogosLogoBoundary {
	s.Top = &v
	return s
}

func (s *DetectImageLogosResponseBodyLogosLogoBoundary) SetWidth(v int32) *DetectImageLogosResponseBodyLogosLogoBoundary {
	s.Width = &v
	return s
}

func (s *DetectImageLogosResponseBodyLogosLogoBoundary) SetHeight(v int32) *DetectImageLogosResponseBodyLogosLogoBoundary {
	s.Height = &v
	return s
}

type DetectImageLogosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectImageLogosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageLogosResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLogosResponse) GoString() string {
	return s.String()
}

func (s *DetectImageLogosResponse) SetHeaders(v map[string]*string) *DetectImageLogosResponse {
	s.Headers = v
	return s
}

func (s *DetectImageLogosResponse) SetBody(v *DetectImageLogosResponseBody) *DetectImageLogosResponse {
	s.Body = v
	return s
}

type DetectImageQRCodesRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s DetectImageQRCodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesRequest) SetProject(v string) *DetectImageQRCodesRequest {
	s.Project = &v
	return s
}

func (s *DetectImageQRCodesRequest) SetImageUri(v string) *DetectImageQRCodesRequest {
	s.ImageUri = &v
	return s
}

type DetectImageQRCodesResponseBody struct {
	ImageUri  *string                                  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QRCodes   []*DetectImageQRCodesResponseBodyQRCodes `json:"QRCodes,omitempty" xml:"QRCodes,omitempty" type:"Repeated"`
}

func (s DetectImageQRCodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesResponseBody) SetImageUri(v string) *DetectImageQRCodesResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageQRCodesResponseBody) SetRequestId(v string) *DetectImageQRCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageQRCodesResponseBody) SetQRCodes(v []*DetectImageQRCodesResponseBodyQRCodes) *DetectImageQRCodesResponseBody {
	s.QRCodes = v
	return s
}

type DetectImageQRCodesResponseBodyQRCodes struct {
	Content        *string                                              `json:"Content,omitempty" xml:"Content,omitempty"`
	QRCodeBoundary *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary `json:"QRCodeBoundary,omitempty" xml:"QRCodeBoundary,omitempty" type:"Struct"`
}

func (s DetectImageQRCodesResponseBodyQRCodes) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesResponseBodyQRCodes) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesResponseBodyQRCodes) SetContent(v string) *DetectImageQRCodesResponseBodyQRCodes {
	s.Content = &v
	return s
}

func (s *DetectImageQRCodesResponseBodyQRCodes) SetQRCodeBoundary(v *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) *DetectImageQRCodesResponseBodyQRCodes {
	s.QRCodeBoundary = v
	return s
}

type DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) SetLeft(v int32) *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary {
	s.Left = &v
	return s
}

func (s *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) SetTop(v int32) *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary {
	s.Top = &v
	return s
}

func (s *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) SetWidth(v int32) *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary {
	s.Width = &v
	return s
}

func (s *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) SetHeight(v int32) *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary {
	s.Height = &v
	return s
}

type DetectImageQRCodesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectImageQRCodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageQRCodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesResponse) SetHeaders(v map[string]*string) *DetectImageQRCodesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageQRCodesResponse) SetBody(v *DetectImageQRCodesResponseBody) *DetectImageQRCodesResponse {
	s.Body = v
	return s
}

type DetectImageTagsRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s DetectImageTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageTagsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageTagsRequest) SetProject(v string) *DetectImageTagsRequest {
	s.Project = &v
	return s
}

func (s *DetectImageTagsRequest) SetImageUri(v string) *DetectImageTagsRequest {
	s.ImageUri = &v
	return s
}

type DetectImageTagsResponseBody struct {
	ImageUri  *string                            `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*DetectImageTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DetectImageTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageTagsResponseBody) SetImageUri(v string) *DetectImageTagsResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageTagsResponseBody) SetRequestId(v string) *DetectImageTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageTagsResponseBody) SetTags(v []*DetectImageTagsResponseBodyTags) *DetectImageTagsResponseBody {
	s.Tags = v
	return s
}

type DetectImageTagsResponseBodyTags struct {
	ParentTagEnName *string  `json:"ParentTagEnName,omitempty" xml:"ParentTagEnName,omitempty"`
	TagName         *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagConfidence   *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagEnName       *string  `json:"TagEnName,omitempty" xml:"TagEnName,omitempty"`
	TagLevel        *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	ParentTagName   *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
}

func (s DetectImageTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DetectImageTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DetectImageTagsResponseBodyTags) SetParentTagEnName(v string) *DetectImageTagsResponseBodyTags {
	s.ParentTagEnName = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetTagName(v string) *DetectImageTagsResponseBodyTags {
	s.TagName = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetTagConfidence(v float32) *DetectImageTagsResponseBodyTags {
	s.TagConfidence = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetTagEnName(v string) *DetectImageTagsResponseBodyTags {
	s.TagEnName = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetTagLevel(v int32) *DetectImageTagsResponseBodyTags {
	s.TagLevel = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetParentTagName(v string) *DetectImageTagsResponseBodyTags {
	s.ParentTagName = &v
	return s
}

type DetectImageTagsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectImageTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageTagsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageTagsResponse) SetHeaders(v map[string]*string) *DetectImageTagsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageTagsResponse) SetBody(v *DetectImageTagsResponseBody) *DetectImageTagsResponse {
	s.Body = v
	return s
}

type DetectQRCodesRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcUris *string `json:"SrcUris,omitempty" xml:"SrcUris,omitempty"`
}

func (s DetectQRCodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesRequest) GoString() string {
	return s.String()
}

func (s *DetectQRCodesRequest) SetProject(v string) *DetectQRCodesRequest {
	s.Project = &v
	return s
}

func (s *DetectQRCodesRequest) SetSrcUris(v string) *DetectQRCodesRequest {
	s.SrcUris = &v
	return s
}

type DetectQRCodesResponseBody struct {
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessDetails []*DetectQRCodesResponseBodySuccessDetails `json:"SuccessDetails,omitempty" xml:"SuccessDetails,omitempty" type:"Repeated"`
	FailDetails    []*DetectQRCodesResponseBodyFailDetails    `json:"FailDetails,omitempty" xml:"FailDetails,omitempty" type:"Repeated"`
}

func (s DetectQRCodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBody) SetRequestId(v string) *DetectQRCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectQRCodesResponseBody) SetSuccessDetails(v []*DetectQRCodesResponseBodySuccessDetails) *DetectQRCodesResponseBody {
	s.SuccessDetails = v
	return s
}

func (s *DetectQRCodesResponseBody) SetFailDetails(v []*DetectQRCodesResponseBodyFailDetails) *DetectQRCodesResponseBody {
	s.FailDetails = v
	return s
}

type DetectQRCodesResponseBodySuccessDetails struct {
	SrcUri  *string                                           `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	QRCodes []*DetectQRCodesResponseBodySuccessDetailsQRCodes `json:"QRCodes,omitempty" xml:"QRCodes,omitempty" type:"Repeated"`
}

func (s DetectQRCodesResponseBodySuccessDetails) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBodySuccessDetails) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBodySuccessDetails) SetSrcUri(v string) *DetectQRCodesResponseBodySuccessDetails {
	s.SrcUri = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetails) SetQRCodes(v []*DetectQRCodesResponseBodySuccessDetailsQRCodes) *DetectQRCodesResponseBodySuccessDetails {
	s.QRCodes = v
	return s
}

type DetectQRCodesResponseBodySuccessDetailsQRCodes struct {
	Content          *string                                                         `json:"Content,omitempty" xml:"Content,omitempty"`
	QRCodesRectangle *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle `json:"QRCodesRectangle,omitempty" xml:"QRCodesRectangle,omitempty" type:"Struct"`
}

func (s DetectQRCodesResponseBodySuccessDetailsQRCodes) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBodySuccessDetailsQRCodes) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodes) SetContent(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodes {
	s.Content = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodes) SetQRCodesRectangle(v *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) *DetectQRCodesResponseBodySuccessDetailsQRCodes {
	s.QRCodesRectangle = v
	return s
}

type DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle struct {
	Left   *string `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *string `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *string `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) SetLeft(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle {
	s.Left = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) SetTop(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle {
	s.Top = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) SetWidth(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle {
	s.Width = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) SetHeight(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle {
	s.Height = &v
	return s
}

type DetectQRCodesResponseBodyFailDetails struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	SrcUri       *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s DetectQRCodesResponseBodyFailDetails) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBodyFailDetails) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBodyFailDetails) SetErrorMessage(v string) *DetectQRCodesResponseBodyFailDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DetectQRCodesResponseBodyFailDetails) SetSrcUri(v string) *DetectQRCodesResponseBodyFailDetails {
	s.SrcUri = &v
	return s
}

func (s *DetectQRCodesResponseBodyFailDetails) SetErrorCode(v string) *DetectQRCodesResponseBodyFailDetails {
	s.ErrorCode = &v
	return s
}

type DetectQRCodesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectQRCodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectQRCodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponse) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponse) SetHeaders(v map[string]*string) *DetectQRCodesResponse {
	s.Headers = v
	return s
}

func (s *DetectQRCodesResponse) SetBody(v *DetectQRCodesResponseBody) *DetectQRCodesResponse {
	s.Body = v
	return s
}

type EncodeBlindWatermarkRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri        *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	WatermarkUri    *string `json:"WatermarkUri,omitempty" xml:"WatermarkUri,omitempty"`
	TargetUri       *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
	ImageQuality    *string `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	TargetImageType *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s EncodeBlindWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s EncodeBlindWatermarkRequest) GoString() string {
	return s.String()
}

func (s *EncodeBlindWatermarkRequest) SetProject(v string) *EncodeBlindWatermarkRequest {
	s.Project = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetImageUri(v string) *EncodeBlindWatermarkRequest {
	s.ImageUri = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetWatermarkUri(v string) *EncodeBlindWatermarkRequest {
	s.WatermarkUri = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetTargetUri(v string) *EncodeBlindWatermarkRequest {
	s.TargetUri = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetImageQuality(v string) *EncodeBlindWatermarkRequest {
	s.ImageQuality = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetContent(v string) *EncodeBlindWatermarkRequest {
	s.Content = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetTargetImageType(v string) *EncodeBlindWatermarkRequest {
	s.TargetImageType = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetModel(v string) *EncodeBlindWatermarkRequest {
	s.Model = &v
	return s
}

type EncodeBlindWatermarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TargetUri *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s EncodeBlindWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EncodeBlindWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *EncodeBlindWatermarkResponseBody) SetRequestId(v string) *EncodeBlindWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *EncodeBlindWatermarkResponseBody) SetTargetUri(v string) *EncodeBlindWatermarkResponseBody {
	s.TargetUri = &v
	return s
}

func (s *EncodeBlindWatermarkResponseBody) SetContent(v string) *EncodeBlindWatermarkResponseBody {
	s.Content = &v
	return s
}

type EncodeBlindWatermarkResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EncodeBlindWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EncodeBlindWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s EncodeBlindWatermarkResponse) GoString() string {
	return s.String()
}

func (s *EncodeBlindWatermarkResponse) SetHeaders(v map[string]*string) *EncodeBlindWatermarkResponse {
	s.Headers = v
	return s
}

func (s *EncodeBlindWatermarkResponse) SetBody(v *EncodeBlindWatermarkResponseBody) *EncodeBlindWatermarkResponse {
	s.Body = v
	return s
}

type FindImagesRequest struct {
	Project                  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId                    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageSizeRange           *string `json:"ImageSizeRange,omitempty" xml:"ImageSizeRange,omitempty"`
	ImageTimeRange           *string `json:"ImageTimeRange,omitempty" xml:"ImageTimeRange,omitempty"`
	CreateTimeRange          *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	ModifyTimeRange          *string `json:"ModifyTimeRange,omitempty" xml:"ModifyTimeRange,omitempty"`
	SourceType               *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUriPrefix          *string `json:"SourceUriPrefix,omitempty" xml:"SourceUriPrefix,omitempty"`
	RemarksAPrefix           *string `json:"RemarksAPrefix,omitempty" xml:"RemarksAPrefix,omitempty"`
	RemarksBPrefix           *string `json:"RemarksBPrefix,omitempty" xml:"RemarksBPrefix,omitempty"`
	TagNames                 *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
	OCRContentsMatch         *string `json:"OCRContentsMatch,omitempty" xml:"OCRContentsMatch,omitempty"`
	AgeRange                 *string `json:"AgeRange,omitempty" xml:"AgeRange,omitempty"`
	Gender                   *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Emotion                  *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	OrderBy                  *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Order                    *string `json:"Order,omitempty" xml:"Order,omitempty"`
	Marker                   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	LocationBoundary         *string `json:"LocationBoundary,omitempty" xml:"LocationBoundary,omitempty"`
	RemarksCPrefix           *string `json:"RemarksCPrefix,omitempty" xml:"RemarksCPrefix,omitempty"`
	RemarksDPrefix           *string `json:"RemarksDPrefix,omitempty" xml:"RemarksDPrefix,omitempty"`
	ExternalId               *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	GroupId                  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Limit                    *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	FacesModifyTimeRange     *string `json:"FacesModifyTimeRange,omitempty" xml:"FacesModifyTimeRange,omitempty"`
	TagsModifyTimeRange      *string `json:"TagsModifyTimeRange,omitempty" xml:"TagsModifyTimeRange,omitempty"`
	AddressLineContentsMatch *string `json:"AddressLineContentsMatch,omitempty" xml:"AddressLineContentsMatch,omitempty"`
	RemarksArrayAIn          *string `json:"RemarksArrayAIn,omitempty" xml:"RemarksArrayAIn,omitempty"`
	RemarksArrayBIn          *string `json:"RemarksArrayBIn,omitempty" xml:"RemarksArrayBIn,omitempty"`
}

func (s FindImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindImagesRequest) GoString() string {
	return s.String()
}

func (s *FindImagesRequest) SetProject(v string) *FindImagesRequest {
	s.Project = &v
	return s
}

func (s *FindImagesRequest) SetSetId(v string) *FindImagesRequest {
	s.SetId = &v
	return s
}

func (s *FindImagesRequest) SetImageSizeRange(v string) *FindImagesRequest {
	s.ImageSizeRange = &v
	return s
}

func (s *FindImagesRequest) SetImageTimeRange(v string) *FindImagesRequest {
	s.ImageTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetCreateTimeRange(v string) *FindImagesRequest {
	s.CreateTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetModifyTimeRange(v string) *FindImagesRequest {
	s.ModifyTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetSourceType(v string) *FindImagesRequest {
	s.SourceType = &v
	return s
}

func (s *FindImagesRequest) SetSourceUriPrefix(v string) *FindImagesRequest {
	s.SourceUriPrefix = &v
	return s
}

func (s *FindImagesRequest) SetRemarksAPrefix(v string) *FindImagesRequest {
	s.RemarksAPrefix = &v
	return s
}

func (s *FindImagesRequest) SetRemarksBPrefix(v string) *FindImagesRequest {
	s.RemarksBPrefix = &v
	return s
}

func (s *FindImagesRequest) SetTagNames(v string) *FindImagesRequest {
	s.TagNames = &v
	return s
}

func (s *FindImagesRequest) SetOCRContentsMatch(v string) *FindImagesRequest {
	s.OCRContentsMatch = &v
	return s
}

func (s *FindImagesRequest) SetAgeRange(v string) *FindImagesRequest {
	s.AgeRange = &v
	return s
}

func (s *FindImagesRequest) SetGender(v string) *FindImagesRequest {
	s.Gender = &v
	return s
}

func (s *FindImagesRequest) SetEmotion(v string) *FindImagesRequest {
	s.Emotion = &v
	return s
}

func (s *FindImagesRequest) SetOrderBy(v string) *FindImagesRequest {
	s.OrderBy = &v
	return s
}

func (s *FindImagesRequest) SetOrder(v string) *FindImagesRequest {
	s.Order = &v
	return s
}

func (s *FindImagesRequest) SetMarker(v string) *FindImagesRequest {
	s.Marker = &v
	return s
}

func (s *FindImagesRequest) SetLocationBoundary(v string) *FindImagesRequest {
	s.LocationBoundary = &v
	return s
}

func (s *FindImagesRequest) SetRemarksCPrefix(v string) *FindImagesRequest {
	s.RemarksCPrefix = &v
	return s
}

func (s *FindImagesRequest) SetRemarksDPrefix(v string) *FindImagesRequest {
	s.RemarksDPrefix = &v
	return s
}

func (s *FindImagesRequest) SetExternalId(v string) *FindImagesRequest {
	s.ExternalId = &v
	return s
}

func (s *FindImagesRequest) SetGroupId(v string) *FindImagesRequest {
	s.GroupId = &v
	return s
}

func (s *FindImagesRequest) SetLimit(v int32) *FindImagesRequest {
	s.Limit = &v
	return s
}

func (s *FindImagesRequest) SetFacesModifyTimeRange(v string) *FindImagesRequest {
	s.FacesModifyTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetTagsModifyTimeRange(v string) *FindImagesRequest {
	s.TagsModifyTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetAddressLineContentsMatch(v string) *FindImagesRequest {
	s.AddressLineContentsMatch = &v
	return s
}

func (s *FindImagesRequest) SetRemarksArrayAIn(v string) *FindImagesRequest {
	s.RemarksArrayAIn = &v
	return s
}

func (s *FindImagesRequest) SetRemarksArrayBIn(v string) *FindImagesRequest {
	s.RemarksArrayBIn = &v
	return s
}

type FindImagesResponseBody struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextMarker *string                         `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	SetId      *string                         `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Images     []*FindImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s FindImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBody) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBody) SetRequestId(v string) *FindImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindImagesResponseBody) SetNextMarker(v string) *FindImagesResponseBody {
	s.NextMarker = &v
	return s
}

func (s *FindImagesResponseBody) SetSetId(v string) *FindImagesResponseBody {
	s.SetId = &v
	return s
}

func (s *FindImagesResponseBody) SetImages(v []*FindImagesResponseBodyImages) *FindImagesResponseBody {
	s.Images = v
	return s
}

type FindImagesResponseBodyImages struct {
	CroppingSuggestionStatus     *string                                           `json:"CroppingSuggestionStatus,omitempty" xml:"CroppingSuggestionStatus,omitempty"`
	ImageQualityModifyTime       *string                                           `json:"ImageQualityModifyTime,omitempty" xml:"ImageQualityModifyTime,omitempty"`
	TagsFailReason               *string                                           `json:"TagsFailReason,omitempty" xml:"TagsFailReason,omitempty"`
	RemarksC                     *string                                           `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	CreateTime                   *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SourceType                   *string                                           `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	FacesFailReason              *string                                           `json:"FacesFailReason,omitempty" xml:"FacesFailReason,omitempty"`
	FacesModifyTime              *string                                           `json:"FacesModifyTime,omitempty" xml:"FacesModifyTime,omitempty"`
	ImageTime                    *string                                           `json:"ImageTime,omitempty" xml:"ImageTime,omitempty"`
	OCRModifyTime                *string                                           `json:"OCRModifyTime,omitempty" xml:"OCRModifyTime,omitempty"`
	AddressModifyTime            *string                                           `json:"AddressModifyTime,omitempty" xml:"AddressModifyTime,omitempty"`
	ImageQualityFailReason       *string                                           `json:"ImageQualityFailReason,omitempty" xml:"ImageQualityFailReason,omitempty"`
	FacesStatus                  *string                                           `json:"FacesStatus,omitempty" xml:"FacesStatus,omitempty"`
	RemarksArrayA                *string                                           `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	ImageHeight                  *int32                                            `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ExternalId                   *string                                           `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	SourceUri                    *string                                           `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	FileSize                     *int32                                            `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	ModifyTime                   *string                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SourcePosition               *string                                           `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	ImageQualityStatus           *string                                           `json:"ImageQualityStatus,omitempty" xml:"ImageQualityStatus,omitempty"`
	OCRFailReason                *string                                           `json:"OCRFailReason,omitempty" xml:"OCRFailReason,omitempty"`
	AddressFailReason            *string                                           `json:"AddressFailReason,omitempty" xml:"AddressFailReason,omitempty"`
	CroppingSuggestionModifyTime *string                                           `json:"CroppingSuggestionModifyTime,omitempty" xml:"CroppingSuggestionModifyTime,omitempty"`
	ImageFormat                  *string                                           `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	ImageWidth                   *int32                                            `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	RemarksArrayB                *string                                           `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	Orientation                  *string                                           `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	RemarksD                     *string                                           `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	TagsStatus                   *string                                           `json:"TagsStatus,omitempty" xml:"TagsStatus,omitempty"`
	CroppingSuggestionFailReason *string                                           `json:"CroppingSuggestionFailReason,omitempty" xml:"CroppingSuggestionFailReason,omitempty"`
	RemarksA                     *string                                           `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	ImageUri                     *string                                           `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	TagsModifyTime               *string                                           `json:"TagsModifyTime,omitempty" xml:"TagsModifyTime,omitempty"`
	OCRStatus                    *string                                           `json:"OCRStatus,omitempty" xml:"OCRStatus,omitempty"`
	AddressStatus                *string                                           `json:"AddressStatus,omitempty" xml:"AddressStatus,omitempty"`
	Exif                         *string                                           `json:"Exif,omitempty" xml:"Exif,omitempty"`
	Location                     *string                                           `json:"Location,omitempty" xml:"Location,omitempty"`
	RemarksB                     *string                                           `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	CroppingSuggestion           []*FindImagesResponseBodyImagesCroppingSuggestion `json:"CroppingSuggestion,omitempty" xml:"CroppingSuggestion,omitempty" type:"Repeated"`
	Faces                        []*FindImagesResponseBodyImagesFaces              `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	Tags                         []*FindImagesResponseBodyImagesTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	OCR                          []*FindImagesResponseBodyImagesOCR                `json:"OCR,omitempty" xml:"OCR,omitempty" type:"Repeated"`
	ImageQuality                 *FindImagesResponseBodyImagesImageQuality         `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty" type:"Struct"`
	Address                      *FindImagesResponseBodyImagesAddress              `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
}

func (s FindImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImages) SetCroppingSuggestionStatus(v string) *FindImagesResponseBodyImages {
	s.CroppingSuggestionStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageQualityModifyTime(v string) *FindImagesResponseBodyImages {
	s.ImageQualityModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetTagsFailReason(v string) *FindImagesResponseBodyImages {
	s.TagsFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksC(v string) *FindImagesResponseBodyImages {
	s.RemarksC = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetCreateTime(v string) *FindImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetSourceType(v string) *FindImagesResponseBodyImages {
	s.SourceType = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetFacesFailReason(v string) *FindImagesResponseBodyImages {
	s.FacesFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetFacesModifyTime(v string) *FindImagesResponseBodyImages {
	s.FacesModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageTime(v string) *FindImagesResponseBodyImages {
	s.ImageTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetOCRModifyTime(v string) *FindImagesResponseBodyImages {
	s.OCRModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetAddressModifyTime(v string) *FindImagesResponseBodyImages {
	s.AddressModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageQualityFailReason(v string) *FindImagesResponseBodyImages {
	s.ImageQualityFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetFacesStatus(v string) *FindImagesResponseBodyImages {
	s.FacesStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksArrayA(v string) *FindImagesResponseBodyImages {
	s.RemarksArrayA = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageHeight(v int32) *FindImagesResponseBodyImages {
	s.ImageHeight = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetExternalId(v string) *FindImagesResponseBodyImages {
	s.ExternalId = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetSourceUri(v string) *FindImagesResponseBodyImages {
	s.SourceUri = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetFileSize(v int32) *FindImagesResponseBodyImages {
	s.FileSize = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetModifyTime(v string) *FindImagesResponseBodyImages {
	s.ModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetSourcePosition(v string) *FindImagesResponseBodyImages {
	s.SourcePosition = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageQualityStatus(v string) *FindImagesResponseBodyImages {
	s.ImageQualityStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetOCRFailReason(v string) *FindImagesResponseBodyImages {
	s.OCRFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetAddressFailReason(v string) *FindImagesResponseBodyImages {
	s.AddressFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetCroppingSuggestionModifyTime(v string) *FindImagesResponseBodyImages {
	s.CroppingSuggestionModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageFormat(v string) *FindImagesResponseBodyImages {
	s.ImageFormat = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageWidth(v int32) *FindImagesResponseBodyImages {
	s.ImageWidth = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksArrayB(v string) *FindImagesResponseBodyImages {
	s.RemarksArrayB = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetOrientation(v string) *FindImagesResponseBodyImages {
	s.Orientation = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksD(v string) *FindImagesResponseBodyImages {
	s.RemarksD = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetTagsStatus(v string) *FindImagesResponseBodyImages {
	s.TagsStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetCroppingSuggestionFailReason(v string) *FindImagesResponseBodyImages {
	s.CroppingSuggestionFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksA(v string) *FindImagesResponseBodyImages {
	s.RemarksA = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageUri(v string) *FindImagesResponseBodyImages {
	s.ImageUri = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetTagsModifyTime(v string) *FindImagesResponseBodyImages {
	s.TagsModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetOCRStatus(v string) *FindImagesResponseBodyImages {
	s.OCRStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetAddressStatus(v string) *FindImagesResponseBodyImages {
	s.AddressStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetExif(v string) *FindImagesResponseBodyImages {
	s.Exif = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetLocation(v string) *FindImagesResponseBodyImages {
	s.Location = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksB(v string) *FindImagesResponseBodyImages {
	s.RemarksB = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetCroppingSuggestion(v []*FindImagesResponseBodyImagesCroppingSuggestion) *FindImagesResponseBodyImages {
	s.CroppingSuggestion = v
	return s
}

func (s *FindImagesResponseBodyImages) SetFaces(v []*FindImagesResponseBodyImagesFaces) *FindImagesResponseBodyImages {
	s.Faces = v
	return s
}

func (s *FindImagesResponseBodyImages) SetTags(v []*FindImagesResponseBodyImagesTags) *FindImagesResponseBodyImages {
	s.Tags = v
	return s
}

func (s *FindImagesResponseBodyImages) SetOCR(v []*FindImagesResponseBodyImagesOCR) *FindImagesResponseBodyImages {
	s.OCR = v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageQuality(v *FindImagesResponseBodyImagesImageQuality) *FindImagesResponseBodyImages {
	s.ImageQuality = v
	return s
}

func (s *FindImagesResponseBodyImages) SetAddress(v *FindImagesResponseBodyImagesAddress) *FindImagesResponseBodyImages {
	s.Address = v
	return s
}

type FindImagesResponseBodyImagesCroppingSuggestion struct {
	Score            *float32                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	AspectRatio      *string                                                         `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	CroppingBoundary *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary `json:"CroppingBoundary,omitempty" xml:"CroppingBoundary,omitempty" type:"Struct"`
}

func (s FindImagesResponseBodyImagesCroppingSuggestion) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesCroppingSuggestion) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesCroppingSuggestion) SetScore(v float32) *FindImagesResponseBodyImagesCroppingSuggestion {
	s.Score = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestion) SetAspectRatio(v string) *FindImagesResponseBodyImagesCroppingSuggestion {
	s.AspectRatio = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestion) SetCroppingBoundary(v *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) *FindImagesResponseBodyImagesCroppingSuggestion {
	s.CroppingBoundary = v
	return s
}

type FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetLeft(v int32) *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Left = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetTop(v int32) *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Top = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetWidth(v int32) *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Width = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetHeight(v int32) *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Height = &v
	return s
}

type FindImagesResponseBodyImagesFaces struct {
	EmotionConfidence *float32                                         `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	Attractive        *float32                                         `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	GroupId           *string                                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Gender            *string                                          `json:"Gender,omitempty" xml:"Gender,omitempty"`
	FaceId            *string                                          `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	GenderConfidence  *float32                                         `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	FaceQuality       *float32                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Emotion           *string                                          `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	Age               *int32                                           `json:"Age,omitempty" xml:"Age,omitempty"`
	FaceConfidence    *float32                                         `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	EmotionDetails    *FindImagesResponseBodyImagesFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
	FaceAttributes    *FindImagesResponseBodyImagesFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
}

func (s FindImagesResponseBodyImagesFaces) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFaces) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFaces) SetEmotionConfidence(v float32) *FindImagesResponseBodyImagesFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetAttractive(v float32) *FindImagesResponseBodyImagesFaces {
	s.Attractive = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetGroupId(v string) *FindImagesResponseBodyImagesFaces {
	s.GroupId = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetGender(v string) *FindImagesResponseBodyImagesFaces {
	s.Gender = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetFaceId(v string) *FindImagesResponseBodyImagesFaces {
	s.FaceId = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetGenderConfidence(v float32) *FindImagesResponseBodyImagesFaces {
	s.GenderConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetFaceQuality(v float32) *FindImagesResponseBodyImagesFaces {
	s.FaceQuality = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetEmotion(v string) *FindImagesResponseBodyImagesFaces {
	s.Emotion = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetAge(v int32) *FindImagesResponseBodyImagesFaces {
	s.Age = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetFaceConfidence(v float32) *FindImagesResponseBodyImagesFaces {
	s.FaceConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetEmotionDetails(v *FindImagesResponseBodyImagesFacesEmotionDetails) *FindImagesResponseBodyImagesFaces {
	s.EmotionDetails = v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetFaceAttributes(v *FindImagesResponseBodyImagesFacesFaceAttributes) *FindImagesResponseBodyImagesFaces {
	s.FaceAttributes = v
	return s
}

type FindImagesResponseBodyImagesFacesEmotionDetails struct {
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
}

func (s FindImagesResponseBodyImagesFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetHAPPY(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetSURPRISED(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetCALM(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetDISGUSTED(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetANGRY(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetSAD(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetSCARED(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.SCARED = &v
	return s
}

type FindImagesResponseBodyImagesFacesFaceAttributes struct {
	GlassesConfidence *float32                                                     `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	Glasses           *string                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Mask              *string                                                      `json:"Mask,omitempty" xml:"Mask,omitempty"`
	BeardConfidence   *float32                                                     `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	MaskConfidence    *float32                                                     `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	Beard             *string                                                      `json:"Beard,omitempty" xml:"Beard,omitempty"`
	FaceBoundary      *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	HeadPose          *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
}

func (s FindImagesResponseBodyImagesFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetGlassesConfidence(v float32) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetGlasses(v string) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetMask(v string) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetBeardConfidence(v float32) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetMaskConfidence(v float32) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetBeard(v string) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.Beard = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetFaceBoundary(v *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetHeadPose(v *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.HeadPose = v
	return s
}

type FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetLeft(v int32) *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetTop(v int32) *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetWidth(v int32) *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetHeight(v int32) *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

type FindImagesResponseBodyImagesFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetPitch(v float32) *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetRoll(v float32) *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetYaw(v float32) *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type FindImagesResponseBodyImagesTags struct {
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s FindImagesResponseBodyImagesTags) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesTags) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesTags) SetTagLevel(v int32) *FindImagesResponseBodyImagesTags {
	s.TagLevel = &v
	return s
}

func (s *FindImagesResponseBodyImagesTags) SetParentTagName(v string) *FindImagesResponseBodyImagesTags {
	s.ParentTagName = &v
	return s
}

func (s *FindImagesResponseBodyImagesTags) SetTagConfidence(v float32) *FindImagesResponseBodyImagesTags {
	s.TagConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesTags) SetTagName(v string) *FindImagesResponseBodyImagesTags {
	s.TagName = &v
	return s
}

type FindImagesResponseBodyImagesOCR struct {
	OCRConfidence *float32                                    `json:"OCRConfidence,omitempty" xml:"OCRConfidence,omitempty"`
	OCRContents   *string                                     `json:"OCRContents,omitempty" xml:"OCRContents,omitempty"`
	OCRBoundary   *FindImagesResponseBodyImagesOCROCRBoundary `json:"OCRBoundary,omitempty" xml:"OCRBoundary,omitempty" type:"Struct"`
}

func (s FindImagesResponseBodyImagesOCR) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesOCR) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesOCR) SetOCRConfidence(v float32) *FindImagesResponseBodyImagesOCR {
	s.OCRConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCR) SetOCRContents(v string) *FindImagesResponseBodyImagesOCR {
	s.OCRContents = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCR) SetOCRBoundary(v *FindImagesResponseBodyImagesOCROCRBoundary) *FindImagesResponseBodyImagesOCR {
	s.OCRBoundary = v
	return s
}

type FindImagesResponseBodyImagesOCROCRBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s FindImagesResponseBodyImagesOCROCRBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesOCROCRBoundary) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesOCROCRBoundary) SetLeft(v int32) *FindImagesResponseBodyImagesOCROCRBoundary {
	s.Left = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCROCRBoundary) SetTop(v int32) *FindImagesResponseBodyImagesOCROCRBoundary {
	s.Top = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCROCRBoundary) SetWidth(v int32) *FindImagesResponseBodyImagesOCROCRBoundary {
	s.Width = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCROCRBoundary) SetHeight(v int32) *FindImagesResponseBodyImagesOCROCRBoundary {
	s.Height = &v
	return s
}

type FindImagesResponseBodyImagesImageQuality struct {
	OverallScore     *float32 `json:"OverallScore,omitempty" xml:"OverallScore,omitempty"`
	Color            *float32 `json:"Color,omitempty" xml:"Color,omitempty"`
	ColorScore       *float32 `json:"ColorScore,omitempty" xml:"ColorScore,omitempty"`
	ContrastScore    *float32 `json:"ContrastScore,omitempty" xml:"ContrastScore,omitempty"`
	Contrast         *float32 `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	ExposureScore    *float32 `json:"ExposureScore,omitempty" xml:"ExposureScore,omitempty"`
	ClarityScore     *float32 `json:"ClarityScore,omitempty" xml:"ClarityScore,omitempty"`
	Clarity          *float32 `json:"Clarity,omitempty" xml:"Clarity,omitempty"`
	Exposure         *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
	CompositionScore *float32 `json:"CompositionScore,omitempty" xml:"CompositionScore,omitempty"`
}

func (s FindImagesResponseBodyImagesImageQuality) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesImageQuality) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesImageQuality) SetOverallScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.OverallScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetColor(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.Color = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetColorScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.ColorScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetContrastScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.ContrastScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetContrast(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.Contrast = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetExposureScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.ExposureScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetClarityScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.ClarityScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetClarity(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.Clarity = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetExposure(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.Exposure = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetCompositionScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.CompositionScore = &v
	return s
}

type FindImagesResponseBodyImagesAddress struct {
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s FindImagesResponseBodyImagesAddress) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesAddress) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesAddress) SetTownship(v string) *FindImagesResponseBodyImagesAddress {
	s.Township = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetDistrict(v string) *FindImagesResponseBodyImagesAddress {
	s.District = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetAddressLine(v string) *FindImagesResponseBodyImagesAddress {
	s.AddressLine = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetCountry(v string) *FindImagesResponseBodyImagesAddress {
	s.Country = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetCity(v string) *FindImagesResponseBodyImagesAddress {
	s.City = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetProvince(v string) *FindImagesResponseBodyImagesAddress {
	s.Province = &v
	return s
}

type FindImagesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponse) GoString() string {
	return s.String()
}

func (s *FindImagesResponse) SetHeaders(v map[string]*string) *FindImagesResponse {
	s.Headers = v
	return s
}

func (s *FindImagesResponse) SetBody(v *FindImagesResponseBody) *FindImagesResponse {
	s.Body = v
	return s
}

type FindSimilarFacesRequest struct {
	Project        *string  `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId          *string  `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageUri       *string  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	FaceId         *string  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	Limit          *int32   `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MinSimilarity  *float32 `json:"MinSimilarity,omitempty" xml:"MinSimilarity,omitempty"`
	ResponseFormat *string  `json:"ResponseFormat,omitempty" xml:"ResponseFormat,omitempty"`
}

func (s FindSimilarFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesRequest) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesRequest) SetProject(v string) *FindSimilarFacesRequest {
	s.Project = &v
	return s
}

func (s *FindSimilarFacesRequest) SetSetId(v string) *FindSimilarFacesRequest {
	s.SetId = &v
	return s
}

func (s *FindSimilarFacesRequest) SetImageUri(v string) *FindSimilarFacesRequest {
	s.ImageUri = &v
	return s
}

func (s *FindSimilarFacesRequest) SetFaceId(v string) *FindSimilarFacesRequest {
	s.FaceId = &v
	return s
}

func (s *FindSimilarFacesRequest) SetLimit(v int32) *FindSimilarFacesRequest {
	s.Limit = &v
	return s
}

func (s *FindSimilarFacesRequest) SetMinSimilarity(v float32) *FindSimilarFacesRequest {
	s.MinSimilarity = &v
	return s
}

func (s *FindSimilarFacesRequest) SetResponseFormat(v string) *FindSimilarFacesRequest {
	s.ResponseFormat = &v
	return s
}

type FindSimilarFacesResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Faces     []*FindSimilarFacesResponseBodyFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
}

func (s FindSimilarFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBody) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBody) SetRequestId(v string) *FindSimilarFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindSimilarFacesResponseBody) SetFaces(v []*FindSimilarFacesResponseBodyFaces) *FindSimilarFacesResponseBody {
	s.Faces = v
	return s
}

type FindSimilarFacesResponseBodyFaces struct {
	ExternalId     *string                                          `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Similarity     *float32                                         `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
	FaceId         *string                                          `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	ImageUri       *string                                          `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	SimilarFaces   []*FindSimilarFacesResponseBodyFacesSimilarFaces `json:"SimilarFaces,omitempty" xml:"SimilarFaces,omitempty" type:"Repeated"`
	FaceAttributes *FindSimilarFacesResponseBodyFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
}

func (s FindSimilarFacesResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFaces) SetExternalId(v string) *FindSimilarFacesResponseBodyFaces {
	s.ExternalId = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetSimilarity(v float32) *FindSimilarFacesResponseBodyFaces {
	s.Similarity = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetFaceId(v string) *FindSimilarFacesResponseBodyFaces {
	s.FaceId = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetImageUri(v string) *FindSimilarFacesResponseBodyFaces {
	s.ImageUri = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetSimilarFaces(v []*FindSimilarFacesResponseBodyFacesSimilarFaces) *FindSimilarFacesResponseBodyFaces {
	s.SimilarFaces = v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetFaceAttributes(v *FindSimilarFacesResponseBodyFacesFaceAttributes) *FindSimilarFacesResponseBodyFaces {
	s.FaceAttributes = v
	return s
}

type FindSimilarFacesResponseBodyFacesSimilarFaces struct {
	ExternalId     *string                                                      `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Similarity     *float32                                                     `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
	FaceId         *string                                                      `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	ImageUri       *string                                                      `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	FaceAttributes *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
}

func (s FindSimilarFacesResponseBodyFacesSimilarFaces) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesSimilarFaces) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetExternalId(v string) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.ExternalId = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetSimilarity(v float32) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.Similarity = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetFaceId(v string) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.FaceId = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetImageUri(v string) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.ImageUri = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetFaceAttributes(v *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.FaceAttributes = v
	return s
}

type FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes struct {
	FaceBoundary *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
}

func (s FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes) SetFaceBoundary(v *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

type FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) SetLeft(v int32) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) SetTop(v int32) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) SetWidth(v int32) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) SetHeight(v int32) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

type FindSimilarFacesResponseBodyFacesFaceAttributes struct {
	FaceBoundary *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
}

func (s FindSimilarFacesResponseBodyFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributes) SetFaceBoundary(v *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) *FindSimilarFacesResponseBodyFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

type FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) SetLeft(v int32) *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) SetTop(v int32) *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) SetWidth(v int32) *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) SetHeight(v int32) *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

type FindSimilarFacesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindSimilarFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindSimilarFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponse) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponse) SetHeaders(v map[string]*string) *FindSimilarFacesResponse {
	s.Headers = v
	return s
}

func (s *FindSimilarFacesResponse) SetBody(v *FindSimilarFacesResponseBody) *FindSimilarFacesResponse {
	s.Body = v
	return s
}

type GetContentKeyRequest struct {
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	DRMServerId *string `json:"DRMServerId,omitempty" xml:"DRMServerId,omitempty"`
	KeyIds      *string `json:"KeyIds,omitempty" xml:"KeyIds,omitempty"`
}

func (s GetContentKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContentKeyRequest) GoString() string {
	return s.String()
}

func (s *GetContentKeyRequest) SetProject(v string) *GetContentKeyRequest {
	s.Project = &v
	return s
}

func (s *GetContentKeyRequest) SetVersionId(v string) *GetContentKeyRequest {
	s.VersionId = &v
	return s
}

func (s *GetContentKeyRequest) SetDRMServerId(v string) *GetContentKeyRequest {
	s.DRMServerId = &v
	return s
}

func (s *GetContentKeyRequest) SetKeyIds(v string) *GetContentKeyRequest {
	s.KeyIds = &v
	return s
}

type GetContentKeyResponseBody struct {
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KeyInfos  *string `json:"KeyInfos,omitempty" xml:"KeyInfos,omitempty"`
}

func (s GetContentKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContentKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetContentKeyResponseBody) SetVersionId(v string) *GetContentKeyResponseBody {
	s.VersionId = &v
	return s
}

func (s *GetContentKeyResponseBody) SetRequestId(v string) *GetContentKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContentKeyResponseBody) SetKeyInfos(v string) *GetContentKeyResponseBody {
	s.KeyInfos = &v
	return s
}

type GetContentKeyResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetContentKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetContentKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContentKeyResponse) GoString() string {
	return s.String()
}

func (s *GetContentKeyResponse) SetHeaders(v map[string]*string) *GetContentKeyResponse {
	s.Headers = v
	return s
}

func (s *GetContentKeyResponse) SetBody(v *GetContentKeyResponseBody) *GetContentKeyResponse {
	s.Body = v
	return s
}

type GetDRMLicenseRequest struct {
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
	DRMType    *string `json:"DRMType,omitempty" xml:"DRMType,omitempty"`
	DRMLicense *string `json:"DRMLicense,omitempty" xml:"DRMLicense,omitempty"`
}

func (s GetDRMLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDRMLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseRequest) SetProject(v string) *GetDRMLicenseRequest {
	s.Project = &v
	return s
}

func (s *GetDRMLicenseRequest) SetDRMType(v string) *GetDRMLicenseRequest {
	s.DRMType = &v
	return s
}

func (s *GetDRMLicenseRequest) SetDRMLicense(v string) *GetDRMLicenseRequest {
	s.DRMLicense = &v
	return s
}

type GetDRMLicenseResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceInfo *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	DRMData    *string `json:"DRMData,omitempty" xml:"DRMData,omitempty"`
}

func (s GetDRMLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDRMLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseResponseBody) SetRequestId(v string) *GetDRMLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetDeviceInfo(v string) *GetDRMLicenseResponseBody {
	s.DeviceInfo = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetDRMData(v string) *GetDRMLicenseResponseBody {
	s.DRMData = &v
	return s
}

type GetDRMLicenseResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDRMLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDRMLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDRMLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseResponse) SetHeaders(v map[string]*string) *GetDRMLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetDRMLicenseResponse) SetBody(v *GetDRMLicenseResponseBody) *GetDRMLicenseResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetProject(v string) *GetImageRequest {
	s.Project = &v
	return s
}

func (s *GetImageRequest) SetSetId(v string) *GetImageRequest {
	s.SetId = &v
	return s
}

func (s *GetImageRequest) SetImageUri(v string) *GetImageRequest {
	s.ImageUri = &v
	return s
}

type GetImageResponseBody struct {
	ImageQuality                 *GetImageResponseBodyImageQuality         `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty" type:"Struct"`
	ModifyTime                   *string                                   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Address                      *GetImageResponseBodyAddress              `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	SourceType                   *string                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri                    *string                                   `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	FacesFailReason              *string                                   `json:"FacesFailReason,omitempty" xml:"FacesFailReason,omitempty"`
	CroppingSuggestionStatus     *string                                   `json:"CroppingSuggestionStatus,omitempty" xml:"CroppingSuggestionStatus,omitempty"`
	CroppingSuggestionFailReason *string                                   `json:"CroppingSuggestionFailReason,omitempty" xml:"CroppingSuggestionFailReason,omitempty"`
	AddressFailReason            *string                                   `json:"AddressFailReason,omitempty" xml:"AddressFailReason,omitempty"`
	RemarksA                     *string                                   `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	AddressModifyTime            *string                                   `json:"AddressModifyTime,omitempty" xml:"AddressModifyTime,omitempty"`
	RemarksB                     *string                                   `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	ImageFormat                  *string                                   `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	TagsFailReason               *string                                   `json:"TagsFailReason,omitempty" xml:"TagsFailReason,omitempty"`
	RemarksArrayB                *string                                   `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	FacesModifyTime              *string                                   `json:"FacesModifyTime,omitempty" xml:"FacesModifyTime,omitempty"`
	Exif                         *string                                   `json:"Exif,omitempty" xml:"Exif,omitempty"`
	RemarksC                     *string                                   `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD                     *string                                   `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	ImageWidth                   *int32                                    `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	RemarksArrayA                *string                                   `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	SourcePosition               *string                                   `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	Tags                         []*GetImageResponseBodyTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Faces                        []*GetImageResponseBodyFaces              `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	AddressStatus                *string                                   `json:"AddressStatus,omitempty" xml:"AddressStatus,omitempty"`
	FacesStatus                  *string                                   `json:"FacesStatus,omitempty" xml:"FacesStatus,omitempty"`
	ImageQualityModifyTime       *string                                   `json:"ImageQualityModifyTime,omitempty" xml:"ImageQualityModifyTime,omitempty"`
	CroppingSuggestion           []*GetImageResponseBodyCroppingSuggestion `json:"CroppingSuggestion,omitempty" xml:"CroppingSuggestion,omitempty" type:"Repeated"`
	RequestId                    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime                   *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId                   *string                                   `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	TagsModifyTime               *string                                   `json:"TagsModifyTime,omitempty" xml:"TagsModifyTime,omitempty"`
	ImageQualityFailReason       *string                                   `json:"ImageQualityFailReason,omitempty" xml:"ImageQualityFailReason,omitempty"`
	Orientation                  *string                                   `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	ImageUri                     *string                                   `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	OCRStatus                    *string                                   `json:"OCRStatus,omitempty" xml:"OCRStatus,omitempty"`
	OCRModifyTime                *string                                   `json:"OCRModifyTime,omitempty" xml:"OCRModifyTime,omitempty"`
	ImageTime                    *string                                   `json:"ImageTime,omitempty" xml:"ImageTime,omitempty"`
	CroppingSuggestionModifyTime *string                                   `json:"CroppingSuggestionModifyTime,omitempty" xml:"CroppingSuggestionModifyTime,omitempty"`
	ImageHeight                  *int32                                    `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageQualityStatus           *string                                   `json:"ImageQualityStatus,omitempty" xml:"ImageQualityStatus,omitempty"`
	TagsStatus                   *string                                   `json:"TagsStatus,omitempty" xml:"TagsStatus,omitempty"`
	OCRFailReason                *string                                   `json:"OCRFailReason,omitempty" xml:"OCRFailReason,omitempty"`
	SetId                        *string                                   `json:"SetId,omitempty" xml:"SetId,omitempty"`
	FileSize                     *int32                                    `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Location                     *string                                   `json:"Location,omitempty" xml:"Location,omitempty"`
	OCR                          []*GetImageResponseBodyOCR                `json:"OCR,omitempty" xml:"OCR,omitempty" type:"Repeated"`
}

func (s GetImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) SetImageQuality(v *GetImageResponseBodyImageQuality) *GetImageResponseBody {
	s.ImageQuality = v
	return s
}

func (s *GetImageResponseBody) SetModifyTime(v string) *GetImageResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetAddress(v *GetImageResponseBodyAddress) *GetImageResponseBody {
	s.Address = v
	return s
}

func (s *GetImageResponseBody) SetSourceType(v string) *GetImageResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetImageResponseBody) SetSourceUri(v string) *GetImageResponseBody {
	s.SourceUri = &v
	return s
}

func (s *GetImageResponseBody) SetFacesFailReason(v string) *GetImageResponseBody {
	s.FacesFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetCroppingSuggestionStatus(v string) *GetImageResponseBody {
	s.CroppingSuggestionStatus = &v
	return s
}

func (s *GetImageResponseBody) SetCroppingSuggestionFailReason(v string) *GetImageResponseBody {
	s.CroppingSuggestionFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetAddressFailReason(v string) *GetImageResponseBody {
	s.AddressFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksA(v string) *GetImageResponseBody {
	s.RemarksA = &v
	return s
}

func (s *GetImageResponseBody) SetAddressModifyTime(v string) *GetImageResponseBody {
	s.AddressModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksB(v string) *GetImageResponseBody {
	s.RemarksB = &v
	return s
}

func (s *GetImageResponseBody) SetImageFormat(v string) *GetImageResponseBody {
	s.ImageFormat = &v
	return s
}

func (s *GetImageResponseBody) SetTagsFailReason(v string) *GetImageResponseBody {
	s.TagsFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksArrayB(v string) *GetImageResponseBody {
	s.RemarksArrayB = &v
	return s
}

func (s *GetImageResponseBody) SetFacesModifyTime(v string) *GetImageResponseBody {
	s.FacesModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetExif(v string) *GetImageResponseBody {
	s.Exif = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksC(v string) *GetImageResponseBody {
	s.RemarksC = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksD(v string) *GetImageResponseBody {
	s.RemarksD = &v
	return s
}

func (s *GetImageResponseBody) SetImageWidth(v int32) *GetImageResponseBody {
	s.ImageWidth = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksArrayA(v string) *GetImageResponseBody {
	s.RemarksArrayA = &v
	return s
}

func (s *GetImageResponseBody) SetSourcePosition(v string) *GetImageResponseBody {
	s.SourcePosition = &v
	return s
}

func (s *GetImageResponseBody) SetTags(v []*GetImageResponseBodyTags) *GetImageResponseBody {
	s.Tags = v
	return s
}

func (s *GetImageResponseBody) SetFaces(v []*GetImageResponseBodyFaces) *GetImageResponseBody {
	s.Faces = v
	return s
}

func (s *GetImageResponseBody) SetAddressStatus(v string) *GetImageResponseBody {
	s.AddressStatus = &v
	return s
}

func (s *GetImageResponseBody) SetFacesStatus(v string) *GetImageResponseBody {
	s.FacesStatus = &v
	return s
}

func (s *GetImageResponseBody) SetImageQualityModifyTime(v string) *GetImageResponseBody {
	s.ImageQualityModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetCroppingSuggestion(v []*GetImageResponseBodyCroppingSuggestion) *GetImageResponseBody {
	s.CroppingSuggestion = v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetCreateTime(v string) *GetImageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetImageResponseBody) SetExternalId(v string) *GetImageResponseBody {
	s.ExternalId = &v
	return s
}

func (s *GetImageResponseBody) SetTagsModifyTime(v string) *GetImageResponseBody {
	s.TagsModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetImageQualityFailReason(v string) *GetImageResponseBody {
	s.ImageQualityFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetOrientation(v string) *GetImageResponseBody {
	s.Orientation = &v
	return s
}

func (s *GetImageResponseBody) SetImageUri(v string) *GetImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageResponseBody) SetOCRStatus(v string) *GetImageResponseBody {
	s.OCRStatus = &v
	return s
}

func (s *GetImageResponseBody) SetOCRModifyTime(v string) *GetImageResponseBody {
	s.OCRModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetImageTime(v string) *GetImageResponseBody {
	s.ImageTime = &v
	return s
}

func (s *GetImageResponseBody) SetCroppingSuggestionModifyTime(v string) *GetImageResponseBody {
	s.CroppingSuggestionModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetImageHeight(v int32) *GetImageResponseBody {
	s.ImageHeight = &v
	return s
}

func (s *GetImageResponseBody) SetImageQualityStatus(v string) *GetImageResponseBody {
	s.ImageQualityStatus = &v
	return s
}

func (s *GetImageResponseBody) SetTagsStatus(v string) *GetImageResponseBody {
	s.TagsStatus = &v
	return s
}

func (s *GetImageResponseBody) SetOCRFailReason(v string) *GetImageResponseBody {
	s.OCRFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetSetId(v string) *GetImageResponseBody {
	s.SetId = &v
	return s
}

func (s *GetImageResponseBody) SetFileSize(v int32) *GetImageResponseBody {
	s.FileSize = &v
	return s
}

func (s *GetImageResponseBody) SetLocation(v string) *GetImageResponseBody {
	s.Location = &v
	return s
}

func (s *GetImageResponseBody) SetOCR(v []*GetImageResponseBodyOCR) *GetImageResponseBody {
	s.OCR = v
	return s
}

type GetImageResponseBodyImageQuality struct {
	OverallScore     *float32 `json:"OverallScore,omitempty" xml:"OverallScore,omitempty"`
	Color            *float32 `json:"Color,omitempty" xml:"Color,omitempty"`
	ColorScore       *float32 `json:"ColorScore,omitempty" xml:"ColorScore,omitempty"`
	ContrastScore    *float32 `json:"ContrastScore,omitempty" xml:"ContrastScore,omitempty"`
	Contrast         *float32 `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	ExposureScore    *float32 `json:"ExposureScore,omitempty" xml:"ExposureScore,omitempty"`
	ClarityScore     *float32 `json:"ClarityScore,omitempty" xml:"ClarityScore,omitempty"`
	Clarity          *float32 `json:"Clarity,omitempty" xml:"Clarity,omitempty"`
	Exposure         *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
	CompositionScore *float32 `json:"CompositionScore,omitempty" xml:"CompositionScore,omitempty"`
}

func (s GetImageResponseBodyImageQuality) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImageQuality) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageQuality) SetOverallScore(v float32) *GetImageResponseBodyImageQuality {
	s.OverallScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetColor(v float32) *GetImageResponseBodyImageQuality {
	s.Color = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetColorScore(v float32) *GetImageResponseBodyImageQuality {
	s.ColorScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetContrastScore(v float32) *GetImageResponseBodyImageQuality {
	s.ContrastScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetContrast(v float32) *GetImageResponseBodyImageQuality {
	s.Contrast = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetExposureScore(v float32) *GetImageResponseBodyImageQuality {
	s.ExposureScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetClarityScore(v float32) *GetImageResponseBodyImageQuality {
	s.ClarityScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetClarity(v float32) *GetImageResponseBodyImageQuality {
	s.Clarity = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetExposure(v float32) *GetImageResponseBodyImageQuality {
	s.Exposure = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetCompositionScore(v float32) *GetImageResponseBodyImageQuality {
	s.CompositionScore = &v
	return s
}

type GetImageResponseBodyAddress struct {
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s GetImageResponseBodyAddress) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyAddress) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyAddress) SetTownship(v string) *GetImageResponseBodyAddress {
	s.Township = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetDistrict(v string) *GetImageResponseBodyAddress {
	s.District = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetAddressLine(v string) *GetImageResponseBodyAddress {
	s.AddressLine = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetCountry(v string) *GetImageResponseBodyAddress {
	s.Country = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetCity(v string) *GetImageResponseBodyAddress {
	s.City = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetProvince(v string) *GetImageResponseBodyAddress {
	s.Province = &v
	return s
}

type GetImageResponseBodyTags struct {
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
}

func (s GetImageResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyTags) SetTagName(v string) *GetImageResponseBodyTags {
	s.TagName = &v
	return s
}

func (s *GetImageResponseBodyTags) SetTagConfidence(v float32) *GetImageResponseBodyTags {
	s.TagConfidence = &v
	return s
}

func (s *GetImageResponseBodyTags) SetTagLevel(v int32) *GetImageResponseBodyTags {
	s.TagLevel = &v
	return s
}

func (s *GetImageResponseBodyTags) SetParentTagName(v string) *GetImageResponseBodyTags {
	s.ParentTagName = &v
	return s
}

type GetImageResponseBodyFaces struct {
	Gender            *string                                  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GenderConfidence  *float32                                 `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	FaceId            *string                                  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceAttributes    *GetImageResponseBodyFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceQuality       *float32                                 `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Emotion           *string                                  `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	Age               *string                                  `json:"Age,omitempty" xml:"Age,omitempty"`
	FaceConfidence    *float32                                 `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	EmotionConfidence *float32                                 `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	Attractive        *float32                                 `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	GroupId           *string                                  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	EmotionDetails    *GetImageResponseBodyFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
}

func (s GetImageResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFaces) SetGender(v string) *GetImageResponseBodyFaces {
	s.Gender = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetGenderConfidence(v float32) *GetImageResponseBodyFaces {
	s.GenderConfidence = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetFaceId(v string) *GetImageResponseBodyFaces {
	s.FaceId = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetFaceAttributes(v *GetImageResponseBodyFacesFaceAttributes) *GetImageResponseBodyFaces {
	s.FaceAttributes = v
	return s
}

func (s *GetImageResponseBodyFaces) SetFaceQuality(v float32) *GetImageResponseBodyFaces {
	s.FaceQuality = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetEmotion(v string) *GetImageResponseBodyFaces {
	s.Emotion = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetAge(v string) *GetImageResponseBodyFaces {
	s.Age = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetFaceConfidence(v float32) *GetImageResponseBodyFaces {
	s.FaceConfidence = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetEmotionConfidence(v float32) *GetImageResponseBodyFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetAttractive(v float32) *GetImageResponseBodyFaces {
	s.Attractive = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetGroupId(v string) *GetImageResponseBodyFaces {
	s.GroupId = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetEmotionDetails(v *GetImageResponseBodyFacesEmotionDetails) *GetImageResponseBodyFaces {
	s.EmotionDetails = v
	return s
}

type GetImageResponseBodyFacesFaceAttributes struct {
	GlassesConfidence *float32                                             `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	Glasses           *string                                              `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Mask              *string                                              `json:"Mask,omitempty" xml:"Mask,omitempty"`
	BeardConfidence   *float32                                             `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	MaskConfidence    *float32                                             `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	FaceBoundary      *GetImageResponseBodyFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	HeadPose          *GetImageResponseBodyFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
	Beard             *string                                              `json:"Beard,omitempty" xml:"Beard,omitempty"`
}

func (s GetImageResponseBodyFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetGlassesConfidence(v float32) *GetImageResponseBodyFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetGlasses(v string) *GetImageResponseBodyFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetMask(v string) *GetImageResponseBodyFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetBeardConfidence(v float32) *GetImageResponseBodyFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetMaskConfidence(v float32) *GetImageResponseBodyFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetFaceBoundary(v *GetImageResponseBodyFacesFaceAttributesFaceBoundary) *GetImageResponseBodyFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetHeadPose(v *GetImageResponseBodyFacesFaceAttributesHeadPose) *GetImageResponseBodyFacesFaceAttributes {
	s.HeadPose = v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetBeard(v string) *GetImageResponseBodyFacesFaceAttributes {
	s.Beard = &v
	return s
}

type GetImageResponseBodyFacesFaceAttributesFaceBoundary struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s GetImageResponseBodyFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFacesFaceAttributesFaceBoundary) SetTop(v int32) *GetImageResponseBodyFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesFaceBoundary) SetWidth(v int32) *GetImageResponseBodyFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesFaceBoundary) SetHeight(v int32) *GetImageResponseBodyFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesFaceBoundary) SetLeft(v int32) *GetImageResponseBodyFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

type GetImageResponseBodyFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s GetImageResponseBodyFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFacesFaceAttributesHeadPose) SetPitch(v float32) *GetImageResponseBodyFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesHeadPose) SetRoll(v float32) *GetImageResponseBodyFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesHeadPose) SetYaw(v float32) *GetImageResponseBodyFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type GetImageResponseBodyFacesEmotionDetails struct {
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
}

func (s GetImageResponseBodyFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetHAPPY(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetCALM(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetSURPRISED(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetDISGUSTED(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetANGRY(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetSAD(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetSCARED(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.SCARED = &v
	return s
}

type GetImageResponseBodyCroppingSuggestion struct {
	Score            *float32                                                `json:"Score,omitempty" xml:"Score,omitempty"`
	CroppingBoundary *GetImageResponseBodyCroppingSuggestionCroppingBoundary `json:"CroppingBoundary,omitempty" xml:"CroppingBoundary,omitempty" type:"Struct"`
	AspectRatio      *string                                                 `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
}

func (s GetImageResponseBodyCroppingSuggestion) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyCroppingSuggestion) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyCroppingSuggestion) SetScore(v float32) *GetImageResponseBodyCroppingSuggestion {
	s.Score = &v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestion) SetCroppingBoundary(v *GetImageResponseBodyCroppingSuggestionCroppingBoundary) *GetImageResponseBodyCroppingSuggestion {
	s.CroppingBoundary = v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestion) SetAspectRatio(v string) *GetImageResponseBodyCroppingSuggestion {
	s.AspectRatio = &v
	return s
}

type GetImageResponseBodyCroppingSuggestionCroppingBoundary struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s GetImageResponseBodyCroppingSuggestionCroppingBoundary) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyCroppingSuggestionCroppingBoundary) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyCroppingSuggestionCroppingBoundary) SetTop(v int32) *GetImageResponseBodyCroppingSuggestionCroppingBoundary {
	s.Top = &v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestionCroppingBoundary) SetWidth(v int32) *GetImageResponseBodyCroppingSuggestionCroppingBoundary {
	s.Width = &v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestionCroppingBoundary) SetHeight(v int32) *GetImageResponseBodyCroppingSuggestionCroppingBoundary {
	s.Height = &v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestionCroppingBoundary) SetLeft(v int32) *GetImageResponseBodyCroppingSuggestionCroppingBoundary {
	s.Left = &v
	return s
}

type GetImageResponseBodyOCR struct {
	OCRConfidence *float32                            `json:"OCRConfidence,omitempty" xml:"OCRConfidence,omitempty"`
	OCRContents   *string                             `json:"OCRContents,omitempty" xml:"OCRContents,omitempty"`
	OCRBoundary   *GetImageResponseBodyOCROCRBoundary `json:"OCRBoundary,omitempty" xml:"OCRBoundary,omitempty" type:"Struct"`
}

func (s GetImageResponseBodyOCR) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyOCR) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyOCR) SetOCRConfidence(v float32) *GetImageResponseBodyOCR {
	s.OCRConfidence = &v
	return s
}

func (s *GetImageResponseBodyOCR) SetOCRContents(v string) *GetImageResponseBodyOCR {
	s.OCRContents = &v
	return s
}

func (s *GetImageResponseBodyOCR) SetOCRBoundary(v *GetImageResponseBodyOCROCRBoundary) *GetImageResponseBodyOCR {
	s.OCRBoundary = v
	return s
}

type GetImageResponseBodyOCROCRBoundary struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s GetImageResponseBodyOCROCRBoundary) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyOCROCRBoundary) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyOCROCRBoundary) SetTop(v int32) *GetImageResponseBodyOCROCRBoundary {
	s.Top = &v
	return s
}

func (s *GetImageResponseBodyOCROCRBoundary) SetWidth(v int32) *GetImageResponseBodyOCROCRBoundary {
	s.Width = &v
	return s
}

func (s *GetImageResponseBodyOCROCRBoundary) SetHeight(v int32) *GetImageResponseBodyOCROCRBoundary {
	s.Height = &v
	return s
}

func (s *GetImageResponseBodyOCROCRBoundary) SetLeft(v int32) *GetImageResponseBodyOCROCRBoundary {
	s.Left = &v
	return s
}

type GetImageResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s *GetImageResponse) SetHeaders(v map[string]*string) *GetImageResponse {
	s.Headers = v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetImageCroppingSuggestionsRequest struct {
	Project      *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri     *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	AspectRatios *string `json:"AspectRatios,omitempty" xml:"AspectRatios,omitempty"`
}

func (s GetImageCroppingSuggestionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsRequest) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsRequest) SetProject(v string) *GetImageCroppingSuggestionsRequest {
	s.Project = &v
	return s
}

func (s *GetImageCroppingSuggestionsRequest) SetImageUri(v string) *GetImageCroppingSuggestionsRequest {
	s.ImageUri = &v
	return s
}

func (s *GetImageCroppingSuggestionsRequest) SetAspectRatios(v string) *GetImageCroppingSuggestionsRequest {
	s.AspectRatios = &v
	return s
}

type GetImageCroppingSuggestionsResponseBody struct {
	ImageUri            *string                                                       `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CroppingSuggestions []*GetImageCroppingSuggestionsResponseBodyCroppingSuggestions `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
}

func (s GetImageCroppingSuggestionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsResponseBody) SetImageUri(v string) *GetImageCroppingSuggestionsResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBody) SetRequestId(v string) *GetImageCroppingSuggestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBody) SetCroppingSuggestions(v []*GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) *GetImageCroppingSuggestionsResponseBody {
	s.CroppingSuggestions = v
	return s
}

type GetImageCroppingSuggestionsResponseBodyCroppingSuggestions struct {
	Score            *float32                                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	AspectRatio      *string                                                                     `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	CroppingBoundary *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary `json:"CroppingBoundary,omitempty" xml:"CroppingBoundary,omitempty" type:"Struct"`
}

func (s GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) SetScore(v float32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions {
	s.Score = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) SetAspectRatio(v string) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions {
	s.AspectRatio = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) SetCroppingBoundary(v *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions {
	s.CroppingBoundary = v
	return s
}

type GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) SetLeft(v int32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary {
	s.Left = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) SetTop(v int32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary {
	s.Top = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) SetWidth(v int32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary {
	s.Width = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) SetHeight(v int32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary {
	s.Height = &v
	return s
}

type GetImageCroppingSuggestionsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageCroppingSuggestionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageCroppingSuggestionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsResponse) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsResponse) SetHeaders(v map[string]*string) *GetImageCroppingSuggestionsResponse {
	s.Headers = v
	return s
}

func (s *GetImageCroppingSuggestionsResponse) SetBody(v *GetImageCroppingSuggestionsResponseBody) *GetImageCroppingSuggestionsResponse {
	s.Body = v
	return s
}

type GetImageQualityRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
}

func (s GetImageQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageQualityRequest) GoString() string {
	return s.String()
}

func (s *GetImageQualityRequest) SetProject(v string) *GetImageQualityRequest {
	s.Project = &v
	return s
}

func (s *GetImageQualityRequest) SetImageUri(v string) *GetImageQualityRequest {
	s.ImageUri = &v
	return s
}

type GetImageQualityResponseBody struct {
	ImageUri     *string                                  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ImageQuality *GetImageQualityResponseBodyImageQuality `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty" type:"Struct"`
}

func (s GetImageQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageQualityResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageQualityResponseBody) SetImageUri(v string) *GetImageQualityResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageQualityResponseBody) SetRequestId(v string) *GetImageQualityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageQualityResponseBody) SetImageQuality(v *GetImageQualityResponseBodyImageQuality) *GetImageQualityResponseBody {
	s.ImageQuality = v
	return s
}

type GetImageQualityResponseBodyImageQuality struct {
	OverallScore     *float32 `json:"OverallScore,omitempty" xml:"OverallScore,omitempty"`
	Color            *float32 `json:"Color,omitempty" xml:"Color,omitempty"`
	ColorScore       *float32 `json:"ColorScore,omitempty" xml:"ColorScore,omitempty"`
	ContrastScore    *float32 `json:"ContrastScore,omitempty" xml:"ContrastScore,omitempty"`
	Contrast         *float32 `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	ExposureScore    *float32 `json:"ExposureScore,omitempty" xml:"ExposureScore,omitempty"`
	ClarityScore     *float32 `json:"ClarityScore,omitempty" xml:"ClarityScore,omitempty"`
	Clarity          *float32 `json:"Clarity,omitempty" xml:"Clarity,omitempty"`
	Exposure         *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
	CompositionScore *float32 `json:"CompositionScore,omitempty" xml:"CompositionScore,omitempty"`
}

func (s GetImageQualityResponseBodyImageQuality) String() string {
	return tea.Prettify(s)
}

func (s GetImageQualityResponseBodyImageQuality) GoString() string {
	return s.String()
}

func (s *GetImageQualityResponseBodyImageQuality) SetOverallScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.OverallScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetColor(v float32) *GetImageQualityResponseBodyImageQuality {
	s.Color = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetColorScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.ColorScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetContrastScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.ContrastScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetContrast(v float32) *GetImageQualityResponseBodyImageQuality {
	s.Contrast = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetExposureScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.ExposureScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetClarityScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.ClarityScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetClarity(v float32) *GetImageQualityResponseBodyImageQuality {
	s.Clarity = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetExposure(v float32) *GetImageQualityResponseBodyImageQuality {
	s.Exposure = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetCompositionScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.CompositionScore = &v
	return s
}

type GetImageQualityResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageQualityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageQualityResponse) GoString() string {
	return s.String()
}

func (s *GetImageQualityResponse) SetHeaders(v map[string]*string) *GetImageQualityResponse {
	s.Headers = v
	return s
}

func (s *GetImageQualityResponse) SetBody(v *GetImageQualityResponseBody) *GetImageQualityResponse {
	s.Body = v
	return s
}

type GetMediaMetaRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	MediaUri *string `json:"MediaUri,omitempty" xml:"MediaUri,omitempty"`
}

func (s GetMediaMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaRequest) GoString() string {
	return s.String()
}

func (s *GetMediaMetaRequest) SetProject(v string) *GetMediaMetaRequest {
	s.Project = &v
	return s
}

func (s *GetMediaMetaRequest) SetMediaUri(v string) *GetMediaMetaRequest {
	s.MediaUri = &v
	return s
}

type GetMediaMetaResponseBody struct {
	MediaUri  *string                            `json:"MediaUri,omitempty" xml:"MediaUri,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MediaMeta *GetMediaMetaResponseBodyMediaMeta `json:"MediaMeta,omitempty" xml:"MediaMeta,omitempty" type:"Struct"`
}

func (s GetMediaMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBody) SetMediaUri(v string) *GetMediaMetaResponseBody {
	s.MediaUri = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetRequestId(v string) *GetMediaMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetMediaMeta(v *GetMediaMetaResponseBodyMediaMeta) *GetMediaMetaResponseBody {
	s.MediaMeta = v
	return s
}

type GetMediaMetaResponseBodyMediaMeta struct {
	MediaFormat  *GetMediaMetaResponseBodyMediaMetaMediaFormat  `json:"MediaFormat,omitempty" xml:"MediaFormat,omitempty" type:"Struct"`
	MediaStreams *GetMediaMetaResponseBodyMediaMetaMediaStreams `json:"MediaStreams,omitempty" xml:"MediaStreams,omitempty" type:"Struct"`
}

func (s GetMediaMetaResponseBodyMediaMeta) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMeta) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMeta) SetMediaFormat(v *GetMediaMetaResponseBodyMediaMetaMediaFormat) *GetMediaMetaResponseBodyMediaMeta {
	s.MediaFormat = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMeta) SetMediaStreams(v *GetMediaMetaResponseBodyMediaMetaMediaStreams) *GetMediaMetaResponseBodyMediaMeta {
	s.MediaStreams = v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaFormat struct {
	CreationTime   *string                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	NumberPrograms *int32                                               `json:"NumberPrograms,omitempty" xml:"NumberPrograms,omitempty"`
	NumberStreams  *int32                                               `json:"NumberStreams,omitempty" xml:"NumberStreams,omitempty"`
	Tag            *GetMediaMetaResponseBodyMediaMetaMediaFormatTag     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	Bitrate        *string                                              `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	StartTime      *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Size           *string                                              `json:"Size,omitempty" xml:"Size,omitempty"`
	Address        *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	FormatLongName *string                                              `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	Duration       *string                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FormatName     *string                                              `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	Location       *string                                              `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormat) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormat) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetCreationTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.CreationTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetNumberPrograms(v int32) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.NumberPrograms = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetNumberStreams(v int32) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.NumberStreams = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetTag(v *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Tag = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetBitrate(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Bitrate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetStartTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.StartTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetSize(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Size = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetAddress(v *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Address = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetFormatLongName(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.FormatLongName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetDuration(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Duration = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetFormatName(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.FormatName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetLocation(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Location = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaFormatTag struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Album        *string `json:"Album,omitempty" xml:"Album,omitempty"`
	AlbumArtist  *string `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	Performer    *string `json:"Performer,omitempty" xml:"Performer,omitempty"`
	Composer     *string `json:"Composer,omitempty" xml:"Composer,omitempty"`
	Artist       *string `json:"Artist,omitempty" xml:"Artist,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Language     *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormatTag) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormatTag) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetCreationTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.CreationTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetAlbum(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Album = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetAlbumArtist(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.AlbumArtist = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetPerformer(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Performer = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetComposer(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Composer = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetArtist(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Artist = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetTitle(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Title = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetLanguage(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Language = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaFormatAddress struct {
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetTownship(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.Township = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetDistrict(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.District = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetAddressLine(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.AddressLine = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetCountry(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.Country = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetCity(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.City = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetProvince(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.Province = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaStreams struct {
	VideoStreams    []*GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams    `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
	AudioStreams    []*GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams    `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	SubtitleStreams []*GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams `json:"SubtitleStreams,omitempty" xml:"SubtitleStreams,omitempty" type:"Repeated"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreams) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreams) SetVideoStreams(v []*GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) *GetMediaMetaResponseBodyMediaMetaMediaStreams {
	s.VideoStreams = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreams) SetAudioStreams(v []*GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) *GetMediaMetaResponseBodyMediaMetaMediaStreams {
	s.AudioStreams = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreams) SetSubtitleStreams(v []*GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) *GetMediaMetaResponseBodyMediaMetaMediaStreams {
	s.SubtitleStreams = v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams struct {
	Index              *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	CodecLongName      *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	Height             *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	SampleAspectRatio  *string `json:"SampleAspectRatio,omitempty" xml:"SampleAspectRatio,omitempty"`
	AverageFrameRate   *string `json:"AverageFrameRate,omitempty" xml:"AverageFrameRate,omitempty"`
	Bitrate            *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Rotate             *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	CodecTagString     *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	Language           *string `json:"Language,omitempty" xml:"Language,omitempty"`
	HasBFrames         *int32  `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	FrameRrate         *string `json:"FrameRrate,omitempty" xml:"FrameRrate,omitempty"`
	Profile            *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Frames             *string `json:"Frames,omitempty" xml:"Frames,omitempty"`
	CodecName          *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	Width              *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	Duration           *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DisplayAspectRatio *string `json:"DisplayAspectRatio,omitempty" xml:"DisplayAspectRatio,omitempty"`
	CodecTag           *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTimeBase      *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	TimeBase           *string `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
	Level              *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	PixelFormat        *string `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetIndex(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Index = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecLongName(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetHeight(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Height = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetSampleAspectRatio(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.SampleAspectRatio = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetAverageFrameRate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.AverageFrameRate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetBitrate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Bitrate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetRotate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Rotate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecTagString(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetLanguage(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Language = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetHasBFrames(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.HasBFrames = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetFrameRrate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.FrameRrate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetProfile(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Profile = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetStartTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.StartTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetFrames(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Frames = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecName(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetWidth(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Width = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetDuration(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Duration = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetDisplayAspectRatio(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.DisplayAspectRatio = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecTag(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecTag = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecTimeBase(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetTimeBase(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.TimeBase = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetLevel(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Level = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetPixelFormat(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.PixelFormat = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams struct {
	Index          *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	SampleRate     *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	ChannelLayout  *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	CodecLongName  *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	Channels       *int32  `json:"Channels,omitempty" xml:"Channels,omitempty"`
	Bitrate        *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SampleFormat   *string `json:"SampleFormat,omitempty" xml:"SampleFormat,omitempty"`
	Frames         *string `json:"Frames,omitempty" xml:"Frames,omitempty"`
	CodecName      *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	CodecTag       *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTimeBase  *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	TimeBase       *string `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetIndex(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Index = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetSampleRate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.SampleRate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetChannelLayout(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.ChannelLayout = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecLongName(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetChannels(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Channels = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetBitrate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Bitrate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecTagString(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetLanguage(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Language = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetStartTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.StartTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetSampleFormat(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.SampleFormat = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetFrames(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Frames = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecName(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetDuration(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Duration = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecTag(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecTag = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecTimeBase(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetTimeBase(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.TimeBase = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams struct {
	Index    *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) SetIndex(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams {
	s.Index = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) SetLanguage(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams {
	s.Language = &v
	return s
}

type GetMediaMetaResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponse) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponse) SetHeaders(v map[string]*string) *GetMediaMetaResponse {
	s.Headers = v
	return s
}

func (s *GetMediaMetaResponse) SetBody(v *GetMediaMetaResponseBody) *GetMediaMetaResponse {
	s.Body = v
	return s
}

type GetOfficeConversionTaskRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetOfficeConversionTaskRequest) SetProject(v string) *GetOfficeConversionTaskRequest {
	s.Project = &v
	return s
}

func (s *GetOfficeConversionTaskRequest) SetTaskId(v string) *GetOfficeConversionTaskRequest {
	s.TaskId = &v
	return s
}

type GetOfficeConversionTaskResponseBody struct {
	Status          *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent         *int32                                         `json:"Percent,omitempty" xml:"Percent,omitempty"`
	FinishTime      *string                                        `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	CreateTime      *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PageCount       *int32                                         `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	NotifyTopicName *string                                        `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NotifyEndpoint  *string                                        `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	SrcUri          *string                                        `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	TgtType         *string                                        `json:"TgtType,omitempty" xml:"TgtType,omitempty"`
	TgtUri          *string                                        `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	ImageSpec       *string                                        `json:"ImageSpec,omitempty" xml:"ImageSpec,omitempty"`
	ExternalID      *string                                        `json:"ExternalID,omitempty" xml:"ExternalID,omitempty"`
	TaskId          *string                                        `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	FailDetail      *GetOfficeConversionTaskResponseBodyFailDetail `json:"FailDetail,omitempty" xml:"FailDetail,omitempty" type:"Struct"`
}

func (s GetOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficeConversionTaskResponseBody) SetStatus(v string) *GetOfficeConversionTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetPercent(v int32) *GetOfficeConversionTaskResponseBody {
	s.Percent = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetFinishTime(v string) *GetOfficeConversionTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetCreateTime(v string) *GetOfficeConversionTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetPageCount(v int32) *GetOfficeConversionTaskResponseBody {
	s.PageCount = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetNotifyTopicName(v string) *GetOfficeConversionTaskResponseBody {
	s.NotifyTopicName = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetRequestId(v string) *GetOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetNotifyEndpoint(v string) *GetOfficeConversionTaskResponseBody {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetSrcUri(v string) *GetOfficeConversionTaskResponseBody {
	s.SrcUri = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetTgtType(v string) *GetOfficeConversionTaskResponseBody {
	s.TgtType = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetTgtUri(v string) *GetOfficeConversionTaskResponseBody {
	s.TgtUri = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetImageSpec(v string) *GetOfficeConversionTaskResponseBody {
	s.ImageSpec = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetExternalID(v string) *GetOfficeConversionTaskResponseBody {
	s.ExternalID = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetTaskId(v string) *GetOfficeConversionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetFailDetail(v *GetOfficeConversionTaskResponseBodyFailDetail) *GetOfficeConversionTaskResponseBody {
	s.FailDetail = v
	return s
}

type GetOfficeConversionTaskResponseBodyFailDetail struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetOfficeConversionTaskResponseBodyFailDetail) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeConversionTaskResponseBodyFailDetail) GoString() string {
	return s.String()
}

func (s *GetOfficeConversionTaskResponseBodyFailDetail) SetCode(v string) *GetOfficeConversionTaskResponseBodyFailDetail {
	s.Code = &v
	return s
}

type GetOfficeConversionTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *GetOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetOfficeConversionTaskResponse) SetBody(v *GetOfficeConversionTaskResponseBody) *GetOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type GetOfficeEditURLRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcUri          *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	SrcType         *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	FileID          *string `json:"FileID,omitempty" xml:"FileID,omitempty"`
	TgtUri          *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	UserID          *string `json:"UserID,omitempty" xml:"UserID,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetOfficeEditURLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeEditURLRequest) GoString() string {
	return s.String()
}

func (s *GetOfficeEditURLRequest) SetProject(v string) *GetOfficeEditURLRequest {
	s.Project = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetSrcUri(v string) *GetOfficeEditURLRequest {
	s.SrcUri = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetSrcType(v string) *GetOfficeEditURLRequest {
	s.SrcType = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetFileID(v string) *GetOfficeEditURLRequest {
	s.FileID = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetTgtUri(v string) *GetOfficeEditURLRequest {
	s.TgtUri = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetUserID(v string) *GetOfficeEditURLRequest {
	s.UserID = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetUserName(v string) *GetOfficeEditURLRequest {
	s.UserName = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetNotifyEndpoint(v string) *GetOfficeEditURLRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetNotifyTopicName(v string) *GetOfficeEditURLRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetOfficeEditURLRequest) SetFileName(v string) *GetOfficeEditURLRequest {
	s.FileName = &v
	return s
}

type GetOfficeEditURLResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	EditURL                 *string `json:"EditURL,omitempty" xml:"EditURL,omitempty"`
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
}

func (s GetOfficeEditURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeEditURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficeEditURLResponseBody) SetRequestId(v string) *GetOfficeEditURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfficeEditURLResponseBody) SetAccessTokenExpiredTime(v string) *GetOfficeEditURLResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetOfficeEditURLResponseBody) SetEditURL(v string) *GetOfficeEditURLResponseBody {
	s.EditURL = &v
	return s
}

func (s *GetOfficeEditURLResponseBody) SetAccessToken(v string) *GetOfficeEditURLResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetOfficeEditURLResponseBody) SetRefreshToken(v string) *GetOfficeEditURLResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetOfficeEditURLResponseBody) SetRefreshTokenExpiredTime(v string) *GetOfficeEditURLResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

type GetOfficeEditURLResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficeEditURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficeEditURLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeEditURLResponse) GoString() string {
	return s.String()
}

func (s *GetOfficeEditURLResponse) SetHeaders(v map[string]*string) *GetOfficeEditURLResponse {
	s.Headers = v
	return s
}

func (s *GetOfficeEditURLResponse) SetBody(v *GetOfficeEditURLResponseBody) *GetOfficeEditURLResponse {
	s.Body = v
	return s
}

type GetOfficePreviewURLRequest struct {
	Project             *string  `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcUri              *string  `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	SrcType             *string  `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	WatermarkType       *int32   `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkValue      *string  `json:"WatermarkValue,omitempty" xml:"WatermarkValue,omitempty"`
	WatermarkFillStyle  *string  `json:"WatermarkFillStyle,omitempty" xml:"WatermarkFillStyle,omitempty"`
	WatermarkFont       *string  `json:"WatermarkFont,omitempty" xml:"WatermarkFont,omitempty"`
	WatermarkRotate     *float32 `json:"WatermarkRotate,omitempty" xml:"WatermarkRotate,omitempty"`
	WatermarkHorizontal *int32   `json:"WatermarkHorizontal,omitempty" xml:"WatermarkHorizontal,omitempty"`
	WatermarkVertical   *int32   `json:"WatermarkVertical,omitempty" xml:"WatermarkVertical,omitempty"`
}

func (s GetOfficePreviewURLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficePreviewURLRequest) GoString() string {
	return s.String()
}

func (s *GetOfficePreviewURLRequest) SetProject(v string) *GetOfficePreviewURLRequest {
	s.Project = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetSrcUri(v string) *GetOfficePreviewURLRequest {
	s.SrcUri = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetSrcType(v string) *GetOfficePreviewURLRequest {
	s.SrcType = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkType(v int32) *GetOfficePreviewURLRequest {
	s.WatermarkType = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkValue(v string) *GetOfficePreviewURLRequest {
	s.WatermarkValue = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkFillStyle(v string) *GetOfficePreviewURLRequest {
	s.WatermarkFillStyle = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkFont(v string) *GetOfficePreviewURLRequest {
	s.WatermarkFont = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkRotate(v float32) *GetOfficePreviewURLRequest {
	s.WatermarkRotate = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkHorizontal(v int32) *GetOfficePreviewURLRequest {
	s.WatermarkHorizontal = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkVertical(v int32) *GetOfficePreviewURLRequest {
	s.WatermarkVertical = &v
	return s
}

type GetOfficePreviewURLResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreviewURL              *string `json:"PreviewURL,omitempty" xml:"PreviewURL,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
}

func (s GetOfficePreviewURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficePreviewURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficePreviewURLResponseBody) SetRequestId(v string) *GetOfficePreviewURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetPreviewURL(v string) *GetOfficePreviewURLResponseBody {
	s.PreviewURL = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetAccessTokenExpiredTime(v string) *GetOfficePreviewURLResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetAccessToken(v string) *GetOfficePreviewURLResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetRefreshToken(v string) *GetOfficePreviewURLResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetRefreshTokenExpiredTime(v string) *GetOfficePreviewURLResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

type GetOfficePreviewURLResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficePreviewURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficePreviewURLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficePreviewURLResponse) GoString() string {
	return s.String()
}

func (s *GetOfficePreviewURLResponse) SetHeaders(v map[string]*string) *GetOfficePreviewURLResponse {
	s.Headers = v
	return s
}

func (s *GetOfficePreviewURLResponse) SetBody(v *GetOfficePreviewURLResponseBody) *GetOfficePreviewURLResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetProject(v string) *GetProjectRequest {
	s.Project = &v
	return s
}

type GetProjectResponseBody struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CU          *int32  `json:"CU,omitempty" xml:"CU,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetType(v string) *GetProjectResponseBody {
	s.Type = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) SetCU(v int32) *GetProjectResponseBody {
	s.CU = &v
	return s
}

func (s *GetProjectResponseBody) SetCreateTime(v string) *GetProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetEndpoint(v string) *GetProjectResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetProjectResponseBody) SetServiceRole(v string) *GetProjectResponseBody {
	s.ServiceRole = &v
	return s
}

func (s *GetProjectResponseBody) SetProject(v string) *GetProjectResponseBody {
	s.Project = &v
	return s
}

func (s *GetProjectResponseBody) SetRegionId(v string) *GetProjectResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetProjectResponseBody) SetBillingType(v string) *GetProjectResponseBody {
	s.BillingType = &v
	return s
}

func (s *GetProjectResponseBody) SetModifyTime(v string) *GetProjectResponseBody {
	s.ModifyTime = &v
	return s
}

type GetProjectResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetSetRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s GetSetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSetRequest) GoString() string {
	return s.String()
}

func (s *GetSetRequest) SetProject(v string) *GetSetRequest {
	s.Project = &v
	return s
}

func (s *GetSetRequest) SetSetId(v string) *GetSetRequest {
	s.SetId = &v
	return s
}

type GetSetResponseBody struct {
	VideoCount  *int32  `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VideoLength *int32  `json:"VideoLength,omitempty" xml:"VideoLength,omitempty"`
	SetId       *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageCount  *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	FaceCount   *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	SetName     *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s GetSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetSetResponseBody) SetVideoCount(v int32) *GetSetResponseBody {
	s.VideoCount = &v
	return s
}

func (s *GetSetResponseBody) SetRequestId(v string) *GetSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSetResponseBody) SetCreateTime(v string) *GetSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSetResponseBody) SetVideoLength(v int32) *GetSetResponseBody {
	s.VideoLength = &v
	return s
}

func (s *GetSetResponseBody) SetSetId(v string) *GetSetResponseBody {
	s.SetId = &v
	return s
}

func (s *GetSetResponseBody) SetImageCount(v int32) *GetSetResponseBody {
	s.ImageCount = &v
	return s
}

func (s *GetSetResponseBody) SetFaceCount(v int32) *GetSetResponseBody {
	s.FaceCount = &v
	return s
}

func (s *GetSetResponseBody) SetSetName(v string) *GetSetResponseBody {
	s.SetName = &v
	return s
}

func (s *GetSetResponseBody) SetModifyTime(v string) *GetSetResponseBody {
	s.ModifyTime = &v
	return s
}

type GetSetResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSetResponse) GoString() string {
	return s.String()
}

func (s *GetSetResponse) SetHeaders(v map[string]*string) *GetSetResponse {
	s.Headers = v
	return s
}

func (s *GetSetResponse) SetBody(v *GetSetResponseBody) *GetSetResponse {
	s.Body = v
	return s
}

type GetVideoRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
}

func (s GetVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoRequest) SetProject(v string) *GetVideoRequest {
	s.Project = &v
	return s
}

func (s *GetVideoRequest) SetSetId(v string) *GetVideoRequest {
	s.SetId = &v
	return s
}

func (s *GetVideoRequest) SetVideoUri(v string) *GetVideoRequest {
	s.VideoUri = &v
	return s
}

type GetVideoResponseBody struct {
	ModifyTime               *string                          `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ProcessStatus            *string                          `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	VideoWidth               *int32                           `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
	SourceType               *string                          `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri                *string                          `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	VideoInfo                *string                          `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty"`
	VideoFrameTagsModifyTime *string                          `json:"VideoFrameTagsModifyTime,omitempty" xml:"VideoFrameTagsModifyTime,omitempty"`
	RemarksA                 *string                          `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	VideoFacesFailReason     *string                          `json:"VideoFacesFailReason,omitempty" xml:"VideoFacesFailReason,omitempty"`
	RemarksB                 *string                          `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	VideoFacesStatus         *string                          `json:"VideoFacesStatus,omitempty" xml:"VideoFacesStatus,omitempty"`
	RemarksC                 *string                          `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	VideoOCRModifyTime       *string                          `json:"VideoOCRModifyTime,omitempty" xml:"VideoOCRModifyTime,omitempty"`
	RemarksD                 *string                          `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	VideoHeight              *int32                           `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	SourcePosition           *string                          `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	VideoOCRFailReason       *string                          `json:"VideoOCRFailReason,omitempty" xml:"VideoOCRFailReason,omitempty"`
	VideoFrameTagsStatus     *string                          `json:"VideoFrameTagsStatus,omitempty" xml:"VideoFrameTagsStatus,omitempty"`
	VideoTagsFailReason      *string                          `json:"VideoTagsFailReason,omitempty" xml:"VideoTagsFailReason,omitempty"`
	VideoTagsModifyTime      *string                          `json:"VideoTagsModifyTime,omitempty" xml:"VideoTagsModifyTime,omitempty"`
	VideoOCRStatus           *string                          `json:"VideoOCRStatus,omitempty" xml:"VideoOCRStatus,omitempty"`
	VideoFrames              *int32                           `json:"VideoFrames,omitempty" xml:"VideoFrames,omitempty"`
	RequestId                *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProcessModifyTime        *string                          `json:"ProcessModifyTime,omitempty" xml:"ProcessModifyTime,omitempty"`
	VideoSTTModifyTime       *string                          `json:"VideoSTTModifyTime,omitempty" xml:"VideoSTTModifyTime,omitempty"`
	ProcessFailReason        *string                          `json:"ProcessFailReason,omitempty" xml:"ProcessFailReason,omitempty"`
	CreateTime               *string                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId               *string                          `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	VideoSTTFailReason       *string                          `json:"VideoSTTFailReason,omitempty" xml:"VideoSTTFailReason,omitempty"`
	VideoUri                 *string                          `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	VideoFrameTagsFailReason *string                          `json:"VideoFrameTagsFailReason,omitempty" xml:"VideoFrameTagsFailReason,omitempty"`
	VideoFormat              *string                          `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	VideoSTTStatus           *string                          `json:"VideoSTTStatus,omitempty" xml:"VideoSTTStatus,omitempty"`
	VideoFacesModifyTime     *string                          `json:"VideoFacesModifyTime,omitempty" xml:"VideoFacesModifyTime,omitempty"`
	VideoTags                []*GetVideoResponseBodyVideoTags `json:"VideoTags,omitempty" xml:"VideoTags,omitempty" type:"Repeated"`
	VideoDuration            *float32                         `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	SetId                    *string                          `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoTagsStatus          *string                          `json:"VideoTagsStatus,omitempty" xml:"VideoTagsStatus,omitempty"`
	FileSize                 *int32                           `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
}

func (s GetVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoResponseBody) SetModifyTime(v string) *GetVideoResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetProcessStatus(v string) *GetVideoResponseBody {
	s.ProcessStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoWidth(v int32) *GetVideoResponseBody {
	s.VideoWidth = &v
	return s
}

func (s *GetVideoResponseBody) SetSourceType(v string) *GetVideoResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetVideoResponseBody) SetSourceUri(v string) *GetVideoResponseBody {
	s.SourceUri = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoInfo(v string) *GetVideoResponseBody {
	s.VideoInfo = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFrameTagsModifyTime(v string) *GetVideoResponseBody {
	s.VideoFrameTagsModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetRemarksA(v string) *GetVideoResponseBody {
	s.RemarksA = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFacesFailReason(v string) *GetVideoResponseBody {
	s.VideoFacesFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetRemarksB(v string) *GetVideoResponseBody {
	s.RemarksB = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFacesStatus(v string) *GetVideoResponseBody {
	s.VideoFacesStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetRemarksC(v string) *GetVideoResponseBody {
	s.RemarksC = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoOCRModifyTime(v string) *GetVideoResponseBody {
	s.VideoOCRModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetRemarksD(v string) *GetVideoResponseBody {
	s.RemarksD = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoHeight(v int32) *GetVideoResponseBody {
	s.VideoHeight = &v
	return s
}

func (s *GetVideoResponseBody) SetSourcePosition(v string) *GetVideoResponseBody {
	s.SourcePosition = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoOCRFailReason(v string) *GetVideoResponseBody {
	s.VideoOCRFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFrameTagsStatus(v string) *GetVideoResponseBody {
	s.VideoFrameTagsStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoTagsFailReason(v string) *GetVideoResponseBody {
	s.VideoTagsFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoTagsModifyTime(v string) *GetVideoResponseBody {
	s.VideoTagsModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoOCRStatus(v string) *GetVideoResponseBody {
	s.VideoOCRStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFrames(v int32) *GetVideoResponseBody {
	s.VideoFrames = &v
	return s
}

func (s *GetVideoResponseBody) SetRequestId(v string) *GetVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoResponseBody) SetProcessModifyTime(v string) *GetVideoResponseBody {
	s.ProcessModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoSTTModifyTime(v string) *GetVideoResponseBody {
	s.VideoSTTModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetProcessFailReason(v string) *GetVideoResponseBody {
	s.ProcessFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetCreateTime(v string) *GetVideoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVideoResponseBody) SetExternalId(v string) *GetVideoResponseBody {
	s.ExternalId = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoSTTFailReason(v string) *GetVideoResponseBody {
	s.VideoSTTFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoUri(v string) *GetVideoResponseBody {
	s.VideoUri = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFrameTagsFailReason(v string) *GetVideoResponseBody {
	s.VideoFrameTagsFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFormat(v string) *GetVideoResponseBody {
	s.VideoFormat = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoSTTStatus(v string) *GetVideoResponseBody {
	s.VideoSTTStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFacesModifyTime(v string) *GetVideoResponseBody {
	s.VideoFacesModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoTags(v []*GetVideoResponseBodyVideoTags) *GetVideoResponseBody {
	s.VideoTags = v
	return s
}

func (s *GetVideoResponseBody) SetVideoDuration(v float32) *GetVideoResponseBody {
	s.VideoDuration = &v
	return s
}

func (s *GetVideoResponseBody) SetSetId(v string) *GetVideoResponseBody {
	s.SetId = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoTagsStatus(v string) *GetVideoResponseBody {
	s.VideoTagsStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetFileSize(v int32) *GetVideoResponseBody {
	s.FileSize = &v
	return s
}

type GetVideoResponseBodyVideoTags struct {
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
}

func (s GetVideoResponseBodyVideoTags) String() string {
	return tea.Prettify(s)
}

func (s GetVideoResponseBodyVideoTags) GoString() string {
	return s.String()
}

func (s *GetVideoResponseBodyVideoTags) SetTagName(v string) *GetVideoResponseBodyVideoTags {
	s.TagName = &v
	return s
}

func (s *GetVideoResponseBodyVideoTags) SetTagConfidence(v float32) *GetVideoResponseBodyVideoTags {
	s.TagConfidence = &v
	return s
}

func (s *GetVideoResponseBodyVideoTags) SetTagLevel(v int32) *GetVideoResponseBodyVideoTags {
	s.TagLevel = &v
	return s
}

func (s *GetVideoResponseBodyVideoTags) SetParentTagName(v string) *GetVideoResponseBodyVideoTags {
	s.ParentTagName = &v
	return s
}

type GetVideoResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoResponse) SetHeaders(v map[string]*string) *GetVideoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoResponse) SetBody(v *GetVideoResponseBody) *GetVideoResponse {
	s.Body = v
	return s
}

type GetVideoTaskRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *GetVideoTaskRequest) SetProject(v string) *GetVideoTaskRequest {
	s.Project = &v
	return s
}

func (s *GetVideoTaskRequest) SetTaskType(v string) *GetVideoTaskRequest {
	s.TaskType = &v
	return s
}

func (s *GetVideoTaskRequest) SetTaskId(v string) *GetVideoTaskRequest {
	s.TaskId = &v
	return s
}

type GetVideoTaskResponseBody struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Progress        *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	Parameters      *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ErrorMessage    *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Result          *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoTaskResponseBody) SetStatus(v string) *GetVideoTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetProgress(v int32) *GetVideoTaskResponseBody {
	s.Progress = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetNotifyEndpoint(v string) *GetVideoTaskResponseBody {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetParameters(v string) *GetVideoTaskResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetTaskId(v string) *GetVideoTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetEndTime(v string) *GetVideoTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetRequestId(v string) *GetVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetTaskType(v string) *GetVideoTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetStartTime(v string) *GetVideoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetNotifyTopicName(v string) *GetVideoTaskResponseBody {
	s.NotifyTopicName = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetErrorMessage(v string) *GetVideoTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetResult(v string) *GetVideoTaskResponseBody {
	s.Result = &v
	return s
}

type GetVideoTaskResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *GetVideoTaskResponse) SetHeaders(v map[string]*string) *GetVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *GetVideoTaskResponse) SetBody(v *GetVideoTaskResponseBody) *GetVideoTaskResponse {
	s.Body = v
	return s
}

type GetWebofficeURLRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcType         *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	FileID          *string `json:"FileID,omitempty" xml:"FileID,omitempty"`
	User            *string `json:"User,omitempty" xml:"User,omitempty"`
	Permission      *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	File            *string `json:"File,omitempty" xml:"File,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Watermark       *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Hidecmb         *bool   `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
}

func (s GetWebofficeURLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLRequest) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLRequest) SetProject(v string) *GetWebofficeURLRequest {
	s.Project = &v
	return s
}

func (s *GetWebofficeURLRequest) SetSrcType(v string) *GetWebofficeURLRequest {
	s.SrcType = &v
	return s
}

func (s *GetWebofficeURLRequest) SetFileID(v string) *GetWebofficeURLRequest {
	s.FileID = &v
	return s
}

func (s *GetWebofficeURLRequest) SetUser(v string) *GetWebofficeURLRequest {
	s.User = &v
	return s
}

func (s *GetWebofficeURLRequest) SetPermission(v string) *GetWebofficeURLRequest {
	s.Permission = &v
	return s
}

func (s *GetWebofficeURLRequest) SetFile(v string) *GetWebofficeURLRequest {
	s.File = &v
	return s
}

func (s *GetWebofficeURLRequest) SetNotifyEndpoint(v string) *GetWebofficeURLRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetWebofficeURLRequest) SetNotifyTopicName(v string) *GetWebofficeURLRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetWebofficeURLRequest) SetWatermark(v string) *GetWebofficeURLRequest {
	s.Watermark = &v
	return s
}

func (s *GetWebofficeURLRequest) SetHidecmb(v bool) *GetWebofficeURLRequest {
	s.Hidecmb = &v
	return s
}

type GetWebofficeURLResponseBody struct {
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	WebofficeURL            *string `json:"WebofficeURL,omitempty" xml:"WebofficeURL,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
}

func (s GetWebofficeURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLResponseBody) SetRefreshToken(v string) *GetWebofficeURLResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetRequestId(v string) *GetWebofficeURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetAccessToken(v string) *GetWebofficeURLResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetRefreshTokenExpiredTime(v string) *GetWebofficeURLResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetWebofficeURL(v string) *GetWebofficeURLResponseBody {
	s.WebofficeURL = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetAccessTokenExpiredTime(v string) *GetWebofficeURLResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

type GetWebofficeURLResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebofficeURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebofficeURLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLResponse) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLResponse) SetHeaders(v map[string]*string) *GetWebofficeURLResponse {
	s.Headers = v
	return s
}

func (s *GetWebofficeURLResponse) SetBody(v *GetWebofficeURLResponseBody) *GetWebofficeURLResponse {
	s.Body = v
	return s
}

type IndexImageRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageUri        *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RemarksA        *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB        *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri       *string `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	SourcePosition  *string `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	RemarksC        *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD        *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	ExternalId      *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	RemarksArrayA   *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB   *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
}

func (s IndexImageRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexImageRequest) GoString() string {
	return s.String()
}

func (s *IndexImageRequest) SetProject(v string) *IndexImageRequest {
	s.Project = &v
	return s
}

func (s *IndexImageRequest) SetSetId(v string) *IndexImageRequest {
	s.SetId = &v
	return s
}

func (s *IndexImageRequest) SetImageUri(v string) *IndexImageRequest {
	s.ImageUri = &v
	return s
}

func (s *IndexImageRequest) SetRemarksA(v string) *IndexImageRequest {
	s.RemarksA = &v
	return s
}

func (s *IndexImageRequest) SetRemarksB(v string) *IndexImageRequest {
	s.RemarksB = &v
	return s
}

func (s *IndexImageRequest) SetSourceType(v string) *IndexImageRequest {
	s.SourceType = &v
	return s
}

func (s *IndexImageRequest) SetSourceUri(v string) *IndexImageRequest {
	s.SourceUri = &v
	return s
}

func (s *IndexImageRequest) SetSourcePosition(v string) *IndexImageRequest {
	s.SourcePosition = &v
	return s
}

func (s *IndexImageRequest) SetRemarksC(v string) *IndexImageRequest {
	s.RemarksC = &v
	return s
}

func (s *IndexImageRequest) SetRemarksD(v string) *IndexImageRequest {
	s.RemarksD = &v
	return s
}

func (s *IndexImageRequest) SetExternalId(v string) *IndexImageRequest {
	s.ExternalId = &v
	return s
}

func (s *IndexImageRequest) SetNotifyEndpoint(v string) *IndexImageRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *IndexImageRequest) SetNotifyTopicName(v string) *IndexImageRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *IndexImageRequest) SetRemarksArrayA(v string) *IndexImageRequest {
	s.RemarksArrayA = &v
	return s
}

func (s *IndexImageRequest) SetRemarksArrayB(v string) *IndexImageRequest {
	s.RemarksArrayB = &v
	return s
}

type IndexImageResponseBody struct {
	RemarksArrayB *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	ModifyTime    *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RemarksC      *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD      *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId    *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	RemarksArrayA *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksA      *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	ImageUri      *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	SetId         *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	RemarksB      *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
}

func (s IndexImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndexImageResponseBody) GoString() string {
	return s.String()
}

func (s *IndexImageResponseBody) SetRemarksArrayB(v string) *IndexImageResponseBody {
	s.RemarksArrayB = &v
	return s
}

func (s *IndexImageResponseBody) SetModifyTime(v string) *IndexImageResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksC(v string) *IndexImageResponseBody {
	s.RemarksC = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksD(v string) *IndexImageResponseBody {
	s.RemarksD = &v
	return s
}

func (s *IndexImageResponseBody) SetRequestId(v string) *IndexImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *IndexImageResponseBody) SetCreateTime(v string) *IndexImageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *IndexImageResponseBody) SetExternalId(v string) *IndexImageResponseBody {
	s.ExternalId = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksArrayA(v string) *IndexImageResponseBody {
	s.RemarksArrayA = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksA(v string) *IndexImageResponseBody {
	s.RemarksA = &v
	return s
}

func (s *IndexImageResponseBody) SetImageUri(v string) *IndexImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *IndexImageResponseBody) SetSetId(v string) *IndexImageResponseBody {
	s.SetId = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksB(v string) *IndexImageResponseBody {
	s.RemarksB = &v
	return s
}

type IndexImageResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndexImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndexImageResponse) String() string {
	return tea.Prettify(s)
}

func (s IndexImageResponse) GoString() string {
	return s.String()
}

func (s *IndexImageResponse) SetHeaders(v map[string]*string) *IndexImageResponse {
	s.Headers = v
	return s
}

func (s *IndexImageResponse) SetBody(v *IndexImageResponseBody) *IndexImageResponse {
	s.Body = v
	return s
}

type IndexVideoRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri        *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	RemarksA        *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB        *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	TgtUri          *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	RemarksC        *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD        *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	ExternalId      *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
}

func (s IndexVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexVideoRequest) GoString() string {
	return s.String()
}

func (s *IndexVideoRequest) SetProject(v string) *IndexVideoRequest {
	s.Project = &v
	return s
}

func (s *IndexVideoRequest) SetSetId(v string) *IndexVideoRequest {
	s.SetId = &v
	return s
}

func (s *IndexVideoRequest) SetVideoUri(v string) *IndexVideoRequest {
	s.VideoUri = &v
	return s
}

func (s *IndexVideoRequest) SetRemarksA(v string) *IndexVideoRequest {
	s.RemarksA = &v
	return s
}

func (s *IndexVideoRequest) SetRemarksB(v string) *IndexVideoRequest {
	s.RemarksB = &v
	return s
}

func (s *IndexVideoRequest) SetTgtUri(v string) *IndexVideoRequest {
	s.TgtUri = &v
	return s
}

func (s *IndexVideoRequest) SetRemarksC(v string) *IndexVideoRequest {
	s.RemarksC = &v
	return s
}

func (s *IndexVideoRequest) SetRemarksD(v string) *IndexVideoRequest {
	s.RemarksD = &v
	return s
}

func (s *IndexVideoRequest) SetExternalId(v string) *IndexVideoRequest {
	s.ExternalId = &v
	return s
}

func (s *IndexVideoRequest) SetNotifyTopicName(v string) *IndexVideoRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *IndexVideoRequest) SetNotifyEndpoint(v string) *IndexVideoRequest {
	s.NotifyEndpoint = &v
	return s
}

type IndexVideoResponseBody struct {
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	VideoUri   *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	RemarksA   *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB   *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC   *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD   *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	SetId      *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s IndexVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndexVideoResponseBody) GoString() string {
	return s.String()
}

func (s *IndexVideoResponseBody) SetModifyTime(v string) *IndexVideoResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *IndexVideoResponseBody) SetRequestId(v string) *IndexVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *IndexVideoResponseBody) SetCreateTime(v string) *IndexVideoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *IndexVideoResponseBody) SetExternalId(v string) *IndexVideoResponseBody {
	s.ExternalId = &v
	return s
}

func (s *IndexVideoResponseBody) SetVideoUri(v string) *IndexVideoResponseBody {
	s.VideoUri = &v
	return s
}

func (s *IndexVideoResponseBody) SetRemarksA(v string) *IndexVideoResponseBody {
	s.RemarksA = &v
	return s
}

func (s *IndexVideoResponseBody) SetRemarksB(v string) *IndexVideoResponseBody {
	s.RemarksB = &v
	return s
}

func (s *IndexVideoResponseBody) SetRemarksC(v string) *IndexVideoResponseBody {
	s.RemarksC = &v
	return s
}

func (s *IndexVideoResponseBody) SetRemarksD(v string) *IndexVideoResponseBody {
	s.RemarksD = &v
	return s
}

func (s *IndexVideoResponseBody) SetSetId(v string) *IndexVideoResponseBody {
	s.SetId = &v
	return s
}

type IndexVideoResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndexVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndexVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s IndexVideoResponse) GoString() string {
	return s.String()
}

func (s *IndexVideoResponse) SetHeaders(v map[string]*string) *IndexVideoResponse {
	s.Headers = v
	return s
}

func (s *IndexVideoResponse) SetBody(v *IndexVideoResponseBody) *IndexVideoResponse {
	s.Body = v
	return s
}

type ListFaceGroupsRequest struct {
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId              *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Marker             *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Limit              *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Order              *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy            *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	RemarksAQuery      *string `json:"RemarksAQuery,omitempty" xml:"RemarksAQuery,omitempty"`
	RemarksBQuery      *string `json:"RemarksBQuery,omitempty" xml:"RemarksBQuery,omitempty"`
	RemarksCQuery      *string `json:"RemarksCQuery,omitempty" xml:"RemarksCQuery,omitempty"`
	RemarksDQuery      *string `json:"RemarksDQuery,omitempty" xml:"RemarksDQuery,omitempty"`
	RemarksArrayAQuery *string `json:"RemarksArrayAQuery,omitempty" xml:"RemarksArrayAQuery,omitempty"`
	RemarksArrayBQuery *string `json:"RemarksArrayBQuery,omitempty" xml:"RemarksArrayBQuery,omitempty"`
	ExternalId         *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
}

func (s ListFaceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsRequest) SetProject(v string) *ListFaceGroupsRequest {
	s.Project = &v
	return s
}

func (s *ListFaceGroupsRequest) SetSetId(v string) *ListFaceGroupsRequest {
	s.SetId = &v
	return s
}

func (s *ListFaceGroupsRequest) SetMarker(v string) *ListFaceGroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListFaceGroupsRequest) SetLimit(v int32) *ListFaceGroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListFaceGroupsRequest) SetOrder(v string) *ListFaceGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListFaceGroupsRequest) SetOrderBy(v string) *ListFaceGroupsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksAQuery(v string) *ListFaceGroupsRequest {
	s.RemarksAQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksBQuery(v string) *ListFaceGroupsRequest {
	s.RemarksBQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksCQuery(v string) *ListFaceGroupsRequest {
	s.RemarksCQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksDQuery(v string) *ListFaceGroupsRequest {
	s.RemarksDQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksArrayAQuery(v string) *ListFaceGroupsRequest {
	s.RemarksArrayAQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksArrayBQuery(v string) *ListFaceGroupsRequest {
	s.RemarksArrayBQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetExternalId(v string) *ListFaceGroupsRequest {
	s.ExternalId = &v
	return s
}

type ListFaceGroupsResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextMarker *string                                 `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	FaceGroups []*ListFaceGroupsResponseBodyFaceGroups `json:"FaceGroups,omitempty" xml:"FaceGroups,omitempty" type:"Repeated"`
}

func (s ListFaceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBody) SetRequestId(v string) *ListFaceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceGroupsResponseBody) SetNextMarker(v string) *ListFaceGroupsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListFaceGroupsResponseBody) SetFaceGroups(v []*ListFaceGroupsResponseBodyFaceGroups) *ListFaceGroupsResponseBody {
	s.FaceGroups = v
	return s
}

type ListFaceGroupsResponseBodyFaceGroups struct {
	Gender         *string                                             `json:"Gender,omitempty" xml:"Gender,omitempty"`
	CreateTime     *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RemarksC       *string                                             `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	GroupCoverFace *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace `json:"GroupCoverFace,omitempty" xml:"GroupCoverFace,omitempty" type:"Struct"`
	FaceCount      *int32                                              `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	RemarksArrayB  *string                                             `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksD       *string                                             `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	MaxAge         *float32                                            `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
	GroupId        *string                                             `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName      *string                                             `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RemarksA       *string                                             `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	AverageAge     *float32                                            `json:"AverageAge,omitempty" xml:"AverageAge,omitempty"`
	RemarksArrayA  *string                                             `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	MinAge         *float32                                            `json:"MinAge,omitempty" xml:"MinAge,omitempty"`
	ImageCount     *int32                                              `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	ExternalId     *string                                             `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	RemarksB       *string                                             `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	ModifyTime     *string                                             `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s ListFaceGroupsResponseBodyFaceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBodyFaceGroups) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetGender(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.Gender = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetCreateTime(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.CreateTime = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksC(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksC = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetGroupCoverFace(v *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) *ListFaceGroupsResponseBodyFaceGroups {
	s.GroupCoverFace = v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetFaceCount(v int32) *ListFaceGroupsResponseBodyFaceGroups {
	s.FaceCount = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksArrayB(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksArrayB = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksD(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksD = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetMaxAge(v float32) *ListFaceGroupsResponseBodyFaceGroups {
	s.MaxAge = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetGroupId(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.GroupId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetGroupName(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.GroupName = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksA(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksA = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetAverageAge(v float32) *ListFaceGroupsResponseBodyFaceGroups {
	s.AverageAge = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksArrayA(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksArrayA = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetMinAge(v float32) *ListFaceGroupsResponseBodyFaceGroups {
	s.MinAge = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetImageCount(v int32) *ListFaceGroupsResponseBodyFaceGroups {
	s.ImageCount = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetExternalId(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.ExternalId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksB(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksB = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetModifyTime(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.ModifyTime = &v
	return s
}

type ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace struct {
	FaceId       *string                                                         `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	ImageUri     *string                                                         `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	FaceBoundary *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	ExternalId   *string                                                         `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ImageHeight  *int64                                                          `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageWidth   *int64                                                          `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
}

func (s ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetFaceId(v string) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.FaceId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetImageUri(v string) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.ImageUri = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetFaceBoundary(v *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.FaceBoundary = v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetExternalId(v string) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.ExternalId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetImageHeight(v int64) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.ImageHeight = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetImageWidth(v int64) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.ImageWidth = &v
	return s
}

type ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) SetTop(v int32) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary {
	s.Top = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) SetWidth(v int32) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary {
	s.Width = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) SetHeight(v int32) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary {
	s.Height = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) SetLeft(v int32) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary {
	s.Left = &v
	return s
}

type ListFaceGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFaceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponse) SetHeaders(v map[string]*string) *ListFaceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListFaceGroupsResponse) SetBody(v *ListFaceGroupsResponseBody) *ListFaceGroupsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	Marker          *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetProject(v string) *ListImagesRequest {
	s.Project = &v
	return s
}

func (s *ListImagesRequest) SetSetId(v string) *ListImagesRequest {
	s.SetId = &v
	return s
}

func (s *ListImagesRequest) SetCreateTimeStart(v string) *ListImagesRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListImagesRequest) SetMarker(v string) *ListImagesRequest {
	s.Marker = &v
	return s
}

func (s *ListImagesRequest) SetLimit(v int32) *ListImagesRequest {
	s.Limit = &v
	return s
}

type ListImagesResponseBody struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextMarker *string                         `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	SetId      *string                         `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Images     []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetNextMarker(v string) *ListImagesResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListImagesResponseBody) SetSetId(v string) *ListImagesResponseBody {
	s.SetId = &v
	return s
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

type ListImagesResponseBodyImages struct {
	CroppingSuggestionStatus     *string                                           `json:"CroppingSuggestionStatus,omitempty" xml:"CroppingSuggestionStatus,omitempty"`
	ImageQualityModifyTime       *string                                           `json:"ImageQualityModifyTime,omitempty" xml:"ImageQualityModifyTime,omitempty"`
	TagsFailReason               *string                                           `json:"TagsFailReason,omitempty" xml:"TagsFailReason,omitempty"`
	RemarksC                     *string                                           `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	CreateTime                   *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SourceType                   *string                                           `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	FacesFailReason              *string                                           `json:"FacesFailReason,omitempty" xml:"FacesFailReason,omitempty"`
	FacesModifyTime              *string                                           `json:"FacesModifyTime,omitempty" xml:"FacesModifyTime,omitempty"`
	ImageTime                    *string                                           `json:"ImageTime,omitempty" xml:"ImageTime,omitempty"`
	OCRModifyTime                *string                                           `json:"OCRModifyTime,omitempty" xml:"OCRModifyTime,omitempty"`
	AddressModifyTime            *string                                           `json:"AddressModifyTime,omitempty" xml:"AddressModifyTime,omitempty"`
	ImageQualityFailReason       *string                                           `json:"ImageQualityFailReason,omitempty" xml:"ImageQualityFailReason,omitempty"`
	FacesStatus                  *string                                           `json:"FacesStatus,omitempty" xml:"FacesStatus,omitempty"`
	RemarksArrayA                *string                                           `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	ImageHeight                  *int32                                            `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ExternalId                   *string                                           `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	SourceUri                    *string                                           `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	FileSize                     *int32                                            `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	ModifyTime                   *string                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SourcePosition               *string                                           `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	ImageQualityStatus           *string                                           `json:"ImageQualityStatus,omitempty" xml:"ImageQualityStatus,omitempty"`
	OCRFailReason                *string                                           `json:"OCRFailReason,omitempty" xml:"OCRFailReason,omitempty"`
	AddressFailReason            *string                                           `json:"AddressFailReason,omitempty" xml:"AddressFailReason,omitempty"`
	CroppingSuggestionModifyTime *string                                           `json:"CroppingSuggestionModifyTime,omitempty" xml:"CroppingSuggestionModifyTime,omitempty"`
	ImageFormat                  *string                                           `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	ImageWidth                   *int32                                            `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	RemarksArrayB                *string                                           `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	Orientation                  *string                                           `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	RemarksD                     *string                                           `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	TagsStatus                   *string                                           `json:"TagsStatus,omitempty" xml:"TagsStatus,omitempty"`
	CroppingSuggestionFailReason *string                                           `json:"CroppingSuggestionFailReason,omitempty" xml:"CroppingSuggestionFailReason,omitempty"`
	RemarksA                     *string                                           `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	ImageUri                     *string                                           `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	TagsModifyTime               *string                                           `json:"TagsModifyTime,omitempty" xml:"TagsModifyTime,omitempty"`
	OCRStatus                    *string                                           `json:"OCRStatus,omitempty" xml:"OCRStatus,omitempty"`
	AddressStatus                *string                                           `json:"AddressStatus,omitempty" xml:"AddressStatus,omitempty"`
	Exif                         *string                                           `json:"Exif,omitempty" xml:"Exif,omitempty"`
	Location                     *string                                           `json:"Location,omitempty" xml:"Location,omitempty"`
	RemarksB                     *string                                           `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	CroppingSuggestion           []*ListImagesResponseBodyImagesCroppingSuggestion `json:"CroppingSuggestion,omitempty" xml:"CroppingSuggestion,omitempty" type:"Repeated"`
	Faces                        []*ListImagesResponseBodyImagesFaces              `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	Tags                         []*ListImagesResponseBodyImagesTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	OCR                          []*ListImagesResponseBodyImagesOCR                `json:"OCR,omitempty" xml:"OCR,omitempty" type:"Repeated"`
	ImageQuality                 *ListImagesResponseBodyImagesImageQuality         `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty" type:"Struct"`
	Address                      *ListImagesResponseBodyImagesAddress              `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetCroppingSuggestionStatus(v string) *ListImagesResponseBodyImages {
	s.CroppingSuggestionStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageQualityModifyTime(v string) *ListImagesResponseBodyImages {
	s.ImageQualityModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetTagsFailReason(v string) *ListImagesResponseBodyImages {
	s.TagsFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksC(v string) *ListImagesResponseBodyImages {
	s.RemarksC = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCreateTime(v string) *ListImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceType(v string) *ListImagesResponseBodyImages {
	s.SourceType = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetFacesFailReason(v string) *ListImagesResponseBodyImages {
	s.FacesFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetFacesModifyTime(v string) *ListImagesResponseBodyImages {
	s.FacesModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageTime(v string) *ListImagesResponseBodyImages {
	s.ImageTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOCRModifyTime(v string) *ListImagesResponseBodyImages {
	s.OCRModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetAddressModifyTime(v string) *ListImagesResponseBodyImages {
	s.AddressModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageQualityFailReason(v string) *ListImagesResponseBodyImages {
	s.ImageQualityFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetFacesStatus(v string) *ListImagesResponseBodyImages {
	s.FacesStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksArrayA(v string) *ListImagesResponseBodyImages {
	s.RemarksArrayA = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageHeight(v int32) *ListImagesResponseBodyImages {
	s.ImageHeight = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetExternalId(v string) *ListImagesResponseBodyImages {
	s.ExternalId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceUri(v string) *ListImagesResponseBodyImages {
	s.SourceUri = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetFileSize(v int32) *ListImagesResponseBodyImages {
	s.FileSize = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetModifyTime(v string) *ListImagesResponseBodyImages {
	s.ModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourcePosition(v string) *ListImagesResponseBodyImages {
	s.SourcePosition = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageQualityStatus(v string) *ListImagesResponseBodyImages {
	s.ImageQualityStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOCRFailReason(v string) *ListImagesResponseBodyImages {
	s.OCRFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetAddressFailReason(v string) *ListImagesResponseBodyImages {
	s.AddressFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCroppingSuggestionModifyTime(v string) *ListImagesResponseBodyImages {
	s.CroppingSuggestionModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageFormat(v string) *ListImagesResponseBodyImages {
	s.ImageFormat = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageWidth(v int32) *ListImagesResponseBodyImages {
	s.ImageWidth = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksArrayB(v string) *ListImagesResponseBodyImages {
	s.RemarksArrayB = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOrientation(v string) *ListImagesResponseBodyImages {
	s.Orientation = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksD(v string) *ListImagesResponseBodyImages {
	s.RemarksD = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetTagsStatus(v string) *ListImagesResponseBodyImages {
	s.TagsStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCroppingSuggestionFailReason(v string) *ListImagesResponseBodyImages {
	s.CroppingSuggestionFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksA(v string) *ListImagesResponseBodyImages {
	s.RemarksA = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageUri(v string) *ListImagesResponseBodyImages {
	s.ImageUri = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetTagsModifyTime(v string) *ListImagesResponseBodyImages {
	s.TagsModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOCRStatus(v string) *ListImagesResponseBodyImages {
	s.OCRStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetAddressStatus(v string) *ListImagesResponseBodyImages {
	s.AddressStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetExif(v string) *ListImagesResponseBodyImages {
	s.Exif = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetLocation(v string) *ListImagesResponseBodyImages {
	s.Location = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksB(v string) *ListImagesResponseBodyImages {
	s.RemarksB = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCroppingSuggestion(v []*ListImagesResponseBodyImagesCroppingSuggestion) *ListImagesResponseBodyImages {
	s.CroppingSuggestion = v
	return s
}

func (s *ListImagesResponseBodyImages) SetFaces(v []*ListImagesResponseBodyImagesFaces) *ListImagesResponseBodyImages {
	s.Faces = v
	return s
}

func (s *ListImagesResponseBodyImages) SetTags(v []*ListImagesResponseBodyImagesTags) *ListImagesResponseBodyImages {
	s.Tags = v
	return s
}

func (s *ListImagesResponseBodyImages) SetOCR(v []*ListImagesResponseBodyImagesOCR) *ListImagesResponseBodyImages {
	s.OCR = v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageQuality(v *ListImagesResponseBodyImagesImageQuality) *ListImagesResponseBodyImages {
	s.ImageQuality = v
	return s
}

func (s *ListImagesResponseBodyImages) SetAddress(v *ListImagesResponseBodyImagesAddress) *ListImagesResponseBodyImages {
	s.Address = v
	return s
}

type ListImagesResponseBodyImagesCroppingSuggestion struct {
	Score            *float32                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	AspectRatio      *string                                                         `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	CroppingBoundary *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary `json:"CroppingBoundary,omitempty" xml:"CroppingBoundary,omitempty" type:"Struct"`
}

func (s ListImagesResponseBodyImagesCroppingSuggestion) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesCroppingSuggestion) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesCroppingSuggestion) SetScore(v float32) *ListImagesResponseBodyImagesCroppingSuggestion {
	s.Score = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestion) SetAspectRatio(v string) *ListImagesResponseBodyImagesCroppingSuggestion {
	s.AspectRatio = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestion) SetCroppingBoundary(v *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) *ListImagesResponseBodyImagesCroppingSuggestion {
	s.CroppingBoundary = v
	return s
}

type ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetLeft(v int32) *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Left = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetTop(v int32) *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Top = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetWidth(v int32) *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Width = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetHeight(v int32) *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Height = &v
	return s
}

type ListImagesResponseBodyImagesFaces struct {
	EmotionConfidence *float32                                         `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	Attractive        *float32                                         `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	GroupId           *string                                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Gender            *string                                          `json:"Gender,omitempty" xml:"Gender,omitempty"`
	FaceId            *string                                          `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	GenderConfidence  *float32                                         `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	FaceQuality       *float32                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Emotion           *string                                          `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	Age               *int32                                           `json:"Age,omitempty" xml:"Age,omitempty"`
	FaceConfidence    *float32                                         `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	EmotionDetails    *ListImagesResponseBodyImagesFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
	FaceAttributes    *ListImagesResponseBodyImagesFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
}

func (s ListImagesResponseBodyImagesFaces) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFaces) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFaces) SetEmotionConfidence(v float32) *ListImagesResponseBodyImagesFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetAttractive(v float32) *ListImagesResponseBodyImagesFaces {
	s.Attractive = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetGroupId(v string) *ListImagesResponseBodyImagesFaces {
	s.GroupId = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetGender(v string) *ListImagesResponseBodyImagesFaces {
	s.Gender = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetFaceId(v string) *ListImagesResponseBodyImagesFaces {
	s.FaceId = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetGenderConfidence(v float32) *ListImagesResponseBodyImagesFaces {
	s.GenderConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetFaceQuality(v float32) *ListImagesResponseBodyImagesFaces {
	s.FaceQuality = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetEmotion(v string) *ListImagesResponseBodyImagesFaces {
	s.Emotion = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetAge(v int32) *ListImagesResponseBodyImagesFaces {
	s.Age = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetFaceConfidence(v float32) *ListImagesResponseBodyImagesFaces {
	s.FaceConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetEmotionDetails(v *ListImagesResponseBodyImagesFacesEmotionDetails) *ListImagesResponseBodyImagesFaces {
	s.EmotionDetails = v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetFaceAttributes(v *ListImagesResponseBodyImagesFacesFaceAttributes) *ListImagesResponseBodyImagesFaces {
	s.FaceAttributes = v
	return s
}

type ListImagesResponseBodyImagesFacesEmotionDetails struct {
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
}

func (s ListImagesResponseBodyImagesFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetHAPPY(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetSURPRISED(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetCALM(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetDISGUSTED(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetANGRY(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetSAD(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetSCARED(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.SCARED = &v
	return s
}

type ListImagesResponseBodyImagesFacesFaceAttributes struct {
	GlassesConfidence *float32                                                     `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	Glasses           *string                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Mask              *string                                                      `json:"Mask,omitempty" xml:"Mask,omitempty"`
	BeardConfidence   *float32                                                     `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	MaskConfidence    *float32                                                     `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	Beard             *string                                                      `json:"Beard,omitempty" xml:"Beard,omitempty"`
	FaceBoundary      *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	HeadPose          *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
}

func (s ListImagesResponseBodyImagesFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetGlassesConfidence(v float32) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetGlasses(v string) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetMask(v string) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetBeardConfidence(v float32) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetMaskConfidence(v float32) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetBeard(v string) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.Beard = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetFaceBoundary(v *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetHeadPose(v *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.HeadPose = v
	return s
}

type ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetLeft(v int32) *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetTop(v int32) *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetWidth(v int32) *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetHeight(v int32) *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

type ListImagesResponseBodyImagesFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetPitch(v float32) *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetRoll(v float32) *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetYaw(v float32) *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type ListImagesResponseBodyImagesTags struct {
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListImagesResponseBodyImagesTags) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesTags) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesTags) SetTagLevel(v int32) *ListImagesResponseBodyImagesTags {
	s.TagLevel = &v
	return s
}

func (s *ListImagesResponseBodyImagesTags) SetParentTagName(v string) *ListImagesResponseBodyImagesTags {
	s.ParentTagName = &v
	return s
}

func (s *ListImagesResponseBodyImagesTags) SetTagConfidence(v float32) *ListImagesResponseBodyImagesTags {
	s.TagConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesTags) SetTagName(v string) *ListImagesResponseBodyImagesTags {
	s.TagName = &v
	return s
}

type ListImagesResponseBodyImagesOCR struct {
	OCRConfidence *float32                                    `json:"OCRConfidence,omitempty" xml:"OCRConfidence,omitempty"`
	OCRContents   *string                                     `json:"OCRContents,omitempty" xml:"OCRContents,omitempty"`
	OCRBoundary   *ListImagesResponseBodyImagesOCROCRBoundary `json:"OCRBoundary,omitempty" xml:"OCRBoundary,omitempty" type:"Struct"`
}

func (s ListImagesResponseBodyImagesOCR) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesOCR) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesOCR) SetOCRConfidence(v float32) *ListImagesResponseBodyImagesOCR {
	s.OCRConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCR) SetOCRContents(v string) *ListImagesResponseBodyImagesOCR {
	s.OCRContents = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCR) SetOCRBoundary(v *ListImagesResponseBodyImagesOCROCRBoundary) *ListImagesResponseBodyImagesOCR {
	s.OCRBoundary = v
	return s
}

type ListImagesResponseBodyImagesOCROCRBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListImagesResponseBodyImagesOCROCRBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesOCROCRBoundary) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesOCROCRBoundary) SetLeft(v int32) *ListImagesResponseBodyImagesOCROCRBoundary {
	s.Left = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCROCRBoundary) SetTop(v int32) *ListImagesResponseBodyImagesOCROCRBoundary {
	s.Top = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCROCRBoundary) SetWidth(v int32) *ListImagesResponseBodyImagesOCROCRBoundary {
	s.Width = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCROCRBoundary) SetHeight(v int32) *ListImagesResponseBodyImagesOCROCRBoundary {
	s.Height = &v
	return s
}

type ListImagesResponseBodyImagesImageQuality struct {
	OverallScore     *float32 `json:"OverallScore,omitempty" xml:"OverallScore,omitempty"`
	Color            *float32 `json:"Color,omitempty" xml:"Color,omitempty"`
	ColorScore       *float32 `json:"ColorScore,omitempty" xml:"ColorScore,omitempty"`
	ContrastScore    *float32 `json:"ContrastScore,omitempty" xml:"ContrastScore,omitempty"`
	Contrast         *float32 `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	ExposureScore    *float32 `json:"ExposureScore,omitempty" xml:"ExposureScore,omitempty"`
	ClarityScore     *float32 `json:"ClarityScore,omitempty" xml:"ClarityScore,omitempty"`
	Clarity          *float32 `json:"Clarity,omitempty" xml:"Clarity,omitempty"`
	Exposure         *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
	CompositionScore *float32 `json:"CompositionScore,omitempty" xml:"CompositionScore,omitempty"`
}

func (s ListImagesResponseBodyImagesImageQuality) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesImageQuality) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesImageQuality) SetOverallScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.OverallScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetColor(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.Color = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetColorScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.ColorScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetContrastScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.ContrastScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetContrast(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.Contrast = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetExposureScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.ExposureScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetClarityScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.ClarityScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetClarity(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.Clarity = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetExposure(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.Exposure = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetCompositionScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.CompositionScore = &v
	return s
}

type ListImagesResponseBodyImagesAddress struct {
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListImagesResponseBodyImagesAddress) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesAddress) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesAddress) SetTownship(v string) *ListImagesResponseBodyImagesAddress {
	s.Township = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetDistrict(v string) *ListImagesResponseBodyImagesAddress {
	s.District = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetAddressLine(v string) *ListImagesResponseBodyImagesAddress {
	s.AddressLine = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetCountry(v string) *ListImagesResponseBodyImagesAddress {
	s.Country = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetCity(v string) *ListImagesResponseBodyImagesAddress {
	s.City = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetProvince(v string) *ListImagesResponseBodyImagesAddress {
	s.Province = &v
	return s
}

type ListImagesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListOfficeConversionTaskRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Marker  *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys *int32  `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
}

func (s ListOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeConversionTaskRequest) SetProject(v string) *ListOfficeConversionTaskRequest {
	s.Project = &v
	return s
}

func (s *ListOfficeConversionTaskRequest) SetMarker(v string) *ListOfficeConversionTaskRequest {
	s.Marker = &v
	return s
}

func (s *ListOfficeConversionTaskRequest) SetMaxKeys(v int32) *ListOfficeConversionTaskRequest {
	s.MaxKeys = &v
	return s
}

type ListOfficeConversionTaskResponseBody struct {
	NextMarker *string                                      `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*ListOfficeConversionTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfficeConversionTaskResponseBody) SetNextMarker(v string) *ListOfficeConversionTaskResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBody) SetRequestId(v string) *ListOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBody) SetTasks(v []*ListOfficeConversionTaskResponseBodyTasks) *ListOfficeConversionTaskResponseBody {
	s.Tasks = v
	return s
}

type ListOfficeConversionTaskResponseBodyTasks struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent         *int32  `json:"Percent,omitempty" xml:"Percent,omitempty"`
	FinishTime      *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PageCount       *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	SrcUri          *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	TgtType         *string `json:"TgtType,omitempty" xml:"TgtType,omitempty"`
	TgtUri          *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	ImageSpec       *string `json:"ImageSpec,omitempty" xml:"ImageSpec,omitempty"`
	ExternalID      *string `json:"ExternalID,omitempty" xml:"ExternalID,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListOfficeConversionTaskResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeConversionTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetStatus(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetPercent(v int32) *ListOfficeConversionTaskResponseBodyTasks {
	s.Percent = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetFinishTime(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.FinishTime = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetCreateTime(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetPageCount(v int32) *ListOfficeConversionTaskResponseBodyTasks {
	s.PageCount = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetNotifyTopicName(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.NotifyTopicName = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetNotifyEndpoint(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.NotifyEndpoint = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetSrcUri(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.SrcUri = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetTgtType(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.TgtType = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetTgtUri(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.TgtUri = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetImageSpec(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.ImageSpec = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetExternalID(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.ExternalID = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetTaskId(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

type ListOfficeConversionTaskResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *ListOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *ListOfficeConversionTaskResponse) SetBody(v *ListOfficeConversionTaskResponseBody) *ListOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type ListProjectAPIsRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s ListProjectAPIsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAPIsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectAPIsRequest) SetProject(v string) *ListProjectAPIsRequest {
	s.Project = &v
	return s
}

type ListProjectAPIsResponseBody struct {
	Project   *string   `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	APIs      []*string `json:"APIs,omitempty" xml:"APIs,omitempty" type:"Repeated"`
}

func (s ListProjectAPIsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAPIsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectAPIsResponseBody) SetProject(v string) *ListProjectAPIsResponseBody {
	s.Project = &v
	return s
}

func (s *ListProjectAPIsResponseBody) SetRequestId(v string) *ListProjectAPIsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectAPIsResponseBody) SetAPIs(v []*string) *ListProjectAPIsResponseBody {
	s.APIs = v
	return s
}

type ListProjectAPIsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectAPIsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectAPIsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAPIsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectAPIsResponse) SetHeaders(v map[string]*string) *ListProjectAPIsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectAPIsResponse) SetBody(v *ListProjectAPIsResponseBody) *ListProjectAPIsResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	Marker  *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys *int32  `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetMarker(v string) *ListProjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListProjectsRequest) SetMaxKeys(v int32) *ListProjectsRequest {
	s.MaxKeys = &v
	return s
}

type ListProjectsResponseBody struct {
	NextMarker *string                             `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Projects   []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetNextMarker(v string) *ListProjectsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

type ListProjectsResponseBodyProjects struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CU          *int32  `json:"CU,omitempty" xml:"CU,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetType(v string) *ListProjectsResponseBodyProjects {
	s.Type = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCU(v int32) *ListProjectsResponseBodyProjects {
	s.CU = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCreateTime(v string) *ListProjectsResponseBodyProjects {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetServiceRole(v string) *ListProjectsResponseBodyProjects {
	s.ServiceRole = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetEndpoint(v string) *ListProjectsResponseBodyProjects {
	s.Endpoint = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProject(v string) *ListProjectsResponseBodyProjects {
	s.Project = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetRegionId(v string) *ListProjectsResponseBodyProjects {
	s.RegionId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetBillingType(v string) *ListProjectsResponseBodyProjects {
	s.BillingType = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetModifyTime(v string) *ListProjectsResponseBodyProjects {
	s.ModifyTime = &v
	return s
}

type ListProjectsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListSetsRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Marker  *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
}

func (s ListSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSetsRequest) GoString() string {
	return s.String()
}

func (s *ListSetsRequest) SetProject(v string) *ListSetsRequest {
	s.Project = &v
	return s
}

func (s *ListSetsRequest) SetMarker(v string) *ListSetsRequest {
	s.Marker = &v
	return s
}

type ListSetsResponseBody struct {
	NextMarker *string                     `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sets       []*ListSetsResponseBodySets `json:"Sets,omitempty" xml:"Sets,omitempty" type:"Repeated"`
}

func (s ListSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSetsResponseBody) SetNextMarker(v string) *ListSetsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListSetsResponseBody) SetRequestId(v string) *ListSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSetsResponseBody) SetSets(v []*ListSetsResponseBodySets) *ListSetsResponseBody {
	s.Sets = v
	return s
}

type ListSetsResponseBodySets struct {
	VideoCount  *int32  `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VideoLength *int32  `json:"VideoLength,omitempty" xml:"VideoLength,omitempty"`
	SetId       *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageCount  *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	FaceCount   *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	SetName     *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s ListSetsResponseBodySets) String() string {
	return tea.Prettify(s)
}

func (s ListSetsResponseBodySets) GoString() string {
	return s.String()
}

func (s *ListSetsResponseBodySets) SetVideoCount(v int32) *ListSetsResponseBodySets {
	s.VideoCount = &v
	return s
}

func (s *ListSetsResponseBodySets) SetCreateTime(v string) *ListSetsResponseBodySets {
	s.CreateTime = &v
	return s
}

func (s *ListSetsResponseBodySets) SetVideoLength(v int32) *ListSetsResponseBodySets {
	s.VideoLength = &v
	return s
}

func (s *ListSetsResponseBodySets) SetSetId(v string) *ListSetsResponseBodySets {
	s.SetId = &v
	return s
}

func (s *ListSetsResponseBodySets) SetImageCount(v int32) *ListSetsResponseBodySets {
	s.ImageCount = &v
	return s
}

func (s *ListSetsResponseBodySets) SetFaceCount(v int32) *ListSetsResponseBodySets {
	s.FaceCount = &v
	return s
}

func (s *ListSetsResponseBodySets) SetSetName(v string) *ListSetsResponseBodySets {
	s.SetName = &v
	return s
}

func (s *ListSetsResponseBodySets) SetModifyTime(v string) *ListSetsResponseBodySets {
	s.ModifyTime = &v
	return s
}

type ListSetsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSetsResponse) GoString() string {
	return s.String()
}

func (s *ListSetsResponse) SetHeaders(v map[string]*string) *ListSetsResponse {
	s.Headers = v
	return s
}

func (s *ListSetsResponse) SetBody(v *ListSetsResponseBody) *ListSetsResponse {
	s.Body = v
	return s
}

type ListSetTagsRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s ListSetTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSetTagsRequest) GoString() string {
	return s.String()
}

func (s *ListSetTagsRequest) SetProject(v string) *ListSetTagsRequest {
	s.Project = &v
	return s
}

func (s *ListSetTagsRequest) SetSetId(v string) *ListSetTagsRequest {
	s.SetId = &v
	return s
}

type ListSetTagsResponseBody struct {
	SetId     *string                        `json:"SetId,omitempty" xml:"SetId,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListSetTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListSetTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSetTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSetTagsResponseBody) SetSetId(v string) *ListSetTagsResponseBody {
	s.SetId = &v
	return s
}

func (s *ListSetTagsResponseBody) SetRequestId(v string) *ListSetTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSetTagsResponseBody) SetTags(v []*ListSetTagsResponseBodyTags) *ListSetTagsResponseBody {
	s.Tags = v
	return s
}

type ListSetTagsResponseBodyTags struct {
	TagLevel *int32  `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	TagName  *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
}

func (s ListSetTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListSetTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListSetTagsResponseBodyTags) SetTagLevel(v int32) *ListSetTagsResponseBodyTags {
	s.TagLevel = &v
	return s
}

func (s *ListSetTagsResponseBodyTags) SetTagName(v string) *ListSetTagsResponseBodyTags {
	s.TagName = &v
	return s
}

func (s *ListSetTagsResponseBodyTags) SetTagCount(v int32) *ListSetTagsResponseBodyTags {
	s.TagCount = &v
	return s
}

type ListSetTagsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSetTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSetTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSetTagsResponse) GoString() string {
	return s.String()
}

func (s *ListSetTagsResponse) SetHeaders(v map[string]*string) *ListSetTagsResponse {
	s.Headers = v
	return s
}

func (s *ListSetTagsResponse) SetBody(v *ListSetTagsResponseBody) *ListSetTagsResponse {
	s.Body = v
	return s
}

type ListVideoAudiosRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
}

func (s ListVideoAudiosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVideoAudiosRequest) GoString() string {
	return s.String()
}

func (s *ListVideoAudiosRequest) SetProject(v string) *ListVideoAudiosRequest {
	s.Project = &v
	return s
}

func (s *ListVideoAudiosRequest) SetSetId(v string) *ListVideoAudiosRequest {
	s.SetId = &v
	return s
}

func (s *ListVideoAudiosRequest) SetVideoUri(v string) *ListVideoAudiosRequest {
	s.VideoUri = &v
	return s
}

func (s *ListVideoAudiosRequest) SetMarker(v string) *ListVideoAudiosRequest {
	s.Marker = &v
	return s
}

type ListVideoAudiosResponseBody struct {
	VideoUri   *string                              `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextMarker *string                              `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	SetId      *string                              `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Audios     []*ListVideoAudiosResponseBodyAudios `json:"Audios,omitempty" xml:"Audios,omitempty" type:"Repeated"`
}

func (s ListVideoAudiosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVideoAudiosResponseBody) GoString() string {
	return s.String()
}

func (s *ListVideoAudiosResponseBody) SetVideoUri(v string) *ListVideoAudiosResponseBody {
	s.VideoUri = &v
	return s
}

func (s *ListVideoAudiosResponseBody) SetRequestId(v string) *ListVideoAudiosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVideoAudiosResponseBody) SetNextMarker(v string) *ListVideoAudiosResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListVideoAudiosResponseBody) SetSetId(v string) *ListVideoAudiosResponseBody {
	s.SetId = &v
	return s
}

func (s *ListVideoAudiosResponseBody) SetAudios(v []*ListVideoAudiosResponseBodyAudios) *ListVideoAudiosResponseBody {
	s.Audios = v
	return s
}

type ListVideoAudiosResponseBodyAudios struct {
	SourcePosition       *string                                        `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	RemarksC             *string                                        `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	CreateTime           *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SourceType           *string                                        `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	AudioDuration        *float32                                       `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	AudioTextsStatus     *string                                        `json:"AudioTextsStatus,omitempty" xml:"AudioTextsStatus,omitempty"`
	AudioFormat          *string                                        `json:"AudioFormat,omitempty" xml:"AudioFormat,omitempty"`
	RemarksD             *string                                        `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	ProcessFailReason    *string                                        `json:"ProcessFailReason,omitempty" xml:"ProcessFailReason,omitempty"`
	ProcessModifyTime    *string                                        `json:"ProcessModifyTime,omitempty" xml:"ProcessModifyTime,omitempty"`
	AudioRate            *int32                                         `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	AudioUri             *string                                        `json:"AudioUri,omitempty" xml:"AudioUri,omitempty"`
	AudioTextsModifyTime *string                                        `json:"AudioTextsModifyTime,omitempty" xml:"AudioTextsModifyTime,omitempty"`
	RemarksA             *string                                        `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	ExternalId           *string                                        `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	SourceUri            *string                                        `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	ProcessStatus        *string                                        `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	AudioTextsFailReason *string                                        `json:"AudioTextsFailReason,omitempty" xml:"AudioTextsFailReason,omitempty"`
	RemarksB             *string                                        `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	FileSize             *int32                                         `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	ModifyTime           *string                                        `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	AudioTexts           []*ListVideoAudiosResponseBodyAudiosAudioTexts `json:"AudioTexts,omitempty" xml:"AudioTexts,omitempty" type:"Repeated"`
}

func (s ListVideoAudiosResponseBodyAudios) String() string {
	return tea.Prettify(s)
}

func (s ListVideoAudiosResponseBodyAudios) GoString() string {
	return s.String()
}

func (s *ListVideoAudiosResponseBodyAudios) SetSourcePosition(v string) *ListVideoAudiosResponseBodyAudios {
	s.SourcePosition = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetRemarksC(v string) *ListVideoAudiosResponseBodyAudios {
	s.RemarksC = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetCreateTime(v string) *ListVideoAudiosResponseBodyAudios {
	s.CreateTime = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetSourceType(v string) *ListVideoAudiosResponseBodyAudios {
	s.SourceType = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetAudioDuration(v float32) *ListVideoAudiosResponseBodyAudios {
	s.AudioDuration = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetAudioTextsStatus(v string) *ListVideoAudiosResponseBodyAudios {
	s.AudioTextsStatus = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetAudioFormat(v string) *ListVideoAudiosResponseBodyAudios {
	s.AudioFormat = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetRemarksD(v string) *ListVideoAudiosResponseBodyAudios {
	s.RemarksD = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetProcessFailReason(v string) *ListVideoAudiosResponseBodyAudios {
	s.ProcessFailReason = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetProcessModifyTime(v string) *ListVideoAudiosResponseBodyAudios {
	s.ProcessModifyTime = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetAudioRate(v int32) *ListVideoAudiosResponseBodyAudios {
	s.AudioRate = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetAudioUri(v string) *ListVideoAudiosResponseBodyAudios {
	s.AudioUri = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetAudioTextsModifyTime(v string) *ListVideoAudiosResponseBodyAudios {
	s.AudioTextsModifyTime = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetRemarksA(v string) *ListVideoAudiosResponseBodyAudios {
	s.RemarksA = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetExternalId(v string) *ListVideoAudiosResponseBodyAudios {
	s.ExternalId = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetSourceUri(v string) *ListVideoAudiosResponseBodyAudios {
	s.SourceUri = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetProcessStatus(v string) *ListVideoAudiosResponseBodyAudios {
	s.ProcessStatus = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetAudioTextsFailReason(v string) *ListVideoAudiosResponseBodyAudios {
	s.AudioTextsFailReason = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetRemarksB(v string) *ListVideoAudiosResponseBodyAudios {
	s.RemarksB = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetFileSize(v int32) *ListVideoAudiosResponseBodyAudios {
	s.FileSize = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetModifyTime(v string) *ListVideoAudiosResponseBodyAudios {
	s.ModifyTime = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudios) SetAudioTexts(v []*ListVideoAudiosResponseBodyAudiosAudioTexts) *ListVideoAudiosResponseBodyAudios {
	s.AudioTexts = v
	return s
}

type ListVideoAudiosResponseBodyAudiosAudioTexts struct {
	EndTime         *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Library         *string  `json:"Library,omitempty" xml:"Library,omitempty"`
	Confidence      *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	BeginTime       *float32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ChannelId       *int32   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EmotionValue    *float32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	SpeechRate      *int32   `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Text            *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	Person          *string  `json:"Person,omitempty" xml:"Person,omitempty"`
	SilenceDuration *float32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
}

func (s ListVideoAudiosResponseBodyAudiosAudioTexts) String() string {
	return tea.Prettify(s)
}

func (s ListVideoAudiosResponseBodyAudiosAudioTexts) GoString() string {
	return s.String()
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetEndTime(v float32) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.EndTime = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetLibrary(v string) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.Library = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetConfidence(v float32) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.Confidence = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetBeginTime(v float32) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.BeginTime = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetChannelId(v int32) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.ChannelId = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetEmotionValue(v float32) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.EmotionValue = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetSpeechRate(v int32) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.SpeechRate = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetText(v string) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.Text = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetPerson(v string) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.Person = &v
	return s
}

func (s *ListVideoAudiosResponseBodyAudiosAudioTexts) SetSilenceDuration(v float32) *ListVideoAudiosResponseBodyAudiosAudioTexts {
	s.SilenceDuration = &v
	return s
}

type ListVideoAudiosResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVideoAudiosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVideoAudiosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVideoAudiosResponse) GoString() string {
	return s.String()
}

func (s *ListVideoAudiosResponse) SetHeaders(v map[string]*string) *ListVideoAudiosResponse {
	s.Headers = v
	return s
}

func (s *ListVideoAudiosResponse) SetBody(v *ListVideoAudiosResponseBody) *ListVideoAudiosResponse {
	s.Body = v
	return s
}

type ListVideoFramesRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
}

func (s ListVideoFramesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesRequest) GoString() string {
	return s.String()
}

func (s *ListVideoFramesRequest) SetProject(v string) *ListVideoFramesRequest {
	s.Project = &v
	return s
}

func (s *ListVideoFramesRequest) SetSetId(v string) *ListVideoFramesRequest {
	s.SetId = &v
	return s
}

func (s *ListVideoFramesRequest) SetVideoUri(v string) *ListVideoFramesRequest {
	s.VideoUri = &v
	return s
}

func (s *ListVideoFramesRequest) SetMarker(v string) *ListVideoFramesRequest {
	s.Marker = &v
	return s
}

type ListVideoFramesResponseBody struct {
	VideoUri   *string                              `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextMarker *string                              `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	SetId      *string                              `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Frames     []*ListVideoFramesResponseBodyFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
}

func (s ListVideoFramesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBody) SetVideoUri(v string) *ListVideoFramesResponseBody {
	s.VideoUri = &v
	return s
}

func (s *ListVideoFramesResponseBody) SetRequestId(v string) *ListVideoFramesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVideoFramesResponseBody) SetNextMarker(v string) *ListVideoFramesResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListVideoFramesResponseBody) SetSetId(v string) *ListVideoFramesResponseBody {
	s.SetId = &v
	return s
}

func (s *ListVideoFramesResponseBody) SetFrames(v []*ListVideoFramesResponseBodyFrames) *ListVideoFramesResponseBody {
	s.Frames = v
	return s
}

type ListVideoFramesResponseBodyFrames struct {
	TagsFailReason  *string                                   `json:"TagsFailReason,omitempty" xml:"TagsFailReason,omitempty"`
	RemarksC        *string                                   `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	CreateTime      *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SourceType      *string                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	FacesFailReason *string                                   `json:"FacesFailReason,omitempty" xml:"FacesFailReason,omitempty"`
	FacesModifyTime *string                                   `json:"FacesModifyTime,omitempty" xml:"FacesModifyTime,omitempty"`
	ImageTime       *string                                   `json:"ImageTime,omitempty" xml:"ImageTime,omitempty"`
	OCRModifyTime   *string                                   `json:"OCRModifyTime,omitempty" xml:"OCRModifyTime,omitempty"`
	FacesStatus     *string                                   `json:"FacesStatus,omitempty" xml:"FacesStatus,omitempty"`
	ImageHeight     *int32                                    `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ExternalId      *string                                   `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	SourceUri       *string                                   `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	ModifyTime      *string                                   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	FileSize        *int32                                    `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	SourcePosition  *string                                   `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	OCRFailReason   *string                                   `json:"OCRFailReason,omitempty" xml:"OCRFailReason,omitempty"`
	ImageFormat     *string                                   `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	ImageWidth      *int32                                    `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	Orientation     *string                                   `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	RemarksD        *string                                   `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	TagsStatus      *string                                   `json:"TagsStatus,omitempty" xml:"TagsStatus,omitempty"`
	RemarksA        *string                                   `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	ImageUri        *string                                   `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	TagsModifyTime  *string                                   `json:"TagsModifyTime,omitempty" xml:"TagsModifyTime,omitempty"`
	OCRStatus       *string                                   `json:"OCRStatus,omitempty" xml:"OCRStatus,omitempty"`
	Exif            *string                                   `json:"Exif,omitempty" xml:"Exif,omitempty"`
	Location        *string                                   `json:"Location,omitempty" xml:"Location,omitempty"`
	RemarksB        *string                                   `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	Faces           []*ListVideoFramesResponseBodyFramesFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	Tags            []*ListVideoFramesResponseBodyFramesTags  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	OCR             []*ListVideoFramesResponseBodyFramesOCR   `json:"OCR,omitempty" xml:"OCR,omitempty" type:"Repeated"`
}

func (s ListVideoFramesResponseBodyFrames) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFrames) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFrames) SetTagsFailReason(v string) *ListVideoFramesResponseBodyFrames {
	s.TagsFailReason = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetRemarksC(v string) *ListVideoFramesResponseBodyFrames {
	s.RemarksC = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetCreateTime(v string) *ListVideoFramesResponseBodyFrames {
	s.CreateTime = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetSourceType(v string) *ListVideoFramesResponseBodyFrames {
	s.SourceType = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetFacesFailReason(v string) *ListVideoFramesResponseBodyFrames {
	s.FacesFailReason = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetFacesModifyTime(v string) *ListVideoFramesResponseBodyFrames {
	s.FacesModifyTime = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetImageTime(v string) *ListVideoFramesResponseBodyFrames {
	s.ImageTime = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetOCRModifyTime(v string) *ListVideoFramesResponseBodyFrames {
	s.OCRModifyTime = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetFacesStatus(v string) *ListVideoFramesResponseBodyFrames {
	s.FacesStatus = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetImageHeight(v int32) *ListVideoFramesResponseBodyFrames {
	s.ImageHeight = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetExternalId(v string) *ListVideoFramesResponseBodyFrames {
	s.ExternalId = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetSourceUri(v string) *ListVideoFramesResponseBodyFrames {
	s.SourceUri = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetModifyTime(v string) *ListVideoFramesResponseBodyFrames {
	s.ModifyTime = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetFileSize(v int32) *ListVideoFramesResponseBodyFrames {
	s.FileSize = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetSourcePosition(v string) *ListVideoFramesResponseBodyFrames {
	s.SourcePosition = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetOCRFailReason(v string) *ListVideoFramesResponseBodyFrames {
	s.OCRFailReason = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetImageFormat(v string) *ListVideoFramesResponseBodyFrames {
	s.ImageFormat = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetImageWidth(v int32) *ListVideoFramesResponseBodyFrames {
	s.ImageWidth = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetOrientation(v string) *ListVideoFramesResponseBodyFrames {
	s.Orientation = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetRemarksD(v string) *ListVideoFramesResponseBodyFrames {
	s.RemarksD = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetTagsStatus(v string) *ListVideoFramesResponseBodyFrames {
	s.TagsStatus = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetRemarksA(v string) *ListVideoFramesResponseBodyFrames {
	s.RemarksA = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetImageUri(v string) *ListVideoFramesResponseBodyFrames {
	s.ImageUri = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetTagsModifyTime(v string) *ListVideoFramesResponseBodyFrames {
	s.TagsModifyTime = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetOCRStatus(v string) *ListVideoFramesResponseBodyFrames {
	s.OCRStatus = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetExif(v string) *ListVideoFramesResponseBodyFrames {
	s.Exif = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetLocation(v string) *ListVideoFramesResponseBodyFrames {
	s.Location = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetRemarksB(v string) *ListVideoFramesResponseBodyFrames {
	s.RemarksB = &v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetFaces(v []*ListVideoFramesResponseBodyFramesFaces) *ListVideoFramesResponseBodyFrames {
	s.Faces = v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetTags(v []*ListVideoFramesResponseBodyFramesTags) *ListVideoFramesResponseBodyFrames {
	s.Tags = v
	return s
}

func (s *ListVideoFramesResponseBodyFrames) SetOCR(v []*ListVideoFramesResponseBodyFramesOCR) *ListVideoFramesResponseBodyFrames {
	s.OCR = v
	return s
}

type ListVideoFramesResponseBodyFramesFaces struct {
	EmotionConfidence *float32                                              `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	Attractive        *float32                                              `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	GroupId           *string                                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Gender            *string                                               `json:"Gender,omitempty" xml:"Gender,omitempty"`
	FaceId            *string                                               `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	GenderConfidence  *float32                                              `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	FaceQuality       *float32                                              `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Emotion           *string                                               `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	Age               *int32                                                `json:"Age,omitempty" xml:"Age,omitempty"`
	FaceConfidence    *float32                                              `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	EmotionDetails    *ListVideoFramesResponseBodyFramesFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
	FaceAttributes    *ListVideoFramesResponseBodyFramesFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
}

func (s ListVideoFramesResponseBodyFramesFaces) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFramesFaces) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetEmotionConfidence(v float32) *ListVideoFramesResponseBodyFramesFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetAttractive(v float32) *ListVideoFramesResponseBodyFramesFaces {
	s.Attractive = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetGroupId(v string) *ListVideoFramesResponseBodyFramesFaces {
	s.GroupId = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetGender(v string) *ListVideoFramesResponseBodyFramesFaces {
	s.Gender = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetFaceId(v string) *ListVideoFramesResponseBodyFramesFaces {
	s.FaceId = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetGenderConfidence(v float32) *ListVideoFramesResponseBodyFramesFaces {
	s.GenderConfidence = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetFaceQuality(v float32) *ListVideoFramesResponseBodyFramesFaces {
	s.FaceQuality = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetEmotion(v string) *ListVideoFramesResponseBodyFramesFaces {
	s.Emotion = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetAge(v int32) *ListVideoFramesResponseBodyFramesFaces {
	s.Age = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetFaceConfidence(v float32) *ListVideoFramesResponseBodyFramesFaces {
	s.FaceConfidence = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetEmotionDetails(v *ListVideoFramesResponseBodyFramesFacesEmotionDetails) *ListVideoFramesResponseBodyFramesFaces {
	s.EmotionDetails = v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFaces) SetFaceAttributes(v *ListVideoFramesResponseBodyFramesFacesFaceAttributes) *ListVideoFramesResponseBodyFramesFaces {
	s.FaceAttributes = v
	return s
}

type ListVideoFramesResponseBodyFramesFacesEmotionDetails struct {
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
}

func (s ListVideoFramesResponseBodyFramesFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFramesFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFramesFacesEmotionDetails) SetHAPPY(v float32) *ListVideoFramesResponseBodyFramesFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesEmotionDetails) SetSURPRISED(v float32) *ListVideoFramesResponseBodyFramesFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesEmotionDetails) SetCALM(v float32) *ListVideoFramesResponseBodyFramesFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesEmotionDetails) SetDISGUSTED(v float32) *ListVideoFramesResponseBodyFramesFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesEmotionDetails) SetANGRY(v float32) *ListVideoFramesResponseBodyFramesFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesEmotionDetails) SetSAD(v float32) *ListVideoFramesResponseBodyFramesFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesEmotionDetails) SetSCARED(v float32) *ListVideoFramesResponseBodyFramesFacesEmotionDetails {
	s.SCARED = &v
	return s
}

type ListVideoFramesResponseBodyFramesFacesFaceAttributes struct {
	GlassesConfidence *float32                                                          `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	Glasses           *string                                                           `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Mask              *string                                                           `json:"Mask,omitempty" xml:"Mask,omitempty"`
	BeardConfidence   *float32                                                          `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	MaskConfidence    *float32                                                          `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	Beard             *string                                                           `json:"Beard,omitempty" xml:"Beard,omitempty"`
	FaceBoundary      *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	HeadPose          *ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
}

func (s ListVideoFramesResponseBodyFramesFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFramesFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributes) SetGlassesConfidence(v float32) *ListVideoFramesResponseBodyFramesFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributes) SetGlasses(v string) *ListVideoFramesResponseBodyFramesFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributes) SetMask(v string) *ListVideoFramesResponseBodyFramesFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributes) SetBeardConfidence(v float32) *ListVideoFramesResponseBodyFramesFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributes) SetMaskConfidence(v float32) *ListVideoFramesResponseBodyFramesFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributes) SetBeard(v string) *ListVideoFramesResponseBodyFramesFacesFaceAttributes {
	s.Beard = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributes) SetFaceBoundary(v *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary) *ListVideoFramesResponseBodyFramesFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributes) SetHeadPose(v *ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose) *ListVideoFramesResponseBodyFramesFacesFaceAttributes {
	s.HeadPose = v
	return s
}

type ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary) SetLeft(v int32) *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary) SetTop(v int32) *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary) SetWidth(v int32) *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary) SetHeight(v int32) *ListVideoFramesResponseBodyFramesFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

type ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose) SetPitch(v float32) *ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose) SetRoll(v float32) *ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose) SetYaw(v float32) *ListVideoFramesResponseBodyFramesFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type ListVideoFramesResponseBodyFramesTags struct {
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListVideoFramesResponseBodyFramesTags) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFramesTags) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFramesTags) SetTagLevel(v int32) *ListVideoFramesResponseBodyFramesTags {
	s.TagLevel = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesTags) SetParentTagName(v string) *ListVideoFramesResponseBodyFramesTags {
	s.ParentTagName = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesTags) SetTagConfidence(v float32) *ListVideoFramesResponseBodyFramesTags {
	s.TagConfidence = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesTags) SetTagName(v string) *ListVideoFramesResponseBodyFramesTags {
	s.TagName = &v
	return s
}

type ListVideoFramesResponseBodyFramesOCR struct {
	OCRConfidence *float32                                         `json:"OCRConfidence,omitempty" xml:"OCRConfidence,omitempty"`
	OCRContents   *string                                          `json:"OCRContents,omitempty" xml:"OCRContents,omitempty"`
	OCRBoundary   *ListVideoFramesResponseBodyFramesOCROCRBoundary `json:"OCRBoundary,omitempty" xml:"OCRBoundary,omitempty" type:"Struct"`
}

func (s ListVideoFramesResponseBodyFramesOCR) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFramesOCR) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFramesOCR) SetOCRConfidence(v float32) *ListVideoFramesResponseBodyFramesOCR {
	s.OCRConfidence = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesOCR) SetOCRContents(v string) *ListVideoFramesResponseBodyFramesOCR {
	s.OCRContents = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesOCR) SetOCRBoundary(v *ListVideoFramesResponseBodyFramesOCROCRBoundary) *ListVideoFramesResponseBodyFramesOCR {
	s.OCRBoundary = v
	return s
}

type ListVideoFramesResponseBodyFramesOCROCRBoundary struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListVideoFramesResponseBodyFramesOCROCRBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponseBodyFramesOCROCRBoundary) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponseBodyFramesOCROCRBoundary) SetLeft(v int32) *ListVideoFramesResponseBodyFramesOCROCRBoundary {
	s.Left = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesOCROCRBoundary) SetTop(v int32) *ListVideoFramesResponseBodyFramesOCROCRBoundary {
	s.Top = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesOCROCRBoundary) SetWidth(v int32) *ListVideoFramesResponseBodyFramesOCROCRBoundary {
	s.Width = &v
	return s
}

func (s *ListVideoFramesResponseBodyFramesOCROCRBoundary) SetHeight(v int32) *ListVideoFramesResponseBodyFramesOCROCRBoundary {
	s.Height = &v
	return s
}

type ListVideoFramesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVideoFramesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVideoFramesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVideoFramesResponse) GoString() string {
	return s.String()
}

func (s *ListVideoFramesResponse) SetHeaders(v map[string]*string) *ListVideoFramesResponse {
	s.Headers = v
	return s
}

func (s *ListVideoFramesResponse) SetBody(v *ListVideoFramesResponseBody) *ListVideoFramesResponse {
	s.Body = v
	return s
}

type ListVideosRequest struct {
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	Marker          *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
}

func (s ListVideosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVideosRequest) GoString() string {
	return s.String()
}

func (s *ListVideosRequest) SetProject(v string) *ListVideosRequest {
	s.Project = &v
	return s
}

func (s *ListVideosRequest) SetSetId(v string) *ListVideosRequest {
	s.SetId = &v
	return s
}

func (s *ListVideosRequest) SetCreateTimeStart(v string) *ListVideosRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListVideosRequest) SetMarker(v string) *ListVideosRequest {
	s.Marker = &v
	return s
}

type ListVideosResponseBody struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextMarker *string                         `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	SetId      *string                         `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Videos     []*ListVideosResponseBodyVideos `json:"Videos,omitempty" xml:"Videos,omitempty" type:"Repeated"`
}

func (s ListVideosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVideosResponseBody) GoString() string {
	return s.String()
}

func (s *ListVideosResponseBody) SetRequestId(v string) *ListVideosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVideosResponseBody) SetNextMarker(v string) *ListVideosResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListVideosResponseBody) SetSetId(v string) *ListVideosResponseBody {
	s.SetId = &v
	return s
}

func (s *ListVideosResponseBody) SetVideos(v []*ListVideosResponseBodyVideos) *ListVideosResponseBody {
	s.Videos = v
	return s
}

type ListVideosResponseBodyVideos struct {
	CreateTime          *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RemarksC            *string                                  `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	VideoTagsFailReason *string                                  `json:"VideoTagsFailReason,omitempty" xml:"VideoTagsFailReason,omitempty"`
	SourceType          *string                                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VideoDuration       *float32                                 `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	ProcessModifyTime   *string                                  `json:"ProcessModifyTime,omitempty" xml:"ProcessModifyTime,omitempty"`
	VideoFrames         *int32                                   `json:"VideoFrames,omitempty" xml:"VideoFrames,omitempty"`
	VideoTagsStatus     *string                                  `json:"VideoTagsStatus,omitempty" xml:"VideoTagsStatus,omitempty"`
	ExternalId          *string                                  `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	SourceUri           *string                                  `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	ModifyTime          *string                                  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	FileSize            *int32                                   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	SourcePosition      *string                                  `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	VideoWidth          *int32                                   `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
	VideoHeight         *int32                                   `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoFormat         *string                                  `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	RemarksD            *string                                  `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	VideoUri            *string                                  `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	VideoTagsModifyTime *string                                  `json:"VideoTagsModifyTime,omitempty" xml:"VideoTagsModifyTime,omitempty"`
	ProcessFailReason   *string                                  `json:"ProcessFailReason,omitempty" xml:"ProcessFailReason,omitempty"`
	RemarksA            *string                                  `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	ProcessStatus       *string                                  `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	RemarksB            *string                                  `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	VideoTags           []*ListVideosResponseBodyVideosVideoTags `json:"VideoTags,omitempty" xml:"VideoTags,omitempty" type:"Repeated"`
}

func (s ListVideosResponseBodyVideos) String() string {
	return tea.Prettify(s)
}

func (s ListVideosResponseBodyVideos) GoString() string {
	return s.String()
}

func (s *ListVideosResponseBodyVideos) SetCreateTime(v string) *ListVideosResponseBodyVideos {
	s.CreateTime = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetRemarksC(v string) *ListVideosResponseBodyVideos {
	s.RemarksC = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoTagsFailReason(v string) *ListVideosResponseBodyVideos {
	s.VideoTagsFailReason = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetSourceType(v string) *ListVideosResponseBodyVideos {
	s.SourceType = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoDuration(v float32) *ListVideosResponseBodyVideos {
	s.VideoDuration = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetProcessModifyTime(v string) *ListVideosResponseBodyVideos {
	s.ProcessModifyTime = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoFrames(v int32) *ListVideosResponseBodyVideos {
	s.VideoFrames = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoTagsStatus(v string) *ListVideosResponseBodyVideos {
	s.VideoTagsStatus = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetExternalId(v string) *ListVideosResponseBodyVideos {
	s.ExternalId = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetSourceUri(v string) *ListVideosResponseBodyVideos {
	s.SourceUri = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetModifyTime(v string) *ListVideosResponseBodyVideos {
	s.ModifyTime = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetFileSize(v int32) *ListVideosResponseBodyVideos {
	s.FileSize = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetSourcePosition(v string) *ListVideosResponseBodyVideos {
	s.SourcePosition = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoWidth(v int32) *ListVideosResponseBodyVideos {
	s.VideoWidth = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoHeight(v int32) *ListVideosResponseBodyVideos {
	s.VideoHeight = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoFormat(v string) *ListVideosResponseBodyVideos {
	s.VideoFormat = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetRemarksD(v string) *ListVideosResponseBodyVideos {
	s.RemarksD = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoUri(v string) *ListVideosResponseBodyVideos {
	s.VideoUri = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoTagsModifyTime(v string) *ListVideosResponseBodyVideos {
	s.VideoTagsModifyTime = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetProcessFailReason(v string) *ListVideosResponseBodyVideos {
	s.ProcessFailReason = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetRemarksA(v string) *ListVideosResponseBodyVideos {
	s.RemarksA = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetProcessStatus(v string) *ListVideosResponseBodyVideos {
	s.ProcessStatus = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetRemarksB(v string) *ListVideosResponseBodyVideos {
	s.RemarksB = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoTags(v []*ListVideosResponseBodyVideosVideoTags) *ListVideosResponseBodyVideos {
	s.VideoTags = v
	return s
}

type ListVideosResponseBodyVideosVideoTags struct {
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
}

func (s ListVideosResponseBodyVideosVideoTags) String() string {
	return tea.Prettify(s)
}

func (s ListVideosResponseBodyVideosVideoTags) GoString() string {
	return s.String()
}

func (s *ListVideosResponseBodyVideosVideoTags) SetTagLevel(v int32) *ListVideosResponseBodyVideosVideoTags {
	s.TagLevel = &v
	return s
}

func (s *ListVideosResponseBodyVideosVideoTags) SetParentTagName(v string) *ListVideosResponseBodyVideosVideoTags {
	s.ParentTagName = &v
	return s
}

func (s *ListVideosResponseBodyVideosVideoTags) SetTagName(v string) *ListVideosResponseBodyVideosVideoTags {
	s.TagName = &v
	return s
}

func (s *ListVideosResponseBodyVideosVideoTags) SetTagConfidence(v float32) *ListVideosResponseBodyVideosVideoTags {
	s.TagConfidence = &v
	return s
}

type ListVideosResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVideosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVideosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVideosResponse) GoString() string {
	return s.String()
}

func (s *ListVideosResponse) SetHeaders(v map[string]*string) *ListVideosResponse {
	s.Headers = v
	return s
}

func (s *ListVideosResponse) SetBody(v *ListVideosResponseBody) *ListVideosResponse {
	s.Body = v
	return s
}

type ListVideoTasksRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys  *int32  `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListVideoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVideoTasksRequest) GoString() string {
	return s.String()
}

func (s *ListVideoTasksRequest) SetProject(v string) *ListVideoTasksRequest {
	s.Project = &v
	return s
}

func (s *ListVideoTasksRequest) SetMarker(v string) *ListVideoTasksRequest {
	s.Marker = &v
	return s
}

func (s *ListVideoTasksRequest) SetMaxKeys(v int32) *ListVideoTasksRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListVideoTasksRequest) SetTaskType(v string) *ListVideoTasksRequest {
	s.TaskType = &v
	return s
}

type ListVideoTasksResponseBody struct {
	NextMarker *string                            `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*ListVideoTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListVideoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVideoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListVideoTasksResponseBody) SetNextMarker(v string) *ListVideoTasksResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListVideoTasksResponseBody) SetRequestId(v string) *ListVideoTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVideoTasksResponseBody) SetTasks(v []*ListVideoTasksResponseBodyTasks) *ListVideoTasksResponseBody {
	s.Tasks = v
	return s
}

type ListVideoTasksResponseBodyTasks struct {
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Progress        *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	ErrorMessage    *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Parameters      *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Result          *string `json:"Result,omitempty" xml:"Result,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
}

func (s ListVideoTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListVideoTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListVideoTasksResponseBodyTasks) SetEndTime(v string) *ListVideoTasksResponseBodyTasks {
	s.EndTime = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetStatus(v string) *ListVideoTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetStartTime(v string) *ListVideoTasksResponseBodyTasks {
	s.StartTime = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetTaskType(v string) *ListVideoTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetProgress(v int32) *ListVideoTasksResponseBodyTasks {
	s.Progress = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetNotifyEndpoint(v string) *ListVideoTasksResponseBodyTasks {
	s.NotifyEndpoint = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetErrorMessage(v string) *ListVideoTasksResponseBodyTasks {
	s.ErrorMessage = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetParameters(v string) *ListVideoTasksResponseBodyTasks {
	s.Parameters = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetResult(v string) *ListVideoTasksResponseBodyTasks {
	s.Result = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetTaskId(v string) *ListVideoTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetNotifyTopicName(v string) *ListVideoTasksResponseBodyTasks {
	s.NotifyTopicName = &v
	return s
}

type ListVideoTasksResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVideoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVideoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVideoTasksResponse) GoString() string {
	return s.String()
}

func (s *ListVideoTasksResponse) SetHeaders(v map[string]*string) *ListVideoTasksResponse {
	s.Headers = v
	return s
}

func (s *ListVideoTasksResponse) SetBody(v *ListVideoTasksResponseBody) *ListVideoTasksResponse {
	s.Body = v
	return s
}

type OpenImmServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenImmServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenImmServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenImmServiceRequest) SetOwnerId(v int64) *OpenImmServiceRequest {
	s.OwnerId = &v
	return s
}

type OpenImmServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenImmServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenImmServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenImmServiceResponseBody) SetOrderId(v string) *OpenImmServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenImmServiceResponseBody) SetRequestId(v string) *OpenImmServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenImmServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenImmServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenImmServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenImmServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenImmServiceResponse) SetHeaders(v map[string]*string) *OpenImmServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenImmServiceResponse) SetBody(v *OpenImmServiceResponseBody) *OpenImmServiceResponse {
	s.Body = v
	return s
}

type PutProjectRequest struct {
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
}

func (s PutProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PutProjectRequest) GoString() string {
	return s.String()
}

func (s *PutProjectRequest) SetProject(v string) *PutProjectRequest {
	s.Project = &v
	return s
}

func (s *PutProjectRequest) SetServiceRole(v string) *PutProjectRequest {
	s.ServiceRole = &v
	return s
}

type PutProjectResponseBody struct {
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CU          *int32  `json:"CU,omitempty" xml:"CU,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
}

func (s PutProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PutProjectResponseBody) SetProject(v string) *PutProjectResponseBody {
	s.Project = &v
	return s
}

func (s *PutProjectResponseBody) SetModifyTime(v string) *PutProjectResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *PutProjectResponseBody) SetType(v string) *PutProjectResponseBody {
	s.Type = &v
	return s
}

func (s *PutProjectResponseBody) SetCU(v int32) *PutProjectResponseBody {
	s.CU = &v
	return s
}

func (s *PutProjectResponseBody) SetServiceRole(v string) *PutProjectResponseBody {
	s.ServiceRole = &v
	return s
}

func (s *PutProjectResponseBody) SetRequestId(v string) *PutProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutProjectResponseBody) SetEndpoint(v string) *PutProjectResponseBody {
	s.Endpoint = &v
	return s
}

func (s *PutProjectResponseBody) SetCreateTime(v string) *PutProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *PutProjectResponseBody) SetRegionId(v string) *PutProjectResponseBody {
	s.RegionId = &v
	return s
}

func (s *PutProjectResponseBody) SetBillingType(v string) *PutProjectResponseBody {
	s.BillingType = &v
	return s
}

type PutProjectResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PutProjectResponse) GoString() string {
	return s.String()
}

func (s *PutProjectResponse) SetHeaders(v map[string]*string) *PutProjectResponse {
	s.Headers = v
	return s
}

func (s *PutProjectResponse) SetBody(v *PutProjectResponseBody) *PutProjectResponse {
	s.Body = v
	return s
}

type RefreshOfficeEditTokenRequest struct {
	Project      *string `json:"Project,omitempty" xml:"Project,omitempty"`
	AccessToken  *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s RefreshOfficeEditTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficeEditTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshOfficeEditTokenRequest) SetProject(v string) *RefreshOfficeEditTokenRequest {
	s.Project = &v
	return s
}

func (s *RefreshOfficeEditTokenRequest) SetAccessToken(v string) *RefreshOfficeEditTokenRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshOfficeEditTokenRequest) SetRefreshToken(v string) *RefreshOfficeEditTokenRequest {
	s.RefreshToken = &v
	return s
}

type RefreshOfficeEditTokenResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
}

func (s RefreshOfficeEditTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficeEditTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshOfficeEditTokenResponseBody) SetRequestId(v string) *RefreshOfficeEditTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshOfficeEditTokenResponseBody) SetAccessTokenExpiredTime(v string) *RefreshOfficeEditTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *RefreshOfficeEditTokenResponseBody) SetRefreshToken(v string) *RefreshOfficeEditTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *RefreshOfficeEditTokenResponseBody) SetAccessToken(v string) *RefreshOfficeEditTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *RefreshOfficeEditTokenResponseBody) SetRefreshTokenExpiredTime(v string) *RefreshOfficeEditTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

type RefreshOfficeEditTokenResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshOfficeEditTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshOfficeEditTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficeEditTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshOfficeEditTokenResponse) SetHeaders(v map[string]*string) *RefreshOfficeEditTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshOfficeEditTokenResponse) SetBody(v *RefreshOfficeEditTokenResponseBody) *RefreshOfficeEditTokenResponse {
	s.Body = v
	return s
}

type RefreshOfficePreviewTokenRequest struct {
	Project      *string `json:"Project,omitempty" xml:"Project,omitempty"`
	AccessToken  *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s RefreshOfficePreviewTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficePreviewTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshOfficePreviewTokenRequest) SetProject(v string) *RefreshOfficePreviewTokenRequest {
	s.Project = &v
	return s
}

func (s *RefreshOfficePreviewTokenRequest) SetAccessToken(v string) *RefreshOfficePreviewTokenRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshOfficePreviewTokenRequest) SetRefreshToken(v string) *RefreshOfficePreviewTokenRequest {
	s.RefreshToken = &v
	return s
}

type RefreshOfficePreviewTokenResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
}

func (s RefreshOfficePreviewTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficePreviewTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshOfficePreviewTokenResponseBody) SetRequestId(v string) *RefreshOfficePreviewTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponseBody) SetAccessTokenExpiredTime(v string) *RefreshOfficePreviewTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponseBody) SetRefreshToken(v string) *RefreshOfficePreviewTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponseBody) SetAccessToken(v string) *RefreshOfficePreviewTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponseBody) SetRefreshTokenExpiredTime(v string) *RefreshOfficePreviewTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

type RefreshOfficePreviewTokenResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshOfficePreviewTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshOfficePreviewTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficePreviewTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshOfficePreviewTokenResponse) SetHeaders(v map[string]*string) *RefreshOfficePreviewTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshOfficePreviewTokenResponse) SetBody(v *RefreshOfficePreviewTokenResponseBody) *RefreshOfficePreviewTokenResponse {
	s.Body = v
	return s
}

type RefreshWebofficeTokenRequest struct {
	Project      *string `json:"Project,omitempty" xml:"Project,omitempty"`
	AccessToken  *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s RefreshWebofficeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenRequest) SetProject(v string) *RefreshWebofficeTokenRequest {
	s.Project = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetAccessToken(v string) *RefreshWebofficeTokenRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetRefreshToken(v string) *RefreshWebofficeTokenRequest {
	s.RefreshToken = &v
	return s
}

type RefreshWebofficeTokenResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
}

func (s RefreshWebofficeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponseBody) SetRequestId(v string) *RefreshWebofficeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshToken(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessToken(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

type RefreshWebofficeTokenResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshWebofficeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshWebofficeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponse) SetHeaders(v map[string]*string) *RefreshWebofficeTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshWebofficeTokenResponse) SetBody(v *RefreshWebofficeTokenResponseBody) *RefreshWebofficeTokenResponse {
	s.Body = v
	return s
}

type UpdateFaceGroupRequest struct {
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId            *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupCoverFaceId *string `json:"GroupCoverFaceId,omitempty" xml:"GroupCoverFaceId,omitempty"`
	RemarksA         *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB         *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC         *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD         *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	RemarksArrayA    *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB    *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	ExternalId       *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ResetItems       *string `json:"ResetItems,omitempty" xml:"ResetItems,omitempty"`
}

func (s UpdateFaceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceGroupRequest) SetProject(v string) *UpdateFaceGroupRequest {
	s.Project = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetSetId(v string) *UpdateFaceGroupRequest {
	s.SetId = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetGroupId(v string) *UpdateFaceGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetGroupName(v string) *UpdateFaceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetGroupCoverFaceId(v string) *UpdateFaceGroupRequest {
	s.GroupCoverFaceId = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksA(v string) *UpdateFaceGroupRequest {
	s.RemarksA = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksB(v string) *UpdateFaceGroupRequest {
	s.RemarksB = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksC(v string) *UpdateFaceGroupRequest {
	s.RemarksC = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksD(v string) *UpdateFaceGroupRequest {
	s.RemarksD = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksArrayA(v string) *UpdateFaceGroupRequest {
	s.RemarksArrayA = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksArrayB(v string) *UpdateFaceGroupRequest {
	s.RemarksArrayB = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetExternalId(v string) *UpdateFaceGroupRequest {
	s.ExternalId = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetResetItems(v string) *UpdateFaceGroupRequest {
	s.ResetItems = &v
	return s
}

type UpdateFaceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s UpdateFaceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceGroupResponseBody) SetRequestId(v string) *UpdateFaceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFaceGroupResponseBody) SetGroupId(v string) *UpdateFaceGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *UpdateFaceGroupResponseBody) SetSetId(v string) *UpdateFaceGroupResponseBody {
	s.SetId = &v
	return s
}

type UpdateFaceGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFaceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFaceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceGroupResponse) SetHeaders(v map[string]*string) *UpdateFaceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceGroupResponse) SetBody(v *UpdateFaceGroupResponseBody) *UpdateFaceGroupResponse {
	s.Body = v
	return s
}

type UpdateImageRequest struct {
	Project        *string                    `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId          *string                    `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageUri       *string                    `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RemarksA       *string                    `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB       *string                    `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	SourceType     *string                    `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri      *string                    `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	SourcePosition *string                    `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	RemarksC       *string                    `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD       *string                    `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	ExternalId     *string                    `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	RemarksArrayA  *string                    `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB  *string                    `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	Faces          []*UpdateImageRequestFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
}

func (s UpdateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) SetProject(v string) *UpdateImageRequest {
	s.Project = &v
	return s
}

func (s *UpdateImageRequest) SetSetId(v string) *UpdateImageRequest {
	s.SetId = &v
	return s
}

func (s *UpdateImageRequest) SetImageUri(v string) *UpdateImageRequest {
	s.ImageUri = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksA(v string) *UpdateImageRequest {
	s.RemarksA = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksB(v string) *UpdateImageRequest {
	s.RemarksB = &v
	return s
}

func (s *UpdateImageRequest) SetSourceType(v string) *UpdateImageRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateImageRequest) SetSourceUri(v string) *UpdateImageRequest {
	s.SourceUri = &v
	return s
}

func (s *UpdateImageRequest) SetSourcePosition(v string) *UpdateImageRequest {
	s.SourcePosition = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksC(v string) *UpdateImageRequest {
	s.RemarksC = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksD(v string) *UpdateImageRequest {
	s.RemarksD = &v
	return s
}

func (s *UpdateImageRequest) SetExternalId(v string) *UpdateImageRequest {
	s.ExternalId = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksArrayA(v string) *UpdateImageRequest {
	s.RemarksArrayA = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksArrayB(v string) *UpdateImageRequest {
	s.RemarksArrayB = &v
	return s
}

func (s *UpdateImageRequest) SetFaces(v []*UpdateImageRequestFaces) *UpdateImageRequest {
	s.Faces = v
	return s
}

type UpdateImageRequestFaces struct {
	FaceId  *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UpdateImageRequestFaces) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageRequestFaces) GoString() string {
	return s.String()
}

func (s *UpdateImageRequestFaces) SetFaceId(v string) *UpdateImageRequestFaces {
	s.FaceId = &v
	return s
}

func (s *UpdateImageRequestFaces) SetGroupId(v string) *UpdateImageRequestFaces {
	s.GroupId = &v
	return s
}

type UpdateImageShrinkRequest struct {
	Project        *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId          *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	ImageUri       *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RemarksA       *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB       *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	SourceType     *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri      *string `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	SourcePosition *string `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	RemarksC       *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD       *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	ExternalId     *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	RemarksArrayA  *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB  *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	FacesShrink    *string `json:"Faces,omitempty" xml:"Faces,omitempty"`
}

func (s UpdateImageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageShrinkRequest) SetProject(v string) *UpdateImageShrinkRequest {
	s.Project = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSetId(v string) *UpdateImageShrinkRequest {
	s.SetId = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetImageUri(v string) *UpdateImageShrinkRequest {
	s.ImageUri = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksA(v string) *UpdateImageShrinkRequest {
	s.RemarksA = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksB(v string) *UpdateImageShrinkRequest {
	s.RemarksB = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSourceType(v string) *UpdateImageShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSourceUri(v string) *UpdateImageShrinkRequest {
	s.SourceUri = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSourcePosition(v string) *UpdateImageShrinkRequest {
	s.SourcePosition = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksC(v string) *UpdateImageShrinkRequest {
	s.RemarksC = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksD(v string) *UpdateImageShrinkRequest {
	s.RemarksD = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetExternalId(v string) *UpdateImageShrinkRequest {
	s.ExternalId = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksArrayA(v string) *UpdateImageShrinkRequest {
	s.RemarksArrayA = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksArrayB(v string) *UpdateImageShrinkRequest {
	s.RemarksArrayB = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetFacesShrink(v string) *UpdateImageShrinkRequest {
	s.FacesShrink = &v
	return s
}

type UpdateImageResponseBody struct {
	RemarksArrayB *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	ModifyTime    *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RemarksC      *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD      *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId    *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	RemarksArrayA *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksA      *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	ImageUri      *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	SetId         *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	RemarksB      *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) SetRemarksArrayB(v string) *UpdateImageResponseBody {
	s.RemarksArrayB = &v
	return s
}

func (s *UpdateImageResponseBody) SetModifyTime(v string) *UpdateImageResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksC(v string) *UpdateImageResponseBody {
	s.RemarksC = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksD(v string) *UpdateImageResponseBody {
	s.RemarksD = &v
	return s
}

func (s *UpdateImageResponseBody) SetRequestId(v string) *UpdateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageResponseBody) SetCreateTime(v string) *UpdateImageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UpdateImageResponseBody) SetExternalId(v string) *UpdateImageResponseBody {
	s.ExternalId = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksArrayA(v string) *UpdateImageResponseBody {
	s.RemarksArrayA = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksA(v string) *UpdateImageResponseBody {
	s.RemarksA = &v
	return s
}

func (s *UpdateImageResponseBody) SetImageUri(v string) *UpdateImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *UpdateImageResponseBody) SetSetId(v string) *UpdateImageResponseBody {
	s.SetId = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksB(v string) *UpdateImageResponseBody {
	s.RemarksB = &v
	return s
}

type UpdateImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageResponse) SetHeaders(v map[string]*string) *UpdateImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageResponse) SetBody(v *UpdateImageResponseBody) *UpdateImageResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	Project        *string `json:"Project,omitempty" xml:"Project,omitempty"`
	NewCU          *int32  `json:"NewCU,omitempty" xml:"NewCU,omitempty"`
	NewServiceRole *string `json:"NewServiceRole,omitempty" xml:"NewServiceRole,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetProject(v string) *UpdateProjectRequest {
	s.Project = &v
	return s
}

func (s *UpdateProjectRequest) SetNewCU(v int32) *UpdateProjectRequest {
	s.NewCU = &v
	return s
}

func (s *UpdateProjectRequest) SetNewServiceRole(v string) *UpdateProjectRequest {
	s.NewServiceRole = &v
	return s
}

type UpdateProjectResponseBody struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CU          *int32  `json:"CU,omitempty" xml:"CU,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetType(v string) *UpdateProjectResponseBody {
	s.Type = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetCU(v int32) *UpdateProjectResponseBody {
	s.CU = &v
	return s
}

func (s *UpdateProjectResponseBody) SetCreateTime(v string) *UpdateProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UpdateProjectResponseBody) SetServiceRole(v string) *UpdateProjectResponseBody {
	s.ServiceRole = &v
	return s
}

func (s *UpdateProjectResponseBody) SetProject(v string) *UpdateProjectResponseBody {
	s.Project = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRegionId(v string) *UpdateProjectResponseBody {
	s.RegionId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetModifyTime(v string) *UpdateProjectResponseBody {
	s.ModifyTime = &v
	return s
}

type UpdateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateSetRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
}

func (s UpdateSetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateSetRequest) SetProject(v string) *UpdateSetRequest {
	s.Project = &v
	return s
}

func (s *UpdateSetRequest) SetSetId(v string) *UpdateSetRequest {
	s.SetId = &v
	return s
}

func (s *UpdateSetRequest) SetSetName(v string) *UpdateSetRequest {
	s.SetName = &v
	return s
}

type UpdateSetResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SetName    *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SetId      *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s UpdateSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSetResponseBody) SetRequestId(v string) *UpdateSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSetResponseBody) SetCreateTime(v string) *UpdateSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UpdateSetResponseBody) SetSetName(v string) *UpdateSetResponseBody {
	s.SetName = &v
	return s
}

func (s *UpdateSetResponseBody) SetModifyTime(v string) *UpdateSetResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *UpdateSetResponseBody) SetSetId(v string) *UpdateSetResponseBody {
	s.SetId = &v
	return s
}

type UpdateSetResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateSetResponse) SetHeaders(v map[string]*string) *UpdateSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateSetResponse) SetBody(v *UpdateSetResponseBody) *UpdateSetResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing-gov-1": tea.String("imm-vpc.cn-beijing-gov-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CompareImageFacesWithOptions(request *CompareImageFacesRequest, runtime *util.RuntimeOptions) (_result *CompareImageFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CompareImageFacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CompareImageFaces"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareImageFaces(request *CompareImageFacesRequest) (_result *CompareImageFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareImageFacesResponse{}
	_body, _err := client.CompareImageFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertOfficeFormatWithOptions(request *ConvertOfficeFormatRequest, runtime *util.RuntimeOptions) (_result *ConvertOfficeFormatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConvertOfficeFormatResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConvertOfficeFormat"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertOfficeFormat(request *ConvertOfficeFormatRequest) (_result *ConvertOfficeFormatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertOfficeFormatResponse{}
	_body, _err := client.ConvertOfficeFormatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGrabFrameTaskWithOptions(request *CreateGrabFrameTaskRequest, runtime *util.RuntimeOptions) (_result *CreateGrabFrameTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGrabFrameTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGrabFrameTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGrabFrameTask(request *CreateGrabFrameTaskRequest) (_result *CreateGrabFrameTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGrabFrameTaskResponse{}
	_body, _err := client.CreateGrabFrameTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupFacesJobWithOptions(request *CreateGroupFacesJobRequest, runtime *util.RuntimeOptions) (_result *CreateGroupFacesJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGroupFacesJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGroupFacesJob"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupFacesJob(request *CreateGroupFacesJobRequest) (_result *CreateGroupFacesJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupFacesJobResponse{}
	_body, _err := client.CreateGroupFacesJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageProcessTaskWithOptions(request *CreateImageProcessTaskRequest, runtime *util.RuntimeOptions) (_result *CreateImageProcessTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateImageProcessTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateImageProcessTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageProcessTask(request *CreateImageProcessTaskRequest) (_result *CreateImageProcessTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageProcessTaskResponse{}
	_body, _err := client.CreateImageProcessTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMediaComplexTaskWithOptions(request *CreateMediaComplexTaskRequest, runtime *util.RuntimeOptions) (_result *CreateMediaComplexTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMediaComplexTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMediaComplexTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMediaComplexTask(request *CreateMediaComplexTaskRequest) (_result *CreateMediaComplexTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMediaComplexTaskResponse{}
	_body, _err := client.CreateMediaComplexTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMergeFaceGroupsJobWithOptions(request *CreateMergeFaceGroupsJobRequest, runtime *util.RuntimeOptions) (_result *CreateMergeFaceGroupsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMergeFaceGroupsJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMergeFaceGroupsJob"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMergeFaceGroupsJob(request *CreateMergeFaceGroupsJobRequest) (_result *CreateMergeFaceGroupsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMergeFaceGroupsJobResponse{}
	_body, _err := client.CreateMergeFaceGroupsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOfficeConversionTaskWithOptions(request *CreateOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *CreateOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOfficeConversionTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOfficeConversionTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOfficeConversionTask(request *CreateOfficeConversionTaskRequest) (_result *CreateOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOfficeConversionTaskResponse{}
	_body, _err := client.CreateOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSetWithOptions(request *CreateSetRequest, runtime *util.RuntimeOptions) (_result *CreateSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSet"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSet(request *CreateSetRequest) (_result *CreateSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSetResponse{}
	_body, _err := client.CreateSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoAbstractTaskWithOptions(request *CreateVideoAbstractTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoAbstractTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVideoAbstractTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVideoAbstractTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoAbstractTask(request *CreateVideoAbstractTaskRequest) (_result *CreateVideoAbstractTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoAbstractTaskResponse{}
	_body, _err := client.CreateVideoAbstractTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoAnalyseTaskWithOptions(request *CreateVideoAnalyseTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoAnalyseTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVideoAnalyseTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVideoAnalyseTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoAnalyseTask(request *CreateVideoAnalyseTaskRequest) (_result *CreateVideoAnalyseTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoAnalyseTaskResponse{}
	_body, _err := client.CreateVideoAnalyseTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoCompressTaskWithOptions(request *CreateVideoCompressTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoCompressTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVideoCompressTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVideoCompressTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoCompressTask(request *CreateVideoCompressTaskRequest) (_result *CreateVideoCompressTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoCompressTaskResponse{}
	_body, _err := client.CreateVideoCompressTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoProduceTaskWithOptions(request *CreateVideoProduceTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoProduceTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVideoProduceTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVideoProduceTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoProduceTask(request *CreateVideoProduceTaskRequest) (_result *CreateVideoProduceTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoProduceTaskResponse{}
	_body, _err := client.CreateVideoProduceTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DecodeBlindWatermarkWithOptions(request *DecodeBlindWatermarkRequest, runtime *util.RuntimeOptions) (_result *DecodeBlindWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DecodeBlindWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DecodeBlindWatermark"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DecodeBlindWatermark(request *DecodeBlindWatermarkRequest) (_result *DecodeBlindWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DecodeBlindWatermarkResponse{}
	_body, _err := client.DecodeBlindWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImage"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageJobWithOptions(request *DeleteImageJobRequest, runtime *util.RuntimeOptions) (_result *DeleteImageJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImageJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImageJob"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImageJob(request *DeleteImageJobRequest) (_result *DeleteImageJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageJobResponse{}
	_body, _err := client.DeleteImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOfficeConversionTaskWithOptions(request *DeleteOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteOfficeConversionTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteOfficeConversionTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOfficeConversionTask(request *DeleteOfficeConversionTaskRequest) (_result *DeleteOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOfficeConversionTaskResponse{}
	_body, _err := client.DeleteOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProject"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSetWithOptions(request *DeleteSetRequest, runtime *util.RuntimeOptions) (_result *DeleteSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSet"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSet(request *DeleteSetRequest) (_result *DeleteSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSetResponse{}
	_body, _err := client.DeleteSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVideoWithOptions(request *DeleteVideoRequest, runtime *util.RuntimeOptions) (_result *DeleteVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVideo"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVideo(request *DeleteVideoRequest) (_result *DeleteVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVideoResponse{}
	_body, _err := client.DeleteVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVideoTaskWithOptions(request *DeleteVideoTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVideoTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVideoTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVideoTask(request *DeleteVideoTaskRequest) (_result *DeleteVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVideoTaskResponse{}
	_body, _err := client.DeleteVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageBodiesWithOptions(request *DetectImageBodiesRequest, runtime *util.RuntimeOptions) (_result *DetectImageBodiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectImageBodiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectImageBodies"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageBodies(request *DetectImageBodiesRequest) (_result *DetectImageBodiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageBodiesResponse{}
	_body, _err := client.DetectImageBodiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageFacesWithOptions(request *DetectImageFacesRequest, runtime *util.RuntimeOptions) (_result *DetectImageFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectImageFacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectImageFaces"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageFaces(request *DetectImageFacesRequest) (_result *DetectImageFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageFacesResponse{}
	_body, _err := client.DetectImageFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageLogosWithOptions(request *DetectImageLogosRequest, runtime *util.RuntimeOptions) (_result *DetectImageLogosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectImageLogosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectImageLogos"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageLogos(request *DetectImageLogosRequest) (_result *DetectImageLogosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageLogosResponse{}
	_body, _err := client.DetectImageLogosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageQRCodesWithOptions(request *DetectImageQRCodesRequest, runtime *util.RuntimeOptions) (_result *DetectImageQRCodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectImageQRCodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectImageQRCodes"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageQRCodes(request *DetectImageQRCodesRequest) (_result *DetectImageQRCodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageQRCodesResponse{}
	_body, _err := client.DetectImageQRCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageTagsWithOptions(request *DetectImageTagsRequest, runtime *util.RuntimeOptions) (_result *DetectImageTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectImageTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectImageTags"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageTags(request *DetectImageTagsRequest) (_result *DetectImageTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageTagsResponse{}
	_body, _err := client.DetectImageTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectQRCodesWithOptions(request *DetectQRCodesRequest, runtime *util.RuntimeOptions) (_result *DetectQRCodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectQRCodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectQRCodes"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectQRCodes(request *DetectQRCodesRequest) (_result *DetectQRCodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectQRCodesResponse{}
	_body, _err := client.DetectQRCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EncodeBlindWatermarkWithOptions(request *EncodeBlindWatermarkRequest, runtime *util.RuntimeOptions) (_result *EncodeBlindWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EncodeBlindWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EncodeBlindWatermark"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EncodeBlindWatermark(request *EncodeBlindWatermarkRequest) (_result *EncodeBlindWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EncodeBlindWatermarkResponse{}
	_body, _err := client.EncodeBlindWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindImagesWithOptions(request *FindImagesRequest, runtime *util.RuntimeOptions) (_result *FindImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindImages"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindImages(request *FindImagesRequest) (_result *FindImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindImagesResponse{}
	_body, _err := client.FindImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindSimilarFacesWithOptions(request *FindSimilarFacesRequest, runtime *util.RuntimeOptions) (_result *FindSimilarFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindSimilarFacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindSimilarFaces"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindSimilarFaces(request *FindSimilarFacesRequest) (_result *FindSimilarFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindSimilarFacesResponse{}
	_body, _err := client.FindSimilarFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetContentKeyWithOptions(request *GetContentKeyRequest, runtime *util.RuntimeOptions) (_result *GetContentKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetContentKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetContentKey"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetContentKey(request *GetContentKeyRequest) (_result *GetContentKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetContentKeyResponse{}
	_body, _err := client.GetContentKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDRMLicenseWithOptions(request *GetDRMLicenseRequest, runtime *util.RuntimeOptions) (_result *GetDRMLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDRMLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDRMLicense"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDRMLicense(request *GetDRMLicenseRequest) (_result *GetDRMLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDRMLicenseResponse{}
	_body, _err := client.GetDRMLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageWithOptions(request *GetImageRequest, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetImage"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImage(request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageCroppingSuggestionsWithOptions(request *GetImageCroppingSuggestionsRequest, runtime *util.RuntimeOptions) (_result *GetImageCroppingSuggestionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetImageCroppingSuggestionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetImageCroppingSuggestions"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageCroppingSuggestions(request *GetImageCroppingSuggestionsRequest) (_result *GetImageCroppingSuggestionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageCroppingSuggestionsResponse{}
	_body, _err := client.GetImageCroppingSuggestionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageQualityWithOptions(request *GetImageQualityRequest, runtime *util.RuntimeOptions) (_result *GetImageQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetImageQualityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetImageQuality"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageQuality(request *GetImageQualityRequest) (_result *GetImageQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageQualityResponse{}
	_body, _err := client.GetImageQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaMetaWithOptions(request *GetMediaMetaRequest, runtime *util.RuntimeOptions) (_result *GetMediaMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaMeta"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaMeta(request *GetMediaMetaRequest) (_result *GetMediaMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaMetaResponse{}
	_body, _err := client.GetMediaMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficeConversionTaskWithOptions(request *GetOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *GetOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOfficeConversionTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOfficeConversionTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficeConversionTask(request *GetOfficeConversionTaskRequest) (_result *GetOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOfficeConversionTaskResponse{}
	_body, _err := client.GetOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficeEditURLWithOptions(request *GetOfficeEditURLRequest, runtime *util.RuntimeOptions) (_result *GetOfficeEditURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOfficeEditURLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOfficeEditURL"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficeEditURL(request *GetOfficeEditURLRequest) (_result *GetOfficeEditURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOfficeEditURLResponse{}
	_body, _err := client.GetOfficeEditURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficePreviewURLWithOptions(request *GetOfficePreviewURLRequest, runtime *util.RuntimeOptions) (_result *GetOfficePreviewURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOfficePreviewURLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOfficePreviewURL"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficePreviewURL(request *GetOfficePreviewURLRequest) (_result *GetOfficePreviewURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOfficePreviewURLResponse{}
	_body, _err := client.GetOfficePreviewURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProject"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSetWithOptions(request *GetSetRequest, runtime *util.RuntimeOptions) (_result *GetSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSet"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSet(request *GetSetRequest) (_result *GetSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSetResponse{}
	_body, _err := client.GetSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoWithOptions(request *GetVideoRequest, runtime *util.RuntimeOptions) (_result *GetVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideo"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideo(request *GetVideoRequest) (_result *GetVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoResponse{}
	_body, _err := client.GetVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoTaskWithOptions(request *GetVideoTaskRequest, runtime *util.RuntimeOptions) (_result *GetVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideoTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoTask(request *GetVideoTaskRequest) (_result *GetVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoTaskResponse{}
	_body, _err := client.GetVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebofficeURLWithOptions(request *GetWebofficeURLRequest, runtime *util.RuntimeOptions) (_result *GetWebofficeURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWebofficeURLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWebofficeURL"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebofficeURL(request *GetWebofficeURLRequest) (_result *GetWebofficeURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebofficeURLResponse{}
	_body, _err := client.GetWebofficeURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndexImageWithOptions(request *IndexImageRequest, runtime *util.RuntimeOptions) (_result *IndexImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IndexImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IndexImage"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndexImage(request *IndexImageRequest) (_result *IndexImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IndexImageResponse{}
	_body, _err := client.IndexImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndexVideoWithOptions(request *IndexVideoRequest, runtime *util.RuntimeOptions) (_result *IndexVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IndexVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IndexVideo"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndexVideo(request *IndexVideoRequest) (_result *IndexVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IndexVideoResponse{}
	_body, _err := client.IndexVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceGroupsWithOptions(request *ListFaceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListFaceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFaceGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFaceGroups"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceGroups(request *ListFaceGroupsRequest) (_result *ListFaceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceGroupsResponse{}
	_body, _err := client.ListFaceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListImages"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfficeConversionTaskWithOptions(request *ListOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *ListOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOfficeConversionTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOfficeConversionTask"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfficeConversionTask(request *ListOfficeConversionTaskRequest) (_result *ListOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfficeConversionTaskResponse{}
	_body, _err := client.ListOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectAPIsWithOptions(request *ListProjectAPIsRequest, runtime *util.RuntimeOptions) (_result *ListProjectAPIsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectAPIsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjectAPIs"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectAPIs(request *ListProjectAPIsRequest) (_result *ListProjectAPIsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectAPIsResponse{}
	_body, _err := client.ListProjectAPIsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjects"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSetsWithOptions(request *ListSetsRequest, runtime *util.RuntimeOptions) (_result *ListSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSets"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSets(request *ListSetsRequest) (_result *ListSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSetsResponse{}
	_body, _err := client.ListSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSetTagsWithOptions(request *ListSetTagsRequest, runtime *util.RuntimeOptions) (_result *ListSetTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSetTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSetTags"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSetTags(request *ListSetTagsRequest) (_result *ListSetTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSetTagsResponse{}
	_body, _err := client.ListSetTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVideoAudiosWithOptions(request *ListVideoAudiosRequest, runtime *util.RuntimeOptions) (_result *ListVideoAudiosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVideoAudiosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVideoAudios"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVideoAudios(request *ListVideoAudiosRequest) (_result *ListVideoAudiosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVideoAudiosResponse{}
	_body, _err := client.ListVideoAudiosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVideoFramesWithOptions(request *ListVideoFramesRequest, runtime *util.RuntimeOptions) (_result *ListVideoFramesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVideoFramesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVideoFrames"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVideoFrames(request *ListVideoFramesRequest) (_result *ListVideoFramesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVideoFramesResponse{}
	_body, _err := client.ListVideoFramesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVideosWithOptions(request *ListVideosRequest, runtime *util.RuntimeOptions) (_result *ListVideosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVideosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVideos"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVideos(request *ListVideosRequest) (_result *ListVideosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVideosResponse{}
	_body, _err := client.ListVideosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVideoTasksWithOptions(request *ListVideoTasksRequest, runtime *util.RuntimeOptions) (_result *ListVideoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVideoTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVideoTasks"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVideoTasks(request *ListVideoTasksRequest) (_result *ListVideoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVideoTasksResponse{}
	_body, _err := client.ListVideoTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenImmServiceWithOptions(request *OpenImmServiceRequest, runtime *util.RuntimeOptions) (_result *OpenImmServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenImmServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenImmService"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenImmService(request *OpenImmServiceRequest) (_result *OpenImmServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenImmServiceResponse{}
	_body, _err := client.OpenImmServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutProjectWithOptions(request *PutProjectRequest, runtime *util.RuntimeOptions) (_result *PutProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutProject"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutProject(request *PutProjectRequest) (_result *PutProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutProjectResponse{}
	_body, _err := client.PutProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshOfficeEditTokenWithOptions(request *RefreshOfficeEditTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshOfficeEditTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshOfficeEditTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshOfficeEditToken"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshOfficeEditToken(request *RefreshOfficeEditTokenRequest) (_result *RefreshOfficeEditTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshOfficeEditTokenResponse{}
	_body, _err := client.RefreshOfficeEditTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshOfficePreviewTokenWithOptions(request *RefreshOfficePreviewTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshOfficePreviewTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshOfficePreviewTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshOfficePreviewToken"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshOfficePreviewToken(request *RefreshOfficePreviewTokenRequest) (_result *RefreshOfficePreviewTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshOfficePreviewTokenResponse{}
	_body, _err := client.RefreshOfficePreviewTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshWebofficeTokenWithOptions(request *RefreshWebofficeTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshWebofficeTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshWebofficeToken"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshWebofficeToken(request *RefreshWebofficeTokenRequest) (_result *RefreshWebofficeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.RefreshWebofficeTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFaceGroupWithOptions(request *UpdateFaceGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateFaceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFaceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFaceGroup"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFaceGroup(request *UpdateFaceGroupRequest) (_result *UpdateFaceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaceGroupResponse{}
	_body, _err := client.UpdateFaceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateImageWithOptions(tmpReq *UpdateImageRequest, runtime *util.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Faces)) {
		request.FacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Faces, tea.String("Faces"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateImage"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateImage(request *UpdateImageRequest) (_result *UpdateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageResponse{}
	_body, _err := client.UpdateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProject"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSetWithOptions(request *UpdateSetRequest, runtime *util.RuntimeOptions) (_result *UpdateSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSet"), tea.String("2017-09-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSet(request *UpdateSetRequest) (_result *UpdateSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSetResponse{}
	_body, _err := client.UpdateSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
