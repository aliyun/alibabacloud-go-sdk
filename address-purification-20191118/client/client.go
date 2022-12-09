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
)

type CorrectAddressRequest struct {
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s CorrectAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CorrectAddressRequest) GoString() string {
	return s.String()
}

func (s *CorrectAddressRequest) SetAppKey(v string) *CorrectAddressRequest {
	s.AppKey = &v
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

func (s *CorrectAddressRequest) SetDefaultProvince(v string) *CorrectAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *CorrectAddressRequest) SetServiceCode(v string) *CorrectAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *CorrectAddressRequest) SetText(v string) *CorrectAddressRequest {
	s.Text = &v
	return s
}

type CorrectAddressResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CorrectAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CorrectAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CorrectAddressResponseBody) SetData(v string) *CorrectAddressResponseBody {
	s.Data = &v
	return s
}

func (s *CorrectAddressResponseBody) SetRequestId(v string) *CorrectAddressResponseBody {
	s.RequestId = &v
	return s
}

type CorrectAddressResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CorrectAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CorrectAddressResponse) SetStatusCode(v int32) *CorrectAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CorrectAddressResponse) SetBody(v *CorrectAddressResponseBody) *CorrectAddressResponse {
	s.Body = v
	return s
}

type ExtractAddressRequest struct {
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ExtractAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractAddressRequest) GoString() string {
	return s.String()
}

func (s *ExtractAddressRequest) SetAppKey(v string) *ExtractAddressRequest {
	s.AppKey = &v
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

func (s *ExtractAddressRequest) SetDefaultProvince(v string) *ExtractAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractAddressRequest) SetServiceCode(v string) *ExtractAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractAddressRequest) SetText(v string) *ExtractAddressRequest {
	s.Text = &v
	return s
}

type ExtractAddressResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtractAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractAddressResponseBody) SetData(v string) *ExtractAddressResponseBody {
	s.Data = &v
	return s
}

func (s *ExtractAddressResponseBody) SetRequestId(v string) *ExtractAddressResponseBody {
	s.RequestId = &v
	return s
}

type ExtractAddressResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtractAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExtractAddressResponse) SetStatusCode(v int32) *ExtractAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtractAddressResponse) SetBody(v *ExtractAddressResponseBody) *ExtractAddressResponse {
	s.Body = v
	return s
}

type ExtractNameRequest struct {
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ExtractNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractNameRequest) GoString() string {
	return s.String()
}

func (s *ExtractNameRequest) SetAppKey(v string) *ExtractNameRequest {
	s.AppKey = &v
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

func (s *ExtractNameRequest) SetDefaultProvince(v string) *ExtractNameRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractNameRequest) SetServiceCode(v string) *ExtractNameRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractNameRequest) SetText(v string) *ExtractNameRequest {
	s.Text = &v
	return s
}

type ExtractNameResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtractNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractNameResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractNameResponseBody) SetData(v string) *ExtractNameResponseBody {
	s.Data = &v
	return s
}

func (s *ExtractNameResponseBody) SetRequestId(v string) *ExtractNameResponseBody {
	s.RequestId = &v
	return s
}

type ExtractNameResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtractNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExtractNameResponse) SetStatusCode(v int32) *ExtractNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtractNameResponse) SetBody(v *ExtractNameResponseBody) *ExtractNameResponse {
	s.Body = v
	return s
}

type ExtractPhoneRequest struct {
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ExtractPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractPhoneRequest) GoString() string {
	return s.String()
}

func (s *ExtractPhoneRequest) SetAppKey(v string) *ExtractPhoneRequest {
	s.AppKey = &v
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

func (s *ExtractPhoneRequest) SetDefaultProvince(v string) *ExtractPhoneRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractPhoneRequest) SetServiceCode(v string) *ExtractPhoneRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractPhoneRequest) SetText(v string) *ExtractPhoneRequest {
	s.Text = &v
	return s
}

type ExtractPhoneResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtractPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractPhoneResponseBody) SetData(v string) *ExtractPhoneResponseBody {
	s.Data = &v
	return s
}

func (s *ExtractPhoneResponseBody) SetRequestId(v string) *ExtractPhoneResponseBody {
	s.RequestId = &v
	return s
}

type ExtractPhoneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtractPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExtractPhoneResponse) SetStatusCode(v int32) *ExtractPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtractPhoneResponse) SetBody(v *ExtractPhoneResponseBody) *ExtractPhoneResponse {
	s.Body = v
	return s
}

type GetAddressDivisionCodeRequest struct {
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetAddressDivisionCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressDivisionCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAddressDivisionCodeRequest) SetAppKey(v string) *GetAddressDivisionCodeRequest {
	s.AppKey = &v
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

func (s *GetAddressDivisionCodeRequest) SetDefaultProvince(v string) *GetAddressDivisionCodeRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetServiceCode(v string) *GetAddressDivisionCodeRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetText(v string) *GetAddressDivisionCodeRequest {
	s.Text = &v
	return s
}

type GetAddressDivisionCodeResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAddressDivisionCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressDivisionCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressDivisionCodeResponseBody) SetData(v string) *GetAddressDivisionCodeResponseBody {
	s.Data = &v
	return s
}

func (s *GetAddressDivisionCodeResponseBody) SetRequestId(v string) *GetAddressDivisionCodeResponseBody {
	s.RequestId = &v
	return s
}

type GetAddressDivisionCodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAddressDivisionCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAddressDivisionCodeResponse) SetStatusCode(v int32) *GetAddressDivisionCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddressDivisionCodeResponse) SetBody(v *GetAddressDivisionCodeResponseBody) *GetAddressDivisionCodeResponse {
	s.Body = v
	return s
}

type GetAddressSimilarityRequest struct {
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetAddressSimilarityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSimilarityRequest) GoString() string {
	return s.String()
}

func (s *GetAddressSimilarityRequest) SetAppKey(v string) *GetAddressSimilarityRequest {
	s.AppKey = &v
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

func (s *GetAddressSimilarityRequest) SetDefaultProvince(v string) *GetAddressSimilarityRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetServiceCode(v string) *GetAddressSimilarityRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetText(v string) *GetAddressSimilarityRequest {
	s.Text = &v
	return s
}

type GetAddressSimilarityResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAddressSimilarityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSimilarityResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressSimilarityResponseBody) SetData(v string) *GetAddressSimilarityResponseBody {
	s.Data = &v
	return s
}

func (s *GetAddressSimilarityResponseBody) SetRequestId(v string) *GetAddressSimilarityResponseBody {
	s.RequestId = &v
	return s
}

type GetAddressSimilarityResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAddressSimilarityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAddressSimilarityResponse) SetStatusCode(v int32) *GetAddressSimilarityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddressSimilarityResponse) SetBody(v *GetAddressSimilarityResponseBody) *GetAddressSimilarityResponse {
	s.Body = v
	return s
}

type GetZipcodeRequest struct {
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetZipcodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetZipcodeRequest) GoString() string {
	return s.String()
}

func (s *GetZipcodeRequest) SetAppKey(v string) *GetZipcodeRequest {
	s.AppKey = &v
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

func (s *GetZipcodeRequest) SetDefaultProvince(v string) *GetZipcodeRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetZipcodeRequest) SetServiceCode(v string) *GetZipcodeRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetZipcodeRequest) SetText(v string) *GetZipcodeRequest {
	s.Text = &v
	return s
}

type GetZipcodeResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetZipcodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetZipcodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetZipcodeResponseBody) SetData(v string) *GetZipcodeResponseBody {
	s.Data = &v
	return s
}

func (s *GetZipcodeResponseBody) SetRequestId(v string) *GetZipcodeResponseBody {
	s.RequestId = &v
	return s
}

type GetZipcodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetZipcodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetZipcodeResponse) SetStatusCode(v int32) *GetZipcodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetZipcodeResponse) SetBody(v *GetZipcodeResponseBody) *GetZipcodeResponse {
	s.Body = v
	return s
}

type StructureAddressRequest struct {
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s StructureAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s StructureAddressRequest) GoString() string {
	return s.String()
}

func (s *StructureAddressRequest) SetAppKey(v string) *StructureAddressRequest {
	s.AppKey = &v
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

func (s *StructureAddressRequest) SetDefaultProvince(v string) *StructureAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *StructureAddressRequest) SetServiceCode(v string) *StructureAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *StructureAddressRequest) SetText(v string) *StructureAddressRequest {
	s.Text = &v
	return s
}

type StructureAddressResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StructureAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StructureAddressResponseBody) GoString() string {
	return s.String()
}

func (s *StructureAddressResponseBody) SetData(v string) *StructureAddressResponseBody {
	s.Data = &v
	return s
}

func (s *StructureAddressResponseBody) SetRequestId(v string) *StructureAddressResponseBody {
	s.RequestId = &v
	return s
}

type StructureAddressResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StructureAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StructureAddressResponse) SetStatusCode(v int32) *StructureAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *StructureAddressResponse) SetBody(v *StructureAddressResponseBody) *StructureAddressResponse {
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

func (client *Client) CorrectAddressWithOptions(request *CorrectAddressRequest, runtime *util.RuntimeOptions) (_result *CorrectAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCity)) {
		body["DefaultCity"] = request.DefaultCity
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDistrict)) {
		body["DefaultDistrict"] = request.DefaultDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultProvince)) {
		body["DefaultProvince"] = request.DefaultProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CorrectAddress"),
		Version:     tea.String("2019-11-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CorrectAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ExtractAddressWithOptions(request *ExtractAddressRequest, runtime *util.RuntimeOptions) (_result *ExtractAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCity)) {
		body["DefaultCity"] = request.DefaultCity
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDistrict)) {
		body["DefaultDistrict"] = request.DefaultDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultProvince)) {
		body["DefaultProvince"] = request.DefaultProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtractAddress"),
		Version:     tea.String("2019-11-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtractAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ExtractNameWithOptions(request *ExtractNameRequest, runtime *util.RuntimeOptions) (_result *ExtractNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCity)) {
		body["DefaultCity"] = request.DefaultCity
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDistrict)) {
		body["DefaultDistrict"] = request.DefaultDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultProvince)) {
		body["DefaultProvince"] = request.DefaultProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtractName"),
		Version:     tea.String("2019-11-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtractNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ExtractPhoneWithOptions(request *ExtractPhoneRequest, runtime *util.RuntimeOptions) (_result *ExtractPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCity)) {
		body["DefaultCity"] = request.DefaultCity
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDistrict)) {
		body["DefaultDistrict"] = request.DefaultDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultProvince)) {
		body["DefaultProvince"] = request.DefaultProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtractPhone"),
		Version:     tea.String("2019-11-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtractPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetAddressDivisionCodeWithOptions(request *GetAddressDivisionCodeRequest, runtime *util.RuntimeOptions) (_result *GetAddressDivisionCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCity)) {
		body["DefaultCity"] = request.DefaultCity
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDistrict)) {
		body["DefaultDistrict"] = request.DefaultDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultProvince)) {
		body["DefaultProvince"] = request.DefaultProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAddressDivisionCode"),
		Version:     tea.String("2019-11-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAddressDivisionCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetAddressSimilarityWithOptions(request *GetAddressSimilarityRequest, runtime *util.RuntimeOptions) (_result *GetAddressSimilarityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCity)) {
		body["DefaultCity"] = request.DefaultCity
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDistrict)) {
		body["DefaultDistrict"] = request.DefaultDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultProvince)) {
		body["DefaultProvince"] = request.DefaultProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAddressSimilarity"),
		Version:     tea.String("2019-11-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAddressSimilarityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetZipcodeWithOptions(request *GetZipcodeRequest, runtime *util.RuntimeOptions) (_result *GetZipcodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCity)) {
		body["DefaultCity"] = request.DefaultCity
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDistrict)) {
		body["DefaultDistrict"] = request.DefaultDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultProvince)) {
		body["DefaultProvince"] = request.DefaultProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetZipcode"),
		Version:     tea.String("2019-11-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetZipcodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) StructureAddressWithOptions(request *StructureAddressRequest, runtime *util.RuntimeOptions) (_result *StructureAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCity)) {
		body["DefaultCity"] = request.DefaultCity
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDistrict)) {
		body["DefaultDistrict"] = request.DefaultDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultProvince)) {
		body["DefaultProvince"] = request.DefaultProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StructureAddress"),
		Version:     tea.String("2019-11-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StructureAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
