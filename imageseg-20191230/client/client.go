// This file is auto-generated, don't edit it. Thanks.
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

type ChangeSkyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ChangeSky/ChangeSky2.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ChangeSky/ChangeSky6.jpg
	ReplaceImageURL *string `json:"ReplaceImageURL,omitempty" xml:"ReplaceImageURL,omitempty"`
}

func (s ChangeSkyRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyRequest) GoString() string {
	return s.String()
}

func (s *ChangeSkyRequest) SetImageURL(v string) *ChangeSkyRequest {
	s.ImageURL = &v
	return s
}

func (s *ChangeSkyRequest) SetReplaceImageURL(v string) *ChangeSkyRequest {
	s.ReplaceImageURL = &v
	return s
}

type ChangeSkyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ChangeSky/ChangeSky2.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ChangeSky/ChangeSky6.jpg
	ReplaceImageURLObject io.Reader `json:"ReplaceImageURL,omitempty" xml:"ReplaceImageURL,omitempty"`
}

func (s ChangeSkyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeSkyAdvanceRequest) SetImageURLObject(v io.Reader) *ChangeSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ChangeSkyAdvanceRequest) SetReplaceImageURLObject(v io.Reader) *ChangeSkyAdvanceRequest {
	s.ReplaceImageURLObject = v
	return s
}

type ChangeSkyResponseBody struct {
	Data *ChangeSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// F9D60817-EC5A-4BAC-9092-4AD42220CFA8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeSkyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeSkyResponseBody) SetData(v *ChangeSkyResponseBodyData) *ChangeSkyResponseBody {
	s.Data = v
	return s
}

func (s *ChangeSkyResponseBody) SetRequestId(v string) *ChangeSkyResponseBody {
	s.RequestId = &v
	return s
}

type ChangeSkyResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_skySegmentator/2020-7-24/invi_skySegmentator_015955791588111000000_5pn2QM.jpg?Expires=1595580958&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Sq4po8h3WAj%2BBFrCgTP3ghlXn4****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ChangeSkyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeSkyResponseBodyData) SetImageURL(v string) *ChangeSkyResponseBodyData {
	s.ImageURL = &v
	return s
}

type ChangeSkyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeSkyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeSkyResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyResponse) GoString() string {
	return s.String()
}

func (s *ChangeSkyResponse) SetHeaders(v map[string]*string) *ChangeSkyResponse {
	s.Headers = v
	return s
}

func (s *ChangeSkyResponse) SetStatusCode(v int32) *ChangeSkyResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeSkyResponse) SetBody(v *ChangeSkyResponseBody) *ChangeSkyResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// E75FE679-0303-4DD1-8252-1143B4FA8A27
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 43A0AEB6-45F4-4138-8E89-E1A5D63200E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// paramsIllegal
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 49E2CC28-ED1D-4CC5-854D-7D0AE2B20976
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {"ImageUrl":"http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_/2020-4-2/invi__015858226731531000018_UE7B9p.png?Expires=1585824473&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSR****&Signature=etyeYQQ%2BWAyQTqQKd8Xq0GiOW****"}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// PROCESS_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type ParseFaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ParseFace/ParseFace1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ParseFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceRequest) GoString() string {
	return s.String()
}

func (s *ParseFaceRequest) SetImageURL(v string) *ParseFaceRequest {
	s.ImageURL = &v
	return s
}

type ParseFaceAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ParseFace/ParseFace1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ParseFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ParseFaceAdvanceRequest) SetImageURLObject(v io.Reader) *ParseFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type ParseFaceResponseBody struct {
	Data *ParseFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6C24839-91A7-41DA-B31F-98F08EF80CC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ParseFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceResponseBody) GoString() string {
	return s.String()
}

func (s *ParseFaceResponseBody) SetData(v *ParseFaceResponseBodyData) *ParseFaceResponseBody {
	s.Data = v
	return s
}

func (s *ParseFaceResponseBody) SetRequestId(v string) *ParseFaceResponseBody {
	s.RequestId = &v
	return s
}

type ParseFaceResponseBodyData struct {
	Elements []*ParseFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ParseFace/ParseFace1.jpg
	OriginImageURL *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
}

func (s ParseFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ParseFaceResponseBodyData) SetElements(v []*ParseFaceResponseBodyDataElements) *ParseFaceResponseBodyData {
	s.Elements = v
	return s
}

func (s *ParseFaceResponseBodyData) SetOriginImageURL(v string) *ParseFaceResponseBodyData {
	s.OriginImageURL = &v
	return s
}

type ParseFaceResponseBodyDataElements struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/fivesensesegmenter/prod/560FA2E7-FDC6-59A5-ABDD-D62A05146734/skin/_18dd_20211231-040658.png?Expires=1640925418&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=2g0M88wZl%2Bn4t4gzQX%2BTIskpWB****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// skin
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ParseFaceResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *ParseFaceResponseBodyDataElements) SetImageURL(v string) *ParseFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *ParseFaceResponseBodyDataElements) SetName(v string) *ParseFaceResponseBodyDataElements {
	s.Name = &v
	return s
}

type ParseFaceResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ParseFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ParseFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceResponse) GoString() string {
	return s.String()
}

func (s *ParseFaceResponse) SetHeaders(v map[string]*string) *ParseFaceResponse {
	s.Headers = v
	return s
}

func (s *ParseFaceResponse) SetStatusCode(v int32) *ParseFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ParseFaceResponse) SetBody(v *ParseFaceResponseBody) *ParseFaceResponse {
	s.Body = v
	return s
}

type RefineMaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/RefineMask/RefineMask1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/RefineMask/RefineMask6.jpg
	MaskImageURL *string `json:"MaskImageURL,omitempty" xml:"MaskImageURL,omitempty"`
}

func (s RefineMaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskRequest) GoString() string {
	return s.String()
}

func (s *RefineMaskRequest) SetImageURL(v string) *RefineMaskRequest {
	s.ImageURL = &v
	return s
}

func (s *RefineMaskRequest) SetMaskImageURL(v string) *RefineMaskRequest {
	s.MaskImageURL = &v
	return s
}

type RefineMaskAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/RefineMask/RefineMask1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/RefineMask/RefineMask6.jpg
	MaskImageURLObject io.Reader `json:"MaskImageURL,omitempty" xml:"MaskImageURL,omitempty"`
}

func (s RefineMaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RefineMaskAdvanceRequest) SetImageURLObject(v io.Reader) *RefineMaskAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RefineMaskAdvanceRequest) SetMaskImageURLObject(v io.Reader) *RefineMaskAdvanceRequest {
	s.MaskImageURLObject = v
	return s
}

type RefineMaskResponseBody struct {
	Data *RefineMaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E948F80B-86D9-54E0-9FF9-ACF3B1DA83C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefineMaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskResponseBody) GoString() string {
	return s.String()
}

func (s *RefineMaskResponseBody) SetData(v *RefineMaskResponseBodyData) *RefineMaskResponseBody {
	s.Data = v
	return s
}

func (s *RefineMaskResponseBody) SetRequestId(v string) *RefineMaskResponseBody {
	s.RequestId = &v
	return s
}

type RefineMaskResponseBodyData struct {
	Elements []*RefineMaskResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RefineMaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefineMaskResponseBodyData) SetElements(v []*RefineMaskResponseBodyDataElements) *RefineMaskResponseBodyData {
	s.Elements = v
	return s
}

type RefineMaskResponseBodyDataElements struct {
	// example:
	//
	// http://algo-app-taobao-mm-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/pixelai-portrait-beauty%2F2020_03_04%2F61f544a1a5004c88a2bf29452db494e9.jpeg?OSSAccessKeyId=LTAI4Fmdm1gQonFLrghJ****&Expires=1583406122&Signature=Heet1ivG0xFP3YlO6usvd0pmrH****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RefineMaskResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RefineMaskResponseBodyDataElements) SetImageURL(v string) *RefineMaskResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

type RefineMaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefineMaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefineMaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskResponse) GoString() string {
	return s.String()
}

func (s *RefineMaskResponse) SetHeaders(v map[string]*string) *RefineMaskResponse {
	s.Headers = v
	return s
}

func (s *RefineMaskResponse) SetStatusCode(v int32) *RefineMaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RefineMaskResponse) SetBody(v *RefineMaskResponseBody) *RefineMaskResponse {
	s.Body = v
	return s
}

type SegmentBodyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentBody/SegmentBody1.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentBodyRequest) SetImageURL(v string) *SegmentBodyRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentBodyRequest) SetReturnForm(v string) *SegmentBodyRequest {
	s.ReturnForm = &v
	return s
}

type SegmentBodyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentBody/SegmentBody1.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentBodyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentBodyAdvanceRequest) SetReturnForm(v string) *SegmentBodyAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentBodyResponseBody struct {
	Data *SegmentBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 30EDCEEA-2806-44C6-AF0B-0988849106FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponseBody) SetData(v *SegmentBodyResponseBodyData) *SegmentBodyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentBodyResponseBody) SetRequestId(v string) *SegmentBodyResponseBody {
	s.RequestId = &v
	return s
}

type SegmentBodyResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_humansegmenter/2021-3-31/invi_humansegmenter_016171823500001081370_Ej0WwO.jpg?Expires=1617184150&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=ZwaWXpAOMzHar%2B1wVO7zeSD83r****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponseBodyData) SetImageURL(v string) *SegmentBodyResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentBodyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentBodyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponse) SetHeaders(v map[string]*string) *SegmentBodyResponse {
	s.Headers = v
	return s
}

func (s *SegmentBodyResponse) SetStatusCode(v int32) *SegmentBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentBodyResponse) SetBody(v *SegmentBodyResponseBody) *SegmentBodyResponse {
	s.Body = v
	return s
}

type SegmentClothRequest struct {
	ClothClass []*string `json:"ClothClass,omitempty" xml:"ClothClass,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCloth/SegmentCloth1.jpg
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OutMode    *int64  `json:"OutMode,omitempty" xml:"OutMode,omitempty"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentClothRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothRequest) GoString() string {
	return s.String()
}

func (s *SegmentClothRequest) SetClothClass(v []*string) *SegmentClothRequest {
	s.ClothClass = v
	return s
}

func (s *SegmentClothRequest) SetImageURL(v string) *SegmentClothRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentClothRequest) SetOutMode(v int64) *SegmentClothRequest {
	s.OutMode = &v
	return s
}

func (s *SegmentClothRequest) SetReturnForm(v string) *SegmentClothRequest {
	s.ReturnForm = &v
	return s
}

type SegmentClothAdvanceRequest struct {
	ClothClass []*string `json:"ClothClass,omitempty" xml:"ClothClass,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCloth/SegmentCloth1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OutMode        *int64    `json:"OutMode,omitempty" xml:"OutMode,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentClothAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentClothAdvanceRequest) SetClothClass(v []*string) *SegmentClothAdvanceRequest {
	s.ClothClass = v
	return s
}

func (s *SegmentClothAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentClothAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentClothAdvanceRequest) SetOutMode(v int64) *SegmentClothAdvanceRequest {
	s.OutMode = &v
	return s
}

func (s *SegmentClothAdvanceRequest) SetReturnForm(v string) *SegmentClothAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentClothResponseBody struct {
	Data *SegmentClothResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// BCE049A3-FE69-41CF-A870-5970156388A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentClothResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseBody) SetData(v *SegmentClothResponseBodyData) *SegmentClothResponseBody {
	s.Data = v
	return s
}

func (s *SegmentClothResponseBody) SetRequestId(v string) *SegmentClothResponseBody {
	s.RequestId = &v
	return s
}

type SegmentClothResponseBodyData struct {
	Elements []*SegmentClothResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentClothResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseBodyData) SetElements(v []*SegmentClothResponseBodyDataElements) *SegmentClothResponseBodyData {
	s.Elements = v
	return s
}

type SegmentClothResponseBodyDataElements struct {
	ClassUrl map[string]*string `json:"ClassUrl,omitempty" xml:"ClassUrl,omitempty"`
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/clothingsegmentation-2020-06-17-16-54-40-688c84cbbd-hnqfq/2020-6-18/invi__015924459307821000041_IIVHoM.png?Expires=1592447730&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Hy8pn3IQj8nuKN0LEaC57cee9L****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentClothResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseBodyDataElements) SetClassUrl(v map[string]*string) *SegmentClothResponseBodyDataElements {
	s.ClassUrl = v
	return s
}

func (s *SegmentClothResponseBodyDataElements) SetImageURL(v string) *SegmentClothResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

type SegmentClothResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentClothResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentClothResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothResponse) GoString() string {
	return s.String()
}

func (s *SegmentClothResponse) SetHeaders(v map[string]*string) *SegmentClothResponse {
	s.Headers = v
	return s
}

func (s *SegmentClothResponse) SetStatusCode(v int32) *SegmentClothResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentClothResponse) SetBody(v *SegmentClothResponseBody) *SegmentClothResponse {
	s.Body = v
	return s
}

type SegmentCommodityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCommodity/SegmentCommodity1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommodityRequest) SetImageURL(v string) *SegmentCommodityRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentCommodityRequest) SetReturnForm(v string) *SegmentCommodityRequest {
	s.ReturnForm = &v
	return s
}

type SegmentCommodityAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCommodity/SegmentCommodity1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommodityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommodityAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentCommodityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentCommodityAdvanceRequest) SetReturnForm(v string) *SegmentCommodityAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentCommodityResponseBody struct {
	Data *SegmentCommodityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6C24839-91A7-41DA-B31F-98F08EF80CC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponseBody) SetData(v *SegmentCommodityResponseBodyData) *SegmentCommodityResponseBody {
	s.Data = v
	return s
}

func (s *SegmentCommodityResponseBody) SetRequestId(v string) *SegmentCommodityResponseBody {
	s.RequestId = &v
	return s
}

type SegmentCommodityResponseBodyData struct {
	// example:
	//
	// http://luban-vgd-invi.oss-cn-hangzhou.aliyuncs.com/upload/result_segmenter/2019-12-20/invi_segmenter_015768355410261076021_Z3t0fc.png?Expires=1577094741&OSSAccessKeyId=LTAI4Fc5SVvzUQ19K1Cz****&Signature=pkaKK3VlfsTR2r%2BYycJzTVEEos****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentCommodityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponseBodyData) SetImageURL(v string) *SegmentCommodityResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentCommodityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityResponse) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponse) SetHeaders(v map[string]*string) *SegmentCommodityResponse {
	s.Headers = v
	return s
}

func (s *SegmentCommodityResponse) SetStatusCode(v int32) *SegmentCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentCommodityResponse) SetBody(v *SegmentCommodityResponseBody) *SegmentCommodityResponse {
	s.Body = v
	return s
}

type SegmentCommonImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCommonImage/SegmentCommonImage1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommonImageRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageRequest) SetImageURL(v string) *SegmentCommonImageRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentCommonImageRequest) SetReturnForm(v string) *SegmentCommonImageRequest {
	s.ReturnForm = &v
	return s
}

type SegmentCommonImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCommonImage/SegmentCommonImage1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommonImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentCommonImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentCommonImageAdvanceRequest) SetReturnForm(v string) *SegmentCommonImageAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentCommonImageResponseBody struct {
	Data *SegmentCommonImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1B8BEF02-0672-44CA-A974-4D6FAC392497
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentCommonImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageResponseBody) SetData(v *SegmentCommonImageResponseBodyData) *SegmentCommonImageResponseBody {
	s.Data = v
	return s
}

func (s *SegmentCommonImageResponseBody) SetRequestId(v string) *SegmentCommonImageResponseBody {
	s.RequestId = &v
	return s
}

type SegmentCommonImageResponseBodyData struct {
	// example:
	//
	// http://luban-vgd-invi.oss-cn-hangzhou.aliyuncs.com/upload/result_segmenter/2019-12-20/invi_segmenter_015768355410261076021_Z3t0fc.png?Expires=1577094741&OSSAccessKeyId=LTAI4Fc5SVvzUQ19K1Cz****&Signature=pkaKK3VlfsTR2r%2BYycJzTVEEos****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentCommonImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageResponseBodyData) SetImageURL(v string) *SegmentCommonImageResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentCommonImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentCommonImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentCommonImageResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageResponse) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageResponse) SetHeaders(v map[string]*string) *SegmentCommonImageResponse {
	s.Headers = v
	return s
}

func (s *SegmentCommonImageResponse) SetStatusCode(v int32) *SegmentCommonImageResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentCommonImageResponse) SetBody(v *SegmentCommonImageResponseBody) *SegmentCommonImageResponse {
	s.Body = v
	return s
}

type SegmentFoodRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentFood/SegmentFood5.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentFoodRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodRequest) GoString() string {
	return s.String()
}

func (s *SegmentFoodRequest) SetImageURL(v string) *SegmentFoodRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentFoodRequest) SetReturnForm(v string) *SegmentFoodRequest {
	s.ReturnForm = &v
	return s
}

type SegmentFoodAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentFood/SegmentFood5.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentFoodAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFoodAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentFoodAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentFoodAdvanceRequest) SetReturnForm(v string) *SegmentFoodAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentFoodResponseBody struct {
	Data *SegmentFoodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 38265D08-AD0F-4752-8E96-D1D9FB96C3D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentFoodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentFoodResponseBody) SetData(v *SegmentFoodResponseBodyData) *SegmentFoodResponseBody {
	s.Data = v
	return s
}

func (s *SegmentFoodResponseBody) SetRequestId(v string) *SegmentFoodResponseBody {
	s.RequestId = &v
	return s
}

type SegmentFoodResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/foodsegmenter-2020-06-17-15-24-00-8658fc85b8-8ds8k/2020-6-18/invi__015924442076191000002_WqJ99N.png?Expires=1592446007&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=5ITSd6ndSuP7pUfoDFpgLPUOGg****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentFoodResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentFoodResponseBodyData) SetImageURL(v string) *SegmentFoodResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentFoodResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentFoodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentFoodResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodResponse) GoString() string {
	return s.String()
}

func (s *SegmentFoodResponse) SetHeaders(v map[string]*string) *SegmentFoodResponse {
	s.Headers = v
	return s
}

func (s *SegmentFoodResponse) SetStatusCode(v int32) *SegmentFoodResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentFoodResponse) SetBody(v *SegmentFoodResponseBody) *SegmentFoodResponse {
	s.Body = v
	return s
}

type SegmentHDBodyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDBody/SegmentHDBody1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyRequest) SetImageURL(v string) *SegmentHDBodyRequest {
	s.ImageURL = &v
	return s
}

type SegmentHDBodyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDBody/SegmentHDBody1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHDBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentHDBodyResponseBody struct {
	Data *SegmentHDBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A8D3F5C3-E414-4981-8D84-E2CADF0B7CBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHDBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyResponseBody) SetData(v *SegmentHDBodyResponseBodyData) *SegmentHDBodyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHDBodyResponseBody) SetRequestId(v string) *SegmentHDBodyResponseBody {
	s.RequestId = &v
	return s
}

type SegmentHDBodyResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/segmenthdbody-2020-05-18-16-27-45-675d9884d7-kd9dz/2020-5-18/invi_humansegmenter_015897914589851000001_wQbLq9.png?Expires=1589793259&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Lx6xSS0t7lqEvy5Qd1keccIAjL****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyResponseBodyData) SetImageURL(v string) *SegmentHDBodyResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentHDBodyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHDBodyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHDBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyResponse) SetHeaders(v map[string]*string) *SegmentHDBodyResponse {
	s.Headers = v
	return s
}

func (s *SegmentHDBodyResponse) SetStatusCode(v int32) *SegmentHDBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHDBodyResponse) SetBody(v *SegmentHDBodyResponseBody) *SegmentHDBodyResponse {
	s.Body = v
	return s
}

type SegmentHDCommonImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDCommonImage/SegmentHDCommonImage1.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s SegmentHDCommonImageRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageRequest) SetImageUrl(v string) *SegmentHDCommonImageRequest {
	s.ImageUrl = &v
	return s
}

type SegmentHDCommonImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDCommonImage/SegmentHDCommonImage1.jpg
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s SegmentHDCommonImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageAdvanceRequest) SetImageUrlObject(v io.Reader) *SegmentHDCommonImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

type SegmentHDCommonImageResponseBody struct {
	Data    *SegmentHDCommonImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EC994171-7964-44B3-85AF-A715CB56747D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHDCommonImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageResponseBody) SetData(v *SegmentHDCommonImageResponseBodyData) *SegmentHDCommonImageResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHDCommonImageResponseBody) SetMessage(v string) *SegmentHDCommonImageResponseBody {
	s.Message = &v
	return s
}

func (s *SegmentHDCommonImageResponseBody) SetRequestId(v string) *SegmentHDCommonImageResponseBody {
	s.RequestId = &v
	return s
}

type SegmentHDCommonImageResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_commoditysegmenter/2020-10-27/invi_commoditysegmenter_016037842193171000000_5pn2QM.png?Expires=1603786019&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=HwUztguGBYXmXGEmuTh%2FL3ztoh****
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s SegmentHDCommonImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageResponseBodyData) SetImageUrl(v string) *SegmentHDCommonImageResponseBodyData {
	s.ImageUrl = &v
	return s
}

type SegmentHDCommonImageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHDCommonImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHDCommonImageResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageResponse) SetHeaders(v map[string]*string) *SegmentHDCommonImageResponse {
	s.Headers = v
	return s
}

func (s *SegmentHDCommonImageResponse) SetStatusCode(v int32) *SegmentHDCommonImageResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHDCommonImageResponse) SetBody(v *SegmentHDCommonImageResponseBody) *SegmentHDCommonImageResponse {
	s.Body = v
	return s
}

type SegmentHDSkyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDSky/SegmentHDSky4.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDSkyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyRequest) SetImageURL(v string) *SegmentHDSkyRequest {
	s.ImageURL = &v
	return s
}

type SegmentHDSkyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDSky/SegmentHDSky4.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDSkyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHDSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentHDSkyResponseBody struct {
	Data *SegmentHDSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1173F38F-B4F4-4A07-AB2E-D490C01347E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHDSkyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyResponseBody) SetData(v *SegmentHDSkyResponseBodyData) *SegmentHDSkyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHDSkyResponseBody) SetRequestId(v string) *SegmentHDSkyResponseBody {
	s.RequestId = &v
	return s
}

type SegmentHDSkyResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/sky-segmentation-hd/res/1173F38F-B4F4-4A07-AB2E-D490C01347E5_0d56_20201027-061858.jpg?Expires=1603781339&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=2F8%2Bj%2FWruWOMqDezwpnJOkcNJD****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDSkyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyResponseBodyData) SetImageURL(v string) *SegmentHDSkyResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentHDSkyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHDSkyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHDSkyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyResponse) SetHeaders(v map[string]*string) *SegmentHDSkyResponse {
	s.Headers = v
	return s
}

func (s *SegmentHDSkyResponse) SetStatusCode(v int32) *SegmentHDSkyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHDSkyResponse) SetBody(v *SegmentHDSkyResponseBody) *SegmentHDSkyResponse {
	s.Body = v
	return s
}

type SegmentHairRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHair/SegmentHair1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHairRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairRequest) GoString() string {
	return s.String()
}

func (s *SegmentHairRequest) SetImageURL(v string) *SegmentHairRequest {
	s.ImageURL = &v
	return s
}

type SegmentHairAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHair/SegmentHair1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHairAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHairAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHairAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentHairResponseBody struct {
	Data *SegmentHairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6C24839-91A7-41DA-B31F-98F08EF80CC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHairResponseBody) SetData(v *SegmentHairResponseBodyData) *SegmentHairResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHairResponseBody) SetRequestId(v string) *SegmentHairResponseBody {
	s.RequestId = &v
	return s
}

type SegmentHairResponseBodyData struct {
	Elements []*SegmentHairResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentHairResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHairResponseBodyData) SetElements(v []*SegmentHairResponseBodyDataElements) *SegmentHairResponseBodyData {
	s.Elements = v
	return s
}

type SegmentHairResponseBodyDataElements struct {
	// example:
	//
	// 180
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_HeadSegmenter/2021-12-31/invi_HeadSegmenter_016409228383064285967296_iPLUXA.png?Expires=1640924638&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=wpKOqSar1bYvGmlTMryfEH2Q9I****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 113
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 446
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 102
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SegmentHairResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentHairResponseBodyDataElements) SetHeight(v int32) *SegmentHairResponseBodyDataElements {
	s.Height = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) SetImageURL(v string) *SegmentHairResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) SetWidth(v int32) *SegmentHairResponseBodyDataElements {
	s.Width = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) SetX(v int32) *SegmentHairResponseBodyDataElements {
	s.X = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) SetY(v int32) *SegmentHairResponseBodyDataElements {
	s.Y = &v
	return s
}

type SegmentHairResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHairResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairResponse) GoString() string {
	return s.String()
}

func (s *SegmentHairResponse) SetHeaders(v map[string]*string) *SegmentHairResponse {
	s.Headers = v
	return s
}

func (s *SegmentHairResponse) SetStatusCode(v int32) *SegmentHairResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHairResponse) SetBody(v *SegmentHairResponseBody) *SegmentHairResponse {
	s.Body = v
	return s
}

type SegmentHeadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHead/SegmentHead1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentHeadRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadRequest) GoString() string {
	return s.String()
}

func (s *SegmentHeadRequest) SetImageURL(v string) *SegmentHeadRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentHeadRequest) SetReturnForm(v string) *SegmentHeadRequest {
	s.ReturnForm = &v
	return s
}

type SegmentHeadAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHead/SegmentHead1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentHeadAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHeadAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHeadAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentHeadAdvanceRequest) SetReturnForm(v string) *SegmentHeadAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentHeadResponseBody struct {
	Data *SegmentHeadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 89BBDA42-E1CA-426E-9A46-C703D8DBB1E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHeadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponseBody) SetData(v *SegmentHeadResponseBodyData) *SegmentHeadResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHeadResponseBody) SetRequestId(v string) *SegmentHeadResponseBody {
	s.RequestId = &v
	return s
}

type SegmentHeadResponseBodyData struct {
	Elements []*SegmentHeadResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentHeadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponseBodyData) SetElements(v []*SegmentHeadResponseBodyDataElements) *SegmentHeadResponseBodyData {
	s.Elements = v
	return s
}

type SegmentHeadResponseBodyDataElements struct {
	// example:
	//
	// 180
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_headsegmenter/2020-6-2/invi_headsegmenter_015910809094981099086_TAamhR.png?Expires=1591082709&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=xuUE%2FbYeB9LpR18VXawOVeutQU****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 116
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 445
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 102
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SegmentHeadResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponseBodyDataElements) SetHeight(v int32) *SegmentHeadResponseBodyDataElements {
	s.Height = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) SetImageURL(v string) *SegmentHeadResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) SetWidth(v int32) *SegmentHeadResponseBodyDataElements {
	s.Width = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) SetX(v int32) *SegmentHeadResponseBodyDataElements {
	s.X = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) SetY(v int32) *SegmentHeadResponseBodyDataElements {
	s.Y = &v
	return s
}

type SegmentHeadResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHeadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHeadResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadResponse) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponse) SetHeaders(v map[string]*string) *SegmentHeadResponse {
	s.Headers = v
	return s
}

func (s *SegmentHeadResponse) SetStatusCode(v int32) *SegmentHeadResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHeadResponse) SetBody(v *SegmentHeadResponseBody) *SegmentHeadResponse {
	s.Body = v
	return s
}

type SegmentSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentScene/SegmentScene1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneRequest) GoString() string {
	return s.String()
}

func (s *SegmentSceneRequest) SetImageURL(v string) *SegmentSceneRequest {
	s.ImageURL = &v
	return s
}

type SegmentSceneAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentScene/SegmentScene1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSceneAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentSceneAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentSceneAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentSceneResponseBody struct {
	Data *SegmentSceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 78EC13BB-74C5-4FBE-916E-C98BD721ED61
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentSceneResponseBody) SetData(v *SegmentSceneResponseBodyData) *SegmentSceneResponseBody {
	s.Data = v
	return s
}

func (s *SegmentSceneResponseBody) SetRequestId(v string) *SegmentSceneResponseBody {
	s.RequestId = &v
	return s
}

type SegmentSceneResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_skySegmentator/2020-7-24/invi_skySegmentator_015955807385661000002_WqJ99N.jpg?Expires=1595582538&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=sBP5bQCErDolM4QQG5%2B0HozNoO****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSceneResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentSceneResponseBodyData) SetImageURL(v string) *SegmentSceneResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneResponse) GoString() string {
	return s.String()
}

func (s *SegmentSceneResponse) SetHeaders(v map[string]*string) *SegmentSceneResponse {
	s.Headers = v
	return s
}

func (s *SegmentSceneResponse) SetStatusCode(v int32) *SegmentSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentSceneResponse) SetBody(v *SegmentSceneResponseBody) *SegmentSceneResponse {
	s.Body = v
	return s
}

type SegmentSkinRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentSkin/SegmentSkin2.jpg
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentSkinRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkinRequest) SetURL(v string) *SegmentSkinRequest {
	s.URL = &v
	return s
}

type SegmentSkinAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentSkin/SegmentSkin2.jpg
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentSkinAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkinAdvanceRequest) SetURLObject(v io.Reader) *SegmentSkinAdvanceRequest {
	s.URLObject = v
	return s
}

type SegmentSkinResponseBody struct {
	Data *SegmentSkinResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DA007354-6CF5-45BE-8333-E06318D848C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentSkinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentSkinResponseBody) SetData(v *SegmentSkinResponseBodyData) *SegmentSkinResponseBody {
	s.Data = v
	return s
}

func (s *SegmentSkinResponseBody) SetRequestId(v string) *SegmentSkinResponseBody {
	s.RequestId = &v
	return s
}

type SegmentSkinResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_skinsegmenter/2020-9-27/invi_skinsegmenter_016011971641871000001_wQbLq9.jpg?Expires=1601198964&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=xjKc%2BScprmB86cxtI%2B1T0R6TlE****
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentSkinResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentSkinResponseBodyData) SetURL(v string) *SegmentSkinResponseBodyData {
	s.URL = &v
	return s
}

type SegmentSkinResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentSkinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentSkinResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinResponse) GoString() string {
	return s.String()
}

func (s *SegmentSkinResponse) SetHeaders(v map[string]*string) *SegmentSkinResponse {
	s.Headers = v
	return s
}

func (s *SegmentSkinResponse) SetStatusCode(v int32) *SegmentSkinResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentSkinResponse) SetBody(v *SegmentSkinResponseBody) *SegmentSkinResponse {
	s.Body = v
	return s
}

type SegmentSkyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentSky/SegmentSky5.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSkyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkyRequest) SetImageURL(v string) *SegmentSkyRequest {
	s.ImageURL = &v
	return s
}

type SegmentSkyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentSky/SegmentSky5.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSkyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentSkyResponseBody struct {
	Data *SegmentSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 80E9D0A0-0330-4210-9986-CAC50C922FF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentSkyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentSkyResponseBody) SetData(v *SegmentSkyResponseBodyData) *SegmentSkyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentSkyResponseBody) SetRequestId(v string) *SegmentSkyResponseBody {
	s.RequestId = &v
	return s
}

type SegmentSkyResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/skysegmentation-2020-05-18-10-44-16-5bc8dc79f9-92b7z/2020-5-18/invi_skySegmentator_015897703560961000003_SqZLDv.png?Expires=1589772156&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=gXrzAUl%2BvIdYbQ9XKdho54MlkX****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSkyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentSkyResponseBodyData) SetImageURL(v string) *SegmentSkyResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentSkyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentSkyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentSkyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyResponse) GoString() string {
	return s.String()
}

func (s *SegmentSkyResponse) SetHeaders(v map[string]*string) *SegmentSkyResponse {
	s.Headers = v
	return s
}

func (s *SegmentSkyResponse) SetStatusCode(v int32) *SegmentSkyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentSkyResponse) SetBody(v *SegmentSkyResponseBody) *SegmentSkyResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imageseg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - ChangeSkyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeSkyResponse
func (client *Client) ChangeSkyWithOptions(request *ChangeSkyRequest, runtime *util.RuntimeOptions) (_result *ChangeSkyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ReplaceImageURL)) {
		query["ReplaceImageURL"] = request.ReplaceImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeSky"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeSkyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeSkyRequest
//
// @return ChangeSkyResponse
func (client *Client) ChangeSky(request *ChangeSkyRequest) (_result *ChangeSkyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeSkyResponse{}
	_body, _err := client.ChangeSkyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeSkyAdvance(request *ChangeSkyAdvanceRequest, runtime *util.RuntimeOptions) (_result *ChangeSkyResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	changeSkyReq := &ChangeSkyRequest{}
	openapiutil.Convert(request, changeSkyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		changeSkyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.ReplaceImageURLObject)) {
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
			Content:     request.ReplaceImageURLObject,
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
		changeSkyReq.ReplaceImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	changeSkyResp, _err := client.ChangeSkyWithOptions(changeSkyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = changeSkyResp
	return _result, _err
}

// @param request - GetAsyncJobResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncJobResultResponse
func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAsyncJobResultRequest
//
// @return GetAsyncJobResultResponse
func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ParseFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseFaceResponse
func (client *Client) ParseFaceWithOptions(request *ParseFaceRequest, runtime *util.RuntimeOptions) (_result *ParseFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ParseFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ParseFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ParseFaceRequest
//
// @return ParseFaceResponse
func (client *Client) ParseFace(request *ParseFaceRequest) (_result *ParseFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ParseFaceResponse{}
	_body, _err := client.ParseFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ParseFaceAdvance(request *ParseFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *ParseFaceResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	parseFaceReq := &ParseFaceRequest{}
	openapiutil.Convert(request, parseFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		parseFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	parseFaceResp, _err := client.ParseFaceWithOptions(parseFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = parseFaceResp
	return _result, _err
}

// @param request - RefineMaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefineMaskResponse
func (client *Client) RefineMaskWithOptions(request *RefineMaskRequest, runtime *util.RuntimeOptions) (_result *RefineMaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.MaskImageURL)) {
		body["MaskImageURL"] = request.MaskImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RefineMask"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefineMaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefineMaskRequest
//
// @return RefineMaskResponse
func (client *Client) RefineMask(request *RefineMaskRequest) (_result *RefineMaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefineMaskResponse{}
	_body, _err := client.RefineMaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefineMaskAdvance(request *RefineMaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *RefineMaskResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	refineMaskReq := &RefineMaskRequest{}
	openapiutil.Convert(request, refineMaskReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		refineMaskReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.MaskImageURLObject)) {
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
			Content:     request.MaskImageURLObject,
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
		refineMaskReq.MaskImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	refineMaskResp, _err := client.RefineMaskWithOptions(refineMaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = refineMaskResp
	return _result, _err
}

// @param request - SegmentBodyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentBodyResponse
func (client *Client) SegmentBodyWithOptions(request *SegmentBodyRequest, runtime *util.RuntimeOptions) (_result *SegmentBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnForm)) {
		query["ReturnForm"] = request.ReturnForm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentBody"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentBodyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentBodyRequest
//
// @return SegmentBodyResponse
func (client *Client) SegmentBody(request *SegmentBodyRequest) (_result *SegmentBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentBodyResponse{}
	_body, _err := client.SegmentBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentBodyAdvance(request *SegmentBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentBodyResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentBodyReq := &SegmentBodyRequest{}
	openapiutil.Convert(request, segmentBodyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentBodyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentBodyResp, _err := client.SegmentBodyWithOptions(segmentBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentBodyResp
	return _result, _err
}

// @param request - SegmentClothRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentClothResponse
func (client *Client) SegmentClothWithOptions(request *SegmentClothRequest, runtime *util.RuntimeOptions) (_result *SegmentClothResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClothClass)) {
		query["ClothClass"] = request.ClothClass
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.OutMode)) {
		query["OutMode"] = request.OutMode
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnForm)) {
		query["ReturnForm"] = request.ReturnForm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentCloth"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentClothResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentClothRequest
//
// @return SegmentClothResponse
func (client *Client) SegmentCloth(request *SegmentClothRequest) (_result *SegmentClothResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentClothResponse{}
	_body, _err := client.SegmentClothWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentClothAdvance(request *SegmentClothAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentClothResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentClothReq := &SegmentClothRequest{}
	openapiutil.Convert(request, segmentClothReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentClothReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentClothResp, _err := client.SegmentClothWithOptions(segmentClothReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentClothResp
	return _result, _err
}

// @param request - SegmentCommodityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentCommodityResponse
func (client *Client) SegmentCommodityWithOptions(request *SegmentCommodityRequest, runtime *util.RuntimeOptions) (_result *SegmentCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnForm)) {
		query["ReturnForm"] = request.ReturnForm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentCommodity"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentCommodityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentCommodityRequest
//
// @return SegmentCommodityResponse
func (client *Client) SegmentCommodity(request *SegmentCommodityRequest) (_result *SegmentCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentCommodityResponse{}
	_body, _err := client.SegmentCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentCommodityAdvance(request *SegmentCommodityAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentCommodityResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentCommodityReq := &SegmentCommodityRequest{}
	openapiutil.Convert(request, segmentCommodityReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentCommodityReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentCommodityResp, _err := client.SegmentCommodityWithOptions(segmentCommodityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentCommodityResp
	return _result, _err
}

// @param request - SegmentCommonImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentCommonImageResponse
func (client *Client) SegmentCommonImageWithOptions(request *SegmentCommonImageRequest, runtime *util.RuntimeOptions) (_result *SegmentCommonImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnForm)) {
		query["ReturnForm"] = request.ReturnForm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentCommonImage"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentCommonImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentCommonImageRequest
//
// @return SegmentCommonImageResponse
func (client *Client) SegmentCommonImage(request *SegmentCommonImageRequest) (_result *SegmentCommonImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentCommonImageResponse{}
	_body, _err := client.SegmentCommonImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentCommonImageAdvance(request *SegmentCommonImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentCommonImageResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentCommonImageReq := &SegmentCommonImageRequest{}
	openapiutil.Convert(request, segmentCommonImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentCommonImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentCommonImageResp, _err := client.SegmentCommonImageWithOptions(segmentCommonImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentCommonImageResp
	return _result, _err
}

// @param request - SegmentFoodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentFoodResponse
func (client *Client) SegmentFoodWithOptions(request *SegmentFoodRequest, runtime *util.RuntimeOptions) (_result *SegmentFoodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnForm)) {
		query["ReturnForm"] = request.ReturnForm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentFood"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentFoodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentFoodRequest
//
// @return SegmentFoodResponse
func (client *Client) SegmentFood(request *SegmentFoodRequest) (_result *SegmentFoodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentFoodResponse{}
	_body, _err := client.SegmentFoodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentFoodAdvance(request *SegmentFoodAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentFoodResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentFoodReq := &SegmentFoodRequest{}
	openapiutil.Convert(request, segmentFoodReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentFoodReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentFoodResp, _err := client.SegmentFoodWithOptions(segmentFoodReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentFoodResp
	return _result, _err
}

// @param request - SegmentHDBodyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentHDBodyResponse
func (client *Client) SegmentHDBodyWithOptions(request *SegmentHDBodyRequest, runtime *util.RuntimeOptions) (_result *SegmentHDBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentHDBody"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentHDBodyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentHDBodyRequest
//
// @return SegmentHDBodyResponse
func (client *Client) SegmentHDBody(request *SegmentHDBodyRequest) (_result *SegmentHDBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHDBodyResponse{}
	_body, _err := client.SegmentHDBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHDBodyAdvance(request *SegmentHDBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHDBodyResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentHDBodyReq := &SegmentHDBodyRequest{}
	openapiutil.Convert(request, segmentHDBodyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentHDBodyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentHDBodyResp, _err := client.SegmentHDBodyWithOptions(segmentHDBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHDBodyResp
	return _result, _err
}

// @param request - SegmentHDCommonImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentHDCommonImageResponse
func (client *Client) SegmentHDCommonImageWithOptions(request *SegmentHDCommonImageRequest, runtime *util.RuntimeOptions) (_result *SegmentHDCommonImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentHDCommonImage"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentHDCommonImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentHDCommonImageRequest
//
// @return SegmentHDCommonImageResponse
func (client *Client) SegmentHDCommonImage(request *SegmentHDCommonImageRequest) (_result *SegmentHDCommonImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHDCommonImageResponse{}
	_body, _err := client.SegmentHDCommonImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHDCommonImageAdvance(request *SegmentHDCommonImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHDCommonImageResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentHDCommonImageReq := &SegmentHDCommonImageRequest{}
	openapiutil.Convert(request, segmentHDCommonImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
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
			Content:     request.ImageUrlObject,
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
		segmentHDCommonImageReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentHDCommonImageResp, _err := client.SegmentHDCommonImageWithOptions(segmentHDCommonImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHDCommonImageResp
	return _result, _err
}

// @param request - SegmentHDSkyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentHDSkyResponse
func (client *Client) SegmentHDSkyWithOptions(request *SegmentHDSkyRequest, runtime *util.RuntimeOptions) (_result *SegmentHDSkyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentHDSky"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentHDSkyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentHDSkyRequest
//
// @return SegmentHDSkyResponse
func (client *Client) SegmentHDSky(request *SegmentHDSkyRequest) (_result *SegmentHDSkyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHDSkyResponse{}
	_body, _err := client.SegmentHDSkyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHDSkyAdvance(request *SegmentHDSkyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHDSkyResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentHDSkyReq := &SegmentHDSkyRequest{}
	openapiutil.Convert(request, segmentHDSkyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentHDSkyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentHDSkyResp, _err := client.SegmentHDSkyWithOptions(segmentHDSkyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHDSkyResp
	return _result, _err
}

// @param request - SegmentHairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentHairResponse
func (client *Client) SegmentHairWithOptions(request *SegmentHairRequest, runtime *util.RuntimeOptions) (_result *SegmentHairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentHair"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentHairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentHairRequest
//
// @return SegmentHairResponse
func (client *Client) SegmentHair(request *SegmentHairRequest) (_result *SegmentHairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHairResponse{}
	_body, _err := client.SegmentHairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHairAdvance(request *SegmentHairAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHairResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentHairReq := &SegmentHairRequest{}
	openapiutil.Convert(request, segmentHairReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentHairReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentHairResp, _err := client.SegmentHairWithOptions(segmentHairReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHairResp
	return _result, _err
}

// @param request - SegmentHeadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentHeadResponse
func (client *Client) SegmentHeadWithOptions(request *SegmentHeadRequest, runtime *util.RuntimeOptions) (_result *SegmentHeadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnForm)) {
		query["ReturnForm"] = request.ReturnForm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentHead"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentHeadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentHeadRequest
//
// @return SegmentHeadResponse
func (client *Client) SegmentHead(request *SegmentHeadRequest) (_result *SegmentHeadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHeadResponse{}
	_body, _err := client.SegmentHeadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHeadAdvance(request *SegmentHeadAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHeadResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentHeadReq := &SegmentHeadRequest{}
	openapiutil.Convert(request, segmentHeadReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentHeadReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentHeadResp, _err := client.SegmentHeadWithOptions(segmentHeadReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHeadResp
	return _result, _err
}

// @param request - SegmentSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentSceneResponse
func (client *Client) SegmentSceneWithOptions(request *SegmentSceneRequest, runtime *util.RuntimeOptions) (_result *SegmentSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentScene"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentSceneRequest
//
// @return SegmentSceneResponse
func (client *Client) SegmentScene(request *SegmentSceneRequest) (_result *SegmentSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentSceneResponse{}
	_body, _err := client.SegmentSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentSceneAdvance(request *SegmentSceneAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentSceneResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentSceneReq := &SegmentSceneRequest{}
	openapiutil.Convert(request, segmentSceneReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentSceneReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentSceneResp, _err := client.SegmentSceneWithOptions(segmentSceneReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentSceneResp
	return _result, _err
}

// @param request - SegmentSkinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentSkinResponse
func (client *Client) SegmentSkinWithOptions(request *SegmentSkinRequest, runtime *util.RuntimeOptions) (_result *SegmentSkinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.URL)) {
		body["URL"] = request.URL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentSkin"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentSkinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentSkinRequest
//
// @return SegmentSkinResponse
func (client *Client) SegmentSkin(request *SegmentSkinRequest) (_result *SegmentSkinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentSkinResponse{}
	_body, _err := client.SegmentSkinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentSkinAdvance(request *SegmentSkinAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentSkinResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentSkinReq := &SegmentSkinRequest{}
	openapiutil.Convert(request, segmentSkinReq)
	if !tea.BoolValue(util.IsUnset(request.URLObject)) {
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
			Content:     request.URLObject,
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
		segmentSkinReq.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentSkinResp, _err := client.SegmentSkinWithOptions(segmentSkinReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentSkinResp
	return _result, _err
}

// @param request - SegmentSkyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SegmentSkyResponse
func (client *Client) SegmentSkyWithOptions(request *SegmentSkyRequest, runtime *util.RuntimeOptions) (_result *SegmentSkyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentSky"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentSkyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SegmentSkyRequest
//
// @return SegmentSkyResponse
func (client *Client) SegmentSky(request *SegmentSkyRequest) (_result *SegmentSkyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentSkyResponse{}
	_body, _err := client.SegmentSkyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentSkyAdvance(request *SegmentSkyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentSkyResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	segmentSkyReq := &SegmentSkyRequest{}
	openapiutil.Convert(request, segmentSkyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		segmentSkyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentSkyResp, _err := client.SegmentSkyWithOptions(segmentSkyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentSkyResp
	return _result, _err
}
