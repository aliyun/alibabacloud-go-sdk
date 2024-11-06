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

type DrivingDirectionNovaRequest struct {
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
}

func (s DrivingDirectionNovaRequest) String() string {
	return tea.Prettify(s)
}

func (s DrivingDirectionNovaRequest) GoString() string {
	return s.String()
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
