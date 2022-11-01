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

type ChangeSkyRequest struct {
	ImageURL        *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	ImageURLObject  io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReplaceImageURL *string   `json:"ReplaceImageURL,omitempty" xml:"ReplaceImageURL,omitempty"`
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

func (s *ChangeSkyAdvanceRequest) SetReplaceImageURL(v string) *ChangeSkyAdvanceRequest {
	s.ReplaceImageURL = &v
	return s
}

type ChangeSkyResponseBody struct {
	Data      *ChangeSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeSkyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Data      *ParseFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Elements       []*ParseFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	OriginImageURL *string                              `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
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
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ParseFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ImageURL     *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	ImageURLObject     io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *RefineMaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefineMaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SegmentAnimalRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentAnimalRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalRequest) GoString() string {
	return s.String()
}

func (s *SegmentAnimalRequest) SetImageURL(v string) *SegmentAnimalRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentAnimalRequest) SetReturnForm(v string) *SegmentAnimalRequest {
	s.ReturnForm = &v
	return s
}

type SegmentAnimalAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentAnimalAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentAnimalAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentAnimalAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentAnimalAdvanceRequest) SetReturnForm(v string) *SegmentAnimalAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentAnimalResponseBody struct {
	Data      *SegmentAnimalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentAnimalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentAnimalResponseBody) SetData(v *SegmentAnimalResponseBodyData) *SegmentAnimalResponseBody {
	s.Data = v
	return s
}

func (s *SegmentAnimalResponseBody) SetRequestId(v string) *SegmentAnimalResponseBody {
	s.RequestId = &v
	return s
}

type SegmentAnimalResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentAnimalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentAnimalResponseBodyData) SetImageURL(v string) *SegmentAnimalResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentAnimalResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentAnimalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentAnimalResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalResponse) GoString() string {
	return s.String()
}

func (s *SegmentAnimalResponse) SetHeaders(v map[string]*string) *SegmentAnimalResponse {
	s.Headers = v
	return s
}

func (s *SegmentAnimalResponse) SetStatusCode(v int32) *SegmentAnimalResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentAnimalResponse) SetBody(v *SegmentAnimalResponseBody) *SegmentAnimalResponse {
	s.Body = v
	return s
}

type SegmentBodyRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
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
	Data      *SegmentBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentClothRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothRequest) GoString() string {
	return s.String()
}

func (s *SegmentClothRequest) SetImageURL(v string) *SegmentClothRequest {
	s.ImageURL = &v
	return s
}

type SegmentClothAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentClothAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentClothAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentClothAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentClothResponseBody struct {
	Data      *SegmentClothResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentClothResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseBodyDataElements) SetImageURL(v string) *SegmentClothResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

type SegmentClothResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentClothResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
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
	Data      *SegmentCommodityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
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
	Data      *SegmentCommonImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentCommonImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SegmentFaceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFaceRequest) SetImageURL(v string) *SegmentFaceRequest {
	s.ImageURL = &v
	return s
}

type SegmentFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFaceAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentFaceResponseBody struct {
	Data      *SegmentFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentFaceResponseBody) SetData(v *SegmentFaceResponseBodyData) *SegmentFaceResponseBody {
	s.Data = v
	return s
}

func (s *SegmentFaceResponseBody) SetRequestId(v string) *SegmentFaceResponseBody {
	s.RequestId = &v
	return s
}

type SegmentFaceResponseBodyData struct {
	Elements []*SegmentFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentFaceResponseBodyData) SetElements(v []*SegmentFaceResponseBodyDataElements) *SegmentFaceResponseBodyData {
	s.Elements = v
	return s
}

type SegmentFaceResponseBodyDataElements struct {
	Height   *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Width    *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	X        *int32  `json:"X,omitempty" xml:"X,omitempty"`
	Y        *int32  `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SegmentFaceResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentFaceResponseBodyDataElements) SetHeight(v int32) *SegmentFaceResponseBodyDataElements {
	s.Height = &v
	return s
}

func (s *SegmentFaceResponseBodyDataElements) SetImageURL(v string) *SegmentFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentFaceResponseBodyDataElements) SetWidth(v int32) *SegmentFaceResponseBodyDataElements {
	s.Width = &v
	return s
}

func (s *SegmentFaceResponseBodyDataElements) SetX(v int32) *SegmentFaceResponseBodyDataElements {
	s.X = &v
	return s
}

func (s *SegmentFaceResponseBodyDataElements) SetY(v int32) *SegmentFaceResponseBodyDataElements {
	s.Y = &v
	return s
}

type SegmentFaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceResponse) GoString() string {
	return s.String()
}

func (s *SegmentFaceResponse) SetHeaders(v map[string]*string) *SegmentFaceResponse {
	s.Headers = v
	return s
}

func (s *SegmentFaceResponse) SetStatusCode(v int32) *SegmentFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentFaceResponse) SetBody(v *SegmentFaceResponseBody) *SegmentFaceResponse {
	s.Body = v
	return s
}

type SegmentFoodRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
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
	Data      *SegmentFoodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentFoodResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SegmentFurnitureRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentFurnitureRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureRequest) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureRequest) SetImageURL(v string) *SegmentFurnitureRequest {
	s.ImageURL = &v
	return s
}

type SegmentFurnitureAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentFurnitureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentFurnitureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentFurnitureResponseBody struct {
	Data      *SegmentFurnitureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentFurnitureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureResponseBody) SetData(v *SegmentFurnitureResponseBodyData) *SegmentFurnitureResponseBody {
	s.Data = v
	return s
}

func (s *SegmentFurnitureResponseBody) SetRequestId(v string) *SegmentFurnitureResponseBody {
	s.RequestId = &v
	return s
}

type SegmentFurnitureResponseBodyData struct {
	Elements []*SegmentFurnitureResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentFurnitureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureResponseBodyData) SetElements(v []*SegmentFurnitureResponseBodyDataElements) *SegmentFurnitureResponseBodyData {
	s.Elements = v
	return s
}

type SegmentFurnitureResponseBodyDataElements struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentFurnitureResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureResponseBodyDataElements) SetImageURL(v string) *SegmentFurnitureResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

type SegmentFurnitureResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentFurnitureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentFurnitureResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureResponse) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureResponse) SetHeaders(v map[string]*string) *SegmentFurnitureResponse {
	s.Headers = v
	return s
}

func (s *SegmentFurnitureResponse) SetStatusCode(v int32) *SegmentFurnitureResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentFurnitureResponse) SetBody(v *SegmentFurnitureResponseBody) *SegmentFurnitureResponse {
	s.Body = v
	return s
}

type SegmentHDBodyRequest struct {
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
	Data      *SegmentHDBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentHDBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Data      *SegmentHDCommonImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentHDCommonImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Data      *SegmentHDSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentHDSkyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Data      *SegmentHairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Height   *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Width    *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	X        *int32  `json:"X,omitempty" xml:"X,omitempty"`
	Y        *int32  `json:"Y,omitempty" xml:"Y,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentHairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
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
	Data      *SegmentHeadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Height   *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Width    *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	X        *int32  `json:"X,omitempty" xml:"X,omitempty"`
	Y        *int32  `json:"Y,omitempty" xml:"Y,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentHeadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SegmentLogoRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentLogoRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoRequest) GoString() string {
	return s.String()
}

func (s *SegmentLogoRequest) SetImageURL(v string) *SegmentLogoRequest {
	s.ImageURL = &v
	return s
}

type SegmentLogoAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentLogoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentLogoAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentLogoAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentLogoResponseBody struct {
	Data      *SegmentLogoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentLogoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentLogoResponseBody) SetData(v *SegmentLogoResponseBodyData) *SegmentLogoResponseBody {
	s.Data = v
	return s
}

func (s *SegmentLogoResponseBody) SetRequestId(v string) *SegmentLogoResponseBody {
	s.RequestId = &v
	return s
}

type SegmentLogoResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentLogoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentLogoResponseBodyData) SetImageURL(v string) *SegmentLogoResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentLogoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentLogoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentLogoResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoResponse) GoString() string {
	return s.String()
}

func (s *SegmentLogoResponse) SetHeaders(v map[string]*string) *SegmentLogoResponse {
	s.Headers = v
	return s
}

func (s *SegmentLogoResponse) SetStatusCode(v int32) *SegmentLogoResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentLogoResponse) SetBody(v *SegmentLogoResponseBody) *SegmentLogoResponse {
	s.Body = v
	return s
}

type SegmentSceneRequest struct {
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
	Data      *SegmentSceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Data      *SegmentSkinResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentSkinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Data      *SegmentSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentSkyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SegmentVehicleRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentVehicleRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleRequest) GoString() string {
	return s.String()
}

func (s *SegmentVehicleRequest) SetImageURL(v string) *SegmentVehicleRequest {
	s.ImageURL = &v
	return s
}

type SegmentVehicleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentVehicleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentVehicleAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentVehicleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentVehicleResponseBody struct {
	Data      *SegmentVehicleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentVehicleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentVehicleResponseBody) SetData(v *SegmentVehicleResponseBodyData) *SegmentVehicleResponseBody {
	s.Data = v
	return s
}

func (s *SegmentVehicleResponseBody) SetRequestId(v string) *SegmentVehicleResponseBody {
	s.RequestId = &v
	return s
}

type SegmentVehicleResponseBodyData struct {
	Elements []*SegmentVehicleResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentVehicleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentVehicleResponseBodyData) SetElements(v []*SegmentVehicleResponseBodyDataElements) *SegmentVehicleResponseBodyData {
	s.Elements = v
	return s
}

type SegmentVehicleResponseBodyDataElements struct {
	ImageURL       *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OriginImageURL *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
}

func (s SegmentVehicleResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentVehicleResponseBodyDataElements) SetImageURL(v string) *SegmentVehicleResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentVehicleResponseBodyDataElements) SetOriginImageURL(v string) *SegmentVehicleResponseBodyDataElements {
	s.OriginImageURL = &v
	return s
}

type SegmentVehicleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentVehicleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentVehicleResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleResponse) GoString() string {
	return s.String()
}

func (s *SegmentVehicleResponse) SetHeaders(v map[string]*string) *SegmentVehicleResponse {
	s.Headers = v
	return s
}

func (s *SegmentVehicleResponse) SetStatusCode(v int32) *SegmentVehicleResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentVehicleResponse) SetBody(v *SegmentVehicleResponseBody) *SegmentVehicleResponse {
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
		Product:  tea.String("imageseg"),
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

	changeSkyResp, _err := client.ChangeSkyWithOptions(changeSkyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = changeSkyResp
	return _result, _err
}

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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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

func (client *Client) SegmentAnimalWithOptions(request *SegmentAnimalRequest, runtime *util.RuntimeOptions) (_result *SegmentAnimalResponse, _err error) {
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
		Action:      tea.String("SegmentAnimal"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentAnimalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentAnimal(request *SegmentAnimalRequest) (_result *SegmentAnimalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentAnimalResponse{}
	_body, _err := client.SegmentAnimalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentAnimalAdvance(request *SegmentAnimalAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentAnimalResponse, _err error) {
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
		Product:  tea.String("imageseg"),
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
	segmentAnimalReq := &SegmentAnimalRequest{}
	openapiutil.Convert(request, segmentAnimalReq)
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
		segmentAnimalReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentAnimalResp, _err := client.SegmentAnimalWithOptions(segmentAnimalReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentAnimalResp
	return _result, _err
}

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
		Product:  tea.String("imageseg"),
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

func (client *Client) SegmentClothWithOptions(request *SegmentClothRequest, runtime *util.RuntimeOptions) (_result *SegmentClothResponse, _err error) {
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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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

func (client *Client) SegmentFaceWithOptions(request *SegmentFaceRequest, runtime *util.RuntimeOptions) (_result *SegmentFaceResponse, _err error) {
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
		Action:      tea.String("SegmentFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentFace(request *SegmentFaceRequest) (_result *SegmentFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentFaceResponse{}
	_body, _err := client.SegmentFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentFaceAdvance(request *SegmentFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentFaceResponse, _err error) {
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
		Product:  tea.String("imageseg"),
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
	segmentFaceReq := &SegmentFaceRequest{}
	openapiutil.Convert(request, segmentFaceReq)
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
		segmentFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentFaceResp, _err := client.SegmentFaceWithOptions(segmentFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentFaceResp
	return _result, _err
}

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
		Product:  tea.String("imageseg"),
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

func (client *Client) SegmentFurnitureWithOptions(request *SegmentFurnitureRequest, runtime *util.RuntimeOptions) (_result *SegmentFurnitureResponse, _err error) {
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
		Action:      tea.String("SegmentFurniture"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentFurnitureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentFurniture(request *SegmentFurnitureRequest) (_result *SegmentFurnitureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentFurnitureResponse{}
	_body, _err := client.SegmentFurnitureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentFurnitureAdvance(request *SegmentFurnitureAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentFurnitureResponse, _err error) {
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
		Product:  tea.String("imageseg"),
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
	segmentFurnitureReq := &SegmentFurnitureRequest{}
	openapiutil.Convert(request, segmentFurnitureReq)
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
		segmentFurnitureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentFurnitureResp, _err := client.SegmentFurnitureWithOptions(segmentFurnitureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentFurnitureResp
	return _result, _err
}

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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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

func (client *Client) SegmentLogoWithOptions(request *SegmentLogoRequest, runtime *util.RuntimeOptions) (_result *SegmentLogoResponse, _err error) {
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
		Action:      tea.String("SegmentLogo"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentLogoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentLogo(request *SegmentLogoRequest) (_result *SegmentLogoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentLogoResponse{}
	_body, _err := client.SegmentLogoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentLogoAdvance(request *SegmentLogoAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentLogoResponse, _err error) {
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
		Product:  tea.String("imageseg"),
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
	segmentLogoReq := &SegmentLogoRequest{}
	openapiutil.Convert(request, segmentLogoReq)
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
		segmentLogoReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentLogoResp, _err := client.SegmentLogoWithOptions(segmentLogoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentLogoResp
	return _result, _err
}

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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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
		Product:  tea.String("imageseg"),
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

func (client *Client) SegmentVehicleWithOptions(request *SegmentVehicleRequest, runtime *util.RuntimeOptions) (_result *SegmentVehicleResponse, _err error) {
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
		Action:      tea.String("SegmentVehicle"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentVehicleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentVehicle(request *SegmentVehicleRequest) (_result *SegmentVehicleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentVehicleResponse{}
	_body, _err := client.SegmentVehicleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentVehicleAdvance(request *SegmentVehicleAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentVehicleResponse, _err error) {
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
		Product:  tea.String("imageseg"),
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
	segmentVehicleReq := &SegmentVehicleRequest{}
	openapiutil.Convert(request, segmentVehicleReq)
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
		segmentVehicleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	segmentVehicleResp, _err := client.SegmentVehicleWithOptions(segmentVehicleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentVehicleResp
	return _result, _err
}
