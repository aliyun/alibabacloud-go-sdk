// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateLicenseRequest struct {
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	LicensePublisher *string `json:"LicensePublisher,omitempty" xml:"LicensePublisher,omitempty"`
	LicenseCode      *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
}

func (s ActivateLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActivateLicenseRequest) SetBizType(v string) *ActivateLicenseRequest {
	s.BizType = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicensePublisher(v string) *ActivateLicenseRequest {
	s.LicensePublisher = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicenseCode(v string) *ActivateLicenseRequest {
	s.LicenseCode = &v
	return s
}

type ActivateLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ActivateLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBody) SetRequestId(v string) *ActivateLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetData(v bool) *ActivateLicenseResponseBody {
	s.Data = &v
	return s
}

type ActivateLicenseResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ActivateLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivateLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponse) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponse) SetHeaders(v map[string]*string) *ActivateLicenseResponse {
	s.Headers = v
	return s
}

func (s *ActivateLicenseResponse) SetBody(v *ActivateLicenseResponseBody) *ActivateLicenseResponse {
	s.Body = v
	return s
}

type BusinessLicenseOcrRequest struct {
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizCode       *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	FileInfo      *string `json:"FileInfo,omitempty" xml:"FileInfo,omitempty"`
	FileStoreType *string `json:"FileStoreType,omitempty" xml:"FileStoreType,omitempty"`
}

func (s BusinessLicenseOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s BusinessLicenseOcrRequest) GoString() string {
	return s.String()
}

func (s *BusinessLicenseOcrRequest) SetLang(v string) *BusinessLicenseOcrRequest {
	s.Lang = &v
	return s
}

func (s *BusinessLicenseOcrRequest) SetBizCode(v string) *BusinessLicenseOcrRequest {
	s.BizCode = &v
	return s
}

func (s *BusinessLicenseOcrRequest) SetFileInfo(v string) *BusinessLicenseOcrRequest {
	s.FileInfo = &v
	return s
}

func (s *BusinessLicenseOcrRequest) SetFileStoreType(v string) *BusinessLicenseOcrRequest {
	s.FileStoreType = &v
	return s
}

type BusinessLicenseOcrResponseBody struct {
	RegisterNumber *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ValidPeriod    *string `json:"ValidPeriod,omitempty" xml:"ValidPeriod,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Address        *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Capital        *string `json:"Capital,omitempty" xml:"Capital,omitempty"`
	LegalPerson    *string `json:"LegalPerson,omitempty" xml:"LegalPerson,omitempty"`
	EstablishDate  *string `json:"EstablishDate,omitempty" xml:"EstablishDate,omitempty"`
	Nationality    *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Business       *string `json:"Business,omitempty" xml:"Business,omitempty"`
	TrackId        *string `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s BusinessLicenseOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BusinessLicenseOcrResponseBody) GoString() string {
	return s.String()
}

func (s *BusinessLicenseOcrResponseBody) SetRegisterNumber(v string) *BusinessLicenseOcrResponseBody {
	s.RegisterNumber = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetType(v string) *BusinessLicenseOcrResponseBody {
	s.Type = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetValidPeriod(v string) *BusinessLicenseOcrResponseBody {
	s.ValidPeriod = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetRequestId(v string) *BusinessLicenseOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetAddress(v string) *BusinessLicenseOcrResponseBody {
	s.Address = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetCapital(v string) *BusinessLicenseOcrResponseBody {
	s.Capital = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetLegalPerson(v string) *BusinessLicenseOcrResponseBody {
	s.LegalPerson = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetEstablishDate(v string) *BusinessLicenseOcrResponseBody {
	s.EstablishDate = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetNationality(v string) *BusinessLicenseOcrResponseBody {
	s.Nationality = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetName(v string) *BusinessLicenseOcrResponseBody {
	s.Name = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetBusiness(v string) *BusinessLicenseOcrResponseBody {
	s.Business = &v
	return s
}

func (s *BusinessLicenseOcrResponseBody) SetTrackId(v string) *BusinessLicenseOcrResponseBody {
	s.TrackId = &v
	return s
}

type BusinessLicenseOcrResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BusinessLicenseOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BusinessLicenseOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s BusinessLicenseOcrResponse) GoString() string {
	return s.String()
}

func (s *BusinessLicenseOcrResponse) SetHeaders(v map[string]*string) *BusinessLicenseOcrResponse {
	s.Headers = v
	return s
}

func (s *BusinessLicenseOcrResponse) SetBody(v *BusinessLicenseOcrResponseBody) *BusinessLicenseOcrResponse {
	s.Body = v
	return s
}

type CertificateQualityRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizCode         *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	FileInfo        *string `json:"FileInfo,omitempty" xml:"FileInfo,omitempty"`
	FileStoreType   *string `json:"FileStoreType,omitempty" xml:"FileStoreType,omitempty"`
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
}

func (s CertificateQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificateQualityRequest) GoString() string {
	return s.String()
}

func (s *CertificateQualityRequest) SetLang(v string) *CertificateQualityRequest {
	s.Lang = &v
	return s
}

func (s *CertificateQualityRequest) SetBizCode(v string) *CertificateQualityRequest {
	s.BizCode = &v
	return s
}

func (s *CertificateQualityRequest) SetFileInfo(v string) *CertificateQualityRequest {
	s.FileInfo = &v
	return s
}

func (s *CertificateQualityRequest) SetFileStoreType(v string) *CertificateQualityRequest {
	s.FileStoreType = &v
	return s
}

func (s *CertificateQualityRequest) SetCertificateType(v string) *CertificateQualityRequest {
	s.CertificateType = &v
	return s
}

type CertificateQualityResponseBody struct {
	ContainSeal      *string `json:"ContainSeal,omitempty" xml:"ContainSeal,omitempty"`
	ContainWatermark *string `json:"ContainWatermark,omitempty" xml:"ContainWatermark,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Copy             *string `json:"Copy,omitempty" xml:"Copy,omitempty"`
	Complete         *string `json:"Complete,omitempty" xml:"Complete,omitempty"`
	NationalEmblem   *string `json:"NationalEmblem,omitempty" xml:"NationalEmblem,omitempty"`
	TargetType       *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	Reflection       *string `json:"Reflection,omitempty" xml:"Reflection,omitempty"`
	Electronic       *string `json:"Electronic,omitempty" xml:"Electronic,omitempty"`
	ContainFront     *string `json:"ContainFront,omitempty" xml:"ContainFront,omitempty"`
	TextClarity      *string `json:"TextClarity,omitempty" xml:"TextClarity,omitempty"`
}

func (s CertificateQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertificateQualityResponseBody) GoString() string {
	return s.String()
}

func (s *CertificateQualityResponseBody) SetContainSeal(v string) *CertificateQualityResponseBody {
	s.ContainSeal = &v
	return s
}

func (s *CertificateQualityResponseBody) SetContainWatermark(v string) *CertificateQualityResponseBody {
	s.ContainWatermark = &v
	return s
}

func (s *CertificateQualityResponseBody) SetRequestId(v string) *CertificateQualityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificateQualityResponseBody) SetCopy(v string) *CertificateQualityResponseBody {
	s.Copy = &v
	return s
}

func (s *CertificateQualityResponseBody) SetComplete(v string) *CertificateQualityResponseBody {
	s.Complete = &v
	return s
}

func (s *CertificateQualityResponseBody) SetNationalEmblem(v string) *CertificateQualityResponseBody {
	s.NationalEmblem = &v
	return s
}

func (s *CertificateQualityResponseBody) SetTargetType(v string) *CertificateQualityResponseBody {
	s.TargetType = &v
	return s
}

func (s *CertificateQualityResponseBody) SetReflection(v string) *CertificateQualityResponseBody {
	s.Reflection = &v
	return s
}

func (s *CertificateQualityResponseBody) SetElectronic(v string) *CertificateQualityResponseBody {
	s.Electronic = &v
	return s
}

func (s *CertificateQualityResponseBody) SetContainFront(v string) *CertificateQualityResponseBody {
	s.ContainFront = &v
	return s
}

func (s *CertificateQualityResponseBody) SetTextClarity(v string) *CertificateQualityResponseBody {
	s.TextClarity = &v
	return s
}

type CertificateQualityResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CertificateQualityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CertificateQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificateQualityResponse) GoString() string {
	return s.String()
}

func (s *CertificateQualityResponse) SetHeaders(v map[string]*string) *CertificateQualityResponse {
	s.Headers = v
	return s
}

func (s *CertificateQualityResponse) SetBody(v *CertificateQualityResponseBody) *CertificateQualityResponse {
	s.Body = v
	return s
}

type DescribeAgreementStatusRequest struct {
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s DescribeAgreementStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgreementStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusRequest) SetAgreementCode(v string) *DescribeAgreementStatusRequest {
	s.AgreementCode = &v
	return s
}

type DescribeAgreementStatusResponseBody struct {
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s DescribeAgreementStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgreementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusResponseBody) SetStatus(v int32) *DescribeAgreementStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetRequestId(v string) *DescribeAgreementStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetUserId(v string) *DescribeAgreementStatusResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetAgreementCode(v string) *DescribeAgreementStatusResponseBody {
	s.AgreementCode = &v
	return s
}

type DescribeAgreementStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAgreementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAgreementStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgreementStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusResponse) SetHeaders(v map[string]*string) *DescribeAgreementStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgreementStatusResponse) SetBody(v *DescribeAgreementStatusResponseBody) *DescribeAgreementStatusResponse {
	s.Body = v
	return s
}

type IdentityCardOcrRequest struct {
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizCode       *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	FileInfo      *string `json:"FileInfo,omitempty" xml:"FileInfo,omitempty"`
	FileStoreType *string `json:"FileStoreType,omitempty" xml:"FileStoreType,omitempty"`
}

func (s IdentityCardOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s IdentityCardOcrRequest) GoString() string {
	return s.String()
}

func (s *IdentityCardOcrRequest) SetLang(v string) *IdentityCardOcrRequest {
	s.Lang = &v
	return s
}

func (s *IdentityCardOcrRequest) SetBizCode(v string) *IdentityCardOcrRequest {
	s.BizCode = &v
	return s
}

func (s *IdentityCardOcrRequest) SetFileInfo(v string) *IdentityCardOcrRequest {
	s.FileInfo = &v
	return s
}

func (s *IdentityCardOcrRequest) SetFileStoreType(v string) *IdentityCardOcrRequest {
	s.FileStoreType = &v
	return s
}

type IdentityCardOcrResponseBody struct {
	Issue       *string `json:"Issue,omitempty" xml:"Issue,omitempty"`
	ValidDate   *string `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	Gender      *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Nationality *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	BirthDate   *string `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	TrackId     *string `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s IdentityCardOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IdentityCardOcrResponseBody) GoString() string {
	return s.String()
}

func (s *IdentityCardOcrResponseBody) SetIssue(v string) *IdentityCardOcrResponseBody {
	s.Issue = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetValidDate(v string) *IdentityCardOcrResponseBody {
	s.ValidDate = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetRequestId(v string) *IdentityCardOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetAddress(v string) *IdentityCardOcrResponseBody {
	s.Address = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetIdNumber(v string) *IdentityCardOcrResponseBody {
	s.IdNumber = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetGender(v string) *IdentityCardOcrResponseBody {
	s.Gender = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetNationality(v string) *IdentityCardOcrResponseBody {
	s.Nationality = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetBirthDate(v string) *IdentityCardOcrResponseBody {
	s.BirthDate = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetTrackId(v string) *IdentityCardOcrResponseBody {
	s.TrackId = &v
	return s
}

func (s *IdentityCardOcrResponseBody) SetName(v string) *IdentityCardOcrResponseBody {
	s.Name = &v
	return s
}

type IdentityCardOcrResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IdentityCardOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IdentityCardOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s IdentityCardOcrResponse) GoString() string {
	return s.String()
}

func (s *IdentityCardOcrResponse) SetHeaders(v map[string]*string) *IdentityCardOcrResponse {
	s.Headers = v
	return s
}

func (s *IdentityCardOcrResponse) SetBody(v *IdentityCardOcrResponseBody) *IdentityCardOcrResponse {
	s.Body = v
	return s
}

type UpdateAgreementStatusRequest struct {
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s UpdateAgreementStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgreementStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgreementStatusRequest) SetAgreementCode(v string) *UpdateAgreementStatusRequest {
	s.AgreementCode = &v
	return s
}

type UpdateAgreementStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAgreementStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgreementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgreementStatusResponseBody) SetRequestId(v string) *UpdateAgreementStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAgreementStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAgreementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAgreementStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgreementStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgreementStatusResponse) SetHeaders(v map[string]*string) *UpdateAgreementStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgreementStatusResponse) SetBody(v *UpdateAgreementStatusResponseBody) *UpdateAgreementStatusResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("mseap"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ActivateLicenseWithOptions(request *ActivateLicenseRequest, runtime *util.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ActivateLicense"), tea.String("2021-01-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivateLicense(request *ActivateLicenseRequest) (_result *ActivateLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.ActivateLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BusinessLicenseOcrWithOptions(request *BusinessLicenseOcrRequest, runtime *util.RuntimeOptions) (_result *BusinessLicenseOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BusinessLicenseOcrResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BusinessLicenseOcr"), tea.String("2021-01-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BusinessLicenseOcr(request *BusinessLicenseOcrRequest) (_result *BusinessLicenseOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BusinessLicenseOcrResponse{}
	_body, _err := client.BusinessLicenseOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CertificateQualityWithOptions(request *CertificateQualityRequest, runtime *util.RuntimeOptions) (_result *CertificateQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CertificateQualityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CertificateQuality"), tea.String("2021-01-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CertificateQuality(request *CertificateQualityRequest) (_result *CertificateQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CertificateQualityResponse{}
	_body, _err := client.CertificateQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAgreementStatusWithOptions(request *DescribeAgreementStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeAgreementStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAgreementStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAgreementStatus"), tea.String("2021-01-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAgreementStatus(request *DescribeAgreementStatusRequest) (_result *DescribeAgreementStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAgreementStatusResponse{}
	_body, _err := client.DescribeAgreementStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IdentityCardOcrWithOptions(request *IdentityCardOcrRequest, runtime *util.RuntimeOptions) (_result *IdentityCardOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IdentityCardOcrResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IdentityCardOcr"), tea.String("2021-01-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IdentityCardOcr(request *IdentityCardOcrRequest) (_result *IdentityCardOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IdentityCardOcrResponse{}
	_body, _err := client.IdentityCardOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAgreementStatusWithOptions(request *UpdateAgreementStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateAgreementStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAgreementStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAgreementStatus"), tea.String("2021-01-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAgreementStatus(request *UpdateAgreementStatusRequest) (_result *UpdateAgreementStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAgreementStatusResponse{}
	_body, _err := client.UpdateAgreementStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
