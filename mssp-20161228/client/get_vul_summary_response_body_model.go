// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVulSummaryResponseBody
	GetCode() *string
	SetData(v *GetVulSummaryResponseBodyData) *GetVulSummaryResponseBody
	GetData() *GetVulSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetVulSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVulSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVulSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVulSummaryResponseBody
	GetSuccess() *bool
}

type GetVulSummaryResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetVulSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the response result.
	//
	// example:
	//
	// system error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVulSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVulSummaryResponseBody) GetData() *GetVulSummaryResponseBodyData {
	return s.Data
}

func (s *GetVulSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVulSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVulSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVulSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVulSummaryResponseBody) SetCode(v string) *GetVulSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetData(v *GetVulSummaryResponseBodyData) *GetVulSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetVulSummaryResponseBody) SetHttpStatusCode(v int32) *GetVulSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetMessage(v string) *GetVulSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetRequestId(v string) *GetVulSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetSuccess(v bool) *GetVulSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetVulSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVulSummaryResponseBodyData struct {
	// Number of completed items.
	//
	// example:
	//
	// 10
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// Risk convergence rate.
	//
	// example:
	//
	// 50
	DealRate *string `json:"DealRate,omitempty" xml:"DealRate,omitempty"`
	// Collection of vulnerability trend nodes.
	TrendList []*GetVulSummaryResponseBodyDataTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
	// Number of unhandled items.
	//
	// example:
	//
	// 5
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetVulSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVulSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBodyData) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *GetVulSummaryResponseBodyData) GetDealRate() *string {
	return s.DealRate
}

func (s *GetVulSummaryResponseBodyData) GetTrendList() []*GetVulSummaryResponseBodyDataTrendList {
	return s.TrendList
}

func (s *GetVulSummaryResponseBodyData) GetWaitHandleCount() *int64 {
	return s.WaitHandleCount
}

func (s *GetVulSummaryResponseBodyData) SetCompletedCount(v int64) *GetVulSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetDealRate(v string) *GetVulSummaryResponseBodyData {
	s.DealRate = &v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetTrendList(v []*GetVulSummaryResponseBodyDataTrendList) *GetVulSummaryResponseBodyData {
	s.TrendList = v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetVulSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

func (s *GetVulSummaryResponseBodyData) Validate() error {
	if s.TrendList != nil {
		for _, item := range s.TrendList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVulSummaryResponseBodyDataTrendList struct {
	// Time point.
	//
	// example:
	//
	// 202407或者20240701
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Number of handled items.
	//
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// Number of discovered items.
	//
	// example:
	//
	// 15
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetVulSummaryResponseBodyDataTrendList) String() string {
	return dara.Prettify(s)
}

func (s GetVulSummaryResponseBodyDataTrendList) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBodyDataTrendList) GetDate() *string {
	return s.Date
}

func (s *GetVulSummaryResponseBodyDataTrendList) GetDealCount() *int64 {
	return s.DealCount
}

func (s *GetVulSummaryResponseBodyDataTrendList) GetFindCount() *int64 {
	return s.FindCount
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetDate(v string) *GetVulSummaryResponseBodyDataTrendList {
	s.Date = &v
	return s
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetDealCount(v int64) *GetVulSummaryResponseBodyDataTrendList {
	s.DealCount = &v
	return s
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetFindCount(v int64) *GetVulSummaryResponseBodyDataTrendList {
	s.FindCount = &v
	return s
}

func (s *GetVulSummaryResponseBodyDataTrendList) Validate() error {
	return dara.Validate(s)
}
