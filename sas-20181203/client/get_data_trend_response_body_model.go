// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataTrendResponseBodyData) *GetDataTrendResponseBody
	GetData() *GetDataTrendResponseBodyData
	SetRequestId(v string) *GetDataTrendResponseBody
	GetRequestId() *string
}

type GetDataTrendResponseBody struct {
	// The response parameters.
	Data *GetDataTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C8E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataTrendResponseBody) GetData() *GetDataTrendResponseBodyData {
	return s.Data
}

func (s *GetDataTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataTrendResponseBody) SetData(v *GetDataTrendResponseBodyData) *GetDataTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetDataTrendResponseBody) SetRequestId(v string) *GetDataTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataTrendResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataTrendResponseBodyData struct {
	// The statistical timestamps of the trend data.
	DateList []*int32 `json:"DateList,omitempty" xml:"DateList,omitempty" type:"Repeated"`
	// The statistical dates and time for the trend data.
	DateStrList []*string `json:"DateStrList,omitempty" xml:"DateStrList,omitempty" type:"Repeated"`
	// The returned data.
	ItemList []*GetDataTrendResponseBodyDataItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s GetDataTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataTrendResponseBodyData) GetDateList() []*int32 {
	return s.DateList
}

func (s *GetDataTrendResponseBodyData) GetDateStrList() []*string {
	return s.DateStrList
}

func (s *GetDataTrendResponseBodyData) GetItemList() []*GetDataTrendResponseBodyDataItemList {
	return s.ItemList
}

func (s *GetDataTrendResponseBodyData) SetDateList(v []*int32) *GetDataTrendResponseBodyData {
	s.DateList = v
	return s
}

func (s *GetDataTrendResponseBodyData) SetDateStrList(v []*string) *GetDataTrendResponseBodyData {
	s.DateStrList = v
	return s
}

func (s *GetDataTrendResponseBodyData) SetItemList(v []*GetDataTrendResponseBodyDataItemList) *GetDataTrendResponseBodyData {
	s.ItemList = v
	return s
}

func (s *GetDataTrendResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDataTrendResponseBodyDataItemList struct {
	// The statistical values of the trend data.
	CountList []*int64 `json:"CountList,omitempty" xml:"CountList,omitempty" type:"Repeated"`
	// The type of the security data. Valid values:
	//
	// 	- **HC_NEW**: the number of new baseline risks.
	//
	// 	- **HC_OPERATE**: the number of handled baseline risks.
	//
	// 	- **VUL_NEW**: the number of new vulnerabilities.
	//
	// 	- **VUL_OPERATE**: the number of handled vulnerabilities.
	//
	// 	- **SUSP_NEW**: the number of new alerts.
	//
	// 	- **SUSP_OPERATE**: the number of handled alerts.
	//
	// example:
	//
	// HC_NEW
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
}

func (s GetDataTrendResponseBodyDataItemList) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrendResponseBodyDataItemList) GoString() string {
	return s.String()
}

func (s *GetDataTrendResponseBodyDataItemList) GetCountList() []*int64 {
	return s.CountList
}

func (s *GetDataTrendResponseBodyDataItemList) GetKeyName() *string {
	return s.KeyName
}

func (s *GetDataTrendResponseBodyDataItemList) SetCountList(v []*int64) *GetDataTrendResponseBodyDataItemList {
	s.CountList = v
	return s
}

func (s *GetDataTrendResponseBodyDataItemList) SetKeyName(v string) *GetDataTrendResponseBodyDataItemList {
	s.KeyName = &v
	return s
}

func (s *GetDataTrendResponseBodyDataItemList) Validate() error {
	return dara.Validate(s)
}
