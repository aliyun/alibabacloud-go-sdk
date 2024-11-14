// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AgentBaseQuery struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s AgentBaseQuery) String() string {
	return tea.Prettify(s)
}

func (s AgentBaseQuery) GoString() string {
	return s.String()
}

func (s *AgentBaseQuery) SetQuery(v string) *AgentBaseQuery {
	s.Query = &v
	return s
}

type CommonAgentQuery struct {
	Query              *string `json:"query,omitempty" xml:"query,omitempty"`
	QuerySceneEnumCode *string `json:"querySceneEnumCode,omitempty" xml:"querySceneEnumCode,omitempty"`
}

func (s CommonAgentQuery) String() string {
	return tea.Prettify(s)
}

func (s CommonAgentQuery) GoString() string {
	return s.String()
}

func (s *CommonAgentQuery) SetQuery(v string) *CommonAgentQuery {
	s.Query = &v
	return s
}

func (s *CommonAgentQuery) SetQuerySceneEnumCode(v string) *CommonAgentQuery {
	s.QuerySceneEnumCode = &v
	return s
}

type QueryResult struct {
	Data []*QueryResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s QueryResult) String() string {
	return tea.Prettify(s)
}

func (s QueryResult) GoString() string {
	return s.String()
}

func (s *QueryResult) SetData(v []*QueryResultData) *QueryResult {
	s.Data = v
	return s
}

type QueryResultData struct {
	Address      *string                  `json:"address,omitempty" xml:"address,omitempty"`
	CityCode     *string                  `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	CityName     *string                  `json:"cityName,omitempty" xml:"cityName,omitempty"`
	DistrictCode *string                  `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	DistrictName *string                  `json:"districtName,omitempty" xml:"districtName,omitempty"`
	Id           *string                  `json:"id,omitempty" xml:"id,omitempty"`
	Images       []*QueryResultDataImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Latitude     *string                  `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude    *string                  `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Metadata     *QueryResultDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	Name         *string                  `json:"name,omitempty" xml:"name,omitempty"`
	ProvinceCode *string                  `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	ProvinceName *string                  `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
	TypeCode     *string                  `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
	Types        *string                  `json:"types,omitempty" xml:"types,omitempty"`
}

func (s QueryResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryResultData) GoString() string {
	return s.String()
}

func (s *QueryResultData) SetAddress(v string) *QueryResultData {
	s.Address = &v
	return s
}

func (s *QueryResultData) SetCityCode(v string) *QueryResultData {
	s.CityCode = &v
	return s
}

func (s *QueryResultData) SetCityName(v string) *QueryResultData {
	s.CityName = &v
	return s
}

func (s *QueryResultData) SetDistrictCode(v string) *QueryResultData {
	s.DistrictCode = &v
	return s
}

func (s *QueryResultData) SetDistrictName(v string) *QueryResultData {
	s.DistrictName = &v
	return s
}

func (s *QueryResultData) SetId(v string) *QueryResultData {
	s.Id = &v
	return s
}

func (s *QueryResultData) SetImages(v []*QueryResultDataImages) *QueryResultData {
	s.Images = v
	return s
}

func (s *QueryResultData) SetLatitude(v string) *QueryResultData {
	s.Latitude = &v
	return s
}

func (s *QueryResultData) SetLongitude(v string) *QueryResultData {
	s.Longitude = &v
	return s
}

func (s *QueryResultData) SetMetadata(v *QueryResultDataMetadata) *QueryResultData {
	s.Metadata = v
	return s
}

func (s *QueryResultData) SetName(v string) *QueryResultData {
	s.Name = &v
	return s
}

func (s *QueryResultData) SetProvinceCode(v string) *QueryResultData {
	s.ProvinceCode = &v
	return s
}

func (s *QueryResultData) SetProvinceName(v string) *QueryResultData {
	s.ProvinceName = &v
	return s
}

func (s *QueryResultData) SetTypeCode(v string) *QueryResultData {
	s.TypeCode = &v
	return s
}

func (s *QueryResultData) SetTypes(v string) *QueryResultData {
	s.Types = &v
	return s
}

type QueryResultDataImages struct {
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	Url   *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryResultDataImages) String() string {
	return tea.Prettify(s)
}

func (s QueryResultDataImages) GoString() string {
	return s.String()
}

func (s *QueryResultDataImages) SetTitle(v string) *QueryResultDataImages {
	s.Title = &v
	return s
}

func (s *QueryResultDataImages) SetUrl(v string) *QueryResultDataImages {
	s.Url = &v
	return s
}

type QueryResultDataMetadata struct {
	BusinessArea      *string `json:"businessArea,omitempty" xml:"businessArea,omitempty"`
	DailyOpeningHours *string `json:"dailyOpeningHours,omitempty" xml:"dailyOpeningHours,omitempty"`
	MainTag           *string `json:"mainTag,omitempty" xml:"mainTag,omitempty"`
	Phone             *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Score             *string `json:"score,omitempty" xml:"score,omitempty"`
	WeeklyOpeningDays *string `json:"weeklyOpeningDays,omitempty" xml:"weeklyOpeningDays,omitempty"`
}

func (s QueryResultDataMetadata) String() string {
	return tea.Prettify(s)
}

func (s QueryResultDataMetadata) GoString() string {
	return s.String()
}

func (s *QueryResultDataMetadata) SetBusinessArea(v string) *QueryResultDataMetadata {
	s.BusinessArea = &v
	return s
}

func (s *QueryResultDataMetadata) SetDailyOpeningHours(v string) *QueryResultDataMetadata {
	s.DailyOpeningHours = &v
	return s
}

func (s *QueryResultDataMetadata) SetMainTag(v string) *QueryResultDataMetadata {
	s.MainTag = &v
	return s
}

func (s *QueryResultDataMetadata) SetPhone(v string) *QueryResultDataMetadata {
	s.Phone = &v
	return s
}

func (s *QueryResultDataMetadata) SetScore(v string) *QueryResultDataMetadata {
	s.Score = &v
	return s
}

func (s *QueryResultDataMetadata) SetWeeklyOpeningDays(v string) *QueryResultDataMetadata {
	s.WeeklyOpeningDays = &v
	return s
}

type BicyclingDirectionRequest struct {
	// example:
	//
	// 39.896463
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s BicyclingDirectionRequest) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionRequest) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionRequest) SetDestinationLatitude(v string) *BicyclingDirectionRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *BicyclingDirectionRequest) SetDestinationLongitude(v string) *BicyclingDirectionRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *BicyclingDirectionRequest) SetOriginLatitude(v string) *BicyclingDirectionRequest {
	s.OriginLatitude = &v
	return s
}

func (s *BicyclingDirectionRequest) SetOriginLongitude(v string) *BicyclingDirectionRequest {
	s.OriginLongitude = &v
	return s
}

type BicyclingDirectionResponseBody struct {
	Data []*BicyclingDirectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// <title>502 Bad Gateway</title>
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BicyclingDirectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionResponseBody) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionResponseBody) SetData(v []*BicyclingDirectionResponseBodyData) *BicyclingDirectionResponseBody {
	s.Data = v
	return s
}

func (s *BicyclingDirectionResponseBody) SetErrorCode(v string) *BicyclingDirectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BicyclingDirectionResponseBody) SetErrorMessage(v string) *BicyclingDirectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BicyclingDirectionResponseBody) SetRequestId(v string) *BicyclingDirectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BicyclingDirectionResponseBody) SetSuccess(v bool) *BicyclingDirectionResponseBody {
	s.Success = &v
	return s
}

type BicyclingDirectionResponseBodyData struct {
	// example:
	//
	// 445
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 38434
	DurationSecond *string                                    `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	Steps          []*BicyclingDirectionResponseBodyDataSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s BicyclingDirectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionResponseBodyData) SetDistanceMeter(v string) *BicyclingDirectionResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *BicyclingDirectionResponseBodyData) SetDurationSecond(v string) *BicyclingDirectionResponseBodyData {
	s.DurationSecond = &v
	return s
}

func (s *BicyclingDirectionResponseBodyData) SetSteps(v []*BicyclingDirectionResponseBodyDataSteps) *BicyclingDirectionResponseBodyData {
	s.Steps = v
	return s
}

type BicyclingDirectionResponseBodyDataSteps struct {
	Cost        *BicyclingDirectionResponseBodyDataStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                      `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                      `json:"orientation,omitempty" xml:"orientation,omitempty"`
	RoadName    *string                                      `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 12939
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s BicyclingDirectionResponseBodyDataSteps) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionResponseBodyDataSteps) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionResponseBodyDataSteps) SetCost(v *BicyclingDirectionResponseBodyDataStepsCost) *BicyclingDirectionResponseBodyDataSteps {
	s.Cost = v
	return s
}

func (s *BicyclingDirectionResponseBodyDataSteps) SetInstruction(v string) *BicyclingDirectionResponseBodyDataSteps {
	s.Instruction = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataSteps) SetOrientation(v string) *BicyclingDirectionResponseBodyDataSteps {
	s.Orientation = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataSteps) SetRoadName(v string) *BicyclingDirectionResponseBodyDataSteps {
	s.RoadName = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataSteps) SetStepDistanceMeter(v string) *BicyclingDirectionResponseBodyDataSteps {
	s.StepDistanceMeter = &v
	return s
}

type BicyclingDirectionResponseBodyDataStepsCost struct {
	// example:
	//
	// 27647
	DurationSecond    *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	// example:
	//
	// xxx
	TollRoads *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls     *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 4
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s BicyclingDirectionResponseBodyDataStepsCost) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionResponseBodyDataStepsCost) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionResponseBodyDataStepsCost) SetDurationSecond(v string) *BicyclingDirectionResponseBodyDataStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataStepsCost) SetTaxiFee(v string) *BicyclingDirectionResponseBodyDataStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataStepsCost) SetTollDistanceMeter(v string) *BicyclingDirectionResponseBodyDataStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataStepsCost) SetTollRoads(v string) *BicyclingDirectionResponseBodyDataStepsCost {
	s.TollRoads = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataStepsCost) SetTolls(v string) *BicyclingDirectionResponseBodyDataStepsCost {
	s.Tolls = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataStepsCost) SetTrafficLights(v string) *BicyclingDirectionResponseBodyDataStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *BicyclingDirectionResponseBodyDataStepsCost) SetTransitFee(v string) *BicyclingDirectionResponseBodyDataStepsCost {
	s.TransitFee = &v
	return s
}

type BicyclingDirectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BicyclingDirectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BicyclingDirectionResponse) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionResponse) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionResponse) SetHeaders(v map[string]*string) *BicyclingDirectionResponse {
	s.Headers = v
	return s
}

func (s *BicyclingDirectionResponse) SetStatusCode(v int32) *BicyclingDirectionResponse {
	s.StatusCode = &v
	return s
}

func (s *BicyclingDirectionResponse) SetBody(v *BicyclingDirectionResponseBody) *BicyclingDirectionResponse {
	s.Body = v
	return s
}

type BicyclingDirectionNovaRequest struct {
	// example:
	//
	// 39.995197
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 117.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s BicyclingDirectionNovaRequest) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionNovaRequest) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionNovaRequest) SetDestinationLatitude(v string) *BicyclingDirectionNovaRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *BicyclingDirectionNovaRequest) SetDestinationLongitude(v string) *BicyclingDirectionNovaRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *BicyclingDirectionNovaRequest) SetOriginLatitude(v string) *BicyclingDirectionNovaRequest {
	s.OriginLatitude = &v
	return s
}

func (s *BicyclingDirectionNovaRequest) SetOriginLongitude(v string) *BicyclingDirectionNovaRequest {
	s.OriginLongitude = &v
	return s
}

type BicyclingDirectionNovaResponseBody struct {
	Data *BicyclingDirectionNovaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// <title>502 Bad Gateway</title>
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BicyclingDirectionNovaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionNovaResponseBody) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionNovaResponseBody) SetData(v *BicyclingDirectionNovaResponseBodyData) *BicyclingDirectionNovaResponseBody {
	s.Data = v
	return s
}

func (s *BicyclingDirectionNovaResponseBody) SetErrorCode(v int32) *BicyclingDirectionNovaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBody) SetErrorMessage(v string) *BicyclingDirectionNovaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBody) SetSuccess(v bool) *BicyclingDirectionNovaResponseBody {
	s.Success = &v
	return s
}

type BicyclingDirectionNovaResponseBodyData struct {
	Count *string `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 39.995197
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string                                        `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
	Paths           []*BicyclingDirectionNovaResponseBodyDataPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// example:
	//
	// 23
	TaxiCost *string `json:"taxiCost,omitempty" xml:"taxiCost,omitempty"`
}

func (s BicyclingDirectionNovaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionNovaResponseBodyData) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionNovaResponseBodyData) SetCount(v string) *BicyclingDirectionNovaResponseBodyData {
	s.Count = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyData) SetDestinationLatitude(v string) *BicyclingDirectionNovaResponseBodyData {
	s.DestinationLatitude = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyData) SetDestinationLongitude(v string) *BicyclingDirectionNovaResponseBodyData {
	s.DestinationLongitude = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyData) SetOriginLatitude(v string) *BicyclingDirectionNovaResponseBodyData {
	s.OriginLatitude = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyData) SetOriginLongitude(v string) *BicyclingDirectionNovaResponseBodyData {
	s.OriginLongitude = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyData) SetPaths(v []*BicyclingDirectionNovaResponseBodyDataPaths) *BicyclingDirectionNovaResponseBodyData {
	s.Paths = v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyData) SetTaxiCost(v string) *BicyclingDirectionNovaResponseBodyData {
	s.TaxiCost = &v
	return s
}

type BicyclingDirectionNovaResponseBodyDataPaths struct {
	Cost *BicyclingDirectionNovaResponseBodyDataPathsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 96375
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 300
	DurationSecond *string                                             `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	Restriction    *string                                             `json:"restriction,omitempty" xml:"restriction,omitempty"`
	Steps          []*BicyclingDirectionNovaResponseBodyDataPathsSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s BicyclingDirectionNovaResponseBodyDataPaths) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionNovaResponseBodyDataPaths) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionNovaResponseBodyDataPaths) SetCost(v *BicyclingDirectionNovaResponseBodyDataPathsCost) *BicyclingDirectionNovaResponseBodyDataPaths {
	s.Cost = v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPaths) SetDistanceMeter(v string) *BicyclingDirectionNovaResponseBodyDataPaths {
	s.DistanceMeter = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPaths) SetDurationSecond(v string) *BicyclingDirectionNovaResponseBodyDataPaths {
	s.DurationSecond = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPaths) SetRestriction(v string) *BicyclingDirectionNovaResponseBodyDataPaths {
	s.Restriction = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPaths) SetSteps(v []*BicyclingDirectionNovaResponseBodyDataPathsSteps) *BicyclingDirectionNovaResponseBodyDataPaths {
	s.Steps = v
	return s
}

type BicyclingDirectionNovaResponseBodyDataPathsCost struct {
	// example:
	//
	// 39233
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights     *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 4
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s BicyclingDirectionNovaResponseBodyDataPathsCost) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionNovaResponseBodyDataPathsCost) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsCost) SetDurationSecond(v string) *BicyclingDirectionNovaResponseBodyDataPathsCost {
	s.DurationSecond = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsCost) SetTaxiFee(v string) *BicyclingDirectionNovaResponseBodyDataPathsCost {
	s.TaxiFee = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsCost) SetTollDistanceMeter(v string) *BicyclingDirectionNovaResponseBodyDataPathsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsCost) SetTollRoads(v string) *BicyclingDirectionNovaResponseBodyDataPathsCost {
	s.TollRoads = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsCost) SetTolls(v string) *BicyclingDirectionNovaResponseBodyDataPathsCost {
	s.Tolls = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsCost) SetTrafficLights(v string) *BicyclingDirectionNovaResponseBodyDataPathsCost {
	s.TrafficLights = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsCost) SetTransitFee(v string) *BicyclingDirectionNovaResponseBodyDataPathsCost {
	s.TransitFee = &v
	return s
}

type BicyclingDirectionNovaResponseBodyDataPathsSteps struct {
	Cost        *BicyclingDirectionNovaResponseBodyDataPathsStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                               `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                               `json:"orientation,omitempty" xml:"orientation,omitempty"`
	Polyline    *string                                               `json:"polyline,omitempty" xml:"polyline,omitempty"`
	RoadName    *string                                               `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 3000
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s BicyclingDirectionNovaResponseBodyDataPathsSteps) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionNovaResponseBodyDataPathsSteps) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsSteps) SetCost(v *BicyclingDirectionNovaResponseBodyDataPathsStepsCost) *BicyclingDirectionNovaResponseBodyDataPathsSteps {
	s.Cost = v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsSteps) SetInstruction(v string) *BicyclingDirectionNovaResponseBodyDataPathsSteps {
	s.Instruction = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsSteps) SetOrientation(v string) *BicyclingDirectionNovaResponseBodyDataPathsSteps {
	s.Orientation = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsSteps) SetPolyline(v string) *BicyclingDirectionNovaResponseBodyDataPathsSteps {
	s.Polyline = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsSteps) SetRoadName(v string) *BicyclingDirectionNovaResponseBodyDataPathsSteps {
	s.RoadName = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsSteps) SetStepDistanceMeter(v string) *BicyclingDirectionNovaResponseBodyDataPathsSteps {
	s.StepDistanceMeter = &v
	return s
}

type BicyclingDirectionNovaResponseBodyDataPathsStepsCost struct {
	// example:
	//
	// 2000
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights     *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee        *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s BicyclingDirectionNovaResponseBodyDataPathsStepsCost) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionNovaResponseBodyDataPathsStepsCost) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsStepsCost) SetDurationSecond(v string) *BicyclingDirectionNovaResponseBodyDataPathsStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsStepsCost) SetTaxiFee(v string) *BicyclingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsStepsCost) SetTollDistanceMeter(v string) *BicyclingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsStepsCost) SetTollRoads(v string) *BicyclingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TollRoads = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsStepsCost) SetTolls(v string) *BicyclingDirectionNovaResponseBodyDataPathsStepsCost {
	s.Tolls = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsStepsCost) SetTrafficLights(v string) *BicyclingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *BicyclingDirectionNovaResponseBodyDataPathsStepsCost) SetTransitFee(v string) *BicyclingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TransitFee = &v
	return s
}

type BicyclingDirectionNovaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BicyclingDirectionNovaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BicyclingDirectionNovaResponse) String() string {
	return tea.Prettify(s)
}

func (s BicyclingDirectionNovaResponse) GoString() string {
	return s.String()
}

func (s *BicyclingDirectionNovaResponse) SetHeaders(v map[string]*string) *BicyclingDirectionNovaResponse {
	s.Headers = v
	return s
}

func (s *BicyclingDirectionNovaResponse) SetStatusCode(v int32) *BicyclingDirectionNovaResponse {
	s.StatusCode = &v
	return s
}

func (s *BicyclingDirectionNovaResponse) SetBody(v *BicyclingDirectionNovaResponseBody) *BicyclingDirectionNovaResponse {
	s.Body = v
	return s
}

type CircleTrafficStatusRequest struct {
	// example:
	//
	// 39.98641364
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// example:
	//
	// 116.3057764
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// example:
	//
	// 1000
	Radius *string `json:"radius,omitempty" xml:"radius,omitempty"`
	// example:
	//
	// HIGHWAY
	RoadLevel *string `json:"roadLevel,omitempty" xml:"roadLevel,omitempty"`
}

func (s CircleTrafficStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CircleTrafficStatusRequest) GoString() string {
	return s.String()
}

func (s *CircleTrafficStatusRequest) SetLatitude(v string) *CircleTrafficStatusRequest {
	s.Latitude = &v
	return s
}

func (s *CircleTrafficStatusRequest) SetLongitude(v string) *CircleTrafficStatusRequest {
	s.Longitude = &v
	return s
}

func (s *CircleTrafficStatusRequest) SetRadius(v string) *CircleTrafficStatusRequest {
	s.Radius = &v
	return s
}

func (s *CircleTrafficStatusRequest) SetRoadLevel(v string) *CircleTrafficStatusRequest {
	s.RoadLevel = &v
	return s
}

type CircleTrafficStatusResponseBody struct {
	Data []*CircleTrafficStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CircleTrafficStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CircleTrafficStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CircleTrafficStatusResponseBody) SetData(v []*CircleTrafficStatusResponseBodyData) *CircleTrafficStatusResponseBody {
	s.Data = v
	return s
}

func (s *CircleTrafficStatusResponseBody) SetErrorCode(v string) *CircleTrafficStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CircleTrafficStatusResponseBody) SetErrorMessage(v string) *CircleTrafficStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CircleTrafficStatusResponseBody) SetRequestId(v string) *CircleTrafficStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CircleTrafficStatusResponseBody) SetSuccess(v bool) *CircleTrafficStatusResponseBody {
	s.Success = &v
	return s
}

type CircleTrafficStatusResponseBodyData struct {
	Description *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	Evaluation  *CircleTrafficStatusResponseBodyDataEvaluation `json:"evaluation,omitempty" xml:"evaluation,omitempty" type:"Struct"`
}

func (s CircleTrafficStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CircleTrafficStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CircleTrafficStatusResponseBodyData) SetDescription(v string) *CircleTrafficStatusResponseBodyData {
	s.Description = &v
	return s
}

func (s *CircleTrafficStatusResponseBodyData) SetEvaluation(v *CircleTrafficStatusResponseBodyDataEvaluation) *CircleTrafficStatusResponseBodyData {
	s.Evaluation = v
	return s
}

type CircleTrafficStatusResponseBodyDataEvaluation struct {
	// example:
	//
	// 0.00%
	BlockedPercentage *string `json:"blockedPercentage,omitempty" xml:"blockedPercentage,omitempty"`
	// example:
	//
	// 54.55%
	ClearPercentage *string `json:"clearPercentage,omitempty" xml:"clearPercentage,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 18.18%
	HeavyPercentage *string `json:"heavyPercentage,omitempty" xml:"heavyPercentage,omitempty"`
	// example:
	//
	// RUNNING_EXECUTION
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 27.27%
	UnknownPercentage *string `json:"unknownPercentage,omitempty" xml:"unknownPercentage,omitempty"`
}

func (s CircleTrafficStatusResponseBodyDataEvaluation) String() string {
	return tea.Prettify(s)
}

func (s CircleTrafficStatusResponseBodyDataEvaluation) GoString() string {
	return s.String()
}

func (s *CircleTrafficStatusResponseBodyDataEvaluation) SetBlockedPercentage(v string) *CircleTrafficStatusResponseBodyDataEvaluation {
	s.BlockedPercentage = &v
	return s
}

func (s *CircleTrafficStatusResponseBodyDataEvaluation) SetClearPercentage(v string) *CircleTrafficStatusResponseBodyDataEvaluation {
	s.ClearPercentage = &v
	return s
}

func (s *CircleTrafficStatusResponseBodyDataEvaluation) SetDescription(v string) *CircleTrafficStatusResponseBodyDataEvaluation {
	s.Description = &v
	return s
}

func (s *CircleTrafficStatusResponseBodyDataEvaluation) SetHeavyPercentage(v string) *CircleTrafficStatusResponseBodyDataEvaluation {
	s.HeavyPercentage = &v
	return s
}

func (s *CircleTrafficStatusResponseBodyDataEvaluation) SetStatus(v string) *CircleTrafficStatusResponseBodyDataEvaluation {
	s.Status = &v
	return s
}

func (s *CircleTrafficStatusResponseBodyDataEvaluation) SetUnknownPercentage(v string) *CircleTrafficStatusResponseBodyDataEvaluation {
	s.UnknownPercentage = &v
	return s
}

type CircleTrafficStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CircleTrafficStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CircleTrafficStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CircleTrafficStatusResponse) GoString() string {
	return s.String()
}

func (s *CircleTrafficStatusResponse) SetHeaders(v map[string]*string) *CircleTrafficStatusResponse {
	s.Headers = v
	return s
}

func (s *CircleTrafficStatusResponse) SetStatusCode(v int32) *CircleTrafficStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CircleTrafficStatusResponse) SetBody(v *CircleTrafficStatusResponseBody) *CircleTrafficStatusResponse {
	s.Body = v
	return s
}

type CommonQueryBySceneRequest struct {
	Body *CommonAgentQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommonQueryBySceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CommonQueryBySceneRequest) GoString() string {
	return s.String()
}

func (s *CommonQueryBySceneRequest) SetBody(v *CommonAgentQuery) *CommonQueryBySceneRequest {
	s.Body = v
	return s
}

type CommonQueryBySceneResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResult       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommonQueryBySceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CommonQueryBySceneResponse) GoString() string {
	return s.String()
}

func (s *CommonQueryBySceneResponse) SetHeaders(v map[string]*string) *CommonQueryBySceneResponse {
	s.Headers = v
	return s
}

func (s *CommonQueryBySceneResponse) SetStatusCode(v int32) *CommonQueryBySceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CommonQueryBySceneResponse) SetBody(v *QueryResult) *CommonQueryBySceneResponse {
	s.Body = v
	return s
}

type DrivingDirectionRequest struct {
	// example:
	//
	// 39.896463
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s DrivingDirectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionRequest) GoString() string {
	return s.String()
}

func (s *DrivingDirectionRequest) SetDestinationLatitude(v string) *DrivingDirectionRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *DrivingDirectionRequest) SetDestinationLongitude(v string) *DrivingDirectionRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *DrivingDirectionRequest) SetOriginLatitude(v string) *DrivingDirectionRequest {
	s.OriginLatitude = &v
	return s
}

func (s *DrivingDirectionRequest) SetOriginLongitude(v string) *DrivingDirectionRequest {
	s.OriginLongitude = &v
	return s
}

type DrivingDirectionResponseBody struct {
	Data []*DrivingDirectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// <title>502 Bad Gateway</title>
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DrivingDirectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionResponseBody) GoString() string {
	return s.String()
}

func (s *DrivingDirectionResponseBody) SetData(v []*DrivingDirectionResponseBodyData) *DrivingDirectionResponseBody {
	s.Data = v
	return s
}

func (s *DrivingDirectionResponseBody) SetErrorCode(v string) *DrivingDirectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DrivingDirectionResponseBody) SetErrorMessage(v string) *DrivingDirectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DrivingDirectionResponseBody) SetRequestId(v string) *DrivingDirectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DrivingDirectionResponseBody) SetSuccess(v bool) *DrivingDirectionResponseBody {
	s.Success = &v
	return s
}

type DrivingDirectionResponseBodyData struct {
	Cost *DrivingDirectionResponseBodyDataCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 445
	DistanceMeter *string                                  `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	Restriction   *string                                  `json:"restriction,omitempty" xml:"restriction,omitempty"`
	Steps         []*DrivingDirectionResponseBodyDataSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s DrivingDirectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DrivingDirectionResponseBodyData) SetCost(v *DrivingDirectionResponseBodyDataCost) *DrivingDirectionResponseBodyData {
	s.Cost = v
	return s
}

func (s *DrivingDirectionResponseBodyData) SetDistanceMeter(v string) *DrivingDirectionResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *DrivingDirectionResponseBodyData) SetRestriction(v string) *DrivingDirectionResponseBodyData {
	s.Restriction = &v
	return s
}

func (s *DrivingDirectionResponseBodyData) SetSteps(v []*DrivingDirectionResponseBodyDataSteps) *DrivingDirectionResponseBodyData {
	s.Steps = v
	return s
}

type DrivingDirectionResponseBodyDataCost struct {
	// example:
	//
	// 1231
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 6
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	// example:
	//
	// 23
	Tolls         *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s DrivingDirectionResponseBodyDataCost) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionResponseBodyDataCost) GoString() string {
	return s.String()
}

func (s *DrivingDirectionResponseBodyDataCost) SetDurationSecond(v string) *DrivingDirectionResponseBodyDataCost {
	s.DurationSecond = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataCost) SetTaxiFee(v string) *DrivingDirectionResponseBodyDataCost {
	s.TaxiFee = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataCost) SetTollDistanceMeter(v string) *DrivingDirectionResponseBodyDataCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataCost) SetTollRoads(v string) *DrivingDirectionResponseBodyDataCost {
	s.TollRoads = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataCost) SetTolls(v string) *DrivingDirectionResponseBodyDataCost {
	s.Tolls = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataCost) SetTrafficLights(v string) *DrivingDirectionResponseBodyDataCost {
	s.TrafficLights = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataCost) SetTransitFee(v string) *DrivingDirectionResponseBodyDataCost {
	s.TransitFee = &v
	return s
}

type DrivingDirectionResponseBodyDataSteps struct {
	Cost        *DrivingDirectionResponseBodyDataStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                    `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                    `json:"orientation,omitempty" xml:"orientation,omitempty"`
	RoadName    *string                                    `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 500
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s DrivingDirectionResponseBodyDataSteps) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionResponseBodyDataSteps) GoString() string {
	return s.String()
}

func (s *DrivingDirectionResponseBodyDataSteps) SetCost(v *DrivingDirectionResponseBodyDataStepsCost) *DrivingDirectionResponseBodyDataSteps {
	s.Cost = v
	return s
}

func (s *DrivingDirectionResponseBodyDataSteps) SetInstruction(v string) *DrivingDirectionResponseBodyDataSteps {
	s.Instruction = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataSteps) SetOrientation(v string) *DrivingDirectionResponseBodyDataSteps {
	s.Orientation = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataSteps) SetRoadName(v string) *DrivingDirectionResponseBodyDataSteps {
	s.RoadName = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataSteps) SetStepDistanceMeter(v string) *DrivingDirectionResponseBodyDataSteps {
	s.StepDistanceMeter = &v
	return s
}

type DrivingDirectionResponseBodyDataStepsCost struct {
	// example:
	//
	// 27647
	DurationSecond    *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	// example:
	//
	// xxx
	TollRoads *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls     *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 5
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s DrivingDirectionResponseBodyDataStepsCost) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionResponseBodyDataStepsCost) GoString() string {
	return s.String()
}

func (s *DrivingDirectionResponseBodyDataStepsCost) SetDurationSecond(v string) *DrivingDirectionResponseBodyDataStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataStepsCost) SetTaxiFee(v string) *DrivingDirectionResponseBodyDataStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataStepsCost) SetTollDistanceMeter(v string) *DrivingDirectionResponseBodyDataStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataStepsCost) SetTollRoads(v string) *DrivingDirectionResponseBodyDataStepsCost {
	s.TollRoads = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataStepsCost) SetTolls(v string) *DrivingDirectionResponseBodyDataStepsCost {
	s.Tolls = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataStepsCost) SetTrafficLights(v string) *DrivingDirectionResponseBodyDataStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *DrivingDirectionResponseBodyDataStepsCost) SetTransitFee(v string) *DrivingDirectionResponseBodyDataStepsCost {
	s.TransitFee = &v
	return s
}

type DrivingDirectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DrivingDirectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DrivingDirectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionResponse) GoString() string {
	return s.String()
}

func (s *DrivingDirectionResponse) SetHeaders(v map[string]*string) *DrivingDirectionResponse {
	s.Headers = v
	return s
}

func (s *DrivingDirectionResponse) SetStatusCode(v int32) *DrivingDirectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DrivingDirectionResponse) SetBody(v *DrivingDirectionResponseBody) *DrivingDirectionResponse {
	s.Body = v
	return s
}

type DrivingDirectionNovaRequest struct {
	CarType *string `json:"carType,omitempty" xml:"carType,omitempty"`
	// example:
	//
	// 43.345456
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
	Plate           *string `json:"plate,omitempty" xml:"plate,omitempty"`
}

func (s DrivingDirectionNovaRequest) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaRequest) GoString() string {
	return s.String()
}

func (s *DrivingDirectionNovaRequest) SetCarType(v string) *DrivingDirectionNovaRequest {
	s.CarType = &v
	return s
}

func (s *DrivingDirectionNovaRequest) SetDestinationLatitude(v string) *DrivingDirectionNovaRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *DrivingDirectionNovaRequest) SetDestinationLongitude(v string) *DrivingDirectionNovaRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *DrivingDirectionNovaRequest) SetOriginLatitude(v string) *DrivingDirectionNovaRequest {
	s.OriginLatitude = &v
	return s
}

func (s *DrivingDirectionNovaRequest) SetOriginLongitude(v string) *DrivingDirectionNovaRequest {
	s.OriginLongitude = &v
	return s
}

func (s *DrivingDirectionNovaRequest) SetPlate(v string) *DrivingDirectionNovaRequest {
	s.Plate = &v
	return s
}

type DrivingDirectionNovaResponseBody struct {
	Data *DrivingDirectionNovaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied, message: No such namespace namespaces/general-perf-cn-shenzhen-e-default.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DrivingDirectionNovaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaResponseBody) GoString() string {
	return s.String()
}

func (s *DrivingDirectionNovaResponseBody) SetData(v *DrivingDirectionNovaResponseBodyData) *DrivingDirectionNovaResponseBody {
	s.Data = v
	return s
}

func (s *DrivingDirectionNovaResponseBody) SetErrorCode(v int32) *DrivingDirectionNovaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DrivingDirectionNovaResponseBody) SetErrorMessage(v string) *DrivingDirectionNovaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DrivingDirectionNovaResponseBody) SetSuccess(v bool) *DrivingDirectionNovaResponseBody {
	s.Success = &v
	return s
}

type DrivingDirectionNovaResponseBodyData struct {
	Count *string `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 40.345456
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 117.466485
	OriginLongitude *string                                      `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
	Paths           []*DrivingDirectionNovaResponseBodyDataPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TaxiCost *string `json:"taxiCost,omitempty" xml:"taxiCost,omitempty"`
}

func (s DrivingDirectionNovaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaResponseBodyData) GoString() string {
	return s.String()
}

func (s *DrivingDirectionNovaResponseBodyData) SetCount(v string) *DrivingDirectionNovaResponseBodyData {
	s.Count = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyData) SetDestinationLatitude(v string) *DrivingDirectionNovaResponseBodyData {
	s.DestinationLatitude = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyData) SetDestinationLongitude(v string) *DrivingDirectionNovaResponseBodyData {
	s.DestinationLongitude = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyData) SetOriginLatitude(v string) *DrivingDirectionNovaResponseBodyData {
	s.OriginLatitude = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyData) SetOriginLongitude(v string) *DrivingDirectionNovaResponseBodyData {
	s.OriginLongitude = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyData) SetPaths(v []*DrivingDirectionNovaResponseBodyDataPaths) *DrivingDirectionNovaResponseBodyData {
	s.Paths = v
	return s
}

func (s *DrivingDirectionNovaResponseBodyData) SetTaxiCost(v string) *DrivingDirectionNovaResponseBodyData {
	s.TaxiCost = &v
	return s
}

type DrivingDirectionNovaResponseBodyDataPaths struct {
	Cost *DrivingDirectionNovaResponseBodyDataPathsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 96375
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 39223
	DurationSecond *string                                           `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	Restriction    *string                                           `json:"restriction,omitempty" xml:"restriction,omitempty"`
	Steps          []*DrivingDirectionNovaResponseBodyDataPathsSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s DrivingDirectionNovaResponseBodyDataPaths) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaResponseBodyDataPaths) GoString() string {
	return s.String()
}

func (s *DrivingDirectionNovaResponseBodyDataPaths) SetCost(v *DrivingDirectionNovaResponseBodyDataPathsCost) *DrivingDirectionNovaResponseBodyDataPaths {
	s.Cost = v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPaths) SetDistanceMeter(v string) *DrivingDirectionNovaResponseBodyDataPaths {
	s.DistanceMeter = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPaths) SetDurationSecond(v string) *DrivingDirectionNovaResponseBodyDataPaths {
	s.DurationSecond = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPaths) SetRestriction(v string) *DrivingDirectionNovaResponseBodyDataPaths {
	s.Restriction = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPaths) SetSteps(v []*DrivingDirectionNovaResponseBodyDataPathsSteps) *DrivingDirectionNovaResponseBodyDataPaths {
	s.Steps = v
	return s
}

type DrivingDirectionNovaResponseBodyDataPathsCost struct {
	// example:
	//
	// 39233
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 3
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 10
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s DrivingDirectionNovaResponseBodyDataPathsCost) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaResponseBodyDataPathsCost) GoString() string {
	return s.String()
}

func (s *DrivingDirectionNovaResponseBodyDataPathsCost) SetDurationSecond(v string) *DrivingDirectionNovaResponseBodyDataPathsCost {
	s.DurationSecond = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsCost) SetTaxiFee(v string) *DrivingDirectionNovaResponseBodyDataPathsCost {
	s.TaxiFee = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsCost) SetTollDistanceMeter(v string) *DrivingDirectionNovaResponseBodyDataPathsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsCost) SetTollRoads(v string) *DrivingDirectionNovaResponseBodyDataPathsCost {
	s.TollRoads = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsCost) SetTolls(v string) *DrivingDirectionNovaResponseBodyDataPathsCost {
	s.Tolls = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsCost) SetTrafficLights(v string) *DrivingDirectionNovaResponseBodyDataPathsCost {
	s.TrafficLights = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsCost) SetTransitFee(v string) *DrivingDirectionNovaResponseBodyDataPathsCost {
	s.TransitFee = &v
	return s
}

type DrivingDirectionNovaResponseBodyDataPathsSteps struct {
	Cost        *DrivingDirectionNovaResponseBodyDataPathsStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                             `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                             `json:"orientation,omitempty" xml:"orientation,omitempty"`
	Polyline    *string                                             `json:"polyline,omitempty" xml:"polyline,omitempty"`
	RoadName    *string                                             `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 3000
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s DrivingDirectionNovaResponseBodyDataPathsSteps) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaResponseBodyDataPathsSteps) GoString() string {
	return s.String()
}

func (s *DrivingDirectionNovaResponseBodyDataPathsSteps) SetCost(v *DrivingDirectionNovaResponseBodyDataPathsStepsCost) *DrivingDirectionNovaResponseBodyDataPathsSteps {
	s.Cost = v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsSteps) SetInstruction(v string) *DrivingDirectionNovaResponseBodyDataPathsSteps {
	s.Instruction = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsSteps) SetOrientation(v string) *DrivingDirectionNovaResponseBodyDataPathsSteps {
	s.Orientation = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsSteps) SetPolyline(v string) *DrivingDirectionNovaResponseBodyDataPathsSteps {
	s.Polyline = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsSteps) SetRoadName(v string) *DrivingDirectionNovaResponseBodyDataPathsSteps {
	s.RoadName = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsSteps) SetStepDistanceMeter(v string) *DrivingDirectionNovaResponseBodyDataPathsSteps {
	s.StepDistanceMeter = &v
	return s
}

type DrivingDirectionNovaResponseBodyDataPathsStepsCost struct {
	// example:
	//
	// 1234
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights     *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee        *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s DrivingDirectionNovaResponseBodyDataPathsStepsCost) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaResponseBodyDataPathsStepsCost) GoString() string {
	return s.String()
}

func (s *DrivingDirectionNovaResponseBodyDataPathsStepsCost) SetDurationSecond(v string) *DrivingDirectionNovaResponseBodyDataPathsStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsStepsCost) SetTaxiFee(v string) *DrivingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsStepsCost) SetTollDistanceMeter(v string) *DrivingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsStepsCost) SetTollRoads(v string) *DrivingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TollRoads = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsStepsCost) SetTolls(v string) *DrivingDirectionNovaResponseBodyDataPathsStepsCost {
	s.Tolls = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsStepsCost) SetTrafficLights(v string) *DrivingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *DrivingDirectionNovaResponseBodyDataPathsStepsCost) SetTransitFee(v string) *DrivingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TransitFee = &v
	return s
}

type DrivingDirectionNovaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DrivingDirectionNovaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DrivingDirectionNovaResponse) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaResponse) GoString() string {
	return s.String()
}

func (s *DrivingDirectionNovaResponse) SetHeaders(v map[string]*string) *DrivingDirectionNovaResponse {
	s.Headers = v
	return s
}

func (s *DrivingDirectionNovaResponse) SetStatusCode(v int32) *DrivingDirectionNovaResponse {
	s.StatusCode = &v
	return s
}

func (s *DrivingDirectionNovaResponse) SetBody(v *DrivingDirectionNovaResponseBody) *DrivingDirectionNovaResponse {
	s.Body = v
	return s
}

type ElectrobikeDirectionRequest struct {
	// example:
	//
	// 39.896463
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s ElectrobikeDirectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionRequest) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionRequest) SetDestinationLatitude(v string) *ElectrobikeDirectionRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *ElectrobikeDirectionRequest) SetDestinationLongitude(v string) *ElectrobikeDirectionRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *ElectrobikeDirectionRequest) SetOriginLatitude(v string) *ElectrobikeDirectionRequest {
	s.OriginLatitude = &v
	return s
}

func (s *ElectrobikeDirectionRequest) SetOriginLongitude(v string) *ElectrobikeDirectionRequest {
	s.OriginLongitude = &v
	return s
}

type ElectrobikeDirectionResponseBody struct {
	Data []*ElectrobikeDirectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied, message: Unauthorized.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ElectrobikeDirectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionResponseBody) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionResponseBody) SetData(v []*ElectrobikeDirectionResponseBodyData) *ElectrobikeDirectionResponseBody {
	s.Data = v
	return s
}

func (s *ElectrobikeDirectionResponseBody) SetErrorCode(v string) *ElectrobikeDirectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ElectrobikeDirectionResponseBody) SetErrorMessage(v string) *ElectrobikeDirectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ElectrobikeDirectionResponseBody) SetRequestId(v string) *ElectrobikeDirectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ElectrobikeDirectionResponseBody) SetSuccess(v bool) *ElectrobikeDirectionResponseBody {
	s.Success = &v
	return s
}

type ElectrobikeDirectionResponseBodyData struct {
	// example:
	//
	// 445
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 2345
	DurationSecond *string                                      `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	Steps          []*ElectrobikeDirectionResponseBodyDataSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s ElectrobikeDirectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionResponseBodyData) SetDistanceMeter(v string) *ElectrobikeDirectionResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyData) SetDurationSecond(v string) *ElectrobikeDirectionResponseBodyData {
	s.DurationSecond = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyData) SetSteps(v []*ElectrobikeDirectionResponseBodyDataSteps) *ElectrobikeDirectionResponseBodyData {
	s.Steps = v
	return s
}

type ElectrobikeDirectionResponseBodyDataSteps struct {
	Cost        *ElectrobikeDirectionResponseBodyDataStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                        `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                        `json:"orientation,omitempty" xml:"orientation,omitempty"`
	RoadName    *string                                        `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 500
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s ElectrobikeDirectionResponseBodyDataSteps) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionResponseBodyDataSteps) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionResponseBodyDataSteps) SetCost(v *ElectrobikeDirectionResponseBodyDataStepsCost) *ElectrobikeDirectionResponseBodyDataSteps {
	s.Cost = v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataSteps) SetInstruction(v string) *ElectrobikeDirectionResponseBodyDataSteps {
	s.Instruction = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataSteps) SetOrientation(v string) *ElectrobikeDirectionResponseBodyDataSteps {
	s.Orientation = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataSteps) SetRoadName(v string) *ElectrobikeDirectionResponseBodyDataSteps {
	s.RoadName = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataSteps) SetStepDistanceMeter(v string) *ElectrobikeDirectionResponseBodyDataSteps {
	s.StepDistanceMeter = &v
	return s
}

type ElectrobikeDirectionResponseBodyDataStepsCost struct {
	// example:
	//
	// 2002
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	TaxiFee        *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	// example:
	//
	// 1000
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	// example:
	//
	// xxx
	TollRoads *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls     *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 5
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s ElectrobikeDirectionResponseBodyDataStepsCost) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionResponseBodyDataStepsCost) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionResponseBodyDataStepsCost) SetDurationSecond(v string) *ElectrobikeDirectionResponseBodyDataStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataStepsCost) SetTaxiFee(v string) *ElectrobikeDirectionResponseBodyDataStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataStepsCost) SetTollDistanceMeter(v string) *ElectrobikeDirectionResponseBodyDataStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataStepsCost) SetTollRoads(v string) *ElectrobikeDirectionResponseBodyDataStepsCost {
	s.TollRoads = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataStepsCost) SetTolls(v string) *ElectrobikeDirectionResponseBodyDataStepsCost {
	s.Tolls = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataStepsCost) SetTrafficLights(v string) *ElectrobikeDirectionResponseBodyDataStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *ElectrobikeDirectionResponseBodyDataStepsCost) SetTransitFee(v string) *ElectrobikeDirectionResponseBodyDataStepsCost {
	s.TransitFee = &v
	return s
}

type ElectrobikeDirectionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ElectrobikeDirectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ElectrobikeDirectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionResponse) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionResponse) SetHeaders(v map[string]*string) *ElectrobikeDirectionResponse {
	s.Headers = v
	return s
}

func (s *ElectrobikeDirectionResponse) SetStatusCode(v int32) *ElectrobikeDirectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ElectrobikeDirectionResponse) SetBody(v *ElectrobikeDirectionResponseBody) *ElectrobikeDirectionResponse {
	s.Body = v
	return s
}

type ElectrobikeDirectionNovaRequest struct {
	// example:
	//
	// 40.234564
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.345456
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s ElectrobikeDirectionNovaRequest) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionNovaRequest) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionNovaRequest) SetDestinationLatitude(v string) *ElectrobikeDirectionNovaRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *ElectrobikeDirectionNovaRequest) SetDestinationLongitude(v string) *ElectrobikeDirectionNovaRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *ElectrobikeDirectionNovaRequest) SetOriginLatitude(v string) *ElectrobikeDirectionNovaRequest {
	s.OriginLatitude = &v
	return s
}

func (s *ElectrobikeDirectionNovaRequest) SetOriginLongitude(v string) *ElectrobikeDirectionNovaRequest {
	s.OriginLongitude = &v
	return s
}

type ElectrobikeDirectionNovaResponseBody struct {
	Data *ElectrobikeDirectionNovaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// <title>502 Bad Gateway</title>
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ElectrobikeDirectionNovaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionNovaResponseBody) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionNovaResponseBody) SetData(v *ElectrobikeDirectionNovaResponseBodyData) *ElectrobikeDirectionNovaResponseBody {
	s.Data = v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBody) SetErrorCode(v int32) *ElectrobikeDirectionNovaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBody) SetErrorMessage(v string) *ElectrobikeDirectionNovaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBody) SetSuccess(v bool) *ElectrobikeDirectionNovaResponseBody {
	s.Success = &v
	return s
}

type ElectrobikeDirectionNovaResponseBodyData struct {
	Count *string `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 40.345456
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.345456
	OriginLongitude *string                                          `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
	Paths           []*ElectrobikeDirectionNovaResponseBodyDataPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TaxiCost *string `json:"taxiCost,omitempty" xml:"taxiCost,omitempty"`
}

func (s ElectrobikeDirectionNovaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionNovaResponseBodyData) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionNovaResponseBodyData) SetCount(v string) *ElectrobikeDirectionNovaResponseBodyData {
	s.Count = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyData) SetDestinationLatitude(v string) *ElectrobikeDirectionNovaResponseBodyData {
	s.DestinationLatitude = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyData) SetDestinationLongitude(v string) *ElectrobikeDirectionNovaResponseBodyData {
	s.DestinationLongitude = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyData) SetOriginLatitude(v string) *ElectrobikeDirectionNovaResponseBodyData {
	s.OriginLatitude = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyData) SetOriginLongitude(v string) *ElectrobikeDirectionNovaResponseBodyData {
	s.OriginLongitude = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyData) SetPaths(v []*ElectrobikeDirectionNovaResponseBodyDataPaths) *ElectrobikeDirectionNovaResponseBodyData {
	s.Paths = v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyData) SetTaxiCost(v string) *ElectrobikeDirectionNovaResponseBodyData {
	s.TaxiCost = &v
	return s
}

type ElectrobikeDirectionNovaResponseBodyDataPaths struct {
	Cost *ElectrobikeDirectionNovaResponseBodyDataPathsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 12000
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 300
	DurationSecond *string                                               `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	Restriction    *string                                               `json:"restriction,omitempty" xml:"restriction,omitempty"`
	Steps          []*ElectrobikeDirectionNovaResponseBodyDataPathsSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s ElectrobikeDirectionNovaResponseBodyDataPaths) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionNovaResponseBodyDataPaths) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPaths) SetCost(v *ElectrobikeDirectionNovaResponseBodyDataPathsCost) *ElectrobikeDirectionNovaResponseBodyDataPaths {
	s.Cost = v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPaths) SetDistanceMeter(v string) *ElectrobikeDirectionNovaResponseBodyDataPaths {
	s.DistanceMeter = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPaths) SetDurationSecond(v string) *ElectrobikeDirectionNovaResponseBodyDataPaths {
	s.DurationSecond = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPaths) SetRestriction(v string) *ElectrobikeDirectionNovaResponseBodyDataPaths {
	s.Restriction = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPaths) SetSteps(v []*ElectrobikeDirectionNovaResponseBodyDataPathsSteps) *ElectrobikeDirectionNovaResponseBodyDataPaths {
	s.Steps = v
	return s
}

type ElectrobikeDirectionNovaResponseBodyDataPathsCost struct {
	// example:
	//
	// 500
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 4
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 4
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s ElectrobikeDirectionNovaResponseBodyDataPathsCost) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionNovaResponseBodyDataPathsCost) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsCost) SetDurationSecond(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsCost {
	s.DurationSecond = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsCost) SetTaxiFee(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsCost {
	s.TaxiFee = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsCost) SetTollDistanceMeter(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsCost) SetTollRoads(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsCost {
	s.TollRoads = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsCost) SetTolls(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsCost {
	s.Tolls = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsCost) SetTrafficLights(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsCost {
	s.TrafficLights = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsCost) SetTransitFee(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsCost {
	s.TransitFee = &v
	return s
}

type ElectrobikeDirectionNovaResponseBodyDataPathsSteps struct {
	Cost        *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                                 `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                                 `json:"orientation,omitempty" xml:"orientation,omitempty"`
	Polyline    *string                                                 `json:"polyline,omitempty" xml:"polyline,omitempty"`
	RoadName    *string                                                 `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 3000
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s ElectrobikeDirectionNovaResponseBodyDataPathsSteps) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionNovaResponseBodyDataPathsSteps) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsSteps) SetCost(v *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) *ElectrobikeDirectionNovaResponseBodyDataPathsSteps {
	s.Cost = v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsSteps) SetInstruction(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsSteps {
	s.Instruction = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsSteps) SetOrientation(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsSteps {
	s.Orientation = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsSteps) SetPolyline(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsSteps {
	s.Polyline = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsSteps) SetRoadName(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsSteps {
	s.RoadName = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsSteps) SetStepDistanceMeter(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsSteps {
	s.StepDistanceMeter = &v
	return s
}

type ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost struct {
	// example:
	//
	// 2000
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 5
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) SetDurationSecond(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) SetTaxiFee(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) SetTollDistanceMeter(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) SetTollRoads(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost {
	s.TollRoads = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) SetTolls(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost {
	s.Tolls = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) SetTrafficLights(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost) SetTransitFee(v string) *ElectrobikeDirectionNovaResponseBodyDataPathsStepsCost {
	s.TransitFee = &v
	return s
}

type ElectrobikeDirectionNovaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ElectrobikeDirectionNovaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ElectrobikeDirectionNovaResponse) String() string {
	return tea.Prettify(s)
}

func (s ElectrobikeDirectionNovaResponse) GoString() string {
	return s.String()
}

func (s *ElectrobikeDirectionNovaResponse) SetHeaders(v map[string]*string) *ElectrobikeDirectionNovaResponse {
	s.Headers = v
	return s
}

func (s *ElectrobikeDirectionNovaResponse) SetStatusCode(v int32) *ElectrobikeDirectionNovaResponse {
	s.StatusCode = &v
	return s
}

func (s *ElectrobikeDirectionNovaResponse) SetBody(v *ElectrobikeDirectionNovaResponseBody) *ElectrobikeDirectionNovaResponse {
	s.Body = v
	return s
}

type GeoCodeRequest struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	City    *string `json:"city,omitempty" xml:"city,omitempty"`
}

func (s GeoCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GeoCodeRequest) GoString() string {
	return s.String()
}

func (s *GeoCodeRequest) SetAddress(v string) *GeoCodeRequest {
	s.Address = &v
	return s
}

func (s *GeoCodeRequest) SetCity(v string) *GeoCodeRequest {
	s.City = &v
	return s
}

type GeoCodeResponseBody struct {
	Data []*GeoCodeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check log.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GeoCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GeoCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GeoCodeResponseBody) SetData(v []*GeoCodeResponseBodyData) *GeoCodeResponseBody {
	s.Data = v
	return s
}

func (s *GeoCodeResponseBody) SetErrorCode(v int32) *GeoCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GeoCodeResponseBody) SetErrorMessage(v string) *GeoCodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GeoCodeResponseBody) SetSuccess(v bool) *GeoCodeResponseBody {
	s.Success = &v
	return s
}

type GeoCodeResponseBodyData struct {
	Building     *GeoCodeResponseBodyDataBuilding `json:"building,omitempty" xml:"building,omitempty" type:"Struct"`
	City         *string                          `json:"city,omitempty" xml:"city,omitempty"`
	CityCode     *string                          `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	District     *string                          `json:"district,omitempty" xml:"district,omitempty"`
	DistrictCode *string                          `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	Latitude     *string                          `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Level        *string                          `json:"level,omitempty" xml:"level,omitempty"`
	Longitude    *string                          `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// example:
	//
	// 12201281024024
	Number   *string `json:"number,omitempty" xml:"number,omitempty"`
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	Street   *string `json:"street,omitempty" xml:"street,omitempty"`
}

func (s GeoCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GeoCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GeoCodeResponseBodyData) SetBuilding(v *GeoCodeResponseBodyDataBuilding) *GeoCodeResponseBodyData {
	s.Building = v
	return s
}

func (s *GeoCodeResponseBodyData) SetCity(v string) *GeoCodeResponseBodyData {
	s.City = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetCityCode(v string) *GeoCodeResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetDistrict(v string) *GeoCodeResponseBodyData {
	s.District = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetDistrictCode(v string) *GeoCodeResponseBodyData {
	s.DistrictCode = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetLatitude(v string) *GeoCodeResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetLevel(v string) *GeoCodeResponseBodyData {
	s.Level = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetLongitude(v string) *GeoCodeResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetNumber(v string) *GeoCodeResponseBodyData {
	s.Number = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetProvince(v string) *GeoCodeResponseBodyData {
	s.Province = &v
	return s
}

func (s *GeoCodeResponseBodyData) SetStreet(v string) *GeoCodeResponseBodyData {
	s.Street = &v
	return s
}

type GeoCodeResponseBodyDataBuilding struct {
	// example:
	//
	// timeliness_ms
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// internal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GeoCodeResponseBodyDataBuilding) String() string {
	return tea.Prettify(s)
}

func (s GeoCodeResponseBodyDataBuilding) GoString() string {
	return s.String()
}

func (s *GeoCodeResponseBodyDataBuilding) SetName(v string) *GeoCodeResponseBodyDataBuilding {
	s.Name = &v
	return s
}

func (s *GeoCodeResponseBodyDataBuilding) SetType(v string) *GeoCodeResponseBodyDataBuilding {
	s.Type = &v
	return s
}

type GeoCodeResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GeoCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GeoCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GeoCodeResponse) GoString() string {
	return s.String()
}

func (s *GeoCodeResponse) SetHeaders(v map[string]*string) *GeoCodeResponse {
	s.Headers = v
	return s
}

func (s *GeoCodeResponse) SetStatusCode(v int32) *GeoCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GeoCodeResponse) SetBody(v *GeoCodeResponseBody) *GeoCodeResponse {
	s.Body = v
	return s
}

type NearbySearchRequest struct {
	// 搜索的关键词
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// 搜索范围中心的纬度坐标。小数精度均不得超过六位
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 搜索范围中心的经度坐标。小数精度均不得超过六位
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 搜索的页数
	//
	// example:
	//
	// {\\"total_count\\": 6851, \\"page_number\\": 54, \\"page_size\\": 100}
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// 搜索的范围，以米为单位。
	Radius *int64 `json:"radius,omitempty" xml:"radius,omitempty"`
	// 搜索结果每页所包含的结果数量
	//
	// example:
	//
	// 812788
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// 目标类型的枚举值，以 `|` 标记分隔
	//
	// example:
	//
	// [\\"synonym\\",\\"stopword\\",\\"correction\\",\\"ner\\",\\"term_weighting\\",\\"category_prediction\\"]
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s NearbySearchRequest) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchRequest) GoString() string {
	return s.String()
}

func (s *NearbySearchRequest) SetKeywords(v string) *NearbySearchRequest {
	s.Keywords = &v
	return s
}

func (s *NearbySearchRequest) SetLatitude(v string) *NearbySearchRequest {
	s.Latitude = &v
	return s
}

func (s *NearbySearchRequest) SetLongitude(v string) *NearbySearchRequest {
	s.Longitude = &v
	return s
}

func (s *NearbySearchRequest) SetPage(v int32) *NearbySearchRequest {
	s.Page = &v
	return s
}

func (s *NearbySearchRequest) SetRadius(v int64) *NearbySearchRequest {
	s.Radius = &v
	return s
}

func (s *NearbySearchRequest) SetSize(v int32) *NearbySearchRequest {
	s.Size = &v
	return s
}

func (s *NearbySearchRequest) SetTypes(v string) *NearbySearchRequest {
	s.Types = &v
	return s
}

type NearbySearchResponseBody struct {
	Data []*NearbySearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied, message: No such namespace namespaces/general-perf-cn-shenzhen-e-default.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s NearbySearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchResponseBody) GoString() string {
	return s.String()
}

func (s *NearbySearchResponseBody) SetData(v []*NearbySearchResponseBodyData) *NearbySearchResponseBody {
	s.Data = v
	return s
}

func (s *NearbySearchResponseBody) SetErrorCode(v string) *NearbySearchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NearbySearchResponseBody) SetErrorMessage(v string) *NearbySearchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NearbySearchResponseBody) SetSuccess(v bool) *NearbySearchResponseBody {
	s.Success = &v
	return s
}

type NearbySearchResponseBodyData struct {
	Address       *string `json:"address,omitempty" xml:"address,omitempty"`
	CityCode      *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	CityName      *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	DistrictCode  *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	DistrictName  *string `json:"districtName,omitempty" xml:"districtName,omitempty"`
	// example:
	//
	// 38865
	Id        *string                               `json:"id,omitempty" xml:"id,omitempty"`
	Images    []*NearbySearchResponseBodyDataImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Latitude  *string                               `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *string                               `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Metadata  *NearbySearchResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// example:
	//
	// hydro-project/hydro-res-auth
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
	TypeCode     *string `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
	// example:
	//
	// [\\"synonym\\",\\"stopword\\",\\"correction\\",\\"category_prediction\\",\\"ner\\",\\"term_weighting\\"]
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s NearbySearchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *NearbySearchResponseBodyData) SetAddress(v string) *NearbySearchResponseBodyData {
	s.Address = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetCityCode(v string) *NearbySearchResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetCityName(v string) *NearbySearchResponseBodyData {
	s.CityName = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetDistanceMeter(v string) *NearbySearchResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetDistrictCode(v string) *NearbySearchResponseBodyData {
	s.DistrictCode = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetDistrictName(v string) *NearbySearchResponseBodyData {
	s.DistrictName = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetId(v string) *NearbySearchResponseBodyData {
	s.Id = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetImages(v []*NearbySearchResponseBodyDataImages) *NearbySearchResponseBodyData {
	s.Images = v
	return s
}

func (s *NearbySearchResponseBodyData) SetLatitude(v string) *NearbySearchResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetLongitude(v string) *NearbySearchResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetMetadata(v *NearbySearchResponseBodyDataMetadata) *NearbySearchResponseBodyData {
	s.Metadata = v
	return s
}

func (s *NearbySearchResponseBodyData) SetName(v string) *NearbySearchResponseBodyData {
	s.Name = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetProvinceCode(v string) *NearbySearchResponseBodyData {
	s.ProvinceCode = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetProvinceName(v string) *NearbySearchResponseBodyData {
	s.ProvinceName = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetTypeCode(v string) *NearbySearchResponseBodyData {
	s.TypeCode = &v
	return s
}

func (s *NearbySearchResponseBodyData) SetTypes(v string) *NearbySearchResponseBodyData {
	s.Types = &v
	return s
}

type NearbySearchResponseBodyDataImages struct {
	// example:
	//
	// https://meeting.dingtalk.com/j/mblzc4zTBWp
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s NearbySearchResponseBodyDataImages) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchResponseBodyDataImages) GoString() string {
	return s.String()
}

func (s *NearbySearchResponseBodyDataImages) SetUrl(v string) *NearbySearchResponseBodyDataImages {
	s.Url = &v
	return s
}

type NearbySearchResponseBodyDataMetadata struct {
	AverageSpend      *string `json:"averageSpend,omitempty" xml:"averageSpend,omitempty"`
	BusinessArea      *string `json:"businessArea,omitempty" xml:"businessArea,omitempty"`
	DailyOpeningHours *string `json:"dailyOpeningHours,omitempty" xml:"dailyOpeningHours,omitempty"`
	MainTag           *string `json:"mainTag,omitempty" xml:"mainTag,omitempty"`
	Phone             *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Score             *string `json:"score,omitempty" xml:"score,omitempty"`
	Tag               *string `json:"tag,omitempty" xml:"tag,omitempty"`
	WeeklyOpeningDays *string `json:"weeklyOpeningDays,omitempty" xml:"weeklyOpeningDays,omitempty"`
}

func (s NearbySearchResponseBodyDataMetadata) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *NearbySearchResponseBodyDataMetadata) SetAverageSpend(v string) *NearbySearchResponseBodyDataMetadata {
	s.AverageSpend = &v
	return s
}

func (s *NearbySearchResponseBodyDataMetadata) SetBusinessArea(v string) *NearbySearchResponseBodyDataMetadata {
	s.BusinessArea = &v
	return s
}

func (s *NearbySearchResponseBodyDataMetadata) SetDailyOpeningHours(v string) *NearbySearchResponseBodyDataMetadata {
	s.DailyOpeningHours = &v
	return s
}

func (s *NearbySearchResponseBodyDataMetadata) SetMainTag(v string) *NearbySearchResponseBodyDataMetadata {
	s.MainTag = &v
	return s
}

func (s *NearbySearchResponseBodyDataMetadata) SetPhone(v string) *NearbySearchResponseBodyDataMetadata {
	s.Phone = &v
	return s
}

func (s *NearbySearchResponseBodyDataMetadata) SetScore(v string) *NearbySearchResponseBodyDataMetadata {
	s.Score = &v
	return s
}

func (s *NearbySearchResponseBodyDataMetadata) SetTag(v string) *NearbySearchResponseBodyDataMetadata {
	s.Tag = &v
	return s
}

func (s *NearbySearchResponseBodyDataMetadata) SetWeeklyOpeningDays(v string) *NearbySearchResponseBodyDataMetadata {
	s.WeeklyOpeningDays = &v
	return s
}

type NearbySearchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NearbySearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NearbySearchResponse) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchResponse) GoString() string {
	return s.String()
}

func (s *NearbySearchResponse) SetHeaders(v map[string]*string) *NearbySearchResponse {
	s.Headers = v
	return s
}

func (s *NearbySearchResponse) SetStatusCode(v int32) *NearbySearchResponse {
	s.StatusCode = &v
	return s
}

func (s *NearbySearchResponse) SetBody(v *NearbySearchResponseBody) *NearbySearchResponse {
	s.Body = v
	return s
}

type NearbySearchNovaRequest struct {
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// example:
	//
	// 39.992873
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// example:
	//
	// 116.310918
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 3000
	Radius *int32 `json:"radius,omitempty" xml:"radius,omitempty"`
	// example:
	//
	// 5
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// GAS_STATION|RESTAURANT|HOTEL|ATTRACTION
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s NearbySearchNovaRequest) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchNovaRequest) GoString() string {
	return s.String()
}

func (s *NearbySearchNovaRequest) SetKeywords(v string) *NearbySearchNovaRequest {
	s.Keywords = &v
	return s
}

func (s *NearbySearchNovaRequest) SetLatitude(v string) *NearbySearchNovaRequest {
	s.Latitude = &v
	return s
}

func (s *NearbySearchNovaRequest) SetLongitude(v string) *NearbySearchNovaRequest {
	s.Longitude = &v
	return s
}

func (s *NearbySearchNovaRequest) SetPage(v int32) *NearbySearchNovaRequest {
	s.Page = &v
	return s
}

func (s *NearbySearchNovaRequest) SetRadius(v int32) *NearbySearchNovaRequest {
	s.Radius = &v
	return s
}

func (s *NearbySearchNovaRequest) SetSize(v int32) *NearbySearchNovaRequest {
	s.Size = &v
	return s
}

func (s *NearbySearchNovaRequest) SetTypes(v string) *NearbySearchNovaRequest {
	s.Types = &v
	return s
}

type NearbySearchNovaResponseBody struct {
	Data []*NearbySearchNovaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// <title>502 Bad Gateway</title>
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s NearbySearchNovaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchNovaResponseBody) GoString() string {
	return s.String()
}

func (s *NearbySearchNovaResponseBody) SetData(v []*NearbySearchNovaResponseBodyData) *NearbySearchNovaResponseBody {
	s.Data = v
	return s
}

func (s *NearbySearchNovaResponseBody) SetErrorCode(v int32) *NearbySearchNovaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NearbySearchNovaResponseBody) SetErrorMessage(v string) *NearbySearchNovaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NearbySearchNovaResponseBody) SetSuccess(v bool) *NearbySearchNovaResponseBody {
	s.Success = &v
	return s
}

type NearbySearchNovaResponseBodyData struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 010
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// example:
	//
	// 445
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	DistrictCode  *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	DistrictName  *string `json:"districtName,omitempty" xml:"districtName,omitempty"`
	// example:
	//
	// 34360
	Id     *string                                   `json:"id,omitempty" xml:"id,omitempty"`
	Images []*NearbySearchNovaResponseBodyDataImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// example:
	//
	// 39.994135
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// example:
	//
	// 108.970162
	Longitude *string                                   `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Metadata  *NearbySearchNovaResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	Name      *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 110000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
	// example:
	//
	// 110200
	TypeCode *string `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
	Types    *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s NearbySearchNovaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchNovaResponseBodyData) GoString() string {
	return s.String()
}

func (s *NearbySearchNovaResponseBodyData) SetAddress(v string) *NearbySearchNovaResponseBodyData {
	s.Address = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetCityCode(v string) *NearbySearchNovaResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetCityName(v string) *NearbySearchNovaResponseBodyData {
	s.CityName = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetDistanceMeter(v string) *NearbySearchNovaResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetDistrictCode(v string) *NearbySearchNovaResponseBodyData {
	s.DistrictCode = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetDistrictName(v string) *NearbySearchNovaResponseBodyData {
	s.DistrictName = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetId(v string) *NearbySearchNovaResponseBodyData {
	s.Id = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetImages(v []*NearbySearchNovaResponseBodyDataImages) *NearbySearchNovaResponseBodyData {
	s.Images = v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetLatitude(v string) *NearbySearchNovaResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetLongitude(v string) *NearbySearchNovaResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetMetadata(v *NearbySearchNovaResponseBodyDataMetadata) *NearbySearchNovaResponseBodyData {
	s.Metadata = v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetName(v string) *NearbySearchNovaResponseBodyData {
	s.Name = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetProvinceCode(v string) *NearbySearchNovaResponseBodyData {
	s.ProvinceCode = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetProvinceName(v string) *NearbySearchNovaResponseBodyData {
	s.ProvinceName = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetTypeCode(v string) *NearbySearchNovaResponseBodyData {
	s.TypeCode = &v
	return s
}

func (s *NearbySearchNovaResponseBodyData) SetTypes(v string) *NearbySearchNovaResponseBodyData {
	s.Types = &v
	return s
}

type NearbySearchNovaResponseBodyDataImages struct {
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://alidocs.dingtalk.com/i/team/nb9XJVAdyvMrOXyA/docs/b9XJlRRKq1BQaGyA
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s NearbySearchNovaResponseBodyDataImages) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchNovaResponseBodyDataImages) GoString() string {
	return s.String()
}

func (s *NearbySearchNovaResponseBodyDataImages) SetTitle(v string) *NearbySearchNovaResponseBodyDataImages {
	s.Title = &v
	return s
}

func (s *NearbySearchNovaResponseBodyDataImages) SetUrl(v string) *NearbySearchNovaResponseBodyDataImages {
	s.Url = &v
	return s
}

type NearbySearchNovaResponseBodyDataMetadata struct {
	// example:
	//
	// 22.00
	AverageSpend *string `json:"averageSpend,omitempty" xml:"averageSpend,omitempty"`
	BusinessArea *string `json:"businessArea,omitempty" xml:"businessArea,omitempty"`
	// example:
	//
	// 11:00-14:00 17:00-21:00
	DailyOpeningHours *string `json:"dailyOpeningHours,omitempty" xml:"dailyOpeningHours,omitempty"`
	MainTag           *string `json:"mainTag,omitempty" xml:"mainTag,omitempty"`
	// example:
	//
	// 029-87983745
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// 4.5
	Score             *string `json:"score,omitempty" xml:"score,omitempty"`
	Tag               *string `json:"tag,omitempty" xml:"tag,omitempty"`
	WeeklyOpeningDays *string `json:"weeklyOpeningDays,omitempty" xml:"weeklyOpeningDays,omitempty"`
}

func (s NearbySearchNovaResponseBodyDataMetadata) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchNovaResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *NearbySearchNovaResponseBodyDataMetadata) SetAverageSpend(v string) *NearbySearchNovaResponseBodyDataMetadata {
	s.AverageSpend = &v
	return s
}

func (s *NearbySearchNovaResponseBodyDataMetadata) SetBusinessArea(v string) *NearbySearchNovaResponseBodyDataMetadata {
	s.BusinessArea = &v
	return s
}

func (s *NearbySearchNovaResponseBodyDataMetadata) SetDailyOpeningHours(v string) *NearbySearchNovaResponseBodyDataMetadata {
	s.DailyOpeningHours = &v
	return s
}

func (s *NearbySearchNovaResponseBodyDataMetadata) SetMainTag(v string) *NearbySearchNovaResponseBodyDataMetadata {
	s.MainTag = &v
	return s
}

func (s *NearbySearchNovaResponseBodyDataMetadata) SetPhone(v string) *NearbySearchNovaResponseBodyDataMetadata {
	s.Phone = &v
	return s
}

func (s *NearbySearchNovaResponseBodyDataMetadata) SetScore(v string) *NearbySearchNovaResponseBodyDataMetadata {
	s.Score = &v
	return s
}

func (s *NearbySearchNovaResponseBodyDataMetadata) SetTag(v string) *NearbySearchNovaResponseBodyDataMetadata {
	s.Tag = &v
	return s
}

func (s *NearbySearchNovaResponseBodyDataMetadata) SetWeeklyOpeningDays(v string) *NearbySearchNovaResponseBodyDataMetadata {
	s.WeeklyOpeningDays = &v
	return s
}

type NearbySearchNovaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NearbySearchNovaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NearbySearchNovaResponse) String() string {
	return tea.Prettify(s)
}

func (s NearbySearchNovaResponse) GoString() string {
	return s.String()
}

func (s *NearbySearchNovaResponse) SetHeaders(v map[string]*string) *NearbySearchNovaResponse {
	s.Headers = v
	return s
}

func (s *NearbySearchNovaResponse) SetStatusCode(v int32) *NearbySearchNovaResponse {
	s.StatusCode = &v
	return s
}

func (s *NearbySearchNovaResponse) SetBody(v *NearbySearchNovaResponseBody) *NearbySearchNovaResponse {
	s.Body = v
	return s
}

type PlaceSearchRequest struct {
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// example:
	//
	// 1
	Page   *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 5
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// GAS_STATION|RESTAURANT|HOTEL|ATTRACTION
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s PlaceSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchRequest) GoString() string {
	return s.String()
}

func (s *PlaceSearchRequest) SetKeywords(v string) *PlaceSearchRequest {
	s.Keywords = &v
	return s
}

func (s *PlaceSearchRequest) SetPage(v int32) *PlaceSearchRequest {
	s.Page = &v
	return s
}

func (s *PlaceSearchRequest) SetRegion(v string) *PlaceSearchRequest {
	s.Region = &v
	return s
}

func (s *PlaceSearchRequest) SetSize(v int32) *PlaceSearchRequest {
	s.Size = &v
	return s
}

func (s *PlaceSearchRequest) SetTypes(v string) *PlaceSearchRequest {
	s.Types = &v
	return s
}

type PlaceSearchResponseBody struct {
	Data []*PlaceSearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied, message: No such namespace namespaces/general-perf-cn-shenzhen-e-default.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PlaceSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchResponseBody) GoString() string {
	return s.String()
}

func (s *PlaceSearchResponseBody) SetData(v []*PlaceSearchResponseBodyData) *PlaceSearchResponseBody {
	s.Data = v
	return s
}

func (s *PlaceSearchResponseBody) SetErrorCode(v int32) *PlaceSearchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PlaceSearchResponseBody) SetErrorMessage(v string) *PlaceSearchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PlaceSearchResponseBody) SetSuccess(v bool) *PlaceSearchResponseBody {
	s.Success = &v
	return s
}

type PlaceSearchResponseBodyData struct {
	Address      *string `json:"address,omitempty" xml:"address,omitempty"`
	CityCode     *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	CityName     *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	DistrictName *string `json:"districtName,omitempty" xml:"districtName,omitempty"`
	// example:
	//
	// 100936
	Id        *string                              `json:"id,omitempty" xml:"id,omitempty"`
	Images    []*PlaceSearchResponseBodyDataImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Latitude  *string                              `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *string                              `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Metadata  *PlaceSearchResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// example:
	//
	// hydro-project/hydro-res-auth
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
	TypeCode     *string `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
	Types        *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s PlaceSearchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *PlaceSearchResponseBodyData) SetAddress(v string) *PlaceSearchResponseBodyData {
	s.Address = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetCityCode(v string) *PlaceSearchResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetCityName(v string) *PlaceSearchResponseBodyData {
	s.CityName = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetDistrictCode(v string) *PlaceSearchResponseBodyData {
	s.DistrictCode = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetDistrictName(v string) *PlaceSearchResponseBodyData {
	s.DistrictName = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetId(v string) *PlaceSearchResponseBodyData {
	s.Id = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetImages(v []*PlaceSearchResponseBodyDataImages) *PlaceSearchResponseBodyData {
	s.Images = v
	return s
}

func (s *PlaceSearchResponseBodyData) SetLatitude(v string) *PlaceSearchResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetLongitude(v string) *PlaceSearchResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetMetadata(v *PlaceSearchResponseBodyDataMetadata) *PlaceSearchResponseBodyData {
	s.Metadata = v
	return s
}

func (s *PlaceSearchResponseBodyData) SetName(v string) *PlaceSearchResponseBodyData {
	s.Name = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetProvinceCode(v string) *PlaceSearchResponseBodyData {
	s.ProvinceCode = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetProvinceName(v string) *PlaceSearchResponseBodyData {
	s.ProvinceName = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetTypeCode(v string) *PlaceSearchResponseBodyData {
	s.TypeCode = &v
	return s
}

func (s *PlaceSearchResponseBodyData) SetTypes(v string) *PlaceSearchResponseBodyData {
	s.Types = &v
	return s
}

type PlaceSearchResponseBodyDataImages struct {
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://store.is.autonavi.com/showpic/d3dd18fa5fb617d02cf7f1aabae80b78
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PlaceSearchResponseBodyDataImages) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchResponseBodyDataImages) GoString() string {
	return s.String()
}

func (s *PlaceSearchResponseBodyDataImages) SetTitle(v string) *PlaceSearchResponseBodyDataImages {
	s.Title = &v
	return s
}

func (s *PlaceSearchResponseBodyDataImages) SetUrl(v string) *PlaceSearchResponseBodyDataImages {
	s.Url = &v
	return s
}

type PlaceSearchResponseBodyDataMetadata struct {
	BusinessArea      *string `json:"businessArea,omitempty" xml:"businessArea,omitempty"`
	DailyOpeningHours *string `json:"dailyOpeningHours,omitempty" xml:"dailyOpeningHours,omitempty"`
	MainTag           *string `json:"mainTag,omitempty" xml:"mainTag,omitempty"`
	Tag               *string `json:"tag,omitempty" xml:"tag,omitempty"`
	WeeklyOpeningDays *string `json:"weeklyOpeningDays,omitempty" xml:"weeklyOpeningDays,omitempty"`
}

func (s PlaceSearchResponseBodyDataMetadata) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *PlaceSearchResponseBodyDataMetadata) SetBusinessArea(v string) *PlaceSearchResponseBodyDataMetadata {
	s.BusinessArea = &v
	return s
}

func (s *PlaceSearchResponseBodyDataMetadata) SetDailyOpeningHours(v string) *PlaceSearchResponseBodyDataMetadata {
	s.DailyOpeningHours = &v
	return s
}

func (s *PlaceSearchResponseBodyDataMetadata) SetMainTag(v string) *PlaceSearchResponseBodyDataMetadata {
	s.MainTag = &v
	return s
}

func (s *PlaceSearchResponseBodyDataMetadata) SetTag(v string) *PlaceSearchResponseBodyDataMetadata {
	s.Tag = &v
	return s
}

func (s *PlaceSearchResponseBodyDataMetadata) SetWeeklyOpeningDays(v string) *PlaceSearchResponseBodyDataMetadata {
	s.WeeklyOpeningDays = &v
	return s
}

type PlaceSearchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PlaceSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PlaceSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchResponse) GoString() string {
	return s.String()
}

func (s *PlaceSearchResponse) SetHeaders(v map[string]*string) *PlaceSearchResponse {
	s.Headers = v
	return s
}

func (s *PlaceSearchResponse) SetStatusCode(v int32) *PlaceSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *PlaceSearchResponse) SetBody(v *PlaceSearchResponseBody) *PlaceSearchResponse {
	s.Body = v
	return s
}

type PlaceSearchNovaRequest struct {
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// example:
	//
	// 1
	Page   *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 5
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// GAS_STATION|RESTAURANT|HOTEL|ATTRACTION
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s PlaceSearchNovaRequest) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchNovaRequest) GoString() string {
	return s.String()
}

func (s *PlaceSearchNovaRequest) SetKeywords(v string) *PlaceSearchNovaRequest {
	s.Keywords = &v
	return s
}

func (s *PlaceSearchNovaRequest) SetPage(v int32) *PlaceSearchNovaRequest {
	s.Page = &v
	return s
}

func (s *PlaceSearchNovaRequest) SetRegion(v string) *PlaceSearchNovaRequest {
	s.Region = &v
	return s
}

func (s *PlaceSearchNovaRequest) SetSize(v int32) *PlaceSearchNovaRequest {
	s.Size = &v
	return s
}

func (s *PlaceSearchNovaRequest) SetTypes(v string) *PlaceSearchNovaRequest {
	s.Types = &v
	return s
}

type PlaceSearchNovaResponseBody struct {
	Data []*PlaceSearchNovaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied, message: Unauthorized.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PlaceSearchNovaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchNovaResponseBody) GoString() string {
	return s.String()
}

func (s *PlaceSearchNovaResponseBody) SetData(v []*PlaceSearchNovaResponseBodyData) *PlaceSearchNovaResponseBody {
	s.Data = v
	return s
}

func (s *PlaceSearchNovaResponseBody) SetErrorCode(v int32) *PlaceSearchNovaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PlaceSearchNovaResponseBody) SetErrorMessage(v string) *PlaceSearchNovaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PlaceSearchNovaResponseBody) SetSuccess(v bool) *PlaceSearchNovaResponseBody {
	s.Success = &v
	return s
}

type PlaceSearchNovaResponseBodyData struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 010
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// example:
	//
	// 445
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 110105
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	DistrictName *string `json:"districtName,omitempty" xml:"districtName,omitempty"`
	// example:
	//
	// 34360
	Id     *string                                  `json:"id,omitempty" xml:"id,omitempty"`
	Images []*PlaceSearchNovaResponseBodyDataImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// example:
	//
	// 39.990039
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// example:
	//
	// 116.482145
	Longitude *string                                  `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Metadata  *PlaceSearchNovaResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// example:
	//
	// hydro-project/hydro-res-auth
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 110000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
	// example:
	//
	// 110200
	TypeCode *string `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
	Types    *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s PlaceSearchNovaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchNovaResponseBodyData) GoString() string {
	return s.String()
}

func (s *PlaceSearchNovaResponseBodyData) SetAddress(v string) *PlaceSearchNovaResponseBodyData {
	s.Address = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetCityCode(v string) *PlaceSearchNovaResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetCityName(v string) *PlaceSearchNovaResponseBodyData {
	s.CityName = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetDistanceMeter(v string) *PlaceSearchNovaResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetDistrictCode(v string) *PlaceSearchNovaResponseBodyData {
	s.DistrictCode = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetDistrictName(v string) *PlaceSearchNovaResponseBodyData {
	s.DistrictName = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetId(v string) *PlaceSearchNovaResponseBodyData {
	s.Id = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetImages(v []*PlaceSearchNovaResponseBodyDataImages) *PlaceSearchNovaResponseBodyData {
	s.Images = v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetLatitude(v string) *PlaceSearchNovaResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetLongitude(v string) *PlaceSearchNovaResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetMetadata(v *PlaceSearchNovaResponseBodyDataMetadata) *PlaceSearchNovaResponseBodyData {
	s.Metadata = v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetName(v string) *PlaceSearchNovaResponseBodyData {
	s.Name = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetProvinceCode(v string) *PlaceSearchNovaResponseBodyData {
	s.ProvinceCode = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetProvinceName(v string) *PlaceSearchNovaResponseBodyData {
	s.ProvinceName = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetTypeCode(v string) *PlaceSearchNovaResponseBodyData {
	s.TypeCode = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyData) SetTypes(v string) *PlaceSearchNovaResponseBodyData {
	s.Types = &v
	return s
}

type PlaceSearchNovaResponseBodyDataImages struct {
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://aos-comment.amap.com/B000A8UNZV/comment/f0a5ca9b58a31f63f8af51f459f75e4b_2048_2048_80.jpg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PlaceSearchNovaResponseBodyDataImages) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchNovaResponseBodyDataImages) GoString() string {
	return s.String()
}

func (s *PlaceSearchNovaResponseBodyDataImages) SetTitle(v string) *PlaceSearchNovaResponseBodyDataImages {
	s.Title = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyDataImages) SetUrl(v string) *PlaceSearchNovaResponseBodyDataImages {
	s.Url = &v
	return s
}

type PlaceSearchNovaResponseBodyDataMetadata struct {
	// example:
	//
	// 78
	AverageSpend      *string `json:"averageSpend,omitempty" xml:"averageSpend,omitempty"`
	BusinessArea      *string `json:"businessArea,omitempty" xml:"businessArea,omitempty"`
	DailyOpeningHours *string `json:"dailyOpeningHours,omitempty" xml:"dailyOpeningHours,omitempty"`
	MainTag           *string `json:"mainTag,omitempty" xml:"mainTag,omitempty"`
	// example:
	//
	// 010-83847583
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// [{\\"value\\":\\"nttd\\",\\"key\\":\\"owner\\"}]
	Tag               *string `json:"tag,omitempty" xml:"tag,omitempty"`
	WeeklyOpeningDays *string `json:"weeklyOpeningDays,omitempty" xml:"weeklyOpeningDays,omitempty"`
}

func (s PlaceSearchNovaResponseBodyDataMetadata) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchNovaResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *PlaceSearchNovaResponseBodyDataMetadata) SetAverageSpend(v string) *PlaceSearchNovaResponseBodyDataMetadata {
	s.AverageSpend = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyDataMetadata) SetBusinessArea(v string) *PlaceSearchNovaResponseBodyDataMetadata {
	s.BusinessArea = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyDataMetadata) SetDailyOpeningHours(v string) *PlaceSearchNovaResponseBodyDataMetadata {
	s.DailyOpeningHours = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyDataMetadata) SetMainTag(v string) *PlaceSearchNovaResponseBodyDataMetadata {
	s.MainTag = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyDataMetadata) SetPhone(v string) *PlaceSearchNovaResponseBodyDataMetadata {
	s.Phone = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyDataMetadata) SetTag(v string) *PlaceSearchNovaResponseBodyDataMetadata {
	s.Tag = &v
	return s
}

func (s *PlaceSearchNovaResponseBodyDataMetadata) SetWeeklyOpeningDays(v string) *PlaceSearchNovaResponseBodyDataMetadata {
	s.WeeklyOpeningDays = &v
	return s
}

type PlaceSearchNovaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PlaceSearchNovaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PlaceSearchNovaResponse) String() string {
	return tea.Prettify(s)
}

func (s PlaceSearchNovaResponse) GoString() string {
	return s.String()
}

func (s *PlaceSearchNovaResponse) SetHeaders(v map[string]*string) *PlaceSearchNovaResponse {
	s.Headers = v
	return s
}

func (s *PlaceSearchNovaResponse) SetStatusCode(v int32) *PlaceSearchNovaResponse {
	s.StatusCode = &v
	return s
}

func (s *PlaceSearchNovaResponse) SetBody(v *PlaceSearchNovaResponseBody) *PlaceSearchNovaResponse {
	s.Body = v
	return s
}

type QueryAttractionsRequest struct {
	Body *AgentBaseQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAttractionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAttractionsRequest) GoString() string {
	return s.String()
}

func (s *QueryAttractionsRequest) SetBody(v *AgentBaseQuery) *QueryAttractionsRequest {
	s.Body = v
	return s
}

type QueryAttractionsResponseBody struct {
	QueryResult *QueryResult `json:"queryResult,omitempty" xml:"queryResult,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryAttractionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAttractionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAttractionsResponseBody) SetQueryResult(v *QueryResult) *QueryAttractionsResponseBody {
	s.QueryResult = v
	return s
}

func (s *QueryAttractionsResponseBody) SetRequestId(v string) *QueryAttractionsResponseBody {
	s.RequestId = &v
	return s
}

type QueryAttractionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAttractionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAttractionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAttractionsResponse) GoString() string {
	return s.String()
}

func (s *QueryAttractionsResponse) SetHeaders(v map[string]*string) *QueryAttractionsResponse {
	s.Headers = v
	return s
}

func (s *QueryAttractionsResponse) SetStatusCode(v int32) *QueryAttractionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAttractionsResponse) SetBody(v *QueryAttractionsResponseBody) *QueryAttractionsResponse {
	s.Body = v
	return s
}

type QueryHotelsRequest struct {
	Body *AgentBaseQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotelsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelsRequest) GoString() string {
	return s.String()
}

func (s *QueryHotelsRequest) SetBody(v *AgentBaseQuery) *QueryHotelsRequest {
	s.Body = v
	return s
}

type QueryHotelsResponseBody struct {
	QueryResult *QueryResult `json:"queryResult,omitempty" xml:"queryResult,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 78032F8B-0157-53F9-B1A8-3BD86ADE38D0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryHotelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotelsResponseBody) SetQueryResult(v *QueryResult) *QueryHotelsResponseBody {
	s.QueryResult = v
	return s
}

func (s *QueryHotelsResponseBody) SetRequestId(v string) *QueryHotelsResponseBody {
	s.RequestId = &v
	return s
}

type QueryHotelsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHotelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotelsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelsResponse) GoString() string {
	return s.String()
}

func (s *QueryHotelsResponse) SetHeaders(v map[string]*string) *QueryHotelsResponse {
	s.Headers = v
	return s
}

func (s *QueryHotelsResponse) SetStatusCode(v int32) *QueryHotelsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotelsResponse) SetBody(v *QueryHotelsResponseBody) *QueryHotelsResponse {
	s.Body = v
	return s
}

type QueryRestaurantsRequest struct {
	Body *AgentBaseQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRestaurantsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRestaurantsRequest) GoString() string {
	return s.String()
}

func (s *QueryRestaurantsRequest) SetBody(v *AgentBaseQuery) *QueryRestaurantsRequest {
	s.Body = v
	return s
}

type QueryRestaurantsResponseBody struct {
	QueryResult *QueryResult `json:"queryResult,omitempty" xml:"queryResult,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D78E16C0-4D44-5065-BFF7-127F7859F687
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryRestaurantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRestaurantsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRestaurantsResponseBody) SetQueryResult(v *QueryResult) *QueryRestaurantsResponseBody {
	s.QueryResult = v
	return s
}

func (s *QueryRestaurantsResponseBody) SetRequestId(v string) *QueryRestaurantsResponseBody {
	s.RequestId = &v
	return s
}

type QueryRestaurantsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRestaurantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRestaurantsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRestaurantsResponse) GoString() string {
	return s.String()
}

func (s *QueryRestaurantsResponse) SetHeaders(v map[string]*string) *QueryRestaurantsResponse {
	s.Headers = v
	return s
}

func (s *QueryRestaurantsResponse) SetStatusCode(v int32) *QueryRestaurantsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRestaurantsResponse) SetBody(v *QueryRestaurantsResponseBody) *QueryRestaurantsResponse {
	s.Body = v
	return s
}

type RectangleTrafficStatusRequest struct {
	// example:
	//
	// 39.966309
	LowerLeftLatitude *string `json:"lowerLeftLatitude,omitempty" xml:"lowerLeftLatitude,omitempty"`
	// example:
	//
	// 116.351147
	LowerLeftLongitude *string `json:"lowerLeftLongitude,omitempty" xml:"lowerLeftLongitude,omitempty"`
	// example:
	//
	// HIGHWAY
	RoadLevel *string `json:"roadLevel,omitempty" xml:"roadLevel,omitempty"`
	// example:
	//
	// 39.968739
	UpperRightLatitude *string `json:"upperRightLatitude,omitempty" xml:"upperRightLatitude,omitempty"`
	// example:
	//
	// 116.35164
	UpperRightLongitude *string `json:"upperRightLongitude,omitempty" xml:"upperRightLongitude,omitempty"`
}

func (s RectangleTrafficStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s RectangleTrafficStatusRequest) GoString() string {
	return s.String()
}

func (s *RectangleTrafficStatusRequest) SetLowerLeftLatitude(v string) *RectangleTrafficStatusRequest {
	s.LowerLeftLatitude = &v
	return s
}

func (s *RectangleTrafficStatusRequest) SetLowerLeftLongitude(v string) *RectangleTrafficStatusRequest {
	s.LowerLeftLongitude = &v
	return s
}

func (s *RectangleTrafficStatusRequest) SetRoadLevel(v string) *RectangleTrafficStatusRequest {
	s.RoadLevel = &v
	return s
}

func (s *RectangleTrafficStatusRequest) SetUpperRightLatitude(v string) *RectangleTrafficStatusRequest {
	s.UpperRightLatitude = &v
	return s
}

func (s *RectangleTrafficStatusRequest) SetUpperRightLongitude(v string) *RectangleTrafficStatusRequest {
	s.UpperRightLongitude = &v
	return s
}

type RectangleTrafficStatusResponseBody struct {
	Data []*RectangleTrafficStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 78032F8B-0157-53F9-B1A8-3BD86ADE38D0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RectangleTrafficStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RectangleTrafficStatusResponseBody) GoString() string {
	return s.String()
}

func (s *RectangleTrafficStatusResponseBody) SetData(v []*RectangleTrafficStatusResponseBodyData) *RectangleTrafficStatusResponseBody {
	s.Data = v
	return s
}

func (s *RectangleTrafficStatusResponseBody) SetErrorCode(v string) *RectangleTrafficStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RectangleTrafficStatusResponseBody) SetErrorMessage(v string) *RectangleTrafficStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RectangleTrafficStatusResponseBody) SetRequestId(v string) *RectangleTrafficStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *RectangleTrafficStatusResponseBody) SetSuccess(v bool) *RectangleTrafficStatusResponseBody {
	s.Success = &v
	return s
}

type RectangleTrafficStatusResponseBodyData struct {
	Description *string                                           `json:"description,omitempty" xml:"description,omitempty"`
	Evaluation  *RectangleTrafficStatusResponseBodyDataEvaluation `json:"evaluation,omitempty" xml:"evaluation,omitempty" type:"Struct"`
}

func (s RectangleTrafficStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RectangleTrafficStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *RectangleTrafficStatusResponseBodyData) SetDescription(v string) *RectangleTrafficStatusResponseBodyData {
	s.Description = &v
	return s
}

func (s *RectangleTrafficStatusResponseBodyData) SetEvaluation(v *RectangleTrafficStatusResponseBodyDataEvaluation) *RectangleTrafficStatusResponseBodyData {
	s.Evaluation = v
	return s
}

type RectangleTrafficStatusResponseBodyDataEvaluation struct {
	// example:
	//
	// 0.00%
	BlockedPercentage *string `json:"blockedPercentage,omitempty" xml:"blockedPercentage,omitempty"`
	// example:
	//
	// 30.77%
	ClearPercentage *string `json:"clearPercentage,omitempty" xml:"clearPercentage,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 7.59%
	HeavyPercentage *string `json:"heavyPercentage,omitempty" xml:"heavyPercentage,omitempty"`
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 61.45%
	UnknownPercentage *string `json:"unknownPercentage,omitempty" xml:"unknownPercentage,omitempty"`
}

func (s RectangleTrafficStatusResponseBodyDataEvaluation) String() string {
	return tea.Prettify(s)
}

func (s RectangleTrafficStatusResponseBodyDataEvaluation) GoString() string {
	return s.String()
}

func (s *RectangleTrafficStatusResponseBodyDataEvaluation) SetBlockedPercentage(v string) *RectangleTrafficStatusResponseBodyDataEvaluation {
	s.BlockedPercentage = &v
	return s
}

func (s *RectangleTrafficStatusResponseBodyDataEvaluation) SetClearPercentage(v string) *RectangleTrafficStatusResponseBodyDataEvaluation {
	s.ClearPercentage = &v
	return s
}

func (s *RectangleTrafficStatusResponseBodyDataEvaluation) SetDescription(v string) *RectangleTrafficStatusResponseBodyDataEvaluation {
	s.Description = &v
	return s
}

func (s *RectangleTrafficStatusResponseBodyDataEvaluation) SetHeavyPercentage(v string) *RectangleTrafficStatusResponseBodyDataEvaluation {
	s.HeavyPercentage = &v
	return s
}

func (s *RectangleTrafficStatusResponseBodyDataEvaluation) SetStatus(v string) *RectangleTrafficStatusResponseBodyDataEvaluation {
	s.Status = &v
	return s
}

func (s *RectangleTrafficStatusResponseBodyDataEvaluation) SetUnknownPercentage(v string) *RectangleTrafficStatusResponseBodyDataEvaluation {
	s.UnknownPercentage = &v
	return s
}

type RectangleTrafficStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RectangleTrafficStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RectangleTrafficStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s RectangleTrafficStatusResponse) GoString() string {
	return s.String()
}

func (s *RectangleTrafficStatusResponse) SetHeaders(v map[string]*string) *RectangleTrafficStatusResponse {
	s.Headers = v
	return s
}

func (s *RectangleTrafficStatusResponse) SetStatusCode(v int32) *RectangleTrafficStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *RectangleTrafficStatusResponse) SetBody(v *RectangleTrafficStatusResponseBody) *RectangleTrafficStatusResponse {
	s.Body = v
	return s
}

type RgeoCodeRequest struct {
	// example:
	//
	// 39.989027
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// example:
	//
	// 116.310918
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

func (s RgeoCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RgeoCodeRequest) GoString() string {
	return s.String()
}

func (s *RgeoCodeRequest) SetLatitude(v string) *RgeoCodeRequest {
	s.Latitude = &v
	return s
}

func (s *RgeoCodeRequest) SetLongitude(v string) *RgeoCodeRequest {
	s.Longitude = &v
	return s
}

type RgeoCodeResponseBody struct {
	Data *RgeoCodeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode    *int32  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RgeoCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RgeoCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RgeoCodeResponseBody) SetData(v *RgeoCodeResponseBodyData) *RgeoCodeResponseBody {
	s.Data = v
	return s
}

func (s *RgeoCodeResponseBody) SetErrorCode(v int32) *RgeoCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RgeoCodeResponseBody) SetErrorMessage(v string) *RgeoCodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RgeoCodeResponseBody) SetSuccess(v bool) *RgeoCodeResponseBody {
	s.Success = &v
	return s
}

type RgeoCodeResponseBodyData struct {
	Building      *RgeoCodeResponseBodyDataBuilding        `json:"building,omitempty" xml:"building,omitempty" type:"Struct"`
	BusinessAreas []*RgeoCodeResponseBodyDataBusinessAreas `json:"businessAreas,omitempty" xml:"businessAreas,omitempty" type:"Repeated"`
	City          *string                                  `json:"city,omitempty" xml:"city,omitempty"`
	// example:
	//
	// 010
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	Country  *string `json:"country,omitempty" xml:"country,omitempty"`
	District *string `json:"district,omitempty" xml:"district,omitempty"`
	// example:
	//
	// 110108
	DistrictCode     *string                               `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	FormattedAddress *string                               `json:"formattedAddress,omitempty" xml:"formattedAddress,omitempty"`
	Neighborhood     *RgeoCodeResponseBodyDataNeighborhood `json:"neighborhood,omitempty" xml:"neighborhood,omitempty" type:"Struct"`
	Province         *string                               `json:"province,omitempty" xml:"province,omitempty"`
	StreetNumber     *RgeoCodeResponseBodyDataStreetNumber `json:"streetNumber,omitempty" xml:"streetNumber,omitempty" type:"Struct"`
	// example:
	//
	// 110108015000
	TownCode *string `json:"townCode,omitempty" xml:"townCode,omitempty"`
	TownShip *string `json:"townShip,omitempty" xml:"townShip,omitempty"`
}

func (s RgeoCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RgeoCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RgeoCodeResponseBodyData) SetBuilding(v *RgeoCodeResponseBodyDataBuilding) *RgeoCodeResponseBodyData {
	s.Building = v
	return s
}

func (s *RgeoCodeResponseBodyData) SetBusinessAreas(v []*RgeoCodeResponseBodyDataBusinessAreas) *RgeoCodeResponseBodyData {
	s.BusinessAreas = v
	return s
}

func (s *RgeoCodeResponseBodyData) SetCity(v string) *RgeoCodeResponseBodyData {
	s.City = &v
	return s
}

func (s *RgeoCodeResponseBodyData) SetCityCode(v string) *RgeoCodeResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *RgeoCodeResponseBodyData) SetCountry(v string) *RgeoCodeResponseBodyData {
	s.Country = &v
	return s
}

func (s *RgeoCodeResponseBodyData) SetDistrict(v string) *RgeoCodeResponseBodyData {
	s.District = &v
	return s
}

func (s *RgeoCodeResponseBodyData) SetDistrictCode(v string) *RgeoCodeResponseBodyData {
	s.DistrictCode = &v
	return s
}

func (s *RgeoCodeResponseBodyData) SetFormattedAddress(v string) *RgeoCodeResponseBodyData {
	s.FormattedAddress = &v
	return s
}

func (s *RgeoCodeResponseBodyData) SetNeighborhood(v *RgeoCodeResponseBodyDataNeighborhood) *RgeoCodeResponseBodyData {
	s.Neighborhood = v
	return s
}

func (s *RgeoCodeResponseBodyData) SetProvince(v string) *RgeoCodeResponseBodyData {
	s.Province = &v
	return s
}

func (s *RgeoCodeResponseBodyData) SetStreetNumber(v *RgeoCodeResponseBodyDataStreetNumber) *RgeoCodeResponseBodyData {
	s.StreetNumber = v
	return s
}

func (s *RgeoCodeResponseBodyData) SetTownCode(v string) *RgeoCodeResponseBodyData {
	s.TownCode = &v
	return s
}

func (s *RgeoCodeResponseBodyData) SetTownShip(v string) *RgeoCodeResponseBodyData {
	s.TownShip = &v
	return s
}

type RgeoCodeResponseBodyDataBuilding struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RgeoCodeResponseBodyDataBuilding) String() string {
	return tea.Prettify(s)
}

func (s RgeoCodeResponseBodyDataBuilding) GoString() string {
	return s.String()
}

func (s *RgeoCodeResponseBodyDataBuilding) SetName(v string) *RgeoCodeResponseBodyDataBuilding {
	s.Name = &v
	return s
}

func (s *RgeoCodeResponseBodyDataBuilding) SetType(v string) *RgeoCodeResponseBodyDataBuilding {
	s.Type = &v
	return s
}

type RgeoCodeResponseBodyDataBusinessAreas struct {
	// example:
	//
	// 110108
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 39.996850
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// example:
	//
	// 116.294214
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RgeoCodeResponseBodyDataBusinessAreas) String() string {
	return tea.Prettify(s)
}

func (s RgeoCodeResponseBodyDataBusinessAreas) GoString() string {
	return s.String()
}

func (s *RgeoCodeResponseBodyDataBusinessAreas) SetId(v string) *RgeoCodeResponseBodyDataBusinessAreas {
	s.Id = &v
	return s
}

func (s *RgeoCodeResponseBodyDataBusinessAreas) SetLatitude(v string) *RgeoCodeResponseBodyDataBusinessAreas {
	s.Latitude = &v
	return s
}

func (s *RgeoCodeResponseBodyDataBusinessAreas) SetLongitude(v string) *RgeoCodeResponseBodyDataBusinessAreas {
	s.Longitude = &v
	return s
}

func (s *RgeoCodeResponseBodyDataBusinessAreas) SetName(v string) *RgeoCodeResponseBodyDataBusinessAreas {
	s.Name = &v
	return s
}

type RgeoCodeResponseBodyDataNeighborhood struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RgeoCodeResponseBodyDataNeighborhood) String() string {
	return tea.Prettify(s)
}

func (s RgeoCodeResponseBodyDataNeighborhood) GoString() string {
	return s.String()
}

func (s *RgeoCodeResponseBodyDataNeighborhood) SetName(v string) *RgeoCodeResponseBodyDataNeighborhood {
	s.Name = &v
	return s
}

func (s *RgeoCodeResponseBodyDataNeighborhood) SetType(v string) *RgeoCodeResponseBodyDataNeighborhood {
	s.Type = &v
	return s
}

type RgeoCodeResponseBodyDataStreetNumber struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// example:
	//
	// 289.832
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 39.986542
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// example:
	//
	// 116.311943
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Number    *string `json:"number,omitempty" xml:"number,omitempty"`
	Street    *string `json:"street,omitempty" xml:"street,omitempty"`
}

func (s RgeoCodeResponseBodyDataStreetNumber) String() string {
	return tea.Prettify(s)
}

func (s RgeoCodeResponseBodyDataStreetNumber) GoString() string {
	return s.String()
}

func (s *RgeoCodeResponseBodyDataStreetNumber) SetDirection(v string) *RgeoCodeResponseBodyDataStreetNumber {
	s.Direction = &v
	return s
}

func (s *RgeoCodeResponseBodyDataStreetNumber) SetDistanceMeter(v string) *RgeoCodeResponseBodyDataStreetNumber {
	s.DistanceMeter = &v
	return s
}

func (s *RgeoCodeResponseBodyDataStreetNumber) SetLatitude(v string) *RgeoCodeResponseBodyDataStreetNumber {
	s.Latitude = &v
	return s
}

func (s *RgeoCodeResponseBodyDataStreetNumber) SetLongitude(v string) *RgeoCodeResponseBodyDataStreetNumber {
	s.Longitude = &v
	return s
}

func (s *RgeoCodeResponseBodyDataStreetNumber) SetNumber(v string) *RgeoCodeResponseBodyDataStreetNumber {
	s.Number = &v
	return s
}

func (s *RgeoCodeResponseBodyDataStreetNumber) SetStreet(v string) *RgeoCodeResponseBodyDataStreetNumber {
	s.Street = &v
	return s
}

type RgeoCodeResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RgeoCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RgeoCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RgeoCodeResponse) GoString() string {
	return s.String()
}

func (s *RgeoCodeResponse) SetHeaders(v map[string]*string) *RgeoCodeResponse {
	s.Headers = v
	return s
}

func (s *RgeoCodeResponse) SetStatusCode(v int32) *RgeoCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RgeoCodeResponse) SetBody(v *RgeoCodeResponseBody) *RgeoCodeResponse {
	s.Body = v
	return s
}

type RoadTrafficStatusRequest struct {
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// example:
	//
	// HIGHWAY
	RoadLevel *string `json:"roadLevel,omitempty" xml:"roadLevel,omitempty"`
	RoadName  *string `json:"roadName,omitempty" xml:"roadName,omitempty"`
}

func (s RoadTrafficStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s RoadTrafficStatusRequest) GoString() string {
	return s.String()
}

func (s *RoadTrafficStatusRequest) SetCity(v string) *RoadTrafficStatusRequest {
	s.City = &v
	return s
}

func (s *RoadTrafficStatusRequest) SetRoadLevel(v string) *RoadTrafficStatusRequest {
	s.RoadLevel = &v
	return s
}

func (s *RoadTrafficStatusRequest) SetRoadName(v string) *RoadTrafficStatusRequest {
	s.RoadName = &v
	return s
}

type RoadTrafficStatusResponseBody struct {
	Data []*RoadTrafficStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 504 gateway error
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RoadTrafficStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RoadTrafficStatusResponseBody) GoString() string {
	return s.String()
}

func (s *RoadTrafficStatusResponseBody) SetData(v []*RoadTrafficStatusResponseBodyData) *RoadTrafficStatusResponseBody {
	s.Data = v
	return s
}

func (s *RoadTrafficStatusResponseBody) SetErrorCode(v string) *RoadTrafficStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RoadTrafficStatusResponseBody) SetErrorMessage(v string) *RoadTrafficStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RoadTrafficStatusResponseBody) SetRequestId(v string) *RoadTrafficStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *RoadTrafficStatusResponseBody) SetSuccess(v bool) *RoadTrafficStatusResponseBody {
	s.Success = &v
	return s
}

type RoadTrafficStatusResponseBodyData struct {
	Description *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	Evaluation  *RoadTrafficStatusResponseBodyDataEvaluation `json:"evaluation,omitempty" xml:"evaluation,omitempty" type:"Struct"`
}

func (s RoadTrafficStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RoadTrafficStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *RoadTrafficStatusResponseBodyData) SetDescription(v string) *RoadTrafficStatusResponseBodyData {
	s.Description = &v
	return s
}

func (s *RoadTrafficStatusResponseBodyData) SetEvaluation(v *RoadTrafficStatusResponseBodyDataEvaluation) *RoadTrafficStatusResponseBodyData {
	s.Evaluation = v
	return s
}

type RoadTrafficStatusResponseBodyDataEvaluation struct {
	// example:
	//
	// 0.00%
	BlockedPercentage *string `json:"blockedPercentage,omitempty" xml:"blockedPercentage,omitempty"`
	// example:
	//
	// 100.00%
	ClearPercentage *string `json:"clearPercentage,omitempty" xml:"clearPercentage,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 0.00%
	HeavyPercentage *string `json:"heavyPercentage,omitempty" xml:"heavyPercentage,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 0.00%
	UnknownPercentage *string `json:"unknownPercentage,omitempty" xml:"unknownPercentage,omitempty"`
}

func (s RoadTrafficStatusResponseBodyDataEvaluation) String() string {
	return tea.Prettify(s)
}

func (s RoadTrafficStatusResponseBodyDataEvaluation) GoString() string {
	return s.String()
}

func (s *RoadTrafficStatusResponseBodyDataEvaluation) SetBlockedPercentage(v string) *RoadTrafficStatusResponseBodyDataEvaluation {
	s.BlockedPercentage = &v
	return s
}

func (s *RoadTrafficStatusResponseBodyDataEvaluation) SetClearPercentage(v string) *RoadTrafficStatusResponseBodyDataEvaluation {
	s.ClearPercentage = &v
	return s
}

func (s *RoadTrafficStatusResponseBodyDataEvaluation) SetDescription(v string) *RoadTrafficStatusResponseBodyDataEvaluation {
	s.Description = &v
	return s
}

func (s *RoadTrafficStatusResponseBodyDataEvaluation) SetHeavyPercentage(v string) *RoadTrafficStatusResponseBodyDataEvaluation {
	s.HeavyPercentage = &v
	return s
}

func (s *RoadTrafficStatusResponseBodyDataEvaluation) SetStatus(v string) *RoadTrafficStatusResponseBodyDataEvaluation {
	s.Status = &v
	return s
}

func (s *RoadTrafficStatusResponseBodyDataEvaluation) SetUnknownPercentage(v string) *RoadTrafficStatusResponseBodyDataEvaluation {
	s.UnknownPercentage = &v
	return s
}

type RoadTrafficStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RoadTrafficStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RoadTrafficStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s RoadTrafficStatusResponse) GoString() string {
	return s.String()
}

func (s *RoadTrafficStatusResponse) SetHeaders(v map[string]*string) *RoadTrafficStatusResponse {
	s.Headers = v
	return s
}

func (s *RoadTrafficStatusResponse) SetStatusCode(v int32) *RoadTrafficStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *RoadTrafficStatusResponse) SetBody(v *RoadTrafficStatusResponseBody) *RoadTrafficStatusResponse {
	s.Body = v
	return s
}

type TransitIntegratedDirectionRequest struct {
	DestinationCity *string `json:"destinationCity,omitempty" xml:"destinationCity,omitempty"`
	// example:
	//
	// 40.345456
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	OriginCity           *string `json:"originCity,omitempty" xml:"originCity,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s TransitIntegratedDirectionRequest) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionRequest) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionRequest) SetDestinationCity(v string) *TransitIntegratedDirectionRequest {
	s.DestinationCity = &v
	return s
}

func (s *TransitIntegratedDirectionRequest) SetDestinationLatitude(v string) *TransitIntegratedDirectionRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *TransitIntegratedDirectionRequest) SetDestinationLongitude(v string) *TransitIntegratedDirectionRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *TransitIntegratedDirectionRequest) SetOriginCity(v string) *TransitIntegratedDirectionRequest {
	s.OriginCity = &v
	return s
}

func (s *TransitIntegratedDirectionRequest) SetOriginLatitude(v string) *TransitIntegratedDirectionRequest {
	s.OriginLatitude = &v
	return s
}

func (s *TransitIntegratedDirectionRequest) SetOriginLongitude(v string) *TransitIntegratedDirectionRequest {
	s.OriginLongitude = &v
	return s
}

type TransitIntegratedDirectionResponseBody struct {
	Data *TransitIntegratedDirectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// <title>502 Bad Gateway</title>
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TransitIntegratedDirectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBody) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBody) SetData(v *TransitIntegratedDirectionResponseBodyData) *TransitIntegratedDirectionResponseBody {
	s.Data = v
	return s
}

func (s *TransitIntegratedDirectionResponseBody) SetErrorCode(v int32) *TransitIntegratedDirectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBody) SetErrorMessage(v string) *TransitIntegratedDirectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBody) SetSuccess(v bool) *TransitIntegratedDirectionResponseBody {
	s.Success = &v
	return s
}

type TransitIntegratedDirectionResponseBodyData struct {
	Cost *TransitIntegratedDirectionResponseBodyDataCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 5
	Count *string `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 39.995197
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 445
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string                                            `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
	Paths           []*TransitIntegratedDirectionResponseBodyDataPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
}

func (s TransitIntegratedDirectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyData) SetCost(v *TransitIntegratedDirectionResponseBodyDataCost) *TransitIntegratedDirectionResponseBodyData {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyData) SetCount(v string) *TransitIntegratedDirectionResponseBodyData {
	s.Count = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyData) SetDestinationLatitude(v string) *TransitIntegratedDirectionResponseBodyData {
	s.DestinationLatitude = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyData) SetDestinationLongitude(v string) *TransitIntegratedDirectionResponseBodyData {
	s.DestinationLongitude = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyData) SetDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyData) SetOriginLatitude(v string) *TransitIntegratedDirectionResponseBodyData {
	s.OriginLatitude = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyData) SetOriginLongitude(v string) *TransitIntegratedDirectionResponseBodyData {
	s.OriginLongitude = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyData) SetPaths(v []*TransitIntegratedDirectionResponseBodyDataPaths) *TransitIntegratedDirectionResponseBodyData {
	s.Paths = v
	return s
}

type TransitIntegratedDirectionResponseBodyDataCost struct {
	// example:
	//
	// 1231
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 6
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	// example:
	//
	// 23
	Tolls         *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataCost) SetDurationSecond(v string) *TransitIntegratedDirectionResponseBodyDataCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataCost) SetTaxiFee(v string) *TransitIntegratedDirectionResponseBodyDataCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataCost) SetTollRoads(v string) *TransitIntegratedDirectionResponseBodyDataCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataCost) SetTolls(v string) *TransitIntegratedDirectionResponseBodyDataCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataCost) SetTrafficLights(v string) *TransitIntegratedDirectionResponseBodyDataCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataCost) SetTransitFee(v string) *TransitIntegratedDirectionResponseBodyDataCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPaths struct {
	Cost *TransitIntegratedDirectionResponseBodyDataPathsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 12000
	DistanceMeter *string                                                    `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	Nightflag     *string                                                    `json:"nightflag,omitempty" xml:"nightflag,omitempty"`
	Segments      []*TransitIntegratedDirectionResponseBodyDataPathsSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	// example:
	//
	// 23435
	WalkingDistanceMeter *string `json:"walkingDistanceMeter,omitempty" xml:"walkingDistanceMeter,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPaths) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPaths) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPaths) SetCost(v *TransitIntegratedDirectionResponseBodyDataPathsCost) *TransitIntegratedDirectionResponseBodyDataPaths {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPaths) SetDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPaths {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPaths) SetNightflag(v string) *TransitIntegratedDirectionResponseBodyDataPaths {
	s.Nightflag = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPaths) SetSegments(v []*TransitIntegratedDirectionResponseBodyDataPathsSegments) *TransitIntegratedDirectionResponseBodyDataPaths {
	s.Segments = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPaths) SetWalkingDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPaths {
	s.WalkingDistanceMeter = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsCost struct {
	// example:
	//
	// 39233
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	// example:
	//
	// 2000
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights     *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 4
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsCost) SetDurationSecond(v string) *TransitIntegratedDirectionResponseBodyDataPathsCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsCost) SetTaxiFee(v string) *TransitIntegratedDirectionResponseBodyDataPathsCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsCost) SetTollRoads(v string) *TransitIntegratedDirectionResponseBodyDataPathsCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsCost) SetTolls(v string) *TransitIntegratedDirectionResponseBodyDataPathsCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsCost) SetTrafficLights(v string) *TransitIntegratedDirectionResponseBodyDataPathsCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsCost) SetTransitFee(v string) *TransitIntegratedDirectionResponseBodyDataPathsCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegments struct {
	Bus     *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBus     `json:"bus,omitempty" xml:"bus,omitempty" type:"Struct"`
	Railway *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway `json:"railway,omitempty" xml:"railway,omitempty" type:"Struct"`
	Taxi    *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi    `json:"taxi,omitempty" xml:"taxi,omitempty" type:"Struct"`
	Walking *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking `json:"walking,omitempty" xml:"walking,omitempty" type:"Struct"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegments) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegments) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegments) SetBus(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBus) *TransitIntegratedDirectionResponseBodyDataPathsSegments {
	s.Bus = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegments) SetRailway(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) *TransitIntegratedDirectionResponseBodyDataPathsSegments {
	s.Railway = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegments) SetTaxi(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) *TransitIntegratedDirectionResponseBodyDataPathsSegments {
	s.Taxi = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegments) SetWalking(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking) *TransitIntegratedDirectionResponseBodyDataPathsSegments {
	s.Walking = v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBus struct {
	Buslines []*TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines `json:"buslines,omitempty" xml:"buslines,omitempty" type:"Repeated"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBus) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBus) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBus) SetBuslines(v []*TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBus {
	s.Buslines = v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines struct {
	ArrivalStop *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop `json:"arrivalStop,omitempty" xml:"arrivalStop,omitempty" type:"Struct"`
	BusTimeTips *string                                                                        `json:"busTimeTips,omitempty" xml:"busTimeTips,omitempty"`
	// example:
	//
	// 0
	Bustimetag    *string                                                                          `json:"bustimetag,omitempty" xml:"bustimetag,omitempty"`
	Cost          *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost          `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	DepartureStop *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop `json:"departureStop,omitempty" xml:"departureStop,omitempty" type:"Struct"`
	// example:
	//
	// 13322
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 2259
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 900000028907
	Id       *string                                                                     `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string                                                                     `json:"name,omitempty" xml:"name,omitempty"`
	Polyline *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesPolyline `json:"polyline,omitempty" xml:"polyline,omitempty" type:"Struct"`
	// example:
	//
	// 0509
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 8
	ViaNum   *string                                                                       `json:"viaNum,omitempty" xml:"viaNum,omitempty"`
	ViaStops []*TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops `json:"viaStops,omitempty" xml:"viaStops,omitempty" type:"Repeated"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetArrivalStop(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.ArrivalStop = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetBusTimeTips(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.BusTimeTips = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetBustimetag(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.Bustimetag = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetCost(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetDepartureStop(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.DepartureStop = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetEndTime(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.EndTime = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.Name = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetPolyline(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesPolyline) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.Polyline = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetStartTime(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.StartTime = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetType(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.Type = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetViaNum(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.ViaNum = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines) SetViaStops(v []*TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslines {
	s.ViaStops = v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop struct {
	Exit *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit `json:"exit,omitempty" xml:"exit,omitempty" type:"Struct"`
	// example:
	//
	// 100935
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 119.82416178385417,30.27139702690972
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop) SetExit(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop {
	s.Exit = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop) SetLocation(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStop {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit struct {
	// example:
	//
	// 900000028907015
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.468213,39.998876
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit) SetLocation(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesArrivalStopExit {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost struct {
	// example:
	//
	// 1521
	DurationSecond    *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights     *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee        *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) SetDurationSecond(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) SetTaxiFee(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) SetTollRoads(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) SetTolls(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) SetTrafficLights(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost) SetTransitFee(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop struct {
	Entrance *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance `json:"entrance,omitempty" xml:"entrance,omitempty" type:"Struct"`
	// example:
	//
	// 60852
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.468213,39.998876
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop) SetEntrance(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop {
	s.Entrance = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop) SetLocation(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStop {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance struct {
	// example:
	//
	// 99088345834
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.468213,39.998876
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance) SetLocation(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesDepartureStopEntrance {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesPolyline struct {
	// example:
	//
	// 116.471544,39.991835
	Polyline *string `json:"polyline,omitempty" xml:"polyline,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesPolyline) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesPolyline) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesPolyline) SetPolyline(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesPolyline {
	s.Polyline = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops struct {
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops) SetLocation(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsBusBuslinesViaStops {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway struct {
	ArrivalStop   *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop   `json:"arrivalStop,omitempty" xml:"arrivalStop,omitempty" type:"Struct"`
	DepartureStop *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop `json:"departureStop,omitempty" xml:"departureStop,omitempty" type:"Struct"`
	// example:
	//
	// 398734
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 434508
	Id     *string                                                                 `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string                                                                 `json:"name,omitempty" xml:"name,omitempty"`
	Spaces []*TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-28 10:07:22
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	Trip *string `json:"trip,omitempty" xml:"trip,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetArrivalStop(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.ArrivalStop = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetDepartureStop(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.DepartureStop = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.Name = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetSpaces(v []*TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.Spaces = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetTime(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.Time = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetTrip(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.Trip = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway) SetType(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailway {
	s.Type = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop struct {
	// example:
	//
	// 023
	Adcode *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	// end
	//
	// example:
	//
	// 1699410466578
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 8234837534
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 101.45625135633681,25.08939480251736
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2024-10-28 10:10:32
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) SetAdcode(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop {
	s.Adcode = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) SetEnd(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop {
	s.End = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) SetLocation(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop {
	s.Name = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop) SetTime(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayArrivalStop {
	s.Time = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop struct {
	// example:
	//
	// 029
	Adcode *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	// example:
	//
	// 100937
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 101.45625135633681,25.08939480251736
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1729440000000
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
	// example:
	//
	// 2024-09-30 10:04:13
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) SetAdcode(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop {
	s.Adcode = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) SetId(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) SetLocation(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) SetName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop {
	s.Name = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) SetStart(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop {
	s.Start = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop) SetTime(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwayDepartureStop {
	s.Time = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 150
	Cost *string `json:"cost,omitempty" xml:"cost,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces) SetCode(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces {
	s.Code = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces) SetCost(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsRailwaySpaces {
	s.Cost = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi struct {
	// example:
	//
	// xxx
	DestinationName *string `json:"destinationName,omitempty" xml:"destinationName,omitempty"`
	// example:
	//
	// 116.464297,39.896458
	DestinationPoint *string `json:"destinationPoint,omitempty" xml:"destinationPoint,omitempty"`
	// example:
	//
	// 1772
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 720
	DriveTimeSecond *string `json:"driveTimeSecond,omitempty" xml:"driveTimeSecond,omitempty"`
	OriginName      *string `json:"originName,omitempty" xml:"originName,omitempty"`
	// example:
	//
	// 116.476597,39.893420
	OriginPoint *string `json:"originPoint,omitempty" xml:"originPoint,omitempty"`
	// example:
	//
	// 13.5
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) SetDestinationName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi {
	s.DestinationName = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) SetDestinationPoint(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi {
	s.DestinationPoint = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) SetDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) SetDriveTimeSecond(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi {
	s.DriveTimeSecond = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) SetOriginName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi {
	s.OriginName = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) SetOriginPoint(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi {
	s.OriginPoint = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi) SetPrice(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsTaxi {
	s.Price = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking struct {
	Cost *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 116.468208,39.998875
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
	// example:
	//
	// 546
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 116.466568,39.995552
	Origin *string                                                                `json:"origin,omitempty" xml:"origin,omitempty"`
	Steps  []*TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking) SetCost(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking) SetDestination(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking {
	s.Destination = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking) SetDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking) SetOrigin(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking {
	s.Origin = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking) SetSteps(v []*TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalking {
	s.Steps = v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost struct {
	// example:
	//
	// 468
	DurationSecond    *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights     *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 3
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) SetDurationSecond(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) SetTaxiFee(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) SetTollRoads(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) SetTolls(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) SetTrafficLights(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost) SetTransitFee(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps struct {
	Cost              *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost     `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction       *string                                                                      `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation       *string                                                                      `json:"orientation,omitempty" xml:"orientation,omitempty"`
	Polyline          *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsPolyline `json:"polyline,omitempty" xml:"polyline,omitempty" type:"Struct"`
	RoadName          *string                                                                      `json:"roadName,omitempty" xml:"roadName,omitempty"`
	StepDistanceMeter *string                                                                      `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) SetCost(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) SetInstruction(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps {
	s.Instruction = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) SetOrientation(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps {
	s.Orientation = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) SetPolyline(v *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsPolyline) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps {
	s.Polyline = v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) SetRoadName(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps {
	s.RoadName = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps) SetStepDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingSteps {
	s.StepDistanceMeter = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost struct {
	// example:
	//
	// 435
	DurationSecond    *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights     *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee        *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) SetDurationSecond(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) SetTaxiFee(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) SetTollRoads(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) SetTolls(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) SetTrafficLights(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost) SetTransitFee(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsPolyline struct {
	// example:
	//
	// 116.467751,39.997631;116.467430,39.997631
	Polyline *string `json:"polyline,omitempty" xml:"polyline,omitempty"`
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsPolyline) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsPolyline) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsPolyline) SetPolyline(v string) *TransitIntegratedDirectionResponseBodyDataPathsSegmentsWalkingStepsPolyline {
	s.Polyline = &v
	return s
}

type TransitIntegratedDirectionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransitIntegratedDirectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransitIntegratedDirectionResponse) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionResponse) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionResponse) SetHeaders(v map[string]*string) *TransitIntegratedDirectionResponse {
	s.Headers = v
	return s
}

func (s *TransitIntegratedDirectionResponse) SetStatusCode(v int32) *TransitIntegratedDirectionResponse {
	s.StatusCode = &v
	return s
}

func (s *TransitIntegratedDirectionResponse) SetBody(v *TransitIntegratedDirectionResponseBody) *TransitIntegratedDirectionResponse {
	s.Body = v
	return s
}

type TransitIntegratedDirectionOldRequest struct {
	DestinationCity *string `json:"destinationCity,omitempty" xml:"destinationCity,omitempty"`
	// example:
	//
	// 39.896463
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	OriginCity           *string `json:"originCity,omitempty" xml:"originCity,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s TransitIntegratedDirectionOldRequest) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldRequest) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldRequest) SetDestinationCity(v string) *TransitIntegratedDirectionOldRequest {
	s.DestinationCity = &v
	return s
}

func (s *TransitIntegratedDirectionOldRequest) SetDestinationLatitude(v string) *TransitIntegratedDirectionOldRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *TransitIntegratedDirectionOldRequest) SetDestinationLongitude(v string) *TransitIntegratedDirectionOldRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *TransitIntegratedDirectionOldRequest) SetOriginCity(v string) *TransitIntegratedDirectionOldRequest {
	s.OriginCity = &v
	return s
}

func (s *TransitIntegratedDirectionOldRequest) SetOriginLatitude(v string) *TransitIntegratedDirectionOldRequest {
	s.OriginLatitude = &v
	return s
}

func (s *TransitIntegratedDirectionOldRequest) SetOriginLongitude(v string) *TransitIntegratedDirectionOldRequest {
	s.OriginLongitude = &v
	return s
}

type TransitIntegratedDirectionOldResponseBody struct {
	Data []*TransitIntegratedDirectionOldResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// <title>502 Bad Gateway</title>
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBody) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBody) SetData(v []*TransitIntegratedDirectionOldResponseBodyData) *TransitIntegratedDirectionOldResponseBody {
	s.Data = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBody) SetErrorCode(v int32) *TransitIntegratedDirectionOldResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBody) SetErrorMessage(v string) *TransitIntegratedDirectionOldResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBody) SetRequestId(v string) *TransitIntegratedDirectionOldResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBody) SetSuccess(v bool) *TransitIntegratedDirectionOldResponseBody {
	s.Success = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyData struct {
	Cost *TransitIntegratedDirectionOldResponseBodyDataCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 445
	DistanceMeter *string                                                  `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	Nightflag     *string                                                  `json:"nightflag,omitempty" xml:"nightflag,omitempty"`
	Segments      []*TransitIntegratedDirectionOldResponseBodyDataSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	// example:
	//
	// 11000
	WalkingDistanceMeter *string `json:"walkingDistanceMeter,omitempty" xml:"walkingDistanceMeter,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyData) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyData) SetCost(v *TransitIntegratedDirectionOldResponseBodyDataCost) *TransitIntegratedDirectionOldResponseBodyData {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyData) SetDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyData) SetNightflag(v string) *TransitIntegratedDirectionOldResponseBodyData {
	s.Nightflag = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyData) SetSegments(v []*TransitIntegratedDirectionOldResponseBodyDataSegments) *TransitIntegratedDirectionOldResponseBodyData {
	s.Segments = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyData) SetWalkingDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyData {
	s.WalkingDistanceMeter = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataCost struct {
	// example:
	//
	// 1231
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 6
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	// example:
	//
	// 23
	Tolls         *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataCost) SetDurationSecond(v string) *TransitIntegratedDirectionOldResponseBodyDataCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataCost) SetTaxiFee(v string) *TransitIntegratedDirectionOldResponseBodyDataCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataCost) SetTollRoads(v string) *TransitIntegratedDirectionOldResponseBodyDataCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataCost) SetTolls(v string) *TransitIntegratedDirectionOldResponseBodyDataCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataCost) SetTrafficLights(v string) *TransitIntegratedDirectionOldResponseBodyDataCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataCost) SetTransitFee(v string) *TransitIntegratedDirectionOldResponseBodyDataCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegments struct {
	Bus     *TransitIntegratedDirectionOldResponseBodyDataSegmentsBus     `json:"bus,omitempty" xml:"bus,omitempty" type:"Struct"`
	Railway *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway `json:"railway,omitempty" xml:"railway,omitempty" type:"Struct"`
	Taxi    *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi    `json:"taxi,omitempty" xml:"taxi,omitempty" type:"Struct"`
	Walking *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking `json:"walking,omitempty" xml:"walking,omitempty" type:"Struct"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegments) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegments) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegments) SetBus(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsBus) *TransitIntegratedDirectionOldResponseBodyDataSegments {
	s.Bus = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegments) SetRailway(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) *TransitIntegratedDirectionOldResponseBodyDataSegments {
	s.Railway = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegments) SetTaxi(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) *TransitIntegratedDirectionOldResponseBodyDataSegments {
	s.Taxi = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegments) SetWalking(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking) *TransitIntegratedDirectionOldResponseBodyDataSegments {
	s.Walking = v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsBus struct {
	Buslines []*TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines `json:"buslines,omitempty" xml:"buslines,omitempty" type:"Repeated"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBus) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBus) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBus) SetBuslines(v []*TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBus {
	s.Buslines = v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines struct {
	ArrivalStop *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop `json:"arrivalStop,omitempty" xml:"arrivalStop,omitempty" type:"Struct"`
	BusTimeTips *string                                                                      `json:"busTimeTips,omitempty" xml:"busTimeTips,omitempty"`
	// example:
	//
	// 09:00---21:00
	Bustimetag    *string                                                                        `json:"bustimetag,omitempty" xml:"bustimetag,omitempty"`
	Cost          *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost          `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	DepartureStop *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop `json:"departureStop,omitempty" xml:"departureStop,omitempty" type:"Struct"`
	// example:
	//
	// 3000
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 1710432000000
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1841625
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// Sheet1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1730448000000
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// cluster
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 3
	ViaNum   *string                                                                     `json:"viaNum,omitempty" xml:"viaNum,omitempty"`
	ViaStops []*TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops `json:"viaStops,omitempty" xml:"viaStops,omitempty" type:"Repeated"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetArrivalStop(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.ArrivalStop = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetBusTimeTips(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.BusTimeTips = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetBustimetag(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.Bustimetag = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetCost(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetDepartureStop(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.DepartureStop = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetEndTime(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.EndTime = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.Name = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetStartTime(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.StartTime = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetType(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.Type = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetViaNum(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.ViaNum = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines) SetViaStops(v []*TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslines {
	s.ViaStops = v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop struct {
	Exit *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit `json:"exit,omitempty" xml:"exit,omitempty" type:"Struct"`
	// example:
	//
	// rpe-fu181hrp24i8fc6r38fbsnn
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.467430,39.997627
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// znzmo_cqz
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop) SetExit(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop {
	s.Exit = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop) SetLocation(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStop {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit struct {
	// example:
	//
	// 99900009
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.467430,39.997627
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// review
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit) SetLocation(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesArrivalStopExit {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost struct {
	// example:
	//
	// 348394
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 12
	TaxiFee *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	// example:
	//
	// 444
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	// example:
	//
	// xxx
	TollRoads *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	// example:
	//
	// 34
	Tolls *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 4
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 3
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) SetDurationSecond(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) SetTaxiFee(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) SetTollRoads(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) SetTolls(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) SetTrafficLights(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost) SetTransitFee(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop struct {
	Entrance *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance `json:"entrance,omitempty" xml:"entrance,omitempty" type:"Struct"`
	// example:
	//
	// st-628f075f-11638689
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.467430,39.997627
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 1597871211794192-1700571612867-LcFKnaDJRsoAsYpltZzfrc7Ts0TY1DYOrID+QVy0WxxxPOY8+bod62GWhHB1NK1lITs5tu2zIwknT3R9S8XCAA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop) SetEntrance(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop {
	s.Entrance = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop) SetLocation(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStop {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance struct {
	// example:
	//
	// 124100037
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.467430,39.997627
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// tyc_cust
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance) SetLocation(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesDepartureStopEntrance {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops struct {
	// example:
	//
	// 136200191
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 114.23980387369792,30.60205837673611
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// WordsWeightTest
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops) SetLocation(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsBusBuslinesViaStops {
	s.Name = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway struct {
	ArrivalStop   *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop   `json:"arrivalStop,omitempty" xml:"arrivalStop,omitempty" type:"Struct"`
	DepartureStop *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop `json:"departureStop,omitempty" xml:"departureStop,omitempty" type:"Struct"`
	// example:
	//
	// 1200
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// rpe-kw141hrpoh8vttte2lrvhdb
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// hbh_rectify
	Name   *string                                                               `json:"name,omitempty" xml:"name,omitempty"`
	Spaces []*TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-11-04 10:26:41
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 4002
	Trip *string `json:"trip,omitempty" xml:"trip,omitempty"`
	// example:
	//
	// OFFLINE_DICT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetArrivalStop(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.ArrivalStop = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetDepartureStop(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.DepartureStop = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.Name = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetSpaces(v []*TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.Spaces = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetTime(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.Time = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetTrip(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.Trip = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway) SetType(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailway {
	s.Type = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop struct {
	// example:
	//
	// 100200
	Adcode *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	// end
	//
	// example:
	//
	// 1728553266042
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 138700037
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.467430,39.997627
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 1634343014745990-1699252808337-stack_10340126_20231106144004.log
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2024-11-09 10:00:46
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) SetAdcode(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop {
	s.Adcode = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) SetEnd(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop {
	s.End = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) SetLocation(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop {
	s.Name = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop) SetTime(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayArrivalStop {
	s.Time = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop struct {
	// example:
	//
	// 1002200
	Adcode *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	// example:
	//
	// 148100008
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 116.467430,39.997627
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// stop_words
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1727141356337
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
	// example:
	//
	// 2024-11-04 10:26:41
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) SetAdcode(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop {
	s.Adcode = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) SetId(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop {
	s.Id = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) SetLocation(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop {
	s.Location = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) SetName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop {
	s.Name = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) SetStart(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop {
	s.Start = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop) SetTime(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwayDepartureStop {
	s.Time = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// None
	Cost *string `json:"cost,omitempty" xml:"cost,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces) SetCode(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces {
	s.Code = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces) SetCost(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsRailwaySpaces {
	s.Cost = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi struct {
	DestinationName *string `json:"destinationName,omitempty" xml:"destinationName,omitempty"`
	// example:
	//
	// 116.461662,39.944807
	DestinationPoint *string `json:"destinationPoint,omitempty" xml:"destinationPoint,omitempty"`
	// example:
	//
	// 1000
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 1200
	DriveTimeSecond *string `json:"driveTimeSecond,omitempty" xml:"driveTimeSecond,omitempty"`
	OriginName      *string `json:"originName,omitempty" xml:"originName,omitempty"`
	// example:
	//
	// 116.461662,39.944807
	OriginPoint *string `json:"originPoint,omitempty" xml:"originPoint,omitempty"`
	// example:
	//
	// 241000
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) SetDestinationName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi {
	s.DestinationName = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) SetDestinationPoint(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi {
	s.DestinationPoint = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) SetDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) SetDriveTimeSecond(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi {
	s.DriveTimeSecond = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) SetOriginName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi {
	s.OriginName = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) SetOriginPoint(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi {
	s.OriginPoint = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi) SetPrice(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsTaxi {
	s.Price = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking struct {
	Cost *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 116.468208,39.998875
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
	// example:
	//
	// 2000
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 116.466568,39.995552
	Origin *string                                                              `json:"origin,omitempty" xml:"origin,omitempty"`
	Steps  []*TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking) SetCost(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking) SetDestination(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking {
	s.Destination = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking) SetDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking {
	s.DistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking) SetOrigin(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking {
	s.Origin = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking) SetSteps(v []*TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalking {
	s.Steps = v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost struct {
	// example:
	//
	// 1293
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	// example:
	//
	// 4000
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	// example:
	//
	// xxx
	TollRoads *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	// example:
	//
	// 1203
	Tolls *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 5
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 3
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) SetDurationSecond(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) SetTaxiFee(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) SetTollRoads(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) SetTolls(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) SetTrafficLights(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost) SetTransitFee(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps struct {
	Cost        *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                                                `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                                                `json:"orientation,omitempty" xml:"orientation,omitempty"`
	// example:
	//
	// xxx
	RoadName *string `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 3000
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps) SetCost(v *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps {
	s.Cost = v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps) SetInstruction(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps {
	s.Instruction = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps) SetOrientation(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps {
	s.Orientation = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps) SetRoadName(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps {
	s.RoadName = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps) SetStepDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingSteps {
	s.StepDistanceMeter = &v
	return s
}

type TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost struct {
	// example:
	//
	// 3435
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 10
	TaxiFee *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	// example:
	//
	// 2000
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	// example:
	//
	// xxx
	TollRoads *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	// example:
	//
	// 3
	Tolls *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 3
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 2
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) SetDurationSecond(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) SetTaxiFee(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) SetTollDistanceMeter(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) SetTollRoads(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost {
	s.TollRoads = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) SetTolls(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost {
	s.Tolls = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) SetTrafficLights(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost) SetTransitFee(v string) *TransitIntegratedDirectionOldResponseBodyDataSegmentsWalkingStepsCost {
	s.TransitFee = &v
	return s
}

type TransitIntegratedDirectionOldResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransitIntegratedDirectionOldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransitIntegratedDirectionOldResponse) String() string {
	return tea.Prettify(s)
}

func (s TransitIntegratedDirectionOldResponse) GoString() string {
	return s.String()
}

func (s *TransitIntegratedDirectionOldResponse) SetHeaders(v map[string]*string) *TransitIntegratedDirectionOldResponse {
	s.Headers = v
	return s
}

func (s *TransitIntegratedDirectionOldResponse) SetStatusCode(v int32) *TransitIntegratedDirectionOldResponse {
	s.StatusCode = &v
	return s
}

func (s *TransitIntegratedDirectionOldResponse) SetBody(v *TransitIntegratedDirectionOldResponseBody) *TransitIntegratedDirectionOldResponse {
	s.Body = v
	return s
}

type WalkingDirectionRequest struct {
	// example:
	//
	// 39.896463
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s WalkingDirectionRequest) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionRequest) GoString() string {
	return s.String()
}

func (s *WalkingDirectionRequest) SetDestinationLatitude(v string) *WalkingDirectionRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *WalkingDirectionRequest) SetDestinationLongitude(v string) *WalkingDirectionRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *WalkingDirectionRequest) SetOriginLatitude(v string) *WalkingDirectionRequest {
	s.OriginLatitude = &v
	return s
}

func (s *WalkingDirectionRequest) SetOriginLongitude(v string) *WalkingDirectionRequest {
	s.OriginLongitude = &v
	return s
}

type WalkingDirectionResponseBody struct {
	Data []*WalkingDirectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied, message: Unauthorized.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WalkingDirectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionResponseBody) GoString() string {
	return s.String()
}

func (s *WalkingDirectionResponseBody) SetData(v []*WalkingDirectionResponseBodyData) *WalkingDirectionResponseBody {
	s.Data = v
	return s
}

func (s *WalkingDirectionResponseBody) SetErrorCode(v string) *WalkingDirectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WalkingDirectionResponseBody) SetErrorMessage(v string) *WalkingDirectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *WalkingDirectionResponseBody) SetRequestId(v string) *WalkingDirectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *WalkingDirectionResponseBody) SetSuccess(v bool) *WalkingDirectionResponseBody {
	s.Success = &v
	return s
}

type WalkingDirectionResponseBodyData struct {
	Cost *WalkingDirectionResponseBodyDataCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 445
	DistanceMeter *string                                  `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	Steps         []*WalkingDirectionResponseBodyDataSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s WalkingDirectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *WalkingDirectionResponseBodyData) SetCost(v *WalkingDirectionResponseBodyDataCost) *WalkingDirectionResponseBodyData {
	s.Cost = v
	return s
}

func (s *WalkingDirectionResponseBodyData) SetDistanceMeter(v string) *WalkingDirectionResponseBodyData {
	s.DistanceMeter = &v
	return s
}

func (s *WalkingDirectionResponseBodyData) SetSteps(v []*WalkingDirectionResponseBodyDataSteps) *WalkingDirectionResponseBodyData {
	s.Steps = v
	return s
}

type WalkingDirectionResponseBodyDataCost struct {
	// example:
	//
	// 1231
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 6
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	// example:
	//
	// 23
	Tolls         *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s WalkingDirectionResponseBodyDataCost) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionResponseBodyDataCost) GoString() string {
	return s.String()
}

func (s *WalkingDirectionResponseBodyDataCost) SetDurationSecond(v string) *WalkingDirectionResponseBodyDataCost {
	s.DurationSecond = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataCost) SetTaxiFee(v string) *WalkingDirectionResponseBodyDataCost {
	s.TaxiFee = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataCost) SetTollDistanceMeter(v string) *WalkingDirectionResponseBodyDataCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataCost) SetTollRoads(v string) *WalkingDirectionResponseBodyDataCost {
	s.TollRoads = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataCost) SetTolls(v string) *WalkingDirectionResponseBodyDataCost {
	s.Tolls = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataCost) SetTrafficLights(v string) *WalkingDirectionResponseBodyDataCost {
	s.TrafficLights = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataCost) SetTransitFee(v string) *WalkingDirectionResponseBodyDataCost {
	s.TransitFee = &v
	return s
}

type WalkingDirectionResponseBodyDataSteps struct {
	Cost        *WalkingDirectionResponseBodyDataStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                    `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                    `json:"orientation,omitempty" xml:"orientation,omitempty"`
	// example:
	//
	// xxx
	RoadName *string `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 16
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s WalkingDirectionResponseBodyDataSteps) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionResponseBodyDataSteps) GoString() string {
	return s.String()
}

func (s *WalkingDirectionResponseBodyDataSteps) SetCost(v *WalkingDirectionResponseBodyDataStepsCost) *WalkingDirectionResponseBodyDataSteps {
	s.Cost = v
	return s
}

func (s *WalkingDirectionResponseBodyDataSteps) SetInstruction(v string) *WalkingDirectionResponseBodyDataSteps {
	s.Instruction = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataSteps) SetOrientation(v string) *WalkingDirectionResponseBodyDataSteps {
	s.Orientation = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataSteps) SetRoadName(v string) *WalkingDirectionResponseBodyDataSteps {
	s.RoadName = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataSteps) SetStepDistanceMeter(v string) *WalkingDirectionResponseBodyDataSteps {
	s.StepDistanceMeter = &v
	return s
}

type WalkingDirectionResponseBodyDataStepsCost struct {
	// example:
	//
	// 2345
	DurationSecond    *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 3
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee    *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s WalkingDirectionResponseBodyDataStepsCost) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionResponseBodyDataStepsCost) GoString() string {
	return s.String()
}

func (s *WalkingDirectionResponseBodyDataStepsCost) SetDurationSecond(v string) *WalkingDirectionResponseBodyDataStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataStepsCost) SetTaxiFee(v string) *WalkingDirectionResponseBodyDataStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataStepsCost) SetTollDistanceMeter(v string) *WalkingDirectionResponseBodyDataStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataStepsCost) SetTollRoads(v string) *WalkingDirectionResponseBodyDataStepsCost {
	s.TollRoads = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataStepsCost) SetTolls(v string) *WalkingDirectionResponseBodyDataStepsCost {
	s.Tolls = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataStepsCost) SetTrafficLights(v string) *WalkingDirectionResponseBodyDataStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *WalkingDirectionResponseBodyDataStepsCost) SetTransitFee(v string) *WalkingDirectionResponseBodyDataStepsCost {
	s.TransitFee = &v
	return s
}

type WalkingDirectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WalkingDirectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WalkingDirectionResponse) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionResponse) GoString() string {
	return s.String()
}

func (s *WalkingDirectionResponse) SetHeaders(v map[string]*string) *WalkingDirectionResponse {
	s.Headers = v
	return s
}

func (s *WalkingDirectionResponse) SetStatusCode(v int32) *WalkingDirectionResponse {
	s.StatusCode = &v
	return s
}

func (s *WalkingDirectionResponse) SetBody(v *WalkingDirectionResponseBody) *WalkingDirectionResponse {
	s.Body = v
	return s
}

type WalkingDirectionNovaRequest struct {
	// example:
	//
	// 40.345456
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
}

func (s WalkingDirectionNovaRequest) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionNovaRequest) GoString() string {
	return s.String()
}

func (s *WalkingDirectionNovaRequest) SetDestinationLatitude(v string) *WalkingDirectionNovaRequest {
	s.DestinationLatitude = &v
	return s
}

func (s *WalkingDirectionNovaRequest) SetDestinationLongitude(v string) *WalkingDirectionNovaRequest {
	s.DestinationLongitude = &v
	return s
}

func (s *WalkingDirectionNovaRequest) SetOriginLatitude(v string) *WalkingDirectionNovaRequest {
	s.OriginLatitude = &v
	return s
}

func (s *WalkingDirectionNovaRequest) SetOriginLongitude(v string) *WalkingDirectionNovaRequest {
	s.OriginLongitude = &v
	return s
}

type WalkingDirectionNovaResponseBody struct {
	Data *WalkingDirectionNovaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Access was denied
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WalkingDirectionNovaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionNovaResponseBody) GoString() string {
	return s.String()
}

func (s *WalkingDirectionNovaResponseBody) SetData(v *WalkingDirectionNovaResponseBodyData) *WalkingDirectionNovaResponseBody {
	s.Data = v
	return s
}

func (s *WalkingDirectionNovaResponseBody) SetErrorCode(v int32) *WalkingDirectionNovaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WalkingDirectionNovaResponseBody) SetErrorMessage(v string) *WalkingDirectionNovaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *WalkingDirectionNovaResponseBody) SetSuccess(v bool) *WalkingDirectionNovaResponseBody {
	s.Success = &v
	return s
}

type WalkingDirectionNovaResponseBodyData struct {
	Count *string `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 40.020642
	DestinationLatitude *string `json:"destinationLatitude,omitempty" xml:"destinationLatitude,omitempty"`
	// example:
	//
	// 116.46424
	DestinationLongitude *string `json:"destinationLongitude,omitempty" xml:"destinationLongitude,omitempty"`
	// example:
	//
	// 39.995197
	OriginLatitude *string `json:"originLatitude,omitempty" xml:"originLatitude,omitempty"`
	// example:
	//
	// 116.466485
	OriginLongitude *string                                      `json:"originLongitude,omitempty" xml:"originLongitude,omitempty"`
	Paths           []*WalkingDirectionNovaResponseBodyDataPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	TaxiCost *string `json:"taxiCost,omitempty" xml:"taxiCost,omitempty"`
}

func (s WalkingDirectionNovaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionNovaResponseBodyData) GoString() string {
	return s.String()
}

func (s *WalkingDirectionNovaResponseBodyData) SetCount(v string) *WalkingDirectionNovaResponseBodyData {
	s.Count = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyData) SetDestinationLatitude(v string) *WalkingDirectionNovaResponseBodyData {
	s.DestinationLatitude = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyData) SetDestinationLongitude(v string) *WalkingDirectionNovaResponseBodyData {
	s.DestinationLongitude = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyData) SetOriginLatitude(v string) *WalkingDirectionNovaResponseBodyData {
	s.OriginLatitude = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyData) SetOriginLongitude(v string) *WalkingDirectionNovaResponseBodyData {
	s.OriginLongitude = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyData) SetPaths(v []*WalkingDirectionNovaResponseBodyDataPaths) *WalkingDirectionNovaResponseBodyData {
	s.Paths = v
	return s
}

func (s *WalkingDirectionNovaResponseBodyData) SetTaxiCost(v string) *WalkingDirectionNovaResponseBodyData {
	s.TaxiCost = &v
	return s
}

type WalkingDirectionNovaResponseBodyDataPaths struct {
	Cost *WalkingDirectionNovaResponseBodyDataPathsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	// example:
	//
	// 12000
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// example:
	//
	// 39223
	DurationSecond *string                                           `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	Restriction    *string                                           `json:"restriction,omitempty" xml:"restriction,omitempty"`
	Steps          []*WalkingDirectionNovaResponseBodyDataPathsSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s WalkingDirectionNovaResponseBodyDataPaths) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionNovaResponseBodyDataPaths) GoString() string {
	return s.String()
}

func (s *WalkingDirectionNovaResponseBodyDataPaths) SetCost(v *WalkingDirectionNovaResponseBodyDataPathsCost) *WalkingDirectionNovaResponseBodyDataPaths {
	s.Cost = v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPaths) SetDistanceMeter(v string) *WalkingDirectionNovaResponseBodyDataPaths {
	s.DistanceMeter = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPaths) SetDurationSecond(v string) *WalkingDirectionNovaResponseBodyDataPaths {
	s.DurationSecond = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPaths) SetRestriction(v string) *WalkingDirectionNovaResponseBodyDataPaths {
	s.Restriction = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPaths) SetSteps(v []*WalkingDirectionNovaResponseBodyDataPathsSteps) *WalkingDirectionNovaResponseBodyDataPaths {
	s.Steps = v
	return s
}

type WalkingDirectionNovaResponseBodyDataPathsCost struct {
	// example:
	//
	// 1232
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	// example:
	//
	// 3
	TrafficLights *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	// example:
	//
	// 10
	TransitFee *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s WalkingDirectionNovaResponseBodyDataPathsCost) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionNovaResponseBodyDataPathsCost) GoString() string {
	return s.String()
}

func (s *WalkingDirectionNovaResponseBodyDataPathsCost) SetDurationSecond(v string) *WalkingDirectionNovaResponseBodyDataPathsCost {
	s.DurationSecond = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsCost) SetTaxiFee(v string) *WalkingDirectionNovaResponseBodyDataPathsCost {
	s.TaxiFee = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsCost) SetTollDistanceMeter(v string) *WalkingDirectionNovaResponseBodyDataPathsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsCost) SetTollRoads(v string) *WalkingDirectionNovaResponseBodyDataPathsCost {
	s.TollRoads = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsCost) SetTolls(v string) *WalkingDirectionNovaResponseBodyDataPathsCost {
	s.Tolls = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsCost) SetTrafficLights(v string) *WalkingDirectionNovaResponseBodyDataPathsCost {
	s.TrafficLights = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsCost) SetTransitFee(v string) *WalkingDirectionNovaResponseBodyDataPathsCost {
	s.TransitFee = &v
	return s
}

type WalkingDirectionNovaResponseBodyDataPathsSteps struct {
	Cost        *WalkingDirectionNovaResponseBodyDataPathsStepsCost `json:"cost,omitempty" xml:"cost,omitempty" type:"Struct"`
	Instruction *string                                             `json:"instruction,omitempty" xml:"instruction,omitempty"`
	Orientation *string                                             `json:"orientation,omitempty" xml:"orientation,omitempty"`
	Polyline    *string                                             `json:"polyline,omitempty" xml:"polyline,omitempty"`
	RoadName    *string                                             `json:"roadName,omitempty" xml:"roadName,omitempty"`
	// example:
	//
	// 1665
	StepDistanceMeter *string `json:"stepDistanceMeter,omitempty" xml:"stepDistanceMeter,omitempty"`
}

func (s WalkingDirectionNovaResponseBodyDataPathsSteps) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionNovaResponseBodyDataPathsSteps) GoString() string {
	return s.String()
}

func (s *WalkingDirectionNovaResponseBodyDataPathsSteps) SetCost(v *WalkingDirectionNovaResponseBodyDataPathsStepsCost) *WalkingDirectionNovaResponseBodyDataPathsSteps {
	s.Cost = v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsSteps) SetInstruction(v string) *WalkingDirectionNovaResponseBodyDataPathsSteps {
	s.Instruction = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsSteps) SetOrientation(v string) *WalkingDirectionNovaResponseBodyDataPathsSteps {
	s.Orientation = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsSteps) SetPolyline(v string) *WalkingDirectionNovaResponseBodyDataPathsSteps {
	s.Polyline = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsSteps) SetRoadName(v string) *WalkingDirectionNovaResponseBodyDataPathsSteps {
	s.RoadName = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsSteps) SetStepDistanceMeter(v string) *WalkingDirectionNovaResponseBodyDataPathsSteps {
	s.StepDistanceMeter = &v
	return s
}

type WalkingDirectionNovaResponseBodyDataPathsStepsCost struct {
	// example:
	//
	// 13
	DurationSecond *string `json:"durationSecond,omitempty" xml:"durationSecond,omitempty"`
	// example:
	//
	// 20
	TaxiFee           *string `json:"taxiFee,omitempty" xml:"taxiFee,omitempty"`
	TollDistanceMeter *string `json:"tollDistanceMeter,omitempty" xml:"tollDistanceMeter,omitempty"`
	TollRoads         *string `json:"tollRoads,omitempty" xml:"tollRoads,omitempty"`
	Tolls             *string `json:"tolls,omitempty" xml:"tolls,omitempty"`
	TrafficLights     *string `json:"trafficLights,omitempty" xml:"trafficLights,omitempty"`
	TransitFee        *string `json:"transitFee,omitempty" xml:"transitFee,omitempty"`
}

func (s WalkingDirectionNovaResponseBodyDataPathsStepsCost) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionNovaResponseBodyDataPathsStepsCost) GoString() string {
	return s.String()
}

func (s *WalkingDirectionNovaResponseBodyDataPathsStepsCost) SetDurationSecond(v string) *WalkingDirectionNovaResponseBodyDataPathsStepsCost {
	s.DurationSecond = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsStepsCost) SetTaxiFee(v string) *WalkingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TaxiFee = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsStepsCost) SetTollDistanceMeter(v string) *WalkingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TollDistanceMeter = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsStepsCost) SetTollRoads(v string) *WalkingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TollRoads = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsStepsCost) SetTolls(v string) *WalkingDirectionNovaResponseBodyDataPathsStepsCost {
	s.Tolls = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsStepsCost) SetTrafficLights(v string) *WalkingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TrafficLights = &v
	return s
}

func (s *WalkingDirectionNovaResponseBodyDataPathsStepsCost) SetTransitFee(v string) *WalkingDirectionNovaResponseBodyDataPathsStepsCost {
	s.TransitFee = &v
	return s
}

type WalkingDirectionNovaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WalkingDirectionNovaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WalkingDirectionNovaResponse) String() string {
	return tea.Prettify(s)
}

func (s WalkingDirectionNovaResponse) GoString() string {
	return s.String()
}

func (s *WalkingDirectionNovaResponse) SetHeaders(v map[string]*string) *WalkingDirectionNovaResponse {
	s.Headers = v
	return s
}

func (s *WalkingDirectionNovaResponse) SetStatusCode(v int32) *WalkingDirectionNovaResponse {
	s.StatusCode = &v
	return s
}

func (s *WalkingDirectionNovaResponse) SetBody(v *WalkingDirectionNovaResponseBody) *WalkingDirectionNovaResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("iqs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 根据起终点坐标检索符合条件的骑行路线规划方案
//
// @param request - BicyclingDirectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BicyclingDirectionResponse
func (client *Client) BicyclingDirectionWithOptions(request *BicyclingDirectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BicyclingDirectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BicyclingDirection"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/direction/bicycling"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BicyclingDirectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的骑行路线规划方案
//
// @param request - BicyclingDirectionRequest
//
// @return BicyclingDirectionResponse
func (client *Client) BicyclingDirection(request *BicyclingDirectionRequest) (_result *BicyclingDirectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BicyclingDirectionResponse{}
	_body, _err := client.BicyclingDirectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的骑行路线规划方案
//
// @param request - BicyclingDirectionNovaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BicyclingDirectionNovaResponse
func (client *Client) BicyclingDirectionNovaWithOptions(request *BicyclingDirectionNovaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BicyclingDirectionNovaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BicyclingDirectionNova"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v2/direction/bicycling"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BicyclingDirectionNovaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的骑行路线规划方案
//
// @param request - BicyclingDirectionNovaRequest
//
// @return BicyclingDirectionNovaResponse
func (client *Client) BicyclingDirectionNova(request *BicyclingDirectionNovaRequest) (_result *BicyclingDirectionNovaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BicyclingDirectionNovaResponse{}
	_body, _err := client.BicyclingDirectionNovaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实时查询圆形区域内的交通信息查询
//
// @param request - CircleTrafficStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CircleTrafficStatusResponse
func (client *Client) CircleTrafficStatusWithOptions(request *CircleTrafficStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CircleTrafficStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["longitude"] = request.Longitude
	}

	if !tea.BoolValue(util.IsUnset(request.Radius)) {
		query["radius"] = request.Radius
	}

	if !tea.BoolValue(util.IsUnset(request.RoadLevel)) {
		query["roadLevel"] = request.RoadLevel
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CircleTrafficStatus"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/traffic/status/circle"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CircleTrafficStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时查询圆形区域内的交通信息查询
//
// @param request - CircleTrafficStatusRequest
//
// @return CircleTrafficStatusResponse
func (client *Client) CircleTrafficStatus(request *CircleTrafficStatusRequest) (_result *CircleTrafficStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CircleTrafficStatusResponse{}
	_body, _err := client.CircleTrafficStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自然语言通用查询
//
// @param request - CommonQueryBySceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommonQueryBySceneResponse
func (client *Client) CommonQueryBySceneWithOptions(request *CommonQueryBySceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CommonQueryBySceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CommonQueryByScene"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/amap-function-call-agent/iqs-agent-service/v2/nl/common"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CommonQueryBySceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自然语言通用查询
//
// @param request - CommonQueryBySceneRequest
//
// @return CommonQueryBySceneResponse
func (client *Client) CommonQueryByScene(request *CommonQueryBySceneRequest) (_result *CommonQueryBySceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CommonQueryBySceneResponse{}
	_body, _err := client.CommonQueryBySceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的驾车路线规划方案
//
// @param request - DrivingDirectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DrivingDirectionResponse
func (client *Client) DrivingDirectionWithOptions(request *DrivingDirectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DrivingDirectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DrivingDirection"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/direction/driving"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DrivingDirectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的驾车路线规划方案
//
// @param request - DrivingDirectionRequest
//
// @return DrivingDirectionResponse
func (client *Client) DrivingDirection(request *DrivingDirectionRequest) (_result *DrivingDirectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DrivingDirectionResponse{}
	_body, _err := client.DrivingDirectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的驾车路线规划方案
//
// @param request - DrivingDirectionNovaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DrivingDirectionNovaResponse
func (client *Client) DrivingDirectionNovaWithOptions(request *DrivingDirectionNovaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DrivingDirectionNovaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CarType)) {
		query["carType"] = request.CarType
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.Plate)) {
		query["plate"] = request.Plate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DrivingDirectionNova"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v2/direction/driving"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DrivingDirectionNovaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的驾车路线规划方案
//
// @param request - DrivingDirectionNovaRequest
//
// @return DrivingDirectionNovaResponse
func (client *Client) DrivingDirectionNova(request *DrivingDirectionNovaRequest) (_result *DrivingDirectionNovaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DrivingDirectionNovaResponse{}
	_body, _err := client.DrivingDirectionNovaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的电动车路线规划方案
//
// @param request - ElectrobikeDirectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ElectrobikeDirectionResponse
func (client *Client) ElectrobikeDirectionWithOptions(request *ElectrobikeDirectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ElectrobikeDirectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ElectrobikeDirection"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/direction/electrobike"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ElectrobikeDirectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的电动车路线规划方案
//
// @param request - ElectrobikeDirectionRequest
//
// @return ElectrobikeDirectionResponse
func (client *Client) ElectrobikeDirection(request *ElectrobikeDirectionRequest) (_result *ElectrobikeDirectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ElectrobikeDirectionResponse{}
	_body, _err := client.ElectrobikeDirectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 电动车路线规划方案V2
//
// @param request - ElectrobikeDirectionNovaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ElectrobikeDirectionNovaResponse
func (client *Client) ElectrobikeDirectionNovaWithOptions(request *ElectrobikeDirectionNovaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ElectrobikeDirectionNovaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ElectrobikeDirectionNova"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v2/direction/electrobike"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ElectrobikeDirectionNovaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 电动车路线规划方案V2
//
// @param request - ElectrobikeDirectionNovaRequest
//
// @return ElectrobikeDirectionNovaResponse
func (client *Client) ElectrobikeDirectionNova(request *ElectrobikeDirectionNovaRequest) (_result *ElectrobikeDirectionNovaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ElectrobikeDirectionNovaResponse{}
	_body, _err := client.ElectrobikeDirectionNovaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 地理编码，将详细的结构化地址转换为高德经纬度坐标
//
// @param request - GeoCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GeoCodeResponse
func (client *Client) GeoCodeWithOptions(request *GeoCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GeoCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["city"] = request.City
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GeoCode"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/geocode/geo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GeoCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 地理编码，将详细的结构化地址转换为高德经纬度坐标
//
// @param request - GeoCodeRequest
//
// @return GeoCodeResponse
func (client *Client) GeoCode(request *GeoCodeRequest) (_result *GeoCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GeoCodeResponse{}
	_body, _err := client.GeoCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过经纬度查询附近的地点
//
// @param request - NearbySearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NearbySearchResponse
func (client *Client) NearbySearchWithOptions(request *NearbySearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *NearbySearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["longitude"] = request.Longitude
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Radius)) {
		query["radius"] = request.Radius
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NearbySearch"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/pois/nearby"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &NearbySearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过经纬度查询附近的地点
//
// @param request - NearbySearchRequest
//
// @return NearbySearchResponse
func (client *Client) NearbySearch(request *NearbySearchRequest) (_result *NearbySearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &NearbySearchResponse{}
	_body, _err := client.NearbySearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过经纬度查询附近的地点
//
// @param request - NearbySearchNovaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NearbySearchNovaResponse
func (client *Client) NearbySearchNovaWithOptions(request *NearbySearchNovaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *NearbySearchNovaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["longitude"] = request.Longitude
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Radius)) {
		query["radius"] = request.Radius
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NearbySearchNova"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v2/pois/nearby"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &NearbySearchNovaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过经纬度查询附近的地点
//
// @param request - NearbySearchNovaRequest
//
// @return NearbySearchNovaResponse
func (client *Client) NearbySearchNova(request *NearbySearchNovaRequest) (_result *NearbySearchNovaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &NearbySearchNovaResponse{}
	_body, _err := client.NearbySearchNovaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过关键词搜索地点
//
// @param request - PlaceSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PlaceSearchResponse
func (client *Client) PlaceSearchWithOptions(request *PlaceSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PlaceSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PlaceSearch"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/pois/place"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PlaceSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过关键词搜索地点
//
// @param request - PlaceSearchRequest
//
// @return PlaceSearchResponse
func (client *Client) PlaceSearch(request *PlaceSearchRequest) (_result *PlaceSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PlaceSearchResponse{}
	_body, _err := client.PlaceSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过关键词搜索地点
//
// @param request - PlaceSearchNovaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PlaceSearchNovaResponse
func (client *Client) PlaceSearchNovaWithOptions(request *PlaceSearchNovaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PlaceSearchNovaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PlaceSearchNova"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v2/pois/place"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PlaceSearchNovaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过关键词搜索地点
//
// @param request - PlaceSearchNovaRequest
//
// @return PlaceSearchNovaResponse
func (client *Client) PlaceSearchNova(request *PlaceSearchNovaRequest) (_result *PlaceSearchNovaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PlaceSearchNovaResponse{}
	_body, _err := client.PlaceSearchNovaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 景点查询
//
// @param request - QueryAttractionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAttractionsResponse
func (client *Client) QueryAttractionsWithOptions(request *QueryAttractionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryAttractionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAttractions"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/amap-function-call-agent/iqs-agent-service/v1/nl/attractions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAttractionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 景点查询
//
// @param request - QueryAttractionsRequest
//
// @return QueryAttractionsResponse
func (client *Client) QueryAttractions(request *QueryAttractionsRequest) (_result *QueryAttractionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAttractionsResponse{}
	_body, _err := client.QueryAttractionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店查询
//
// @param request - QueryHotelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHotelsResponse
func (client *Client) QueryHotelsWithOptions(request *QueryHotelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryHotelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHotels"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/amap-function-call-agent/iqs-agent-service/v1/nl/hotels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHotelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店查询
//
// @param request - QueryHotelsRequest
//
// @return QueryHotelsResponse
func (client *Client) QueryHotels(request *QueryHotelsRequest) (_result *QueryHotelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryHotelsResponse{}
	_body, _err := client.QueryHotelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 餐厅查询
//
// @param request - QueryRestaurantsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRestaurantsResponse
func (client *Client) QueryRestaurantsWithOptions(request *QueryRestaurantsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryRestaurantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRestaurants"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/amap-function-call-agent/iqs-agent-service/v1/nl/restaurants"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRestaurantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 餐厅查询
//
// @param request - QueryRestaurantsRequest
//
// @return QueryRestaurantsResponse
func (client *Client) QueryRestaurants(request *QueryRestaurantsRequest) (_result *QueryRestaurantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRestaurantsResponse{}
	_body, _err := client.QueryRestaurantsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实时查询矩形区域内的交通信息查询
//
// @param request - RectangleTrafficStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RectangleTrafficStatusResponse
func (client *Client) RectangleTrafficStatusWithOptions(request *RectangleTrafficStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RectangleTrafficStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LowerLeftLatitude)) {
		query["lowerLeftLatitude"] = request.LowerLeftLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.LowerLeftLongitude)) {
		query["lowerLeftLongitude"] = request.LowerLeftLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.RoadLevel)) {
		query["roadLevel"] = request.RoadLevel
	}

	if !tea.BoolValue(util.IsUnset(request.UpperRightLatitude)) {
		query["upperRightLatitude"] = request.UpperRightLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.UpperRightLongitude)) {
		query["upperRightLongitude"] = request.UpperRightLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RectangleTrafficStatus"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/traffic/status/rectangle"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RectangleTrafficStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时查询矩形区域内的交通信息查询
//
// @param request - RectangleTrafficStatusRequest
//
// @return RectangleTrafficStatusResponse
func (client *Client) RectangleTrafficStatus(request *RectangleTrafficStatusRequest) (_result *RectangleTrafficStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RectangleTrafficStatusResponse{}
	_body, _err := client.RectangleTrafficStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 逆地理编码，将经纬度转换为详细结构化的地址信息
//
// @param request - RgeoCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RgeoCodeResponse
func (client *Client) RgeoCodeWithOptions(request *RgeoCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RgeoCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["longitude"] = request.Longitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RgeoCode"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/geocode/regeo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RgeoCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 逆地理编码，将经纬度转换为详细结构化的地址信息
//
// @param request - RgeoCodeRequest
//
// @return RgeoCodeResponse
func (client *Client) RgeoCode(request *RgeoCodeRequest) (_result *RgeoCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RgeoCodeResponse{}
	_body, _err := client.RgeoCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实时查询指定线路的交通信息
//
// @param request - RoadTrafficStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RoadTrafficStatusResponse
func (client *Client) RoadTrafficStatusWithOptions(request *RoadTrafficStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RoadTrafficStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["city"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.RoadLevel)) {
		query["roadLevel"] = request.RoadLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RoadName)) {
		query["roadName"] = request.RoadName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RoadTrafficStatus"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/traffic/status/road"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RoadTrafficStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时查询指定线路的交通信息
//
// @param request - RoadTrafficStatusRequest
//
// @return RoadTrafficStatusResponse
func (client *Client) RoadTrafficStatus(request *RoadTrafficStatusRequest) (_result *RoadTrafficStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RoadTrafficStatusResponse{}
	_body, _err := client.RoadTrafficStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 公共交通路线规划方案
//
// @param request - TransitIntegratedDirectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransitIntegratedDirectionResponse
func (client *Client) TransitIntegratedDirectionWithOptions(request *TransitIntegratedDirectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TransitIntegratedDirectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCity)) {
		query["destinationCity"] = request.DestinationCity
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginCity)) {
		query["originCity"] = request.OriginCity
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransitIntegratedDirection"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v2/direction/transit/integrated"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TransitIntegratedDirectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 公共交通路线规划方案
//
// @param request - TransitIntegratedDirectionRequest
//
// @return TransitIntegratedDirectionResponse
func (client *Client) TransitIntegratedDirection(request *TransitIntegratedDirectionRequest) (_result *TransitIntegratedDirectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TransitIntegratedDirectionResponse{}
	_body, _err := client.TransitIntegratedDirectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的公共交通路线规划方案
//
// @param request - TransitIntegratedDirectionOldRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransitIntegratedDirectionOldResponse
func (client *Client) TransitIntegratedDirectionOldWithOptions(request *TransitIntegratedDirectionOldRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TransitIntegratedDirectionOldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCity)) {
		query["destinationCity"] = request.DestinationCity
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginCity)) {
		query["originCity"] = request.OriginCity
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransitIntegratedDirectionOld"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/direction/transit/integrated"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TransitIntegratedDirectionOldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的公共交通路线规划方案
//
// @param request - TransitIntegratedDirectionOldRequest
//
// @return TransitIntegratedDirectionOldResponse
func (client *Client) TransitIntegratedDirectionOld(request *TransitIntegratedDirectionOldRequest) (_result *TransitIntegratedDirectionOldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TransitIntegratedDirectionOldResponse{}
	_body, _err := client.TransitIntegratedDirectionOldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的步行路线规划方案
//
// @param request - WalkingDirectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WalkingDirectionResponse
func (client *Client) WalkingDirectionWithOptions(request *WalkingDirectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *WalkingDirectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("WalkingDirection"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v1/direction/walking"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &WalkingDirectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的步行路线规划方案
//
// @param request - WalkingDirectionRequest
//
// @return WalkingDirectionResponse
func (client *Client) WalkingDirection(request *WalkingDirectionRequest) (_result *WalkingDirectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WalkingDirectionResponse{}
	_body, _err := client.WalkingDirectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的步行路线规划方案
//
// @param request - WalkingDirectionNovaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WalkingDirectionNovaResponse
func (client *Client) WalkingDirectionNovaWithOptions(request *WalkingDirectionNovaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *WalkingDirectionNovaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationLatitude)) {
		query["destinationLatitude"] = request.DestinationLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationLongitude)) {
		query["destinationLongitude"] = request.DestinationLongitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLatitude)) {
		query["originLatitude"] = request.OriginLatitude
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLongitude)) {
		query["originLongitude"] = request.OriginLongitude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("WalkingDirectionNova"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ipaas/v2/direction/walking"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &WalkingDirectionNovaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起终点坐标检索符合条件的步行路线规划方案
//
// @param request - WalkingDirectionNovaRequest
//
// @return WalkingDirectionNovaResponse
func (client *Client) WalkingDirectionNova(request *WalkingDirectionNovaRequest) (_result *WalkingDirectionNovaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WalkingDirectionNovaResponse{}
	_body, _err := client.WalkingDirectionNovaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
