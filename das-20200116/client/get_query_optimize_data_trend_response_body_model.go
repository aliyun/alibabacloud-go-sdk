// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeDataTrendResponseBody
	GetCode() *string
	SetData(v *GetQueryOptimizeDataTrendResponseBodyData) *GetQueryOptimizeDataTrendResponseBody
	GetData() *GetQueryOptimizeDataTrendResponseBodyData
	SetMessage(v string) *GetQueryOptimizeDataTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeDataTrendResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeDataTrendResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeDataTrendResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *GetQueryOptimizeDataTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeDataTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeDataTrendResponseBody) GetData() *GetQueryOptimizeDataTrendResponseBodyData {
	return s.Data
}

func (s *GetQueryOptimizeDataTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeDataTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeDataTrendResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetCode(v string) *GetQueryOptimizeDataTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetData(v *GetQueryOptimizeDataTrendResponseBodyData) *GetQueryOptimizeDataTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetMessage(v string) *GetQueryOptimizeDataTrendResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetRequestId(v string) *GetQueryOptimizeDataTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetSuccess(v string) *GetQueryOptimizeDataTrendResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueryOptimizeDataTrendResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The details of the trend data.
	List []*GetQueryOptimizeDataTrendResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeDataTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) GetList() []*GetQueryOptimizeDataTrendResponseBodyDataList {
	return s.List
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetExtra(v string) *GetQueryOptimizeDataTrendResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetList(v []*GetQueryOptimizeDataTrendResponseBodyDataList) *GetQueryOptimizeDataTrendResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeDataTrendResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeDataTrendResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetTotal(v int64) *GetQueryOptimizeDataTrendResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQueryOptimizeDataTrendResponseBodyDataList struct {
	// The name of the metric. Valid values:
	//
	// 	- **sqlExecuteCount**: the number of executions of slow SQL queries.
	//
	// 	- **sqlExecuteCountDiff**: the difference in the number of executions of slow SQL queries compared to the previous day.
	//
	// 	- **sqlCount**: the number of slow SQL templates.
	//
	// 	- **sqlCountDiff**: the difference in the number of slow SQL templates compared to the previous day.
	//
	// 	- **optimizedSqlExecuteCount**: the number of optimizable executions of slow SQL queries.
	//
	// 	- **optimizedSqlExecuteCountDiff**: the difference in the number of optimizable executions of slow SQL queries compared to the previous day.
	//
	// 	- **optimizedSqlCount**: the number of optimizable slow SQL templates.
	//
	// 	- **optimizedSqlCountDiff**: the difference in the number of optimizable slow SQL templates compared to the previous day.
	//
	// example:
	//
	// sqlExecuteCount
	Kpi *string `json:"Kpi,omitempty" xml:"Kpi,omitempty"`
	// The data timestamp. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1643040000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetQueryOptimizeDataTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) GetKpi() *string {
	return s.Kpi
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) GetValue() *float64 {
	return s.Value
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) SetKpi(v string) *GetQueryOptimizeDataTrendResponseBodyDataList {
	s.Kpi = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) SetTimestamp(v int64) *GetQueryOptimizeDataTrendResponseBodyDataList {
	s.Timestamp = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) SetValue(v float64) *GetQueryOptimizeDataTrendResponseBodyDataList {
	s.Value = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
