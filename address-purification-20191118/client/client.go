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

type GetAddressDivisionCodeRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetAddressDivisionCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressDivisionCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAddressDivisionCodeRequest) SetServiceCode(v string) *GetAddressDivisionCodeRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetText(v string) *GetAddressDivisionCodeRequest {
	s.Text = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetDefaultProvince(v string) *GetAddressDivisionCodeRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetDefaultCity(v string) *GetAddressDivisionCodeRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetDefaultDistrict(v string) *GetAddressDivisionCodeRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetAppKey(v string) *GetAddressDivisionCodeRequest {
	s.AppKey = &v
	return s
}

type GetAddressDivisionCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAddressDivisionCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressDivisionCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressDivisionCodeResponseBody) SetRequestId(v string) *GetAddressDivisionCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddressDivisionCodeResponseBody) SetData(v string) *GetAddressDivisionCodeResponseBody {
	s.Data = &v
	return s
}

type GetAddressDivisionCodeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAddressDivisionCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAddressDivisionCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressDivisionCodeResponse) GoString() string {
	return s.String()
}

func (s *GetAddressDivisionCodeResponse) SetHeaders(v map[string]*string) *GetAddressDivisionCodeResponse {
	s.Headers = v
	return s
}

func (s *GetAddressDivisionCodeResponse) SetBody(v *GetAddressDivisionCodeResponseBody) *GetAddressDivisionCodeResponse {
	s.Body = v
	return s
}

type StructureAddressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s StructureAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s StructureAddressRequest) GoString() string {
	return s.String()
}

func (s *StructureAddressRequest) SetServiceCode(v string) *StructureAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *StructureAddressRequest) SetText(v string) *StructureAddressRequest {
	s.Text = &v
	return s
}

func (s *StructureAddressRequest) SetDefaultProvince(v string) *StructureAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *StructureAddressRequest) SetDefaultCity(v string) *StructureAddressRequest {
	s.DefaultCity = &v
	return s
}

func (s *StructureAddressRequest) SetDefaultDistrict(v string) *StructureAddressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *StructureAddressRequest) SetAppKey(v string) *StructureAddressRequest {
	s.AppKey = &v
	return s
}

type StructureAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s StructureAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StructureAddressResponseBody) GoString() string {
	return s.String()
}

func (s *StructureAddressResponseBody) SetRequestId(v string) *StructureAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *StructureAddressResponseBody) SetData(v string) *StructureAddressResponseBody {
	s.Data = &v
	return s
}

type StructureAddressResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StructureAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StructureAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s StructureAddressResponse) GoString() string {
	return s.String()
}

func (s *StructureAddressResponse) SetHeaders(v map[string]*string) *StructureAddressResponse {
	s.Headers = v
	return s
}

func (s *StructureAddressResponse) SetBody(v *StructureAddressResponseBody) *StructureAddressResponse {
	s.Body = v
	return s
}

type ExtractExpressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ExtractExpressRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractExpressRequest) GoString() string {
	return s.String()
}

func (s *ExtractExpressRequest) SetServiceCode(v string) *ExtractExpressRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractExpressRequest) SetText(v string) *ExtractExpressRequest {
	s.Text = &v
	return s
}

func (s *ExtractExpressRequest) SetDefaultProvince(v string) *ExtractExpressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractExpressRequest) SetDefaultCity(v string) *ExtractExpressRequest {
	s.DefaultCity = &v
	return s
}

func (s *ExtractExpressRequest) SetDefaultDistrict(v string) *ExtractExpressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ExtractExpressRequest) SetAppKey(v string) *ExtractExpressRequest {
	s.AppKey = &v
	return s
}

type ExtractExpressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ExtractExpressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractExpressResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractExpressResponseBody) SetRequestId(v string) *ExtractExpressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtractExpressResponseBody) SetData(v string) *ExtractExpressResponseBody {
	s.Data = &v
	return s
}

type ExtractExpressResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtractExpressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractExpressResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractExpressResponse) GoString() string {
	return s.String()
}

func (s *ExtractExpressResponse) SetHeaders(v map[string]*string) *ExtractExpressResponse {
	s.Headers = v
	return s
}

func (s *ExtractExpressResponse) SetBody(v *ExtractExpressResponseBody) *ExtractExpressResponse {
	s.Body = v
	return s
}

type ExtractNameRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ExtractNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractNameRequest) GoString() string {
	return s.String()
}

func (s *ExtractNameRequest) SetServiceCode(v string) *ExtractNameRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractNameRequest) SetText(v string) *ExtractNameRequest {
	s.Text = &v
	return s
}

func (s *ExtractNameRequest) SetDefaultProvince(v string) *ExtractNameRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractNameRequest) SetDefaultCity(v string) *ExtractNameRequest {
	s.DefaultCity = &v
	return s
}

func (s *ExtractNameRequest) SetDefaultDistrict(v string) *ExtractNameRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ExtractNameRequest) SetAppKey(v string) *ExtractNameRequest {
	s.AppKey = &v
	return s
}

type ExtractNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ExtractNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractNameResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractNameResponseBody) SetRequestId(v string) *ExtractNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtractNameResponseBody) SetData(v string) *ExtractNameResponseBody {
	s.Data = &v
	return s
}

type ExtractNameResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtractNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractNameResponse) GoString() string {
	return s.String()
}

func (s *ExtractNameResponse) SetHeaders(v map[string]*string) *ExtractNameResponse {
	s.Headers = v
	return s
}

func (s *ExtractNameResponse) SetBody(v *ExtractNameResponseBody) *ExtractNameResponse {
	s.Body = v
	return s
}

type GetAddressBlockMappingRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetAddressBlockMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressBlockMappingRequest) GoString() string {
	return s.String()
}

func (s *GetAddressBlockMappingRequest) SetServiceCode(v string) *GetAddressBlockMappingRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressBlockMappingRequest) SetText(v string) *GetAddressBlockMappingRequest {
	s.Text = &v
	return s
}

func (s *GetAddressBlockMappingRequest) SetDefaultProvince(v string) *GetAddressBlockMappingRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressBlockMappingRequest) SetDefaultCity(v string) *GetAddressBlockMappingRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressBlockMappingRequest) SetDefaultDistrict(v string) *GetAddressBlockMappingRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressBlockMappingRequest) SetAppKey(v string) *GetAddressBlockMappingRequest {
	s.AppKey = &v
	return s
}

type GetAddressBlockMappingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAddressBlockMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressBlockMappingResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressBlockMappingResponseBody) SetRequestId(v string) *GetAddressBlockMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddressBlockMappingResponseBody) SetData(v string) *GetAddressBlockMappingResponseBody {
	s.Data = &v
	return s
}

type GetAddressBlockMappingResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAddressBlockMappingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAddressBlockMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressBlockMappingResponse) GoString() string {
	return s.String()
}

func (s *GetAddressBlockMappingResponse) SetHeaders(v map[string]*string) *GetAddressBlockMappingResponse {
	s.Headers = v
	return s
}

func (s *GetAddressBlockMappingResponse) SetBody(v *GetAddressBlockMappingResponseBody) *GetAddressBlockMappingResponse {
	s.Body = v
	return s
}

type GetAddressSearchRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetAddressSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSearchRequest) GoString() string {
	return s.String()
}

func (s *GetAddressSearchRequest) SetServiceCode(v string) *GetAddressSearchRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressSearchRequest) SetText(v string) *GetAddressSearchRequest {
	s.Text = &v
	return s
}

func (s *GetAddressSearchRequest) SetDefaultProvince(v string) *GetAddressSearchRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressSearchRequest) SetDefaultCity(v string) *GetAddressSearchRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressSearchRequest) SetDefaultDistrict(v string) *GetAddressSearchRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressSearchRequest) SetAppKey(v string) *GetAddressSearchRequest {
	s.AppKey = &v
	return s
}

type GetAddressSearchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAddressSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSearchResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressSearchResponseBody) SetRequestId(v string) *GetAddressSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddressSearchResponseBody) SetData(v string) *GetAddressSearchResponseBody {
	s.Data = &v
	return s
}

type GetAddressSearchResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAddressSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAddressSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSearchResponse) GoString() string {
	return s.String()
}

func (s *GetAddressSearchResponse) SetHeaders(v map[string]*string) *GetAddressSearchResponse {
	s.Headers = v
	return s
}

func (s *GetAddressSearchResponse) SetBody(v *GetAddressSearchResponseBody) *GetAddressSearchResponse {
	s.Body = v
	return s
}

type PredictPOIRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s PredictPOIRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictPOIRequest) GoString() string {
	return s.String()
}

func (s *PredictPOIRequest) SetServiceCode(v string) *PredictPOIRequest {
	s.ServiceCode = &v
	return s
}

func (s *PredictPOIRequest) SetText(v string) *PredictPOIRequest {
	s.Text = &v
	return s
}

func (s *PredictPOIRequest) SetDefaultProvince(v string) *PredictPOIRequest {
	s.DefaultProvince = &v
	return s
}

func (s *PredictPOIRequest) SetDefaultCity(v string) *PredictPOIRequest {
	s.DefaultCity = &v
	return s
}

func (s *PredictPOIRequest) SetDefaultDistrict(v string) *PredictPOIRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *PredictPOIRequest) SetAppKey(v string) *PredictPOIRequest {
	s.AppKey = &v
	return s
}

type PredictPOIResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s PredictPOIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictPOIResponseBody) GoString() string {
	return s.String()
}

func (s *PredictPOIResponseBody) SetRequestId(v string) *PredictPOIResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictPOIResponseBody) SetData(v string) *PredictPOIResponseBody {
	s.Data = &v
	return s
}

type PredictPOIResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PredictPOIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PredictPOIResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictPOIResponse) GoString() string {
	return s.String()
}

func (s *PredictPOIResponse) SetHeaders(v map[string]*string) *PredictPOIResponse {
	s.Headers = v
	return s
}

func (s *PredictPOIResponse) SetBody(v *PredictPOIResponseBody) *PredictPOIResponse {
	s.Body = v
	return s
}

type ClassifyPOIRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ClassifyPOIRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyPOIRequest) GoString() string {
	return s.String()
}

func (s *ClassifyPOIRequest) SetServiceCode(v string) *ClassifyPOIRequest {
	s.ServiceCode = &v
	return s
}

func (s *ClassifyPOIRequest) SetText(v string) *ClassifyPOIRequest {
	s.Text = &v
	return s
}

func (s *ClassifyPOIRequest) SetDefaultProvince(v string) *ClassifyPOIRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ClassifyPOIRequest) SetDefaultCity(v string) *ClassifyPOIRequest {
	s.DefaultCity = &v
	return s
}

func (s *ClassifyPOIRequest) SetDefaultDistrict(v string) *ClassifyPOIRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ClassifyPOIRequest) SetAppKey(v string) *ClassifyPOIRequest {
	s.AppKey = &v
	return s
}

type ClassifyPOIResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ClassifyPOIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClassifyPOIResponseBody) GoString() string {
	return s.String()
}

func (s *ClassifyPOIResponseBody) SetRequestId(v string) *ClassifyPOIResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClassifyPOIResponseBody) SetData(v string) *ClassifyPOIResponseBody {
	s.Data = &v
	return s
}

type ClassifyPOIResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClassifyPOIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClassifyPOIResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyPOIResponse) GoString() string {
	return s.String()
}

func (s *ClassifyPOIResponse) SetHeaders(v map[string]*string) *ClassifyPOIResponse {
	s.Headers = v
	return s
}

func (s *ClassifyPOIResponse) SetBody(v *ClassifyPOIResponseBody) *ClassifyPOIResponse {
	s.Body = v
	return s
}

type CorrectAddressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s CorrectAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CorrectAddressRequest) GoString() string {
	return s.String()
}

func (s *CorrectAddressRequest) SetServiceCode(v string) *CorrectAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *CorrectAddressRequest) SetText(v string) *CorrectAddressRequest {
	s.Text = &v
	return s
}

func (s *CorrectAddressRequest) SetDefaultProvince(v string) *CorrectAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *CorrectAddressRequest) SetDefaultCity(v string) *CorrectAddressRequest {
	s.DefaultCity = &v
	return s
}

func (s *CorrectAddressRequest) SetDefaultDistrict(v string) *CorrectAddressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *CorrectAddressRequest) SetAppKey(v string) *CorrectAddressRequest {
	s.AppKey = &v
	return s
}

type CorrectAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CorrectAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CorrectAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CorrectAddressResponseBody) SetRequestId(v string) *CorrectAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CorrectAddressResponseBody) SetData(v string) *CorrectAddressResponseBody {
	s.Data = &v
	return s
}

type CorrectAddressResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CorrectAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CorrectAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CorrectAddressResponse) GoString() string {
	return s.String()
}

func (s *CorrectAddressResponse) SetHeaders(v map[string]*string) *CorrectAddressResponse {
	s.Headers = v
	return s
}

func (s *CorrectAddressResponse) SetBody(v *CorrectAddressResponseBody) *CorrectAddressResponse {
	s.Body = v
	return s
}

type GetZipcodeRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetZipcodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetZipcodeRequest) GoString() string {
	return s.String()
}

func (s *GetZipcodeRequest) SetServiceCode(v string) *GetZipcodeRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetZipcodeRequest) SetText(v string) *GetZipcodeRequest {
	s.Text = &v
	return s
}

func (s *GetZipcodeRequest) SetDefaultProvince(v string) *GetZipcodeRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetZipcodeRequest) SetDefaultCity(v string) *GetZipcodeRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetZipcodeRequest) SetDefaultDistrict(v string) *GetZipcodeRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetZipcodeRequest) SetAppKey(v string) *GetZipcodeRequest {
	s.AppKey = &v
	return s
}

type GetZipcodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetZipcodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetZipcodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetZipcodeResponseBody) SetRequestId(v string) *GetZipcodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetZipcodeResponseBody) SetData(v string) *GetZipcodeResponseBody {
	s.Data = &v
	return s
}

type GetZipcodeResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetZipcodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetZipcodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetZipcodeResponse) GoString() string {
	return s.String()
}

func (s *GetZipcodeResponse) SetHeaders(v map[string]*string) *GetZipcodeResponse {
	s.Headers = v
	return s
}

func (s *GetZipcodeResponse) SetBody(v *GetZipcodeResponseBody) *GetZipcodeResponse {
	s.Body = v
	return s
}

type CompleteAddressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s CompleteAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteAddressRequest) GoString() string {
	return s.String()
}

func (s *CompleteAddressRequest) SetServiceCode(v string) *CompleteAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *CompleteAddressRequest) SetText(v string) *CompleteAddressRequest {
	s.Text = &v
	return s
}

func (s *CompleteAddressRequest) SetDefaultProvince(v string) *CompleteAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *CompleteAddressRequest) SetDefaultCity(v string) *CompleteAddressRequest {
	s.DefaultCity = &v
	return s
}

func (s *CompleteAddressRequest) SetDefaultDistrict(v string) *CompleteAddressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *CompleteAddressRequest) SetAppKey(v string) *CompleteAddressRequest {
	s.AppKey = &v
	return s
}

type CompleteAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CompleteAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompleteAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteAddressResponseBody) SetRequestId(v string) *CompleteAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteAddressResponseBody) SetData(v string) *CompleteAddressResponseBody {
	s.Data = &v
	return s
}

type CompleteAddressResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CompleteAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompleteAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteAddressResponse) GoString() string {
	return s.String()
}

func (s *CompleteAddressResponse) SetHeaders(v map[string]*string) *CompleteAddressResponse {
	s.Headers = v
	return s
}

func (s *CompleteAddressResponse) SetBody(v *CompleteAddressResponseBody) *CompleteAddressResponse {
	s.Body = v
	return s
}

type GetAddressSimilarityRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetAddressSimilarityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSimilarityRequest) GoString() string {
	return s.String()
}

func (s *GetAddressSimilarityRequest) SetServiceCode(v string) *GetAddressSimilarityRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetText(v string) *GetAddressSimilarityRequest {
	s.Text = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetDefaultProvince(v string) *GetAddressSimilarityRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetDefaultCity(v string) *GetAddressSimilarityRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetDefaultDistrict(v string) *GetAddressSimilarityRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetAppKey(v string) *GetAddressSimilarityRequest {
	s.AppKey = &v
	return s
}

type GetAddressSimilarityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAddressSimilarityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSimilarityResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressSimilarityResponseBody) SetRequestId(v string) *GetAddressSimilarityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddressSimilarityResponseBody) SetData(v string) *GetAddressSimilarityResponseBody {
	s.Data = &v
	return s
}

type GetAddressSimilarityResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAddressSimilarityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAddressSimilarityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSimilarityResponse) GoString() string {
	return s.String()
}

func (s *GetAddressSimilarityResponse) SetHeaders(v map[string]*string) *GetAddressSimilarityResponse {
	s.Headers = v
	return s
}

func (s *GetAddressSimilarityResponse) SetBody(v *GetAddressSimilarityResponseBody) *GetAddressSimilarityResponse {
	s.Body = v
	return s
}

type GetAddressGeocodeRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetAddressGeocodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressGeocodeRequest) GoString() string {
	return s.String()
}

func (s *GetAddressGeocodeRequest) SetServiceCode(v string) *GetAddressGeocodeRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetText(v string) *GetAddressGeocodeRequest {
	s.Text = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetDefaultProvince(v string) *GetAddressGeocodeRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetDefaultCity(v string) *GetAddressGeocodeRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetDefaultDistrict(v string) *GetAddressGeocodeRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetAppKey(v string) *GetAddressGeocodeRequest {
	s.AppKey = &v
	return s
}

type GetAddressGeocodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAddressGeocodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressGeocodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressGeocodeResponseBody) SetRequestId(v string) *GetAddressGeocodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddressGeocodeResponseBody) SetData(v string) *GetAddressGeocodeResponseBody {
	s.Data = &v
	return s
}

type GetAddressGeocodeResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAddressGeocodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAddressGeocodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressGeocodeResponse) GoString() string {
	return s.String()
}

func (s *GetAddressGeocodeResponse) SetHeaders(v map[string]*string) *GetAddressGeocodeResponse {
	s.Headers = v
	return s
}

func (s *GetAddressGeocodeResponse) SetBody(v *GetAddressGeocodeResponseBody) *GetAddressGeocodeResponse {
	s.Body = v
	return s
}

type TransferCoordRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	SrcCoord        *string `json:"SrcCoord,omitempty" xml:"SrcCoord,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s TransferCoordRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCoordRequest) GoString() string {
	return s.String()
}

func (s *TransferCoordRequest) SetServiceCode(v string) *TransferCoordRequest {
	s.ServiceCode = &v
	return s
}

func (s *TransferCoordRequest) SetText(v string) *TransferCoordRequest {
	s.Text = &v
	return s
}

func (s *TransferCoordRequest) SetSrcCoord(v string) *TransferCoordRequest {
	s.SrcCoord = &v
	return s
}

func (s *TransferCoordRequest) SetDefaultProvince(v string) *TransferCoordRequest {
	s.DefaultProvince = &v
	return s
}

func (s *TransferCoordRequest) SetDefaultCity(v string) *TransferCoordRequest {
	s.DefaultCity = &v
	return s
}

func (s *TransferCoordRequest) SetDefaultDistrict(v string) *TransferCoordRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *TransferCoordRequest) SetAppKey(v string) *TransferCoordRequest {
	s.AppKey = &v
	return s
}

type TransferCoordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s TransferCoordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferCoordResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCoordResponseBody) SetRequestId(v string) *TransferCoordResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCoordResponseBody) SetData(v string) *TransferCoordResponseBody {
	s.Data = &v
	return s
}

type TransferCoordResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferCoordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferCoordResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCoordResponse) GoString() string {
	return s.String()
}

func (s *TransferCoordResponse) SetHeaders(v map[string]*string) *TransferCoordResponse {
	s.Headers = v
	return s
}

func (s *TransferCoordResponse) SetBody(v *TransferCoordResponseBody) *TransferCoordResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	ServiceCode *string                `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Parameters  map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetServiceCode(v string) *UpdateProjectRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateProjectRequest) SetParameters(v map[string]interface{}) *UpdateProjectRequest {
	s.Parameters = v
	return s
}

type UpdateProjectShrinkRequest struct {
	ServiceCode      *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s UpdateProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectShrinkRequest) SetServiceCode(v string) *UpdateProjectShrinkRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetParametersShrink(v string) *UpdateProjectShrinkRequest {
	s.ParametersShrink = &v
	return s
}

type UpdateProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetData(v string) *UpdateProjectResponseBody {
	s.Data = &v
	return s
}

type UpdateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type ExtractPhoneRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ExtractPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractPhoneRequest) GoString() string {
	return s.String()
}

func (s *ExtractPhoneRequest) SetServiceCode(v string) *ExtractPhoneRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractPhoneRequest) SetText(v string) *ExtractPhoneRequest {
	s.Text = &v
	return s
}

func (s *ExtractPhoneRequest) SetDefaultProvince(v string) *ExtractPhoneRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractPhoneRequest) SetDefaultCity(v string) *ExtractPhoneRequest {
	s.DefaultCity = &v
	return s
}

func (s *ExtractPhoneRequest) SetDefaultDistrict(v string) *ExtractPhoneRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ExtractPhoneRequest) SetAppKey(v string) *ExtractPhoneRequest {
	s.AppKey = &v
	return s
}

type ExtractPhoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ExtractPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractPhoneResponseBody) SetRequestId(v string) *ExtractPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtractPhoneResponseBody) SetData(v string) *ExtractPhoneResponseBody {
	s.Data = &v
	return s
}

type ExtractPhoneResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtractPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractPhoneResponse) GoString() string {
	return s.String()
}

func (s *ExtractPhoneResponse) SetHeaders(v map[string]*string) *ExtractPhoneResponse {
	s.Headers = v
	return s
}

func (s *ExtractPhoneResponse) SetBody(v *ExtractPhoneResponseBody) *ExtractPhoneResponse {
	s.Body = v
	return s
}

type GetInputSearchRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetInputSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInputSearchRequest) GoString() string {
	return s.String()
}

func (s *GetInputSearchRequest) SetServiceCode(v string) *GetInputSearchRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetInputSearchRequest) SetText(v string) *GetInputSearchRequest {
	s.Text = &v
	return s
}

func (s *GetInputSearchRequest) SetDefaultProvince(v string) *GetInputSearchRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetInputSearchRequest) SetDefaultCity(v string) *GetInputSearchRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetInputSearchRequest) SetDefaultDistrict(v string) *GetInputSearchRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetInputSearchRequest) SetAppKey(v string) *GetInputSearchRequest {
	s.AppKey = &v
	return s
}

type GetInputSearchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetInputSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInputSearchResponseBody) GoString() string {
	return s.String()
}

func (s *GetInputSearchResponseBody) SetRequestId(v string) *GetInputSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInputSearchResponseBody) SetData(v string) *GetInputSearchResponseBody {
	s.Data = &v
	return s
}

type GetInputSearchResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInputSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInputSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInputSearchResponse) GoString() string {
	return s.String()
}

func (s *GetInputSearchResponse) SetHeaders(v map[string]*string) *GetInputSearchResponse {
	s.Headers = v
	return s
}

func (s *GetInputSearchResponse) SetBody(v *GetInputSearchResponseBody) *GetInputSearchResponse {
	s.Body = v
	return s
}

type GetAddressEvaluateRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetAddressEvaluateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressEvaluateRequest) GoString() string {
	return s.String()
}

func (s *GetAddressEvaluateRequest) SetServiceCode(v string) *GetAddressEvaluateRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressEvaluateRequest) SetText(v string) *GetAddressEvaluateRequest {
	s.Text = &v
	return s
}

func (s *GetAddressEvaluateRequest) SetDefaultProvince(v string) *GetAddressEvaluateRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressEvaluateRequest) SetDefaultCity(v string) *GetAddressEvaluateRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressEvaluateRequest) SetDefaultDistrict(v string) *GetAddressEvaluateRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressEvaluateRequest) SetAppKey(v string) *GetAddressEvaluateRequest {
	s.AppKey = &v
	return s
}

type GetAddressEvaluateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAddressEvaluateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressEvaluateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressEvaluateResponseBody) SetRequestId(v string) *GetAddressEvaluateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddressEvaluateResponseBody) SetData(v string) *GetAddressEvaluateResponseBody {
	s.Data = &v
	return s
}

type GetAddressEvaluateResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAddressEvaluateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAddressEvaluateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressEvaluateResponse) GoString() string {
	return s.String()
}

func (s *GetAddressEvaluateResponse) SetHeaders(v map[string]*string) *GetAddressEvaluateResponse {
	s.Headers = v
	return s
}

func (s *GetAddressEvaluateResponse) SetBody(v *GetAddressEvaluateResponseBody) *GetAddressEvaluateResponse {
	s.Body = v
	return s
}

type ExtractAddressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ExtractAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractAddressRequest) GoString() string {
	return s.String()
}

func (s *ExtractAddressRequest) SetServiceCode(v string) *ExtractAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractAddressRequest) SetText(v string) *ExtractAddressRequest {
	s.Text = &v
	return s
}

func (s *ExtractAddressRequest) SetDefaultProvince(v string) *ExtractAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractAddressRequest) SetDefaultCity(v string) *ExtractAddressRequest {
	s.DefaultCity = &v
	return s
}

func (s *ExtractAddressRequest) SetDefaultDistrict(v string) *ExtractAddressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ExtractAddressRequest) SetAppKey(v string) *ExtractAddressRequest {
	s.AppKey = &v
	return s
}

type ExtractAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ExtractAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractAddressResponseBody) SetRequestId(v string) *ExtractAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtractAddressResponseBody) SetData(v string) *ExtractAddressResponseBody {
	s.Data = &v
	return s
}

type ExtractAddressResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtractAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractAddressResponse) GoString() string {
	return s.String()
}

func (s *ExtractAddressResponse) SetHeaders(v map[string]*string) *ExtractAddressResponse {
	s.Headers = v
	return s
}

func (s *ExtractAddressResponse) SetBody(v *ExtractAddressResponseBody) *ExtractAddressResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("address-purification"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetAddressDivisionCodeWithOptions(request *GetAddressDivisionCodeRequest, runtime *util.RuntimeOptions) (_result *GetAddressDivisionCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAddressDivisionCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAddressDivisionCode"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressDivisionCode(request *GetAddressDivisionCodeRequest) (_result *GetAddressDivisionCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressDivisionCodeResponse{}
	_body, _err := client.GetAddressDivisionCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StructureAddressWithOptions(request *StructureAddressRequest, runtime *util.RuntimeOptions) (_result *StructureAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StructureAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StructureAddress"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StructureAddress(request *StructureAddressRequest) (_result *StructureAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StructureAddressResponse{}
	_body, _err := client.StructureAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractExpressWithOptions(request *ExtractExpressRequest, runtime *util.RuntimeOptions) (_result *ExtractExpressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExtractExpressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExtractExpress"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractExpress(request *ExtractExpressRequest) (_result *ExtractExpressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractExpressResponse{}
	_body, _err := client.ExtractExpressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractNameWithOptions(request *ExtractNameRequest, runtime *util.RuntimeOptions) (_result *ExtractNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExtractNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExtractName"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractName(request *ExtractNameRequest) (_result *ExtractNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractNameResponse{}
	_body, _err := client.ExtractNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAddressBlockMappingWithOptions(request *GetAddressBlockMappingRequest, runtime *util.RuntimeOptions) (_result *GetAddressBlockMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAddressBlockMappingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAddressBlockMapping"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressBlockMapping(request *GetAddressBlockMappingRequest) (_result *GetAddressBlockMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressBlockMappingResponse{}
	_body, _err := client.GetAddressBlockMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAddressSearchWithOptions(request *GetAddressSearchRequest, runtime *util.RuntimeOptions) (_result *GetAddressSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAddressSearchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAddressSearch"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressSearch(request *GetAddressSearchRequest) (_result *GetAddressSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressSearchResponse{}
	_body, _err := client.GetAddressSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PredictPOIWithOptions(request *PredictPOIRequest, runtime *util.RuntimeOptions) (_result *PredictPOIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PredictPOIResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PredictPOI"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PredictPOI(request *PredictPOIRequest) (_result *PredictPOIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictPOIResponse{}
	_body, _err := client.PredictPOIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyPOIWithOptions(request *ClassifyPOIRequest, runtime *util.RuntimeOptions) (_result *ClassifyPOIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClassifyPOIResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClassifyPOI"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyPOI(request *ClassifyPOIRequest) (_result *ClassifyPOIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClassifyPOIResponse{}
	_body, _err := client.ClassifyPOIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CorrectAddressWithOptions(request *CorrectAddressRequest, runtime *util.RuntimeOptions) (_result *CorrectAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CorrectAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CorrectAddress"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CorrectAddress(request *CorrectAddressRequest) (_result *CorrectAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CorrectAddressResponse{}
	_body, _err := client.CorrectAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetZipcodeWithOptions(request *GetZipcodeRequest, runtime *util.RuntimeOptions) (_result *GetZipcodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetZipcodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetZipcode"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetZipcode(request *GetZipcodeRequest) (_result *GetZipcodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetZipcodeResponse{}
	_body, _err := client.GetZipcodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteAddressWithOptions(request *CompleteAddressRequest, runtime *util.RuntimeOptions) (_result *CompleteAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CompleteAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CompleteAddress"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteAddress(request *CompleteAddressRequest) (_result *CompleteAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompleteAddressResponse{}
	_body, _err := client.CompleteAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAddressSimilarityWithOptions(request *GetAddressSimilarityRequest, runtime *util.RuntimeOptions) (_result *GetAddressSimilarityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAddressSimilarityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAddressSimilarity"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressSimilarity(request *GetAddressSimilarityRequest) (_result *GetAddressSimilarityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressSimilarityResponse{}
	_body, _err := client.GetAddressSimilarityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAddressGeocodeWithOptions(request *GetAddressGeocodeRequest, runtime *util.RuntimeOptions) (_result *GetAddressGeocodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAddressGeocodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAddressGeocode"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressGeocode(request *GetAddressGeocodeRequest) (_result *GetAddressGeocodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressGeocodeResponse{}
	_body, _err := client.GetAddressGeocodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferCoordWithOptions(request *TransferCoordRequest, runtime *util.RuntimeOptions) (_result *TransferCoordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferCoordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferCoord"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferCoord(request *TransferCoordRequest) (_result *TransferCoordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferCoordResponse{}
	_body, _err := client.TransferCoordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(tmpReq *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProject"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractPhoneWithOptions(request *ExtractPhoneRequest, runtime *util.RuntimeOptions) (_result *ExtractPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExtractPhoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExtractPhone"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractPhone(request *ExtractPhoneRequest) (_result *ExtractPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractPhoneResponse{}
	_body, _err := client.ExtractPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInputSearchWithOptions(request *GetInputSearchRequest, runtime *util.RuntimeOptions) (_result *GetInputSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetInputSearchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInputSearch"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInputSearch(request *GetInputSearchRequest) (_result *GetInputSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInputSearchResponse{}
	_body, _err := client.GetInputSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAddressEvaluateWithOptions(request *GetAddressEvaluateRequest, runtime *util.RuntimeOptions) (_result *GetAddressEvaluateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAddressEvaluateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAddressEvaluate"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressEvaluate(request *GetAddressEvaluateRequest) (_result *GetAddressEvaluateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressEvaluateResponse{}
	_body, _err := client.GetAddressEvaluateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractAddressWithOptions(request *ExtractAddressRequest, runtime *util.RuntimeOptions) (_result *ExtractAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExtractAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExtractAddress"), tea.String("2019-11-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractAddress(request *ExtractAddressRequest) (_result *ExtractAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractAddressResponse{}
	_body, _err := client.ExtractAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
