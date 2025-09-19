// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChartDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllChartSubTypeList(v []*DescribeChartDataResponseBodyAllChartSubTypeList) *DescribeChartDataResponseBody
	GetAllChartSubTypeList() []*DescribeChartDataResponseBodyAllChartSubTypeList
	SetChartDataType(v string) *DescribeChartDataResponseBody
	GetChartDataType() *string
	SetChartSubTypeList(v []*string) *DescribeChartDataResponseBody
	GetChartSubTypeList() []*string
	SetChartType(v string) *DescribeChartDataResponseBody
	GetChartType() *string
	SetCoordinateData(v *DescribeChartDataResponseBodyCoordinateData) *DescribeChartDataResponseBody
	GetCoordinateData() *DescribeChartDataResponseBodyCoordinateData
	SetMultipleData(v []*DescribeChartDataResponseBodyMultipleData) *DescribeChartDataResponseBody
	GetMultipleData() []*DescribeChartDataResponseBodyMultipleData
	SetPropertyArrayValue(v string) *DescribeChartDataResponseBody
	GetPropertyArrayValue() *string
	SetPropertyValue(v string) *DescribeChartDataResponseBody
	GetPropertyValue() *string
	SetProperyArrayValue(v string) *DescribeChartDataResponseBody
	GetProperyArrayValue() *string
	SetRequestId(v string) *DescribeChartDataResponseBody
	GetRequestId() *string
	SetSingleData(v *DescribeChartDataResponseBodySingleData) *DescribeChartDataResponseBody
	GetSingleData() *DescribeChartDataResponseBodySingleData
}

type DescribeChartDataResponseBody struct {
	// The valid values for all subtypes of the chart.
	AllChartSubTypeList []*DescribeChartDataResponseBodyAllChartSubTypeList `json:"AllChartSubTypeList,omitempty" xml:"AllChartSubTypeList,omitempty" type:"Repeated"`
	// The data type of the chart. Valid values:
	//
	// 	- **commonCoordinate**
	//
	// 	- **timeCoordinate**
	//
	// 	- **multipleValue**
	//
	// 	- **singleValue**
	//
	// 	- **propertyValue**
	//
	// 	- **propertyArrayValue**
	//
	// example:
	//
	// propertyArrayValue
	ChartDataType *string `json:"ChartDataType,omitempty" xml:"ChartDataType,omitempty"`
	// The subtype values in which the chart is selected.
	ChartSubTypeList []*string `json:"ChartSubTypeList,omitempty" xml:"ChartSubTypeList,omitempty" type:"Repeated"`
	// The type of the chart. Valid values:
	//
	// 	- **timeLine**
	//
	// 	- **timeBar**
	//
	// 	- **bar**
	//
	// 	- **line**
	//
	// 	- **pie**
	//
	// 	- **gauge**
	//
	// 	- **table**
	//
	// 	- **text**
	//
	// example:
	//
	// timeBar
	ChartType *string `json:"ChartType,omitempty" xml:"ChartType,omitempty"`
	// The coordinate data.
	CoordinateData *DescribeChartDataResponseBodyCoordinateData `json:"CoordinateData,omitempty" xml:"CoordinateData,omitempty" type:"Struct"`
	// The values in the multi-value charts.
	MultipleData []*DescribeChartDataResponseBodyMultipleData `json:"MultipleData,omitempty" xml:"MultipleData,omitempty" type:"Repeated"`
	// The attribute value of the array chart.
	//
	// example:
	//
	// []
	PropertyArrayValue *string `json:"PropertyArrayValue,omitempty" xml:"PropertyArrayValue,omitempty"`
	// The data of the chart.
	//
	// example:
	//
	// {\\"totalCount\\": \\"0\\"}
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The array data of the chart.
	//
	// example:
	//
	// []
	ProperyArrayValue *string `json:"ProperyArrayValue,omitempty" xml:"ProperyArrayValue,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 50CCE62A-2BC4-5CF8-B976-E4F62A31****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data of the single value chart.
	SingleData *DescribeChartDataResponseBodySingleData `json:"SingleData,omitempty" xml:"SingleData,omitempty" type:"Struct"`
}

func (s DescribeChartDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChartDataResponseBody) GetAllChartSubTypeList() []*DescribeChartDataResponseBodyAllChartSubTypeList {
	return s.AllChartSubTypeList
}

func (s *DescribeChartDataResponseBody) GetChartDataType() *string {
	return s.ChartDataType
}

func (s *DescribeChartDataResponseBody) GetChartSubTypeList() []*string {
	return s.ChartSubTypeList
}

func (s *DescribeChartDataResponseBody) GetChartType() *string {
	return s.ChartType
}

func (s *DescribeChartDataResponseBody) GetCoordinateData() *DescribeChartDataResponseBodyCoordinateData {
	return s.CoordinateData
}

func (s *DescribeChartDataResponseBody) GetMultipleData() []*DescribeChartDataResponseBodyMultipleData {
	return s.MultipleData
}

func (s *DescribeChartDataResponseBody) GetPropertyArrayValue() *string {
	return s.PropertyArrayValue
}

func (s *DescribeChartDataResponseBody) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *DescribeChartDataResponseBody) GetProperyArrayValue() *string {
	return s.ProperyArrayValue
}

func (s *DescribeChartDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChartDataResponseBody) GetSingleData() *DescribeChartDataResponseBodySingleData {
	return s.SingleData
}

func (s *DescribeChartDataResponseBody) SetAllChartSubTypeList(v []*DescribeChartDataResponseBodyAllChartSubTypeList) *DescribeChartDataResponseBody {
	s.AllChartSubTypeList = v
	return s
}

func (s *DescribeChartDataResponseBody) SetChartDataType(v string) *DescribeChartDataResponseBody {
	s.ChartDataType = &v
	return s
}

func (s *DescribeChartDataResponseBody) SetChartSubTypeList(v []*string) *DescribeChartDataResponseBody {
	s.ChartSubTypeList = v
	return s
}

func (s *DescribeChartDataResponseBody) SetChartType(v string) *DescribeChartDataResponseBody {
	s.ChartType = &v
	return s
}

func (s *DescribeChartDataResponseBody) SetCoordinateData(v *DescribeChartDataResponseBodyCoordinateData) *DescribeChartDataResponseBody {
	s.CoordinateData = v
	return s
}

func (s *DescribeChartDataResponseBody) SetMultipleData(v []*DescribeChartDataResponseBodyMultipleData) *DescribeChartDataResponseBody {
	s.MultipleData = v
	return s
}

func (s *DescribeChartDataResponseBody) SetPropertyArrayValue(v string) *DescribeChartDataResponseBody {
	s.PropertyArrayValue = &v
	return s
}

func (s *DescribeChartDataResponseBody) SetPropertyValue(v string) *DescribeChartDataResponseBody {
	s.PropertyValue = &v
	return s
}

func (s *DescribeChartDataResponseBody) SetProperyArrayValue(v string) *DescribeChartDataResponseBody {
	s.ProperyArrayValue = &v
	return s
}

func (s *DescribeChartDataResponseBody) SetRequestId(v string) *DescribeChartDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChartDataResponseBody) SetSingleData(v *DescribeChartDataResponseBodySingleData) *DescribeChartDataResponseBody {
	s.SingleData = v
	return s
}

func (s *DescribeChartDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChartDataResponseBodyAllChartSubTypeList struct {
	// The subtype of the chart.
	//
	// example:
	//
	// CID_SUSPICIOUS_TREND-ALL
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The name of the chart subtype.
	//
	// example:
	//
	// All Alerts
	SubTypeName *string `json:"SubTypeName,omitempty" xml:"SubTypeName,omitempty"`
}

func (s DescribeChartDataResponseBodyAllChartSubTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartDataResponseBodyAllChartSubTypeList) GoString() string {
	return s.String()
}

func (s *DescribeChartDataResponseBodyAllChartSubTypeList) GetSubType() *string {
	return s.SubType
}

func (s *DescribeChartDataResponseBodyAllChartSubTypeList) GetSubTypeName() *string {
	return s.SubTypeName
}

func (s *DescribeChartDataResponseBodyAllChartSubTypeList) SetSubType(v string) *DescribeChartDataResponseBodyAllChartSubTypeList {
	s.SubType = &v
	return s
}

func (s *DescribeChartDataResponseBodyAllChartSubTypeList) SetSubTypeName(v string) *DescribeChartDataResponseBodyAllChartSubTypeList {
	s.SubTypeName = &v
	return s
}

func (s *DescribeChartDataResponseBodyAllChartSubTypeList) Validate() error {
	return dara.Validate(s)
}

type DescribeChartDataResponseBodyCoordinateData struct {
	// The x-axis values.
	XAxis []*string `json:"XAxis,omitempty" xml:"XAxis,omitempty" type:"Repeated"`
	// The y-axis values.
	YAxisList []*DescribeChartDataResponseBodyCoordinateDataYAxisList `json:"YAxisList,omitempty" xml:"YAxisList,omitempty" type:"Repeated"`
}

func (s DescribeChartDataResponseBodyCoordinateData) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartDataResponseBodyCoordinateData) GoString() string {
	return s.String()
}

func (s *DescribeChartDataResponseBodyCoordinateData) GetXAxis() []*string {
	return s.XAxis
}

func (s *DescribeChartDataResponseBodyCoordinateData) GetYAxisList() []*DescribeChartDataResponseBodyCoordinateDataYAxisList {
	return s.YAxisList
}

func (s *DescribeChartDataResponseBodyCoordinateData) SetXAxis(v []*string) *DescribeChartDataResponseBodyCoordinateData {
	s.XAxis = v
	return s
}

func (s *DescribeChartDataResponseBodyCoordinateData) SetYAxisList(v []*DescribeChartDataResponseBodyCoordinateDataYAxisList) *DescribeChartDataResponseBodyCoordinateData {
	s.YAxisList = v
	return s
}

func (s *DescribeChartDataResponseBodyCoordinateData) Validate() error {
	return dara.Validate(s)
}

type DescribeChartDataResponseBodyCoordinateDataYAxisList struct {
	// The name of the data type.
	//
	// example:
	//
	// Port
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subtype data of the chart.
	//
	// example:
	//
	// CID_SUSPICIOUS_TREND-AL
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The type of the data.
	//
	// example:
	//
	// high
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The values of the y-axis that corresponds to x-axis points.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeChartDataResponseBodyCoordinateDataYAxisList) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartDataResponseBodyCoordinateDataYAxisList) GoString() string {
	return s.String()
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) GetName() *string {
	return s.Name
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) GetSubType() *string {
	return s.SubType
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) GetType() *string {
	return s.Type
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) GetValue() []*string {
	return s.Value
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) SetName(v string) *DescribeChartDataResponseBodyCoordinateDataYAxisList {
	s.Name = &v
	return s
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) SetSubType(v string) *DescribeChartDataResponseBodyCoordinateDataYAxisList {
	s.SubType = &v
	return s
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) SetType(v string) *DescribeChartDataResponseBodyCoordinateDataYAxisList {
	s.Type = &v
	return s
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) SetValue(v []*string) *DescribeChartDataResponseBodyCoordinateDataYAxisList {
	s.Value = v
	return s
}

func (s *DescribeChartDataResponseBodyCoordinateDataYAxisList) Validate() error {
	return dara.Validate(s)
}

type DescribeChartDataResponseBodyMultipleData struct {
	// The font color, which is an RGBA value.
	//
	// example:
	//
	// #FFA800
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The name of the data type.
	//
	// example:
	//
	// Safety
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data.
	//
	// example:
	//
	// safe
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// 0
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeChartDataResponseBodyMultipleData) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartDataResponseBodyMultipleData) GoString() string {
	return s.String()
}

func (s *DescribeChartDataResponseBodyMultipleData) GetColor() *string {
	return s.Color
}

func (s *DescribeChartDataResponseBodyMultipleData) GetName() *string {
	return s.Name
}

func (s *DescribeChartDataResponseBodyMultipleData) GetType() *string {
	return s.Type
}

func (s *DescribeChartDataResponseBodyMultipleData) GetValue() *int64 {
	return s.Value
}

func (s *DescribeChartDataResponseBodyMultipleData) SetColor(v string) *DescribeChartDataResponseBodyMultipleData {
	s.Color = &v
	return s
}

func (s *DescribeChartDataResponseBodyMultipleData) SetName(v string) *DescribeChartDataResponseBodyMultipleData {
	s.Name = &v
	return s
}

func (s *DescribeChartDataResponseBodyMultipleData) SetType(v string) *DescribeChartDataResponseBodyMultipleData {
	s.Type = &v
	return s
}

func (s *DescribeChartDataResponseBodyMultipleData) SetValue(v int64) *DescribeChartDataResponseBodyMultipleData {
	s.Value = &v
	return s
}

func (s *DescribeChartDataResponseBodyMultipleData) Validate() error {
	return dara.Validate(s)
}

type DescribeChartDataResponseBodySingleData struct {
	// The name of the data type.
	//
	// example:
	//
	// suspicious
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data.
	//
	// example:
	//
	// suspicious
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value in the single value chart.
	//
	// example:
	//
	// 172
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeChartDataResponseBodySingleData) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartDataResponseBodySingleData) GoString() string {
	return s.String()
}

func (s *DescribeChartDataResponseBodySingleData) GetName() *string {
	return s.Name
}

func (s *DescribeChartDataResponseBodySingleData) GetType() *string {
	return s.Type
}

func (s *DescribeChartDataResponseBodySingleData) GetValue() *int64 {
	return s.Value
}

func (s *DescribeChartDataResponseBodySingleData) SetName(v string) *DescribeChartDataResponseBodySingleData {
	s.Name = &v
	return s
}

func (s *DescribeChartDataResponseBodySingleData) SetType(v string) *DescribeChartDataResponseBodySingleData {
	s.Type = &v
	return s
}

func (s *DescribeChartDataResponseBodySingleData) SetValue(v int64) *DescribeChartDataResponseBodySingleData {
	s.Value = &v
	return s
}

func (s *DescribeChartDataResponseBodySingleData) Validate() error {
	return dara.Validate(s)
}
