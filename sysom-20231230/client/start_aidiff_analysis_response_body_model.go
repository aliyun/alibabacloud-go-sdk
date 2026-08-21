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
	// The status code.
	//
	// - `code == Success` indicates that authorization is successful.
	//
	// - Other status codes indicate authorization failed. Check the `message` field for the detailed fault information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The AI differential analysis result data.
	//
	// example:
	//
	// "[
	//
	//   {
	//
	//       name: "xxx", #operator name
	//
	//       before_time: 2, # total time of the former
	//
	//       after_time: 4, # total time of the latter
	//
	//       time_diff: 2,  # time difference
	//
	//       before_time_perc: "80%", # total time percentage of the former
	//
	//       after_time_perc: "23%", # total time percentage of the latter
	//
	//       time_perc_diff: "-54%",  # time percentage difference
	//
	//       before_count: 1, # total call count of the former
	//
	//       after_count: 2,  # total call count of the latter
	//
	//       count_diff: 1,   # call count difference
	//
	//       before_count_perc: "56%", # total call percentage of the former
	//
	//       after_count_perc: "32%",  # total call percentage of the latter
	//
	//       count_perc_diff: "44%",   # call percentage difference
	//
	//   },
	//
	//   {...}
	//
	// ]"
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code description. This field is empty if no error occurs.
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
