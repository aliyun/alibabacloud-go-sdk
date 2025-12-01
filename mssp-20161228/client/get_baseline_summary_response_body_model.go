// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBaselineSummaryResponseBody
	GetCode() *string
	SetData(v *GetBaselineSummaryResponseBodyData) *GetBaselineSummaryResponseBody
	GetData() *GetBaselineSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetBaselineSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBaselineSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBaselineSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBaselineSummaryResponseBody
	GetSuccess() *bool
}

type GetBaselineSummaryResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetBaselineSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 67D61738-5E38-5164-947A-34E3850D493A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Values: true: success; false: failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBaselineSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBaselineSummaryResponseBody) GetData() *GetBaselineSummaryResponseBodyData {
	return s.Data
}

func (s *GetBaselineSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBaselineSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBaselineSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBaselineSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBaselineSummaryResponseBody) SetCode(v string) *GetBaselineSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetData(v *GetBaselineSummaryResponseBodyData) *GetBaselineSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetHttpStatusCode(v int32) *GetBaselineSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetMessage(v string) *GetBaselineSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetRequestId(v string) *GetBaselineSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetSuccess(v bool) *GetBaselineSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBaselineSummaryResponseBodyData struct {
	// Collection of baseline statistical data.
	TrendDTOList []*GetBaselineSummaryResponseBodyDataTrendDTOList `json:"TrendDTOList,omitempty" xml:"TrendDTOList,omitempty" type:"Repeated"`
}

func (s GetBaselineSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBodyData) GetTrendDTOList() []*GetBaselineSummaryResponseBodyDataTrendDTOList {
	return s.TrendDTOList
}

func (s *GetBaselineSummaryResponseBodyData) SetTrendDTOList(v []*GetBaselineSummaryResponseBodyDataTrendDTOList) *GetBaselineSummaryResponseBodyData {
	s.TrendDTOList = v
	return s
}

func (s *GetBaselineSummaryResponseBodyData) Validate() error {
	if s.TrendDTOList != nil {
		for _, item := range s.TrendDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBaselineSummaryResponseBodyDataTrendDTOList struct {
	// Date point.
	//
	// example:
	//
	// 202408或者20240801
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Number of processed items.
	//
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// Number of discovered items.
	//
	// example:
	//
	// 12
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetBaselineSummaryResponseBodyDataTrendDTOList) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineSummaryResponseBodyDataTrendDTOList) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) GetDate() *string {
	return s.Date
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) GetDealCount() *int64 {
	return s.DealCount
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) GetFindCount() *int64 {
	return s.FindCount
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetDate(v string) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.Date = &v
	return s
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetDealCount(v int64) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.DealCount = &v
	return s
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetFindCount(v int64) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.FindCount = &v
	return s
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) Validate() error {
	return dara.Validate(s)
}
