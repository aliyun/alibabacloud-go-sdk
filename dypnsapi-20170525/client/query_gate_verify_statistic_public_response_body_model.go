// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyStatisticPublicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryGateVerifyStatisticPublicResponseBody
	GetCode() *string
	SetData(v *QueryGateVerifyStatisticPublicResponseBodyData) *QueryGateVerifyStatisticPublicResponseBody
	GetData() *QueryGateVerifyStatisticPublicResponseBodyData
	SetMessage(v string) *QueryGateVerifyStatisticPublicResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryGateVerifyStatisticPublicResponseBody
	GetRequestId() *string
}

type QueryGateVerifyStatisticPublicResponseBody struct {
	// The response code. Valid values:
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the calls of Phone Number Verification Service, including the total calls, the successful calls, failed calls, unknown calls, and daily calls within the statistical date range.
	Data *QueryGateVerifyStatisticPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryGateVerifyStatisticPublicResponseBody) GetData() *QueryGateVerifyStatisticPublicResponseBodyData {
	return s.Data
}

func (s *QueryGateVerifyStatisticPublicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryGateVerifyStatisticPublicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetCode(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetData(v *QueryGateVerifyStatisticPublicResponseBodyData) *QueryGateVerifyStatisticPublicResponseBody {
	s.Data = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetMessage(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetRequestId(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryGateVerifyStatisticPublicResponseBodyData struct {
	// The information about the daily calls.
	DayStatistic []*QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic `json:"DayStatistic,omitempty" xml:"DayStatistic,omitempty" type:"Repeated"`
	// The total calls.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The failed calls.
	//
	// example:
	//
	// 20
	TotalFail *int64 `json:"TotalFail,omitempty" xml:"TotalFail,omitempty"`
	// The successful calls.
	//
	// example:
	//
	// 0
	TotalSuccess *int64 `json:"TotalSuccess,omitempty" xml:"TotalSuccess,omitempty"`
	// The unknown calls.
	//
	// example:
	//
	// 0
	TotalUnknown *int64 `json:"TotalUnknown,omitempty" xml:"TotalUnknown,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) GetDayStatistic() []*QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	return s.DayStatistic
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) GetTotalFail() *int64 {
	return s.TotalFail
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) GetTotalSuccess() *int64 {
	return s.TotalSuccess
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) GetTotalUnknown() *int64 {
	return s.TotalUnknown
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetDayStatistic(v []*QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.DayStatistic = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotal(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalFail(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalFail = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalSuccess(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalSuccess = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalUnknown(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalUnknown = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic struct {
	// The date. This field is accurate to the day. The value of this field is in the YYYYMMDD format. Example: 20220103.
	//
	// example:
	//
	// 20220103
	StatisticDateStr *string `json:"StatisticDateStr,omitempty" xml:"StatisticDateStr,omitempty"`
	// The failed calls on the day.
	//
	// example:
	//
	// 20
	TotalFail *int64 `json:"TotalFail,omitempty" xml:"TotalFail,omitempty"`
	// The successful calls on the day.
	//
	// example:
	//
	// 0
	TotalSuccess *int64 `json:"TotalSuccess,omitempty" xml:"TotalSuccess,omitempty"`
	// The unknown calls on the day.
	//
	// example:
	//
	// 0
	TotalUnknown *int64 `json:"TotalUnknown,omitempty" xml:"TotalUnknown,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) GetStatisticDateStr() *string {
	return s.StatisticDateStr
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) GetTotalFail() *int64 {
	return s.TotalFail
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) GetTotalSuccess() *int64 {
	return s.TotalSuccess
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) GetTotalUnknown() *int64 {
	return s.TotalUnknown
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetStatisticDateStr(v string) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.StatisticDateStr = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalFail(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalFail = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalSuccess(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalSuccess = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalUnknown(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalUnknown = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) Validate() error {
	return dara.Validate(s)
}
