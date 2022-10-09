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

type AccessTokenRequest struct {
	AppKey    *string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
}

func (s AccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s AccessTokenRequest) GoString() string {
	return s.String()
}

func (s *AccessTokenRequest) SetAppKey(v string) *AccessTokenRequest {
	s.AppKey = &v
	return s
}

func (s *AccessTokenRequest) SetAppSecret(v string) *AccessTokenRequest {
	s.AppSecret = &v
	return s
}

type AccessTokenResponseBody struct {
	Code      *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data      *AccessTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                      `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TraceId   *string                      `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *AccessTokenResponseBody) SetCode(v string) *AccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *AccessTokenResponseBody) SetData(v *AccessTokenResponseBodyData) *AccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *AccessTokenResponseBody) SetMessage(v string) *AccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *AccessTokenResponseBody) SetRequestId(v string) *AccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccessTokenResponseBody) SetTraceId(v string) *AccessTokenResponseBody {
	s.TraceId = &v
	return s
}

type AccessTokenResponseBodyData struct {
	Expire *int64  `json:"expire,omitempty" xml:"expire,omitempty"`
	Token  *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s AccessTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccessTokenResponseBodyData) SetExpire(v int64) *AccessTokenResponseBodyData {
	s.Expire = &v
	return s
}

func (s *AccessTokenResponseBodyData) SetToken(v string) *AccessTokenResponseBodyData {
	s.Token = &v
	return s
}

type AccessTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s AccessTokenResponse) GoString() string {
	return s.String()
}

func (s *AccessTokenResponse) SetHeaders(v map[string]*string) *AccessTokenResponse {
	s.Headers = v
	return s
}

func (s *AccessTokenResponse) SetStatusCode(v int32) *AccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *AccessTokenResponse) SetBody(v *AccessTokenResponseBody) *AccessTokenResponse {
	s.Body = v
	return s
}

type AddressGetHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s AddressGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddressGetHeaders) GoString() string {
	return s.String()
}

func (s *AddressGetHeaders) SetCommonHeaders(v map[string]*string) *AddressGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddressGetHeaders) SetXAcsBtripSoCorpToken(v string) *AddressGetHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type AddressGetRequest struct {
	ActionType  *int32  `json:"action_type,omitempty" xml:"action_type,omitempty"`
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Type        *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s AddressGetRequest) String() string {
	return tea.Prettify(s)
}

func (s AddressGetRequest) GoString() string {
	return s.String()
}

func (s *AddressGetRequest) SetActionType(v int32) *AddressGetRequest {
	s.ActionType = &v
	return s
}

func (s *AddressGetRequest) SetItineraryId(v string) *AddressGetRequest {
	s.ItineraryId = &v
	return s
}

func (s *AddressGetRequest) SetPhone(v string) *AddressGetRequest {
	s.Phone = &v
	return s
}

func (s *AddressGetRequest) SetType(v int32) *AddressGetRequest {
	s.Type = &v
	return s
}

func (s *AddressGetRequest) SetUserId(v string) *AddressGetRequest {
	s.UserId = &v
	return s
}

type AddressGetResponseBody struct {
	Code      *int32                        `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                       `json:"message,omitempty" xml:"message,omitempty"`
	Module    *AddressGetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                         `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                       `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AddressGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddressGetResponseBody) GoString() string {
	return s.String()
}

func (s *AddressGetResponseBody) SetCode(v int32) *AddressGetResponseBody {
	s.Code = &v
	return s
}

func (s *AddressGetResponseBody) SetMessage(v string) *AddressGetResponseBody {
	s.Message = &v
	return s
}

func (s *AddressGetResponseBody) SetModule(v *AddressGetResponseBodyModule) *AddressGetResponseBody {
	s.Module = v
	return s
}

func (s *AddressGetResponseBody) SetRequestId(v string) *AddressGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddressGetResponseBody) SetSuccess(v bool) *AddressGetResponseBody {
	s.Success = &v
	return s
}

func (s *AddressGetResponseBody) SetTraceId(v string) *AddressGetResponseBody {
	s.TraceId = &v
	return s
}

type AddressGetResponseBodyModule struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddressGetResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s AddressGetResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AddressGetResponseBodyModule) SetUrl(v string) *AddressGetResponseBodyModule {
	s.Url = &v
	return s
}

type AddressGetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddressGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddressGetResponse) String() string {
	return tea.Prettify(s)
}

func (s AddressGetResponse) GoString() string {
	return s.String()
}

func (s *AddressGetResponse) SetHeaders(v map[string]*string) *AddressGetResponse {
	s.Headers = v
	return s
}

func (s *AddressGetResponse) SetStatusCode(v int32) *AddressGetResponse {
	s.StatusCode = &v
	return s
}

func (s *AddressGetResponse) SetBody(v *AddressGetResponseBody) *AddressGetResponse {
	s.Body = v
	return s
}

type AirportSearchHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s AirportSearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s AirportSearchHeaders) GoString() string {
	return s.String()
}

func (s *AirportSearchHeaders) SetCommonHeaders(v map[string]*string) *AirportSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AirportSearchHeaders) SetXAcsBtripSoCorpToken(v string) *AirportSearchHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type AirportSearchRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Type    *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AirportSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s AirportSearchRequest) GoString() string {
	return s.String()
}

func (s *AirportSearchRequest) SetKeyword(v string) *AirportSearchRequest {
	s.Keyword = &v
	return s
}

func (s *AirportSearchRequest) SetType(v int32) *AirportSearchRequest {
	s.Type = &v
	return s
}

type AirportSearchResponseBody struct {
	Code      *int32                           `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                          `json:"message,omitempty" xml:"message,omitempty"`
	Module    *AirportSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                            `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                          `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AirportSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AirportSearchResponseBody) GoString() string {
	return s.String()
}

func (s *AirportSearchResponseBody) SetCode(v int32) *AirportSearchResponseBody {
	s.Code = &v
	return s
}

func (s *AirportSearchResponseBody) SetMessage(v string) *AirportSearchResponseBody {
	s.Message = &v
	return s
}

func (s *AirportSearchResponseBody) SetModule(v *AirportSearchResponseBodyModule) *AirportSearchResponseBody {
	s.Module = v
	return s
}

func (s *AirportSearchResponseBody) SetRequestId(v string) *AirportSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *AirportSearchResponseBody) SetSuccess(v bool) *AirportSearchResponseBody {
	s.Success = &v
	return s
}

func (s *AirportSearchResponseBody) SetTraceId(v string) *AirportSearchResponseBody {
	s.TraceId = &v
	return s
}

type AirportSearchResponseBodyModule struct {
	Cities []*AirportSearchResponseBodyModuleCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
	Nearby *bool                                    `json:"nearby,omitempty" xml:"nearby,omitempty"`
}

func (s AirportSearchResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s AirportSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AirportSearchResponseBodyModule) SetCities(v []*AirportSearchResponseBodyModuleCities) *AirportSearchResponseBodyModule {
	s.Cities = v
	return s
}

func (s *AirportSearchResponseBodyModule) SetNearby(v bool) *AirportSearchResponseBodyModule {
	s.Nearby = &v
	return s
}

type AirportSearchResponseBodyModuleCities struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Distance   *int32  `json:"distance,omitempty" xml:"distance,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	TravelName *string `json:"travel_name,omitempty" xml:"travel_name,omitempty"`
}

func (s AirportSearchResponseBodyModuleCities) String() string {
	return tea.Prettify(s)
}

func (s AirportSearchResponseBodyModuleCities) GoString() string {
	return s.String()
}

func (s *AirportSearchResponseBodyModuleCities) SetCode(v string) *AirportSearchResponseBodyModuleCities {
	s.Code = &v
	return s
}

func (s *AirportSearchResponseBodyModuleCities) SetDistance(v int32) *AirportSearchResponseBodyModuleCities {
	s.Distance = &v
	return s
}

func (s *AirportSearchResponseBodyModuleCities) SetName(v string) *AirportSearchResponseBodyModuleCities {
	s.Name = &v
	return s
}

func (s *AirportSearchResponseBodyModuleCities) SetTravelName(v string) *AirportSearchResponseBodyModuleCities {
	s.TravelName = &v
	return s
}

type AirportSearchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AirportSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AirportSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s AirportSearchResponse) GoString() string {
	return s.String()
}

func (s *AirportSearchResponse) SetHeaders(v map[string]*string) *AirportSearchResponse {
	s.Headers = v
	return s
}

func (s *AirportSearchResponse) SetStatusCode(v int32) *AirportSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *AirportSearchResponse) SetBody(v *AirportSearchResponseBody) *AirportSearchResponse {
	s.Body = v
	return s
}

type ApplyAddHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddHeaders) GoString() string {
	return s.String()
}

func (s *ApplyAddHeaders) SetCommonHeaders(v map[string]*string) *ApplyAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyAddHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ApplyAddRequest struct {
	Budget                    *int64                                   `json:"budget,omitempty" xml:"budget,omitempty"`
	BudgetMerge               *int32                                   `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CorpName                  *string                                  `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId                  *string                                  `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName                *string                                  `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExternalTravelerList      []*ApplyAddRequestExternalTravelerList   `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty" type:"Repeated"`
	ExternalTravelerStandard  *ApplyAddRequestExternalTravelerStandard `json:"external_traveler_standard,omitempty" xml:"external_traveler_standard,omitempty" type:"Struct"`
	FlightBudget              *int64                                   `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	HotelBudget               *int64                                   `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShare                *ApplyAddRequestHotelShare               `json:"hotel_share,omitempty" xml:"hotel_share,omitempty" type:"Struct"`
	InternationalFlightCabins *string                                  `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	ItineraryList             []*ApplyAddRequestItineraryList          `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	ItineraryRule             *int32                                   `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetList          []*ApplyAddRequestItinerarySetList       `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty" type:"Repeated"`
	LimitTraveler             *int32                                   `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	Status                    *int32                                   `json:"status,omitempty" xml:"status,omitempty"`
	ThirdpartApplyId          *string                                  `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId       *string                                  `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	TogetherBookRule          *int32                                   `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	TrainBudget               *int64                                   `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerList              []*ApplyAddRequestTravelerList           `json:"traveler_list,omitempty" xml:"traveler_list,omitempty" type:"Repeated"`
	TravelerStandard          []*ApplyAddRequestTravelerStandard       `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty" type:"Repeated"`
	TripCause                 *string                                  `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripDay                   *int32                                   `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	TripTitle                 *string                                  `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	Type                      *int32                                   `json:"type,omitempty" xml:"type,omitempty"`
	UnionNo                   *string                                  `json:"union_no,omitempty" xml:"union_no,omitempty"`
	UserId                    *string                                  `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName                  *string                                  `json:"user_name,omitempty" xml:"user_name,omitempty"`
	VehicleBudget             *int64                                   `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyAddRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequest) GoString() string {
	return s.String()
}

func (s *ApplyAddRequest) SetBudget(v int64) *ApplyAddRequest {
	s.Budget = &v
	return s
}

func (s *ApplyAddRequest) SetBudgetMerge(v int32) *ApplyAddRequest {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyAddRequest) SetCorpName(v string) *ApplyAddRequest {
	s.CorpName = &v
	return s
}

func (s *ApplyAddRequest) SetDepartId(v string) *ApplyAddRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyAddRequest) SetDepartName(v string) *ApplyAddRequest {
	s.DepartName = &v
	return s
}

func (s *ApplyAddRequest) SetExternalTravelerList(v []*ApplyAddRequestExternalTravelerList) *ApplyAddRequest {
	s.ExternalTravelerList = v
	return s
}

func (s *ApplyAddRequest) SetExternalTravelerStandard(v *ApplyAddRequestExternalTravelerStandard) *ApplyAddRequest {
	s.ExternalTravelerStandard = v
	return s
}

func (s *ApplyAddRequest) SetFlightBudget(v int64) *ApplyAddRequest {
	s.FlightBudget = &v
	return s
}

func (s *ApplyAddRequest) SetHotelBudget(v int64) *ApplyAddRequest {
	s.HotelBudget = &v
	return s
}

func (s *ApplyAddRequest) SetHotelShare(v *ApplyAddRequestHotelShare) *ApplyAddRequest {
	s.HotelShare = v
	return s
}

func (s *ApplyAddRequest) SetInternationalFlightCabins(v string) *ApplyAddRequest {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyAddRequest) SetItineraryList(v []*ApplyAddRequestItineraryList) *ApplyAddRequest {
	s.ItineraryList = v
	return s
}

func (s *ApplyAddRequest) SetItineraryRule(v int32) *ApplyAddRequest {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyAddRequest) SetItinerarySetList(v []*ApplyAddRequestItinerarySetList) *ApplyAddRequest {
	s.ItinerarySetList = v
	return s
}

func (s *ApplyAddRequest) SetLimitTraveler(v int32) *ApplyAddRequest {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyAddRequest) SetStatus(v int32) *ApplyAddRequest {
	s.Status = &v
	return s
}

func (s *ApplyAddRequest) SetThirdpartApplyId(v string) *ApplyAddRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyAddRequest) SetThirdpartBusinessId(v string) *ApplyAddRequest {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyAddRequest) SetTogetherBookRule(v int32) *ApplyAddRequest {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyAddRequest) SetTrainBudget(v int64) *ApplyAddRequest {
	s.TrainBudget = &v
	return s
}

func (s *ApplyAddRequest) SetTravelerList(v []*ApplyAddRequestTravelerList) *ApplyAddRequest {
	s.TravelerList = v
	return s
}

func (s *ApplyAddRequest) SetTravelerStandard(v []*ApplyAddRequestTravelerStandard) *ApplyAddRequest {
	s.TravelerStandard = v
	return s
}

func (s *ApplyAddRequest) SetTripCause(v string) *ApplyAddRequest {
	s.TripCause = &v
	return s
}

func (s *ApplyAddRequest) SetTripDay(v int32) *ApplyAddRequest {
	s.TripDay = &v
	return s
}

func (s *ApplyAddRequest) SetTripTitle(v string) *ApplyAddRequest {
	s.TripTitle = &v
	return s
}

func (s *ApplyAddRequest) SetType(v int32) *ApplyAddRequest {
	s.Type = &v
	return s
}

func (s *ApplyAddRequest) SetUnionNo(v string) *ApplyAddRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyAddRequest) SetUserId(v string) *ApplyAddRequest {
	s.UserId = &v
	return s
}

func (s *ApplyAddRequest) SetUserName(v string) *ApplyAddRequest {
	s.UserName = &v
	return s
}

func (s *ApplyAddRequest) SetVehicleBudget(v int64) *ApplyAddRequest {
	s.VehicleBudget = &v
	return s
}

type ApplyAddRequestExternalTravelerList struct {
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyAddRequestExternalTravelerList) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestExternalTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestExternalTravelerList) SetUserName(v string) *ApplyAddRequestExternalTravelerList {
	s.UserName = &v
	return s
}

type ApplyAddRequestExternalTravelerStandard struct {
	BusinessDiscount *int32                                               `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	EconomyDiscount  *int32                                               `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount    *int32                                               `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins     *string                                              `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	HotelCitys       []*ApplyAddRequestExternalTravelerStandardHotelCitys `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	ReserveType      *int32                                               `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainSeats       *string                                              `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
}

func (s ApplyAddRequestExternalTravelerStandard) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestExternalTravelerStandard) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestExternalTravelerStandard) SetBusinessDiscount(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetEconomyDiscount(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetFirstDiscount(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetFlightCabins(v string) *ApplyAddRequestExternalTravelerStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetHotelCitys(v []*ApplyAddRequestExternalTravelerStandardHotelCitys) *ApplyAddRequestExternalTravelerStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetReserveType(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetTrainSeats(v string) *ApplyAddRequestExternalTravelerStandard {
	s.TrainSeats = &v
	return s
}

type ApplyAddRequestExternalTravelerStandardHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyAddRequestExternalTravelerStandardHotelCitys) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestExternalTravelerStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) SetCityCode(v string) *ApplyAddRequestExternalTravelerStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) SetCityName(v string) *ApplyAddRequestExternalTravelerStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) SetFee(v int64) *ApplyAddRequestExternalTravelerStandardHotelCitys {
	s.Fee = &v
	return s
}

type ApplyAddRequestHotelShare struct {
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyAddRequestHotelShare) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestHotelShare) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestHotelShare) SetParam(v string) *ApplyAddRequestHotelShare {
	s.Param = &v
	return s
}

func (s *ApplyAddRequestHotelShare) SetType(v string) *ApplyAddRequestHotelShare {
	s.Type = &v
	return s
}

type ApplyAddRequestItineraryList struct {
	ArrCity               *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode           *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrDate               *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CostCenterId          *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	DepCity               *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode           *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepDate               *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceId             *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	ItineraryId           *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	NeedHotel             *bool   `json:"need_hotel,omitempty" xml:"need_hotel,omitempty"`
	NeedTraffic           *bool   `json:"need_traffic,omitempty" xml:"need_traffic,omitempty"`
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	TrafficType           *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	TripWay               *int32  `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s ApplyAddRequestItineraryList) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestItineraryList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestItineraryList) SetArrCity(v string) *ApplyAddRequestItineraryList {
	s.ArrCity = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetArrCityCode(v string) *ApplyAddRequestItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetArrDate(v string) *ApplyAddRequestItineraryList {
	s.ArrDate = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetCostCenterId(v int64) *ApplyAddRequestItineraryList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetDepCity(v string) *ApplyAddRequestItineraryList {
	s.DepCity = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetDepCityCode(v string) *ApplyAddRequestItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetDepDate(v string) *ApplyAddRequestItineraryList {
	s.DepDate = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetInvoiceId(v int64) *ApplyAddRequestItineraryList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetItineraryId(v string) *ApplyAddRequestItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetNeedHotel(v bool) *ApplyAddRequestItineraryList {
	s.NeedHotel = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetNeedTraffic(v bool) *ApplyAddRequestItineraryList {
	s.NeedTraffic = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetProjectCode(v string) *ApplyAddRequestItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetProjectTitle(v string) *ApplyAddRequestItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetThirdPartInvoiceId(v string) *ApplyAddRequestItineraryList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetThirdpartCostCenterId(v string) *ApplyAddRequestItineraryList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetTrafficType(v int32) *ApplyAddRequestItineraryList {
	s.TrafficType = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetTripWay(v int32) *ApplyAddRequestItineraryList {
	s.TripWay = &v
	return s
}

type ApplyAddRequestItinerarySetList struct {
	ArrDate               *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CityCodeSet           *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	CitySet               *string `json:"city_set,omitempty" xml:"city_set,omitempty"`
	CostCenterId          *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	DepDate               *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceId             *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	ItineraryId           *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	TrafficType           *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s ApplyAddRequestItinerarySetList) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestItinerarySetList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestItinerarySetList) SetArrDate(v string) *ApplyAddRequestItinerarySetList {
	s.ArrDate = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetCityCodeSet(v string) *ApplyAddRequestItinerarySetList {
	s.CityCodeSet = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetCitySet(v string) *ApplyAddRequestItinerarySetList {
	s.CitySet = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetCostCenterId(v int64) *ApplyAddRequestItinerarySetList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetDepDate(v string) *ApplyAddRequestItinerarySetList {
	s.DepDate = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetInvoiceId(v int64) *ApplyAddRequestItinerarySetList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetItineraryId(v string) *ApplyAddRequestItinerarySetList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetProjectCode(v string) *ApplyAddRequestItinerarySetList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetProjectTitle(v string) *ApplyAddRequestItinerarySetList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetThirdPartInvoiceId(v string) *ApplyAddRequestItinerarySetList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetThirdpartCostCenterId(v string) *ApplyAddRequestItinerarySetList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetTrafficType(v int32) *ApplyAddRequestItinerarySetList {
	s.TrafficType = &v
	return s
}

type ApplyAddRequestTravelerList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyAddRequestTravelerList) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerList) SetUserId(v string) *ApplyAddRequestTravelerList {
	s.UserId = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetUserName(v string) *ApplyAddRequestTravelerList {
	s.UserName = &v
	return s
}

type ApplyAddRequestTravelerStandard struct {
	BusinessDiscount *int32                                       `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	EconomyDiscount  *int32                                       `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount    *int32                                       `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins     *string                                      `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	HotelCitys       []*ApplyAddRequestTravelerStandardHotelCitys `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	ReserveType      *int32                                       `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainSeats       *string                                      `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	UserId           *string                                      `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ApplyAddRequestTravelerStandard) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestTravelerStandard) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandard) SetBusinessDiscount(v int32) *ApplyAddRequestTravelerStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetEconomyDiscount(v int32) *ApplyAddRequestTravelerStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetFirstDiscount(v int32) *ApplyAddRequestTravelerStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetFlightCabins(v string) *ApplyAddRequestTravelerStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetHotelCitys(v []*ApplyAddRequestTravelerStandardHotelCitys) *ApplyAddRequestTravelerStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetReserveType(v int32) *ApplyAddRequestTravelerStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetTrainSeats(v string) *ApplyAddRequestTravelerStandard {
	s.TrainSeats = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetUserId(v string) *ApplyAddRequestTravelerStandard {
	s.UserId = &v
	return s
}

type ApplyAddRequestTravelerStandardHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyAddRequestTravelerStandardHotelCitys) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) SetCityCode(v string) *ApplyAddRequestTravelerStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) SetCityName(v string) *ApplyAddRequestTravelerStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) SetFee(v int64) *ApplyAddRequestTravelerStandardHotelCitys {
	s.Fee = &v
	return s
}

type ApplyAddShrinkRequest struct {
	Budget                         *int64  `json:"budget,omitempty" xml:"budget,omitempty"`
	BudgetMerge                    *int32  `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CorpName                       *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId                       *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName                     *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExternalTravelerListShrink     *string `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty"`
	ExternalTravelerStandardShrink *string `json:"external_traveler_standard,omitempty" xml:"external_traveler_standard,omitempty"`
	FlightBudget                   *int64  `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	HotelBudget                    *int64  `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShareShrink               *string `json:"hotel_share,omitempty" xml:"hotel_share,omitempty"`
	InternationalFlightCabins      *string `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	ItineraryListShrink            *string `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty"`
	ItineraryRule                  *int32  `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetListShrink         *string `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty"`
	LimitTraveler                  *int32  `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	Status                         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdpartApplyId               *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId            *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	TogetherBookRule               *int32  `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	TrainBudget                    *int64  `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerListShrink             *string `json:"traveler_list,omitempty" xml:"traveler_list,omitempty"`
	TravelerStandardShrink         *string `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty"`
	TripCause                      *string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripDay                        *int32  `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	TripTitle                      *string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	Type                           *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UnionNo                        *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	UserId                         *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName                       *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	VehicleBudget                  *int64  `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyAddShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyAddShrinkRequest) SetBudget(v int64) *ApplyAddShrinkRequest {
	s.Budget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetBudgetMerge(v int32) *ApplyAddShrinkRequest {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetCorpName(v string) *ApplyAddShrinkRequest {
	s.CorpName = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetDepartId(v string) *ApplyAddShrinkRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetDepartName(v string) *ApplyAddShrinkRequest {
	s.DepartName = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetExternalTravelerListShrink(v string) *ApplyAddShrinkRequest {
	s.ExternalTravelerListShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetExternalTravelerStandardShrink(v string) *ApplyAddShrinkRequest {
	s.ExternalTravelerStandardShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetFlightBudget(v int64) *ApplyAddShrinkRequest {
	s.FlightBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetHotelBudget(v int64) *ApplyAddShrinkRequest {
	s.HotelBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetHotelShareShrink(v string) *ApplyAddShrinkRequest {
	s.HotelShareShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetInternationalFlightCabins(v string) *ApplyAddShrinkRequest {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetItineraryListShrink(v string) *ApplyAddShrinkRequest {
	s.ItineraryListShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetItineraryRule(v int32) *ApplyAddShrinkRequest {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetItinerarySetListShrink(v string) *ApplyAddShrinkRequest {
	s.ItinerarySetListShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetLimitTraveler(v int32) *ApplyAddShrinkRequest {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetStatus(v int32) *ApplyAddShrinkRequest {
	s.Status = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetThirdpartApplyId(v string) *ApplyAddShrinkRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetThirdpartBusinessId(v string) *ApplyAddShrinkRequest {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTogetherBookRule(v int32) *ApplyAddShrinkRequest {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTrainBudget(v int64) *ApplyAddShrinkRequest {
	s.TrainBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTravelerListShrink(v string) *ApplyAddShrinkRequest {
	s.TravelerListShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTravelerStandardShrink(v string) *ApplyAddShrinkRequest {
	s.TravelerStandardShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTripCause(v string) *ApplyAddShrinkRequest {
	s.TripCause = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTripDay(v int32) *ApplyAddShrinkRequest {
	s.TripDay = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTripTitle(v string) *ApplyAddShrinkRequest {
	s.TripTitle = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetType(v int32) *ApplyAddShrinkRequest {
	s.Type = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetUnionNo(v string) *ApplyAddShrinkRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetUserId(v string) *ApplyAddShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetUserName(v string) *ApplyAddShrinkRequest {
	s.UserName = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetVehicleBudget(v int64) *ApplyAddShrinkRequest {
	s.VehicleBudget = &v
	return s
}

type ApplyAddResponseBody struct {
	Code      *int32                      `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                     `json:"message,omitempty" xml:"message,omitempty"`
	Module    *ApplyAddResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                       `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                     `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAddResponseBody) SetCode(v int32) *ApplyAddResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyAddResponseBody) SetMessage(v string) *ApplyAddResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyAddResponseBody) SetModule(v *ApplyAddResponseBodyModule) *ApplyAddResponseBody {
	s.Module = v
	return s
}

func (s *ApplyAddResponseBody) SetRequestId(v string) *ApplyAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyAddResponseBody) SetSuccess(v bool) *ApplyAddResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyAddResponseBody) SetTraceId(v string) *ApplyAddResponseBody {
	s.TraceId = &v
	return s
}

type ApplyAddResponseBodyModule struct {
	ApplyId             *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
}

func (s ApplyAddResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ApplyAddResponseBodyModule) SetApplyId(v int64) *ApplyAddResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *ApplyAddResponseBodyModule) SetThirdpartApplyId(v string) *ApplyAddResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyAddResponseBodyModule) SetThirdpartBusinessId(v string) *ApplyAddResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

type ApplyAddResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyAddResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyAddResponse) GoString() string {
	return s.String()
}

func (s *ApplyAddResponse) SetHeaders(v map[string]*string) *ApplyAddResponse {
	s.Headers = v
	return s
}

func (s *ApplyAddResponse) SetStatusCode(v int32) *ApplyAddResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAddResponse) SetBody(v *ApplyAddResponseBody) *ApplyAddResponse {
	s.Body = v
	return s
}

type ApplyApproveHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyApproveHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyApproveHeaders) GoString() string {
	return s.String()
}

func (s *ApplyApproveHeaders) SetCommonHeaders(v map[string]*string) *ApplyApproveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyApproveHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyApproveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ApplyApproveRequest struct {
	ApplyId     *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	Note        *string `json:"note,omitempty" xml:"note,omitempty"`
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	Status      *int32  `json:"status,omitempty" xml:"status,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName    *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyApproveRequest) GoString() string {
	return s.String()
}

func (s *ApplyApproveRequest) SetApplyId(v string) *ApplyApproveRequest {
	s.ApplyId = &v
	return s
}

func (s *ApplyApproveRequest) SetNote(v string) *ApplyApproveRequest {
	s.Note = &v
	return s
}

func (s *ApplyApproveRequest) SetOperateTime(v string) *ApplyApproveRequest {
	s.OperateTime = &v
	return s
}

func (s *ApplyApproveRequest) SetStatus(v int32) *ApplyApproveRequest {
	s.Status = &v
	return s
}

func (s *ApplyApproveRequest) SetUserId(v string) *ApplyApproveRequest {
	s.UserId = &v
	return s
}

func (s *ApplyApproveRequest) SetUserName(v string) *ApplyApproveRequest {
	s.UserName = &v
	return s
}

type ApplyApproveResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *string `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
}

func (s ApplyApproveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyApproveResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyApproveResponseBody) SetCode(v int32) *ApplyApproveResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyApproveResponseBody) SetMessage(v string) *ApplyApproveResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyApproveResponseBody) SetModule(v string) *ApplyApproveResponseBody {
	s.Module = &v
	return s
}

func (s *ApplyApproveResponseBody) SetRequestId(v string) *ApplyApproveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyApproveResponseBody) SetSuccess(v bool) *ApplyApproveResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyApproveResponseBody) SetTraceId(v string) *ApplyApproveResponseBody {
	s.TraceId = &v
	return s
}

type ApplyApproveResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyApproveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyApproveResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyApproveResponse) GoString() string {
	return s.String()
}

func (s *ApplyApproveResponse) SetHeaders(v map[string]*string) *ApplyApproveResponse {
	s.Headers = v
	return s
}

func (s *ApplyApproveResponse) SetStatusCode(v int32) *ApplyApproveResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyApproveResponse) SetBody(v *ApplyApproveResponseBody) *ApplyApproveResponse {
	s.Body = v
	return s
}

type ApplyListQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyListQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryHeaders) GoString() string {
	return s.String()
}

func (s *ApplyListQueryHeaders) SetCommonHeaders(v map[string]*string) *ApplyListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ApplyListQueryRequest struct {
	AllApply         *bool   `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	DepartId         *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	GmtModified      *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	OnlyShangLvApply *bool   `json:"only_shang_lv_apply,omitempty" xml:"only_shang_lv_apply,omitempty"`
	Page             *int32  `json:"page,omitempty" xml:"page,omitempty"`
	PageSize         *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Type             *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UnionNo          *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ApplyListQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryRequest) GoString() string {
	return s.String()
}

func (s *ApplyListQueryRequest) SetAllApply(v bool) *ApplyListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *ApplyListQueryRequest) SetDepartId(v string) *ApplyListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyListQueryRequest) SetEndTime(v string) *ApplyListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *ApplyListQueryRequest) SetGmtModified(v string) *ApplyListQueryRequest {
	s.GmtModified = &v
	return s
}

func (s *ApplyListQueryRequest) SetOnlyShangLvApply(v bool) *ApplyListQueryRequest {
	s.OnlyShangLvApply = &v
	return s
}

func (s *ApplyListQueryRequest) SetPage(v int32) *ApplyListQueryRequest {
	s.Page = &v
	return s
}

func (s *ApplyListQueryRequest) SetPageSize(v int32) *ApplyListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *ApplyListQueryRequest) SetStartTime(v string) *ApplyListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *ApplyListQueryRequest) SetType(v int32) *ApplyListQueryRequest {
	s.Type = &v
	return s
}

func (s *ApplyListQueryRequest) SetUnionNo(v string) *ApplyListQueryRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyListQueryRequest) SetUserId(v string) *ApplyListQueryRequest {
	s.UserId = &v
	return s
}

type ApplyListQueryResponseBody struct {
	Code       *int32                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message    *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	ModuleList []*ApplyListQueryResponseBodyModuleList `json:"module_list,omitempty" xml:"module_list,omitempty" type:"Repeated"`
	RequestId  *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success    *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId    *string                                 `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyListQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBody) SetCode(v int32) *ApplyListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyListQueryResponseBody) SetMessage(v string) *ApplyListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyListQueryResponseBody) SetModuleList(v []*ApplyListQueryResponseBodyModuleList) *ApplyListQueryResponseBody {
	s.ModuleList = v
	return s
}

func (s *ApplyListQueryResponseBody) SetRequestId(v string) *ApplyListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyListQueryResponseBody) SetSuccess(v bool) *ApplyListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyListQueryResponseBody) SetTraceId(v string) *ApplyListQueryResponseBody {
	s.TraceId = &v
	return s
}

type ApplyListQueryResponseBodyModuleList struct {
	ApplyShowId          *string                                                     `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	ApproverList         []*ApplyListQueryResponseBodyModuleListApproverList         `json:"approver_list,omitempty" xml:"approver_list,omitempty" type:"Repeated"`
	CorpId               *string                                                     `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName             *string                                                     `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId             *string                                                     `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName           *string                                                     `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExternalTravelerList []*ApplyListQueryResponseBodyModuleListExternalTravelerList `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty" type:"Repeated"`
	FlowCode             *string                                                     `json:"flow_code,omitempty" xml:"flow_code,omitempty"`
	GmtCreate            *string                                                     `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModified          *string                                                     `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	Id                   *int64                                                      `json:"id,omitempty" xml:"id,omitempty"`
	ItineraryList        []*ApplyListQueryResponseBodyModuleListItineraryList        `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	ItineraryRule        *int32                                                      `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetList     []*ApplyListQueryResponseBodyModuleListItinerarySetList     `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty" type:"Repeated"`
	Status               *int32                                                      `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc           *string                                                     `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	ThirdpartBusinessId  *string                                                     `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	ThirdpartId          *string                                                     `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	TravelerList         []*ApplyListQueryResponseBodyModuleListTravelerList         `json:"traveler_list,omitempty" xml:"traveler_list,omitempty" type:"Repeated"`
	TripCause            *string                                                     `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripDay              *int32                                                      `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	TripTitle            *string                                                     `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	Type                 *int32                                                      `json:"type,omitempty" xml:"type,omitempty"`
	UnionNo              *string                                                     `json:"union_no,omitempty" xml:"union_no,omitempty"`
	UserId               *string                                                     `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName             *string                                                     `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleList) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleList) SetApplyShowId(v string) *ApplyListQueryResponseBodyModuleList {
	s.ApplyShowId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetApproverList(v []*ApplyListQueryResponseBodyModuleListApproverList) *ApplyListQueryResponseBodyModuleList {
	s.ApproverList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetCorpId(v string) *ApplyListQueryResponseBodyModuleList {
	s.CorpId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetCorpName(v string) *ApplyListQueryResponseBodyModuleList {
	s.CorpName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetDepartId(v string) *ApplyListQueryResponseBodyModuleList {
	s.DepartId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetDepartName(v string) *ApplyListQueryResponseBodyModuleList {
	s.DepartName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetExternalTravelerList(v []*ApplyListQueryResponseBodyModuleListExternalTravelerList) *ApplyListQueryResponseBodyModuleList {
	s.ExternalTravelerList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetFlowCode(v string) *ApplyListQueryResponseBodyModuleList {
	s.FlowCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetGmtCreate(v string) *ApplyListQueryResponseBodyModuleList {
	s.GmtCreate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetGmtModified(v string) *ApplyListQueryResponseBodyModuleList {
	s.GmtModified = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetId(v int64) *ApplyListQueryResponseBodyModuleList {
	s.Id = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetItineraryList(v []*ApplyListQueryResponseBodyModuleListItineraryList) *ApplyListQueryResponseBodyModuleList {
	s.ItineraryList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetItineraryRule(v int32) *ApplyListQueryResponseBodyModuleList {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetItinerarySetList(v []*ApplyListQueryResponseBodyModuleListItinerarySetList) *ApplyListQueryResponseBodyModuleList {
	s.ItinerarySetList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetStatus(v int32) *ApplyListQueryResponseBodyModuleList {
	s.Status = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetStatusDesc(v string) *ApplyListQueryResponseBodyModuleList {
	s.StatusDesc = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetThirdpartBusinessId(v string) *ApplyListQueryResponseBodyModuleList {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetThirdpartId(v string) *ApplyListQueryResponseBodyModuleList {
	s.ThirdpartId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetTravelerList(v []*ApplyListQueryResponseBodyModuleListTravelerList) *ApplyListQueryResponseBodyModuleList {
	s.TravelerList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetTripCause(v string) *ApplyListQueryResponseBodyModuleList {
	s.TripCause = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetTripDay(v int32) *ApplyListQueryResponseBodyModuleList {
	s.TripDay = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetTripTitle(v string) *ApplyListQueryResponseBodyModuleList {
	s.TripTitle = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetType(v int32) *ApplyListQueryResponseBodyModuleList {
	s.Type = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetUnionNo(v string) *ApplyListQueryResponseBodyModuleList {
	s.UnionNo = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetUserId(v string) *ApplyListQueryResponseBodyModuleList {
	s.UserId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetUserName(v string) *ApplyListQueryResponseBodyModuleList {
	s.UserName = &v
	return s
}

type ApplyListQueryResponseBodyModuleListApproverList struct {
	Note        *string `json:"note,omitempty" xml:"note,omitempty"`
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	Order       *int32  `json:"order,omitempty" xml:"order,omitempty"`
	Status      *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc  *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName    *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListApproverList) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListApproverList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetNote(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.Note = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetOperateTime(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.OperateTime = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetOrder(v int32) *ApplyListQueryResponseBodyModuleListApproverList {
	s.Order = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetStatus(v int32) *ApplyListQueryResponseBodyModuleListApproverList {
	s.Status = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetStatusDesc(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.StatusDesc = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetUserId(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.UserId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetUserName(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.UserName = &v
	return s
}

type ApplyListQueryResponseBodyModuleListExternalTravelerList struct {
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListExternalTravelerList) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListExternalTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetUserName(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.UserName = &v
	return s
}

type ApplyListQueryResponseBodyModuleListItineraryList struct {
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrDate        *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepDate        *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceName    *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	ItineraryId    *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ProjectCode    *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle   *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	TrafficType    *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	TripWay        *int32  `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListItineraryList) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListItineraryList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetArrCity(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ArrCity = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetArrDate(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ArrDate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetCostCenterName(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetDepCity(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.DepCity = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetDepDate(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.DepDate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetInvoiceName(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetItineraryId(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetProjectCode(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetProjectTitle(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetTrafficType(v int32) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.TrafficType = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetTripWay(v int32) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.TripWay = &v
	return s
}

type ApplyListQueryResponseBodyModuleListItinerarySetList struct {
	ArrDate        *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CityCodeSet    *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	CitySet        *string `json:"city_set,omitempty" xml:"city_set,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepDate        *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceName    *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	ItineraryId    *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ProjectCode    *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle   *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	TrafficType    *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListItinerarySetList) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListItinerarySetList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetArrDate(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ArrDate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetCityCodeSet(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.CityCodeSet = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetCitySet(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.CitySet = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetCostCenterName(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetDepDate(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.DepDate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetInvoiceName(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetItineraryId(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetProjectCode(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetProjectTitle(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetTrafficType(v int32) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.TrafficType = &v
	return s
}

type ApplyListQueryResponseBodyModuleListTravelerList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListTravelerList) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetUserId(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.UserId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetUserName(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.UserName = &v
	return s
}

type ApplyListQueryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyListQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyListQueryResponse) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponse) SetHeaders(v map[string]*string) *ApplyListQueryResponse {
	s.Headers = v
	return s
}

func (s *ApplyListQueryResponse) SetStatusCode(v int32) *ApplyListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyListQueryResponse) SetBody(v *ApplyListQueryResponseBody) *ApplyListQueryResponse {
	s.Body = v
	return s
}

type ApplyModifyHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyModifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyHeaders) GoString() string {
	return s.String()
}

func (s *ApplyModifyHeaders) SetCommonHeaders(v map[string]*string) *ApplyModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyModifyHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ApplyModifyRequest struct {
	Budget                   *int64                                      `json:"budget,omitempty" xml:"budget,omitempty"`
	BudgetMerge              *int32                                      `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CorpName                 *string                                     `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId                 *string                                     `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName               *string                                     `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExternalTravelerList     []*ApplyModifyRequestExternalTravelerList   `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty" type:"Repeated"`
	ExternalTravelerStandard *ApplyModifyRequestExternalTravelerStandard `json:"external_traveler_standard,omitempty" xml:"external_traveler_standard,omitempty" type:"Struct"`
	FlightBudget             *int64                                      `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	HotelBudget              *int64                                      `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShare               *ApplyModifyRequestHotelShare               `json:"hotel_share,omitempty" xml:"hotel_share,omitempty" type:"Struct"`
	ItineraryList            []*ApplyModifyRequestItineraryList          `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	ItineraryRule            *int32                                      `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetList         []*ApplyModifyRequestItinerarySetList       `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty" type:"Repeated"`
	LimitTraveler            *int32                                      `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	Status                   *int32                                      `json:"status,omitempty" xml:"status,omitempty"`
	ThirdpartApplyId         *string                                     `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId      *string                                     `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	TogetherBookRule         *int32                                      `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	TrainBudget              *int64                                      `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerList             []*ApplyModifyRequestTravelerList           `json:"traveler_list,omitempty" xml:"traveler_list,omitempty" type:"Repeated"`
	TravelerStandard         []*ApplyModifyRequestTravelerStandard       `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty" type:"Repeated"`
	TripCause                *string                                     `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripDay                  *int32                                      `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	TripTitle                *string                                     `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	UnionNo                  *string                                     `json:"union_no,omitempty" xml:"union_no,omitempty"`
	UserId                   *string                                     `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName                 *string                                     `json:"user_name,omitempty" xml:"user_name,omitempty"`
	VehicleBudget            *int64                                      `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequest) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequest) SetBudget(v int64) *ApplyModifyRequest {
	s.Budget = &v
	return s
}

func (s *ApplyModifyRequest) SetBudgetMerge(v int32) *ApplyModifyRequest {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyModifyRequest) SetCorpName(v string) *ApplyModifyRequest {
	s.CorpName = &v
	return s
}

func (s *ApplyModifyRequest) SetDepartId(v string) *ApplyModifyRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyModifyRequest) SetDepartName(v string) *ApplyModifyRequest {
	s.DepartName = &v
	return s
}

func (s *ApplyModifyRequest) SetExternalTravelerList(v []*ApplyModifyRequestExternalTravelerList) *ApplyModifyRequest {
	s.ExternalTravelerList = v
	return s
}

func (s *ApplyModifyRequest) SetExternalTravelerStandard(v *ApplyModifyRequestExternalTravelerStandard) *ApplyModifyRequest {
	s.ExternalTravelerStandard = v
	return s
}

func (s *ApplyModifyRequest) SetFlightBudget(v int64) *ApplyModifyRequest {
	s.FlightBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetHotelBudget(v int64) *ApplyModifyRequest {
	s.HotelBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetHotelShare(v *ApplyModifyRequestHotelShare) *ApplyModifyRequest {
	s.HotelShare = v
	return s
}

func (s *ApplyModifyRequest) SetItineraryList(v []*ApplyModifyRequestItineraryList) *ApplyModifyRequest {
	s.ItineraryList = v
	return s
}

func (s *ApplyModifyRequest) SetItineraryRule(v int32) *ApplyModifyRequest {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyModifyRequest) SetItinerarySetList(v []*ApplyModifyRequestItinerarySetList) *ApplyModifyRequest {
	s.ItinerarySetList = v
	return s
}

func (s *ApplyModifyRequest) SetLimitTraveler(v int32) *ApplyModifyRequest {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyModifyRequest) SetStatus(v int32) *ApplyModifyRequest {
	s.Status = &v
	return s
}

func (s *ApplyModifyRequest) SetThirdpartApplyId(v string) *ApplyModifyRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyModifyRequest) SetThirdpartBusinessId(v string) *ApplyModifyRequest {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyModifyRequest) SetTogetherBookRule(v int32) *ApplyModifyRequest {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyModifyRequest) SetTrainBudget(v int64) *ApplyModifyRequest {
	s.TrainBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetTravelerList(v []*ApplyModifyRequestTravelerList) *ApplyModifyRequest {
	s.TravelerList = v
	return s
}

func (s *ApplyModifyRequest) SetTravelerStandard(v []*ApplyModifyRequestTravelerStandard) *ApplyModifyRequest {
	s.TravelerStandard = v
	return s
}

func (s *ApplyModifyRequest) SetTripCause(v string) *ApplyModifyRequest {
	s.TripCause = &v
	return s
}

func (s *ApplyModifyRequest) SetTripDay(v int32) *ApplyModifyRequest {
	s.TripDay = &v
	return s
}

func (s *ApplyModifyRequest) SetTripTitle(v string) *ApplyModifyRequest {
	s.TripTitle = &v
	return s
}

func (s *ApplyModifyRequest) SetUnionNo(v string) *ApplyModifyRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyModifyRequest) SetUserId(v string) *ApplyModifyRequest {
	s.UserId = &v
	return s
}

func (s *ApplyModifyRequest) SetUserName(v string) *ApplyModifyRequest {
	s.UserName = &v
	return s
}

func (s *ApplyModifyRequest) SetVehicleBudget(v int64) *ApplyModifyRequest {
	s.VehicleBudget = &v
	return s
}

type ApplyModifyRequestExternalTravelerList struct {
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyModifyRequestExternalTravelerList) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestExternalTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestExternalTravelerList) SetUserName(v string) *ApplyModifyRequestExternalTravelerList {
	s.UserName = &v
	return s
}

type ApplyModifyRequestExternalTravelerStandard struct {
	BusinessDiscount *int32                                                  `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	EconomyDiscount  *int32                                                  `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount    *int32                                                  `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins     *string                                                 `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	HotelCitys       []*ApplyModifyRequestExternalTravelerStandardHotelCitys `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	ReserveType      *int32                                                  `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainSeats       *string                                                 `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
}

func (s ApplyModifyRequestExternalTravelerStandard) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestExternalTravelerStandard) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetBusinessDiscount(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetEconomyDiscount(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetFirstDiscount(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetFlightCabins(v string) *ApplyModifyRequestExternalTravelerStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetHotelCitys(v []*ApplyModifyRequestExternalTravelerStandardHotelCitys) *ApplyModifyRequestExternalTravelerStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetReserveType(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetTrainSeats(v string) *ApplyModifyRequestExternalTravelerStandard {
	s.TrainSeats = &v
	return s
}

type ApplyModifyRequestExternalTravelerStandardHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyModifyRequestExternalTravelerStandardHotelCitys) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestExternalTravelerStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) SetCityCode(v string) *ApplyModifyRequestExternalTravelerStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) SetCityName(v string) *ApplyModifyRequestExternalTravelerStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) SetFee(v int64) *ApplyModifyRequestExternalTravelerStandardHotelCitys {
	s.Fee = &v
	return s
}

type ApplyModifyRequestHotelShare struct {
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyModifyRequestHotelShare) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestHotelShare) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestHotelShare) SetParam(v string) *ApplyModifyRequestHotelShare {
	s.Param = &v
	return s
}

func (s *ApplyModifyRequestHotelShare) SetType(v string) *ApplyModifyRequestHotelShare {
	s.Type = &v
	return s
}

type ApplyModifyRequestItineraryList struct {
	ArrCity               *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode           *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrDate               *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CostCenterId          *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	DepCity               *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode           *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepDate               *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceId             *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	ItineraryId           *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	NeedHotel             *bool   `json:"need_hotel,omitempty" xml:"need_hotel,omitempty"`
	NeedTraffic           *bool   `json:"need_traffic,omitempty" xml:"need_traffic,omitempty"`
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	TrafficType           *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	TripWay               *int32  `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s ApplyModifyRequestItineraryList) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestItineraryList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestItineraryList) SetArrCity(v string) *ApplyModifyRequestItineraryList {
	s.ArrCity = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetArrCityCode(v string) *ApplyModifyRequestItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetArrDate(v string) *ApplyModifyRequestItineraryList {
	s.ArrDate = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetCostCenterId(v int64) *ApplyModifyRequestItineraryList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetDepCity(v string) *ApplyModifyRequestItineraryList {
	s.DepCity = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetDepCityCode(v string) *ApplyModifyRequestItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetDepDate(v string) *ApplyModifyRequestItineraryList {
	s.DepDate = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetInvoiceId(v int64) *ApplyModifyRequestItineraryList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetItineraryId(v string) *ApplyModifyRequestItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetNeedHotel(v bool) *ApplyModifyRequestItineraryList {
	s.NeedHotel = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetNeedTraffic(v bool) *ApplyModifyRequestItineraryList {
	s.NeedTraffic = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetProjectCode(v string) *ApplyModifyRequestItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetProjectTitle(v string) *ApplyModifyRequestItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetThirdPartInvoiceId(v string) *ApplyModifyRequestItineraryList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetThirdpartCostCenterId(v string) *ApplyModifyRequestItineraryList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetTrafficType(v int32) *ApplyModifyRequestItineraryList {
	s.TrafficType = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetTripWay(v int32) *ApplyModifyRequestItineraryList {
	s.TripWay = &v
	return s
}

type ApplyModifyRequestItinerarySetList struct {
	ArrDate               *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CityCodeSet           *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	CitySet               *string `json:"city_set,omitempty" xml:"city_set,omitempty"`
	CostCenterId          *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	DepDate               *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceId             *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	ItineraryId           *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	TrafficType           *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s ApplyModifyRequestItinerarySetList) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestItinerarySetList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestItinerarySetList) SetArrDate(v string) *ApplyModifyRequestItinerarySetList {
	s.ArrDate = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetCityCodeSet(v string) *ApplyModifyRequestItinerarySetList {
	s.CityCodeSet = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetCitySet(v string) *ApplyModifyRequestItinerarySetList {
	s.CitySet = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetCostCenterId(v int64) *ApplyModifyRequestItinerarySetList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetDepDate(v string) *ApplyModifyRequestItinerarySetList {
	s.DepDate = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetInvoiceId(v int64) *ApplyModifyRequestItinerarySetList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetItineraryId(v string) *ApplyModifyRequestItinerarySetList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetProjectCode(v string) *ApplyModifyRequestItinerarySetList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetProjectTitle(v string) *ApplyModifyRequestItinerarySetList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetThirdPartInvoiceId(v string) *ApplyModifyRequestItinerarySetList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetThirdpartCostCenterId(v string) *ApplyModifyRequestItinerarySetList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetTrafficType(v int32) *ApplyModifyRequestItinerarySetList {
	s.TrafficType = &v
	return s
}

type ApplyModifyRequestTravelerList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyModifyRequestTravelerList) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerList) SetUserId(v string) *ApplyModifyRequestTravelerList {
	s.UserId = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetUserName(v string) *ApplyModifyRequestTravelerList {
	s.UserName = &v
	return s
}

type ApplyModifyRequestTravelerStandard struct {
	BusinessDiscount *int32                                          `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	EconomyDiscount  *int32                                          `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount    *int32                                          `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins     *string                                         `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	HotelCitys       []*ApplyModifyRequestTravelerStandardHotelCitys `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	ReserveType      *int32                                          `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainSeats       *string                                         `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	UserId           *string                                         `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ApplyModifyRequestTravelerStandard) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandard) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandard) SetBusinessDiscount(v int32) *ApplyModifyRequestTravelerStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetEconomyDiscount(v int32) *ApplyModifyRequestTravelerStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetFirstDiscount(v int32) *ApplyModifyRequestTravelerStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetFlightCabins(v string) *ApplyModifyRequestTravelerStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetHotelCitys(v []*ApplyModifyRequestTravelerStandardHotelCitys) *ApplyModifyRequestTravelerStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetReserveType(v int32) *ApplyModifyRequestTravelerStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetTrainSeats(v string) *ApplyModifyRequestTravelerStandard {
	s.TrainSeats = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetUserId(v string) *ApplyModifyRequestTravelerStandard {
	s.UserId = &v
	return s
}

type ApplyModifyRequestTravelerStandardHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardHotelCitys) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) SetCityCode(v string) *ApplyModifyRequestTravelerStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) SetCityName(v string) *ApplyModifyRequestTravelerStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) SetFee(v int64) *ApplyModifyRequestTravelerStandardHotelCitys {
	s.Fee = &v
	return s
}

type ApplyModifyShrinkRequest struct {
	Budget                         *int64  `json:"budget,omitempty" xml:"budget,omitempty"`
	BudgetMerge                    *int32  `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CorpName                       *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId                       *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName                     *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExternalTravelerListShrink     *string `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty"`
	ExternalTravelerStandardShrink *string `json:"external_traveler_standard,omitempty" xml:"external_traveler_standard,omitempty"`
	FlightBudget                   *int64  `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	HotelBudget                    *int64  `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShareShrink               *string `json:"hotel_share,omitempty" xml:"hotel_share,omitempty"`
	ItineraryListShrink            *string `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty"`
	ItineraryRule                  *int32  `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetListShrink         *string `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty"`
	LimitTraveler                  *int32  `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	Status                         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdpartApplyId               *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId            *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	TogetherBookRule               *int32  `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	TrainBudget                    *int64  `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerListShrink             *string `json:"traveler_list,omitempty" xml:"traveler_list,omitempty"`
	TravelerStandardShrink         *string `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty"`
	TripCause                      *string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripDay                        *int32  `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	TripTitle                      *string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	UnionNo                        *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	UserId                         *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName                       *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	VehicleBudget                  *int64  `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyModifyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyModifyShrinkRequest) SetBudget(v int64) *ApplyModifyShrinkRequest {
	s.Budget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetBudgetMerge(v int32) *ApplyModifyShrinkRequest {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetCorpName(v string) *ApplyModifyShrinkRequest {
	s.CorpName = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetDepartId(v string) *ApplyModifyShrinkRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetDepartName(v string) *ApplyModifyShrinkRequest {
	s.DepartName = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetExternalTravelerListShrink(v string) *ApplyModifyShrinkRequest {
	s.ExternalTravelerListShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetExternalTravelerStandardShrink(v string) *ApplyModifyShrinkRequest {
	s.ExternalTravelerStandardShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetFlightBudget(v int64) *ApplyModifyShrinkRequest {
	s.FlightBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetHotelBudget(v int64) *ApplyModifyShrinkRequest {
	s.HotelBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetHotelShareShrink(v string) *ApplyModifyShrinkRequest {
	s.HotelShareShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetItineraryListShrink(v string) *ApplyModifyShrinkRequest {
	s.ItineraryListShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetItineraryRule(v int32) *ApplyModifyShrinkRequest {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetItinerarySetListShrink(v string) *ApplyModifyShrinkRequest {
	s.ItinerarySetListShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetLimitTraveler(v int32) *ApplyModifyShrinkRequest {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetStatus(v int32) *ApplyModifyShrinkRequest {
	s.Status = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetThirdpartApplyId(v string) *ApplyModifyShrinkRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetThirdpartBusinessId(v string) *ApplyModifyShrinkRequest {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTogetherBookRule(v int32) *ApplyModifyShrinkRequest {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTrainBudget(v int64) *ApplyModifyShrinkRequest {
	s.TrainBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTravelerListShrink(v string) *ApplyModifyShrinkRequest {
	s.TravelerListShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTravelerStandardShrink(v string) *ApplyModifyShrinkRequest {
	s.TravelerStandardShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTripCause(v string) *ApplyModifyShrinkRequest {
	s.TripCause = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTripDay(v int32) *ApplyModifyShrinkRequest {
	s.TripDay = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTripTitle(v string) *ApplyModifyShrinkRequest {
	s.TripTitle = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetUnionNo(v string) *ApplyModifyShrinkRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetUserId(v string) *ApplyModifyShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetUserName(v string) *ApplyModifyShrinkRequest {
	s.UserName = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetVehicleBudget(v int64) *ApplyModifyShrinkRequest {
	s.VehicleBudget = &v
	return s
}

type ApplyModifyResponseBody struct {
	Code      *int32                         `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                        `json:"message,omitempty" xml:"message,omitempty"`
	Module    *ApplyModifyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                          `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                        `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyModifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyModifyResponseBody) SetCode(v int32) *ApplyModifyResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyModifyResponseBody) SetMessage(v string) *ApplyModifyResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyModifyResponseBody) SetModule(v *ApplyModifyResponseBodyModule) *ApplyModifyResponseBody {
	s.Module = v
	return s
}

func (s *ApplyModifyResponseBody) SetRequestId(v string) *ApplyModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyModifyResponseBody) SetSuccess(v bool) *ApplyModifyResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyModifyResponseBody) SetTraceId(v string) *ApplyModifyResponseBody {
	s.TraceId = &v
	return s
}

type ApplyModifyResponseBodyModule struct {
	ApplyId             *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
}

func (s ApplyModifyResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ApplyModifyResponseBodyModule) SetApplyId(v int64) *ApplyModifyResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *ApplyModifyResponseBodyModule) SetThirdpartApplyId(v string) *ApplyModifyResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyModifyResponseBodyModule) SetThirdpartBusinessId(v string) *ApplyModifyResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

type ApplyModifyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyModifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyModifyResponse) GoString() string {
	return s.String()
}

func (s *ApplyModifyResponse) SetHeaders(v map[string]*string) *ApplyModifyResponse {
	s.Headers = v
	return s
}

func (s *ApplyModifyResponse) SetStatusCode(v int32) *ApplyModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyModifyResponse) SetBody(v *ApplyModifyResponseBody) *ApplyModifyResponse {
	s.Body = v
	return s
}

type ApplyQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *ApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *ApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ApplyQueryRequest struct {
	ApplyId          *int32  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApplyShowId      *string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	Type             *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *ApplyQueryRequest) SetApplyId(v int32) *ApplyQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *ApplyQueryRequest) SetApplyShowId(v string) *ApplyQueryRequest {
	s.ApplyShowId = &v
	return s
}

func (s *ApplyQueryRequest) SetThirdpartApplyId(v string) *ApplyQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyQueryRequest) SetType(v int32) *ApplyQueryRequest {
	s.Type = &v
	return s
}

type ApplyQueryResponseBody struct {
	Code      *int32                        `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                       `json:"message,omitempty" xml:"message,omitempty"`
	Module    *ApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                         `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                       `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBody) SetCode(v int32) *ApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyQueryResponseBody) SetMessage(v string) *ApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyQueryResponseBody) SetModule(v *ApplyQueryResponseBodyModule) *ApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *ApplyQueryResponseBody) SetRequestId(v string) *ApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyQueryResponseBody) SetSuccess(v bool) *ApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyQueryResponseBody) SetTraceId(v string) *ApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

type ApplyQueryResponseBodyModule struct {
	ApplyShowId          *string                                             `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	ApproverList         []*ApplyQueryResponseBodyModuleApproverList         `json:"approver_list,omitempty" xml:"approver_list,omitempty" type:"Repeated"`
	Budget               *int64                                              `json:"budget,omitempty" xml:"budget,omitempty"`
	BudgetMerge          *int32                                              `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CorpId               *string                                             `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName             *string                                             `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId             *string                                             `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName           *string                                             `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExternalTravelerList []*ApplyQueryResponseBodyModuleExternalTravelerList `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty" type:"Repeated"`
	FlightBudget         *int64                                              `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	GmtCreate            *string                                             `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModified          *string                                             `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	HotelBudget          *int64                                              `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShare           *ApplyQueryResponseBodyModuleHotelShare             `json:"hotel_share,omitempty" xml:"hotel_share,omitempty" type:"Struct"`
	Id                   *int64                                              `json:"id,omitempty" xml:"id,omitempty"`
	ItineraryList        []*ApplyQueryResponseBodyModuleItineraryList        `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	ItineraryRule        *int32                                              `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetList     []*ApplyQueryResponseBodyModuleItinerarySetList     `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty" type:"Repeated"`
	LimitTraveler        *int32                                              `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	Status               *int32                                              `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc           *string                                             `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	ThirdpartBusinessId  *string                                             `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	ThirdpartId          *string                                             `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	TogetherBookRule     *int32                                              `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	TrainBudget          *int64                                              `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerList         []*ApplyQueryResponseBodyModuleTravelerList         `json:"traveler_list,omitempty" xml:"traveler_list,omitempty" type:"Repeated"`
	TripCause            *string                                             `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripDay              *int32                                              `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	TripTitle            *string                                             `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	Type                 *int32                                              `json:"type,omitempty" xml:"type,omitempty"`
	UnionNo              *string                                             `json:"union_no,omitempty" xml:"union_no,omitempty"`
	UserId               *string                                             `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName             *string                                             `json:"user_name,omitempty" xml:"user_name,omitempty"`
	VehicleBudget        *int64                                              `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModule) SetApplyShowId(v string) *ApplyQueryResponseBodyModule {
	s.ApplyShowId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetApproverList(v []*ApplyQueryResponseBodyModuleApproverList) *ApplyQueryResponseBodyModule {
	s.ApproverList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetBudget(v int64) *ApplyQueryResponseBodyModule {
	s.Budget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetBudgetMerge(v int32) *ApplyQueryResponseBodyModule {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetCorpId(v string) *ApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetCorpName(v string) *ApplyQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetDepartId(v string) *ApplyQueryResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetDepartName(v string) *ApplyQueryResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetExternalTravelerList(v []*ApplyQueryResponseBodyModuleExternalTravelerList) *ApplyQueryResponseBodyModule {
	s.ExternalTravelerList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetFlightBudget(v int64) *ApplyQueryResponseBodyModule {
	s.FlightBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetGmtCreate(v string) *ApplyQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetGmtModified(v string) *ApplyQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetHotelBudget(v int64) *ApplyQueryResponseBodyModule {
	s.HotelBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetHotelShare(v *ApplyQueryResponseBodyModuleHotelShare) *ApplyQueryResponseBodyModule {
	s.HotelShare = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetId(v int64) *ApplyQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetItineraryList(v []*ApplyQueryResponseBodyModuleItineraryList) *ApplyQueryResponseBodyModule {
	s.ItineraryList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetItineraryRule(v int32) *ApplyQueryResponseBodyModule {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetItinerarySetList(v []*ApplyQueryResponseBodyModuleItinerarySetList) *ApplyQueryResponseBodyModule {
	s.ItinerarySetList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetLimitTraveler(v int32) *ApplyQueryResponseBodyModule {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetStatus(v int32) *ApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetStatusDesc(v string) *ApplyQueryResponseBodyModule {
	s.StatusDesc = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetThirdpartBusinessId(v string) *ApplyQueryResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetThirdpartId(v string) *ApplyQueryResponseBodyModule {
	s.ThirdpartId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTogetherBookRule(v int32) *ApplyQueryResponseBodyModule {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTrainBudget(v int64) *ApplyQueryResponseBodyModule {
	s.TrainBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTravelerList(v []*ApplyQueryResponseBodyModuleTravelerList) *ApplyQueryResponseBodyModule {
	s.TravelerList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTripCause(v string) *ApplyQueryResponseBodyModule {
	s.TripCause = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTripDay(v int32) *ApplyQueryResponseBodyModule {
	s.TripDay = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTripTitle(v string) *ApplyQueryResponseBodyModule {
	s.TripTitle = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetType(v int32) *ApplyQueryResponseBodyModule {
	s.Type = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetUnionNo(v string) *ApplyQueryResponseBodyModule {
	s.UnionNo = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetUserId(v string) *ApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetUserName(v string) *ApplyQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetVehicleBudget(v int64) *ApplyQueryResponseBodyModule {
	s.VehicleBudget = &v
	return s
}

type ApplyQueryResponseBodyModuleApproverList struct {
	Note        *string `json:"note,omitempty" xml:"note,omitempty"`
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	Order       *int32  `json:"order,omitempty" xml:"order,omitempty"`
	Status      *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc  *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName    *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyQueryResponseBodyModuleApproverList) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleApproverList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetNote(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.Note = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetOperateTime(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.OperateTime = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetOrder(v int32) *ApplyQueryResponseBodyModuleApproverList {
	s.Order = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetStatus(v int32) *ApplyQueryResponseBodyModuleApproverList {
	s.Status = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetStatusDesc(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.StatusDesc = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetUserId(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.UserId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetUserName(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.UserName = &v
	return s
}

type ApplyQueryResponseBodyModuleExternalTravelerList struct {
	BusinessDiscount *int32                                                        `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	EconomyDiscount  *int32                                                        `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount    *int32                                                        `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins     *string                                                       `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	HotelCitys       []*ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	ReserveType      *int32                                                        `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainSeats       *string                                                       `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	UserName         *string                                                       `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyQueryResponseBodyModuleExternalTravelerList) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleExternalTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetBusinessDiscount(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetEconomyDiscount(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetFirstDiscount(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetFlightCabins(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.FlightCabins = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetHotelCitys(v []*ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.HotelCitys = v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetReserveType(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.ReserveType = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetTrainSeats(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.TrainSeats = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetUserName(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.UserName = &v
	return s
}

type ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) SetCityCode(v string) *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) SetCityName(v string) *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) SetFee(v int64) *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys {
	s.Fee = &v
	return s
}

type ApplyQueryResponseBodyModuleHotelShare struct {
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyQueryResponseBodyModuleHotelShare) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleHotelShare) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleHotelShare) SetParam(v string) *ApplyQueryResponseBodyModuleHotelShare {
	s.Param = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleHotelShare) SetType(v string) *ApplyQueryResponseBodyModuleHotelShare {
	s.Type = &v
	return s
}

type ApplyQueryResponseBodyModuleItineraryList struct {
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode    *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrDate        *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode    *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepDate        *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceName    *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	ItineraryId    *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ProjectCode    *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle   *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	TrafficType    *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	TripWay        *int32  `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s ApplyQueryResponseBodyModuleItineraryList) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleItineraryList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetArrCity(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ArrCity = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetArrCityCode(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetArrDate(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ArrDate = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetCostCenterName(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetDepCity(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.DepCity = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetDepCityCode(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetDepDate(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.DepDate = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetInvoiceName(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetItineraryId(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetProjectCode(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetProjectTitle(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetTrafficType(v int32) *ApplyQueryResponseBodyModuleItineraryList {
	s.TrafficType = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetTripWay(v int32) *ApplyQueryResponseBodyModuleItineraryList {
	s.TripWay = &v
	return s
}

type ApplyQueryResponseBodyModuleItinerarySetList struct {
	ArrDate        *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CityCodeSet    *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	CitySet        *string `json:"city_set,omitempty" xml:"city_set,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepDate        *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceName    *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	ItineraryId    *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ProjectCode    *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle   *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	TrafficType    *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s ApplyQueryResponseBodyModuleItinerarySetList) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleItinerarySetList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetArrDate(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ArrDate = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetCityCodeSet(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.CityCodeSet = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetCitySet(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.CitySet = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetCostCenterName(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetDepDate(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.DepDate = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetInvoiceName(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetItineraryId(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetProjectCode(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetProjectTitle(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetTrafficType(v int32) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.TrafficType = &v
	return s
}

type ApplyQueryResponseBodyModuleTravelerList struct {
	BusinessDiscount *int32                                                `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	EconomyDiscount  *int32                                                `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount    *int32                                                `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins     *string                                               `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	HotelCitys       []*ApplyQueryResponseBodyModuleTravelerListHotelCitys `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	ReserveType      *int32                                                `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainSeats       *string                                               `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	UserId           *string                                               `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName         *string                                               `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyQueryResponseBodyModuleTravelerList) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetBusinessDiscount(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetEconomyDiscount(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetFirstDiscount(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetFlightCabins(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.FlightCabins = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetHotelCitys(v []*ApplyQueryResponseBodyModuleTravelerListHotelCitys) *ApplyQueryResponseBodyModuleTravelerList {
	s.HotelCitys = v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetReserveType(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.ReserveType = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetTrainSeats(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.TrainSeats = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetUserId(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.UserId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetUserName(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.UserName = &v
	return s
}

type ApplyQueryResponseBodyModuleTravelerListHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyQueryResponseBodyModuleTravelerListHotelCitys) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleTravelerListHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) SetCityCode(v string) *ApplyQueryResponseBodyModuleTravelerListHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) SetCityName(v string) *ApplyQueryResponseBodyModuleTravelerListHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) SetFee(v int64) *ApplyQueryResponseBodyModuleTravelerListHotelCitys {
	s.Fee = &v
	return s
}

type ApplyQueryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponse) SetHeaders(v map[string]*string) *ApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *ApplyQueryResponse) SetStatusCode(v int32) *ApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyQueryResponse) SetBody(v *ApplyQueryResponseBody) *ApplyQueryResponse {
	s.Body = v
	return s
}

type CarApplyAddHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarApplyAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s CarApplyAddHeaders) GoString() string {
	return s.String()
}

func (s *CarApplyAddHeaders) SetCommonHeaders(v map[string]*string) *CarApplyAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarApplyAddHeaders) SetXAcsBtripSoCorpToken(v string) *CarApplyAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CarApplyAddRequest struct {
	Cause                 *string `json:"cause,omitempty" xml:"cause,omitempty"`
	City                  *string `json:"city,omitempty" xml:"city,omitempty"`
	Date                  *string `json:"date,omitempty" xml:"date,omitempty"`
	FinishedDate          *string `json:"finished_date,omitempty" xml:"finished_date,omitempty"`
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName           *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	Status                *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdPartApplyId      *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	TimesTotal            *int32  `json:"times_total,omitempty" xml:"times_total,omitempty"`
	TimesType             *int32  `json:"times_type,omitempty" xml:"times_type,omitempty"`
	TimesUsed             *int32  `json:"times_used,omitempty" xml:"times_used,omitempty"`
	Title                 *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId                *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyAddRequest) String() string {
	return tea.Prettify(s)
}

func (s CarApplyAddRequest) GoString() string {
	return s.String()
}

func (s *CarApplyAddRequest) SetCause(v string) *CarApplyAddRequest {
	s.Cause = &v
	return s
}

func (s *CarApplyAddRequest) SetCity(v string) *CarApplyAddRequest {
	s.City = &v
	return s
}

func (s *CarApplyAddRequest) SetDate(v string) *CarApplyAddRequest {
	s.Date = &v
	return s
}

func (s *CarApplyAddRequest) SetFinishedDate(v string) *CarApplyAddRequest {
	s.FinishedDate = &v
	return s
}

func (s *CarApplyAddRequest) SetProjectCode(v string) *CarApplyAddRequest {
	s.ProjectCode = &v
	return s
}

func (s *CarApplyAddRequest) SetProjectName(v string) *CarApplyAddRequest {
	s.ProjectName = &v
	return s
}

func (s *CarApplyAddRequest) SetStatus(v int32) *CarApplyAddRequest {
	s.Status = &v
	return s
}

func (s *CarApplyAddRequest) SetThirdPartApplyId(v string) *CarApplyAddRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *CarApplyAddRequest) SetThirdPartCostCenterId(v string) *CarApplyAddRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *CarApplyAddRequest) SetThirdPartInvoiceId(v string) *CarApplyAddRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *CarApplyAddRequest) SetTimesTotal(v int32) *CarApplyAddRequest {
	s.TimesTotal = &v
	return s
}

func (s *CarApplyAddRequest) SetTimesType(v int32) *CarApplyAddRequest {
	s.TimesType = &v
	return s
}

func (s *CarApplyAddRequest) SetTimesUsed(v int32) *CarApplyAddRequest {
	s.TimesUsed = &v
	return s
}

func (s *CarApplyAddRequest) SetTitle(v string) *CarApplyAddRequest {
	s.Title = &v
	return s
}

func (s *CarApplyAddRequest) SetUserId(v string) *CarApplyAddRequest {
	s.UserId = &v
	return s
}

type CarApplyAddResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *int64  `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CarApplyAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CarApplyAddResponseBody) GoString() string {
	return s.String()
}

func (s *CarApplyAddResponseBody) SetCode(v int32) *CarApplyAddResponseBody {
	s.Code = &v
	return s
}

func (s *CarApplyAddResponseBody) SetMessage(v string) *CarApplyAddResponseBody {
	s.Message = &v
	return s
}

func (s *CarApplyAddResponseBody) SetModule(v int64) *CarApplyAddResponseBody {
	s.Module = &v
	return s
}

func (s *CarApplyAddResponseBody) SetRequestId(v string) *CarApplyAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarApplyAddResponseBody) SetSuccess(v bool) *CarApplyAddResponseBody {
	s.Success = &v
	return s
}

func (s *CarApplyAddResponseBody) SetTraceId(v string) *CarApplyAddResponseBody {
	s.TraceId = &v
	return s
}

type CarApplyAddResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CarApplyAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CarApplyAddResponse) String() string {
	return tea.Prettify(s)
}

func (s CarApplyAddResponse) GoString() string {
	return s.String()
}

func (s *CarApplyAddResponse) SetHeaders(v map[string]*string) *CarApplyAddResponse {
	s.Headers = v
	return s
}

func (s *CarApplyAddResponse) SetStatusCode(v int32) *CarApplyAddResponse {
	s.StatusCode = &v
	return s
}

func (s *CarApplyAddResponse) SetBody(v *CarApplyAddResponseBody) *CarApplyAddResponse {
	s.Body = v
	return s
}

type CarApplyModifyHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarApplyModifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s CarApplyModifyHeaders) GoString() string {
	return s.String()
}

func (s *CarApplyModifyHeaders) SetCommonHeaders(v map[string]*string) *CarApplyModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarApplyModifyHeaders) SetXAcsBtripSoCorpToken(v string) *CarApplyModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CarApplyModifyRequest struct {
	OperateTime      *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	Remark           *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Status           *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CarApplyModifyRequest) GoString() string {
	return s.String()
}

func (s *CarApplyModifyRequest) SetOperateTime(v string) *CarApplyModifyRequest {
	s.OperateTime = &v
	return s
}

func (s *CarApplyModifyRequest) SetRemark(v string) *CarApplyModifyRequest {
	s.Remark = &v
	return s
}

func (s *CarApplyModifyRequest) SetStatus(v int32) *CarApplyModifyRequest {
	s.Status = &v
	return s
}

func (s *CarApplyModifyRequest) SetThirdPartApplyId(v string) *CarApplyModifyRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *CarApplyModifyRequest) SetUserId(v string) *CarApplyModifyRequest {
	s.UserId = &v
	return s
}

type CarApplyModifyResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *bool   `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CarApplyModifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CarApplyModifyResponseBody) GoString() string {
	return s.String()
}

func (s *CarApplyModifyResponseBody) SetCode(v int32) *CarApplyModifyResponseBody {
	s.Code = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetMessage(v string) *CarApplyModifyResponseBody {
	s.Message = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetModule(v bool) *CarApplyModifyResponseBody {
	s.Module = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetRequestId(v string) *CarApplyModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetSuccess(v bool) *CarApplyModifyResponseBody {
	s.Success = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetTraceId(v string) *CarApplyModifyResponseBody {
	s.TraceId = &v
	return s
}

type CarApplyModifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CarApplyModifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CarApplyModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CarApplyModifyResponse) GoString() string {
	return s.String()
}

func (s *CarApplyModifyResponse) SetHeaders(v map[string]*string) *CarApplyModifyResponse {
	s.Headers = v
	return s
}

func (s *CarApplyModifyResponse) SetStatusCode(v int32) *CarApplyModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CarApplyModifyResponse) SetBody(v *CarApplyModifyResponseBody) *CarApplyModifyResponse {
	s.Body = v
	return s
}

type CarApplyQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarApplyQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CarApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *CarApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *CarApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CarApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CarApplyQueryRequest struct {
	CreatedEndAt     *string `json:"created_end_at,omitempty" xml:"created_end_at,omitempty"`
	CreatedStartAt   *string `json:"created_start_at,omitempty" xml:"created_start_at,omitempty"`
	PageNumber       *int32  `json:"page_number,omitempty" xml:"page_number,omitempty"`
	PageSize         *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CarApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *CarApplyQueryRequest) SetCreatedEndAt(v string) *CarApplyQueryRequest {
	s.CreatedEndAt = &v
	return s
}

func (s *CarApplyQueryRequest) SetCreatedStartAt(v string) *CarApplyQueryRequest {
	s.CreatedStartAt = &v
	return s
}

func (s *CarApplyQueryRequest) SetPageNumber(v int32) *CarApplyQueryRequest {
	s.PageNumber = &v
	return s
}

func (s *CarApplyQueryRequest) SetPageSize(v int32) *CarApplyQueryRequest {
	s.PageSize = &v
	return s
}

func (s *CarApplyQueryRequest) SetThirdPartApplyId(v string) *CarApplyQueryRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *CarApplyQueryRequest) SetUserId(v string) *CarApplyQueryRequest {
	s.UserId = &v
	return s
}

type CarApplyQueryResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ApplyList []*CarApplyQueryResponseBodyApplyList `json:"apply_list,omitempty" xml:"apply_list,omitempty" type:"Repeated"`
	Code      *int32                                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Success   *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
	Total     *int32                                `json:"total,omitempty" xml:"total,omitempty"`
	TraceId   *string                               `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CarApplyQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CarApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBody) SetRequestId(v string) *CarApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetApplyList(v []*CarApplyQueryResponseBodyApplyList) *CarApplyQueryResponseBody {
	s.ApplyList = v
	return s
}

func (s *CarApplyQueryResponseBody) SetCode(v int32) *CarApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetMessage(v string) *CarApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetSuccess(v bool) *CarApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetTotal(v int32) *CarApplyQueryResponseBody {
	s.Total = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetTraceId(v string) *CarApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

type CarApplyQueryResponseBodyApplyList struct {
	ApproverList  []*CarApplyQueryResponseBodyApplyListApproverList  `json:"approver_list,omitempty" xml:"approver_list,omitempty" type:"Repeated"`
	DepartId      *string                                            `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName    *string                                            `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	GmtCreate     *string                                            `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModified   *string                                            `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	ItineraryList []*CarApplyQueryResponseBodyApplyListItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	Status        *int32                                             `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc    *string                                            `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	ThirdpartId   *string                                            `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	TripCause     *string                                            `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripTitle     *string                                            `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	UserId        *string                                            `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName      *string                                            `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CarApplyQueryResponseBodyApplyList) String() string {
	return tea.Prettify(s)
}

func (s CarApplyQueryResponseBodyApplyList) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBodyApplyList) SetApproverList(v []*CarApplyQueryResponseBodyApplyListApproverList) *CarApplyQueryResponseBodyApplyList {
	s.ApproverList = v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetDepartId(v string) *CarApplyQueryResponseBodyApplyList {
	s.DepartId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetDepartName(v string) *CarApplyQueryResponseBodyApplyList {
	s.DepartName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetGmtCreate(v string) *CarApplyQueryResponseBodyApplyList {
	s.GmtCreate = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetGmtModified(v string) *CarApplyQueryResponseBodyApplyList {
	s.GmtModified = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetItineraryList(v []*CarApplyQueryResponseBodyApplyListItineraryList) *CarApplyQueryResponseBodyApplyList {
	s.ItineraryList = v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetStatus(v int32) *CarApplyQueryResponseBodyApplyList {
	s.Status = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetStatusDesc(v string) *CarApplyQueryResponseBodyApplyList {
	s.StatusDesc = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetThirdpartId(v string) *CarApplyQueryResponseBodyApplyList {
	s.ThirdpartId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetTripCause(v string) *CarApplyQueryResponseBodyApplyList {
	s.TripCause = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetTripTitle(v string) *CarApplyQueryResponseBodyApplyList {
	s.TripTitle = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetUserId(v string) *CarApplyQueryResponseBodyApplyList {
	s.UserId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetUserName(v string) *CarApplyQueryResponseBodyApplyList {
	s.UserName = &v
	return s
}

type CarApplyQueryResponseBodyApplyListApproverList struct {
	Note        *string `json:"note,omitempty" xml:"note,omitempty"`
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	Order       *int32  `json:"order,omitempty" xml:"order,omitempty"`
	Status      *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc  *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName    *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CarApplyQueryResponseBodyApplyListApproverList) String() string {
	return tea.Prettify(s)
}

func (s CarApplyQueryResponseBodyApplyListApproverList) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetNote(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.Note = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetOperateTime(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.OperateTime = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetOrder(v int32) *CarApplyQueryResponseBodyApplyListApproverList {
	s.Order = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetStatus(v int32) *CarApplyQueryResponseBodyApplyListApproverList {
	s.Status = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetStatusDesc(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.StatusDesc = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetUserId(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.UserId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetUserName(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.UserName = &v
	return s
}

type CarApplyQueryResponseBodyApplyListItineraryList struct {
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode    *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrDate        *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CostCenterId   *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode    *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepDate        *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceId      *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	InvoiceName    *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	ItineraryId    *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ProjectCode    *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle   *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	TrafficType    *int32  `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s CarApplyQueryResponseBodyApplyListItineraryList) String() string {
	return tea.Prettify(s)
}

func (s CarApplyQueryResponseBodyApplyListItineraryList) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetArrCity(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ArrCity = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetArrCityCode(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetArrDate(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ArrDate = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetCostCenterId(v int64) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.CostCenterId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetCostCenterName(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.CostCenterName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetDepCity(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.DepCity = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetDepCityCode(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetDepDate(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.DepDate = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetInvoiceId(v int64) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.InvoiceId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetInvoiceName(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.InvoiceName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetItineraryId(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetProjectCode(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetProjectTitle(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetTrafficType(v int32) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.TrafficType = &v
	return s
}

type CarApplyQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CarApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CarApplyQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CarApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponse) SetHeaders(v map[string]*string) *CarApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *CarApplyQueryResponse) SetStatusCode(v int32) *CarApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CarApplyQueryResponse) SetBody(v *CarApplyQueryResponseBody) *CarApplyQueryResponse {
	s.Body = v
	return s
}

type CarBillSettlementQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarBillSettlementQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CarBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *CarBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CarBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CarBillSettlementQueryRequest struct {
	PageNo      *int32  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize    *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PeriodEnd   *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
}

func (s CarBillSettlementQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CarBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryRequest) SetPageNo(v int32) *CarBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetPageSize(v int32) *CarBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetPeriodEnd(v string) *CarBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetPeriodStart(v string) *CarBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

type CarBillSettlementQueryResponseBody struct {
	Code      *int32                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module    *CarBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                   `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CarBillSettlementQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CarBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryResponseBody) SetCode(v int32) *CarBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetMessage(v string) *CarBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetModule(v *CarBillSettlementQueryResponseBodyModule) *CarBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetRequestId(v string) *CarBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetSuccess(v bool) *CarBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetTraceId(v string) *CarBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

type CarBillSettlementQueryResponseBodyModule struct {
	Category    *int32                                              `json:"category,omitempty" xml:"category,omitempty"`
	CorpId      *string                                             `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList    []*CarBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
	PeriodEnd   *string                                             `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string                                             `json:"period_start,omitempty" xml:"period_start,omitempty"`
	TotalNum    *int64                                              `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s CarBillSettlementQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CarBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryResponseBodyModule) SetCategory(v int32) *CarBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetCorpId(v string) *CarBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetDataList(v []*CarBillSettlementQueryResponseBodyModuleDataList) *CarBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *CarBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *CarBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *CarBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

type CarBillSettlementQueryResponseBodyModuleDataList struct {
	AlipayTradeNo         *string  `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyId               *string  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrCity               *string  `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrDate               *string  `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	ArrLocation           *string  `json:"arr_location,omitempty" xml:"arr_location,omitempty"`
	ArrTime               *string  `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BillRecordTime        *string  `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookTime              *string  `json:"book_time,omitempty" xml:"book_time,omitempty"`
	BookerId              *string  `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	BookerJobNo           *string  `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName            *string  `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	BusinessCategory      *string  `json:"business_category,omitempty" xml:"business_category,omitempty"`
	CapitalDirection      *string  `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CarLevel              *string  `json:"car_level,omitempty" xml:"car_level,omitempty"`
	CascadeDepartment     *string  `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CostCenter            *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	CostCenterNumber      *string  `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	Coupon                *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	CouponPrice           *float64 `json:"coupon_price,omitempty" xml:"coupon_price,omitempty"`
	Department            *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId          *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptCity              *string  `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	DeptDate              *string  `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptLocation          *string  `json:"dept_location,omitempty" xml:"dept_location,omitempty"`
	DeptTime              *string  `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	EstimateDriveDistance *string  `json:"estimate_drive_distance,omitempty" xml:"estimate_drive_distance,omitempty"`
	EstimatePrice         *float64 `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
	FeeType               *string  `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	Index                 *string  `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle          *string  `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	Memo                  *string  `json:"memo,omitempty" xml:"memo,omitempty"`
	OrderId               *string  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderPrice            *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OverApplyId           *string  `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PersonSettleFee       *float64 `json:"person_settle_fee,omitempty" xml:"person_settle_fee,omitempty"`
	PrimaryId             *int64   `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProjectCode           *string  `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName           *string  `json:"project_name,omitempty" xml:"project_name,omitempty"`
	ProviderName          *string  `json:"provider_name,omitempty" xml:"provider_name,omitempty"`
	RealDriveDistance     *string  `json:"real_drive_distance,omitempty" xml:"real_drive_distance,omitempty"`
	RealFromAddr          *string  `json:"real_from_addr,omitempty" xml:"real_from_addr,omitempty"`
	RealToAddr            *string  `json:"real_to_addr,omitempty" xml:"real_to_addr,omitempty"`
	Remark                *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	ServiceFee            *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettlementFee         *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	SettlementGrantFee    *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	SettlementTime        *string  `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	SettlementType        *string  `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	SpecialOrder          *string  `json:"special_order,omitempty" xml:"special_order,omitempty"`
	SpecialReason         *string  `json:"special_reason,omitempty" xml:"special_reason,omitempty"`
	Status                *int32   `json:"status,omitempty" xml:"status,omitempty"`
	SubOrderId            *string  `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	TravelerId            *string  `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	TravelerJobNo         *string  `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerName          *string  `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	UserConfirmDesc       *string  `json:"user_confirm_desc,omitempty" xml:"user_confirm_desc,omitempty"`
	VoucherType           *int32   `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}

func (s CarBillSettlementQueryResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s CarBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrCity(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCity = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrDate(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrLocation(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrLocation = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBusinessCategory(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BusinessCategory = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCarLevel(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CarLevel = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCoupon(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCouponPrice(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CouponPrice = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeptCity(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeptCity = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeptDate(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeptLocation(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeptLocation = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeptTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetEstimateDriveDistance(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.EstimateDriveDistance = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetEstimatePrice(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.EstimatePrice = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetMemo(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Memo = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetOrderPrice(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetPersonSettleFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.PersonSettleFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetProviderName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ProviderName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetRealDriveDistance(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.RealDriveDistance = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetRealFromAddr(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.RealFromAddr = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetRealToAddr(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.RealToAddr = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSpecialOrder(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SpecialOrder = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSpecialReason(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SpecialReason = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSubOrderId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SubOrderId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetUserConfirmDesc(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.UserConfirmDesc = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

type CarBillSettlementQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CarBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CarBillSettlementQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CarBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryResponse) SetHeaders(v map[string]*string) *CarBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *CarBillSettlementQueryResponse) SetStatusCode(v int32) *CarBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CarBillSettlementQueryResponse) SetBody(v *CarBillSettlementQueryResponseBody) *CarBillSettlementQueryResponse {
	s.Body = v
	return s
}

type CarOrderListQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarOrderListQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CarOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *CarOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarOrderListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CarOrderListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CarOrderListQueryRequest struct {
	AllApply         *bool   `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	ApplyId          *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	DepartId         *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Page             *int32  `json:"page,omitempty" xml:"page,omitempty"`
	PageSize         *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	UpdateEndTime    *string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	UpdateStartTime  *string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarOrderListQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CarOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryRequest) SetAllApply(v bool) *CarOrderListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *CarOrderListQueryRequest) SetApplyId(v int64) *CarOrderListQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *CarOrderListQueryRequest) SetDepartId(v string) *CarOrderListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *CarOrderListQueryRequest) SetEndTime(v string) *CarOrderListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *CarOrderListQueryRequest) SetPage(v int32) *CarOrderListQueryRequest {
	s.Page = &v
	return s
}

func (s *CarOrderListQueryRequest) SetPageSize(v int32) *CarOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *CarOrderListQueryRequest) SetStartTime(v string) *CarOrderListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *CarOrderListQueryRequest) SetThirdpartApplyId(v string) *CarOrderListQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *CarOrderListQueryRequest) SetUpdateEndTime(v string) *CarOrderListQueryRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *CarOrderListQueryRequest) SetUpdateStartTime(v string) *CarOrderListQueryRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *CarOrderListQueryRequest) SetUserId(v string) *CarOrderListQueryRequest {
	s.UserId = &v
	return s
}

type CarOrderListQueryResponseBody struct {
	Code      *int32                                 `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*CarOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo  *CarOrderListQueryResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CarOrderListQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CarOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBody) SetCode(v int32) *CarOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CarOrderListQueryResponseBody) SetMessage(v string) *CarOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CarOrderListQueryResponseBody) SetModule(v []*CarOrderListQueryResponseBodyModule) *CarOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *CarOrderListQueryResponseBody) SetPageInfo(v *CarOrderListQueryResponseBodyPageInfo) *CarOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *CarOrderListQueryResponseBody) SetRequestId(v string) *CarOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarOrderListQueryResponseBody) SetSuccess(v bool) *CarOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CarOrderListQueryResponseBody) SetTraceId(v string) *CarOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

type CarOrderListQueryResponseBodyModule struct {
	ApplyId              *int64                                                  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApplyShowId          *string                                                 `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	BtripTitle           *string                                                 `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	BusinessCategory     *string                                                 `json:"business_category,omitempty" xml:"business_category,omitempty"`
	CancelTime           *string                                                 `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	CarInfo              *string                                                 `json:"car_info,omitempty" xml:"car_info,omitempty"`
	CarLevel             *int32                                                  `json:"car_level,omitempty" xml:"car_level,omitempty"`
	CorpId               *string                                                 `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName             *string                                                 `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	CostCenterId         *int64                                                  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName       *string                                                 `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	CostCenterNumber     *string                                                 `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	DeptId               *int64                                                  `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	DeptName             *string                                                 `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	DriverConfirmTime    *string                                                 `json:"driver_confirm_time,omitempty" xml:"driver_confirm_time,omitempty"`
	EstimatePrice        *float64                                                `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
	FromAddress          *string                                                 `json:"from_address,omitempty" xml:"from_address,omitempty"`
	FromCityName         *string                                                 `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	GmtCreate            *string                                                 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModified          *string                                                 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	Id                   *int64                                                  `json:"id,omitempty" xml:"id,omitempty"`
	InvoiceId            *int64                                                  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	InvoiceTitle         *string                                                 `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	IsSpecial            *bool                                                   `json:"is_special,omitempty" xml:"is_special,omitempty"`
	Memo                 *string                                                 `json:"memo,omitempty" xml:"memo,omitempty"`
	OrderStatus          *int32                                                  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	PassengerName        *string                                                 `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PayTime              *string                                                 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	PriceInfoList        []*CarOrderListQueryResponseBodyModulePriceInfoList     `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	ProjectCode          *string                                                 `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectId            *int64                                                  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle         *string                                                 `json:"project_title,omitempty" xml:"project_title,omitempty"`
	Provider             *int32                                                  `json:"provider,omitempty" xml:"provider,omitempty"`
	PublishTime          *string                                                 `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	RealFromAddress      *string                                                 `json:"real_from_address,omitempty" xml:"real_from_address,omitempty"`
	RealFromCityName     *string                                                 `json:"real_from_city_name,omitempty" xml:"real_from_city_name,omitempty"`
	RealToAddress        *string                                                 `json:"real_to_address,omitempty" xml:"real_to_address,omitempty"`
	RealToCityName       *string                                                 `json:"real_to_city_name,omitempty" xml:"real_to_city_name,omitempty"`
	ServiceType          *int32                                                  `json:"service_type,omitempty" xml:"service_type,omitempty"`
	SpecialTypes         []*string                                               `json:"special_types,omitempty" xml:"special_types,omitempty" type:"Repeated"`
	TakenTime            *string                                                 `json:"taken_time,omitempty" xml:"taken_time,omitempty"`
	ThirdpartApplyId     *string                                                 `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartItineraryId *string                                                 `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	ToAddress            *string                                                 `json:"to_address,omitempty" xml:"to_address,omitempty"`
	ToCityName           *string                                                 `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
	TravelDistance       *float64                                                `json:"travel_distance,omitempty" xml:"travel_distance,omitempty"`
	UserAffiliateList    []*CarOrderListQueryResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	UserConfirm          *int32                                                  `json:"user_confirm,omitempty" xml:"user_confirm,omitempty"`
	UserId               *string                                                 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName             *string                                                 `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CarOrderListQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CarOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBodyModule) SetApplyId(v int64) *CarOrderListQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetApplyShowId(v string) *CarOrderListQueryResponseBodyModule {
	s.ApplyShowId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetBtripTitle(v string) *CarOrderListQueryResponseBodyModule {
	s.BtripTitle = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetBusinessCategory(v string) *CarOrderListQueryResponseBodyModule {
	s.BusinessCategory = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCancelTime(v string) *CarOrderListQueryResponseBodyModule {
	s.CancelTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCarInfo(v string) *CarOrderListQueryResponseBodyModule {
	s.CarInfo = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCarLevel(v int32) *CarOrderListQueryResponseBodyModule {
	s.CarLevel = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCorpId(v string) *CarOrderListQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCorpName(v string) *CarOrderListQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCostCenterId(v int64) *CarOrderListQueryResponseBodyModule {
	s.CostCenterId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCostCenterName(v string) *CarOrderListQueryResponseBodyModule {
	s.CostCenterName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCostCenterNumber(v string) *CarOrderListQueryResponseBodyModule {
	s.CostCenterNumber = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetDeptId(v int64) *CarOrderListQueryResponseBodyModule {
	s.DeptId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetDeptName(v string) *CarOrderListQueryResponseBodyModule {
	s.DeptName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetDriverConfirmTime(v string) *CarOrderListQueryResponseBodyModule {
	s.DriverConfirmTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetEstimatePrice(v float64) *CarOrderListQueryResponseBodyModule {
	s.EstimatePrice = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetFromAddress(v string) *CarOrderListQueryResponseBodyModule {
	s.FromAddress = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetFromCityName(v string) *CarOrderListQueryResponseBodyModule {
	s.FromCityName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetGmtCreate(v string) *CarOrderListQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetGmtModified(v string) *CarOrderListQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetId(v int64) *CarOrderListQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetInvoiceId(v int64) *CarOrderListQueryResponseBodyModule {
	s.InvoiceId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetInvoiceTitle(v string) *CarOrderListQueryResponseBodyModule {
	s.InvoiceTitle = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetIsSpecial(v bool) *CarOrderListQueryResponseBodyModule {
	s.IsSpecial = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetMemo(v string) *CarOrderListQueryResponseBodyModule {
	s.Memo = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetOrderStatus(v int32) *CarOrderListQueryResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetPassengerName(v string) *CarOrderListQueryResponseBodyModule {
	s.PassengerName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetPayTime(v string) *CarOrderListQueryResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetPriceInfoList(v []*CarOrderListQueryResponseBodyModulePriceInfoList) *CarOrderListQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetProjectCode(v string) *CarOrderListQueryResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetProjectId(v int64) *CarOrderListQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetProjectTitle(v string) *CarOrderListQueryResponseBodyModule {
	s.ProjectTitle = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetProvider(v int32) *CarOrderListQueryResponseBodyModule {
	s.Provider = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetPublishTime(v string) *CarOrderListQueryResponseBodyModule {
	s.PublishTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealFromAddress(v string) *CarOrderListQueryResponseBodyModule {
	s.RealFromAddress = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealFromCityName(v string) *CarOrderListQueryResponseBodyModule {
	s.RealFromCityName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealToAddress(v string) *CarOrderListQueryResponseBodyModule {
	s.RealToAddress = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealToCityName(v string) *CarOrderListQueryResponseBodyModule {
	s.RealToCityName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetServiceType(v int32) *CarOrderListQueryResponseBodyModule {
	s.ServiceType = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetSpecialTypes(v []*string) *CarOrderListQueryResponseBodyModule {
	s.SpecialTypes = v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetTakenTime(v string) *CarOrderListQueryResponseBodyModule {
	s.TakenTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetThirdpartApplyId(v string) *CarOrderListQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetThirdpartItineraryId(v string) *CarOrderListQueryResponseBodyModule {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetToAddress(v string) *CarOrderListQueryResponseBodyModule {
	s.ToAddress = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetToCityName(v string) *CarOrderListQueryResponseBodyModule {
	s.ToCityName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetTravelDistance(v float64) *CarOrderListQueryResponseBodyModule {
	s.TravelDistance = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetUserAffiliateList(v []*CarOrderListQueryResponseBodyModuleUserAffiliateList) *CarOrderListQueryResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetUserConfirm(v int32) *CarOrderListQueryResponseBodyModule {
	s.UserConfirm = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetUserId(v string) *CarOrderListQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetUserName(v string) *CarOrderListQueryResponseBodyModule {
	s.UserName = &v
	return s
}

type CarOrderListQueryResponseBodyModulePriceInfoList struct {
	CategoryCode  *int32   `json:"category_code,omitempty" xml:"category_code,omitempty"`
	CategoryType  *int32   `json:"category_type,omitempty" xml:"category_type,omitempty"`
	GmtCreate     *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string  `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PayType       *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	PersonPrice   *float64 `json:"person_price,omitempty" xml:"person_price,omitempty"`
	Price         *float64 `json:"price,omitempty" xml:"price,omitempty"`
	TradeId       *string  `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	Type          *int32   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CarOrderListQueryResponseBodyModulePriceInfoList) String() string {
	return tea.Prettify(s)
}

func (s CarOrderListQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetCategoryType(v int32) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetPersonPrice(v float64) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.PersonPrice = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetType(v int32) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

type CarOrderListQueryResponseBodyModuleUserAffiliateList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CarOrderListQueryResponseBodyModuleUserAffiliateList) String() string {
	return tea.Prettify(s)
}

func (s CarOrderListQueryResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBodyModuleUserAffiliateList) SetUserId(v string) *CarOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModuleUserAffiliateList) SetUserName(v string) *CarOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

type CarOrderListQueryResponseBodyPageInfo struct {
	Page        *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageSize    *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s CarOrderListQueryResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s CarOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBodyPageInfo) SetPage(v int32) *CarOrderListQueryResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *CarOrderListQueryResponseBodyPageInfo) SetPageSize(v int32) *CarOrderListQueryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *CarOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *CarOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

type CarOrderListQueryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CarOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CarOrderListQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CarOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponse) SetHeaders(v map[string]*string) *CarOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *CarOrderListQueryResponse) SetStatusCode(v int32) *CarOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CarOrderListQueryResponse) SetBody(v *CarOrderListQueryResponseBody) *CarOrderListQueryResponse {
	s.Body = v
	return s
}

type CitySearchHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CitySearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s CitySearchHeaders) GoString() string {
	return s.String()
}

func (s *CitySearchHeaders) SetCommonHeaders(v map[string]*string) *CitySearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CitySearchHeaders) SetXAcsBtripSoCorpToken(v string) *CitySearchHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CitySearchRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

func (s CitySearchRequest) String() string {
	return tea.Prettify(s)
}

func (s CitySearchRequest) GoString() string {
	return s.String()
}

func (s *CitySearchRequest) SetKeyword(v string) *CitySearchRequest {
	s.Keyword = &v
	return s
}

type CitySearchResponseBody struct {
	Code      *int32                        `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                       `json:"message,omitempty" xml:"message,omitempty"`
	Module    *CitySearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                         `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                       `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CitySearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CitySearchResponseBody) GoString() string {
	return s.String()
}

func (s *CitySearchResponseBody) SetCode(v int32) *CitySearchResponseBody {
	s.Code = &v
	return s
}

func (s *CitySearchResponseBody) SetMessage(v string) *CitySearchResponseBody {
	s.Message = &v
	return s
}

func (s *CitySearchResponseBody) SetModule(v *CitySearchResponseBodyModule) *CitySearchResponseBody {
	s.Module = v
	return s
}

func (s *CitySearchResponseBody) SetRequestId(v string) *CitySearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CitySearchResponseBody) SetSuccess(v bool) *CitySearchResponseBody {
	s.Success = &v
	return s
}

func (s *CitySearchResponseBody) SetTraceId(v string) *CitySearchResponseBody {
	s.TraceId = &v
	return s
}

type CitySearchResponseBodyModule struct {
	Cities []*CitySearchResponseBodyModuleCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
}

func (s CitySearchResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CitySearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CitySearchResponseBodyModule) SetCities(v []*CitySearchResponseBodyModuleCities) *CitySearchResponseBodyModule {
	s.Cities = v
	return s
}

type CitySearchResponseBodyModuleCities struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Region *int32  `json:"region,omitempty" xml:"region,omitempty"`
}

func (s CitySearchResponseBodyModuleCities) String() string {
	return tea.Prettify(s)
}

func (s CitySearchResponseBodyModuleCities) GoString() string {
	return s.String()
}

func (s *CitySearchResponseBodyModuleCities) SetCode(v string) *CitySearchResponseBodyModuleCities {
	s.Code = &v
	return s
}

func (s *CitySearchResponseBodyModuleCities) SetName(v string) *CitySearchResponseBodyModuleCities {
	s.Name = &v
	return s
}

func (s *CitySearchResponseBodyModuleCities) SetRegion(v int32) *CitySearchResponseBodyModuleCities {
	s.Region = &v
	return s
}

type CitySearchResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CitySearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CitySearchResponse) String() string {
	return tea.Prettify(s)
}

func (s CitySearchResponse) GoString() string {
	return s.String()
}

func (s *CitySearchResponse) SetHeaders(v map[string]*string) *CitySearchResponse {
	s.Headers = v
	return s
}

func (s *CitySearchResponse) SetStatusCode(v int32) *CitySearchResponse {
	s.StatusCode = &v
	return s
}

func (s *CitySearchResponse) SetBody(v *CitySearchResponseBody) *CitySearchResponse {
	s.Body = v
	return s
}

type CommonApplyQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CommonApplyQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CommonApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *CommonApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommonApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CommonApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CommonApplyQueryRequest struct {
	ApplyId     *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BizCategory *int32  `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CommonApplyQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CommonApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryRequest) SetApplyId(v int64) *CommonApplyQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *CommonApplyQueryRequest) SetBizCategory(v int32) *CommonApplyQueryRequest {
	s.BizCategory = &v
	return s
}

func (s *CommonApplyQueryRequest) SetUserId(v string) *CommonApplyQueryRequest {
	s.UserId = &v
	return s
}

type CommonApplyQueryResponseBody struct {
	Code      *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module    *CommonApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                             `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CommonApplyQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommonApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryResponseBody) SetCode(v int32) *CommonApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CommonApplyQueryResponseBody) SetMessage(v string) *CommonApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CommonApplyQueryResponseBody) SetModule(v *CommonApplyQueryResponseBodyModule) *CommonApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *CommonApplyQueryResponseBody) SetRequestId(v string) *CommonApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommonApplyQueryResponseBody) SetSuccess(v bool) *CommonApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CommonApplyQueryResponseBody) SetTraceId(v string) *CommonApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

type CommonApplyQueryResponseBodyModule struct {
	ApplyId         *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BizCategory     *int32  `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
	Cause           *string `json:"cause,omitempty" xml:"cause,omitempty"`
	CorpId          *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	ExtendValue     *string `json:"extend_value,omitempty" xml:"extend_value,omitempty"`
	GmtCreate       *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	Status          *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdpartCorpId *string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	ThirdpartId     *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	TripCause       *string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	UserId          *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CommonApplyQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CommonApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryResponseBodyModule) SetApplyId(v int64) *CommonApplyQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetBizCategory(v int32) *CommonApplyQueryResponseBodyModule {
	s.BizCategory = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetCause(v string) *CommonApplyQueryResponseBodyModule {
	s.Cause = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetCorpId(v string) *CommonApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetExtendValue(v string) *CommonApplyQueryResponseBodyModule {
	s.ExtendValue = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetGmtCreate(v string) *CommonApplyQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetStatus(v int32) *CommonApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetThirdpartCorpId(v string) *CommonApplyQueryResponseBodyModule {
	s.ThirdpartCorpId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetThirdpartId(v string) *CommonApplyQueryResponseBodyModule {
	s.ThirdpartId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetTripCause(v string) *CommonApplyQueryResponseBodyModule {
	s.TripCause = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetUserId(v string) *CommonApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

type CommonApplyQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CommonApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommonApplyQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CommonApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryResponse) SetHeaders(v map[string]*string) *CommonApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *CommonApplyQueryResponse) SetStatusCode(v int32) *CommonApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CommonApplyQueryResponse) SetBody(v *CommonApplyQueryResponseBody) *CommonApplyQueryResponse {
	s.Body = v
	return s
}

type CommonApplySyncHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CommonApplySyncHeaders) String() string {
	return tea.Prettify(s)
}

func (s CommonApplySyncHeaders) GoString() string {
	return s.String()
}

func (s *CommonApplySyncHeaders) SetCommonHeaders(v map[string]*string) *CommonApplySyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommonApplySyncHeaders) SetXAcsBtripSoCorpToken(v string) *CommonApplySyncHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CommonApplySyncRequest struct {
	ApplyId          *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BizCategory      *int32  `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
	Remark           *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Status           *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdpartyFlowId *string `json:"thirdparty_flow_id,omitempty" xml:"thirdparty_flow_id,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CommonApplySyncRequest) String() string {
	return tea.Prettify(s)
}

func (s CommonApplySyncRequest) GoString() string {
	return s.String()
}

func (s *CommonApplySyncRequest) SetApplyId(v int64) *CommonApplySyncRequest {
	s.ApplyId = &v
	return s
}

func (s *CommonApplySyncRequest) SetBizCategory(v int32) *CommonApplySyncRequest {
	s.BizCategory = &v
	return s
}

func (s *CommonApplySyncRequest) SetRemark(v string) *CommonApplySyncRequest {
	s.Remark = &v
	return s
}

func (s *CommonApplySyncRequest) SetStatus(v int32) *CommonApplySyncRequest {
	s.Status = &v
	return s
}

func (s *CommonApplySyncRequest) SetThirdpartyFlowId(v string) *CommonApplySyncRequest {
	s.ThirdpartyFlowId = &v
	return s
}

func (s *CommonApplySyncRequest) SetUserId(v string) *CommonApplySyncRequest {
	s.UserId = &v
	return s
}

type CommonApplySyncResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *bool   `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CommonApplySyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommonApplySyncResponseBody) GoString() string {
	return s.String()
}

func (s *CommonApplySyncResponseBody) SetCode(v int32) *CommonApplySyncResponseBody {
	s.Code = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetMessage(v string) *CommonApplySyncResponseBody {
	s.Message = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetModule(v bool) *CommonApplySyncResponseBody {
	s.Module = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetRequestId(v string) *CommonApplySyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetSuccess(v bool) *CommonApplySyncResponseBody {
	s.Success = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetTraceId(v string) *CommonApplySyncResponseBody {
	s.TraceId = &v
	return s
}

type CommonApplySyncResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CommonApplySyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommonApplySyncResponse) String() string {
	return tea.Prettify(s)
}

func (s CommonApplySyncResponse) GoString() string {
	return s.String()
}

func (s *CommonApplySyncResponse) SetHeaders(v map[string]*string) *CommonApplySyncResponse {
	s.Headers = v
	return s
}

func (s *CommonApplySyncResponse) SetStatusCode(v int32) *CommonApplySyncResponse {
	s.StatusCode = &v
	return s
}

func (s *CommonApplySyncResponse) SetBody(v *CommonApplySyncResponseBody) *CommonApplySyncResponse {
	s.Body = v
	return s
}

type CorpTokenHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripAccessToken *string            `json:"x-acs-btrip-access-token,omitempty" xml:"x-acs-btrip-access-token,omitempty"`
}

func (s CorpTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s CorpTokenHeaders) GoString() string {
	return s.String()
}

func (s *CorpTokenHeaders) SetCommonHeaders(v map[string]*string) *CorpTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CorpTokenHeaders) SetXAcsBtripAccessToken(v string) *CorpTokenHeaders {
	s.XAcsBtripAccessToken = &v
	return s
}

type CorpTokenRequest struct {
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Type   *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CorpTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CorpTokenRequest) GoString() string {
	return s.String()
}

func (s *CorpTokenRequest) SetCorpId(v string) *CorpTokenRequest {
	s.CorpId = &v
	return s
}

func (s *CorpTokenRequest) SetType(v int32) *CorpTokenRequest {
	s.Type = &v
	return s
}

type CorpTokenResponseBody struct {
	Code      *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data      *CorpTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                    `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TraceId   *string                    `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CorpTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CorpTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CorpTokenResponseBody) SetCode(v string) *CorpTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CorpTokenResponseBody) SetData(v *CorpTokenResponseBodyData) *CorpTokenResponseBody {
	s.Data = v
	return s
}

func (s *CorpTokenResponseBody) SetMessage(v string) *CorpTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CorpTokenResponseBody) SetRequestId(v string) *CorpTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CorpTokenResponseBody) SetTraceId(v string) *CorpTokenResponseBody {
	s.TraceId = &v
	return s
}

type CorpTokenResponseBodyData struct {
	Expire *int64  `json:"expire,omitempty" xml:"expire,omitempty"`
	Token  *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s CorpTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CorpTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CorpTokenResponseBodyData) SetExpire(v int64) *CorpTokenResponseBodyData {
	s.Expire = &v
	return s
}

func (s *CorpTokenResponseBodyData) SetToken(v string) *CorpTokenResponseBodyData {
	s.Token = &v
	return s
}

type CorpTokenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CorpTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CorpTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CorpTokenResponse) GoString() string {
	return s.String()
}

func (s *CorpTokenResponse) SetHeaders(v map[string]*string) *CorpTokenResponse {
	s.Headers = v
	return s
}

func (s *CorpTokenResponse) SetStatusCode(v int32) *CorpTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CorpTokenResponse) SetBody(v *CorpTokenResponseBody) *CorpTokenResponse {
	s.Body = v
	return s
}

type CostCenterDeleteHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CostCenterDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s CostCenterDeleteHeaders) GoString() string {
	return s.String()
}

func (s *CostCenterDeleteHeaders) SetCommonHeaders(v map[string]*string) *CostCenterDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CostCenterDeleteHeaders) SetXAcsBtripSoCorpToken(v string) *CostCenterDeleteHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CostCenterDeleteRequest struct {
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s CostCenterDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CostCenterDeleteRequest) GoString() string {
	return s.String()
}

func (s *CostCenterDeleteRequest) SetThirdpartId(v string) *CostCenterDeleteRequest {
	s.ThirdpartId = &v
	return s
}

type CostCenterDeleteResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CostCenterDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CostCenterDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CostCenterDeleteResponseBody) SetCode(v int32) *CostCenterDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *CostCenterDeleteResponseBody) SetMessage(v string) *CostCenterDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *CostCenterDeleteResponseBody) SetRequestId(v string) *CostCenterDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CostCenterDeleteResponseBody) SetSuccess(v bool) *CostCenterDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *CostCenterDeleteResponseBody) SetTraceId(v string) *CostCenterDeleteResponseBody {
	s.TraceId = &v
	return s
}

type CostCenterDeleteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CostCenterDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CostCenterDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CostCenterDeleteResponse) GoString() string {
	return s.String()
}

func (s *CostCenterDeleteResponse) SetHeaders(v map[string]*string) *CostCenterDeleteResponse {
	s.Headers = v
	return s
}

func (s *CostCenterDeleteResponse) SetStatusCode(v int32) *CostCenterDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CostCenterDeleteResponse) SetBody(v *CostCenterDeleteResponseBody) *CostCenterDeleteResponse {
	s.Body = v
	return s
}

type CostCenterModifyHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CostCenterModifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s CostCenterModifyHeaders) GoString() string {
	return s.String()
}

func (s *CostCenterModifyHeaders) SetCommonHeaders(v map[string]*string) *CostCenterModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CostCenterModifyHeaders) SetXAcsBtripSoCorpToken(v string) *CostCenterModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CostCenterModifyRequest struct {
	AlipayNo    *string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	Number      *string `json:"number,omitempty" xml:"number,omitempty"`
	Scope       *int64  `json:"scope,omitempty" xml:"scope,omitempty"`
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CostCenterModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CostCenterModifyRequest) GoString() string {
	return s.String()
}

func (s *CostCenterModifyRequest) SetAlipayNo(v string) *CostCenterModifyRequest {
	s.AlipayNo = &v
	return s
}

func (s *CostCenterModifyRequest) SetNumber(v string) *CostCenterModifyRequest {
	s.Number = &v
	return s
}

func (s *CostCenterModifyRequest) SetScope(v int64) *CostCenterModifyRequest {
	s.Scope = &v
	return s
}

func (s *CostCenterModifyRequest) SetThirdpartId(v string) *CostCenterModifyRequest {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterModifyRequest) SetTitle(v string) *CostCenterModifyRequest {
	s.Title = &v
	return s
}

type CostCenterModifyResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CostCenterModifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CostCenterModifyResponseBody) GoString() string {
	return s.String()
}

func (s *CostCenterModifyResponseBody) SetCode(v int32) *CostCenterModifyResponseBody {
	s.Code = &v
	return s
}

func (s *CostCenterModifyResponseBody) SetMessage(v string) *CostCenterModifyResponseBody {
	s.Message = &v
	return s
}

func (s *CostCenterModifyResponseBody) SetRequestId(v string) *CostCenterModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CostCenterModifyResponseBody) SetSuccess(v bool) *CostCenterModifyResponseBody {
	s.Success = &v
	return s
}

func (s *CostCenterModifyResponseBody) SetTraceId(v string) *CostCenterModifyResponseBody {
	s.TraceId = &v
	return s
}

type CostCenterModifyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CostCenterModifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CostCenterModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CostCenterModifyResponse) GoString() string {
	return s.String()
}

func (s *CostCenterModifyResponse) SetHeaders(v map[string]*string) *CostCenterModifyResponse {
	s.Headers = v
	return s
}

func (s *CostCenterModifyResponse) SetStatusCode(v int32) *CostCenterModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CostCenterModifyResponse) SetBody(v *CostCenterModifyResponseBody) *CostCenterModifyResponse {
	s.Body = v
	return s
}

type CostCenterQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CostCenterQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CostCenterQueryHeaders) GoString() string {
	return s.String()
}

func (s *CostCenterQueryHeaders) SetCommonHeaders(v map[string]*string) *CostCenterQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CostCenterQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CostCenterQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CostCenterQueryRequest struct {
	NeedOrgEntity *bool   `json:"need_org_entity,omitempty" xml:"need_org_entity,omitempty"`
	ThirdpartId   *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	Title         *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId        *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CostCenterQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CostCenterQueryRequest) GoString() string {
	return s.String()
}

func (s *CostCenterQueryRequest) SetNeedOrgEntity(v bool) *CostCenterQueryRequest {
	s.NeedOrgEntity = &v
	return s
}

func (s *CostCenterQueryRequest) SetThirdpartId(v string) *CostCenterQueryRequest {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterQueryRequest) SetTitle(v string) *CostCenterQueryRequest {
	s.Title = &v
	return s
}

func (s *CostCenterQueryRequest) SetUserId(v string) *CostCenterQueryRequest {
	s.UserId = &v
	return s
}

type CostCenterQueryResponseBody struct {
	Code      *int32                               `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*CostCenterQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	MorePage  *bool                                `json:"more_page,omitempty" xml:"more_page,omitempty"`
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                              `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CostCenterQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CostCenterQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CostCenterQueryResponseBody) SetCode(v int32) *CostCenterQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetMessage(v string) *CostCenterQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetModule(v []*CostCenterQueryResponseBodyModule) *CostCenterQueryResponseBody {
	s.Module = v
	return s
}

func (s *CostCenterQueryResponseBody) SetMorePage(v bool) *CostCenterQueryResponseBody {
	s.MorePage = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetRequestId(v string) *CostCenterQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetSuccess(v bool) *CostCenterQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetTraceId(v string) *CostCenterQueryResponseBody {
	s.TraceId = &v
	return s
}

type CostCenterQueryResponseBodyModule struct {
	AlipayNo    *string                                       `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	CorpId      *string                                       `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	EntityDOS   []*CostCenterQueryResponseBodyModuleEntityDOS `json:"entity_d_o_s,omitempty" xml:"entity_d_o_s,omitempty" type:"Repeated"`
	Id          *int64                                        `json:"id,omitempty" xml:"id,omitempty"`
	Number      *string                                       `json:"number,omitempty" xml:"number,omitempty"`
	RuleCode    *int64                                        `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
	Scope       *int64                                        `json:"scope,omitempty" xml:"scope,omitempty"`
	ThirdpartId *string                                       `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	Title       *string                                       `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CostCenterQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CostCenterQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CostCenterQueryResponseBodyModule) SetAlipayNo(v string) *CostCenterQueryResponseBodyModule {
	s.AlipayNo = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetCorpId(v string) *CostCenterQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetEntityDOS(v []*CostCenterQueryResponseBodyModuleEntityDOS) *CostCenterQueryResponseBodyModule {
	s.EntityDOS = v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetId(v int64) *CostCenterQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetNumber(v string) *CostCenterQueryResponseBodyModule {
	s.Number = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetRuleCode(v int64) *CostCenterQueryResponseBodyModule {
	s.RuleCode = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetScope(v int64) *CostCenterQueryResponseBodyModule {
	s.Scope = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetThirdpartId(v string) *CostCenterQueryResponseBodyModule {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetTitle(v string) *CostCenterQueryResponseBodyModule {
	s.Title = &v
	return s
}

type CostCenterQueryResponseBodyModuleEntityDOS struct {
	CorpId     *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	EntityId   *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	UserNum    *int32  `json:"user_num,omitempty" xml:"user_num,omitempty"`
}

func (s CostCenterQueryResponseBodyModuleEntityDOS) String() string {
	return tea.Prettify(s)
}

func (s CostCenterQueryResponseBodyModuleEntityDOS) GoString() string {
	return s.String()
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetCorpId(v string) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.CorpId = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetEntityId(v string) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.EntityId = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetEntityType(v string) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.EntityType = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetName(v string) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.Name = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetUserNum(v int32) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.UserNum = &v
	return s
}

type CostCenterQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CostCenterQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CostCenterQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CostCenterQueryResponse) GoString() string {
	return s.String()
}

func (s *CostCenterQueryResponse) SetHeaders(v map[string]*string) *CostCenterQueryResponse {
	s.Headers = v
	return s
}

func (s *CostCenterQueryResponse) SetStatusCode(v int32) *CostCenterQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CostCenterQueryResponse) SetBody(v *CostCenterQueryResponseBody) *CostCenterQueryResponse {
	s.Body = v
	return s
}

type CostCenterSaveHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CostCenterSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s CostCenterSaveHeaders) GoString() string {
	return s.String()
}

func (s *CostCenterSaveHeaders) SetCommonHeaders(v map[string]*string) *CostCenterSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CostCenterSaveHeaders) SetXAcsBtripSoCorpToken(v string) *CostCenterSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type CostCenterSaveRequest struct {
	AlipayNo    *string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	Number      *string `json:"number,omitempty" xml:"number,omitempty"`
	Scope       *int64  `json:"scope,omitempty" xml:"scope,omitempty"`
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CostCenterSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s CostCenterSaveRequest) GoString() string {
	return s.String()
}

func (s *CostCenterSaveRequest) SetAlipayNo(v string) *CostCenterSaveRequest {
	s.AlipayNo = &v
	return s
}

func (s *CostCenterSaveRequest) SetNumber(v string) *CostCenterSaveRequest {
	s.Number = &v
	return s
}

func (s *CostCenterSaveRequest) SetScope(v int64) *CostCenterSaveRequest {
	s.Scope = &v
	return s
}

func (s *CostCenterSaveRequest) SetThirdpartId(v string) *CostCenterSaveRequest {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterSaveRequest) SetTitle(v string) *CostCenterSaveRequest {
	s.Title = &v
	return s
}

type CostCenterSaveResponseBody struct {
	Code      *int32                            `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module    *CostCenterSaveResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                             `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                           `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CostCenterSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CostCenterSaveResponseBody) GoString() string {
	return s.String()
}

func (s *CostCenterSaveResponseBody) SetCode(v int32) *CostCenterSaveResponseBody {
	s.Code = &v
	return s
}

func (s *CostCenterSaveResponseBody) SetMessage(v string) *CostCenterSaveResponseBody {
	s.Message = &v
	return s
}

func (s *CostCenterSaveResponseBody) SetModule(v *CostCenterSaveResponseBodyModule) *CostCenterSaveResponseBody {
	s.Module = v
	return s
}

func (s *CostCenterSaveResponseBody) SetRequestId(v string) *CostCenterSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CostCenterSaveResponseBody) SetSuccess(v bool) *CostCenterSaveResponseBody {
	s.Success = &v
	return s
}

func (s *CostCenterSaveResponseBody) SetTraceId(v string) *CostCenterSaveResponseBody {
	s.TraceId = &v
	return s
}

type CostCenterSaveResponseBodyModule struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CostCenterSaveResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CostCenterSaveResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CostCenterSaveResponseBodyModule) SetId(v int64) *CostCenterSaveResponseBodyModule {
	s.Id = &v
	return s
}

type CostCenterSaveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CostCenterSaveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CostCenterSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s CostCenterSaveResponse) GoString() string {
	return s.String()
}

func (s *CostCenterSaveResponse) SetHeaders(v map[string]*string) *CostCenterSaveResponse {
	s.Headers = v
	return s
}

func (s *CostCenterSaveResponse) SetStatusCode(v int32) *CostCenterSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *CostCenterSaveResponse) SetBody(v *CostCenterSaveResponseBody) *CostCenterSaveResponse {
	s.Body = v
	return s
}

type DepartmentSaveHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s DepartmentSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s DepartmentSaveHeaders) GoString() string {
	return s.String()
}

func (s *DepartmentSaveHeaders) SetCommonHeaders(v map[string]*string) *DepartmentSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DepartmentSaveHeaders) SetXAcsBtripSoCorpToken(v string) *DepartmentSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type DepartmentSaveRequest struct {
	DepartList []*DepartmentSaveRequestDepartList `json:"depart_list,omitempty" xml:"depart_list,omitempty" type:"Repeated"`
}

func (s DepartmentSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s DepartmentSaveRequest) GoString() string {
	return s.String()
}

func (s *DepartmentSaveRequest) SetDepartList(v []*DepartmentSaveRequestDepartList) *DepartmentSaveRequest {
	s.DepartList = v
	return s
}

type DepartmentSaveRequestDepartList struct {
	DepartId       *int64  `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName     *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	DepartPid      *int64  `json:"depart_pid,omitempty" xml:"depart_pid,omitempty"`
	ManagerIds     *string `json:"manager_ids,omitempty" xml:"manager_ids,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdDepartId  *string `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
	ThirdDepartPid *string `json:"third_depart_pid,omitempty" xml:"third_depart_pid,omitempty"`
}

func (s DepartmentSaveRequestDepartList) String() string {
	return tea.Prettify(s)
}

func (s DepartmentSaveRequestDepartList) GoString() string {
	return s.String()
}

func (s *DepartmentSaveRequestDepartList) SetDepartId(v int64) *DepartmentSaveRequestDepartList {
	s.DepartId = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetDepartName(v string) *DepartmentSaveRequestDepartList {
	s.DepartName = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetDepartPid(v int64) *DepartmentSaveRequestDepartList {
	s.DepartPid = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetManagerIds(v string) *DepartmentSaveRequestDepartList {
	s.ManagerIds = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetStatus(v int32) *DepartmentSaveRequestDepartList {
	s.Status = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetThirdDepartId(v string) *DepartmentSaveRequestDepartList {
	s.ThirdDepartId = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetThirdDepartPid(v string) *DepartmentSaveRequestDepartList {
	s.ThirdDepartPid = &v
	return s
}

type DepartmentSaveShrinkRequest struct {
	DepartListShrink *string `json:"depart_list,omitempty" xml:"depart_list,omitempty"`
}

func (s DepartmentSaveShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DepartmentSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *DepartmentSaveShrinkRequest) SetDepartListShrink(v string) *DepartmentSaveShrinkRequest {
	s.DepartListShrink = &v
	return s
}

type DepartmentSaveResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *string `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DepartmentSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DepartmentSaveResponseBody) GoString() string {
	return s.String()
}

func (s *DepartmentSaveResponseBody) SetCode(v int32) *DepartmentSaveResponseBody {
	s.Code = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetMessage(v string) *DepartmentSaveResponseBody {
	s.Message = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetModule(v string) *DepartmentSaveResponseBody {
	s.Module = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetRequestId(v string) *DepartmentSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetSuccess(v bool) *DepartmentSaveResponseBody {
	s.Success = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetTraceId(v string) *DepartmentSaveResponseBody {
	s.TraceId = &v
	return s
}

type DepartmentSaveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DepartmentSaveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DepartmentSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s DepartmentSaveResponse) GoString() string {
	return s.String()
}

func (s *DepartmentSaveResponse) SetHeaders(v map[string]*string) *DepartmentSaveResponse {
	s.Headers = v
	return s
}

func (s *DepartmentSaveResponse) SetStatusCode(v int32) *DepartmentSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *DepartmentSaveResponse) SetBody(v *DepartmentSaveResponseBody) *DepartmentSaveResponse {
	s.Body = v
	return s
}

type EntityAddHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s EntityAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s EntityAddHeaders) GoString() string {
	return s.String()
}

func (s *EntityAddHeaders) SetCommonHeaders(v map[string]*string) *EntityAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EntityAddHeaders) SetXAcsBtripSoCorpToken(v string) *EntityAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type EntityAddRequest struct {
	EntityDOList []*EntityAddRequestEntityDOList `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty" type:"Repeated"`
	ThirdpartId  *string                         `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntityAddRequest) String() string {
	return tea.Prettify(s)
}

func (s EntityAddRequest) GoString() string {
	return s.String()
}

func (s *EntityAddRequest) SetEntityDOList(v []*EntityAddRequestEntityDOList) *EntityAddRequest {
	s.EntityDOList = v
	return s
}

func (s *EntityAddRequest) SetThirdpartId(v string) *EntityAddRequest {
	s.ThirdpartId = &v
	return s
}

type EntityAddRequestEntityDOList struct {
	EntityId   *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s EntityAddRequestEntityDOList) String() string {
	return tea.Prettify(s)
}

func (s EntityAddRequestEntityDOList) GoString() string {
	return s.String()
}

func (s *EntityAddRequestEntityDOList) SetEntityId(v string) *EntityAddRequestEntityDOList {
	s.EntityId = &v
	return s
}

func (s *EntityAddRequestEntityDOList) SetEntityType(v string) *EntityAddRequestEntityDOList {
	s.EntityType = &v
	return s
}

type EntityAddShrinkRequest struct {
	EntityDOListShrink *string `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty"`
	ThirdpartId        *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntityAddShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EntityAddShrinkRequest) GoString() string {
	return s.String()
}

func (s *EntityAddShrinkRequest) SetEntityDOListShrink(v string) *EntityAddShrinkRequest {
	s.EntityDOListShrink = &v
	return s
}

func (s *EntityAddShrinkRequest) SetThirdpartId(v string) *EntityAddShrinkRequest {
	s.ThirdpartId = &v
	return s
}

type EntityAddResponseBody struct {
	Code      *int32                       `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                      `json:"message,omitempty" xml:"message,omitempty"`
	Module    *EntityAddResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                        `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                      `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s EntityAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntityAddResponseBody) GoString() string {
	return s.String()
}

func (s *EntityAddResponseBody) SetCode(v int32) *EntityAddResponseBody {
	s.Code = &v
	return s
}

func (s *EntityAddResponseBody) SetMessage(v string) *EntityAddResponseBody {
	s.Message = &v
	return s
}

func (s *EntityAddResponseBody) SetModule(v *EntityAddResponseBodyModule) *EntityAddResponseBody {
	s.Module = v
	return s
}

func (s *EntityAddResponseBody) SetRequestId(v string) *EntityAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntityAddResponseBody) SetSuccess(v bool) *EntityAddResponseBody {
	s.Success = &v
	return s
}

func (s *EntityAddResponseBody) SetTraceId(v string) *EntityAddResponseBody {
	s.TraceId = &v
	return s
}

type EntityAddResponseBodyModule struct {
	AddNum          *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s EntityAddResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s EntityAddResponseBodyModule) GoString() string {
	return s.String()
}

func (s *EntityAddResponseBodyModule) SetAddNum(v int32) *EntityAddResponseBodyModule {
	s.AddNum = &v
	return s
}

func (s *EntityAddResponseBodyModule) SetSelectedUserNum(v int32) *EntityAddResponseBodyModule {
	s.SelectedUserNum = &v
	return s
}

type EntityAddResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EntityAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EntityAddResponse) String() string {
	return tea.Prettify(s)
}

func (s EntityAddResponse) GoString() string {
	return s.String()
}

func (s *EntityAddResponse) SetHeaders(v map[string]*string) *EntityAddResponse {
	s.Headers = v
	return s
}

func (s *EntityAddResponse) SetStatusCode(v int32) *EntityAddResponse {
	s.StatusCode = &v
	return s
}

func (s *EntityAddResponse) SetBody(v *EntityAddResponseBody) *EntityAddResponse {
	s.Body = v
	return s
}

type EntityDeleteHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s EntityDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s EntityDeleteHeaders) GoString() string {
	return s.String()
}

func (s *EntityDeleteHeaders) SetCommonHeaders(v map[string]*string) *EntityDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EntityDeleteHeaders) SetXAcsBtripSoCorpToken(v string) *EntityDeleteHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type EntityDeleteRequest struct {
	DelAll       *bool                              `json:"del_all,omitempty" xml:"del_all,omitempty"`
	EntityDOList []*EntityDeleteRequestEntityDOList `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty" type:"Repeated"`
	ThirdpartId  *string                            `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntityDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s EntityDeleteRequest) GoString() string {
	return s.String()
}

func (s *EntityDeleteRequest) SetDelAll(v bool) *EntityDeleteRequest {
	s.DelAll = &v
	return s
}

func (s *EntityDeleteRequest) SetEntityDOList(v []*EntityDeleteRequestEntityDOList) *EntityDeleteRequest {
	s.EntityDOList = v
	return s
}

func (s *EntityDeleteRequest) SetThirdpartId(v string) *EntityDeleteRequest {
	s.ThirdpartId = &v
	return s
}

type EntityDeleteRequestEntityDOList struct {
	EntityId   *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s EntityDeleteRequestEntityDOList) String() string {
	return tea.Prettify(s)
}

func (s EntityDeleteRequestEntityDOList) GoString() string {
	return s.String()
}

func (s *EntityDeleteRequestEntityDOList) SetEntityId(v string) *EntityDeleteRequestEntityDOList {
	s.EntityId = &v
	return s
}

func (s *EntityDeleteRequestEntityDOList) SetEntityType(v string) *EntityDeleteRequestEntityDOList {
	s.EntityType = &v
	return s
}

type EntityDeleteShrinkRequest struct {
	DelAll             *bool   `json:"del_all,omitempty" xml:"del_all,omitempty"`
	EntityDOListShrink *string `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty"`
	ThirdpartId        *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntityDeleteShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EntityDeleteShrinkRequest) GoString() string {
	return s.String()
}

func (s *EntityDeleteShrinkRequest) SetDelAll(v bool) *EntityDeleteShrinkRequest {
	s.DelAll = &v
	return s
}

func (s *EntityDeleteShrinkRequest) SetEntityDOListShrink(v string) *EntityDeleteShrinkRequest {
	s.EntityDOListShrink = &v
	return s
}

func (s *EntityDeleteShrinkRequest) SetThirdpartId(v string) *EntityDeleteShrinkRequest {
	s.ThirdpartId = &v
	return s
}

type EntityDeleteResponseBody struct {
	Code      *int32                          `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                         `json:"message,omitempty" xml:"message,omitempty"`
	Module    *EntityDeleteResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	MorePage  *bool                           `json:"more_page,omitempty" xml:"more_page,omitempty"`
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                           `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                         `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s EntityDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntityDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *EntityDeleteResponseBody) SetCode(v int32) *EntityDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *EntityDeleteResponseBody) SetMessage(v string) *EntityDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *EntityDeleteResponseBody) SetModule(v *EntityDeleteResponseBodyModule) *EntityDeleteResponseBody {
	s.Module = v
	return s
}

func (s *EntityDeleteResponseBody) SetMorePage(v bool) *EntityDeleteResponseBody {
	s.MorePage = &v
	return s
}

func (s *EntityDeleteResponseBody) SetRequestId(v string) *EntityDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntityDeleteResponseBody) SetSuccess(v bool) *EntityDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *EntityDeleteResponseBody) SetTraceId(v string) *EntityDeleteResponseBody {
	s.TraceId = &v
	return s
}

type EntityDeleteResponseBodyModule struct {
	RemoveNum       *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
	SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s EntityDeleteResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s EntityDeleteResponseBodyModule) GoString() string {
	return s.String()
}

func (s *EntityDeleteResponseBodyModule) SetRemoveNum(v int32) *EntityDeleteResponseBodyModule {
	s.RemoveNum = &v
	return s
}

func (s *EntityDeleteResponseBodyModule) SetSelectedUserNum(v int32) *EntityDeleteResponseBodyModule {
	s.SelectedUserNum = &v
	return s
}

type EntityDeleteResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EntityDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EntityDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s EntityDeleteResponse) GoString() string {
	return s.String()
}

func (s *EntityDeleteResponse) SetHeaders(v map[string]*string) *EntityDeleteResponse {
	s.Headers = v
	return s
}

func (s *EntityDeleteResponse) SetStatusCode(v int32) *EntityDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *EntityDeleteResponse) SetBody(v *EntityDeleteResponseBody) *EntityDeleteResponse {
	s.Body = v
	return s
}

type EntitySetHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s EntitySetHeaders) String() string {
	return tea.Prettify(s)
}

func (s EntitySetHeaders) GoString() string {
	return s.String()
}

func (s *EntitySetHeaders) SetCommonHeaders(v map[string]*string) *EntitySetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EntitySetHeaders) SetXAcsBtripSoCorpToken(v string) *EntitySetHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type EntitySetRequest struct {
	EntityDOList []*EntitySetRequestEntityDOList `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty" type:"Repeated"`
	ThirdpartId  *string                         `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntitySetRequest) String() string {
	return tea.Prettify(s)
}

func (s EntitySetRequest) GoString() string {
	return s.String()
}

func (s *EntitySetRequest) SetEntityDOList(v []*EntitySetRequestEntityDOList) *EntitySetRequest {
	s.EntityDOList = v
	return s
}

func (s *EntitySetRequest) SetThirdpartId(v string) *EntitySetRequest {
	s.ThirdpartId = &v
	return s
}

type EntitySetRequestEntityDOList struct {
	EntityId   *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s EntitySetRequestEntityDOList) String() string {
	return tea.Prettify(s)
}

func (s EntitySetRequestEntityDOList) GoString() string {
	return s.String()
}

func (s *EntitySetRequestEntityDOList) SetEntityId(v string) *EntitySetRequestEntityDOList {
	s.EntityId = &v
	return s
}

func (s *EntitySetRequestEntityDOList) SetEntityType(v string) *EntitySetRequestEntityDOList {
	s.EntityType = &v
	return s
}

type EntitySetShrinkRequest struct {
	EntityDOListShrink *string `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty"`
	ThirdpartId        *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntitySetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EntitySetShrinkRequest) GoString() string {
	return s.String()
}

func (s *EntitySetShrinkRequest) SetEntityDOListShrink(v string) *EntitySetShrinkRequest {
	s.EntityDOListShrink = &v
	return s
}

func (s *EntitySetShrinkRequest) SetThirdpartId(v string) *EntitySetShrinkRequest {
	s.ThirdpartId = &v
	return s
}

type EntitySetResponseBody struct {
	Code      *int32                       `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                      `json:"message,omitempty" xml:"message,omitempty"`
	Module    *EntitySetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	MorePage  *bool                        `json:"more_page,omitempty" xml:"more_page,omitempty"`
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                        `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                      `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s EntitySetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntitySetResponseBody) GoString() string {
	return s.String()
}

func (s *EntitySetResponseBody) SetCode(v int32) *EntitySetResponseBody {
	s.Code = &v
	return s
}

func (s *EntitySetResponseBody) SetMessage(v string) *EntitySetResponseBody {
	s.Message = &v
	return s
}

func (s *EntitySetResponseBody) SetModule(v *EntitySetResponseBodyModule) *EntitySetResponseBody {
	s.Module = v
	return s
}

func (s *EntitySetResponseBody) SetMorePage(v bool) *EntitySetResponseBody {
	s.MorePage = &v
	return s
}

func (s *EntitySetResponseBody) SetRequestId(v string) *EntitySetResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntitySetResponseBody) SetSuccess(v bool) *EntitySetResponseBody {
	s.Success = &v
	return s
}

func (s *EntitySetResponseBody) SetTraceId(v string) *EntitySetResponseBody {
	s.TraceId = &v
	return s
}

type EntitySetResponseBodyModule struct {
	AddNum          *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	RemoveNum       *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
	SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s EntitySetResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s EntitySetResponseBodyModule) GoString() string {
	return s.String()
}

func (s *EntitySetResponseBodyModule) SetAddNum(v int32) *EntitySetResponseBodyModule {
	s.AddNum = &v
	return s
}

func (s *EntitySetResponseBodyModule) SetRemoveNum(v int32) *EntitySetResponseBodyModule {
	s.RemoveNum = &v
	return s
}

func (s *EntitySetResponseBodyModule) SetSelectedUserNum(v int32) *EntitySetResponseBodyModule {
	s.SelectedUserNum = &v
	return s
}

type EntitySetResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EntitySetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EntitySetResponse) String() string {
	return tea.Prettify(s)
}

func (s EntitySetResponse) GoString() string {
	return s.String()
}

func (s *EntitySetResponse) SetHeaders(v map[string]*string) *EntitySetResponse {
	s.Headers = v
	return s
}

func (s *EntitySetResponse) SetStatusCode(v int32) *EntitySetResponse {
	s.StatusCode = &v
	return s
}

func (s *EntitySetResponse) SetBody(v *EntitySetResponseBody) *EntitySetResponse {
	s.Body = v
	return s
}

type ExceedApplySyncHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ExceedApplySyncHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExceedApplySyncHeaders) GoString() string {
	return s.String()
}

func (s *ExceedApplySyncHeaders) SetCommonHeaders(v map[string]*string) *ExceedApplySyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExceedApplySyncHeaders) SetXAcsBtripSoCorpToken(v string) *ExceedApplySyncHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ExceedApplySyncRequest struct {
	ApplyId          *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BizCategory      *int32  `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
	Remark           *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Status           *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdpartyFlowId *string `json:"thirdparty_flow_id,omitempty" xml:"thirdparty_flow_id,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ExceedApplySyncRequest) String() string {
	return tea.Prettify(s)
}

func (s ExceedApplySyncRequest) GoString() string {
	return s.String()
}

func (s *ExceedApplySyncRequest) SetApplyId(v int64) *ExceedApplySyncRequest {
	s.ApplyId = &v
	return s
}

func (s *ExceedApplySyncRequest) SetBizCategory(v int32) *ExceedApplySyncRequest {
	s.BizCategory = &v
	return s
}

func (s *ExceedApplySyncRequest) SetRemark(v string) *ExceedApplySyncRequest {
	s.Remark = &v
	return s
}

func (s *ExceedApplySyncRequest) SetStatus(v int32) *ExceedApplySyncRequest {
	s.Status = &v
	return s
}

func (s *ExceedApplySyncRequest) SetThirdpartyFlowId(v string) *ExceedApplySyncRequest {
	s.ThirdpartyFlowId = &v
	return s
}

func (s *ExceedApplySyncRequest) SetUserId(v string) *ExceedApplySyncRequest {
	s.UserId = &v
	return s
}

type ExceedApplySyncResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *bool   `json:"module,omitempty" xml:"module,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ExceedApplySyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExceedApplySyncResponseBody) GoString() string {
	return s.String()
}

func (s *ExceedApplySyncResponseBody) SetRequestId(v string) *ExceedApplySyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExceedApplySyncResponseBody) SetCode(v int32) *ExceedApplySyncResponseBody {
	s.Code = &v
	return s
}

func (s *ExceedApplySyncResponseBody) SetMessage(v string) *ExceedApplySyncResponseBody {
	s.Message = &v
	return s
}

func (s *ExceedApplySyncResponseBody) SetModule(v bool) *ExceedApplySyncResponseBody {
	s.Module = &v
	return s
}

func (s *ExceedApplySyncResponseBody) SetSuccess(v bool) *ExceedApplySyncResponseBody {
	s.Success = &v
	return s
}

func (s *ExceedApplySyncResponseBody) SetTraceId(v string) *ExceedApplySyncResponseBody {
	s.TraceId = &v
	return s
}

type ExceedApplySyncResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExceedApplySyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExceedApplySyncResponse) String() string {
	return tea.Prettify(s)
}

func (s ExceedApplySyncResponse) GoString() string {
	return s.String()
}

func (s *ExceedApplySyncResponse) SetHeaders(v map[string]*string) *ExceedApplySyncResponse {
	s.Headers = v
	return s
}

func (s *ExceedApplySyncResponse) SetStatusCode(v int32) *ExceedApplySyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ExceedApplySyncResponse) SetBody(v *ExceedApplySyncResponseBody) *ExceedApplySyncResponse {
	s.Body = v
	return s
}

type FlightBillSettlementQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightBillSettlementQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s FlightBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type FlightBillSettlementQueryRequest struct {
	PageNo      *int32  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize    *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PeriodEnd   *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
}

func (s FlightBillSettlementQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s FlightBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryRequest) SetPageNo(v int32) *FlightBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetPageSize(v int32) *FlightBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetPeriodEnd(v string) *FlightBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetPeriodStart(v string) *FlightBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

type FlightBillSettlementQueryResponseBody struct {
	Code      *int32                                       `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                      `json:"message,omitempty" xml:"message,omitempty"`
	Module    *FlightBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                      `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightBillSettlementQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlightBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryResponseBody) SetCode(v int32) *FlightBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetMessage(v string) *FlightBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetModule(v *FlightBillSettlementQueryResponseBodyModule) *FlightBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetRequestId(v string) *FlightBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetSuccess(v bool) *FlightBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetTraceId(v string) *FlightBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

type FlightBillSettlementQueryResponseBodyModule struct {
	Category    *int32                                                 `json:"category,omitempty" xml:"category,omitempty"`
	CorpId      *string                                                `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList    []*FlightBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
	PeriodEnd   *string                                                `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string                                                `json:"period_start,omitempty" xml:"period_start,omitempty"`
	TotalNum    *int64                                                 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s FlightBillSettlementQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s FlightBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetCategory(v int32) *FlightBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetCorpId(v string) *FlightBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetDataList(v []*FlightBillSettlementQueryResponseBodyModuleDataList) *FlightBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *FlightBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *FlightBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *FlightBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

type FlightBillSettlementQueryResponseBodyModuleDataList struct {
	AdvanceDay             *int32   `json:"advance_day,omitempty" xml:"advance_day,omitempty"`
	AirlineCorpCode        *string  `json:"airline_corp_code,omitempty" xml:"airline_corp_code,omitempty"`
	AirlineCorpName        *string  `json:"airline_corp_name,omitempty" xml:"airline_corp_name,omitempty"`
	AlipayTradeNo          *string  `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyId                *string  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrAirportCode         *string  `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrCity                *string  `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrDate                *string  `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	ArrStation             *string  `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	ArrTime                *string  `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BillRecordTime         *string  `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookTime               *string  `json:"book_time,omitempty" xml:"book_time,omitempty"`
	BookerId               *string  `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	BookerJobNo            *string  `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName             *string  `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	BtripCouponFee         *float64 `json:"btrip_coupon_fee,omitempty" xml:"btrip_coupon_fee,omitempty"`
	BuildFee               *float64 `json:"build_fee,omitempty" xml:"build_fee,omitempty"`
	Cabin                  *string  `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass             *string  `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CapitalDirection       *string  `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment      *string  `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	ChangeFee              *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	CorpPayOrderFee        *float64 `json:"corp_pay_order_fee,omitempty" xml:"corp_pay_order_fee,omitempty"`
	CostCenter             *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	CostCenterNumber       *string  `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	Coupon                 *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	DepAirportCode         *string  `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	Department             *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId           *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptCity               *string  `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	DeptDate               *string  `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptStation            *string  `json:"dept_station,omitempty" xml:"dept_station,omitempty"`
	DeptTime               *string  `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	Discount               *string  `json:"discount,omitempty" xml:"discount,omitempty"`
	FeeType                *string  `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FlightNo               *string  `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	Index                  *string  `json:"index,omitempty" xml:"index,omitempty"`
	InsuranceFee           *float64 `json:"insurance_fee,omitempty" xml:"insurance_fee,omitempty"`
	InvoiceTitle           *string  `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	ItineraryNum           *string  `json:"itinerary_num,omitempty" xml:"itinerary_num,omitempty"`
	ItineraryPrice         *float64 `json:"itinerary_price,omitempty" xml:"itinerary_price,omitempty"`
	MostDifferenceDeptTime *string  `json:"most_difference_dept_time,omitempty" xml:"most_difference_dept_time,omitempty"`
	MostDifferenceDiscount *string  `json:"most_difference_discount,omitempty" xml:"most_difference_discount,omitempty"`
	MostDifferenceFlightNo *string  `json:"most_difference_flight_no,omitempty" xml:"most_difference_flight_no,omitempty"`
	MostDifferencePrice    *float64 `json:"most_difference_price,omitempty" xml:"most_difference_price,omitempty"`
	MostDifferenceReason   *string  `json:"most_difference_reason,omitempty" xml:"most_difference_reason,omitempty"`
	MostPrice              *float64 `json:"most_price,omitempty" xml:"most_price,omitempty"`
	NegotiationCouponFee   *float64 `json:"negotiation_coupon_fee,omitempty" xml:"negotiation_coupon_fee,omitempty"`
	OilFee                 *float64 `json:"oil_fee,omitempty" xml:"oil_fee,omitempty"`
	OrderId                *string  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OverApplyId            *string  `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PrimaryId              *int64   `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProjectCode            *string  `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName            *string  `json:"project_name,omitempty" xml:"project_name,omitempty"`
	RefundFee              *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	RefundUpgradeCost      *float64 `json:"refund_upgrade_cost,omitempty" xml:"refund_upgrade_cost,omitempty"`
	Remark                 *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	RepeatRefund           *string  `json:"repeat_refund,omitempty" xml:"repeat_refund,omitempty"`
	SealPrice              *float64 `json:"seal_price,omitempty" xml:"seal_price,omitempty"`
	ServiceFee             *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettlementFee          *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	SettlementGrantFee     *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	SettlementTime         *string  `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	SettlementType         *string  `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	Status                 *int32   `json:"status,omitempty" xml:"status,omitempty"`
	TicketId               *string  `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	TravelerId             *string  `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	TravelerJobNo          *string  `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerName           *string  `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	UpgradeCost            *float64 `json:"upgrade_cost,omitempty" xml:"upgrade_cost,omitempty"`
	VoucherType            *int32   `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}

func (s FlightBillSettlementQueryResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s FlightBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAdvanceDay(v int32) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AdvanceDay = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAirlineCorpCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AirlineCorpCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAirlineCorpName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AirlineCorpName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrAirportCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrCity(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCity = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrDate(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrStation(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrStation = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBtripCouponFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BtripCouponFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBuildFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BuildFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCabin(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Cabin = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCabinClass(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CabinClass = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetChangeFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCorpPayOrderFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CorpPayOrderFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCoupon(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDepAirportCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeptCity(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptCity = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeptDate(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeptStation(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptStation = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeptTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDiscount(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Discount = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetFlightNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.FlightNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetInsuranceFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetItineraryNum(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ItineraryNum = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetItineraryPrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ItineraryPrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceDeptTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceDeptTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceDiscount(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceDiscount = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceFlightNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceFlightNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferencePrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferencePrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceReason(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceReason = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostPrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostPrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetNegotiationCouponFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.NegotiationCouponFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetOilFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.OilFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRefundFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRefundUpgradeCost(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundUpgradeCost = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRepeatRefund(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RepeatRefund = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSealPrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SealPrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTicketId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TicketId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetUpgradeCost(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.UpgradeCost = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

type FlightBillSettlementQueryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FlightBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlightBillSettlementQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s FlightBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryResponse) SetHeaders(v map[string]*string) *FlightBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightBillSettlementQueryResponse) SetStatusCode(v int32) *FlightBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponse) SetBody(v *FlightBillSettlementQueryResponseBody) *FlightBillSettlementQueryResponse {
	s.Body = v
	return s
}

type FlightExceedApplyQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightExceedApplyQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s FlightExceedApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightExceedApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightExceedApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightExceedApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type FlightExceedApplyQueryRequest struct {
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
}

func (s FlightExceedApplyQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s FlightExceedApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryRequest) SetApplyId(v int64) *FlightExceedApplyQueryRequest {
	s.ApplyId = &v
	return s
}

type FlightExceedApplyQueryResponseBody struct {
	Code      *int32                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module    *FlightExceedApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                   `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightExceedApplyQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlightExceedApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponseBody) SetCode(v int32) *FlightExceedApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetMessage(v string) *FlightExceedApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetModule(v *FlightExceedApplyQueryResponseBodyModule) *FlightExceedApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetRequestId(v string) *FlightExceedApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetSuccess(v bool) *FlightExceedApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetTraceId(v string) *FlightExceedApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

type FlightExceedApplyQueryResponseBodyModule struct {
	ApplyId              *int64                                                        `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApplyIntentionInfoDo *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo `json:"apply_intention_info_do,omitempty" xml:"apply_intention_info_do,omitempty" type:"Struct"`
	BtripCause           *string                                                       `json:"btrip_cause,omitempty" xml:"btrip_cause,omitempty"`
	CorpId               *string                                                       `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	ExceedReason         *string                                                       `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	ExceedType           *int32                                                        `json:"exceed_type,omitempty" xml:"exceed_type,omitempty"`
	OriginStandard       *string                                                       `json:"origin_standard,omitempty" xml:"origin_standard,omitempty"`
	Status               *int32                                                        `json:"status,omitempty" xml:"status,omitempty"`
	SubmitTime           *string                                                       `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
	ThirdpartApplyId     *string                                                       `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartCorpId      *string                                                       `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	UserId               *string                                                       `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightExceedApplyQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s FlightExceedApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetApplyId(v int64) *FlightExceedApplyQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetApplyIntentionInfoDo(v *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) *FlightExceedApplyQueryResponseBodyModule {
	s.ApplyIntentionInfoDo = v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetBtripCause(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.BtripCause = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetCorpId(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetExceedReason(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.ExceedReason = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetExceedType(v int32) *FlightExceedApplyQueryResponseBodyModule {
	s.ExceedType = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetOriginStandard(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.OriginStandard = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetStatus(v int32) *FlightExceedApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetSubmitTime(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.SubmitTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetThirdpartApplyId(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetThirdpartCorpId(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.ThirdpartCorpId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetUserId(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

type FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo struct {
	ArrCity       *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityName   *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTime       *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	Cabin         *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass    *int32  `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassStr *string `json:"cabin_class_str,omitempty" xml:"cabin_class_str,omitempty"`
	DepCity       *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityName   *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTime       *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Discount      *string `json:"discount,omitempty" xml:"discount,omitempty"`
	FlightNo      *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	Price         *int64  `json:"price,omitempty" xml:"price,omitempty"`
	Type          *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) String() string {
	return tea.Prettify(s)
}

func (s FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetArrCity(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.ArrCity = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetArrCityName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.ArrCityName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetArrTime(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.ArrTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCabin(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Cabin = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCabinClass(v int32) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CabinClass = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCabinClassStr(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CabinClassStr = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDepCity(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.DepCity = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDepCityName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.DepCityName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDepTime(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.DepTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDiscount(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Discount = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetFlightNo(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.FlightNo = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetPrice(v int64) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Price = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetType(v int32) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Type = &v
	return s
}

type FlightExceedApplyQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FlightExceedApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlightExceedApplyQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s FlightExceedApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponse) SetHeaders(v map[string]*string) *FlightExceedApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightExceedApplyQueryResponse) SetStatusCode(v int32) *FlightExceedApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightExceedApplyQueryResponse) SetBody(v *FlightExceedApplyQueryResponseBody) *FlightExceedApplyQueryResponse {
	s.Body = v
	return s
}

type FlightOrderListQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightOrderListQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightOrderListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightOrderListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type FlightOrderListQueryRequest struct {
	AllApply         *bool   `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	ApplyId          *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	DepartId         *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Page             *int32  `json:"page,omitempty" xml:"page,omitempty"`
	PageSize         *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	UpdateEndTime    *string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	UpdateStartTime  *string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderListQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryRequest) SetAllApply(v bool) *FlightOrderListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetApplyId(v int64) *FlightOrderListQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetDepartId(v string) *FlightOrderListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetEndTime(v string) *FlightOrderListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetPage(v int32) *FlightOrderListQueryRequest {
	s.Page = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetPageSize(v int32) *FlightOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetStartTime(v string) *FlightOrderListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetThirdpartApplyId(v string) *FlightOrderListQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetUpdateEndTime(v string) *FlightOrderListQueryRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetUpdateStartTime(v string) *FlightOrderListQueryRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetUserId(v string) *FlightOrderListQueryRequest {
	s.UserId = &v
	return s
}

type FlightOrderListQueryResponseBody struct {
	Code      *int32                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*FlightOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo  *FlightOrderListQueryResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                   `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightOrderListQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBody) SetCode(v int32) *FlightOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetMessage(v string) *FlightOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetModule(v []*FlightOrderListQueryResponseBodyModule) *FlightOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetPageInfo(v *FlightOrderListQueryResponseBodyPageInfo) *FlightOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetRequestId(v string) *FlightOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetSuccess(v bool) *FlightOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetTraceId(v string) *FlightOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

type FlightOrderListQueryResponseBodyModule struct {
	ApplyId              *int64                                                     `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrAirport           *string                                                    `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	ArrCity              *string                                                    `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	BtripTitle           *string                                                    `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	CabinClass           *string                                                    `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	ContactName          *string                                                    `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId               *string                                                    `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName             *string                                                    `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	CostCenter           *FlightOrderListQueryResponseBodyModuleCostCenter          `json:"cost_center,omitempty" xml:"cost_center,omitempty" type:"Struct"`
	DepAirport           *string                                                    `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	DepCity              *string                                                    `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepDate              *string                                                    `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	DepartId             *string                                                    `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName           *string                                                    `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	Discount             *string                                                    `json:"discount,omitempty" xml:"discount,omitempty"`
	FlightNo             *string                                                    `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	GmtCreate            *string                                                    `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModified          *string                                                    `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	Id                   *int64                                                     `json:"id,omitempty" xml:"id,omitempty"`
	InsureInfoList       []*FlightOrderListQueryResponseBodyModuleInsureInfoList    `json:"insure_info_list,omitempty" xml:"insure_info_list,omitempty" type:"Repeated"`
	Invoice              *FlightOrderListQueryResponseBodyModuleInvoice             `json:"invoice,omitempty" xml:"invoice,omitempty" type:"Struct"`
	PassengerCount       *int32                                                     `json:"passenger_count,omitempty" xml:"passenger_count,omitempty"`
	PassengerName        *string                                                    `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PriceInfoList        []*FlightOrderListQueryResponseBodyModulePriceInfoList     `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	ProjectCode          *string                                                    `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectId            *int64                                                     `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle         *string                                                    `json:"project_title,omitempty" xml:"project_title,omitempty"`
	RetDate              *string                                                    `json:"ret_date,omitempty" xml:"ret_date,omitempty"`
	Status               *int32                                                     `json:"status,omitempty" xml:"status,omitempty"`
	ThirdPartProjectId   *string                                                    `json:"third_part_project_id,omitempty" xml:"third_part_project_id,omitempty"`
	ThirdpartApplyId     *string                                                    `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartItineraryId *string                                                    `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	TripType             *int32                                                     `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	UserAffiliateList    []*FlightOrderListQueryResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	UserId               *string                                                    `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName             *string                                                    `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModule) SetApplyId(v int64) *FlightOrderListQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetArrAirport(v string) *FlightOrderListQueryResponseBodyModule {
	s.ArrAirport = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetArrCity(v string) *FlightOrderListQueryResponseBodyModule {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetBtripTitle(v string) *FlightOrderListQueryResponseBodyModule {
	s.BtripTitle = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetCabinClass(v string) *FlightOrderListQueryResponseBodyModule {
	s.CabinClass = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetContactName(v string) *FlightOrderListQueryResponseBodyModule {
	s.ContactName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetCorpId(v string) *FlightOrderListQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetCorpName(v string) *FlightOrderListQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetCostCenter(v *FlightOrderListQueryResponseBodyModuleCostCenter) *FlightOrderListQueryResponseBodyModule {
	s.CostCenter = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepAirport(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepAirport = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepCity(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepCity = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepDate(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepDate = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepartId(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepartName(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDiscount(v string) *FlightOrderListQueryResponseBodyModule {
	s.Discount = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetFlightNo(v string) *FlightOrderListQueryResponseBodyModule {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetGmtCreate(v string) *FlightOrderListQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetGmtModified(v string) *FlightOrderListQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetId(v int64) *FlightOrderListQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetInsureInfoList(v []*FlightOrderListQueryResponseBodyModuleInsureInfoList) *FlightOrderListQueryResponseBodyModule {
	s.InsureInfoList = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetInvoice(v *FlightOrderListQueryResponseBodyModuleInvoice) *FlightOrderListQueryResponseBodyModule {
	s.Invoice = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetPassengerCount(v int32) *FlightOrderListQueryResponseBodyModule {
	s.PassengerCount = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetPassengerName(v string) *FlightOrderListQueryResponseBodyModule {
	s.PassengerName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetPriceInfoList(v []*FlightOrderListQueryResponseBodyModulePriceInfoList) *FlightOrderListQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetProjectCode(v string) *FlightOrderListQueryResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetProjectId(v int64) *FlightOrderListQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetProjectTitle(v string) *FlightOrderListQueryResponseBodyModule {
	s.ProjectTitle = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetRetDate(v string) *FlightOrderListQueryResponseBodyModule {
	s.RetDate = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetStatus(v int32) *FlightOrderListQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetThirdPartProjectId(v string) *FlightOrderListQueryResponseBodyModule {
	s.ThirdPartProjectId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetThirdpartApplyId(v string) *FlightOrderListQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetThirdpartItineraryId(v string) *FlightOrderListQueryResponseBodyModule {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetTripType(v int32) *FlightOrderListQueryResponseBodyModule {
	s.TripType = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetUserAffiliateList(v []*FlightOrderListQueryResponseBodyModuleUserAffiliateList) *FlightOrderListQueryResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetUserId(v string) *FlightOrderListQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetUserName(v string) *FlightOrderListQueryResponseBodyModule {
	s.UserName = &v
	return s
}

type FlightOrderListQueryResponseBodyModuleCostCenter struct {
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Id     *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModuleCostCenter) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModuleCostCenter) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) SetCorpId(v string) *FlightOrderListQueryResponseBodyModuleCostCenter {
	s.CorpId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) SetId(v int64) *FlightOrderListQueryResponseBodyModuleCostCenter {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) SetName(v string) *FlightOrderListQueryResponseBodyModuleCostCenter {
	s.Name = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) SetNumber(v string) *FlightOrderListQueryResponseBodyModuleCostCenter {
	s.Number = &v
	return s
}

type FlightOrderListQueryResponseBodyModuleInsureInfoList struct {
	InsureNo *string `json:"insure_no,omitempty" xml:"insure_no,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Status   *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModuleInsureInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModuleInsureInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) SetInsureNo(v string) *FlightOrderListQueryResponseBodyModuleInsureInfoList {
	s.InsureNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) SetName(v string) *FlightOrderListQueryResponseBodyModuleInsureInfoList {
	s.Name = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) SetStatus(v int32) *FlightOrderListQueryResponseBodyModuleInsureInfoList {
	s.Status = &v
	return s
}

type FlightOrderListQueryResponseBodyModuleInvoice struct {
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModuleInvoice) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModuleInvoice) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModuleInvoice) SetId(v int64) *FlightOrderListQueryResponseBodyModuleInvoice {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleInvoice) SetTitle(v string) *FlightOrderListQueryResponseBodyModuleInvoice {
	s.Title = &v
	return s
}

type FlightOrderListQueryResponseBodyModulePriceInfoList struct {
	CategoryCode     *int32   `json:"category_code,omitempty" xml:"category_code,omitempty"`
	CategoryType     *int32   `json:"category_type,omitempty" xml:"category_type,omitempty"`
	ChangeFlightNo   *string  `json:"change_flight_no,omitempty" xml:"change_flight_no,omitempty"`
	Discount         *string  `json:"discount,omitempty" xml:"discount,omitempty"`
	EndTime          *string  `json:"end_time,omitempty" xml:"end_time,omitempty"`
	GmtCreate        *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	OriginalTicketNo *string  `json:"original_ticket_no,omitempty" xml:"original_ticket_no,omitempty"`
	PassengerName    *string  `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PayType          *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	Price            *float64 `json:"price,omitempty" xml:"price,omitempty"`
	StartTime        *string  `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TicketNo         *string  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TradeId          *string  `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	Type             *int32   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModulePriceInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetCategoryType(v int32) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetChangeFlightNo(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.ChangeFlightNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetDiscount(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.Discount = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetEndTime(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.EndTime = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetOriginalTicketNo(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.OriginalTicketNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetStartTime(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.StartTime = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetTicketNo(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetType(v int32) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

type FlightOrderListQueryResponseBodyModuleUserAffiliateList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModuleUserAffiliateList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModuleUserAffiliateList) SetUserId(v string) *FlightOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleUserAffiliateList) SetUserName(v string) *FlightOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

type FlightOrderListQueryResponseBodyPageInfo struct {
	Page        *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageSize    *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s FlightOrderListQueryResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyPageInfo) SetPage(v int32) *FlightOrderListQueryResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyPageInfo) SetPageSize(v int32) *FlightOrderListQueryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *FlightOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

type FlightOrderListQueryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FlightOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlightOrderListQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponse) SetHeaders(v map[string]*string) *FlightOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightOrderListQueryResponse) SetStatusCode(v int32) *FlightOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightOrderListQueryResponse) SetBody(v *FlightOrderListQueryResponseBody) *FlightOrderListQueryResponse {
	s.Body = v
	return s
}

type FlightOrderQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightOrderQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightOrderQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightOrderQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightOrderQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type FlightOrderQueryRequest struct {
	OrderId *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	UserId  *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryRequest) SetOrderId(v int64) *FlightOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *FlightOrderQueryRequest) SetUserId(v string) *FlightOrderQueryRequest {
	s.UserId = &v
	return s
}

type FlightOrderQueryResponseBody struct {
	Code      *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module    *FlightOrderQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                             `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightOrderQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBody) SetCode(v int32) *FlightOrderQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOrderQueryResponseBody) SetMessage(v string) *FlightOrderQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOrderQueryResponseBody) SetModule(v *FlightOrderQueryResponseBodyModule) *FlightOrderQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightOrderQueryResponseBody) SetRequestId(v string) *FlightOrderQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOrderQueryResponseBody) SetSuccess(v bool) *FlightOrderQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOrderQueryResponseBody) SetTraceId(v string) *FlightOrderQueryResponseBody {
	s.TraceId = &v
	return s
}

type FlightOrderQueryResponseBodyModule struct {
	FlightChangeTicketInfoList []*FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList `json:"flight_change_ticket_info_list,omitempty" xml:"flight_change_ticket_info_list,omitempty" type:"Repeated"`
	FlightInfoList             []*FlightOrderQueryResponseBodyModuleFlightInfoList             `json:"flight_info_list,omitempty" xml:"flight_info_list,omitempty" type:"Repeated"`
	FlightRefundTicketInfoList []*FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList `json:"flight_refund_ticket_info_list,omitempty" xml:"flight_refund_ticket_info_list,omitempty" type:"Repeated"`
	FlightTicketInfoList       []*FlightOrderQueryResponseBodyModuleFlightTicketInfoList       `json:"flight_ticket_info_list,omitempty" xml:"flight_ticket_info_list,omitempty" type:"Repeated"`
	InsuranceInfoList          []*FlightOrderQueryResponseBodyModuleInsuranceInfoList          `json:"insurance_info_list,omitempty" xml:"insurance_info_list,omitempty" type:"Repeated"`
	InvoiceInfo                *FlightOrderQueryResponseBodyModuleInvoiceInfo                  `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	OrderBaseInfo              *FlightOrderQueryResponseBodyModuleOrderBaseInfo                `json:"order_base_info,omitempty" xml:"order_base_info,omitempty" type:"Struct"`
	PassengerInfoList          []*FlightOrderQueryResponseBodyModulePassengerInfoList          `json:"passenger_info_list,omitempty" xml:"passenger_info_list,omitempty" type:"Repeated"`
	PriceInfoList              []*FlightOrderQueryResponseBodyModulePriceInfoList              `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
}

func (s FlightOrderQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModule) SetFlightChangeTicketInfoList(v []*FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) *FlightOrderQueryResponseBodyModule {
	s.FlightChangeTicketInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetFlightInfoList(v []*FlightOrderQueryResponseBodyModuleFlightInfoList) *FlightOrderQueryResponseBodyModule {
	s.FlightInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetFlightRefundTicketInfoList(v []*FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) *FlightOrderQueryResponseBodyModule {
	s.FlightRefundTicketInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetFlightTicketInfoList(v []*FlightOrderQueryResponseBodyModuleFlightTicketInfoList) *FlightOrderQueryResponseBodyModule {
	s.FlightTicketInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetInsuranceInfoList(v []*FlightOrderQueryResponseBodyModuleInsuranceInfoList) *FlightOrderQueryResponseBodyModule {
	s.InsuranceInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetInvoiceInfo(v *FlightOrderQueryResponseBodyModuleInvoiceInfo) *FlightOrderQueryResponseBodyModule {
	s.InvoiceInfo = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetOrderBaseInfo(v *FlightOrderQueryResponseBodyModuleOrderBaseInfo) *FlightOrderQueryResponseBodyModule {
	s.OrderBaseInfo = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetPassengerInfoList(v []*FlightOrderQueryResponseBodyModulePassengerInfoList) *FlightOrderQueryResponseBodyModule {
	s.PassengerInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetPriceInfoList(v []*FlightOrderQueryResponseBodyModulePriceInfoList) *FlightOrderQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

type FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList struct {
	ArrTime          *string  `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	ChangeCabin      *string  `json:"change_cabin,omitempty" xml:"change_cabin,omitempty"`
	ChangeCabinLevel *string  `json:"change_cabin_level,omitempty" xml:"change_cabin_level,omitempty"`
	ChangeFee        *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	ChangeFlightNo   *string  `json:"change_flight_no,omitempty" xml:"change_flight_no,omitempty"`
	ChangeOrderId    *int64   `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	ChangeReason     *string  `json:"change_reason,omitempty" xml:"change_reason,omitempty"`
	ChangeType       *int32   `json:"change_type,omitempty" xml:"change_type,omitempty"`
	DepTime          *string  `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	GmtCreate        *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify        *string  `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	OriginTicketNo   *string  `json:"origin_ticket_no,omitempty" xml:"origin_ticket_no,omitempty"`
	TicketNo         *string  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	UpgradeFee       *float64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetArrTime(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeCabin(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeCabin = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeCabinLevel(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeCabinLevel = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeFee(v float64) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeFee = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeFlightNo(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeFlightNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeOrderId(v int64) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeOrderId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeReason(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeReason = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeType(v int32) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetDepTime(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetGmtModify(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetOriginTicketNo(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.OriginTicketNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetTicketNo(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetUpgradeFee(v float64) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.UpgradeFee = &v
	return s
}

type FlightOrderQueryResponseBodyModuleFlightInfoList struct {
	AirlineCode    *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName    *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrAirportName *string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
	ArrCityCode    *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName    *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTime        *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	Cabin          *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinLevel     *string `json:"cabin_level,omitempty" xml:"cabin_level,omitempty"`
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepAirportName *string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
	DepCityCode    *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName    *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTime        *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FlightMile     *int32  `json:"flight_mile,omitempty" xml:"flight_mile,omitempty"`
	FlightNo       *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetAirlineCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.AirlineCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetAirlineName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.AirlineName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrAirportName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrAirportName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrTime(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetCabin(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.Cabin = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetCabinLevel(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.CabinLevel = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepAirportName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepAirportName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepTime(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetFlightMile(v int32) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.FlightMile = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetFlightNo(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.FlightNo = &v
	return s
}

type FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList struct {
	GmtCreate       *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify       *string  `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	RefundOrderId   *int64   `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	RefundReason    *string  `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	RefundTicketFee *float64 `json:"refund_ticket_fee,omitempty" xml:"refund_ticket_fee,omitempty"`
	RefundType      *int32   `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	TicketNo        *string  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetGmtModify(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetRefundOrderId(v int64) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.RefundOrderId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetRefundReason(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.RefundReason = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetRefundTicketFee(v float64) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.RefundTicketFee = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetRefundType(v int32) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.RefundType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetTicketNo(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.TicketNo = &v
	return s
}

type FlightOrderQueryResponseBodyModuleFlightTicketInfoList struct {
	BuildPrice       *float64 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	Changed          *bool    `json:"changed,omitempty" xml:"changed,omitempty"`
	Discount         *int32   `json:"discount,omitempty" xml:"discount,omitempty"`
	GmtCreate        *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify        *string  `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	OilPrice         *float64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	PayType          *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	SettlePrice      *float64 `json:"settle_price,omitempty" xml:"settle_price,omitempty"`
	TicketNo         *string  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketPrice      *float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	TicketStatus     *string  `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	TicketStatusCode *int32   `json:"ticket_status_code,omitempty" xml:"ticket_status_code,omitempty"`
	UserId           *string  `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightTicketInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetBuildPrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.BuildPrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetChanged(v bool) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.Changed = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetDiscount(v int32) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.Discount = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetGmtModify(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetOilPrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.OilPrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetPayType(v int32) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.PayType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetSettlePrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.SettlePrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetTicketNo(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetTicketPrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.TicketPrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetTicketStatus(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.TicketStatus = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetTicketStatusCode(v int32) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.TicketStatusCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetUserId(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.UserId = &v
	return s
}

type FlightOrderQueryResponseBodyModuleInsuranceInfoList struct {
	Amount      *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	InsuranceNo *string  `json:"insurance_no,omitempty" xml:"insurance_no,omitempty"`
	Status      *int32   `json:"status,omitempty" xml:"status,omitempty"`
	Type        *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleInsuranceInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleInsuranceInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetAmount(v float64) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.Amount = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetInsuranceNo(v string) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.InsuranceNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetStatus(v int32) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.Status = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetType(v string) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.Type = &v
	return s
}

type FlightOrderQueryResponseBodyModuleInvoiceInfo struct {
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleInvoiceInfo) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleInvoiceInfo) SetId(v int64) *FlightOrderQueryResponseBodyModuleInvoiceInfo {
	s.Id = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInvoiceInfo) SetTitle(v string) *FlightOrderQueryResponseBodyModuleInvoiceInfo {
	s.Title = &v
	return s
}

type FlightOrderQueryResponseBodyModuleOrderBaseInfo struct {
	ApplyId              *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripTitle           *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	ContactName          *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId               *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName             *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId             *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName           *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	GmtCreate            *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify            *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	ItineraryId          *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	OrderId              *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderStatus          *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	ThirdpartApplyId     *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartCorpId      *string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	TripType             *int32  `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	UserId               *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleOrderBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetApplyId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetBtripTitle(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.BtripTitle = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetContactName(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ContactName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpName(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartName(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtModify(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtModify = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetItineraryId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ItineraryId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderId(v int64) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderStatus(v int32) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderStatus = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartApplyId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartCorpId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartCorpId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartItineraryId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetTripType(v int32) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.TripType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetUserId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.UserId = &v
	return s
}

type FlightOrderQueryResponseBodyModulePassengerInfoList struct {
	CostCenterId       *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName     *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	CostCenterNumber   *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	ProjectCode        *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectId          *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle       *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdpartProjectId *string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	UserId             *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName           *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	UserType           *int32  `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s FlightOrderQueryResponseBodyModulePassengerInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModulePassengerInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterId(v int64) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterName(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterNumber(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterNumber = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetProjectCode(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetProjectId(v int64) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetProjectTitle(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectTitle = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetThirdpartProjectId(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.ThirdpartProjectId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetUserId(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.UserId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetUserName(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.UserName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetUserType(v int32) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.UserType = &v
	return s
}

type FlightOrderQueryResponseBodyModulePriceInfoList struct {
	CategoryCode  *int32   `json:"category_code,omitempty" xml:"category_code,omitempty"`
	GmtCreate     *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string  `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PayType       *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	Price         *float64 `json:"price,omitempty" xml:"price,omitempty"`
	TradeId       *string  `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	Type          *int32   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderQueryResponseBodyModulePriceInfoList) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetType(v int32) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

type FlightOrderQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FlightOrderQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlightOrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s FlightOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponse) SetHeaders(v map[string]*string) *FlightOrderQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightOrderQueryResponse) SetStatusCode(v int32) *FlightOrderQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightOrderQueryResponse) SetBody(v *FlightOrderQueryResponseBody) *FlightOrderQueryResponse {
	s.Body = v
	return s
}

type HotelBillSettlementQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s HotelBillSettlementQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HotelBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *HotelBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type HotelBillSettlementQueryRequest struct {
	PageNo      *int32  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize    *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PeriodEnd   *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
}

func (s HotelBillSettlementQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HotelBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryRequest) SetPageNo(v int32) *HotelBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetPageSize(v int32) *HotelBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetPeriodEnd(v string) *HotelBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetPeriodStart(v string) *HotelBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

type HotelBillSettlementQueryResponseBody struct {
	Code      *int32                                      `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	Module    *HotelBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                     `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelBillSettlementQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HotelBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryResponseBody) SetCode(v int32) *HotelBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetMessage(v string) *HotelBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetModule(v *HotelBillSettlementQueryResponseBodyModule) *HotelBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetRequestId(v string) *HotelBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetSuccess(v bool) *HotelBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetTraceId(v string) *HotelBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

type HotelBillSettlementQueryResponseBodyModule struct {
	Category    *int32                                                `json:"category,omitempty" xml:"category,omitempty"`
	CorpId      *string                                               `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList    []*HotelBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
	PeriodEnd   *string                                               `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string                                               `json:"period_start,omitempty" xml:"period_start,omitempty"`
	TotalNum    *int64                                                `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s HotelBillSettlementQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s HotelBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetCategory(v int32) *HotelBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetCorpId(v string) *HotelBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetDataList(v []*HotelBillSettlementQueryResponseBodyModuleDataList) *HotelBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *HotelBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *HotelBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *HotelBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

type HotelBillSettlementQueryResponseBodyModuleDataList struct {
	AlipayTradeNo      *string  `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyId            *string  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BillRecordTime     *string  `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookTime           *string  `json:"book_time,omitempty" xml:"book_time,omitempty"`
	BookerId           *string  `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	BookerJobNo        *string  `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName         *string  `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	CapitalDirection   *string  `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment  *string  `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CheckInDate        *string  `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	CheckoutDate       *string  `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	City               *string  `json:"city,omitempty" xml:"city,omitempty"`
	CityCode           *string  `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CorpRefundFee      *float64 `json:"corp_refund_fee,omitempty" xml:"corp_refund_fee,omitempty"`
	CorpTotalFee       *float64 `json:"corp_total_fee,omitempty" xml:"corp_total_fee,omitempty"`
	CostCenter         *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	CostCenterNumber   *string  `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	Department         *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId       *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	FeeType            *string  `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	Fees               *float64 `json:"fees,omitempty" xml:"fees,omitempty"`
	FuPointFee         *float64 `json:"fu_point_fee,omitempty" xml:"fu_point_fee,omitempty"`
	HotelName          *string  `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	Index              *string  `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle       *string  `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	IsNegotiation      *string  `json:"is_negotiation,omitempty" xml:"is_negotiation,omitempty"`
	IsShareStr         *string  `json:"is_share_str,omitempty" xml:"is_share_str,omitempty"`
	Nights             *int32   `json:"nights,omitempty" xml:"nights,omitempty"`
	OrderId            *string  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderPrice         *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OrderType          *string  `json:"order_type,omitempty" xml:"order_type,omitempty"`
	OverApplyId        *string  `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PersonRefundFee    *float64 `json:"person_refund_fee,omitempty" xml:"person_refund_fee,omitempty"`
	PersonSettlePrice  *float64 `json:"person_settle_price,omitempty" xml:"person_settle_price,omitempty"`
	PrimaryId          *int64   `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProjectCode        *string  `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName        *string  `json:"project_name,omitempty" xml:"project_name,omitempty"`
	PromotionFee       *float64 `json:"promotion_fee,omitempty" xml:"promotion_fee,omitempty"`
	Remark             *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	RoomNumber         *int32   `json:"room_number,omitempty" xml:"room_number,omitempty"`
	RoomPrice          *float64 `json:"room_price,omitempty" xml:"room_price,omitempty"`
	RoomType           *string  `json:"room_type,omitempty" xml:"room_type,omitempty"`
	ServiceFee         *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettlementFee      *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	SettlementGrantFee *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	SettlementTime     *string  `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	SettlementType     *string  `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	Status             *int32   `json:"status,omitempty" xml:"status,omitempty"`
	TotalNights        *int32   `json:"total_nights,omitempty" xml:"total_nights,omitempty"`
	TravelerId         *string  `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	TravelerJobNo      *string  `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerName       *string  `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	VoucherType        *int32   `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}

func (s HotelBillSettlementQueryResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s HotelBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCheckInDate(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CheckInDate = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCheckoutDate(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CheckoutDate = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCity(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.City = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCityCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CityCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCorpRefundFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CorpRefundFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCorpTotalFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CorpTotalFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetFees(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Fees = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetFuPointFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.FuPointFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetHotelName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.HotelName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetIsNegotiation(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.IsNegotiation = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetIsShareStr(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.IsShareStr = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetNights(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Nights = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOrderPrice(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOrderType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPersonRefundFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PersonRefundFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPersonSettlePrice(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PersonSettlePrice = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPromotionFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PromotionFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRoomNumber(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomNumber = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRoomPrice(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomPrice = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRoomType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTotalNights(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TotalNights = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

type HotelBillSettlementQueryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HotelBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HotelBillSettlementQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HotelBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryResponse) SetHeaders(v map[string]*string) *HotelBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelBillSettlementQueryResponse) SetStatusCode(v int32) *HotelBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponse) SetBody(v *HotelBillSettlementQueryResponseBody) *HotelBillSettlementQueryResponse {
	s.Body = v
	return s
}

type HotelExceedApplyQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s HotelExceedApplyQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HotelExceedApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelExceedApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelExceedApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *HotelExceedApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type HotelExceedApplyQueryRequest struct {
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
}

func (s HotelExceedApplyQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HotelExceedApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryRequest) SetApplyId(v int64) *HotelExceedApplyQueryRequest {
	s.ApplyId = &v
	return s
}

type HotelExceedApplyQueryResponseBody struct {
	Code      *int32                                   `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module    *HotelExceedApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelExceedApplyQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HotelExceedApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryResponseBody) SetCode(v int32) *HotelExceedApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetMessage(v string) *HotelExceedApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetModule(v *HotelExceedApplyQueryResponseBodyModule) *HotelExceedApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetRequestId(v string) *HotelExceedApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetSuccess(v bool) *HotelExceedApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetTraceId(v string) *HotelExceedApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

type HotelExceedApplyQueryResponseBodyModule struct {
	ApplyId              *int64                                                       `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApplyIntentionInfoDo *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo `json:"apply_intention_info_do,omitempty" xml:"apply_intention_info_do,omitempty" type:"Struct"`
	BtripCause           *string                                                      `json:"btrip_cause,omitempty" xml:"btrip_cause,omitempty"`
	CorpId               *string                                                      `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	ExceedReason         *string                                                      `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	ExceedType           *int32                                                       `json:"exceed_type,omitempty" xml:"exceed_type,omitempty"`
	OriginStandard       *string                                                      `json:"origin_standard,omitempty" xml:"origin_standard,omitempty"`
	Status               *int32                                                       `json:"status,omitempty" xml:"status,omitempty"`
	SubmitTime           *string                                                      `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
	ThirdpartApplyId     *string                                                      `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartCorpId      *string                                                      `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	UserId               *string                                                      `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s HotelExceedApplyQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s HotelExceedApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetApplyId(v int64) *HotelExceedApplyQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetApplyIntentionInfoDo(v *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) *HotelExceedApplyQueryResponseBodyModule {
	s.ApplyIntentionInfoDo = v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetBtripCause(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.BtripCause = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetCorpId(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetExceedReason(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.ExceedReason = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetExceedType(v int32) *HotelExceedApplyQueryResponseBodyModule {
	s.ExceedType = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetOriginStandard(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.OriginStandard = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetStatus(v int32) *HotelExceedApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetSubmitTime(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.SubmitTime = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetThirdpartApplyId(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetThirdpartCorpId(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.ThirdpartCorpId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetUserId(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

type HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo struct {
	CheckIn  *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Price    *int64  `json:"price,omitempty" xml:"price,omitempty"`
	Together *bool   `json:"together,omitempty" xml:"together,omitempty"`
	Type     *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) String() string {
	return tea.Prettify(s)
}

func (s HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCheckIn(v string) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CheckIn = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCheckOut(v string) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CheckOut = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCityCode(v string) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CityCode = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCityName(v string) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CityName = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetPrice(v int64) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Price = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetTogether(v bool) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Together = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetType(v int32) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Type = &v
	return s
}

type HotelExceedApplyQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HotelExceedApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HotelExceedApplyQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HotelExceedApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryResponse) SetHeaders(v map[string]*string) *HotelExceedApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelExceedApplyQueryResponse) SetStatusCode(v int32) *HotelExceedApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelExceedApplyQueryResponse) SetBody(v *HotelExceedApplyQueryResponseBody) *HotelExceedApplyQueryResponse {
	s.Body = v
	return s
}

type HotelOrderListQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s HotelOrderListQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *HotelOrderListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type HotelOrderListQueryRequest struct {
	AllApply         *bool   `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	ApplyId          *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	DepartId         *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Page             *int32  `json:"page,omitempty" xml:"page,omitempty"`
	PageSize         *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	UpdateEndTime    *string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	UpdateStartTime  *string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s HotelOrderListQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryRequest) SetAllApply(v bool) *HotelOrderListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetApplyId(v int64) *HotelOrderListQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetDepartId(v string) *HotelOrderListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetEndTime(v string) *HotelOrderListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetPage(v int32) *HotelOrderListQueryRequest {
	s.Page = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetPageSize(v int32) *HotelOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetStartTime(v string) *HotelOrderListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetThirdpartApplyId(v string) *HotelOrderListQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetUpdateEndTime(v string) *HotelOrderListQueryRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetUpdateStartTime(v string) *HotelOrderListQueryRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetUserId(v string) *HotelOrderListQueryRequest {
	s.UserId = &v
	return s
}

type HotelOrderListQueryResponseBody struct {
	Code      *int32                                   `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*HotelOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo  *HotelOrderListQueryResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelOrderListQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBody) SetCode(v int32) *HotelOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetMessage(v string) *HotelOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetModule(v []*HotelOrderListQueryResponseBodyModule) *HotelOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetPageInfo(v *HotelOrderListQueryResponseBodyPageInfo) *HotelOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetRequestId(v string) *HotelOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetSuccess(v bool) *HotelOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetTraceId(v string) *HotelOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

type HotelOrderListQueryResponseBodyModule struct {
	ApplyId                    *int64                                                    `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripTitle                 *string                                                   `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	CheckIn                    *string                                                   `json:"check_in,omitempty" xml:"check_in,omitempty"`
	CheckOut                   *string                                                   `json:"check_out,omitempty" xml:"check_out,omitempty"`
	City                       *string                                                   `json:"city,omitempty" xml:"city,omitempty"`
	ContactName                *string                                                   `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId                     *string                                                   `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName                   *string                                                   `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	CostCenter                 *HotelOrderListQueryResponseBodyModuleCostCenter          `json:"cost_center,omitempty" xml:"cost_center,omitempty" type:"Struct"`
	DepartId                   *string                                                   `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName                 *string                                                   `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	GmtCreate                  *string                                                   `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModified                *string                                                   `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	Guest                      *string                                                   `json:"guest,omitempty" xml:"guest,omitempty"`
	HotelName                  *string                                                   `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	HotelSupportVatInvoiceType *int32                                                    `json:"hotel_support_vat_invoice_type,omitempty" xml:"hotel_support_vat_invoice_type,omitempty"`
	Id                         *int64                                                    `json:"id,omitempty" xml:"id,omitempty"`
	Invoice                    *HotelOrderListQueryResponseBodyModuleInvoice             `json:"invoice,omitempty" xml:"invoice,omitempty" type:"Struct"`
	Night                      *int32                                                    `json:"night,omitempty" xml:"night,omitempty"`
	OrderStatus                *int32                                                    `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OrderStatusDesc            *string                                                   `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	OrderType                  *int32                                                    `json:"order_type,omitempty" xml:"order_type,omitempty"`
	OrderTypeDesc              *string                                                   `json:"order_type_desc,omitempty" xml:"order_type_desc,omitempty"`
	PriceInfoList              []*HotelOrderListQueryResponseBodyModulePriceInfoList     `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	ProjectCode                *string                                                   `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectId                  *int64                                                    `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle               *string                                                   `json:"project_title,omitempty" xml:"project_title,omitempty"`
	RoomNum                    *int32                                                    `json:"room_num,omitempty" xml:"room_num,omitempty"`
	RoomType                   *string                                                   `json:"room_type,omitempty" xml:"room_type,omitempty"`
	ThirdpartApplyId           *string                                                   `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartItineraryId       *string                                                   `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	ThirdpartProjectId         *string                                                   `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	UserAffiliateList          []*HotelOrderListQueryResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	UserId                     *string                                                   `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName                   *string                                                   `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModule) SetApplyId(v int64) *HotelOrderListQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetBtripTitle(v string) *HotelOrderListQueryResponseBodyModule {
	s.BtripTitle = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCheckIn(v string) *HotelOrderListQueryResponseBodyModule {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCheckOut(v string) *HotelOrderListQueryResponseBodyModule {
	s.CheckOut = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCity(v string) *HotelOrderListQueryResponseBodyModule {
	s.City = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetContactName(v string) *HotelOrderListQueryResponseBodyModule {
	s.ContactName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCorpId(v string) *HotelOrderListQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCorpName(v string) *HotelOrderListQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCostCenter(v *HotelOrderListQueryResponseBodyModuleCostCenter) *HotelOrderListQueryResponseBodyModule {
	s.CostCenter = v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetDepartId(v string) *HotelOrderListQueryResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetDepartName(v string) *HotelOrderListQueryResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetGmtCreate(v string) *HotelOrderListQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetGmtModified(v string) *HotelOrderListQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetGuest(v string) *HotelOrderListQueryResponseBodyModule {
	s.Guest = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetHotelName(v string) *HotelOrderListQueryResponseBodyModule {
	s.HotelName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetHotelSupportVatInvoiceType(v int32) *HotelOrderListQueryResponseBodyModule {
	s.HotelSupportVatInvoiceType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetId(v int64) *HotelOrderListQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetInvoice(v *HotelOrderListQueryResponseBodyModuleInvoice) *HotelOrderListQueryResponseBodyModule {
	s.Invoice = v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetNight(v int32) *HotelOrderListQueryResponseBodyModule {
	s.Night = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetOrderStatus(v int32) *HotelOrderListQueryResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetOrderStatusDesc(v string) *HotelOrderListQueryResponseBodyModule {
	s.OrderStatusDesc = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetOrderType(v int32) *HotelOrderListQueryResponseBodyModule {
	s.OrderType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetOrderTypeDesc(v string) *HotelOrderListQueryResponseBodyModule {
	s.OrderTypeDesc = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetPriceInfoList(v []*HotelOrderListQueryResponseBodyModulePriceInfoList) *HotelOrderListQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetProjectCode(v string) *HotelOrderListQueryResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetProjectId(v int64) *HotelOrderListQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetProjectTitle(v string) *HotelOrderListQueryResponseBodyModule {
	s.ProjectTitle = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetRoomNum(v int32) *HotelOrderListQueryResponseBodyModule {
	s.RoomNum = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetRoomType(v string) *HotelOrderListQueryResponseBodyModule {
	s.RoomType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetThirdpartApplyId(v string) *HotelOrderListQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetThirdpartItineraryId(v string) *HotelOrderListQueryResponseBodyModule {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetThirdpartProjectId(v string) *HotelOrderListQueryResponseBodyModule {
	s.ThirdpartProjectId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetUserAffiliateList(v []*HotelOrderListQueryResponseBodyModuleUserAffiliateList) *HotelOrderListQueryResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetUserId(v string) *HotelOrderListQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetUserName(v string) *HotelOrderListQueryResponseBodyModule {
	s.UserName = &v
	return s
}

type HotelOrderListQueryResponseBodyModuleCostCenter struct {
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Id     *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModuleCostCenter) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModuleCostCenter) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) SetCorpId(v string) *HotelOrderListQueryResponseBodyModuleCostCenter {
	s.CorpId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) SetId(v int64) *HotelOrderListQueryResponseBodyModuleCostCenter {
	s.Id = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) SetName(v string) *HotelOrderListQueryResponseBodyModuleCostCenter {
	s.Name = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) SetNumber(v string) *HotelOrderListQueryResponseBodyModuleCostCenter {
	s.Number = &v
	return s
}

type HotelOrderListQueryResponseBodyModuleInvoice struct {
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	InvoiceType *int32  `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModuleInvoice) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModuleInvoice) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) SetId(v int64) *HotelOrderListQueryResponseBodyModuleInvoice {
	s.Id = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) SetInvoiceType(v int32) *HotelOrderListQueryResponseBodyModuleInvoice {
	s.InvoiceType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) SetTitle(v string) *HotelOrderListQueryResponseBodyModuleInvoice {
	s.Title = &v
	return s
}

type HotelOrderListQueryResponseBodyModulePriceInfoList struct {
	CategoryCode  *int32   `json:"category_code,omitempty" xml:"category_code,omitempty"`
	CategoryType  *int32   `json:"category_type,omitempty" xml:"category_type,omitempty"`
	GmtCreate     *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string  `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PayType       *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	Price         *float64 `json:"price,omitempty" xml:"price,omitempty"`
	TradeId       *string  `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	Type          *int32   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModulePriceInfoList) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetCategoryType(v int32) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetType(v int32) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

type HotelOrderListQueryResponseBodyModuleUserAffiliateList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModuleUserAffiliateList) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModuleUserAffiliateList) SetUserId(v string) *HotelOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleUserAffiliateList) SetUserName(v string) *HotelOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

type HotelOrderListQueryResponseBodyPageInfo struct {
	Page        *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageSize    *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s HotelOrderListQueryResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyPageInfo) SetPage(v int32) *HotelOrderListQueryResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyPageInfo) SetPageSize(v int32) *HotelOrderListQueryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *HotelOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

type HotelOrderListQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HotelOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HotelOrderListQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponse) SetHeaders(v map[string]*string) *HotelOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderListQueryResponse) SetStatusCode(v int32) *HotelOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderListQueryResponse) SetBody(v *HotelOrderListQueryResponseBody) *HotelOrderListQueryResponse {
	s.Body = v
	return s
}

type IeFlightBillSettlementQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s IeFlightBillSettlementQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s IeFlightBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *IeFlightBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IeFlightBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *IeFlightBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type IeFlightBillSettlementQueryRequest struct {
	PageNo      *int32  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize    *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PeriodEnd   *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
}

func (s IeFlightBillSettlementQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s IeFlightBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryRequest) SetPageNo(v int32) *IeFlightBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetPageSize(v int32) *IeFlightBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetPeriodEnd(v string) *IeFlightBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetPeriodStart(v string) *IeFlightBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

type IeFlightBillSettlementQueryResponseBody struct {
	Code      *int32                                         `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                        `json:"message,omitempty" xml:"message,omitempty"`
	Module    *IeFlightBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	MorePage  *bool                                          `json:"more_page,omitempty" xml:"more_page,omitempty"`
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                          `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                        `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IeFlightBillSettlementQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IeFlightBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryResponseBody) SetCode(v int32) *IeFlightBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetMessage(v string) *IeFlightBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetModule(v *IeFlightBillSettlementQueryResponseBodyModule) *IeFlightBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetMorePage(v bool) *IeFlightBillSettlementQueryResponseBody {
	s.MorePage = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetRequestId(v string) *IeFlightBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetSuccess(v bool) *IeFlightBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetTraceId(v string) *IeFlightBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

type IeFlightBillSettlementQueryResponseBodyModule struct {
	Category    *int32                                                   `json:"category,omitempty" xml:"category,omitempty"`
	CorpId      *string                                                  `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList    []*IeFlightBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
	PeriodEnd   *string                                                  `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string                                                  `json:"period_start,omitempty" xml:"period_start,omitempty"`
	TotalNum    *int64                                                   `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s IeFlightBillSettlementQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s IeFlightBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetCategory(v int32) *IeFlightBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetCorpId(v string) *IeFlightBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetDataList(v []*IeFlightBillSettlementQueryResponseBodyModuleDataList) *IeFlightBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *IeFlightBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *IeFlightBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *IeFlightBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

type IeFlightBillSettlementQueryResponseBodyModuleDataList struct {
	AdvanceDay             *int32   `json:"advance_day,omitempty" xml:"advance_day,omitempty"`
	AirlineCorpCode        *string  `json:"airline_corp_code,omitempty" xml:"airline_corp_code,omitempty"`
	AirlineCorpName        *string  `json:"airline_corp_name,omitempty" xml:"airline_corp_name,omitempty"`
	AlipayTradeNo          *string  `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyId                *string  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrAirportCode         *string  `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrCity                *string  `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrDate                *string  `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	ArrStation             *string  `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	ArrTime                *string  `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BillRecordTime         *string  `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookMode               *string  `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	BookTime               *string  `json:"book_time,omitempty" xml:"book_time,omitempty"`
	BookerId               *string  `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	BookerJobNo            *string  `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName             *string  `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	BtripCouponFee         *float64 `json:"btrip_coupon_fee,omitempty" xml:"btrip_coupon_fee,omitempty"`
	Cabin                  *string  `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass             *string  `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CapitalDirection       *string  `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment      *string  `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	ChangeFee              *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	CorpPayOrderFee        *float64 `json:"corp_pay_order_fee,omitempty" xml:"corp_pay_order_fee,omitempty"`
	CostCenter             *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	CostCenterNumber       *string  `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	Coupon                 *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	DepAirportCode         *string  `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	Department             *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId           *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptCity               *string  `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	DeptDate               *string  `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptStation            *string  `json:"dept_station,omitempty" xml:"dept_station,omitempty"`
	DeptTime               *string  `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	Discount               *string  `json:"discount,omitempty" xml:"discount,omitempty"`
	FeeType                *string  `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FlightNo               *string  `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	Index                  *string  `json:"index,omitempty" xml:"index,omitempty"`
	InsuranceFee           *float64 `json:"insurance_fee,omitempty" xml:"insurance_fee,omitempty"`
	InsuranceNumber        *string  `json:"insurance_number,omitempty" xml:"insurance_number,omitempty"`
	InvoiceTitle           *string  `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	MostDifferenceDeptTime *string  `json:"most_difference_dept_time,omitempty" xml:"most_difference_dept_time,omitempty"`
	MostDifferenceDiscount *string  `json:"most_difference_discount,omitempty" xml:"most_difference_discount,omitempty"`
	MostDifferenceFlightNo *string  `json:"most_difference_flight_no,omitempty" xml:"most_difference_flight_no,omitempty"`
	MostDifferencePrice    *float64 `json:"most_difference_price,omitempty" xml:"most_difference_price,omitempty"`
	MostDifferenceReason   *string  `json:"most_difference_reason,omitempty" xml:"most_difference_reason,omitempty"`
	MostPrice              *float64 `json:"most_price,omitempty" xml:"most_price,omitempty"`
	NegotiationCouponFee   *float64 `json:"negotiation_coupon_fee,omitempty" xml:"negotiation_coupon_fee,omitempty"`
	OrderId                *string  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderStatusDesc        *string  `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	OverApplyId            *string  `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PrimaryId              *int64   `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProjectCode            *string  `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName            *string  `json:"project_name,omitempty" xml:"project_name,omitempty"`
	RefundFee              *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	Remark                 *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	RepeatRefund           *string  `json:"repeat_refund,omitempty" xml:"repeat_refund,omitempty"`
	SealPrice              *float64 `json:"seal_price,omitempty" xml:"seal_price,omitempty"`
	SegmentType            *string  `json:"segment_type,omitempty" xml:"segment_type,omitempty"`
	ServiceFee             *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettlementFee          *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	SettlementGrantFee     *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	SettlementTime         *string  `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	SettlementType         *string  `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	Status                 *int32   `json:"status,omitempty" xml:"status,omitempty"`
	SubOrderId             *string  `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	TaxFee                 *float64 `json:"tax_fee,omitempty" xml:"tax_fee,omitempty"`
	TicketId               *string  `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	Trade                  *string  `json:"trade,omitempty" xml:"trade,omitempty"`
	TravelerId             *string  `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	TravelerJobNo          *string  `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerName           *string  `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	VoucherType            *int32   `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}

func (s IeFlightBillSettlementQueryResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s IeFlightBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAdvanceDay(v int32) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AdvanceDay = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAirlineCorpCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AirlineCorpCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAirlineCorpName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AirlineCorpName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrAirportCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrAirportCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrCity(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCity = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrDate(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrStation(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrStation = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookMode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookMode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBtripCouponFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BtripCouponFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCabin(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Cabin = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCabinClass(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CabinClass = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetChangeFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCorpPayOrderFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CorpPayOrderFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCoupon(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepAirportCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepAirportCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeptCity(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptCity = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeptDate(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeptStation(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptStation = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeptTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDiscount(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Discount = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetFlightNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.FlightNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetInsuranceFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetInsuranceNumber(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceNumber = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceDeptTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceDeptTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceDiscount(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceDiscount = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceFlightNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceFlightNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferencePrice(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferencePrice = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceReason(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceReason = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostPrice(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostPrice = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetNegotiationCouponFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.NegotiationCouponFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetOrderStatusDesc(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.OrderStatusDesc = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetRefundFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetRepeatRefund(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.RepeatRefund = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSealPrice(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SealPrice = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSegmentType(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SegmentType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSubOrderId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SubOrderId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTaxFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TaxFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTicketId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TicketId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTrade(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Trade = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

type IeFlightBillSettlementQueryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IeFlightBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IeFlightBillSettlementQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s IeFlightBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryResponse) SetHeaders(v map[string]*string) *IeFlightBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *IeFlightBillSettlementQueryResponse) SetStatusCode(v int32) *IeFlightBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponse) SetBody(v *IeFlightBillSettlementQueryResponseBody) *IeFlightBillSettlementQueryResponse {
	s.Body = v
	return s
}

type InvoiceAddHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvoiceAddHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceAddHeaders) SetCommonHeaders(v map[string]*string) *InvoiceAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceAddHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type InvoiceAddRequest struct {
	Address     *string `json:"address,omitempty" xml:"address,omitempty"`
	BankName    *string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	BankNo      *string `json:"bank_no,omitempty" xml:"bank_no,omitempty"`
	TaxNo       *string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
	Tel         *string `json:"tel,omitempty" xml:"tel,omitempty"`
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Type        *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvoiceAddRequest) String() string {
	return tea.Prettify(s)
}

func (s InvoiceAddRequest) GoString() string {
	return s.String()
}

func (s *InvoiceAddRequest) SetAddress(v string) *InvoiceAddRequest {
	s.Address = &v
	return s
}

func (s *InvoiceAddRequest) SetBankName(v string) *InvoiceAddRequest {
	s.BankName = &v
	return s
}

func (s *InvoiceAddRequest) SetBankNo(v string) *InvoiceAddRequest {
	s.BankNo = &v
	return s
}

func (s *InvoiceAddRequest) SetTaxNo(v string) *InvoiceAddRequest {
	s.TaxNo = &v
	return s
}

func (s *InvoiceAddRequest) SetTel(v string) *InvoiceAddRequest {
	s.Tel = &v
	return s
}

func (s *InvoiceAddRequest) SetThirdPartId(v string) *InvoiceAddRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceAddRequest) SetTitle(v string) *InvoiceAddRequest {
	s.Title = &v
	return s
}

func (s *InvoiceAddRequest) SetType(v int32) *InvoiceAddRequest {
	s.Type = &v
	return s
}

type InvoiceAddResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InvoiceAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvoiceAddResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceAddResponseBody) SetCode(v int32) *InvoiceAddResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceAddResponseBody) SetMessage(v string) *InvoiceAddResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceAddResponseBody) SetRequestId(v string) *InvoiceAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceAddResponseBody) SetSuccess(v bool) *InvoiceAddResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceAddResponseBody) SetTraceId(v string) *InvoiceAddResponseBody {
	s.TraceId = &v
	return s
}

type InvoiceAddResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvoiceAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvoiceAddResponse) String() string {
	return tea.Prettify(s)
}

func (s InvoiceAddResponse) GoString() string {
	return s.String()
}

func (s *InvoiceAddResponse) SetHeaders(v map[string]*string) *InvoiceAddResponse {
	s.Headers = v
	return s
}

func (s *InvoiceAddResponse) SetStatusCode(v int32) *InvoiceAddResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceAddResponse) SetBody(v *InvoiceAddResponseBody) *InvoiceAddResponse {
	s.Body = v
	return s
}

type InvoiceDeleteHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvoiceDeleteHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceDeleteHeaders) SetCommonHeaders(v map[string]*string) *InvoiceDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceDeleteHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceDeleteHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type InvoiceDeleteRequest struct {
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s InvoiceDeleteRequest) GoString() string {
	return s.String()
}

func (s *InvoiceDeleteRequest) SetThirdPartId(v string) *InvoiceDeleteRequest {
	s.ThirdPartId = &v
	return s
}

type InvoiceDeleteResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InvoiceDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvoiceDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceDeleteResponseBody) SetCode(v int32) *InvoiceDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceDeleteResponseBody) SetMessage(v string) *InvoiceDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceDeleteResponseBody) SetRequestId(v string) *InvoiceDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceDeleteResponseBody) SetSuccess(v bool) *InvoiceDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceDeleteResponseBody) SetTraceId(v string) *InvoiceDeleteResponseBody {
	s.TraceId = &v
	return s
}

type InvoiceDeleteResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvoiceDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvoiceDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s InvoiceDeleteResponse) GoString() string {
	return s.String()
}

func (s *InvoiceDeleteResponse) SetHeaders(v map[string]*string) *InvoiceDeleteResponse {
	s.Headers = v
	return s
}

func (s *InvoiceDeleteResponse) SetStatusCode(v int32) *InvoiceDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceDeleteResponse) SetBody(v *InvoiceDeleteResponseBody) *InvoiceDeleteResponse {
	s.Body = v
	return s
}

type InvoiceModifyHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceModifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvoiceModifyHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceModifyHeaders) SetCommonHeaders(v map[string]*string) *InvoiceModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceModifyHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type InvoiceModifyRequest struct {
	Address     *string `json:"address,omitempty" xml:"address,omitempty"`
	BankName    *string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	BankNo      *string `json:"bank_no,omitempty" xml:"bank_no,omitempty"`
	TaxNo       *string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
	Tel         *string `json:"tel,omitempty" xml:"tel,omitempty"`
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Type        *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvoiceModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InvoiceModifyRequest) GoString() string {
	return s.String()
}

func (s *InvoiceModifyRequest) SetAddress(v string) *InvoiceModifyRequest {
	s.Address = &v
	return s
}

func (s *InvoiceModifyRequest) SetBankName(v string) *InvoiceModifyRequest {
	s.BankName = &v
	return s
}

func (s *InvoiceModifyRequest) SetBankNo(v string) *InvoiceModifyRequest {
	s.BankNo = &v
	return s
}

func (s *InvoiceModifyRequest) SetTaxNo(v string) *InvoiceModifyRequest {
	s.TaxNo = &v
	return s
}

func (s *InvoiceModifyRequest) SetTel(v string) *InvoiceModifyRequest {
	s.Tel = &v
	return s
}

func (s *InvoiceModifyRequest) SetThirdPartId(v string) *InvoiceModifyRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceModifyRequest) SetTitle(v string) *InvoiceModifyRequest {
	s.Title = &v
	return s
}

func (s *InvoiceModifyRequest) SetType(v int32) *InvoiceModifyRequest {
	s.Type = &v
	return s
}

type InvoiceModifyResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InvoiceModifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvoiceModifyResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceModifyResponseBody) SetCode(v int32) *InvoiceModifyResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceModifyResponseBody) SetMessage(v string) *InvoiceModifyResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceModifyResponseBody) SetRequestId(v string) *InvoiceModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceModifyResponseBody) SetSuccess(v bool) *InvoiceModifyResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceModifyResponseBody) SetTraceId(v string) *InvoiceModifyResponseBody {
	s.TraceId = &v
	return s
}

type InvoiceModifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvoiceModifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvoiceModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InvoiceModifyResponse) GoString() string {
	return s.String()
}

func (s *InvoiceModifyResponse) SetHeaders(v map[string]*string) *InvoiceModifyResponse {
	s.Headers = v
	return s
}

func (s *InvoiceModifyResponse) SetStatusCode(v int32) *InvoiceModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceModifyResponse) SetBody(v *InvoiceModifyResponseBody) *InvoiceModifyResponse {
	s.Body = v
	return s
}

type InvoiceRuleSaveHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceRuleSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvoiceRuleSaveHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveHeaders) SetCommonHeaders(v map[string]*string) *InvoiceRuleSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceRuleSaveHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceRuleSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type InvoiceRuleSaveRequest struct {
	AllEmploye  *bool                             `json:"all_employe,omitempty" xml:"all_employe,omitempty"`
	Entities    []*InvoiceRuleSaveRequestEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
	ThirdPartId *string                           `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceRuleSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s InvoiceRuleSaveRequest) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveRequest) SetAllEmploye(v bool) *InvoiceRuleSaveRequest {
	s.AllEmploye = &v
	return s
}

func (s *InvoiceRuleSaveRequest) SetEntities(v []*InvoiceRuleSaveRequestEntities) *InvoiceRuleSaveRequest {
	s.Entities = v
	return s
}

func (s *InvoiceRuleSaveRequest) SetThirdPartId(v string) *InvoiceRuleSaveRequest {
	s.ThirdPartId = &v
	return s
}

type InvoiceRuleSaveRequestEntities struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvoiceRuleSaveRequestEntities) String() string {
	return tea.Prettify(s)
}

func (s InvoiceRuleSaveRequestEntities) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveRequestEntities) SetId(v string) *InvoiceRuleSaveRequestEntities {
	s.Id = &v
	return s
}

func (s *InvoiceRuleSaveRequestEntities) SetName(v string) *InvoiceRuleSaveRequestEntities {
	s.Name = &v
	return s
}

func (s *InvoiceRuleSaveRequestEntities) SetType(v int32) *InvoiceRuleSaveRequestEntities {
	s.Type = &v
	return s
}

type InvoiceRuleSaveShrinkRequest struct {
	AllEmploye     *bool   `json:"all_employe,omitempty" xml:"all_employe,omitempty"`
	EntitiesShrink *string `json:"entities,omitempty" xml:"entities,omitempty"`
	ThirdPartId    *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceRuleSaveShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InvoiceRuleSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveShrinkRequest) SetAllEmploye(v bool) *InvoiceRuleSaveShrinkRequest {
	s.AllEmploye = &v
	return s
}

func (s *InvoiceRuleSaveShrinkRequest) SetEntitiesShrink(v string) *InvoiceRuleSaveShrinkRequest {
	s.EntitiesShrink = &v
	return s
}

func (s *InvoiceRuleSaveShrinkRequest) SetThirdPartId(v string) *InvoiceRuleSaveShrinkRequest {
	s.ThirdPartId = &v
	return s
}

type InvoiceRuleSaveResponseBody struct {
	Code      *int32                             `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module    *InvoiceRuleSaveResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                              `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                            `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InvoiceRuleSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvoiceRuleSaveResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveResponseBody) SetCode(v int32) *InvoiceRuleSaveResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetMessage(v string) *InvoiceRuleSaveResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetModule(v *InvoiceRuleSaveResponseBodyModule) *InvoiceRuleSaveResponseBody {
	s.Module = v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetRequestId(v string) *InvoiceRuleSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetSuccess(v bool) *InvoiceRuleSaveResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetTraceId(v string) *InvoiceRuleSaveResponseBody {
	s.TraceId = &v
	return s
}

type InvoiceRuleSaveResponseBodyModule struct {
	AddNum    *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	RemoveNum *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
}

func (s InvoiceRuleSaveResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s InvoiceRuleSaveResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveResponseBodyModule) SetAddNum(v int32) *InvoiceRuleSaveResponseBodyModule {
	s.AddNum = &v
	return s
}

func (s *InvoiceRuleSaveResponseBodyModule) SetRemoveNum(v int32) *InvoiceRuleSaveResponseBodyModule {
	s.RemoveNum = &v
	return s
}

type InvoiceRuleSaveResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvoiceRuleSaveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvoiceRuleSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s InvoiceRuleSaveResponse) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveResponse) SetHeaders(v map[string]*string) *InvoiceRuleSaveResponse {
	s.Headers = v
	return s
}

func (s *InvoiceRuleSaveResponse) SetStatusCode(v int32) *InvoiceRuleSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceRuleSaveResponse) SetBody(v *InvoiceRuleSaveResponseBody) *InvoiceRuleSaveResponse {
	s.Body = v
	return s
}

type InvoiceSearchHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceSearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvoiceSearchHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceSearchHeaders) SetCommonHeaders(v map[string]*string) *InvoiceSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceSearchHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceSearchHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type InvoiceSearchRequest struct {
	Title  *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s InvoiceSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s InvoiceSearchRequest) GoString() string {
	return s.String()
}

func (s *InvoiceSearchRequest) SetTitle(v string) *InvoiceSearchRequest {
	s.Title = &v
	return s
}

func (s *InvoiceSearchRequest) SetUserId(v string) *InvoiceSearchRequest {
	s.UserId = &v
	return s
}

type InvoiceSearchResponseBody struct {
	Code      *int32                             `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*InvoiceSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                              `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                            `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InvoiceSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvoiceSearchResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceSearchResponseBody) SetCode(v int32) *InvoiceSearchResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceSearchResponseBody) SetMessage(v string) *InvoiceSearchResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceSearchResponseBody) SetModule(v []*InvoiceSearchResponseBodyModule) *InvoiceSearchResponseBody {
	s.Module = v
	return s
}

func (s *InvoiceSearchResponseBody) SetRequestId(v string) *InvoiceSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceSearchResponseBody) SetSuccess(v bool) *InvoiceSearchResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceSearchResponseBody) SetTraceId(v string) *InvoiceSearchResponseBody {
	s.TraceId = &v
	return s
}

type InvoiceSearchResponseBodyModule struct {
	Id                 *int64  `json:"id,omitempty" xml:"id,omitempty"`
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s InvoiceSearchResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s InvoiceSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InvoiceSearchResponseBodyModule) SetId(v int64) *InvoiceSearchResponseBodyModule {
	s.Id = &v
	return s
}

func (s *InvoiceSearchResponseBodyModule) SetThirdPartInvoiceId(v string) *InvoiceSearchResponseBodyModule {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *InvoiceSearchResponseBodyModule) SetTitle(v string) *InvoiceSearchResponseBodyModule {
	s.Title = &v
	return s
}

type InvoiceSearchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvoiceSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvoiceSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s InvoiceSearchResponse) GoString() string {
	return s.String()
}

func (s *InvoiceSearchResponse) SetHeaders(v map[string]*string) *InvoiceSearchResponse {
	s.Headers = v
	return s
}

func (s *InvoiceSearchResponse) SetStatusCode(v int32) *InvoiceSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceSearchResponse) SetBody(v *InvoiceSearchResponseBody) *InvoiceSearchResponse {
	s.Body = v
	return s
}

type IsvUserSaveHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s IsvUserSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s IsvUserSaveHeaders) GoString() string {
	return s.String()
}

func (s *IsvUserSaveHeaders) SetCommonHeaders(v map[string]*string) *IsvUserSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IsvUserSaveHeaders) SetXAcsBtripSoCorpToken(v string) *IsvUserSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type IsvUserSaveRequest struct {
	UserList []*IsvUserSaveRequestUserList `json:"user_list,omitempty" xml:"user_list,omitempty" type:"Repeated"`
}

func (s IsvUserSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s IsvUserSaveRequest) GoString() string {
	return s.String()
}

func (s *IsvUserSaveRequest) SetUserList(v []*IsvUserSaveRequestUserList) *IsvUserSaveRequest {
	s.UserList = v
	return s
}

type IsvUserSaveRequestUserList struct {
	DepartId          *int64    `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	Email             *string   `json:"email,omitempty" xml:"email,omitempty"`
	JobNo             *string   `json:"job_no,omitempty" xml:"job_no,omitempty"`
	LeaveStatus       *int32    `json:"leave_status,omitempty" xml:"leave_status,omitempty"`
	ManagerUserId     *string   `json:"manager_user_id,omitempty" xml:"manager_user_id,omitempty"`
	Phone             *string   `json:"phone,omitempty" xml:"phone,omitempty"`
	Position          *string   `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel     *string   `json:"position_level,omitempty" xml:"position_level,omitempty"`
	RealNameEn        *string   `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	ThirdDepartId     *string   `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
	ThirdDepartIdList []*string `json:"third_depart_id_list,omitempty" xml:"third_depart_id_list,omitempty" type:"Repeated"`
	UserId            *string   `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName          *string   `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s IsvUserSaveRequestUserList) String() string {
	return tea.Prettify(s)
}

func (s IsvUserSaveRequestUserList) GoString() string {
	return s.String()
}

func (s *IsvUserSaveRequestUserList) SetDepartId(v int64) *IsvUserSaveRequestUserList {
	s.DepartId = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetEmail(v string) *IsvUserSaveRequestUserList {
	s.Email = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetJobNo(v string) *IsvUserSaveRequestUserList {
	s.JobNo = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetLeaveStatus(v int32) *IsvUserSaveRequestUserList {
	s.LeaveStatus = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetManagerUserId(v string) *IsvUserSaveRequestUserList {
	s.ManagerUserId = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetPhone(v string) *IsvUserSaveRequestUserList {
	s.Phone = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetPosition(v string) *IsvUserSaveRequestUserList {
	s.Position = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetPositionLevel(v string) *IsvUserSaveRequestUserList {
	s.PositionLevel = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetRealNameEn(v string) *IsvUserSaveRequestUserList {
	s.RealNameEn = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetThirdDepartId(v string) *IsvUserSaveRequestUserList {
	s.ThirdDepartId = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetThirdDepartIdList(v []*string) *IsvUserSaveRequestUserList {
	s.ThirdDepartIdList = v
	return s
}

func (s *IsvUserSaveRequestUserList) SetUserId(v string) *IsvUserSaveRequestUserList {
	s.UserId = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetUserName(v string) *IsvUserSaveRequestUserList {
	s.UserName = &v
	return s
}

type IsvUserSaveShrinkRequest struct {
	UserListShrink *string `json:"user_list,omitempty" xml:"user_list,omitempty"`
}

func (s IsvUserSaveShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s IsvUserSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *IsvUserSaveShrinkRequest) SetUserListShrink(v string) *IsvUserSaveShrinkRequest {
	s.UserListShrink = &v
	return s
}

type IsvUserSaveResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *string `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IsvUserSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsvUserSaveResponseBody) GoString() string {
	return s.String()
}

func (s *IsvUserSaveResponseBody) SetCode(v int32) *IsvUserSaveResponseBody {
	s.Code = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetMessage(v string) *IsvUserSaveResponseBody {
	s.Message = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetModule(v string) *IsvUserSaveResponseBody {
	s.Module = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetRequestId(v string) *IsvUserSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetSuccess(v bool) *IsvUserSaveResponseBody {
	s.Success = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetTraceId(v string) *IsvUserSaveResponseBody {
	s.TraceId = &v
	return s
}

type IsvUserSaveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IsvUserSaveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IsvUserSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s IsvUserSaveResponse) GoString() string {
	return s.String()
}

func (s *IsvUserSaveResponse) SetHeaders(v map[string]*string) *IsvUserSaveResponse {
	s.Headers = v
	return s
}

func (s *IsvUserSaveResponse) SetStatusCode(v int32) *IsvUserSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *IsvUserSaveResponse) SetBody(v *IsvUserSaveResponseBody) *IsvUserSaveResponse {
	s.Body = v
	return s
}

type MonthBillGetHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s MonthBillGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s MonthBillGetHeaders) GoString() string {
	return s.String()
}

func (s *MonthBillGetHeaders) SetCommonHeaders(v map[string]*string) *MonthBillGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MonthBillGetHeaders) SetXAcsBtripSoCorpToken(v string) *MonthBillGetHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type MonthBillGetRequest struct {
	BillMonth *string `json:"bill_month,omitempty" xml:"bill_month,omitempty"`
}

func (s MonthBillGetRequest) String() string {
	return tea.Prettify(s)
}

func (s MonthBillGetRequest) GoString() string {
	return s.String()
}

func (s *MonthBillGetRequest) SetBillMonth(v string) *MonthBillGetRequest {
	s.BillMonth = &v
	return s
}

type MonthBillGetResponseBody struct {
	Code      *int32                            `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*MonthBillGetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                             `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                           `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MonthBillGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MonthBillGetResponseBody) GoString() string {
	return s.String()
}

func (s *MonthBillGetResponseBody) SetCode(v int32) *MonthBillGetResponseBody {
	s.Code = &v
	return s
}

func (s *MonthBillGetResponseBody) SetMessage(v string) *MonthBillGetResponseBody {
	s.Message = &v
	return s
}

func (s *MonthBillGetResponseBody) SetModule(v []*MonthBillGetResponseBodyModule) *MonthBillGetResponseBody {
	s.Module = v
	return s
}

func (s *MonthBillGetResponseBody) SetRequestId(v string) *MonthBillGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *MonthBillGetResponseBody) SetSuccess(v bool) *MonthBillGetResponseBody {
	s.Success = &v
	return s
}

func (s *MonthBillGetResponseBody) SetTraceId(v string) *MonthBillGetResponseBody {
	s.TraceId = &v
	return s
}

type MonthBillGetResponseBodyModule struct {
	EndDate   *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	StartDate *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MonthBillGetResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s MonthBillGetResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MonthBillGetResponseBodyModule) SetEndDate(v string) *MonthBillGetResponseBodyModule {
	s.EndDate = &v
	return s
}

func (s *MonthBillGetResponseBodyModule) SetStartDate(v string) *MonthBillGetResponseBodyModule {
	s.StartDate = &v
	return s
}

func (s *MonthBillGetResponseBodyModule) SetUrl(v string) *MonthBillGetResponseBodyModule {
	s.Url = &v
	return s
}

type MonthBillGetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MonthBillGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MonthBillGetResponse) String() string {
	return tea.Prettify(s)
}

func (s MonthBillGetResponse) GoString() string {
	return s.String()
}

func (s *MonthBillGetResponse) SetHeaders(v map[string]*string) *MonthBillGetResponse {
	s.Headers = v
	return s
}

func (s *MonthBillGetResponse) SetStatusCode(v int32) *MonthBillGetResponse {
	s.StatusCode = &v
	return s
}

func (s *MonthBillGetResponse) SetBody(v *MonthBillGetResponseBody) *MonthBillGetResponse {
	s.Body = v
	return s
}

type ProjectAddHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ProjectAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProjectAddHeaders) GoString() string {
	return s.String()
}

func (s *ProjectAddHeaders) SetCommonHeaders(v map[string]*string) *ProjectAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProjectAddHeaders) SetXAcsBtripSoCorpToken(v string) *ProjectAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ProjectAddRequest struct {
	Code                  *string `json:"code,omitempty" xml:"code,omitempty"`
	ProjectName           *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	ThirdPartId           *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
}

func (s ProjectAddRequest) String() string {
	return tea.Prettify(s)
}

func (s ProjectAddRequest) GoString() string {
	return s.String()
}

func (s *ProjectAddRequest) SetCode(v string) *ProjectAddRequest {
	s.Code = &v
	return s
}

func (s *ProjectAddRequest) SetProjectName(v string) *ProjectAddRequest {
	s.ProjectName = &v
	return s
}

func (s *ProjectAddRequest) SetThirdPartCostCenterId(v string) *ProjectAddRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *ProjectAddRequest) SetThirdPartId(v string) *ProjectAddRequest {
	s.ThirdPartId = &v
	return s
}

func (s *ProjectAddRequest) SetThirdPartInvoiceId(v string) *ProjectAddRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

type ProjectAddResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *int64  `json:"module,omitempty" xml:"module,omitempty"`
	MorePage  *bool   `json:"more_page,omitempty" xml:"more_page,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ProjectAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProjectAddResponseBody) GoString() string {
	return s.String()
}

func (s *ProjectAddResponseBody) SetCode(v int32) *ProjectAddResponseBody {
	s.Code = &v
	return s
}

func (s *ProjectAddResponseBody) SetMessage(v string) *ProjectAddResponseBody {
	s.Message = &v
	return s
}

func (s *ProjectAddResponseBody) SetModule(v int64) *ProjectAddResponseBody {
	s.Module = &v
	return s
}

func (s *ProjectAddResponseBody) SetMorePage(v bool) *ProjectAddResponseBody {
	s.MorePage = &v
	return s
}

func (s *ProjectAddResponseBody) SetRequestId(v string) *ProjectAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProjectAddResponseBody) SetSuccess(v bool) *ProjectAddResponseBody {
	s.Success = &v
	return s
}

func (s *ProjectAddResponseBody) SetTraceId(v string) *ProjectAddResponseBody {
	s.TraceId = &v
	return s
}

type ProjectAddResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ProjectAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProjectAddResponse) String() string {
	return tea.Prettify(s)
}

func (s ProjectAddResponse) GoString() string {
	return s.String()
}

func (s *ProjectAddResponse) SetHeaders(v map[string]*string) *ProjectAddResponse {
	s.Headers = v
	return s
}

func (s *ProjectAddResponse) SetStatusCode(v int32) *ProjectAddResponse {
	s.StatusCode = &v
	return s
}

func (s *ProjectAddResponse) SetBody(v *ProjectAddResponseBody) *ProjectAddResponse {
	s.Body = v
	return s
}

type ProjectDeleteHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ProjectDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProjectDeleteHeaders) GoString() string {
	return s.String()
}

func (s *ProjectDeleteHeaders) SetCommonHeaders(v map[string]*string) *ProjectDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProjectDeleteHeaders) SetXAcsBtripSoCorpToken(v string) *ProjectDeleteHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ProjectDeleteRequest struct {
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s ProjectDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s ProjectDeleteRequest) GoString() string {
	return s.String()
}

func (s *ProjectDeleteRequest) SetThirdPartId(v string) *ProjectDeleteRequest {
	s.ThirdPartId = &v
	return s
}

type ProjectDeleteResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *bool   `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ProjectDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProjectDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *ProjectDeleteResponseBody) SetCode(v int32) *ProjectDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetMessage(v string) *ProjectDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetModule(v bool) *ProjectDeleteResponseBody {
	s.Module = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetRequestId(v string) *ProjectDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetSuccess(v bool) *ProjectDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetTraceId(v string) *ProjectDeleteResponseBody {
	s.TraceId = &v
	return s
}

type ProjectDeleteResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ProjectDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProjectDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s ProjectDeleteResponse) GoString() string {
	return s.String()
}

func (s *ProjectDeleteResponse) SetHeaders(v map[string]*string) *ProjectDeleteResponse {
	s.Headers = v
	return s
}

func (s *ProjectDeleteResponse) SetStatusCode(v int32) *ProjectDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *ProjectDeleteResponse) SetBody(v *ProjectDeleteResponseBody) *ProjectDeleteResponse {
	s.Body = v
	return s
}

type ProjectModifyHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ProjectModifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProjectModifyHeaders) GoString() string {
	return s.String()
}

func (s *ProjectModifyHeaders) SetCommonHeaders(v map[string]*string) *ProjectModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProjectModifyHeaders) SetXAcsBtripSoCorpToken(v string) *ProjectModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type ProjectModifyRequest struct {
	Code                  *string `json:"code,omitempty" xml:"code,omitempty"`
	ProjectName           *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	ThirdPartId           *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
}

func (s ProjectModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ProjectModifyRequest) GoString() string {
	return s.String()
}

func (s *ProjectModifyRequest) SetCode(v string) *ProjectModifyRequest {
	s.Code = &v
	return s
}

func (s *ProjectModifyRequest) SetProjectName(v string) *ProjectModifyRequest {
	s.ProjectName = &v
	return s
}

func (s *ProjectModifyRequest) SetThirdPartCostCenterId(v string) *ProjectModifyRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *ProjectModifyRequest) SetThirdPartId(v string) *ProjectModifyRequest {
	s.ThirdPartId = &v
	return s
}

func (s *ProjectModifyRequest) SetThirdPartInvoiceId(v string) *ProjectModifyRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

type ProjectModifyResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *bool   `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ProjectModifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProjectModifyResponseBody) GoString() string {
	return s.String()
}

func (s *ProjectModifyResponseBody) SetCode(v int32) *ProjectModifyResponseBody {
	s.Code = &v
	return s
}

func (s *ProjectModifyResponseBody) SetMessage(v string) *ProjectModifyResponseBody {
	s.Message = &v
	return s
}

func (s *ProjectModifyResponseBody) SetModule(v bool) *ProjectModifyResponseBody {
	s.Module = &v
	return s
}

func (s *ProjectModifyResponseBody) SetRequestId(v string) *ProjectModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProjectModifyResponseBody) SetSuccess(v bool) *ProjectModifyResponseBody {
	s.Success = &v
	return s
}

func (s *ProjectModifyResponseBody) SetTraceId(v string) *ProjectModifyResponseBody {
	s.TraceId = &v
	return s
}

type ProjectModifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ProjectModifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProjectModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ProjectModifyResponse) GoString() string {
	return s.String()
}

func (s *ProjectModifyResponse) SetHeaders(v map[string]*string) *ProjectModifyResponse {
	s.Headers = v
	return s
}

func (s *ProjectModifyResponse) SetStatusCode(v int32) *ProjectModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ProjectModifyResponse) SetBody(v *ProjectModifyResponseBody) *ProjectModifyResponse {
	s.Body = v
	return s
}

type TrainBillSettlementQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainBillSettlementQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s TrainBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type TrainBillSettlementQueryRequest struct {
	PageNo      *int32  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize    *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PeriodEnd   *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
}

func (s TrainBillSettlementQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TrainBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryRequest) SetPageNo(v int32) *TrainBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetPageSize(v int32) *TrainBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetPeriodEnd(v string) *TrainBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetPeriodStart(v string) *TrainBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

type TrainBillSettlementQueryResponseBody struct {
	Code      *int32                                      `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	Module    *TrainBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                     `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainBillSettlementQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrainBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryResponseBody) SetCode(v int32) *TrainBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetMessage(v string) *TrainBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetModule(v *TrainBillSettlementQueryResponseBodyModule) *TrainBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetRequestId(v string) *TrainBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetSuccess(v bool) *TrainBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetTraceId(v string) *TrainBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

type TrainBillSettlementQueryResponseBodyModule struct {
	Category    *int32                                                `json:"category,omitempty" xml:"category,omitempty"`
	CorpId      *string                                               `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList    []*TrainBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
	PeriodEnd   *string                                               `json:"period_end,omitempty" xml:"period_end,omitempty"`
	PeriodStart *string                                               `json:"period_start,omitempty" xml:"period_start,omitempty"`
	TotalNum    *int64                                                `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s TrainBillSettlementQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s TrainBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetCategory(v int32) *TrainBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetCorpId(v string) *TrainBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetDataList(v []*TrainBillSettlementQueryResponseBodyModuleDataList) *TrainBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *TrainBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *TrainBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *TrainBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

type TrainBillSettlementQueryResponseBodyModuleDataList struct {
	AlipayTradeNo      *string  `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyId            *string  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrDate            *string  `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	ArrStation         *string  `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	ArrTime            *string  `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BillRecordTime     *string  `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookTime           *string  `json:"book_time,omitempty" xml:"book_time,omitempty"`
	BookerId           *string  `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	BookerJobNo        *string  `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName         *string  `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	CapitalDirection   *string  `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment  *string  `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	ChangeFee          *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	CostCenter         *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	CostCenterNumber   *string  `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	Coupon             *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	Department         *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId       *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptDate           *string  `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptStation        *string  `json:"dept_station,omitempty" xml:"dept_station,omitempty"`
	DeptTime           *string  `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	FeeType            *string  `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	Index              *string  `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle       *string  `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	OrderId            *string  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderPrice         *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OverApplyId        *string  `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PrimaryId          *int64   `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProjectCode        *string  `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName        *string  `json:"project_name,omitempty" xml:"project_name,omitempty"`
	RefundFee          *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	Remark             *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	RunTime            *string  `json:"run_time,omitempty" xml:"run_time,omitempty"`
	SeatNo             *string  `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	SeatType           *string  `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	ServiceFee         *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettlementFee      *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	SettlementGrantFee *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	SettlementTime     *string  `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	SettlementType     *string  `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	Status             *int32   `json:"status,omitempty" xml:"status,omitempty"`
	TicketNo           *string  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketPrice        *float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	TrainNo            *string  `json:"train_no,omitempty" xml:"train_no,omitempty"`
	TrainType          *string  `json:"train_type,omitempty" xml:"train_type,omitempty"`
	TravelerId         *string  `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	TravelerJobNo      *string  `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerName       *string  `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	VoucherType        *int32   `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}

func (s TrainBillSettlementQueryResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s TrainBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrDate(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrStation(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrStation = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetChangeFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCoupon(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDeptDate(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDeptStation(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DeptStation = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDeptTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetOrderPrice(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRefundFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.RefundFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRunTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.RunTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSeatNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SeatNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSeatType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SeatType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTicketNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TicketNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTicketPrice(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TicketPrice = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTrainNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TrainNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTrainType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TrainType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

type TrainBillSettlementQueryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TrainBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TrainBillSettlementQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TrainBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryResponse) SetHeaders(v map[string]*string) *TrainBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainBillSettlementQueryResponse) SetStatusCode(v int32) *TrainBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponse) SetBody(v *TrainBillSettlementQueryResponseBody) *TrainBillSettlementQueryResponse {
	s.Body = v
	return s
}

type TrainExceedApplyQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainExceedApplyQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s TrainExceedApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainExceedApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainExceedApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainExceedApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type TrainExceedApplyQueryRequest struct {
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
}

func (s TrainExceedApplyQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TrainExceedApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryRequest) SetApplyId(v int64) *TrainExceedApplyQueryRequest {
	s.ApplyId = &v
	return s
}

type TrainExceedApplyQueryResponseBody struct {
	Code      *int32                                   `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module    *TrainExceedApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainExceedApplyQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrainExceedApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryResponseBody) SetCode(v int32) *TrainExceedApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetMessage(v string) *TrainExceedApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetModule(v *TrainExceedApplyQueryResponseBodyModule) *TrainExceedApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetRequestId(v string) *TrainExceedApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetSuccess(v bool) *TrainExceedApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetTraceId(v string) *TrainExceedApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

type TrainExceedApplyQueryResponseBodyModule struct {
	ApplyId              *int64                                                       `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApplyIntentionInfoDO *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO `json:"apply_intention_info_d_o,omitempty" xml:"apply_intention_info_d_o,omitempty" type:"Struct"`
	BtripCause           *string                                                      `json:"btrip_cause,omitempty" xml:"btrip_cause,omitempty"`
	CorpId               *string                                                      `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	ExceedReason         *string                                                      `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	ExceedType           *int32                                                       `json:"exceed_type,omitempty" xml:"exceed_type,omitempty"`
	OriginStandard       *string                                                      `json:"origin_standard,omitempty" xml:"origin_standard,omitempty"`
	Status               *int32                                                       `json:"status,omitempty" xml:"status,omitempty"`
	SubmitTime           *string                                                      `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
	ThirdpartApplyId     *string                                                      `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartCorpId      *string                                                      `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	UserId               *string                                                      `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainExceedApplyQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s TrainExceedApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetApplyId(v int64) *TrainExceedApplyQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetApplyIntentionInfoDO(v *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) *TrainExceedApplyQueryResponseBodyModule {
	s.ApplyIntentionInfoDO = v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetBtripCause(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.BtripCause = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetCorpId(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetExceedReason(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.ExceedReason = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetExceedType(v int32) *TrainExceedApplyQueryResponseBodyModule {
	s.ExceedType = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetOriginStandard(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.OriginStandard = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetStatus(v int32) *TrainExceedApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetSubmitTime(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.SubmitTime = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetThirdpartApplyId(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetThirdpartCorpId(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.ThirdpartCorpId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetUserId(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

type TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO struct {
	ArrCity       *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityName   *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrStation    *string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	ArrTime       *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepCity       *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityName   *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepStation    *string `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	DepTime       *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Price         *int64  `json:"price,omitempty" xml:"price,omitempty"`
	SeatName      *string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	TrainNo       *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	TrainTypeDesc *string `json:"train_type_desc,omitempty" xml:"train_type_desc,omitempty"`
	Type          *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) String() string {
	return tea.Prettify(s)
}

func (s TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetArrCity(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.ArrCity = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetArrCityName(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.ArrCityName = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetArrStation(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.ArrStation = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetArrTime(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.ArrTime = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetDepCity(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.DepCity = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetDepCityName(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.DepCityName = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetDepStation(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.DepStation = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetDepTime(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.DepTime = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetPrice(v int64) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.Price = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetSeatName(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.SeatName = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetTrainNo(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.TrainNo = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetTrainTypeDesc(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.TrainTypeDesc = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetType(v int32) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.Type = &v
	return s
}

type TrainExceedApplyQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TrainExceedApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TrainExceedApplyQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TrainExceedApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryResponse) SetHeaders(v map[string]*string) *TrainExceedApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainExceedApplyQueryResponse) SetStatusCode(v int32) *TrainExceedApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainExceedApplyQueryResponse) SetBody(v *TrainExceedApplyQueryResponseBody) *TrainExceedApplyQueryResponse {
	s.Body = v
	return s
}

type TrainOrderListQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainOrderListQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainOrderListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type TrainOrderListQueryRequest struct {
	AllApply         *bool   `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	ApplyId          *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	DepartId         *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Page             *int32  `json:"page,omitempty" xml:"page,omitempty"`
	PageSize         *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	UpdateEndTime    *string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	UpdateStartTime  *string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderListQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryRequest) SetAllApply(v bool) *TrainOrderListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetApplyId(v int64) *TrainOrderListQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetDepartId(v string) *TrainOrderListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetEndTime(v string) *TrainOrderListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetPage(v int32) *TrainOrderListQueryRequest {
	s.Page = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetPageSize(v int32) *TrainOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetStartTime(v string) *TrainOrderListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetThirdpartApplyId(v string) *TrainOrderListQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetUpdateEndTime(v string) *TrainOrderListQueryRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetUpdateStartTime(v string) *TrainOrderListQueryRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetUserId(v string) *TrainOrderListQueryRequest {
	s.UserId = &v
	return s
}

type TrainOrderListQueryResponseBody struct {
	Code      *int32                                   `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*TrainOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo  *TrainOrderListQueryResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainOrderListQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBody) SetCode(v int32) *TrainOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetMessage(v string) *TrainOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetModule(v []*TrainOrderListQueryResponseBodyModule) *TrainOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetPageInfo(v *TrainOrderListQueryResponseBodyPageInfo) *TrainOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetRequestId(v string) *TrainOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetSuccess(v bool) *TrainOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetTraceId(v string) *TrainOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

type TrainOrderListQueryResponseBodyModule struct {
	ApplyId              *int64                                                    `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrCity              *string                                                   `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrStation           *string                                                   `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	ArrTime              *string                                                   `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BtripTitle           *string                                                   `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	ContactName          *string                                                   `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId               *string                                                   `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName             *string                                                   `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	CostCenter           *TrainOrderListQueryResponseBodyModuleCostCenter          `json:"cost_center,omitempty" xml:"cost_center,omitempty" type:"Struct"`
	DepCity              *string                                                   `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepStation           *string                                                   `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	DepTime              *string                                                   `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	DepartId             *string                                                   `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName           *string                                                   `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	GmtCreate            *string                                                   `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModified          *string                                                   `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	Id                   *int64                                                    `json:"id,omitempty" xml:"id,omitempty"`
	Invoice              *TrainOrderListQueryResponseBodyModuleInvoice             `json:"invoice,omitempty" xml:"invoice,omitempty" type:"Struct"`
	PriceInfoList        []*TrainOrderListQueryResponseBodyModulePriceInfoList     `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	ProjectCode          *string                                                   `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectId            *int64                                                    `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle         *string                                                   `json:"project_title,omitempty" xml:"project_title,omitempty"`
	RiderName            *string                                                   `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	RunTime              *string                                                   `json:"run_time,omitempty" xml:"run_time,omitempty"`
	SeatType             *string                                                   `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	Status               *int32                                                    `json:"status,omitempty" xml:"status,omitempty"`
	ThirdPartProjectId   *string                                                   `json:"third_part_project_id,omitempty" xml:"third_part_project_id,omitempty"`
	ThirdpartApplyId     *string                                                   `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartItineraryId *string                                                   `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	TicketCount          *int32                                                    `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	TicketNo12306        *string                                                   `json:"ticket_no12306,omitempty" xml:"ticket_no12306,omitempty"`
	TrainNumber          *string                                                   `json:"train_number,omitempty" xml:"train_number,omitempty"`
	TrainType            *string                                                   `json:"train_type,omitempty" xml:"train_type,omitempty"`
	UserAffiliateList    []*TrainOrderListQueryResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	UserId               *string                                                   `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName             *string                                                   `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModule) SetApplyId(v int64) *TrainOrderListQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetArrCity(v string) *TrainOrderListQueryResponseBodyModule {
	s.ArrCity = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetArrStation(v string) *TrainOrderListQueryResponseBodyModule {
	s.ArrStation = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetArrTime(v string) *TrainOrderListQueryResponseBodyModule {
	s.ArrTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetBtripTitle(v string) *TrainOrderListQueryResponseBodyModule {
	s.BtripTitle = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetContactName(v string) *TrainOrderListQueryResponseBodyModule {
	s.ContactName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetCorpId(v string) *TrainOrderListQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetCorpName(v string) *TrainOrderListQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetCostCenter(v *TrainOrderListQueryResponseBodyModuleCostCenter) *TrainOrderListQueryResponseBodyModule {
	s.CostCenter = v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepCity(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepCity = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepStation(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepStation = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepTime(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepartId(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepartName(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetGmtCreate(v string) *TrainOrderListQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetGmtModified(v string) *TrainOrderListQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetId(v int64) *TrainOrderListQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetInvoice(v *TrainOrderListQueryResponseBodyModuleInvoice) *TrainOrderListQueryResponseBodyModule {
	s.Invoice = v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetPriceInfoList(v []*TrainOrderListQueryResponseBodyModulePriceInfoList) *TrainOrderListQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetProjectCode(v string) *TrainOrderListQueryResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetProjectId(v int64) *TrainOrderListQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetProjectTitle(v string) *TrainOrderListQueryResponseBodyModule {
	s.ProjectTitle = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetRiderName(v string) *TrainOrderListQueryResponseBodyModule {
	s.RiderName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetRunTime(v string) *TrainOrderListQueryResponseBodyModule {
	s.RunTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetSeatType(v string) *TrainOrderListQueryResponseBodyModule {
	s.SeatType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetStatus(v int32) *TrainOrderListQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetThirdPartProjectId(v string) *TrainOrderListQueryResponseBodyModule {
	s.ThirdPartProjectId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetThirdpartApplyId(v string) *TrainOrderListQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetThirdpartItineraryId(v string) *TrainOrderListQueryResponseBodyModule {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetTicketCount(v int32) *TrainOrderListQueryResponseBodyModule {
	s.TicketCount = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetTicketNo12306(v string) *TrainOrderListQueryResponseBodyModule {
	s.TicketNo12306 = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetTrainNumber(v string) *TrainOrderListQueryResponseBodyModule {
	s.TrainNumber = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetTrainType(v string) *TrainOrderListQueryResponseBodyModule {
	s.TrainType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetUserAffiliateList(v []*TrainOrderListQueryResponseBodyModuleUserAffiliateList) *TrainOrderListQueryResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetUserId(v string) *TrainOrderListQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetUserName(v string) *TrainOrderListQueryResponseBodyModule {
	s.UserName = &v
	return s
}

type TrainOrderListQueryResponseBodyModuleCostCenter struct {
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Id     *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModuleCostCenter) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModuleCostCenter) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) SetCorpId(v string) *TrainOrderListQueryResponseBodyModuleCostCenter {
	s.CorpId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) SetId(v int64) *TrainOrderListQueryResponseBodyModuleCostCenter {
	s.Id = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) SetName(v string) *TrainOrderListQueryResponseBodyModuleCostCenter {
	s.Name = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) SetNumber(v string) *TrainOrderListQueryResponseBodyModuleCostCenter {
	s.Number = &v
	return s
}

type TrainOrderListQueryResponseBodyModuleInvoice struct {
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModuleInvoice) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModuleInvoice) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModuleInvoice) SetId(v int64) *TrainOrderListQueryResponseBodyModuleInvoice {
	s.Id = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleInvoice) SetTitle(v string) *TrainOrderListQueryResponseBodyModuleInvoice {
	s.Title = &v
	return s
}

type TrainOrderListQueryResponseBodyModulePriceInfoList struct {
	CategoryCode    *int32   `json:"category_code,omitempty" xml:"category_code,omitempty"`
	CategoryType    *int32   `json:"category_type,omitempty" xml:"category_type,omitempty"`
	EndCity         *string  `json:"end_city,omitempty" xml:"end_city,omitempty"`
	EndTime         *string  `json:"end_time,omitempty" xml:"end_time,omitempty"`
	GmtCreate       *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	OriginalTrainNo *string  `json:"original_train_no,omitempty" xml:"original_train_no,omitempty"`
	PassengerName   *string  `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PayType         *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	Price           *float64 `json:"price,omitempty" xml:"price,omitempty"`
	SeatType        *string  `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	StartCity       *string  `json:"start_city,omitempty" xml:"start_city,omitempty"`
	StartTime       *string  `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TradeId         *string  `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	TrainNo         *string  `json:"train_no,omitempty" xml:"train_no,omitempty"`
	Type            *int32   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModulePriceInfoList) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetCategoryType(v int32) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetEndCity(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.EndCity = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetEndTime(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.EndTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetOriginalTrainNo(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.OriginalTrainNo = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetSeatType(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.SeatType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetStartCity(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.StartCity = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetStartTime(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.StartTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetTrainNo(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.TrainNo = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetType(v int32) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

type TrainOrderListQueryResponseBodyModuleUserAffiliateList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModuleUserAffiliateList) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModuleUserAffiliateList) SetUserId(v string) *TrainOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleUserAffiliateList) SetUserName(v string) *TrainOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

type TrainOrderListQueryResponseBodyPageInfo struct {
	Page        *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageSize    *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s TrainOrderListQueryResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyPageInfo) SetPage(v int32) *TrainOrderListQueryResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyPageInfo) SetPageSize(v int32) *TrainOrderListQueryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *TrainOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

type TrainOrderListQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TrainOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TrainOrderListQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponse) SetHeaders(v map[string]*string) *TrainOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderListQueryResponse) SetStatusCode(v int32) *TrainOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderListQueryResponse) SetBody(v *TrainOrderListQueryResponseBody) *TrainOrderListQueryResponse {
	s.Body = v
	return s
}

type TrainOrderQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainOrderQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainOrderQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type TrainOrderQueryRequest struct {
	OrderId *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	UserId  *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryRequest) SetOrderId(v int64) *TrainOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *TrainOrderQueryRequest) SetUserId(v string) *TrainOrderQueryRequest {
	s.UserId = &v
	return s
}

type TrainOrderQueryResponseBody struct {
	Module     *TrainOrderQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId  *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultCode *int32                             `json:"result_code,omitempty" xml:"result_code,omitempty"`
	ResultMsg  *string                            `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	Success    *bool                              `json:"success,omitempty" xml:"success,omitempty"`
	TraceId    *string                            `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainOrderQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBody) SetModule(v *TrainOrderQueryResponseBodyModule) *TrainOrderQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderQueryResponseBody) SetRequestId(v string) *TrainOrderQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderQueryResponseBody) SetResultCode(v int32) *TrainOrderQueryResponseBody {
	s.ResultCode = &v
	return s
}

func (s *TrainOrderQueryResponseBody) SetResultMsg(v string) *TrainOrderQueryResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *TrainOrderQueryResponseBody) SetSuccess(v bool) *TrainOrderQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderQueryResponseBody) SetTraceId(v string) *TrainOrderQueryResponseBody {
	s.TraceId = &v
	return s
}

type TrainOrderQueryResponseBodyModule struct {
	ChangeTicketInfoList []*TrainOrderQueryResponseBodyModuleChangeTicketInfoList `json:"change_ticket_info_list,omitempty" xml:"change_ticket_info_list,omitempty" type:"Repeated"`
	InvoiceInfo          *TrainOrderQueryResponseBodyModuleInvoiceInfo            `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	OrderBaseInfo        *TrainOrderQueryResponseBodyModuleOrderBaseInfo          `json:"order_base_info,omitempty" xml:"order_base_info,omitempty" type:"Struct"`
	PassengerInfoList    []*TrainOrderQueryResponseBodyModulePassengerInfoList    `json:"passenger_info_list,omitempty" xml:"passenger_info_list,omitempty" type:"Repeated"`
	PriceInfoList        []*TrainOrderQueryResponseBodyModulePriceInfoList        `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	RefundTicketInfoList []*TrainOrderQueryResponseBodyModuleRefundTicketInfoList `json:"refund_ticket_info_list,omitempty" xml:"refund_ticket_info_list,omitempty" type:"Repeated"`
	TicketInfoList       []*TrainOrderQueryResponseBodyModuleTicketInfoList       `json:"ticket_info_list,omitempty" xml:"ticket_info_list,omitempty" type:"Repeated"`
	TrainInfo            *TrainOrderQueryResponseBodyModuleTrainInfo              `json:"train_info,omitempty" xml:"train_info,omitempty" type:"Struct"`
}

func (s TrainOrderQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModule) SetChangeTicketInfoList(v []*TrainOrderQueryResponseBodyModuleChangeTicketInfoList) *TrainOrderQueryResponseBodyModule {
	s.ChangeTicketInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetInvoiceInfo(v *TrainOrderQueryResponseBodyModuleInvoiceInfo) *TrainOrderQueryResponseBodyModule {
	s.InvoiceInfo = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetOrderBaseInfo(v *TrainOrderQueryResponseBodyModuleOrderBaseInfo) *TrainOrderQueryResponseBodyModule {
	s.OrderBaseInfo = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetPassengerInfoList(v []*TrainOrderQueryResponseBodyModulePassengerInfoList) *TrainOrderQueryResponseBodyModule {
	s.PassengerInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetPriceInfoList(v []*TrainOrderQueryResponseBodyModulePriceInfoList) *TrainOrderQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetRefundTicketInfoList(v []*TrainOrderQueryResponseBodyModuleRefundTicketInfoList) *TrainOrderQueryResponseBodyModule {
	s.RefundTicketInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetTicketInfoList(v []*TrainOrderQueryResponseBodyModuleTicketInfoList) *TrainOrderQueryResponseBodyModule {
	s.TicketInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetTrainInfo(v *TrainOrderQueryResponseBodyModuleTrainInfo) *TrainOrderQueryResponseBodyModule {
	s.TrainInfo = v
	return s
}

type TrainOrderQueryResponseBodyModuleChangeTicketInfoList struct {
	ChangeCoachNo       *string  `json:"change_coach_no,omitempty" xml:"change_coach_no,omitempty"`
	ChangeGapFee        *float64 `json:"change_gap_fee,omitempty" xml:"change_gap_fee,omitempty"`
	ChangeHandlingFee   *float64 `json:"change_handling_fee,omitempty" xml:"change_handling_fee,omitempty"`
	ChangeSeatNo        *string  `json:"change_seat_no,omitempty" xml:"change_seat_no,omitempty"`
	ChangeSeatTypeName  *string  `json:"change_seat_type_name,omitempty" xml:"change_seat_type_name,omitempty"`
	ChangeServiceFee    *float64 `json:"change_service_fee,omitempty" xml:"change_service_fee,omitempty"`
	ChangeTrainNo       *string  `json:"change_train_no,omitempty" xml:"change_train_no,omitempty"`
	ChangeTrainTypeName *string  `json:"change_train_type_name,omitempty" xml:"change_train_type_name,omitempty"`
	CheckInTime         *string  `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	CheckOutTime        *string  `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	EndTime             *string  `json:"end_time,omitempty" xml:"end_time,omitempty"`
	FromStationName     *string  `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	GmtCreate           *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify           *string  `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	OriginTicketNo      *string  `json:"origin_ticket_no,omitempty" xml:"origin_ticket_no,omitempty"`
	OutTicketStatus     *string  `json:"out_ticket_status,omitempty" xml:"out_ticket_status,omitempty"`
	StartTime           *string  `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TicketNo            *string  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	ToStationName       *string  `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleChangeTicketInfoList) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeCoachNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeCoachNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeGapFee(v float64) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeGapFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeHandlingFee(v float64) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeHandlingFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeSeatNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeSeatNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeSeatTypeName(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeSeatTypeName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeServiceFee(v float64) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeServiceFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeTrainNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeTrainNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeTrainTypeName(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeTrainTypeName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetCheckInTime(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.CheckInTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetCheckOutTime(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.CheckOutTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetEndTime(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.EndTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetFromStationName(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.FromStationName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetGmtModify(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetOriginTicketNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.OriginTicketNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetOutTicketStatus(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.OutTicketStatus = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetStartTime(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.StartTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetTicketNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetToStationName(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ToStationName = &v
	return s
}

type TrainOrderQueryResponseBodyModuleInvoiceInfo struct {
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleInvoiceInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleInvoiceInfo) SetId(v int64) *TrainOrderQueryResponseBodyModuleInvoiceInfo {
	s.Id = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleInvoiceInfo) SetTitle(v string) *TrainOrderQueryResponseBodyModuleInvoiceInfo {
	s.Title = &v
	return s
}

type TrainOrderQueryResponseBodyModuleOrderBaseInfo struct {
	ApplyId              *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripTitle           *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	ContactName          *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId               *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName             *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId             *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName           *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	GmtCreate            *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify            *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	ItineraryId          *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	OrderId              *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderStatus          *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	ThirdpartApplyId     *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartCorpId      *string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	TripType             *int32  `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	UserId               *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleOrderBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetApplyId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ApplyId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetBtripTitle(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.BtripTitle = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetContactName(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ContactName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpName(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartName(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtModify(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetItineraryId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ItineraryId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderId(v int64) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderStatus(v int32) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderStatus = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartApplyId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartCorpId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartCorpId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartItineraryId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetTripType(v int32) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.TripType = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetUserId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.UserId = &v
	return s
}

type TrainOrderQueryResponseBodyModulePassengerInfoList struct {
	CostCenterId       *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName     *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	CostCenterNumber   *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	ProjectCode        *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectId          *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle       *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdpartProjectId *string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	UserId             *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName           *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	UserType           *int32  `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s TrainOrderQueryResponseBodyModulePassengerInfoList) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModulePassengerInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterId(v int64) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterName(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterNumber(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterNumber = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetProjectCode(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectCode = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetProjectId(v int64) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetProjectTitle(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectTitle = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetThirdpartProjectId(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.ThirdpartProjectId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetUserId(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetUserName(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.UserName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetUserType(v int32) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.UserType = &v
	return s
}

type TrainOrderQueryResponseBodyModulePriceInfoList struct {
	CategoryCode  *int32   `json:"category_code,omitempty" xml:"category_code,omitempty"`
	GmtCreate     *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string  `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PayType       *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	Price         *float64 `json:"price,omitempty" xml:"price,omitempty"`
	TradeId       *string  `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	Type          *int32   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TrainOrderQueryResponseBodyModulePriceInfoList) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetType(v int32) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

type TrainOrderQueryResponseBodyModuleRefundTicketInfoList struct {
	GmtCreate        *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify        *string  `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	RefundFee        *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	RefundServiceFee *float64 `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
	TicketNo         *string  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleRefundTicketInfoList) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleRefundTicketInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetGmtModify(v string) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetRefundFee(v float64) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.RefundFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetRefundServiceFee(v float64) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.RefundServiceFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetTicketNo(v string) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.TicketNo = &v
	return s
}

type TrainOrderQueryResponseBodyModuleTicketInfoList struct {
	Changed         *bool    `json:"changed,omitempty" xml:"changed,omitempty"`
	CheckInTime     *string  `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	CheckOutTime    *string  `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	CoachNo         *string  `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
	EndTime         *string  `json:"end_time,omitempty" xml:"end_time,omitempty"`
	GmtCreate       *string  `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify       *string  `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	OutTicketStatus *string  `json:"out_ticket_status,omitempty" xml:"out_ticket_status,omitempty"`
	PayType         *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	SeatNo          *string  `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	SeatTypeName    *string  `json:"seat_type_name,omitempty" xml:"seat_type_name,omitempty"`
	ServiceFee      *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	StartTime       *string  `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TicketNo        *string  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketPrice     *float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	TicketStatus    *int32   `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	TrainTypeName   *string  `json:"train_type_name,omitempty" xml:"train_type_name,omitempty"`
	UserId          *string  `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleTicketInfoList) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleTicketInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetChanged(v bool) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.Changed = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetCheckInTime(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.CheckInTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetCheckOutTime(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.CheckOutTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetCoachNo(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.CoachNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetEndTime(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.EndTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetGmtModify(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetOutTicketStatus(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.OutTicketStatus = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetPayType(v int32) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.PayType = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetSeatNo(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.SeatNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetSeatTypeName(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.SeatTypeName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetServiceFee(v float64) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.ServiceFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetStartTime(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.StartTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetTicketNo(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetTicketPrice(v float64) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.TicketPrice = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetTicketStatus(v int32) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.TicketStatus = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetTrainTypeName(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.TrainTypeName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetUserId(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.UserId = &v
	return s
}

type TrainOrderQueryResponseBodyModuleTrainInfo struct {
	ArrTime         *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepTime         *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FromStationName *string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	RunTime         *int64  `json:"run_time,omitempty" xml:"run_time,omitempty"`
	ToStationName   *string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	TrainNo         *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleTrainInfo) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleTrainInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetArrTime(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.ArrTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetDepTime(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.DepTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetFromStationName(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.FromStationName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetRunTime(v int64) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.RunTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetToStationName(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.ToStationName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetTrainNo(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.TrainNo = &v
	return s
}

type TrainOrderQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TrainOrderQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TrainOrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TrainOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponse) SetHeaders(v map[string]*string) *TrainOrderQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderQueryResponse) SetStatusCode(v int32) *TrainOrderQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderQueryResponse) SetBody(v *TrainOrderQueryResponseBody) *TrainOrderQueryResponse {
	s.Body = v
	return s
}

type TrainStationSearchHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainStationSearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s TrainStationSearchHeaders) GoString() string {
	return s.String()
}

func (s *TrainStationSearchHeaders) SetCommonHeaders(v map[string]*string) *TrainStationSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainStationSearchHeaders) SetXAcsBtripSoCorpToken(v string) *TrainStationSearchHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type TrainStationSearchRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

func (s TrainStationSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s TrainStationSearchRequest) GoString() string {
	return s.String()
}

func (s *TrainStationSearchRequest) SetKeyword(v string) *TrainStationSearchRequest {
	s.Keyword = &v
	return s
}

type TrainStationSearchResponseBody struct {
	Code      *int32                                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Module    *TrainStationSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                               `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainStationSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrainStationSearchResponseBody) GoString() string {
	return s.String()
}

func (s *TrainStationSearchResponseBody) SetCode(v int32) *TrainStationSearchResponseBody {
	s.Code = &v
	return s
}

func (s *TrainStationSearchResponseBody) SetMessage(v string) *TrainStationSearchResponseBody {
	s.Message = &v
	return s
}

func (s *TrainStationSearchResponseBody) SetModule(v *TrainStationSearchResponseBodyModule) *TrainStationSearchResponseBody {
	s.Module = v
	return s
}

func (s *TrainStationSearchResponseBody) SetRequestId(v string) *TrainStationSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainStationSearchResponseBody) SetSuccess(v bool) *TrainStationSearchResponseBody {
	s.Success = &v
	return s
}

func (s *TrainStationSearchResponseBody) SetTraceId(v string) *TrainStationSearchResponseBody {
	s.TraceId = &v
	return s
}

type TrainStationSearchResponseBodyModule struct {
	Cities []*TrainStationSearchResponseBodyModuleCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
}

func (s TrainStationSearchResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s TrainStationSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainStationSearchResponseBodyModule) SetCities(v []*TrainStationSearchResponseBodyModuleCities) *TrainStationSearchResponseBodyModule {
	s.Cities = v
	return s
}

type TrainStationSearchResponseBodyModuleCities struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TrainStationSearchResponseBodyModuleCities) String() string {
	return tea.Prettify(s)
}

func (s TrainStationSearchResponseBodyModuleCities) GoString() string {
	return s.String()
}

func (s *TrainStationSearchResponseBodyModuleCities) SetCode(v string) *TrainStationSearchResponseBodyModuleCities {
	s.Code = &v
	return s
}

func (s *TrainStationSearchResponseBodyModuleCities) SetName(v string) *TrainStationSearchResponseBodyModuleCities {
	s.Name = &v
	return s
}

type TrainStationSearchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TrainStationSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TrainStationSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s TrainStationSearchResponse) GoString() string {
	return s.String()
}

func (s *TrainStationSearchResponse) SetHeaders(v map[string]*string) *TrainStationSearchResponse {
	s.Headers = v
	return s
}

func (s *TrainStationSearchResponse) SetStatusCode(v int32) *TrainStationSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainStationSearchResponse) SetBody(v *TrainStationSearchResponseBody) *TrainStationSearchResponse {
	s.Body = v
	return s
}

type UserQueryHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s UserQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s UserQueryHeaders) GoString() string {
	return s.String()
}

func (s *UserQueryHeaders) SetCommonHeaders(v map[string]*string) *UserQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UserQueryHeaders) SetXAcsBtripSoCorpToken(v string) *UserQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

type UserQueryRequest struct {
	ModifiedTimeGreaterOrEqualThan *string `json:"modified_time_greater_or_equal_than,omitempty" xml:"modified_time_greater_or_equal_than,omitempty"`
	PageSize                       *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PageToken                      *string `json:"page_token,omitempty" xml:"page_token,omitempty"`
	ThirdPartJobNo                 *string `json:"third_part_job_no,omitempty" xml:"third_part_job_no,omitempty"`
}

func (s UserQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s UserQueryRequest) GoString() string {
	return s.String()
}

func (s *UserQueryRequest) SetModifiedTimeGreaterOrEqualThan(v string) *UserQueryRequest {
	s.ModifiedTimeGreaterOrEqualThan = &v
	return s
}

func (s *UserQueryRequest) SetPageSize(v int32) *UserQueryRequest {
	s.PageSize = &v
	return s
}

func (s *UserQueryRequest) SetPageToken(v string) *UserQueryRequest {
	s.PageToken = &v
	return s
}

func (s *UserQueryRequest) SetThirdPartJobNo(v string) *UserQueryRequest {
	s.ThirdPartJobNo = &v
	return s
}

type UserQueryResponseBody struct {
	Code      *int32                       `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                      `json:"message,omitempty" xml:"message,omitempty"`
	Module    *UserQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                        `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                      `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s UserQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UserQueryResponseBody) GoString() string {
	return s.String()
}

func (s *UserQueryResponseBody) SetCode(v int32) *UserQueryResponseBody {
	s.Code = &v
	return s
}

func (s *UserQueryResponseBody) SetMessage(v string) *UserQueryResponseBody {
	s.Message = &v
	return s
}

func (s *UserQueryResponseBody) SetModule(v *UserQueryResponseBodyModule) *UserQueryResponseBody {
	s.Module = v
	return s
}

func (s *UserQueryResponseBody) SetRequestId(v string) *UserQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UserQueryResponseBody) SetSuccess(v bool) *UserQueryResponseBody {
	s.Success = &v
	return s
}

func (s *UserQueryResponseBody) SetTraceId(v string) *UserQueryResponseBody {
	s.TraceId = &v
	return s
}

type UserQueryResponseBodyModule struct {
	HasMore   *bool                               `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Items     []*UserQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageToken *string                             `json:"page_token,omitempty" xml:"page_token,omitempty"`
	Total     *int64                              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s UserQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s UserQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UserQueryResponseBodyModule) SetHasMore(v bool) *UserQueryResponseBodyModule {
	s.HasMore = &v
	return s
}

func (s *UserQueryResponseBodyModule) SetItems(v []*UserQueryResponseBodyModuleItems) *UserQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *UserQueryResponseBodyModule) SetPageToken(v string) *UserQueryResponseBodyModule {
	s.PageToken = &v
	return s
}

func (s *UserQueryResponseBodyModule) SetTotal(v int64) *UserQueryResponseBodyModule {
	s.Total = &v
	return s
}

type UserQueryResponseBodyModuleItems struct {
	EmployeeNick        *string `json:"employee_nick,omitempty" xml:"employee_nick,omitempty"`
	ThirdPartEmployeeId *string `json:"third_part_employee_id,omitempty" xml:"third_part_employee_id,omitempty"`
	ThirdPartJobNo      *string `json:"third_part_job_no,omitempty" xml:"third_part_job_no,omitempty"`
}

func (s UserQueryResponseBodyModuleItems) String() string {
	return tea.Prettify(s)
}

func (s UserQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *UserQueryResponseBodyModuleItems) SetEmployeeNick(v string) *UserQueryResponseBodyModuleItems {
	s.EmployeeNick = &v
	return s
}

func (s *UserQueryResponseBodyModuleItems) SetThirdPartEmployeeId(v string) *UserQueryResponseBodyModuleItems {
	s.ThirdPartEmployeeId = &v
	return s
}

func (s *UserQueryResponseBodyModuleItems) SetThirdPartJobNo(v string) *UserQueryResponseBodyModuleItems {
	s.ThirdPartJobNo = &v
	return s
}

type UserQueryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UserQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UserQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s UserQueryResponse) GoString() string {
	return s.String()
}

func (s *UserQueryResponse) SetHeaders(v map[string]*string) *UserQueryResponse {
	s.Headers = v
	return s
}

func (s *UserQueryResponse) SetStatusCode(v int32) *UserQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *UserQueryResponse) SetBody(v *UserQueryResponseBody) *UserQueryResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("btripopen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AccessToken(request *AccessTokenRequest) (_result *AccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AccessTokenResponse{}
	_body, _err := client.AccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AccessTokenWithOptions(request *AccessTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["app_key"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		query["app_secret"] = request.AppSecret
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AccessToken"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/btrip-open-auth/v1/access-token/action/take"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddressGet(request *AddressGetRequest) (_result *AddressGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddressGetHeaders{}
	_result = &AddressGetResponse{}
	_body, _err := client.AddressGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddressGetWithOptions(request *AddressGetRequest, headers *AddressGetHeaders, runtime *util.RuntimeOptions) (_result *AddressGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["action_type"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.ItineraryId)) {
		query["itinerary_id"] = request.ItineraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddressGet"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/v1/address"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddressGetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AirportSearch(request *AirportSearchRequest) (_result *AirportSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AirportSearchHeaders{}
	_result = &AirportSearchResponse{}
	_body, _err := client.AirportSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AirportSearchWithOptions(request *AirportSearchRequest, headers *AirportSearchHeaders, runtime *util.RuntimeOptions) (_result *AirportSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AirportSearch"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/city/v1/airport"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AirportSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyAdd(request *ApplyAddRequest) (_result *ApplyAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyAddHeaders{}
	_result = &ApplyAddResponse{}
	_body, _err := client.ApplyAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyAddWithOptions(tmpReq *ApplyAddRequest, headers *ApplyAddHeaders, runtime *util.RuntimeOptions) (_result *ApplyAddResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ApplyAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExternalTravelerList)) {
		request.ExternalTravelerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalTravelerList, tea.String("external_traveler_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ExternalTravelerStandard))) {
		request.ExternalTravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ExternalTravelerStandard), tea.String("external_traveler_standard"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.HotelShare))) {
		request.HotelShareShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.HotelShare), tea.String("hotel_share"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ItineraryList)) {
		request.ItineraryListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItineraryList, tea.String("itinerary_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ItinerarySetList)) {
		request.ItinerarySetListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItinerarySetList, tea.String("itinerary_set_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TravelerList)) {
		request.TravelerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerList, tea.String("traveler_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TravelerStandard)) {
		request.TravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerStandard, tea.String("traveler_standard"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InternationalFlightCabins)) {
		query["international_flight_cabins"] = request.InternationalFlightCabins
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Budget)) {
		body["budget"] = request.Budget
	}

	if !tea.BoolValue(util.IsUnset(request.BudgetMerge)) {
		body["budget_merge"] = request.BudgetMerge
	}

	if !tea.BoolValue(util.IsUnset(request.CorpName)) {
		body["corp_name"] = request.CorpName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartId)) {
		body["depart_id"] = request.DepartId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartName)) {
		body["depart_name"] = request.DepartName
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalTravelerListShrink)) {
		body["external_traveler_list"] = request.ExternalTravelerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalTravelerStandardShrink)) {
		body["external_traveler_standard"] = request.ExternalTravelerStandardShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FlightBudget)) {
		body["flight_budget"] = request.FlightBudget
	}

	if !tea.BoolValue(util.IsUnset(request.HotelBudget)) {
		body["hotel_budget"] = request.HotelBudget
	}

	if !tea.BoolValue(util.IsUnset(request.HotelShareShrink)) {
		body["hotel_share"] = request.HotelShareShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ItineraryListShrink)) {
		body["itinerary_list"] = request.ItineraryListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ItineraryRule)) {
		body["itinerary_rule"] = request.ItineraryRule
	}

	if !tea.BoolValue(util.IsUnset(request.ItinerarySetListShrink)) {
		body["itinerary_set_list"] = request.ItinerarySetListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LimitTraveler)) {
		body["limit_traveler"] = request.LimitTraveler
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartApplyId)) {
		body["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartBusinessId)) {
		body["thirdpart_business_id"] = request.ThirdpartBusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.TogetherBookRule)) {
		body["together_book_rule"] = request.TogetherBookRule
	}

	if !tea.BoolValue(util.IsUnset(request.TrainBudget)) {
		body["train_budget"] = request.TrainBudget
	}

	if !tea.BoolValue(util.IsUnset(request.TravelerListShrink)) {
		body["traveler_list"] = request.TravelerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TravelerStandardShrink)) {
		body["traveler_standard"] = request.TravelerStandardShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TripCause)) {
		body["trip_cause"] = request.TripCause
	}

	if !tea.BoolValue(util.IsUnset(request.TripDay)) {
		body["trip_day"] = request.TripDay
	}

	if !tea.BoolValue(util.IsUnset(request.TripTitle)) {
		body["trip_title"] = request.TripTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UnionNo)) {
		body["union_no"] = request.UnionNo
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["user_name"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleBudget)) {
		body["vehicle_budget"] = request.VehicleBudget
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyAdd"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/biz-trip"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyAddResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyApprove(request *ApplyApproveRequest) (_result *ApplyApproveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyApproveHeaders{}
	_result = &ApplyApproveResponse{}
	_body, _err := client.ApplyApproveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyApproveWithOptions(request *ApplyApproveRequest, headers *ApplyApproveHeaders, runtime *util.RuntimeOptions) (_result *ApplyApproveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		body["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.OperateTime)) {
		body["operate_time"] = request.OperateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["user_name"] = request.UserName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyApprove"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/biz-trip/action/approve"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyApproveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyListQuery(request *ApplyListQueryRequest) (_result *ApplyListQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyListQueryHeaders{}
	_result = &ApplyListQueryResponse{}
	_body, _err := client.ApplyListQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyListQueryWithOptions(request *ApplyListQueryRequest, headers *ApplyListQueryHeaders, runtime *util.RuntimeOptions) (_result *ApplyListQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllApply)) {
		query["all_apply"] = request.AllApply
	}

	if !tea.BoolValue(util.IsUnset(request.DepartId)) {
		query["depart_id"] = request.DepartId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["end_time"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GmtModified)) {
		query["gmt_modified"] = request.GmtModified
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyShangLvApply)) {
		query["only_shang_lv_apply"] = request.OnlyShangLvApply
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["start_time"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UnionNo)) {
		query["union_no"] = request.UnionNo
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyListQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/biz-trips"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyListQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyModify(request *ApplyModifyRequest) (_result *ApplyModifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyModifyHeaders{}
	_result = &ApplyModifyResponse{}
	_body, _err := client.ApplyModifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyModifyWithOptions(tmpReq *ApplyModifyRequest, headers *ApplyModifyHeaders, runtime *util.RuntimeOptions) (_result *ApplyModifyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ApplyModifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExternalTravelerList)) {
		request.ExternalTravelerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalTravelerList, tea.String("external_traveler_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ExternalTravelerStandard))) {
		request.ExternalTravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ExternalTravelerStandard), tea.String("external_traveler_standard"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.HotelShare))) {
		request.HotelShareShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.HotelShare), tea.String("hotel_share"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ItineraryList)) {
		request.ItineraryListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItineraryList, tea.String("itinerary_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ItinerarySetList)) {
		request.ItinerarySetListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItinerarySetList, tea.String("itinerary_set_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TravelerList)) {
		request.TravelerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerList, tea.String("traveler_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TravelerStandard)) {
		request.TravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerStandard, tea.String("traveler_standard"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Budget)) {
		body["budget"] = request.Budget
	}

	if !tea.BoolValue(util.IsUnset(request.BudgetMerge)) {
		body["budget_merge"] = request.BudgetMerge
	}

	if !tea.BoolValue(util.IsUnset(request.CorpName)) {
		body["corp_name"] = request.CorpName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartId)) {
		body["depart_id"] = request.DepartId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartName)) {
		body["depart_name"] = request.DepartName
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalTravelerListShrink)) {
		body["external_traveler_list"] = request.ExternalTravelerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalTravelerStandardShrink)) {
		body["external_traveler_standard"] = request.ExternalTravelerStandardShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FlightBudget)) {
		body["flight_budget"] = request.FlightBudget
	}

	if !tea.BoolValue(util.IsUnset(request.HotelBudget)) {
		body["hotel_budget"] = request.HotelBudget
	}

	if !tea.BoolValue(util.IsUnset(request.HotelShareShrink)) {
		body["hotel_share"] = request.HotelShareShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ItineraryListShrink)) {
		body["itinerary_list"] = request.ItineraryListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ItineraryRule)) {
		body["itinerary_rule"] = request.ItineraryRule
	}

	if !tea.BoolValue(util.IsUnset(request.ItinerarySetListShrink)) {
		body["itinerary_set_list"] = request.ItinerarySetListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LimitTraveler)) {
		body["limit_traveler"] = request.LimitTraveler
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartApplyId)) {
		body["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartBusinessId)) {
		body["thirdpart_business_id"] = request.ThirdpartBusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.TogetherBookRule)) {
		body["together_book_rule"] = request.TogetherBookRule
	}

	if !tea.BoolValue(util.IsUnset(request.TrainBudget)) {
		body["train_budget"] = request.TrainBudget
	}

	if !tea.BoolValue(util.IsUnset(request.TravelerListShrink)) {
		body["traveler_list"] = request.TravelerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TravelerStandardShrink)) {
		body["traveler_standard"] = request.TravelerStandardShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TripCause)) {
		body["trip_cause"] = request.TripCause
	}

	if !tea.BoolValue(util.IsUnset(request.TripDay)) {
		body["trip_day"] = request.TripDay
	}

	if !tea.BoolValue(util.IsUnset(request.TripTitle)) {
		body["trip_title"] = request.TripTitle
	}

	if !tea.BoolValue(util.IsUnset(request.UnionNo)) {
		body["union_no"] = request.UnionNo
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["user_name"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VehicleBudget)) {
		body["vehicle_budget"] = request.VehicleBudget
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyModify"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/biz-trip"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyQuery(request *ApplyQueryRequest) (_result *ApplyQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyQueryHeaders{}
	_result = &ApplyQueryResponse{}
	_body, _err := client.ApplyQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyQueryWithOptions(request *ApplyQueryRequest, headers *ApplyQueryHeaders, runtime *util.RuntimeOptions) (_result *ApplyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyShowId)) {
		query["apply_show_id"] = request.ApplyShowId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartApplyId)) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/biz-trip"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CarApplyAdd(request *CarApplyAddRequest) (_result *CarApplyAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CarApplyAddHeaders{}
	_result = &CarApplyAddResponse{}
	_body, _err := client.CarApplyAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CarApplyAddWithOptions(request *CarApplyAddRequest, headers *CarApplyAddHeaders, runtime *util.RuntimeOptions) (_result *CarApplyAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cause)) {
		body["cause"] = request.Cause
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		body["city"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Date)) {
		body["date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.FinishedDate)) {
		body["finished_date"] = request.FinishedDate
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectCode)) {
		body["project_code"] = request.ProjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["project_name"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartApplyId)) {
		body["third_part_apply_id"] = request.ThirdPartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartCostCenterId)) {
		body["third_part_cost_center_id"] = request.ThirdPartCostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartInvoiceId)) {
		body["third_part_invoice_id"] = request.ThirdPartInvoiceId
	}

	if !tea.BoolValue(util.IsUnset(request.TimesTotal)) {
		body["times_total"] = request.TimesTotal
	}

	if !tea.BoolValue(util.IsUnset(request.TimesType)) {
		body["times_type"] = request.TimesType
	}

	if !tea.BoolValue(util.IsUnset(request.TimesUsed)) {
		body["times_used"] = request.TimesUsed
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CarApplyAdd"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/car"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CarApplyAddResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CarApplyModify(request *CarApplyModifyRequest) (_result *CarApplyModifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CarApplyModifyHeaders{}
	_result = &CarApplyModifyResponse{}
	_body, _err := client.CarApplyModifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CarApplyModifyWithOptions(request *CarApplyModifyRequest, headers *CarApplyModifyHeaders, runtime *util.RuntimeOptions) (_result *CarApplyModifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperateTime)) {
		body["operate_time"] = request.OperateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartApplyId)) {
		body["third_part_apply_id"] = request.ThirdPartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CarApplyModify"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/car"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CarApplyModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CarApplyQuery(request *CarApplyQueryRequest) (_result *CarApplyQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CarApplyQueryHeaders{}
	_result = &CarApplyQueryResponse{}
	_body, _err := client.CarApplyQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CarApplyQueryWithOptions(request *CarApplyQueryRequest, headers *CarApplyQueryHeaders, runtime *util.RuntimeOptions) (_result *CarApplyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatedEndAt)) {
		query["created_end_at"] = request.CreatedEndAt
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedStartAt)) {
		query["created_start_at"] = request.CreatedStartAt
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["page_number"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartApplyId)) {
		query["third_part_apply_id"] = request.ThirdPartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CarApplyQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/car"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CarApplyQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CarBillSettlementQuery(request *CarBillSettlementQueryRequest) (_result *CarBillSettlementQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CarBillSettlementQueryHeaders{}
	_result = &CarBillSettlementQueryResponse{}
	_body, _err := client.CarBillSettlementQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CarBillSettlementQueryWithOptions(request *CarBillSettlementQueryRequest, headers *CarBillSettlementQueryHeaders, runtime *util.RuntimeOptions) (_result *CarBillSettlementQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["page_no"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["period_end"] = request.PeriodEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["period_start"] = request.PeriodStart
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CarBillSettlementQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/car/v1/bill-settlement"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CarBillSettlementQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CarOrderListQuery(request *CarOrderListQueryRequest) (_result *CarOrderListQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CarOrderListQueryHeaders{}
	_result = &CarOrderListQueryResponse{}
	_body, _err := client.CarOrderListQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CarOrderListQueryWithOptions(request *CarOrderListQueryRequest, headers *CarOrderListQueryHeaders, runtime *util.RuntimeOptions) (_result *CarOrderListQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllApply)) {
		query["all_apply"] = request.AllApply
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartId)) {
		query["depart_id"] = request.DepartId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["end_time"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["start_time"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartApplyId)) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateEndTime)) {
		query["update_end_time"] = request.UpdateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStartTime)) {
		query["update_start_time"] = request.UpdateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CarOrderListQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/car/v1/order-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CarOrderListQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CitySearch(request *CitySearchRequest) (_result *CitySearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CitySearchHeaders{}
	_result = &CitySearchResponse{}
	_body, _err := client.CitySearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CitySearchWithOptions(request *CitySearchRequest, headers *CitySearchHeaders, runtime *util.RuntimeOptions) (_result *CitySearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CitySearch"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/city/v1/city"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CitySearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommonApplyQuery(request *CommonApplyQueryRequest) (_result *CommonApplyQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CommonApplyQueryHeaders{}
	_result = &CommonApplyQueryResponse{}
	_body, _err := client.CommonApplyQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CommonApplyQueryWithOptions(request *CommonApplyQueryRequest, headers *CommonApplyQueryHeaders, runtime *util.RuntimeOptions) (_result *CommonApplyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCategory)) {
		query["biz_category"] = request.BizCategory
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CommonApplyQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/common"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CommonApplyQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommonApplySync(request *CommonApplySyncRequest) (_result *CommonApplySyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CommonApplySyncHeaders{}
	_result = &CommonApplySyncResponse{}
	_body, _err := client.CommonApplySyncWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CommonApplySyncWithOptions(request *CommonApplySyncRequest, headers *CommonApplySyncHeaders, runtime *util.RuntimeOptions) (_result *CommonApplySyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCategory)) {
		query["biz_category"] = request.BizCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyFlowId)) {
		query["thirdparty_flow_id"] = request.ThirdpartyFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CommonApplySync"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/syn-common"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CommonApplySyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CorpToken(request *CorpTokenRequest) (_result *CorpTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CorpTokenHeaders{}
	_result = &CorpTokenResponse{}
	_body, _err := client.CorpTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CorpTokenWithOptions(request *CorpTokenRequest, headers *CorpTokenHeaders, runtime *util.RuntimeOptions) (_result *CorpTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corp_id"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripAccessToken)) {
		realHeaders["x-acs-btrip-access-token"] = util.ToJSONString(headers.XAcsBtripAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CorpToken"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/btrip-open-auth/v1/corp-token/action/take"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CorpTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CostCenterDelete(request *CostCenterDeleteRequest) (_result *CostCenterDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CostCenterDeleteHeaders{}
	_result = &CostCenterDeleteResponse{}
	_body, _err := client.CostCenterDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CostCenterDeleteWithOptions(request *CostCenterDeleteRequest, headers *CostCenterDeleteHeaders, runtime *util.RuntimeOptions) (_result *CostCenterDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ThirdpartId)) {
		query["thirdpart_id"] = request.ThirdpartId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CostCenterDelete"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/costcenter/v1/delete-costcenter"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CostCenterDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CostCenterModify(request *CostCenterModifyRequest) (_result *CostCenterModifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CostCenterModifyHeaders{}
	_result = &CostCenterModifyResponse{}
	_body, _err := client.CostCenterModifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CostCenterModifyWithOptions(request *CostCenterModifyRequest, headers *CostCenterModifyHeaders, runtime *util.RuntimeOptions) (_result *CostCenterModifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlipayNo)) {
		body["alipay_no"] = request.AlipayNo
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		body["number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartId)) {
		body["thirdpart_id"] = request.ThirdpartId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CostCenterModify"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/costcenter/v1/modify-costcenter"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CostCenterModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CostCenterQuery(request *CostCenterQueryRequest) (_result *CostCenterQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CostCenterQueryHeaders{}
	_result = &CostCenterQueryResponse{}
	_body, _err := client.CostCenterQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CostCenterQueryWithOptions(request *CostCenterQueryRequest, headers *CostCenterQueryHeaders, runtime *util.RuntimeOptions) (_result *CostCenterQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedOrgEntity)) {
		query["need_org_entity"] = request.NeedOrgEntity
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartId)) {
		query["thirdpart_id"] = request.ThirdpartId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CostCenterQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/costcenter/v1/costcenter"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CostCenterQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CostCenterSave(request *CostCenterSaveRequest) (_result *CostCenterSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CostCenterSaveHeaders{}
	_result = &CostCenterSaveResponse{}
	_body, _err := client.CostCenterSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CostCenterSaveWithOptions(request *CostCenterSaveRequest, headers *CostCenterSaveHeaders, runtime *util.RuntimeOptions) (_result *CostCenterSaveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlipayNo)) {
		body["alipay_no"] = request.AlipayNo
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		body["number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartId)) {
		body["thirdpart_id"] = request.ThirdpartId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CostCenterSave"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/costcenter/v1/save-costcenter"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CostCenterSaveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DepartmentSave(request *DepartmentSaveRequest) (_result *DepartmentSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DepartmentSaveHeaders{}
	_result = &DepartmentSaveResponse{}
	_body, _err := client.DepartmentSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DepartmentSaveWithOptions(tmpReq *DepartmentSaveRequest, headers *DepartmentSaveHeaders, runtime *util.RuntimeOptions) (_result *DepartmentSaveResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DepartmentSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepartList)) {
		request.DepartListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepartList, tea.String("depart_list"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartListShrink)) {
		body["depart_list"] = request.DepartListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DepartmentSave"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/department/v1/department"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DepartmentSaveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EntityAdd(request *EntityAddRequest) (_result *EntityAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EntityAddHeaders{}
	_result = &EntityAddResponse{}
	_body, _err := client.EntityAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EntityAddWithOptions(tmpReq *EntityAddRequest, headers *EntityAddHeaders, runtime *util.RuntimeOptions) (_result *EntityAddResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EntityAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityDOList)) {
		request.EntityDOListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityDOList, tea.String("entity_d_o_list"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityDOListShrink)) {
		body["entity_d_o_list"] = request.EntityDOListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartId)) {
		body["thirdpart_id"] = request.ThirdpartId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EntityAdd"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/costcenter/v1/add-entity"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EntityAddResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EntityDelete(request *EntityDeleteRequest) (_result *EntityDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EntityDeleteHeaders{}
	_result = &EntityDeleteResponse{}
	_body, _err := client.EntityDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EntityDeleteWithOptions(tmpReq *EntityDeleteRequest, headers *EntityDeleteHeaders, runtime *util.RuntimeOptions) (_result *EntityDeleteResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EntityDeleteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityDOList)) {
		request.EntityDOListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityDOList, tea.String("entity_d_o_list"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DelAll)) {
		query["del_all"] = request.DelAll
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartId)) {
		query["thirdpart_id"] = request.ThirdpartId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityDOListShrink)) {
		body["entity_d_o_list"] = request.EntityDOListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EntityDelete"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/costcenter/v1/entity/action/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EntityDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EntitySet(request *EntitySetRequest) (_result *EntitySetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EntitySetHeaders{}
	_result = &EntitySetResponse{}
	_body, _err := client.EntitySetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EntitySetWithOptions(tmpReq *EntitySetRequest, headers *EntitySetHeaders, runtime *util.RuntimeOptions) (_result *EntitySetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EntitySetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityDOList)) {
		request.EntityDOListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityDOList, tea.String("entity_d_o_list"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityDOListShrink)) {
		body["entity_d_o_list"] = request.EntityDOListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartId)) {
		body["thirdpart_id"] = request.ThirdpartId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EntitySet"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/costcenter/v1/set-entity"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EntitySetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExceedApplySync(request *ExceedApplySyncRequest) (_result *ExceedApplySyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExceedApplySyncHeaders{}
	_result = &ExceedApplySyncResponse{}
	_body, _err := client.ExceedApplySyncWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExceedApplySyncWithOptions(request *ExceedApplySyncRequest, headers *ExceedApplySyncHeaders, runtime *util.RuntimeOptions) (_result *ExceedApplySyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCategory)) {
		query["biz_category"] = request.BizCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyFlowId)) {
		query["thirdparty_flow_id"] = request.ThirdpartyFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExceedApplySync"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/syn-exceed"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExceedApplySyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlightBillSettlementQuery(request *FlightBillSettlementQueryRequest) (_result *FlightBillSettlementQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FlightBillSettlementQueryHeaders{}
	_result = &FlightBillSettlementQueryResponse{}
	_body, _err := client.FlightBillSettlementQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlightBillSettlementQueryWithOptions(request *FlightBillSettlementQueryRequest, headers *FlightBillSettlementQueryHeaders, runtime *util.RuntimeOptions) (_result *FlightBillSettlementQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["page_no"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["period_end"] = request.PeriodEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["period_start"] = request.PeriodStart
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FlightBillSettlementQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/flight/v1/bill-settlement"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FlightBillSettlementQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlightExceedApplyQuery(request *FlightExceedApplyQueryRequest) (_result *FlightExceedApplyQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FlightExceedApplyQueryHeaders{}
	_result = &FlightExceedApplyQueryResponse{}
	_body, _err := client.FlightExceedApplyQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlightExceedApplyQueryWithOptions(request *FlightExceedApplyQueryRequest, headers *FlightExceedApplyQueryHeaders, runtime *util.RuntimeOptions) (_result *FlightExceedApplyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FlightExceedApplyQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/flight-exceed"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FlightExceedApplyQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlightOrderListQuery(request *FlightOrderListQueryRequest) (_result *FlightOrderListQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FlightOrderListQueryHeaders{}
	_result = &FlightOrderListQueryResponse{}
	_body, _err := client.FlightOrderListQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlightOrderListQueryWithOptions(request *FlightOrderListQueryRequest, headers *FlightOrderListQueryHeaders, runtime *util.RuntimeOptions) (_result *FlightOrderListQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllApply)) {
		query["all_apply"] = request.AllApply
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartId)) {
		query["depart_id"] = request.DepartId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["end_time"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["start_time"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartApplyId)) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateEndTime)) {
		query["update_end_time"] = request.UpdateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStartTime)) {
		query["update_start_time"] = request.UpdateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FlightOrderListQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/flight/v1/order-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FlightOrderListQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlightOrderQuery(request *FlightOrderQueryRequest) (_result *FlightOrderQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FlightOrderQueryHeaders{}
	_result = &FlightOrderQueryResponse{}
	_body, _err := client.FlightOrderQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlightOrderQueryWithOptions(request *FlightOrderQueryRequest, headers *FlightOrderQueryHeaders, runtime *util.RuntimeOptions) (_result *FlightOrderQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["order_id"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FlightOrderQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/flight/v1/order"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FlightOrderQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HotelBillSettlementQuery(request *HotelBillSettlementQueryRequest) (_result *HotelBillSettlementQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HotelBillSettlementQueryHeaders{}
	_result = &HotelBillSettlementQueryResponse{}
	_body, _err := client.HotelBillSettlementQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HotelBillSettlementQueryWithOptions(request *HotelBillSettlementQueryRequest, headers *HotelBillSettlementQueryHeaders, runtime *util.RuntimeOptions) (_result *HotelBillSettlementQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["page_no"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["period_end"] = request.PeriodEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["period_start"] = request.PeriodStart
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HotelBillSettlementQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/hotel/v1/bill-settlement"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &HotelBillSettlementQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HotelExceedApplyQuery(request *HotelExceedApplyQueryRequest) (_result *HotelExceedApplyQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HotelExceedApplyQueryHeaders{}
	_result = &HotelExceedApplyQueryResponse{}
	_body, _err := client.HotelExceedApplyQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HotelExceedApplyQueryWithOptions(request *HotelExceedApplyQueryRequest, headers *HotelExceedApplyQueryHeaders, runtime *util.RuntimeOptions) (_result *HotelExceedApplyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HotelExceedApplyQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/hotel-exceed"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &HotelExceedApplyQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HotelOrderListQuery(request *HotelOrderListQueryRequest) (_result *HotelOrderListQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HotelOrderListQueryHeaders{}
	_result = &HotelOrderListQueryResponse{}
	_body, _err := client.HotelOrderListQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HotelOrderListQueryWithOptions(request *HotelOrderListQueryRequest, headers *HotelOrderListQueryHeaders, runtime *util.RuntimeOptions) (_result *HotelOrderListQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllApply)) {
		query["all_apply"] = request.AllApply
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartId)) {
		query["depart_id"] = request.DepartId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["end_time"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["start_time"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartApplyId)) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateEndTime)) {
		query["update_end_time"] = request.UpdateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStartTime)) {
		query["update_start_time"] = request.UpdateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HotelOrderListQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/hotel/v1/order-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &HotelOrderListQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IeFlightBillSettlementQuery(request *IeFlightBillSettlementQueryRequest) (_result *IeFlightBillSettlementQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IeFlightBillSettlementQueryHeaders{}
	_result = &IeFlightBillSettlementQueryResponse{}
	_body, _err := client.IeFlightBillSettlementQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IeFlightBillSettlementQueryWithOptions(request *IeFlightBillSettlementQueryRequest, headers *IeFlightBillSettlementQueryHeaders, runtime *util.RuntimeOptions) (_result *IeFlightBillSettlementQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["page_no"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["period_end"] = request.PeriodEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["period_start"] = request.PeriodStart
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IeFlightBillSettlementQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ie-flight/v1/bill-settlement"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &IeFlightBillSettlementQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvoiceAdd(request *InvoiceAddRequest) (_result *InvoiceAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvoiceAddHeaders{}
	_result = &InvoiceAddResponse{}
	_body, _err := client.InvoiceAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvoiceAddWithOptions(request *InvoiceAddRequest, headers *InvoiceAddHeaders, runtime *util.RuntimeOptions) (_result *InvoiceAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.BankName)) {
		body["bank_name"] = request.BankName
	}

	if !tea.BoolValue(util.IsUnset(request.BankNo)) {
		body["bank_no"] = request.BankNo
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNo)) {
		body["tax_no"] = request.TaxNo
	}

	if !tea.BoolValue(util.IsUnset(request.Tel)) {
		body["tel"] = request.Tel
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		body["third_part_id"] = request.ThirdPartId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvoiceAdd"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/invoice/v1/add-invoice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvoiceAddResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvoiceDelete(request *InvoiceDeleteRequest) (_result *InvoiceDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvoiceDeleteHeaders{}
	_result = &InvoiceDeleteResponse{}
	_body, _err := client.InvoiceDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvoiceDeleteWithOptions(request *InvoiceDeleteRequest, headers *InvoiceDeleteHeaders, runtime *util.RuntimeOptions) (_result *InvoiceDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		query["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvoiceDelete"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/invoice/v1/invoice"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InvoiceDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvoiceModify(request *InvoiceModifyRequest) (_result *InvoiceModifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvoiceModifyHeaders{}
	_result = &InvoiceModifyResponse{}
	_body, _err := client.InvoiceModifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvoiceModifyWithOptions(request *InvoiceModifyRequest, headers *InvoiceModifyHeaders, runtime *util.RuntimeOptions) (_result *InvoiceModifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.BankName)) {
		body["bank_name"] = request.BankName
	}

	if !tea.BoolValue(util.IsUnset(request.BankNo)) {
		body["bank_no"] = request.BankNo
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNo)) {
		body["tax_no"] = request.TaxNo
	}

	if !tea.BoolValue(util.IsUnset(request.Tel)) {
		body["tel"] = request.Tel
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		body["third_part_id"] = request.ThirdPartId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvoiceModify"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/invoice/v1/invoice"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvoiceModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvoiceRuleSave(request *InvoiceRuleSaveRequest) (_result *InvoiceRuleSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvoiceRuleSaveHeaders{}
	_result = &InvoiceRuleSaveResponse{}
	_body, _err := client.InvoiceRuleSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvoiceRuleSaveWithOptions(tmpReq *InvoiceRuleSaveRequest, headers *InvoiceRuleSaveHeaders, runtime *util.RuntimeOptions) (_result *InvoiceRuleSaveResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InvoiceRuleSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Entities)) {
		request.EntitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Entities, tea.String("entities"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllEmploye)) {
		body["all_employe"] = request.AllEmploye
	}

	if !tea.BoolValue(util.IsUnset(request.EntitiesShrink)) {
		body["entities"] = request.EntitiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		body["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvoiceRuleSave"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/invoice/v1/invoice-rule"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvoiceRuleSaveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvoiceSearch(request *InvoiceSearchRequest) (_result *InvoiceSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvoiceSearchHeaders{}
	_result = &InvoiceSearchResponse{}
	_body, _err := client.InvoiceSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvoiceSearchWithOptions(request *InvoiceSearchRequest, headers *InvoiceSearchHeaders, runtime *util.RuntimeOptions) (_result *InvoiceSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvoiceSearch"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/invoice/v1/invoice"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InvoiceSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsvUserSave(request *IsvUserSaveRequest) (_result *IsvUserSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IsvUserSaveHeaders{}
	_result = &IsvUserSaveResponse{}
	_body, _err := client.IsvUserSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsvUserSaveWithOptions(tmpReq *IsvUserSaveRequest, headers *IsvUserSaveHeaders, runtime *util.RuntimeOptions) (_result *IsvUserSaveResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &IsvUserSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserList)) {
		request.UserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserList, tea.String("user_list"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserListShrink)) {
		body["user_list"] = request.UserListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IsvUserSave"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/isvuser/v1/isvuser"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IsvUserSaveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MonthBillGet(request *MonthBillGetRequest) (_result *MonthBillGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MonthBillGetHeaders{}
	_result = &MonthBillGetResponse{}
	_body, _err := client.MonthBillGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MonthBillGetWithOptions(request *MonthBillGetRequest, headers *MonthBillGetHeaders, runtime *util.RuntimeOptions) (_result *MonthBillGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillMonth)) {
		query["bill_month"] = request.BillMonth
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MonthBillGet"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/v1/month-bill"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MonthBillGetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProjectAdd(request *ProjectAddRequest) (_result *ProjectAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProjectAddHeaders{}
	_result = &ProjectAddResponse{}
	_body, _err := client.ProjectAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProjectAddWithOptions(request *ProjectAddRequest, headers *ProjectAddHeaders, runtime *util.RuntimeOptions) (_result *ProjectAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["project_name"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartCostCenterId)) {
		body["third_part_cost_center_id"] = request.ThirdPartCostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		body["third_part_id"] = request.ThirdPartId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartInvoiceId)) {
		body["third_part_invoice_id"] = request.ThirdPartInvoiceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ProjectAdd"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/cost/v1/project"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ProjectAddResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProjectDelete(request *ProjectDeleteRequest) (_result *ProjectDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProjectDeleteHeaders{}
	_result = &ProjectDeleteResponse{}
	_body, _err := client.ProjectDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProjectDeleteWithOptions(request *ProjectDeleteRequest, headers *ProjectDeleteHeaders, runtime *util.RuntimeOptions) (_result *ProjectDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		query["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ProjectDelete"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/cost/v1/project"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ProjectDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProjectModify(request *ProjectModifyRequest) (_result *ProjectModifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProjectModifyHeaders{}
	_result = &ProjectModifyResponse{}
	_body, _err := client.ProjectModifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProjectModifyWithOptions(request *ProjectModifyRequest, headers *ProjectModifyHeaders, runtime *util.RuntimeOptions) (_result *ProjectModifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["project_name"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartCostCenterId)) {
		body["third_part_cost_center_id"] = request.ThirdPartCostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		body["third_part_id"] = request.ThirdPartId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartInvoiceId)) {
		body["third_part_invoice_id"] = request.ThirdPartInvoiceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ProjectModify"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/cost/v1/project"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ProjectModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TrainBillSettlementQuery(request *TrainBillSettlementQueryRequest) (_result *TrainBillSettlementQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TrainBillSettlementQueryHeaders{}
	_result = &TrainBillSettlementQueryResponse{}
	_body, _err := client.TrainBillSettlementQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TrainBillSettlementQueryWithOptions(request *TrainBillSettlementQueryRequest, headers *TrainBillSettlementQueryHeaders, runtime *util.RuntimeOptions) (_result *TrainBillSettlementQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["page_no"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["period_end"] = request.PeriodEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["period_start"] = request.PeriodStart
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TrainBillSettlementQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/train/v1/bill-settlement"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TrainBillSettlementQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TrainExceedApplyQuery(request *TrainExceedApplyQueryRequest) (_result *TrainExceedApplyQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TrainExceedApplyQueryHeaders{}
	_result = &TrainExceedApplyQueryResponse{}
	_body, _err := client.TrainExceedApplyQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TrainExceedApplyQueryWithOptions(request *TrainExceedApplyQueryRequest, headers *TrainExceedApplyQueryHeaders, runtime *util.RuntimeOptions) (_result *TrainExceedApplyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TrainExceedApplyQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apply/v1/train-exceed"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TrainExceedApplyQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TrainOrderListQuery(request *TrainOrderListQueryRequest) (_result *TrainOrderListQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TrainOrderListQueryHeaders{}
	_result = &TrainOrderListQueryResponse{}
	_body, _err := client.TrainOrderListQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TrainOrderListQueryWithOptions(request *TrainOrderListQueryRequest, headers *TrainOrderListQueryHeaders, runtime *util.RuntimeOptions) (_result *TrainOrderListQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllApply)) {
		query["all_apply"] = request.AllApply
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["apply_id"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartId)) {
		query["depart_id"] = request.DepartId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["end_time"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["start_time"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartApplyId)) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateEndTime)) {
		query["update_end_time"] = request.UpdateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStartTime)) {
		query["update_start_time"] = request.UpdateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TrainOrderListQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/train/v1/order-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TrainOrderListQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TrainOrderQuery(request *TrainOrderQueryRequest) (_result *TrainOrderQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TrainOrderQueryHeaders{}
	_result = &TrainOrderQueryResponse{}
	_body, _err := client.TrainOrderQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TrainOrderQueryWithOptions(request *TrainOrderQueryRequest, headers *TrainOrderQueryHeaders, runtime *util.RuntimeOptions) (_result *TrainOrderQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["order_id"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TrainOrderQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/train/v1/order"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TrainOrderQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TrainStationSearch(request *TrainStationSearchRequest) (_result *TrainStationSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TrainStationSearchHeaders{}
	_result = &TrainStationSearchResponse{}
	_body, _err := client.TrainStationSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TrainStationSearchWithOptions(request *TrainStationSearchRequest, headers *TrainStationSearchHeaders, runtime *util.RuntimeOptions) (_result *TrainStationSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TrainStationSearch"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/city/v1/train"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TrainStationSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UserQuery(request *UserQueryRequest) (_result *UserQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UserQueryHeaders{}
	_result = &UserQueryResponse{}
	_body, _err := client.UserQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UserQueryWithOptions(request *UserQueryRequest, headers *UserQueryHeaders, runtime *util.RuntimeOptions) (_result *UserQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModifiedTimeGreaterOrEqualThan)) {
		query["modified_time_greater_or_equal_than"] = request.ModifiedTimeGreaterOrEqualThan
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["page_token"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartJobNo)) {
		query["third_part_job_no"] = request.ThirdPartJobNo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsBtripSoCorpToken)) {
		realHeaders["x-acs-btrip-so-corp-token"] = util.ToJSONString(headers.XAcsBtripSoCorpToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UserQuery"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/v1/user"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UserQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
