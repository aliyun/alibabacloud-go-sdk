// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeSolutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeSolutionResponseBody
	GetCode() *string
	SetData(v *GetQueryOptimizeSolutionResponseBodyData) *GetQueryOptimizeSolutionResponseBody
	GetData() *GetQueryOptimizeSolutionResponseBodyData
	SetMessage(v string) *GetQueryOptimizeSolutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeSolutionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeSolutionResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeSolutionResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetQueryOptimizeSolutionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 30FF4E40-17F3-5A51-AB23-43F30D9B****
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

func (s GetQueryOptimizeSolutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeSolutionResponseBody) GetData() *GetQueryOptimizeSolutionResponseBodyData {
	return s.Data
}

func (s *GetQueryOptimizeSolutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeSolutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeSolutionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeSolutionResponseBody) SetCode(v string) *GetQueryOptimizeSolutionResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) SetData(v *GetQueryOptimizeSolutionResponseBodyData) *GetQueryOptimizeSolutionResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) SetMessage(v string) *GetQueryOptimizeSolutionResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) SetRequestId(v string) *GetQueryOptimizeSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) SetSuccess(v string) *GetQueryOptimizeSolutionResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQueryOptimizeSolutionResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The optimization suggestions.
	List []*GetQueryOptimizeSolutionResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeSolutionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeSolutionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetQueryOptimizeSolutionResponseBodyData) GetList() []*GetQueryOptimizeSolutionResponseBodyDataList {
	return s.List
}

func (s *GetQueryOptimizeSolutionResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQueryOptimizeSolutionResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueryOptimizeSolutionResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetExtra(v string) *GetQueryOptimizeSolutionResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetList(v []*GetQueryOptimizeSolutionResponseBodyDataList) *GetQueryOptimizeSolutionResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeSolutionResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeSolutionResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetTotal(v int64) *GetQueryOptimizeSolutionResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetQueryOptimizeSolutionResponseBodyDataList struct {
	// The severity level. Valid values:
	//
	// 	- **INFO**
	//
	// 	- **WARN**
	//
	// example:
	//
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// LARGE_ROWS_EXAMINED
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The suggestion.
	//
	// example:
	//
	// LARGE_ROWS_EXAMINED_SOLUTION
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	SolutionExt *string `json:"SolutionExt,omitempty" xml:"SolutionExt,omitempty"`
}

func (s GetQueryOptimizeSolutionResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeSolutionResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) GetLevel() *string {
	return s.Level
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) GetRuleId() *string {
	return s.RuleId
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) GetSolution() *string {
	return s.Solution
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) GetSolutionExt() *string {
	return s.SolutionExt
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) SetLevel(v string) *GetQueryOptimizeSolutionResponseBodyDataList {
	s.Level = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) SetRuleId(v string) *GetQueryOptimizeSolutionResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) SetSolution(v string) *GetQueryOptimizeSolutionResponseBodyDataList {
	s.Solution = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) SetSolutionExt(v string) *GetQueryOptimizeSolutionResponseBodyDataList {
	s.SolutionExt = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
