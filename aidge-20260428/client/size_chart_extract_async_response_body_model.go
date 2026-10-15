// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartExtractAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SizeChartExtractAsyncResponseBody
	GetCode() *string
	SetData(v *SizeChartExtractAsyncResponseBodyData) *SizeChartExtractAsyncResponseBody
	GetData() *SizeChartExtractAsyncResponseBodyData
	SetMessage(v string) *SizeChartExtractAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *SizeChartExtractAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SizeChartExtractAsyncResponseBody
	GetSuccess() *bool
}

type SizeChartExtractAsyncResponseBody struct {
	// The error code. This parameter is not returned when the call is successful.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result of the asynchronous task for size chart extraction.
	Data *SizeChartExtractAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned when the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to identify a unique request.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true and false. A value of true indicates that the call is successful. A value of false indicates that the call fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SizeChartExtractAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *SizeChartExtractAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *SizeChartExtractAsyncResponseBody) GetData() *SizeChartExtractAsyncResponseBodyData {
	return s.Data
}

func (s *SizeChartExtractAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SizeChartExtractAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SizeChartExtractAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SizeChartExtractAsyncResponseBody) SetCode(v string) *SizeChartExtractAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *SizeChartExtractAsyncResponseBody) SetData(v *SizeChartExtractAsyncResponseBodyData) *SizeChartExtractAsyncResponseBody {
	s.Data = v
	return s
}

func (s *SizeChartExtractAsyncResponseBody) SetMessage(v string) *SizeChartExtractAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *SizeChartExtractAsyncResponseBody) SetRequestId(v string) *SizeChartExtractAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *SizeChartExtractAsyncResponseBody) SetSuccess(v bool) *SizeChartExtractAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *SizeChartExtractAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SizeChartExtractAsyncResponseBodyData struct {
	// The ID of the asynchronous task used to query the result.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SizeChartExtractAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *SizeChartExtractAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SizeChartExtractAsyncResponseBodyData) SetTaskId(v string) *SizeChartExtractAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SizeChartExtractAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
