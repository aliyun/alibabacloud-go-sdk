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
	"io"
)

type RecognizeAdvancedRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// 是否需要去除印章功能，默认不需要。true：需要 false：不需要
	NoStamp *bool `json:"NoStamp,omitempty" xml:"NoStamp,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要图案检测功能，默认不需要。true：需要 false：不需要
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 是否需要分段功能，默认不需要。true：需要 false：不需要
	Paragraph *bool `json:"Paragraph,omitempty" xml:"Paragraph,omitempty"`
	// 是否需要成行返回功能，默认不需要
	Row *bool `json:"Row,omitempty" xml:"Row,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAdvancedRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAdvancedRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAdvancedRequest) SetNeedRotate(v bool) *RecognizeAdvancedRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetNeedSortPage(v bool) *RecognizeAdvancedRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetNoStamp(v bool) *RecognizeAdvancedRequest {
	s.NoStamp = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetOutputCharInfo(v bool) *RecognizeAdvancedRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetOutputFigure(v bool) *RecognizeAdvancedRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetOutputTable(v bool) *RecognizeAdvancedRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetParagraph(v bool) *RecognizeAdvancedRequest {
	s.Paragraph = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetRow(v bool) *RecognizeAdvancedRequest {
	s.Row = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetUrl(v string) *RecognizeAdvancedRequest {
	s.Url = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetBody(v io.Reader) *RecognizeAdvancedRequest {
	s.Body = v
	return s
}

type RecognizeAdvancedResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeAdvancedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAdvancedResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAdvancedResponseBody) SetCode(v string) *RecognizeAdvancedResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetData(v string) *RecognizeAdvancedResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetMessage(v string) *RecognizeAdvancedResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetRequestId(v string) *RecognizeAdvancedResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeAdvancedResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeAdvancedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeAdvancedResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAdvancedResponse) GoString() string {
	return s.String()
}

func (s *RecognizeAdvancedResponse) SetHeaders(v map[string]*string) *RecognizeAdvancedResponse {
	s.Headers = v
	return s
}

func (s *RecognizeAdvancedResponse) SetBody(v *RecognizeAdvancedResponseBody) *RecognizeAdvancedResponse {
	s.Body = v
	return s
}

type RecognizeAirItineraryRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAirItineraryRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAirItineraryRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAirItineraryRequest) SetUrl(v string) *RecognizeAirItineraryRequest {
	s.Url = &v
	return s
}

func (s *RecognizeAirItineraryRequest) SetBody(v io.Reader) *RecognizeAirItineraryRequest {
	s.Body = v
	return s
}

type RecognizeAirItineraryResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeAirItineraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAirItineraryResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAirItineraryResponseBody) SetCode(v string) *RecognizeAirItineraryResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetData(v string) *RecognizeAirItineraryResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetMessage(v string) *RecognizeAirItineraryResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetRequestId(v string) *RecognizeAirItineraryResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeAirItineraryResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeAirItineraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeAirItineraryResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAirItineraryResponse) GoString() string {
	return s.String()
}

func (s *RecognizeAirItineraryResponse) SetHeaders(v map[string]*string) *RecognizeAirItineraryResponse {
	s.Headers = v
	return s
}

func (s *RecognizeAirItineraryResponse) SetBody(v *RecognizeAirItineraryResponseBody) *RecognizeAirItineraryResponse {
	s.Body = v
	return s
}

type RecognizeBankAccountLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankAccountLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAccountLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankAccountLicenseRequest) SetUrl(v string) *RecognizeBankAccountLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBankAccountLicenseRequest) SetBody(v io.Reader) *RecognizeBankAccountLicenseRequest {
	s.Body = v
	return s
}

type RecognizeBankAccountLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBankAccountLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAccountLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankAccountLicenseResponseBody) SetCode(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetData(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetMessage(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetRequestId(v string) *RecognizeBankAccountLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBankAccountLicenseResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBankAccountLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBankAccountLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAccountLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBankAccountLicenseResponse) SetHeaders(v map[string]*string) *RecognizeBankAccountLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBankAccountLicenseResponse) SetBody(v *RecognizeBankAccountLicenseResponseBody) *RecognizeBankAccountLicenseResponse {
	s.Body = v
	return s
}

type RecognizeBankCardRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardRequest) SetUrl(v string) *RecognizeBankCardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBankCardRequest) SetBody(v io.Reader) *RecognizeBankCardRequest {
	s.Body = v
	return s
}

type RecognizeBankCardResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBankCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBody) SetCode(v string) *RecognizeBankCardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetData(v string) *RecognizeBankCardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetMessage(v string) *RecognizeBankCardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetRequestId(v string) *RecognizeBankCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBankCardResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBankCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBankCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponse) SetHeaders(v map[string]*string) *RecognizeBankCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBankCardResponse) SetBody(v *RecognizeBankCardResponseBody) *RecognizeBankCardResponse {
	s.Body = v
	return s
}

type RecognizeBasicRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBasicRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBasicRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBasicRequest) SetUrl(v string) *RecognizeBasicRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBasicRequest) SetBody(v io.Reader) *RecognizeBasicRequest {
	s.Body = v
	return s
}

type RecognizeBasicResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBasicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBasicResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBasicResponseBody) SetCode(v string) *RecognizeBasicResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetData(v string) *RecognizeBasicResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetMessage(v string) *RecognizeBasicResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetRequestId(v string) *RecognizeBasicResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBasicResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBasicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBasicResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBasicResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBasicResponse) SetHeaders(v map[string]*string) *RecognizeBasicResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBasicResponse) SetBody(v *RecognizeBasicResponseBody) *RecognizeBasicResponse {
	s.Body = v
	return s
}

type RecognizeBatchRecognizeRequest struct {
	// 图片名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 图片识别op类型
	ImageOp *string `json:"ImageOp,omitempty" xml:"ImageOp,omitempty"`
	// 图片存入oss中的key
	ImageOssKey *string `json:"ImageOssKey,omitempty" xml:"ImageOssKey,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeBatchRecognizeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBatchRecognizeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBatchRecognizeRequest) SetImageName(v string) *RecognizeBatchRecognizeRequest {
	s.ImageName = &v
	return s
}

func (s *RecognizeBatchRecognizeRequest) SetImageOp(v string) *RecognizeBatchRecognizeRequest {
	s.ImageOp = &v
	return s
}

func (s *RecognizeBatchRecognizeRequest) SetImageOssKey(v string) *RecognizeBatchRecognizeRequest {
	s.ImageOssKey = &v
	return s
}

func (s *RecognizeBatchRecognizeRequest) SetNeedRotate(v bool) *RecognizeBatchRecognizeRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeBatchRecognizeRequest) SetNeedSortPage(v bool) *RecognizeBatchRecognizeRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeBatchRecognizeRequest) SetOutputCharInfo(v bool) *RecognizeBatchRecognizeRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeBatchRecognizeRequest) SetOutputTable(v bool) *RecognizeBatchRecognizeRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeBatchRecognizeRequest) SetUrl(v string) *RecognizeBatchRecognizeRequest {
	s.Url = &v
	return s
}

type RecognizeBatchRecognizeResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBatchRecognizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBatchRecognizeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBatchRecognizeResponseBody) SetCode(v string) *RecognizeBatchRecognizeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBatchRecognizeResponseBody) SetData(v string) *RecognizeBatchRecognizeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBatchRecognizeResponseBody) SetMessage(v string) *RecognizeBatchRecognizeResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBatchRecognizeResponseBody) SetRequestId(v string) *RecognizeBatchRecognizeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBatchRecognizeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBatchRecognizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBatchRecognizeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBatchRecognizeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBatchRecognizeResponse) SetHeaders(v map[string]*string) *RecognizeBatchRecognizeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBatchRecognizeResponse) SetBody(v *RecognizeBatchRecognizeResponseBody) *RecognizeBatchRecognizeResponse {
	s.Body = v
	return s
}

type RecognizeBirthCertificationRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBirthCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBirthCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBirthCertificationRequest) SetUrl(v string) *RecognizeBirthCertificationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBirthCertificationRequest) SetBody(v io.Reader) *RecognizeBirthCertificationRequest {
	s.Body = v
	return s
}

type RecognizeBirthCertificationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBirthCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBirthCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBirthCertificationResponseBody) SetCode(v string) *RecognizeBirthCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetData(v string) *RecognizeBirthCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetMessage(v string) *RecognizeBirthCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetRequestId(v string) *RecognizeBirthCertificationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBirthCertificationResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBirthCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBirthCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBirthCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBirthCertificationResponse) SetHeaders(v map[string]*string) *RecognizeBirthCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBirthCertificationResponse) SetBody(v *RecognizeBirthCertificationResponseBody) *RecognizeBirthCertificationResponse {
	s.Body = v
	return s
}

type RecognizeBusShipTicketRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBusShipTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusShipTicketRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusShipTicketRequest) SetUrl(v string) *RecognizeBusShipTicketRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBusShipTicketRequest) SetBody(v io.Reader) *RecognizeBusShipTicketRequest {
	s.Body = v
	return s
}

type RecognizeBusShipTicketResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBusShipTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusShipTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusShipTicketResponseBody) SetCode(v string) *RecognizeBusShipTicketResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBusShipTicketResponseBody) SetData(v string) *RecognizeBusShipTicketResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBusShipTicketResponseBody) SetMessage(v string) *RecognizeBusShipTicketResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBusShipTicketResponseBody) SetRequestId(v string) *RecognizeBusShipTicketResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBusShipTicketResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBusShipTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBusShipTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusShipTicketResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBusShipTicketResponse) SetHeaders(v map[string]*string) *RecognizeBusShipTicketResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBusShipTicketResponse) SetBody(v *RecognizeBusShipTicketResponseBody) *RecognizeBusShipTicketResponse {
	s.Body = v
	return s
}

type RecognizeBusinessLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBusinessLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseRequest) SetUrl(v string) *RecognizeBusinessLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBusinessLicenseRequest) SetBody(v io.Reader) *RecognizeBusinessLicenseRequest {
	s.Body = v
	return s
}

type RecognizeBusinessLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBody) SetCode(v string) *RecognizeBusinessLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetData(v string) *RecognizeBusinessLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetMessage(v string) *RecognizeBusinessLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetRequestId(v string) *RecognizeBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBusinessLicenseResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBusinessLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponse) SetHeaders(v map[string]*string) *RecognizeBusinessLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBusinessLicenseResponse) SetBody(v *RecognizeBusinessLicenseResponseBody) *RecognizeBusinessLicenseResponse {
	s.Body = v
	return s
}

type RecognizeCarInvoiceRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCarInvoiceRequest) SetUrl(v string) *RecognizeCarInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCarInvoiceRequest) SetBody(v io.Reader) *RecognizeCarInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeCarInvoiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCarInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarInvoiceResponseBody) SetCode(v string) *RecognizeCarInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetData(v string) *RecognizeCarInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetMessage(v string) *RecognizeCarInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetRequestId(v string) *RecognizeCarInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCarInvoiceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeCarInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeCarInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCarInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeCarInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCarInvoiceResponse) SetBody(v *RecognizeCarInvoiceResponseBody) *RecognizeCarInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeCarNumberRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarNumberRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCarNumberRequest) SetUrl(v string) *RecognizeCarNumberRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCarNumberRequest) SetBody(v io.Reader) *RecognizeCarNumberRequest {
	s.Body = v
	return s
}

type RecognizeCarNumberResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCarNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarNumberResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarNumberResponseBody) SetCode(v string) *RecognizeCarNumberResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetData(v string) *RecognizeCarNumberResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetMessage(v string) *RecognizeCarNumberResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetRequestId(v string) *RecognizeCarNumberResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCarNumberResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeCarNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeCarNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarNumberResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCarNumberResponse) SetHeaders(v map[string]*string) *RecognizeCarNumberResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCarNumberResponse) SetBody(v *RecognizeCarNumberResponseBody) *RecognizeCarNumberResponse {
	s.Body = v
	return s
}

type RecognizeCarVinCodeRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarVinCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarVinCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCarVinCodeRequest) SetUrl(v string) *RecognizeCarVinCodeRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCarVinCodeRequest) SetBody(v io.Reader) *RecognizeCarVinCodeRequest {
	s.Body = v
	return s
}

type RecognizeCarVinCodeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCarVinCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarVinCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarVinCodeResponseBody) SetCode(v string) *RecognizeCarVinCodeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetData(v string) *RecognizeCarVinCodeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetMessage(v string) *RecognizeCarVinCodeResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetRequestId(v string) *RecognizeCarVinCodeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCarVinCodeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeCarVinCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeCarVinCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarVinCodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCarVinCodeResponse) SetHeaders(v map[string]*string) *RecognizeCarVinCodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCarVinCodeResponse) SetBody(v *RecognizeCarVinCodeResponseBody) *RecognizeCarVinCodeResponse {
	s.Body = v
	return s
}

type RecognizeChinesePassportRequest struct {
	// 是否需要图案检测功能，默认需要
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeChinesePassportRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinesePassportRequest) GoString() string {
	return s.String()
}

func (s *RecognizeChinesePassportRequest) SetOutputFigure(v bool) *RecognizeChinesePassportRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeChinesePassportRequest) SetUrl(v string) *RecognizeChinesePassportRequest {
	s.Url = &v
	return s
}

func (s *RecognizeChinesePassportRequest) SetBody(v io.Reader) *RecognizeChinesePassportRequest {
	s.Body = v
	return s
}

type RecognizeChinesePassportResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeChinesePassportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinesePassportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeChinesePassportResponseBody) SetCode(v string) *RecognizeChinesePassportResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeChinesePassportResponseBody) SetData(v string) *RecognizeChinesePassportResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeChinesePassportResponseBody) SetMessage(v string) *RecognizeChinesePassportResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeChinesePassportResponseBody) SetRequestId(v string) *RecognizeChinesePassportResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeChinesePassportResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeChinesePassportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeChinesePassportResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinesePassportResponse) GoString() string {
	return s.String()
}

func (s *RecognizeChinesePassportResponse) SetHeaders(v map[string]*string) *RecognizeChinesePassportResponse {
	s.Headers = v
	return s
}

func (s *RecognizeChinesePassportResponse) SetBody(v *RecognizeChinesePassportResponseBody) *RecognizeChinesePassportResponse {
	s.Body = v
	return s
}

type RecognizeCommonPrintedInvoiceRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCommonPrintedInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCommonPrintedInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCommonPrintedInvoiceRequest) SetUrl(v string) *RecognizeCommonPrintedInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceRequest) SetBody(v io.Reader) *RecognizeCommonPrintedInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeCommonPrintedInvoiceResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCommonPrintedInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCommonPrintedInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCommonPrintedInvoiceResponseBody) SetCode(v string) *RecognizeCommonPrintedInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponseBody) SetData(v string) *RecognizeCommonPrintedInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponseBody) SetMessage(v string) *RecognizeCommonPrintedInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponseBody) SetRequestId(v string) *RecognizeCommonPrintedInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCommonPrintedInvoiceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeCommonPrintedInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeCommonPrintedInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCommonPrintedInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCommonPrintedInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeCommonPrintedInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponse) SetBody(v *RecognizeCommonPrintedInvoiceResponseBody) *RecognizeCommonPrintedInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeCtwoMedicalDeviceManageLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCtwoMedicalDeviceManageLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCtwoMedicalDeviceManageLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseRequest) SetUrl(v string) *RecognizeCtwoMedicalDeviceManageLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseRequest) SetBody(v io.Reader) *RecognizeCtwoMedicalDeviceManageLicenseRequest {
	s.Body = v
	return s
}

type RecognizeCtwoMedicalDeviceManageLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetCode(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetData(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetMessage(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetRequestId(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCtwoMedicalDeviceManageLicenseResponse struct {
	Headers map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeCtwoMedicalDeviceManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponse) SetHeaders(v map[string]*string) *RecognizeCtwoMedicalDeviceManageLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponse) SetBody(v *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) *RecognizeCtwoMedicalDeviceManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeDeleteExcelRecordRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RecognizeDeleteExcelRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDeleteExcelRecordRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDeleteExcelRecordRequest) SetId(v string) *RecognizeDeleteExcelRecordRequest {
	s.Id = &v
	return s
}

type RecognizeDeleteExcelRecordResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDeleteExcelRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDeleteExcelRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDeleteExcelRecordResponseBody) SetCode(v string) *RecognizeDeleteExcelRecordResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeDeleteExcelRecordResponseBody) SetData(v string) *RecognizeDeleteExcelRecordResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeDeleteExcelRecordResponseBody) SetMessage(v string) *RecognizeDeleteExcelRecordResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeDeleteExcelRecordResponseBody) SetRequestId(v string) *RecognizeDeleteExcelRecordResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeDeleteExcelRecordResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeDeleteExcelRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeDeleteExcelRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDeleteExcelRecordResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDeleteExcelRecordResponse) SetHeaders(v map[string]*string) *RecognizeDeleteExcelRecordResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDeleteExcelRecordResponse) SetBody(v *RecognizeDeleteExcelRecordResponseBody) *RecognizeDeleteExcelRecordResponse {
	s.Body = v
	return s
}

type RecognizeDrivingLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeDrivingLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseRequest) SetUrl(v string) *RecognizeDrivingLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeDrivingLicenseRequest) SetBody(v io.Reader) *RecognizeDrivingLicenseRequest {
	s.Body = v
	return s
}

type RecognizeDrivingLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBody) SetCode(v string) *RecognizeDrivingLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetData(v string) *RecognizeDrivingLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetMessage(v string) *RecognizeDrivingLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetRequestId(v string) *RecognizeDrivingLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeDrivingLicenseResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeDrivingLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeDrivingLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponse) SetHeaders(v map[string]*string) *RecognizeDrivingLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDrivingLicenseResponse) SetBody(v *RecognizeDrivingLicenseResponseBody) *RecognizeDrivingLicenseResponse {
	s.Body = v
	return s
}

type RecognizeEduFormulaRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduFormulaRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduFormulaRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduFormulaRequest) SetUrl(v string) *RecognizeEduFormulaRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduFormulaRequest) SetBody(v io.Reader) *RecognizeEduFormulaRequest {
	s.Body = v
	return s
}

type RecognizeEduFormulaResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduFormulaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduFormulaResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduFormulaResponseBody) SetCode(v string) *RecognizeEduFormulaResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetData(v string) *RecognizeEduFormulaResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetMessage(v string) *RecognizeEduFormulaResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetRequestId(v string) *RecognizeEduFormulaResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduFormulaResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeEduFormulaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeEduFormulaResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduFormulaResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduFormulaResponse) SetHeaders(v map[string]*string) *RecognizeEduFormulaResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduFormulaResponse) SetBody(v *RecognizeEduFormulaResponseBody) *RecognizeEduFormulaResponse {
	s.Body = v
	return s
}

type RecognizeEduOralCalculationRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduOralCalculationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduOralCalculationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduOralCalculationRequest) SetUrl(v string) *RecognizeEduOralCalculationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduOralCalculationRequest) SetBody(v io.Reader) *RecognizeEduOralCalculationRequest {
	s.Body = v
	return s
}

type RecognizeEduOralCalculationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduOralCalculationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduOralCalculationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduOralCalculationResponseBody) SetCode(v string) *RecognizeEduOralCalculationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetData(v string) *RecognizeEduOralCalculationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetMessage(v string) *RecognizeEduOralCalculationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetRequestId(v string) *RecognizeEduOralCalculationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduOralCalculationResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeEduOralCalculationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeEduOralCalculationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduOralCalculationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduOralCalculationResponse) SetHeaders(v map[string]*string) *RecognizeEduOralCalculationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduOralCalculationResponse) SetBody(v *RecognizeEduOralCalculationResponseBody) *RecognizeEduOralCalculationResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperCutRequest struct {
	// 切题类型
	CutType *string `json:"CutType,omitempty" xml:"CutType,omitempty"`
	// 图片类型
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// 年级学科
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperCutRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperCutRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperCutRequest) SetCutType(v string) *RecognizeEduPaperCutRequest {
	s.CutType = &v
	return s
}

func (s *RecognizeEduPaperCutRequest) SetImageType(v string) *RecognizeEduPaperCutRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeEduPaperCutRequest) SetSubject(v string) *RecognizeEduPaperCutRequest {
	s.Subject = &v
	return s
}

func (s *RecognizeEduPaperCutRequest) SetUrl(v string) *RecognizeEduPaperCutRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduPaperCutRequest) SetBody(v io.Reader) *RecognizeEduPaperCutRequest {
	s.Body = v
	return s
}

type RecognizeEduPaperCutResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduPaperCutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperCutResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperCutResponseBody) SetCode(v string) *RecognizeEduPaperCutResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetData(v string) *RecognizeEduPaperCutResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetMessage(v string) *RecognizeEduPaperCutResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetRequestId(v string) *RecognizeEduPaperCutResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduPaperCutResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeEduPaperCutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeEduPaperCutResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperCutResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperCutResponse) SetHeaders(v map[string]*string) *RecognizeEduPaperCutResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduPaperCutResponse) SetBody(v *RecognizeEduPaperCutResponseBody) *RecognizeEduPaperCutResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperOcrRequest struct {
	// 图片类型
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// 是否输出原图坐标信息(如果图片被做过旋转，图片校正等处理)
	OutputOricoord *bool `json:"OutputOricoord,omitempty" xml:"OutputOricoord,omitempty"`
	// 年级学科
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperOcrRequest) SetImageType(v string) *RecognizeEduPaperOcrRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetOutputOricoord(v bool) *RecognizeEduPaperOcrRequest {
	s.OutputOricoord = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetSubject(v string) *RecognizeEduPaperOcrRequest {
	s.Subject = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetUrl(v string) *RecognizeEduPaperOcrRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetBody(v io.Reader) *RecognizeEduPaperOcrRequest {
	s.Body = v
	return s
}

type RecognizeEduPaperOcrResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduPaperOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperOcrResponseBody) SetCode(v string) *RecognizeEduPaperOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetData(v string) *RecognizeEduPaperOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetMessage(v string) *RecognizeEduPaperOcrResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetRequestId(v string) *RecognizeEduPaperOcrResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduPaperOcrResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeEduPaperOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeEduPaperOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperOcrResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperOcrResponse) SetHeaders(v map[string]*string) *RecognizeEduPaperOcrResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduPaperOcrResponse) SetBody(v *RecognizeEduPaperOcrResponseBody) *RecognizeEduPaperOcrResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperStructedRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 学科类型
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperStructedRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperStructedRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperStructedRequest) SetNeedRotate(v bool) *RecognizeEduPaperStructedRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeEduPaperStructedRequest) SetSubject(v string) *RecognizeEduPaperStructedRequest {
	s.Subject = &v
	return s
}

func (s *RecognizeEduPaperStructedRequest) SetUrl(v string) *RecognizeEduPaperStructedRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduPaperStructedRequest) SetBody(v io.Reader) *RecognizeEduPaperStructedRequest {
	s.Body = v
	return s
}

type RecognizeEduPaperStructedResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduPaperStructedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperStructedResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperStructedResponseBody) SetCode(v string) *RecognizeEduPaperStructedResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetData(v string) *RecognizeEduPaperStructedResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetMessage(v string) *RecognizeEduPaperStructedResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetRequestId(v string) *RecognizeEduPaperStructedResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduPaperStructedResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeEduPaperStructedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeEduPaperStructedResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperStructedResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperStructedResponse) SetHeaders(v map[string]*string) *RecognizeEduPaperStructedResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduPaperStructedResponse) SetBody(v *RecognizeEduPaperStructedResponseBody) *RecognizeEduPaperStructedResponse {
	s.Body = v
	return s
}

type RecognizeEduQuestionOcrRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduQuestionOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduQuestionOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduQuestionOcrRequest) SetNeedRotate(v bool) *RecognizeEduQuestionOcrRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeEduQuestionOcrRequest) SetUrl(v string) *RecognizeEduQuestionOcrRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduQuestionOcrRequest) SetBody(v io.Reader) *RecognizeEduQuestionOcrRequest {
	s.Body = v
	return s
}

type RecognizeEduQuestionOcrResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduQuestionOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduQuestionOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduQuestionOcrResponseBody) SetCode(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetData(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetMessage(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetRequestId(v string) *RecognizeEduQuestionOcrResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduQuestionOcrResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeEduQuestionOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeEduQuestionOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduQuestionOcrResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduQuestionOcrResponse) SetHeaders(v map[string]*string) *RecognizeEduQuestionOcrResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduQuestionOcrResponse) SetBody(v *RecognizeEduQuestionOcrResponseBody) *RecognizeEduQuestionOcrResponse {
	s.Body = v
	return s
}

type RecognizeEnglishRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEnglishRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEnglishRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEnglishRequest) SetNeedRotate(v bool) *RecognizeEnglishRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeEnglishRequest) SetOutputTable(v bool) *RecognizeEnglishRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeEnglishRequest) SetUrl(v string) *RecognizeEnglishRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEnglishRequest) SetBody(v io.Reader) *RecognizeEnglishRequest {
	s.Body = v
	return s
}

type RecognizeEnglishResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEnglishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEnglishResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEnglishResponseBody) SetCode(v string) *RecognizeEnglishResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetData(v string) *RecognizeEnglishResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetMessage(v string) *RecognizeEnglishResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetRequestId(v string) *RecognizeEnglishResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEnglishResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeEnglishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeEnglishResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEnglishResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEnglishResponse) SetHeaders(v map[string]*string) *RecognizeEnglishResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEnglishResponse) SetBody(v *RecognizeEnglishResponseBody) *RecognizeEnglishResponse {
	s.Body = v
	return s
}

type RecognizeEstateCertificationRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEstateCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEstateCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEstateCertificationRequest) SetUrl(v string) *RecognizeEstateCertificationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEstateCertificationRequest) SetBody(v io.Reader) *RecognizeEstateCertificationRequest {
	s.Body = v
	return s
}

type RecognizeEstateCertificationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEstateCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEstateCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEstateCertificationResponseBody) SetCode(v string) *RecognizeEstateCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetData(v string) *RecognizeEstateCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetMessage(v string) *RecognizeEstateCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetRequestId(v string) *RecognizeEstateCertificationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEstateCertificationResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeEstateCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeEstateCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEstateCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEstateCertificationResponse) SetHeaders(v map[string]*string) *RecognizeEstateCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEstateCertificationResponse) SetBody(v *RecognizeEstateCertificationResponseBody) *RecognizeEstateCertificationResponse {
	s.Body = v
	return s
}

type RecognizeExcelExportRequest struct {
	// 文件名称
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 图片识别op类型
	ImageOp *string `json:"ImageOp,omitempty" xml:"ImageOp,omitempty"`
	// 识别图片数量
	OcrImageCount *int64 `json:"OcrImageCount,omitempty" xml:"OcrImageCount,omitempty"`
	// 图片识别结果集
	OcrResult *string `json:"OcrResult,omitempty" xml:"OcrResult,omitempty"`
	// 票证类型
	OcrType *string `json:"OcrType,omitempty" xml:"OcrType,omitempty"`
}

func (s RecognizeExcelExportRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExcelExportRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExcelExportRequest) SetFileName(v string) *RecognizeExcelExportRequest {
	s.FileName = &v
	return s
}

func (s *RecognizeExcelExportRequest) SetImageOp(v string) *RecognizeExcelExportRequest {
	s.ImageOp = &v
	return s
}

func (s *RecognizeExcelExportRequest) SetOcrImageCount(v int64) *RecognizeExcelExportRequest {
	s.OcrImageCount = &v
	return s
}

func (s *RecognizeExcelExportRequest) SetOcrResult(v string) *RecognizeExcelExportRequest {
	s.OcrResult = &v
	return s
}

func (s *RecognizeExcelExportRequest) SetOcrType(v string) *RecognizeExcelExportRequest {
	s.OcrType = &v
	return s
}

type RecognizeExcelExportResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeExcelExportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExcelExportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExcelExportResponseBody) SetCode(v string) *RecognizeExcelExportResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeExcelExportResponseBody) SetData(v string) *RecognizeExcelExportResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeExcelExportResponseBody) SetMessage(v string) *RecognizeExcelExportResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeExcelExportResponseBody) SetRequestId(v string) *RecognizeExcelExportResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeExcelExportResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeExcelExportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeExcelExportResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExcelExportResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExcelExportResponse) SetHeaders(v map[string]*string) *RecognizeExcelExportResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExcelExportResponse) SetBody(v *RecognizeExcelExportResponseBody) *RecognizeExcelExportResponse {
	s.Body = v
	return s
}

type RecognizeExcelRecordRequest struct {
	// 页码
	CurrPage *int64 `json:"CurrPage,omitempty" xml:"CurrPage,omitempty"`
	// 每页数据
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s RecognizeExcelRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExcelRecordRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExcelRecordRequest) SetCurrPage(v int64) *RecognizeExcelRecordRequest {
	s.CurrPage = &v
	return s
}

func (s *RecognizeExcelRecordRequest) SetPageSize(v int64) *RecognizeExcelRecordRequest {
	s.PageSize = &v
	return s
}

type RecognizeExcelRecordResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeExcelRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExcelRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExcelRecordResponseBody) SetCode(v string) *RecognizeExcelRecordResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeExcelRecordResponseBody) SetData(v string) *RecognizeExcelRecordResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeExcelRecordResponseBody) SetMessage(v string) *RecognizeExcelRecordResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeExcelRecordResponseBody) SetRequestId(v string) *RecognizeExcelRecordResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeExcelRecordResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeExcelRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeExcelRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExcelRecordResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExcelRecordResponse) SetHeaders(v map[string]*string) *RecognizeExcelRecordResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExcelRecordResponse) SetBody(v *RecognizeExcelRecordResponseBody) *RecognizeExcelRecordResponse {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToHKRequest struct {
	// 图案坐标信息输出，针对结构化，如身份证人脸头像
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeExitEntryPermitToHKRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToHKRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToHKRequest) SetOutputFigure(v bool) *RecognizeExitEntryPermitToHKRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKRequest) SetUrl(v string) *RecognizeExitEntryPermitToHKRequest {
	s.Url = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKRequest) SetBody(v io.Reader) *RecognizeExitEntryPermitToHKRequest {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToHKResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeExitEntryPermitToHKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToHKResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToHKResponseBody) SetCode(v string) *RecognizeExitEntryPermitToHKResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponseBody) SetData(v string) *RecognizeExitEntryPermitToHKResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponseBody) SetMessage(v string) *RecognizeExitEntryPermitToHKResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponseBody) SetRequestId(v string) *RecognizeExitEntryPermitToHKResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeExitEntryPermitToHKResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeExitEntryPermitToHKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeExitEntryPermitToHKResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToHKResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToHKResponse) SetHeaders(v map[string]*string) *RecognizeExitEntryPermitToHKResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponse) SetBody(v *RecognizeExitEntryPermitToHKResponseBody) *RecognizeExitEntryPermitToHKResponse {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToMainlandRequest struct {
	// 图案坐标信息输出，针对结构化，如身份证人脸头像
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeExitEntryPermitToMainlandRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToMainlandRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToMainlandRequest) SetOutputFigure(v bool) *RecognizeExitEntryPermitToMainlandRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandRequest) SetUrl(v string) *RecognizeExitEntryPermitToMainlandRequest {
	s.Url = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandRequest) SetBody(v io.Reader) *RecognizeExitEntryPermitToMainlandRequest {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToMainlandResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeExitEntryPermitToMainlandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToMainlandResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToMainlandResponseBody) SetCode(v string) *RecognizeExitEntryPermitToMainlandResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponseBody) SetData(v string) *RecognizeExitEntryPermitToMainlandResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponseBody) SetMessage(v string) *RecognizeExitEntryPermitToMainlandResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponseBody) SetRequestId(v string) *RecognizeExitEntryPermitToMainlandResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeExitEntryPermitToMainlandResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeExitEntryPermitToMainlandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeExitEntryPermitToMainlandResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToMainlandResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToMainlandResponse) SetHeaders(v map[string]*string) *RecognizeExitEntryPermitToMainlandResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponse) SetBody(v *RecognizeExitEntryPermitToMainlandResponseBody) *RecognizeExitEntryPermitToMainlandResponse {
	s.Body = v
	return s
}

type RecognizeFoodManageLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeFoodManageLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodManageLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFoodManageLicenseRequest) SetUrl(v string) *RecognizeFoodManageLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeFoodManageLicenseRequest) SetBody(v io.Reader) *RecognizeFoodManageLicenseRequest {
	s.Body = v
	return s
}

type RecognizeFoodManageLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFoodManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFoodManageLicenseResponseBody) SetCode(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetData(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetMessage(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetRequestId(v string) *RecognizeFoodManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeFoodManageLicenseResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeFoodManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFoodManageLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodManageLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFoodManageLicenseResponse) SetHeaders(v map[string]*string) *RecognizeFoodManageLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFoodManageLicenseResponse) SetBody(v *RecognizeFoodManageLicenseResponseBody) *RecognizeFoodManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeFoodProduceLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeFoodProduceLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodProduceLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFoodProduceLicenseRequest) SetUrl(v string) *RecognizeFoodProduceLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeFoodProduceLicenseRequest) SetBody(v io.Reader) *RecognizeFoodProduceLicenseRequest {
	s.Body = v
	return s
}

type RecognizeFoodProduceLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFoodProduceLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodProduceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetCode(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetData(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetMessage(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetRequestId(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeFoodProduceLicenseResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeFoodProduceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFoodProduceLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodProduceLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFoodProduceLicenseResponse) SetHeaders(v map[string]*string) *RecognizeFoodProduceLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFoodProduceLicenseResponse) SetBody(v *RecognizeFoodProduceLicenseResponseBody) *RecognizeFoodProduceLicenseResponse {
	s.Body = v
	return s
}

type RecognizeGeneralRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeGeneralRequest) GoString() string {
	return s.String()
}

func (s *RecognizeGeneralRequest) SetUrl(v string) *RecognizeGeneralRequest {
	s.Url = &v
	return s
}

func (s *RecognizeGeneralRequest) SetBody(v io.Reader) *RecognizeGeneralRequest {
	s.Body = v
	return s
}

type RecognizeGeneralResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeGeneralResponseBody) SetCode(v string) *RecognizeGeneralResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetData(v string) *RecognizeGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetMessage(v string) *RecognizeGeneralResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetRequestId(v string) *RecognizeGeneralResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeGeneralResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeGeneralResponse) GoString() string {
	return s.String()
}

func (s *RecognizeGeneralResponse) SetHeaders(v map[string]*string) *RecognizeGeneralResponse {
	s.Headers = v
	return s
}

func (s *RecognizeGeneralResponse) SetBody(v *RecognizeGeneralResponseBody) *RecognizeGeneralResponse {
	s.Body = v
	return s
}

type RecognizeHandwritingRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHandwritingRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandwritingRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHandwritingRequest) SetNeedRotate(v bool) *RecognizeHandwritingRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetNeedSortPage(v bool) *RecognizeHandwritingRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetOutputCharInfo(v bool) *RecognizeHandwritingRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetOutputTable(v bool) *RecognizeHandwritingRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetUrl(v string) *RecognizeHandwritingRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetBody(v io.Reader) *RecognizeHandwritingRequest {
	s.Body = v
	return s
}

type RecognizeHandwritingResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHandwritingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandwritingResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHandwritingResponseBody) SetCode(v string) *RecognizeHandwritingResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetData(v string) *RecognizeHandwritingResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetMessage(v string) *RecognizeHandwritingResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetRequestId(v string) *RecognizeHandwritingResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHandwritingResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeHandwritingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeHandwritingResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandwritingResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHandwritingResponse) SetHeaders(v map[string]*string) *RecognizeHandwritingResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHandwritingResponse) SetBody(v *RecognizeHandwritingResponseBody) *RecognizeHandwritingResponse {
	s.Body = v
	return s
}

type RecognizeHouseholdRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHouseholdRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseholdRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHouseholdRequest) SetUrl(v string) *RecognizeHouseholdRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHouseholdRequest) SetBody(v io.Reader) *RecognizeHouseholdRequest {
	s.Body = v
	return s
}

type RecognizeHouseholdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHouseholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseholdResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHouseholdResponseBody) SetCode(v string) *RecognizeHouseholdResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetData(v string) *RecognizeHouseholdResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetMessage(v string) *RecognizeHouseholdResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetRequestId(v string) *RecognizeHouseholdResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHouseholdResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeHouseholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeHouseholdResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseholdResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHouseholdResponse) SetHeaders(v map[string]*string) *RecognizeHouseholdResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHouseholdResponse) SetBody(v *RecognizeHouseholdResponseBody) *RecognizeHouseholdResponse {
	s.Body = v
	return s
}

type RecognizeIdcardRequest struct {
	// 是否需要图案检测功能，默认不需要
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeIdcardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdcardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIdcardRequest) SetOutputFigure(v bool) *RecognizeIdcardRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeIdcardRequest) SetUrl(v string) *RecognizeIdcardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeIdcardRequest) SetBody(v io.Reader) *RecognizeIdcardRequest {
	s.Body = v
	return s
}

type RecognizeIdcardResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeIdcardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdcardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIdcardResponseBody) SetCode(v string) *RecognizeIdcardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetData(v string) *RecognizeIdcardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetMessage(v string) *RecognizeIdcardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetRequestId(v string) *RecognizeIdcardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeIdcardResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeIdcardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeIdcardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdcardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeIdcardResponse) SetHeaders(v map[string]*string) *RecognizeIdcardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeIdcardResponse) SetBody(v *RecognizeIdcardResponseBody) *RecognizeIdcardResponse {
	s.Body = v
	return s
}

type RecognizeInvoiceRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeInvoiceRequest) SetUrl(v string) *RecognizeInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeInvoiceRequest) SetBody(v io.Reader) *RecognizeInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeInvoiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeInvoiceResponseBody) SetCode(v string) *RecognizeInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetData(v string) *RecognizeInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetMessage(v string) *RecognizeInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetRequestId(v string) *RecognizeInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeInvoiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeInvoiceResponse) SetBody(v *RecognizeInvoiceResponseBody) *RecognizeInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeJanpaneseRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeJanpaneseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeJanpaneseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeJanpaneseRequest) SetNeedRotate(v bool) *RecognizeJanpaneseRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetOutputCharInfo(v bool) *RecognizeJanpaneseRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetOutputTable(v bool) *RecognizeJanpaneseRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetUrl(v string) *RecognizeJanpaneseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetBody(v io.Reader) *RecognizeJanpaneseRequest {
	s.Body = v
	return s
}

type RecognizeJanpaneseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeJanpaneseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeJanpaneseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeJanpaneseResponseBody) SetCode(v string) *RecognizeJanpaneseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetData(v string) *RecognizeJanpaneseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetMessage(v string) *RecognizeJanpaneseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetRequestId(v string) *RecognizeJanpaneseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeJanpaneseResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeJanpaneseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeJanpaneseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeJanpaneseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeJanpaneseResponse) SetHeaders(v map[string]*string) *RecognizeJanpaneseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeJanpaneseResponse) SetBody(v *RecognizeJanpaneseResponseBody) *RecognizeJanpaneseResponse {
	s.Body = v
	return s
}

type RecognizeKoreanRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeKoreanRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeKoreanRequest) GoString() string {
	return s.String()
}

func (s *RecognizeKoreanRequest) SetNeedRotate(v bool) *RecognizeKoreanRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeKoreanRequest) SetOutputCharInfo(v bool) *RecognizeKoreanRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeKoreanRequest) SetOutputTable(v bool) *RecognizeKoreanRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeKoreanRequest) SetUrl(v string) *RecognizeKoreanRequest {
	s.Url = &v
	return s
}

func (s *RecognizeKoreanRequest) SetBody(v io.Reader) *RecognizeKoreanRequest {
	s.Body = v
	return s
}

type RecognizeKoreanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeKoreanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeKoreanResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeKoreanResponseBody) SetCode(v string) *RecognizeKoreanResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetData(v string) *RecognizeKoreanResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetMessage(v string) *RecognizeKoreanResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetRequestId(v string) *RecognizeKoreanResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeKoreanResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeKoreanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeKoreanResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeKoreanResponse) GoString() string {
	return s.String()
}

func (s *RecognizeKoreanResponse) SetHeaders(v map[string]*string) *RecognizeKoreanResponse {
	s.Headers = v
	return s
}

func (s *RecognizeKoreanResponse) SetBody(v *RecognizeKoreanResponseBody) *RecognizeKoreanResponse {
	s.Body = v
	return s
}

type RecognizeLatinRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeLatinRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLatinRequest) GoString() string {
	return s.String()
}

func (s *RecognizeLatinRequest) SetNeedRotate(v bool) *RecognizeLatinRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeLatinRequest) SetOutputCharInfo(v bool) *RecognizeLatinRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeLatinRequest) SetOutputTable(v bool) *RecognizeLatinRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeLatinRequest) SetUrl(v string) *RecognizeLatinRequest {
	s.Url = &v
	return s
}

func (s *RecognizeLatinRequest) SetBody(v io.Reader) *RecognizeLatinRequest {
	s.Body = v
	return s
}

type RecognizeLatinResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeLatinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLatinResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeLatinResponseBody) SetCode(v string) *RecognizeLatinResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetData(v string) *RecognizeLatinResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetMessage(v string) *RecognizeLatinResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetRequestId(v string) *RecognizeLatinResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeLatinResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeLatinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeLatinResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLatinResponse) GoString() string {
	return s.String()
}

func (s *RecognizeLatinResponse) SetHeaders(v map[string]*string) *RecognizeLatinResponse {
	s.Headers = v
	return s
}

func (s *RecognizeLatinResponse) SetBody(v *RecognizeLatinResponseBody) *RecognizeLatinResponse {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceManageLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMedicalDeviceManageLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceManageLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceManageLicenseRequest) SetUrl(v string) *RecognizeMedicalDeviceManageLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseRequest) SetBody(v io.Reader) *RecognizeMedicalDeviceManageLicenseRequest {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceManageLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMedicalDeviceManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetCode(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetData(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetMessage(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetRequestId(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMedicalDeviceManageLicenseResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeMedicalDeviceManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeMedicalDeviceManageLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceManageLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceManageLicenseResponse) SetHeaders(v map[string]*string) *RecognizeMedicalDeviceManageLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponse) SetBody(v *RecognizeMedicalDeviceManageLicenseResponseBody) *RecognizeMedicalDeviceManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceProduceLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMedicalDeviceProduceLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceProduceLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceProduceLicenseRequest) SetUrl(v string) *RecognizeMedicalDeviceProduceLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseRequest) SetBody(v io.Reader) *RecognizeMedicalDeviceProduceLicenseRequest {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceProduceLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMedicalDeviceProduceLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceProduceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetCode(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetData(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetMessage(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetRequestId(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMedicalDeviceProduceLicenseResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeMedicalDeviceProduceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeMedicalDeviceProduceLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceProduceLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceProduceLicenseResponse) SetHeaders(v map[string]*string) *RecognizeMedicalDeviceProduceLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponse) SetBody(v *RecognizeMedicalDeviceProduceLicenseResponseBody) *RecognizeMedicalDeviceProduceLicenseResponse {
	s.Body = v
	return s
}

type RecognizeMixedInvoicesRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMixedInvoicesRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedInvoicesRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMixedInvoicesRequest) SetUrl(v string) *RecognizeMixedInvoicesRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMixedInvoicesRequest) SetBody(v io.Reader) *RecognizeMixedInvoicesRequest {
	s.Body = v
	return s
}

type RecognizeMixedInvoicesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMixedInvoicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedInvoicesResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMixedInvoicesResponseBody) SetCode(v string) *RecognizeMixedInvoicesResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetData(v string) *RecognizeMixedInvoicesResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetMessage(v string) *RecognizeMixedInvoicesResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetRequestId(v string) *RecognizeMixedInvoicesResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMixedInvoicesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeMixedInvoicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeMixedInvoicesResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedInvoicesResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMixedInvoicesResponse) SetHeaders(v map[string]*string) *RecognizeMixedInvoicesResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMixedInvoicesResponse) SetBody(v *RecognizeMixedInvoicesResponseBody) *RecognizeMixedInvoicesResponse {
	s.Body = v
	return s
}

type RecognizeMultiLanguageRequest struct {
	// 识别语种
	Languages []*string `json:"Languages,omitempty" xml:"Languages,omitempty" type:"Repeated"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMultiLanguageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageRequest) SetLanguages(v []*string) *RecognizeMultiLanguageRequest {
	s.Languages = v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetNeedRotate(v bool) *RecognizeMultiLanguageRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetNeedSortPage(v bool) *RecognizeMultiLanguageRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetOutputCharInfo(v bool) *RecognizeMultiLanguageRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetOutputTable(v bool) *RecognizeMultiLanguageRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetUrl(v string) *RecognizeMultiLanguageRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetBody(v io.Reader) *RecognizeMultiLanguageRequest {
	s.Body = v
	return s
}

type RecognizeMultiLanguageShrinkRequest struct {
	// 识别语种
	LanguagesShrink *string `json:"Languages,omitempty" xml:"Languages,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMultiLanguageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageShrinkRequest) SetLanguagesShrink(v string) *RecognizeMultiLanguageShrinkRequest {
	s.LanguagesShrink = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetNeedRotate(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetNeedSortPage(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetOutputCharInfo(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetOutputTable(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetUrl(v string) *RecognizeMultiLanguageShrinkRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetBody(v io.Reader) *RecognizeMultiLanguageShrinkRequest {
	s.Body = v
	return s
}

type RecognizeMultiLanguageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMultiLanguageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageResponseBody) SetCode(v string) *RecognizeMultiLanguageResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetData(v string) *RecognizeMultiLanguageResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetMessage(v string) *RecognizeMultiLanguageResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetRequestId(v string) *RecognizeMultiLanguageResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMultiLanguageResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeMultiLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeMultiLanguageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageResponse) SetHeaders(v map[string]*string) *RecognizeMultiLanguageResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMultiLanguageResponse) SetBody(v *RecognizeMultiLanguageResponseBody) *RecognizeMultiLanguageResponse {
	s.Body = v
	return s
}

type RecognizePassportRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePassportRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportRequest) GoString() string {
	return s.String()
}

func (s *RecognizePassportRequest) SetUrl(v string) *RecognizePassportRequest {
	s.Url = &v
	return s
}

func (s *RecognizePassportRequest) SetBody(v io.Reader) *RecognizePassportRequest {
	s.Body = v
	return s
}

type RecognizePassportResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePassportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePassportResponseBody) SetCode(v string) *RecognizePassportResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizePassportResponseBody) SetData(v string) *RecognizePassportResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizePassportResponseBody) SetMessage(v string) *RecognizePassportResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizePassportResponseBody) SetRequestId(v string) *RecognizePassportResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePassportResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizePassportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizePassportResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportResponse) GoString() string {
	return s.String()
}

func (s *RecognizePassportResponse) SetHeaders(v map[string]*string) *RecognizePassportResponse {
	s.Headers = v
	return s
}

func (s *RecognizePassportResponse) SetBody(v *RecognizePassportResponseBody) *RecognizePassportResponse {
	s.Body = v
	return s
}

type RecognizeQuotaInvoiceRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeQuotaInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceRequest) SetUrl(v string) *RecognizeQuotaInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeQuotaInvoiceRequest) SetBody(v io.Reader) *RecognizeQuotaInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeQuotaInvoiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBody) SetCode(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetData(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetMessage(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetRequestId(v string) *RecognizeQuotaInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeQuotaInvoiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeQuotaInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeQuotaInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeQuotaInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) SetBody(v *RecognizeQuotaInvoiceResponseBody) *RecognizeQuotaInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeRideHailingItineraryRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRideHailingItineraryRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRideHailingItineraryRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRideHailingItineraryRequest) SetUrl(v string) *RecognizeRideHailingItineraryRequest {
	s.Url = &v
	return s
}

func (s *RecognizeRideHailingItineraryRequest) SetBody(v io.Reader) *RecognizeRideHailingItineraryRequest {
	s.Body = v
	return s
}

type RecognizeRideHailingItineraryResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeRideHailingItineraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRideHailingItineraryResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRideHailingItineraryResponseBody) SetCode(v string) *RecognizeRideHailingItineraryResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeRideHailingItineraryResponseBody) SetData(v string) *RecognizeRideHailingItineraryResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeRideHailingItineraryResponseBody) SetMessage(v string) *RecognizeRideHailingItineraryResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeRideHailingItineraryResponseBody) SetRequestId(v string) *RecognizeRideHailingItineraryResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeRideHailingItineraryResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeRideHailingItineraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeRideHailingItineraryResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRideHailingItineraryResponse) GoString() string {
	return s.String()
}

func (s *RecognizeRideHailingItineraryResponse) SetHeaders(v map[string]*string) *RecognizeRideHailingItineraryResponse {
	s.Headers = v
	return s
}

func (s *RecognizeRideHailingItineraryResponse) SetBody(v *RecognizeRideHailingItineraryResponseBody) *RecognizeRideHailingItineraryResponse {
	s.Body = v
	return s
}

type RecognizeRollTicketRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRollTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRollTicketRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRollTicketRequest) SetUrl(v string) *RecognizeRollTicketRequest {
	s.Url = &v
	return s
}

func (s *RecognizeRollTicketRequest) SetBody(v io.Reader) *RecognizeRollTicketRequest {
	s.Body = v
	return s
}

type RecognizeRollTicketResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeRollTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRollTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRollTicketResponseBody) SetCode(v string) *RecognizeRollTicketResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetData(v string) *RecognizeRollTicketResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetMessage(v string) *RecognizeRollTicketResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetRequestId(v string) *RecognizeRollTicketResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeRollTicketResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeRollTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeRollTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRollTicketResponse) GoString() string {
	return s.String()
}

func (s *RecognizeRollTicketResponse) SetHeaders(v map[string]*string) *RecognizeRollTicketResponse {
	s.Headers = v
	return s
}

func (s *RecognizeRollTicketResponse) SetBody(v *RecognizeRollTicketResponseBody) *RecognizeRollTicketResponse {
	s.Body = v
	return s
}

type RecognizeRussianRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRussianRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussianRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRussianRequest) SetNeedRotate(v bool) *RecognizeRussianRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeRussianRequest) SetOutputCharInfo(v bool) *RecognizeRussianRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeRussianRequest) SetOutputTable(v bool) *RecognizeRussianRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeRussianRequest) SetUrl(v string) *RecognizeRussianRequest {
	s.Url = &v
	return s
}

func (s *RecognizeRussianRequest) SetBody(v io.Reader) *RecognizeRussianRequest {
	s.Body = v
	return s
}

type RecognizeRussianResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeRussianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussianResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRussianResponseBody) SetCode(v string) *RecognizeRussianResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetData(v string) *RecognizeRussianResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetMessage(v string) *RecognizeRussianResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetRequestId(v string) *RecognizeRussianResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeRussianResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeRussianResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeRussianResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussianResponse) GoString() string {
	return s.String()
}

func (s *RecognizeRussianResponse) SetHeaders(v map[string]*string) *RecognizeRussianResponse {
	s.Headers = v
	return s
}

func (s *RecognizeRussianResponse) SetBody(v *RecognizeRussianResponseBody) *RecognizeRussianResponse {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeSocialSecurityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardRequest) SetUrl(v string) *RecognizeSocialSecurityCardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeSocialSecurityCardRequest) SetBody(v io.Reader) *RecognizeSocialSecurityCardRequest {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeSocialSecurityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardResponseBody) SetCode(v string) *RecognizeSocialSecurityCardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeSocialSecurityCardResponseBody) SetData(v string) *RecognizeSocialSecurityCardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeSocialSecurityCardResponseBody) SetMessage(v string) *RecognizeSocialSecurityCardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeSocialSecurityCardResponseBody) SetRequestId(v string) *RecognizeSocialSecurityCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeSocialSecurityCardResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeSocialSecurityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeSocialSecurityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardResponse) SetHeaders(v map[string]*string) *RecognizeSocialSecurityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeSocialSecurityCardResponse) SetBody(v *RecognizeSocialSecurityCardResponseBody) *RecognizeSocialSecurityCardResponse {
	s.Body = v
	return s
}

type RecognizeTableOcrRequest struct {
	// 是否无线条
	LineLess *bool `json:"LineLess,omitempty" xml:"LineLess,omitempty"`
	// 是否需要自动旋转功能，默认需要
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否跳过表格识别，如果没有检测到表格，可以设置"skip_detection":true
	SkipDetection *bool `json:"SkipDetection,omitempty" xml:"SkipDetection,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTableOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableOcrRequest) SetLineLess(v bool) *RecognizeTableOcrRequest {
	s.LineLess = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetNeedRotate(v bool) *RecognizeTableOcrRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetSkipDetection(v bool) *RecognizeTableOcrRequest {
	s.SkipDetection = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetUrl(v string) *RecognizeTableOcrRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetBody(v io.Reader) *RecognizeTableOcrRequest {
	s.Body = v
	return s
}

type RecognizeTableOcrResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTableOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTableOcrResponseBody) SetCode(v string) *RecognizeTableOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetData(v string) *RecognizeTableOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetMessage(v string) *RecognizeTableOcrResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetRequestId(v string) *RecognizeTableOcrResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTableOcrResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTableOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTableOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableOcrResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTableOcrResponse) SetHeaders(v map[string]*string) *RecognizeTableOcrResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTableOcrResponse) SetBody(v *RecognizeTableOcrResponseBody) *RecognizeTableOcrResponse {
	s.Body = v
	return s
}

type RecognizeTaxClearanceCertificateRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTaxClearanceCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxClearanceCertificateRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxClearanceCertificateRequest) SetUrl(v string) *RecognizeTaxClearanceCertificateRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateRequest) SetBody(v io.Reader) *RecognizeTaxClearanceCertificateRequest {
	s.Body = v
	return s
}

type RecognizeTaxClearanceCertificateResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTaxClearanceCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxClearanceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTaxClearanceCertificateResponseBody) SetCode(v string) *RecognizeTaxClearanceCertificateResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponseBody) SetData(v string) *RecognizeTaxClearanceCertificateResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponseBody) SetMessage(v string) *RecognizeTaxClearanceCertificateResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponseBody) SetRequestId(v string) *RecognizeTaxClearanceCertificateResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTaxClearanceCertificateResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTaxClearanceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTaxClearanceCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxClearanceCertificateResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTaxClearanceCertificateResponse) SetHeaders(v map[string]*string) *RecognizeTaxClearanceCertificateResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponse) SetBody(v *RecognizeTaxClearanceCertificateResponseBody) *RecognizeTaxClearanceCertificateResponse {
	s.Body = v
	return s
}

type RecognizeTaxiInvoiceRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTaxiInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceRequest) SetUrl(v string) *RecognizeTaxiInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTaxiInvoiceRequest) SetBody(v io.Reader) *RecognizeTaxiInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeTaxiInvoiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBody) SetCode(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetData(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetMessage(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetRequestId(v string) *RecognizeTaxiInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTaxiInvoiceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTaxiInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTaxiInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTaxiInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) SetBody(v *RecognizeTaxiInvoiceResponseBody) *RecognizeTaxiInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeThaiRequest struct {
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeThaiRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeThaiRequest) GoString() string {
	return s.String()
}

func (s *RecognizeThaiRequest) SetNeedRotate(v bool) *RecognizeThaiRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeThaiRequest) SetOutputCharInfo(v bool) *RecognizeThaiRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeThaiRequest) SetOutputTable(v bool) *RecognizeThaiRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeThaiRequest) SetUrl(v string) *RecognizeThaiRequest {
	s.Url = &v
	return s
}

func (s *RecognizeThaiRequest) SetBody(v io.Reader) *RecognizeThaiRequest {
	s.Body = v
	return s
}

type RecognizeThaiResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeThaiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeThaiResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeThaiResponseBody) SetCode(v string) *RecognizeThaiResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetData(v string) *RecognizeThaiResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetMessage(v string) *RecognizeThaiResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetRequestId(v string) *RecognizeThaiResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeThaiResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeThaiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeThaiResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeThaiResponse) GoString() string {
	return s.String()
}

func (s *RecognizeThaiResponse) SetHeaders(v map[string]*string) *RecognizeThaiResponse {
	s.Headers = v
	return s
}

func (s *RecognizeThaiResponse) SetBody(v *RecognizeThaiResponseBody) *RecognizeThaiResponse {
	s.Body = v
	return s
}

type RecognizeTollInvoiceRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTollInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTollInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTollInvoiceRequest) SetUrl(v string) *RecognizeTollInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTollInvoiceRequest) SetBody(v io.Reader) *RecognizeTollInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeTollInvoiceResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTollInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTollInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTollInvoiceResponseBody) SetCode(v string) *RecognizeTollInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTollInvoiceResponseBody) SetData(v string) *RecognizeTollInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTollInvoiceResponseBody) SetMessage(v string) *RecognizeTollInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTollInvoiceResponseBody) SetRequestId(v string) *RecognizeTollInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTollInvoiceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTollInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTollInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTollInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTollInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTollInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTollInvoiceResponse) SetBody(v *RecognizeTollInvoiceResponseBody) *RecognizeTollInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeTradeMarkCertificationRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTradeMarkCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTradeMarkCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTradeMarkCertificationRequest) SetUrl(v string) *RecognizeTradeMarkCertificationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTradeMarkCertificationRequest) SetBody(v io.Reader) *RecognizeTradeMarkCertificationRequest {
	s.Body = v
	return s
}

type RecognizeTradeMarkCertificationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTradeMarkCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTradeMarkCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetCode(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetData(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetMessage(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetRequestId(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTradeMarkCertificationResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTradeMarkCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTradeMarkCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTradeMarkCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTradeMarkCertificationResponse) SetHeaders(v map[string]*string) *RecognizeTradeMarkCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTradeMarkCertificationResponse) SetBody(v *RecognizeTradeMarkCertificationResponseBody) *RecognizeTradeMarkCertificationResponse {
	s.Body = v
	return s
}

type RecognizeTrainInvoiceRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTrainInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTrainInvoiceRequest) SetUrl(v string) *RecognizeTrainInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTrainInvoiceRequest) SetBody(v io.Reader) *RecognizeTrainInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeTrainInvoiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTrainInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTrainInvoiceResponseBody) SetCode(v string) *RecognizeTrainInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetData(v string) *RecognizeTrainInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetMessage(v string) *RecognizeTrainInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetRequestId(v string) *RecognizeTrainInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTrainInvoiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTrainInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTrainInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTrainInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTrainInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTrainInvoiceResponse) SetBody(v *RecognizeTrainInvoiceResponseBody) *RecognizeTrainInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeUsedCarInvoiceRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeUsedCarInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUsedCarInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeUsedCarInvoiceRequest) SetUrl(v string) *RecognizeUsedCarInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeUsedCarInvoiceRequest) SetBody(v io.Reader) *RecognizeUsedCarInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeUsedCarInvoiceResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeUsedCarInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUsedCarInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeUsedCarInvoiceResponseBody) SetCode(v string) *RecognizeUsedCarInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeUsedCarInvoiceResponseBody) SetData(v string) *RecognizeUsedCarInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeUsedCarInvoiceResponseBody) SetMessage(v string) *RecognizeUsedCarInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeUsedCarInvoiceResponseBody) SetRequestId(v string) *RecognizeUsedCarInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeUsedCarInvoiceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeUsedCarInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeUsedCarInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUsedCarInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeUsedCarInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeUsedCarInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeUsedCarInvoiceResponse) SetBody(v *RecognizeUsedCarInvoiceResponseBody) *RecognizeUsedCarInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeVehicleCertificationRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleCertificationRequest) SetUrl(v string) *RecognizeVehicleCertificationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeVehicleCertificationRequest) SetBody(v io.Reader) *RecognizeVehicleCertificationRequest {
	s.Body = v
	return s
}

type RecognizeVehicleCertificationResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleCertificationResponseBody) SetCode(v string) *RecognizeVehicleCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeVehicleCertificationResponseBody) SetData(v string) *RecognizeVehicleCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeVehicleCertificationResponseBody) SetMessage(v string) *RecognizeVehicleCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeVehicleCertificationResponseBody) SetRequestId(v string) *RecognizeVehicleCertificationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleCertificationResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVehicleCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVehicleCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleCertificationResponse) SetHeaders(v map[string]*string) *RecognizeVehicleCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleCertificationResponse) SetBody(v *RecognizeVehicleCertificationResponseBody) *RecognizeVehicleCertificationResponse {
	s.Body = v
	return s
}

type RecognizeVehicleLicenseRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleLicenseRequest) SetUrl(v string) *RecognizeVehicleLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeVehicleLicenseRequest) SetBody(v io.Reader) *RecognizeVehicleLicenseRequest {
	s.Body = v
	return s
}

type RecognizeVehicleLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleLicenseResponseBody) SetCode(v string) *RecognizeVehicleLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetData(v string) *RecognizeVehicleLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetMessage(v string) *RecognizeVehicleLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetRequestId(v string) *RecognizeVehicleLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleLicenseResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVehicleLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVehicleLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleLicenseResponse) SetHeaders(v map[string]*string) *RecognizeVehicleLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleLicenseResponse) SetBody(v *RecognizeVehicleLicenseResponseBody) *RecognizeVehicleLicenseResponse {
	s.Body = v
	return s
}

type RecognizeVehicleRegistrationRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleRegistrationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleRegistrationRequest) SetUrl(v string) *RecognizeVehicleRegistrationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeVehicleRegistrationRequest) SetBody(v io.Reader) *RecognizeVehicleRegistrationRequest {
	s.Body = v
	return s
}

type RecognizeVehicleRegistrationResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleRegistrationResponseBody) SetCode(v string) *RecognizeVehicleRegistrationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeVehicleRegistrationResponseBody) SetData(v string) *RecognizeVehicleRegistrationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeVehicleRegistrationResponseBody) SetMessage(v string) *RecognizeVehicleRegistrationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeVehicleRegistrationResponseBody) SetRequestId(v string) *RecognizeVehicleRegistrationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleRegistrationResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVehicleRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVehicleRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleRegistrationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleRegistrationResponse) SetHeaders(v map[string]*string) *RecognizeVehicleRegistrationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleRegistrationResponse) SetBody(v *RecognizeVehicleRegistrationResponseBody) *RecognizeVehicleRegistrationResponse {
	s.Body = v
	return s
}

type RecognizeWaybillRequest struct {
	// 图片链接（长度不超 2048，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片二进制字节流，最大10MB
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeWaybillRequest) GoString() string {
	return s.String()
}

func (s *RecognizeWaybillRequest) SetUrl(v string) *RecognizeWaybillRequest {
	s.Url = &v
	return s
}

func (s *RecognizeWaybillRequest) SetBody(v io.Reader) *RecognizeWaybillRequest {
	s.Body = v
	return s
}

type RecognizeWaybillResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeWaybillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeWaybillResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeWaybillResponseBody) SetCode(v string) *RecognizeWaybillResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetData(v string) *RecognizeWaybillResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetMessage(v string) *RecognizeWaybillResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetRequestId(v string) *RecognizeWaybillResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeWaybillResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeWaybillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeWaybillResponse) GoString() string {
	return s.String()
}

func (s *RecognizeWaybillResponse) SetHeaders(v map[string]*string) *RecognizeWaybillResponse {
	s.Headers = v
	return s
}

func (s *RecognizeWaybillResponse) SetBody(v *RecognizeWaybillResponseBody) *RecognizeWaybillResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ocr-api"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) RecognizeAdvancedWithOptions(request *RecognizeAdvancedRequest, runtime *util.RuntimeOptions) (_result *RecognizeAdvancedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSortPage)) {
		query["NeedSortPage"] = request.NeedSortPage
	}

	if !tea.BoolValue(util.IsUnset(request.NoStamp)) {
		query["NoStamp"] = request.NoStamp
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Paragraph)) {
		query["Paragraph"] = request.Paragraph
	}

	if !tea.BoolValue(util.IsUnset(request.Row)) {
		query["Row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeAdvanced"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeAdvancedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeAdvanced(request *RecognizeAdvancedRequest) (_result *RecognizeAdvancedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeAdvancedResponse{}
	_body, _err := client.RecognizeAdvancedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeAirItineraryWithOptions(request *RecognizeAirItineraryRequest, runtime *util.RuntimeOptions) (_result *RecognizeAirItineraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeAirItinerary"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeAirItineraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeAirItinerary(request *RecognizeAirItineraryRequest) (_result *RecognizeAirItineraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeAirItineraryResponse{}
	_body, _err := client.RecognizeAirItineraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBankAccountLicenseWithOptions(request *RecognizeBankAccountLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankAccountLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBankAccountLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBankAccountLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBankAccountLicense(request *RecognizeBankAccountLicenseRequest) (_result *RecognizeBankAccountLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBankAccountLicenseResponse{}
	_body, _err := client.RecognizeBankAccountLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBankCardWithOptions(request *RecognizeBankCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBankCard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBankCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBankCard(request *RecognizeBankCardRequest) (_result *RecognizeBankCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBankCardResponse{}
	_body, _err := client.RecognizeBankCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBasicWithOptions(request *RecognizeBasicRequest, runtime *util.RuntimeOptions) (_result *RecognizeBasicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBasic"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBasicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBasic(request *RecognizeBasicRequest) (_result *RecognizeBasicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBasicResponse{}
	_body, _err := client.RecognizeBasicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBatchRecognizeWithOptions(request *RecognizeBatchRecognizeRequest, runtime *util.RuntimeOptions) (_result *RecognizeBatchRecognizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageOp)) {
		query["ImageOp"] = request.ImageOp
	}

	if !tea.BoolValue(util.IsUnset(request.ImageOssKey)) {
		query["ImageOssKey"] = request.ImageOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSortPage)) {
		query["NeedSortPage"] = request.NeedSortPage
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBatchRecognize"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBatchRecognizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBatchRecognize(request *RecognizeBatchRecognizeRequest) (_result *RecognizeBatchRecognizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBatchRecognizeResponse{}
	_body, _err := client.RecognizeBatchRecognizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBirthCertificationWithOptions(request *RecognizeBirthCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeBirthCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBirthCertification"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBirthCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBirthCertification(request *RecognizeBirthCertificationRequest) (_result *RecognizeBirthCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBirthCertificationResponse{}
	_body, _err := client.RecognizeBirthCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBusShipTicketWithOptions(request *RecognizeBusShipTicketRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusShipTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBusShipTicket"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBusShipTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBusShipTicket(request *RecognizeBusShipTicketRequest) (_result *RecognizeBusShipTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBusShipTicketResponse{}
	_body, _err := client.RecognizeBusShipTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBusinessLicenseWithOptions(request *RecognizeBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBusinessLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBusinessLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBusinessLicense(request *RecognizeBusinessLicenseRequest) (_result *RecognizeBusinessLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBusinessLicenseResponse{}
	_body, _err := client.RecognizeBusinessLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeCarInvoiceWithOptions(request *RecognizeCarInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCarInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCarInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeCarInvoice(request *RecognizeCarInvoiceRequest) (_result *RecognizeCarInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCarInvoiceResponse{}
	_body, _err := client.RecognizeCarInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeCarNumberWithOptions(request *RecognizeCarNumberRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCarNumber"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCarNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeCarNumber(request *RecognizeCarNumberRequest) (_result *RecognizeCarNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCarNumberResponse{}
	_body, _err := client.RecognizeCarNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeCarVinCodeWithOptions(request *RecognizeCarVinCodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarVinCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCarVinCode"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCarVinCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeCarVinCode(request *RecognizeCarVinCodeRequest) (_result *RecognizeCarVinCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCarVinCodeResponse{}
	_body, _err := client.RecognizeCarVinCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeChinesePassportWithOptions(request *RecognizeChinesePassportRequest, runtime *util.RuntimeOptions) (_result *RecognizeChinesePassportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeChinesePassport"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeChinesePassportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeChinesePassport(request *RecognizeChinesePassportRequest) (_result *RecognizeChinesePassportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeChinesePassportResponse{}
	_body, _err := client.RecognizeChinesePassportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeCommonPrintedInvoiceWithOptions(request *RecognizeCommonPrintedInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeCommonPrintedInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCommonPrintedInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCommonPrintedInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeCommonPrintedInvoice(request *RecognizeCommonPrintedInvoiceRequest) (_result *RecognizeCommonPrintedInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCommonPrintedInvoiceResponse{}
	_body, _err := client.RecognizeCommonPrintedInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeCtwoMedicalDeviceManageLicenseWithOptions(request *RecognizeCtwoMedicalDeviceManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeCtwoMedicalDeviceManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCtwoMedicalDeviceManageLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCtwoMedicalDeviceManageLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeCtwoMedicalDeviceManageLicense(request *RecognizeCtwoMedicalDeviceManageLicenseRequest) (_result *RecognizeCtwoMedicalDeviceManageLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCtwoMedicalDeviceManageLicenseResponse{}
	_body, _err := client.RecognizeCtwoMedicalDeviceManageLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeDeleteExcelRecordWithOptions(request *RecognizeDeleteExcelRecordRequest, runtime *util.RuntimeOptions) (_result *RecognizeDeleteExcelRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeDeleteExcelRecord"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeDeleteExcelRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeDeleteExcelRecord(request *RecognizeDeleteExcelRecordRequest) (_result *RecognizeDeleteExcelRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeDeleteExcelRecordResponse{}
	_body, _err := client.RecognizeDeleteExcelRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeDrivingLicenseWithOptions(request *RecognizeDrivingLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeDrivingLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeDrivingLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeDrivingLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeDrivingLicense(request *RecognizeDrivingLicenseRequest) (_result *RecognizeDrivingLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeDrivingLicenseResponse{}
	_body, _err := client.RecognizeDrivingLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeEduFormulaWithOptions(request *RecognizeEduFormulaRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduFormulaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduFormula"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduFormulaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeEduFormula(request *RecognizeEduFormulaRequest) (_result *RecognizeEduFormulaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduFormulaResponse{}
	_body, _err := client.RecognizeEduFormulaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeEduOralCalculationWithOptions(request *RecognizeEduOralCalculationRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduOralCalculationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduOralCalculation"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduOralCalculationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeEduOralCalculation(request *RecognizeEduOralCalculationRequest) (_result *RecognizeEduOralCalculationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduOralCalculationResponse{}
	_body, _err := client.RecognizeEduOralCalculationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeEduPaperCutWithOptions(request *RecognizeEduPaperCutRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperCutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CutType)) {
		query["CutType"] = request.CutType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduPaperCut"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduPaperCutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeEduPaperCut(request *RecognizeEduPaperCutRequest) (_result *RecognizeEduPaperCutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduPaperCutResponse{}
	_body, _err := client.RecognizeEduPaperCutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeEduPaperOcrWithOptions(request *RecognizeEduPaperOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.OutputOricoord)) {
		query["OutputOricoord"] = request.OutputOricoord
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduPaperOcr"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduPaperOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeEduPaperOcr(request *RecognizeEduPaperOcrRequest) (_result *RecognizeEduPaperOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduPaperOcrResponse{}
	_body, _err := client.RecognizeEduPaperOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeEduPaperStructedWithOptions(request *RecognizeEduPaperStructedRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperStructedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduPaperStructed"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduPaperStructedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeEduPaperStructed(request *RecognizeEduPaperStructedRequest) (_result *RecognizeEduPaperStructedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduPaperStructedResponse{}
	_body, _err := client.RecognizeEduPaperStructedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeEduQuestionOcrWithOptions(request *RecognizeEduQuestionOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduQuestionOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduQuestionOcr"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduQuestionOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeEduQuestionOcr(request *RecognizeEduQuestionOcrRequest) (_result *RecognizeEduQuestionOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduQuestionOcrResponse{}
	_body, _err := client.RecognizeEduQuestionOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeEnglishWithOptions(request *RecognizeEnglishRequest, runtime *util.RuntimeOptions) (_result *RecognizeEnglishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEnglish"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEnglishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeEnglish(request *RecognizeEnglishRequest) (_result *RecognizeEnglishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEnglishResponse{}
	_body, _err := client.RecognizeEnglishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeEstateCertificationWithOptions(request *RecognizeEstateCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeEstateCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEstateCertification"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEstateCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeEstateCertification(request *RecognizeEstateCertificationRequest) (_result *RecognizeEstateCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEstateCertificationResponse{}
	_body, _err := client.RecognizeEstateCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeExcelExportWithOptions(request *RecognizeExcelExportRequest, runtime *util.RuntimeOptions) (_result *RecognizeExcelExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageOp)) {
		query["ImageOp"] = request.ImageOp
	}

	if !tea.BoolValue(util.IsUnset(request.OcrImageCount)) {
		query["OcrImageCount"] = request.OcrImageCount
	}

	if !tea.BoolValue(util.IsUnset(request.OcrResult)) {
		query["OcrResult"] = request.OcrResult
	}

	if !tea.BoolValue(util.IsUnset(request.OcrType)) {
		query["OcrType"] = request.OcrType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeExcelExport"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeExcelExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeExcelExport(request *RecognizeExcelExportRequest) (_result *RecognizeExcelExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeExcelExportResponse{}
	_body, _err := client.RecognizeExcelExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeExcelRecordWithOptions(request *RecognizeExcelRecordRequest, runtime *util.RuntimeOptions) (_result *RecognizeExcelRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrPage)) {
		query["CurrPage"] = request.CurrPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeExcelRecord"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeExcelRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeExcelRecord(request *RecognizeExcelRecordRequest) (_result *RecognizeExcelRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeExcelRecordResponse{}
	_body, _err := client.RecognizeExcelRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeExitEntryPermitToHKWithOptions(request *RecognizeExitEntryPermitToHKRequest, runtime *util.RuntimeOptions) (_result *RecognizeExitEntryPermitToHKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeExitEntryPermitToHK"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeExitEntryPermitToHKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeExitEntryPermitToHK(request *RecognizeExitEntryPermitToHKRequest) (_result *RecognizeExitEntryPermitToHKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeExitEntryPermitToHKResponse{}
	_body, _err := client.RecognizeExitEntryPermitToHKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeExitEntryPermitToMainlandWithOptions(request *RecognizeExitEntryPermitToMainlandRequest, runtime *util.RuntimeOptions) (_result *RecognizeExitEntryPermitToMainlandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeExitEntryPermitToMainland"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeExitEntryPermitToMainlandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeExitEntryPermitToMainland(request *RecognizeExitEntryPermitToMainlandRequest) (_result *RecognizeExitEntryPermitToMainlandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeExitEntryPermitToMainlandResponse{}
	_body, _err := client.RecognizeExitEntryPermitToMainlandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFoodManageLicenseWithOptions(request *RecognizeFoodManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeFoodManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeFoodManageLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeFoodManageLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFoodManageLicense(request *RecognizeFoodManageLicenseRequest) (_result *RecognizeFoodManageLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFoodManageLicenseResponse{}
	_body, _err := client.RecognizeFoodManageLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFoodProduceLicenseWithOptions(request *RecognizeFoodProduceLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeFoodProduceLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeFoodProduceLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeFoodProduceLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFoodProduceLicense(request *RecognizeFoodProduceLicenseRequest) (_result *RecognizeFoodProduceLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFoodProduceLicenseResponse{}
	_body, _err := client.RecognizeFoodProduceLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeGeneralWithOptions(request *RecognizeGeneralRequest, runtime *util.RuntimeOptions) (_result *RecognizeGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeGeneral"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeGeneral(request *RecognizeGeneralRequest) (_result *RecognizeGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeGeneralResponse{}
	_body, _err := client.RecognizeGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeHandwritingWithOptions(request *RecognizeHandwritingRequest, runtime *util.RuntimeOptions) (_result *RecognizeHandwritingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSortPage)) {
		query["NeedSortPage"] = request.NeedSortPage
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeHandwriting"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHandwritingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeHandwriting(request *RecognizeHandwritingRequest) (_result *RecognizeHandwritingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHandwritingResponse{}
	_body, _err := client.RecognizeHandwritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeHouseholdWithOptions(request *RecognizeHouseholdRequest, runtime *util.RuntimeOptions) (_result *RecognizeHouseholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeHousehold"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHouseholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeHousehold(request *RecognizeHouseholdRequest) (_result *RecognizeHouseholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHouseholdResponse{}
	_body, _err := client.RecognizeHouseholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeIdcardWithOptions(request *RecognizeIdcardRequest, runtime *util.RuntimeOptions) (_result *RecognizeIdcardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeIdcard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeIdcardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeIdcard(request *RecognizeIdcardRequest) (_result *RecognizeIdcardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeIdcardResponse{}
	_body, _err := client.RecognizeIdcardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeInvoiceWithOptions(request *RecognizeInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeInvoice(request *RecognizeInvoiceRequest) (_result *RecognizeInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeInvoiceResponse{}
	_body, _err := client.RecognizeInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeJanpaneseWithOptions(request *RecognizeJanpaneseRequest, runtime *util.RuntimeOptions) (_result *RecognizeJanpaneseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeJanpanese"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeJanpaneseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeJanpanese(request *RecognizeJanpaneseRequest) (_result *RecognizeJanpaneseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeJanpaneseResponse{}
	_body, _err := client.RecognizeJanpaneseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeKoreanWithOptions(request *RecognizeKoreanRequest, runtime *util.RuntimeOptions) (_result *RecognizeKoreanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeKorean"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeKoreanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeKorean(request *RecognizeKoreanRequest) (_result *RecognizeKoreanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeKoreanResponse{}
	_body, _err := client.RecognizeKoreanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeLatinWithOptions(request *RecognizeLatinRequest, runtime *util.RuntimeOptions) (_result *RecognizeLatinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeLatin"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeLatinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeLatin(request *RecognizeLatinRequest) (_result *RecognizeLatinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeLatinResponse{}
	_body, _err := client.RecognizeLatinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeMedicalDeviceManageLicenseWithOptions(request *RecognizeMedicalDeviceManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeMedicalDeviceManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMedicalDeviceManageLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMedicalDeviceManageLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeMedicalDeviceManageLicense(request *RecognizeMedicalDeviceManageLicenseRequest) (_result *RecognizeMedicalDeviceManageLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMedicalDeviceManageLicenseResponse{}
	_body, _err := client.RecognizeMedicalDeviceManageLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeMedicalDeviceProduceLicenseWithOptions(request *RecognizeMedicalDeviceProduceLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeMedicalDeviceProduceLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMedicalDeviceProduceLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMedicalDeviceProduceLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeMedicalDeviceProduceLicense(request *RecognizeMedicalDeviceProduceLicenseRequest) (_result *RecognizeMedicalDeviceProduceLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMedicalDeviceProduceLicenseResponse{}
	_body, _err := client.RecognizeMedicalDeviceProduceLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeMixedInvoicesWithOptions(request *RecognizeMixedInvoicesRequest, runtime *util.RuntimeOptions) (_result *RecognizeMixedInvoicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMixedInvoices"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMixedInvoicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeMixedInvoices(request *RecognizeMixedInvoicesRequest) (_result *RecognizeMixedInvoicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMixedInvoicesResponse{}
	_body, _err := client.RecognizeMixedInvoicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeMultiLanguageWithOptions(tmpReq *RecognizeMultiLanguageRequest, runtime *util.RuntimeOptions) (_result *RecognizeMultiLanguageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RecognizeMultiLanguageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Languages)) {
		request.LanguagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Languages, tea.String("Languages"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LanguagesShrink)) {
		query["Languages"] = request.LanguagesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSortPage)) {
		query["NeedSortPage"] = request.NeedSortPage
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMultiLanguage"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMultiLanguageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeMultiLanguage(request *RecognizeMultiLanguageRequest) (_result *RecognizeMultiLanguageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMultiLanguageResponse{}
	_body, _err := client.RecognizeMultiLanguageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizePassportWithOptions(request *RecognizePassportRequest, runtime *util.RuntimeOptions) (_result *RecognizePassportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizePassport"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePassportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizePassport(request *RecognizePassportRequest) (_result *RecognizePassportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePassportResponse{}
	_body, _err := client.RecognizePassportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeQuotaInvoiceWithOptions(request *RecognizeQuotaInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeQuotaInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeQuotaInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeQuotaInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeQuotaInvoice(request *RecognizeQuotaInvoiceRequest) (_result *RecognizeQuotaInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeQuotaInvoiceResponse{}
	_body, _err := client.RecognizeQuotaInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeRideHailingItineraryWithOptions(request *RecognizeRideHailingItineraryRequest, runtime *util.RuntimeOptions) (_result *RecognizeRideHailingItineraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeRideHailingItinerary"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeRideHailingItineraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeRideHailingItinerary(request *RecognizeRideHailingItineraryRequest) (_result *RecognizeRideHailingItineraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeRideHailingItineraryResponse{}
	_body, _err := client.RecognizeRideHailingItineraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeRollTicketWithOptions(request *RecognizeRollTicketRequest, runtime *util.RuntimeOptions) (_result *RecognizeRollTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeRollTicket"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeRollTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeRollTicket(request *RecognizeRollTicketRequest) (_result *RecognizeRollTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeRollTicketResponse{}
	_body, _err := client.RecognizeRollTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeRussianWithOptions(request *RecognizeRussianRequest, runtime *util.RuntimeOptions) (_result *RecognizeRussianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeRussian"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeRussianResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeRussian(request *RecognizeRussianRequest) (_result *RecognizeRussianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeRussianResponse{}
	_body, _err := client.RecognizeRussianWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeSocialSecurityCardWithOptions(request *RecognizeSocialSecurityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeSocialSecurityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeSocialSecurityCard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeSocialSecurityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeSocialSecurityCard(request *RecognizeSocialSecurityCardRequest) (_result *RecognizeSocialSecurityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeSocialSecurityCardResponse{}
	_body, _err := client.RecognizeSocialSecurityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTableOcrWithOptions(request *RecognizeTableOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeTableOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LineLess)) {
		query["LineLess"] = request.LineLess
	}

	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDetection)) {
		query["SkipDetection"] = request.SkipDetection
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTableOcr"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTableOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTableOcr(request *RecognizeTableOcrRequest) (_result *RecognizeTableOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTableOcrResponse{}
	_body, _err := client.RecognizeTableOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTaxClearanceCertificateWithOptions(request *RecognizeTaxClearanceCertificateRequest, runtime *util.RuntimeOptions) (_result *RecognizeTaxClearanceCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTaxClearanceCertificate"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTaxClearanceCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTaxClearanceCertificate(request *RecognizeTaxClearanceCertificateRequest) (_result *RecognizeTaxClearanceCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTaxClearanceCertificateResponse{}
	_body, _err := client.RecognizeTaxClearanceCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTaxiInvoiceWithOptions(request *RecognizeTaxiInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTaxiInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTaxiInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTaxiInvoice(request *RecognizeTaxiInvoiceRequest) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTaxiInvoiceResponse{}
	_body, _err := client.RecognizeTaxiInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeThaiWithOptions(request *RecognizeThaiRequest, runtime *util.RuntimeOptions) (_result *RecognizeThaiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeThai"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeThaiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeThai(request *RecognizeThaiRequest) (_result *RecognizeThaiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeThaiResponse{}
	_body, _err := client.RecognizeThaiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTollInvoiceWithOptions(request *RecognizeTollInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTollInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTollInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTollInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTollInvoice(request *RecognizeTollInvoiceRequest) (_result *RecognizeTollInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTollInvoiceResponse{}
	_body, _err := client.RecognizeTollInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTradeMarkCertificationWithOptions(request *RecognizeTradeMarkCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeTradeMarkCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTradeMarkCertification"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTradeMarkCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTradeMarkCertification(request *RecognizeTradeMarkCertificationRequest) (_result *RecognizeTradeMarkCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTradeMarkCertificationResponse{}
	_body, _err := client.RecognizeTradeMarkCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTrainInvoiceWithOptions(request *RecognizeTrainInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTrainInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTrainInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTrainInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTrainInvoice(request *RecognizeTrainInvoiceRequest) (_result *RecognizeTrainInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTrainInvoiceResponse{}
	_body, _err := client.RecognizeTrainInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeUsedCarInvoiceWithOptions(request *RecognizeUsedCarInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeUsedCarInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeUsedCarInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeUsedCarInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeUsedCarInvoice(request *RecognizeUsedCarInvoiceRequest) (_result *RecognizeUsedCarInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeUsedCarInvoiceResponse{}
	_body, _err := client.RecognizeUsedCarInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVehicleCertificationWithOptions(request *RecognizeVehicleCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleCertification"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVehicleCertification(request *RecognizeVehicleCertificationRequest) (_result *RecognizeVehicleCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleCertificationResponse{}
	_body, _err := client.RecognizeVehicleCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVehicleLicenseWithOptions(request *RecognizeVehicleLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVehicleLicense(request *RecognizeVehicleLicenseRequest) (_result *RecognizeVehicleLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleLicenseResponse{}
	_body, _err := client.RecognizeVehicleLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVehicleRegistrationWithOptions(request *RecognizeVehicleRegistrationRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleRegistration"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVehicleRegistration(request *RecognizeVehicleRegistrationRequest) (_result *RecognizeVehicleRegistrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleRegistrationResponse{}
	_body, _err := client.RecognizeVehicleRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeWaybillWithOptions(request *RecognizeWaybillRequest, runtime *util.RuntimeOptions) (_result *RecognizeWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeWaybill"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeWaybillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeWaybill(request *RecognizeWaybillRequest) (_result *RecognizeWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeWaybillResponse{}
	_body, _err := client.RecognizeWaybillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
