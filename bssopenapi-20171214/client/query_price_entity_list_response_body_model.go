// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPriceEntityListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPriceEntityListResponseBody
	GetCode() *string
	SetData(v *QueryPriceEntityListResponseBodyData) *QueryPriceEntityListResponseBody
	GetData() *QueryPriceEntityListResponseBodyData
	SetMessage(v string) *QueryPriceEntityListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPriceEntityListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPriceEntityListResponseBody
	GetSuccess() *bool
}

type QueryPriceEntityListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *QueryPriceEntityListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPriceEntityListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceEntityListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPriceEntityListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPriceEntityListResponseBody) GetData() *QueryPriceEntityListResponseBodyData {
	return s.Data
}

func (s *QueryPriceEntityListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPriceEntityListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPriceEntityListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPriceEntityListResponseBody) SetCode(v string) *QueryPriceEntityListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPriceEntityListResponseBody) SetData(v *QueryPriceEntityListResponseBodyData) *QueryPriceEntityListResponseBody {
	s.Data = v
	return s
}

func (s *QueryPriceEntityListResponseBody) SetMessage(v string) *QueryPriceEntityListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPriceEntityListResponseBody) SetRequestId(v string) *QueryPriceEntityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPriceEntityListResponseBody) SetSuccess(v bool) *QueryPriceEntityListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPriceEntityListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPriceEntityListResponseBodyData struct {
	// The information about the billable items.
	PriceEntityInfoList []*QueryPriceEntityListResponseBodyDataPriceEntityInfoList `json:"PriceEntityInfoList,omitempty" xml:"PriceEntityInfoList,omitempty" type:"Repeated"`
}

func (s QueryPriceEntityListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceEntityListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPriceEntityListResponseBodyData) GetPriceEntityInfoList() []*QueryPriceEntityListResponseBodyDataPriceEntityInfoList {
	return s.PriceEntityInfoList
}

func (s *QueryPriceEntityListResponseBodyData) SetPriceEntityInfoList(v []*QueryPriceEntityListResponseBodyDataPriceEntityInfoList) *QueryPriceEntityListResponseBodyData {
	s.PriceEntityInfoList = v
	return s
}

func (s *QueryPriceEntityListResponseBodyData) Validate() error {
	if s.PriceEntityInfoList != nil {
		for _, item := range s.PriceEntityInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPriceEntityListResponseBodyDataPriceEntityInfoList struct {
	// The code of the billable item.
	//
	// example:
	//
	// instance_type
	PriceEntityCode *string `json:"PriceEntityCode,omitempty" xml:"PriceEntityCode,omitempty"`
	// The name of the billable item.
	//
	// example:
	//
	// Elastic Compute Service (ECS) instance
	PriceEntityName *string `json:"PriceEntityName,omitempty" xml:"PriceEntityName,omitempty"`
	// The factors of the billable item.
	PriceFactorList []*QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList `json:"PriceFactorList,omitempty" xml:"PriceFactorList,omitempty" type:"Repeated"`
}

func (s QueryPriceEntityListResponseBodyDataPriceEntityInfoList) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceEntityListResponseBodyDataPriceEntityInfoList) GoString() string {
	return s.String()
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoList) GetPriceEntityCode() *string {
	return s.PriceEntityCode
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoList) GetPriceEntityName() *string {
	return s.PriceEntityName
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoList) GetPriceFactorList() []*QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList {
	return s.PriceFactorList
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoList) SetPriceEntityCode(v string) *QueryPriceEntityListResponseBodyDataPriceEntityInfoList {
	s.PriceEntityCode = &v
	return s
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoList) SetPriceEntityName(v string) *QueryPriceEntityListResponseBodyDataPriceEntityInfoList {
	s.PriceEntityName = &v
	return s
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoList) SetPriceFactorList(v []*QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) *QueryPriceEntityListResponseBodyDataPriceEntityInfoList {
	s.PriceFactorList = v
	return s
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoList) Validate() error {
	if s.PriceFactorList != nil {
		for _, item := range s.PriceFactorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList struct {
	// The code of the factor.
	//
	// example:
	//
	// vm_region_no
	PriceFactorCode *string `json:"PriceFactorCode,omitempty" xml:"PriceFactorCode,omitempty"`
	// The name of the factor.
	//
	// example:
	//
	// Region
	PriceFactorName *string `json:"PriceFactorName,omitempty" xml:"PriceFactorName,omitempty"`
	// The values of the factor.
	PriceFactorValueList []*string `json:"PriceFactorValueList,omitempty" xml:"PriceFactorValueList,omitempty" type:"Repeated"`
}

func (s QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) GoString() string {
	return s.String()
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) GetPriceFactorCode() *string {
	return s.PriceFactorCode
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) GetPriceFactorName() *string {
	return s.PriceFactorName
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) GetPriceFactorValueList() []*string {
	return s.PriceFactorValueList
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) SetPriceFactorCode(v string) *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList {
	s.PriceFactorCode = &v
	return s
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) SetPriceFactorName(v string) *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList {
	s.PriceFactorName = &v
	return s
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) SetPriceFactorValueList(v []*string) *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList {
	s.PriceFactorValueList = v
	return s
}

func (s *QueryPriceEntityListResponseBodyDataPriceEntityInfoListPriceFactorList) Validate() error {
	return dara.Validate(s)
}
