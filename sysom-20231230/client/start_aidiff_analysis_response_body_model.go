// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIDiffAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartAIDiffAnalysisResponseBody
	GetCode() *string
	SetData(v string) *StartAIDiffAnalysisResponseBody
	GetData() *string
	SetMessage(v string) *StartAIDiffAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartAIDiffAnalysisResponseBody
	GetRequestId() *string
}

type StartAIDiffAnalysisResponseBody struct {
	// Status code
	//
	// - `code == Success` indicates that authorization succeeded.
	//
	// - Other status codes indicate that authorization failed. When authorization fails, view the `message` field to obtain detailed error information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// AI differential analysis result data
	//
	// example:
	//
	// "[
	//
	//   {
	//
	//       name: "xxx", #算子名称
	//
	//       before_time: 2, # 前者总耗时
	//
	//       after_time: 4, # 后者总耗时
	//
	//       time_diff: 2,  # 耗时差异
	//
	//       before_time_perc: "80%", # 前者总耗时百分比
	//
	//       after_time_perc: "23%", # 后者总耗时百分比
	//
	//       time_perc_diff: "-54%",  # 耗时差异
	//
	//       before_count: 1, # 前者总调用差异
	//
	//       after_count: 2,  # 后者总调用差异
	//
	//       count_diff: 1,   # 调用差异
	//
	//       before_count_perc: "56%", # 前者总调用差异
	//
	//       after_count_perc: "32%",  # 后者总调用差异
	//
	//       count_perc_diff: "44%",   # 调用差异
	//
	//   },
	//
	//   {...}
	//
	// ]"
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Error message description; empty if no error occurred
	//
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartAIDiffAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAIDiffAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *StartAIDiffAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartAIDiffAnalysisResponseBody) GetData() *string {
	return s.Data
}

func (s *StartAIDiffAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartAIDiffAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAIDiffAnalysisResponseBody) SetCode(v string) *StartAIDiffAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *StartAIDiffAnalysisResponseBody) SetData(v string) *StartAIDiffAnalysisResponseBody {
	s.Data = &v
	return s
}

func (s *StartAIDiffAnalysisResponseBody) SetMessage(v string) *StartAIDiffAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *StartAIDiffAnalysisResponseBody) SetRequestId(v string) *StartAIDiffAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAIDiffAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
