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

type Address struct {
	CityCode     *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	Detail       *string `json:"detail,omitempty" xml:"detail,omitempty"`
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s Address) String() string {
	return tea.Prettify(s)
}

func (s Address) GoString() string {
	return s.String()
}

func (s *Address) SetCityCode(v string) *Address {
	s.CityCode = &v
	return s
}

func (s *Address) SetDetail(v string) *Address {
	s.Detail = &v
	return s
}

func (s *Address) SetDistrictCode(v string) *Address {
	s.DistrictCode = &v
	return s
}

func (s *Address) SetProvinceCode(v string) *Address {
	s.ProvinceCode = &v
	return s
}

type Company struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Uscc *string `json:"uscc,omitempty" xml:"uscc,omitempty"`
}

func (s Company) String() string {
	return tea.Prettify(s)
}

func (s Company) GoString() string {
	return s.String()
}

func (s *Company) SetName(v string) *Company {
	s.Name = &v
	return s
}

func (s *Company) SetUscc(v string) *Company {
	s.Uscc = &v
	return s
}

type Contact struct {
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s Contact) String() string {
	return tea.Prettify(s)
}

func (s Contact) GoString() string {
	return s.String()
}

func (s *Contact) SetEmail(v string) *Contact {
	s.Email = &v
	return s
}

func (s *Contact) SetName(v string) *Contact {
	s.Name = &v
	return s
}

func (s *Contact) SetPhone(v string) *Contact {
	s.Phone = &v
	return s
}

type ExtendInfo struct {
	DepositAmount *float64 `json:"depositAmount,omitempty" xml:"depositAmount,omitempty"`
	Desc          *string  `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s ExtendInfo) String() string {
	return tea.Prettify(s)
}

func (s ExtendInfo) GoString() string {
	return s.String()
}

func (s *ExtendInfo) SetDepositAmount(v float64) *ExtendInfo {
	s.DepositAmount = &v
	return s
}

func (s *ExtendInfo) SetDesc(v string) *ExtendInfo {
	s.Desc = &v
	return s
}

type SubjectExtendInfo struct {
	DeliveryDesc *string `json:"deliveryDesc,omitempty" xml:"deliveryDesc,omitempty"`
	Desc         *string `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s SubjectExtendInfo) String() string {
	return tea.Prettify(s)
}

func (s SubjectExtendInfo) GoString() string {
	return s.String()
}

func (s *SubjectExtendInfo) SetDeliveryDesc(v string) *SubjectExtendInfo {
	s.DeliveryDesc = &v
	return s
}

func (s *SubjectExtendInfo) SetDesc(v string) *SubjectExtendInfo {
	s.Desc = &v
	return s
}

type CreateSourcingProjectRequest struct {
	Address    *CreateSourcingProjectRequestAddress    `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	BizId      *string                                 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizNo      *string                                 `json:"BizNo,omitempty" xml:"BizNo,omitempty"`
	BizType    *string                                 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Company    *CreateSourcingProjectRequestCompany    `json:"Company,omitempty" xml:"Company,omitempty" type:"Struct"`
	Contact    *CreateSourcingProjectRequestContact    `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	CreateTime *string                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime *string                                 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExtendInfo map[string]*string                      `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	SourceUrl  *string                                 `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	SubBizType *string                                 `json:"SubBizType,omitempty" xml:"SubBizType,omitempty"`
	Subjects   []*CreateSourcingProjectRequestSubjects `json:"Subjects,omitempty" xml:"Subjects,omitempty" type:"Repeated"`
	Title      *string                                 `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateSourcingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectRequest) SetAddress(v *CreateSourcingProjectRequestAddress) *CreateSourcingProjectRequest {
	s.Address = v
	return s
}

func (s *CreateSourcingProjectRequest) SetBizId(v string) *CreateSourcingProjectRequest {
	s.BizId = &v
	return s
}

func (s *CreateSourcingProjectRequest) SetBizNo(v string) *CreateSourcingProjectRequest {
	s.BizNo = &v
	return s
}

func (s *CreateSourcingProjectRequest) SetBizType(v string) *CreateSourcingProjectRequest {
	s.BizType = &v
	return s
}

func (s *CreateSourcingProjectRequest) SetCompany(v *CreateSourcingProjectRequestCompany) *CreateSourcingProjectRequest {
	s.Company = v
	return s
}

func (s *CreateSourcingProjectRequest) SetContact(v *CreateSourcingProjectRequestContact) *CreateSourcingProjectRequest {
	s.Contact = v
	return s
}

func (s *CreateSourcingProjectRequest) SetCreateTime(v string) *CreateSourcingProjectRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateSourcingProjectRequest) SetExpireTime(v string) *CreateSourcingProjectRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateSourcingProjectRequest) SetExtendInfo(v map[string]*string) *CreateSourcingProjectRequest {
	s.ExtendInfo = v
	return s
}

func (s *CreateSourcingProjectRequest) SetSourceUrl(v string) *CreateSourcingProjectRequest {
	s.SourceUrl = &v
	return s
}

func (s *CreateSourcingProjectRequest) SetSubBizType(v string) *CreateSourcingProjectRequest {
	s.SubBizType = &v
	return s
}

func (s *CreateSourcingProjectRequest) SetSubjects(v []*CreateSourcingProjectRequestSubjects) *CreateSourcingProjectRequest {
	s.Subjects = v
	return s
}

func (s *CreateSourcingProjectRequest) SetTitle(v string) *CreateSourcingProjectRequest {
	s.Title = &v
	return s
}

type CreateSourcingProjectRequestAddress struct {
	CityCode     *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	Detail       *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	DistrictCode *string `json:"DistrictCode,omitempty" xml:"DistrictCode,omitempty"`
	ProvinceCode *string `json:"ProvinceCode,omitempty" xml:"ProvinceCode,omitempty"`
}

func (s CreateSourcingProjectRequestAddress) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectRequestAddress) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectRequestAddress) SetCityCode(v string) *CreateSourcingProjectRequestAddress {
	s.CityCode = &v
	return s
}

func (s *CreateSourcingProjectRequestAddress) SetDetail(v string) *CreateSourcingProjectRequestAddress {
	s.Detail = &v
	return s
}

func (s *CreateSourcingProjectRequestAddress) SetDistrictCode(v string) *CreateSourcingProjectRequestAddress {
	s.DistrictCode = &v
	return s
}

func (s *CreateSourcingProjectRequestAddress) SetProvinceCode(v string) *CreateSourcingProjectRequestAddress {
	s.ProvinceCode = &v
	return s
}

type CreateSourcingProjectRequestCompany struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Uscc *string `json:"Uscc,omitempty" xml:"Uscc,omitempty"`
}

func (s CreateSourcingProjectRequestCompany) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectRequestCompany) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectRequestCompany) SetName(v string) *CreateSourcingProjectRequestCompany {
	s.Name = &v
	return s
}

func (s *CreateSourcingProjectRequestCompany) SetUscc(v string) *CreateSourcingProjectRequestCompany {
	s.Uscc = &v
	return s
}

type CreateSourcingProjectRequestContact struct {
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s CreateSourcingProjectRequestContact) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectRequestContact) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectRequestContact) SetEmail(v string) *CreateSourcingProjectRequestContact {
	s.Email = &v
	return s
}

func (s *CreateSourcingProjectRequestContact) SetName(v string) *CreateSourcingProjectRequestContact {
	s.Name = &v
	return s
}

func (s *CreateSourcingProjectRequestContact) SetPhone(v string) *CreateSourcingProjectRequestContact {
	s.Phone = &v
	return s
}

type CreateSourcingProjectRequestSubjects struct {
	Address    *CreateSourcingProjectRequestSubjectsAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	ExtendInfo map[string]*string                           `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	Name       *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Quantity   *float64                                     `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Spec       *string                                      `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Unit       *string                                      `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s CreateSourcingProjectRequestSubjects) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectRequestSubjects) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectRequestSubjects) SetAddress(v *CreateSourcingProjectRequestSubjectsAddress) *CreateSourcingProjectRequestSubjects {
	s.Address = v
	return s
}

func (s *CreateSourcingProjectRequestSubjects) SetCode(v string) *CreateSourcingProjectRequestSubjects {
	s.Code = &v
	return s
}

func (s *CreateSourcingProjectRequestSubjects) SetExtendInfo(v map[string]*string) *CreateSourcingProjectRequestSubjects {
	s.ExtendInfo = v
	return s
}

func (s *CreateSourcingProjectRequestSubjects) SetName(v string) *CreateSourcingProjectRequestSubjects {
	s.Name = &v
	return s
}

func (s *CreateSourcingProjectRequestSubjects) SetQuantity(v float64) *CreateSourcingProjectRequestSubjects {
	s.Quantity = &v
	return s
}

func (s *CreateSourcingProjectRequestSubjects) SetSpec(v string) *CreateSourcingProjectRequestSubjects {
	s.Spec = &v
	return s
}

func (s *CreateSourcingProjectRequestSubjects) SetUnit(v string) *CreateSourcingProjectRequestSubjects {
	s.Unit = &v
	return s
}

type CreateSourcingProjectRequestSubjectsAddress struct {
	CityCode     *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	Detail       *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	DistrictCode *string `json:"DistrictCode,omitempty" xml:"DistrictCode,omitempty"`
	ProvinceCode *string `json:"ProvinceCode,omitempty" xml:"ProvinceCode,omitempty"`
}

func (s CreateSourcingProjectRequestSubjectsAddress) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectRequestSubjectsAddress) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectRequestSubjectsAddress) SetCityCode(v string) *CreateSourcingProjectRequestSubjectsAddress {
	s.CityCode = &v
	return s
}

func (s *CreateSourcingProjectRequestSubjectsAddress) SetDetail(v string) *CreateSourcingProjectRequestSubjectsAddress {
	s.Detail = &v
	return s
}

func (s *CreateSourcingProjectRequestSubjectsAddress) SetDistrictCode(v string) *CreateSourcingProjectRequestSubjectsAddress {
	s.DistrictCode = &v
	return s
}

func (s *CreateSourcingProjectRequestSubjectsAddress) SetProvinceCode(v string) *CreateSourcingProjectRequestSubjectsAddress {
	s.ProvinceCode = &v
	return s
}

type CreateSourcingProjectShrinkRequest struct {
	AddressShrink    *string `json:"Address,omitempty" xml:"Address,omitempty"`
	BizId            *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizNo            *string `json:"BizNo,omitempty" xml:"BizNo,omitempty"`
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CompanyShrink    *string `json:"Company,omitempty" xml:"Company,omitempty"`
	ContactShrink    *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExtendInfoShrink *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	SourceUrl        *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	SubBizType       *string `json:"SubBizType,omitempty" xml:"SubBizType,omitempty"`
	SubjectsShrink   *string `json:"Subjects,omitempty" xml:"Subjects,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateSourcingProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectShrinkRequest) SetAddressShrink(v string) *CreateSourcingProjectShrinkRequest {
	s.AddressShrink = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetBizId(v string) *CreateSourcingProjectShrinkRequest {
	s.BizId = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetBizNo(v string) *CreateSourcingProjectShrinkRequest {
	s.BizNo = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetBizType(v string) *CreateSourcingProjectShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetCompanyShrink(v string) *CreateSourcingProjectShrinkRequest {
	s.CompanyShrink = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetContactShrink(v string) *CreateSourcingProjectShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetCreateTime(v string) *CreateSourcingProjectShrinkRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetExpireTime(v string) *CreateSourcingProjectShrinkRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetExtendInfoShrink(v string) *CreateSourcingProjectShrinkRequest {
	s.ExtendInfoShrink = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetSourceUrl(v string) *CreateSourcingProjectShrinkRequest {
	s.SourceUrl = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetSubBizType(v string) *CreateSourcingProjectShrinkRequest {
	s.SubBizType = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetSubjectsShrink(v string) *CreateSourcingProjectShrinkRequest {
	s.SubjectsShrink = &v
	return s
}

func (s *CreateSourcingProjectShrinkRequest) SetTitle(v string) *CreateSourcingProjectShrinkRequest {
	s.Title = &v
	return s
}

type CreateSourcingProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSourcingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectResponseBody) SetCode(v string) *CreateSourcingProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSourcingProjectResponseBody) SetData(v string) *CreateSourcingProjectResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSourcingProjectResponseBody) SetMessage(v string) *CreateSourcingProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSourcingProjectResponseBody) SetRequestId(v string) *CreateSourcingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSourcingProjectResponseBody) SetSuccess(v bool) *CreateSourcingProjectResponseBody {
	s.Success = &v
	return s
}

type CreateSourcingProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSourcingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSourcingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSourcingProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateSourcingProjectResponse) SetHeaders(v map[string]*string) *CreateSourcingProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateSourcingProjectResponse) SetStatusCode(v int32) *CreateSourcingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSourcingProjectResponse) SetBody(v *CreateSourcingProjectResponseBody) *CreateSourcingProjectResponse {
	s.Body = v
	return s
}

type UpdateSourcingProjectRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateSourcingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSourcingProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateSourcingProjectRequest) SetBizId(v string) *UpdateSourcingProjectRequest {
	s.BizId = &v
	return s
}

func (s *UpdateSourcingProjectRequest) SetStatus(v string) *UpdateSourcingProjectRequest {
	s.Status = &v
	return s
}

func (s *UpdateSourcingProjectRequest) SetUpdateTime(v string) *UpdateSourcingProjectRequest {
	s.UpdateTime = &v
	return s
}

type UpdateSourcingProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSourcingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSourcingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSourcingProjectResponseBody) SetCode(v string) *UpdateSourcingProjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSourcingProjectResponseBody) SetData(v string) *UpdateSourcingProjectResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSourcingProjectResponseBody) SetMessage(v string) *UpdateSourcingProjectResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSourcingProjectResponseBody) SetRequestId(v string) *UpdateSourcingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSourcingProjectResponseBody) SetSuccess(v bool) *UpdateSourcingProjectResponseBody {
	s.Success = &v
	return s
}

type UpdateSourcingProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSourcingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSourcingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSourcingProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateSourcingProjectResponse) SetHeaders(v map[string]*string) *UpdateSourcingProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateSourcingProjectResponse) SetStatusCode(v int32) *UpdateSourcingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSourcingProjectResponse) SetBody(v *UpdateSourcingProjectResponseBody) *UpdateSourcingProjectResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("alicloudproc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateSourcingProjectWithOptions(tmpReq *CreateSourcingProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSourcingProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSourcingProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Address)) {
		request.AddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Address, tea.String("Address"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Company)) {
		request.CompanyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Company, tea.String("Company"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Contact)) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, tea.String("Contact"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendInfo)) {
		request.ExtendInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendInfo, tea.String("ExtendInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Subjects)) {
		request.SubjectsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Subjects, tea.String("Subjects"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressShrink)) {
		query["Address"] = request.AddressShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizNo)) {
		query["BizNo"] = request.BizNo
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CompanyShrink)) {
		query["Company"] = request.CompanyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ContactShrink)) {
		query["Contact"] = request.ContactShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		query["CreateTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendInfoShrink)) {
		query["ExtendInfo"] = request.ExtendInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUrl)) {
		query["SourceUrl"] = request.SourceUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SubBizType)) {
		query["SubBizType"] = request.SubBizType
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectsShrink)) {
		query["Subjects"] = request.SubjectsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSourcingProject"),
		Version:     tea.String("2024-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/srm/lite/sourcing/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSourcingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSourcingProject(request *CreateSourcingProjectRequest) (_result *CreateSourcingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSourcingProjectResponse{}
	_body, _err := client.CreateSourcingProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSourcingProjectWithOptions(request *UpdateSourcingProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSourcingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTime)) {
		query["UpdateTime"] = request.UpdateTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSourcingProject"),
		Version:     tea.String("2024-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/srm/lite/sourcing/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSourcingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSourcingProject(request *UpdateSourcingProjectRequest) (_result *UpdateSourcingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSourcingProjectResponse{}
	_body, _err := client.UpdateSourcingProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
