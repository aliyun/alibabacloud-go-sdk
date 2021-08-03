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

type RecognizeDrivingLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeDrivingLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBody) SetRequestId(v string) *RecognizeDrivingLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetData(v string) *RecognizeDrivingLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetCode(v string) *RecognizeDrivingLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetMessage(v string) *RecognizeDrivingLicenseResponseBody {
	s.Message = &v
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

type RecognizeKoreanRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
}

func (s RecognizeKoreanRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeKoreanRequest) GoString() string {
	return s.String()
}

func (s *RecognizeKoreanRequest) SetUrl(v string) *RecognizeKoreanRequest {
	s.Url = &v
	return s
}

func (s *RecognizeKoreanRequest) SetOutputCharInfo(v bool) *RecognizeKoreanRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeKoreanRequest) SetNeedRotate(v bool) *RecognizeKoreanRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeKoreanRequest) SetOutputTable(v bool) *RecognizeKoreanRequest {
	s.OutputTable = &v
	return s
}

type RecognizeKoreanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeKoreanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeKoreanResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeKoreanResponseBody) SetRequestId(v string) *RecognizeKoreanResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetData(v string) *RecognizeKoreanResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetCode(v string) *RecognizeKoreanResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetMessage(v string) *RecognizeKoreanResponseBody {
	s.Message = &v
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

type RecognizeCarInvoiceRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeCarInvoiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeCarInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarInvoiceResponseBody) SetRequestId(v string) *RecognizeCarInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetData(v string) *RecognizeCarInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetCode(v string) *RecognizeCarInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetMessage(v string) *RecognizeCarInvoiceResponseBody {
	s.Message = &v
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

type RecognizeMixedInvoicesRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeMixedInvoicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeMixedInvoicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedInvoicesResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMixedInvoicesResponseBody) SetRequestId(v string) *RecognizeMixedInvoicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetData(v string) *RecognizeMixedInvoicesResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetCode(v string) *RecognizeMixedInvoicesResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetMessage(v string) *RecognizeMixedInvoicesResponseBody {
	s.Message = &v
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

type RecognizeEstateCertificationRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeEstateCertificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeEstateCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEstateCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEstateCertificationResponseBody) SetRequestId(v string) *RecognizeEstateCertificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetData(v string) *RecognizeEstateCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetCode(v string) *RecognizeEstateCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetMessage(v string) *RecognizeEstateCertificationResponseBody {
	s.Message = &v
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

type RecognizeCarNumberRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeCarNumberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeCarNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarNumberResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarNumberResponseBody) SetRequestId(v string) *RecognizeCarNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetData(v string) *RecognizeCarNumberResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetCode(v string) *RecognizeCarNumberResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetMessage(v string) *RecognizeCarNumberResponseBody {
	s.Message = &v
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

type RecognizeEduPaperOcrRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片类型
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// 年级学科
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// 是否输出原图坐标信息(如果图片被做过旋转，图片校正等处理)
	OutputOricoord *bool `json:"OutputOricoord,omitempty" xml:"OutputOricoord,omitempty"`
}

func (s RecognizeEduPaperOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperOcrRequest) SetUrl(v string) *RecognizeEduPaperOcrRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetImageType(v string) *RecognizeEduPaperOcrRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetSubject(v string) *RecognizeEduPaperOcrRequest {
	s.Subject = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetOutputOricoord(v bool) *RecognizeEduPaperOcrRequest {
	s.OutputOricoord = &v
	return s
}

type RecognizeEduPaperOcrResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeEduPaperOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperOcrResponseBody) SetRequestId(v string) *RecognizeEduPaperOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetData(v string) *RecognizeEduPaperOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetCode(v string) *RecognizeEduPaperOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetMessage(v string) *RecognizeEduPaperOcrResponseBody {
	s.Message = &v
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

type RecognizeTradeMarkCertificationRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeTradeMarkCertificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeTradeMarkCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTradeMarkCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetRequestId(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetData(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetCode(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetMessage(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Message = &v
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

type RecognizeBirthCertificationRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeBirthCertificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeBirthCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBirthCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBirthCertificationResponseBody) SetRequestId(v string) *RecognizeBirthCertificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetData(v string) *RecognizeBirthCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetCode(v string) *RecognizeBirthCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetMessage(v string) *RecognizeBirthCertificationResponseBody {
	s.Message = &v
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

type RecognizeDoctypeRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeDoctypeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDoctypeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDoctypeRequest) SetUrl(v string) *RecognizeDoctypeRequest {
	s.Url = &v
	return s
}

type RecognizeDoctypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeDoctypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDoctypeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDoctypeResponseBody) SetRequestId(v string) *RecognizeDoctypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeDoctypeResponseBody) SetData(v string) *RecognizeDoctypeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeDoctypeResponseBody) SetCode(v string) *RecognizeDoctypeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeDoctypeResponseBody) SetMessage(v string) *RecognizeDoctypeResponseBody {
	s.Message = &v
	return s
}

type RecognizeDoctypeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeDoctypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeDoctypeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDoctypeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDoctypeResponse) SetHeaders(v map[string]*string) *RecognizeDoctypeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDoctypeResponse) SetBody(v *RecognizeDoctypeResponseBody) *RecognizeDoctypeResponse {
	s.Body = v
	return s
}

type RecognizeBankAccountLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeBankAccountLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeBankAccountLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAccountLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankAccountLicenseResponseBody) SetRequestId(v string) *RecognizeBankAccountLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetData(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetCode(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetMessage(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Message = &v
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

type RecognizeFoodManageLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeFoodManageLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeFoodManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFoodManageLicenseResponseBody) SetRequestId(v string) *RecognizeFoodManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetData(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetCode(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetMessage(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Message = &v
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

type RecognizeRollTicketRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeRollTicketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeRollTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRollTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRollTicketResponseBody) SetRequestId(v string) *RecognizeRollTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetData(v string) *RecognizeRollTicketResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetCode(v string) *RecognizeRollTicketResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetMessage(v string) *RecognizeRollTicketResponseBody {
	s.Message = &v
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

type RecognizeEduFormulaRequest struct {
	// 图片链接（不支持 base64）,图片最大尺寸 1024*1024
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeEduFormulaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeEduFormulaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduFormulaResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduFormulaResponseBody) SetRequestId(v string) *RecognizeEduFormulaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetData(v string) *RecognizeEduFormulaResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetCode(v string) *RecognizeEduFormulaResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetMessage(v string) *RecognizeEduFormulaResponseBody {
	s.Message = &v
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

type RecognizePassportRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizePassportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizePassportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePassportResponseBody) SetRequestId(v string) *RecognizePassportResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizePassportResponseBody) SetData(v string) *RecognizePassportResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizePassportResponseBody) SetCode(v string) *RecognizePassportResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizePassportResponseBody) SetMessage(v string) *RecognizePassportResponseBody {
	s.Message = &v
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

type RecognizeTaxiInvoiceRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeTaxiInvoiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBody) SetRequestId(v string) *RecognizeTaxiInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetData(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetCode(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetMessage(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Message = &v
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

type RecognizeFoodProduceLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeFoodProduceLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeFoodProduceLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodProduceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetRequestId(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetData(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetCode(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetMessage(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Message = &v
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

type RecognizeMedicalDeviceProduceLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeMedicalDeviceProduceLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeMedicalDeviceProduceLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceProduceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetRequestId(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetData(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetCode(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetMessage(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Message = &v
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

type RecognizeHandwritingRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
}

func (s RecognizeHandwritingRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandwritingRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHandwritingRequest) SetUrl(v string) *RecognizeHandwritingRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetOutputCharInfo(v bool) *RecognizeHandwritingRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetNeedRotate(v bool) *RecognizeHandwritingRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetOutputTable(v bool) *RecognizeHandwritingRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetNeedSortPage(v bool) *RecognizeHandwritingRequest {
	s.NeedSortPage = &v
	return s
}

type RecognizeHandwritingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeHandwritingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandwritingResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHandwritingResponseBody) SetRequestId(v string) *RecognizeHandwritingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetData(v string) *RecognizeHandwritingResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetCode(v string) *RecognizeHandwritingResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetMessage(v string) *RecognizeHandwritingResponseBody {
	s.Message = &v
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

type RecognizeAirItineraryRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeAirItineraryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeAirItineraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAirItineraryResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAirItineraryResponseBody) SetRequestId(v string) *RecognizeAirItineraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetData(v string) *RecognizeAirItineraryResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetCode(v string) *RecognizeAirItineraryResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetMessage(v string) *RecognizeAirItineraryResponseBody {
	s.Message = &v
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

type RecognizeJanpaneseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
}

func (s RecognizeJanpaneseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeJanpaneseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeJanpaneseRequest) SetUrl(v string) *RecognizeJanpaneseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetOutputCharInfo(v bool) *RecognizeJanpaneseRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetNeedRotate(v bool) *RecognizeJanpaneseRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetOutputTable(v bool) *RecognizeJanpaneseRequest {
	s.OutputTable = &v
	return s
}

type RecognizeJanpaneseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeJanpaneseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeJanpaneseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeJanpaneseResponseBody) SetRequestId(v string) *RecognizeJanpaneseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetData(v string) *RecognizeJanpaneseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetCode(v string) *RecognizeJanpaneseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetMessage(v string) *RecognizeJanpaneseResponseBody {
	s.Message = &v
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

type RecognizeCtwoMedicalDeviceManageLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeCtwoMedicalDeviceManageLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetRequestId(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetData(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetCode(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetMessage(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Message = &v
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

type RecognizeThaiRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
}

func (s RecognizeThaiRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeThaiRequest) GoString() string {
	return s.String()
}

func (s *RecognizeThaiRequest) SetUrl(v string) *RecognizeThaiRequest {
	s.Url = &v
	return s
}

func (s *RecognizeThaiRequest) SetOutputCharInfo(v bool) *RecognizeThaiRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeThaiRequest) SetNeedRotate(v bool) *RecognizeThaiRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeThaiRequest) SetOutputTable(v bool) *RecognizeThaiRequest {
	s.OutputTable = &v
	return s
}

type RecognizeThaiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeThaiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeThaiResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeThaiResponseBody) SetRequestId(v string) *RecognizeThaiResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetData(v string) *RecognizeThaiResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetCode(v string) *RecognizeThaiResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetMessage(v string) *RecognizeThaiResponseBody {
	s.Message = &v
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

type RecognizeMedicalDeviceManageLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeMedicalDeviceManageLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeMedicalDeviceManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetRequestId(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetData(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetCode(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetMessage(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Message = &v
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

type RecognizeLatinRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
}

func (s RecognizeLatinRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLatinRequest) GoString() string {
	return s.String()
}

func (s *RecognizeLatinRequest) SetUrl(v string) *RecognizeLatinRequest {
	s.Url = &v
	return s
}

func (s *RecognizeLatinRequest) SetOutputCharInfo(v bool) *RecognizeLatinRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeLatinRequest) SetNeedRotate(v bool) *RecognizeLatinRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeLatinRequest) SetOutputTable(v bool) *RecognizeLatinRequest {
	s.OutputTable = &v
	return s
}

type RecognizeLatinResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeLatinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLatinResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeLatinResponseBody) SetRequestId(v string) *RecognizeLatinResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetData(v string) *RecognizeLatinResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetCode(v string) *RecognizeLatinResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetMessage(v string) *RecognizeLatinResponseBody {
	s.Message = &v
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

type RecognizeInvoiceRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeInvoiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeInvoiceResponseBody) SetRequestId(v string) *RecognizeInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetData(v string) *RecognizeInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetCode(v string) *RecognizeInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetMessage(v string) *RecognizeInvoiceResponseBody {
	s.Message = &v
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

type RecognizeMixedCardsRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeMixedCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedCardsRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMixedCardsRequest) SetUrl(v string) *RecognizeMixedCardsRequest {
	s.Url = &v
	return s
}

type RecognizeMixedCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeMixedCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedCardsResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMixedCardsResponseBody) SetRequestId(v string) *RecognizeMixedCardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeMixedCardsResponseBody) SetData(v string) *RecognizeMixedCardsResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMixedCardsResponseBody) SetCode(v string) *RecognizeMixedCardsResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMixedCardsResponseBody) SetMessage(v string) *RecognizeMixedCardsResponseBody {
	s.Message = &v
	return s
}

type RecognizeMixedCardsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeMixedCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeMixedCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedCardsResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMixedCardsResponse) SetHeaders(v map[string]*string) *RecognizeMixedCardsResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMixedCardsResponse) SetBody(v *RecognizeMixedCardsResponseBody) *RecognizeMixedCardsResponse {
	s.Body = v
	return s
}

type RecognizeWaybillRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeWaybillResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeWaybillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeWaybillResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeWaybillResponseBody) SetRequestId(v string) *RecognizeWaybillResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetData(v string) *RecognizeWaybillResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetCode(v string) *RecognizeWaybillResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetMessage(v string) *RecognizeWaybillResponseBody {
	s.Message = &v
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

type RecognizeCarVinCodeRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeCarVinCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeCarVinCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarVinCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarVinCodeResponseBody) SetRequestId(v string) *RecognizeCarVinCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetData(v string) *RecognizeCarVinCodeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetCode(v string) *RecognizeCarVinCodeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetMessage(v string) *RecognizeCarVinCodeResponseBody {
	s.Message = &v
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

type RecognizeAdvancedRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
}

func (s RecognizeAdvancedRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAdvancedRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAdvancedRequest) SetUrl(v string) *RecognizeAdvancedRequest {
	s.Url = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetOutputCharInfo(v bool) *RecognizeAdvancedRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetNeedRotate(v bool) *RecognizeAdvancedRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetOutputTable(v bool) *RecognizeAdvancedRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetNeedSortPage(v bool) *RecognizeAdvancedRequest {
	s.NeedSortPage = &v
	return s
}

type RecognizeAdvancedResponseBody struct {
	// 请求唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误提示
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeAdvancedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAdvancedResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAdvancedResponseBody) SetRequestId(v string) *RecognizeAdvancedResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetData(v string) *RecognizeAdvancedResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetCode(v string) *RecognizeAdvancedResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetMessage(v string) *RecognizeAdvancedResponseBody {
	s.Message = &v
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

type RecognizeVehicleLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeVehicleLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeVehicleLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleLicenseResponseBody) SetRequestId(v string) *RecognizeVehicleLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetData(v string) *RecognizeVehicleLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetCode(v string) *RecognizeVehicleLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetMessage(v string) *RecognizeVehicleLicenseResponseBody {
	s.Message = &v
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

type RecognizeRussianRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
}

func (s RecognizeRussianRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussianRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRussianRequest) SetUrl(v string) *RecognizeRussianRequest {
	s.Url = &v
	return s
}

func (s *RecognizeRussianRequest) SetOutputCharInfo(v bool) *RecognizeRussianRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeRussianRequest) SetNeedRotate(v bool) *RecognizeRussianRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeRussianRequest) SetOutputTable(v bool) *RecognizeRussianRequest {
	s.OutputTable = &v
	return s
}

type RecognizeRussianResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeRussianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussianResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRussianResponseBody) SetRequestId(v string) *RecognizeRussianResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetData(v string) *RecognizeRussianResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetCode(v string) *RecognizeRussianResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetMessage(v string) *RecognizeRussianResponseBody {
	s.Message = &v
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

type RecognizeHouseCertificationRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeHouseCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHouseCertificationRequest) SetUrl(v string) *RecognizeHouseCertificationRequest {
	s.Url = &v
	return s
}

type RecognizeHouseCertificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeHouseCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHouseCertificationResponseBody) SetRequestId(v string) *RecognizeHouseCertificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeHouseCertificationResponseBody) SetData(v string) *RecognizeHouseCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHouseCertificationResponseBody) SetCode(v string) *RecognizeHouseCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHouseCertificationResponseBody) SetMessage(v string) *RecognizeHouseCertificationResponseBody {
	s.Message = &v
	return s
}

type RecognizeHouseCertificationResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeHouseCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeHouseCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHouseCertificationResponse) SetHeaders(v map[string]*string) *RecognizeHouseCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHouseCertificationResponse) SetBody(v *RecognizeHouseCertificationResponseBody) *RecognizeHouseCertificationResponse {
	s.Body = v
	return s
}

type RecognizeBasicRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeBasicResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeBasicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBasicResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBasicResponseBody) SetRequestId(v string) *RecognizeBasicResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetData(v string) *RecognizeBasicResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetCode(v string) *RecognizeBasicResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetMessage(v string) *RecognizeBasicResponseBody {
	s.Message = &v
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

type RecognizeBusinessLicenseRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeBusinessLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBody) SetRequestId(v string) *RecognizeBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetData(v string) *RecognizeBusinessLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetCode(v string) *RecognizeBusinessLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetMessage(v string) *RecognizeBusinessLicenseResponseBody {
	s.Message = &v
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

type RecognizeBankCardRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeBankCardResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeBankCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBody) SetRequestId(v string) *RecognizeBankCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetData(v string) *RecognizeBankCardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetCode(v string) *RecognizeBankCardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetMessage(v string) *RecognizeBankCardResponseBody {
	s.Message = &v
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

type RecognizeEduPaperCutRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 切题类型
	CutType *string `json:"CutType,omitempty" xml:"CutType,omitempty"`
	// 图片类型
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// 年级学科
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// 是否输出原图坐标信息(如果图片被做过旋转，图片校正等处理)
	OutputOricoord *bool `json:"OutputOricoord,omitempty" xml:"OutputOricoord,omitempty"`
}

func (s RecognizeEduPaperCutRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperCutRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperCutRequest) SetUrl(v string) *RecognizeEduPaperCutRequest {
	s.Url = &v
	return s
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

func (s *RecognizeEduPaperCutRequest) SetOutputOricoord(v bool) *RecognizeEduPaperCutRequest {
	s.OutputOricoord = &v
	return s
}

type RecognizeEduPaperCutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeEduPaperCutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperCutResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperCutResponseBody) SetRequestId(v string) *RecognizeEduPaperCutResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetData(v string) *RecognizeEduPaperCutResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetCode(v string) *RecognizeEduPaperCutResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetMessage(v string) *RecognizeEduPaperCutResponseBody {
	s.Message = &v
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

type RecognizeHouseholdRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeHouseholdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeHouseholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseholdResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHouseholdResponseBody) SetRequestId(v string) *RecognizeHouseholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetData(v string) *RecognizeHouseholdResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetCode(v string) *RecognizeHouseholdResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetMessage(v string) *RecognizeHouseholdResponseBody {
	s.Message = &v
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

type RecognizeEduQuestionOcrRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
}

func (s RecognizeEduQuestionOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduQuestionOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduQuestionOcrRequest) SetUrl(v string) *RecognizeEduQuestionOcrRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduQuestionOcrRequest) SetNeedRotate(v bool) *RecognizeEduQuestionOcrRequest {
	s.NeedRotate = &v
	return s
}

type RecognizeEduQuestionOcrResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeEduQuestionOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduQuestionOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduQuestionOcrResponseBody) SetRequestId(v string) *RecognizeEduQuestionOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetData(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetCode(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetMessage(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Message = &v
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

type RecognizeTrainInvoiceRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeTrainInvoiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeTrainInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTrainInvoiceResponseBody) SetRequestId(v string) *RecognizeTrainInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetData(v string) *RecognizeTrainInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetCode(v string) *RecognizeTrainInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetMessage(v string) *RecognizeTrainInvoiceResponseBody {
	s.Message = &v
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

type RecognizeTableOcrRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeTableOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableOcrRequest) SetUrl(v string) *RecognizeTableOcrRequest {
	s.Url = &v
	return s
}

type RecognizeTableOcrResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeTableOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTableOcrResponseBody) SetRequestId(v string) *RecognizeTableOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetData(v string) *RecognizeTableOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetCode(v string) *RecognizeTableOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetMessage(v string) *RecognizeTableOcrResponseBody {
	s.Message = &v
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

type RecognizeEnglishRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
}

func (s RecognizeEnglishRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEnglishRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEnglishRequest) SetUrl(v string) *RecognizeEnglishRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEnglishRequest) SetNeedRotate(v bool) *RecognizeEnglishRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeEnglishRequest) SetOutputTable(v bool) *RecognizeEnglishRequest {
	s.OutputTable = &v
	return s
}

type RecognizeEnglishResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeEnglishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEnglishResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEnglishResponseBody) SetRequestId(v string) *RecognizeEnglishResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetData(v string) *RecognizeEnglishResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetCode(v string) *RecognizeEnglishResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetMessage(v string) *RecognizeEnglishResponseBody {
	s.Message = &v
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

type RecognizeMultiLanguageRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 识别语种
	Languages []*string `json:"Languages,omitempty" xml:"Languages,omitempty" type:"Repeated"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
}

func (s RecognizeMultiLanguageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageRequest) SetUrl(v string) *RecognizeMultiLanguageRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetLanguages(v []*string) *RecognizeMultiLanguageRequest {
	s.Languages = v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetOutputCharInfo(v bool) *RecognizeMultiLanguageRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetNeedRotate(v bool) *RecognizeMultiLanguageRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetOutputTable(v bool) *RecognizeMultiLanguageRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetNeedSortPage(v bool) *RecognizeMultiLanguageRequest {
	s.NeedSortPage = &v
	return s
}

type RecognizeMultiLanguageShrinkRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 识别语种
	LanguagesShrink *string `json:"Languages,omitempty" xml:"Languages,omitempty"`
	// 是否输出单字识别结果
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// 是否输出表格识别结果，包含单元格信息
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// 是否按顺序输出文字块。false表示从左往右，从上到下的顺序；true表示从上到下，从左往右的顺序
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
}

func (s RecognizeMultiLanguageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageShrinkRequest) SetUrl(v string) *RecognizeMultiLanguageShrinkRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetLanguagesShrink(v string) *RecognizeMultiLanguageShrinkRequest {
	s.LanguagesShrink = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetOutputCharInfo(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetNeedRotate(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetOutputTable(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetNeedSortPage(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.NeedSortPage = &v
	return s
}

type RecognizeMultiLanguageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeMultiLanguageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageResponseBody) SetRequestId(v string) *RecognizeMultiLanguageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetData(v string) *RecognizeMultiLanguageResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetCode(v string) *RecognizeMultiLanguageResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetMessage(v string) *RecognizeMultiLanguageResponseBody {
	s.Message = &v
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

type RecognizeEduOralCalculationRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeEduOralCalculationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeEduOralCalculationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduOralCalculationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduOralCalculationResponseBody) SetRequestId(v string) *RecognizeEduOralCalculationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetData(v string) *RecognizeEduOralCalculationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetCode(v string) *RecognizeEduOralCalculationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetMessage(v string) *RecognizeEduOralCalculationResponseBody {
	s.Message = &v
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

type RecognizeQuotaInvoiceRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeQuotaInvoiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBody) SetRequestId(v string) *RecognizeQuotaInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetData(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetCode(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetMessage(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Message = &v
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

type RecognizeGeneralRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type RecognizeGeneralResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeGeneralResponseBody) SetRequestId(v string) *RecognizeGeneralResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetData(v string) *RecognizeGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetCode(v string) *RecognizeGeneralResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetMessage(v string) *RecognizeGeneralResponseBody {
	s.Message = &v
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

type RecognizeEduPaperStructedRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 是否需要自动旋转功能(结构化检测、混贴场景、教育相关场景会自动做旋转，无需设置)，返回角度信息
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
}

func (s RecognizeEduPaperStructedRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperStructedRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperStructedRequest) SetUrl(v string) *RecognizeEduPaperStructedRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduPaperStructedRequest) SetNeedRotate(v bool) *RecognizeEduPaperStructedRequest {
	s.NeedRotate = &v
	return s
}

type RecognizeEduPaperStructedResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeEduPaperStructedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperStructedResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperStructedResponseBody) SetRequestId(v string) *RecognizeEduPaperStructedResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetData(v string) *RecognizeEduPaperStructedResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetCode(v string) *RecognizeEduPaperStructedResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetMessage(v string) *RecognizeEduPaperStructedResponseBody {
	s.Message = &v
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

type RecognizeIdcardRequest struct {
	// 图片链接（长度不超 1014，不支持 base64）
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeIdcardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdcardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIdcardRequest) SetUrl(v string) *RecognizeIdcardRequest {
	s.Url = &v
	return s
}

type RecognizeIdcardResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecognizeIdcardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdcardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIdcardResponseBody) SetRequestId(v string) *RecognizeIdcardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetData(v string) *RecognizeIdcardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetCode(v string) *RecognizeIdcardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetMessage(v string) *RecognizeIdcardResponseBody {
	s.Message = &v
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

func (client *Client) RecognizeDrivingLicenseWithOptions(request *RecognizeDrivingLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeDrivingLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeDrivingLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeDrivingLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeKoreanWithOptions(request *RecognizeKoreanRequest, runtime *util.RuntimeOptions) (_result *RecognizeKoreanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeKoreanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeKorean"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeCarInvoiceWithOptions(request *RecognizeCarInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeCarInvoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeCarInvoice"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeMixedInvoicesWithOptions(request *RecognizeMixedInvoicesRequest, runtime *util.RuntimeOptions) (_result *RecognizeMixedInvoicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeMixedInvoicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeMixedInvoices"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeEstateCertificationWithOptions(request *RecognizeEstateCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeEstateCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeEstateCertificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeEstateCertification"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeCarNumberWithOptions(request *RecognizeCarNumberRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeCarNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeCarNumber"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeEduPaperOcrWithOptions(request *RecognizeEduPaperOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeEduPaperOcrResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeEduPaperOcr"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeTradeMarkCertificationWithOptions(request *RecognizeTradeMarkCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeTradeMarkCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeTradeMarkCertificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeTradeMarkCertification"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeBirthCertificationWithOptions(request *RecognizeBirthCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeBirthCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeBirthCertificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeBirthCertification"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeDoctypeWithOptions(request *RecognizeDoctypeRequest, runtime *util.RuntimeOptions) (_result *RecognizeDoctypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeDoctypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeDoctype"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeDoctype(request *RecognizeDoctypeRequest) (_result *RecognizeDoctypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeDoctypeResponse{}
	_body, _err := client.RecognizeDoctypeWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeBankAccountLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeBankAccountLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeFoodManageLicenseWithOptions(request *RecognizeFoodManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeFoodManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeFoodManageLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeFoodManageLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeRollTicketWithOptions(request *RecognizeRollTicketRequest, runtime *util.RuntimeOptions) (_result *RecognizeRollTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeRollTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeRollTicket"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeEduFormulaWithOptions(request *RecognizeEduFormulaRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduFormulaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeEduFormulaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeEduFormula"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizePassportWithOptions(request *RecognizePassportRequest, runtime *util.RuntimeOptions) (_result *RecognizePassportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizePassportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizePassport"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeTaxiInvoiceWithOptions(request *RecognizeTaxiInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeTaxiInvoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeTaxiInvoice"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeFoodProduceLicenseWithOptions(request *RecognizeFoodProduceLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeFoodProduceLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeFoodProduceLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeFoodProduceLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeMedicalDeviceProduceLicenseWithOptions(request *RecognizeMedicalDeviceProduceLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeMedicalDeviceProduceLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeMedicalDeviceProduceLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeMedicalDeviceProduceLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeHandwritingWithOptions(request *RecognizeHandwritingRequest, runtime *util.RuntimeOptions) (_result *RecognizeHandwritingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeHandwritingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeHandwriting"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeAirItineraryWithOptions(request *RecognizeAirItineraryRequest, runtime *util.RuntimeOptions) (_result *RecognizeAirItineraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeAirItineraryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeAirItinerary"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeJanpaneseWithOptions(request *RecognizeJanpaneseRequest, runtime *util.RuntimeOptions) (_result *RecognizeJanpaneseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeJanpaneseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeJanpanese"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeCtwoMedicalDeviceManageLicenseWithOptions(request *RecognizeCtwoMedicalDeviceManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeCtwoMedicalDeviceManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeCtwoMedicalDeviceManageLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeCtwoMedicalDeviceManageLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeThaiWithOptions(request *RecognizeThaiRequest, runtime *util.RuntimeOptions) (_result *RecognizeThaiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeThaiResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeThai"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeMedicalDeviceManageLicenseWithOptions(request *RecognizeMedicalDeviceManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeMedicalDeviceManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeMedicalDeviceManageLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeMedicalDeviceManageLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeLatinWithOptions(request *RecognizeLatinRequest, runtime *util.RuntimeOptions) (_result *RecognizeLatinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeLatinResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeLatin"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeInvoiceWithOptions(request *RecognizeInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeInvoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeInvoice"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeMixedCardsWithOptions(request *RecognizeMixedCardsRequest, runtime *util.RuntimeOptions) (_result *RecognizeMixedCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeMixedCardsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeMixedCards"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeMixedCards(request *RecognizeMixedCardsRequest) (_result *RecognizeMixedCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMixedCardsResponse{}
	_body, _err := client.RecognizeMixedCardsWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeWaybillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeWaybill"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeCarVinCodeWithOptions(request *RecognizeCarVinCodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarVinCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeCarVinCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeCarVinCode"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeAdvancedWithOptions(request *RecognizeAdvancedRequest, runtime *util.RuntimeOptions) (_result *RecognizeAdvancedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeAdvancedResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeAdvanced"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeVehicleLicenseWithOptions(request *RecognizeVehicleLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeVehicleLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeVehicleLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeRussianWithOptions(request *RecognizeRussianRequest, runtime *util.RuntimeOptions) (_result *RecognizeRussianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeRussianResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeRussian"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeHouseCertificationWithOptions(request *RecognizeHouseCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeHouseCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeHouseCertificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeHouseCertification"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeHouseCertification(request *RecognizeHouseCertificationRequest) (_result *RecognizeHouseCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHouseCertificationResponse{}
	_body, _err := client.RecognizeHouseCertificationWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeBasicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeBasic"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeBusinessLicenseWithOptions(request *RecognizeBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeBusinessLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeBusinessLicense"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeBankCardWithOptions(request *RecognizeBankCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeBankCardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeBankCard"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeEduPaperCutWithOptions(request *RecognizeEduPaperCutRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperCutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeEduPaperCutResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeEduPaperCut"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeHouseholdWithOptions(request *RecognizeHouseholdRequest, runtime *util.RuntimeOptions) (_result *RecognizeHouseholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeHouseholdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeHousehold"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeEduQuestionOcrWithOptions(request *RecognizeEduQuestionOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduQuestionOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeEduQuestionOcrResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeEduQuestionOcr"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeTrainInvoiceWithOptions(request *RecognizeTrainInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTrainInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeTrainInvoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeTrainInvoice"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeTableOcrWithOptions(request *RecognizeTableOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeTableOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeTableOcrResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeTableOcr"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeEnglishWithOptions(request *RecognizeEnglishRequest, runtime *util.RuntimeOptions) (_result *RecognizeEnglishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeEnglishResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeEnglish"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeMultiLanguageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeMultiLanguage"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeEduOralCalculationWithOptions(request *RecognizeEduOralCalculationRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduOralCalculationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeEduOralCalculationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeEduOralCalculation"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeQuotaInvoiceWithOptions(request *RecognizeQuotaInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeQuotaInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeQuotaInvoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeQuotaInvoice"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeGeneralWithOptions(request *RecognizeGeneralRequest, runtime *util.RuntimeOptions) (_result *RecognizeGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeGeneralResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeGeneral"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeEduPaperStructedWithOptions(request *RecognizeEduPaperStructedRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperStructedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeEduPaperStructedResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeEduPaperStructed"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeIdcardWithOptions(request *RecognizeIdcardRequest, runtime *util.RuntimeOptions) (_result *RecognizeIdcardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecognizeIdcardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeIdcard"), tea.String("2021-07-07"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
