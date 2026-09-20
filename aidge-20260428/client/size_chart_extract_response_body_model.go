// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartExtractResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SizeChartExtractResponseBody
	GetCode() *string
	SetData(v *SizeChartExtractResponseBodyData) *SizeChartExtractResponseBody
	GetData() *SizeChartExtractResponseBodyData
	SetMessage(v string) *SizeChartExtractResponseBody
	GetMessage() *string
	SetRequestId(v string) *SizeChartExtractResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SizeChartExtractResponseBody
	GetSuccess() *bool
}

type SizeChartExtractResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The size chart extraction result.
	Data *SizeChartExtractResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SizeChartExtractResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractResponseBody) GoString() string {
	return s.String()
}

func (s *SizeChartExtractResponseBody) GetCode() *string {
	return s.Code
}

func (s *SizeChartExtractResponseBody) GetData() *SizeChartExtractResponseBodyData {
	return s.Data
}

func (s *SizeChartExtractResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SizeChartExtractResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SizeChartExtractResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SizeChartExtractResponseBody) SetCode(v string) *SizeChartExtractResponseBody {
	s.Code = &v
	return s
}

func (s *SizeChartExtractResponseBody) SetData(v *SizeChartExtractResponseBodyData) *SizeChartExtractResponseBody {
	s.Data = v
	return s
}

func (s *SizeChartExtractResponseBody) SetMessage(v string) *SizeChartExtractResponseBody {
	s.Message = &v
	return s
}

func (s *SizeChartExtractResponseBody) SetRequestId(v string) *SizeChartExtractResponseBody {
	s.RequestId = &v
	return s
}

func (s *SizeChartExtractResponseBody) SetSuccess(v bool) *SizeChartExtractResponseBody {
	s.Success = &v
	return s
}

func (s *SizeChartExtractResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SizeChartExtractResponseBodyData struct {
	// The asynchronous task ID, which is used to query the result later.
	//
	// example:
	//
	// task-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SizeChartExtractResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractResponseBodyData) GoString() string {
	return s.String()
}

func (s *SizeChartExtractResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SizeChartExtractResponseBodyData) SetTaskId(v string) *SizeChartExtractResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SizeChartExtractResponseBodyData) Validate() error {
	return dara.Validate(s)
}
