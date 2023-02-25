// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type Box struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s Box) String() string {
	return tea.Prettify(s)
}

func (s Box) GoString() string {
	return s.String()
}

func (s *Box) SetHeight(v int32) *Box {
	s.Height = &v
	return s
}

func (s *Box) SetLeft(v int32) *Box {
	s.Left = &v
	return s
}

func (s *Box) SetTop(v int32) *Box {
	s.Top = &v
	return s
}

func (s *Box) SetWidth(v int32) *Box {
	s.Width = &v
	return s
}

type AddImageRequest struct {
	Boxes         []*Box  `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	DetectLimit   *int32  `json:"DetectLimit,omitempty" xml:"DetectLimit,omitempty"`
	IntAttr       *int32  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	PicContent    *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicName       *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	PicUrl        *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	StrAttr       *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetBoxes(v []*Box) *AddImageRequest {
	s.Boxes = v
	return s
}

func (s *AddImageRequest) SetCustomContent(v string) *AddImageRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageRequest) SetDetectLimit(v int32) *AddImageRequest {
	s.DetectLimit = &v
	return s
}

func (s *AddImageRequest) SetIntAttr(v int32) *AddImageRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageRequest) SetPicContent(v string) *AddImageRequest {
	s.PicContent = &v
	return s
}

func (s *AddImageRequest) SetPicName(v string) *AddImageRequest {
	s.PicName = &v
	return s
}

func (s *AddImageRequest) SetPicUrl(v string) *AddImageRequest {
	s.PicUrl = &v
	return s
}

func (s *AddImageRequest) SetStrAttr(v string) *AddImageRequest {
	s.StrAttr = &v
	return s
}

type AddImageAdvanceRequest struct {
	Boxes            []*Box    `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	CustomContent    *string   `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	DetectLimit      *int32    `json:"DetectLimit,omitempty" xml:"DetectLimit,omitempty"`
	IntAttr          *int32    `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	PicContentObject io.Reader `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicName          *string   `json:"PicName,omitempty" xml:"PicName,omitempty"`
	PicUrl           *string   `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	StrAttr          *string   `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s AddImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddImageAdvanceRequest) SetBoxes(v []*Box) *AddImageAdvanceRequest {
	s.Boxes = v
	return s
}

func (s *AddImageAdvanceRequest) SetCustomContent(v string) *AddImageAdvanceRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageAdvanceRequest) SetDetectLimit(v int32) *AddImageAdvanceRequest {
	s.DetectLimit = &v
	return s
}

func (s *AddImageAdvanceRequest) SetIntAttr(v int32) *AddImageAdvanceRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageAdvanceRequest) SetPicContentObject(v io.Reader) *AddImageAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *AddImageAdvanceRequest) SetPicName(v string) *AddImageAdvanceRequest {
	s.PicName = &v
	return s
}

func (s *AddImageAdvanceRequest) SetPicUrl(v string) *AddImageAdvanceRequest {
	s.PicUrl = &v
	return s
}

func (s *AddImageAdvanceRequest) SetStrAttr(v string) *AddImageAdvanceRequest {
	s.StrAttr = &v
	return s
}

type AddImageShrinkRequest struct {
	BoxesShrink   *string `json:"Boxes,omitempty" xml:"Boxes,omitempty"`
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	DetectLimit   *int32  `json:"DetectLimit,omitempty" xml:"DetectLimit,omitempty"`
	IntAttr       *int32  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	PicContent    *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicName       *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	PicUrl        *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	StrAttr       *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s AddImageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddImageShrinkRequest) SetBoxesShrink(v string) *AddImageShrinkRequest {
	s.BoxesShrink = &v
	return s
}

func (s *AddImageShrinkRequest) SetCustomContent(v string) *AddImageShrinkRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageShrinkRequest) SetDetectLimit(v int32) *AddImageShrinkRequest {
	s.DetectLimit = &v
	return s
}

func (s *AddImageShrinkRequest) SetIntAttr(v int32) *AddImageShrinkRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageShrinkRequest) SetPicContent(v string) *AddImageShrinkRequest {
	s.PicContent = &v
	return s
}

func (s *AddImageShrinkRequest) SetPicName(v string) *AddImageShrinkRequest {
	s.PicName = &v
	return s
}

func (s *AddImageShrinkRequest) SetPicUrl(v string) *AddImageShrinkRequest {
	s.PicUrl = &v
	return s
}

func (s *AddImageShrinkRequest) SetStrAttr(v string) *AddImageShrinkRequest {
	s.StrAttr = &v
	return s
}

type AddImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetSuccess(v bool) *AddImageResponseBody {
	s.Success = &v
	return s
}

type AddImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddImageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponse) GoString() string {
	return s.String()
}

func (s *AddImageResponse) SetHeaders(v map[string]*string) *AddImageResponse {
	s.Headers = v
	return s
}

func (s *AddImageResponse) SetStatusCode(v int32) *AddImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type CheckOssIncrementMetaExistRequest struct {
	// oss bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// oss path
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s CheckOssIncrementMetaExistRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckOssIncrementMetaExistRequest) GoString() string {
	return s.String()
}

func (s *CheckOssIncrementMetaExistRequest) SetBucketName(v string) *CheckOssIncrementMetaExistRequest {
	s.BucketName = &v
	return s
}

func (s *CheckOssIncrementMetaExistRequest) SetKey(v string) *CheckOssIncrementMetaExistRequest {
	s.Key = &v
	return s
}

func (s *CheckOssIncrementMetaExistRequest) SetOssPath(v string) *CheckOssIncrementMetaExistRequest {
	s.OssPath = &v
	return s
}

type CheckOssIncrementMetaExistResponseBody struct {
	Data      *CheckOssIncrementMetaExistResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckOssIncrementMetaExistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckOssIncrementMetaExistResponseBody) GoString() string {
	return s.String()
}

func (s *CheckOssIncrementMetaExistResponseBody) SetData(v *CheckOssIncrementMetaExistResponseBodyData) *CheckOssIncrementMetaExistResponseBody {
	s.Data = v
	return s
}

func (s *CheckOssIncrementMetaExistResponseBody) SetRequestId(v string) *CheckOssIncrementMetaExistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckOssIncrementMetaExistResponseBody) SetSuccess(v bool) *CheckOssIncrementMetaExistResponseBody {
	s.Success = &v
	return s
}

type CheckOssIncrementMetaExistResponseBodyData struct {
	Exist *bool `json:"Exist,omitempty" xml:"Exist,omitempty"`
}

func (s CheckOssIncrementMetaExistResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckOssIncrementMetaExistResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckOssIncrementMetaExistResponseBodyData) SetExist(v bool) *CheckOssIncrementMetaExistResponseBodyData {
	s.Exist = &v
	return s
}

type CheckOssIncrementMetaExistResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckOssIncrementMetaExistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckOssIncrementMetaExistResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckOssIncrementMetaExistResponse) GoString() string {
	return s.String()
}

func (s *CheckOssIncrementMetaExistResponse) SetHeaders(v map[string]*string) *CheckOssIncrementMetaExistResponse {
	s.Headers = v
	return s
}

func (s *CheckOssIncrementMetaExistResponse) SetStatusCode(v int32) *CheckOssIncrementMetaExistResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckOssIncrementMetaExistResponse) SetBody(v *CheckOssIncrementMetaExistResponseBody) *CheckOssIncrementMetaExistResponse {
	s.Body = v
	return s
}

type CreateBatchTasksRequest struct {
	// oss bucket
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CallbackAddress *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	// oss path
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s CreateBatchTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTasksRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTasksRequest) SetBucketName(v string) *CreateBatchTasksRequest {
	s.BucketName = &v
	return s
}

func (s *CreateBatchTasksRequest) SetCallbackAddress(v string) *CreateBatchTasksRequest {
	s.CallbackAddress = &v
	return s
}

func (s *CreateBatchTasksRequest) SetOssPath(v string) *CreateBatchTasksRequest {
	s.OssPath = &v
	return s
}

type CreateBatchTasksResponseBody struct {
	Data      *CreateBatchTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBatchTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchTasksResponseBody) SetData(v *CreateBatchTasksResponseBodyData) *CreateBatchTasksResponseBody {
	s.Data = v
	return s
}

func (s *CreateBatchTasksResponseBody) SetRequestId(v string) *CreateBatchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchTasksResponseBody) SetSuccess(v bool) *CreateBatchTasksResponseBody {
	s.Success = &v
	return s
}

type CreateBatchTasksResponseBodyData struct {
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateBatchTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBatchTasksResponseBodyData) SetId(v int64) *CreateBatchTasksResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateBatchTasksResponseBodyData) SetStatus(v string) *CreateBatchTasksResponseBodyData {
	s.Status = &v
	return s
}

type CreateBatchTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBatchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBatchTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTasksResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchTasksResponse) SetHeaders(v map[string]*string) *CreateBatchTasksResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchTasksResponse) SetStatusCode(v int32) *CreateBatchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchTasksResponse) SetBody(v *CreateBatchTasksResponseBody) *CreateBatchTasksResponse {
	s.Body = v
	return s
}

type CreateDumpMetaResponseBody struct {
	Data      *CreateDumpMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDumpMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDumpMetaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDumpMetaResponseBody) SetData(v *CreateDumpMetaResponseBodyData) *CreateDumpMetaResponseBody {
	s.Data = v
	return s
}

func (s *CreateDumpMetaResponseBody) SetRequestId(v string) *CreateDumpMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDumpMetaResponseBody) SetSuccess(v bool) *CreateDumpMetaResponseBody {
	s.Success = &v
	return s
}

type CreateDumpMetaResponseBodyData struct {
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDumpMetaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDumpMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDumpMetaResponseBodyData) SetId(v int64) *CreateDumpMetaResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateDumpMetaResponseBodyData) SetStatus(v string) *CreateDumpMetaResponseBodyData {
	s.Status = &v
	return s
}

type CreateDumpMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDumpMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDumpMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDumpMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateDumpMetaResponse) SetHeaders(v map[string]*string) *CreateDumpMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateDumpMetaResponse) SetStatusCode(v int32) *CreateDumpMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDumpMetaResponse) SetBody(v *CreateDumpMetaResponseBody) *CreateDumpMetaResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetPicName(v string) *DeleteImageRequest {
	s.PicName = &v
	return s
}

type DeleteImageResponseBody struct {
	Data      *DeleteImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetData(v *DeleteImageResponseBodyData) *DeleteImageResponseBody {
	s.Data = v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetSuccess(v bool) *DeleteImageResponseBody {
	s.Success = &v
	return s
}

type DeleteImageResponseBodyData struct {
	PicNames []*string `json:"PicNames,omitempty" xml:"PicNames,omitempty" type:"Repeated"`
}

func (s DeleteImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBodyData) SetPicNames(v []*string) *DeleteImageResponseBodyData {
	s.PicNames = v
	return s
}

type DeleteImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteImageResponse) SetStatusCode(v int32) *DeleteImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetPicName(v string) *GetImageRequest {
	s.PicName = &v
	return s
}

type GetImageResponseBody struct {
	Data      *GetImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) SetData(v *GetImageResponseBodyData) *GetImageResponseBody {
	s.Data = v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSuccess(v bool) *GetImageResponseBody {
	s.Success = &v
	return s
}

type GetImageResponseBodyData struct {
	Boxes         []*GetImageResponseBodyDataBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	CustomContent *string                          `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr       *int32                           `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	PicName       *string                          `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId     *string                          `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	StrAttr       *string                          `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s GetImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyData) SetBoxes(v []*GetImageResponseBodyDataBoxes) *GetImageResponseBodyData {
	s.Boxes = v
	return s
}

func (s *GetImageResponseBodyData) SetCustomContent(v string) *GetImageResponseBodyData {
	s.CustomContent = &v
	return s
}

func (s *GetImageResponseBodyData) SetIntAttr(v int32) *GetImageResponseBodyData {
	s.IntAttr = &v
	return s
}

func (s *GetImageResponseBodyData) SetPicName(v string) *GetImageResponseBodyData {
	s.PicName = &v
	return s
}

func (s *GetImageResponseBodyData) SetProductId(v string) *GetImageResponseBodyData {
	s.ProductId = &v
	return s
}

func (s *GetImageResponseBodyData) SetStrAttr(v string) *GetImageResponseBodyData {
	s.StrAttr = &v
	return s
}

type GetImageResponseBodyDataBoxes struct {
	Bbox       []*int32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s GetImageResponseBodyDataBoxes) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyDataBoxes) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyDataBoxes) SetBbox(v []*int32) *GetImageResponseBodyDataBoxes {
	s.Bbox = v
	return s
}

func (s *GetImageResponseBodyDataBoxes) SetConfidence(v float32) *GetImageResponseBodyDataBoxes {
	s.Confidence = &v
	return s
}

type GetImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetImageResponse) SetStatusCode(v int32) *GetImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	Data      *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

type GetInstanceResponseBodyData struct {
	Capacity      *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Qps           *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceType   *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCount    *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UtcCreateTime *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
	UtcExpireTime *int64  `json:"UtcExpireTime,omitempty" xml:"UtcExpireTime,omitempty"`
}

func (s GetInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) SetCapacity(v int32) *GetInstanceResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceId(v string) *GetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceName(v string) *GetInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetQps(v string) *GetInstanceResponseBodyData {
	s.Qps = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRegion(v string) *GetInstanceResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetServiceType(v string) *GetInstanceResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetTotalCount(v int64) *GetInstanceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUtcCreateTime(v int64) *GetInstanceResponseBodyData {
	s.UtcCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUtcExpireTime(v int64) *GetInstanceResponseBodyData {
	s.UtcExpireTime = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type ListBatchTasksRequest struct {
	// oss bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// oss path
	OssPath    *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBatchTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTasksRequest) GoString() string {
	return s.String()
}

func (s *ListBatchTasksRequest) SetBucketName(v string) *ListBatchTasksRequest {
	s.BucketName = &v
	return s
}

func (s *ListBatchTasksRequest) SetId(v int64) *ListBatchTasksRequest {
	s.Id = &v
	return s
}

func (s *ListBatchTasksRequest) SetOssPath(v string) *ListBatchTasksRequest {
	s.OssPath = &v
	return s
}

func (s *ListBatchTasksRequest) SetPageNumber(v int32) *ListBatchTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBatchTasksRequest) SetPageSize(v int32) *ListBatchTasksRequest {
	s.PageSize = &v
	return s
}

type ListBatchTasksResponseBody struct {
	Data      *ListBatchTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBatchTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchTasksResponseBody) SetData(v *ListBatchTasksResponseBodyData) *ListBatchTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListBatchTasksResponseBody) SetRequestId(v string) *ListBatchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchTasksResponseBody) SetSuccess(v bool) *ListBatchTasksResponseBody {
	s.Success = &v
	return s
}

type ListBatchTasksResponseBodyData struct {
	Increments []*ListBatchTasksResponseBodyDataIncrements `json:"Increments,omitempty" xml:"Increments,omitempty" type:"Repeated"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListBatchTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBatchTasksResponseBodyData) SetIncrements(v []*ListBatchTasksResponseBodyDataIncrements) *ListBatchTasksResponseBodyData {
	s.Increments = v
	return s
}

func (s *ListBatchTasksResponseBodyData) SetPageNumber(v int32) *ListBatchTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBatchTasksResponseBodyData) SetPageSize(v int32) *ListBatchTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBatchTasksResponseBodyData) SetTotal(v int64) *ListBatchTasksResponseBodyData {
	s.Total = &v
	return s
}

type ListBatchTasksResponseBodyDataIncrements struct {
	BucketName         *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CallbackAddress    *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	CallbackTaskStatus *string `json:"CallbackTaskStatus,omitempty" xml:"CallbackTaskStatus,omitempty"`
	ErrorUrl           *string `json:"ErrorUrl,omitempty" xml:"ErrorUrl,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Msg                *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	OssPath            *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UtcCreateTime      *string `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
	UtcUpdateTime      *string `json:"UtcUpdateTime,omitempty" xml:"UtcUpdateTime,omitempty"`
}

func (s ListBatchTasksResponseBodyDataIncrements) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTasksResponseBodyDataIncrements) GoString() string {
	return s.String()
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetBucketName(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.BucketName = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetCallbackAddress(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.CallbackAddress = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetCallbackTaskStatus(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.CallbackTaskStatus = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetErrorUrl(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.ErrorUrl = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetId(v int64) *ListBatchTasksResponseBodyDataIncrements {
	s.Id = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetMsg(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.Msg = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetOssPath(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.OssPath = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetStatus(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.Status = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetUtcCreateTime(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.UtcCreateTime = &v
	return s
}

func (s *ListBatchTasksResponseBodyDataIncrements) SetUtcUpdateTime(v string) *ListBatchTasksResponseBodyDataIncrements {
	s.UtcUpdateTime = &v
	return s
}

type ListBatchTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBatchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBatchTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTasksResponse) GoString() string {
	return s.String()
}

func (s *ListBatchTasksResponse) SetHeaders(v map[string]*string) *ListBatchTasksResponse {
	s.Headers = v
	return s
}

func (s *ListBatchTasksResponse) SetStatusCode(v int32) *ListBatchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchTasksResponse) SetBody(v *ListBatchTasksResponseBody) *ListBatchTasksResponse {
	s.Body = v
	return s
}

type ListDumpMetaRequest struct {
	Id         *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDumpMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDumpMetaRequest) GoString() string {
	return s.String()
}

func (s *ListDumpMetaRequest) SetId(v int64) *ListDumpMetaRequest {
	s.Id = &v
	return s
}

func (s *ListDumpMetaRequest) SetPageNumber(v int32) *ListDumpMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDumpMetaRequest) SetPageSize(v int32) *ListDumpMetaRequest {
	s.PageSize = &v
	return s
}

type ListDumpMetaResponseBody struct {
	Data      *ListDumpMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDumpMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDumpMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ListDumpMetaResponseBody) SetData(v *ListDumpMetaResponseBodyData) *ListDumpMetaResponseBody {
	s.Data = v
	return s
}

func (s *ListDumpMetaResponseBody) SetRequestId(v string) *ListDumpMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDumpMetaResponseBody) SetSuccess(v bool) *ListDumpMetaResponseBody {
	s.Success = &v
	return s
}

type ListDumpMetaResponseBodyData struct {
	DumpMetaList []*ListDumpMetaResponseBodyDataDumpMetaList `json:"DumpMetaList,omitempty" xml:"DumpMetaList,omitempty" type:"Repeated"`
	PageNumber   *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total        *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDumpMetaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDumpMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDumpMetaResponseBodyData) SetDumpMetaList(v []*ListDumpMetaResponseBodyDataDumpMetaList) *ListDumpMetaResponseBodyData {
	s.DumpMetaList = v
	return s
}

func (s *ListDumpMetaResponseBodyData) SetPageNumber(v int32) *ListDumpMetaResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDumpMetaResponseBodyData) SetPageSize(v int32) *ListDumpMetaResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDumpMetaResponseBodyData) SetTotal(v int64) *ListDumpMetaResponseBodyData {
	s.Total = &v
	return s
}

type ListDumpMetaResponseBodyDataDumpMetaList struct {
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MetaUrl       *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	Msg           *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UtcCreateTime *string `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
	UtcUpdateTime *string `json:"UtcUpdateTime,omitempty" xml:"UtcUpdateTime,omitempty"`
}

func (s ListDumpMetaResponseBodyDataDumpMetaList) String() string {
	return tea.Prettify(s)
}

func (s ListDumpMetaResponseBodyDataDumpMetaList) GoString() string {
	return s.String()
}

func (s *ListDumpMetaResponseBodyDataDumpMetaList) SetId(v int64) *ListDumpMetaResponseBodyDataDumpMetaList {
	s.Id = &v
	return s
}

func (s *ListDumpMetaResponseBodyDataDumpMetaList) SetMetaUrl(v string) *ListDumpMetaResponseBodyDataDumpMetaList {
	s.MetaUrl = &v
	return s
}

func (s *ListDumpMetaResponseBodyDataDumpMetaList) SetMsg(v string) *ListDumpMetaResponseBodyDataDumpMetaList {
	s.Msg = &v
	return s
}

func (s *ListDumpMetaResponseBodyDataDumpMetaList) SetStatus(v string) *ListDumpMetaResponseBodyDataDumpMetaList {
	s.Status = &v
	return s
}

func (s *ListDumpMetaResponseBodyDataDumpMetaList) SetUtcCreateTime(v string) *ListDumpMetaResponseBodyDataDumpMetaList {
	s.UtcCreateTime = &v
	return s
}

func (s *ListDumpMetaResponseBodyDataDumpMetaList) SetUtcUpdateTime(v string) *ListDumpMetaResponseBodyDataDumpMetaList {
	s.UtcUpdateTime = &v
	return s
}

type ListDumpMetaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDumpMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDumpMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDumpMetaResponse) GoString() string {
	return s.String()
}

func (s *ListDumpMetaResponse) SetHeaders(v map[string]*string) *ListDumpMetaResponse {
	s.Headers = v
	return s
}

func (s *ListDumpMetaResponse) SetStatusCode(v int32) *ListDumpMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDumpMetaResponse) SetBody(v *ListDumpMetaResponseBody) *ListDumpMetaResponse {
	s.Body = v
	return s
}

type ListInstanceRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServiceType  *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) SetInstanceName(v string) *ListInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceRequest) SetPageNumber(v int32) *ListInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceRequest) SetPageSize(v int32) *ListInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRequest) SetServiceType(v string) *ListInstanceRequest {
	s.ServiceType = &v
	return s
}

type ListInstanceResponseBody struct {
	Data      *ListInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) SetData(v *ListInstanceResponseBodyData) *ListInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetSuccess(v bool) *ListInstanceResponseBody {
	s.Success = &v
	return s
}

type ListInstanceResponseBodyData struct {
	Instances  []*ListInstanceResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int64                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyData) SetInstances(v []*ListInstanceResponseBodyDataInstances) *ListInstanceResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListInstanceResponseBodyData) SetPageNumber(v int32) *ListInstanceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetPageSize(v int32) *ListInstanceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetTotal(v int64) *ListInstanceResponseBodyData {
	s.Total = &v
	return s
}

type ListInstanceResponseBodyDataInstances struct {
	Capacity      *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Qps           *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceType   *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UtcCreateTime *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
	UtcExpireTime *int64  `json:"UtcExpireTime,omitempty" xml:"UtcExpireTime,omitempty"`
}

func (s ListInstanceResponseBodyDataInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyDataInstances) SetCapacity(v int32) *ListInstanceResponseBodyDataInstances {
	s.Capacity = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetInstanceId(v string) *ListInstanceResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetInstanceName(v string) *ListInstanceResponseBodyDataInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetQps(v string) *ListInstanceResponseBodyDataInstances {
	s.Qps = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetRegion(v string) *ListInstanceResponseBodyDataInstances {
	s.Region = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetServiceType(v string) *ListInstanceResponseBodyDataInstances {
	s.ServiceType = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetStatus(v string) *ListInstanceResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetUtcCreateTime(v int64) *ListInstanceResponseBodyDataInstances {
	s.UtcCreateTime = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetUtcExpireTime(v int64) *ListInstanceResponseBodyDataInstances {
	s.UtcExpireTime = &v
	return s
}

type ListInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetStatusCode(v int32) *ListInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

type ListOssBucketAndPathRequest struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	OssPath    *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s ListOssBucketAndPathRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOssBucketAndPathRequest) GoString() string {
	return s.String()
}

func (s *ListOssBucketAndPathRequest) SetBucketName(v string) *ListOssBucketAndPathRequest {
	s.BucketName = &v
	return s
}

func (s *ListOssBucketAndPathRequest) SetOssPath(v string) *ListOssBucketAndPathRequest {
	s.OssPath = &v
	return s
}

type ListOssBucketAndPathResponseBody struct {
	Data      *ListOssBucketAndPathResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOssBucketAndPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOssBucketAndPathResponseBody) GoString() string {
	return s.String()
}

func (s *ListOssBucketAndPathResponseBody) SetData(v *ListOssBucketAndPathResponseBodyData) *ListOssBucketAndPathResponseBody {
	s.Data = v
	return s
}

func (s *ListOssBucketAndPathResponseBody) SetRequestId(v string) *ListOssBucketAndPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOssBucketAndPathResponseBody) SetSuccess(v bool) *ListOssBucketAndPathResponseBody {
	s.Success = &v
	return s
}

type ListOssBucketAndPathResponseBodyData struct {
	BucketList []*string `json:"BucketList,omitempty" xml:"BucketList,omitempty" type:"Repeated"`
	PathList   []*string `json:"PathList,omitempty" xml:"PathList,omitempty" type:"Repeated"`
}

func (s ListOssBucketAndPathResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListOssBucketAndPathResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOssBucketAndPathResponseBodyData) SetBucketList(v []*string) *ListOssBucketAndPathResponseBodyData {
	s.BucketList = v
	return s
}

func (s *ListOssBucketAndPathResponseBodyData) SetPathList(v []*string) *ListOssBucketAndPathResponseBodyData {
	s.PathList = v
	return s
}

type ListOssBucketAndPathResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOssBucketAndPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOssBucketAndPathResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOssBucketAndPathResponse) GoString() string {
	return s.String()
}

func (s *ListOssBucketAndPathResponse) SetHeaders(v map[string]*string) *ListOssBucketAndPathResponse {
	s.Headers = v
	return s
}

func (s *ListOssBucketAndPathResponse) SetStatusCode(v int32) *ListOssBucketAndPathResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOssBucketAndPathResponse) SetBody(v *ListOssBucketAndPathResponseBody) *ListOssBucketAndPathResponse {
	s.Body = v
	return s
}

type ResetInstanceResponseBody struct {
	Data      *ResetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponseBody) SetData(v *ResetInstanceResponseBodyData) *ResetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ResetInstanceResponseBody) SetRequestId(v string) *ResetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetInstanceResponseBody) SetSuccess(v bool) *ResetInstanceResponseBody {
	s.Success = &v
	return s
}

type ResetInstanceResponseBodyData struct {
	Capacity      *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ServiceType   *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UtcCreateTime *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
	UtcExpireTime *int64  `json:"UtcExpireTime,omitempty" xml:"UtcExpireTime,omitempty"`
}

func (s ResetInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ResetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponseBodyData) SetCapacity(v int32) *ResetInstanceResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *ResetInstanceResponseBodyData) SetInstanceId(v string) *ResetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ResetInstanceResponseBodyData) SetInstanceName(v string) *ResetInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ResetInstanceResponseBodyData) SetServiceType(v string) *ResetInstanceResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *ResetInstanceResponseBodyData) SetStatus(v string) *ResetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ResetInstanceResponseBodyData) SetUtcCreateTime(v int64) *ResetInstanceResponseBodyData {
	s.UtcCreateTime = &v
	return s
}

func (s *ResetInstanceResponseBodyData) SetUtcExpireTime(v int64) *ResetInstanceResponseBodyData {
	s.UtcExpireTime = &v
	return s
}

type ResetInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponse) SetHeaders(v map[string]*string) *ResetInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResetInstanceResponse) SetStatusCode(v int32) *ResetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetInstanceResponse) SetBody(v *ResetInstanceResponseBody) *ResetInstanceResponse {
	s.Body = v
	return s
}

type SearchImageByNameRequest struct {
	Filter    *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Num       *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PicName   *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Start     *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchImageByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByNameRequest) SetFilter(v string) *SearchImageByNameRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByNameRequest) SetNum(v int32) *SearchImageByNameRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByNameRequest) SetPicName(v string) *SearchImageByNameRequest {
	s.PicName = &v
	return s
}

func (s *SearchImageByNameRequest) SetProductId(v string) *SearchImageByNameRequest {
	s.ProductId = &v
	return s
}

func (s *SearchImageByNameRequest) SetStart(v int32) *SearchImageByNameRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByNameRequest) SetText(v string) *SearchImageByNameRequest {
	s.Text = &v
	return s
}

type SearchImageByNameResponseBody struct {
	Data      *SearchImageByNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBody) SetData(v *SearchImageByNameResponseBodyData) *SearchImageByNameResponseBody {
	s.Data = v
	return s
}

func (s *SearchImageByNameResponseBody) SetRequestId(v string) *SearchImageByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetSuccess(v bool) *SearchImageByNameResponseBody {
	s.Success = &v
	return s
}

type SearchImageByNameResponseBodyData struct {
	PicInfos []*SearchImageByNameResponseBodyDataPicInfos `json:"PicInfos,omitempty" xml:"PicInfos,omitempty" type:"Repeated"`
	PicList  []*SearchImageByNameResponseBodyDataPicList  `json:"PicList,omitempty" xml:"PicList,omitempty" type:"Repeated"`
}

func (s SearchImageByNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyData) SetPicInfos(v []*SearchImageByNameResponseBodyDataPicInfos) *SearchImageByNameResponseBodyData {
	s.PicInfos = v
	return s
}

func (s *SearchImageByNameResponseBodyData) SetPicList(v []*SearchImageByNameResponseBodyDataPicList) *SearchImageByNameResponseBodyData {
	s.PicList = v
	return s
}

type SearchImageByNameResponseBodyDataPicInfos struct {
	Bbox       []*int32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s SearchImageByNameResponseBodyDataPicInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyDataPicInfos) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyDataPicInfos) SetBbox(v []*int32) *SearchImageByNameResponseBodyDataPicInfos {
	s.Bbox = v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicInfos) SetConfidence(v float32) *SearchImageByNameResponseBodyDataPicInfos {
	s.Confidence = &v
	return s
}

type SearchImageByNameResponseBodyDataPicList struct {
	CustomContent *string                                                 `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr       *int32                                                  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	OtherBoxes    []*SearchImageByNameResponseBodyDataPicListOtherBoxes   `json:"OtherBoxes,omitempty" xml:"OtherBoxes,omitempty" type:"Repeated"`
	PicName       *string                                                 `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId     *string                                                 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Score         *float32                                                `json:"Score,omitempty" xml:"Score,omitempty"`
	SimilarBoxes  []*SearchImageByNameResponseBodyDataPicListSimilarBoxes `json:"SimilarBoxes,omitempty" xml:"SimilarBoxes,omitempty" type:"Repeated"`
	StrAttr       *string                                                 `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s SearchImageByNameResponseBodyDataPicList) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyDataPicList) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyDataPicList) SetCustomContent(v string) *SearchImageByNameResponseBodyDataPicList {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicList) SetIntAttr(v int32) *SearchImageByNameResponseBodyDataPicList {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicList) SetOtherBoxes(v []*SearchImageByNameResponseBodyDataPicListOtherBoxes) *SearchImageByNameResponseBodyDataPicList {
	s.OtherBoxes = v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicList) SetPicName(v string) *SearchImageByNameResponseBodyDataPicList {
	s.PicName = &v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicList) SetProductId(v string) *SearchImageByNameResponseBodyDataPicList {
	s.ProductId = &v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicList) SetScore(v float32) *SearchImageByNameResponseBodyDataPicList {
	s.Score = &v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicList) SetSimilarBoxes(v []*SearchImageByNameResponseBodyDataPicListSimilarBoxes) *SearchImageByNameResponseBodyDataPicList {
	s.SimilarBoxes = v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicList) SetStrAttr(v string) *SearchImageByNameResponseBodyDataPicList {
	s.StrAttr = &v
	return s
}

type SearchImageByNameResponseBodyDataPicListOtherBoxes struct {
	Bbox       []*int32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s SearchImageByNameResponseBodyDataPicListOtherBoxes) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyDataPicListOtherBoxes) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyDataPicListOtherBoxes) SetBbox(v []*int32) *SearchImageByNameResponseBodyDataPicListOtherBoxes {
	s.Bbox = v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicListOtherBoxes) SetConfidence(v float32) *SearchImageByNameResponseBodyDataPicListOtherBoxes {
	s.Confidence = &v
	return s
}

type SearchImageByNameResponseBodyDataPicListSimilarBoxes struct {
	Bbox       []*int32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Score      *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchImageByNameResponseBodyDataPicListSimilarBoxes) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyDataPicListSimilarBoxes) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyDataPicListSimilarBoxes) SetBbox(v []*int32) *SearchImageByNameResponseBodyDataPicListSimilarBoxes {
	s.Bbox = v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicListSimilarBoxes) SetConfidence(v float32) *SearchImageByNameResponseBodyDataPicListSimilarBoxes {
	s.Confidence = &v
	return s
}

func (s *SearchImageByNameResponseBodyDataPicListSimilarBoxes) SetScore(v float32) *SearchImageByNameResponseBodyDataPicListSimilarBoxes {
	s.Score = &v
	return s
}

type SearchImageByNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchImageByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchImageByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponse) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponse) SetHeaders(v map[string]*string) *SearchImageByNameResponse {
	s.Headers = v
	return s
}

func (s *SearchImageByNameResponse) SetStatusCode(v int32) *SearchImageByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageByNameResponse) SetBody(v *SearchImageByNameResponseBody) *SearchImageByNameResponse {
	s.Body = v
	return s
}

type SearchImageByPicRequest struct {
	Box         *Box    `json:"Box,omitempty" xml:"Box,omitempty"`
	DetectLimit *int32  `json:"DetectLimit,omitempty" xml:"DetectLimit,omitempty"`
	Filter      *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Num         *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PicContent  *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicUrl      *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Start       *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchImageByPicRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByPicRequest) SetBox(v *Box) *SearchImageByPicRequest {
	s.Box = v
	return s
}

func (s *SearchImageByPicRequest) SetDetectLimit(v int32) *SearchImageByPicRequest {
	s.DetectLimit = &v
	return s
}

func (s *SearchImageByPicRequest) SetFilter(v string) *SearchImageByPicRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByPicRequest) SetNum(v int32) *SearchImageByPicRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByPicRequest) SetPicContent(v string) *SearchImageByPicRequest {
	s.PicContent = &v
	return s
}

func (s *SearchImageByPicRequest) SetPicUrl(v string) *SearchImageByPicRequest {
	s.PicUrl = &v
	return s
}

func (s *SearchImageByPicRequest) SetStart(v int32) *SearchImageByPicRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByPicRequest) SetText(v string) *SearchImageByPicRequest {
	s.Text = &v
	return s
}

type SearchImageByPicAdvanceRequest struct {
	Box              *Box      `json:"Box,omitempty" xml:"Box,omitempty"`
	DetectLimit      *int32    `json:"DetectLimit,omitempty" xml:"DetectLimit,omitempty"`
	Filter           *string   `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Num              *int32    `json:"Num,omitempty" xml:"Num,omitempty"`
	PicContentObject io.Reader `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicUrl           *string   `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Start            *int32    `json:"Start,omitempty" xml:"Start,omitempty"`
	Text             *string   `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchImageByPicAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByPicAdvanceRequest) SetBox(v *Box) *SearchImageByPicAdvanceRequest {
	s.Box = v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetDetectLimit(v int32) *SearchImageByPicAdvanceRequest {
	s.DetectLimit = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetFilter(v string) *SearchImageByPicAdvanceRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetNum(v int32) *SearchImageByPicAdvanceRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetPicContentObject(v io.Reader) *SearchImageByPicAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetPicUrl(v string) *SearchImageByPicAdvanceRequest {
	s.PicUrl = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetStart(v int32) *SearchImageByPicAdvanceRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetText(v string) *SearchImageByPicAdvanceRequest {
	s.Text = &v
	return s
}

type SearchImageByPicShrinkRequest struct {
	BoxShrink   *string `json:"Box,omitempty" xml:"Box,omitempty"`
	DetectLimit *int32  `json:"DetectLimit,omitempty" xml:"DetectLimit,omitempty"`
	Filter      *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Num         *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PicContent  *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicUrl      *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Start       *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchImageByPicShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByPicShrinkRequest) SetBoxShrink(v string) *SearchImageByPicShrinkRequest {
	s.BoxShrink = &v
	return s
}

func (s *SearchImageByPicShrinkRequest) SetDetectLimit(v int32) *SearchImageByPicShrinkRequest {
	s.DetectLimit = &v
	return s
}

func (s *SearchImageByPicShrinkRequest) SetFilter(v string) *SearchImageByPicShrinkRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByPicShrinkRequest) SetNum(v int32) *SearchImageByPicShrinkRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByPicShrinkRequest) SetPicContent(v string) *SearchImageByPicShrinkRequest {
	s.PicContent = &v
	return s
}

func (s *SearchImageByPicShrinkRequest) SetPicUrl(v string) *SearchImageByPicShrinkRequest {
	s.PicUrl = &v
	return s
}

func (s *SearchImageByPicShrinkRequest) SetStart(v int32) *SearchImageByPicShrinkRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByPicShrinkRequest) SetText(v string) *SearchImageByPicShrinkRequest {
	s.Text = &v
	return s
}

type SearchImageByPicResponseBody struct {
	Data      *SearchImageByPicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByPicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBody) SetData(v *SearchImageByPicResponseBodyData) *SearchImageByPicResponseBody {
	s.Data = v
	return s
}

func (s *SearchImageByPicResponseBody) SetRequestId(v string) *SearchImageByPicResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetSuccess(v bool) *SearchImageByPicResponseBody {
	s.Success = &v
	return s
}

type SearchImageByPicResponseBodyData struct {
	PicInfos []*SearchImageByPicResponseBodyDataPicInfos `json:"PicInfos,omitempty" xml:"PicInfos,omitempty" type:"Repeated"`
	PicList  []*SearchImageByPicResponseBodyDataPicList  `json:"PicList,omitempty" xml:"PicList,omitempty" type:"Repeated"`
}

func (s SearchImageByPicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyData) SetPicInfos(v []*SearchImageByPicResponseBodyDataPicInfos) *SearchImageByPicResponseBodyData {
	s.PicInfos = v
	return s
}

func (s *SearchImageByPicResponseBodyData) SetPicList(v []*SearchImageByPicResponseBodyDataPicList) *SearchImageByPicResponseBodyData {
	s.PicList = v
	return s
}

type SearchImageByPicResponseBodyDataPicInfos struct {
	Bbox       []*int32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s SearchImageByPicResponseBodyDataPicInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyDataPicInfos) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyDataPicInfos) SetBbox(v []*int32) *SearchImageByPicResponseBodyDataPicInfos {
	s.Bbox = v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicInfos) SetConfidence(v float32) *SearchImageByPicResponseBodyDataPicInfos {
	s.Confidence = &v
	return s
}

type SearchImageByPicResponseBodyDataPicList struct {
	CustomContent *string                                                `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr       *int32                                                 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	OtherBoxes    []*SearchImageByPicResponseBodyDataPicListOtherBoxes   `json:"OtherBoxes,omitempty" xml:"OtherBoxes,omitempty" type:"Repeated"`
	PicName       *string                                                `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId     *string                                                `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Score         *float32                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	SimilarBoxes  []*SearchImageByPicResponseBodyDataPicListSimilarBoxes `json:"SimilarBoxes,omitempty" xml:"SimilarBoxes,omitempty" type:"Repeated"`
	StrAttr       *string                                                `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s SearchImageByPicResponseBodyDataPicList) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyDataPicList) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyDataPicList) SetCustomContent(v string) *SearchImageByPicResponseBodyDataPicList {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicList) SetIntAttr(v int32) *SearchImageByPicResponseBodyDataPicList {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicList) SetOtherBoxes(v []*SearchImageByPicResponseBodyDataPicListOtherBoxes) *SearchImageByPicResponseBodyDataPicList {
	s.OtherBoxes = v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicList) SetPicName(v string) *SearchImageByPicResponseBodyDataPicList {
	s.PicName = &v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicList) SetProductId(v string) *SearchImageByPicResponseBodyDataPicList {
	s.ProductId = &v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicList) SetScore(v float32) *SearchImageByPicResponseBodyDataPicList {
	s.Score = &v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicList) SetSimilarBoxes(v []*SearchImageByPicResponseBodyDataPicListSimilarBoxes) *SearchImageByPicResponseBodyDataPicList {
	s.SimilarBoxes = v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicList) SetStrAttr(v string) *SearchImageByPicResponseBodyDataPicList {
	s.StrAttr = &v
	return s
}

type SearchImageByPicResponseBodyDataPicListOtherBoxes struct {
	Bbox       []*int32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s SearchImageByPicResponseBodyDataPicListOtherBoxes) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyDataPicListOtherBoxes) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyDataPicListOtherBoxes) SetBbox(v []*int32) *SearchImageByPicResponseBodyDataPicListOtherBoxes {
	s.Bbox = v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicListOtherBoxes) SetConfidence(v float32) *SearchImageByPicResponseBodyDataPicListOtherBoxes {
	s.Confidence = &v
	return s
}

type SearchImageByPicResponseBodyDataPicListSimilarBoxes struct {
	Bbox       []*int32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Score      *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchImageByPicResponseBodyDataPicListSimilarBoxes) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyDataPicListSimilarBoxes) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyDataPicListSimilarBoxes) SetBbox(v []*int32) *SearchImageByPicResponseBodyDataPicListSimilarBoxes {
	s.Bbox = v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicListSimilarBoxes) SetConfidence(v float32) *SearchImageByPicResponseBodyDataPicListSimilarBoxes {
	s.Confidence = &v
	return s
}

func (s *SearchImageByPicResponseBodyDataPicListSimilarBoxes) SetScore(v float32) *SearchImageByPicResponseBodyDataPicListSimilarBoxes {
	s.Score = &v
	return s
}

type SearchImageByPicResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchImageByPicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchImageByPicResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponse) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponse) SetHeaders(v map[string]*string) *SearchImageByPicResponse {
	s.Headers = v
	return s
}

func (s *SearchImageByPicResponse) SetStatusCode(v int32) *SearchImageByPicResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageByPicResponse) SetBody(v *SearchImageByPicResponseBody) *SearchImageByPicResponse {
	s.Body = v
	return s
}

type StopBatchTasksResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopBatchTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopBatchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *StopBatchTasksResponseBody) SetRequestId(v string) *StopBatchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopBatchTasksResponseBody) SetSuccess(v bool) *StopBatchTasksResponseBody {
	s.Success = &v
	return s
}

type StopBatchTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopBatchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopBatchTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s StopBatchTasksResponse) GoString() string {
	return s.String()
}

func (s *StopBatchTasksResponse) SetHeaders(v map[string]*string) *StopBatchTasksResponse {
	s.Headers = v
	return s
}

func (s *StopBatchTasksResponse) SetStatusCode(v int32) *StopBatchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *StopBatchTasksResponse) SetBody(v *StopBatchTasksResponseBody) *StopBatchTasksResponse {
	s.Body = v
	return s
}

type StopDumpMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDumpMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDumpMetaResponseBody) GoString() string {
	return s.String()
}

func (s *StopDumpMetaResponseBody) SetRequestId(v string) *StopDumpMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDumpMetaResponseBody) SetSuccess(v bool) *StopDumpMetaResponseBody {
	s.Success = &v
	return s
}

type StopDumpMetaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopDumpMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDumpMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDumpMetaResponse) GoString() string {
	return s.String()
}

func (s *StopDumpMetaResponse) SetHeaders(v map[string]*string) *StopDumpMetaResponse {
	s.Headers = v
	return s
}

func (s *StopDumpMetaResponse) SetStatusCode(v int32) *StopDumpMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDumpMetaResponse) SetBody(v *StopDumpMetaResponseBody) *StopDumpMetaResponse {
	s.Body = v
	return s
}

type UpdateImageRequest struct {
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr       *int32  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	PicName       *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	StrAttr       *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s UpdateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) SetCustomContent(v string) *UpdateImageRequest {
	s.CustomContent = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr(v int32) *UpdateImageRequest {
	s.IntAttr = &v
	return s
}

func (s *UpdateImageRequest) SetPicName(v string) *UpdateImageRequest {
	s.PicName = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr(v string) *UpdateImageRequest {
	s.StrAttr = &v
	return s
}

type UpdateImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) SetRequestId(v string) *UpdateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageResponseBody) SetSuccess(v bool) *UpdateImageResponseBody {
	s.Success = &v
	return s
}

type UpdateImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateImageResponse) SetStatusCode(v int32) *UpdateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageResponse) SetBody(v *UpdateImageResponseBody) *UpdateImageResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aisearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddImageWithOptions(InstanceName *string, ProductId *string, tmpReq *AddImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Boxes)) {
		request.BoxesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Boxes, tea.String("Boxes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomContent)) {
		query["CustomContent"] = request.CustomContent
	}

	if !tea.BoolValue(util.IsUnset(request.DetectLimit)) {
		query["DetectLimit"] = request.DetectLimit
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr)) {
		query["IntAttr"] = request.IntAttr
	}

	if !tea.BoolValue(util.IsUnset(request.PicContent)) {
		query["PicContent"] = request.PicContent
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		query["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		query["PicUrl"] = request.PicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr)) {
		query["StrAttr"] = request.StrAttr
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BoxesShrink)) {
		body["Boxes"] = request.BoxesShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImage"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/entity/" + tea.StringValue(openapiutil.GetEncodeParam(ProductId)) + "/image"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddImage(InstanceName *string, ProductId *string, request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(InstanceName, ProductId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddImageAdvance(InstanceName *string, ProductId *string, request *AddImageAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("aisearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	addImageReq := &AddImageRequest{}
	openapiutil.Convert(request, addImageReq)
	if !tea.BoolValue(util.IsUnset(request.PicContentObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.PicContentObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		addImageReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	addImageResp, _err := client.AddImageWithOptions(InstanceName, ProductId, addImageReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addImageResp
	return _result, _err
}

func (client *Client) CheckOssIncrementMetaExistWithOptions(request *CheckOssIncrementMetaExistRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckOssIncrementMetaExistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OssPath)) {
		query["OssPath"] = request.OssPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckOssIncrementMetaExist"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/oss/check-increment-metafile"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckOssIncrementMetaExistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckOssIncrementMetaExist(request *CheckOssIncrementMetaExistRequest) (_result *CheckOssIncrementMetaExistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckOssIncrementMetaExistResponse{}
	_body, _err := client.CheckOssIncrementMetaExistWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBatchTasksWithOptions(InstanceName *string, request *CreateBatchTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBatchTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackAddress)) {
		query["CallbackAddress"] = request.CallbackAddress
	}

	if !tea.BoolValue(util.IsUnset(request.OssPath)) {
		query["OssPath"] = request.OssPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBatchTasks"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/batch-task"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBatchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBatchTasks(InstanceName *string, request *CreateBatchTasksRequest) (_result *CreateBatchTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBatchTasksResponse{}
	_body, _err := client.CreateBatchTasksWithOptions(InstanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDumpMetaWithOptions(InstanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDumpMetaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDumpMeta"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/dump-meta"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDumpMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDumpMeta(InstanceName *string) (_result *CreateDumpMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDumpMetaResponse{}
	_body, _err := client.CreateDumpMetaWithOptions(InstanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(InstanceName *string, ProductId *string, request *DeleteImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		query["PicName"] = request.PicName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImage"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/entity/" + tea.StringValue(openapiutil.GetEncodeParam(ProductId)) + "/image"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(InstanceName *string, ProductId *string, request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(InstanceName, ProductId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageWithOptions(InstanceName *string, ProductId *string, request *GetImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		query["PicName"] = request.PicName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImage"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/entity/" + tea.StringValue(openapiutil.GetEncodeParam(ProductId)) + "/image"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImage(InstanceName *string, ProductId *string, request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(InstanceName, ProductId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(InstanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(InstanceName *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(InstanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBatchTasksWithOptions(InstanceName *string, request *ListBatchTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListBatchTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OssPath)) {
		query["OssPath"] = request.OssPath
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBatchTasks"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/batch-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBatchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBatchTasks(InstanceName *string, request *ListBatchTasksRequest) (_result *ListBatchTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBatchTasksResponse{}
	_body, _err := client.ListBatchTasksWithOptions(InstanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDumpMetaWithOptions(InstanceName *string, request *ListDumpMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDumpMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDumpMeta"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/dump-metas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDumpMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDumpMeta(InstanceName *string, request *ListDumpMetaRequest) (_result *ListDumpMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDumpMetaResponse{}
	_body, _err := client.ListDumpMetaWithOptions(InstanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstance"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOssBucketAndPathWithOptions(request *ListOssBucketAndPathRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOssBucketAndPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssPath)) {
		query["OssPath"] = request.OssPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOssBucketAndPath"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/oss/buckets-and-path"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOssBucketAndPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOssBucketAndPath(request *ListOssBucketAndPathRequest) (_result *ListOssBucketAndPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOssBucketAndPathResponse{}
	_body, _err := client.ListOssBucketAndPathWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetInstanceWithOptions(InstanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ResetInstance"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/reset"),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetInstance(InstanceName *string) (_result *ResetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetInstanceResponse{}
	_body, _err := client.ResetInstanceWithOptions(InstanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageByNameWithOptions(InstanceName *string, request *SearchImageByNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchImageByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		query["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchImageByName"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/search-by-name"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchImageByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchImageByName(InstanceName *string, request *SearchImageByNameRequest) (_result *SearchImageByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchImageByNameResponse{}
	_body, _err := client.SearchImageByNameWithOptions(InstanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageByPicWithOptions(InstanceName *string, tmpReq *SearchImageByPicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchImageByPicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Box)) {
		request.BoxShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Box, tea.String("Box"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BoxShrink)) {
		query["Box"] = request.BoxShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DetectLimit)) {
		query["DetectLimit"] = request.DetectLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PicContent)) {
		query["PicContent"] = request.PicContent
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		query["PicUrl"] = request.PicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchImageByPic"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/search-by-pic"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchImageByPicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchImageByPic(InstanceName *string, request *SearchImageByPicRequest) (_result *SearchImageByPicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchImageByPicResponse{}
	_body, _err := client.SearchImageByPicWithOptions(InstanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageByPicAdvance(InstanceName *string, request *SearchImageByPicAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("aisearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	searchImageByPicReq := &SearchImageByPicRequest{}
	openapiutil.Convert(request, searchImageByPicReq)
	if !tea.BoolValue(util.IsUnset(request.PicContentObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.PicContentObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		searchImageByPicReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	searchImageByPicResp, _err := client.SearchImageByPicWithOptions(InstanceName, searchImageByPicReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchImageByPicResp
	return _result, _err
}

func (client *Client) StopBatchTasksWithOptions(InstanceName *string, Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopBatchTasksResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopBatchTasks"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/batch-task/" + tea.StringValue(openapiutil.GetEncodeParam(Id)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopBatchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopBatchTasks(InstanceName *string, Id *string) (_result *StopBatchTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopBatchTasksResponse{}
	_body, _err := client.StopBatchTasksWithOptions(InstanceName, Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDumpMetaWithOptions(InstanceName *string, Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopDumpMetaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopDumpMeta"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/dump-meta/" + tea.StringValue(openapiutil.GetEncodeParam(Id)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDumpMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDumpMeta(InstanceName *string, Id *string) (_result *StopDumpMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopDumpMetaResponse{}
	_body, _err := client.StopDumpMetaWithOptions(InstanceName, Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateImageWithOptions(InstanceName *string, ProductId *string, request *UpdateImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomContent)) {
		query["CustomContent"] = request.CustomContent
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr)) {
		query["IntAttr"] = request.IntAttr
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		query["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr)) {
		query["StrAttr"] = request.StrAttr
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImage"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/entity/" + tea.StringValue(openapiutil.GetEncodeParam(ProductId)) + "/image"),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateImage(InstanceName *string, ProductId *string, request *UpdateImageRequest) (_result *UpdateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateImageResponse{}
	_body, _err := client.UpdateImageWithOptions(InstanceName, ProductId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
