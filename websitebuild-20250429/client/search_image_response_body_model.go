// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SearchImageResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SearchImageResponseBody
	GetErrorMsg() *string
	SetImageResponse(v *SearchImageResponseBodyImageResponse) *SearchImageResponseBody
	GetImageResponse() *SearchImageResponseBodyImageResponse
	SetRequestId(v string) *SearchImageResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SearchImageResponseBody
	GetSuccess() *string
}

type SearchImageResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg      *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	ImageResponse *SearchImageResponseBodyImageResponse `json:"ImageResponse,omitempty" xml:"ImageResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchImageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchImageResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SearchImageResponseBody) GetImageResponse() *SearchImageResponseBodyImageResponse {
	return s.ImageResponse
}

func (s *SearchImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchImageResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchImageResponseBody) SetErrorCode(v string) *SearchImageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchImageResponseBody) SetErrorMsg(v string) *SearchImageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SearchImageResponseBody) SetImageResponse(v *SearchImageResponseBodyImageResponse) *SearchImageResponseBody {
	s.ImageResponse = v
	return s
}

func (s *SearchImageResponseBody) SetRequestId(v string) *SearchImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageResponseBody) SetSuccess(v string) *SearchImageResponseBody {
	s.Success = &v
	return s
}

func (s *SearchImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchImageResponseBodyImageResponse struct {
	ImageList []*SearchImageResponseBodyImageResponseImageList `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2c6b65b6f9d625d4e2514a628bde8eb2e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s SearchImageResponseBodyImageResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchImageResponseBodyImageResponse) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBodyImageResponse) GetImageList() []*SearchImageResponseBodyImageResponseImageList {
	return s.ImageList
}

func (s *SearchImageResponseBodyImageResponse) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchImageResponseBodyImageResponse) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchImageResponseBodyImageResponse) SetImageList(v []*SearchImageResponseBodyImageResponseImageList) *SearchImageResponseBodyImageResponse {
	s.ImageList = v
	return s
}

func (s *SearchImageResponseBodyImageResponse) SetMaxResults(v int32) *SearchImageResponseBodyImageResponse {
	s.MaxResults = &v
	return s
}

func (s *SearchImageResponseBodyImageResponse) SetNextToken(v string) *SearchImageResponseBodyImageResponse {
	s.NextToken = &v
	return s
}

func (s *SearchImageResponseBodyImageResponse) Validate() error {
	return dara.Validate(s)
}

type SearchImageResponseBodyImageResponseImageList struct {
	DescriptiveTones *string `json:"DescriptiveTones,omitempty" xml:"DescriptiveTones,omitempty"`
	// example:
	//
	// 1000
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// normal
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	ImageRatio    *string `json:"ImageRatio,omitempty" xml:"ImageRatio,omitempty"`
	// example:
	//
	// 70687446-821c-4ebd-9be6-b57dc0a3500a
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// oss key
	//
	// example:
	//
	// 1ad5/c728/cb40/2410/e01d/b24c/9acd/7d51/1ad5c728cb402410e01db24c9acd7d51
	OssKey              *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	QuantitativePalette *string `json:"QuantitativePalette,omitempty" xml:"QuantitativePalette,omitempty"`
	TagsFromImage       *string `json:"TagsFromImage,omitempty" xml:"TagsFromImage,omitempty"`
	// example:
	//
	// https://other-general-huadong1.oss-cn-hangzhou.aliyuncs.com/uploadWidget/RollTicket_01.jpeg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 154
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SearchImageResponseBodyImageResponseImageList) String() string {
	return dara.Prettify(s)
}

func (s SearchImageResponseBodyImageResponseImageList) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBodyImageResponseImageList) GetDescriptiveTones() *string {
	return s.DescriptiveTones
}

func (s *SearchImageResponseBodyImageResponseImageList) GetHeight() *int32 {
	return s.Height
}

func (s *SearchImageResponseBodyImageResponseImageList) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *SearchImageResponseBodyImageResponseImageList) GetImageRatio() *string {
	return s.ImageRatio
}

func (s *SearchImageResponseBodyImageResponseImageList) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *SearchImageResponseBodyImageResponseImageList) GetOssKey() *string {
	return s.OssKey
}

func (s *SearchImageResponseBodyImageResponseImageList) GetQuantitativePalette() *string {
	return s.QuantitativePalette
}

func (s *SearchImageResponseBodyImageResponseImageList) GetTagsFromImage() *string {
	return s.TagsFromImage
}

func (s *SearchImageResponseBodyImageResponseImageList) GetUrl() *string {
	return s.Url
}

func (s *SearchImageResponseBodyImageResponseImageList) GetWidth() *int32 {
	return s.Width
}

func (s *SearchImageResponseBodyImageResponseImageList) SetDescriptiveTones(v string) *SearchImageResponseBodyImageResponseImageList {
	s.DescriptiveTones = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetHeight(v int32) *SearchImageResponseBodyImageResponseImageList {
	s.Height = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetImageCategory(v string) *SearchImageResponseBodyImageResponseImageList {
	s.ImageCategory = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetImageRatio(v string) *SearchImageResponseBodyImageResponseImageList {
	s.ImageRatio = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetImageUuid(v string) *SearchImageResponseBodyImageResponseImageList {
	s.ImageUuid = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetOssKey(v string) *SearchImageResponseBodyImageResponseImageList {
	s.OssKey = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetQuantitativePalette(v string) *SearchImageResponseBodyImageResponseImageList {
	s.QuantitativePalette = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetTagsFromImage(v string) *SearchImageResponseBodyImageResponseImageList {
	s.TagsFromImage = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetUrl(v string) *SearchImageResponseBodyImageResponseImageList {
	s.Url = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) SetWidth(v int32) *SearchImageResponseBodyImageResponseImageList {
	s.Width = &v
	return s
}

func (s *SearchImageResponseBodyImageResponseImageList) Validate() error {
	return dara.Validate(s)
}
