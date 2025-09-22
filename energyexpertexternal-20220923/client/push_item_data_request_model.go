// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushItemDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PushItemDataRequest
	GetCode() *string
	SetItems(v *PushItemDataRequestItems) *PushItemDataRequest
	GetItems() *PushItemDataRequestItems
	SetYear(v string) *PushItemDataRequest
	GetYear() *string
}

type PushItemDataRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20210223-01
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// List of data to be pushed.
	//
	// This parameter is required.
	Items *PushItemDataRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Struct"`
	// The year of the data created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s PushItemDataRequest) String() string {
	return dara.Prettify(s)
}

func (s PushItemDataRequest) GoString() string {
	return s.String()
}

func (s *PushItemDataRequest) GetCode() *string {
	return s.Code
}

func (s *PushItemDataRequest) GetItems() *PushItemDataRequestItems {
	return s.Items
}

func (s *PushItemDataRequest) GetYear() *string {
	return s.Year
}

func (s *PushItemDataRequest) SetCode(v string) *PushItemDataRequest {
	s.Code = &v
	return s
}

func (s *PushItemDataRequest) SetItems(v *PushItemDataRequestItems) *PushItemDataRequest {
	s.Items = v
	return s
}

func (s *PushItemDataRequest) SetYear(v string) *PushItemDataRequest {
	s.Year = &v
	return s
}

func (s *PushItemDataRequest) Validate() error {
	return dara.Validate(s)
}

type PushItemDataRequestItems struct {
	// API data identification.<props="intl">For details: [GetDataItemList ](https://www.alibabacloud.com/help/en/energy-expert/developer-reference/api-energyexpertexternal-2022-09-23-getdataitemlist)
	//
	// This parameter is required.
	//
	// example:
	//
	// demo_api_code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Month *string `json:"month,omitempty" xml:"month,omitempty"`
	// The value of the data item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.11
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PushItemDataRequestItems) String() string {
	return dara.Prettify(s)
}

func (s PushItemDataRequestItems) GoString() string {
	return s.String()
}

func (s *PushItemDataRequestItems) GetCode() *string {
	return s.Code
}

func (s *PushItemDataRequestItems) GetMonth() *string {
	return s.Month
}

func (s *PushItemDataRequestItems) GetValue() *float64 {
	return s.Value
}

func (s *PushItemDataRequestItems) SetCode(v string) *PushItemDataRequestItems {
	s.Code = &v
	return s
}

func (s *PushItemDataRequestItems) SetMonth(v string) *PushItemDataRequestItems {
	s.Month = &v
	return s
}

func (s *PushItemDataRequestItems) SetValue(v float64) *PushItemDataRequestItems {
	s.Value = &v
	return s
}

func (s *PushItemDataRequestItems) Validate() error {
	return dara.Validate(s)
}
