// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkTaskSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkTaskSummaryResponseBody
	GetCode() *string
	SetData(v *GetWorkTaskSummaryResponseBodyData) *GetWorkTaskSummaryResponseBody
	GetData() *GetWorkTaskSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkTaskSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkTaskSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkTaskSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkTaskSummaryResponseBody
	GetSuccess() *bool
}

type GetWorkTaskSummaryResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetWorkTaskSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Successful!
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

func (s GetWorkTaskSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkTaskSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkTaskSummaryResponseBody) GetData() *GetWorkTaskSummaryResponseBodyData {
	return s.Data
}

func (s *GetWorkTaskSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkTaskSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkTaskSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkTaskSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkTaskSummaryResponseBody) SetCode(v string) *GetWorkTaskSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetData(v *GetWorkTaskSummaryResponseBodyData) *GetWorkTaskSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetHttpStatusCode(v int32) *GetWorkTaskSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetMessage(v string) *GetWorkTaskSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetRequestId(v string) *GetWorkTaskSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetSuccess(v bool) *GetWorkTaskSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkTaskSummaryResponseBodyData struct {
	// Average response time (in minutes).
	//
	// example:
	//
	// 60
	DealAverageDuration *int64 `json:"DealAverageDuration,omitempty" xml:"DealAverageDuration,omitempty"`
	// Year-over-year growth rate of average response time.
	//
	// example:
	//
	// 20
	DealAverageDurationGrowthRate *string `json:"DealAverageDurationGrowthRate,omitempty" xml:"DealAverageDurationGrowthRate,omitempty"`
	// Number of work orders responded to.
	//
	// example:
	//
	// 100
	DealWorkTaskCount *int64 `json:"DealWorkTaskCount,omitempty" xml:"DealWorkTaskCount,omitempty"`
	// Year-over-year growth rate of the number of work orders responded to.
	//
	// example:
	//
	// 20
	DealWorkTaskCountRate *string `json:"DealWorkTaskCountRate,omitempty" xml:"DealWorkTaskCountRate,omitempty"`
	// Number of service responses.
	//
	// example:
	//
	// 10
	WorkTaskCount *int64 `json:"WorkTaskCount,omitempty" xml:"WorkTaskCount,omitempty"`
	// Problem closure rate.
	//
	// example:
	//
	// 90
	WorkTaskDealRate *string `json:"WorkTaskDealRate,omitempty" xml:"WorkTaskDealRate,omitempty"`
	// Year-over-year growth rate of problem closure rate.
	//
	// example:
	//
	// 20
	WorkTaskDealRateGrowthRate *string `json:"WorkTaskDealRateGrowthRate,omitempty" xml:"WorkTaskDealRateGrowthRate,omitempty"`
	// Year-over-year growth rate of service responses.
	//
	// example:
	//
	// 20
	WorkTaskGrowthRate *string `json:"WorkTaskGrowthRate,omitempty" xml:"WorkTaskGrowthRate,omitempty"`
}

func (s GetWorkTaskSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkTaskSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponseBodyData) GetDealAverageDuration() *int64 {
	return s.DealAverageDuration
}

func (s *GetWorkTaskSummaryResponseBodyData) GetDealAverageDurationGrowthRate() *string {
	return s.DealAverageDurationGrowthRate
}

func (s *GetWorkTaskSummaryResponseBodyData) GetDealWorkTaskCount() *int64 {
	return s.DealWorkTaskCount
}

func (s *GetWorkTaskSummaryResponseBodyData) GetDealWorkTaskCountRate() *string {
	return s.DealWorkTaskCountRate
}

func (s *GetWorkTaskSummaryResponseBodyData) GetWorkTaskCount() *int64 {
	return s.WorkTaskCount
}

func (s *GetWorkTaskSummaryResponseBodyData) GetWorkTaskDealRate() *string {
	return s.WorkTaskDealRate
}

func (s *GetWorkTaskSummaryResponseBodyData) GetWorkTaskDealRateGrowthRate() *string {
	return s.WorkTaskDealRateGrowthRate
}

func (s *GetWorkTaskSummaryResponseBodyData) GetWorkTaskGrowthRate() *string {
	return s.WorkTaskGrowthRate
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealAverageDuration(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.DealAverageDuration = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealAverageDurationGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.DealAverageDurationGrowthRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealWorkTaskCount(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.DealWorkTaskCount = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealWorkTaskCountRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.DealWorkTaskCountRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskCount(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskCount = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskDealRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskDealRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskDealRateGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskDealRateGrowthRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskGrowthRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
