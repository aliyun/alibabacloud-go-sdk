// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type DeleteImageRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ProductId    *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	PicName      *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
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

func (s *DeleteImageRequest) SetProductId(v string) *DeleteImageRequest {
	s.ProductId = &v
	return s
}

func (s *DeleteImageRequest) SetPicName(v string) *DeleteImageRequest {
	s.PicName = &v
	return s
}

type DeleteImageResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetMessage(v string) *DeleteImageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetCode(v int32) *DeleteImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageResponseBody) SetSuccess(v bool) *DeleteImageResponseBody {
	s.Success = &v
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

type AddImageRequest struct {
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	CategoryId    *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ProductId     *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	PicName       *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	PicContent    *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	Crop          *bool   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr       *int32  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	StrAttr       *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetInstanceName(v string) *AddImageRequest {
	s.InstanceName = &v
	return s
}

func (s *AddImageRequest) SetCategoryId(v int32) *AddImageRequest {
	s.CategoryId = &v
	return s
}

func (s *AddImageRequest) SetProductId(v string) *AddImageRequest {
	s.ProductId = &v
	return s
}

func (s *AddImageRequest) SetPicName(v string) *AddImageRequest {
	s.PicName = &v
	return s
}

func (s *AddImageRequest) SetPicContent(v string) *AddImageRequest {
	s.PicContent = &v
	return s
}

func (s *AddImageRequest) SetCrop(v bool) *AddImageRequest {
	s.Crop = &v
	return s
}

func (s *AddImageRequest) SetRegion(v string) *AddImageRequest {
	s.Region = &v
	return s
}

func (s *AddImageRequest) SetCustomContent(v string) *AddImageRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageRequest) SetIntAttr(v int32) *AddImageRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageRequest) SetStrAttr(v string) *AddImageRequest {
	s.StrAttr = &v
	return s
}

type AddImageAdvanceRequest struct {
	PicContentObject io.Reader `json:"PicContentObject,omitempty" xml:"PicContentObject,omitempty" require:"true"`
	InstanceName     *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	CategoryId       *int32    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ProductId        *string   `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	PicName          *string   `json:"PicName,omitempty" xml:"PicName,omitempty"`
	Crop             *bool     `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Region           *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	CustomContent    *string   `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr          *int32    `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	StrAttr          *string   `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s AddImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddImageAdvanceRequest) SetPicContentObject(v io.Reader) *AddImageAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *AddImageAdvanceRequest) SetInstanceName(v string) *AddImageAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *AddImageAdvanceRequest) SetCategoryId(v int32) *AddImageAdvanceRequest {
	s.CategoryId = &v
	return s
}

func (s *AddImageAdvanceRequest) SetProductId(v string) *AddImageAdvanceRequest {
	s.ProductId = &v
	return s
}

func (s *AddImageAdvanceRequest) SetPicName(v string) *AddImageAdvanceRequest {
	s.PicName = &v
	return s
}

func (s *AddImageAdvanceRequest) SetCrop(v bool) *AddImageAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *AddImageAdvanceRequest) SetRegion(v string) *AddImageAdvanceRequest {
	s.Region = &v
	return s
}

func (s *AddImageAdvanceRequest) SetCustomContent(v string) *AddImageAdvanceRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageAdvanceRequest) SetIntAttr(v int32) *AddImageAdvanceRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageAdvanceRequest) SetStrAttr(v string) *AddImageAdvanceRequest {
	s.StrAttr = &v
	return s
}

type AddImageResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	PicInfo   *AddImageResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetMessage(v string) *AddImageResponseBody {
	s.Message = &v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetCode(v int32) *AddImageResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageResponseBody) SetPicInfo(v *AddImageResponseBodyPicInfo) *AddImageResponseBody {
	s.PicInfo = v
	return s
}

func (s *AddImageResponseBody) SetSuccess(v bool) *AddImageResponseBody {
	s.Success = &v
	return s
}

type AddImageResponseBodyPicInfo struct {
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	CategoryId *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s AddImageResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *AddImageResponseBodyPicInfo) SetRegion(v string) *AddImageResponseBodyPicInfo {
	s.Region = &v
	return s
}

func (s *AddImageResponseBodyPicInfo) SetCategoryId(v int32) *AddImageResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

type AddImageResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type SearchImageByPicRequest struct {
	CategoryId   *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PicContent   *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	Crop         *bool   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Num          *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Start        *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
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

func (s *SearchImageByPicRequest) SetInstanceName(v string) *SearchImageByPicRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByPicRequest) SetPicContent(v string) *SearchImageByPicRequest {
	s.PicContent = &v
	return s
}

func (s *SearchImageByPicRequest) SetCrop(v bool) *SearchImageByPicRequest {
	s.Crop = &v
	return s
}

func (s *SearchImageByPicRequest) SetRegion(v string) *SearchImageByPicRequest {
	s.Region = &v
	return s
}

func (s *SearchImageByPicRequest) SetNum(v int32) *SearchImageByPicRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByPicRequest) SetStart(v int32) *SearchImageByPicRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByPicRequest) SetType(v string) *SearchImageByPicRequest {
	s.Type = &v
	return s
}

func (s *SearchImageByPicRequest) SetFilter(v string) *SearchImageByPicRequest {
	s.Filter = &v
	return s
}

type SearchImageByPicAdvanceRequest struct {
	PicContentObject io.Reader `json:"PicContentObject,omitempty" xml:"PicContentObject,omitempty" require:"true"`
	CategoryId       *int32    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	InstanceName     *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Crop             *bool     `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Region           *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	Num              *int32    `json:"Num,omitempty" xml:"Num,omitempty"`
	Start            *int32    `json:"Start,omitempty" xml:"Start,omitempty"`
	Type             *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Filter           *string   `json:"Filter,omitempty" xml:"Filter,omitempty"`
}

func (s SearchImageByPicAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByPicAdvanceRequest) SetPicContentObject(v io.Reader) *SearchImageByPicAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetCategoryId(v int32) *SearchImageByPicAdvanceRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetInstanceName(v string) *SearchImageByPicAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetCrop(v bool) *SearchImageByPicAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetRegion(v string) *SearchImageByPicAdvanceRequest {
	s.Region = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetNum(v int32) *SearchImageByPicAdvanceRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetStart(v int32) *SearchImageByPicAdvanceRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetType(v string) *SearchImageByPicAdvanceRequest {
	s.Type = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetFilter(v string) *SearchImageByPicAdvanceRequest {
	s.Filter = &v
	return s
}

type SearchImageByPicResponseBody struct {
	Msg       *string                                 `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Head      *SearchImageByPicResponseBodyHead       `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Auctions  []*SearchImageByPicResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	PicInfo   *SearchImageByPicResponseBodyPicInfo    `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByPicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBody) SetMsg(v string) *SearchImageByPicResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetHead(v *SearchImageByPicResponseBodyHead) *SearchImageByPicResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByPicResponseBody) SetRequestId(v string) *SearchImageByPicResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetAuctions(v []*SearchImageByPicResponseBodyAuctions) *SearchImageByPicResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByPicResponseBody) SetCode(v int32) *SearchImageByPicResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetPicInfo(v *SearchImageByPicResponseBodyPicInfo) *SearchImageByPicResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchImageByPicResponseBody) SetSuccess(v bool) *SearchImageByPicResponseBody {
	s.Success = &v
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

type SearchImageByPicResponseBodyAuctions struct {
	PicName        *string  `json:"PicName,omitempty" xml:"PicName,omitempty"`
	IntAttr        *int32   `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	CategoryId     *int32   `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ProductId      *string  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	StrAttr        *string  `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	SortExprValues *string  `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	CustomContent  *string  `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchImageByPicResponseBodyAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyAuctions) SetPicName(v string) *SearchImageByPicResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByPicResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetProductId(v string) *SearchImageByPicResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetSortExprValues(v string) *SearchImageByPicResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetCustomContent(v string) *SearchImageByPicResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetScore(v float32) *SearchImageByPicResponseBodyAuctions {
	s.Score = &v
	return s
}

type SearchImageByPicResponseBodyPicInfo struct {
	Region        *string                                             `json:"Region,omitempty" xml:"Region,omitempty"`
	CategoryId    *int32                                              `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	MultiRegion   []*SearchImageByPicResponseBodyPicInfoMultiRegion   `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	AllCategories []*SearchImageByPicResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
}

func (s SearchImageByPicResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfo) SetRegion(v string) *SearchImageByPicResponseBodyPicInfo {
	s.Region = &v
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

func (s *SearchImageByPicResponseBodyPicInfo) SetAllCategories(v []*SearchImageByPicResponseBodyPicInfoAllCategories) *SearchImageByPicResponseBodyPicInfo {
	s.AllCategories = v
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

type SearchImageByPicResponseBodyPicInfoAllCategories struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfoAllCategories) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByPicResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByPicResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

type SearchImageByPicResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchImageByPicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchImageByPicResponse) SetBody(v *SearchImageByPicResponseBody) *SearchImageByPicResponse {
	s.Body = v
	return s
}

type SearchImageByNameRequest struct {
	CategoryId   *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ProductId    *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	PicName      *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	Num          *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Start        *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
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

func (s *SearchImageByNameRequest) SetInstanceName(v string) *SearchImageByNameRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByNameRequest) SetProductId(v string) *SearchImageByNameRequest {
	s.ProductId = &v
	return s
}

func (s *SearchImageByNameRequest) SetPicName(v string) *SearchImageByNameRequest {
	s.PicName = &v
	return s
}

func (s *SearchImageByNameRequest) SetNum(v int32) *SearchImageByNameRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByNameRequest) SetStart(v int32) *SearchImageByNameRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByNameRequest) SetType(v string) *SearchImageByNameRequest {
	s.Type = &v
	return s
}

func (s *SearchImageByNameRequest) SetFilter(v string) *SearchImageByNameRequest {
	s.Filter = &v
	return s
}

type SearchImageByNameResponseBody struct {
	Msg       *string                                  `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Head      *SearchImageByNameResponseBodyHead       `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Auctions  []*SearchImageByNameResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	PicInfo   *SearchImageByNameResponseBodyPicInfo    `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBody) SetMsg(v string) *SearchImageByNameResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetHead(v *SearchImageByNameResponseBodyHead) *SearchImageByNameResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByNameResponseBody) SetRequestId(v string) *SearchImageByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetAuctions(v []*SearchImageByNameResponseBodyAuctions) *SearchImageByNameResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByNameResponseBody) SetCode(v int32) *SearchImageByNameResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetPicInfo(v *SearchImageByNameResponseBodyPicInfo) *SearchImageByNameResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchImageByNameResponseBody) SetSuccess(v bool) *SearchImageByNameResponseBody {
	s.Success = &v
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

type SearchImageByNameResponseBodyAuctions struct {
	PicName        *string  `json:"PicName,omitempty" xml:"PicName,omitempty"`
	IntAttr        *int32   `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	CategoryId     *int32   `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ProductId      *string  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	StrAttr        *string  `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	SortExprValues *string  `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	CustomContent  *string  `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchImageByNameResponseBodyAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyAuctions) SetPicName(v string) *SearchImageByNameResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByNameResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetProductId(v string) *SearchImageByNameResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetSortExprValues(v string) *SearchImageByNameResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetCustomContent(v string) *SearchImageByNameResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetScore(v float32) *SearchImageByNameResponseBodyAuctions {
	s.Score = &v
	return s
}

type SearchImageByNameResponseBodyPicInfo struct {
	Region        *string                                              `json:"Region,omitempty" xml:"Region,omitempty"`
	CategoryId    *int32                                               `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	MultiRegion   []*SearchImageByNameResponseBodyPicInfoMultiRegion   `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	AllCategories []*SearchImageByNameResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
}

func (s SearchImageByNameResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfo) SetRegion(v string) *SearchImageByNameResponseBodyPicInfo {
	s.Region = &v
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

func (s *SearchImageByNameResponseBodyPicInfo) SetAllCategories(v []*SearchImageByNameResponseBodyPicInfoAllCategories) *SearchImageByNameResponseBodyPicInfo {
	s.AllCategories = v
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

type SearchImageByNameResponseBodyPicInfoAllCategories struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfoAllCategories) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByNameResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByNameResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

type SearchImageByNameResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchImageByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchImageByNameResponse) SetBody(v *SearchImageByNameResponseBody) *SearchImageByNameResponse {
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

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImage"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) AddImageWithOptions(request *AddImageRequest, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddImage"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.PicContentObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	addImageReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	addImageResp, _err := client.AddImageWithOptions(addImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addImageResp
	return _result, _err
}

func (client *Client) SearchImageByPicWithOptions(request *SearchImageByPicRequest, runtime *util.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchImageByPicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchImageByPic"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.PicContentObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	searchImageByPicReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	searchImageByPicResp, _err := client.SearchImageByPicWithOptions(searchImageByPicReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchImageByPicResp
	return _result, _err
}

func (client *Client) SearchImageByNameWithOptions(request *SearchImageByNameRequest, runtime *util.RuntimeOptions) (_result *SearchImageByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchImageByNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchImageByName"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
