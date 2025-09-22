// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCarbonEmissionTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCarbonEmissionTrendResponseBodyData) *GetCarbonEmissionTrendResponseBody
	GetData() *GetCarbonEmissionTrendResponseBodyData
	SetRequestId(v string) *GetCarbonEmissionTrendResponseBody
	GetRequestId() *string
}

type GetCarbonEmissionTrendResponseBody struct {
	// The response parameters.
	Data *GetCarbonEmissionTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request.
	//
	// example:
	//
	// 9bc20a5a-b26b-4c28-922a-7cd10b61f96f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBody) GetData() *GetCarbonEmissionTrendResponseBodyData {
	return s.Data
}

func (s *GetCarbonEmissionTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCarbonEmissionTrendResponseBody) SetData(v *GetCarbonEmissionTrendResponseBodyData) *GetCarbonEmissionTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBody) SetRequestId(v string) *GetCarbonEmissionTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCarbonEmissionTrendResponseBodyData struct {
	// Actual emission list.
	ActualEmissionList []*GetCarbonEmissionTrendResponseBodyDataActualEmissionList `json:"actualEmissionList,omitempty" xml:"actualEmissionList,omitempty" type:"Repeated"`
	// Target Emission List.
	TargetEmissionList []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionList `json:"targetEmissionList,omitempty" xml:"targetEmissionList,omitempty" type:"Repeated"`
}

func (s GetCarbonEmissionTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyData) GetActualEmissionList() []*GetCarbonEmissionTrendResponseBodyDataActualEmissionList {
	return s.ActualEmissionList
}

func (s *GetCarbonEmissionTrendResponseBodyData) GetTargetEmissionList() []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionList {
	return s.TargetEmissionList
}

func (s *GetCarbonEmissionTrendResponseBodyData) SetActualEmissionList(v []*GetCarbonEmissionTrendResponseBodyDataActualEmissionList) *GetCarbonEmissionTrendResponseBodyData {
	s.ActualEmissionList = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyData) SetTargetEmissionList(v []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) *GetCarbonEmissionTrendResponseBodyData {
	s.TargetEmissionList = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetCarbonEmissionTrendResponseBodyDataActualEmissionList struct {
	// Data item list.
	Items []*GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The year.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBodyDataActualEmissionList) String() string {
	return dara.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyDataActualEmissionList) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionList) GetItems() []*GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems {
	return s.Items
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionList) GetYear() *string {
	return s.Year
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionList) SetItems(v []*GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) *GetCarbonEmissionTrendResponseBodyDataActualEmissionList {
	s.Items = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionList) SetYear(v string) *GetCarbonEmissionTrendResponseBodyDataActualEmissionList {
	s.Year = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionList) Validate() error {
	return dara.Validate(s)
}

type GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems struct {
	// Carbon emissions.
	//
	// example:
	//
	// 20.22
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The month.
	//
	// example:
	//
	// 11
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// The year.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) String() string {
	return dara.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) GetMonth() *int32 {
	return s.Month
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) GetYear() *string {
	return s.Year
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) SetCarbonEmissionData(v float64) *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) SetMonth(v int32) *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems {
	s.Month = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) SetYear(v string) *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems {
	s.Year = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) Validate() error {
	return dara.Validate(s)
}

type GetCarbonEmissionTrendResponseBodyDataTargetEmissionList struct {
	// Data item list.
	Items []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The year.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) String() string {
	return dara.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) GetItems() []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems {
	return s.Items
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) GetYear() *string {
	return s.Year
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) SetItems(v []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList {
	s.Items = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) SetYear(v string) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList {
	s.Year = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) Validate() error {
	return dara.Validate(s)
}

type GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems struct {
	// Carbon emissions.
	//
	// example:
	//
	// 20.22
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The month.
	//
	// example:
	//
	// 10
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// The year.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) String() string {
	return dara.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) GetMonth() *int32 {
	return s.Month
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) GetYear() *string {
	return s.Year
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) SetCarbonEmissionData(v float64) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) SetMonth(v int32) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems {
	s.Month = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) SetYear(v string) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems {
	s.Year = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) Validate() error {
	return dara.Validate(s)
}
