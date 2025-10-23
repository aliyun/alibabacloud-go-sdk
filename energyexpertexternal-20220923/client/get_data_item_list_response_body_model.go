// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataItemListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetDataItemListResponseBodyData) *GetDataItemListResponseBody
	GetData() []*GetDataItemListResponseBodyData
	SetRequestId(v string) *GetDataItemListResponseBody
	GetRequestId() *string
}

type GetDataItemListResponseBody struct {
	// Response parameters.
	Data []*GetDataItemListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDataItemListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataItemListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataItemListResponseBody) GetData() []*GetDataItemListResponseBodyData {
	return s.Data
}

func (s *GetDataItemListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataItemListResponseBody) SetData(v []*GetDataItemListResponseBodyData) *GetDataItemListResponseBody {
	s.Data = v
	return s
}

func (s *GetDataItemListResponseBody) SetRequestId(v string) *GetDataItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataItemListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataItemListResponseBodyData struct {
	// The identifier of the data item.
	//
	// example:
	//
	// demo_api_code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The name of the data item.
	//
	// example:
	//
	// name_bbb
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Data filling method, 1: monthly value 2: annual value.
	//
	// example:
	//
	// 1
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The data item unit.
	//
	// example:
	//
	// kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetDataItemListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataItemListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataItemListResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetDataItemListResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetDataItemListResponseBodyData) GetPeriod() *int32 {
	return s.Period
}

func (s *GetDataItemListResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *GetDataItemListResponseBodyData) SetCode(v string) *GetDataItemListResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDataItemListResponseBodyData) SetName(v string) *GetDataItemListResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDataItemListResponseBodyData) SetPeriod(v int32) *GetDataItemListResponseBodyData {
	s.Period = &v
	return s
}

func (s *GetDataItemListResponseBodyData) SetUnit(v string) *GetDataItemListResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetDataItemListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
