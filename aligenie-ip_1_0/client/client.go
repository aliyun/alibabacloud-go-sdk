// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddCartoonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddCartoonHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCartoonHeaders) GoString() string {
	return s.String()
}

func (s *AddCartoonHeaders) SetCommonHeaders(v map[string]*string) *AddCartoonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCartoonHeaders) SetXAcsAligenieAccessToken(v string) *AddCartoonHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddCartoonHeaders) SetAuthorization(v string) *AddCartoonHeaders {
	s.Authorization = &v
	return s
}

type AddCartoonRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0***eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40c804***697
	StartVideoMd5 *string `json:"StartVideoMd5,omitempty" xml:"StartVideoMd5,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://***.mp4
	StartVideoUrl *string `json:"StartVideoUrl,omitempty" xml:"StartVideoUrl,omitempty"`
}

func (s AddCartoonRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCartoonRequest) GoString() string {
	return s.String()
}

func (s *AddCartoonRequest) SetHotelId(v string) *AddCartoonRequest {
	s.HotelId = &v
	return s
}

func (s *AddCartoonRequest) SetStartVideoMd5(v string) *AddCartoonRequest {
	s.StartVideoMd5 = &v
	return s
}

func (s *AddCartoonRequest) SetStartVideoUrl(v string) *AddCartoonRequest {
	s.StartVideoUrl = &v
	return s
}

type AddCartoonResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddCartoonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCartoonResponseBody) GoString() string {
	return s.String()
}

func (s *AddCartoonResponseBody) SetMessage(v string) *AddCartoonResponseBody {
	s.Message = &v
	return s
}

func (s *AddCartoonResponseBody) SetRequestId(v string) *AddCartoonResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCartoonResponseBody) SetResult(v bool) *AddCartoonResponseBody {
	s.Result = &v
	return s
}

func (s *AddCartoonResponseBody) SetStatusCode(v int32) *AddCartoonResponseBody {
	s.StatusCode = &v
	return s
}

type AddCartoonResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCartoonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCartoonResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCartoonResponse) GoString() string {
	return s.String()
}

func (s *AddCartoonResponse) SetHeaders(v map[string]*string) *AddCartoonResponse {
	s.Headers = v
	return s
}

func (s *AddCartoonResponse) SetStatusCode(v int32) *AddCartoonResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCartoonResponse) SetBody(v *AddCartoonResponseBody) *AddCartoonResponse {
	s.Body = v
	return s
}

type AddCustomQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddCustomQAHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAHeaders) GoString() string {
	return s.String()
}

func (s *AddCustomQAHeaders) SetCommonHeaders(v map[string]*string) *AddCustomQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCustomQAHeaders) SetXAcsAligenieAccessToken(v string) *AddCustomQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddCustomQAHeaders) SetAuthorization(v string) *AddCustomQAHeaders {
	s.Authorization = &v
	return s
}

type AddCustomQARequest struct {
	Answers []*string `json:"Answers,omitempty" xml:"Answers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId  *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWords []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// example:
	//
	// ***
	MajorQuestion          *string   `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestions []*string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty" type:"Repeated"`
}

func (s AddCustomQARequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQARequest) GoString() string {
	return s.String()
}

func (s *AddCustomQARequest) SetAnswers(v []*string) *AddCustomQARequest {
	s.Answers = v
	return s
}

func (s *AddCustomQARequest) SetHotelId(v string) *AddCustomQARequest {
	s.HotelId = &v
	return s
}

func (s *AddCustomQARequest) SetKeyWords(v []*string) *AddCustomQARequest {
	s.KeyWords = v
	return s
}

func (s *AddCustomQARequest) SetMajorQuestion(v string) *AddCustomQARequest {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQARequest) SetSupplementaryQuestions(v []*string) *AddCustomQARequest {
	s.SupplementaryQuestions = v
	return s
}

type AddCustomQAShrinkRequest struct {
	AnswersShrink *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId        *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWordsShrink *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// ***
	MajorQuestion                *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestionsShrink *string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty"`
}

func (s AddCustomQAShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCustomQAShrinkRequest) SetAnswersShrink(v string) *AddCustomQAShrinkRequest {
	s.AnswersShrink = &v
	return s
}

func (s *AddCustomQAShrinkRequest) SetHotelId(v string) *AddCustomQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddCustomQAShrinkRequest) SetKeyWordsShrink(v string) *AddCustomQAShrinkRequest {
	s.KeyWordsShrink = &v
	return s
}

func (s *AddCustomQAShrinkRequest) SetMajorQuestion(v string) *AddCustomQAShrinkRequest {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQAShrinkRequest) SetSupplementaryQuestionsShrink(v string) *AddCustomQAShrinkRequest {
	s.SupplementaryQuestionsShrink = &v
	return s
}

type AddCustomQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddCustomQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomQAResponseBody) SetMessage(v string) *AddCustomQAResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomQAResponseBody) SetRequestId(v string) *AddCustomQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomQAResponseBody) SetResult(v bool) *AddCustomQAResponseBody {
	s.Result = &v
	return s
}

func (s *AddCustomQAResponseBody) SetStatusCode(v int32) *AddCustomQAResponseBody {
	s.StatusCode = &v
	return s
}

type AddCustomQAResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomQAResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAResponse) GoString() string {
	return s.String()
}

func (s *AddCustomQAResponse) SetHeaders(v map[string]*string) *AddCustomQAResponse {
	s.Headers = v
	return s
}

func (s *AddCustomQAResponse) SetStatusCode(v int32) *AddCustomQAResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomQAResponse) SetBody(v *AddCustomQAResponseBody) *AddCustomQAResponse {
	s.Body = v
	return s
}

type AddCustomQAV2Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddCustomQAV2Headers) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAV2Headers) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2Headers) SetCommonHeaders(v map[string]*string) *AddCustomQAV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *AddCustomQAV2Headers) SetXAcsAligenieAccessToken(v string) *AddCustomQAV2Headers {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddCustomQAV2Headers) SetAuthorization(v string) *AddCustomQAV2Headers {
	s.Authorization = &v
	return s
}

type AddCustomQAV2Request struct {
	// This parameter is required.
	Answers []*string `json:"Answers,omitempty" xml:"Answers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId                *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWords               []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	MajorQuestion          *string   `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestions []*string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty" type:"Repeated"`
}

func (s AddCustomQAV2Request) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAV2Request) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2Request) SetAnswers(v []*string) *AddCustomQAV2Request {
	s.Answers = v
	return s
}

func (s *AddCustomQAV2Request) SetHotelId(v string) *AddCustomQAV2Request {
	s.HotelId = &v
	return s
}

func (s *AddCustomQAV2Request) SetKeyWords(v []*string) *AddCustomQAV2Request {
	s.KeyWords = v
	return s
}

func (s *AddCustomQAV2Request) SetMajorQuestion(v string) *AddCustomQAV2Request {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQAV2Request) SetSupplementaryQuestions(v []*string) *AddCustomQAV2Request {
	s.SupplementaryQuestions = v
	return s
}

type AddCustomQAV2ShrinkRequest struct {
	// This parameter is required.
	AnswersShrink *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId                      *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWordsShrink               *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	MajorQuestion                *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestionsShrink *string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty"`
}

func (s AddCustomQAV2ShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2ShrinkRequest) SetAnswersShrink(v string) *AddCustomQAV2ShrinkRequest {
	s.AnswersShrink = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) SetHotelId(v string) *AddCustomQAV2ShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) SetKeyWordsShrink(v string) *AddCustomQAV2ShrinkRequest {
	s.KeyWordsShrink = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) SetMajorQuestion(v string) *AddCustomQAV2ShrinkRequest {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) SetSupplementaryQuestionsShrink(v string) *AddCustomQAV2ShrinkRequest {
	s.SupplementaryQuestionsShrink = &v
	return s
}

type AddCustomQAV2ResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FAFCD152-4791-5F2F-B0BE-2DC06FD4F05B
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddCustomQAV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddCustomQAV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAV2ResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2ResponseBody) SetMessage(v string) *AddCustomQAV2ResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomQAV2ResponseBody) SetRequestId(v string) *AddCustomQAV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomQAV2ResponseBody) SetResult(v *AddCustomQAV2ResponseBodyResult) *AddCustomQAV2ResponseBody {
	s.Result = v
	return s
}

func (s *AddCustomQAV2ResponseBody) SetStatusCode(v int32) *AddCustomQAV2ResponseBody {
	s.StatusCode = &v
	return s
}

type AddCustomQAV2ResponseBodyResult struct {
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// example:
	//
	// 2023-01-10 10:01:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// a7***83
	HotelId  *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// 8xxx9
	LastOperator  *string `json:"LastOperator,omitempty" xml:"LastOperator,omitempty"`
	MajorQuestion *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	// qaID
	//
	// example:
	//
	// 1
	QaId *string `json:"QaId,omitempty" xml:"QaId,omitempty"`
	// example:
	//
	// 0
	Status                *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplementaryQuestion *string `json:"SupplementaryQuestion,omitempty" xml:"SupplementaryQuestion,omitempty"`
	// example:
	//
	// 2023-01-10 10:01:59
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s AddCustomQAV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2ResponseBodyResult) SetAnswers(v string) *AddCustomQAV2ResponseBodyResult {
	s.Answers = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetCreateTime(v string) *AddCustomQAV2ResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetHotelId(v string) *AddCustomQAV2ResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetKeyWords(v string) *AddCustomQAV2ResponseBodyResult {
	s.KeyWords = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetLastOperator(v string) *AddCustomQAV2ResponseBodyResult {
	s.LastOperator = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetMajorQuestion(v string) *AddCustomQAV2ResponseBodyResult {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetQaId(v string) *AddCustomQAV2ResponseBodyResult {
	s.QaId = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetStatus(v int32) *AddCustomQAV2ResponseBodyResult {
	s.Status = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetSupplementaryQuestion(v string) *AddCustomQAV2ResponseBodyResult {
	s.SupplementaryQuestion = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetUpdateTime(v string) *AddCustomQAV2ResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type AddCustomQAV2Response struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomQAV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomQAV2Response) String() string {
	return tea.Prettify(s)
}

func (s AddCustomQAV2Response) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2Response) SetHeaders(v map[string]*string) *AddCustomQAV2Response {
	s.Headers = v
	return s
}

func (s *AddCustomQAV2Response) SetStatusCode(v int32) *AddCustomQAV2Response {
	s.StatusCode = &v
	return s
}

func (s *AddCustomQAV2Response) SetBody(v *AddCustomQAV2ResponseBody) *AddCustomQAV2Response {
	s.Body = v
	return s
}

type AddMessageTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddMessageTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddMessageTemplateHeaders) GoString() string {
	return s.String()
}

func (s *AddMessageTemplateHeaders) SetCommonHeaders(v map[string]*string) *AddMessageTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMessageTemplateHeaders) SetXAcsAligenieAccessToken(v string) *AddMessageTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddMessageTemplateHeaders) SetAuthorization(v string) *AddMessageTemplateHeaders {
	s.Authorization = &v
	return s
}

type AddMessageTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 这是${hotel}的一个测试模板
	TemplateDetail *string `json:"TemplateDetail,omitempty" xml:"TemplateDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s AddMessageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMessageTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddMessageTemplateRequest) SetTemplateDetail(v string) *AddMessageTemplateRequest {
	s.TemplateDetail = &v
	return s
}

func (s *AddMessageTemplateRequest) SetTemplateName(v string) *AddMessageTemplateRequest {
	s.TemplateName = &v
	return s
}

type AddMessageTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 11
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AddMessageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMessageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddMessageTemplateResponseBody) SetCode(v int32) *AddMessageTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *AddMessageTemplateResponseBody) SetMessage(v string) *AddMessageTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *AddMessageTemplateResponseBody) SetRequestId(v string) *AddMessageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMessageTemplateResponseBody) SetResult(v int64) *AddMessageTemplateResponseBody {
	s.Result = &v
	return s
}

type AddMessageTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMessageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMessageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMessageTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddMessageTemplateResponse) SetHeaders(v map[string]*string) *AddMessageTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddMessageTemplateResponse) SetStatusCode(v int32) *AddMessageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMessageTemplateResponse) SetBody(v *AddMessageTemplateResponseBody) *AddMessageTemplateResponse {
	s.Body = v
	return s
}

type AddOrUpdateDisPlayModesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddOrUpdateDisPlayModesHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateDisPlayModesHeaders) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesHeaders) SetCommonHeaders(v map[string]*string) *AddOrUpdateDisPlayModesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrUpdateDisPlayModesHeaders) SetXAcsAligenieAccessToken(v string) *AddOrUpdateDisPlayModesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddOrUpdateDisPlayModesHeaders) SetAuthorization(v string) *AddOrUpdateDisPlayModesHeaders {
	s.Authorization = &v
	return s
}

type AddOrUpdateDisPlayModesRequest struct {
	// This parameter is required.
	HotelDeviceModeList []*string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s AddOrUpdateDisPlayModesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateDisPlayModesRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesRequest) SetHotelDeviceModeList(v []*string) *AddOrUpdateDisPlayModesRequest {
	s.HotelDeviceModeList = v
	return s
}

func (s *AddOrUpdateDisPlayModesRequest) SetHotelId(v string) *AddOrUpdateDisPlayModesRequest {
	s.HotelId = &v
	return s
}

type AddOrUpdateDisPlayModesShrinkRequest struct {
	// This parameter is required.
	HotelDeviceModeListShrink *string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s AddOrUpdateDisPlayModesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateDisPlayModesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesShrinkRequest) SetHotelDeviceModeListShrink(v string) *AddOrUpdateDisPlayModesShrinkRequest {
	s.HotelDeviceModeListShrink = &v
	return s
}

func (s *AddOrUpdateDisPlayModesShrinkRequest) SetHotelId(v string) *AddOrUpdateDisPlayModesShrinkRequest {
	s.HotelId = &v
	return s
}

type AddOrUpdateDisPlayModesResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddOrUpdateDisPlayModesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateDisPlayModesResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesResponseBody) SetMessage(v string) *AddOrUpdateDisPlayModesResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponseBody) SetRequestId(v string) *AddOrUpdateDisPlayModesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponseBody) SetResult(v bool) *AddOrUpdateDisPlayModesResponseBody {
	s.Result = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponseBody) SetStatusCode(v int32) *AddOrUpdateDisPlayModesResponseBody {
	s.StatusCode = &v
	return s
}

type AddOrUpdateDisPlayModesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateDisPlayModesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateDisPlayModesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateDisPlayModesResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesResponse) SetHeaders(v map[string]*string) *AddOrUpdateDisPlayModesResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateDisPlayModesResponse) SetStatusCode(v int32) *AddOrUpdateDisPlayModesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponse) SetBody(v *AddOrUpdateDisPlayModesResponseBody) *AddOrUpdateDisPlayModesResponse {
	s.Body = v
	return s
}

type AddOrUpdateHotelSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddOrUpdateHotelSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateHotelSettingHeaders) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingHeaders) SetCommonHeaders(v map[string]*string) *AddOrUpdateHotelSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrUpdateHotelSettingHeaders) SetXAcsAligenieAccessToken(v string) *AddOrUpdateHotelSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddOrUpdateHotelSettingHeaders) SetAuthorization(v string) *AddOrUpdateHotelSettingHeaders {
	s.Authorization = &v
	return s
}

type AddOrUpdateHotelSettingRequest struct {
	HotelDeviceModeList []*string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty" type:"Repeated"`
	// example:
	//
	// a7a3***013
	HotelId          *string                                         `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelScreenSaver *AddOrUpdateHotelSettingRequestHotelScreenSaver `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty" type:"Struct"`
	NightMode        *AddOrUpdateHotelSettingRequestNightMode        `json:"NightMode,omitempty" xml:"NightMode,omitempty" type:"Struct"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddOrUpdateHotelSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateHotelSettingRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingRequest) SetHotelDeviceModeList(v []*string) *AddOrUpdateHotelSettingRequest {
	s.HotelDeviceModeList = v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetHotelId(v string) *AddOrUpdateHotelSettingRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetHotelScreenSaver(v *AddOrUpdateHotelSettingRequestHotelScreenSaver) *AddOrUpdateHotelSettingRequest {
	s.HotelScreenSaver = v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetNightMode(v *AddOrUpdateHotelSettingRequestNightMode) *AddOrUpdateHotelSettingRequest {
	s.NightMode = v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetSettingType(v string) *AddOrUpdateHotelSettingRequest {
	s.SettingType = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetValue(v string) *AddOrUpdateHotelSettingRequest {
	s.Value = &v
	return s
}

type AddOrUpdateHotelSettingRequestHotelScreenSaver struct {
	// example:
	//
	// https://a****jpg
	ScreenSaverPicUrl *string `json:"ScreenSaverPicUrl,omitempty" xml:"ScreenSaverPicUrl,omitempty"`
	// example:
	//
	// common-weather
	ScreenSaverStyle *string `json:"ScreenSaverStyle,omitempty" xml:"ScreenSaverStyle,omitempty"`
}

func (s AddOrUpdateHotelSettingRequestHotelScreenSaver) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateHotelSettingRequestHotelScreenSaver) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingRequestHotelScreenSaver) SetScreenSaverPicUrl(v string) *AddOrUpdateHotelSettingRequestHotelScreenSaver {
	s.ScreenSaverPicUrl = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestHotelScreenSaver) SetScreenSaverStyle(v string) *AddOrUpdateHotelSettingRequestHotelScreenSaver {
	s.ScreenSaverStyle = &v
	return s
}

type AddOrUpdateHotelSettingRequestNightMode struct {
	DefaultBright *string `json:"DefaultBright,omitempty" xml:"DefaultBright,omitempty"`
	DefaultVolume *string `json:"DefaultVolume,omitempty" xml:"DefaultVolume,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 22:00
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// screenoff
	StandbyAction *string `json:"StandbyAction,omitempty" xml:"StandbyAction,omitempty"`
	// example:
	//
	// 7:00
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s AddOrUpdateHotelSettingRequestNightMode) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateHotelSettingRequestNightMode) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetDefaultBright(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.DefaultBright = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetDefaultVolume(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.DefaultVolume = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetEnable(v bool) *AddOrUpdateHotelSettingRequestNightMode {
	s.Enable = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetEnd(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.End = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetStandbyAction(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.StandbyAction = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetStart(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.Start = &v
	return s
}

type AddOrUpdateHotelSettingShrinkRequest struct {
	HotelDeviceModeListShrink *string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty"`
	// example:
	//
	// a7a3***013
	HotelId                *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelScreenSaverShrink *string `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty"`
	NightModeShrink        *string `json:"NightMode,omitempty" xml:"NightMode,omitempty"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddOrUpdateHotelSettingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateHotelSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetHotelDeviceModeListShrink(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.HotelDeviceModeListShrink = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetHotelId(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetHotelScreenSaverShrink(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.HotelScreenSaverShrink = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetNightModeShrink(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.NightModeShrink = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetSettingType(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.SettingType = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetValue(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.Value = &v
	return s
}

type AddOrUpdateHotelSettingResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddOrUpdateHotelSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateHotelSettingResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingResponseBody) SetMessage(v string) *AddOrUpdateHotelSettingResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponseBody) SetRequestId(v string) *AddOrUpdateHotelSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponseBody) SetResult(v bool) *AddOrUpdateHotelSettingResponseBody {
	s.Result = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponseBody) SetStatusCode(v int32) *AddOrUpdateHotelSettingResponseBody {
	s.StatusCode = &v
	return s
}

type AddOrUpdateHotelSettingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateHotelSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateHotelSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateHotelSettingResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingResponse) SetHeaders(v map[string]*string) *AddOrUpdateHotelSettingResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateHotelSettingResponse) SetStatusCode(v int32) *AddOrUpdateHotelSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponse) SetBody(v *AddOrUpdateHotelSettingResponseBody) *AddOrUpdateHotelSettingResponse {
	s.Body = v
	return s
}

type AddOrUpdateScreenSaverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddOrUpdateScreenSaverHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateScreenSaverHeaders) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverHeaders) SetCommonHeaders(v map[string]*string) *AddOrUpdateScreenSaverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrUpdateScreenSaverHeaders) SetXAcsAligenieAccessToken(v string) *AddOrUpdateScreenSaverHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddOrUpdateScreenSaverHeaders) SetAuthorization(v string) *AddOrUpdateScreenSaverHeaders {
	s.Authorization = &v
	return s
}

type AddOrUpdateScreenSaverRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	HotelScreenSaver *AddOrUpdateScreenSaverRequestHotelScreenSaver `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty" type:"Struct"`
}

func (s AddOrUpdateScreenSaverRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateScreenSaverRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverRequest) SetHotelId(v string) *AddOrUpdateScreenSaverRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateScreenSaverRequest) SetHotelScreenSaver(v *AddOrUpdateScreenSaverRequestHotelScreenSaver) *AddOrUpdateScreenSaverRequest {
	s.HotelScreenSaver = v
	return s
}

type AddOrUpdateScreenSaverRequestHotelScreenSaver struct {
	// example:
	//
	// xxx.png
	ScreenSaverPicUrl *string `json:"ScreenSaverPicUrl,omitempty" xml:"ScreenSaverPicUrl,omitempty"`
	// example:
	//
	// common-weather
	ScreenSaverStyle *string `json:"ScreenSaverStyle,omitempty" xml:"ScreenSaverStyle,omitempty"`
}

func (s AddOrUpdateScreenSaverRequestHotelScreenSaver) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateScreenSaverRequestHotelScreenSaver) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverRequestHotelScreenSaver) SetScreenSaverPicUrl(v string) *AddOrUpdateScreenSaverRequestHotelScreenSaver {
	s.ScreenSaverPicUrl = &v
	return s
}

func (s *AddOrUpdateScreenSaverRequestHotelScreenSaver) SetScreenSaverStyle(v string) *AddOrUpdateScreenSaverRequestHotelScreenSaver {
	s.ScreenSaverStyle = &v
	return s
}

type AddOrUpdateScreenSaverShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	HotelScreenSaverShrink *string `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty"`
}

func (s AddOrUpdateScreenSaverShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateScreenSaverShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverShrinkRequest) SetHotelId(v string) *AddOrUpdateScreenSaverShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateScreenSaverShrinkRequest) SetHotelScreenSaverShrink(v string) *AddOrUpdateScreenSaverShrinkRequest {
	s.HotelScreenSaverShrink = &v
	return s
}

type AddOrUpdateScreenSaverResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4EED***9661
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddOrUpdateScreenSaverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateScreenSaverResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverResponseBody) SetMessage(v string) *AddOrUpdateScreenSaverResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponseBody) SetRequestId(v string) *AddOrUpdateScreenSaverResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponseBody) SetResult(v bool) *AddOrUpdateScreenSaverResponseBody {
	s.Result = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponseBody) SetStatusCode(v int32) *AddOrUpdateScreenSaverResponseBody {
	s.StatusCode = &v
	return s
}

type AddOrUpdateScreenSaverResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateScreenSaverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateScreenSaverResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateScreenSaverResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverResponse) SetHeaders(v map[string]*string) *AddOrUpdateScreenSaverResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateScreenSaverResponse) SetStatusCode(v int32) *AddOrUpdateScreenSaverResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponse) SetBody(v *AddOrUpdateScreenSaverResponseBody) *AddOrUpdateScreenSaverResponse {
	s.Body = v
	return s
}

type AddOrUpdateWelcomeTextHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddOrUpdateWelcomeTextHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateWelcomeTextHeaders) GoString() string {
	return s.String()
}

func (s *AddOrUpdateWelcomeTextHeaders) SetCommonHeaders(v map[string]*string) *AddOrUpdateWelcomeTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrUpdateWelcomeTextHeaders) SetXAcsAligenieAccessToken(v string) *AddOrUpdateWelcomeTextHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddOrUpdateWelcomeTextHeaders) SetAuthorization(v string) *AddOrUpdateWelcomeTextHeaders {
	s.Authorization = &v
	return s
}

type AddOrUpdateWelcomeTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://ailabsaicloudservice.alicdn.com/tmp/a.wav
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	// This parameter is required.
	WelcomeText *string `json:"WelcomeText,omitempty" xml:"WelcomeText,omitempty"`
}

func (s AddOrUpdateWelcomeTextRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateWelcomeTextRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateWelcomeTextRequest) SetHotelId(v string) *AddOrUpdateWelcomeTextRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateWelcomeTextRequest) SetMusicUrl(v string) *AddOrUpdateWelcomeTextRequest {
	s.MusicUrl = &v
	return s
}

func (s *AddOrUpdateWelcomeTextRequest) SetWelcomeText(v string) *AddOrUpdateWelcomeTextRequest {
	s.WelcomeText = &v
	return s
}

type AddOrUpdateWelcomeTextResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddOrUpdateWelcomeTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateWelcomeTextResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetExtentions(v map[string]interface{}) *AddOrUpdateWelcomeTextResponseBody {
	s.Extentions = v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetMessage(v string) *AddOrUpdateWelcomeTextResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetRequestId(v string) *AddOrUpdateWelcomeTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetResult(v bool) *AddOrUpdateWelcomeTextResponseBody {
	s.Result = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetStatusCode(v int32) *AddOrUpdateWelcomeTextResponseBody {
	s.StatusCode = &v
	return s
}

type AddOrUpdateWelcomeTextResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateWelcomeTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateWelcomeTextResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrUpdateWelcomeTextResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateWelcomeTextResponse) SetHeaders(v map[string]*string) *AddOrUpdateWelcomeTextResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateWelcomeTextResponse) SetStatusCode(v int32) *AddOrUpdateWelcomeTextResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponse) SetBody(v *AddOrUpdateWelcomeTextResponseBody) *AddOrUpdateWelcomeTextResponse {
	s.Body = v
	return s
}

type AuditHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuditHotelHeaders) String() string {
	return tea.Prettify(s)
}

func (s AuditHotelHeaders) GoString() string {
	return s.String()
}

func (s *AuditHotelHeaders) SetCommonHeaders(v map[string]*string) *AuditHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuditHotelHeaders) SetXAcsAligenieAccessToken(v string) *AuditHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuditHotelHeaders) SetAuthorization(v string) *AuditHotelHeaders {
	s.Authorization = &v
	return s
}

type AuditHotelRequest struct {
	// This parameter is required.
	AuditHotelReq *AuditHotelRequestAuditHotelReq `json:"AuditHotelReq,omitempty" xml:"AuditHotelReq,omitempty" type:"Struct"`
}

func (s AuditHotelRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditHotelRequest) GoString() string {
	return s.String()
}

func (s *AuditHotelRequest) SetAuditHotelReq(v *AuditHotelRequestAuditHotelReq) *AuditHotelRequest {
	s.AuditHotelReq = v
	return s
}

type AuditHotelRequestAuditHotelReq struct {
	// example:
	//
	// 同意
	AuditOpinion *string `json:"AuditOpinion,omitempty" xml:"AuditOpinion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AuditHotelRequestAuditHotelReq) String() string {
	return tea.Prettify(s)
}

func (s AuditHotelRequestAuditHotelReq) GoString() string {
	return s.String()
}

func (s *AuditHotelRequestAuditHotelReq) SetAuditOpinion(v string) *AuditHotelRequestAuditHotelReq {
	s.AuditOpinion = &v
	return s
}

func (s *AuditHotelRequestAuditHotelReq) SetHotelId(v string) *AuditHotelRequestAuditHotelReq {
	s.HotelId = &v
	return s
}

func (s *AuditHotelRequestAuditHotelReq) SetStatus(v int32) *AuditHotelRequestAuditHotelReq {
	s.Status = &v
	return s
}

type AuditHotelShrinkRequest struct {
	// This parameter is required.
	AuditHotelReqShrink *string `json:"AuditHotelReq,omitempty" xml:"AuditHotelReq,omitempty"`
}

func (s AuditHotelShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditHotelShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuditHotelShrinkRequest) SetAuditHotelReqShrink(v string) *AuditHotelShrinkRequest {
	s.AuditHotelReqShrink = &v
	return s
}

type AuditHotelResponseBody struct {
	// example:
	//
	// 200
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AuditHotelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuditHotelResponseBody) GoString() string {
	return s.String()
}

func (s *AuditHotelResponseBody) SetCode(v int32) *AuditHotelResponseBody {
	s.Code = &v
	return s
}

func (s *AuditHotelResponseBody) SetMessage(v string) *AuditHotelResponseBody {
	s.Message = &v
	return s
}

func (s *AuditHotelResponseBody) SetRequestId(v string) *AuditHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuditHotelResponseBody) SetResult(v bool) *AuditHotelResponseBody {
	s.Result = &v
	return s
}

type AuditHotelResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditHotelResponse) String() string {
	return tea.Prettify(s)
}

func (s AuditHotelResponse) GoString() string {
	return s.String()
}

func (s *AuditHotelResponse) SetHeaders(v map[string]*string) *AuditHotelResponse {
	s.Headers = v
	return s
}

func (s *AuditHotelResponse) SetStatusCode(v int32) *AuditHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditHotelResponse) SetBody(v *AuditHotelResponseBody) *AuditHotelResponse {
	s.Body = v
	return s
}

type BatchAddHotelRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s BatchAddHotelRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchAddHotelRoomHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomHeaders) SetCommonHeaders(v map[string]*string) *BatchAddHotelRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddHotelRoomHeaders) SetXAcsAligenieAccessToken(v string) *BatchAddHotelRoomHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *BatchAddHotelRoomHeaders) SetAuthorization(v string) *BatchAddHotelRoomHeaders {
	s.Authorization = &v
	return s
}

type BatchAddHotelRoomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNoList []*string `json:"RoomNoList,omitempty" xml:"RoomNoList,omitempty" type:"Repeated"`
}

func (s BatchAddHotelRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddHotelRoomRequest) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomRequest) SetHotelId(v string) *BatchAddHotelRoomRequest {
	s.HotelId = &v
	return s
}

func (s *BatchAddHotelRoomRequest) SetRoomNoList(v []*string) *BatchAddHotelRoomRequest {
	s.RoomNoList = v
	return s
}

type BatchAddHotelRoomShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNoListShrink *string `json:"RoomNoList,omitempty" xml:"RoomNoList,omitempty"`
}

func (s BatchAddHotelRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddHotelRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomShrinkRequest) SetHotelId(v string) *BatchAddHotelRoomShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *BatchAddHotelRoomShrinkRequest) SetRoomNoListShrink(v string) *BatchAddHotelRoomShrinkRequest {
	s.RoomNoListShrink = &v
	return s
}

type BatchAddHotelRoomResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BatchAddHotelRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddHotelRoomResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomResponseBody) SetCode(v int32) *BatchAddHotelRoomResponseBody {
	s.Code = &v
	return s
}

func (s *BatchAddHotelRoomResponseBody) SetMessage(v string) *BatchAddHotelRoomResponseBody {
	s.Message = &v
	return s
}

func (s *BatchAddHotelRoomResponseBody) SetRequestId(v string) *BatchAddHotelRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddHotelRoomResponseBody) SetResult(v bool) *BatchAddHotelRoomResponseBody {
	s.Result = &v
	return s
}

type BatchAddHotelRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddHotelRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddHotelRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddHotelRoomResponse) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomResponse) SetHeaders(v map[string]*string) *BatchAddHotelRoomResponse {
	s.Headers = v
	return s
}

func (s *BatchAddHotelRoomResponse) SetStatusCode(v int32) *BatchAddHotelRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddHotelRoomResponse) SetBody(v *BatchAddHotelRoomResponseBody) *BatchAddHotelRoomResponse {
	s.Body = v
	return s
}

type BatchDeleteHotelRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s BatchDeleteHotelRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteHotelRoomHeaders) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomHeaders) SetCommonHeaders(v map[string]*string) *BatchDeleteHotelRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchDeleteHotelRoomHeaders) SetXAcsAligenieAccessToken(v string) *BatchDeleteHotelRoomHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *BatchDeleteHotelRoomHeaders) SetAuthorization(v string) *BatchDeleteHotelRoomHeaders {
	s.Authorization = &v
	return s
}

type BatchDeleteHotelRoomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNoList []*string `json:"RoomNoList,omitempty" xml:"RoomNoList,omitempty" type:"Repeated"`
}

func (s BatchDeleteHotelRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteHotelRoomRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomRequest) SetHotelId(v string) *BatchDeleteHotelRoomRequest {
	s.HotelId = &v
	return s
}

func (s *BatchDeleteHotelRoomRequest) SetRoomNoList(v []*string) *BatchDeleteHotelRoomRequest {
	s.RoomNoList = v
	return s
}

type BatchDeleteHotelRoomShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNoListShrink *string `json:"RoomNoList,omitempty" xml:"RoomNoList,omitempty"`
}

func (s BatchDeleteHotelRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteHotelRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomShrinkRequest) SetHotelId(v string) *BatchDeleteHotelRoomShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *BatchDeleteHotelRoomShrinkRequest) SetRoomNoListShrink(v string) *BatchDeleteHotelRoomShrinkRequest {
	s.RoomNoListShrink = &v
	return s
}

type BatchDeleteHotelRoomResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BatchDeleteHotelRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteHotelRoomResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomResponseBody) SetCode(v int32) *BatchDeleteHotelRoomResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteHotelRoomResponseBody) SetMessage(v string) *BatchDeleteHotelRoomResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteHotelRoomResponseBody) SetRequestId(v string) *BatchDeleteHotelRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteHotelRoomResponseBody) SetResult(v bool) *BatchDeleteHotelRoomResponseBody {
	s.Result = &v
	return s
}

type BatchDeleteHotelRoomResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteHotelRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteHotelRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteHotelRoomResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomResponse) SetHeaders(v map[string]*string) *BatchDeleteHotelRoomResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteHotelRoomResponse) SetStatusCode(v int32) *BatchDeleteHotelRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteHotelRoomResponse) SetBody(v *BatchDeleteHotelRoomResponseBody) *BatchDeleteHotelRoomResponse {
	s.Body = v
	return s
}

type CheckoutWithAKHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CheckoutWithAKHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckoutWithAKHeaders) GoString() string {
	return s.String()
}

func (s *CheckoutWithAKHeaders) SetCommonHeaders(v map[string]*string) *CheckoutWithAKHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckoutWithAKHeaders) SetXAcsAligenieAccessToken(v string) *CheckoutWithAKHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CheckoutWithAKHeaders) SetAuthorization(v string) *CheckoutWithAKHeaders {
	s.Authorization = &v
	return s
}

type CheckoutWithAKRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s CheckoutWithAKRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckoutWithAKRequest) GoString() string {
	return s.String()
}

func (s *CheckoutWithAKRequest) SetHotelId(v string) *CheckoutWithAKRequest {
	s.HotelId = &v
	return s
}

func (s *CheckoutWithAKRequest) SetRoomNo(v string) *CheckoutWithAKRequest {
	s.RoomNo = &v
	return s
}

type CheckoutWithAKResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CheckoutWithAKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckoutWithAKResponseBody) GoString() string {
	return s.String()
}

func (s *CheckoutWithAKResponseBody) SetMessage(v string) *CheckoutWithAKResponseBody {
	s.Message = &v
	return s
}

func (s *CheckoutWithAKResponseBody) SetRequestId(v string) *CheckoutWithAKResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckoutWithAKResponseBody) SetResult(v bool) *CheckoutWithAKResponseBody {
	s.Result = &v
	return s
}

func (s *CheckoutWithAKResponseBody) SetStatusCode(v int32) *CheckoutWithAKResponseBody {
	s.StatusCode = &v
	return s
}

type CheckoutWithAKResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckoutWithAKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckoutWithAKResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckoutWithAKResponse) GoString() string {
	return s.String()
}

func (s *CheckoutWithAKResponse) SetHeaders(v map[string]*string) *CheckoutWithAKResponse {
	s.Headers = v
	return s
}

func (s *CheckoutWithAKResponse) SetStatusCode(v int32) *CheckoutWithAKResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckoutWithAKResponse) SetBody(v *CheckoutWithAKResponseBody) *CheckoutWithAKResponse {
	s.Body = v
	return s
}

type ChildAccountAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ChildAccountAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChildAccountAuthHeaders) GoString() string {
	return s.String()
}

func (s *ChildAccountAuthHeaders) SetCommonHeaders(v map[string]*string) *ChildAccountAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChildAccountAuthHeaders) SetXAcsAligenieAccessToken(v string) *ChildAccountAuthHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ChildAccountAuthHeaders) SetAuthorization(v string) *ChildAccountAuthHeaders {
	s.Authorization = &v
	return s
}

type ChildAccountAuthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// lee
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30471753
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEV***E3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s ChildAccountAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s ChildAccountAuthRequest) GoString() string {
	return s.String()
}

func (s *ChildAccountAuthRequest) SetAccount(v string) *ChildAccountAuthRequest {
	s.Account = &v
	return s
}

func (s *ChildAccountAuthRequest) SetAppKey(v string) *ChildAccountAuthRequest {
	s.AppKey = &v
	return s
}

func (s *ChildAccountAuthRequest) SetHotelId(v string) *ChildAccountAuthRequest {
	s.HotelId = &v
	return s
}

func (s *ChildAccountAuthRequest) SetTbOpenId(v string) *ChildAccountAuthRequest {
	s.TbOpenId = &v
	return s
}

type ChildAccountAuthResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3DB51A10-327C-58D3-91DF-3A5A471C51E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ChildAccountAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChildAccountAuthResponseBody) GoString() string {
	return s.String()
}

func (s *ChildAccountAuthResponseBody) SetExtentions(v map[string]interface{}) *ChildAccountAuthResponseBody {
	s.Extentions = v
	return s
}

func (s *ChildAccountAuthResponseBody) SetMessage(v string) *ChildAccountAuthResponseBody {
	s.Message = &v
	return s
}

func (s *ChildAccountAuthResponseBody) SetRequestId(v string) *ChildAccountAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChildAccountAuthResponseBody) SetResult(v bool) *ChildAccountAuthResponseBody {
	s.Result = &v
	return s
}

func (s *ChildAccountAuthResponseBody) SetStatusCode(v int32) *ChildAccountAuthResponseBody {
	s.StatusCode = &v
	return s
}

type ChildAccountAuthResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChildAccountAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChildAccountAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s ChildAccountAuthResponse) GoString() string {
	return s.String()
}

func (s *ChildAccountAuthResponse) SetHeaders(v map[string]*string) *ChildAccountAuthResponse {
	s.Headers = v
	return s
}

func (s *ChildAccountAuthResponse) SetStatusCode(v int32) *ChildAccountAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *ChildAccountAuthResponse) SetBody(v *ChildAccountAuthResponseBody) *ChildAccountAuthResponse {
	s.Body = v
	return s
}

type ControlRoomDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ControlRoomDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ControlRoomDeviceHeaders) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceHeaders) SetCommonHeaders(v map[string]*string) *ControlRoomDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ControlRoomDeviceHeaders) SetXAcsAligenieAccessToken(v string) *ControlRoomDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ControlRoomDeviceHeaders) SetAuthorization(v string) *ControlRoomDeviceHeaders {
	s.Authorization = &v
	return s
}

type ControlRoomDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// thing.attribute.set
	//
	// thing.attribute.adjust
	Cmd         *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	DeviceIndex *int32  `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INFRARED49122575595
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ControlRoomDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ControlRoomDeviceRequest) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceRequest) SetCmd(v string) *ControlRoomDeviceRequest {
	s.Cmd = &v
	return s
}

func (s *ControlRoomDeviceRequest) SetDeviceIndex(v int32) *ControlRoomDeviceRequest {
	s.DeviceIndex = &v
	return s
}

func (s *ControlRoomDeviceRequest) SetDeviceNumber(v string) *ControlRoomDeviceRequest {
	s.DeviceNumber = &v
	return s
}

func (s *ControlRoomDeviceRequest) SetHotelId(v string) *ControlRoomDeviceRequest {
	s.HotelId = &v
	return s
}

func (s *ControlRoomDeviceRequest) SetProperties(v map[string]*string) *ControlRoomDeviceRequest {
	s.Properties = v
	return s
}

func (s *ControlRoomDeviceRequest) SetRoomNo(v string) *ControlRoomDeviceRequest {
	s.RoomNo = &v
	return s
}

type ControlRoomDeviceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// thing.attribute.set
	//
	// thing.attribute.adjust
	Cmd         *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	DeviceIndex *int32  `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INFRARED49122575595
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ControlRoomDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ControlRoomDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceShrinkRequest) SetCmd(v string) *ControlRoomDeviceShrinkRequest {
	s.Cmd = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetDeviceIndex(v int32) *ControlRoomDeviceShrinkRequest {
	s.DeviceIndex = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetDeviceNumber(v string) *ControlRoomDeviceShrinkRequest {
	s.DeviceNumber = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetHotelId(v string) *ControlRoomDeviceShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetPropertiesShrink(v string) *ControlRoomDeviceShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetRoomNo(v string) *ControlRoomDeviceShrinkRequest {
	s.RoomNo = &v
	return s
}

type ControlRoomDeviceResponseBody struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ControlRoomDeviceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ControlRoomDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ControlRoomDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceResponseBody) SetCode(v int32) *ControlRoomDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ControlRoomDeviceResponseBody) SetMessage(v string) *ControlRoomDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ControlRoomDeviceResponseBody) SetRequestId(v string) *ControlRoomDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ControlRoomDeviceResponseBody) SetResult(v *ControlRoomDeviceResponseBodyResult) *ControlRoomDeviceResponseBody {
	s.Result = v
	return s
}

type ControlRoomDeviceResponseBodyResult struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ControlRoomDeviceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ControlRoomDeviceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceResponseBodyResult) SetMessage(v string) *ControlRoomDeviceResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ControlRoomDeviceResponseBodyResult) SetStatus(v int32) *ControlRoomDeviceResponseBodyResult {
	s.Status = &v
	return s
}

type ControlRoomDeviceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ControlRoomDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ControlRoomDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ControlRoomDeviceResponse) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceResponse) SetHeaders(v map[string]*string) *ControlRoomDeviceResponse {
	s.Headers = v
	return s
}

func (s *ControlRoomDeviceResponse) SetStatusCode(v int32) *ControlRoomDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ControlRoomDeviceResponse) SetBody(v *ControlRoomDeviceResponseBody) *ControlRoomDeviceResponse {
	s.Body = v
	return s
}

type CreateHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateHotelHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelHeaders) GoString() string {
	return s.String()
}

func (s *CreateHotelHeaders) SetCommonHeaders(v map[string]*string) *CreateHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateHotelHeaders) SetXAcsAligenieAccessToken(v string) *CreateHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateHotelHeaders) SetAuthorization(v string) *CreateHotelHeaders {
	s.Authorization = &v
	return s
}

type CreateHotelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 333566791
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-1 00:00:00
	EstOpenTime *string `json:"EstOpenTime,omitempty" xml:"EstOpenTime,omitempty"`
	// This parameter is required.
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test@hotel.com
	HotelEmail *string `json:"HotelEmail,omitempty" xml:"HotelEmail,omitempty"`
	// This parameter is required.
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13xxxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// jTO****Rw
	RelatedPk *string `json:"RelatedPk,omitempty" xml:"RelatedPk,omitempty"`
	// 酒店关联产品列表
	RelatedPks []*string `json:"RelatedPks,omitempty" xml:"RelatedPks,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	RoomNum *int32 `json:"RoomNum,omitempty" xml:"RoomNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEV***E3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s CreateHotelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelRequest) GoString() string {
	return s.String()
}

func (s *CreateHotelRequest) SetAppKey(v string) *CreateHotelRequest {
	s.AppKey = &v
	return s
}

func (s *CreateHotelRequest) SetEstOpenTime(v string) *CreateHotelRequest {
	s.EstOpenTime = &v
	return s
}

func (s *CreateHotelRequest) SetHotelAddress(v string) *CreateHotelRequest {
	s.HotelAddress = &v
	return s
}

func (s *CreateHotelRequest) SetHotelEmail(v string) *CreateHotelRequest {
	s.HotelEmail = &v
	return s
}

func (s *CreateHotelRequest) SetHotelName(v string) *CreateHotelRequest {
	s.HotelName = &v
	return s
}

func (s *CreateHotelRequest) SetPhoneNumber(v string) *CreateHotelRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateHotelRequest) SetRelatedPk(v string) *CreateHotelRequest {
	s.RelatedPk = &v
	return s
}

func (s *CreateHotelRequest) SetRelatedPks(v []*string) *CreateHotelRequest {
	s.RelatedPks = v
	return s
}

func (s *CreateHotelRequest) SetRemark(v string) *CreateHotelRequest {
	s.Remark = &v
	return s
}

func (s *CreateHotelRequest) SetRoomNum(v int32) *CreateHotelRequest {
	s.RoomNum = &v
	return s
}

func (s *CreateHotelRequest) SetTbOpenId(v string) *CreateHotelRequest {
	s.TbOpenId = &v
	return s
}

type CreateHotelShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 333566791
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-1 00:00:00
	EstOpenTime *string `json:"EstOpenTime,omitempty" xml:"EstOpenTime,omitempty"`
	// This parameter is required.
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test@hotel.com
	HotelEmail *string `json:"HotelEmail,omitempty" xml:"HotelEmail,omitempty"`
	// This parameter is required.
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13xxxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// jTO****Rw
	RelatedPk *string `json:"RelatedPk,omitempty" xml:"RelatedPk,omitempty"`
	// 酒店关联产品列表
	RelatedPksShrink *string `json:"RelatedPks,omitempty" xml:"RelatedPks,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	RoomNum *int32 `json:"RoomNum,omitempty" xml:"RoomNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEV***E3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s CreateHotelShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHotelShrinkRequest) SetAppKey(v string) *CreateHotelShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetEstOpenTime(v string) *CreateHotelShrinkRequest {
	s.EstOpenTime = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetHotelAddress(v string) *CreateHotelShrinkRequest {
	s.HotelAddress = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetHotelEmail(v string) *CreateHotelShrinkRequest {
	s.HotelEmail = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetHotelName(v string) *CreateHotelShrinkRequest {
	s.HotelName = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetPhoneNumber(v string) *CreateHotelShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetRelatedPk(v string) *CreateHotelShrinkRequest {
	s.RelatedPk = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetRelatedPksShrink(v string) *CreateHotelShrinkRequest {
	s.RelatedPksShrink = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetRemark(v string) *CreateHotelShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetRoomNum(v int32) *CreateHotelShrinkRequest {
	s.RoomNum = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetTbOpenId(v string) *CreateHotelShrinkRequest {
	s.TbOpenId = &v
	return s
}

type CreateHotelResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5abfd9***2c38661
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CreateHotelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHotelResponseBody) SetExtentions(v map[string]interface{}) *CreateHotelResponseBody {
	s.Extentions = v
	return s
}

func (s *CreateHotelResponseBody) SetMessage(v string) *CreateHotelResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHotelResponseBody) SetRequestId(v string) *CreateHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHotelResponseBody) SetResult(v string) *CreateHotelResponseBody {
	s.Result = &v
	return s
}

func (s *CreateHotelResponseBody) SetStatusCode(v int32) *CreateHotelResponseBody {
	s.StatusCode = &v
	return s
}

type CreateHotelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHotelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelResponse) GoString() string {
	return s.String()
}

func (s *CreateHotelResponse) SetHeaders(v map[string]*string) *CreateHotelResponse {
	s.Headers = v
	return s
}

func (s *CreateHotelResponse) SetStatusCode(v int32) *CreateHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHotelResponse) SetBody(v *CreateHotelResponseBody) *CreateHotelResponse {
	s.Body = v
	return s
}

type CreateHotelAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateHotelAlarmHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmHeaders) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmHeaders) SetCommonHeaders(v map[string]*string) *CreateHotelAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateHotelAlarmHeaders) SetXAcsAligenieAccessToken(v string) *CreateHotelAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateHotelAlarmHeaders) SetAuthorization(v string) *CreateHotelAlarmHeaders {
	s.Authorization = &v
	return s
}

type CreateHotelAlarmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cf2446fc9d144c85aaee4f9ae20a96e7
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// DOU_YIN
	MusicType *string `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// This parameter is required.
	Rooms []*string `json:"Rooms,omitempty" xml:"Rooms,omitempty" type:"Repeated"`
	// This parameter is required.
	ScheduleInfo *CreateHotelAlarmRequestScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
}

func (s CreateHotelAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmRequest) SetHotelId(v string) *CreateHotelAlarmRequest {
	s.HotelId = &v
	return s
}

func (s *CreateHotelAlarmRequest) SetMusicType(v string) *CreateHotelAlarmRequest {
	s.MusicType = &v
	return s
}

func (s *CreateHotelAlarmRequest) SetRooms(v []*string) *CreateHotelAlarmRequest {
	s.Rooms = v
	return s
}

func (s *CreateHotelAlarmRequest) SetScheduleInfo(v *CreateHotelAlarmRequestScheduleInfo) *CreateHotelAlarmRequest {
	s.ScheduleInfo = v
	return s
}

type CreateHotelAlarmRequestScheduleInfo struct {
	Once *CreateHotelAlarmRequestScheduleInfoOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// ONCE, WEEKLY
	//
	// This parameter is required.
	Type   *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly *CreateHotelAlarmRequestScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s CreateHotelAlarmRequestScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmRequestScheduleInfo) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmRequestScheduleInfo) SetOnce(v *CreateHotelAlarmRequestScheduleInfoOnce) *CreateHotelAlarmRequestScheduleInfo {
	s.Once = v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfo) SetType(v string) *CreateHotelAlarmRequestScheduleInfo {
	s.Type = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfo) SetWeekly(v *CreateHotelAlarmRequestScheduleInfoWeekly) *CreateHotelAlarmRequestScheduleInfo {
	s.Weekly = v
	return s
}

type CreateHotelAlarmRequestScheduleInfoOnce struct {
	// example:
	//
	// 20
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// example:
	//
	// 19
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 30
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 9
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 2022
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s CreateHotelAlarmRequestScheduleInfoOnce) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmRequestScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetDay(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetHour(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetMinute(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetMonth(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetYear(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Year = &v
	return s
}

type CreateHotelAlarmRequestScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 30
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s CreateHotelAlarmRequestScheduleInfoWeekly) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmRequestScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *CreateHotelAlarmRequestScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) SetHour(v int32) *CreateHotelAlarmRequestScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) SetMinute(v int32) *CreateHotelAlarmRequestScheduleInfoWeekly {
	s.Minute = &v
	return s
}

type CreateHotelAlarmShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cf2446fc9d144c85aaee4f9ae20a96e7
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// DOU_YIN
	MusicType *string `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// This parameter is required.
	RoomsShrink *string `json:"Rooms,omitempty" xml:"Rooms,omitempty"`
	// This parameter is required.
	ScheduleInfoShrink *string `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty"`
}

func (s CreateHotelAlarmShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmShrinkRequest) SetHotelId(v string) *CreateHotelAlarmShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *CreateHotelAlarmShrinkRequest) SetMusicType(v string) *CreateHotelAlarmShrinkRequest {
	s.MusicType = &v
	return s
}

func (s *CreateHotelAlarmShrinkRequest) SetRoomsShrink(v string) *CreateHotelAlarmShrinkRequest {
	s.RoomsShrink = &v
	return s
}

func (s *CreateHotelAlarmShrinkRequest) SetScheduleInfoShrink(v string) *CreateHotelAlarmShrinkRequest {
	s.ScheduleInfoShrink = &v
	return s
}

type CreateHotelAlarmResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***86881
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*CreateHotelAlarmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CreateHotelAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmResponseBody) SetExtentions(v map[string]interface{}) *CreateHotelAlarmResponseBody {
	s.Extentions = v
	return s
}

func (s *CreateHotelAlarmResponseBody) SetMessage(v string) *CreateHotelAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHotelAlarmResponseBody) SetRequestId(v string) *CreateHotelAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHotelAlarmResponseBody) SetResult(v []*CreateHotelAlarmResponseBodyResult) *CreateHotelAlarmResponseBody {
	s.Result = v
	return s
}

func (s *CreateHotelAlarmResponseBody) SetStatusCode(v int32) *CreateHotelAlarmResponseBody {
	s.StatusCode = &v
	return s
}

type CreateHotelAlarmResponseBodyResult struct {
	// example:
	//
	// 94
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// example:
	//
	// Pvk***TA==
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	FailMsg      *string `json:"FailMsg,omitempty" xml:"FailMsg,omitempty"`
	// example:
	//
	// 101
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// mg***Qd
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s CreateHotelAlarmResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmResponseBodyResult) SetAlarmId(v int64) *CreateHotelAlarmResponseBodyResult {
	s.AlarmId = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) SetDeviceOpenId(v string) *CreateHotelAlarmResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) SetFailMsg(v string) *CreateHotelAlarmResponseBodyResult {
	s.FailMsg = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) SetRoomNo(v string) *CreateHotelAlarmResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) SetUserOpenId(v string) *CreateHotelAlarmResponseBodyResult {
	s.UserOpenId = &v
	return s
}

type CreateHotelAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHotelAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHotelAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHotelAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmResponse) SetHeaders(v map[string]*string) *CreateHotelAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateHotelAlarmResponse) SetStatusCode(v int32) *CreateHotelAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHotelAlarmResponse) SetBody(v *CreateHotelAlarmResponseBody) *CreateHotelAlarmResponse {
	s.Body = v
	return s
}

type CreateRcuSceneHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateRcuSceneHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateRcuSceneHeaders) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneHeaders) SetCommonHeaders(v map[string]*string) *CreateRcuSceneHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRcuSceneHeaders) SetXAcsAligenieAccessToken(v string) *CreateRcuSceneHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateRcuSceneHeaders) SetAuthorization(v string) *CreateRcuSceneHeaders {
	s.Authorization = &v
	return s
}

type CreateRcuSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yoga
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	SceneRelationExtDTO *CreateRcuSceneRequestSceneRelationExtDTO `json:"SceneRelationExtDTO,omitempty" xml:"SceneRelationExtDTO,omitempty" type:"Struct"`
}

func (s CreateRcuSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRcuSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneRequest) SetHotelId(v string) *CreateRcuSceneRequest {
	s.HotelId = &v
	return s
}

func (s *CreateRcuSceneRequest) SetSceneId(v string) *CreateRcuSceneRequest {
	s.SceneId = &v
	return s
}

func (s *CreateRcuSceneRequest) SetSceneRelationExtDTO(v *CreateRcuSceneRequestSceneRelationExtDTO) *CreateRcuSceneRequest {
	s.SceneRelationExtDTO = v
	return s
}

type CreateRcuSceneRequestSceneRelationExtDTO struct {
	// This parameter is required.
	CorpusList []*string `json:"CorpusList,omitempty" xml:"CorpusList,omitempty" type:"Repeated"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingmoshi/shuimian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRcuSceneRequestSceneRelationExtDTO) String() string {
	return tea.Prettify(s)
}

func (s CreateRcuSceneRequestSceneRelationExtDTO) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) SetCorpusList(v []*string) *CreateRcuSceneRequestSceneRelationExtDTO {
	s.CorpusList = v
	return s
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) SetDescription(v string) *CreateRcuSceneRequestSceneRelationExtDTO {
	s.Description = &v
	return s
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) SetIcon(v string) *CreateRcuSceneRequestSceneRelationExtDTO {
	s.Icon = &v
	return s
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) SetName(v string) *CreateRcuSceneRequestSceneRelationExtDTO {
	s.Name = &v
	return s
}

type CreateRcuSceneShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yoga
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	SceneRelationExtDTOShrink *string `json:"SceneRelationExtDTO,omitempty" xml:"SceneRelationExtDTO,omitempty"`
}

func (s CreateRcuSceneShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRcuSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneShrinkRequest) SetHotelId(v string) *CreateRcuSceneShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *CreateRcuSceneShrinkRequest) SetSceneId(v string) *CreateRcuSceneShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *CreateRcuSceneShrinkRequest) SetSceneRelationExtDTOShrink(v string) *CreateRcuSceneShrinkRequest {
	s.SceneRelationExtDTOShrink = &v
	return s
}

type CreateRcuSceneResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 844BA5CE-E30A-53CB-8A11-DE1F344C846D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CreateRcuSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRcuSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneResponseBody) SetMessage(v string) *CreateRcuSceneResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRcuSceneResponseBody) SetRequestId(v string) *CreateRcuSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRcuSceneResponseBody) SetResult(v bool) *CreateRcuSceneResponseBody {
	s.Result = &v
	return s
}

func (s *CreateRcuSceneResponseBody) SetStatusCode(v int32) *CreateRcuSceneResponseBody {
	s.StatusCode = &v
	return s
}

type CreateRcuSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRcuSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRcuSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRcuSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneResponse) SetHeaders(v map[string]*string) *CreateRcuSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateRcuSceneResponse) SetStatusCode(v int32) *CreateRcuSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRcuSceneResponse) SetBody(v *CreateRcuSceneResponseBody) *CreateRcuSceneResponse {
	s.Body = v
	return s
}

type DeleteCartoonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteCartoonHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCartoonHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCartoonHeaders) SetCommonHeaders(v map[string]*string) *DeleteCartoonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCartoonHeaders) SetXAcsAligenieAccessToken(v string) *DeleteCartoonHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteCartoonHeaders) SetAuthorization(v string) *DeleteCartoonHeaders {
	s.Authorization = &v
	return s
}

type DeleteCartoonRequest struct {
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteCartoonRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCartoonRequest) GoString() string {
	return s.String()
}

func (s *DeleteCartoonRequest) SetHotelId(v string) *DeleteCartoonRequest {
	s.HotelId = &v
	return s
}

type DeleteCartoonResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteCartoonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCartoonResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCartoonResponseBody) SetMessage(v string) *DeleteCartoonResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCartoonResponseBody) SetRequestId(v string) *DeleteCartoonResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCartoonResponseBody) SetResult(v bool) *DeleteCartoonResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCartoonResponseBody) SetStatusCode(v int32) *DeleteCartoonResponseBody {
	s.StatusCode = &v
	return s
}

type DeleteCartoonResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCartoonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCartoonResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCartoonResponse) GoString() string {
	return s.String()
}

func (s *DeleteCartoonResponse) SetHeaders(v map[string]*string) *DeleteCartoonResponse {
	s.Headers = v
	return s
}

func (s *DeleteCartoonResponse) SetStatusCode(v int32) *DeleteCartoonResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCartoonResponse) SetBody(v *DeleteCartoonResponseBody) *DeleteCartoonResponse {
	s.Body = v
	return s
}

type DeleteCustomQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteCustomQAHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomQAHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCustomQAHeaders) SetCommonHeaders(v map[string]*string) *DeleteCustomQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCustomQAHeaders) SetXAcsAligenieAccessToken(v string) *DeleteCustomQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteCustomQAHeaders) SetAuthorization(v string) *DeleteCustomQAHeaders {
	s.Authorization = &v
	return s
}

type DeleteCustomQARequest struct {
	CustomQAIds []*string `json:"CustomQAIds,omitempty" xml:"CustomQAIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteCustomQARequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomQARequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomQARequest) SetCustomQAIds(v []*string) *DeleteCustomQARequest {
	s.CustomQAIds = v
	return s
}

func (s *DeleteCustomQARequest) SetHotelId(v string) *DeleteCustomQARequest {
	s.HotelId = &v
	return s
}

type DeleteCustomQAShrinkRequest struct {
	CustomQAIdsShrink *string `json:"CustomQAIds,omitempty" xml:"CustomQAIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteCustomQAShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomQAShrinkRequest) SetCustomQAIdsShrink(v string) *DeleteCustomQAShrinkRequest {
	s.CustomQAIdsShrink = &v
	return s
}

func (s *DeleteCustomQAShrinkRequest) SetHotelId(v string) *DeleteCustomQAShrinkRequest {
	s.HotelId = &v
	return s
}

type DeleteCustomQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteCustomQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomQAResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomQAResponseBody) SetMessage(v string) *DeleteCustomQAResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomQAResponseBody) SetRequestId(v string) *DeleteCustomQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomQAResponseBody) SetResult(v bool) *DeleteCustomQAResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCustomQAResponseBody) SetStatusCode(v int32) *DeleteCustomQAResponseBody {
	s.StatusCode = &v
	return s
}

type DeleteCustomQAResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomQAResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomQAResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomQAResponse) SetHeaders(v map[string]*string) *DeleteCustomQAResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomQAResponse) SetStatusCode(v int32) *DeleteCustomQAResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomQAResponse) SetBody(v *DeleteCustomQAResponseBody) *DeleteCustomQAResponse {
	s.Body = v
	return s
}

type DeleteHotelAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteHotelAlarmHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelAlarmHeaders) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmHeaders) SetCommonHeaders(v map[string]*string) *DeleteHotelAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteHotelAlarmHeaders) SetXAcsAligenieAccessToken(v string) *DeleteHotelAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteHotelAlarmHeaders) SetAuthorization(v string) *DeleteHotelAlarmHeaders {
	s.Authorization = &v
	return s
}

type DeleteHotelAlarmRequest struct {
	// This parameter is required.
	Alarms []*DeleteHotelAlarmRequestAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteHotelAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelAlarmRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmRequest) SetAlarms(v []*DeleteHotelAlarmRequestAlarms) *DeleteHotelAlarmRequest {
	s.Alarms = v
	return s
}

func (s *DeleteHotelAlarmRequest) SetHotelId(v string) *DeleteHotelAlarmRequest {
	s.HotelId = &v
	return s
}

type DeleteHotelAlarmRequestAlarms struct {
	// This parameter is required.
	//
	// example:
	//
	// 5029
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PvkB***TA==
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	// example:
	//
	// 101
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mgw/k***HQd
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s DeleteHotelAlarmRequestAlarms) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelAlarmRequestAlarms) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmRequestAlarms) SetAlarmId(v int64) *DeleteHotelAlarmRequestAlarms {
	s.AlarmId = &v
	return s
}

func (s *DeleteHotelAlarmRequestAlarms) SetDeviceOpenId(v string) *DeleteHotelAlarmRequestAlarms {
	s.DeviceOpenId = &v
	return s
}

func (s *DeleteHotelAlarmRequestAlarms) SetRoomNo(v string) *DeleteHotelAlarmRequestAlarms {
	s.RoomNo = &v
	return s
}

func (s *DeleteHotelAlarmRequestAlarms) SetUserOpenId(v string) *DeleteHotelAlarmRequestAlarms {
	s.UserOpenId = &v
	return s
}

type DeleteHotelAlarmShrinkRequest struct {
	// This parameter is required.
	AlarmsShrink *string `json:"Alarms,omitempty" xml:"Alarms,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteHotelAlarmShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmShrinkRequest) SetAlarmsShrink(v string) *DeleteHotelAlarmShrinkRequest {
	s.AlarmsShrink = &v
	return s
}

func (s *DeleteHotelAlarmShrinkRequest) SetHotelId(v string) *DeleteHotelAlarmShrinkRequest {
	s.HotelId = &v
	return s
}

type DeleteHotelAlarmResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteHotelAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmResponseBody) SetExtentions(v map[string]interface{}) *DeleteHotelAlarmResponseBody {
	s.Extentions = v
	return s
}

func (s *DeleteHotelAlarmResponseBody) SetMessage(v string) *DeleteHotelAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotelAlarmResponseBody) SetRequestId(v string) *DeleteHotelAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotelAlarmResponseBody) SetResult(v int32) *DeleteHotelAlarmResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteHotelAlarmResponseBody) SetStatusCode(v int32) *DeleteHotelAlarmResponseBody {
	s.StatusCode = &v
	return s
}

type DeleteHotelAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHotelAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHotelAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelAlarmResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmResponse) SetHeaders(v map[string]*string) *DeleteHotelAlarmResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotelAlarmResponse) SetStatusCode(v int32) *DeleteHotelAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotelAlarmResponse) SetBody(v *DeleteHotelAlarmResponseBody) *DeleteHotelAlarmResponse {
	s.Body = v
	return s
}

type DeleteHotelSceneBookItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteHotelSceneBookItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelSceneBookItemHeaders) GoString() string {
	return s.String()
}

func (s *DeleteHotelSceneBookItemHeaders) SetCommonHeaders(v map[string]*string) *DeleteHotelSceneBookItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteHotelSceneBookItemHeaders) SetXAcsAligenieAccessToken(v string) *DeleteHotelSceneBookItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteHotelSceneBookItemHeaders) SetAuthorization(v string) *DeleteHotelSceneBookItemHeaders {
	s.Authorization = &v
	return s
}

type DeleteHotelSceneBookItemRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 11823
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteHotelSceneBookItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelSceneBookItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotelSceneBookItemRequest) SetHotelId(v string) *DeleteHotelSceneBookItemRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteHotelSceneBookItemRequest) SetId(v int64) *DeleteHotelSceneBookItemRequest {
	s.Id = &v
	return s
}

func (s *DeleteHotelSceneBookItemRequest) SetName(v string) *DeleteHotelSceneBookItemRequest {
	s.Name = &v
	return s
}

type DeleteHotelSceneBookItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteHotelSceneBookItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelSceneBookItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotelSceneBookItemResponseBody) SetCode(v int32) *DeleteHotelSceneBookItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponseBody) SetMessage(v string) *DeleteHotelSceneBookItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponseBody) SetRequestId(v string) *DeleteHotelSceneBookItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponseBody) SetResult(v bool) *DeleteHotelSceneBookItemResponseBody {
	s.Result = &v
	return s
}

type DeleteHotelSceneBookItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHotelSceneBookItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHotelSceneBookItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelSceneBookItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotelSceneBookItemResponse) SetHeaders(v map[string]*string) *DeleteHotelSceneBookItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotelSceneBookItemResponse) SetStatusCode(v int32) *DeleteHotelSceneBookItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponse) SetBody(v *DeleteHotelSceneBookItemResponseBody) *DeleteHotelSceneBookItemResponse {
	s.Body = v
	return s
}

type DeleteHotelSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteHotelSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelSettingHeaders) GoString() string {
	return s.String()
}

func (s *DeleteHotelSettingHeaders) SetCommonHeaders(v map[string]*string) *DeleteHotelSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteHotelSettingHeaders) SetXAcsAligenieAccessToken(v string) *DeleteHotelSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteHotelSettingHeaders) SetAuthorization(v string) *DeleteHotelSettingHeaders {
	s.Authorization = &v
	return s
}

type DeleteHotelSettingRequest struct {
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
}

func (s DeleteHotelSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelSettingRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotelSettingRequest) SetHotelId(v string) *DeleteHotelSettingRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteHotelSettingRequest) SetSettingType(v string) *DeleteHotelSettingRequest {
	s.SettingType = &v
	return s
}

type DeleteHotelSettingResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67****BB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteHotelSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotelSettingResponseBody) SetMessage(v string) *DeleteHotelSettingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotelSettingResponseBody) SetRequestId(v string) *DeleteHotelSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotelSettingResponseBody) SetResult(v bool) *DeleteHotelSettingResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteHotelSettingResponseBody) SetStatusCode(v int32) *DeleteHotelSettingResponseBody {
	s.StatusCode = &v
	return s
}

type DeleteHotelSettingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHotelSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHotelSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotelSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotelSettingResponse) SetHeaders(v map[string]*string) *DeleteHotelSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotelSettingResponse) SetStatusCode(v int32) *DeleteHotelSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotelSettingResponse) SetBody(v *DeleteHotelSettingResponseBody) *DeleteHotelSettingResponse {
	s.Body = v
	return s
}

type DeleteMessageTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteMessageTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageTemplateHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMessageTemplateHeaders) SetCommonHeaders(v map[string]*string) *DeleteMessageTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMessageTemplateHeaders) SetXAcsAligenieAccessToken(v string) *DeleteMessageTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteMessageTemplateHeaders) SetAuthorization(v string) *DeleteMessageTemplateHeaders {
	s.Authorization = &v
	return s
}

type DeleteMessageTemplateRequest struct {
	// example:
	//
	// 234
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMessageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageTemplateRequest) SetTemplateId(v int64) *DeleteMessageTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteMessageTemplateResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F7E2****B7C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteMessageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessageTemplateResponseBody) SetMessage(v string) *DeleteMessageTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMessageTemplateResponseBody) SetRequestId(v string) *DeleteMessageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMessageTemplateResponseBody) SetResult(v bool) *DeleteMessageTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteMessageTemplateResponseBody) SetStatusCode(v int32) *DeleteMessageTemplateResponseBody {
	s.StatusCode = &v
	return s
}

type DeleteMessageTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageTemplateResponse) SetHeaders(v map[string]*string) *DeleteMessageTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageTemplateResponse) SetStatusCode(v int32) *DeleteMessageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageTemplateResponse) SetBody(v *DeleteMessageTemplateResponseBody) *DeleteMessageTemplateResponse {
	s.Body = v
	return s
}

type DeleteRcuSceneHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteRcuSceneHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRcuSceneHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRcuSceneHeaders) SetCommonHeaders(v map[string]*string) *DeleteRcuSceneHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRcuSceneHeaders) SetXAcsAligenieAccessToken(v string) *DeleteRcuSceneHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteRcuSceneHeaders) SetAuthorization(v string) *DeleteRcuSceneHeaders {
	s.Authorization = &v
	return s
}

type DeleteRcuSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yoga
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeleteRcuSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRcuSceneRequest) GoString() string {
	return s.String()
}

func (s *DeleteRcuSceneRequest) SetHotelId(v string) *DeleteRcuSceneRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteRcuSceneRequest) SetSceneId(v string) *DeleteRcuSceneRequest {
	s.SceneId = &v
	return s
}

type DeleteRcuSceneResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F61A7B7-409C-525D-AFDB-238A4E88925A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteRcuSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRcuSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRcuSceneResponseBody) SetMessage(v string) *DeleteRcuSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRcuSceneResponseBody) SetRequestId(v string) *DeleteRcuSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRcuSceneResponseBody) SetResult(v bool) *DeleteRcuSceneResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteRcuSceneResponseBody) SetStatusCode(v int32) *DeleteRcuSceneResponseBody {
	s.StatusCode = &v
	return s
}

type DeleteRcuSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRcuSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRcuSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRcuSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteRcuSceneResponse) SetHeaders(v map[string]*string) *DeleteRcuSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteRcuSceneResponse) SetStatusCode(v int32) *DeleteRcuSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRcuSceneResponse) SetBody(v *DeleteRcuSceneResponseBody) *DeleteRcuSceneResponse {
	s.Body = v
	return s
}

type DeviceControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeviceControlHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlHeaders) GoString() string {
	return s.String()
}

func (s *DeviceControlHeaders) SetCommonHeaders(v map[string]*string) *DeviceControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeviceControlHeaders) SetXAcsAligenieAccessToken(v string) *DeviceControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeviceControlHeaders) SetAuthorization(v string) *DeviceControlHeaders {
	s.Authorization = &v
	return s
}

type DeviceControlRequest struct {
	Payload  *DeviceControlRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *DeviceControlRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeviceControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlRequest) SetPayload(v *DeviceControlRequestPayload) *DeviceControlRequest {
	s.Payload = v
	return s
}

func (s *DeviceControlRequest) SetUserInfo(v *DeviceControlRequestUserInfo) *DeviceControlRequest {
	s.UserInfo = v
	return s
}

type DeviceControlRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// aircondition
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open
	Cmd *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// air_condition
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// example:
	//
	// {}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// room
	Location   *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s DeviceControlRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequestPayload) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestPayload) SetCategory(v string) *DeviceControlRequestPayload {
	s.Category = &v
	return s
}

func (s *DeviceControlRequestPayload) SetCmd(v string) *DeviceControlRequestPayload {
	s.Cmd = &v
	return s
}

func (s *DeviceControlRequestPayload) SetDeviceNumber(v string) *DeviceControlRequestPayload {
	s.DeviceNumber = &v
	return s
}

func (s *DeviceControlRequestPayload) SetExtendInfo(v string) *DeviceControlRequestPayload {
	s.ExtendInfo = &v
	return s
}

func (s *DeviceControlRequestPayload) SetLocation(v string) *DeviceControlRequestPayload {
	s.Location = &v
	return s
}

func (s *DeviceControlRequestPayload) SetProperties(v map[string]*string) *DeviceControlRequestPayload {
	s.Properties = v
	return s
}

type DeviceControlRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeviceControlRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestUserInfo) SetEncodeKey(v string) *DeviceControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetEncodeType(v string) *DeviceControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetId(v string) *DeviceControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetIdType(v string) *DeviceControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetOrganizationId(v string) *DeviceControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type DeviceControlShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeviceControlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlShrinkRequest) SetPayloadShrink(v string) *DeviceControlShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeviceControlShrinkRequest) SetUserInfoShrink(v string) *DeviceControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type DeviceControlResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeviceControlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeviceControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceControlResponseBody) SetCode(v int32) *DeviceControlResponseBody {
	s.Code = &v
	return s
}

func (s *DeviceControlResponseBody) SetMessage(v string) *DeviceControlResponseBody {
	s.Message = &v
	return s
}

func (s *DeviceControlResponseBody) SetRequestId(v string) *DeviceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceControlResponseBody) SetResult(v *DeviceControlResponseBodyResult) *DeviceControlResponseBody {
	s.Result = v
	return s
}

type DeviceControlResponseBodyResult struct {
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeviceControlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeviceControlResponseBodyResult) SetStatus(v string) *DeviceControlResponseBodyResult {
	s.Status = &v
	return s
}

type DeviceControlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeviceControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeviceControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlResponse) GoString() string {
	return s.String()
}

func (s *DeviceControlResponse) SetHeaders(v map[string]*string) *DeviceControlResponse {
	s.Headers = v
	return s
}

func (s *DeviceControlResponse) SetStatusCode(v int32) *DeviceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceControlResponse) SetBody(v *DeviceControlResponseBody) *DeviceControlResponse {
	s.Body = v
	return s
}

type GetBasicInfoQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetBasicInfoQAHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoQAHeaders) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQAHeaders) SetCommonHeaders(v map[string]*string) *GetBasicInfoQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBasicInfoQAHeaders) SetXAcsAligenieAccessToken(v string) *GetBasicInfoQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetBasicInfoQAHeaders) SetAuthorization(v string) *GetBasicInfoQAHeaders {
	s.Authorization = &v
	return s
}

type GetBasicInfoQARequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s GetBasicInfoQARequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoQARequest) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQARequest) SetHotelId(v string) *GetBasicInfoQARequest {
	s.HotelId = &v
	return s
}

type GetBasicInfoQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetBasicInfoQAResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetBasicInfoQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoQAResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQAResponseBody) SetMessage(v string) *GetBasicInfoQAResponseBody {
	s.Message = &v
	return s
}

func (s *GetBasicInfoQAResponseBody) SetRequestId(v string) *GetBasicInfoQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicInfoQAResponseBody) SetResult(v *GetBasicInfoQAResponseBodyResult) *GetBasicInfoQAResponseBody {
	s.Result = v
	return s
}

func (s *GetBasicInfoQAResponseBody) SetStatusCode(v int32) *GetBasicInfoQAResponseBody {
	s.StatusCode = &v
	return s
}

type GetBasicInfoQAResponseBodyResult struct {
	// example:
	//
	// 11:11
	CheckInTime *string `json:"CheckInTime,omitempty" xml:"CheckInTime,omitempty"`
	// example:
	//
	// 11:11
	CheckOutTime      *string `json:"CheckOutTime,omitempty" xml:"CheckOutTime,omitempty"`
	HotelAddress      *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	HotelIntroduction *string `json:"HotelIntroduction,omitempty" xml:"HotelIntroduction,omitempty"`
	HotelMember       *string `json:"HotelMember,omitempty" xml:"HotelMember,omitempty"`
	HotelService      *string `json:"HotelService,omitempty" xml:"HotelService,omitempty"`
	ParkingExpenses   *string `json:"ParkingExpenses,omitempty" xml:"ParkingExpenses,omitempty"`
	ParkingPosition   *string `json:"ParkingPosition,omitempty" xml:"ParkingPosition,omitempty"`
	// example:
	//
	// 123***
	PhoneNumber  *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	WifiName     *string `json:"WifiName,omitempty" xml:"WifiName,omitempty"`
	WifiPassword *string `json:"WifiPassword,omitempty" xml:"WifiPassword,omitempty"`
}

func (s GetBasicInfoQAResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoQAResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQAResponseBodyResult) SetCheckInTime(v string) *GetBasicInfoQAResponseBodyResult {
	s.CheckInTime = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetCheckOutTime(v string) *GetBasicInfoQAResponseBodyResult {
	s.CheckOutTime = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetHotelAddress(v string) *GetBasicInfoQAResponseBodyResult {
	s.HotelAddress = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetHotelIntroduction(v string) *GetBasicInfoQAResponseBodyResult {
	s.HotelIntroduction = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetHotelMember(v string) *GetBasicInfoQAResponseBodyResult {
	s.HotelMember = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetHotelService(v string) *GetBasicInfoQAResponseBodyResult {
	s.HotelService = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetParkingExpenses(v string) *GetBasicInfoQAResponseBodyResult {
	s.ParkingExpenses = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetParkingPosition(v string) *GetBasicInfoQAResponseBodyResult {
	s.ParkingPosition = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetPhoneNumber(v string) *GetBasicInfoQAResponseBodyResult {
	s.PhoneNumber = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetWifiName(v string) *GetBasicInfoQAResponseBodyResult {
	s.WifiName = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetWifiPassword(v string) *GetBasicInfoQAResponseBodyResult {
	s.WifiPassword = &v
	return s
}

type GetBasicInfoQAResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicInfoQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicInfoQAResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoQAResponse) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQAResponse) SetHeaders(v map[string]*string) *GetBasicInfoQAResponse {
	s.Headers = v
	return s
}

func (s *GetBasicInfoQAResponse) SetStatusCode(v int32) *GetBasicInfoQAResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicInfoQAResponse) SetBody(v *GetBasicInfoQAResponseBody) *GetBasicInfoQAResponse {
	s.Body = v
	return s
}

type GetCartoonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetCartoonHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCartoonHeaders) GoString() string {
	return s.String()
}

func (s *GetCartoonHeaders) SetCommonHeaders(v map[string]*string) *GetCartoonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCartoonHeaders) SetXAcsAligenieAccessToken(v string) *GetCartoonHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetCartoonHeaders) SetAuthorization(v string) *GetCartoonHeaders {
	s.Authorization = &v
	return s
}

type GetCartoonRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s GetCartoonRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCartoonRequest) GoString() string {
	return s.String()
}

func (s *GetCartoonRequest) SetHotelId(v string) *GetCartoonRequest {
	s.HotelId = &v
	return s
}

type GetCartoonResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetCartoonResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetCartoonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCartoonResponseBody) GoString() string {
	return s.String()
}

func (s *GetCartoonResponseBody) SetMessage(v string) *GetCartoonResponseBody {
	s.Message = &v
	return s
}

func (s *GetCartoonResponseBody) SetRequestId(v string) *GetCartoonResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCartoonResponseBody) SetResult(v *GetCartoonResponseBodyResult) *GetCartoonResponseBody {
	s.Result = v
	return s
}

func (s *GetCartoonResponseBody) SetStatusCode(v int32) *GetCartoonResponseBody {
	s.StatusCode = &v
	return s
}

type GetCartoonResponseBodyResult struct {
	// example:
	//
	// https://ai***.mp4
	StartVideoMd5 *string `json:"StartVideoMd5,omitempty" xml:"StartVideoMd5,omitempty"`
	// example:
	//
	// 40c8***97
	StartVideoUrl *string `json:"StartVideoUrl,omitempty" xml:"StartVideoUrl,omitempty"`
}

func (s GetCartoonResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCartoonResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCartoonResponseBodyResult) SetStartVideoMd5(v string) *GetCartoonResponseBodyResult {
	s.StartVideoMd5 = &v
	return s
}

func (s *GetCartoonResponseBodyResult) SetStartVideoUrl(v string) *GetCartoonResponseBodyResult {
	s.StartVideoUrl = &v
	return s
}

type GetCartoonResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCartoonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCartoonResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCartoonResponse) GoString() string {
	return s.String()
}

func (s *GetCartoonResponse) SetHeaders(v map[string]*string) *GetCartoonResponse {
	s.Headers = v
	return s
}

func (s *GetCartoonResponse) SetStatusCode(v int32) *GetCartoonResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCartoonResponse) SetBody(v *GetCartoonResponseBody) *GetCartoonResponse {
	s.Body = v
	return s
}

type GetHotelContactByGenieDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelContactByGenieDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByGenieDeviceHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceHeaders) SetCommonHeaders(v map[string]*string) *GetHotelContactByGenieDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelContactByGenieDeviceHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelContactByGenieDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelContactByGenieDeviceHeaders) SetAuthorization(v string) *GetHotelContactByGenieDeviceHeaders {
	s.Authorization = &v
	return s
}

type GetHotelContactByGenieDeviceRequest struct {
	DeviceInfo *GetHotelContactByGenieDeviceRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo   *GetHotelContactByGenieDeviceRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelContactByGenieDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByGenieDeviceRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceRequest) SetDeviceInfo(v *GetHotelContactByGenieDeviceRequestDeviceInfo) *GetHotelContactByGenieDeviceRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetHotelContactByGenieDeviceRequest) SetUserInfo(v *GetHotelContactByGenieDeviceRequestUserInfo) *GetHotelContactByGenieDeviceRequest {
	s.UserInfo = v
	return s
}

type GetHotelContactByGenieDeviceRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelContactByGenieDeviceRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByGenieDeviceRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetEncodeKey(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetEncodeType(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetId(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetIdType(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetOrganizationId(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelContactByGenieDeviceRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1***2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelContactByGenieDeviceRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByGenieDeviceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetEncodeKey(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetEncodeType(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetId(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetIdType(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetOrganizationId(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelContactByGenieDeviceShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelContactByGenieDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByGenieDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceShrinkRequest) SetDeviceInfoShrink(v string) *GetHotelContactByGenieDeviceShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetHotelContactByGenieDeviceShrinkRequest) SetUserInfoShrink(v string) *GetHotelContactByGenieDeviceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelContactByGenieDeviceResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelContactByGenieDeviceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelContactByGenieDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByGenieDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceResponseBody) SetMessage(v string) *GetHotelContactByGenieDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBody) SetRequestId(v string) *GetHotelContactByGenieDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBody) SetResult(v *GetHotelContactByGenieDeviceResponseBodyResult) *GetHotelContactByGenieDeviceResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBody) SetStatusCode(v int32) *GetHotelContactByGenieDeviceResponseBody {
	s.StatusCode = &v
	return s
}

type GetHotelContactByGenieDeviceResponseBodyResult struct {
	// example:
	//
	// 1649472283046
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// example:
	//
	// 2022-07-21 20:02:12
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-07-27 14:06:27
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// xxx.icon
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 1
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 101
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// group
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2E57***D45F9
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetHotelContactByGenieDeviceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByGenieDeviceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetExpireAt(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.ExpireAt = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetGmtCreate(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetGmtModified(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetHotelId(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetIcon(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetId(v int64) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetName(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetNumber(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Number = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetStatus(v int32) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetType(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetUuid(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Uuid = &v
	return s
}

type GetHotelContactByGenieDeviceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelContactByGenieDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelContactByGenieDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByGenieDeviceResponse) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceResponse) SetHeaders(v map[string]*string) *GetHotelContactByGenieDeviceResponse {
	s.Headers = v
	return s
}

func (s *GetHotelContactByGenieDeviceResponse) SetStatusCode(v int32) *GetHotelContactByGenieDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponse) SetBody(v *GetHotelContactByGenieDeviceResponseBody) *GetHotelContactByGenieDeviceResponse {
	s.Body = v
	return s
}

type GetHotelContactByNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelContactByNumberHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByNumberHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberHeaders) SetCommonHeaders(v map[string]*string) *GetHotelContactByNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelContactByNumberHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelContactByNumberHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelContactByNumberHeaders) SetAuthorization(v string) *GetHotelContactByNumberHeaders {
	s.Authorization = &v
	return s
}

type GetHotelContactByNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101
	Number   *string                                 `json:"Number,omitempty" xml:"Number,omitempty"`
	UserInfo *GetHotelContactByNumberRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelContactByNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByNumberRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberRequest) SetNumber(v string) *GetHotelContactByNumberRequest {
	s.Number = &v
	return s
}

func (s *GetHotelContactByNumberRequest) SetUserInfo(v *GetHotelContactByNumberRequestUserInfo) *GetHotelContactByNumberRequest {
	s.UserInfo = v
	return s
}

type GetHotelContactByNumberRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelContactByNumberRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByNumberRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberRequestUserInfo) SetEncodeKey(v string) *GetHotelContactByNumberRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) SetEncodeType(v string) *GetHotelContactByNumberRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) SetId(v string) *GetHotelContactByNumberRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) SetIdType(v string) *GetHotelContactByNumberRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) SetOrganizationId(v string) *GetHotelContactByNumberRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelContactByNumberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101
	Number         *string `json:"Number,omitempty" xml:"Number,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelContactByNumberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberShrinkRequest) SetNumber(v string) *GetHotelContactByNumberShrinkRequest {
	s.Number = &v
	return s
}

func (s *GetHotelContactByNumberShrinkRequest) SetUserInfoShrink(v string) *GetHotelContactByNumberShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelContactByNumberResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelContactByNumberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelContactByNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberResponseBody) SetMessage(v string) *GetHotelContactByNumberResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelContactByNumberResponseBody) SetRequestId(v string) *GetHotelContactByNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelContactByNumberResponseBody) SetResult(v *GetHotelContactByNumberResponseBodyResult) *GetHotelContactByNumberResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelContactByNumberResponseBody) SetStatusCode(v int32) *GetHotelContactByNumberResponseBody {
	s.StatusCode = &v
	return s
}

type GetHotelContactByNumberResponseBodyResult struct {
	// example:
	//
	// 1649316479098
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// xxx.icon
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 101
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// group
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2E57***D45F9
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetHotelContactByNumberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByNumberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberResponseBodyResult) SetExpireAt(v string) *GetHotelContactByNumberResponseBodyResult {
	s.ExpireAt = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetHotelId(v string) *GetHotelContactByNumberResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetIcon(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetName(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetNumber(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Number = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetStatus(v int32) *GetHotelContactByNumberResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetType(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetUuid(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Uuid = &v
	return s
}

type GetHotelContactByNumberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelContactByNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelContactByNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactByNumberResponse) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberResponse) SetHeaders(v map[string]*string) *GetHotelContactByNumberResponse {
	s.Headers = v
	return s
}

func (s *GetHotelContactByNumberResponse) SetStatusCode(v int32) *GetHotelContactByNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactByNumberResponse) SetBody(v *GetHotelContactByNumberResponseBody) *GetHotelContactByNumberResponse {
	s.Body = v
	return s
}

type GetHotelContactsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelContactsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactsHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelContactsHeaders) SetCommonHeaders(v map[string]*string) *GetHotelContactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelContactsHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelContactsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelContactsHeaders) SetAuthorization(v string) *GetHotelContactsHeaders {
	s.Authorization = &v
	return s
}

type GetHotelContactsRequest struct {
	UserInfo *GetHotelContactsRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactsRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactsRequest) SetUserInfo(v *GetHotelContactsRequestUserInfo) *GetHotelContactsRequest {
	s.UserInfo = v
	return s
}

type GetHotelContactsRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelContactsRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelContactsRequestUserInfo) SetEncodeKey(v string) *GetHotelContactsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) SetEncodeType(v string) *GetHotelContactsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) SetId(v string) *GetHotelContactsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) SetIdType(v string) *GetHotelContactsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) SetOrganizationId(v string) *GetHotelContactsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelContactsShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelContactsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactsShrinkRequest) SetUserInfoShrink(v string) *GetHotelContactsShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelContactsResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelContactsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelContactsResponseBody) SetMessage(v string) *GetHotelContactsResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelContactsResponseBody) SetRequestId(v string) *GetHotelContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelContactsResponseBody) SetResult(v []*GetHotelContactsResponseBodyResult) *GetHotelContactsResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelContactsResponseBody) SetStatusCode(v int32) *GetHotelContactsResponseBody {
	s.StatusCode = &v
	return s
}

type GetHotelContactsResponseBodyResult struct {
	// example:
	//
	// 1649472283046
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// example:
	//
	// cf24***96e7
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// xxx.icon
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 101
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// group
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0862***A809
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetHotelContactsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelContactsResponseBodyResult) SetExpireAt(v string) *GetHotelContactsResponseBodyResult {
	s.ExpireAt = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetHotelId(v string) *GetHotelContactsResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetIcon(v string) *GetHotelContactsResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetName(v string) *GetHotelContactsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetNumber(v string) *GetHotelContactsResponseBodyResult {
	s.Number = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetStatus(v int32) *GetHotelContactsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetType(v string) *GetHotelContactsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetUuid(v string) *GetHotelContactsResponseBodyResult {
	s.Uuid = &v
	return s
}

type GetHotelContactsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelContactsResponse) GoString() string {
	return s.String()
}

func (s *GetHotelContactsResponse) SetHeaders(v map[string]*string) *GetHotelContactsResponse {
	s.Headers = v
	return s
}

func (s *GetHotelContactsResponse) SetStatusCode(v int32) *GetHotelContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactsResponse) SetBody(v *GetHotelContactsResponseBody) *GetHotelContactsResponse {
	s.Body = v
	return s
}

type GetHotelHomeBackImageAndModesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelHomeBackImageAndModesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetCommonHeaders(v map[string]*string) *GetHotelHomeBackImageAndModesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelHomeBackImageAndModesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetAuthorization(v string) *GetHotelHomeBackImageAndModesHeaders {
	s.Authorization = &v
	return s
}

type GetHotelHomeBackImageAndModesRequest struct {
	// This parameter is required.
	UserInfo *GetHotelHomeBackImageAndModesRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelHomeBackImageAndModesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesRequest) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesRequest) SetUserInfo(v *GetHotelHomeBackImageAndModesRequestUserInfo) *GetHotelHomeBackImageAndModesRequest {
	s.UserInfo = v
	return s
}

type GetHotelHomeBackImageAndModesRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1248494721591392955
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mFU6VtVU+pgA8lx6rYMo7SPl11t+8b+8ALrn10MIPEdpK/HI9wELAEppYhPI1cYRDa4og8AMjAEBZKbLUwFjFA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelHomeBackImageAndModesRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetEncodeKey(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetEncodeType(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetId(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetIdType(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetOrganizationId(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelHomeBackImageAndModesShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelHomeBackImageAndModesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesShrinkRequest) SetUserInfoShrink(v string) *GetHotelHomeBackImageAndModesShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelHomeBackImageAndModesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 394450FC-9035-1B7C-8829-BC88832473FC
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelHomeBackImageAndModesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetHotelHomeBackImageAndModesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetCode(v int32) *GetHotelHomeBackImageAndModesResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetMessage(v string) *GetHotelHomeBackImageAndModesResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetRequestId(v string) *GetHotelHomeBackImageAndModesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetResult(v *GetHotelHomeBackImageAndModesResponseBodyResult) *GetHotelHomeBackImageAndModesResponseBody {
	s.Result = v
	return s
}

type GetHotelHomeBackImageAndModesResponseBodyResult struct {
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/3d4fe6d66ec49d9789635f66627f0339/welcome_audios/976210a6532150f49c2677a8b7dbc105/l6fspbhn.jpg
	BackgroundImage *string `json:"BackgroundImage,omitempty" xml:"BackgroundImage,omitempty"`
	// example:
	//
	// 宣雍测试橙蜂酒店
	HotelName *string                                                    `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	ModeList  []*GetHotelHomeBackImageAndModesResponseBodyResultModeList `json:"ModeList,omitempty" xml:"ModeList,omitempty" type:"Repeated"`
}

func (s GetHotelHomeBackImageAndModesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetBackgroundImage(v string) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.BackgroundImage = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetHotelName(v string) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.HotelName = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetModeList(v []*GetHotelHomeBackImageAndModesResponseBodyResultModeList) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.ModeList = v
	return s
}

type GetHotelHomeBackImageAndModesResponseBodyResultModeList struct {
	// example:
	//
	// 浪漫模式
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// example:
	//
	// romantic
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingmoshi/langman.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
}

func (s GetHotelHomeBackImageAndModesResponseBodyResultModeList) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBodyResultModeList) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetCnName(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.CnName = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetCode(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.Code = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetIcon(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.Icon = &v
	return s
}

type GetHotelHomeBackImageAndModesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelHomeBackImageAndModesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelHomeBackImageAndModesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponse) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponse) SetHeaders(v map[string]*string) *GetHotelHomeBackImageAndModesResponse {
	s.Headers = v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponse) SetStatusCode(v int32) *GetHotelHomeBackImageAndModesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponse) SetBody(v *GetHotelHomeBackImageAndModesResponseBody) *GetHotelHomeBackImageAndModesResponse {
	s.Body = v
	return s
}

type GetHotelNoticeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelNoticeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeHeaders) SetCommonHeaders(v map[string]*string) *GetHotelNoticeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelNoticeHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelNoticeHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelNoticeHeaders) SetAuthorization(v string) *GetHotelNoticeHeaders {
	s.Authorization = &v
	return s
}

type GetHotelNoticeRequest struct {
	// This parameter is required.
	UserInfo *GetHotelNoticeRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeRequest) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeRequest) SetUserInfo(v *GetHotelNoticeRequestUserInfo) *GetHotelNoticeRequest {
	s.UserInfo = v
	return s
}

type GetHotelNoticeRequestUserInfo struct {
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelNoticeRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeRequestUserInfo) SetEncodeKey(v string) *GetHotelNoticeRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) SetEncodeType(v string) *GetHotelNoticeRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) SetId(v string) *GetHotelNoticeRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) SetIdType(v string) *GetHotelNoticeRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) SetOrganizationId(v string) *GetHotelNoticeRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelNoticeShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelNoticeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeShrinkRequest) SetUserInfoShrink(v string) *GetHotelNoticeShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelNoticeResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test notice...
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetHotelNoticeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeResponseBody) SetCode(v int32) *GetHotelNoticeResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelNoticeResponseBody) SetMessage(v string) *GetHotelNoticeResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelNoticeResponseBody) SetRequestId(v string) *GetHotelNoticeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelNoticeResponseBody) SetResult(v string) *GetHotelNoticeResponseBody {
	s.Result = &v
	return s
}

type GetHotelNoticeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeResponse) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeResponse) SetHeaders(v map[string]*string) *GetHotelNoticeResponse {
	s.Headers = v
	return s
}

func (s *GetHotelNoticeResponse) SetStatusCode(v int32) *GetHotelNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelNoticeResponse) SetBody(v *GetHotelNoticeResponseBody) *GetHotelNoticeResponse {
	s.Body = v
	return s
}

type GetHotelNoticeV2Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelNoticeV2Headers) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeV2Headers) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2Headers) SetCommonHeaders(v map[string]*string) *GetHotelNoticeV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelNoticeV2Headers) SetXAcsAligenieAccessToken(v string) *GetHotelNoticeV2Headers {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelNoticeV2Headers) SetAuthorization(v string) *GetHotelNoticeV2Headers {
	s.Authorization = &v
	return s
}

type GetHotelNoticeV2Request struct {
	// This parameter is required.
	UserInfo *GetHotelNoticeV2RequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelNoticeV2Request) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeV2Request) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2Request) SetUserInfo(v *GetHotelNoticeV2RequestUserInfo) *GetHotelNoticeV2Request {
	s.UserInfo = v
	return s
}

type GetHotelNoticeV2RequestUserInfo struct {
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelNoticeV2RequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeV2RequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2RequestUserInfo) SetEncodeKey(v string) *GetHotelNoticeV2RequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) SetEncodeType(v string) *GetHotelNoticeV2RequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) SetId(v string) *GetHotelNoticeV2RequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) SetIdType(v string) *GetHotelNoticeV2RequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) SetOrganizationId(v string) *GetHotelNoticeV2RequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelNoticeV2ShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelNoticeV2ShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2ShrinkRequest) SetUserInfoShrink(v string) *GetHotelNoticeV2ShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelNoticeV2ResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0D0C***67DB
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelNoticeV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelNoticeV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2ResponseBody) SetMessage(v string) *GetHotelNoticeV2ResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBody) SetRequestId(v string) *GetHotelNoticeV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBody) SetResult(v *GetHotelNoticeV2ResponseBodyResult) *GetHotelNoticeV2ResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelNoticeV2ResponseBody) SetStatusCode(v int32) *GetHotelNoticeV2ResponseBody {
	s.StatusCode = &v
	return s
}

type GetHotelNoticeV2ResponseBodyResult struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotelNoticeV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2ResponseBodyResult) SetContent(v string) *GetHotelNoticeV2ResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBodyResult) SetHotelId(v string) *GetHotelNoticeV2ResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBodyResult) SetTitle(v string) *GetHotelNoticeV2ResponseBodyResult {
	s.Title = &v
	return s
}

type GetHotelNoticeV2Response struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelNoticeV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelNoticeV2Response) String() string {
	return tea.Prettify(s)
}

func (s GetHotelNoticeV2Response) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2Response) SetHeaders(v map[string]*string) *GetHotelNoticeV2Response {
	s.Headers = v
	return s
}

func (s *GetHotelNoticeV2Response) SetStatusCode(v int32) *GetHotelNoticeV2Response {
	s.StatusCode = &v
	return s
}

func (s *GetHotelNoticeV2Response) SetBody(v *GetHotelNoticeV2ResponseBody) *GetHotelNoticeV2Response {
	s.Body = v
	return s
}

type GetHotelOrderDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelOrderDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailHeaders) SetCommonHeaders(v map[string]*string) *GetHotelOrderDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelOrderDetailHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelOrderDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelOrderDetailHeaders) SetAuthorization(v string) *GetHotelOrderDetailHeaders {
	s.Authorization = &v
	return s
}

type GetHotelOrderDetailRequest struct {
	// This parameter is required.
	Payload *GetHotelOrderDetailRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
}

func (s GetHotelOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailRequest) SetPayload(v *GetHotelOrderDetailRequestPayload) *GetHotelOrderDetailRequest {
	s.Payload = v
	return s
}

type GetHotelOrderDetailRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220714150702000168270112410630
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
}

func (s GetHotelOrderDetailRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailRequestPayload) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailRequestPayload) SetOrderNo(v string) *GetHotelOrderDetailRequestPayload {
	s.OrderNo = &v
	return s
}

type GetHotelOrderDetailShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s GetHotelOrderDetailShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailShrinkRequest) SetPayloadShrink(v string) *GetHotelOrderDetailShrinkRequest {
	s.PayloadShrink = &v
	return s
}

type GetHotelOrderDetailResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6F579407-13C4-1708-AFA2-B657BE5FE8F5
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelOrderDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetHotelOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponseBody) SetCode(v int32) *GetHotelOrderDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetMessage(v string) *GetHotelOrderDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetRequestId(v string) *GetHotelOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetResult(v []*GetHotelOrderDetailResponseBodyResult) *GetHotelOrderDetailResponseBody {
	s.Result = v
	return s
}

type GetHotelOrderDetailResponseBodyResult struct {
	// example:
	//
	// 200
	ApplyAmt *int64 `json:"ApplyAmt,omitempty" xml:"ApplyAmt,omitempty"`
	// example:
	//
	// 1659952892000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/weixiu/dianqilei/chuanglian.png
	ItemUrl *string `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	// example:
	//
	// 窗帘
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s GetHotelOrderDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponseBodyResult) SetApplyAmt(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.ApplyAmt = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetGmtCreate(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetItemUrl(v string) *GetHotelOrderDetailResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetName(v string) *GetHotelOrderDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetQuantity(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.Quantity = &v
	return s
}

type GetHotelOrderDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponse) SetHeaders(v map[string]*string) *GetHotelOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotelOrderDetailResponse) SetStatusCode(v int32) *GetHotelOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelOrderDetailResponse) SetBody(v *GetHotelOrderDetailResponseBody) *GetHotelOrderDetailResponse {
	s.Body = v
	return s
}

type GetHotelRoomDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelRoomDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelRoomDeviceHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceHeaders) SetCommonHeaders(v map[string]*string) *GetHotelRoomDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelRoomDeviceHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelRoomDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelRoomDeviceHeaders) SetAuthorization(v string) *GetHotelRoomDeviceHeaders {
	s.Authorization = &v
	return s
}

type GetHotelRoomDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s GetHotelRoomDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelRoomDeviceRequest) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceRequest) SetHotelId(v string) *GetHotelRoomDeviceRequest {
	s.HotelId = &v
	return s
}

func (s *GetHotelRoomDeviceRequest) SetRoomNo(v string) *GetHotelRoomDeviceRequest {
	s.RoomNo = &v
	return s
}

type GetHotelRoomDeviceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// vrehvuifdsgrts
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelRoomDeviceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetHotelRoomDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelRoomDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceResponseBody) SetCode(v int32) *GetHotelRoomDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBody) SetMessage(v string) *GetHotelRoomDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBody) SetRequestId(v string) *GetHotelRoomDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBody) SetResult(v []*GetHotelRoomDeviceResponseBodyResult) *GetHotelRoomDeviceResponseBody {
	s.Result = v
	return s
}

type GetHotelRoomDeviceResponseBodyResult struct {
	// example:
	//
	// 1.0.0-release
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// aa:aa:aa:aa:aa:aa
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1
	OnlineStatus *int32 `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// dsfdsfrgreg
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s GetHotelRoomDeviceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelRoomDeviceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetFirmwareVersion(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.FirmwareVersion = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetHotelId(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetMac(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.Mac = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetOnlineStatus(v int32) *GetHotelRoomDeviceResponseBodyResult {
	s.OnlineStatus = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetRoomNo(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetSn(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.Sn = &v
	return s
}

type GetHotelRoomDeviceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelRoomDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelRoomDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelRoomDeviceResponse) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceResponse) SetHeaders(v map[string]*string) *GetHotelRoomDeviceResponse {
	s.Headers = v
	return s
}

func (s *GetHotelRoomDeviceResponse) SetStatusCode(v int32) *GetHotelRoomDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelRoomDeviceResponse) SetBody(v *GetHotelRoomDeviceResponseBody) *GetHotelRoomDeviceResponse {
	s.Body = v
	return s
}

type GetHotelSampleUtterancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelSampleUtterancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesHeaders) SetCommonHeaders(v map[string]*string) *GetHotelSampleUtterancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelSampleUtterancesHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelSampleUtterancesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelSampleUtterancesHeaders) SetAuthorization(v string) *GetHotelSampleUtterancesHeaders {
	s.Authorization = &v
	return s
}

type GetHotelSampleUtterancesRequest struct {
	UserInfo *GetHotelSampleUtterancesRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelSampleUtterancesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesRequest) SetUserInfo(v *GetHotelSampleUtterancesRequestUserInfo) *GetHotelSampleUtterancesRequest {
	s.UserInfo = v
	return s
}

type GetHotelSampleUtterancesRequestUserInfo struct {
	// This parameter is required.
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelSampleUtterancesRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetEncodeKey(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetEncodeType(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetId(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetIdType(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetOrganizationId(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelSampleUtterancesShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelSampleUtterancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesShrinkRequest) SetUserInfoShrink(v string) *GetHotelSampleUtterancesShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelSampleUtterancesResponseBody struct {
	Code      *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetHotelSampleUtterancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesResponseBody) SetCode(v int32) *GetHotelSampleUtterancesResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetMessage(v string) *GetHotelSampleUtterancesResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetRequestId(v string) *GetHotelSampleUtterancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetResult(v []*string) *GetHotelSampleUtterancesResponseBody {
	s.Result = v
	return s
}

type GetHotelSampleUtterancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelSampleUtterancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelSampleUtterancesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesResponse) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesResponse) SetHeaders(v map[string]*string) *GetHotelSampleUtterancesResponse {
	s.Headers = v
	return s
}

func (s *GetHotelSampleUtterancesResponse) SetStatusCode(v int32) *GetHotelSampleUtterancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelSampleUtterancesResponse) SetBody(v *GetHotelSampleUtterancesResponseBody) *GetHotelSampleUtterancesResponse {
	s.Body = v
	return s
}

type GetHotelSceneItemDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelSceneItemDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSceneItemDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailHeaders) SetCommonHeaders(v map[string]*string) *GetHotelSceneItemDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelSceneItemDetailHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelSceneItemDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelSceneItemDetailHeaders) SetAuthorization(v string) *GetHotelSceneItemDetailHeaders {
	s.Authorization = &v
	return s
}

type GetHotelSceneItemDetailRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 10336
	ItemId *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetHotelSceneItemDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSceneItemDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailRequest) SetHotelId(v string) *GetHotelSceneItemDetailRequest {
	s.HotelId = &v
	return s
}

func (s *GetHotelSceneItemDetailRequest) SetItemId(v int64) *GetHotelSceneItemDetailRequest {
	s.ItemId = &v
	return s
}

func (s *GetHotelSceneItemDetailRequest) SetName(v string) *GetHotelSceneItemDetailRequest {
	s.Name = &v
	return s
}

type GetHotelSceneItemDetailResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelSceneItemDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetHotelSceneItemDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSceneItemDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailResponseBody) SetCode(v int32) *GetHotelSceneItemDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBody) SetMessage(v string) *GetHotelSceneItemDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBody) SetRequestId(v string) *GetHotelSceneItemDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBody) SetResult(v *GetHotelSceneItemDetailResponseBodyResult) *GetHotelSceneItemDetailResponseBody {
	s.Result = v
	return s
}

type GetHotelSceneItemDetailResponseBodyResult struct {
	// example:
	//
	// 客用品类
	Category     *string                                                  `json:"Category,omitempty" xml:"Category,omitempty"`
	DialogueList []*GetHotelSceneItemDetailResponseBodyResultDialogueList `json:"DialogueList,omitempty" xml:"DialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/wupin/keyongpinlei/zhijin.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 10336
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 擦鞋布
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 170
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1666168828
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetHotelSceneItemDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSceneItemDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetCategory(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Category = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetDialogueList(v []*GetHotelSceneItemDetailResponseBodyResultDialogueList) *GetHotelSceneItemDetailResponseBodyResult {
	s.DialogueList = v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetIcon(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetId(v int64) *GetHotelSceneItemDetailResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetName(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetPrice(v int64) *GetHotelSceneItemDetailResponseBodyResult {
	s.Price = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetStatus(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetType(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetUpdateTime(v int64) *GetHotelSceneItemDetailResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type GetHotelSceneItemDetailResponseBodyResultDialogueList struct {
	// example:
	//
	// 1666164774
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 336
	DialogueId *string `json:"DialogueId,omitempty" xml:"DialogueId,omitempty"`
	NoAnswer   *string `json:"NoAnswer,omitempty" xml:"NoAnswer,omitempty"`
	// example:
	//
	// 4
	NoAnswerTemplate *string `json:"NoAnswerTemplate,omitempty" xml:"NoAnswerTemplate,omitempty"`
	// example:
	//
	// 0
	Process  *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// 10336
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 1666164774
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	YesAnswer  *string `json:"YesAnswer,omitempty" xml:"YesAnswer,omitempty"`
	// example:
	//
	// 4
	YesAnswerTemplate *string `json:"YesAnswerTemplate,omitempty" xml:"YesAnswerTemplate,omitempty"`
}

func (s GetHotelSceneItemDetailResponseBodyResultDialogueList) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSceneItemDetailResponseBodyResultDialogueList) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetCreateTime(v int64) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.CreateTime = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetDialogueId(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.DialogueId = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetNoAnswer(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.NoAnswer = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetNoAnswerTemplate(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.NoAnswerTemplate = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetProcess(v int32) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.Process = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetQuestion(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.Question = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetServiceId(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.ServiceId = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetUpdateTime(v int64) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.UpdateTime = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetYesAnswer(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.YesAnswer = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetYesAnswerTemplate(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.YesAnswerTemplate = &v
	return s
}

type GetHotelSceneItemDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelSceneItemDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelSceneItemDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSceneItemDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailResponse) SetHeaders(v map[string]*string) *GetHotelSceneItemDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotelSceneItemDetailResponse) SetStatusCode(v int32) *GetHotelSceneItemDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelSceneItemDetailResponse) SetBody(v *GetHotelSceneItemDetailResponseBody) *GetHotelSceneItemDetailResponse {
	s.Body = v
	return s
}

type GetHotelScreenSaverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelScreenSaverHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverHeaders) SetCommonHeaders(v map[string]*string) *GetHotelScreenSaverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelScreenSaverHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelScreenSaverHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelScreenSaverHeaders) SetAuthorization(v string) *GetHotelScreenSaverHeaders {
	s.Authorization = &v
	return s
}

type GetHotelScreenSaverRequest struct {
	// This parameter is required.
	UserInfo *GetHotelScreenSaverRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelScreenSaverRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverRequest) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverRequest) SetUserInfo(v *GetHotelScreenSaverRequestUserInfo) *GetHotelScreenSaverRequest {
	s.UserInfo = v
	return s
}

type GetHotelScreenSaverRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1248494721591392955
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mFU6VtVU+pgA8lx6rYMo7SPl11t+8b+8ALrn10MIPEdpK/HI9wELAEppYhPI1cYRDa4og8AMjAEBZKbLUwFjFA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelScreenSaverRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverRequestUserInfo) SetEncodeKey(v string) *GetHotelScreenSaverRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetEncodeType(v string) *GetHotelScreenSaverRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetId(v string) *GetHotelScreenSaverRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetIdType(v string) *GetHotelScreenSaverRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetOrganizationId(v string) *GetHotelScreenSaverRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelScreenSaverShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelScreenSaverShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverShrinkRequest) SetUserInfoShrink(v string) *GetHotelScreenSaverShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelScreenSaverResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5F0467E1-19F2-1757-B6D0-B79917BA2E81
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelScreenSaverResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetHotelScreenSaverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponseBody) SetCode(v int32) *GetHotelScreenSaverResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetMessage(v string) *GetHotelScreenSaverResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetRequestId(v string) *GetHotelScreenSaverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetResult(v *GetHotelScreenSaverResponseBodyResult) *GetHotelScreenSaverResponseBody {
	s.Result = v
	return s
}

type GetHotelScreenSaverResponseBodyResult struct {
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/3d4fe6d66ec49d9789635f66627f0339/welcome_audios/976210a6532150f49c2677a8b7dbc105/l6fspbhd.jpg
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	// example:
	//
	// common-weather
	StyleCode *string `json:"StyleCode,omitempty" xml:"StyleCode,omitempty"`
}

func (s GetHotelScreenSaverResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponseBodyResult) SetPicUrl(v string) *GetHotelScreenSaverResponseBodyResult {
	s.PicUrl = &v
	return s
}

func (s *GetHotelScreenSaverResponseBodyResult) SetStyleCode(v string) *GetHotelScreenSaverResponseBodyResult {
	s.StyleCode = &v
	return s
}

type GetHotelScreenSaverResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelScreenSaverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelScreenSaverResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverResponse) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponse) SetHeaders(v map[string]*string) *GetHotelScreenSaverResponse {
	s.Headers = v
	return s
}

func (s *GetHotelScreenSaverResponse) SetStatusCode(v int32) *GetHotelScreenSaverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelScreenSaverResponse) SetBody(v *GetHotelScreenSaverResponseBody) *GetHotelScreenSaverResponse {
	s.Body = v
	return s
}

type GetHotelScreenSaverStyleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelScreenSaverStyleHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverStyleHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleHeaders) SetCommonHeaders(v map[string]*string) *GetHotelScreenSaverStyleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelScreenSaverStyleHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelScreenSaverStyleHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelScreenSaverStyleHeaders) SetAuthorization(v string) *GetHotelScreenSaverStyleHeaders {
	s.Authorization = &v
	return s
}

type GetHotelScreenSaverStyleRequest struct {
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s GetHotelScreenSaverStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverStyleRequest) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleRequest) SetHotelId(v string) *GetHotelScreenSaverStyleRequest {
	s.HotelId = &v
	return s
}

type GetHotelScreenSaverStyleResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67**6FA
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelScreenSaverStyleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelScreenSaverStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleResponseBody) SetMessage(v string) *GetHotelScreenSaverStyleResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBody) SetRequestId(v string) *GetHotelScreenSaverStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBody) SetResult(v []*GetHotelScreenSaverStyleResponseBodyResult) *GetHotelScreenSaverStyleResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBody) SetStatusCode(v int32) *GetHotelScreenSaverStyleResponseBody {
	s.StatusCode = &v
	return s
}

type GetHotelScreenSaverStyleResponseBodyResult struct {
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// example:
	//
	// common-weather
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// common-weather
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// example:
	//
	// https://img.***.png
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s GetHotelScreenSaverStyleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverStyleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) SetCnName(v string) *GetHotelScreenSaverStyleResponseBodyResult {
	s.CnName = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) SetCode(v string) *GetHotelScreenSaverStyleResponseBodyResult {
	s.Code = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) SetEnName(v string) *GetHotelScreenSaverStyleResponseBodyResult {
	s.EnName = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) SetPicUrl(v string) *GetHotelScreenSaverStyleResponseBodyResult {
	s.PicUrl = &v
	return s
}

type GetHotelScreenSaverStyleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelScreenSaverStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelScreenSaverStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverStyleResponse) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleResponse) SetHeaders(v map[string]*string) *GetHotelScreenSaverStyleResponse {
	s.Headers = v
	return s
}

func (s *GetHotelScreenSaverStyleResponse) SetStatusCode(v int32) *GetHotelScreenSaverStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponse) SetBody(v *GetHotelScreenSaverStyleResponseBody) *GetHotelScreenSaverStyleResponse {
	s.Body = v
	return s
}

type GetHotelSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelSettingHeaders) SetCommonHeaders(v map[string]*string) *GetHotelSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelSettingHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelSettingHeaders) SetAuthorization(v string) *GetHotelSettingHeaders {
	s.Authorization = &v
	return s
}

type GetHotelSettingRequest struct {
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
}

func (s GetHotelSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSettingRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSettingRequest) SetHotelId(v string) *GetHotelSettingRequest {
	s.HotelId = &v
	return s
}

func (s *GetHotelSettingRequest) SetSettingType(v string) *GetHotelSettingRequest {
	s.SettingType = &v
	return s
}

type GetHotelSettingResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F7E2****B7C94
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponseBody) SetMessage(v string) *GetHotelSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelSettingResponseBody) SetRequestId(v string) *GetHotelSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelSettingResponseBody) SetResult(v *GetHotelSettingResponseBodyResult) *GetHotelSettingResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelSettingResponseBody) SetStatusCode(v int32) *GetHotelSettingResponseBody {
	s.StatusCode = &v
	return s
}

type GetHotelSettingResponseBodyResult struct {
	// example:
	//
	// 0
	DeleteToken *int64 `json:"DeleteToken,omitempty" xml:"DeleteToken,omitempty"`
	// example:
	//
	// {}
	ExtInfo             *string   `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	HotelDeviceModeList []*string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty" type:"Repeated"`
	// example:
	//
	// af7***536
	HotelId          *string                                            `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelScreenSaver *GetHotelSettingResponseBodyResultHotelScreenSaver `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty" type:"Struct"`
	NightMode        *GetHotelSettingResponseBodyResultNightMode        `json:"NightMode,omitempty" xml:"NightMode,omitempty" type:"Struct"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetHotelSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponseBodyResult) SetDeleteToken(v int64) *GetHotelSettingResponseBodyResult {
	s.DeleteToken = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetExtInfo(v string) *GetHotelSettingResponseBodyResult {
	s.ExtInfo = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetHotelDeviceModeList(v []*string) *GetHotelSettingResponseBodyResult {
	s.HotelDeviceModeList = v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetHotelId(v string) *GetHotelSettingResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetHotelScreenSaver(v *GetHotelSettingResponseBodyResultHotelScreenSaver) *GetHotelSettingResponseBodyResult {
	s.HotelScreenSaver = v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetNightMode(v *GetHotelSettingResponseBodyResultNightMode) *GetHotelSettingResponseBodyResult {
	s.NightMode = v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetSettingType(v string) *GetHotelSettingResponseBodyResult {
	s.SettingType = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetValue(v string) *GetHotelSettingResponseBodyResult {
	s.Value = &v
	return s
}

type GetHotelSettingResponseBodyResultHotelScreenSaver struct {
	// example:
	//
	// https://a***png
	ScreenSaverPicUrl *string `json:"ScreenSaverPicUrl,omitempty" xml:"ScreenSaverPicUrl,omitempty"`
	// example:
	//
	// common-weather
	ScreenSaverStyle *string `json:"ScreenSaverStyle,omitempty" xml:"ScreenSaverStyle,omitempty"`
}

func (s GetHotelSettingResponseBodyResultHotelScreenSaver) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSettingResponseBodyResultHotelScreenSaver) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponseBodyResultHotelScreenSaver) SetScreenSaverPicUrl(v string) *GetHotelSettingResponseBodyResultHotelScreenSaver {
	s.ScreenSaverPicUrl = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultHotelScreenSaver) SetScreenSaverStyle(v string) *GetHotelSettingResponseBodyResultHotelScreenSaver {
	s.ScreenSaverStyle = &v
	return s
}

type GetHotelSettingResponseBodyResultNightMode struct {
	// 夜间模式下的默认亮度
	DefaultBright *string `json:"DefaultBright,omitempty" xml:"DefaultBright,omitempty"`
	// 夜间模式下的默认音量
	DefaultVolume *string `json:"DefaultVolume,omitempty" xml:"DefaultVolume,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 22:00
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// screenoff
	StandbyAction *string `json:"StandbyAction,omitempty" xml:"StandbyAction,omitempty"`
	// example:
	//
	// 07:00
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetHotelSettingResponseBodyResultNightMode) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSettingResponseBodyResultNightMode) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetDefaultBright(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.DefaultBright = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetDefaultVolume(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.DefaultVolume = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetEnable(v bool) *GetHotelSettingResponseBodyResultNightMode {
	s.Enable = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetEnd(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.End = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetStandbyAction(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.StandbyAction = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetStart(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.Start = &v
	return s
}

type GetHotelSettingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSettingResponse) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponse) SetHeaders(v map[string]*string) *GetHotelSettingResponse {
	s.Headers = v
	return s
}

func (s *GetHotelSettingResponse) SetStatusCode(v int32) *GetHotelSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelSettingResponse) SetBody(v *GetHotelSettingResponseBody) *GetHotelSettingResponse {
	s.Body = v
	return s
}

type GetRelationProductListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetRelationProductListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRelationProductListHeaders) GoString() string {
	return s.String()
}

func (s *GetRelationProductListHeaders) SetCommonHeaders(v map[string]*string) *GetRelationProductListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelationProductListHeaders) SetXAcsAligenieAccessToken(v string) *GetRelationProductListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetRelationProductListHeaders) SetAuthorization(v string) *GetRelationProductListHeaders {
	s.Authorization = &v
	return s
}

type GetRelationProductListResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetRelationProductListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetRelationProductListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRelationProductListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelationProductListResponseBody) SetExtentions(v map[string]interface{}) *GetRelationProductListResponseBody {
	s.Extentions = v
	return s
}

func (s *GetRelationProductListResponseBody) SetMessage(v string) *GetRelationProductListResponseBody {
	s.Message = &v
	return s
}

func (s *GetRelationProductListResponseBody) SetRequestId(v string) *GetRelationProductListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRelationProductListResponseBody) SetResult(v []*GetRelationProductListResponseBodyResult) *GetRelationProductListResponseBody {
	s.Result = v
	return s
}

func (s *GetRelationProductListResponseBody) SetStatusCode(v int32) *GetRelationProductListResponseBody {
	s.StatusCode = &v
	return s
}

type GetRelationProductListResponseBodyResult struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// jTOSl***l1odxImRw
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetRelationProductListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRelationProductListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRelationProductListResponseBodyResult) SetName(v string) *GetRelationProductListResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetRelationProductListResponseBodyResult) SetPk(v string) *GetRelationProductListResponseBodyResult {
	s.Pk = &v
	return s
}

type GetRelationProductListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRelationProductListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRelationProductListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRelationProductListResponse) GoString() string {
	return s.String()
}

func (s *GetRelationProductListResponse) SetHeaders(v map[string]*string) *GetRelationProductListResponse {
	s.Headers = v
	return s
}

func (s *GetRelationProductListResponse) SetStatusCode(v int32) *GetRelationProductListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRelationProductListResponse) SetBody(v *GetRelationProductListResponseBody) *GetRelationProductListResponse {
	s.Body = v
	return s
}

type GetUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUnionIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUnionIdHeaders) SetXAcsAligenieAccessToken(v string) *GetUnionIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetUnionIdHeaders) SetAuthorization(v string) *GetUnionIdHeaders {
	s.Authorization = &v
	return s
}

type GetUnionIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 125****0946
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 62a319****abdc
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEVICE_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
}

func (s GetUnionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetUnionIdRequest) SetEncodeKey(v string) *GetUnionIdRequest {
	s.EncodeKey = &v
	return s
}

func (s *GetUnionIdRequest) SetEncodeType(v string) *GetUnionIdRequest {
	s.EncodeType = &v
	return s
}

func (s *GetUnionIdRequest) SetId(v string) *GetUnionIdRequest {
	s.Id = &v
	return s
}

func (s *GetUnionIdRequest) SetIdType(v string) *GetUnionIdRequest {
	s.IdType = &v
	return s
}

type GetUnionIdResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetUnionIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetUnionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnionIdResponseBody) SetMessage(v string) *GetUnionIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetUnionIdResponseBody) SetRequestId(v string) *GetUnionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUnionIdResponseBody) SetResult(v []*GetUnionIdResponseBodyResult) *GetUnionIdResponseBody {
	s.Result = v
	return s
}

func (s *GetUnionIdResponseBody) SetStatusCode(v int32) *GetUnionIdResponseBody {
	s.StatusCode = &v
	return s
}

type GetUnionIdResponseBodyResult struct {
	// example:
	//
	// 4325***765
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// 8bh2****8s8
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetUnionIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUnionIdResponseBodyResult) SetOrganizationId(v string) *GetUnionIdResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *GetUnionIdResponseBodyResult) SetUnionId(v string) *GetUnionIdResponseBodyResult {
	s.UnionId = &v
	return s
}

type GetUnionIdResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUnionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetUnionIdResponse) SetHeaders(v map[string]*string) *GetUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetUnionIdResponse) SetStatusCode(v int32) *GetUnionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnionIdResponse) SetBody(v *GetUnionIdResponseBody) *GetUnionIdResponse {
	s.Body = v
	return s
}

type GetWelcomeTextAndMusicHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetWelcomeTextAndMusicHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWelcomeTextAndMusicHeaders) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicHeaders) SetCommonHeaders(v map[string]*string) *GetWelcomeTextAndMusicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWelcomeTextAndMusicHeaders) SetXAcsAligenieAccessToken(v string) *GetWelcomeTextAndMusicHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetWelcomeTextAndMusicHeaders) SetAuthorization(v string) *GetWelcomeTextAndMusicHeaders {
	s.Authorization = &v
	return s
}

type GetWelcomeTextAndMusicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s GetWelcomeTextAndMusicRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWelcomeTextAndMusicRequest) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicRequest) SetHotelId(v string) *GetWelcomeTextAndMusicRequest {
	s.HotelId = &v
	return s
}

type GetWelcomeTextAndMusicResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetWelcomeTextAndMusicResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetWelcomeTextAndMusicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWelcomeTextAndMusicResponseBody) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicResponseBody) SetExtentions(v map[string]interface{}) *GetWelcomeTextAndMusicResponseBody {
	s.Extentions = v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) SetMessage(v string) *GetWelcomeTextAndMusicResponseBody {
	s.Message = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) SetRequestId(v string) *GetWelcomeTextAndMusicResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) SetResult(v *GetWelcomeTextAndMusicResponseBodyResult) *GetWelcomeTextAndMusicResponseBody {
	s.Result = v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) SetStatusCode(v int32) *GetWelcomeTextAndMusicResponseBody {
	s.StatusCode = &v
	return s
}

type GetWelcomeTextAndMusicResponseBodyResult struct {
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// http://ailabsaicloudservice.alicdn.com/tmp/a.wav
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	Text     *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetWelcomeTextAndMusicResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetWelcomeTextAndMusicResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) SetHotelId(v string) *GetWelcomeTextAndMusicResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) SetMusicUrl(v string) *GetWelcomeTextAndMusicResponseBodyResult {
	s.MusicUrl = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) SetText(v string) *GetWelcomeTextAndMusicResponseBodyResult {
	s.Text = &v
	return s
}

type GetWelcomeTextAndMusicResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWelcomeTextAndMusicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWelcomeTextAndMusicResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWelcomeTextAndMusicResponse) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicResponse) SetHeaders(v map[string]*string) *GetWelcomeTextAndMusicResponse {
	s.Headers = v
	return s
}

func (s *GetWelcomeTextAndMusicResponse) SetStatusCode(v int32) *GetWelcomeTextAndMusicResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponse) SetBody(v *GetWelcomeTextAndMusicResponseBody) *GetWelcomeTextAndMusicResponse {
	s.Body = v
	return s
}

type HotelQrBindHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s HotelQrBindHeaders) String() string {
	return tea.Prettify(s)
}

func (s HotelQrBindHeaders) GoString() string {
	return s.String()
}

func (s *HotelQrBindHeaders) SetCommonHeaders(v map[string]*string) *HotelQrBindHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelQrBindHeaders) SetXAcsAligenieAccessToken(v string) *HotelQrBindHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *HotelQrBindHeaders) SetAuthorization(v string) *HotelQrBindHeaders {
	s.Authorization = &v
	return s
}

type HotelQrBindRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// freuisghrtiesnvfkdsvbfuidslnvfs
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s HotelQrBindRequest) String() string {
	return tea.Prettify(s)
}

func (s HotelQrBindRequest) GoString() string {
	return s.String()
}

func (s *HotelQrBindRequest) SetClientId(v string) *HotelQrBindRequest {
	s.ClientId = &v
	return s
}

func (s *HotelQrBindRequest) SetCode(v string) *HotelQrBindRequest {
	s.Code = &v
	return s
}

func (s *HotelQrBindRequest) SetExtInfo(v string) *HotelQrBindRequest {
	s.ExtInfo = &v
	return s
}

func (s *HotelQrBindRequest) SetHotelId(v string) *HotelQrBindRequest {
	s.HotelId = &v
	return s
}

func (s *HotelQrBindRequest) SetRoomNo(v string) *HotelQrBindRequest {
	s.RoomNo = &v
	return s
}

type HotelQrBindResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73****9-175A-1324-8202-9FAAB*****A
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *HotelQrBindResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s HotelQrBindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HotelQrBindResponseBody) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponseBody) SetMessage(v string) *HotelQrBindResponseBody {
	s.Message = &v
	return s
}

func (s *HotelQrBindResponseBody) SetRequestId(v string) *HotelQrBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelQrBindResponseBody) SetResult(v *HotelQrBindResponseBodyResult) *HotelQrBindResponseBody {
	s.Result = v
	return s
}

func (s *HotelQrBindResponseBody) SetStatusCode(v int32) *HotelQrBindResponseBody {
	s.StatusCode = &v
	return s
}

type HotelQrBindResponseBodyResult struct {
	OpenDeviceInfo *HotelQrBindResponseBodyResultOpenDeviceInfo `json:"OpenDeviceInfo,omitempty" xml:"OpenDeviceInfo,omitempty" type:"Struct"`
	OpenUserInfo   *HotelQrBindResponseBodyResultOpenUserInfo   `json:"OpenUserInfo,omitempty" xml:"OpenUserInfo,omitempty" type:"Struct"`
}

func (s HotelQrBindResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s HotelQrBindResponseBodyResult) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponseBodyResult) SetOpenDeviceInfo(v *HotelQrBindResponseBodyResultOpenDeviceInfo) *HotelQrBindResponseBodyResult {
	s.OpenDeviceInfo = v
	return s
}

func (s *HotelQrBindResponseBodyResult) SetOpenUserInfo(v *HotelQrBindResponseBodyResultOpenUserInfo) *HotelQrBindResponseBodyResult {
	s.OpenUserInfo = v
	return s
}

type HotelQrBindResponseBodyResultOpenDeviceInfo struct {
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// xxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// aaaaaaaa
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s HotelQrBindResponseBodyResultOpenDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelQrBindResponseBodyResultOpenDeviceInfo) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetEncodeKey(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetEncodeType(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetId(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.Id = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetIdType(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.IdType = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetOrganizationId(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.OrganizationId = &v
	return s
}

type HotelQrBindResponseBodyResultOpenUserInfo struct {
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// aaaaaaaa
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s HotelQrBindResponseBodyResultOpenUserInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelQrBindResponseBodyResultOpenUserInfo) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetEncodeKey(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetEncodeType(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.EncodeType = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetId(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.Id = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetIdType(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.IdType = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetOrganizationId(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.OrganizationId = &v
	return s
}

type HotelQrBindResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelQrBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelQrBindResponse) String() string {
	return tea.Prettify(s)
}

func (s HotelQrBindResponse) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponse) SetHeaders(v map[string]*string) *HotelQrBindResponse {
	s.Headers = v
	return s
}

func (s *HotelQrBindResponse) SetStatusCode(v int32) *HotelQrBindResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelQrBindResponse) SetBody(v *HotelQrBindResponseBody) *HotelQrBindResponse {
	s.Body = v
	return s
}

type ImportHotelConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ImportHotelConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s ImportHotelConfigHeaders) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigHeaders) SetCommonHeaders(v map[string]*string) *ImportHotelConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ImportHotelConfigHeaders) SetXAcsAligenieAccessToken(v string) *ImportHotelConfigHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ImportHotelConfigHeaders) SetAuthorization(v string) *ImportHotelConfigHeaders {
	s.Authorization = &v
	return s
}

type ImportHotelConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	ImportHotelConfig *ImportHotelConfigRequestImportHotelConfig `json:"ImportHotelConfig,omitempty" xml:"ImportHotelConfig,omitempty" type:"Struct"`
}

func (s ImportHotelConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportHotelConfigRequest) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigRequest) SetHotelId(v string) *ImportHotelConfigRequest {
	s.HotelId = &v
	return s
}

func (s *ImportHotelConfigRequest) SetImportHotelConfig(v *ImportHotelConfigRequestImportHotelConfig) *ImportHotelConfigRequest {
	s.ImportHotelConfig = v
	return s
}

type ImportHotelConfigRequestImportHotelConfig struct {
	RcuCustomScenes []*ImportHotelConfigRequestImportHotelConfigRcuCustomScenes `json:"RcuCustomScenes,omitempty" xml:"RcuCustomScenes,omitempty" type:"Repeated"`
}

func (s ImportHotelConfigRequestImportHotelConfig) String() string {
	return tea.Prettify(s)
}

func (s ImportHotelConfigRequestImportHotelConfig) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigRequestImportHotelConfig) SetRcuCustomScenes(v []*ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) *ImportHotelConfigRequestImportHotelConfig {
	s.RcuCustomScenes = v
	return s
}

type ImportHotelConfigRequestImportHotelConfigRcuCustomScenes struct {
	// This parameter is required.
	CorpusList  []*string `json:"CorpusList,omitempty" xml:"CorpusList,omitempty" type:"Repeated"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *string   `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) String() string {
	return tea.Prettify(s)
}

func (s ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetCorpusList(v []*string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.CorpusList = v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetDescription(v string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.Description = &v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetIcon(v string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.Icon = &v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetName(v string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.Name = &v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetSceneId(v string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.SceneId = &v
	return s
}

type ImportHotelConfigShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	ImportHotelConfigShrink *string `json:"ImportHotelConfig,omitempty" xml:"ImportHotelConfig,omitempty"`
}

func (s ImportHotelConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportHotelConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigShrinkRequest) SetHotelId(v string) *ImportHotelConfigShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ImportHotelConfigShrinkRequest) SetImportHotelConfigShrink(v string) *ImportHotelConfigShrinkRequest {
	s.ImportHotelConfigShrink = &v
	return s
}

type ImportHotelConfigResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ImportHotelConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportHotelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigResponseBody) SetMessage(v string) *ImportHotelConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ImportHotelConfigResponseBody) SetRequestId(v string) *ImportHotelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportHotelConfigResponseBody) SetResult(v bool) *ImportHotelConfigResponseBody {
	s.Result = &v
	return s
}

func (s *ImportHotelConfigResponseBody) SetStatusCode(v int32) *ImportHotelConfigResponseBody {
	s.StatusCode = &v
	return s
}

type ImportHotelConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportHotelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportHotelConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportHotelConfigResponse) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigResponse) SetHeaders(v map[string]*string) *ImportHotelConfigResponse {
	s.Headers = v
	return s
}

func (s *ImportHotelConfigResponse) SetStatusCode(v int32) *ImportHotelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportHotelConfigResponse) SetBody(v *ImportHotelConfigResponseBody) *ImportHotelConfigResponse {
	s.Body = v
	return s
}

type ImportRoomControlDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ImportRoomControlDevicesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesHeaders) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesHeaders) SetCommonHeaders(v map[string]*string) *ImportRoomControlDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ImportRoomControlDevicesHeaders) SetXAcsAligenieAccessToken(v string) *ImportRoomControlDevicesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ImportRoomControlDevicesHeaders) SetAuthorization(v string) *ImportRoomControlDevicesHeaders {
	s.Authorization = &v
	return s
}

type ImportRoomControlDevicesRequest struct {
	EnableInfraredDeviceImport *string `json:"EnableInfraredDeviceImport,omitempty" xml:"EnableInfraredDeviceImport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vdgrefds
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	LocationDevices []*ImportRoomControlDevicesRequestLocationDevices `json:"LocationDevices,omitempty" xml:"LocationDevices,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ImportRoomControlDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesRequest) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequest) SetEnableInfraredDeviceImport(v string) *ImportRoomControlDevicesRequest {
	s.EnableInfraredDeviceImport = &v
	return s
}

func (s *ImportRoomControlDevicesRequest) SetHotelId(v string) *ImportRoomControlDevicesRequest {
	s.HotelId = &v
	return s
}

func (s *ImportRoomControlDevicesRequest) SetLocationDevices(v []*ImportRoomControlDevicesRequestLocationDevices) *ImportRoomControlDevicesRequest {
	s.LocationDevices = v
	return s
}

func (s *ImportRoomControlDevicesRequest) SetRoomNo(v string) *ImportRoomControlDevicesRequest {
	s.RoomNo = &v
	return s
}

type ImportRoomControlDevicesRequestLocationDevices struct {
	Devices []*ImportRoomControlDevicesRequestLocationDevicesDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// room
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
}

func (s ImportRoomControlDevicesRequestLocationDevices) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesRequestLocationDevices) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequestLocationDevices) SetDevices(v []*ImportRoomControlDevicesRequestLocationDevicesDevices) *ImportRoomControlDevicesRequestLocationDevices {
	s.Devices = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevices) SetLocation(v string) *ImportRoomControlDevicesRequestLocationDevices {
	s.Location = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevices) SetLocationName(v string) *ImportRoomControlDevicesRequestLocationDevices {
	s.LocationName = &v
	return s
}

type ImportRoomControlDevicesRequestLocationDevicesDevices struct {
	AliasList   []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	Brand       *string   `json:"Brand,omitempty" xml:"Brand,omitempty"`
	City        *string   `json:"City,omitempty" xml:"City,omitempty"`
	ConnectType *string   `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// This parameter is required.
	DeviceName        *string                                                                 `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	Dn                *string                                                                 `json:"Dn,omitempty" xml:"Dn,omitempty"`
	InfraredId        *string                                                                 `json:"InfraredId,omitempty" xml:"InfraredId,omitempty"`
	InfraredIndex     *string                                                                 `json:"InfraredIndex,omitempty" xml:"InfraredIndex,omitempty"`
	InfraredVersion   *string                                                                 `json:"InfraredVersion,omitempty" xml:"InfraredVersion,omitempty"`
	MultiKeySwitchExt *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt `json:"MultiKeySwitchExt,omitempty" xml:"MultiKeySwitchExt,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// light
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// night_light
	Number          *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Pk              *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Province        *string `json:"Province,omitempty" xml:"Province,omitempty"`
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevices) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevices) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetAliasList(v []*string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.AliasList = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetBrand(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Brand = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetCity(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.City = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetConnectType(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.ConnectType = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetDeviceName(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.DeviceName = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetDn(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Dn = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetInfraredId(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.InfraredId = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetInfraredIndex(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.InfraredIndex = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetInfraredVersion(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.InfraredVersion = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetMultiKeySwitchExt(v *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.MultiKeySwitchExt = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetName(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Name = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetNumber(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Number = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetPk(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Pk = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetProvince(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Province = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetServiceProvider(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.ServiceProvider = &v
	return s
}

type ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt struct {
	SwitchList []*ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList `json:"SwitchList,omitempty" xml:"SwitchList,omitempty" type:"Repeated"`
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) SetSwitchList(v []*ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt {
	s.SwitchList = v
	return s
}

type ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList struct {
	AliasList   []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	Category    *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeviceIndex *int32    `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	DeviceName  *string   `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	Location    *string   `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetAliasList(v []*string) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.AliasList = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetCategory(v string) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.Category = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetDeviceIndex(v int32) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.DeviceIndex = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetDeviceName(v string) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.DeviceName = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetLocation(v string) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.Location = &v
	return s
}

type ImportRoomControlDevicesShrinkRequest struct {
	EnableInfraredDeviceImport *string `json:"EnableInfraredDeviceImport,omitempty" xml:"EnableInfraredDeviceImport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vdgrefds
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	LocationDevicesShrink *string `json:"LocationDevices,omitempty" xml:"LocationDevices,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ImportRoomControlDevicesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesShrinkRequest) SetEnableInfraredDeviceImport(v string) *ImportRoomControlDevicesShrinkRequest {
	s.EnableInfraredDeviceImport = &v
	return s
}

func (s *ImportRoomControlDevicesShrinkRequest) SetHotelId(v string) *ImportRoomControlDevicesShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ImportRoomControlDevicesShrinkRequest) SetLocationDevicesShrink(v string) *ImportRoomControlDevicesShrinkRequest {
	s.LocationDevicesShrink = &v
	return s
}

func (s *ImportRoomControlDevicesShrinkRequest) SetRoomNo(v string) *ImportRoomControlDevicesShrinkRequest {
	s.RoomNo = &v
	return s
}

type ImportRoomControlDevicesResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// fdsfregtre
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ImportRoomControlDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesResponseBody) SetExtentions(v map[string]interface{}) *ImportRoomControlDevicesResponseBody {
	s.Extentions = v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) SetMessage(v string) *ImportRoomControlDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) SetRequestId(v string) *ImportRoomControlDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) SetResult(v int32) *ImportRoomControlDevicesResponseBody {
	s.Result = &v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) SetStatusCode(v int32) *ImportRoomControlDevicesResponseBody {
	s.StatusCode = &v
	return s
}

type ImportRoomControlDevicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportRoomControlDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportRoomControlDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomControlDevicesResponse) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesResponse) SetHeaders(v map[string]*string) *ImportRoomControlDevicesResponse {
	s.Headers = v
	return s
}

func (s *ImportRoomControlDevicesResponse) SetStatusCode(v int32) *ImportRoomControlDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportRoomControlDevicesResponse) SetBody(v *ImportRoomControlDevicesResponseBody) *ImportRoomControlDevicesResponse {
	s.Body = v
	return s
}

type ImportRoomGenieScenesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ImportRoomGenieScenesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesHeaders) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesHeaders) SetCommonHeaders(v map[string]*string) *ImportRoomGenieScenesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ImportRoomGenieScenesHeaders) SetXAcsAligenieAccessToken(v string) *ImportRoomGenieScenesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ImportRoomGenieScenesHeaders) SetAuthorization(v string) *ImportRoomGenieScenesHeaders {
	s.Authorization = &v
	return s
}

type ImportRoomGenieScenesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo    *string                                  `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	SceneList []*ImportRoomGenieScenesRequestSceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
}

func (s ImportRoomGenieScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesRequest) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequest) SetHotelId(v string) *ImportRoomGenieScenesRequest {
	s.HotelId = &v
	return s
}

func (s *ImportRoomGenieScenesRequest) SetRoomNo(v string) *ImportRoomGenieScenesRequest {
	s.RoomNo = &v
	return s
}

func (s *ImportRoomGenieScenesRequest) SetSceneList(v []*ImportRoomGenieScenesRequestSceneList) *ImportRoomGenieScenesRequest {
	s.SceneList = v
	return s
}

type ImportRoomGenieScenesRequestSceneList struct {
	// This parameter is required.
	Actions     []*ImportRoomGenieScenesRequestSceneListActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Description *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Display *bool `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// http://xxx.com/yyy.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TriggerLogical *int32 `json:"TriggerLogical,omitempty" xml:"TriggerLogical,omitempty"`
	// This parameter is required.
	Triggers []*ImportRoomGenieScenesRequestSceneListTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ImportRoomGenieScenesRequestSceneList) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneList) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneList) SetActions(v []*ImportRoomGenieScenesRequestSceneListActions) *ImportRoomGenieScenesRequestSceneList {
	s.Actions = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetDescription(v string) *ImportRoomGenieScenesRequestSceneList {
	s.Description = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetDisplay(v bool) *ImportRoomGenieScenesRequestSceneList {
	s.Display = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetIcon(v string) *ImportRoomGenieScenesRequestSceneList {
	s.Icon = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetSceneName(v string) *ImportRoomGenieScenesRequestSceneList {
	s.SceneName = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetTriggerLogical(v int32) *ImportRoomGenieScenesRequestSceneList {
	s.TriggerLogical = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetTriggers(v []*ImportRoomGenieScenesRequestSceneListTriggers) *ImportRoomGenieScenesRequestSceneList {
	s.Triggers = v
	return s
}

type ImportRoomGenieScenesRequestSceneListActions struct {
	AttributeValues []*ImportRoomGenieScenesRequestSceneListActionsAttributeValues `json:"AttributeValues,omitempty" xml:"AttributeValues,omitempty" type:"Repeated"`
	Device          *ImportRoomGenieScenesRequestSceneListActionsDevice            `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	Reply           *string                                                        `json:"Reply,omitempty" xml:"Reply,omitempty"`
	Type            *int32                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListActions) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListActions) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListActions) SetAttributeValues(v []*ImportRoomGenieScenesRequestSceneListActionsAttributeValues) *ImportRoomGenieScenesRequestSceneListActions {
	s.AttributeValues = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActions) SetDevice(v *ImportRoomGenieScenesRequestSceneListActionsDevice) *ImportRoomGenieScenesRequestSceneListActions {
	s.Device = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActions) SetReply(v string) *ImportRoomGenieScenesRequestSceneListActions {
	s.Reply = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActions) SetType(v int32) *ImportRoomGenieScenesRequestSceneListActions {
	s.Type = &v
	return s
}

type ImportRoomGenieScenesRequestSceneListActionsAttributeValues struct {
	// This parameter is required.
	//
	// example:
	//
	// powerstate
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListActionsAttributeValues) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListActionsAttributeValues) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListActionsAttributeValues) SetAttributeName(v string) *ImportRoomGenieScenesRequestSceneListActionsAttributeValues {
	s.AttributeName = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsAttributeValues) SetAttributeValue(v string) *ImportRoomGenieScenesRequestSceneListActionsAttributeValues {
	s.AttributeValue = &v
	return s
}

type ImportRoomGenieScenesRequestSceneListActionsDevice struct {
	// This parameter is required.
	//
	// example:
	//
	// light
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 0
	DeviceIndex *int32 `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3c5d2ab8f9ec
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListActionsDevice) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListActionsDevice) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) SetCategory(v string) *ImportRoomGenieScenesRequestSceneListActionsDevice {
	s.Category = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) SetDeviceIndex(v int32) *ImportRoomGenieScenesRequestSceneListActionsDevice {
	s.DeviceIndex = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) SetDeviceNumber(v string) *ImportRoomGenieScenesRequestSceneListActionsDevice {
	s.DeviceNumber = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) SetType(v int32) *ImportRoomGenieScenesRequestSceneListActionsDevice {
	s.Type = &v
	return s
}

type ImportRoomGenieScenesRequestSceneListTriggers struct {
	AttributeValues []*ImportRoomGenieScenesRequestSceneListTriggersAttributeValues `json:"AttributeValues,omitempty" xml:"AttributeValues,omitempty" type:"Repeated"`
	CorpusList      []*string                                                       `json:"CorpusList,omitempty" xml:"CorpusList,omitempty" type:"Repeated"`
	Device          *ImportRoomGenieScenesRequestSceneListTriggersDevice            `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListTriggers) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListTriggers) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) SetAttributeValues(v []*ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) *ImportRoomGenieScenesRequestSceneListTriggers {
	s.AttributeValues = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) SetCorpusList(v []*string) *ImportRoomGenieScenesRequestSceneListTriggers {
	s.CorpusList = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) SetDevice(v *ImportRoomGenieScenesRequestSceneListTriggersDevice) *ImportRoomGenieScenesRequestSceneListTriggers {
	s.Device = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) SetType(v int32) *ImportRoomGenieScenesRequestSceneListTriggers {
	s.Type = &v
	return s
}

type ImportRoomGenieScenesRequestSceneListTriggersAttributeValues struct {
	// This parameter is required.
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// This parameter is required.
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) SetAttributeName(v string) *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues {
	s.AttributeName = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) SetAttributeValue(v string) *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues {
	s.AttributeValue = &v
	return s
}

type ImportRoomGenieScenesRequestSceneListTriggersDevice struct {
	// This parameter is required.
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeviceIndex *string `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// This parameter is required.
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListTriggersDevice) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListTriggersDevice) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) SetCategory(v string) *ImportRoomGenieScenesRequestSceneListTriggersDevice {
	s.Category = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) SetDeviceIndex(v string) *ImportRoomGenieScenesRequestSceneListTriggersDevice {
	s.DeviceIndex = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) SetDeviceNumber(v string) *ImportRoomGenieScenesRequestSceneListTriggersDevice {
	s.DeviceNumber = &v
	return s
}

type ImportRoomGenieScenesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo          *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	SceneListShrink *string `json:"SceneList,omitempty" xml:"SceneList,omitempty"`
}

func (s ImportRoomGenieScenesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesShrinkRequest) SetHotelId(v string) *ImportRoomGenieScenesShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ImportRoomGenieScenesShrinkRequest) SetRoomNo(v string) *ImportRoomGenieScenesShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *ImportRoomGenieScenesShrinkRequest) SetSceneListShrink(v string) *ImportRoomGenieScenesShrinkRequest {
	s.SceneListShrink = &v
	return s
}

type ImportRoomGenieScenesResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ImportRoomGenieScenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesResponseBody) SetMessage(v string) *ImportRoomGenieScenesResponseBody {
	s.Message = &v
	return s
}

func (s *ImportRoomGenieScenesResponseBody) SetRequestId(v string) *ImportRoomGenieScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportRoomGenieScenesResponseBody) SetResult(v bool) *ImportRoomGenieScenesResponseBody {
	s.Result = &v
	return s
}

func (s *ImportRoomGenieScenesResponseBody) SetStatusCode(v int32) *ImportRoomGenieScenesResponseBody {
	s.StatusCode = &v
	return s
}

type ImportRoomGenieScenesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportRoomGenieScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportRoomGenieScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportRoomGenieScenesResponse) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesResponse) SetHeaders(v map[string]*string) *ImportRoomGenieScenesResponse {
	s.Headers = v
	return s
}

func (s *ImportRoomGenieScenesResponse) SetStatusCode(v int32) *ImportRoomGenieScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportRoomGenieScenesResponse) SetBody(v *ImportRoomGenieScenesResponseBody) *ImportRoomGenieScenesResponse {
	s.Body = v
	return s
}

type InsertHotelSceneBookItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s InsertHotelSceneBookItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertHotelSceneBookItemHeaders) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemHeaders) SetCommonHeaders(v map[string]*string) *InsertHotelSceneBookItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertHotelSceneBookItemHeaders) SetXAcsAligenieAccessToken(v string) *InsertHotelSceneBookItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *InsertHotelSceneBookItemHeaders) SetAuthorization(v string) *InsertHotelSceneBookItemHeaders {
	s.Authorization = &v
	return s
}

type InsertHotelSceneBookItemRequest struct {
	// addHotelSceneItemReq
	//
	// This parameter is required.
	AddHotelSceneItemReq *InsertHotelSceneBookItemRequestAddHotelSceneItemReq `json:"AddHotelSceneItemReq,omitempty" xml:"AddHotelSceneItemReq,omitempty" type:"Struct"`
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s InsertHotelSceneBookItemRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertHotelSceneBookItemRequest) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemRequest) SetAddHotelSceneItemReq(v *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) *InsertHotelSceneBookItemRequest {
	s.AddHotelSceneItemReq = v
	return s
}

func (s *InsertHotelSceneBookItemRequest) SetHotelId(v string) *InsertHotelSceneBookItemRequest {
	s.HotelId = &v
	return s
}

type InsertHotelSceneBookItemRequestAddHotelSceneItemReq struct {
	// icon
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/28d7a91e3c05db3855725fc39e0387e7/welcome_audios/aa918294b6ca3aa115c51135bf9b80cb/l9f996sq.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 青椒肉丝
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1250
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FOOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s InsertHotelSceneBookItemRequestAddHotelSceneItemReq) String() string {
	return tea.Prettify(s)
}

func (s InsertHotelSceneBookItemRequestAddHotelSceneItemReq) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) SetIcon(v string) *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	s.Icon = &v
	return s
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) SetName(v string) *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	s.Name = &v
	return s
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) SetPrice(v int64) *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	s.Price = &v
	return s
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) SetType(v string) *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	s.Type = &v
	return s
}

type InsertHotelSceneBookItemShrinkRequest struct {
	// addHotelSceneItemReq
	//
	// This parameter is required.
	AddHotelSceneItemReqShrink *string `json:"AddHotelSceneItemReq,omitempty" xml:"AddHotelSceneItemReq,omitempty"`
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s InsertHotelSceneBookItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertHotelSceneBookItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemShrinkRequest) SetAddHotelSceneItemReqShrink(v string) *InsertHotelSceneBookItemShrinkRequest {
	s.AddHotelSceneItemReqShrink = &v
	return s
}

func (s *InsertHotelSceneBookItemShrinkRequest) SetHotelId(v string) *InsertHotelSceneBookItemShrinkRequest {
	s.HotelId = &v
	return s
}

type InsertHotelSceneBookItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 36FB***80C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InsertHotelSceneBookItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertHotelSceneBookItemResponseBody) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemResponseBody) SetCode(v int32) *InsertHotelSceneBookItemResponseBody {
	s.Code = &v
	return s
}

func (s *InsertHotelSceneBookItemResponseBody) SetMessage(v string) *InsertHotelSceneBookItemResponseBody {
	s.Message = &v
	return s
}

func (s *InsertHotelSceneBookItemResponseBody) SetRequestId(v string) *InsertHotelSceneBookItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertHotelSceneBookItemResponseBody) SetResult(v bool) *InsertHotelSceneBookItemResponseBody {
	s.Result = &v
	return s
}

type InsertHotelSceneBookItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertHotelSceneBookItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertHotelSceneBookItemResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertHotelSceneBookItemResponse) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemResponse) SetHeaders(v map[string]*string) *InsertHotelSceneBookItemResponse {
	s.Headers = v
	return s
}

func (s *InsertHotelSceneBookItemResponse) SetStatusCode(v int32) *InsertHotelSceneBookItemResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertHotelSceneBookItemResponse) SetBody(v *InsertHotelSceneBookItemResponseBody) *InsertHotelSceneBookItemResponse {
	s.Body = v
	return s
}

type InvokeRobotPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s InvokeRobotPushHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvokeRobotPushHeaders) GoString() string {
	return s.String()
}

func (s *InvokeRobotPushHeaders) SetCommonHeaders(v map[string]*string) *InvokeRobotPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeRobotPushHeaders) SetXAcsAligenieAccessToken(v string) *InvokeRobotPushHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *InvokeRobotPushHeaders) SetAuthorization(v string) *InvokeRobotPushHeaders {
	s.Authorization = &v
	return s
}

type InvokeRobotPushRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GET
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s InvokeRobotPushRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeRobotPushRequest) GoString() string {
	return s.String()
}

func (s *InvokeRobotPushRequest) SetHotelId(v string) *InvokeRobotPushRequest {
	s.HotelId = &v
	return s
}

func (s *InvokeRobotPushRequest) SetPushType(v string) *InvokeRobotPushRequest {
	s.PushType = &v
	return s
}

func (s *InvokeRobotPushRequest) SetRoomNo(v string) *InvokeRobotPushRequest {
	s.RoomNo = &v
	return s
}

type InvokeRobotPushResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s InvokeRobotPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeRobotPushResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeRobotPushResponseBody) SetMessage(v string) *InvokeRobotPushResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeRobotPushResponseBody) SetRequestId(v string) *InvokeRobotPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeRobotPushResponseBody) SetResult(v bool) *InvokeRobotPushResponseBody {
	s.Result = &v
	return s
}

func (s *InvokeRobotPushResponseBody) SetStatusCode(v int32) *InvokeRobotPushResponseBody {
	s.StatusCode = &v
	return s
}

type InvokeRobotPushResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeRobotPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeRobotPushResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeRobotPushResponse) GoString() string {
	return s.String()
}

func (s *InvokeRobotPushResponse) SetHeaders(v map[string]*string) *InvokeRobotPushResponse {
	s.Headers = v
	return s
}

func (s *InvokeRobotPushResponse) SetStatusCode(v int32) *InvokeRobotPushResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeRobotPushResponse) SetBody(v *InvokeRobotPushResponseBody) *InvokeRobotPushResponse {
	s.Body = v
	return s
}

type ListAllProvincesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAllProvincesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAllProvincesHeaders) GoString() string {
	return s.String()
}

func (s *ListAllProvincesHeaders) SetCommonHeaders(v map[string]*string) *ListAllProvincesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAllProvincesHeaders) SetXAcsAligenieAccessToken(v string) *ListAllProvincesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListAllProvincesHeaders) SetAuthorization(v string) *ListAllProvincesHeaders {
	s.Authorization = &v
	return s
}

type ListAllProvincesResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 00534880-4397-5134-B212-1030B7A37C27
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListAllProvincesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllProvincesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllProvincesResponseBody) SetMessage(v string) *ListAllProvincesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllProvincesResponseBody) SetRequestId(v string) *ListAllProvincesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllProvincesResponseBody) SetResult(v []*string) *ListAllProvincesResponseBody {
	s.Result = v
	return s
}

func (s *ListAllProvincesResponseBody) SetStatusCode(v int32) *ListAllProvincesResponseBody {
	s.StatusCode = &v
	return s
}

type ListAllProvincesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllProvincesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllProvincesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllProvincesResponse) GoString() string {
	return s.String()
}

func (s *ListAllProvincesResponse) SetHeaders(v map[string]*string) *ListAllProvincesResponse {
	s.Headers = v
	return s
}

func (s *ListAllProvincesResponse) SetStatusCode(v int32) *ListAllProvincesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllProvincesResponse) SetBody(v *ListAllProvincesResponseBody) *ListAllProvincesResponse {
	s.Body = v
	return s
}

type ListCitiesByProvinceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCitiesByProvinceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCitiesByProvinceHeaders) GoString() string {
	return s.String()
}

func (s *ListCitiesByProvinceHeaders) SetCommonHeaders(v map[string]*string) *ListCitiesByProvinceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCitiesByProvinceHeaders) SetXAcsAligenieAccessToken(v string) *ListCitiesByProvinceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCitiesByProvinceHeaders) SetAuthorization(v string) *ListCitiesByProvinceHeaders {
	s.Authorization = &v
	return s
}

type ListCitiesByProvinceRequest struct {
	// This parameter is required.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListCitiesByProvinceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCitiesByProvinceRequest) GoString() string {
	return s.String()
}

func (s *ListCitiesByProvinceRequest) SetProvince(v string) *ListCitiesByProvinceRequest {
	s.Province = &v
	return s
}

type ListCitiesByProvinceResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 860194F7-9593-50EA-8E53-BCEC0D325A00
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListCitiesByProvinceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCitiesByProvinceResponseBody) GoString() string {
	return s.String()
}

func (s *ListCitiesByProvinceResponseBody) SetMessage(v string) *ListCitiesByProvinceResponseBody {
	s.Message = &v
	return s
}

func (s *ListCitiesByProvinceResponseBody) SetRequestId(v string) *ListCitiesByProvinceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCitiesByProvinceResponseBody) SetResult(v []*string) *ListCitiesByProvinceResponseBody {
	s.Result = v
	return s
}

func (s *ListCitiesByProvinceResponseBody) SetStatusCode(v int32) *ListCitiesByProvinceResponseBody {
	s.StatusCode = &v
	return s
}

type ListCitiesByProvinceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCitiesByProvinceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCitiesByProvinceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCitiesByProvinceResponse) GoString() string {
	return s.String()
}

func (s *ListCitiesByProvinceResponse) SetHeaders(v map[string]*string) *ListCitiesByProvinceResponse {
	s.Headers = v
	return s
}

func (s *ListCitiesByProvinceResponse) SetStatusCode(v int32) *ListCitiesByProvinceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCitiesByProvinceResponse) SetBody(v *ListCitiesByProvinceResponseBody) *ListCitiesByProvinceResponse {
	s.Body = v
	return s
}

type ListCustomQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCustomQAHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCustomQAHeaders) GoString() string {
	return s.String()
}

func (s *ListCustomQAHeaders) SetCommonHeaders(v map[string]*string) *ListCustomQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCustomQAHeaders) SetXAcsAligenieAccessToken(v string) *ListCustomQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCustomQAHeaders) SetAuthorization(v string) *ListCustomQAHeaders {
	s.Authorization = &v
	return s
}

type ListCustomQARequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// ***
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	Page *ListCustomQARequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListCustomQARequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomQARequest) GoString() string {
	return s.String()
}

func (s *ListCustomQARequest) SetHotelId(v string) *ListCustomQARequest {
	s.HotelId = &v
	return s
}

func (s *ListCustomQARequest) SetKeyword(v string) *ListCustomQARequest {
	s.Keyword = &v
	return s
}

func (s *ListCustomQARequest) SetPage(v *ListCustomQARequestPage) *ListCustomQARequest {
	s.Page = v
	return s
}

type ListCustomQARequestPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCustomQARequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListCustomQARequestPage) GoString() string {
	return s.String()
}

func (s *ListCustomQARequestPage) SetPageNumber(v int32) *ListCustomQARequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListCustomQARequestPage) SetPageSize(v int32) *ListCustomQARequestPage {
	s.PageSize = &v
	return s
}

type ListCustomQAShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// ***
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s ListCustomQAShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCustomQAShrinkRequest) SetHotelId(v string) *ListCustomQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListCustomQAShrinkRequest) SetKeyword(v string) *ListCustomQAShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListCustomQAShrinkRequest) SetPageShrink(v string) *ListCustomQAShrinkRequest {
	s.PageShrink = &v
	return s
}

type ListCustomQAResponseBody struct {
	// example:
	//
	// success
	Message *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListCustomQAResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCustomQAResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListCustomQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomQAResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomQAResponseBody) SetMessage(v string) *ListCustomQAResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomQAResponseBody) SetPage(v *ListCustomQAResponseBodyPage) *ListCustomQAResponseBody {
	s.Page = v
	return s
}

func (s *ListCustomQAResponseBody) SetRequestId(v string) *ListCustomQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomQAResponseBody) SetResult(v []*ListCustomQAResponseBodyResult) *ListCustomQAResponseBody {
	s.Result = v
	return s
}

func (s *ListCustomQAResponseBody) SetStatusCode(v int32) *ListCustomQAResponseBody {
	s.StatusCode = &v
	return s
}

type ListCustomQAResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 21
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCustomQAResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListCustomQAResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListCustomQAResponseBodyPage) SetPageNumber(v int32) *ListCustomQAResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListCustomQAResponseBodyPage) SetPageSize(v int32) *ListCustomQAResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListCustomQAResponseBodyPage) SetTotal(v int32) *ListCustomQAResponseBodyPage {
	s.Total = &v
	return s
}

type ListCustomQAResponseBodyResult struct {
	// example:
	//
	// 22;11
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// example:
	//
	// 2023-01-10 10:01:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 111
	CustomQAId *string `json:"CustomQAId,omitempty" xml:"CustomQAId,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 22;11
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// ***
	MajorQuestion *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 22;11
	SupplementaryQuestion *string `json:"SupplementaryQuestion,omitempty" xml:"SupplementaryQuestion,omitempty"`
	// example:
	//
	// 2023-01-10 10:01:59
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCustomQAResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCustomQAResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCustomQAResponseBodyResult) SetAnswers(v string) *ListCustomQAResponseBodyResult {
	s.Answers = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetCreateTime(v string) *ListCustomQAResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetCustomQAId(v string) *ListCustomQAResponseBodyResult {
	s.CustomQAId = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetHotelId(v string) *ListCustomQAResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetKeyWords(v string) *ListCustomQAResponseBodyResult {
	s.KeyWords = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetMajorQuestion(v string) *ListCustomQAResponseBodyResult {
	s.MajorQuestion = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetStatus(v int32) *ListCustomQAResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetSupplementaryQuestion(v string) *ListCustomQAResponseBodyResult {
	s.SupplementaryQuestion = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetUpdateTime(v string) *ListCustomQAResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type ListCustomQAResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomQAResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomQAResponse) GoString() string {
	return s.String()
}

func (s *ListCustomQAResponse) SetHeaders(v map[string]*string) *ListCustomQAResponse {
	s.Headers = v
	return s
}

func (s *ListCustomQAResponse) SetStatusCode(v int32) *ListCustomQAResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomQAResponse) SetBody(v *ListCustomQAResponseBody) *ListCustomQAResponse {
	s.Body = v
	return s
}

type ListDialogueTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDialogueTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDialogueTemplateHeaders) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateHeaders) SetCommonHeaders(v map[string]*string) *ListDialogueTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDialogueTemplateHeaders) SetXAcsAligenieAccessToken(v string) *ListDialogueTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDialogueTemplateHeaders) SetAuthorization(v string) *ListDialogueTemplateHeaders {
	s.Authorization = &v
	return s
}

type ListDialogueTemplateRequest struct {
	// hotelId
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ListDialogueTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDialogueTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateRequest) SetHotelId(v string) *ListDialogueTemplateRequest {
	s.HotelId = &v
	return s
}

type ListDialogueTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 0EC7*726E
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDialogueTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDialogueTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDialogueTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBody) SetCode(v int32) *ListDialogueTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListDialogueTemplateResponseBody) SetMessage(v string) *ListDialogueTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListDialogueTemplateResponseBody) SetRequestId(v string) *ListDialogueTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDialogueTemplateResponseBody) SetResult(v []*ListDialogueTemplateResponseBodyResult) *ListDialogueTemplateResponseBody {
	s.Result = v
	return s
}

type ListDialogueTemplateResponseBodyResult struct {
	TemplateDetail *ListDialogueTemplateResponseBodyResultTemplateDetail `json:"TemplateDetail,omitempty" xml:"TemplateDetail,omitempty" type:"Struct"`
	// example:
	//
	// 4
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 物品多轮模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDialogueTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDialogueTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBodyResult) SetTemplateDetail(v *ListDialogueTemplateResponseBodyResultTemplateDetail) *ListDialogueTemplateResponseBodyResult {
	s.TemplateDetail = v
	return s
}

func (s *ListDialogueTemplateResponseBodyResult) SetTemplateId(v int64) *ListDialogueTemplateResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResult) SetTemplateName(v string) *ListDialogueTemplateResponseBodyResult {
	s.TemplateName = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResult) SetType(v string) *ListDialogueTemplateResponseBodyResult {
	s.Type = &v
	return s
}

type ListDialogueTemplateResponseBodyResultTemplateDetail struct {
	FirstDialogueTemplate  *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate  `json:"FirstDialogueTemplate,omitempty" xml:"FirstDialogueTemplate,omitempty" type:"Struct"`
	SecondDialogueTemplate *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate `json:"SecondDialogueTemplate,omitempty" xml:"SecondDialogueTemplate,omitempty" type:"Struct"`
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetail) String() string {
	return tea.Prettify(s)
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetail) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetail) SetFirstDialogueTemplate(v *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) *ListDialogueTemplateResponseBodyResultTemplateDetail {
	s.FirstDialogueTemplate = v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetail) SetSecondDialogueTemplate(v *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) *ListDialogueTemplateResponseBodyResultTemplateDetail {
	s.SecondDialogueTemplate = v
	return s
}

type ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate struct {
	// example:
	//
	// ${goodsName}${price}元，请问需要服务员送来吗？
	NonzeroPriceYesAnswer *string `json:"NonzeroPriceYesAnswer,omitempty" xml:"NonzeroPriceYesAnswer,omitempty"`
	// example:
	//
	// 对不起，暂时不提供此物品。
	ZeroPriceNoAnswer *string `json:"ZeroPriceNoAnswer,omitempty" xml:"ZeroPriceNoAnswer,omitempty"`
	// example:
	//
	// 好的，服务员会尽快送来。
	ZeroPriceYesAnswer *string `json:"ZeroPriceYesAnswer,omitempty" xml:"ZeroPriceYesAnswer,omitempty"`
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) SetNonzeroPriceYesAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate {
	s.NonzeroPriceYesAnswer = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) SetZeroPriceNoAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate {
	s.ZeroPriceNoAnswer = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) SetZeroPriceYesAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate {
	s.ZeroPriceYesAnswer = &v
	return s
}

type ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate struct {
	// example:
	//
	// 好的，已取消。
	NonzeroPriceNoAnswer *string `json:"NonzeroPriceNoAnswer,omitempty" xml:"NonzeroPriceNoAnswer,omitempty"`
	// example:
	//
	// 好的，服务员会尽快送来${goodsName}
	NonzeroPriceYesAnswer *string `json:"NonzeroPriceYesAnswer,omitempty" xml:"NonzeroPriceYesAnswer,omitempty"`
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) SetNonzeroPriceNoAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate {
	s.NonzeroPriceNoAnswer = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) SetNonzeroPriceYesAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate {
	s.NonzeroPriceYesAnswer = &v
	return s
}

type ListDialogueTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDialogueTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDialogueTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDialogueTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponse) SetHeaders(v map[string]*string) *ListDialogueTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListDialogueTemplateResponse) SetStatusCode(v int32) *ListDialogueTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDialogueTemplateResponse) SetBody(v *ListDialogueTemplateResponseBody) *ListDialogueTemplateResponse {
	s.Body = v
	return s
}

type ListHotelAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelAlarmHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmHeaders) SetCommonHeaders(v map[string]*string) *ListHotelAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelAlarmHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelAlarmHeaders) SetAuthorization(v string) *ListHotelAlarmHeaders {
	s.Authorization = &v
	return s
}

type ListHotelAlarmRequest struct {
	// example:
	//
	// a7a3***013
	HotelId *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	Rooms   []*string `json:"Rooms,omitempty" xml:"Rooms,omitempty" type:"Repeated"`
}

func (s ListHotelAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmRequest) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmRequest) SetHotelId(v string) *ListHotelAlarmRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelAlarmRequest) SetRooms(v []*string) *ListHotelAlarmRequest {
	s.Rooms = v
	return s
}

type ListHotelAlarmShrinkRequest struct {
	// example:
	//
	// a7a3***013
	HotelId     *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	RoomsShrink *string `json:"Rooms,omitempty" xml:"Rooms,omitempty"`
}

func (s ListHotelAlarmShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmShrinkRequest) SetHotelId(v string) *ListHotelAlarmShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelAlarmShrinkRequest) SetRoomsShrink(v string) *ListHotelAlarmShrinkRequest {
	s.RoomsShrink = &v
	return s
}

type ListHotelAlarmResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***881
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelAlarmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListHotelAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBody) SetExtentions(v map[string]interface{}) *ListHotelAlarmResponseBody {
	s.Extentions = v
	return s
}

func (s *ListHotelAlarmResponseBody) SetMessage(v string) *ListHotelAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelAlarmResponseBody) SetRequestId(v string) *ListHotelAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelAlarmResponseBody) SetResult(v []*ListHotelAlarmResponseBodyResult) *ListHotelAlarmResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelAlarmResponseBody) SetStatusCode(v int32) *ListHotelAlarmResponseBody {
	s.StatusCode = &v
	return s
}

type ListHotelAlarmResponseBodyResult struct {
	// example:
	//
	// 5039
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// example:
	//
	// PvkB****VVTA==
	DeviceOpenId *string                                       `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	ScheduleInfo *ListHotelAlarmResponseBodyResultScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	// example:
	//
	// mgw/k***HQd
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s ListHotelAlarmResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBodyResult) SetAlarmId(v int64) *ListHotelAlarmResponseBodyResult {
	s.AlarmId = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResult) SetDeviceOpenId(v string) *ListHotelAlarmResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResult) SetScheduleInfo(v *ListHotelAlarmResponseBodyResultScheduleInfo) *ListHotelAlarmResponseBodyResult {
	s.ScheduleInfo = v
	return s
}

func (s *ListHotelAlarmResponseBodyResult) SetUserOpenId(v string) *ListHotelAlarmResponseBodyResult {
	s.UserOpenId = &v
	return s
}

type ListHotelAlarmResponseBodyResultScheduleInfo struct {
	Once *ListHotelAlarmResponseBodyResultScheduleInfoOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// ONCE, WEEKLY
	//
	// example:
	//
	// ONCE
	Type   *string                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly *ListHotelAlarmResponseBodyResultScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s ListHotelAlarmResponseBodyResultScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmResponseBodyResultScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) SetOnce(v *ListHotelAlarmResponseBodyResultScheduleInfoOnce) *ListHotelAlarmResponseBodyResultScheduleInfo {
	s.Once = v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) SetType(v string) *ListHotelAlarmResponseBodyResultScheduleInfo {
	s.Type = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) SetWeekly(v *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) *ListHotelAlarmResponseBodyResultScheduleInfo {
	s.Weekly = v
	return s
}

type ListHotelAlarmResponseBodyResultScheduleInfoOnce struct {
	// example:
	//
	// 20
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 30
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 9
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 2022
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ListHotelAlarmResponseBodyResultScheduleInfoOnce) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmResponseBodyResultScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetDay(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetHour(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetMinute(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetMonth(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetYear(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Year = &v
	return s
}

type ListHotelAlarmResponseBodyResultScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s ListHotelAlarmResponseBodyResultScheduleInfoWeekly) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmResponseBodyResultScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *ListHotelAlarmResponseBodyResultScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) SetHour(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) SetMinute(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoWeekly {
	s.Minute = &v
	return s
}

type ListHotelAlarmResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelAlarmResponse) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponse) SetHeaders(v map[string]*string) *ListHotelAlarmResponse {
	s.Headers = v
	return s
}

func (s *ListHotelAlarmResponse) SetStatusCode(v int32) *ListHotelAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelAlarmResponse) SetBody(v *ListHotelAlarmResponseBody) *ListHotelAlarmResponse {
	s.Body = v
	return s
}

type ListHotelControlDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelControlDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceHeaders) SetCommonHeaders(v map[string]*string) *ListHotelControlDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelControlDeviceHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelControlDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelControlDeviceHeaders) SetAuthorization(v string) *ListHotelControlDeviceHeaders {
	s.Authorization = &v
	return s
}

type ListHotelControlDeviceRequest struct {
	UserInfo *ListHotelControlDeviceRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelControlDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceRequest) SetUserInfo(v *ListHotelControlDeviceRequestUserInfo) *ListHotelControlDeviceRequest {
	s.UserInfo = v
	return s
}

type ListHotelControlDeviceRequestUserInfo struct {
	// This parameter is required.
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListHotelControlDeviceRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceRequestUserInfo) SetEncodeKey(v string) *ListHotelControlDeviceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetEncodeType(v string) *ListHotelControlDeviceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetId(v string) *ListHotelControlDeviceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetIdType(v string) *ListHotelControlDeviceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetOrganizationId(v string) *ListHotelControlDeviceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListHotelControlDeviceShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelControlDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceShrinkRequest) SetUserInfoShrink(v string) *ListHotelControlDeviceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListHotelControlDeviceResponseBody struct {
	Code      *int32               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelControlDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceResponseBody) SetCode(v int32) *ListHotelControlDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetMessage(v string) *ListHotelControlDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetRequestId(v string) *ListHotelControlDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetResult(v []map[string]*string) *ListHotelControlDeviceResponseBody {
	s.Result = v
	return s
}

type ListHotelControlDeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelControlDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelControlDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceResponse) SetHeaders(v map[string]*string) *ListHotelControlDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListHotelControlDeviceResponse) SetStatusCode(v int32) *ListHotelControlDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelControlDeviceResponse) SetBody(v *ListHotelControlDeviceResponseBody) *ListHotelControlDeviceResponse {
	s.Body = v
	return s
}

type ListHotelInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelInfoHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelInfoHeaders) SetCommonHeaders(v map[string]*string) *ListHotelInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelInfoHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelInfoHeaders) SetAuthorization(v string) *ListHotelInfoHeaders {
	s.Authorization = &v
	return s
}

type ListHotelInfoResponseBody struct {
	Extentions map[string]interface{}             `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListHotelInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListHotelInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelInfoResponseBody) SetExtentions(v map[string]interface{}) *ListHotelInfoResponseBody {
	s.Extentions = v
	return s
}

func (s *ListHotelInfoResponseBody) SetMessage(v string) *ListHotelInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelInfoResponseBody) SetRequestId(v string) *ListHotelInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelInfoResponseBody) SetResult(v []*ListHotelInfoResponseBodyResult) *ListHotelInfoResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelInfoResponseBody) SetStatusCode(v int32) *ListHotelInfoResponseBody {
	s.StatusCode = &v
	return s
}

type ListHotelInfoResponseBodyResult struct {
	AuthAccount  []*ListHotelInfoResponseBodyResultAuthAccount `json:"AuthAccount,omitempty" xml:"AuthAccount,omitempty" type:"Repeated"`
	HotelAddress *string                                       `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// example:
	//
	// cf2446fc9d144c85aaee4f9ae20a96e7
	HotelId   *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
}

func (s ListHotelInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelInfoResponseBodyResult) SetAuthAccount(v []*ListHotelInfoResponseBodyResultAuthAccount) *ListHotelInfoResponseBodyResult {
	s.AuthAccount = v
	return s
}

func (s *ListHotelInfoResponseBodyResult) SetHotelAddress(v string) *ListHotelInfoResponseBodyResult {
	s.HotelAddress = &v
	return s
}

func (s *ListHotelInfoResponseBodyResult) SetHotelId(v string) *ListHotelInfoResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *ListHotelInfoResponseBodyResult) SetHotelName(v string) *ListHotelInfoResponseBodyResult {
	s.HotelName = &v
	return s
}

type ListHotelInfoResponseBodyResultAuthAccount struct {
	// example:
	//
	// leetest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListHotelInfoResponseBodyResultAuthAccount) String() string {
	return tea.Prettify(s)
}

func (s ListHotelInfoResponseBodyResultAuthAccount) GoString() string {
	return s.String()
}

func (s *ListHotelInfoResponseBodyResultAuthAccount) SetUserName(v string) *ListHotelInfoResponseBodyResultAuthAccount {
	s.UserName = &v
	return s
}

type ListHotelInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelInfoResponse) GoString() string {
	return s.String()
}

func (s *ListHotelInfoResponse) SetHeaders(v map[string]*string) *ListHotelInfoResponse {
	s.Headers = v
	return s
}

func (s *ListHotelInfoResponse) SetStatusCode(v int32) *ListHotelInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelInfoResponse) SetBody(v *ListHotelInfoResponseBody) *ListHotelInfoResponse {
	s.Body = v
	return s
}

type ListHotelMessageTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelMessageTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelMessageTemplateHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelMessageTemplateHeaders) SetCommonHeaders(v map[string]*string) *ListHotelMessageTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelMessageTemplateHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelMessageTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelMessageTemplateHeaders) SetAuthorization(v string) *ListHotelMessageTemplateHeaders {
	s.Authorization = &v
	return s
}

type ListHotelMessageTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelMessageTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelMessageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelMessageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelMessageTemplateResponseBody) SetCode(v int32) *ListHotelMessageTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBody) SetMessage(v string) *ListHotelMessageTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBody) SetRequestId(v string) *ListHotelMessageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBody) SetResult(v []*ListHotelMessageTemplateResponseBodyResult) *ListHotelMessageTemplateResponseBody {
	s.Result = v
	return s
}

type ListHotelMessageTemplateResponseBodyResult struct {
	// example:
	//
	// 不通过
	AuditMark *string `json:"AuditMark,omitempty" xml:"AuditMark,omitempty"`
	// example:
	//
	// COMMIT
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// 这是${hotel}的一个测试模板
	TemplateDetail *string `json:"TemplateDetail,omitempty" xml:"TemplateDetail,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 测试模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListHotelMessageTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelMessageTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetAuditMark(v string) *ListHotelMessageTemplateResponseBodyResult {
	s.AuditMark = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetAuditStatus(v string) *ListHotelMessageTemplateResponseBodyResult {
	s.AuditStatus = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetTemplateDetail(v string) *ListHotelMessageTemplateResponseBodyResult {
	s.TemplateDetail = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetTemplateId(v int64) *ListHotelMessageTemplateResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetTemplateName(v string) *ListHotelMessageTemplateResponseBodyResult {
	s.TemplateName = &v
	return s
}

type ListHotelMessageTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelMessageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelMessageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelMessageTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListHotelMessageTemplateResponse) SetHeaders(v map[string]*string) *ListHotelMessageTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListHotelMessageTemplateResponse) SetStatusCode(v int32) *ListHotelMessageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelMessageTemplateResponse) SetBody(v *ListHotelMessageTemplateResponseBody) *ListHotelMessageTemplateResponse {
	s.Body = v
	return s
}

type ListHotelOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelOrderHeaders) SetCommonHeaders(v map[string]*string) *ListHotelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelOrderHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelOrderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelOrderHeaders) SetAuthorization(v string) *ListHotelOrderHeaders {
	s.Authorization = &v
	return s
}

type ListHotelOrderRequest struct {
	// This parameter is required.
	Payload *ListHotelOrderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListHotelOrderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderRequest) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequest) SetPayload(v *ListHotelOrderRequestPayload) *ListHotelOrderRequest {
	s.Payload = v
	return s
}

func (s *ListHotelOrderRequest) SetUserInfo(v *ListHotelOrderRequestUserInfo) *ListHotelOrderRequest {
	s.UserInfo = v
	return s
}

type ListHotelOrderRequestPayload struct {
	// This parameter is required.
	Page *ListHotelOrderRequestPayloadPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListHotelOrderRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestPayload) SetPage(v *ListHotelOrderRequestPayloadPage) *ListHotelOrderRequestPayload {
	s.Page = v
	return s
}

type ListHotelOrderRequestPayloadPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotelOrderRequestPayloadPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderRequestPayloadPage) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestPayloadPage) SetPageNumber(v int64) *ListHotelOrderRequestPayloadPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelOrderRequestPayloadPage) SetPageSize(v int64) *ListHotelOrderRequestPayloadPage {
	s.PageSize = &v
	return s
}

type ListHotelOrderRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1248494721591392955
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mFU6VtVU+pgA8lx6rYMo7SPl11t+8b+8ALrn10MIPEdpK/HI9wELAEppYhPI1cYRDa4og8AMjAEBZKbLUwFjFA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListHotelOrderRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestUserInfo) SetEncodeKey(v string) *ListHotelOrderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetEncodeType(v string) *ListHotelOrderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetId(v string) *ListHotelOrderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetIdType(v string) *ListHotelOrderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetOrganizationId(v string) *ListHotelOrderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListHotelOrderShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelOrderShrinkRequest) SetPayloadShrink(v string) *ListHotelOrderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListHotelOrderShrinkRequest) SetUserInfoShrink(v string) *ListHotelOrderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListHotelOrderResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListHotelOrderResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 07F61FDA-606F-10A0-8ED0-C6CE62710A48
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelOrderResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBody) SetCode(v int32) *ListHotelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetMessage(v string) *ListHotelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetPage(v *ListHotelOrderResponseBodyPage) *ListHotelOrderResponseBody {
	s.Page = v
	return s
}

func (s *ListHotelOrderResponseBody) SetRequestId(v string) *ListHotelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetResult(v []*ListHotelOrderResponseBodyResult) *ListHotelOrderResponseBody {
	s.Result = v
	return s
}

type ListHotelOrderResponseBodyPage struct {
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 21
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 7
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelOrderResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBodyPage) SetHasNext(v bool) *ListHotelOrderResponseBodyPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetPageNumber(v int32) *ListHotelOrderResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetPageSize(v int32) *ListHotelOrderResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetTotal(v int32) *ListHotelOrderResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetTotalPage(v int32) *ListHotelOrderResponseBodyPage {
	s.TotalPage = &v
	return s
}

type ListHotelOrderResponseBodyResult struct {
	// example:
	//
	// 21.5
	ApplyAmt *int64 `json:"ApplyAmt,omitempty" xml:"ApplyAmt,omitempty"`
	// example:
	//
	// 1659952892000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 20220808180132000114508555527711
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// 12
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// example:
	//
	// 2001
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// REPAIR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingfenlei/shebeiweixiu.png
	TypeIconUrl *string `json:"TypeIconUrl,omitempty" xml:"TypeIconUrl,omitempty"`
	// example:
	//
	// 设备维修
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListHotelOrderResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBodyResult) SetApplyAmt(v int64) *ListHotelOrderResponseBodyResult {
	s.ApplyAmt = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetGmtCreate(v int64) *ListHotelOrderResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetOrderNo(v string) *ListHotelOrderResponseBodyResult {
	s.OrderNo = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetQuantity(v int64) *ListHotelOrderResponseBodyResult {
	s.Quantity = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetRoomNo(v string) *ListHotelOrderResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetStatus(v string) *ListHotelOrderResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetType(v string) *ListHotelOrderResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetTypeIconUrl(v string) *ListHotelOrderResponseBodyResult {
	s.TypeIconUrl = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetTypeName(v string) *ListHotelOrderResponseBodyResult {
	s.TypeName = &v
	return s
}

type ListHotelOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderResponse) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponse) SetHeaders(v map[string]*string) *ListHotelOrderResponse {
	s.Headers = v
	return s
}

func (s *ListHotelOrderResponse) SetStatusCode(v int32) *ListHotelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelOrderResponse) SetBody(v *ListHotelOrderResponseBody) *ListHotelOrderResponse {
	s.Body = v
	return s
}

type ListHotelRoomsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelRoomsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelRoomsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsHeaders) SetCommonHeaders(v map[string]*string) *ListHotelRoomsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelRoomsHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelRoomsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelRoomsHeaders) SetAuthorization(v string) *ListHotelRoomsHeaders {
	s.Authorization = &v
	return s
}

type ListHotelRoomsRequest struct {
	// if can be null:
	// true
	HotelAdminRoom *ListHotelRoomsRequestHotelAdminRoom `json:"HotelAdminRoom,omitempty" xml:"HotelAdminRoom,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ListHotelRoomsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsRequest) SetHotelAdminRoom(v *ListHotelRoomsRequestHotelAdminRoom) *ListHotelRoomsRequest {
	s.HotelAdminRoom = v
	return s
}

func (s *ListHotelRoomsRequest) SetHotelId(v string) *ListHotelRoomsRequest {
	s.HotelId = &v
	return s
}

type ListHotelRoomsRequestHotelAdminRoom struct {
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ListHotelRoomsRequestHotelAdminRoom) String() string {
	return tea.Prettify(s)
}

func (s ListHotelRoomsRequestHotelAdminRoom) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsRequestHotelAdminRoom) SetRoomNo(v string) *ListHotelRoomsRequestHotelAdminRoom {
	s.RoomNo = &v
	return s
}

type ListHotelRoomsShrinkRequest struct {
	// if can be null:
	// true
	HotelAdminRoomShrink *string `json:"HotelAdminRoom,omitempty" xml:"HotelAdminRoom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ListHotelRoomsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelRoomsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsShrinkRequest) SetHotelAdminRoomShrink(v string) *ListHotelRoomsShrinkRequest {
	s.HotelAdminRoomShrink = &v
	return s
}

func (s *ListHotelRoomsShrinkRequest) SetHotelId(v string) *ListHotelRoomsShrinkRequest {
	s.HotelId = &v
	return s
}

type ListHotelRoomsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelRoomsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelRoomsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsResponseBody) SetCode(v int32) *ListHotelRoomsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelRoomsResponseBody) SetMessage(v string) *ListHotelRoomsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelRoomsResponseBody) SetRequestId(v string) *ListHotelRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelRoomsResponseBody) SetResult(v []*ListHotelRoomsResponseBodyResult) *ListHotelRoomsResponseBody {
	s.Result = v
	return s
}

type ListHotelRoomsResponseBodyResult struct {
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 102
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ListHotelRoomsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelRoomsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsResponseBodyResult) SetHotelId(v string) *ListHotelRoomsResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *ListHotelRoomsResponseBodyResult) SetRoomNo(v string) *ListHotelRoomsResponseBodyResult {
	s.RoomNo = &v
	return s
}

type ListHotelRoomsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelRoomsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsResponse) SetHeaders(v map[string]*string) *ListHotelRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListHotelRoomsResponse) SetStatusCode(v int32) *ListHotelRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelRoomsResponse) SetBody(v *ListHotelRoomsResponseBody) *ListHotelRoomsResponse {
	s.Body = v
	return s
}

type ListHotelSceneBookItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelSceneBookItemsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsHeaders) SetCommonHeaders(v map[string]*string) *ListHotelSceneBookItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelSceneBookItemsHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelSceneBookItemsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelSceneBookItemsHeaders) SetAuthorization(v string) *ListHotelSceneBookItemsHeaders {
	s.Authorization = &v
	return s
}

type ListHotelSceneBookItemsRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	Page *ListHotelSceneBookItemsRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// FOOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneBookItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsRequest) SetHotelId(v string) *ListHotelSceneBookItemsRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelSceneBookItemsRequest) SetPage(v *ListHotelSceneBookItemsRequestPage) *ListHotelSceneBookItemsRequest {
	s.Page = v
	return s
}

func (s *ListHotelSceneBookItemsRequest) SetType(v string) *ListHotelSceneBookItemsRequest {
	s.Type = &v
	return s
}

type ListHotelSceneBookItemsRequestPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotelSceneBookItemsRequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsRequestPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsRequestPage) SetPageNumber(v int32) *ListHotelSceneBookItemsRequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneBookItemsRequestPage) SetPageSize(v int32) *ListHotelSceneBookItemsRequestPage {
	s.PageSize = &v
	return s
}

type ListHotelSceneBookItemsShrinkRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FOOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneBookItemsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsShrinkRequest) SetHotelId(v string) *ListHotelSceneBookItemsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelSceneBookItemsShrinkRequest) SetPageShrink(v string) *ListHotelSceneBookItemsShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListHotelSceneBookItemsShrinkRequest) SetType(v string) *ListHotelSceneBookItemsShrinkRequest {
	s.Type = &v
	return s
}

type ListHotelSceneBookItemsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelSceneBookItemsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelSceneBookItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponseBody) SetCode(v int32) *ListHotelSceneBookItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBody) SetMessage(v string) *ListHotelSceneBookItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBody) SetRequestId(v string) *ListHotelSceneBookItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBody) SetResult(v *ListHotelSceneBookItemsResponseBodyResult) *ListHotelSceneBookItemsResponseBody {
	s.Result = v
	return s
}

type ListHotelSceneBookItemsResponseBodyResult struct {
	Page          *ListHotelSceneBookItemsResponseBodyResultPage            `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	SceneItemList []*ListHotelSceneBookItemsResponseBodyResultSceneItemList `json:"SceneItemList,omitempty" xml:"SceneItemList,omitempty" type:"Repeated"`
}

func (s ListHotelSceneBookItemsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponseBodyResult) SetPage(v *ListHotelSceneBookItemsResponseBodyResultPage) *ListHotelSceneBookItemsResponseBodyResult {
	s.Page = v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResult) SetSceneItemList(v []*ListHotelSceneBookItemsResponseBodyResultSceneItemList) *ListHotelSceneBookItemsResponseBodyResult {
	s.SceneItemList = v
	return s
}

type ListHotelSceneBookItemsResponseBodyResultPage struct {
	// example:
	//
	// True
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 23
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelSceneBookItemsResponseBodyResultPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsResponseBodyResultPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetHasNext(v bool) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetPageNumber(v int32) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetPageSize(v int32) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetTotal(v int32) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.Total = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetTotalPage(v int32) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.TotalPage = &v
	return s
}

type ListHotelSceneBookItemsResponseBodyResultSceneItemList struct {
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/28d7a91e3c05db3855725fc39e0387e7/welcome_audios/aa918294b6ca3aa115c51135bf9b80cb/l9f996sq.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 11824
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 青椒肉丝
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1850
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// FOOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1666161803
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListHotelSceneBookItemsResponseBodyResultSceneItemList) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsResponseBodyResultSceneItemList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetIcon(v string) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Icon = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetId(v int64) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Id = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetName(v string) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Name = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetPrice(v int64) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Price = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetStatus(v string) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Status = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetType(v string) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Type = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetUpdateTime(v int64) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.UpdateTime = &v
	return s
}

type ListHotelSceneBookItemsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelSceneBookItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelSceneBookItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneBookItemsResponse) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponse) SetHeaders(v map[string]*string) *ListHotelSceneBookItemsResponse {
	s.Headers = v
	return s
}

func (s *ListHotelSceneBookItemsResponse) SetStatusCode(v int32) *ListHotelSceneBookItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelSceneBookItemsResponse) SetBody(v *ListHotelSceneBookItemsResponseBody) *ListHotelSceneBookItemsResponse {
	s.Body = v
	return s
}

type ListHotelSceneItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelSceneItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemHeaders) SetCommonHeaders(v map[string]*string) *ListHotelSceneItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelSceneItemHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelSceneItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelSceneItemHeaders) SetAuthorization(v string) *ListHotelSceneItemHeaders {
	s.Authorization = &v
	return s
}

type ListHotelSceneItemRequest struct {
	// This parameter is required.
	Payload *ListHotelSceneItemRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListHotelSceneItemRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequest) SetPayload(v *ListHotelSceneItemRequestPayload) *ListHotelSceneItemRequest {
	s.Payload = v
	return s
}

func (s *ListHotelSceneItemRequest) SetUserInfo(v *ListHotelSceneItemRequestUserInfo) *ListHotelSceneItemRequest {
	s.UserInfo = v
	return s
}

type ListHotelSceneItemRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneItemRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequestPayload) SetType(v string) *ListHotelSceneItemRequestPayload {
	s.Type = &v
	return s
}

type ListHotelSceneItemRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1248494721591392955
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mFU6VtVU+pgA8lx6rYMo7SPl11t+8b+8ALrn10MIPEdpK/HI9wELAEppYhPI1cYRDa4og8AMjAEBZKbLUwFjFA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListHotelSceneItemRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequestUserInfo) SetEncodeKey(v string) *ListHotelSceneItemRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetEncodeType(v string) *ListHotelSceneItemRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetId(v string) *ListHotelSceneItemRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetIdType(v string) *ListHotelSceneItemRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetOrganizationId(v string) *ListHotelSceneItemRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListHotelSceneItemShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelSceneItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemShrinkRequest) SetPayloadShrink(v string) *ListHotelSceneItemShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListHotelSceneItemShrinkRequest) SetUserInfoShrink(v string) *ListHotelSceneItemShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListHotelSceneItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListHotelSceneItemResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// CEADB586-51CB-1B6B-95BD-AB85A7A08E97
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelSceneItemResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBody) SetCode(v int32) *ListHotelSceneItemResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetMessage(v string) *ListHotelSceneItemResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetPage(v *ListHotelSceneItemResponseBodyPage) *ListHotelSceneItemResponseBody {
	s.Page = v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetRequestId(v string) *ListHotelSceneItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetResult(v *ListHotelSceneItemResponseBodyResult) *ListHotelSceneItemResponseBody {
	s.Result = v
	return s
}

type ListHotelSceneItemResponseBodyPage struct {
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 6
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelSceneItemResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyPage) SetHasNext(v bool) *ListHotelSceneItemResponseBodyPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetPageNumber(v int32) *ListHotelSceneItemResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetPageSize(v int32) *ListHotelSceneItemResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetTotal(v int32) *ListHotelSceneItemResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetTotalPage(v int32) *ListHotelSceneItemResponseBodyPage {
	s.TotalPage = &v
	return s
}

type ListHotelSceneItemResponseBodyResult struct {
	SecondCategoryList []*ListHotelSceneItemResponseBodyResultSecondCategoryList `json:"SecondCategoryList,omitempty" xml:"SecondCategoryList,omitempty" type:"Repeated"`
}

func (s ListHotelSceneItemResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResult) SetSecondCategoryList(v []*ListHotelSceneItemResponseBodyResultSecondCategoryList) *ListHotelSceneItemResponseBodyResult {
	s.SecondCategoryList = v
	return s
}

type ListHotelSceneItemResponseBodyResultSecondCategoryList struct {
	ItemList []*ListHotelSceneItemResponseBodyResultSecondCategoryListItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	// example:
	//
	// 客用品类
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryList) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) SetItemList(v []*ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) *ListHotelSceneItemResponseBodyResultSecondCategoryList {
	s.ItemList = v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) SetSecondCategoryName(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryList {
	s.SecondCategoryName = &v
	return s
}

type ListHotelSceneItemResponseBodyResultSecondCategoryListItemList struct {
	// example:
	//
	// 客用品类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/wupin/keyongpinlei/mianqian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 152860
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 棉签
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Price        *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	ResidueLimit *int64 `json:"ResidueLimit,omitempty" xml:"ResidueLimit,omitempty"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetCategory(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Category = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetIcon(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Icon = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetId(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Id = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetName(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Name = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetPrice(v int64) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Price = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetResidueLimit(v int64) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.ResidueLimit = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetStatus(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Status = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetType(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Type = &v
	return s
}

type ListHotelSceneItemResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelSceneItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelSceneItemResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponse) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponse) SetHeaders(v map[string]*string) *ListHotelSceneItemResponse {
	s.Headers = v
	return s
}

func (s *ListHotelSceneItemResponse) SetStatusCode(v int32) *ListHotelSceneItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelSceneItemResponse) SetBody(v *ListHotelSceneItemResponseBody) *ListHotelSceneItemResponse {
	s.Body = v
	return s
}

type ListHotelSceneItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelSceneItemsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsHeaders) SetCommonHeaders(v map[string]*string) *ListHotelSceneItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelSceneItemsHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelSceneItemsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelSceneItemsHeaders) SetAuthorization(v string) *ListHotelSceneItemsHeaders {
	s.Authorization = &v
	return s
}

type ListHotelSceneItemsRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// ListHotelSceneReq
	//
	// This parameter is required.
	ListHotelSceneReq *ListHotelSceneItemsRequestListHotelSceneReq `json:"ListHotelSceneReq,omitempty" xml:"ListHotelSceneReq,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsRequest) SetHotelId(v string) *ListHotelSceneItemsRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelSceneItemsRequest) SetListHotelSceneReq(v *ListHotelSceneItemsRequestListHotelSceneReq) *ListHotelSceneItemsRequest {
	s.ListHotelSceneReq = v
	return s
}

type ListHotelSceneItemsRequestListHotelSceneReq struct {
	// example:
	//
	// 客用品类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 棉签
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// This parameter is required.
	Page *ListHotelSceneItemsRequestListHotelSceneReqPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneItemsRequestListHotelSceneReq) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsRequestListHotelSceneReq) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetCategory(v string) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Category = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetKeywords(v string) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Keywords = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetPage(v *ListHotelSceneItemsRequestListHotelSceneReqPage) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Page = v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetStatus(v string) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Status = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetType(v string) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Type = &v
	return s
}

type ListHotelSceneItemsRequestListHotelSceneReqPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotelSceneItemsRequestListHotelSceneReqPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsRequestListHotelSceneReqPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsRequestListHotelSceneReqPage) SetPageNumber(v int32) *ListHotelSceneItemsRequestListHotelSceneReqPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReqPage) SetPageSize(v int32) *ListHotelSceneItemsRequestListHotelSceneReqPage {
	s.PageSize = &v
	return s
}

type ListHotelSceneItemsShrinkRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// ListHotelSceneReq
	//
	// This parameter is required.
	ListHotelSceneReqShrink *string `json:"ListHotelSceneReq,omitempty" xml:"ListHotelSceneReq,omitempty"`
}

func (s ListHotelSceneItemsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsShrinkRequest) SetHotelId(v string) *ListHotelSceneItemsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelSceneItemsShrinkRequest) SetListHotelSceneReqShrink(v string) *ListHotelSceneItemsShrinkRequest {
	s.ListHotelSceneReqShrink = &v
	return s
}

type ListHotelSceneItemsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelSceneItemsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponseBody) SetCode(v int32) *ListHotelSceneItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelSceneItemsResponseBody) SetMessage(v string) *ListHotelSceneItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelSceneItemsResponseBody) SetRequestId(v string) *ListHotelSceneItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelSceneItemsResponseBody) SetResult(v *ListHotelSceneItemsResponseBodyResult) *ListHotelSceneItemsResponseBody {
	s.Result = v
	return s
}

type ListHotelSceneItemsResponseBodyResult struct {
	Page          *ListHotelSceneItemsResponseBodyResultPage            `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	SceneItemList []*ListHotelSceneItemsResponseBodyResultSceneItemList `json:"SceneItemList,omitempty" xml:"SceneItemList,omitempty" type:"Repeated"`
}

func (s ListHotelSceneItemsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponseBodyResult) SetPage(v *ListHotelSceneItemsResponseBodyResultPage) *ListHotelSceneItemsResponseBodyResult {
	s.Page = v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResult) SetSceneItemList(v []*ListHotelSceneItemsResponseBodyResultSceneItemList) *ListHotelSceneItemsResponseBodyResult {
	s.SceneItemList = v
	return s
}

type ListHotelSceneItemsResponseBodyResultPage struct {
	// example:
	//
	// False
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 23
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelSceneItemsResponseBodyResultPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsResponseBodyResultPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetHasNext(v bool) *ListHotelSceneItemsResponseBodyResultPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetPageNumber(v int32) *ListHotelSceneItemsResponseBodyResultPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetPageSize(v int32) *ListHotelSceneItemsResponseBodyResultPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetTotal(v int32) *ListHotelSceneItemsResponseBodyResultPage {
	s.Total = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetTotalPage(v int32) *ListHotelSceneItemsResponseBodyResultPage {
	s.TotalPage = &v
	return s
}

type ListHotelSceneItemsResponseBodyResultSceneItemList struct {
	BeyondLimitReply *string `json:"BeyondLimitReply,omitempty" xml:"BeyondLimitReply,omitempty"`
	// example:
	//
	// 客用品类
	Category       *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeliveryMethod *string `json:"DeliveryMethod,omitempty" xml:"DeliveryMethod,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/wupin/keyongpinlei/mianqian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// id
	//
	// example:
	//
	// 10336
	Id          *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	LimitNumber *int32 `json:"LimitNumber,omitempty" xml:"LimitNumber,omitempty"`
	LimitSwitch *int32 `json:"LimitSwitch,omitempty" xml:"LimitSwitch,omitempty"`
	// example:
	//
	// 棉签
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// example:
	//
	// 160
	Price     *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1666163226
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListHotelSceneItemsResponseBodyResultSceneItemList) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsResponseBodyResultSceneItemList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetBeyondLimitReply(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.BeyondLimitReply = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetCategory(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Category = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetDeliveryMethod(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.DeliveryMethod = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetIcon(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Icon = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetId(v int64) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Id = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetLimitNumber(v int32) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.LimitNumber = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetLimitSwitch(v int32) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.LimitSwitch = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetName(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Name = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetPaymentMethod(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.PaymentMethod = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetPrice(v int64) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Price = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetRobotName(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.RobotName = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetStatus(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Status = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetType(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Type = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetUpdateTime(v int64) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.UpdateTime = &v
	return s
}

type ListHotelSceneItemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelSceneItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelSceneItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemsResponse) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponse) SetHeaders(v map[string]*string) *ListHotelSceneItemsResponse {
	s.Headers = v
	return s
}

func (s *ListHotelSceneItemsResponse) SetStatusCode(v int32) *ListHotelSceneItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelSceneItemsResponse) SetBody(v *ListHotelSceneItemsResponseBody) *ListHotelSceneItemsResponse {
	s.Body = v
	return s
}

type ListHotelServiceCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelServiceCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryHeaders) SetCommonHeaders(v map[string]*string) *ListHotelServiceCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelServiceCategoryHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelServiceCategoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelServiceCategoryHeaders) SetAuthorization(v string) *ListHotelServiceCategoryHeaders {
	s.Authorization = &v
	return s
}

type ListHotelServiceCategoryRequest struct {
	// This parameter is required.
	Payload *ListHotelServiceCategoryRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
}

func (s ListHotelServiceCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryRequest) SetPayload(v *ListHotelServiceCategoryRequestPayload) *ListHotelServiceCategoryRequest {
	s.Payload = v
	return s
}

type ListHotelServiceCategoryRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// HOTEL_SERVICE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelServiceCategoryRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryRequestPayload) SetType(v string) *ListHotelServiceCategoryRequestPayload {
	s.Type = &v
	return s
}

type ListHotelServiceCategoryShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s ListHotelServiceCategoryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryShrinkRequest) SetPayloadShrink(v string) *ListHotelServiceCategoryShrinkRequest {
	s.PayloadShrink = &v
	return s
}

type ListHotelServiceCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5373C821-65D2-1764-B9F9-951914937FF5
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelServiceCategoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelServiceCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponseBody) SetCode(v int32) *ListHotelServiceCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetMessage(v string) *ListHotelServiceCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetRequestId(v string) *ListHotelServiceCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetResult(v []*ListHotelServiceCategoryResponseBodyResult) *ListHotelServiceCategoryResponseBody {
	s.Result = v
	return s
}

type ListHotelServiceCategoryResponseBodyResult struct {
	// example:
	//
	// GOODS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 生活及洗漱用品
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingfenlei/wupintianjia.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 物品添加
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// HOTEL_SERVICE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelServiceCategoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetCode(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetDesc(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetIcon(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetName(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetType(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Type = &v
	return s
}

type ListHotelServiceCategoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelServiceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelServiceCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponse) SetHeaders(v map[string]*string) *ListHotelServiceCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListHotelServiceCategoryResponse) SetStatusCode(v int32) *ListHotelServiceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelServiceCategoryResponse) SetBody(v *ListHotelServiceCategoryResponseBody) *ListHotelServiceCategoryResponse {
	s.Body = v
	return s
}

type ListHotelsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelsHeaders) SetCommonHeaders(v map[string]*string) *ListHotelsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelsHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelsHeaders) SetAuthorization(v string) *ListHotelsHeaders {
	s.Authorization = &v
	return s
}

type ListHotelsRequest struct {
	// if can be null:
	// true
	HotelRequest *ListHotelsRequestHotelRequest `json:"HotelRequest,omitempty" xml:"HotelRequest,omitempty" type:"Struct"`
	// This parameter is required.
	Page *ListHotelsRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListHotelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsRequest) GoString() string {
	return s.String()
}

func (s *ListHotelsRequest) SetHotelRequest(v *ListHotelsRequestHotelRequest) *ListHotelsRequest {
	s.HotelRequest = v
	return s
}

func (s *ListHotelsRequest) SetPage(v *ListHotelsRequestPage) *ListHotelsRequest {
	s.Page = v
	return s
}

func (s *ListHotelsRequest) SetStatus(v int32) *ListHotelsRequest {
	s.Status = &v
	return s
}

type ListHotelsRequestHotelRequest struct {
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ListHotelsRequestHotelRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsRequestHotelRequest) GoString() string {
	return s.String()
}

func (s *ListHotelsRequestHotelRequest) SetHotelId(v string) *ListHotelsRequestHotelRequest {
	s.HotelId = &v
	return s
}

type ListHotelsRequestPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotelsRequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsRequestPage) GoString() string {
	return s.String()
}

func (s *ListHotelsRequestPage) SetPageNumber(v int32) *ListHotelsRequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelsRequestPage) SetPageSize(v int32) *ListHotelsRequestPage {
	s.PageSize = &v
	return s
}

type ListHotelsShrinkRequest struct {
	// if can be null:
	// true
	HotelRequestShrink *string `json:"HotelRequest,omitempty" xml:"HotelRequest,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListHotelsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelsShrinkRequest) SetHotelRequestShrink(v string) *ListHotelsShrinkRequest {
	s.HotelRequestShrink = &v
	return s
}

func (s *ListHotelsShrinkRequest) SetPageShrink(v string) *ListHotelsShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListHotelsShrinkRequest) SetStatus(v int32) *ListHotelsShrinkRequest {
	s.Status = &v
	return s
}

type ListHotelsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelsResponseBody) SetCode(v int32) *ListHotelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelsResponseBody) SetMessage(v string) *ListHotelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelsResponseBody) SetRequestId(v string) *ListHotelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelsResponseBody) SetResult(v *ListHotelsResponseBodyResult) *ListHotelsResponseBody {
	s.Result = v
	return s
}

type ListHotelsResponseBodyResult struct {
	HotelInfoList []*ListHotelsResponseBodyResultHotelInfoList `json:"HotelInfoList,omitempty" xml:"HotelInfoList,omitempty" type:"Repeated"`
	Page          *ListHotelsResponseBodyResultPage            `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListHotelsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelsResponseBodyResult) SetHotelInfoList(v []*ListHotelsResponseBodyResultHotelInfoList) *ListHotelsResponseBodyResult {
	s.HotelInfoList = v
	return s
}

func (s *ListHotelsResponseBodyResult) SetPage(v *ListHotelsResponseBodyResultPage) *ListHotelsResponseBodyResult {
	s.Page = v
	return s
}

type ListHotelsResponseBodyResultHotelInfoList struct {
	AccountNames []*string `json:"AccountNames,omitempty" xml:"AccountNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1654568802000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 酒店地址
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// example:
	//
	// 73ab1b03018d4da69b5bef17095f569b
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 酒店名称
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// example:
	//
	// 酒店
	IndustryType *string `json:"IndustryType,omitempty" xml:"IndustryType,omitempty"`
	// example:
	//
	// 13312340987
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 测试产品
	RelatedProductName *string `json:"RelatedProductName,omitempty" xml:"RelatedProductName,omitempty"`
	// example:
	//
	// 12
	RoomNum *int32 `json:"RoomNum,omitempty" xml:"RoomNum,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListHotelsResponseBodyResultHotelInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsResponseBodyResultHotelInfoList) GoString() string {
	return s.String()
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetAccountNames(v []*string) *ListHotelsResponseBodyResultHotelInfoList {
	s.AccountNames = v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetCreateTime(v int64) *ListHotelsResponseBodyResultHotelInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetHotelAddress(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.HotelAddress = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetHotelId(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.HotelId = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetHotelName(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.HotelName = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetIndustryType(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.IndustryType = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetPhoneNumber(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.PhoneNumber = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetRelatedProductName(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.RelatedProductName = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetRoomNum(v int32) *ListHotelsResponseBodyResultHotelInfoList {
	s.RoomNum = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetStatus(v int32) *ListHotelsResponseBodyResultHotelInfoList {
	s.Status = &v
	return s
}

type ListHotelsResponseBodyResultPage struct {
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 23
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelsResponseBodyResultPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsResponseBodyResultPage) GoString() string {
	return s.String()
}

func (s *ListHotelsResponseBodyResultPage) SetHasNext(v bool) *ListHotelsResponseBodyResultPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) SetPageNumber(v int32) *ListHotelsResponseBodyResultPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) SetPageSize(v int32) *ListHotelsResponseBodyResultPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) SetTotal(v int32) *ListHotelsResponseBodyResultPage {
	s.Total = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) SetTotalPage(v int32) *ListHotelsResponseBodyResultPage {
	s.TotalPage = &v
	return s
}

type ListHotelsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelsResponse) GoString() string {
	return s.String()
}

func (s *ListHotelsResponse) SetHeaders(v map[string]*string) *ListHotelsResponse {
	s.Headers = v
	return s
}

func (s *ListHotelsResponse) SetStatusCode(v int32) *ListHotelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelsResponse) SetBody(v *ListHotelsResponseBody) *ListHotelsResponse {
	s.Body = v
	return s
}

type ListInfraredDeviceBrandsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListInfraredDeviceBrandsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredDeviceBrandsHeaders) GoString() string {
	return s.String()
}

func (s *ListInfraredDeviceBrandsHeaders) SetCommonHeaders(v map[string]*string) *ListInfraredDeviceBrandsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInfraredDeviceBrandsHeaders) SetXAcsAligenieAccessToken(v string) *ListInfraredDeviceBrandsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListInfraredDeviceBrandsHeaders) SetAuthorization(v string) *ListInfraredDeviceBrandsHeaders {
	s.Authorization = &v
	return s
}

type ListInfraredDeviceBrandsRequest struct {
	// This parameter is required.
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
}

func (s ListInfraredDeviceBrandsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredDeviceBrandsRequest) GoString() string {
	return s.String()
}

func (s *ListInfraredDeviceBrandsRequest) SetCategory(v string) *ListInfraredDeviceBrandsRequest {
	s.Category = &v
	return s
}

func (s *ListInfraredDeviceBrandsRequest) SetServiceProvider(v string) *ListInfraredDeviceBrandsRequest {
	s.ServiceProvider = &v
	return s
}

type ListInfraredDeviceBrandsResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 860194F7-9593-50EA-8E53-BCEC0D325A00
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string][]*string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListInfraredDeviceBrandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredDeviceBrandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInfraredDeviceBrandsResponseBody) SetMessage(v string) *ListInfraredDeviceBrandsResponseBody {
	s.Message = &v
	return s
}

func (s *ListInfraredDeviceBrandsResponseBody) SetRequestId(v string) *ListInfraredDeviceBrandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInfraredDeviceBrandsResponseBody) SetResult(v map[string][]*string) *ListInfraredDeviceBrandsResponseBody {
	s.Result = v
	return s
}

func (s *ListInfraredDeviceBrandsResponseBody) SetStatusCode(v int32) *ListInfraredDeviceBrandsResponseBody {
	s.StatusCode = &v
	return s
}

type ListInfraredDeviceBrandsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInfraredDeviceBrandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInfraredDeviceBrandsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredDeviceBrandsResponse) GoString() string {
	return s.String()
}

func (s *ListInfraredDeviceBrandsResponse) SetHeaders(v map[string]*string) *ListInfraredDeviceBrandsResponse {
	s.Headers = v
	return s
}

func (s *ListInfraredDeviceBrandsResponse) SetStatusCode(v int32) *ListInfraredDeviceBrandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInfraredDeviceBrandsResponse) SetBody(v *ListInfraredDeviceBrandsResponseBody) *ListInfraredDeviceBrandsResponse {
	s.Body = v
	return s
}

type ListInfraredRemoteControllersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListInfraredRemoteControllersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredRemoteControllersHeaders) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersHeaders) SetCommonHeaders(v map[string]*string) *ListInfraredRemoteControllersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInfraredRemoteControllersHeaders) SetXAcsAligenieAccessToken(v string) *ListInfraredRemoteControllersHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListInfraredRemoteControllersHeaders) SetAuthorization(v string) *ListInfraredRemoteControllersHeaders {
	s.Authorization = &v
	return s
}

type ListInfraredRemoteControllersRequest struct {
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// This parameter is required.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId         *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	Province        *string `json:"Province,omitempty" xml:"Province,omitempty"`
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
}

func (s ListInfraredRemoteControllersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredRemoteControllersRequest) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersRequest) SetBrand(v string) *ListInfraredRemoteControllersRequest {
	s.Brand = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetCategory(v string) *ListInfraredRemoteControllersRequest {
	s.Category = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetCity(v string) *ListInfraredRemoteControllersRequest {
	s.City = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetHotelId(v string) *ListInfraredRemoteControllersRequest {
	s.HotelId = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetProvince(v string) *ListInfraredRemoteControllersRequest {
	s.Province = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetServiceProvider(v string) *ListInfraredRemoteControllersRequest {
	s.ServiceProvider = &v
	return s
}

type ListInfraredRemoteControllersResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0C90A059-3653-5356-A78E-8A6BDA606155
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListInfraredRemoteControllersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListInfraredRemoteControllersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredRemoteControllersResponseBody) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersResponseBody) SetMessage(v string) *ListInfraredRemoteControllersResponseBody {
	s.Message = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBody) SetRequestId(v string) *ListInfraredRemoteControllersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBody) SetResult(v []*ListInfraredRemoteControllersResponseBodyResult) *ListInfraredRemoteControllersResponseBody {
	s.Result = v
	return s
}

func (s *ListInfraredRemoteControllersResponseBody) SetStatusCode(v int32) *ListInfraredRemoteControllersResponseBody {
	s.StatusCode = &v
	return s
}

type ListInfraredRemoteControllersResponseBodyResult struct {
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 3747
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListInfraredRemoteControllersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredRemoteControllersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersResponseBodyResult) SetIndex(v int32) *ListInfraredRemoteControllersResponseBodyResult {
	s.Index = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBodyResult) SetRid(v int64) *ListInfraredRemoteControllersResponseBodyResult {
	s.Rid = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBodyResult) SetVersion(v string) *ListInfraredRemoteControllersResponseBodyResult {
	s.Version = &v
	return s
}

type ListInfraredRemoteControllersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInfraredRemoteControllersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInfraredRemoteControllersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInfraredRemoteControllersResponse) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersResponse) SetHeaders(v map[string]*string) *ListInfraredRemoteControllersResponse {
	s.Headers = v
	return s
}

func (s *ListInfraredRemoteControllersResponse) SetStatusCode(v int32) *ListInfraredRemoteControllersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInfraredRemoteControllersResponse) SetBody(v *ListInfraredRemoteControllersResponseBody) *ListInfraredRemoteControllersResponse {
	s.Body = v
	return s
}

type ListSTBServiceProvidersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSTBServiceProvidersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSTBServiceProvidersHeaders) GoString() string {
	return s.String()
}

func (s *ListSTBServiceProvidersHeaders) SetCommonHeaders(v map[string]*string) *ListSTBServiceProvidersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSTBServiceProvidersHeaders) SetXAcsAligenieAccessToken(v string) *ListSTBServiceProvidersHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSTBServiceProvidersHeaders) SetAuthorization(v string) *ListSTBServiceProvidersHeaders {
	s.Authorization = &v
	return s
}

type ListSTBServiceProvidersRequest struct {
	// This parameter is required.
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// This parameter is required.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListSTBServiceProvidersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSTBServiceProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListSTBServiceProvidersRequest) SetCity(v string) *ListSTBServiceProvidersRequest {
	s.City = &v
	return s
}

func (s *ListSTBServiceProvidersRequest) SetProvince(v string) *ListSTBServiceProvidersRequest {
	s.Province = &v
	return s
}

type ListSTBServiceProvidersResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1036C376-7A37-5A73-BE8B-C6DB40107EC1
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string][]*string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListSTBServiceProvidersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSTBServiceProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSTBServiceProvidersResponseBody) SetMessage(v string) *ListSTBServiceProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListSTBServiceProvidersResponseBody) SetRequestId(v string) *ListSTBServiceProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSTBServiceProvidersResponseBody) SetResult(v map[string][]*string) *ListSTBServiceProvidersResponseBody {
	s.Result = v
	return s
}

func (s *ListSTBServiceProvidersResponseBody) SetStatusCode(v int32) *ListSTBServiceProvidersResponseBody {
	s.StatusCode = &v
	return s
}

type ListSTBServiceProvidersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSTBServiceProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSTBServiceProvidersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSTBServiceProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListSTBServiceProvidersResponse) SetHeaders(v map[string]*string) *ListSTBServiceProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListSTBServiceProvidersResponse) SetStatusCode(v int32) *ListSTBServiceProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSTBServiceProvidersResponse) SetBody(v *ListSTBServiceProvidersResponseBody) *ListSTBServiceProvidersResponse {
	s.Body = v
	return s
}

type ListSceneCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSceneCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSceneCategoryHeaders) GoString() string {
	return s.String()
}

func (s *ListSceneCategoryHeaders) SetCommonHeaders(v map[string]*string) *ListSceneCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSceneCategoryHeaders) SetXAcsAligenieAccessToken(v string) *ListSceneCategoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSceneCategoryHeaders) SetAuthorization(v string) *ListSceneCategoryHeaders {
	s.Authorization = &v
	return s
}

type ListSceneCategoryRequest struct {
	// hotelId
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REPAIR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSceneCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSceneCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListSceneCategoryRequest) SetHotelId(v string) *ListSceneCategoryRequest {
	s.HotelId = &v
	return s
}

func (s *ListSceneCategoryRequest) SetType(v string) *ListSceneCategoryRequest {
	s.Type = &v
	return s
}

type ListSceneCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 0EC7*726E
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSceneCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSceneCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListSceneCategoryResponseBody) SetCode(v int32) *ListSceneCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListSceneCategoryResponseBody) SetMessage(v string) *ListSceneCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListSceneCategoryResponseBody) SetRequestId(v string) *ListSceneCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSceneCategoryResponseBody) SetResult(v []*string) *ListSceneCategoryResponseBody {
	s.Result = v
	return s
}

type ListSceneCategoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSceneCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSceneCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSceneCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListSceneCategoryResponse) SetHeaders(v map[string]*string) *ListSceneCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListSceneCategoryResponse) SetStatusCode(v int32) *ListSceneCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSceneCategoryResponse) SetBody(v *ListSceneCategoryResponseBody) *ListSceneCategoryResponse {
	s.Body = v
	return s
}

type ListServiceQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListServiceQAHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListServiceQAHeaders) GoString() string {
	return s.String()
}

func (s *ListServiceQAHeaders) SetCommonHeaders(v map[string]*string) *ListServiceQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListServiceQAHeaders) SetXAcsAligenieAccessToken(v string) *ListServiceQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListServiceQAHeaders) SetAuthorization(v string) *ListServiceQAHeaders {
	s.Authorization = &v
	return s
}

type ListServiceQARequest struct {
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// ***
	Keyword *string                   `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Page    *ListServiceQARequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListServiceQARequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceQARequest) GoString() string {
	return s.String()
}

func (s *ListServiceQARequest) SetActive(v bool) *ListServiceQARequest {
	s.Active = &v
	return s
}

func (s *ListServiceQARequest) SetHotelId(v string) *ListServiceQARequest {
	s.HotelId = &v
	return s
}

func (s *ListServiceQARequest) SetKeyword(v string) *ListServiceQARequest {
	s.Keyword = &v
	return s
}

func (s *ListServiceQARequest) SetPage(v *ListServiceQARequestPage) *ListServiceQARequest {
	s.Page = v
	return s
}

type ListServiceQARequestPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListServiceQARequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListServiceQARequestPage) GoString() string {
	return s.String()
}

func (s *ListServiceQARequestPage) SetPageNumber(v int32) *ListServiceQARequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListServiceQARequestPage) SetPageSize(v int32) *ListServiceQARequestPage {
	s.PageSize = &v
	return s
}

type ListServiceQAShrinkRequest struct {
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// ***
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s ListServiceQAShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListServiceQAShrinkRequest) SetActive(v bool) *ListServiceQAShrinkRequest {
	s.Active = &v
	return s
}

func (s *ListServiceQAShrinkRequest) SetHotelId(v string) *ListServiceQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListServiceQAShrinkRequest) SetKeyword(v string) *ListServiceQAShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListServiceQAShrinkRequest) SetPageShrink(v string) *ListServiceQAShrinkRequest {
	s.PageShrink = &v
	return s
}

type ListServiceQAResponseBody struct {
	// example:
	//
	// success
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListServiceQAResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListServiceQAResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListServiceQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceQAResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceQAResponseBody) SetMessage(v string) *ListServiceQAResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceQAResponseBody) SetPage(v *ListServiceQAResponseBodyPage) *ListServiceQAResponseBody {
	s.Page = v
	return s
}

func (s *ListServiceQAResponseBody) SetRequestId(v string) *ListServiceQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceQAResponseBody) SetResult(v []*ListServiceQAResponseBodyResult) *ListServiceQAResponseBody {
	s.Result = v
	return s
}

func (s *ListServiceQAResponseBody) SetStatusCode(v int32) *ListServiceQAResponseBody {
	s.StatusCode = &v
	return s
}

type ListServiceQAResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListServiceQAResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListServiceQAResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListServiceQAResponseBodyPage) SetPageNumber(v int32) *ListServiceQAResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListServiceQAResponseBodyPage) SetPageSize(v int32) *ListServiceQAResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListServiceQAResponseBodyPage) SetTotal(v int32) *ListServiceQAResponseBodyPage {
	s.Total = &v
	return s
}

type ListServiceQAResponseBodyResult struct {
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// example:
	//
	// ***
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 2022-07-27 14:06:27
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Question    *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// 1
	ServiceQAId *int64  `json:"ServiceQAId,omitempty" xml:"ServiceQAId,omitempty"`
	Templates   *string `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s ListServiceQAResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListServiceQAResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListServiceQAResponseBodyResult) SetActive(v bool) *ListServiceQAResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetAnswer(v string) *ListServiceQAResponseBodyResult {
	s.Answer = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetGmtModified(v string) *ListServiceQAResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetQuestion(v string) *ListServiceQAResponseBodyResult {
	s.Question = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetServiceQAId(v int64) *ListServiceQAResponseBodyResult {
	s.ServiceQAId = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetTemplates(v string) *ListServiceQAResponseBodyResult {
	s.Templates = &v
	return s
}

type ListServiceQAResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceQAResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceQAResponse) GoString() string {
	return s.String()
}

func (s *ListServiceQAResponse) SetHeaders(v map[string]*string) *ListServiceQAResponse {
	s.Headers = v
	return s
}

func (s *ListServiceQAResponse) SetStatusCode(v int32) *ListServiceQAResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceQAResponse) SetBody(v *ListServiceQAResponseBody) *ListServiceQAResponse {
	s.Body = v
	return s
}

type ListTicketsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListTicketsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsHeaders) GoString() string {
	return s.String()
}

func (s *ListTicketsHeaders) SetCommonHeaders(v map[string]*string) *ListTicketsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTicketsHeaders) SetXAcsAligenieAccessToken(v string) *ListTicketsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListTicketsHeaders) SetAuthorization(v string) *ListTicketsHeaders {
	s.Authorization = &v
	return s
}

type ListTicketsRequest struct {
	// example:
	//
	// 2022-09-14 14:23:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// true
	IsDesc *bool `json:"IsDesc,omitempty" xml:"IsDesc,omitempty"`
	// example:
	//
	// false
	IsNeedCallback *bool `json:"IsNeedCallback,omitempty" xml:"IsNeedCallback,omitempty"`
	// example:
	//
	// false
	IsNeedCharges *bool                   `json:"IsNeedCharges,omitempty" xml:"IsNeedCharges,omitempty"`
	Page          *ListTicketsRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// gmtCalled
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// example:
	//
	// 2022-04-08 09:39:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ""
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsRequest) SetEndTime(v string) *ListTicketsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTicketsRequest) SetHotelId(v string) *ListTicketsRequest {
	s.HotelId = &v
	return s
}

func (s *ListTicketsRequest) SetIsDesc(v bool) *ListTicketsRequest {
	s.IsDesc = &v
	return s
}

func (s *ListTicketsRequest) SetIsNeedCallback(v bool) *ListTicketsRequest {
	s.IsNeedCallback = &v
	return s
}

func (s *ListTicketsRequest) SetIsNeedCharges(v bool) *ListTicketsRequest {
	s.IsNeedCharges = &v
	return s
}

func (s *ListTicketsRequest) SetPage(v *ListTicketsRequestPage) *ListTicketsRequest {
	s.Page = v
	return s
}

func (s *ListTicketsRequest) SetRoomNo(v string) *ListTicketsRequest {
	s.RoomNo = &v
	return s
}

func (s *ListTicketsRequest) SetSortField(v string) *ListTicketsRequest {
	s.SortField = &v
	return s
}

func (s *ListTicketsRequest) SetStartTime(v string) *ListTicketsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTicketsRequest) SetStatus(v string) *ListTicketsRequest {
	s.Status = &v
	return s
}

func (s *ListTicketsRequest) SetType(v string) *ListTicketsRequest {
	s.Type = &v
	return s
}

type ListTicketsRequestPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTicketsRequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsRequestPage) GoString() string {
	return s.String()
}

func (s *ListTicketsRequestPage) SetPageNumber(v int32) *ListTicketsRequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsRequestPage) SetPageSize(v int32) *ListTicketsRequestPage {
	s.PageSize = &v
	return s
}

type ListTicketsShrinkRequest struct {
	// example:
	//
	// 2022-09-14 14:23:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// true
	IsDesc *bool `json:"IsDesc,omitempty" xml:"IsDesc,omitempty"`
	// example:
	//
	// false
	IsNeedCallback *bool `json:"IsNeedCallback,omitempty" xml:"IsNeedCallback,omitempty"`
	// example:
	//
	// false
	IsNeedCharges *bool   `json:"IsNeedCharges,omitempty" xml:"IsNeedCharges,omitempty"`
	PageShrink    *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// gmtCalled
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// example:
	//
	// 2022-04-08 09:39:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ""
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTicketsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsShrinkRequest) SetEndTime(v string) *ListTicketsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetHotelId(v string) *ListTicketsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetIsDesc(v bool) *ListTicketsShrinkRequest {
	s.IsDesc = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetIsNeedCallback(v bool) *ListTicketsShrinkRequest {
	s.IsNeedCallback = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetIsNeedCharges(v bool) *ListTicketsShrinkRequest {
	s.IsNeedCharges = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetPageShrink(v string) *ListTicketsShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetRoomNo(v string) *ListTicketsShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetSortField(v string) *ListTicketsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetStartTime(v string) *ListTicketsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetStatus(v string) *ListTicketsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetType(v string) *ListTicketsShrinkRequest {
	s.Type = &v
	return s
}

type ListTicketsResponseBody struct {
	// example:
	//
	// success
	Message *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListTicketsResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListTicketsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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

func (s *ListTicketsResponseBody) SetPage(v *ListTicketsResponseBodyPage) *ListTicketsResponseBody {
	s.Page = v
	return s
}

func (s *ListTicketsResponseBody) SetRequestId(v string) *ListTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketsResponseBody) SetResult(v []*ListTicketsResponseBodyResult) *ListTicketsResponseBody {
	s.Result = v
	return s
}

func (s *ListTicketsResponseBody) SetStatusCode(v int32) *ListTicketsResponseBody {
	s.StatusCode = &v
	return s
}

type ListTicketsResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTicketsResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyPage) SetPageNumber(v int32) *ListTicketsResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsResponseBodyPage) SetPageSize(v int32) *ListTicketsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListTicketsResponseBodyPage) SetTotal(v int32) *ListTicketsResponseBodyPage {
	s.Total = &v
	return s
}

type ListTicketsResponseBodyResult struct {
	// example:
	//
	// false
	Action *bool `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// ***
	AssignedHandler *string `json:"AssignedHandler,omitempty" xml:"AssignedHandler,omitempty"`
	// example:
	//
	// ***
	ChargesRemark *string `json:"ChargesRemark,omitempty" xml:"ChargesRemark,omitempty"`
	// example:
	//
	// ***
	CompleteRemark *string                                 `json:"CompleteRemark,omitempty" xml:"CompleteRemark,omitempty"`
	Dialogs        []*ListTicketsResponseBodyResultDialogs `json:"Dialogs,omitempty" xml:"Dialogs,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-01-09 00:00:00
	GmtCalled *string `json:"GmtCalled,omitempty" xml:"GmtCalled,omitempty"`
	// example:
	//
	// 2023-01-09 00:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2023-01-09 00:00:00
	GmtDelayed *string `json:"GmtDelayed,omitempty" xml:"GmtDelayed,omitempty"`
	// example:
	//
	// 2023-01-09 00:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 2023***93975
	GroupKey *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	// example:
	//
	// 45
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsAcceptedCharges *bool `json:"IsAcceptedCharges,omitempty" xml:"IsAcceptedCharges,omitempty"`
	// example:
	//
	// true
	IsDelayed *bool `json:"IsDelayed,omitempty" xml:"IsDelayed,omitempty"`
	// example:
	//
	// false
	IsNeedCallback *bool `json:"IsNeedCallback,omitempty" xml:"IsNeedCallback,omitempty"`
	// example:
	//
	// false
	IsNeedCharges *bool                    `json:"IsNeedCharges,omitempty" xml:"IsNeedCharges,omitempty"`
	Operations    []map[string]interface{} `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Repeated"`
	// example:
	//
	// ***
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 101
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ""
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTicketsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyResult) SetAction(v bool) *ListTicketsResponseBodyResult {
	s.Action = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetAssignedHandler(v string) *ListTicketsResponseBodyResult {
	s.AssignedHandler = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetChargesRemark(v string) *ListTicketsResponseBodyResult {
	s.ChargesRemark = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetCompleteRemark(v string) *ListTicketsResponseBodyResult {
	s.CompleteRemark = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetDialogs(v []*ListTicketsResponseBodyResultDialogs) *ListTicketsResponseBodyResult {
	s.Dialogs = v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGmtCalled(v string) *ListTicketsResponseBodyResult {
	s.GmtCalled = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGmtCreate(v string) *ListTicketsResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGmtDelayed(v string) *ListTicketsResponseBodyResult {
	s.GmtDelayed = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGmtModified(v string) *ListTicketsResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGroupKey(v string) *ListTicketsResponseBodyResult {
	s.GroupKey = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetId(v int64) *ListTicketsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetIsAcceptedCharges(v bool) *ListTicketsResponseBodyResult {
	s.IsAcceptedCharges = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetIsDelayed(v bool) *ListTicketsResponseBodyResult {
	s.IsDelayed = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetIsNeedCallback(v bool) *ListTicketsResponseBodyResult {
	s.IsNeedCallback = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetIsNeedCharges(v bool) *ListTicketsResponseBodyResult {
	s.IsNeedCharges = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetOperations(v []map[string]interface{}) *ListTicketsResponseBodyResult {
	s.Operations = v
	return s
}

func (s *ListTicketsResponseBodyResult) SetRemark(v string) *ListTicketsResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetRoomNo(v string) *ListTicketsResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetStatus(v string) *ListTicketsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetType(v string) *ListTicketsResponseBodyResult {
	s.Type = &v
	return s
}

type ListTicketsResponseBodyResultDialogs struct {
	Answer   *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
}

func (s ListTicketsResponseBodyResultDialogs) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBodyResultDialogs) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyResultDialogs) SetAnswer(v string) *ListTicketsResponseBodyResultDialogs {
	s.Answer = &v
	return s
}

func (s *ListTicketsResponseBodyResultDialogs) SetQuestion(v string) *ListTicketsResponseBodyResultDialogs {
	s.Question = &v
	return s
}

type ListTicketsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTicketsResponse) SetStatusCode(v int32) *ListTicketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTicketsResponse) SetBody(v *ListTicketsResponseBody) *ListTicketsResponse {
	s.Body = v
	return s
}

type PageGetHotelRoomDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PageGetHotelRoomDevicesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageGetHotelRoomDevicesHeaders) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesHeaders) SetCommonHeaders(v map[string]*string) *PageGetHotelRoomDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageGetHotelRoomDevicesHeaders) SetXAcsAligenieAccessToken(v string) *PageGetHotelRoomDevicesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PageGetHotelRoomDevicesHeaders) SetAuthorization(v string) *PageGetHotelRoomDevicesHeaders {
	s.Authorization = &v
	return s
}

type PageGetHotelRoomDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s PageGetHotelRoomDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s PageGetHotelRoomDevicesRequest) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesRequest) SetHotelId(v string) *PageGetHotelRoomDevicesRequest {
	s.HotelId = &v
	return s
}

func (s *PageGetHotelRoomDevicesRequest) SetPageNumber(v int32) *PageGetHotelRoomDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *PageGetHotelRoomDevicesRequest) SetPageSize(v int32) *PageGetHotelRoomDevicesRequest {
	s.PageSize = &v
	return s
}

type PageGetHotelRoomDevicesResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *PageGetHotelRoomDevicesResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 4EFBDDF4-B19D-526C-8C3D-CD8AB51974EE
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*PageGetHotelRoomDevicesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PageGetHotelRoomDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageGetHotelRoomDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesResponseBody) SetExtentions(v map[string]interface{}) *PageGetHotelRoomDevicesResponseBody {
	s.Extentions = v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetMessage(v string) *PageGetHotelRoomDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetPage(v *PageGetHotelRoomDevicesResponseBodyPage) *PageGetHotelRoomDevicesResponseBody {
	s.Page = v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetRequestId(v string) *PageGetHotelRoomDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetResult(v []*PageGetHotelRoomDevicesResponseBodyResult) *PageGetHotelRoomDevicesResponseBody {
	s.Result = v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetStatusCode(v int32) *PageGetHotelRoomDevicesResponseBody {
	s.StatusCode = &v
	return s
}

type PageGetHotelRoomDevicesResponseBodyPage struct {
	// example:
	//
	// 4
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// False
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// example:
	//
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s PageGetHotelRoomDevicesResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s PageGetHotelRoomDevicesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetEnd(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.End = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetHasNext(v bool) *PageGetHotelRoomDevicesResponseBodyPage {
	s.HasNext = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetPageNumber(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetPageSize(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetStart(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.Start = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetTotal(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.Total = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetTotalPage(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.TotalPage = &v
	return s
}

type PageGetHotelRoomDevicesResponseBodyResult struct {
	// example:
	//
	// V21.10.00.313
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// b4:xx:xx:xx:65:2b
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1
	OnlineStatus *int32 `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// 2001
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// 1200xxx048
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s PageGetHotelRoomDevicesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PageGetHotelRoomDevicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetFirmwareVersion(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.FirmwareVersion = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetHotelId(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetMac(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.Mac = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetOnlineStatus(v int32) *PageGetHotelRoomDevicesResponseBodyResult {
	s.OnlineStatus = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetRoomNo(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetSn(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.Sn = &v
	return s
}

type PageGetHotelRoomDevicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageGetHotelRoomDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageGetHotelRoomDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s PageGetHotelRoomDevicesResponse) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesResponse) SetHeaders(v map[string]*string) *PageGetHotelRoomDevicesResponse {
	s.Headers = v
	return s
}

func (s *PageGetHotelRoomDevicesResponse) SetStatusCode(v int32) *PageGetHotelRoomDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponse) SetBody(v *PageGetHotelRoomDevicesResponseBody) *PageGetHotelRoomDevicesResponse {
	s.Body = v
	return s
}

type PmsEventReportHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PmsEventReportHeaders) String() string {
	return tea.Prettify(s)
}

func (s PmsEventReportHeaders) GoString() string {
	return s.String()
}

func (s *PmsEventReportHeaders) SetCommonHeaders(v map[string]*string) *PmsEventReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PmsEventReportHeaders) SetXAcsAligenieAccessToken(v string) *PmsEventReportHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PmsEventReportHeaders) SetAuthorization(v string) *PmsEventReportHeaders {
	s.Authorization = &v
	return s
}

type PmsEventReportRequest struct {
	// This parameter is required.
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s PmsEventReportRequest) String() string {
	return tea.Prettify(s)
}

func (s PmsEventReportRequest) GoString() string {
	return s.String()
}

func (s *PmsEventReportRequest) SetPayload(v string) *PmsEventReportRequest {
	s.Payload = &v
	return s
}

type PmsEventReportResponseBody struct {
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PmsEventReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PmsEventReportResponseBody) GoString() string {
	return s.String()
}

func (s *PmsEventReportResponseBody) SetMessage(v string) *PmsEventReportResponseBody {
	s.Message = &v
	return s
}

func (s *PmsEventReportResponseBody) SetRequestId(v string) *PmsEventReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *PmsEventReportResponseBody) SetResult(v bool) *PmsEventReportResponseBody {
	s.Result = &v
	return s
}

func (s *PmsEventReportResponseBody) SetStatusCode(v int32) *PmsEventReportResponseBody {
	s.StatusCode = &v
	return s
}

type PmsEventReportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PmsEventReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PmsEventReportResponse) String() string {
	return tea.Prettify(s)
}

func (s PmsEventReportResponse) GoString() string {
	return s.String()
}

func (s *PmsEventReportResponse) SetHeaders(v map[string]*string) *PmsEventReportResponse {
	s.Headers = v
	return s
}

func (s *PmsEventReportResponse) SetStatusCode(v int32) *PmsEventReportResponse {
	s.StatusCode = &v
	return s
}

func (s *PmsEventReportResponse) SetBody(v *PmsEventReportResponseBody) *PmsEventReportResponse {
	s.Body = v
	return s
}

type PushHotelMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushHotelMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushHotelMessageHeaders) GoString() string {
	return s.String()
}

func (s *PushHotelMessageHeaders) SetCommonHeaders(v map[string]*string) *PushHotelMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushHotelMessageHeaders) SetXAcsAligenieAccessToken(v string) *PushHotelMessageHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushHotelMessageHeaders) SetAuthorization(v string) *PushHotelMessageHeaders {
	s.Authorization = &v
	return s
}

type PushHotelMessageRequest struct {
	// pushHotelMessageReq
	//
	// This parameter is required.
	PushHotelMessageReq *PushHotelMessageRequestPushHotelMessageReq `json:"PushHotelMessageReq,omitempty" xml:"PushHotelMessageReq,omitempty" type:"Struct"`
}

func (s PushHotelMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushHotelMessageRequest) GoString() string {
	return s.String()
}

func (s *PushHotelMessageRequest) SetPushHotelMessageReq(v *PushHotelMessageRequestPushHotelMessageReq) *PushHotelMessageRequest {
	s.PushHotelMessageReq = v
	return s
}

type PushHotelMessageRequestPushHotelMessageReq struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId  *string            `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	ParamMap map[string]*string `json:"ParamMap,omitempty" xml:"ParamMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s PushHotelMessageRequestPushHotelMessageReq) String() string {
	return tea.Prettify(s)
}

func (s PushHotelMessageRequestPushHotelMessageReq) GoString() string {
	return s.String()
}

func (s *PushHotelMessageRequestPushHotelMessageReq) SetHotelId(v string) *PushHotelMessageRequestPushHotelMessageReq {
	s.HotelId = &v
	return s
}

func (s *PushHotelMessageRequestPushHotelMessageReq) SetParamMap(v map[string]*string) *PushHotelMessageRequestPushHotelMessageReq {
	s.ParamMap = v
	return s
}

func (s *PushHotelMessageRequestPushHotelMessageReq) SetRoomNo(v string) *PushHotelMessageRequestPushHotelMessageReq {
	s.RoomNo = &v
	return s
}

func (s *PushHotelMessageRequestPushHotelMessageReq) SetTemplateId(v int64) *PushHotelMessageRequestPushHotelMessageReq {
	s.TemplateId = &v
	return s
}

type PushHotelMessageShrinkRequest struct {
	// pushHotelMessageReq
	//
	// This parameter is required.
	PushHotelMessageReqShrink *string `json:"PushHotelMessageReq,omitempty" xml:"PushHotelMessageReq,omitempty"`
}

func (s PushHotelMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PushHotelMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushHotelMessageShrinkRequest) SetPushHotelMessageReqShrink(v string) *PushHotelMessageShrinkRequest {
	s.PushHotelMessageReqShrink = &v
	return s
}

type PushHotelMessageResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PushHotelMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushHotelMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushHotelMessageResponseBody) SetCode(v int32) *PushHotelMessageResponseBody {
	s.Code = &v
	return s
}

func (s *PushHotelMessageResponseBody) SetMessage(v string) *PushHotelMessageResponseBody {
	s.Message = &v
	return s
}

func (s *PushHotelMessageResponseBody) SetRequestId(v string) *PushHotelMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushHotelMessageResponseBody) SetResult(v bool) *PushHotelMessageResponseBody {
	s.Result = &v
	return s
}

type PushHotelMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushHotelMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushHotelMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushHotelMessageResponse) GoString() string {
	return s.String()
}

func (s *PushHotelMessageResponse) SetHeaders(v map[string]*string) *PushHotelMessageResponse {
	s.Headers = v
	return s
}

func (s *PushHotelMessageResponse) SetStatusCode(v int32) *PushHotelMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PushHotelMessageResponse) SetBody(v *PushHotelMessageResponseBody) *PushHotelMessageResponse {
	s.Body = v
	return s
}

type PushVoiceBoxCommandsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushVoiceBoxCommandsHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushVoiceBoxCommandsHeaders) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsHeaders) SetCommonHeaders(v map[string]*string) *PushVoiceBoxCommandsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushVoiceBoxCommandsHeaders) SetXAcsAligenieAccessToken(v string) *PushVoiceBoxCommandsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushVoiceBoxCommandsHeaders) SetAuthorization(v string) *PushVoiceBoxCommandsHeaders {
	s.Authorization = &v
	return s
}

type PushVoiceBoxCommandsRequest struct {
	// This parameter is required.
	Commands []*PushVoiceBoxCommandsRequestCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// This parameter is required.
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s PushVoiceBoxCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s PushVoiceBoxCommandsRequest) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsRequest) SetCommands(v []*PushVoiceBoxCommandsRequestCommands) *PushVoiceBoxCommandsRequest {
	s.Commands = v
	return s
}

func (s *PushVoiceBoxCommandsRequest) SetHotelId(v string) *PushVoiceBoxCommandsRequest {
	s.HotelId = &v
	return s
}

func (s *PushVoiceBoxCommandsRequest) SetRoomNo(v string) *PushVoiceBoxCommandsRequest {
	s.RoomNo = &v
	return s
}

type PushVoiceBoxCommandsRequestCommands struct {
	// This parameter is required.
	CommandDomain *string `json:"CommandDomain,omitempty" xml:"CommandDomain,omitempty"`
	// This parameter is required.
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	Payload     *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s PushVoiceBoxCommandsRequestCommands) String() string {
	return tea.Prettify(s)
}

func (s PushVoiceBoxCommandsRequestCommands) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsRequestCommands) SetCommandDomain(v string) *PushVoiceBoxCommandsRequestCommands {
	s.CommandDomain = &v
	return s
}

func (s *PushVoiceBoxCommandsRequestCommands) SetCommandName(v string) *PushVoiceBoxCommandsRequestCommands {
	s.CommandName = &v
	return s
}

func (s *PushVoiceBoxCommandsRequestCommands) SetPayload(v string) *PushVoiceBoxCommandsRequestCommands {
	s.Payload = &v
	return s
}

type PushVoiceBoxCommandsShrinkRequest struct {
	// This parameter is required.
	CommandsShrink *string `json:"Commands,omitempty" xml:"Commands,omitempty"`
	// This parameter is required.
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s PushVoiceBoxCommandsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PushVoiceBoxCommandsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsShrinkRequest) SetCommandsShrink(v string) *PushVoiceBoxCommandsShrinkRequest {
	s.CommandsShrink = &v
	return s
}

func (s *PushVoiceBoxCommandsShrinkRequest) SetHotelId(v string) *PushVoiceBoxCommandsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *PushVoiceBoxCommandsShrinkRequest) SetRoomNo(v string) *PushVoiceBoxCommandsShrinkRequest {
	s.RoomNo = &v
	return s
}

type PushVoiceBoxCommandsResponseBody struct {
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	StatusCode *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PushVoiceBoxCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushVoiceBoxCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsResponseBody) SetCode(v int32) *PushVoiceBoxCommandsResponseBody {
	s.Code = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) SetMessage(v string) *PushVoiceBoxCommandsResponseBody {
	s.Message = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) SetRequestId(v string) *PushVoiceBoxCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) SetResult(v bool) *PushVoiceBoxCommandsResponseBody {
	s.Result = &v
	return s
}

func (s *PushVoiceBoxCommandsResponseBody) SetStatusCode(v int32) *PushVoiceBoxCommandsResponseBody {
	s.StatusCode = &v
	return s
}

type PushVoiceBoxCommandsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushVoiceBoxCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushVoiceBoxCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s PushVoiceBoxCommandsResponse) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsResponse) SetHeaders(v map[string]*string) *PushVoiceBoxCommandsResponse {
	s.Headers = v
	return s
}

func (s *PushVoiceBoxCommandsResponse) SetStatusCode(v int32) *PushVoiceBoxCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *PushVoiceBoxCommandsResponse) SetBody(v *PushVoiceBoxCommandsResponseBody) *PushVoiceBoxCommandsResponse {
	s.Body = v
	return s
}

type PushWelcomeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushWelcomeHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeHeaders) GoString() string {
	return s.String()
}

func (s *PushWelcomeHeaders) SetCommonHeaders(v map[string]*string) *PushWelcomeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushWelcomeHeaders) SetXAcsAligenieAccessToken(v string) *PushWelcomeHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushWelcomeHeaders) SetAuthorization(v string) *PushWelcomeHeaders {
	s.Authorization = &v
	return s
}

type PushWelcomeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// http://ailabsaicloudservice.alicdn.com/tmp/a.wav
	WelcomeMusicUrl *string `json:"WelcomeMusicUrl,omitempty" xml:"WelcomeMusicUrl,omitempty"`
	// This parameter is required.
	WelcomeText *string `json:"WelcomeText,omitempty" xml:"WelcomeText,omitempty"`
}

func (s PushWelcomeRequest) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeRequest) GoString() string {
	return s.String()
}

func (s *PushWelcomeRequest) SetHotelId(v string) *PushWelcomeRequest {
	s.HotelId = &v
	return s
}

func (s *PushWelcomeRequest) SetRoomNo(v string) *PushWelcomeRequest {
	s.RoomNo = &v
	return s
}

func (s *PushWelcomeRequest) SetWelcomeMusicUrl(v string) *PushWelcomeRequest {
	s.WelcomeMusicUrl = &v
	return s
}

func (s *PushWelcomeRequest) SetWelcomeText(v string) *PushWelcomeRequest {
	s.WelcomeText = &v
	return s
}

type PushWelcomeResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PushWelcomeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeResponseBody) GoString() string {
	return s.String()
}

func (s *PushWelcomeResponseBody) SetMessage(v string) *PushWelcomeResponseBody {
	s.Message = &v
	return s
}

func (s *PushWelcomeResponseBody) SetRequestId(v string) *PushWelcomeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushWelcomeResponseBody) SetResult(v bool) *PushWelcomeResponseBody {
	s.Result = &v
	return s
}

func (s *PushWelcomeResponseBody) SetStatusCode(v int32) *PushWelcomeResponseBody {
	s.StatusCode = &v
	return s
}

type PushWelcomeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushWelcomeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushWelcomeResponse) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeResponse) GoString() string {
	return s.String()
}

func (s *PushWelcomeResponse) SetHeaders(v map[string]*string) *PushWelcomeResponse {
	s.Headers = v
	return s
}

func (s *PushWelcomeResponse) SetStatusCode(v int32) *PushWelcomeResponse {
	s.StatusCode = &v
	return s
}

func (s *PushWelcomeResponse) SetBody(v *PushWelcomeResponseBody) *PushWelcomeResponse {
	s.Body = v
	return s
}

type PushWelcomeTextAndMusicHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushWelcomeTextAndMusicHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeTextAndMusicHeaders) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicHeaders) SetCommonHeaders(v map[string]*string) *PushWelcomeTextAndMusicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushWelcomeTextAndMusicHeaders) SetXAcsAligenieAccessToken(v string) *PushWelcomeTextAndMusicHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushWelcomeTextAndMusicHeaders) SetAuthorization(v string) *PushWelcomeTextAndMusicHeaders {
	s.Authorization = &v
	return s
}

type PushWelcomeTextAndMusicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo           *string            `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	TemplateVariable map[string]*string `json:"TemplateVariable,omitempty" xml:"TemplateVariable,omitempty"`
}

func (s PushWelcomeTextAndMusicRequest) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeTextAndMusicRequest) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicRequest) SetHotelId(v string) *PushWelcomeTextAndMusicRequest {
	s.HotelId = &v
	return s
}

func (s *PushWelcomeTextAndMusicRequest) SetRoomNo(v string) *PushWelcomeTextAndMusicRequest {
	s.RoomNo = &v
	return s
}

func (s *PushWelcomeTextAndMusicRequest) SetTemplateVariable(v map[string]*string) *PushWelcomeTextAndMusicRequest {
	s.TemplateVariable = v
	return s
}

type PushWelcomeTextAndMusicShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo                 *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	TemplateVariableShrink *string `json:"TemplateVariable,omitempty" xml:"TemplateVariable,omitempty"`
}

func (s PushWelcomeTextAndMusicShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeTextAndMusicShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicShrinkRequest) SetHotelId(v string) *PushWelcomeTextAndMusicShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *PushWelcomeTextAndMusicShrinkRequest) SetRoomNo(v string) *PushWelcomeTextAndMusicShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *PushWelcomeTextAndMusicShrinkRequest) SetTemplateVariableShrink(v string) *PushWelcomeTextAndMusicShrinkRequest {
	s.TemplateVariableShrink = &v
	return s
}

type PushWelcomeTextAndMusicResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F7E2****B7C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PushWelcomeTextAndMusicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeTextAndMusicResponseBody) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicResponseBody) SetExtentions(v map[string]interface{}) *PushWelcomeTextAndMusicResponseBody {
	s.Extentions = v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) SetMessage(v string) *PushWelcomeTextAndMusicResponseBody {
	s.Message = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) SetRequestId(v string) *PushWelcomeTextAndMusicResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) SetResult(v bool) *PushWelcomeTextAndMusicResponseBody {
	s.Result = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) SetStatusCode(v int32) *PushWelcomeTextAndMusicResponseBody {
	s.StatusCode = &v
	return s
}

type PushWelcomeTextAndMusicResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushWelcomeTextAndMusicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushWelcomeTextAndMusicResponse) String() string {
	return tea.Prettify(s)
}

func (s PushWelcomeTextAndMusicResponse) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicResponse) SetHeaders(v map[string]*string) *PushWelcomeTextAndMusicResponse {
	s.Headers = v
	return s
}

func (s *PushWelcomeTextAndMusicResponse) SetStatusCode(v int32) *PushWelcomeTextAndMusicResponse {
	s.StatusCode = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponse) SetBody(v *PushWelcomeTextAndMusicResponseBody) *PushWelcomeTextAndMusicResponse {
	s.Body = v
	return s
}

type QueryDeviceStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryDeviceStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceStatusHeaders) SetXAcsAligenieAccessToken(v string) *QueryDeviceStatusHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryDeviceStatusHeaders) SetAuthorization(v string) *QueryDeviceStatusHeaders {
	s.Authorization = &v
	return s
}

type QueryDeviceStatusRequest struct {
	Payload  *QueryDeviceStatusRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *QueryDeviceStatusRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s QueryDeviceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequest) SetPayload(v *QueryDeviceStatusRequestPayload) *QueryDeviceStatusRequest {
	s.Payload = v
	return s
}

func (s *QueryDeviceStatusRequest) SetUserInfo(v *QueryDeviceStatusRequestUserInfo) *QueryDeviceStatusRequest {
	s.UserInfo = v
	return s
}

type QueryDeviceStatusRequestPayload struct {
	LocationDevices []*QueryDeviceStatusRequestPayloadLocationDevices `json:"LocationDevices,omitempty" xml:"LocationDevices,omitempty" type:"Repeated"`
	Properties      map[string]*string                                `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s QueryDeviceStatusRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequestPayload) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestPayload) SetLocationDevices(v []*QueryDeviceStatusRequestPayloadLocationDevices) *QueryDeviceStatusRequestPayload {
	s.LocationDevices = v
	return s
}

func (s *QueryDeviceStatusRequestPayload) SetProperties(v map[string]*string) *QueryDeviceStatusRequestPayload {
	s.Properties = v
	return s
}

type QueryDeviceStatusRequestPayloadLocationDevices struct {
	// example:
	//
	// night_light
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// example:
	//
	// light
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// room
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s QueryDeviceStatusRequestPayloadLocationDevices) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequestPayloadLocationDevices) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetDeviceNumber(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.DeviceNumber = &v
	return s
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetDeviceType(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetLocation(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.Location = &v
	return s
}

type QueryDeviceStatusRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s QueryDeviceStatusRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequestUserInfo) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestUserInfo) SetEncodeKey(v string) *QueryDeviceStatusRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetEncodeType(v string) *QueryDeviceStatusRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetId(v string) *QueryDeviceStatusRequestUserInfo {
	s.Id = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetIdType(v string) *QueryDeviceStatusRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetOrganizationId(v string) *QueryDeviceStatusRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type QueryDeviceStatusShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s QueryDeviceStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusShrinkRequest) SetPayloadShrink(v string) *QueryDeviceStatusShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *QueryDeviceStatusShrinkRequest) SetUserInfoShrink(v string) *QueryDeviceStatusShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type QueryDeviceStatusResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// fdsgrefds
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryDeviceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBody) SetCode(v int32) *QueryDeviceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetMessage(v string) *QueryDeviceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetRequestId(v string) *QueryDeviceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetResult(v []map[string]*string) *QueryDeviceStatusResponseBody {
	s.Result = v
	return s
}

type QueryDeviceStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponse) SetHeaders(v map[string]*string) *QueryDeviceStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceStatusResponse) SetStatusCode(v int32) *QueryDeviceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceStatusResponse) SetBody(v *QueryDeviceStatusResponseBody) *QueryDeviceStatusResponse {
	s.Body = v
	return s
}

type QueryHotelRoomDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryHotelRoomDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryHotelRoomDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHotelRoomDetailHeaders) SetXAcsAligenieAccessToken(v string) *QueryHotelRoomDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryHotelRoomDetailHeaders) SetAuthorization(v string) *QueryHotelRoomDetailHeaders {
	s.Authorization = &v
	return s
}

type QueryHotelRoomDetailRequest struct {
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 38:c8:**:**:f5:22
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// 设备sn信息
	//
	// 注：若在mac uuid sn全都输入的情况下 按照输入正确的内容查询 若全输入都是正确的 则 按照 uuid > mac > sn 优先级查询
	//
	// 传入mac uuid sn其中一个 则酒店id和房间号可不传
	//
	// example:
	//
	// 280**28
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// 588***96j5WU
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryHotelRoomDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailRequest) SetHotelId(v string) *QueryHotelRoomDetailRequest {
	s.HotelId = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) SetMac(v string) *QueryHotelRoomDetailRequest {
	s.Mac = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) SetRoomNo(v string) *QueryHotelRoomDetailRequest {
	s.RoomNo = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) SetSn(v string) *QueryHotelRoomDetailRequest {
	s.Sn = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) SetUuid(v string) *QueryHotelRoomDetailRequest {
	s.Uuid = &v
	return s
}

type QueryHotelRoomDetailResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryHotelRoomDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryHotelRoomDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBody) SetMessage(v string) *QueryHotelRoomDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBody) SetRequestId(v string) *QueryHotelRoomDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBody) SetResult(v *QueryHotelRoomDetailResponseBodyResult) *QueryHotelRoomDetailResponseBody {
	s.Result = v
	return s
}

func (s *QueryHotelRoomDetailResponseBody) SetStatusCode(v int32) *QueryHotelRoomDetailResponseBody {
	s.StatusCode = &v
	return s
}

type QueryHotelRoomDetailResponseBodyResult struct {
	AuthAccounts []*QueryHotelRoomDetailResponseBodyResultAuthAccounts `json:"AuthAccounts,omitempty" xml:"AuthAccounts,omitempty" type:"Repeated"`
	// example:
	//
	// rcu
	ConnectType        *string                                              `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	CreatorAccountName *string                                              `json:"CreatorAccountName,omitempty" xml:"CreatorAccountName,omitempty"`
	DeviceInfos        []*QueryHotelRoomDetailResponseBodyResultDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// a7***83
	HotelId         *string                                                `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelName       *string                                                `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	OtherService    *QueryHotelRoomDetailResponseBodyResultOtherService    `json:"OtherService,omitempty" xml:"OtherService,omitempty" type:"Struct"`
	RoomControlInfo *QueryHotelRoomDetailResponseBodyResultRoomControlInfo `json:"RoomControlInfo,omitempty" xml:"RoomControlInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2001
	RoomNo          *string                                                `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	RoomServiceInfo *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo `json:"RoomServiceInfo,omitempty" xml:"RoomServiceInfo,omitempty" type:"Struct"`
}

func (s QueryHotelRoomDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetAuthAccounts(v []*QueryHotelRoomDetailResponseBodyResultAuthAccounts) *QueryHotelRoomDetailResponseBodyResult {
	s.AuthAccounts = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetConnectType(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.ConnectType = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetCreatorAccountName(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.CreatorAccountName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetDeviceInfos(v []*QueryHotelRoomDetailResponseBodyResultDeviceInfos) *QueryHotelRoomDetailResponseBodyResult {
	s.DeviceInfos = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetHotelId(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetHotelName(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.HotelName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetOtherService(v *QueryHotelRoomDetailResponseBodyResultOtherService) *QueryHotelRoomDetailResponseBodyResult {
	s.OtherService = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetRoomControlInfo(v *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) *QueryHotelRoomDetailResponseBodyResult {
	s.RoomControlInfo = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetRoomNo(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetRoomServiceInfo(v *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) *QueryHotelRoomDetailResponseBodyResult {
	s.RoomServiceInfo = v
	return s
}

type QueryHotelRoomDetailResponseBodyResultAuthAccounts struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 2023-01-01 12:00:00
	AuthTime *string `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultAuthAccounts) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultAuthAccounts) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultAuthAccounts) SetAccountName(v string) *QueryHotelRoomDetailResponseBodyResultAuthAccounts {
	s.AccountName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultAuthAccounts) SetAuthTime(v string) *QueryHotelRoomDetailResponseBodyResultAuthAccounts {
	s.AuthTime = &v
	return s
}

type QueryHotelRoomDetailResponseBodyResultDeviceInfos struct {
	// example:
	//
	// 2023-01-01 12:00:00
	ActiveTime *string `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// 6.1.8-RS-20230425.1806
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	// example:
	//
	// fa:03:23:58:c3:00
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1
	OnlineStatus *int32 `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// sag42dlz4qf
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// 41c95c18a0a643bcb58edf438877def5
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultDeviceInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultDeviceInfos) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetActiveTime(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.ActiveTime = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetDeviceName(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.DeviceName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetFirmwareVersion(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.FirmwareVersion = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetMac(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.Mac = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetOnlineStatus(v int32) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.OnlineStatus = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetSn(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.Sn = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetUuid(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.Uuid = &v
	return s
}

type QueryHotelRoomDetailResponseBodyResultOtherService struct {
	// example:
	//
	// false
	OpenCall *bool `json:"OpenCall,omitempty" xml:"OpenCall,omitempty"`
	// example:
	//
	// 0
	UnhandleTickets *int32 `json:"UnhandleTickets,omitempty" xml:"UnhandleTickets,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultOtherService) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultOtherService) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultOtherService) SetOpenCall(v bool) *QueryHotelRoomDetailResponseBodyResultOtherService {
	s.OpenCall = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultOtherService) SetUnhandleTickets(v int32) *QueryHotelRoomDetailResponseBodyResultOtherService {
	s.UnhandleTickets = &v
	return s
}

type QueryHotelRoomDetailResponseBodyResultRoomControlInfo struct {
	// example:
	//
	// 78
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// app
	AppName     *string                                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DeviceInfos []*QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// http://www.xxx.com
	RcuUrl *string `json:"RcuUrl,omitempty" xml:"RcuUrl,omitempty"`
	// example:
	//
	// 1170
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultRoomControlInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultRoomControlInfo) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetAppId(v int64) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.AppId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetAppName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.AppName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetDeviceInfos(v []*QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.DeviceInfos = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetRcuUrl(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.RcuUrl = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetTemplateId(v int64) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.TemplateId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetTemplateName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.TemplateName = &v
	return s
}

type QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos struct {
	// example:
	//
	// light
	CategoryEnName *string `json:"CategoryEnName,omitempty" xml:"CategoryEnName,omitempty"`
	// example:
	//
	// 3
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// rcu
	DeviceConnectType *string `json:"DeviceConnectType,omitempty" xml:"DeviceConnectType,omitempty"`
	// example:
	//
	// 4
	DeviceCount *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// example:
	//
	// readLight
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// room
	LocationEnName *string `json:"LocationEnName,omitempty" xml:"LocationEnName,omitempty"`
	// example:
	//
	// 1
	LocationId   *int64  `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// example:
	//
	// a1ueWGP6W2L
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetCategoryEnName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.CategoryEnName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetCategoryId(v int64) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.CategoryId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetCategoryName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.CategoryName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetDeviceConnectType(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.DeviceConnectType = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetDeviceCount(v int32) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.DeviceCount = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetDeviceId(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.DeviceId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetDeviceName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.DeviceName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetLocationEnName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.LocationEnName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetLocationId(v int64) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.LocationId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetLocationName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.LocationName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetProductKey(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.ProductKey = &v
	return s
}

type QueryHotelRoomDetailResponseBodyResultRoomServiceInfo struct {
	// example:
	//
	// 0
	BookServiceCnt *int32 `json:"BookServiceCnt,omitempty" xml:"BookServiceCnt,omitempty"`
	// example:
	//
	// 10
	GoodsServiceCnt *int32 `json:"GoodsServiceCnt,omitempty" xml:"GoodsServiceCnt,omitempty"`
	// example:
	//
	// 10
	RepairServiceCnt *int32 `json:"RepairServiceCnt,omitempty" xml:"RepairServiceCnt,omitempty"`
	// example:
	//
	// 12
	RoomServiceCnt *int32 `json:"RoomServiceCnt,omitempty" xml:"RoomServiceCnt,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) SetBookServiceCnt(v int32) *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	s.BookServiceCnt = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) SetGoodsServiceCnt(v int32) *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	s.GoodsServiceCnt = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) SetRepairServiceCnt(v int32) *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	s.RepairServiceCnt = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) SetRoomServiceCnt(v int32) *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	s.RoomServiceCnt = &v
	return s
}

type QueryHotelRoomDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHotelRoomDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotelRoomDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelRoomDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponse) SetHeaders(v map[string]*string) *QueryHotelRoomDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryHotelRoomDetailResponse) SetStatusCode(v int32) *QueryHotelRoomDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotelRoomDetailResponse) SetBody(v *QueryHotelRoomDetailResponseBody) *QueryHotelRoomDetailResponse {
	s.Body = v
	return s
}

type QueryRoomControlDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryRoomControlDevicesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesHeaders) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesHeaders) SetCommonHeaders(v map[string]*string) *QueryRoomControlDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRoomControlDevicesHeaders) SetXAcsAligenieAccessToken(v string) *QueryRoomControlDevicesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryRoomControlDevicesHeaders) SetAuthorization(v string) *QueryRoomControlDevicesHeaders {
	s.Authorization = &v
	return s
}

type QueryRoomControlDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s QueryRoomControlDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesRequest) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesRequest) SetHotelId(v string) *QueryRoomControlDevicesRequest {
	s.HotelId = &v
	return s
}

func (s *QueryRoomControlDevicesRequest) SetRoomNo(v string) *QueryRoomControlDevicesRequest {
	s.RoomNo = &v
	return s
}

type QueryRoomControlDevicesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// fdsgfdscvre
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryRoomControlDevicesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryRoomControlDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBody) SetCode(v int32) *QueryRoomControlDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBody) SetMessage(v string) *QueryRoomControlDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBody) SetRequestId(v string) *QueryRoomControlDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBody) SetResult(v []*QueryRoomControlDevicesResponseBodyResult) *QueryRoomControlDevicesResponseBody {
	s.Result = v
	return s
}

type QueryRoomControlDevicesResponseBodyResult struct {
	Devices []*QueryRoomControlDevicesResponseBodyResultDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// room
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
}

func (s QueryRoomControlDevicesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBodyResult) SetDevices(v []*QueryRoomControlDevicesResponseBodyResultDevices) *QueryRoomControlDevicesResponseBodyResult {
	s.Devices = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResult) SetLocation(v string) *QueryRoomControlDevicesResponseBodyResult {
	s.Location = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResult) SetLocationName(v string) *QueryRoomControlDevicesResponseBodyResult {
	s.LocationName = &v
	return s
}

type QueryRoomControlDevicesResponseBodyResultDevices struct {
	AliasList         []*string                                                          `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	ConnectType       *string                                                            `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	DN                *string                                                            `json:"DN,omitempty" xml:"DN,omitempty"`
	DeviceName        *string                                                            `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceStatus      *string                                                            `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	MultiKeySwitchExt *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt `json:"MultiKeySwitchExt,omitempty" xml:"MultiKeySwitchExt,omitempty" type:"Struct"`
	// example:
	//
	// light
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// night_light
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	PK     *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryRoomControlDevicesResponseBodyResultDevices) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBodyResultDevices) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetAliasList(v []*string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.AliasList = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetConnectType(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.ConnectType = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetDN(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.DN = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetDeviceName(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.DeviceName = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetDeviceStatus(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.DeviceStatus = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetMultiKeySwitchExt(v *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.MultiKeySwitchExt = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetName(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.Name = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetNumber(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.Number = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetPK(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.PK = &v
	return s
}

type QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt struct {
	SwitchList []*QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList `json:"SwitchList,omitempty" xml:"SwitchList,omitempty" type:"Repeated"`
}

func (s QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) SetSwitchList(v []*QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt {
	s.SwitchList = v
	return s
}

type QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList struct {
	AliasList    []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	Category     *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeviceIndex  *int32    `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	DeviceName   *string   `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceStatus *string   `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	ElementCode  *string   `json:"ElementCode,omitempty" xml:"ElementCode,omitempty"`
	Location     *string   `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetAliasList(v []*string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.AliasList = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetCategory(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Category = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceIndex(v int32) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceIndex = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceName(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceName = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceStatus(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceStatus = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetElementCode(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.ElementCode = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetLocation(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Location = &v
	return s
}

type QueryRoomControlDevicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRoomControlDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRoomControlDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesResponse) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponse) SetHeaders(v map[string]*string) *QueryRoomControlDevicesResponse {
	s.Headers = v
	return s
}

func (s *QueryRoomControlDevicesResponse) SetStatusCode(v int32) *QueryRoomControlDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRoomControlDevicesResponse) SetBody(v *QueryRoomControlDevicesResponseBody) *QueryRoomControlDevicesResponse {
	s.Body = v
	return s
}

type QueryRoomControlDevicesAndStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryRoomControlDevicesAndStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusHeaders) SetXAcsAligenieAccessToken(v string) *QueryRoomControlDevicesAndStatusHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusHeaders) SetAuthorization(v string) *QueryRoomControlDevicesAndStatusHeaders {
	s.Authorization = &v
	return s
}

type QueryRoomControlDevicesAndStatusRequest struct {
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusRequest) SetHotelId(v string) *QueryRoomControlDevicesAndStatusRequest {
	s.HotelId = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusRequest) SetRoomNo(v string) *QueryRoomControlDevicesAndStatusRequest {
	s.RoomNo = &v
	return s
}

type QueryRoomControlDevicesAndStatusResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67***6FA
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryRoomControlDevicesAndStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetCode(v int32) *QueryRoomControlDevicesAndStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetMessage(v string) *QueryRoomControlDevicesAndStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetRequestId(v string) *QueryRoomControlDevicesAndStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetResult(v []*QueryRoomControlDevicesAndStatusResponseBodyResult) *QueryRoomControlDevicesAndStatusResponseBody {
	s.Result = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetStatusCode(v int32) *QueryRoomControlDevicesAndStatusResponseBody {
	s.StatusCode = &v
	return s
}

type QueryRoomControlDevicesAndStatusResponseBodyResult struct {
	Devices []*QueryRoomControlDevicesAndStatusResponseBodyResultDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// room
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) SetDevices(v []*QueryRoomControlDevicesAndStatusResponseBodyResultDevices) *QueryRoomControlDevicesAndStatusResponseBodyResult {
	s.Devices = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) SetLocation(v string) *QueryRoomControlDevicesAndStatusResponseBodyResult {
	s.Location = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) SetLocationName(v string) *QueryRoomControlDevicesAndStatusResponseBodyResult {
	s.LocationName = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) SetRoomNo(v string) *QueryRoomControlDevicesAndStatusResponseBodyResult {
	s.RoomNo = &v
	return s
}

type QueryRoomControlDevicesAndStatusResponseBodyResultDevices struct {
	AliasList []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	Brand     *string   `json:"Brand,omitempty" xml:"Brand,omitempty"`
	City      *string   `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// rcu
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	DeviceName  *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// {"powerstate": "1"}
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 3c5d***9ec
	Dn *string `json:"Dn,omitempty" xml:"Dn,omitempty"`
	// example:
	//
	// 9**7
	InfraredId *string `json:"InfraredId,omitempty" xml:"InfraredId,omitempty"`
	// example:
	//
	// 2
	InfraredIndex *string `json:"InfraredIndex,omitempty" xml:"InfraredIndex,omitempty"`
	// example:
	//
	// 3.0
	InfraredVersion   *string                                                                     `json:"InfraredVersion,omitempty" xml:"InfraredVersion,omitempty"`
	MultiKeySwitchExt *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt `json:"MultiKeySwitchExt,omitempty" xml:"MultiKeySwitchExt,omitempty" type:"Struct"`
	// example:
	//
	// light
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// night_light
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 50255129
	Pk              *string            `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Province        *string            `json:"Province,omitempty" xml:"Province,omitempty"`
	ServiceProvider *string            `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
	Status          map[string]*string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevices) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetAliasList(v []*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.AliasList = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetBrand(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Brand = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetCity(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.City = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetConnectType(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.ConnectType = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetDeviceName(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.DeviceName = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetDeviceStatus(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.DeviceStatus = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetDn(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Dn = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetInfraredId(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.InfraredId = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetInfraredIndex(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.InfraredIndex = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetInfraredVersion(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.InfraredVersion = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetMultiKeySwitchExt(v *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.MultiKeySwitchExt = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetName(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Name = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetNumber(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Number = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetPk(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Pk = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetProvince(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Province = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetServiceProvider(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.ServiceProvider = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetStatus(v map[string]*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Status = v
	return s
}

type QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt struct {
	SwitchList []*QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList `json:"SwitchList,omitempty" xml:"SwitchList,omitempty" type:"Repeated"`
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) SetSwitchList(v []*QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt {
	s.SwitchList = v
	return s
}

type QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList struct {
	AliasList []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	// example:
	//
	// light
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 1
	DeviceIndex *int32  `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	DeviceName  *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// {
	//
	//       "powerstate": "0"
	//
	// }
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// e2
	ElementCode *string `json:"ElementCode,omitempty" xml:"ElementCode,omitempty"`
	// example:
	//
	// room
	Location *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	Status   map[string]*string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags     []*string          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetAliasList(v []*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.AliasList = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetCategory(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Category = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceIndex(v int32) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceIndex = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceName(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceName = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceStatus(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceStatus = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetElementCode(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.ElementCode = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetLocation(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Location = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetStatus(v map[string]*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Status = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetTags(v []*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Tags = v
	return s
}

type QueryRoomControlDevicesAndStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRoomControlDevicesAndStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponse) SetHeaders(v map[string]*string) *QueryRoomControlDevicesAndStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponse) SetStatusCode(v int32) *QueryRoomControlDevicesAndStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponse) SetBody(v *QueryRoomControlDevicesAndStatusResponseBody) *QueryRoomControlDevicesAndStatusResponse {
	s.Body = v
	return s
}

type QuerySceneListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QuerySceneListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneListHeaders) GoString() string {
	return s.String()
}

func (s *QuerySceneListHeaders) SetCommonHeaders(v map[string]*string) *QuerySceneListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySceneListHeaders) SetXAcsAligenieAccessToken(v string) *QuerySceneListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QuerySceneListHeaders) SetAuthorization(v string) *QuerySceneListHeaders {
	s.Authorization = &v
	return s
}

type QuerySceneListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId         *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	SceneStates     []*int32  `json:"SceneStates,omitempty" xml:"SceneStates,omitempty" type:"Repeated"`
	SceneTypes      []*string `json:"SceneTypes,omitempty" xml:"SceneTypes,omitempty" type:"Repeated"`
	TemplateInfoIds []*string `json:"TemplateInfoIds,omitempty" xml:"TemplateInfoIds,omitempty" type:"Repeated"`
}

func (s QuerySceneListRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneListRequest) GoString() string {
	return s.String()
}

func (s *QuerySceneListRequest) SetHotelId(v string) *QuerySceneListRequest {
	s.HotelId = &v
	return s
}

func (s *QuerySceneListRequest) SetSceneStates(v []*int32) *QuerySceneListRequest {
	s.SceneStates = v
	return s
}

func (s *QuerySceneListRequest) SetSceneTypes(v []*string) *QuerySceneListRequest {
	s.SceneTypes = v
	return s
}

func (s *QuerySceneListRequest) SetTemplateInfoIds(v []*string) *QuerySceneListRequest {
	s.TemplateInfoIds = v
	return s
}

type QuerySceneListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId               *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	SceneStatesShrink     *string `json:"SceneStates,omitempty" xml:"SceneStates,omitempty"`
	SceneTypesShrink      *string `json:"SceneTypes,omitempty" xml:"SceneTypes,omitempty"`
	TemplateInfoIdsShrink *string `json:"TemplateInfoIds,omitempty" xml:"TemplateInfoIds,omitempty"`
}

func (s QuerySceneListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySceneListShrinkRequest) SetHotelId(v string) *QuerySceneListShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *QuerySceneListShrinkRequest) SetSceneStatesShrink(v string) *QuerySceneListShrinkRequest {
	s.SceneStatesShrink = &v
	return s
}

func (s *QuerySceneListShrinkRequest) SetSceneTypesShrink(v string) *QuerySceneListShrinkRequest {
	s.SceneTypesShrink = &v
	return s
}

func (s *QuerySceneListShrinkRequest) SetTemplateInfoIdsShrink(v string) *QuerySceneListShrinkRequest {
	s.TemplateInfoIdsShrink = &v
	return s
}

type QuerySceneListResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FAFCD152-4791-5F2F-B0BE-2DC06FD4F05B
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*QuerySceneListResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QuerySceneListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponseBody) SetMessage(v string) *QuerySceneListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySceneListResponseBody) SetRequestId(v string) *QuerySceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySceneListResponseBody) SetResults(v []*QuerySceneListResponseBodyResults) *QuerySceneListResponseBody {
	s.Results = v
	return s
}

func (s *QuerySceneListResponseBody) SetStatusCode(v int32) *QuerySceneListResponseBody {
	s.StatusCode = &v
	return s
}

type QuerySceneListResponseBodyResults struct {
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingmoshi/shuimian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 73
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// external
	SceneSource *string `json:"SceneSource,omitempty" xml:"SceneSource,omitempty"`
	// example:
	//
	// 1
	SceneState *int32 `json:"SceneState,omitempty" xml:"SceneState,omitempty"`
	// example:
	//
	// common
	SceneType           *string                                                 `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	TemplateInfoDTOList []*QuerySceneListResponseBodyResultsTemplateInfoDTOList `json:"TemplateInfoDTOList,omitempty" xml:"TemplateInfoDTOList,omitempty" type:"Repeated"`
	UnavailableReason   *string                                                 `json:"UnavailableReason,omitempty" xml:"UnavailableReason,omitempty"`
}

func (s QuerySceneListResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneListResponseBodyResults) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponseBodyResults) SetIcon(v string) *QuerySceneListResponseBodyResults {
	s.Icon = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneId(v string) *QuerySceneListResponseBodyResults {
	s.SceneId = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneName(v string) *QuerySceneListResponseBodyResults {
	s.SceneName = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneSource(v string) *QuerySceneListResponseBodyResults {
	s.SceneSource = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneState(v int32) *QuerySceneListResponseBodyResults {
	s.SceneState = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneType(v string) *QuerySceneListResponseBodyResults {
	s.SceneType = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetTemplateInfoDTOList(v []*QuerySceneListResponseBodyResultsTemplateInfoDTOList) *QuerySceneListResponseBodyResults {
	s.TemplateInfoDTOList = v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetUnavailableReason(v string) *QuerySceneListResponseBodyResults {
	s.UnavailableReason = &v
	return s
}

type QuerySceneListResponseBodyResultsTemplateInfoDTOList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 6962
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 101
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QuerySceneListResponseBodyResultsTemplateInfoDTOList) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneListResponseBodyResultsTemplateInfoDTOList) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) SetDescription(v string) *QuerySceneListResponseBodyResultsTemplateInfoDTOList {
	s.Description = &v
	return s
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) SetId(v int64) *QuerySceneListResponseBodyResultsTemplateInfoDTOList {
	s.Id = &v
	return s
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) SetName(v string) *QuerySceneListResponseBodyResultsTemplateInfoDTOList {
	s.Name = &v
	return s
}

type QuerySceneListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySceneListResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneListResponse) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponse) SetHeaders(v map[string]*string) *QuerySceneListResponse {
	s.Headers = v
	return s
}

func (s *QuerySceneListResponse) SetStatusCode(v int32) *QuerySceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySceneListResponse) SetBody(v *QuerySceneListResponseBody) *QuerySceneListResponse {
	s.Body = v
	return s
}

type RemoveChildAccountAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RemoveChildAccountAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveChildAccountAuthHeaders) GoString() string {
	return s.String()
}

func (s *RemoveChildAccountAuthHeaders) SetCommonHeaders(v map[string]*string) *RemoveChildAccountAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveChildAccountAuthHeaders) SetXAcsAligenieAccessToken(v string) *RemoveChildAccountAuthHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *RemoveChildAccountAuthHeaders) SetAuthorization(v string) *RemoveChildAccountAuthHeaders {
	s.Authorization = &v
	return s
}

type RemoveChildAccountAuthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30**53
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tbxxxx
	ChildAccountName *string `json:"ChildAccountName,omitempty" xml:"ChildAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEV***E3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s RemoveChildAccountAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveChildAccountAuthRequest) GoString() string {
	return s.String()
}

func (s *RemoveChildAccountAuthRequest) SetAppKey(v string) *RemoveChildAccountAuthRequest {
	s.AppKey = &v
	return s
}

func (s *RemoveChildAccountAuthRequest) SetChildAccountName(v string) *RemoveChildAccountAuthRequest {
	s.ChildAccountName = &v
	return s
}

func (s *RemoveChildAccountAuthRequest) SetHotelId(v string) *RemoveChildAccountAuthRequest {
	s.HotelId = &v
	return s
}

func (s *RemoveChildAccountAuthRequest) SetTbOpenId(v string) *RemoveChildAccountAuthRequest {
	s.TbOpenId = &v
	return s
}

type RemoveChildAccountAuthResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F12B***F34E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s RemoveChildAccountAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveChildAccountAuthResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveChildAccountAuthResponseBody) SetMessage(v string) *RemoveChildAccountAuthResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveChildAccountAuthResponseBody) SetRequestId(v string) *RemoveChildAccountAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveChildAccountAuthResponseBody) SetResult(v bool) *RemoveChildAccountAuthResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveChildAccountAuthResponseBody) SetStatusCode(v int32) *RemoveChildAccountAuthResponseBody {
	s.StatusCode = &v
	return s
}

type RemoveChildAccountAuthResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveChildAccountAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveChildAccountAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveChildAccountAuthResponse) GoString() string {
	return s.String()
}

func (s *RemoveChildAccountAuthResponse) SetHeaders(v map[string]*string) *RemoveChildAccountAuthResponse {
	s.Headers = v
	return s
}

func (s *RemoveChildAccountAuthResponse) SetStatusCode(v int32) *RemoveChildAccountAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveChildAccountAuthResponse) SetBody(v *RemoveChildAccountAuthResponseBody) *RemoveChildAccountAuthResponse {
	s.Body = v
	return s
}

type RemoveHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RemoveHotelHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveHotelHeaders) GoString() string {
	return s.String()
}

func (s *RemoveHotelHeaders) SetCommonHeaders(v map[string]*string) *RemoveHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveHotelHeaders) SetXAcsAligenieAccessToken(v string) *RemoveHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *RemoveHotelHeaders) SetAuthorization(v string) *RemoveHotelHeaders {
	s.Authorization = &v
	return s
}

type RemoveHotelRequest struct {
	// appkey
	//
	// This parameter is required.
	//
	// example:
	//
	// 30193305
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEVK***UE3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s RemoveHotelRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveHotelRequest) GoString() string {
	return s.String()
}

func (s *RemoveHotelRequest) SetAppKey(v string) *RemoveHotelRequest {
	s.AppKey = &v
	return s
}

func (s *RemoveHotelRequest) SetHotelId(v string) *RemoveHotelRequest {
	s.HotelId = &v
	return s
}

func (s *RemoveHotelRequest) SetTbOpenId(v string) *RemoveHotelRequest {
	s.TbOpenId = &v
	return s
}

type RemoveHotelResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s RemoveHotelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveHotelResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveHotelResponseBody) SetExtentions(v map[string]interface{}) *RemoveHotelResponseBody {
	s.Extentions = v
	return s
}

func (s *RemoveHotelResponseBody) SetMessage(v string) *RemoveHotelResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveHotelResponseBody) SetRequestId(v string) *RemoveHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveHotelResponseBody) SetResult(v bool) *RemoveHotelResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveHotelResponseBody) SetStatusCode(v int32) *RemoveHotelResponseBody {
	s.StatusCode = &v
	return s
}

type RemoveHotelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveHotelResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveHotelResponse) GoString() string {
	return s.String()
}

func (s *RemoveHotelResponse) SetHeaders(v map[string]*string) *RemoveHotelResponse {
	s.Headers = v
	return s
}

func (s *RemoveHotelResponse) SetStatusCode(v int32) *RemoveHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveHotelResponse) SetBody(v *RemoveHotelResponseBody) *RemoveHotelResponse {
	s.Body = v
	return s
}

type ResetWelcomeTextAndMusicHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ResetWelcomeTextAndMusicHeaders) String() string {
	return tea.Prettify(s)
}

func (s ResetWelcomeTextAndMusicHeaders) GoString() string {
	return s.String()
}

func (s *ResetWelcomeTextAndMusicHeaders) SetCommonHeaders(v map[string]*string) *ResetWelcomeTextAndMusicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ResetWelcomeTextAndMusicHeaders) SetXAcsAligenieAccessToken(v string) *ResetWelcomeTextAndMusicHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ResetWelcomeTextAndMusicHeaders) SetAuthorization(v string) *ResetWelcomeTextAndMusicHeaders {
	s.Authorization = &v
	return s
}

type ResetWelcomeTextAndMusicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ResetWelcomeTextAndMusicRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetWelcomeTextAndMusicRequest) GoString() string {
	return s.String()
}

func (s *ResetWelcomeTextAndMusicRequest) SetHotelId(v string) *ResetWelcomeTextAndMusicRequest {
	s.HotelId = &v
	return s
}

type ResetWelcomeTextAndMusicResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ResetWelcomeTextAndMusicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetWelcomeTextAndMusicResponseBody) GoString() string {
	return s.String()
}

func (s *ResetWelcomeTextAndMusicResponseBody) SetExtentions(v map[string]interface{}) *ResetWelcomeTextAndMusicResponseBody {
	s.Extentions = v
	return s
}

func (s *ResetWelcomeTextAndMusicResponseBody) SetMessage(v string) *ResetWelcomeTextAndMusicResponseBody {
	s.Message = &v
	return s
}

func (s *ResetWelcomeTextAndMusicResponseBody) SetRequestId(v string) *ResetWelcomeTextAndMusicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetWelcomeTextAndMusicResponseBody) SetResult(v bool) *ResetWelcomeTextAndMusicResponseBody {
	s.Result = &v
	return s
}

func (s *ResetWelcomeTextAndMusicResponseBody) SetStatusCode(v int32) *ResetWelcomeTextAndMusicResponseBody {
	s.StatusCode = &v
	return s
}

type ResetWelcomeTextAndMusicResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetWelcomeTextAndMusicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetWelcomeTextAndMusicResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetWelcomeTextAndMusicResponse) GoString() string {
	return s.String()
}

func (s *ResetWelcomeTextAndMusicResponse) SetHeaders(v map[string]*string) *ResetWelcomeTextAndMusicResponse {
	s.Headers = v
	return s
}

func (s *ResetWelcomeTextAndMusicResponse) SetStatusCode(v int32) *ResetWelcomeTextAndMusicResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetWelcomeTextAndMusicResponse) SetBody(v *ResetWelcomeTextAndMusicResponseBody) *ResetWelcomeTextAndMusicResponse {
	s.Body = v
	return s
}

type RoomCheckOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RoomCheckOutHeaders) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutHeaders) GoString() string {
	return s.String()
}

func (s *RoomCheckOutHeaders) SetCommonHeaders(v map[string]*string) *RoomCheckOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RoomCheckOutHeaders) SetXAcsAligenieAccessToken(v string) *RoomCheckOutHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *RoomCheckOutHeaders) SetAuthorization(v string) *RoomCheckOutHeaders {
	s.Authorization = &v
	return s
}

type RoomCheckOutRequest struct {
	DeviceInfo *RoomCheckOutRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo   *RoomCheckOutRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s RoomCheckOutRequest) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutRequest) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequest) SetDeviceInfo(v *RoomCheckOutRequestDeviceInfo) *RoomCheckOutRequest {
	s.DeviceInfo = v
	return s
}

func (s *RoomCheckOutRequest) SetUserInfo(v *RoomCheckOutRequestUserInfo) *RoomCheckOutRequest {
	s.UserInfo = v
	return s
}

type RoomCheckOutRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s RoomCheckOutRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequestDeviceInfo) SetEncodeKey(v string) *RoomCheckOutRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetEncodeType(v string) *RoomCheckOutRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetId(v string) *RoomCheckOutRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetIdType(v string) *RoomCheckOutRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetOrganizationId(v string) *RoomCheckOutRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type RoomCheckOutRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s RoomCheckOutRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutRequestUserInfo) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequestUserInfo) SetEncodeKey(v string) *RoomCheckOutRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetEncodeType(v string) *RoomCheckOutRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetId(v string) *RoomCheckOutRequestUserInfo {
	s.Id = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetIdType(v string) *RoomCheckOutRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetOrganizationId(v string) *RoomCheckOutRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type RoomCheckOutShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s RoomCheckOutShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutShrinkRequest) GoString() string {
	return s.String()
}

func (s *RoomCheckOutShrinkRequest) SetDeviceInfoShrink(v string) *RoomCheckOutShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *RoomCheckOutShrinkRequest) SetUserInfoShrink(v string) *RoomCheckOutShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type RoomCheckOutResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// dsvrevd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RoomCheckOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutResponseBody) GoString() string {
	return s.String()
}

func (s *RoomCheckOutResponseBody) SetCode(v int32) *RoomCheckOutResponseBody {
	s.Code = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetMessage(v string) *RoomCheckOutResponseBody {
	s.Message = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetRequestId(v string) *RoomCheckOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetResult(v bool) *RoomCheckOutResponseBody {
	s.Result = &v
	return s
}

type RoomCheckOutResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RoomCheckOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RoomCheckOutResponse) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutResponse) GoString() string {
	return s.String()
}

func (s *RoomCheckOutResponse) SetHeaders(v map[string]*string) *RoomCheckOutResponse {
	s.Headers = v
	return s
}

func (s *RoomCheckOutResponse) SetStatusCode(v int32) *RoomCheckOutResponse {
	s.StatusCode = &v
	return s
}

func (s *RoomCheckOutResponse) SetBody(v *RoomCheckOutResponseBody) *RoomCheckOutResponse {
	s.Body = v
	return s
}

type SubmitHotelOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SubmitHotelOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderHeaders) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderHeaders) SetCommonHeaders(v map[string]*string) *SubmitHotelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitHotelOrderHeaders) SetXAcsAligenieAccessToken(v string) *SubmitHotelOrderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SubmitHotelOrderHeaders) SetAuthorization(v string) *SubmitHotelOrderHeaders {
	s.Authorization = &v
	return s
}

type SubmitHotelOrderRequest struct {
	// This parameter is required.
	Payload *SubmitHotelOrderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *SubmitHotelOrderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SubmitHotelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequest) SetPayload(v *SubmitHotelOrderRequestPayload) *SubmitHotelOrderRequest {
	s.Payload = v
	return s
}

func (s *SubmitHotelOrderRequest) SetUserInfo(v *SubmitHotelOrderRequestUserInfo) *SubmitHotelOrderRequest {
	s.UserInfo = v
	return s
}

type SubmitHotelOrderRequestPayload struct {
	// This parameter is required.
	ItemList []*SubmitHotelOrderRequestPayloadItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitHotelOrderRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderRequestPayload) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestPayload) SetItemList(v []*SubmitHotelOrderRequestPayloadItemList) *SubmitHotelOrderRequestPayload {
	s.ItemList = v
	return s
}

func (s *SubmitHotelOrderRequestPayload) SetType(v string) *SubmitHotelOrderRequestPayload {
	s.Type = &v
	return s
}

type SubmitHotelOrderRequestPayloadItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// 152860
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s SubmitHotelOrderRequestPayloadItemList) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderRequestPayloadItemList) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestPayloadItemList) SetItemId(v int64) *SubmitHotelOrderRequestPayloadItemList {
	s.ItemId = &v
	return s
}

func (s *SubmitHotelOrderRequestPayloadItemList) SetQuantity(v int64) *SubmitHotelOrderRequestPayloadItemList {
	s.Quantity = &v
	return s
}

type SubmitHotelOrderRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1248494721591392955
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mFU6VtVU+pgA8lx6rYMo7SPl11t+8b+8ALrn10MIPEdpK/HI9wELAEppYhPI1cYRDa4og8AMjAEBZKbLUwFjFA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SubmitHotelOrderRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestUserInfo) SetEncodeKey(v string) *SubmitHotelOrderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetEncodeType(v string) *SubmitHotelOrderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetId(v string) *SubmitHotelOrderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetIdType(v string) *SubmitHotelOrderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetOrganizationId(v string) *SubmitHotelOrderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type SubmitHotelOrderShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SubmitHotelOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderShrinkRequest) SetPayloadShrink(v string) *SubmitHotelOrderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *SubmitHotelOrderShrinkRequest) SetUserInfoShrink(v string) *SubmitHotelOrderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type SubmitHotelOrderResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CCCF2E86-D9B5-12A6-AD25-8A06933D2B0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20220809104752000114671478353329
	Result     *string `json:"Result,omitempty" xml:"Result,omitempty"`
	StatusCode *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitHotelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderResponseBody) SetCode(v int32) *SubmitHotelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetMessage(v string) *SubmitHotelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetRequestId(v string) *SubmitHotelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetResult(v string) *SubmitHotelOrderResponseBody {
	s.Result = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetStatusCode(v int32) *SubmitHotelOrderResponseBody {
	s.StatusCode = &v
	return s
}

type SubmitHotelOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHotelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHotelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderResponse) SetHeaders(v map[string]*string) *SubmitHotelOrderResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotelOrderResponse) SetStatusCode(v int32) *SubmitHotelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotelOrderResponse) SetBody(v *SubmitHotelOrderResponseBody) *SubmitHotelOrderResponse {
	s.Body = v
	return s
}

type SyncDeviceStatusWithAkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SyncDeviceStatusWithAkHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncDeviceStatusWithAkHeaders) GoString() string {
	return s.String()
}

func (s *SyncDeviceStatusWithAkHeaders) SetCommonHeaders(v map[string]*string) *SyncDeviceStatusWithAkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDeviceStatusWithAkHeaders) SetXAcsAligenieAccessToken(v string) *SyncDeviceStatusWithAkHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SyncDeviceStatusWithAkHeaders) SetAuthorization(v string) *SyncDeviceStatusWithAkHeaders {
	s.Authorization = &v
	return s
}

type SyncDeviceStatusWithAkRequest struct {
	CategoryCnName *string `json:"CategoryCnName,omitempty" xml:"CategoryCnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// light
	CategoryEnName *string `json:"CategoryEnName,omitempty" xml:"CategoryEnName,omitempty"`
	DeviceName     *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// room
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationCnName *string `json:"LocationCnName,omitempty" xml:"LocationCnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bedLight
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Switch          *int32  `json:"Switch,omitempty" xml:"Switch,omitempty"`
	FanSpeed        *string `json:"fanSpeed,omitempty" xml:"fanSpeed,omitempty"`
	Mode            *string `json:"mode,omitempty" xml:"mode,omitempty"`
	RoomTemperature *string `json:"roomTemperature,omitempty" xml:"roomTemperature,omitempty"`
	Temperature     *string `json:"temperature,omitempty" xml:"temperature,omitempty"`
	Value           *int32  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SyncDeviceStatusWithAkRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDeviceStatusWithAkRequest) GoString() string {
	return s.String()
}

func (s *SyncDeviceStatusWithAkRequest) SetCategoryCnName(v string) *SyncDeviceStatusWithAkRequest {
	s.CategoryCnName = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetCategoryEnName(v string) *SyncDeviceStatusWithAkRequest {
	s.CategoryEnName = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetDeviceName(v string) *SyncDeviceStatusWithAkRequest {
	s.DeviceName = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetHotelId(v string) *SyncDeviceStatusWithAkRequest {
	s.HotelId = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetLocation(v string) *SyncDeviceStatusWithAkRequest {
	s.Location = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetLocationCnName(v string) *SyncDeviceStatusWithAkRequest {
	s.LocationCnName = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetNumber(v string) *SyncDeviceStatusWithAkRequest {
	s.Number = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetRoomNo(v string) *SyncDeviceStatusWithAkRequest {
	s.RoomNo = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetSwitch(v int32) *SyncDeviceStatusWithAkRequest {
	s.Switch = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetFanSpeed(v string) *SyncDeviceStatusWithAkRequest {
	s.FanSpeed = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetMode(v string) *SyncDeviceStatusWithAkRequest {
	s.Mode = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetRoomTemperature(v string) *SyncDeviceStatusWithAkRequest {
	s.RoomTemperature = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetTemperature(v string) *SyncDeviceStatusWithAkRequest {
	s.Temperature = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetValue(v int32) *SyncDeviceStatusWithAkRequest {
	s.Value = &v
	return s
}

type SyncDeviceStatusWithAkResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SyncDeviceStatusWithAkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDeviceStatusWithAkResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDeviceStatusWithAkResponseBody) SetMessage(v string) *SyncDeviceStatusWithAkResponseBody {
	s.Message = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponseBody) SetResult(v bool) *SyncDeviceStatusWithAkResponseBody {
	s.Result = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponseBody) SetStatusCode(v int32) *SyncDeviceStatusWithAkResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponseBody) SetRequestId(v string) *SyncDeviceStatusWithAkResponseBody {
	s.RequestId = &v
	return s
}

type SyncDeviceStatusWithAkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDeviceStatusWithAkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDeviceStatusWithAkResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDeviceStatusWithAkResponse) GoString() string {
	return s.String()
}

func (s *SyncDeviceStatusWithAkResponse) SetHeaders(v map[string]*string) *SyncDeviceStatusWithAkResponse {
	s.Headers = v
	return s
}

func (s *SyncDeviceStatusWithAkResponse) SetStatusCode(v int32) *SyncDeviceStatusWithAkResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponse) SetBody(v *SyncDeviceStatusWithAkResponseBody) *SyncDeviceStatusWithAkResponse {
	s.Body = v
	return s
}

type UpdateBasicInfoQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateBasicInfoQAHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicInfoQAHeaders) GoString() string {
	return s.String()
}

func (s *UpdateBasicInfoQAHeaders) SetCommonHeaders(v map[string]*string) *UpdateBasicInfoQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateBasicInfoQAHeaders) SetXAcsAligenieAccessToken(v string) *UpdateBasicInfoQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateBasicInfoQAHeaders) SetAuthorization(v string) *UpdateBasicInfoQAHeaders {
	s.Authorization = &v
	return s
}

type UpdateBasicInfoQARequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11:11
	CheckInTime *string `json:"CheckInTime,omitempty" xml:"CheckInTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11:11
	CheckOutTime *string `json:"CheckOutTime,omitempty" xml:"CheckOutTime,omitempty"`
	// This parameter is required.
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId           *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelIntroduction *string `json:"HotelIntroduction,omitempty" xml:"HotelIntroduction,omitempty"`
	HotelMember       *string `json:"HotelMember,omitempty" xml:"HotelMember,omitempty"`
	HotelService      *string `json:"HotelService,omitempty" xml:"HotelService,omitempty"`
	// This parameter is required.
	ParkingExpenses *string `json:"ParkingExpenses,omitempty" xml:"ParkingExpenses,omitempty"`
	// This parameter is required.
	ParkingPosition *string `json:"ParkingPosition,omitempty" xml:"ParkingPosition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	WifiName *string `json:"WifiName,omitempty" xml:"WifiName,omitempty"`
	// This parameter is required.
	WifiPassword *string `json:"WifiPassword,omitempty" xml:"WifiPassword,omitempty"`
}

func (s UpdateBasicInfoQARequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicInfoQARequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicInfoQARequest) SetCheckInTime(v string) *UpdateBasicInfoQARequest {
	s.CheckInTime = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetCheckOutTime(v string) *UpdateBasicInfoQARequest {
	s.CheckOutTime = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelAddress(v string) *UpdateBasicInfoQARequest {
	s.HotelAddress = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelId(v string) *UpdateBasicInfoQARequest {
	s.HotelId = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelIntroduction(v string) *UpdateBasicInfoQARequest {
	s.HotelIntroduction = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelMember(v string) *UpdateBasicInfoQARequest {
	s.HotelMember = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelService(v string) *UpdateBasicInfoQARequest {
	s.HotelService = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetParkingExpenses(v string) *UpdateBasicInfoQARequest {
	s.ParkingExpenses = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetParkingPosition(v string) *UpdateBasicInfoQARequest {
	s.ParkingPosition = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetPhoneNumber(v string) *UpdateBasicInfoQARequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetWifiName(v string) *UpdateBasicInfoQARequest {
	s.WifiName = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetWifiPassword(v string) *UpdateBasicInfoQARequest {
	s.WifiPassword = &v
	return s
}

type UpdateBasicInfoQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateBasicInfoQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicInfoQAResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicInfoQAResponseBody) SetMessage(v string) *UpdateBasicInfoQAResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBasicInfoQAResponseBody) SetRequestId(v string) *UpdateBasicInfoQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBasicInfoQAResponseBody) SetResult(v bool) *UpdateBasicInfoQAResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateBasicInfoQAResponseBody) SetStatusCode(v int32) *UpdateBasicInfoQAResponseBody {
	s.StatusCode = &v
	return s
}

type UpdateBasicInfoQAResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBasicInfoQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBasicInfoQAResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicInfoQAResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicInfoQAResponse) SetHeaders(v map[string]*string) *UpdateBasicInfoQAResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicInfoQAResponse) SetStatusCode(v int32) *UpdateBasicInfoQAResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicInfoQAResponse) SetBody(v *UpdateBasicInfoQAResponseBody) *UpdateBasicInfoQAResponse {
	s.Body = v
	return s
}

type UpdateCustomQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateCustomQAHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomQAHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCustomQAHeaders) SetCommonHeaders(v map[string]*string) *UpdateCustomQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCustomQAHeaders) SetXAcsAligenieAccessToken(v string) *UpdateCustomQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateCustomQAHeaders) SetAuthorization(v string) *UpdateCustomQAHeaders {
	s.Authorization = &v
	return s
}

type UpdateCustomQARequest struct {
	Answers []*string `json:"Answers,omitempty" xml:"Answers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomQAId *string `json:"CustomQAId,omitempty" xml:"CustomQAId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId  *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWords []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// example:
	//
	// ***
	MajorQuestion          *string   `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestions []*string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty" type:"Repeated"`
}

func (s UpdateCustomQARequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomQARequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomQARequest) SetAnswers(v []*string) *UpdateCustomQARequest {
	s.Answers = v
	return s
}

func (s *UpdateCustomQARequest) SetCustomQAId(v string) *UpdateCustomQARequest {
	s.CustomQAId = &v
	return s
}

func (s *UpdateCustomQARequest) SetHotelId(v string) *UpdateCustomQARequest {
	s.HotelId = &v
	return s
}

func (s *UpdateCustomQARequest) SetKeyWords(v []*string) *UpdateCustomQARequest {
	s.KeyWords = v
	return s
}

func (s *UpdateCustomQARequest) SetMajorQuestion(v string) *UpdateCustomQARequest {
	s.MajorQuestion = &v
	return s
}

func (s *UpdateCustomQARequest) SetSupplementaryQuestions(v []*string) *UpdateCustomQARequest {
	s.SupplementaryQuestions = v
	return s
}

type UpdateCustomQAShrinkRequest struct {
	AnswersShrink *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomQAId *string `json:"CustomQAId,omitempty" xml:"CustomQAId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId        *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWordsShrink *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// ***
	MajorQuestion                *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestionsShrink *string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty"`
}

func (s UpdateCustomQAShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomQAShrinkRequest) SetAnswersShrink(v string) *UpdateCustomQAShrinkRequest {
	s.AnswersShrink = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetCustomQAId(v string) *UpdateCustomQAShrinkRequest {
	s.CustomQAId = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetHotelId(v string) *UpdateCustomQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetKeyWordsShrink(v string) *UpdateCustomQAShrinkRequest {
	s.KeyWordsShrink = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetMajorQuestion(v string) *UpdateCustomQAShrinkRequest {
	s.MajorQuestion = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetSupplementaryQuestionsShrink(v string) *UpdateCustomQAShrinkRequest {
	s.SupplementaryQuestionsShrink = &v
	return s
}

type UpdateCustomQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***BB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateCustomQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomQAResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomQAResponseBody) SetMessage(v string) *UpdateCustomQAResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomQAResponseBody) SetRequestId(v string) *UpdateCustomQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomQAResponseBody) SetResult(v bool) *UpdateCustomQAResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateCustomQAResponseBody) SetStatusCode(v int32) *UpdateCustomQAResponseBody {
	s.StatusCode = &v
	return s
}

type UpdateCustomQAResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomQAResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomQAResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomQAResponse) SetHeaders(v map[string]*string) *UpdateCustomQAResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomQAResponse) SetStatusCode(v int32) *UpdateCustomQAResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomQAResponse) SetBody(v *UpdateCustomQAResponseBody) *UpdateCustomQAResponse {
	s.Body = v
	return s
}

type UpdateHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateHotelHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHotelHeaders) SetCommonHeaders(v map[string]*string) *UpdateHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHotelHeaders) SetXAcsAligenieAccessToken(v string) *UpdateHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateHotelHeaders) SetAuthorization(v string) *UpdateHotelHeaders {
	s.Authorization = &v
	return s
}

type UpdateHotelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 31342884
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 2022-02-22 00:00:00
	EstOpenTime  *string `json:"EstOpenTime,omitempty" xml:"EstOpenTime,omitempty"`
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// example:
	//
	// a*****@hotel.com
	HotelEmail *string `json:"HotelEmail,omitempty" xml:"HotelEmail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId   *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// example:
	//
	// 130***
	PhoneNumber *string   `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	RelatedPks  []*string `json:"RelatedPks,omitempty" xml:"RelatedPks,omitempty" type:"Repeated"`
	Remark      *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 4
	RoomNum *int32 `json:"RoomNum,omitempty" xml:"RoomNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEVK***UE3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s UpdateHotelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelRequest) SetAppKey(v string) *UpdateHotelRequest {
	s.AppKey = &v
	return s
}

func (s *UpdateHotelRequest) SetEstOpenTime(v string) *UpdateHotelRequest {
	s.EstOpenTime = &v
	return s
}

func (s *UpdateHotelRequest) SetHotelAddress(v string) *UpdateHotelRequest {
	s.HotelAddress = &v
	return s
}

func (s *UpdateHotelRequest) SetHotelEmail(v string) *UpdateHotelRequest {
	s.HotelEmail = &v
	return s
}

func (s *UpdateHotelRequest) SetHotelId(v string) *UpdateHotelRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelRequest) SetHotelName(v string) *UpdateHotelRequest {
	s.HotelName = &v
	return s
}

func (s *UpdateHotelRequest) SetPhoneNumber(v string) *UpdateHotelRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateHotelRequest) SetRelatedPks(v []*string) *UpdateHotelRequest {
	s.RelatedPks = v
	return s
}

func (s *UpdateHotelRequest) SetRemark(v string) *UpdateHotelRequest {
	s.Remark = &v
	return s
}

func (s *UpdateHotelRequest) SetRoomNum(v int32) *UpdateHotelRequest {
	s.RoomNum = &v
	return s
}

func (s *UpdateHotelRequest) SetTbOpenId(v string) *UpdateHotelRequest {
	s.TbOpenId = &v
	return s
}

type UpdateHotelShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 31342884
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 2022-02-22 00:00:00
	EstOpenTime  *string `json:"EstOpenTime,omitempty" xml:"EstOpenTime,omitempty"`
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// example:
	//
	// a*****@hotel.com
	HotelEmail *string `json:"HotelEmail,omitempty" xml:"HotelEmail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId   *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// example:
	//
	// 130***
	PhoneNumber      *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	RelatedPksShrink *string `json:"RelatedPks,omitempty" xml:"RelatedPks,omitempty"`
	Remark           *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 4
	RoomNum *int32 `json:"RoomNum,omitempty" xml:"RoomNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEVK***UE3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s UpdateHotelShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelShrinkRequest) SetAppKey(v string) *UpdateHotelShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetEstOpenTime(v string) *UpdateHotelShrinkRequest {
	s.EstOpenTime = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetHotelAddress(v string) *UpdateHotelShrinkRequest {
	s.HotelAddress = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetHotelEmail(v string) *UpdateHotelShrinkRequest {
	s.HotelEmail = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetHotelId(v string) *UpdateHotelShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetHotelName(v string) *UpdateHotelShrinkRequest {
	s.HotelName = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetPhoneNumber(v string) *UpdateHotelShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetRelatedPksShrink(v string) *UpdateHotelShrinkRequest {
	s.RelatedPksShrink = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetRemark(v string) *UpdateHotelShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetRoomNum(v int32) *UpdateHotelShrinkRequest {
	s.RoomNum = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetTbOpenId(v string) *UpdateHotelShrinkRequest {
	s.TbOpenId = &v
	return s
}

type UpdateHotelResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 46C53AEB-B19C-5C42-B32E-A726979C126F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateHotelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotelResponseBody) SetExtentions(v map[string]interface{}) *UpdateHotelResponseBody {
	s.Extentions = v
	return s
}

func (s *UpdateHotelResponseBody) SetMessage(v string) *UpdateHotelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHotelResponseBody) SetRequestId(v string) *UpdateHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotelResponseBody) SetResult(v bool) *UpdateHotelResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateHotelResponseBody) SetStatusCode(v int32) *UpdateHotelResponseBody {
	s.StatusCode = &v
	return s
}

type UpdateHotelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotelResponse) SetHeaders(v map[string]*string) *UpdateHotelResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotelResponse) SetStatusCode(v int32) *UpdateHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelResponse) SetBody(v *UpdateHotelResponseBody) *UpdateHotelResponse {
	s.Body = v
	return s
}

type UpdateHotelAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateHotelAlarmHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmHeaders) SetCommonHeaders(v map[string]*string) *UpdateHotelAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHotelAlarmHeaders) SetXAcsAligenieAccessToken(v string) *UpdateHotelAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateHotelAlarmHeaders) SetAuthorization(v string) *UpdateHotelAlarmHeaders {
	s.Authorization = &v
	return s
}

type UpdateHotelAlarmRequest struct {
	// This parameter is required.
	Alarms []*UpdateHotelAlarmRequestAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7a381a668bc485980bed3876a75e013
	HotelId      *string                              `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	ScheduleInfo *UpdateHotelAlarmRequestScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
}

func (s UpdateHotelAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequest) SetAlarms(v []*UpdateHotelAlarmRequestAlarms) *UpdateHotelAlarmRequest {
	s.Alarms = v
	return s
}

func (s *UpdateHotelAlarmRequest) SetHotelId(v string) *UpdateHotelAlarmRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelAlarmRequest) SetScheduleInfo(v *UpdateHotelAlarmRequestScheduleInfo) *UpdateHotelAlarmRequest {
	s.ScheduleInfo = v
	return s
}

type UpdateHotelAlarmRequestAlarms struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pvk***VTA==
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	// example:
	//
	// 101
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mgw/***dHQd
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s UpdateHotelAlarmRequestAlarms) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmRequestAlarms) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequestAlarms) SetAlarmId(v int64) *UpdateHotelAlarmRequestAlarms {
	s.AlarmId = &v
	return s
}

func (s *UpdateHotelAlarmRequestAlarms) SetDeviceOpenId(v string) *UpdateHotelAlarmRequestAlarms {
	s.DeviceOpenId = &v
	return s
}

func (s *UpdateHotelAlarmRequestAlarms) SetRoomNo(v string) *UpdateHotelAlarmRequestAlarms {
	s.RoomNo = &v
	return s
}

func (s *UpdateHotelAlarmRequestAlarms) SetUserOpenId(v string) *UpdateHotelAlarmRequestAlarms {
	s.UserOpenId = &v
	return s
}

type UpdateHotelAlarmRequestScheduleInfo struct {
	Once *UpdateHotelAlarmRequestScheduleInfoOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// ONCE, WEEKLY
	//
	// example:
	//
	// ONCE
	Type   *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly *UpdateHotelAlarmRequestScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s UpdateHotelAlarmRequestScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmRequestScheduleInfo) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequestScheduleInfo) SetOnce(v *UpdateHotelAlarmRequestScheduleInfoOnce) *UpdateHotelAlarmRequestScheduleInfo {
	s.Once = v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfo) SetType(v string) *UpdateHotelAlarmRequestScheduleInfo {
	s.Type = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfo) SetWeekly(v *UpdateHotelAlarmRequestScheduleInfoWeekly) *UpdateHotelAlarmRequestScheduleInfo {
	s.Weekly = v
	return s
}

type UpdateHotelAlarmRequestScheduleInfoOnce struct {
	// example:
	//
	// 20
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 9
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 2022
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s UpdateHotelAlarmRequestScheduleInfoOnce) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmRequestScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetDay(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetHour(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetMinute(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetMonth(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetYear(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Year = &v
	return s
}

type UpdateHotelAlarmRequestScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s UpdateHotelAlarmRequestScheduleInfoWeekly) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmRequestScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *UpdateHotelAlarmRequestScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) SetHour(v int32) *UpdateHotelAlarmRequestScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) SetMinute(v int32) *UpdateHotelAlarmRequestScheduleInfoWeekly {
	s.Minute = &v
	return s
}

type UpdateHotelAlarmShrinkRequest struct {
	// This parameter is required.
	AlarmsShrink *string `json:"Alarms,omitempty" xml:"Alarms,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7a381a668bc485980bed3876a75e013
	HotelId            *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	ScheduleInfoShrink *string `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty"`
}

func (s UpdateHotelAlarmShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmShrinkRequest) SetAlarmsShrink(v string) *UpdateHotelAlarmShrinkRequest {
	s.AlarmsShrink = &v
	return s
}

func (s *UpdateHotelAlarmShrinkRequest) SetHotelId(v string) *UpdateHotelAlarmShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelAlarmShrinkRequest) SetScheduleInfoShrink(v string) *UpdateHotelAlarmShrinkRequest {
	s.ScheduleInfoShrink = &v
	return s
}

type UpdateHotelAlarmResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateHotelAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmResponseBody) SetExtentions(v map[string]interface{}) *UpdateHotelAlarmResponseBody {
	s.Extentions = v
	return s
}

func (s *UpdateHotelAlarmResponseBody) SetMessage(v string) *UpdateHotelAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHotelAlarmResponseBody) SetRequestId(v string) *UpdateHotelAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotelAlarmResponseBody) SetResult(v int32) *UpdateHotelAlarmResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateHotelAlarmResponseBody) SetStatusCode(v int32) *UpdateHotelAlarmResponseBody {
	s.StatusCode = &v
	return s
}

type UpdateHotelAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotelAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotelAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelAlarmResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmResponse) SetHeaders(v map[string]*string) *UpdateHotelAlarmResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotelAlarmResponse) SetStatusCode(v int32) *UpdateHotelAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelAlarmResponse) SetBody(v *UpdateHotelAlarmResponseBody) *UpdateHotelAlarmResponse {
	s.Body = v
	return s
}

type UpdateHotelSceneBookItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateHotelSceneBookItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneBookItemHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemHeaders) SetCommonHeaders(v map[string]*string) *UpdateHotelSceneBookItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHotelSceneBookItemHeaders) SetXAcsAligenieAccessToken(v string) *UpdateHotelSceneBookItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateHotelSceneBookItemHeaders) SetAuthorization(v string) *UpdateHotelSceneBookItemHeaders {
	s.Authorization = &v
	return s
}

type UpdateHotelSceneBookItemRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// updateHotelSceneBookReq
	//
	// This parameter is required.
	UpdateHotelSceneBookReq *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq `json:"UpdateHotelSceneBookReq,omitempty" xml:"UpdateHotelSceneBookReq,omitempty" type:"Struct"`
}

func (s UpdateHotelSceneBookItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneBookItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemRequest) SetHotelId(v string) *UpdateHotelSceneBookItemRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequest) SetUpdateHotelSceneBookReq(v *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) *UpdateHotelSceneBookItemRequest {
	s.UpdateHotelSceneBookReq = v
	return s
}

type UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq struct {
	// icon
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/28d7a91e3c05db3855725fc39e0387e7/welcome_audios/aa918294b6ca3aa115c51135bf9b80cb/l9f996sq.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 青椒肉丝
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1250
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) SetIcon(v string) *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	s.Icon = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) SetId(v int64) *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	s.Id = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) SetName(v string) *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	s.Name = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) SetPrice(v int64) *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	s.Price = &v
	return s
}

type UpdateHotelSceneBookItemShrinkRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// updateHotelSceneBookReq
	//
	// This parameter is required.
	UpdateHotelSceneBookReqShrink *string `json:"UpdateHotelSceneBookReq,omitempty" xml:"UpdateHotelSceneBookReq,omitempty"`
}

func (s UpdateHotelSceneBookItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneBookItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemShrinkRequest) SetHotelId(v string) *UpdateHotelSceneBookItemShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelSceneBookItemShrinkRequest) SetUpdateHotelSceneBookReqShrink(v string) *UpdateHotelSceneBookItemShrinkRequest {
	s.UpdateHotelSceneBookReqShrink = &v
	return s
}

type UpdateHotelSceneBookItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateHotelSceneBookItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneBookItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemResponseBody) SetCode(v int32) *UpdateHotelSceneBookItemResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponseBody) SetMessage(v string) *UpdateHotelSceneBookItemResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponseBody) SetRequestId(v string) *UpdateHotelSceneBookItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponseBody) SetResult(v bool) *UpdateHotelSceneBookItemResponseBody {
	s.Result = &v
	return s
}

type UpdateHotelSceneBookItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotelSceneBookItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotelSceneBookItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneBookItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemResponse) SetHeaders(v map[string]*string) *UpdateHotelSceneBookItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotelSceneBookItemResponse) SetStatusCode(v int32) *UpdateHotelSceneBookItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponse) SetBody(v *UpdateHotelSceneBookItemResponseBody) *UpdateHotelSceneBookItemResponse {
	s.Body = v
	return s
}

type UpdateHotelSceneItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateHotelSceneItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneItemHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemHeaders) SetCommonHeaders(v map[string]*string) *UpdateHotelSceneItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHotelSceneItemHeaders) SetXAcsAligenieAccessToken(v string) *UpdateHotelSceneItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateHotelSceneItemHeaders) SetAuthorization(v string) *UpdateHotelSceneItemHeaders {
	s.Authorization = &v
	return s
}

type UpdateHotelSceneItemRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// updateHotelSceneReq
	//
	// This parameter is required.
	UpdateHotelSceneOperateReq *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq `json:"UpdateHotelSceneOperateReq,omitempty" xml:"UpdateHotelSceneOperateReq,omitempty" type:"Struct"`
	// UpdateHotelSceneReq
	//
	// This parameter is required.
	UpdateHotelSceneReq *UpdateHotelSceneItemRequestUpdateHotelSceneReq `json:"UpdateHotelSceneReq,omitempty" xml:"UpdateHotelSceneReq,omitempty" type:"Struct"`
}

func (s UpdateHotelSceneItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemRequest) SetHotelId(v string) *UpdateHotelSceneItemRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelSceneItemRequest) SetUpdateHotelSceneOperateReq(v *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) *UpdateHotelSceneItemRequest {
	s.UpdateHotelSceneOperateReq = v
	return s
}

func (s *UpdateHotelSceneItemRequest) SetUpdateHotelSceneReq(v *UpdateHotelSceneItemRequestUpdateHotelSceneReq) *UpdateHotelSceneItemRequest {
	s.UpdateHotelSceneReq = v
	return s
}

type UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	IsUseTemplateAnswer *bool `json:"IsUseTemplateAnswer,omitempty" xml:"IsUseTemplateAnswer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) SetIsUseTemplateAnswer(v bool) *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq {
	s.IsUseTemplateAnswer = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) SetOperateType(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq {
	s.OperateType = &v
	return s
}

type UpdateHotelSceneItemRequestUpdateHotelSceneReq struct {
	BeyondLimitReply *string `json:"BeyondLimitReply,omitempty" xml:"BeyondLimitReply,omitempty"`
	DeliveryMethod   *string `json:"DeliveryMethod,omitempty" xml:"DeliveryMethod,omitempty"`
	// This parameter is required.
	DialogueList []*UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList `json:"DialogueList,omitempty" xml:"DialogueList,omitempty" type:"Repeated"`
	// icon
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/wupin/keyongpinlei/mianqian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// itemID
	//
	// example:
	//
	// 10337
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LimitNumber   *int64  `json:"LimitNumber,omitempty" xml:"LimitNumber,omitempty"`
	LimitSwitch   *int32  `json:"LimitSwitch,omitempty" xml:"LimitSwitch,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 165
	Price     *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneReq) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneReq) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetBeyondLimitReply(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.BeyondLimitReply = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetDeliveryMethod(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.DeliveryMethod = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetDialogueList(v []*UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.DialogueList = v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetIcon(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Icon = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetId(v int64) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Id = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetLimitNumber(v int64) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.LimitNumber = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetLimitSwitch(v int32) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.LimitSwitch = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetName(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Name = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetPaymentMethod(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.PaymentMethod = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetPrice(v int64) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Price = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetRobotName(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.RobotName = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetStatus(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Status = &v
	return s
}

type UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList struct {
	// example:
	//
	// 335
	DialogueId *string `json:"DialogueId,omitempty" xml:"DialogueId,omitempty"`
	// example:
	//
	// 对不起，暂时不提供此物品
	NoAnswer *string `json:"NoAnswer,omitempty" xml:"NoAnswer,omitempty"`
	// example:
	//
	// 4
	NoAnswerTemplate *string `json:"NoAnswerTemplate,omitempty" xml:"NoAnswerTemplate,omitempty"`
	// example:
	//
	// 0
	Process  *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// itemId
	//
	// example:
	//
	// 10337
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 纸巾1.5元，请问需要么？
	YesAnswer *string `json:"YesAnswer,omitempty" xml:"YesAnswer,omitempty"`
	// example:
	//
	// 4
	YesAnswerTemplate *string `json:"YesAnswerTemplate,omitempty" xml:"YesAnswerTemplate,omitempty"`
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetDialogueId(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.DialogueId = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetNoAnswer(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.NoAnswer = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetNoAnswerTemplate(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.NoAnswerTemplate = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetProcess(v int32) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.Process = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetQuestion(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.Question = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetServiceId(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.ServiceId = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetYesAnswer(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.YesAnswer = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetYesAnswerTemplate(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.YesAnswerTemplate = &v
	return s
}

type UpdateHotelSceneItemShrinkRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// updateHotelSceneReq
	//
	// This parameter is required.
	UpdateHotelSceneOperateReqShrink *string `json:"UpdateHotelSceneOperateReq,omitempty" xml:"UpdateHotelSceneOperateReq,omitempty"`
	// UpdateHotelSceneReq
	//
	// This parameter is required.
	UpdateHotelSceneReqShrink *string `json:"UpdateHotelSceneReq,omitempty" xml:"UpdateHotelSceneReq,omitempty"`
}

func (s UpdateHotelSceneItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemShrinkRequest) SetHotelId(v string) *UpdateHotelSceneItemShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelSceneItemShrinkRequest) SetUpdateHotelSceneOperateReqShrink(v string) *UpdateHotelSceneItemShrinkRequest {
	s.UpdateHotelSceneOperateReqShrink = &v
	return s
}

func (s *UpdateHotelSceneItemShrinkRequest) SetUpdateHotelSceneReqShrink(v string) *UpdateHotelSceneItemShrinkRequest {
	s.UpdateHotelSceneReqShrink = &v
	return s
}

type UpdateHotelSceneItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0D0C***67DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateHotelSceneItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemResponseBody) SetCode(v int32) *UpdateHotelSceneItemResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHotelSceneItemResponseBody) SetMessage(v string) *UpdateHotelSceneItemResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHotelSceneItemResponseBody) SetRequestId(v string) *UpdateHotelSceneItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotelSceneItemResponseBody) SetResult(v bool) *UpdateHotelSceneItemResponseBody {
	s.Result = &v
	return s
}

type UpdateHotelSceneItemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotelSceneItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotelSceneItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotelSceneItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemResponse) SetHeaders(v map[string]*string) *UpdateHotelSceneItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotelSceneItemResponse) SetStatusCode(v int32) *UpdateHotelSceneItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelSceneItemResponse) SetBody(v *UpdateHotelSceneItemResponseBody) *UpdateHotelSceneItemResponse {
	s.Body = v
	return s
}

type UpdateMessageTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateMessageTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageTemplateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMessageTemplateHeaders) SetCommonHeaders(v map[string]*string) *UpdateMessageTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMessageTemplateHeaders) SetXAcsAligenieAccessToken(v string) *UpdateMessageTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateMessageTemplateHeaders) SetAuthorization(v string) *UpdateMessageTemplateHeaders {
	s.Authorization = &v
	return s
}

type UpdateMessageTemplateRequest struct {
	TemplateDetail *string `json:"TemplateDetail,omitempty" xml:"TemplateDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123123
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateMessageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageTemplateRequest) SetTemplateDetail(v string) *UpdateMessageTemplateRequest {
	s.TemplateDetail = &v
	return s
}

func (s *UpdateMessageTemplateRequest) SetTemplateId(v int64) *UpdateMessageTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateMessageTemplateRequest) SetTemplateName(v string) *UpdateMessageTemplateRequest {
	s.TemplateName = &v
	return s
}

type UpdateMessageTemplateResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateMessageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageTemplateResponseBody) SetMessage(v string) *UpdateMessageTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMessageTemplateResponseBody) SetRequestId(v string) *UpdateMessageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMessageTemplateResponseBody) SetResult(v bool) *UpdateMessageTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateMessageTemplateResponseBody) SetStatusCode(v int32) *UpdateMessageTemplateResponseBody {
	s.StatusCode = &v
	return s
}

type UpdateMessageTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageTemplateResponse) SetHeaders(v map[string]*string) *UpdateMessageTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageTemplateResponse) SetStatusCode(v int32) *UpdateMessageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageTemplateResponse) SetBody(v *UpdateMessageTemplateResponseBody) *UpdateMessageTemplateResponse {
	s.Body = v
	return s
}

type UpdateRcuSceneHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateRcuSceneHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRcuSceneHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneHeaders) SetCommonHeaders(v map[string]*string) *UpdateRcuSceneHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRcuSceneHeaders) SetXAcsAligenieAccessToken(v string) *UpdateRcuSceneHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateRcuSceneHeaders) SetAuthorization(v string) *UpdateRcuSceneHeaders {
	s.Authorization = &v
	return s
}

type UpdateRcuSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yoga
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	SceneRelationExtDTO *UpdateRcuSceneRequestSceneRelationExtDTO `json:"SceneRelationExtDTO,omitempty" xml:"SceneRelationExtDTO,omitempty" type:"Struct"`
}

func (s UpdateRcuSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRcuSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneRequest) SetHotelId(v string) *UpdateRcuSceneRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateRcuSceneRequest) SetSceneId(v string) *UpdateRcuSceneRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateRcuSceneRequest) SetSceneRelationExtDTO(v *UpdateRcuSceneRequestSceneRelationExtDTO) *UpdateRcuSceneRequest {
	s.SceneRelationExtDTO = v
	return s
}

type UpdateRcuSceneRequestSceneRelationExtDTO struct {
	CorpusList  []*string `json:"CorpusList,omitempty" xml:"CorpusList,omitempty" type:"Repeated"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingmoshi/shuimian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateRcuSceneRequestSceneRelationExtDTO) String() string {
	return tea.Prettify(s)
}

func (s UpdateRcuSceneRequestSceneRelationExtDTO) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) SetCorpusList(v []*string) *UpdateRcuSceneRequestSceneRelationExtDTO {
	s.CorpusList = v
	return s
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) SetDescription(v string) *UpdateRcuSceneRequestSceneRelationExtDTO {
	s.Description = &v
	return s
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) SetIcon(v string) *UpdateRcuSceneRequestSceneRelationExtDTO {
	s.Icon = &v
	return s
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) SetName(v string) *UpdateRcuSceneRequestSceneRelationExtDTO {
	s.Name = &v
	return s
}

type UpdateRcuSceneShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yoga
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	SceneRelationExtDTOShrink *string `json:"SceneRelationExtDTO,omitempty" xml:"SceneRelationExtDTO,omitempty"`
}

func (s UpdateRcuSceneShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRcuSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneShrinkRequest) SetHotelId(v string) *UpdateRcuSceneShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateRcuSceneShrinkRequest) SetSceneId(v string) *UpdateRcuSceneShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateRcuSceneShrinkRequest) SetSceneRelationExtDTOShrink(v string) *UpdateRcuSceneShrinkRequest {
	s.SceneRelationExtDTOShrink = &v
	return s
}

type UpdateRcuSceneResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3A680F3A-6672-5A47-AB28-12BBCD80C679
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateRcuSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRcuSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneResponseBody) SetMessage(v string) *UpdateRcuSceneResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRcuSceneResponseBody) SetRequestId(v string) *UpdateRcuSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRcuSceneResponseBody) SetResult(v bool) *UpdateRcuSceneResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateRcuSceneResponseBody) SetStatusCode(v int32) *UpdateRcuSceneResponseBody {
	s.StatusCode = &v
	return s
}

type UpdateRcuSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRcuSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRcuSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRcuSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneResponse) SetHeaders(v map[string]*string) *UpdateRcuSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateRcuSceneResponse) SetStatusCode(v int32) *UpdateRcuSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRcuSceneResponse) SetBody(v *UpdateRcuSceneResponseBody) *UpdateRcuSceneResponse {
	s.Body = v
	return s
}

type UpdateServiceQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateServiceQAHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceQAHeaders) GoString() string {
	return s.String()
}

func (s *UpdateServiceQAHeaders) SetCommonHeaders(v map[string]*string) *UpdateServiceQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateServiceQAHeaders) SetXAcsAligenieAccessToken(v string) *UpdateServiceQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateServiceQAHeaders) SetAuthorization(v string) *UpdateServiceQAHeaders {
	s.Authorization = &v
	return s
}

type UpdateServiceQARequest struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 1
	ServiceQAId *int64 `json:"ServiceQAId,omitempty" xml:"ServiceQAId,omitempty"`
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
}

func (s UpdateServiceQARequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceQARequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceQARequest) SetAnswer(v string) *UpdateServiceQARequest {
	s.Answer = &v
	return s
}

func (s *UpdateServiceQARequest) SetHotelId(v string) *UpdateServiceQARequest {
	s.HotelId = &v
	return s
}

func (s *UpdateServiceQARequest) SetServiceQAId(v int64) *UpdateServiceQARequest {
	s.ServiceQAId = &v
	return s
}

func (s *UpdateServiceQARequest) SetIsActive(v bool) *UpdateServiceQARequest {
	s.IsActive = &v
	return s
}

type UpdateServiceQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67***6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateServiceQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceQAResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceQAResponseBody) SetMessage(v string) *UpdateServiceQAResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceQAResponseBody) SetRequestId(v string) *UpdateServiceQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceQAResponseBody) SetResult(v bool) *UpdateServiceQAResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateServiceQAResponseBody) SetStatusCode(v int32) *UpdateServiceQAResponseBody {
	s.StatusCode = &v
	return s
}

type UpdateServiceQAResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceQAResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceQAResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceQAResponse) SetHeaders(v map[string]*string) *UpdateServiceQAResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceQAResponse) SetStatusCode(v int32) *UpdateServiceQAResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceQAResponse) SetBody(v *UpdateServiceQAResponseBody) *UpdateServiceQAResponse {
	s.Body = v
	return s
}

type UpdateTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTicketHeaders) SetCommonHeaders(v map[string]*string) *UpdateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTicketHeaders) SetXAcsAligenieAccessToken(v string) *UpdateTicketHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateTicketHeaders) SetAuthorization(v string) *UpdateTicketHeaders {
	s.Authorization = &v
	return s
}

type UpdateTicketRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2023***93975
	GroupKey *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequest) SetGroupKey(v string) *UpdateTicketRequest {
	s.GroupKey = &v
	return s
}

func (s *UpdateTicketRequest) SetHotelId(v string) *UpdateTicketRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateTicketRequest) SetStatus(v string) *UpdateTicketRequest {
	s.Status = &v
	return s
}

type UpdateTicketResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponseBody) SetMessage(v string) *UpdateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTicketResponseBody) SetRequestId(v string) *UpdateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketResponseBody) SetResult(v bool) *UpdateTicketResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateTicketResponseBody) SetStatusCode(v int32) *UpdateTicketResponseBody {
	s.StatusCode = &v
	return s
}

type UpdateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponse) SetHeaders(v map[string]*string) *UpdateTicketResponse {
	s.Headers = v
	return s
}

func (s *UpdateTicketResponse) SetStatusCode(v int32) *UpdateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTicketResponse) SetBody(v *UpdateTicketResponseBody) *UpdateTicketResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aligenie"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加动画
//
// @param request - AddCartoonRequest
//
// @param headers - AddCartoonHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCartoonResponse
func (client *Client) AddCartoonWithOptions(request *AddCartoonRequest, headers *AddCartoonHeaders, runtime *util.RuntimeOptions) (_result *AddCartoonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.StartVideoMd5)) {
		body["StartVideoMd5"] = request.StartVideoMd5
	}

	if !tea.BoolValue(util.IsUnset(request.StartVideoUrl)) {
		body["StartVideoUrl"] = request.StartVideoUrl
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCartoon"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/addCartoon"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCartoonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加动画
//
// @param request - AddCartoonRequest
//
// @return AddCartoonResponse
func (client *Client) AddCartoon(request *AddCartoonRequest) (_result *AddCartoonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCartoonHeaders{}
	_result = &AddCartoonResponse{}
	_body, _err := client.AddCartoonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增自定义问答
//
// @param tmpReq - AddCustomQARequest
//
// @param headers - AddCustomQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomQAResponse
func (client *Client) AddCustomQAWithOptions(tmpReq *AddCustomQARequest, headers *AddCustomQAHeaders, runtime *util.RuntimeOptions) (_result *AddCustomQAResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddCustomQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Answers)) {
		request.AnswersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Answers, tea.String("Answers"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.KeyWords)) {
		request.KeyWordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KeyWords, tea.String("KeyWords"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SupplementaryQuestions)) {
		request.SupplementaryQuestionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupplementaryQuestions, tea.String("SupplementaryQuestions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnswersShrink)) {
		body["Answers"] = request.AnswersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWordsShrink)) {
		body["KeyWords"] = request.KeyWordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MajorQuestion)) {
		body["MajorQuestion"] = request.MajorQuestion
	}

	if !tea.BoolValue(util.IsUnset(request.SupplementaryQuestionsShrink)) {
		body["SupplementaryQuestions"] = request.SupplementaryQuestionsShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCustomQA"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/addCustomQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCustomQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增自定义问答
//
// @param request - AddCustomQARequest
//
// @return AddCustomQAResponse
func (client *Client) AddCustomQA(request *AddCustomQARequest) (_result *AddCustomQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCustomQAHeaders{}
	_result = &AddCustomQAResponse{}
	_body, _err := client.AddCustomQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加问答V2
//
// @param tmpReq - AddCustomQAV2Request
//
// @param headers - AddCustomQAV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomQAV2Response
func (client *Client) AddCustomQAV2WithOptions(tmpReq *AddCustomQAV2Request, headers *AddCustomQAV2Headers, runtime *util.RuntimeOptions) (_result *AddCustomQAV2Response, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddCustomQAV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Answers)) {
		request.AnswersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Answers, tea.String("Answers"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.KeyWords)) {
		request.KeyWordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KeyWords, tea.String("KeyWords"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SupplementaryQuestions)) {
		request.SupplementaryQuestionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupplementaryQuestions, tea.String("SupplementaryQuestions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnswersShrink)) {
		body["Answers"] = request.AnswersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWordsShrink)) {
		body["KeyWords"] = request.KeyWordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MajorQuestion)) {
		body["MajorQuestion"] = request.MajorQuestion
	}

	if !tea.BoolValue(util.IsUnset(request.SupplementaryQuestionsShrink)) {
		body["SupplementaryQuestions"] = request.SupplementaryQuestionsShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCustomQAV2"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/addQAV2"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCustomQAV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加问答V2
//
// @param request - AddCustomQAV2Request
//
// @return AddCustomQAV2Response
func (client *Client) AddCustomQAV2(request *AddCustomQAV2Request) (_result *AddCustomQAV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCustomQAV2Headers{}
	_result = &AddCustomQAV2Response{}
	_body, _err := client.AddCustomQAV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加消息模板
//
// @param request - AddMessageTemplateRequest
//
// @param headers - AddMessageTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMessageTemplateResponse
func (client *Client) AddMessageTemplateWithOptions(request *AddMessageTemplateRequest, headers *AddMessageTemplateHeaders, runtime *util.RuntimeOptions) (_result *AddMessageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateDetail)) {
		body["TemplateDetail"] = request.TemplateDetail
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["TemplateName"] = request.TemplateName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMessageTemplate"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/addMessageTemplate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMessageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加消息模板
//
// @param request - AddMessageTemplateRequest
//
// @return AddMessageTemplateResponse
func (client *Client) AddMessageTemplate(request *AddMessageTemplateRequest) (_result *AddMessageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddMessageTemplateHeaders{}
	_result = &AddMessageTemplateResponse{}
	_body, _err := client.AddMessageTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或者编辑带屏展示模式
//
// @param tmpReq - AddOrUpdateDisPlayModesRequest
//
// @param headers - AddOrUpdateDisPlayModesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateDisPlayModesResponse
func (client *Client) AddOrUpdateDisPlayModesWithOptions(tmpReq *AddOrUpdateDisPlayModesRequest, headers *AddOrUpdateDisPlayModesHeaders, runtime *util.RuntimeOptions) (_result *AddOrUpdateDisPlayModesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddOrUpdateDisPlayModesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HotelDeviceModeList)) {
		request.HotelDeviceModeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelDeviceModeList, tea.String("HotelDeviceModeList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelDeviceModeListShrink)) {
		body["HotelDeviceModeList"] = request.HotelDeviceModeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddOrUpdateDisPlayModes"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/addOrUpdateDisPlayModes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOrUpdateDisPlayModesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或者编辑带屏展示模式
//
// @param request - AddOrUpdateDisPlayModesRequest
//
// @return AddOrUpdateDisPlayModesResponse
func (client *Client) AddOrUpdateDisPlayModes(request *AddOrUpdateDisPlayModesRequest) (_result *AddOrUpdateDisPlayModesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOrUpdateDisPlayModesHeaders{}
	_result = &AddOrUpdateDisPlayModesResponse{}
	_body, _err := client.AddOrUpdateDisPlayModesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或者编辑定制配置
//
// @param tmpReq - AddOrUpdateHotelSettingRequest
//
// @param headers - AddOrUpdateHotelSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateHotelSettingResponse
func (client *Client) AddOrUpdateHotelSettingWithOptions(tmpReq *AddOrUpdateHotelSettingRequest, headers *AddOrUpdateHotelSettingHeaders, runtime *util.RuntimeOptions) (_result *AddOrUpdateHotelSettingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddOrUpdateHotelSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HotelDeviceModeList)) {
		request.HotelDeviceModeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelDeviceModeList, tea.String("HotelDeviceModeList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.HotelScreenSaver)) {
		request.HotelScreenSaverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelScreenSaver, tea.String("HotelScreenSaver"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NightMode)) {
		request.NightModeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NightMode, tea.String("NightMode"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelDeviceModeListShrink)) {
		body["HotelDeviceModeList"] = request.HotelDeviceModeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.HotelScreenSaverShrink)) {
		body["HotelScreenSaver"] = request.HotelScreenSaverShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NightModeShrink)) {
		body["NightMode"] = request.NightModeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SettingType)) {
		body["SettingType"] = request.SettingType
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddOrUpdateHotelSetting"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/addOrUpdateHotelSetting"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOrUpdateHotelSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或者编辑定制配置
//
// @param request - AddOrUpdateHotelSettingRequest
//
// @return AddOrUpdateHotelSettingResponse
func (client *Client) AddOrUpdateHotelSetting(request *AddOrUpdateHotelSettingRequest) (_result *AddOrUpdateHotelSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOrUpdateHotelSettingHeaders{}
	_result = &AddOrUpdateHotelSettingResponse{}
	_body, _err := client.AddOrUpdateHotelSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或者编辑带屏屏保
//
// @param tmpReq - AddOrUpdateScreenSaverRequest
//
// @param headers - AddOrUpdateScreenSaverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateScreenSaverResponse
func (client *Client) AddOrUpdateScreenSaverWithOptions(tmpReq *AddOrUpdateScreenSaverRequest, headers *AddOrUpdateScreenSaverHeaders, runtime *util.RuntimeOptions) (_result *AddOrUpdateScreenSaverResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddOrUpdateScreenSaverShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HotelScreenSaver)) {
		request.HotelScreenSaverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelScreenSaver, tea.String("HotelScreenSaver"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.HotelScreenSaverShrink)) {
		body["HotelScreenSaver"] = request.HotelScreenSaverShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddOrUpdateScreenSaver"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/addOrUpdateScreenSaver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOrUpdateScreenSaverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或者编辑带屏屏保
//
// @param request - AddOrUpdateScreenSaverRequest
//
// @return AddOrUpdateScreenSaverResponse
func (client *Client) AddOrUpdateScreenSaver(request *AddOrUpdateScreenSaverRequest) (_result *AddOrUpdateScreenSaverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOrUpdateScreenSaverHeaders{}
	_result = &AddOrUpdateScreenSaverResponse{}
	_body, _err := client.AddOrUpdateScreenSaverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加/更新欢迎语信息
//
// @param request - AddOrUpdateWelcomeTextRequest
//
// @param headers - AddOrUpdateWelcomeTextHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateWelcomeTextResponse
func (client *Client) AddOrUpdateWelcomeTextWithOptions(request *AddOrUpdateWelcomeTextRequest, headers *AddOrUpdateWelcomeTextHeaders, runtime *util.RuntimeOptions) (_result *AddOrUpdateWelcomeTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.MusicUrl)) {
		body["MusicUrl"] = request.MusicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WelcomeText)) {
		body["WelcomeText"] = request.WelcomeText
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddOrUpdateWelcomeText"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/addOrUpdateWelcomeText"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOrUpdateWelcomeTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加/更新欢迎语信息
//
// @param request - AddOrUpdateWelcomeTextRequest
//
// @return AddOrUpdateWelcomeTextResponse
func (client *Client) AddOrUpdateWelcomeText(request *AddOrUpdateWelcomeTextRequest) (_result *AddOrUpdateWelcomeTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOrUpdateWelcomeTextHeaders{}
	_result = &AddOrUpdateWelcomeTextResponse{}
	_body, _err := client.AddOrUpdateWelcomeTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 审批酒店
//
// @param tmpReq - AuditHotelRequest
//
// @param headers - AuditHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuditHotelResponse
func (client *Client) AuditHotelWithOptions(tmpReq *AuditHotelRequest, headers *AuditHotelHeaders, runtime *util.RuntimeOptions) (_result *AuditHotelResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AuditHotelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AuditHotelReq)) {
		request.AuditHotelReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuditHotelReq, tea.String("AuditHotelReq"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditHotelReqShrink)) {
		query["AuditHotelReq"] = request.AuditHotelReqShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuditHotel"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/auditHotel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AuditHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 审批酒店
//
// @param request - AuditHotelRequest
//
// @return AuditHotelResponse
func (client *Client) AuditHotel(request *AuditHotelRequest) (_result *AuditHotelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AuditHotelHeaders{}
	_result = &AuditHotelResponse{}
	_body, _err := client.AuditHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量创建房间
//
// @param tmpReq - BatchAddHotelRoomRequest
//
// @param headers - BatchAddHotelRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddHotelRoomResponse
func (client *Client) BatchAddHotelRoomWithOptions(tmpReq *BatchAddHotelRoomRequest, headers *BatchAddHotelRoomHeaders, runtime *util.RuntimeOptions) (_result *BatchAddHotelRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchAddHotelRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoomNoList)) {
		request.RoomNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomNoList, tea.String("RoomNoList"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNoListShrink)) {
		body["RoomNoList"] = request.RoomNoListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAddHotelRoom"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/batchAddHotelRoom"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddHotelRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建房间
//
// @param request - BatchAddHotelRoomRequest
//
// @return BatchAddHotelRoomResponse
func (client *Client) BatchAddHotelRoom(request *BatchAddHotelRoomRequest) (_result *BatchAddHotelRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchAddHotelRoomHeaders{}
	_result = &BatchAddHotelRoomResponse{}
	_body, _err := client.BatchAddHotelRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除房间
//
// @param tmpReq - BatchDeleteHotelRoomRequest
//
// @param headers - BatchDeleteHotelRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteHotelRoomResponse
func (client *Client) BatchDeleteHotelRoomWithOptions(tmpReq *BatchDeleteHotelRoomRequest, headers *BatchDeleteHotelRoomHeaders, runtime *util.RuntimeOptions) (_result *BatchDeleteHotelRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchDeleteHotelRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoomNoList)) {
		request.RoomNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomNoList, tea.String("RoomNoList"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNoListShrink)) {
		body["RoomNoList"] = request.RoomNoListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteHotelRoom"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/batchDeleteHotelRoom"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteHotelRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除房间
//
// @param request - BatchDeleteHotelRoomRequest
//
// @return BatchDeleteHotelRoomResponse
func (client *Client) BatchDeleteHotelRoom(request *BatchDeleteHotelRoomRequest) (_result *BatchDeleteHotelRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchDeleteHotelRoomHeaders{}
	_result = &BatchDeleteHotelRoomResponse{}
	_body, _err := client.BatchDeleteHotelRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店退房，清楚例如闹钟等定时信息
//
// @param request - CheckoutWithAKRequest
//
// @param headers - CheckoutWithAKHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckoutWithAKResponse
func (client *Client) CheckoutWithAKWithOptions(request *CheckoutWithAKRequest, headers *CheckoutWithAKHeaders, runtime *util.RuntimeOptions) (_result *CheckoutWithAKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckoutWithAK"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/checkoutWithAK"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckoutWithAKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店退房，清楚例如闹钟等定时信息
//
// @param request - CheckoutWithAKRequest
//
// @return CheckoutWithAKResponse
func (client *Client) CheckoutWithAK(request *CheckoutWithAKRequest) (_result *CheckoutWithAKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckoutWithAKHeaders{}
	_result = &CheckoutWithAKResponse{}
	_body, _err := client.CheckoutWithAKWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 子账号授权
//
// @param request - ChildAccountAuthRequest
//
// @param headers - ChildAccountAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChildAccountAuthResponse
func (client *Client) ChildAccountAuthWithOptions(request *ChildAccountAuthRequest, headers *ChildAccountAuthHeaders, runtime *util.RuntimeOptions) (_result *ChildAccountAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.TbOpenId)) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChildAccountAuth"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/childAccountAuth"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChildAccountAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 子账号授权
//
// @param request - ChildAccountAuthRequest
//
// @return ChildAccountAuthResponse
func (client *Client) ChildAccountAuth(request *ChildAccountAuthRequest) (_result *ChildAccountAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChildAccountAuthHeaders{}
	_result = &ChildAccountAuthResponse{}
	_body, _err := client.ChildAccountAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 控制房间内设备
//
// @param tmpReq - ControlRoomDeviceRequest
//
// @param headers - ControlRoomDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ControlRoomDeviceResponse
func (client *Client) ControlRoomDeviceWithOptions(tmpReq *ControlRoomDeviceRequest, headers *ControlRoomDeviceHeaders, runtime *util.RuntimeOptions) (_result *ControlRoomDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ControlRoomDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Properties)) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, tea.String("Properties"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cmd)) {
		body["Cmd"] = request.Cmd
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIndex)) {
		body["DeviceIndex"] = request.DeviceIndex
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceNumber)) {
		body["DeviceNumber"] = request.DeviceNumber
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertiesShrink)) {
		body["Properties"] = request.PropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ControlRoomDevice"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/controlRoomDevice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ControlRoomDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 控制房间内设备
//
// @param request - ControlRoomDeviceRequest
//
// @return ControlRoomDeviceResponse
func (client *Client) ControlRoomDevice(request *ControlRoomDeviceRequest) (_result *ControlRoomDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ControlRoomDeviceHeaders{}
	_result = &ControlRoomDeviceResponse{}
	_body, _err := client.ControlRoomDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建酒店项目
//
// @param tmpReq - CreateHotelRequest
//
// @param headers - CreateHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHotelResponse
func (client *Client) CreateHotelWithOptions(tmpReq *CreateHotelRequest, headers *CreateHotelHeaders, runtime *util.RuntimeOptions) (_result *CreateHotelResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateHotelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedPks)) {
		request.RelatedPksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedPks, tea.String("RelatedPks"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.EstOpenTime)) {
		body["EstOpenTime"] = request.EstOpenTime
	}

	if !tea.BoolValue(util.IsUnset(request.HotelAddress)) {
		body["HotelAddress"] = request.HotelAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HotelEmail)) {
		body["HotelEmail"] = request.HotelEmail
	}

	if !tea.BoolValue(util.IsUnset(request.HotelName)) {
		body["HotelName"] = request.HotelName
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedPk)) {
		body["RelatedPk"] = request.RelatedPk
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedPksShrink)) {
		body["RelatedPks"] = request.RelatedPksShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNum)) {
		body["RoomNum"] = request.RoomNum
	}

	if !tea.BoolValue(util.IsUnset(request.TbOpenId)) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHotel"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/createHotel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建酒店项目
//
// @param request - CreateHotelRequest
//
// @return CreateHotelResponse
func (client *Client) CreateHotel(request *CreateHotelRequest) (_result *CreateHotelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateHotelHeaders{}
	_result = &CreateHotelResponse{}
	_body, _err := client.CreateHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量创建闹钟
//
// @param tmpReq - CreateHotelAlarmRequest
//
// @param headers - CreateHotelAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHotelAlarmResponse
func (client *Client) CreateHotelAlarmWithOptions(tmpReq *CreateHotelAlarmRequest, headers *CreateHotelAlarmHeaders, runtime *util.RuntimeOptions) (_result *CreateHotelAlarmResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateHotelAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Rooms)) {
		request.RoomsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rooms, tea.String("Rooms"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ScheduleInfo)) {
		request.ScheduleInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleInfo, tea.String("ScheduleInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.MusicType)) {
		body["MusicType"] = request.MusicType
	}

	if !tea.BoolValue(util.IsUnset(request.RoomsShrink)) {
		body["Rooms"] = request.RoomsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleInfoShrink)) {
		body["ScheduleInfo"] = request.ScheduleInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHotelAlarm"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/createHotelAlarm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHotelAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建闹钟
//
// @param request - CreateHotelAlarmRequest
//
// @return CreateHotelAlarmResponse
func (client *Client) CreateHotelAlarm(request *CreateHotelAlarmRequest) (_result *CreateHotelAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateHotelAlarmHeaders{}
	_result = &CreateHotelAlarmResponse{}
	_body, _err := client.CreateHotelAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店rcu自定义场景创建
//
// @param tmpReq - CreateRcuSceneRequest
//
// @param headers - CreateRcuSceneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRcuSceneResponse
func (client *Client) CreateRcuSceneWithOptions(tmpReq *CreateRcuSceneRequest, headers *CreateRcuSceneHeaders, runtime *util.RuntimeOptions) (_result *CreateRcuSceneResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRcuSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SceneRelationExtDTO)) {
		request.SceneRelationExtDTOShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneRelationExtDTO, tea.String("SceneRelationExtDTO"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneRelationExtDTOShrink)) {
		body["SceneRelationExtDTO"] = request.SceneRelationExtDTOShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRcuScene"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/createRcuScene"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRcuSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店rcu自定义场景创建
//
// @param request - CreateRcuSceneRequest
//
// @return CreateRcuSceneResponse
func (client *Client) CreateRcuScene(request *CreateRcuSceneRequest) (_result *CreateRcuSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateRcuSceneHeaders{}
	_result = &CreateRcuSceneResponse{}
	_body, _err := client.CreateRcuSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除动画
//
// @param request - DeleteCartoonRequest
//
// @param headers - DeleteCartoonHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCartoonResponse
func (client *Client) DeleteCartoonWithOptions(request *DeleteCartoonRequest, headers *DeleteCartoonHeaders, runtime *util.RuntimeOptions) (_result *DeleteCartoonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCartoon"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deleteCartoon"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCartoonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除动画
//
// @param request - DeleteCartoonRequest
//
// @return DeleteCartoonResponse
func (client *Client) DeleteCartoon(request *DeleteCartoonRequest) (_result *DeleteCartoonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCartoonHeaders{}
	_result = &DeleteCartoonResponse{}
	_body, _err := client.DeleteCartoonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义问答
//
// @param tmpReq - DeleteCustomQARequest
//
// @param headers - DeleteCustomQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomQAResponse
func (client *Client) DeleteCustomQAWithOptions(tmpReq *DeleteCustomQARequest, headers *DeleteCustomQAHeaders, runtime *util.RuntimeOptions) (_result *DeleteCustomQAResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteCustomQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomQAIds)) {
		request.CustomQAIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomQAIds, tea.String("CustomQAIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomQAIdsShrink)) {
		body["CustomQAIds"] = request.CustomQAIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomQA"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deleteCustomQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义问答
//
// @param request - DeleteCustomQARequest
//
// @return DeleteCustomQAResponse
func (client *Client) DeleteCustomQA(request *DeleteCustomQARequest) (_result *DeleteCustomQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCustomQAHeaders{}
	_result = &DeleteCustomQAResponse{}
	_body, _err := client.DeleteCustomQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除酒店闹钟
//
// @param tmpReq - DeleteHotelAlarmRequest
//
// @param headers - DeleteHotelAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHotelAlarmResponse
func (client *Client) DeleteHotelAlarmWithOptions(tmpReq *DeleteHotelAlarmRequest, headers *DeleteHotelAlarmHeaders, runtime *util.RuntimeOptions) (_result *DeleteHotelAlarmResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteHotelAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Alarms)) {
		request.AlarmsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Alarms, tea.String("Alarms"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmsShrink)) {
		body["Alarms"] = request.AlarmsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHotelAlarm"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deleteHotelAlarm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHotelAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除酒店闹钟
//
// @param request - DeleteHotelAlarmRequest
//
// @return DeleteHotelAlarmResponse
func (client *Client) DeleteHotelAlarm(request *DeleteHotelAlarmRequest) (_result *DeleteHotelAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteHotelAlarmHeaders{}
	_result = &DeleteHotelAlarmResponse{}
	_body, _err := client.DeleteHotelAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景预订删除
//
// @param request - DeleteHotelSceneBookItemRequest
//
// @param headers - DeleteHotelSceneBookItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHotelSceneBookItemResponse
func (client *Client) DeleteHotelSceneBookItemWithOptions(request *DeleteHotelSceneBookItemRequest, headers *DeleteHotelSceneBookItemHeaders, runtime *util.RuntimeOptions) (_result *DeleteHotelSceneBookItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHotelSceneBookItem"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deleteHotelSceneBookItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHotelSceneBookItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景预订删除
//
// @param request - DeleteHotelSceneBookItemRequest
//
// @return DeleteHotelSceneBookItemResponse
func (client *Client) DeleteHotelSceneBookItem(request *DeleteHotelSceneBookItemRequest) (_result *DeleteHotelSceneBookItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteHotelSceneBookItemHeaders{}
	_result = &DeleteHotelSceneBookItemResponse{}
	_body, _err := client.DeleteHotelSceneBookItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除定制配置
//
// @param request - DeleteHotelSettingRequest
//
// @param headers - DeleteHotelSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHotelSettingResponse
func (client *Client) DeleteHotelSettingWithOptions(request *DeleteHotelSettingRequest, headers *DeleteHotelSettingHeaders, runtime *util.RuntimeOptions) (_result *DeleteHotelSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.SettingType)) {
		body["SettingType"] = request.SettingType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHotelSetting"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deleteHotelSetting"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHotelSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除定制配置
//
// @param request - DeleteHotelSettingRequest
//
// @return DeleteHotelSettingResponse
func (client *Client) DeleteHotelSetting(request *DeleteHotelSettingRequest) (_result *DeleteHotelSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteHotelSettingHeaders{}
	_result = &DeleteHotelSettingResponse{}
	_body, _err := client.DeleteHotelSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除消息通知模板
//
// @param request - DeleteMessageTemplateRequest
//
// @param headers - DeleteMessageTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageTemplateResponse
func (client *Client) DeleteMessageTemplateWithOptions(request *DeleteMessageTemplateRequest, headers *DeleteMessageTemplateHeaders, runtime *util.RuntimeOptions) (_result *DeleteMessageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMessageTemplate"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deleteMessageTemplate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMessageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消息通知模板
//
// @param request - DeleteMessageTemplateRequest
//
// @return DeleteMessageTemplateResponse
func (client *Client) DeleteMessageTemplate(request *DeleteMessageTemplateRequest) (_result *DeleteMessageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMessageTemplateHeaders{}
	_result = &DeleteMessageTemplateResponse{}
	_body, _err := client.DeleteMessageTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除酒店自定义rcu场景
//
// @param request - DeleteRcuSceneRequest
//
// @param headers - DeleteRcuSceneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRcuSceneResponse
func (client *Client) DeleteRcuSceneWithOptions(request *DeleteRcuSceneRequest, headers *DeleteRcuSceneHeaders, runtime *util.RuntimeOptions) (_result *DeleteRcuSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRcuScene"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deleteRcuScene"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRcuSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除酒店自定义rcu场景
//
// @param request - DeleteRcuSceneRequest
//
// @return DeleteRcuSceneResponse
func (client *Client) DeleteRcuScene(request *DeleteRcuSceneRequest) (_result *DeleteRcuSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRcuSceneHeaders{}
	_result = &DeleteRcuSceneResponse{}
	_body, _err := client.DeleteRcuSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备控制
//
// @param tmpReq - DeviceControlRequest
//
// @param headers - DeviceControlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeviceControlResponse
func (client *Client) DeviceControlWithOptions(tmpReq *DeviceControlRequest, headers *DeviceControlHeaders, runtime *util.RuntimeOptions) (_result *DeviceControlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeviceControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeviceControl"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deviceControl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备控制
//
// @param request - DeviceControlRequest
//
// @return DeviceControlResponse
func (client *Client) DeviceControl(request *DeviceControlRequest) (_result *DeviceControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeviceControlHeaders{}
	_result = &DeviceControlResponse{}
	_body, _err := client.DeviceControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取基础信息问答
//
// @param request - GetBasicInfoQARequest
//
// @param headers - GetBasicInfoQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicInfoQAResponse
func (client *Client) GetBasicInfoQAWithOptions(request *GetBasicInfoQARequest, headers *GetBasicInfoQAHeaders, runtime *util.RuntimeOptions) (_result *GetBasicInfoQAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBasicInfoQA"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getBasicInfoQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicInfoQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取基础信息问答
//
// @param request - GetBasicInfoQARequest
//
// @return GetBasicInfoQAResponse
func (client *Client) GetBasicInfoQA(request *GetBasicInfoQARequest) (_result *GetBasicInfoQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBasicInfoQAHeaders{}
	_result = &GetBasicInfoQAResponse{}
	_body, _err := client.GetBasicInfoQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询动画
//
// @param request - GetCartoonRequest
//
// @param headers - GetCartoonHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCartoonResponse
func (client *Client) GetCartoonWithOptions(request *GetCartoonRequest, headers *GetCartoonHeaders, runtime *util.RuntimeOptions) (_result *GetCartoonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCartoon"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getCartoon"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCartoonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询动画
//
// @param request - GetCartoonRequest
//
// @return GetCartoonResponse
func (client *Client) GetCartoon(request *GetCartoonRequest) (_result *GetCartoonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCartoonHeaders{}
	_result = &GetCartoonResponse{}
	_body, _err := client.GetCartoonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前设备的通话信息
//
// @param tmpReq - GetHotelContactByGenieDeviceRequest
//
// @param headers - GetHotelContactByGenieDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelContactByGenieDeviceResponse
func (client *Client) GetHotelContactByGenieDeviceWithOptions(tmpReq *GetHotelContactByGenieDeviceRequest, headers *GetHotelContactByGenieDeviceHeaders, runtime *util.RuntimeOptions) (_result *GetHotelContactByGenieDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelContactByGenieDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelContactByGenieDevice"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelContactByGenieDevice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelContactByGenieDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前设备的通话信息
//
// @param request - GetHotelContactByGenieDeviceRequest
//
// @return GetHotelContactByGenieDeviceResponse
func (client *Client) GetHotelContactByGenieDevice(request *GetHotelContactByGenieDeviceRequest) (_result *GetHotelContactByGenieDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelContactByGenieDeviceHeaders{}
	_result = &GetHotelContactByGenieDeviceResponse{}
	_body, _err := client.GetHotelContactByGenieDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据号码获取酒店联系人
//
// @param tmpReq - GetHotelContactByNumberRequest
//
// @param headers - GetHotelContactByNumberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelContactByNumberResponse
func (client *Client) GetHotelContactByNumberWithOptions(tmpReq *GetHotelContactByNumberRequest, headers *GetHotelContactByNumberHeaders, runtime *util.RuntimeOptions) (_result *GetHotelContactByNumberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelContactByNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Number)) {
		body["Number"] = request.Number
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelContactByNumber"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelContactByNumber"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelContactByNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据号码获取酒店联系人
//
// @param request - GetHotelContactByNumberRequest
//
// @return GetHotelContactByNumberResponse
func (client *Client) GetHotelContactByNumber(request *GetHotelContactByNumberRequest) (_result *GetHotelContactByNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelContactByNumberHeaders{}
	_result = &GetHotelContactByNumberResponse{}
	_body, _err := client.GetHotelContactByNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店联系人
//
// @param tmpReq - GetHotelContactsRequest
//
// @param headers - GetHotelContactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelContactsResponse
func (client *Client) GetHotelContactsWithOptions(tmpReq *GetHotelContactsRequest, headers *GetHotelContactsHeaders, runtime *util.RuntimeOptions) (_result *GetHotelContactsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelContactsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelContacts"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelContacts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店联系人
//
// @param request - GetHotelContactsRequest
//
// @return GetHotelContactsResponse
func (client *Client) GetHotelContacts(request *GetHotelContactsRequest) (_result *GetHotelContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelContactsHeaders{}
	_result = &GetHotelContactsResponse{}
	_body, _err := client.GetHotelContactsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取首页背景图和场景模式
//
// @param tmpReq - GetHotelHomeBackImageAndModesRequest
//
// @param headers - GetHotelHomeBackImageAndModesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelHomeBackImageAndModesResponse
func (client *Client) GetHotelHomeBackImageAndModesWithOptions(tmpReq *GetHotelHomeBackImageAndModesRequest, headers *GetHotelHomeBackImageAndModesHeaders, runtime *util.RuntimeOptions) (_result *GetHotelHomeBackImageAndModesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelHomeBackImageAndModesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelHomeBackImageAndModes"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelHomeBackImageAndModes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelHomeBackImageAndModesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取首页背景图和场景模式
//
// @param request - GetHotelHomeBackImageAndModesRequest
//
// @return GetHotelHomeBackImageAndModesResponse
func (client *Client) GetHotelHomeBackImageAndModes(request *GetHotelHomeBackImageAndModesRequest) (_result *GetHotelHomeBackImageAndModesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelHomeBackImageAndModesHeaders{}
	_result = &GetHotelHomeBackImageAndModesResponse{}
	_body, _err := client.GetHotelHomeBackImageAndModesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店通知
//
// @param tmpReq - GetHotelNoticeRequest
//
// @param headers - GetHotelNoticeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelNoticeResponse
func (client *Client) GetHotelNoticeWithOptions(tmpReq *GetHotelNoticeRequest, headers *GetHotelNoticeHeaders, runtime *util.RuntimeOptions) (_result *GetHotelNoticeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelNoticeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelNotice"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelNotice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店通知
//
// @param request - GetHotelNoticeRequest
//
// @return GetHotelNoticeResponse
func (client *Client) GetHotelNotice(request *GetHotelNoticeRequest) (_result *GetHotelNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelNoticeHeaders{}
	_result = &GetHotelNoticeResponse{}
	_body, _err := client.GetHotelNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店通知
//
// @param tmpReq - GetHotelNoticeV2Request
//
// @param headers - GetHotelNoticeV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelNoticeV2Response
func (client *Client) GetHotelNoticeV2WithOptions(tmpReq *GetHotelNoticeV2Request, headers *GetHotelNoticeV2Headers, runtime *util.RuntimeOptions) (_result *GetHotelNoticeV2Response, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelNoticeV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelNoticeV2"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelNoticeV2"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelNoticeV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店通知
//
// @param request - GetHotelNoticeV2Request
//
// @return GetHotelNoticeV2Response
func (client *Client) GetHotelNoticeV2(request *GetHotelNoticeV2Request) (_result *GetHotelNoticeV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelNoticeV2Headers{}
	_result = &GetHotelNoticeV2Response{}
	_body, _err := client.GetHotelNoticeV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店订单详情
//
// @param tmpReq - GetHotelOrderDetailRequest
//
// @param headers - GetHotelOrderDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelOrderDetailResponse
func (client *Client) GetHotelOrderDetailWithOptions(tmpReq *GetHotelOrderDetailRequest, headers *GetHotelOrderDetailHeaders, runtime *util.RuntimeOptions) (_result *GetHotelOrderDetailResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelOrderDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelOrderDetail"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelOrderDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店订单详情
//
// @param request - GetHotelOrderDetailRequest
//
// @return GetHotelOrderDetailResponse
func (client *Client) GetHotelOrderDetail(request *GetHotelOrderDetailRequest) (_result *GetHotelOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelOrderDetailHeaders{}
	_result = &GetHotelOrderDetailResponse{}
	_body, _err := client.GetHotelOrderDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店房间猫精设备信息
//
// @param request - GetHotelRoomDeviceRequest
//
// @param headers - GetHotelRoomDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelRoomDeviceResponse
func (client *Client) GetHotelRoomDeviceWithOptions(request *GetHotelRoomDeviceRequest, headers *GetHotelRoomDeviceHeaders, runtime *util.RuntimeOptions) (_result *GetHotelRoomDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		query["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		query["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelRoomDevice"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelRoomDevice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelRoomDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店房间猫精设备信息
//
// @param request - GetHotelRoomDeviceRequest
//
// @return GetHotelRoomDeviceResponse
func (client *Client) GetHotelRoomDevice(request *GetHotelRoomDeviceRequest) (_result *GetHotelRoomDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelRoomDeviceHeaders{}
	_result = &GetHotelRoomDeviceResponse{}
	_body, _err := client.GetHotelRoomDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取推荐语料
//
// @param tmpReq - GetHotelSampleUtterancesRequest
//
// @param headers - GetHotelSampleUtterancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelSampleUtterancesResponse
func (client *Client) GetHotelSampleUtterancesWithOptions(tmpReq *GetHotelSampleUtterancesRequest, headers *GetHotelSampleUtterancesHeaders, runtime *util.RuntimeOptions) (_result *GetHotelSampleUtterancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelSampleUtterancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelSampleUtterances"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelSampleUtterances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelSampleUtterancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取推荐语料
//
// @param request - GetHotelSampleUtterancesRequest
//
// @return GetHotelSampleUtterancesResponse
func (client *Client) GetHotelSampleUtterances(request *GetHotelSampleUtterancesRequest) (_result *GetHotelSampleUtterancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelSampleUtterancesHeaders{}
	_result = &GetHotelSampleUtterancesResponse{}
	_body, _err := client.GetHotelSampleUtterancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景详情
//
// @param request - GetHotelSceneItemDetailRequest
//
// @param headers - GetHotelSceneItemDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelSceneItemDetailResponse
func (client *Client) GetHotelSceneItemDetailWithOptions(request *GetHotelSceneItemDetailRequest, headers *GetHotelSceneItemDetailHeaders, runtime *util.RuntimeOptions) (_result *GetHotelSceneItemDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		body["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelSceneItemDetail"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelSceneItemDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelSceneItemDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景详情
//
// @param request - GetHotelSceneItemDetailRequest
//
// @return GetHotelSceneItemDetailResponse
func (client *Client) GetHotelSceneItemDetail(request *GetHotelSceneItemDetailRequest) (_result *GetHotelSceneItemDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelSceneItemDetailHeaders{}
	_result = &GetHotelSceneItemDetailResponse{}
	_body, _err := client.GetHotelSceneItemDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店屏保
//
// @param tmpReq - GetHotelScreenSaverRequest
//
// @param headers - GetHotelScreenSaverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelScreenSaverResponse
func (client *Client) GetHotelScreenSaverWithOptions(tmpReq *GetHotelScreenSaverRequest, headers *GetHotelScreenSaverHeaders, runtime *util.RuntimeOptions) (_result *GetHotelScreenSaverResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelScreenSaverShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelScreenSaver"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelScreenSaver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelScreenSaverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店屏保
//
// @param request - GetHotelScreenSaverRequest
//
// @return GetHotelScreenSaverResponse
func (client *Client) GetHotelScreenSaver(request *GetHotelScreenSaverRequest) (_result *GetHotelScreenSaverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelScreenSaverHeaders{}
	_result = &GetHotelScreenSaverResponse{}
	_body, _err := client.GetHotelScreenSaverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取屏保列表
//
// @param request - GetHotelScreenSaverStyleRequest
//
// @param headers - GetHotelScreenSaverStyleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelScreenSaverStyleResponse
func (client *Client) GetHotelScreenSaverStyleWithOptions(request *GetHotelScreenSaverStyleRequest, headers *GetHotelScreenSaverStyleHeaders, runtime *util.RuntimeOptions) (_result *GetHotelScreenSaverStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelScreenSaverStyle"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelScreenSaverStyle"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelScreenSaverStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取屏保列表
//
// @param request - GetHotelScreenSaverStyleRequest
//
// @return GetHotelScreenSaverStyleResponse
func (client *Client) GetHotelScreenSaverStyle(request *GetHotelScreenSaverStyleRequest) (_result *GetHotelScreenSaverStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelScreenSaverStyleHeaders{}
	_result = &GetHotelScreenSaverStyleResponse{}
	_body, _err := client.GetHotelScreenSaverStyleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询定制配置
//
// @param request - GetHotelSettingRequest
//
// @param headers - GetHotelSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelSettingResponse
func (client *Client) GetHotelSettingWithOptions(request *GetHotelSettingRequest, headers *GetHotelSettingHeaders, runtime *util.RuntimeOptions) (_result *GetHotelSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.SettingType)) {
		body["SettingType"] = request.SettingType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelSetting"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelSetting"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询定制配置
//
// @param request - GetHotelSettingRequest
//
// @return GetHotelSettingResponse
func (client *Client) GetHotelSetting(request *GetHotelSettingRequest) (_result *GetHotelSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelSettingHeaders{}
	_result = &GetHotelSettingResponse{}
	_body, _err := client.GetHotelSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关联产品列表查看
//
// @param headers - GetRelationProductListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRelationProductListResponse
func (client *Client) GetRelationProductListWithOptions(headers *GetRelationProductListHeaders, runtime *util.RuntimeOptions) (_result *GetRelationProductListResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRelationProductList"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getRelationProductList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRelationProductListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关联产品列表查看
//
// @return GetRelationProductListResponse
func (client *Client) GetRelationProductList() (_result *GetRelationProductListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRelationProductListHeaders{}
	_result = &GetRelationProductListResponse{}
	_body, _err := client.GetRelationProductListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取组织下unionId列表
//
// @param request - GetUnionIdRequest
//
// @param headers - GetUnionIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUnionIdResponse
func (client *Client) GetUnionIdWithOptions(request *GetUnionIdRequest, headers *GetUnionIdHeaders, runtime *util.RuntimeOptions) (_result *GetUnionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodeKey)) {
		body["EncodeKey"] = request.EncodeKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		body["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IdType)) {
		body["IdType"] = request.IdType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUnionId"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getUnionId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUnionIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取组织下unionId列表
//
// @param request - GetUnionIdRequest
//
// @return GetUnionIdResponse
func (client *Client) GetUnionId(request *GetUnionIdRequest) (_result *GetUnionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUnionIdHeaders{}
	_result = &GetUnionIdResponse{}
	_body, _err := client.GetUnionIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询欢迎语信息
//
// @param request - GetWelcomeTextAndMusicRequest
//
// @param headers - GetWelcomeTextAndMusicHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWelcomeTextAndMusicResponse
func (client *Client) GetWelcomeTextAndMusicWithOptions(request *GetWelcomeTextAndMusicRequest, headers *GetWelcomeTextAndMusicHeaders, runtime *util.RuntimeOptions) (_result *GetWelcomeTextAndMusicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWelcomeTextAndMusic"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getWelcomeTextAndMusic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWelcomeTextAndMusicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询欢迎语信息
//
// @param request - GetWelcomeTextAndMusicRequest
//
// @return GetWelcomeTextAndMusicResponse
func (client *Client) GetWelcomeTextAndMusic(request *GetWelcomeTextAndMusicRequest) (_result *GetWelcomeTextAndMusicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWelcomeTextAndMusicHeaders{}
	_result = &GetWelcomeTextAndMusicResponse{}
	_body, _err := client.GetWelcomeTextAndMusicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店带屏设备扫码绑定
//
// @param request - HotelQrBindRequest
//
// @param headers - HotelQrBindHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelQrBindResponse
func (client *Client) HotelQrBindWithOptions(request *HotelQrBindRequest, headers *HotelQrBindHeaders, runtime *util.RuntimeOptions) (_result *HotelQrBindResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HotelQrBind"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/hotelQrBind"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HotelQrBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店带屏设备扫码绑定
//
// @param request - HotelQrBindRequest
//
// @return HotelQrBindResponse
func (client *Client) HotelQrBind(request *HotelQrBindRequest) (_result *HotelQrBindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HotelQrBindHeaders{}
	_result = &HotelQrBindResponse{}
	_body, _err := client.HotelQrBindWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量导入酒店配置
//
// @param tmpReq - ImportHotelConfigRequest
//
// @param headers - ImportHotelConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportHotelConfigResponse
func (client *Client) ImportHotelConfigWithOptions(tmpReq *ImportHotelConfigRequest, headers *ImportHotelConfigHeaders, runtime *util.RuntimeOptions) (_result *ImportHotelConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportHotelConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ImportHotelConfig)) {
		request.ImportHotelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImportHotelConfig, tea.String("ImportHotelConfig"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.ImportHotelConfigShrink)) {
		body["ImportHotelConfig"] = request.ImportHotelConfigShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportHotelConfig"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/importHotelConfig"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportHotelConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导入酒店配置
//
// @param request - ImportHotelConfigRequest
//
// @return ImportHotelConfigResponse
func (client *Client) ImportHotelConfig(request *ImportHotelConfigRequest) (_result *ImportHotelConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ImportHotelConfigHeaders{}
	_result = &ImportHotelConfigResponse{}
	_body, _err := client.ImportHotelConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量导入设备（同时补充房型）
//
// @param tmpReq - ImportRoomControlDevicesRequest
//
// @param headers - ImportRoomControlDevicesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportRoomControlDevicesResponse
func (client *Client) ImportRoomControlDevicesWithOptions(tmpReq *ImportRoomControlDevicesRequest, headers *ImportRoomControlDevicesHeaders, runtime *util.RuntimeOptions) (_result *ImportRoomControlDevicesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportRoomControlDevicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LocationDevices)) {
		request.LocationDevicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationDevices, tea.String("LocationDevices"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableInfraredDeviceImport)) {
		body["EnableInfraredDeviceImport"] = request.EnableInfraredDeviceImport
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.LocationDevicesShrink)) {
		body["LocationDevices"] = request.LocationDevicesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportRoomControlDevices"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/importRoomControlDevices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportRoomControlDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导入设备（同时补充房型）
//
// @param request - ImportRoomControlDevicesRequest
//
// @return ImportRoomControlDevicesResponse
func (client *Client) ImportRoomControlDevices(request *ImportRoomControlDevicesRequest) (_result *ImportRoomControlDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ImportRoomControlDevicesHeaders{}
	_result = &ImportRoomControlDevicesResponse{}
	_body, _err := client.ImportRoomControlDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入房间内精灵场景
//
// @param tmpReq - ImportRoomGenieScenesRequest
//
// @param headers - ImportRoomGenieScenesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportRoomGenieScenesResponse
func (client *Client) ImportRoomGenieScenesWithOptions(tmpReq *ImportRoomGenieScenesRequest, headers *ImportRoomGenieScenesHeaders, runtime *util.RuntimeOptions) (_result *ImportRoomGenieScenesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportRoomGenieScenesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SceneList)) {
		request.SceneListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneList, tea.String("SceneList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	if !tea.BoolValue(util.IsUnset(request.SceneListShrink)) {
		body["SceneList"] = request.SceneListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportRoomGenieScenes"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/importRoomGenieScenes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportRoomGenieScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入房间内精灵场景
//
// @param request - ImportRoomGenieScenesRequest
//
// @return ImportRoomGenieScenesResponse
func (client *Client) ImportRoomGenieScenes(request *ImportRoomGenieScenesRequest) (_result *ImportRoomGenieScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ImportRoomGenieScenesHeaders{}
	_result = &ImportRoomGenieScenesResponse{}
	_body, _err := client.ImportRoomGenieScenesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景预订新增
//
// @param tmpReq - InsertHotelSceneBookItemRequest
//
// @param headers - InsertHotelSceneBookItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertHotelSceneBookItemResponse
func (client *Client) InsertHotelSceneBookItemWithOptions(tmpReq *InsertHotelSceneBookItemRequest, headers *InsertHotelSceneBookItemHeaders, runtime *util.RuntimeOptions) (_result *InsertHotelSceneBookItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InsertHotelSceneBookItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AddHotelSceneItemReq)) {
		request.AddHotelSceneItemReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddHotelSceneItemReq, tea.String("AddHotelSceneItemReq"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddHotelSceneItemReqShrink)) {
		query["AddHotelSceneItemReq"] = request.AddHotelSceneItemReqShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertHotelSceneBookItem"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/insertHotelSceneBookItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertHotelSceneBookItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景预订新增
//
// @param request - InsertHotelSceneBookItemRequest
//
// @return InsertHotelSceneBookItemResponse
func (client *Client) InsertHotelSceneBookItem(request *InsertHotelSceneBookItemRequest) (_result *InsertHotelSceneBookItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertHotelSceneBookItemHeaders{}
	_result = &InsertHotelSceneBookItemResponse{}
	_body, _err := client.InsertHotelSceneBookItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人服务，消息推送
//
// @param request - InvokeRobotPushRequest
//
// @param headers - InvokeRobotPushHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeRobotPushResponse
func (client *Client) InvokeRobotPushWithOptions(request *InvokeRobotPushRequest, headers *InvokeRobotPushHeaders, runtime *util.RuntimeOptions) (_result *InvokeRobotPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.PushType)) {
		body["PushType"] = request.PushType
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeRobotPush"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/invokeRobotPush"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeRobotPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人服务，消息推送
//
// @param request - InvokeRobotPushRequest
//
// @return InvokeRobotPushResponse
func (client *Client) InvokeRobotPush(request *InvokeRobotPushRequest) (_result *InvokeRobotPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvokeRobotPushHeaders{}
	_result = &InvokeRobotPushResponse{}
	_body, _err := client.InvokeRobotPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询省份
//
// @param headers - ListAllProvincesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllProvincesResponse
func (client *Client) ListAllProvincesWithOptions(headers *ListAllProvincesHeaders, runtime *util.RuntimeOptions) (_result *ListAllProvincesResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("ListAllProvinces"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listAllProvinces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllProvincesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询省份
//
// @return ListAllProvincesResponse
func (client *Client) ListAllProvinces() (_result *ListAllProvincesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAllProvincesHeaders{}
	_result = &ListAllProvincesResponse{}
	_body, _err := client.ListAllProvincesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询城市
//
// @param request - ListCitiesByProvinceRequest
//
// @param headers - ListCitiesByProvinceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCitiesByProvinceResponse
func (client *Client) ListCitiesByProvinceWithOptions(request *ListCitiesByProvinceRequest, headers *ListCitiesByProvinceHeaders, runtime *util.RuntimeOptions) (_result *ListCitiesByProvinceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Province)) {
		body["Province"] = request.Province
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCitiesByProvince"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listCitiesByProvince"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCitiesByProvinceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询城市
//
// @param request - ListCitiesByProvinceRequest
//
// @return ListCitiesByProvinceResponse
func (client *Client) ListCitiesByProvince(request *ListCitiesByProvinceRequest) (_result *ListCitiesByProvinceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCitiesByProvinceHeaders{}
	_result = &ListCitiesByProvinceResponse{}
	_body, _err := client.ListCitiesByProvinceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自定义问答列表
//
// @param tmpReq - ListCustomQARequest
//
// @param headers - ListCustomQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomQAResponse
func (client *Client) ListCustomQAWithOptions(tmpReq *ListCustomQARequest, headers *ListCustomQAHeaders, runtime *util.RuntimeOptions) (_result *ListCustomQAResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCustomQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Page)) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, tea.String("Page"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		body["Page"] = request.PageShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomQA"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listCustomQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义问答列表
//
// @param request - ListCustomQARequest
//
// @return ListCustomQAResponse
func (client *Client) ListCustomQA(request *ListCustomQARequest) (_result *ListCustomQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCustomQAHeaders{}
	_result = &ListCustomQAResponse{}
	_body, _err := client.ListCustomQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景对话模板
//
// @param request - ListDialogueTemplateRequest
//
// @param headers - ListDialogueTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDialogueTemplateResponse
func (client *Client) ListDialogueTemplateWithOptions(request *ListDialogueTemplateRequest, headers *ListDialogueTemplateHeaders, runtime *util.RuntimeOptions) (_result *ListDialogueTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDialogueTemplate"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listDialogueTemplate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDialogueTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景对话模板
//
// @param request - ListDialogueTemplateRequest
//
// @return ListDialogueTemplateResponse
func (client *Client) ListDialogueTemplate(request *ListDialogueTemplateRequest) (_result *ListDialogueTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDialogueTemplateHeaders{}
	_result = &ListDialogueTemplateResponse{}
	_body, _err := client.ListDialogueTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店闹钟
//
// @param tmpReq - ListHotelAlarmRequest
//
// @param headers - ListHotelAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelAlarmResponse
func (client *Client) ListHotelAlarmWithOptions(tmpReq *ListHotelAlarmRequest, headers *ListHotelAlarmHeaders, runtime *util.RuntimeOptions) (_result *ListHotelAlarmResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Rooms)) {
		request.RoomsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rooms, tea.String("Rooms"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomsShrink)) {
		body["Rooms"] = request.RoomsShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelAlarm"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelAlarmList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店闹钟
//
// @param request - ListHotelAlarmRequest
//
// @return ListHotelAlarmResponse
func (client *Client) ListHotelAlarm(request *ListHotelAlarmRequest) (_result *ListHotelAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelAlarmHeaders{}
	_result = &ListHotelAlarmResponse{}
	_body, _err := client.ListHotelAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店设备列表
//
// @param tmpReq - ListHotelControlDeviceRequest
//
// @param headers - ListHotelControlDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelControlDeviceResponse
func (client *Client) ListHotelControlDeviceWithOptions(tmpReq *ListHotelControlDeviceRequest, headers *ListHotelControlDeviceHeaders, runtime *util.RuntimeOptions) (_result *ListHotelControlDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelControlDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelControlDevice"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelControlDevice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelControlDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店设备列表
//
// @param request - ListHotelControlDeviceRequest
//
// @return ListHotelControlDeviceResponse
func (client *Client) ListHotelControlDevice(request *ListHotelControlDeviceRequest) (_result *ListHotelControlDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelControlDeviceHeaders{}
	_result = &ListHotelControlDeviceResponse{}
	_body, _err := client.ListHotelControlDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店列表
//
// @param headers - ListHotelInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelInfoResponse
func (client *Client) ListHotelInfoWithOptions(headers *ListHotelInfoHeaders, runtime *util.RuntimeOptions) (_result *ListHotelInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelInfo"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店列表
//
// @return ListHotelInfoResponse
func (client *Client) ListHotelInfo() (_result *ListHotelInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelInfoHeaders{}
	_result = &ListHotelInfoResponse{}
	_body, _err := client.ListHotelInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取消息模板
//
// @param headers - ListHotelMessageTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelMessageTemplateResponse
func (client *Client) ListHotelMessageTemplateWithOptions(headers *ListHotelMessageTemplateHeaders, runtime *util.RuntimeOptions) (_result *ListHotelMessageTemplateResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelMessageTemplate"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelMessageTemplate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelMessageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取消息模板
//
// @return ListHotelMessageTemplateResponse
func (client *Client) ListHotelMessageTemplate() (_result *ListHotelMessageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelMessageTemplateHeaders{}
	_result = &ListHotelMessageTemplateResponse{}
	_body, _err := client.ListHotelMessageTemplateWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店订单列表
//
// @param tmpReq - ListHotelOrderRequest
//
// @param headers - ListHotelOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelOrderResponse
func (client *Client) ListHotelOrderWithOptions(tmpReq *ListHotelOrderRequest, headers *ListHotelOrderHeaders, runtime *util.RuntimeOptions) (_result *ListHotelOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelOrder"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelOrder"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单列表
//
// @param request - ListHotelOrderRequest
//
// @return ListHotelOrderResponse
func (client *Client) ListHotelOrder(request *ListHotelOrderRequest) (_result *ListHotelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelOrderHeaders{}
	_result = &ListHotelOrderResponse{}
	_body, _err := client.ListHotelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店的所有房间
//
// @param tmpReq - ListHotelRoomsRequest
//
// @param headers - ListHotelRoomsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelRoomsResponse
func (client *Client) ListHotelRoomsWithOptions(tmpReq *ListHotelRoomsRequest, headers *ListHotelRoomsHeaders, runtime *util.RuntimeOptions) (_result *ListHotelRoomsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelRoomsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HotelAdminRoom)) {
		request.HotelAdminRoomShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelAdminRoom, tea.String("HotelAdminRoom"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelAdminRoomShrink)) {
		body["HotelAdminRoom"] = request.HotelAdminRoomShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelRooms"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelRooms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelRoomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店的所有房间
//
// @param request - ListHotelRoomsRequest
//
// @return ListHotelRoomsResponse
func (client *Client) ListHotelRooms(request *ListHotelRoomsRequest) (_result *ListHotelRoomsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelRoomsHeaders{}
	_result = &ListHotelRoomsResponse{}
	_body, _err := client.ListHotelRoomsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景预订列表（餐饮/SPA休闲/打车）
//
// @param tmpReq - ListHotelSceneBookItemsRequest
//
// @param headers - ListHotelSceneBookItemsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelSceneBookItemsResponse
func (client *Client) ListHotelSceneBookItemsWithOptions(tmpReq *ListHotelSceneBookItemsRequest, headers *ListHotelSceneBookItemsHeaders, runtime *util.RuntimeOptions) (_result *ListHotelSceneBookItemsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelSceneBookItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Page)) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, tea.String("Page"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		query["Page"] = request.PageShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelSceneBookItems"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelSceneBookItems"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelSceneBookItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景预订列表（餐饮/SPA休闲/打车）
//
// @param request - ListHotelSceneBookItemsRequest
//
// @return ListHotelSceneBookItemsResponse
func (client *Client) ListHotelSceneBookItems(request *ListHotelSceneBookItemsRequest) (_result *ListHotelSceneBookItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelSceneBookItemsHeaders{}
	_result = &ListHotelSceneBookItemsResponse{}
	_body, _err := client.ListHotelSceneBookItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务项目
//
// @param tmpReq - ListHotelSceneItemRequest
//
// @param headers - ListHotelSceneItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelSceneItemResponse
func (client *Client) ListHotelSceneItemWithOptions(tmpReq *ListHotelSceneItemRequest, headers *ListHotelSceneItemHeaders, runtime *util.RuntimeOptions) (_result *ListHotelSceneItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelSceneItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelSceneItem"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelSceneItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelSceneItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务项目
//
// @param request - ListHotelSceneItemRequest
//
// @return ListHotelSceneItemResponse
func (client *Client) ListHotelSceneItem(request *ListHotelSceneItemRequest) (_result *ListHotelSceneItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelSceneItemHeaders{}
	_result = &ListHotelSceneItemResponse{}
	_body, _err := client.ListHotelSceneItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景列表（物品/服务/维修）
//
// @param tmpReq - ListHotelSceneItemsRequest
//
// @param headers - ListHotelSceneItemsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelSceneItemsResponse
func (client *Client) ListHotelSceneItemsWithOptions(tmpReq *ListHotelSceneItemsRequest, headers *ListHotelSceneItemsHeaders, runtime *util.RuntimeOptions) (_result *ListHotelSceneItemsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelSceneItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListHotelSceneReq)) {
		request.ListHotelSceneReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListHotelSceneReq, tea.String("ListHotelSceneReq"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListHotelSceneReqShrink)) {
		query["ListHotelSceneReq"] = request.ListHotelSceneReqShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelSceneItems"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelSceneItems"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelSceneItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景列表（物品/服务/维修）
//
// @param request - ListHotelSceneItemsRequest
//
// @return ListHotelSceneItemsResponse
func (client *Client) ListHotelSceneItems(request *ListHotelSceneItemsRequest) (_result *ListHotelSceneItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelSceneItemsHeaders{}
	_result = &ListHotelSceneItemsResponse{}
	_body, _err := client.ListHotelSceneItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务分类列表
//
// @param tmpReq - ListHotelServiceCategoryRequest
//
// @param headers - ListHotelServiceCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelServiceCategoryResponse
func (client *Client) ListHotelServiceCategoryWithOptions(tmpReq *ListHotelServiceCategoryRequest, headers *ListHotelServiceCategoryHeaders, runtime *util.RuntimeOptions) (_result *ListHotelServiceCategoryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelServiceCategoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelServiceCategory"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelServiceCategory"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelServiceCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务分类列表
//
// @param request - ListHotelServiceCategoryRequest
//
// @return ListHotelServiceCategoryResponse
func (client *Client) ListHotelServiceCategory(request *ListHotelServiceCategoryRequest) (_result *ListHotelServiceCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelServiceCategoryHeaders{}
	_result = &ListHotelServiceCategoryResponse{}
	_body, _err := client.ListHotelServiceCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店列表(待审批/已拒绝/已通过)
//
// @param tmpReq - ListHotelsRequest
//
// @param headers - ListHotelsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelsResponse
func (client *Client) ListHotelsWithOptions(tmpReq *ListHotelsRequest, headers *ListHotelsHeaders, runtime *util.RuntimeOptions) (_result *ListHotelsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HotelRequest)) {
		request.HotelRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelRequest, tea.String("HotelRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Page)) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, tea.String("Page"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelRequestShrink)) {
		query["HotelRequest"] = request.HotelRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		query["Page"] = request.PageShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotels"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店列表(待审批/已拒绝/已通过)
//
// @param request - ListHotelsRequest
//
// @return ListHotelsResponse
func (client *Client) ListHotels(request *ListHotelsRequest) (_result *ListHotelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelsHeaders{}
	_result = &ListHotelsResponse{}
	_body, _err := client.ListHotelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询红外品牌列表
//
// @param request - ListInfraredDeviceBrandsRequest
//
// @param headers - ListInfraredDeviceBrandsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInfraredDeviceBrandsResponse
func (client *Client) ListInfraredDeviceBrandsWithOptions(request *ListInfraredDeviceBrandsRequest, headers *ListInfraredDeviceBrandsHeaders, runtime *util.RuntimeOptions) (_result *ListInfraredDeviceBrandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceProvider)) {
		body["ServiceProvider"] = request.ServiceProvider
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInfraredDeviceBrands"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listInfraredDeviceBrands"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInfraredDeviceBrandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询红外品牌列表
//
// @param request - ListInfraredDeviceBrandsRequest
//
// @return ListInfraredDeviceBrandsResponse
func (client *Client) ListInfraredDeviceBrands(request *ListInfraredDeviceBrandsRequest) (_result *ListInfraredDeviceBrandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListInfraredDeviceBrandsHeaders{}
	_result = &ListInfraredDeviceBrandsResponse{}
	_body, _err := client.ListInfraredDeviceBrandsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询红外码库列表
//
// @param request - ListInfraredRemoteControllersRequest
//
// @param headers - ListInfraredRemoteControllersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInfraredRemoteControllersResponse
func (client *Client) ListInfraredRemoteControllersWithOptions(request *ListInfraredRemoteControllersRequest, headers *ListInfraredRemoteControllersHeaders, runtime *util.RuntimeOptions) (_result *ListInfraredRemoteControllersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Brand)) {
		body["Brand"] = request.Brand
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		body["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		body["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceProvider)) {
		body["ServiceProvider"] = request.ServiceProvider
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInfraredRemoteControllers"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listInfraredRemoteControllers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInfraredRemoteControllersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询红外码库列表
//
// @param request - ListInfraredRemoteControllersRequest
//
// @return ListInfraredRemoteControllersResponse
func (client *Client) ListInfraredRemoteControllers(request *ListInfraredRemoteControllersRequest) (_result *ListInfraredRemoteControllersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListInfraredRemoteControllersHeaders{}
	_result = &ListInfraredRemoteControllersResponse{}
	_body, _err := client.ListInfraredRemoteControllersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务提供商
//
// @param request - ListSTBServiceProvidersRequest
//
// @param headers - ListSTBServiceProvidersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSTBServiceProvidersResponse
func (client *Client) ListSTBServiceProvidersWithOptions(request *ListSTBServiceProvidersRequest, headers *ListSTBServiceProvidersHeaders, runtime *util.RuntimeOptions) (_result *ListSTBServiceProvidersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		body["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		body["Province"] = request.Province
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSTBServiceProviders"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listSTBServiceProviders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSTBServiceProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务提供商
//
// @param request - ListSTBServiceProvidersRequest
//
// @return ListSTBServiceProvidersResponse
func (client *Client) ListSTBServiceProviders(request *ListSTBServiceProvidersRequest) (_result *ListSTBServiceProvidersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSTBServiceProvidersHeaders{}
	_result = &ListSTBServiceProvidersResponse{}
	_body, _err := client.ListSTBServiceProvidersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景分类
//
// @param request - ListSceneCategoryRequest
//
// @param headers - ListSceneCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSceneCategoryResponse
func (client *Client) ListSceneCategoryWithOptions(request *ListSceneCategoryRequest, headers *ListSceneCategoryHeaders, runtime *util.RuntimeOptions) (_result *ListSceneCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSceneCategory"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listSceneCategory"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSceneCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景分类
//
// @param request - ListSceneCategoryRequest
//
// @return ListSceneCategoryResponse
func (client *Client) ListSceneCategory(request *ListSceneCategoryRequest) (_result *ListSceneCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSceneCategoryHeaders{}
	_result = &ListSceneCategoryResponse{}
	_body, _err := client.ListSceneCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务设施问答列表
//
// @param tmpReq - ListServiceQARequest
//
// @param headers - ListServiceQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceQAResponse
func (client *Client) ListServiceQAWithOptions(tmpReq *ListServiceQARequest, headers *ListServiceQAHeaders, runtime *util.RuntimeOptions) (_result *ListServiceQAResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListServiceQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Page)) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, tea.String("Page"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		body["Page"] = request.PageShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceQA"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listServiceQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务设施问答列表
//
// @param request - ListServiceQARequest
//
// @return ListServiceQAResponse
func (client *Client) ListServiceQA(request *ListServiceQARequest) (_result *ListServiceQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListServiceQAHeaders{}
	_result = &ListServiceQAResponse{}
	_body, _err := client.ListServiceQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工单列表
//
// @param tmpReq - ListTicketsRequest
//
// @param headers - ListTicketsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketsResponse
func (client *Client) ListTicketsWithOptions(tmpReq *ListTicketsRequest, headers *ListTicketsHeaders, runtime *util.RuntimeOptions) (_result *ListTicketsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Page)) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, tea.String("Page"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.IsDesc)) {
		body["IsDesc"] = request.IsDesc
	}

	if !tea.BoolValue(util.IsUnset(request.IsNeedCallback)) {
		body["IsNeedCallback"] = request.IsNeedCallback
	}

	if !tea.BoolValue(util.IsUnset(request.IsNeedCharges)) {
		body["IsNeedCharges"] = request.IsNeedCharges
	}

	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		body["Page"] = request.PageShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		body["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTickets"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listTickets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工单列表
//
// @param request - ListTicketsRequest
//
// @return ListTicketsResponse
func (client *Client) ListTickets(request *ListTicketsRequest) (_result *ListTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTicketsHeaders{}
	_result = &ListTicketsResponse{}
	_body, _err := client.ListTicketsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询酒店房间主控设备
//
// @param request - PageGetHotelRoomDevicesRequest
//
// @param headers - PageGetHotelRoomDevicesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageGetHotelRoomDevicesResponse
func (client *Client) PageGetHotelRoomDevicesWithOptions(request *PageGetHotelRoomDevicesRequest, headers *PageGetHotelRoomDevicesHeaders, runtime *util.RuntimeOptions) (_result *PageGetHotelRoomDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageGetHotelRoomDevices"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/pageGetHotelRoomDevices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageGetHotelRoomDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询酒店房间主控设备
//
// @param request - PageGetHotelRoomDevicesRequest
//
// @return PageGetHotelRoomDevicesResponse
func (client *Client) PageGetHotelRoomDevices(request *PageGetHotelRoomDevicesRequest) (_result *PageGetHotelRoomDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageGetHotelRoomDevicesHeaders{}
	_result = &PageGetHotelRoomDevicesResponse{}
	_body, _err := client.PageGetHotelRoomDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// pms事件上报
//
// @param request - PmsEventReportRequest
//
// @param headers - PmsEventReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PmsEventReportResponse
func (client *Client) PmsEventReportWithOptions(request *PmsEventReportRequest, headers *PmsEventReportHeaders, runtime *util.RuntimeOptions) (_result *PmsEventReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Payload)) {
		body["Payload"] = request.Payload
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PmsEventReport"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/pmsEventReport"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PmsEventReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pms事件上报
//
// @param request - PmsEventReportRequest
//
// @return PmsEventReportResponse
func (client *Client) PmsEventReport(request *PmsEventReportRequest) (_result *PmsEventReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PmsEventReportHeaders{}
	_result = &PmsEventReportResponse{}
	_body, _err := client.PmsEventReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送酒店消息
//
// @param tmpReq - PushHotelMessageRequest
//
// @param headers - PushHotelMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushHotelMessageResponse
func (client *Client) PushHotelMessageWithOptions(tmpReq *PushHotelMessageRequest, headers *PushHotelMessageHeaders, runtime *util.RuntimeOptions) (_result *PushHotelMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PushHotelMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PushHotelMessageReq)) {
		request.PushHotelMessageReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PushHotelMessageReq, tea.String("PushHotelMessageReq"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PushHotelMessageReqShrink)) {
		query["PushHotelMessageReq"] = request.PushHotelMessageReqShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushHotelMessage"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/pushHotelMessage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushHotelMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送酒店消息
//
// @param request - PushHotelMessageRequest
//
// @return PushHotelMessageResponse
func (client *Client) PushHotelMessage(request *PushHotelMessageRequest) (_result *PushHotelMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushHotelMessageHeaders{}
	_result = &PushHotelMessageResponse{}
	_body, _err := client.PushHotelMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送音箱指令
//
// @param tmpReq - PushVoiceBoxCommandsRequest
//
// @param headers - PushVoiceBoxCommandsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushVoiceBoxCommandsResponse
func (client *Client) PushVoiceBoxCommandsWithOptions(tmpReq *PushVoiceBoxCommandsRequest, headers *PushVoiceBoxCommandsHeaders, runtime *util.RuntimeOptions) (_result *PushVoiceBoxCommandsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PushVoiceBoxCommandsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Commands)) {
		request.CommandsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commands, tea.String("Commands"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandsShrink)) {
		body["Commands"] = request.CommandsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushVoiceBoxCommands"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/pushVoiceBoxCommands"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushVoiceBoxCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送音箱指令
//
// @param request - PushVoiceBoxCommandsRequest
//
// @return PushVoiceBoxCommandsResponse
func (client *Client) PushVoiceBoxCommands(request *PushVoiceBoxCommandsRequest) (_result *PushVoiceBoxCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushVoiceBoxCommandsHeaders{}
	_result = &PushVoiceBoxCommandsResponse{}
	_body, _err := client.PushVoiceBoxCommandsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 直接推送欢迎语
//
// @param request - PushWelcomeRequest
//
// @param headers - PushWelcomeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushWelcomeResponse
func (client *Client) PushWelcomeWithOptions(request *PushWelcomeRequest, headers *PushWelcomeHeaders, runtime *util.RuntimeOptions) (_result *PushWelcomeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	if !tea.BoolValue(util.IsUnset(request.WelcomeMusicUrl)) {
		body["WelcomeMusicUrl"] = request.WelcomeMusicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WelcomeText)) {
		body["WelcomeText"] = request.WelcomeText
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushWelcome"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/pushWelcome"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushWelcomeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直接推送欢迎语
//
// @param request - PushWelcomeRequest
//
// @return PushWelcomeResponse
func (client *Client) PushWelcome(request *PushWelcomeRequest) (_result *PushWelcomeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushWelcomeHeaders{}
	_result = &PushWelcomeResponse{}
	_body, _err := client.PushWelcomeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送欢迎语
//
// @param tmpReq - PushWelcomeTextAndMusicRequest
//
// @param headers - PushWelcomeTextAndMusicHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushWelcomeTextAndMusicResponse
func (client *Client) PushWelcomeTextAndMusicWithOptions(tmpReq *PushWelcomeTextAndMusicRequest, headers *PushWelcomeTextAndMusicHeaders, runtime *util.RuntimeOptions) (_result *PushWelcomeTextAndMusicResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PushWelcomeTextAndMusicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TemplateVariable)) {
		request.TemplateVariableShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateVariable, tea.String("TemplateVariable"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVariableShrink)) {
		body["TemplateVariable"] = request.TemplateVariableShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushWelcomeTextAndMusic"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/pushWelcomeTextAndMusic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushWelcomeTextAndMusicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送欢迎语
//
// @param request - PushWelcomeTextAndMusicRequest
//
// @return PushWelcomeTextAndMusicResponse
func (client *Client) PushWelcomeTextAndMusic(request *PushWelcomeTextAndMusicRequest) (_result *PushWelcomeTextAndMusicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushWelcomeTextAndMusicHeaders{}
	_result = &PushWelcomeTextAndMusicResponse{}
	_body, _err := client.PushWelcomeTextAndMusicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店设备状态/模式状态查询
//
// @param tmpReq - QueryDeviceStatusRequest
//
// @param headers - QueryDeviceStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceStatusResponse
func (client *Client) QueryDeviceStatusWithOptions(tmpReq *QueryDeviceStatusRequest, headers *QueryDeviceStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryDeviceStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceStatus"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/queryDeviceStatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店设备状态/模式状态查询
//
// @param request - QueryDeviceStatusRequest
//
// @return QueryDeviceStatusResponse
func (client *Client) QueryDeviceStatus(request *QueryDeviceStatusRequest) (_result *QueryDeviceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceStatusHeaders{}
	_result = &QueryDeviceStatusResponse{}
	_body, _err := client.QueryDeviceStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询房间详细信息
//
// @param request - QueryHotelRoomDetailRequest
//
// @param headers - QueryHotelRoomDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHotelRoomDetailResponse
func (client *Client) QueryHotelRoomDetailWithOptions(request *QueryHotelRoomDetailRequest, headers *QueryHotelRoomDetailHeaders, runtime *util.RuntimeOptions) (_result *QueryHotelRoomDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Mac)) {
		body["Mac"] = request.Mac
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["Sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHotelRoomDetail"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/queryHotelRoomDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHotelRoomDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询房间详细信息
//
// @param request - QueryHotelRoomDetailRequest
//
// @return QueryHotelRoomDetailResponse
func (client *Client) QueryHotelRoomDetail(request *QueryHotelRoomDetailRequest) (_result *QueryHotelRoomDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHotelRoomDetailHeaders{}
	_result = &QueryHotelRoomDetailResponse{}
	_body, _err := client.QueryHotelRoomDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店房间客控设备
//
// @param request - QueryRoomControlDevicesRequest
//
// @param headers - QueryRoomControlDevicesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRoomControlDevicesResponse
func (client *Client) QueryRoomControlDevicesWithOptions(request *QueryRoomControlDevicesRequest, headers *QueryRoomControlDevicesHeaders, runtime *util.RuntimeOptions) (_result *QueryRoomControlDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		query["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		query["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRoomControlDevices"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/queryRoomControlDevices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRoomControlDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店房间客控设备
//
// @param request - QueryRoomControlDevicesRequest
//
// @return QueryRoomControlDevicesResponse
func (client *Client) QueryRoomControlDevices(request *QueryRoomControlDevicesRequest) (_result *QueryRoomControlDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRoomControlDevicesHeaders{}
	_result = &QueryRoomControlDevicesResponse{}
	_body, _err := client.QueryRoomControlDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询房间被控设备包含设备状态
//
// @param request - QueryRoomControlDevicesAndStatusRequest
//
// @param headers - QueryRoomControlDevicesAndStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRoomControlDevicesAndStatusResponse
func (client *Client) QueryRoomControlDevicesAndStatusWithOptions(request *QueryRoomControlDevicesAndStatusRequest, headers *QueryRoomControlDevicesAndStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryRoomControlDevicesAndStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRoomControlDevicesAndStatus"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/queryRoomControlDevicesAndStatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRoomControlDevicesAndStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询房间被控设备包含设备状态
//
// @param request - QueryRoomControlDevicesAndStatusRequest
//
// @return QueryRoomControlDevicesAndStatusResponse
func (client *Client) QueryRoomControlDevicesAndStatus(request *QueryRoomControlDevicesAndStatusRequest) (_result *QueryRoomControlDevicesAndStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRoomControlDevicesAndStatusHeaders{}
	_result = &QueryRoomControlDevicesAndStatusResponse{}
	_body, _err := client.QueryRoomControlDevicesAndStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店场景列表
//
// @param tmpReq - QuerySceneListRequest
//
// @param headers - QuerySceneListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySceneListResponse
func (client *Client) QuerySceneListWithOptions(tmpReq *QuerySceneListRequest, headers *QuerySceneListHeaders, runtime *util.RuntimeOptions) (_result *QuerySceneListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QuerySceneListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SceneStates)) {
		request.SceneStatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneStates, tea.String("SceneStates"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SceneTypes)) {
		request.SceneTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneTypes, tea.String("SceneTypes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TemplateInfoIds)) {
		request.TemplateInfoIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateInfoIds, tea.String("TemplateInfoIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneStatesShrink)) {
		body["SceneStates"] = request.SceneStatesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SceneTypesShrink)) {
		body["SceneTypes"] = request.SceneTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateInfoIdsShrink)) {
		body["TemplateInfoIds"] = request.TemplateInfoIdsShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySceneList"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/querySceneList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySceneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店场景列表
//
// @param request - QuerySceneListRequest
//
// @return QuerySceneListResponse
func (client *Client) QuerySceneList(request *QuerySceneListRequest) (_result *QuerySceneListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySceneListHeaders{}
	_result = &QuerySceneListResponse{}
	_body, _err := client.QuerySceneListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除子账号授权
//
// @param request - RemoveChildAccountAuthRequest
//
// @param headers - RemoveChildAccountAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveChildAccountAuthResponse
func (client *Client) RemoveChildAccountAuthWithOptions(request *RemoveChildAccountAuthRequest, headers *RemoveChildAccountAuthHeaders, runtime *util.RuntimeOptions) (_result *RemoveChildAccountAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ChildAccountName)) {
		body["ChildAccountName"] = request.ChildAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.TbOpenId)) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveChildAccountAuth"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/removeChildAccountAuth"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveChildAccountAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除子账号授权
//
// @param request - RemoveChildAccountAuthRequest
//
// @return RemoveChildAccountAuthResponse
func (client *Client) RemoveChildAccountAuth(request *RemoveChildAccountAuthRequest) (_result *RemoveChildAccountAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveChildAccountAuthHeaders{}
	_result = &RemoveChildAccountAuthResponse{}
	_body, _err := client.RemoveChildAccountAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除酒店项目
//
// @param request - RemoveHotelRequest
//
// @param headers - RemoveHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveHotelResponse
func (client *Client) RemoveHotelWithOptions(request *RemoveHotelRequest, headers *RemoveHotelHeaders, runtime *util.RuntimeOptions) (_result *RemoveHotelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.TbOpenId)) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveHotel"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/removeHotel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除酒店项目
//
// @param request - RemoveHotelRequest
//
// @return RemoveHotelResponse
func (client *Client) RemoveHotel(request *RemoveHotelRequest) (_result *RemoveHotelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveHotelHeaders{}
	_result = &RemoveHotelResponse{}
	_body, _err := client.RemoveHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置欢迎语信息
//
// @param request - ResetWelcomeTextAndMusicRequest
//
// @param headers - ResetWelcomeTextAndMusicHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetWelcomeTextAndMusicResponse
func (client *Client) ResetWelcomeTextAndMusicWithOptions(request *ResetWelcomeTextAndMusicRequest, headers *ResetWelcomeTextAndMusicHeaders, runtime *util.RuntimeOptions) (_result *ResetWelcomeTextAndMusicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetWelcomeTextAndMusic"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/resetWelcomeTextAndMusic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetWelcomeTextAndMusicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置欢迎语信息
//
// @param request - ResetWelcomeTextAndMusicRequest
//
// @return ResetWelcomeTextAndMusicResponse
func (client *Client) ResetWelcomeTextAndMusic(request *ResetWelcomeTextAndMusicRequest) (_result *ResetWelcomeTextAndMusicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ResetWelcomeTextAndMusicHeaders{}
	_result = &ResetWelcomeTextAndMusicResponse{}
	_body, _err := client.ResetWelcomeTextAndMusicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 退房
//
// @param tmpReq - RoomCheckOutRequest
//
// @param headers - RoomCheckOutHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RoomCheckOutResponse
func (client *Client) RoomCheckOutWithOptions(tmpReq *RoomCheckOutRequest, headers *RoomCheckOutHeaders, runtime *util.RuntimeOptions) (_result *RoomCheckOutResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RoomCheckOutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RoomCheckOut"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/roomCheckOut"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RoomCheckOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 退房
//
// @param request - RoomCheckOutRequest
//
// @return RoomCheckOutResponse
func (client *Client) RoomCheckOut(request *RoomCheckOutRequest) (_result *RoomCheckOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RoomCheckOutHeaders{}
	_result = &RoomCheckOutResponse{}
	_body, _err := client.RoomCheckOutWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交酒店订单
//
// @param tmpReq - SubmitHotelOrderRequest
//
// @param headers - SubmitHotelOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitHotelOrderResponse
func (client *Client) SubmitHotelOrderWithOptions(tmpReq *SubmitHotelOrderRequest, headers *SubmitHotelOrderHeaders, runtime *util.RuntimeOptions) (_result *SubmitHotelOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitHotelOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitHotelOrder"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/submitHotelOrder"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitHotelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交酒店订单
//
// @param request - SubmitHotelOrderRequest
//
// @return SubmitHotelOrderResponse
func (client *Client) SubmitHotelOrder(request *SubmitHotelOrderRequest) (_result *SubmitHotelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubmitHotelOrderHeaders{}
	_result = &SubmitHotelOrderResponse{}
	_body, _err := client.SubmitHotelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步客控设备状态到主控设备
//
// @param request - SyncDeviceStatusWithAkRequest
//
// @param headers - SyncDeviceStatusWithAkHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDeviceStatusWithAkResponse
func (client *Client) SyncDeviceStatusWithAkWithOptions(request *SyncDeviceStatusWithAkRequest, headers *SyncDeviceStatusWithAkHeaders, runtime *util.RuntimeOptions) (_result *SyncDeviceStatusWithAkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryCnName)) {
		body["CategoryCnName"] = request.CategoryCnName
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryEnName)) {
		body["CategoryEnName"] = request.CategoryEnName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.LocationCnName)) {
		body["LocationCnName"] = request.LocationCnName
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		body["Number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNo)) {
		body["RoomNo"] = request.RoomNo
	}

	if !tea.BoolValue(util.IsUnset(request.Switch)) {
		body["Switch"] = request.Switch
	}

	if !tea.BoolValue(util.IsUnset(request.FanSpeed)) {
		body["fanSpeed"] = request.FanSpeed
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RoomTemperature)) {
		body["roomTemperature"] = request.RoomTemperature
	}

	if !tea.BoolValue(util.IsUnset(request.Temperature)) {
		body["temperature"] = request.Temperature
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncDeviceStatusWithAk"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/syncDeviceStatusWithAk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncDeviceStatusWithAkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步客控设备状态到主控设备
//
// @param request - SyncDeviceStatusWithAkRequest
//
// @return SyncDeviceStatusWithAkResponse
func (client *Client) SyncDeviceStatusWithAk(request *SyncDeviceStatusWithAkRequest) (_result *SyncDeviceStatusWithAkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncDeviceStatusWithAkHeaders{}
	_result = &SyncDeviceStatusWithAkResponse{}
	_body, _err := client.SyncDeviceStatusWithAkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改基础信息问答
//
// @param request - UpdateBasicInfoQARequest
//
// @param headers - UpdateBasicInfoQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBasicInfoQAResponse
func (client *Client) UpdateBasicInfoQAWithOptions(request *UpdateBasicInfoQARequest, headers *UpdateBasicInfoQAHeaders, runtime *util.RuntimeOptions) (_result *UpdateBasicInfoQAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckInTime)) {
		body["CheckInTime"] = request.CheckInTime
	}

	if !tea.BoolValue(util.IsUnset(request.CheckOutTime)) {
		body["CheckOutTime"] = request.CheckOutTime
	}

	if !tea.BoolValue(util.IsUnset(request.HotelAddress)) {
		body["HotelAddress"] = request.HotelAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.HotelIntroduction)) {
		body["HotelIntroduction"] = request.HotelIntroduction
	}

	if !tea.BoolValue(util.IsUnset(request.HotelMember)) {
		body["HotelMember"] = request.HotelMember
	}

	if !tea.BoolValue(util.IsUnset(request.HotelService)) {
		body["HotelService"] = request.HotelService
	}

	if !tea.BoolValue(util.IsUnset(request.ParkingExpenses)) {
		body["ParkingExpenses"] = request.ParkingExpenses
	}

	if !tea.BoolValue(util.IsUnset(request.ParkingPosition)) {
		body["ParkingPosition"] = request.ParkingPosition
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.WifiName)) {
		body["WifiName"] = request.WifiName
	}

	if !tea.BoolValue(util.IsUnset(request.WifiPassword)) {
		body["WifiPassword"] = request.WifiPassword
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBasicInfoQA"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateBasicInfoQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBasicInfoQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改基础信息问答
//
// @param request - UpdateBasicInfoQARequest
//
// @return UpdateBasicInfoQAResponse
func (client *Client) UpdateBasicInfoQA(request *UpdateBasicInfoQARequest) (_result *UpdateBasicInfoQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateBasicInfoQAHeaders{}
	_result = &UpdateBasicInfoQAResponse{}
	_body, _err := client.UpdateBasicInfoQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改自定义问答
//
// @param tmpReq - UpdateCustomQARequest
//
// @param headers - UpdateCustomQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomQAResponse
func (client *Client) UpdateCustomQAWithOptions(tmpReq *UpdateCustomQARequest, headers *UpdateCustomQAHeaders, runtime *util.RuntimeOptions) (_result *UpdateCustomQAResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateCustomQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Answers)) {
		request.AnswersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Answers, tea.String("Answers"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.KeyWords)) {
		request.KeyWordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KeyWords, tea.String("KeyWords"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SupplementaryQuestions)) {
		request.SupplementaryQuestionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupplementaryQuestions, tea.String("SupplementaryQuestions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnswersShrink)) {
		body["Answers"] = request.AnswersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustomQAId)) {
		body["CustomQAId"] = request.CustomQAId
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWordsShrink)) {
		body["KeyWords"] = request.KeyWordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MajorQuestion)) {
		body["MajorQuestion"] = request.MajorQuestion
	}

	if !tea.BoolValue(util.IsUnset(request.SupplementaryQuestionsShrink)) {
		body["SupplementaryQuestions"] = request.SupplementaryQuestionsShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomQA"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateCustomQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义问答
//
// @param request - UpdateCustomQARequest
//
// @return UpdateCustomQAResponse
func (client *Client) UpdateCustomQA(request *UpdateCustomQARequest) (_result *UpdateCustomQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCustomQAHeaders{}
	_result = &UpdateCustomQAResponse{}
	_body, _err := client.UpdateCustomQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改酒店项目
//
// @param tmpReq - UpdateHotelRequest
//
// @param headers - UpdateHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotelResponse
func (client *Client) UpdateHotelWithOptions(tmpReq *UpdateHotelRequest, headers *UpdateHotelHeaders, runtime *util.RuntimeOptions) (_result *UpdateHotelResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateHotelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedPks)) {
		request.RelatedPksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedPks, tea.String("RelatedPks"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.EstOpenTime)) {
		body["EstOpenTime"] = request.EstOpenTime
	}

	if !tea.BoolValue(util.IsUnset(request.HotelAddress)) {
		body["HotelAddress"] = request.HotelAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HotelEmail)) {
		body["HotelEmail"] = request.HotelEmail
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.HotelName)) {
		body["HotelName"] = request.HotelName
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedPksShrink)) {
		body["RelatedPks"] = request.RelatedPksShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RoomNum)) {
		body["RoomNum"] = request.RoomNum
	}

	if !tea.BoolValue(util.IsUnset(request.TbOpenId)) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHotel"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateHotel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改酒店项目
//
// @param request - UpdateHotelRequest
//
// @return UpdateHotelResponse
func (client *Client) UpdateHotel(request *UpdateHotelRequest) (_result *UpdateHotelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateHotelHeaders{}
	_result = &UpdateHotelResponse{}
	_body, _err := client.UpdateHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改酒店闹钟
//
// @param tmpReq - UpdateHotelAlarmRequest
//
// @param headers - UpdateHotelAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotelAlarmResponse
func (client *Client) UpdateHotelAlarmWithOptions(tmpReq *UpdateHotelAlarmRequest, headers *UpdateHotelAlarmHeaders, runtime *util.RuntimeOptions) (_result *UpdateHotelAlarmResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateHotelAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Alarms)) {
		request.AlarmsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Alarms, tea.String("Alarms"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ScheduleInfo)) {
		request.ScheduleInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleInfo, tea.String("ScheduleInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmsShrink)) {
		body["Alarms"] = request.AlarmsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleInfoShrink)) {
		body["ScheduleInfo"] = request.ScheduleInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHotelAlarm"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateHotelAlarm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHotelAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改酒店闹钟
//
// @param request - UpdateHotelAlarmRequest
//
// @return UpdateHotelAlarmResponse
func (client *Client) UpdateHotelAlarm(request *UpdateHotelAlarmRequest) (_result *UpdateHotelAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateHotelAlarmHeaders{}
	_result = &UpdateHotelAlarmResponse{}
	_body, _err := client.UpdateHotelAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景预订编辑
//
// @param tmpReq - UpdateHotelSceneBookItemRequest
//
// @param headers - UpdateHotelSceneBookItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotelSceneBookItemResponse
func (client *Client) UpdateHotelSceneBookItemWithOptions(tmpReq *UpdateHotelSceneBookItemRequest, headers *UpdateHotelSceneBookItemHeaders, runtime *util.RuntimeOptions) (_result *UpdateHotelSceneBookItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateHotelSceneBookItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateHotelSceneBookReq)) {
		request.UpdateHotelSceneBookReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateHotelSceneBookReq, tea.String("UpdateHotelSceneBookReq"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateHotelSceneBookReqShrink)) {
		query["UpdateHotelSceneBookReq"] = request.UpdateHotelSceneBookReqShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHotelSceneBookItem"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateHotelSceneBookItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHotelSceneBookItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景预订编辑
//
// @param request - UpdateHotelSceneBookItemRequest
//
// @return UpdateHotelSceneBookItemResponse
func (client *Client) UpdateHotelSceneBookItem(request *UpdateHotelSceneBookItemRequest) (_result *UpdateHotelSceneBookItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateHotelSceneBookItemHeaders{}
	_result = &UpdateHotelSceneBookItemResponse{}
	_body, _err := client.UpdateHotelSceneBookItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景修改（开启/关闭/编辑）
//
// @param tmpReq - UpdateHotelSceneItemRequest
//
// @param headers - UpdateHotelSceneItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotelSceneItemResponse
func (client *Client) UpdateHotelSceneItemWithOptions(tmpReq *UpdateHotelSceneItemRequest, headers *UpdateHotelSceneItemHeaders, runtime *util.RuntimeOptions) (_result *UpdateHotelSceneItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateHotelSceneItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateHotelSceneOperateReq)) {
		request.UpdateHotelSceneOperateReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateHotelSceneOperateReq, tea.String("UpdateHotelSceneOperateReq"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateHotelSceneReq)) {
		request.UpdateHotelSceneReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateHotelSceneReq, tea.String("UpdateHotelSceneReq"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateHotelSceneOperateReqShrink)) {
		query["UpdateHotelSceneOperateReq"] = request.UpdateHotelSceneOperateReqShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateHotelSceneReqShrink)) {
		query["UpdateHotelSceneReq"] = request.UpdateHotelSceneReqShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHotelSceneItem"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateHotelSceneItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHotelSceneItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景修改（开启/关闭/编辑）
//
// @param request - UpdateHotelSceneItemRequest
//
// @return UpdateHotelSceneItemResponse
func (client *Client) UpdateHotelSceneItem(request *UpdateHotelSceneItemRequest) (_result *UpdateHotelSceneItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateHotelSceneItemHeaders{}
	_result = &UpdateHotelSceneItemResponse{}
	_body, _err := client.UpdateHotelSceneItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新消息通知模板
//
// @param request - UpdateMessageTemplateRequest
//
// @param headers - UpdateMessageTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageTemplateResponse
func (client *Client) UpdateMessageTemplateWithOptions(request *UpdateMessageTemplateRequest, headers *UpdateMessageTemplateHeaders, runtime *util.RuntimeOptions) (_result *UpdateMessageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateDetail)) {
		body["TemplateDetail"] = request.TemplateDetail
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["TemplateName"] = request.TemplateName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMessageTemplate"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateMessageTemplate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMessageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新消息通知模板
//
// @param request - UpdateMessageTemplateRequest
//
// @return UpdateMessageTemplateResponse
func (client *Client) UpdateMessageTemplate(request *UpdateMessageTemplateRequest) (_result *UpdateMessageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMessageTemplateHeaders{}
	_result = &UpdateMessageTemplateResponse{}
	_body, _err := client.UpdateMessageTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改酒店自定义rcu场景
//
// @param tmpReq - UpdateRcuSceneRequest
//
// @param headers - UpdateRcuSceneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRcuSceneResponse
func (client *Client) UpdateRcuSceneWithOptions(tmpReq *UpdateRcuSceneRequest, headers *UpdateRcuSceneHeaders, runtime *util.RuntimeOptions) (_result *UpdateRcuSceneResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateRcuSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SceneRelationExtDTO)) {
		request.SceneRelationExtDTOShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneRelationExtDTO, tea.String("SceneRelationExtDTO"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneRelationExtDTOShrink)) {
		body["SceneRelationExtDTO"] = request.SceneRelationExtDTOShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRcuScene"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateRcuScene"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRcuSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改酒店自定义rcu场景
//
// @param request - UpdateRcuSceneRequest
//
// @return UpdateRcuSceneResponse
func (client *Client) UpdateRcuScene(request *UpdateRcuSceneRequest) (_result *UpdateRcuSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRcuSceneHeaders{}
	_result = &UpdateRcuSceneResponse{}
	_body, _err := client.UpdateRcuSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改服务设施问答
//
// @param request - UpdateServiceQARequest
//
// @param headers - UpdateServiceQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceQAResponse
func (client *Client) UpdateServiceQAWithOptions(request *UpdateServiceQARequest, headers *UpdateServiceQAHeaders, runtime *util.RuntimeOptions) (_result *UpdateServiceQAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Answer)) {
		body["Answer"] = request.Answer
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceQAId)) {
		body["ServiceQAId"] = request.ServiceQAId
	}

	if !tea.BoolValue(util.IsUnset(request.IsActive)) {
		body["isActive"] = request.IsActive
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceQA"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateServiceQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改服务设施问答
//
// @param request - UpdateServiceQARequest
//
// @return UpdateServiceQAResponse
func (client *Client) UpdateServiceQA(request *UpdateServiceQARequest) (_result *UpdateServiceQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateServiceQAHeaders{}
	_result = &UpdateServiceQAResponse{}
	_body, _err := client.UpdateServiceQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改工单
//
// @param request - UpdateTicketRequest
//
// @param headers - UpdateTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketResponse
func (client *Client) UpdateTicketWithOptions(request *UpdateTicketRequest, headers *UpdateTicketHeaders, runtime *util.RuntimeOptions) (_result *UpdateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupKey)) {
		body["GroupKey"] = request.GroupKey
	}

	if !tea.BoolValue(util.IsUnset(request.HotelId)) {
		body["HotelId"] = request.HotelId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTicket"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/updateTicket"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改工单
//
// @param request - UpdateTicketRequest
//
// @return UpdateTicketResponse
func (client *Client) UpdateTicket(request *UpdateTicketRequest) (_result *UpdateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTicketHeaders{}
	_result = &UpdateTicketResponse{}
	_body, _err := client.UpdateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
