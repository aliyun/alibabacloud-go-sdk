// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighRiskPieChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHighRiskPieChartResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeHighRiskPieChartResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeHighRiskPieChartResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHighRiskPieChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeHighRiskPieChartResponseBodyResultObject) *DescribeHighRiskPieChartResponseBody
	GetResultObject() *DescribeHighRiskPieChartResponseBodyResultObject
	SetSuccess(v bool) *DescribeHighRiskPieChartResponseBody
	GetSuccess() *bool
}

type DescribeHighRiskPieChartResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error details
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeHighRiskPieChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHighRiskPieChartResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeHighRiskPieChartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHighRiskPieChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHighRiskPieChartResponseBody) GetResultObject() *DescribeHighRiskPieChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeHighRiskPieChartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHighRiskPieChartResponseBody) SetCode(v string) *DescribeHighRiskPieChartResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBody) SetHttpStatusCode(v string) *DescribeHighRiskPieChartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBody) SetMessage(v string) *DescribeHighRiskPieChartResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBody) SetRequestId(v string) *DescribeHighRiskPieChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBody) SetResultObject(v *DescribeHighRiskPieChartResponseBodyResultObject) *DescribeHighRiskPieChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBody) SetSuccess(v bool) *DescribeHighRiskPieChartResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObject struct {
	// High-risk IP city
	HighRiskIPCity *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity `json:"highRiskIPCity,omitempty" xml:"highRiskIPCity,omitempty" type:"Struct"`
	// High-risk IP归属province
	HighRiskIPProvince *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince `json:"highRiskIPProvince,omitempty" xml:"highRiskIPProvince,omitempty" type:"Struct"`
	// High-risk mobile phone归属city
	HighRiskMobileCity *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity `json:"highRiskMobileCity,omitempty" xml:"highRiskMobileCity,omitempty" type:"Struct"`
	// High-risk mobile phone\\"s province of origin
	HighRiskMobileProvince *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince `json:"highRiskMobileProvince,omitempty" xml:"highRiskMobileProvince,omitempty" type:"Struct"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) GetHighRiskIPCity() *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity {
	return s.HighRiskIPCity
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) GetHighRiskIPProvince() *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince {
	return s.HighRiskIPProvince
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) GetHighRiskMobileCity() *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity {
	return s.HighRiskMobileCity
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) GetHighRiskMobileProvince() *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince {
	return s.HighRiskMobileProvince
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) SetHighRiskIPCity(v *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) *DescribeHighRiskPieChartResponseBodyResultObject {
	s.HighRiskIPCity = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) SetHighRiskIPProvince(v *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) *DescribeHighRiskPieChartResponseBodyResultObject {
	s.HighRiskIPProvince = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) SetHighRiskMobileCity(v *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) *DescribeHighRiskPieChartResponseBodyResultObject {
	s.HighRiskMobileCity = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) SetHighRiskMobileProvince(v *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) *DescribeHighRiskPieChartResponseBodyResultObject {
	s.HighRiskMobileProvince = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity struct {
	// Chart flag, default true
	//
	// example:
	//
	// true
	Animation *bool `json:"animation,omitempty" xml:"animation,omitempty"`
	// Belonging grid.
	Grid *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid `json:"grid,omitempty" xml:"grid,omitempty" type:"Struct"`
	// Chart data
	Series []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) GetAnimation() *bool {
	return s.Animation
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) GetGrid() *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid {
	return s.Grid
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) GetSeries() []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries {
	return s.Series
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) SetAnimation(v bool) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity {
	s.Animation = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) SetGrid(v *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity {
	s.Grid = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) SetSeries(v []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity {
	s.Series = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCity) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid struct {
	// Chart flag, default false
	//
	// example:
	//
	// false
	Show *bool `json:"show,omitempty" xml:"show,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid) GetShow() *bool {
	return s.Show
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid) SetShow(v bool) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid {
	s.Show = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCityGrid) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries struct {
	// Returned data object
	Data []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Field name
	//
	// example:
	//
	// 杭州市
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Chart identifier, default is false
	//
	// example:
	//
	// false
	RoseType *string `json:"roseType,omitempty" xml:"roseType,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) GetData() []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData {
	return s.Data
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) GetName() *string {
	return s.Name
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) GetRoseType() *string {
	return s.RoseType
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) SetData(v []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries {
	s.Data = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) SetName(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries {
	s.Name = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) SetRoseType(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries {
	s.RoseType = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeries) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData struct {
	// Field name
	//
	// example:
	//
	// 杭州市
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Data value
	//
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData) GetName() *string {
	return s.Name
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData) GetValue() *string {
	return s.Value
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData) SetName(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData {
	s.Name = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData) SetValue(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData {
	s.Value = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPCitySeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince struct {
	// Chart identifier, default is true
	//
	// example:
	//
	// true
	Animation *bool `json:"animation,omitempty" xml:"animation,omitempty"`
	// Belonging grid.
	Grid *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid `json:"grid,omitempty" xml:"grid,omitempty" type:"Struct"`
	// Chart data
	Series []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) GetAnimation() *bool {
	return s.Animation
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) GetGrid() *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid {
	return s.Grid
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) GetSeries() []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries {
	return s.Series
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) SetAnimation(v bool) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince {
	s.Animation = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) SetGrid(v *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince {
	s.Grid = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) SetSeries(v []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince {
	s.Series = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvince) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid struct {
	// Chart identifier, default is false
	//
	// example:
	//
	// false
	Show *bool `json:"show,omitempty" xml:"show,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid) GetShow() *bool {
	return s.Show
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid) SetShow(v bool) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid {
	s.Show = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceGrid) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries struct {
	// Returned data object
	Data []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Field name
	//
	// example:
	//
	// 浙江省
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Chart identifier, default is false
	//
	// example:
	//
	// false
	RoseType *string `json:"roseType,omitempty" xml:"roseType,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) GetData() []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData {
	return s.Data
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) GetName() *string {
	return s.Name
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) GetRoseType() *string {
	return s.RoseType
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) SetData(v []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries {
	s.Data = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) SetName(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries {
	s.Name = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) SetRoseType(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries {
	s.RoseType = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData struct {
	// Variable name
	//
	// example:
	//
	// 浙江省
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Data value
	//
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData) GetName() *string {
	return s.Name
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData) GetValue() *string {
	return s.Value
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData) SetName(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData {
	s.Name = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData) SetValue(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData {
	s.Value = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskIPProvinceSeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity struct {
	// Chart flag, default is true
	//
	// example:
	//
	// true
	Animation *bool `json:"animation,omitempty" xml:"animation,omitempty"`
	// Belongs to grid.
	Grid *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid `json:"grid,omitempty" xml:"grid,omitempty" type:"Struct"`
	// Chart data
	Series []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) GetAnimation() *bool {
	return s.Animation
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) GetGrid() *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid {
	return s.Grid
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) GetSeries() []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries {
	return s.Series
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) SetAnimation(v bool) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity {
	s.Animation = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) SetGrid(v *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity {
	s.Grid = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) SetSeries(v []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity {
	s.Series = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCity) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid struct {
	// Chart flag, default is false
	//
	// example:
	//
	// false
	Show *bool `json:"show,omitempty" xml:"show,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid) GetShow() *bool {
	return s.Show
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid) SetShow(v bool) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid {
	s.Show = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCityGrid) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries struct {
	// Returned data object
	Data []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Field name
	//
	// example:
	//
	// 杭州市
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Chart flag, default is false
	//
	// example:
	//
	// false
	RoseType *string `json:"roseType,omitempty" xml:"roseType,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) GetData() []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData {
	return s.Data
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) GetName() *string {
	return s.Name
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) GetRoseType() *string {
	return s.RoseType
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) SetData(v []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries {
	s.Data = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) SetName(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries {
	s.Name = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) SetRoseType(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries {
	s.RoseType = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeries) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData struct {
	// Field name
	//
	// example:
	//
	// 杭州市
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Data value
	//
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData) GetName() *string {
	return s.Name
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData) GetValue() *string {
	return s.Value
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData) SetName(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData {
	s.Name = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData) SetValue(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData {
	s.Value = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileCitySeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince struct {
	// Indicator, default true
	//
	// example:
	//
	// true
	Animation *bool `json:"animation,omitempty" xml:"animation,omitempty"`
	// Belongs to grid.
	Grid *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid `json:"grid,omitempty" xml:"grid,omitempty" type:"Struct"`
	// Chart data
	Series []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) GetAnimation() *bool {
	return s.Animation
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) GetGrid() *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid {
	return s.Grid
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) GetSeries() []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries {
	return s.Series
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) SetAnimation(v bool) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince {
	s.Animation = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) SetGrid(v *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince {
	s.Grid = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) SetSeries(v []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince {
	s.Series = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvince) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid struct {
	// Chart flag, default is false
	//
	// example:
	//
	// false
	Show *bool `json:"show,omitempty" xml:"show,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid) GetShow() *bool {
	return s.Show
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid) SetShow(v bool) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid {
	s.Show = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceGrid) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries struct {
	// High-risk position data.
	Data []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Display title
	//
	// example:
	//
	// 分值区间占比
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Chart identifier, default false
	//
	// example:
	//
	// false
	RoseType *string `json:"roseType,omitempty" xml:"roseType,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) GetData() []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData {
	return s.Data
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) GetName() *string {
	return s.Name
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) GetRoseType() *string {
	return s.RoseType
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) SetData(v []*DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries {
	s.Data = v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) SetName(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries {
	s.Name = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) SetRoseType(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries {
	s.RoseType = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData struct {
	// Field name
	//
	// example:
	//
	// 浙江省
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Data value
	//
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData) GetName() *string {
	return s.Name
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData) GetValue() *string {
	return s.Value
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData) SetName(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData {
	s.Name = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData) SetValue(v string) *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData {
	s.Value = &v
	return s
}

func (s *DescribeHighRiskPieChartResponseBodyResultObjectHighRiskMobileProvinceSeriesData) Validate() error {
	return dara.Validate(s)
}
