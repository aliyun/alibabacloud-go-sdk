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

type AddImageRequest struct {
	CategoryId    *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop          *bool   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IntAttr       *int32  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	IntAttr2      *int32  `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	PicContent    *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicName       *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId     *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	StrAttr       *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	StrAttr2      *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
}

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetCategoryId(v int32) *AddImageRequest {
	s.CategoryId = &v
	return s
}

func (s *AddImageRequest) SetCrop(v bool) *AddImageRequest {
	s.Crop = &v
	return s
}

func (s *AddImageRequest) SetCustomContent(v string) *AddImageRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageRequest) SetInstanceName(v string) *AddImageRequest {
	s.InstanceName = &v
	return s
}

func (s *AddImageRequest) SetIntAttr(v int32) *AddImageRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageRequest) SetIntAttr2(v int32) *AddImageRequest {
	s.IntAttr2 = &v
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

func (s *AddImageRequest) SetProductId(v string) *AddImageRequest {
	s.ProductId = &v
	return s
}

func (s *AddImageRequest) SetRegion(v string) *AddImageRequest {
	s.Region = &v
	return s
}

func (s *AddImageRequest) SetStrAttr(v string) *AddImageRequest {
	s.StrAttr = &v
	return s
}

func (s *AddImageRequest) SetStrAttr2(v string) *AddImageRequest {
	s.StrAttr2 = &v
	return s
}

type AddImageAdvanceRequest struct {
	CategoryId       *int32    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop             *bool     `json:"Crop,omitempty" xml:"Crop,omitempty"`
	CustomContent    *string   `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	InstanceName     *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IntAttr          *int32    `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	IntAttr2         *int32    `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	PicContentObject io.Reader `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicName          *string   `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId        *string   `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Region           *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	StrAttr          *string   `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	StrAttr2         *string   `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
}

func (s AddImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddImageAdvanceRequest) SetCategoryId(v int32) *AddImageAdvanceRequest {
	s.CategoryId = &v
	return s
}

func (s *AddImageAdvanceRequest) SetCrop(v bool) *AddImageAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *AddImageAdvanceRequest) SetCustomContent(v string) *AddImageAdvanceRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageAdvanceRequest) SetInstanceName(v string) *AddImageAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *AddImageAdvanceRequest) SetIntAttr(v int32) *AddImageAdvanceRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageAdvanceRequest) SetIntAttr2(v int32) *AddImageAdvanceRequest {
	s.IntAttr2 = &v
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

func (s *AddImageAdvanceRequest) SetProductId(v string) *AddImageAdvanceRequest {
	s.ProductId = &v
	return s
}

func (s *AddImageAdvanceRequest) SetRegion(v string) *AddImageAdvanceRequest {
	s.Region = &v
	return s
}

func (s *AddImageAdvanceRequest) SetStrAttr(v string) *AddImageAdvanceRequest {
	s.StrAttr = &v
	return s
}

func (s *AddImageAdvanceRequest) SetStrAttr2(v string) *AddImageAdvanceRequest {
	s.StrAttr2 = &v
	return s
}

type AddImageResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PicInfo   *AddImageResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetCode(v int32) *AddImageResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageResponseBody) SetMessage(v string) *AddImageResponseBody {
	s.Message = &v
	return s
}

func (s *AddImageResponseBody) SetPicInfo(v *AddImageResponseBodyPicInfo) *AddImageResponseBody {
	s.PicInfo = v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetSuccess(v bool) *AddImageResponseBody {
	s.Success = &v
	return s
}

type AddImageResponseBodyPicInfo struct {
	CategoryId *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s AddImageResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *AddImageResponseBodyPicInfo) SetCategoryId(v int32) *AddImageResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *AddImageResponseBodyPicInfo) SetRegion(v string) *AddImageResponseBodyPicInfo {
	s.Region = &v
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

type DeleteImageRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PicName      *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId    *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetInstanceName(v string) *DeleteImageRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteImageRequest) SetPicName(v string) *DeleteImageRequest {
	s.PicName = &v
	return s
}

func (s *DeleteImageRequest) SetProductId(v string) *DeleteImageRequest {
	s.ProductId = &v
	return s
}

type DeleteImageResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetCode(v int32) *DeleteImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageResponseBody) SetData(v *DeleteImageResponseBodyData) *DeleteImageResponseBody {
	s.Data = v
	return s
}

func (s *DeleteImageResponseBody) SetMessage(v string) *DeleteImageResponseBody {
	s.Message = &v
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

type DetailRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailRequest) GoString() string {
	return s.String()
}

func (s *DetailRequest) SetInstanceName(v string) *DetailRequest {
	s.InstanceName = &v
	return s
}

type DetailResponseBody struct {
	Instance  *DetailResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailResponseBody) GoString() string {
	return s.String()
}

func (s *DetailResponseBody) SetInstance(v *DetailResponseBodyInstance) *DetailResponseBody {
	s.Instance = v
	return s
}

func (s *DetailResponseBody) SetRequestId(v string) *DetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailResponseBody) SetSuccess(v bool) *DetailResponseBody {
	s.Success = &v
	return s
}

type DetailResponseBodyInstance struct {
	Capacity      *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Qps           *int32  `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceType   *int32  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	TotalCount    *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UtcCreate     *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	UtcExpireTime *string `json:"UtcExpireTime,omitempty" xml:"UtcExpireTime,omitempty"`
}

func (s DetailResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s DetailResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DetailResponseBodyInstance) SetCapacity(v int32) *DetailResponseBodyInstance {
	s.Capacity = &v
	return s
}

func (s *DetailResponseBodyInstance) SetName(v string) *DetailResponseBodyInstance {
	s.Name = &v
	return s
}

func (s *DetailResponseBodyInstance) SetQps(v int32) *DetailResponseBodyInstance {
	s.Qps = &v
	return s
}

func (s *DetailResponseBodyInstance) SetRegion(v string) *DetailResponseBodyInstance {
	s.Region = &v
	return s
}

func (s *DetailResponseBodyInstance) SetServiceType(v int32) *DetailResponseBodyInstance {
	s.ServiceType = &v
	return s
}

func (s *DetailResponseBodyInstance) SetTotalCount(v int64) *DetailResponseBodyInstance {
	s.TotalCount = &v
	return s
}

func (s *DetailResponseBodyInstance) SetUtcCreate(v string) *DetailResponseBodyInstance {
	s.UtcCreate = &v
	return s
}

func (s *DetailResponseBodyInstance) SetUtcExpireTime(v string) *DetailResponseBodyInstance {
	s.UtcExpireTime = &v
	return s
}

type DetailResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DetailResponse) GoString() string {
	return s.String()
}

func (s *DetailResponse) SetHeaders(v map[string]*string) *DetailResponse {
	s.Headers = v
	return s
}

func (s *DetailResponse) SetStatusCode(v int32) *DetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailResponse) SetBody(v *DetailResponseBody) *DetailResponse {
	s.Body = v
	return s
}

type DumpMetaRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DumpMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaRequest) GoString() string {
	return s.String()
}

func (s *DumpMetaRequest) SetInstanceName(v string) *DumpMetaRequest {
	s.InstanceName = &v
	return s
}

type DumpMetaResponseBody struct {
	Data      *DumpMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DumpMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DumpMetaResponseBody) SetData(v *DumpMetaResponseBodyData) *DumpMetaResponseBody {
	s.Data = v
	return s
}

func (s *DumpMetaResponseBody) SetRequestId(v string) *DumpMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DumpMetaResponseBody) SetSuccess(v bool) *DumpMetaResponseBody {
	s.Success = &v
	return s
}

type DumpMetaResponseBodyData struct {
	DumpMetaStatus *string `json:"DumpMetaStatus,omitempty" xml:"DumpMetaStatus,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DumpMetaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *DumpMetaResponseBodyData) SetDumpMetaStatus(v string) *DumpMetaResponseBodyData {
	s.DumpMetaStatus = &v
	return s
}

func (s *DumpMetaResponseBodyData) SetId(v string) *DumpMetaResponseBodyData {
	s.Id = &v
	return s
}

type DumpMetaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DumpMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DumpMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaResponse) GoString() string {
	return s.String()
}

func (s *DumpMetaResponse) SetHeaders(v map[string]*string) *DumpMetaResponse {
	s.Headers = v
	return s
}

func (s *DumpMetaResponse) SetStatusCode(v int32) *DumpMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DumpMetaResponse) SetBody(v *DumpMetaResponseBody) *DumpMetaResponse {
	s.Body = v
	return s
}

type DumpMetaListRequest struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DumpMetaListRequest) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListRequest) GoString() string {
	return s.String()
}

func (s *DumpMetaListRequest) SetId(v int64) *DumpMetaListRequest {
	s.Id = &v
	return s
}

func (s *DumpMetaListRequest) SetInstanceName(v string) *DumpMetaListRequest {
	s.InstanceName = &v
	return s
}

func (s *DumpMetaListRequest) SetPageNumber(v int32) *DumpMetaListRequest {
	s.PageNumber = &v
	return s
}

func (s *DumpMetaListRequest) SetPageSize(v int32) *DumpMetaListRequest {
	s.PageSize = &v
	return s
}

type DumpMetaListResponseBody struct {
	Data      *DumpMetaListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DumpMetaListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBody) SetData(v *DumpMetaListResponseBodyData) *DumpMetaListResponseBody {
	s.Data = v
	return s
}

func (s *DumpMetaListResponseBody) SetRequestId(v string) *DumpMetaListResponseBody {
	s.RequestId = &v
	return s
}

type DumpMetaListResponseBodyData struct {
	DumpMetaList []*DumpMetaListResponseBodyDataDumpMetaList `json:"DumpMetaList,omitempty" xml:"DumpMetaList,omitempty" type:"Repeated"`
	PageNumber   *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DumpMetaListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBodyData) SetDumpMetaList(v []*DumpMetaListResponseBodyDataDumpMetaList) *DumpMetaListResponseBodyData {
	s.DumpMetaList = v
	return s
}

func (s *DumpMetaListResponseBodyData) SetPageNumber(v int32) *DumpMetaListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DumpMetaListResponseBodyData) SetPageSize(v int32) *DumpMetaListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DumpMetaListResponseBodyData) SetTotalCount(v int64) *DumpMetaListResponseBodyData {
	s.TotalCount = &v
	return s
}

type DumpMetaListResponseBodyDataDumpMetaList struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MetaUrl     *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	Msg         *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UtcCreate   *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	UtcModified *int64  `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
}

func (s DumpMetaListResponseBodyDataDumpMetaList) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListResponseBodyDataDumpMetaList) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetCode(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Code = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetId(v int64) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Id = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetMetaUrl(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.MetaUrl = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetMsg(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Msg = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetStatus(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Status = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetUtcCreate(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.UtcCreate = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetUtcModified(v int64) *DumpMetaListResponseBodyDataDumpMetaList {
	s.UtcModified = &v
	return s
}

type DumpMetaListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DumpMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DumpMetaListResponse) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListResponse) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponse) SetHeaders(v map[string]*string) *DumpMetaListResponse {
	s.Headers = v
	return s
}

func (s *DumpMetaListResponse) SetStatusCode(v int32) *DumpMetaListResponse {
	s.StatusCode = &v
	return s
}

func (s *DumpMetaListResponse) SetBody(v *DumpMetaListResponseBody) *DumpMetaListResponse {
	s.Body = v
	return s
}

type IncreaseInstanceRequest struct {
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CallbackAddress *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s IncreaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s IncreaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceRequest) SetBucketName(v string) *IncreaseInstanceRequest {
	s.BucketName = &v
	return s
}

func (s *IncreaseInstanceRequest) SetCallbackAddress(v string) *IncreaseInstanceRequest {
	s.CallbackAddress = &v
	return s
}

func (s *IncreaseInstanceRequest) SetInstanceName(v string) *IncreaseInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *IncreaseInstanceRequest) SetPath(v string) *IncreaseInstanceRequest {
	s.Path = &v
	return s
}

type IncreaseInstanceResponseBody struct {
	Data      *IncreaseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IncreaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IncreaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponseBody) SetData(v *IncreaseInstanceResponseBodyData) *IncreaseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *IncreaseInstanceResponseBody) SetRequestId(v string) *IncreaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *IncreaseInstanceResponseBody) SetSuccess(v bool) *IncreaseInstanceResponseBody {
	s.Success = &v
	return s
}

type IncreaseInstanceResponseBodyData struct {
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IncrementStatus *string `json:"IncrementStatus,omitempty" xml:"IncrementStatus,omitempty"`
}

func (s IncreaseInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IncreaseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponseBodyData) SetId(v string) *IncreaseInstanceResponseBodyData {
	s.Id = &v
	return s
}

func (s *IncreaseInstanceResponseBodyData) SetIncrementStatus(v string) *IncreaseInstanceResponseBodyData {
	s.IncrementStatus = &v
	return s
}

type IncreaseInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IncreaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IncreaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s IncreaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponse) SetHeaders(v map[string]*string) *IncreaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *IncreaseInstanceResponse) SetStatusCode(v int32) *IncreaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *IncreaseInstanceResponse) SetBody(v *IncreaseInstanceResponseBody) *IncreaseInstanceResponse {
	s.Body = v
	return s
}

type IncreaseListRequest struct {
	BucketName   *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s IncreaseListRequest) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListRequest) GoString() string {
	return s.String()
}

func (s *IncreaseListRequest) SetBucketName(v string) *IncreaseListRequest {
	s.BucketName = &v
	return s
}

func (s *IncreaseListRequest) SetId(v int64) *IncreaseListRequest {
	s.Id = &v
	return s
}

func (s *IncreaseListRequest) SetInstanceName(v string) *IncreaseListRequest {
	s.InstanceName = &v
	return s
}

func (s *IncreaseListRequest) SetPageNumber(v int32) *IncreaseListRequest {
	s.PageNumber = &v
	return s
}

func (s *IncreaseListRequest) SetPageSize(v int32) *IncreaseListRequest {
	s.PageSize = &v
	return s
}

func (s *IncreaseListRequest) SetPath(v string) *IncreaseListRequest {
	s.Path = &v
	return s
}

type IncreaseListResponseBody struct {
	Data      *IncreaseListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IncreaseListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponseBody) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBody) SetData(v *IncreaseListResponseBodyData) *IncreaseListResponseBody {
	s.Data = v
	return s
}

func (s *IncreaseListResponseBody) SetRequestId(v string) *IncreaseListResponseBody {
	s.RequestId = &v
	return s
}

type IncreaseListResponseBodyData struct {
	Increments *IncreaseListResponseBodyDataIncrements `json:"Increments,omitempty" xml:"Increments,omitempty" type:"Struct"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s IncreaseListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponseBodyData) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyData) SetIncrements(v *IncreaseListResponseBodyDataIncrements) *IncreaseListResponseBodyData {
	s.Increments = v
	return s
}

func (s *IncreaseListResponseBodyData) SetPageNumber(v int32) *IncreaseListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *IncreaseListResponseBodyData) SetPageSize(v int32) *IncreaseListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *IncreaseListResponseBodyData) SetTotalCount(v int64) *IncreaseListResponseBodyData {
	s.TotalCount = &v
	return s
}

type IncreaseListResponseBodyDataIncrements struct {
	Instance []*IncreaseListResponseBodyDataIncrementsInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s IncreaseListResponseBodyDataIncrements) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponseBodyDataIncrements) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyDataIncrements) SetInstance(v []*IncreaseListResponseBodyDataIncrementsInstance) *IncreaseListResponseBodyDataIncrements {
	s.Instance = v
	return s
}

type IncreaseListResponseBodyDataIncrementsInstance struct {
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CallbackAddress *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorUrl        *string `json:"ErrorUrl,omitempty" xml:"ErrorUrl,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Msg             *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UtcCreate       *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	UtcModified     *int64  `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
}

func (s IncreaseListResponseBodyDataIncrementsInstance) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponseBodyDataIncrementsInstance) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetBucketName(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.BucketName = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetCallbackAddress(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.CallbackAddress = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetCode(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Code = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetErrorUrl(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.ErrorUrl = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetId(v int64) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Id = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetMsg(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Msg = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetPath(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Path = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetStatus(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Status = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetUtcCreate(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.UtcCreate = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetUtcModified(v int64) *IncreaseListResponseBodyDataIncrementsInstance {
	s.UtcModified = &v
	return s
}

type IncreaseListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IncreaseListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IncreaseListResponse) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponse) GoString() string {
	return s.String()
}

func (s *IncreaseListResponse) SetHeaders(v map[string]*string) *IncreaseListResponse {
	s.Headers = v
	return s
}

func (s *IncreaseListResponse) SetStatusCode(v int32) *IncreaseListResponse {
	s.StatusCode = &v
	return s
}

func (s *IncreaseListResponse) SetBody(v *IncreaseListResponseBody) *IncreaseListResponse {
	s.Body = v
	return s
}

type SearchImageByNameRequest struct {
	CategoryId   *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Num          *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PicName      *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId    *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Start        *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByNameRequest) SetCategoryId(v int32) *SearchImageByNameRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameRequest) SetFilter(v string) *SearchImageByNameRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByNameRequest) SetInstanceName(v string) *SearchImageByNameRequest {
	s.InstanceName = &v
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

type SearchImageByNameResponseBody struct {
	Auctions  []*SearchImageByNameResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Head      *SearchImageByNameResponseBodyHead       `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	Msg       *string                                  `json:"Msg,omitempty" xml:"Msg,omitempty"`
	PicInfo   *SearchImageByNameResponseBodyPicInfo    `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBody) SetAuctions(v []*SearchImageByNameResponseBodyAuctions) *SearchImageByNameResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByNameResponseBody) SetCode(v int32) *SearchImageByNameResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetHead(v *SearchImageByNameResponseBodyHead) *SearchImageByNameResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByNameResponseBody) SetMsg(v string) *SearchImageByNameResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetPicInfo(v *SearchImageByNameResponseBodyPicInfo) *SearchImageByNameResponseBody {
	s.PicInfo = v
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

type SearchImageByNameResponseBodyAuctions struct {
	CategoryId     *int32   `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CustomContent  *string  `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr        *int32   `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	IntAttr2       *int32   `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	PicName        *string  `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId      *string  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	SortExprValues *string  `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	StrAttr        *string  `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	StrAttr2       *string  `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
}

func (s SearchImageByNameResponseBodyAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByNameResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetCustomContent(v string) *SearchImageByNameResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr2(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetPicName(v string) *SearchImageByNameResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetProductId(v string) *SearchImageByNameResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetScore(v float32) *SearchImageByNameResponseBodyAuctions {
	s.Score = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetSortExprValues(v string) *SearchImageByNameResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr2(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

type SearchImageByNameResponseBodyHead struct {
	DocsFound  *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	SearchTime *int32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
}

func (s SearchImageByNameResponseBodyHead) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyHead) SetDocsFound(v int32) *SearchImageByNameResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchImageByNameResponseBodyHead) SetDocsReturn(v int32) *SearchImageByNameResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchImageByNameResponseBodyHead) SetSearchTime(v int32) *SearchImageByNameResponseBodyHead {
	s.SearchTime = &v
	return s
}

type SearchImageByNameResponseBodyPicInfo struct {
	AllCategories []*SearchImageByNameResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	CategoryId    *int32                                               `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	MultiRegion   []*SearchImageByNameResponseBodyPicInfoMultiRegion   `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	Region        *string                                              `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfo) SetAllCategories(v []*SearchImageByNameResponseBodyPicInfoAllCategories) *SearchImageByNameResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetCategoryId(v int32) *SearchImageByNameResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetMultiRegion(v []*SearchImageByNameResponseBodyPicInfoMultiRegion) *SearchImageByNameResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetRegion(v string) *SearchImageByNameResponseBodyPicInfo {
	s.Region = &v
	return s
}

type SearchImageByNameResponseBodyPicInfoAllCategories struct {
	Id   *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfoAllCategories) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByNameResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByNameResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

type SearchImageByNameResponseBodyPicInfoMultiRegion struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfoMultiRegion) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfoMultiRegion) SetRegion(v string) *SearchImageByNameResponseBodyPicInfoMultiRegion {
	s.Region = &v
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
	CategoryId   *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop         *bool   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Num          *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PicContent   *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Start        *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByPicRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByPicRequest) SetCategoryId(v int32) *SearchImageByPicRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicRequest) SetCrop(v bool) *SearchImageByPicRequest {
	s.Crop = &v
	return s
}

func (s *SearchImageByPicRequest) SetFilter(v string) *SearchImageByPicRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByPicRequest) SetInstanceName(v string) *SearchImageByPicRequest {
	s.InstanceName = &v
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

func (s *SearchImageByPicRequest) SetRegion(v string) *SearchImageByPicRequest {
	s.Region = &v
	return s
}

func (s *SearchImageByPicRequest) SetStart(v int32) *SearchImageByPicRequest {
	s.Start = &v
	return s
}

type SearchImageByPicAdvanceRequest struct {
	CategoryId       *int32    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop             *bool     `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Filter           *string   `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceName     *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Num              *int32    `json:"Num,omitempty" xml:"Num,omitempty"`
	PicContentObject io.Reader `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	Region           *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	Start            *int32    `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByPicAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByPicAdvanceRequest) SetCategoryId(v int32) *SearchImageByPicAdvanceRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetCrop(v bool) *SearchImageByPicAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetFilter(v string) *SearchImageByPicAdvanceRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetInstanceName(v string) *SearchImageByPicAdvanceRequest {
	s.InstanceName = &v
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

func (s *SearchImageByPicAdvanceRequest) SetRegion(v string) *SearchImageByPicAdvanceRequest {
	s.Region = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetStart(v int32) *SearchImageByPicAdvanceRequest {
	s.Start = &v
	return s
}

type SearchImageByPicResponseBody struct {
	Auctions  []*SearchImageByPicResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Head      *SearchImageByPicResponseBodyHead       `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	Msg       *string                                 `json:"Msg,omitempty" xml:"Msg,omitempty"`
	PicInfo   *SearchImageByPicResponseBodyPicInfo    `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByPicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBody) SetAuctions(v []*SearchImageByPicResponseBodyAuctions) *SearchImageByPicResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByPicResponseBody) SetCode(v int32) *SearchImageByPicResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetHead(v *SearchImageByPicResponseBodyHead) *SearchImageByPicResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByPicResponseBody) SetMsg(v string) *SearchImageByPicResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetPicInfo(v *SearchImageByPicResponseBodyPicInfo) *SearchImageByPicResponseBody {
	s.PicInfo = v
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

type SearchImageByPicResponseBodyAuctions struct {
	CategoryId     *int32   `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CustomContent  *string  `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr        *int32   `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	IntAttr2       *int32   `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	PicName        *string  `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId      *string  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	SortExprValues *string  `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	StrAttr        *string  `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	StrAttr2       *string  `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
}

func (s SearchImageByPicResponseBodyAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByPicResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetCustomContent(v string) *SearchImageByPicResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr2(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetPicName(v string) *SearchImageByPicResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetProductId(v string) *SearchImageByPicResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetScore(v float32) *SearchImageByPicResponseBodyAuctions {
	s.Score = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetSortExprValues(v string) *SearchImageByPicResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr2(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

type SearchImageByPicResponseBodyHead struct {
	DocsFound  *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	SearchTime *int32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
}

func (s SearchImageByPicResponseBodyHead) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyHead) SetDocsFound(v int32) *SearchImageByPicResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchImageByPicResponseBodyHead) SetDocsReturn(v int32) *SearchImageByPicResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchImageByPicResponseBodyHead) SetSearchTime(v int32) *SearchImageByPicResponseBodyHead {
	s.SearchTime = &v
	return s
}

type SearchImageByPicResponseBodyPicInfo struct {
	AllCategories []*SearchImageByPicResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	CategoryId    *int32                                              `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	MultiRegion   []*SearchImageByPicResponseBodyPicInfoMultiRegion   `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	Region        *string                                             `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfo) SetAllCategories(v []*SearchImageByPicResponseBodyPicInfoAllCategories) *SearchImageByPicResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetCategoryId(v int32) *SearchImageByPicResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetMultiRegion(v []*SearchImageByPicResponseBodyPicInfoMultiRegion) *SearchImageByPicResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetRegion(v string) *SearchImageByPicResponseBodyPicInfo {
	s.Region = &v
	return s
}

type SearchImageByPicResponseBodyPicInfoAllCategories struct {
	Id   *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfoAllCategories) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByPicResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByPicResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

type SearchImageByPicResponseBodyPicInfoMultiRegion struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfoMultiRegion) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfoMultiRegion) SetRegion(v string) *SearchImageByPicResponseBodyPicInfoMultiRegion {
	s.Region = &v
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

type UpdateImageRequest struct {
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IntAttr       *int32  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	IntAttr2      *int32  `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	PicName       *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId     *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	StrAttr       *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	StrAttr2      *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
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

func (s *UpdateImageRequest) SetInstanceName(v string) *UpdateImageRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr(v int32) *UpdateImageRequest {
	s.IntAttr = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr2(v int32) *UpdateImageRequest {
	s.IntAttr2 = &v
	return s
}

func (s *UpdateImageRequest) SetPicName(v string) *UpdateImageRequest {
	s.PicName = &v
	return s
}

func (s *UpdateImageRequest) SetProductId(v string) *UpdateImageRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr(v string) *UpdateImageRequest {
	s.StrAttr = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr2(v string) *UpdateImageRequest {
	s.StrAttr2 = &v
	return s
}

type UpdateImageResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) SetCode(v int32) *UpdateImageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageResponseBody) SetMessage(v string) *UpdateImageResponseBody {
	s.Message = &v
	return s
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imagesearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddImageWithOptions(request *AddImageRequest, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.CustomContent)) {
		body["CustomContent"] = request.CustomContent
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr)) {
		body["IntAttr"] = request.IntAttr
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr2)) {
		body["IntAttr2"] = request.IntAttr2
	}

	if !tea.BoolValue(util.IsUnset(request.PicContent)) {
		body["PicContent"] = request.PicContent
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr)) {
		body["StrAttr"] = request.StrAttr
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr2)) {
		body["StrAttr2"] = request.StrAttr2
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImage"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
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

func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddImageAdvance(request *AddImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
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
		Product:  tea.String("ImageSearch"),
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

	addImageResp, _err := client.AddImageWithOptions(addImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addImageResp
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImage"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) DetailWithOptions(request *DetailRequest, runtime *util.RuntimeOptions) (_result *DetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Detail"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Detail(request *DetailRequest) (_result *DetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetailResponse{}
	_body, _err := client.DetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DumpMetaWithOptions(request *DumpMetaRequest, runtime *util.RuntimeOptions) (_result *DumpMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DumpMeta"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DumpMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DumpMeta(request *DumpMetaRequest) (_result *DumpMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DumpMetaResponse{}
	_body, _err := client.DumpMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DumpMetaListWithOptions(request *DumpMetaListRequest, runtime *util.RuntimeOptions) (_result *DumpMetaListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DumpMetaList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DumpMetaListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DumpMetaList(request *DumpMetaListRequest) (_result *DumpMetaListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DumpMetaListResponse{}
	_body, _err := client.DumpMetaListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IncreaseInstanceWithOptions(request *IncreaseInstanceRequest, runtime *util.RuntimeOptions) (_result *IncreaseInstanceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IncreaseInstance"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IncreaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IncreaseInstance(request *IncreaseInstanceRequest) (_result *IncreaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IncreaseInstanceResponse{}
	_body, _err := client.IncreaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IncreaseListWithOptions(request *IncreaseListRequest, runtime *util.RuntimeOptions) (_result *IncreaseListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IncreaseList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IncreaseListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IncreaseList(request *IncreaseListRequest) (_result *IncreaseListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IncreaseListResponse{}
	_body, _err := client.IncreaseListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageByNameWithOptions(request *SearchImageByNameRequest, runtime *util.RuntimeOptions) (_result *SearchImageByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		body["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchImageByName"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) SearchImageByName(request *SearchImageByNameRequest) (_result *SearchImageByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchImageByNameResponse{}
	_body, _err := client.SearchImageByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageByPicWithOptions(request *SearchImageByPicRequest, runtime *util.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		body["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PicContent)) {
		body["PicContent"] = request.PicContent
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchImageByPic"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) SearchImageByPic(request *SearchImageByPicRequest) (_result *SearchImageByPicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchImageByPicResponse{}
	_body, _err := client.SearchImageByPicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageByPicAdvance(request *SearchImageByPicAdvanceRequest, runtime *util.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
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
		Product:  tea.String("ImageSearch"),
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

	searchImageByPicResp, _err := client.SearchImageByPicWithOptions(searchImageByPicReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchImageByPicResp
	return _result, _err
}

func (client *Client) UpdateImageWithOptions(request *UpdateImageRequest, runtime *util.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomContent)) {
		body["CustomContent"] = request.CustomContent
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr)) {
		body["IntAttr"] = request.IntAttr
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr2)) {
		body["IntAttr2"] = request.IntAttr2
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr)) {
		body["StrAttr"] = request.StrAttr
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr2)) {
		body["StrAttr2"] = request.StrAttr2
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImage"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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
