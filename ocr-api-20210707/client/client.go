// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type RecognizeAdvancedRequest struct {
	NeedRotate     *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	NeedSortPage   *bool     `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	NoStamp        *bool     `json:"NoStamp,omitempty" xml:"NoStamp,omitempty"`
	OutputCharInfo *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputFigure   *bool     `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	OutputTable    *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Paragraph      *bool     `json:"Paragraph,omitempty" xml:"Paragraph,omitempty"`
	Row            *bool     `json:"Row,omitempty" xml:"Row,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeAdvancedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeAdvancedResponse) SetStatusCode(v int32) *RecognizeAdvancedResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeAdvancedResponse) SetBody(v *RecognizeAdvancedResponseBody) *RecognizeAdvancedResponse {
	s.Body = v
	return s
}

type RecognizeAirItineraryRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeAirItineraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeAirItineraryResponse) SetStatusCode(v int32) *RecognizeAirItineraryResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeAirItineraryResponse) SetBody(v *RecognizeAirItineraryResponseBody) *RecognizeAirItineraryResponse {
	s.Body = v
	return s
}

type RecognizeBankAcceptanceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankAcceptanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAcceptanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankAcceptanceRequest) SetUrl(v string) *RecognizeBankAcceptanceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBankAcceptanceRequest) SetBody(v io.Reader) *RecognizeBankAcceptanceRequest {
	s.Body = v
	return s
}

type RecognizeBankAcceptanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBankAcceptanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAcceptanceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankAcceptanceResponseBody) SetCode(v string) *RecognizeBankAcceptanceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBankAcceptanceResponseBody) SetData(v string) *RecognizeBankAcceptanceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBankAcceptanceResponseBody) SetMessage(v string) *RecognizeBankAcceptanceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBankAcceptanceResponseBody) SetRequestId(v string) *RecognizeBankAcceptanceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBankAcceptanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBankAcceptanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBankAcceptanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAcceptanceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBankAcceptanceResponse) SetHeaders(v map[string]*string) *RecognizeBankAcceptanceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBankAcceptanceResponse) SetStatusCode(v int32) *RecognizeBankAcceptanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBankAcceptanceResponse) SetBody(v *RecognizeBankAcceptanceResponseBody) *RecognizeBankAcceptanceResponse {
	s.Body = v
	return s
}

type RecognizeBankAccountLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBankAccountLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBankAccountLicenseResponse) SetStatusCode(v int32) *RecognizeBankAccountLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponse) SetBody(v *RecognizeBankAccountLicenseResponseBody) *RecognizeBankAccountLicenseResponse {
	s.Body = v
	return s
}

type RecognizeBankCardRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBankCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBankCardResponse) SetStatusCode(v int32) *RecognizeBankCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBankCardResponse) SetBody(v *RecognizeBankCardResponseBody) *RecognizeBankCardResponse {
	s.Body = v
	return s
}

type RecognizeBasicRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBasicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBasicResponse) SetStatusCode(v int32) *RecognizeBasicResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBasicResponse) SetBody(v *RecognizeBasicResponseBody) *RecognizeBasicResponse {
	s.Body = v
	return s
}

type RecognizeBirthCertificationRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBirthCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBirthCertificationResponse) SetStatusCode(v int32) *RecognizeBirthCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBirthCertificationResponse) SetBody(v *RecognizeBirthCertificationResponseBody) *RecognizeBirthCertificationResponse {
	s.Body = v
	return s
}

type RecognizeBusShipTicketRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBusShipTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBusShipTicketResponse) SetStatusCode(v int32) *RecognizeBusShipTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBusShipTicketResponse) SetBody(v *RecognizeBusShipTicketResponseBody) *RecognizeBusShipTicketResponse {
	s.Body = v
	return s
}

type RecognizeBusinessLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBusinessLicenseResponse) SetStatusCode(v int32) *RecognizeBusinessLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBusinessLicenseResponse) SetBody(v *RecognizeBusinessLicenseResponseBody) *RecognizeBusinessLicenseResponse {
	s.Body = v
	return s
}

type RecognizeCarInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeCarInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeCarInvoiceResponse) SetStatusCode(v int32) *RecognizeCarInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCarInvoiceResponse) SetBody(v *RecognizeCarInvoiceResponseBody) *RecognizeCarInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeCarNumberRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeCarNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeCarNumberResponse) SetStatusCode(v int32) *RecognizeCarNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCarNumberResponse) SetBody(v *RecognizeCarNumberResponseBody) *RecognizeCarNumberResponse {
	s.Body = v
	return s
}

type RecognizeCarVinCodeRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeCarVinCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeCarVinCodeResponse) SetStatusCode(v int32) *RecognizeCarVinCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCarVinCodeResponse) SetBody(v *RecognizeCarVinCodeResponseBody) *RecognizeCarVinCodeResponse {
	s.Body = v
	return s
}

type RecognizeChinesePassportRequest struct {
	OutputFigure *bool     `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	Url          *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body         io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeChinesePassportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeChinesePassportResponse) SetStatusCode(v int32) *RecognizeChinesePassportResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeChinesePassportResponse) SetBody(v *RecognizeChinesePassportResponseBody) *RecognizeChinesePassportResponse {
	s.Body = v
	return s
}

type RecognizeCommonPrintedInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeCommonPrintedInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeCommonPrintedInvoiceResponse) SetStatusCode(v int32) *RecognizeCommonPrintedInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponse) SetBody(v *RecognizeCommonPrintedInvoiceResponseBody) *RecognizeCommonPrintedInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeCosmeticProduceLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCosmeticProduceLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCosmeticProduceLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCosmeticProduceLicenseRequest) SetUrl(v string) *RecognizeCosmeticProduceLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseRequest) SetBody(v io.Reader) *RecognizeCosmeticProduceLicenseRequest {
	s.Body = v
	return s
}

type RecognizeCosmeticProduceLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCosmeticProduceLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCosmeticProduceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCosmeticProduceLicenseResponseBody) SetCode(v string) *RecognizeCosmeticProduceLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponseBody) SetData(v string) *RecognizeCosmeticProduceLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponseBody) SetMessage(v string) *RecognizeCosmeticProduceLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponseBody) SetRequestId(v string) *RecognizeCosmeticProduceLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCosmeticProduceLicenseResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeCosmeticProduceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeCosmeticProduceLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCosmeticProduceLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCosmeticProduceLicenseResponse) SetHeaders(v map[string]*string) *RecognizeCosmeticProduceLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponse) SetStatusCode(v int32) *RecognizeCosmeticProduceLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponse) SetBody(v *RecognizeCosmeticProduceLicenseResponseBody) *RecognizeCosmeticProduceLicenseResponse {
	s.Body = v
	return s
}

type RecognizeCovidTestReportRequest struct {
	MultipleResult *bool     `json:"MultipleResult,omitempty" xml:"MultipleResult,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCovidTestReportRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCovidTestReportRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCovidTestReportRequest) SetMultipleResult(v bool) *RecognizeCovidTestReportRequest {
	s.MultipleResult = &v
	return s
}

func (s *RecognizeCovidTestReportRequest) SetUrl(v string) *RecognizeCovidTestReportRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCovidTestReportRequest) SetBody(v io.Reader) *RecognizeCovidTestReportRequest {
	s.Body = v
	return s
}

type RecognizeCovidTestReportResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCovidTestReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCovidTestReportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCovidTestReportResponseBody) SetCode(v string) *RecognizeCovidTestReportResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCovidTestReportResponseBody) SetData(v string) *RecognizeCovidTestReportResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCovidTestReportResponseBody) SetMessage(v string) *RecognizeCovidTestReportResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCovidTestReportResponseBody) SetRequestId(v string) *RecognizeCovidTestReportResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCovidTestReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeCovidTestReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeCovidTestReportResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCovidTestReportResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCovidTestReportResponse) SetHeaders(v map[string]*string) *RecognizeCovidTestReportResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCovidTestReportResponse) SetStatusCode(v int32) *RecognizeCovidTestReportResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCovidTestReportResponse) SetBody(v *RecognizeCovidTestReportResponseBody) *RecognizeCovidTestReportResponse {
	s.Body = v
	return s
}

type RecognizeCtwoMedicalDeviceManageLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeCtwoMedicalDeviceManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponse) SetStatusCode(v int32) *RecognizeCtwoMedicalDeviceManageLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponse) SetBody(v *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) *RecognizeCtwoMedicalDeviceManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeDocumentStructureRequest struct {
	NeedRotate        *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	NeedSortPage      *bool     `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	NoStamp           *bool     `json:"NoStamp,omitempty" xml:"NoStamp,omitempty"`
	OutputCharInfo    *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable       *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Page              *bool     `json:"Page,omitempty" xml:"Page,omitempty"`
	Paragraph         *bool     `json:"Paragraph,omitempty" xml:"Paragraph,omitempty"`
	Row               *bool     `json:"Row,omitempty" xml:"Row,omitempty"`
	Url               *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	UseNewStyleOutput *bool     `json:"UseNewStyleOutput,omitempty" xml:"UseNewStyleOutput,omitempty"`
	Body              io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeDocumentStructureRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDocumentStructureRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDocumentStructureRequest) SetNeedRotate(v bool) *RecognizeDocumentStructureRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetNeedSortPage(v bool) *RecognizeDocumentStructureRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetNoStamp(v bool) *RecognizeDocumentStructureRequest {
	s.NoStamp = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetOutputCharInfo(v bool) *RecognizeDocumentStructureRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetOutputTable(v bool) *RecognizeDocumentStructureRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetPage(v bool) *RecognizeDocumentStructureRequest {
	s.Page = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetParagraph(v bool) *RecognizeDocumentStructureRequest {
	s.Paragraph = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetRow(v bool) *RecognizeDocumentStructureRequest {
	s.Row = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetUrl(v string) *RecognizeDocumentStructureRequest {
	s.Url = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetUseNewStyleOutput(v bool) *RecognizeDocumentStructureRequest {
	s.UseNewStyleOutput = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetBody(v io.Reader) *RecognizeDocumentStructureRequest {
	s.Body = v
	return s
}

type RecognizeDocumentStructureResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDocumentStructureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDocumentStructureResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDocumentStructureResponseBody) SetCode(v string) *RecognizeDocumentStructureResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeDocumentStructureResponseBody) SetData(v string) *RecognizeDocumentStructureResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeDocumentStructureResponseBody) SetMessage(v string) *RecognizeDocumentStructureResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeDocumentStructureResponseBody) SetRequestId(v string) *RecognizeDocumentStructureResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeDocumentStructureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeDocumentStructureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeDocumentStructureResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDocumentStructureResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDocumentStructureResponse) SetHeaders(v map[string]*string) *RecognizeDocumentStructureResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDocumentStructureResponse) SetStatusCode(v int32) *RecognizeDocumentStructureResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeDocumentStructureResponse) SetBody(v *RecognizeDocumentStructureResponseBody) *RecognizeDocumentStructureResponse {
	s.Body = v
	return s
}

type RecognizeDrivingLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeDrivingLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeDrivingLicenseResponse) SetStatusCode(v int32) *RecognizeDrivingLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeDrivingLicenseResponse) SetBody(v *RecognizeDrivingLicenseResponseBody) *RecognizeDrivingLicenseResponse {
	s.Body = v
	return s
}

type RecognizeEduFormulaRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeEduFormulaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeEduFormulaResponse) SetStatusCode(v int32) *RecognizeEduFormulaResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduFormulaResponse) SetBody(v *RecognizeEduFormulaResponseBody) *RecognizeEduFormulaResponse {
	s.Body = v
	return s
}

type RecognizeEduOralCalculationRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeEduOralCalculationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeEduOralCalculationResponse) SetStatusCode(v int32) *RecognizeEduOralCalculationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduOralCalculationResponse) SetBody(v *RecognizeEduOralCalculationResponseBody) *RecognizeEduOralCalculationResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperCutRequest struct {
	CutType   *string   `json:"CutType,omitempty" xml:"CutType,omitempty"`
	ImageType *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Subject   *string   `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Url       *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body      io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeEduPaperCutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeEduPaperCutResponse) SetStatusCode(v int32) *RecognizeEduPaperCutResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduPaperCutResponse) SetBody(v *RecognizeEduPaperCutResponseBody) *RecognizeEduPaperCutResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperOcrRequest struct {
	ImageType      *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	OutputOricoord *bool     `json:"OutputOricoord,omitempty" xml:"OutputOricoord,omitempty"`
	Subject        *string   `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeEduPaperOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeEduPaperOcrResponse) SetStatusCode(v int32) *RecognizeEduPaperOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduPaperOcrResponse) SetBody(v *RecognizeEduPaperOcrResponseBody) *RecognizeEduPaperOcrResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperStructedRequest struct {
	NeedRotate *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	Subject    *string   `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Url        *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body       io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeEduPaperStructedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeEduPaperStructedResponse) SetStatusCode(v int32) *RecognizeEduPaperStructedResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduPaperStructedResponse) SetBody(v *RecognizeEduPaperStructedResponseBody) *RecognizeEduPaperStructedResponse {
	s.Body = v
	return s
}

type RecognizeEduQuestionOcrRequest struct {
	NeedRotate *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	Url        *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body       io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeEduQuestionOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeEduQuestionOcrResponse) SetStatusCode(v int32) *RecognizeEduQuestionOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponse) SetBody(v *RecognizeEduQuestionOcrResponseBody) *RecognizeEduQuestionOcrResponse {
	s.Body = v
	return s
}

type RecognizeEnglishRequest struct {
	NeedRotate  *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	OutputTable *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url         *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body        io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeEnglishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeEnglishResponse) SetStatusCode(v int32) *RecognizeEnglishResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEnglishResponse) SetBody(v *RecognizeEnglishResponseBody) *RecognizeEnglishResponse {
	s.Body = v
	return s
}

type RecognizeEstateCertificationRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeEstateCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeEstateCertificationResponse) SetStatusCode(v int32) *RecognizeEstateCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEstateCertificationResponse) SetBody(v *RecognizeEstateCertificationResponseBody) *RecognizeEstateCertificationResponse {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToHKRequest struct {
	OutputFigure *bool     `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	Url          *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body         io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeExitEntryPermitToHKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeExitEntryPermitToHKResponse) SetStatusCode(v int32) *RecognizeExitEntryPermitToHKResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponse) SetBody(v *RecognizeExitEntryPermitToHKResponseBody) *RecognizeExitEntryPermitToHKResponse {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToMainlandRequest struct {
	OutputFigure *bool     `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	Url          *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body         io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeExitEntryPermitToMainlandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeExitEntryPermitToMainlandResponse) SetStatusCode(v int32) *RecognizeExitEntryPermitToMainlandResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponse) SetBody(v *RecognizeExitEntryPermitToMainlandResponseBody) *RecognizeExitEntryPermitToMainlandResponse {
	s.Body = v
	return s
}

type RecognizeFoodManageLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeFoodManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeFoodManageLicenseResponse) SetStatusCode(v int32) *RecognizeFoodManageLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponse) SetBody(v *RecognizeFoodManageLicenseResponseBody) *RecognizeFoodManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeFoodProduceLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeFoodProduceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeFoodProduceLicenseResponse) SetStatusCode(v int32) *RecognizeFoodProduceLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponse) SetBody(v *RecognizeFoodProduceLicenseResponseBody) *RecognizeFoodProduceLicenseResponse {
	s.Body = v
	return s
}

type RecognizeGeneralRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeGeneralResponse) SetStatusCode(v int32) *RecognizeGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeGeneralResponse) SetBody(v *RecognizeGeneralResponseBody) *RecognizeGeneralResponse {
	s.Body = v
	return s
}

type RecognizeHandwritingRequest struct {
	NeedRotate     *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	NeedSortPage   *bool     `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	OutputCharInfo *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable    *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeHandwritingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeHandwritingResponse) SetStatusCode(v int32) *RecognizeHandwritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHandwritingResponse) SetBody(v *RecognizeHandwritingResponseBody) *RecognizeHandwritingResponse {
	s.Body = v
	return s
}

type RecognizeHealthCodeRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHealthCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHealthCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHealthCodeRequest) SetUrl(v string) *RecognizeHealthCodeRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHealthCodeRequest) SetBody(v io.Reader) *RecognizeHealthCodeRequest {
	s.Body = v
	return s
}

type RecognizeHealthCodeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHealthCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHealthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHealthCodeResponseBody) SetCode(v string) *RecognizeHealthCodeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHealthCodeResponseBody) SetData(v string) *RecognizeHealthCodeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHealthCodeResponseBody) SetMessage(v string) *RecognizeHealthCodeResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHealthCodeResponseBody) SetRequestId(v string) *RecognizeHealthCodeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHealthCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeHealthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeHealthCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHealthCodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHealthCodeResponse) SetHeaders(v map[string]*string) *RecognizeHealthCodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHealthCodeResponse) SetStatusCode(v int32) *RecognizeHealthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHealthCodeResponse) SetBody(v *RecognizeHealthCodeResponseBody) *RecognizeHealthCodeResponse {
	s.Body = v
	return s
}

type RecognizeHotelConsumeRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHotelConsumeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHotelConsumeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHotelConsumeRequest) SetUrl(v string) *RecognizeHotelConsumeRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHotelConsumeRequest) SetBody(v io.Reader) *RecognizeHotelConsumeRequest {
	s.Body = v
	return s
}

type RecognizeHotelConsumeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHotelConsumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHotelConsumeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHotelConsumeResponseBody) SetCode(v string) *RecognizeHotelConsumeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHotelConsumeResponseBody) SetData(v string) *RecognizeHotelConsumeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHotelConsumeResponseBody) SetMessage(v string) *RecognizeHotelConsumeResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHotelConsumeResponseBody) SetRequestId(v string) *RecognizeHotelConsumeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHotelConsumeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeHotelConsumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeHotelConsumeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHotelConsumeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHotelConsumeResponse) SetHeaders(v map[string]*string) *RecognizeHotelConsumeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHotelConsumeResponse) SetStatusCode(v int32) *RecognizeHotelConsumeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHotelConsumeResponse) SetBody(v *RecognizeHotelConsumeResponseBody) *RecognizeHotelConsumeResponse {
	s.Body = v
	return s
}

type RecognizeHouseholdRequest struct {
	IsResidentPage *bool     `json:"IsResidentPage,omitempty" xml:"IsResidentPage,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHouseholdRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseholdRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHouseholdRequest) SetIsResidentPage(v bool) *RecognizeHouseholdRequest {
	s.IsResidentPage = &v
	return s
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeHouseholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeHouseholdResponse) SetStatusCode(v int32) *RecognizeHouseholdResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHouseholdResponse) SetBody(v *RecognizeHouseholdResponseBody) *RecognizeHouseholdResponse {
	s.Body = v
	return s
}

type RecognizeIdcardRequest struct {
	OutputFigure      *bool     `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	OutputQualityInfo *bool     `json:"OutputQualityInfo,omitempty" xml:"OutputQualityInfo,omitempty"`
	Url               *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body              io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *RecognizeIdcardRequest) SetOutputQualityInfo(v bool) *RecognizeIdcardRequest {
	s.OutputQualityInfo = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeIdcardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeIdcardResponse) SetStatusCode(v int32) *RecognizeIdcardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeIdcardResponse) SetBody(v *RecognizeIdcardResponseBody) *RecognizeIdcardResponse {
	s.Body = v
	return s
}

type RecognizeInternationalBusinessLicenseRequest struct {
	Country *string   `json:"Country,omitempty" xml:"Country,omitempty"`
	Url     *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body    io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInternationalBusinessLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalBusinessLicenseRequest) SetCountry(v string) *RecognizeInternationalBusinessLicenseRequest {
	s.Country = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseRequest) SetUrl(v string) *RecognizeInternationalBusinessLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseRequest) SetBody(v io.Reader) *RecognizeInternationalBusinessLicenseRequest {
	s.Body = v
	return s
}

type RecognizeInternationalBusinessLicenseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeInternationalBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalBusinessLicenseResponseBody) SetCode(v string) *RecognizeInternationalBusinessLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponseBody) SetData(v string) *RecognizeInternationalBusinessLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponseBody) SetMessage(v string) *RecognizeInternationalBusinessLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponseBody) SetRequestId(v string) *RecognizeInternationalBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeInternationalBusinessLicenseResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeInternationalBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeInternationalBusinessLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalBusinessLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalBusinessLicenseResponse) SetHeaders(v map[string]*string) *RecognizeInternationalBusinessLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponse) SetStatusCode(v int32) *RecognizeInternationalBusinessLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponse) SetBody(v *RecognizeInternationalBusinessLicenseResponseBody) *RecognizeInternationalBusinessLicenseResponse {
	s.Body = v
	return s
}

type RecognizeInternationalIdcardRequest struct {
	Country *string   `json:"Country,omitempty" xml:"Country,omitempty"`
	Url     *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body    io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInternationalIdcardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalIdcardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalIdcardRequest) SetCountry(v string) *RecognizeInternationalIdcardRequest {
	s.Country = &v
	return s
}

func (s *RecognizeInternationalIdcardRequest) SetUrl(v string) *RecognizeInternationalIdcardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeInternationalIdcardRequest) SetBody(v io.Reader) *RecognizeInternationalIdcardRequest {
	s.Body = v
	return s
}

type RecognizeInternationalIdcardResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeInternationalIdcardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalIdcardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalIdcardResponseBody) SetCode(v string) *RecognizeInternationalIdcardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeInternationalIdcardResponseBody) SetData(v string) *RecognizeInternationalIdcardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeInternationalIdcardResponseBody) SetMessage(v string) *RecognizeInternationalIdcardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeInternationalIdcardResponseBody) SetRequestId(v string) *RecognizeInternationalIdcardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeInternationalIdcardResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeInternationalIdcardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeInternationalIdcardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalIdcardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalIdcardResponse) SetHeaders(v map[string]*string) *RecognizeInternationalIdcardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeInternationalIdcardResponse) SetStatusCode(v int32) *RecognizeInternationalIdcardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeInternationalIdcardResponse) SetBody(v *RecognizeInternationalIdcardResponseBody) *RecognizeInternationalIdcardResponse {
	s.Body = v
	return s
}

type RecognizeInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeInvoiceResponse) SetStatusCode(v int32) *RecognizeInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeInvoiceResponse) SetBody(v *RecognizeInvoiceResponseBody) *RecognizeInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeJanpaneseRequest struct {
	NeedRotate     *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	OutputCharInfo *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable    *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeJanpaneseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeJanpaneseResponse) SetStatusCode(v int32) *RecognizeJanpaneseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeJanpaneseResponse) SetBody(v *RecognizeJanpaneseResponseBody) *RecognizeJanpaneseResponse {
	s.Body = v
	return s
}

type RecognizeKoreanRequest struct {
	NeedRotate     *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	OutputCharInfo *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable    *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeKoreanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeKoreanResponse) SetStatusCode(v int32) *RecognizeKoreanResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeKoreanResponse) SetBody(v *RecognizeKoreanResponseBody) *RecognizeKoreanResponse {
	s.Body = v
	return s
}

type RecognizeLatinRequest struct {
	NeedRotate     *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	OutputCharInfo *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable    *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeLatinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeLatinResponse) SetStatusCode(v int32) *RecognizeLatinResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeLatinResponse) SetBody(v *RecognizeLatinResponseBody) *RecognizeLatinResponse {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceManageLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeMedicalDeviceManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeMedicalDeviceManageLicenseResponse) SetStatusCode(v int32) *RecognizeMedicalDeviceManageLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponse) SetBody(v *RecognizeMedicalDeviceManageLicenseResponseBody) *RecognizeMedicalDeviceManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceProduceLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeMedicalDeviceProduceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeMedicalDeviceProduceLicenseResponse) SetStatusCode(v int32) *RecognizeMedicalDeviceProduceLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponse) SetBody(v *RecognizeMedicalDeviceProduceLicenseResponseBody) *RecognizeMedicalDeviceProduceLicenseResponse {
	s.Body = v
	return s
}

type RecognizeMixedInvoicesRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeMixedInvoicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeMixedInvoicesResponse) SetStatusCode(v int32) *RecognizeMixedInvoicesResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMixedInvoicesResponse) SetBody(v *RecognizeMixedInvoicesResponseBody) *RecognizeMixedInvoicesResponse {
	s.Body = v
	return s
}

type RecognizeMultiLanguageRequest struct {
	Languages      []*string `json:"Languages,omitempty" xml:"Languages,omitempty" type:"Repeated"`
	NeedRotate     *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	NeedSortPage   *bool     `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	OutputCharInfo *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable    *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	LanguagesShrink *string   `json:"Languages,omitempty" xml:"Languages,omitempty"`
	NeedRotate      *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	NeedSortPage    *bool     `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	OutputCharInfo  *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable     *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url             *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body            io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeMultiLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeMultiLanguageResponse) SetStatusCode(v int32) *RecognizeMultiLanguageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMultiLanguageResponse) SetBody(v *RecognizeMultiLanguageResponseBody) *RecognizeMultiLanguageResponse {
	s.Body = v
	return s
}

type RecognizeNonTaxInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeNonTaxInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeNonTaxInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeNonTaxInvoiceRequest) SetUrl(v string) *RecognizeNonTaxInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeNonTaxInvoiceRequest) SetBody(v io.Reader) *RecognizeNonTaxInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeNonTaxInvoiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeNonTaxInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeNonTaxInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeNonTaxInvoiceResponseBody) SetCode(v string) *RecognizeNonTaxInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeNonTaxInvoiceResponseBody) SetData(v string) *RecognizeNonTaxInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeNonTaxInvoiceResponseBody) SetMessage(v string) *RecognizeNonTaxInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeNonTaxInvoiceResponseBody) SetRequestId(v string) *RecognizeNonTaxInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeNonTaxInvoiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeNonTaxInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeNonTaxInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeNonTaxInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeNonTaxInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeNonTaxInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeNonTaxInvoiceResponse) SetStatusCode(v int32) *RecognizeNonTaxInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeNonTaxInvoiceResponse) SetBody(v *RecognizeNonTaxInvoiceResponseBody) *RecognizeNonTaxInvoiceResponse {
	s.Body = v
	return s
}

type RecognizePassportRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizePassportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizePassportResponse) SetStatusCode(v int32) *RecognizePassportResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePassportResponse) SetBody(v *RecognizePassportResponseBody) *RecognizePassportResponse {
	s.Body = v
	return s
}

type RecognizePaymentRecordRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePaymentRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePaymentRecordRequest) GoString() string {
	return s.String()
}

func (s *RecognizePaymentRecordRequest) SetUrl(v string) *RecognizePaymentRecordRequest {
	s.Url = &v
	return s
}

func (s *RecognizePaymentRecordRequest) SetBody(v io.Reader) *RecognizePaymentRecordRequest {
	s.Body = v
	return s
}

type RecognizePaymentRecordResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePaymentRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePaymentRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePaymentRecordResponseBody) SetCode(v string) *RecognizePaymentRecordResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizePaymentRecordResponseBody) SetData(v string) *RecognizePaymentRecordResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizePaymentRecordResponseBody) SetMessage(v string) *RecognizePaymentRecordResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizePaymentRecordResponseBody) SetRequestId(v string) *RecognizePaymentRecordResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePaymentRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizePaymentRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizePaymentRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePaymentRecordResponse) GoString() string {
	return s.String()
}

func (s *RecognizePaymentRecordResponse) SetHeaders(v map[string]*string) *RecognizePaymentRecordResponse {
	s.Headers = v
	return s
}

func (s *RecognizePaymentRecordResponse) SetStatusCode(v int32) *RecognizePaymentRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePaymentRecordResponse) SetBody(v *RecognizePaymentRecordResponseBody) *RecognizePaymentRecordResponse {
	s.Body = v
	return s
}

type RecognizePurchaseRecordRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePurchaseRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePurchaseRecordRequest) GoString() string {
	return s.String()
}

func (s *RecognizePurchaseRecordRequest) SetUrl(v string) *RecognizePurchaseRecordRequest {
	s.Url = &v
	return s
}

func (s *RecognizePurchaseRecordRequest) SetBody(v io.Reader) *RecognizePurchaseRecordRequest {
	s.Body = v
	return s
}

type RecognizePurchaseRecordResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePurchaseRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePurchaseRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePurchaseRecordResponseBody) SetCode(v string) *RecognizePurchaseRecordResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizePurchaseRecordResponseBody) SetData(v string) *RecognizePurchaseRecordResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizePurchaseRecordResponseBody) SetMessage(v string) *RecognizePurchaseRecordResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizePurchaseRecordResponseBody) SetRequestId(v string) *RecognizePurchaseRecordResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePurchaseRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizePurchaseRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizePurchaseRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePurchaseRecordResponse) GoString() string {
	return s.String()
}

func (s *RecognizePurchaseRecordResponse) SetHeaders(v map[string]*string) *RecognizePurchaseRecordResponse {
	s.Headers = v
	return s
}

func (s *RecognizePurchaseRecordResponse) SetStatusCode(v int32) *RecognizePurchaseRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePurchaseRecordResponse) SetBody(v *RecognizePurchaseRecordResponseBody) *RecognizePurchaseRecordResponse {
	s.Body = v
	return s
}

type RecognizeQuotaInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeQuotaInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeQuotaInvoiceResponse) SetStatusCode(v int32) *RecognizeQuotaInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) SetBody(v *RecognizeQuotaInvoiceResponseBody) *RecognizeQuotaInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeRideHailingItineraryRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeRideHailingItineraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeRideHailingItineraryResponse) SetStatusCode(v int32) *RecognizeRideHailingItineraryResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeRideHailingItineraryResponse) SetBody(v *RecognizeRideHailingItineraryResponseBody) *RecognizeRideHailingItineraryResponse {
	s.Body = v
	return s
}

type RecognizeRollTicketRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeRollTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeRollTicketResponse) SetStatusCode(v int32) *RecognizeRollTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeRollTicketResponse) SetBody(v *RecognizeRollTicketResponseBody) *RecognizeRollTicketResponse {
	s.Body = v
	return s
}

type RecognizeRussianRequest struct {
	NeedRotate     *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	OutputCharInfo *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable    *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeRussianResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeRussianResponse) SetStatusCode(v int32) *RecognizeRussianResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeRussianResponse) SetBody(v *RecognizeRussianResponseBody) *RecognizeRussianResponse {
	s.Body = v
	return s
}

type RecognizeShoppingReceiptRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeShoppingReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeShoppingReceiptRequest) GoString() string {
	return s.String()
}

func (s *RecognizeShoppingReceiptRequest) SetUrl(v string) *RecognizeShoppingReceiptRequest {
	s.Url = &v
	return s
}

func (s *RecognizeShoppingReceiptRequest) SetBody(v io.Reader) *RecognizeShoppingReceiptRequest {
	s.Body = v
	return s
}

type RecognizeShoppingReceiptResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeShoppingReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeShoppingReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeShoppingReceiptResponseBody) SetCode(v string) *RecognizeShoppingReceiptResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeShoppingReceiptResponseBody) SetData(v string) *RecognizeShoppingReceiptResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeShoppingReceiptResponseBody) SetMessage(v string) *RecognizeShoppingReceiptResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeShoppingReceiptResponseBody) SetRequestId(v string) *RecognizeShoppingReceiptResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeShoppingReceiptResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeShoppingReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeShoppingReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeShoppingReceiptResponse) GoString() string {
	return s.String()
}

func (s *RecognizeShoppingReceiptResponse) SetHeaders(v map[string]*string) *RecognizeShoppingReceiptResponse {
	s.Headers = v
	return s
}

func (s *RecognizeShoppingReceiptResponse) SetStatusCode(v int32) *RecognizeShoppingReceiptResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeShoppingReceiptResponse) SetBody(v *RecognizeShoppingReceiptResponseBody) *RecognizeShoppingReceiptResponse {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeSocialSecurityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeSocialSecurityCardResponse) SetStatusCode(v int32) *RecognizeSocialSecurityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeSocialSecurityCardResponse) SetBody(v *RecognizeSocialSecurityCardResponseBody) *RecognizeSocialSecurityCardResponse {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardVersionIIRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeSocialSecurityCardVersionIIRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardVersionIIRequest) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardVersionIIRequest) SetUrl(v string) *RecognizeSocialSecurityCardVersionIIRequest {
	s.Url = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIRequest) SetBody(v io.Reader) *RecognizeSocialSecurityCardVersionIIRequest {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardVersionIIResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeSocialSecurityCardVersionIIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardVersionIIResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardVersionIIResponseBody) SetCode(v string) *RecognizeSocialSecurityCardVersionIIResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponseBody) SetData(v string) *RecognizeSocialSecurityCardVersionIIResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponseBody) SetMessage(v string) *RecognizeSocialSecurityCardVersionIIResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponseBody) SetRequestId(v string) *RecognizeSocialSecurityCardVersionIIResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeSocialSecurityCardVersionIIResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeSocialSecurityCardVersionIIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeSocialSecurityCardVersionIIResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardVersionIIResponse) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardVersionIIResponse) SetHeaders(v map[string]*string) *RecognizeSocialSecurityCardVersionIIResponse {
	s.Headers = v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponse) SetStatusCode(v int32) *RecognizeSocialSecurityCardVersionIIResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponse) SetBody(v *RecognizeSocialSecurityCardVersionIIResponseBody) *RecognizeSocialSecurityCardVersionIIResponse {
	s.Body = v
	return s
}

type RecognizeTableOcrRequest struct {
	LineLess      *bool     `json:"LineLess,omitempty" xml:"LineLess,omitempty"`
	NeedRotate    *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	SkipDetection *bool     `json:"SkipDetection,omitempty" xml:"SkipDetection,omitempty"`
	Url           *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body          io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTableOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTableOcrResponse) SetStatusCode(v int32) *RecognizeTableOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTableOcrResponse) SetBody(v *RecognizeTableOcrResponseBody) *RecognizeTableOcrResponse {
	s.Body = v
	return s
}

type RecognizeTaxClearanceCertificateRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTaxClearanceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTaxClearanceCertificateResponse) SetStatusCode(v int32) *RecognizeTaxClearanceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponse) SetBody(v *RecognizeTaxClearanceCertificateResponseBody) *RecognizeTaxClearanceCertificateResponse {
	s.Body = v
	return s
}

type RecognizeTaxiInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTaxiInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTaxiInvoiceResponse) SetStatusCode(v int32) *RecognizeTaxiInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) SetBody(v *RecognizeTaxiInvoiceResponseBody) *RecognizeTaxiInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeThaiRequest struct {
	NeedRotate     *bool     `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	OutputCharInfo *bool     `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	OutputTable    *bool     `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body           io.Reader `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeThaiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeThaiResponse) SetStatusCode(v int32) *RecognizeThaiResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeThaiResponse) SetBody(v *RecognizeThaiResponseBody) *RecognizeThaiResponse {
	s.Body = v
	return s
}

type RecognizeTollInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTollInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTollInvoiceResponse) SetStatusCode(v int32) *RecognizeTollInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTollInvoiceResponse) SetBody(v *RecognizeTollInvoiceResponseBody) *RecognizeTollInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeTradeMarkCertificationRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTradeMarkCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTradeMarkCertificationResponse) SetStatusCode(v int32) *RecognizeTradeMarkCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponse) SetBody(v *RecognizeTradeMarkCertificationResponseBody) *RecognizeTradeMarkCertificationResponse {
	s.Body = v
	return s
}

type RecognizeTrainInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTrainInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTrainInvoiceResponse) SetStatusCode(v int32) *RecognizeTrainInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTrainInvoiceResponse) SetBody(v *RecognizeTrainInvoiceResponseBody) *RecognizeTrainInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeTravelCardRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTravelCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTravelCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTravelCardRequest) SetUrl(v string) *RecognizeTravelCardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTravelCardRequest) SetBody(v io.Reader) *RecognizeTravelCardRequest {
	s.Body = v
	return s
}

type RecognizeTravelCardResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTravelCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTravelCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTravelCardResponseBody) SetCode(v string) *RecognizeTravelCardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTravelCardResponseBody) SetData(v string) *RecognizeTravelCardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTravelCardResponseBody) SetMessage(v string) *RecognizeTravelCardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTravelCardResponseBody) SetRequestId(v string) *RecognizeTravelCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTravelCardResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTravelCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTravelCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTravelCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTravelCardResponse) SetHeaders(v map[string]*string) *RecognizeTravelCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTravelCardResponse) SetStatusCode(v int32) *RecognizeTravelCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTravelCardResponse) SetBody(v *RecognizeTravelCardResponseBody) *RecognizeTravelCardResponse {
	s.Body = v
	return s
}

type RecognizeUsedCarInvoiceRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeUsedCarInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeUsedCarInvoiceResponse) SetStatusCode(v int32) *RecognizeUsedCarInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeUsedCarInvoiceResponse) SetBody(v *RecognizeUsedCarInvoiceResponseBody) *RecognizeUsedCarInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeVehicleCertificationRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVehicleCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeVehicleCertificationResponse) SetStatusCode(v int32) *RecognizeVehicleCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleCertificationResponse) SetBody(v *RecognizeVehicleCertificationResponseBody) *RecognizeVehicleCertificationResponse {
	s.Body = v
	return s
}

type RecognizeVehicleLicenseRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVehicleLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeVehicleLicenseResponse) SetStatusCode(v int32) *RecognizeVehicleLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleLicenseResponse) SetBody(v *RecognizeVehicleLicenseResponseBody) *RecognizeVehicleLicenseResponse {
	s.Body = v
	return s
}

type RecognizeVehicleRegistrationRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVehicleRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeVehicleRegistrationResponse) SetStatusCode(v int32) *RecognizeVehicleRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleRegistrationResponse) SetBody(v *RecognizeVehicleRegistrationResponseBody) *RecognizeVehicleRegistrationResponse {
	s.Body = v
	return s
}

type RecognizeWaybillRequest struct {
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeWaybillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeWaybillResponse) SetStatusCode(v int32) *RecognizeWaybillResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeWaybillResponse) SetBody(v *RecognizeWaybillResponseBody) *RecognizeWaybillResponse {
	s.Body = v
	return s
}

type VerifyBusinessLicenseRequest struct {
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	CreditCode  *string `json:"CreditCode,omitempty" xml:"CreditCode,omitempty"`
	LegalPerson *string `json:"LegalPerson,omitempty" xml:"LegalPerson,omitempty"`
}

func (s VerifyBusinessLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *VerifyBusinessLicenseRequest) SetCompanyName(v string) *VerifyBusinessLicenseRequest {
	s.CompanyName = &v
	return s
}

func (s *VerifyBusinessLicenseRequest) SetCreditCode(v string) *VerifyBusinessLicenseRequest {
	s.CreditCode = &v
	return s
}

func (s *VerifyBusinessLicenseRequest) SetLegalPerson(v string) *VerifyBusinessLicenseRequest {
	s.LegalPerson = &v
	return s
}

type VerifyBusinessLicenseResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyBusinessLicenseResponseBody) SetData(v string) *VerifyBusinessLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyBusinessLicenseResponseBody) SetRequestId(v string) *VerifyBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

type VerifyBusinessLicenseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyBusinessLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyBusinessLicenseResponse) GoString() string {
	return s.String()
}

func (s *VerifyBusinessLicenseResponse) SetHeaders(v map[string]*string) *VerifyBusinessLicenseResponse {
	s.Headers = v
	return s
}

func (s *VerifyBusinessLicenseResponse) SetStatusCode(v int32) *VerifyBusinessLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyBusinessLicenseResponse) SetBody(v *VerifyBusinessLicenseResponseBody) *VerifyBusinessLicenseResponse {
	s.Body = v
	return s
}

type VerifyVATInvoiceRequest struct {
	InvoiceCode *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceDate *string `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	InvoiceNo   *string `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	InvoiceSum  *string `json:"InvoiceSum,omitempty" xml:"InvoiceSum,omitempty"`
	VerifyCode  *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s VerifyVATInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyVATInvoiceRequest) GoString() string {
	return s.String()
}

func (s *VerifyVATInvoiceRequest) SetInvoiceCode(v string) *VerifyVATInvoiceRequest {
	s.InvoiceCode = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetInvoiceDate(v string) *VerifyVATInvoiceRequest {
	s.InvoiceDate = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetInvoiceNo(v string) *VerifyVATInvoiceRequest {
	s.InvoiceNo = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetInvoiceSum(v string) *VerifyVATInvoiceRequest {
	s.InvoiceSum = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetVerifyCode(v string) *VerifyVATInvoiceRequest {
	s.VerifyCode = &v
	return s
}

type VerifyVATInvoiceResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyVATInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyVATInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyVATInvoiceResponseBody) SetData(v string) *VerifyVATInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyVATInvoiceResponseBody) SetRequestId(v string) *VerifyVATInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type VerifyVATInvoiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyVATInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyVATInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyVATInvoiceResponse) GoString() string {
	return s.String()
}

func (s *VerifyVATInvoiceResponse) SetHeaders(v map[string]*string) *VerifyVATInvoiceResponse {
	s.Headers = v
	return s
}

func (s *VerifyVATInvoiceResponse) SetStatusCode(v int32) *VerifyVATInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyVATInvoiceResponse) SetBody(v *VerifyVATInvoiceResponseBody) *VerifyVATInvoiceResponse {
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

func (client *Client) RecognizeBankAcceptanceWithOptions(request *RecognizeBankAcceptanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankAcceptanceResponse, _err error) {
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
		Action:      tea.String("RecognizeBankAcceptance"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBankAcceptanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBankAcceptance(request *RecognizeBankAcceptanceRequest) (_result *RecognizeBankAcceptanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBankAcceptanceResponse{}
	_body, _err := client.RecognizeBankAcceptanceWithOptions(request, runtime)
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

func (client *Client) RecognizeCosmeticProduceLicenseWithOptions(request *RecognizeCosmeticProduceLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeCosmeticProduceLicenseResponse, _err error) {
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
		Action:      tea.String("RecognizeCosmeticProduceLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCosmeticProduceLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeCosmeticProduceLicense(request *RecognizeCosmeticProduceLicenseRequest) (_result *RecognizeCosmeticProduceLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCosmeticProduceLicenseResponse{}
	_body, _err := client.RecognizeCosmeticProduceLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeCovidTestReportWithOptions(request *RecognizeCovidTestReportRequest, runtime *util.RuntimeOptions) (_result *RecognizeCovidTestReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MultipleResult)) {
		query["MultipleResult"] = request.MultipleResult
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
		Action:      tea.String("RecognizeCovidTestReport"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCovidTestReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeCovidTestReport(request *RecognizeCovidTestReportRequest) (_result *RecognizeCovidTestReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCovidTestReportResponse{}
	_body, _err := client.RecognizeCovidTestReportWithOptions(request, runtime)
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

func (client *Client) RecognizeDocumentStructureWithOptions(request *RecognizeDocumentStructureRequest, runtime *util.RuntimeOptions) (_result *RecognizeDocumentStructureResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
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

	if !tea.BoolValue(util.IsUnset(request.UseNewStyleOutput)) {
		query["UseNewStyleOutput"] = request.UseNewStyleOutput
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeDocumentStructure"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeDocumentStructureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeDocumentStructure(request *RecognizeDocumentStructureRequest) (_result *RecognizeDocumentStructureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeDocumentStructureResponse{}
	_body, _err := client.RecognizeDocumentStructureWithOptions(request, runtime)
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

func (client *Client) RecognizeHealthCodeWithOptions(request *RecognizeHealthCodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeHealthCodeResponse, _err error) {
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
		Action:      tea.String("RecognizeHealthCode"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHealthCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeHealthCode(request *RecognizeHealthCodeRequest) (_result *RecognizeHealthCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHealthCodeResponse{}
	_body, _err := client.RecognizeHealthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeHotelConsumeWithOptions(request *RecognizeHotelConsumeRequest, runtime *util.RuntimeOptions) (_result *RecognizeHotelConsumeResponse, _err error) {
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
		Action:      tea.String("RecognizeHotelConsume"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHotelConsumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeHotelConsume(request *RecognizeHotelConsumeRequest) (_result *RecognizeHotelConsumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHotelConsumeResponse{}
	_body, _err := client.RecognizeHotelConsumeWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.IsResidentPage)) {
		query["IsResidentPage"] = request.IsResidentPage
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

	if !tea.BoolValue(util.IsUnset(request.OutputQualityInfo)) {
		query["OutputQualityInfo"] = request.OutputQualityInfo
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

func (client *Client) RecognizeInternationalBusinessLicenseWithOptions(request *RecognizeInternationalBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeInternationalBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
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
		Action:      tea.String("RecognizeInternationalBusinessLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeInternationalBusinessLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeInternationalBusinessLicense(request *RecognizeInternationalBusinessLicenseRequest) (_result *RecognizeInternationalBusinessLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeInternationalBusinessLicenseResponse{}
	_body, _err := client.RecognizeInternationalBusinessLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeInternationalIdcardWithOptions(request *RecognizeInternationalIdcardRequest, runtime *util.RuntimeOptions) (_result *RecognizeInternationalIdcardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
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
		Action:      tea.String("RecognizeInternationalIdcard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeInternationalIdcardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeInternationalIdcard(request *RecognizeInternationalIdcardRequest) (_result *RecognizeInternationalIdcardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeInternationalIdcardResponse{}
	_body, _err := client.RecognizeInternationalIdcardWithOptions(request, runtime)
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
		Stream: tmpReq.Body,
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

func (client *Client) RecognizeNonTaxInvoiceWithOptions(request *RecognizeNonTaxInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeNonTaxInvoiceResponse, _err error) {
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
		Action:      tea.String("RecognizeNonTaxInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeNonTaxInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeNonTaxInvoice(request *RecognizeNonTaxInvoiceRequest) (_result *RecognizeNonTaxInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeNonTaxInvoiceResponse{}
	_body, _err := client.RecognizeNonTaxInvoiceWithOptions(request, runtime)
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

func (client *Client) RecognizePaymentRecordWithOptions(request *RecognizePaymentRecordRequest, runtime *util.RuntimeOptions) (_result *RecognizePaymentRecordResponse, _err error) {
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
		Action:      tea.String("RecognizePaymentRecord"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePaymentRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizePaymentRecord(request *RecognizePaymentRecordRequest) (_result *RecognizePaymentRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePaymentRecordResponse{}
	_body, _err := client.RecognizePaymentRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizePurchaseRecordWithOptions(request *RecognizePurchaseRecordRequest, runtime *util.RuntimeOptions) (_result *RecognizePurchaseRecordResponse, _err error) {
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
		Action:      tea.String("RecognizePurchaseRecord"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePurchaseRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizePurchaseRecord(request *RecognizePurchaseRecordRequest) (_result *RecognizePurchaseRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePurchaseRecordResponse{}
	_body, _err := client.RecognizePurchaseRecordWithOptions(request, runtime)
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

func (client *Client) RecognizeShoppingReceiptWithOptions(request *RecognizeShoppingReceiptRequest, runtime *util.RuntimeOptions) (_result *RecognizeShoppingReceiptResponse, _err error) {
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
		Action:      tea.String("RecognizeShoppingReceipt"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeShoppingReceiptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeShoppingReceipt(request *RecognizeShoppingReceiptRequest) (_result *RecognizeShoppingReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeShoppingReceiptResponse{}
	_body, _err := client.RecognizeShoppingReceiptWithOptions(request, runtime)
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

func (client *Client) RecognizeSocialSecurityCardVersionIIWithOptions(request *RecognizeSocialSecurityCardVersionIIRequest, runtime *util.RuntimeOptions) (_result *RecognizeSocialSecurityCardVersionIIResponse, _err error) {
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
		Action:      tea.String("RecognizeSocialSecurityCardVersionII"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeSocialSecurityCardVersionIIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeSocialSecurityCardVersionII(request *RecognizeSocialSecurityCardVersionIIRequest) (_result *RecognizeSocialSecurityCardVersionIIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeSocialSecurityCardVersionIIResponse{}
	_body, _err := client.RecognizeSocialSecurityCardVersionIIWithOptions(request, runtime)
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

func (client *Client) RecognizeTravelCardWithOptions(request *RecognizeTravelCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeTravelCardResponse, _err error) {
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
		Action:      tea.String("RecognizeTravelCard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTravelCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTravelCard(request *RecognizeTravelCardRequest) (_result *RecognizeTravelCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTravelCardResponse{}
	_body, _err := client.RecognizeTravelCardWithOptions(request, runtime)
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

func (client *Client) VerifyBusinessLicenseWithOptions(request *VerifyBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *VerifyBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyName)) {
		query["CompanyName"] = request.CompanyName
	}

	if !tea.BoolValue(util.IsUnset(request.CreditCode)) {
		query["CreditCode"] = request.CreditCode
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPerson)) {
		query["LegalPerson"] = request.LegalPerson
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyBusinessLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyBusinessLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyBusinessLicense(request *VerifyBusinessLicenseRequest) (_result *VerifyBusinessLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyBusinessLicenseResponse{}
	_body, _err := client.VerifyBusinessLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyVATInvoiceWithOptions(request *VerifyVATInvoiceRequest, runtime *util.RuntimeOptions) (_result *VerifyVATInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InvoiceCode)) {
		query["InvoiceCode"] = request.InvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceDate)) {
		query["InvoiceDate"] = request.InvoiceDate
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceNo)) {
		query["InvoiceNo"] = request.InvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceSum)) {
		query["InvoiceSum"] = request.InvoiceSum
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyVATInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyVATInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyVATInvoice(request *VerifyVATInvoiceRequest) (_result *VerifyVATInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyVATInvoiceResponse{}
	_body, _err := client.VerifyVATInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
