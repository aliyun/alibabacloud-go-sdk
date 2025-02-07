// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeDeliveryAddressResponseBody struct {
	Addresses []*DescribeDeliveryAddressResponseBodyAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// 72481C12-69AB-5CE1-8A35-A8EFA921****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeliveryAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBody) SetAddresses(v []*DescribeDeliveryAddressResponseBodyAddresses) *DescribeDeliveryAddressResponseBody {
	s.Addresses = v
	return s
}

func (s *DescribeDeliveryAddressResponseBody) SetRequestId(v string) *DescribeDeliveryAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBody) SetTotalCount(v int32) *DescribeDeliveryAddressResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDeliveryAddressResponseBodyAddresses struct {
	Area *DescribeDeliveryAddressResponseBodyAddressesArea `json:"Area,omitempty" xml:"Area,omitempty" type:"Struct"`
	City *DescribeDeliveryAddressResponseBodyAddressesCity `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	// example:
	//
	// Alice
	Contacts *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	// example:
	//
	// true
	DefaultAddress *bool   `json:"DefaultAddress,omitempty" xml:"DefaultAddress,omitempty"`
	Detail         *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// 1381111****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 03****
	PostalCode *string                                               `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province   *DescribeDeliveryAddressResponseBodyAddressesProvince `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	Town       *DescribeDeliveryAddressResponseBodyAddressesTown     `json:"Town,omitempty" xml:"Town,omitempty" type:"Struct"`
}

func (s DescribeDeliveryAddressResponseBodyAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddresses) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetArea(v *DescribeDeliveryAddressResponseBodyAddressesArea) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Area = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetCity(v *DescribeDeliveryAddressResponseBodyAddressesCity) *DescribeDeliveryAddressResponseBodyAddresses {
	s.City = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetContacts(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Contacts = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetDefaultAddress(v bool) *DescribeDeliveryAddressResponseBodyAddresses {
	s.DefaultAddress = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetDetail(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Detail = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetMobile(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Mobile = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetPostalCode(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.PostalCode = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetProvince(v *DescribeDeliveryAddressResponseBodyAddressesProvince) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Province = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetTown(v *DescribeDeliveryAddressResponseBodyAddressesTown) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Town = v
	return s
}

type DescribeDeliveryAddressResponseBodyAddressesArea struct {
	// example:
	//
	// 33****
	AreaId   *int64  `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesArea) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesArea) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) SetAreaId(v int64) *DescribeDeliveryAddressResponseBodyAddressesArea {
	s.AreaId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) SetAreaName(v string) *DescribeDeliveryAddressResponseBodyAddressesArea {
	s.AreaName = &v
	return s
}

type DescribeDeliveryAddressResponseBodyAddressesCity struct {
	// example:
	//
	// 33****
	CityId   *int64  `json:"CityId,omitempty" xml:"CityId,omitempty"`
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesCity) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesCity) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) SetCityId(v int64) *DescribeDeliveryAddressResponseBodyAddressesCity {
	s.CityId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) SetCityName(v string) *DescribeDeliveryAddressResponseBodyAddressesCity {
	s.CityName = &v
	return s
}

type DescribeDeliveryAddressResponseBodyAddressesProvince struct {
	// example:
	//
	// 330000
	ProvinceId   *int64  `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesProvince) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesProvince) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) SetProvinceId(v int64) *DescribeDeliveryAddressResponseBodyAddressesProvince {
	s.ProvinceId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) SetProvinceName(v string) *DescribeDeliveryAddressResponseBodyAddressesProvince {
	s.ProvinceName = &v
	return s
}

type DescribeDeliveryAddressResponseBodyAddressesTown struct {
	// example:
	//
	// 3001****
	TownId   *int64  `json:"TownId,omitempty" xml:"TownId,omitempty"`
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesTown) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesTown) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) SetTownId(v int64) *DescribeDeliveryAddressResponseBodyAddressesTown {
	s.TownId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) SetTownName(v string) *DescribeDeliveryAddressResponseBodyAddressesTown {
	s.TownName = &v
	return s
}

type DescribeDeliveryAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeliveryAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeliveryAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponse) SetHeaders(v map[string]*string) *DescribeDeliveryAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeliveryAddressResponse) SetStatusCode(v int32) *DescribeDeliveryAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeliveryAddressResponse) SetBody(v *DescribeDeliveryAddressResponseBody) *DescribeDeliveryAddressResponse {
	s.Body = v
	return s
}

type DescribePackageDeductionsRequest struct {
	EndTime     *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	PackageIds  []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CorePackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePackageDeductionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageDeductionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsRequest) SetEndTime(v int64) *DescribePackageDeductionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetInstanceIds(v []*string) *DescribePackageDeductionsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPackageIds(v []*string) *DescribePackageDeductionsRequest {
	s.PackageIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPageNum(v int32) *DescribePackageDeductionsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPageSize(v int32) *DescribePackageDeductionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetResourceType(v string) *DescribePackageDeductionsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetStartTime(v int64) *DescribePackageDeductionsRequest {
	s.StartTime = &v
	return s
}

type DescribePackageDeductionsResponseBody struct {
	Deductions []*DescribePackageDeductionsResponseBodyDeductions `json:"Deductions,omitempty" xml:"Deductions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount        *int64   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalUsedCoreTime *float32 `json:"TotalUsedCoreTime,omitempty" xml:"TotalUsedCoreTime,omitempty"`
	TotalUsedTime     *int64   `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
}

func (s DescribePackageDeductionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageDeductionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponseBody) SetDeductions(v []*DescribePackageDeductionsResponseBodyDeductions) *DescribePackageDeductionsResponseBody {
	s.Deductions = v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetPageNum(v int32) *DescribePackageDeductionsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetPageSize(v int32) *DescribePackageDeductionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetRequestId(v string) *DescribePackageDeductionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalCount(v int64) *DescribePackageDeductionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalUsedCoreTime(v float32) *DescribePackageDeductionsResponseBody {
	s.TotalUsedCoreTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalUsedTime(v int64) *DescribePackageDeductionsResponseBody {
	s.TotalUsedTime = &v
	return s
}

type DescribePackageDeductionsResponseBodyDeductions struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// ecd-6wye9optu0kag****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// DemoComputer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// eds.enterprise_office.4c8g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// example:
	//
	// 2024-07-31T03:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Deleted
	InstanceState *string `json:"InstanceState,omitempty" xml:"InstanceState,omitempty"`
	// example:
	//
	// 8192
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 2024-07-31T02:00Z
	StaTime *string `json:"StaTime,omitempty" xml:"StaTime,omitempty"`
	// example:
	//
	// 4.0
	UsedCoreTime *float32 `json:"UsedCoreTime,omitempty" xml:"UsedCoreTime,omitempty"`
	// example:
	//
	// 3600
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s DescribePackageDeductionsResponseBodyDeductions) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageDeductionsResponseBodyDeductions) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetCpu(v int32) *DescribePackageDeductionsResponseBodyDeductions {
	s.Cpu = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopName(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopName = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetEndTime(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.EndTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetInstanceState(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.InstanceState = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetMemory(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.Memory = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetOsType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.OsType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetRegionId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.RegionId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetResourceType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.ResourceType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetStaTime(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.StaTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedCoreTime(v float32) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedCoreTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedTime(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedTime = &v
	return s
}

type DescribePackageDeductionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackageDeductionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackageDeductionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageDeductionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponse) SetHeaders(v map[string]*string) *DescribePackageDeductionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePackageDeductionsResponse) SetStatusCode(v int32) *DescribePackageDeductionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackageDeductionsResponse) SetBody(v *DescribePackageDeductionsResponseBody) *DescribePackageDeductionsResponse {
	s.Body = v
	return s
}

type ModifyInstancePropertiesRequest struct {
	// example:
	//
	// mdp-0c62ayep0nk4v****
	InstanceId  *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// PackageUsedUpStrategy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Postpaid
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstancePropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePropertiesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesRequest) SetInstanceId(v string) *ModifyInstancePropertiesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetInstanceIds(v []*string) *ModifyInstancePropertiesRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetKey(v string) *ModifyInstancePropertiesRequest {
	s.Key = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetResourceType(v string) *ModifyInstancePropertiesRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetValue(v string) *ModifyInstancePropertiesRequest {
	s.Value = &v
	return s
}

type ModifyInstancePropertiesResponseBody struct {
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstancePropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesResponseBody) SetRequestId(v string) *ModifyInstancePropertiesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstancePropertiesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstancePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstancePropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePropertiesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesResponse) SetHeaders(v map[string]*string) *ModifyInstancePropertiesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstancePropertiesResponse) SetStatusCode(v int32) *ModifyInstancePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstancePropertiesResponse) SetBody(v *ModifyInstancePropertiesResponseBody) *ModifyInstancePropertiesResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("wss"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - DescribeDeliveryAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeliveryAddressResponse
func (client *Client) DescribeDeliveryAddressWithOptions(runtime *util.RuntimeOptions) (_result *DescribeDeliveryAddressResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeliveryAddress"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDeliveryAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDeliveryAddressResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return DescribeDeliveryAddressResponse
func (client *Client) DescribeDeliveryAddress() (_result *DescribeDeliveryAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeliveryAddressResponse{}
	_body, _err := client.DescribeDeliveryAddressWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询核时包抵扣明细
//
// @param request - DescribePackageDeductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackageDeductionsResponse
func (client *Client) DescribePackageDeductionsWithOptions(request *DescribePackageDeductionsRequest, runtime *util.RuntimeOptions) (_result *DescribePackageDeductionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.PackageIds)) {
		query["PackageIds"] = request.PackageIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackageDeductions"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePackageDeductionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePackageDeductionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询核时包抵扣明细
//
// @param request - DescribePackageDeductionsRequest
//
// @return DescribePackageDeductionsResponse
func (client *Client) DescribePackageDeductions(request *DescribePackageDeductionsRequest) (_result *DescribePackageDeductionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackageDeductionsResponse{}
	_body, _err := client.DescribePackageDeductionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyInstancePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstancePropertiesResponse
func (client *Client) ModifyInstancePropertiesWithOptions(request *ModifyInstancePropertiesRequest, runtime *util.RuntimeOptions) (_result *ModifyInstancePropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceProperties"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyInstancePropertiesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyInstancePropertiesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyInstancePropertiesRequest
//
// @return ModifyInstancePropertiesResponse
func (client *Client) ModifyInstanceProperties(request *ModifyInstancePropertiesRequest) (_result *ModifyInstancePropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstancePropertiesResponse{}
	_body, _err := client.ModifyInstancePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
