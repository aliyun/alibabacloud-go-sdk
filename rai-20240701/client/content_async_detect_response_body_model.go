// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentAsyncDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ContentAsyncDetectResponseBody
	GetCode() *string
	SetData(v *ContentAsyncDetectResponseBodyData) *ContentAsyncDetectResponseBody
	GetData() *ContentAsyncDetectResponseBodyData
	SetHttpStatusCode(v int32) *ContentAsyncDetectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ContentAsyncDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *ContentAsyncDetectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ContentAsyncDetectResponseBody
	GetSuccess() *bool
}

type ContentAsyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ContentAsyncDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ContentAsyncDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContentAsyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ContentAsyncDetectResponseBody) GetData() *ContentAsyncDetectResponseBodyData {
	return s.Data
}

func (s *ContentAsyncDetectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ContentAsyncDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ContentAsyncDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContentAsyncDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ContentAsyncDetectResponseBody) SetCode(v string) *ContentAsyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetData(v *ContentAsyncDetectResponseBodyData) *ContentAsyncDetectResponseBody {
	s.Data = v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetHttpStatusCode(v int32) *ContentAsyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetMessage(v string) *ContentAsyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetRequestId(v string) *ContentAsyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetSuccess(v bool) *ContentAsyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ContentAsyncDetectResponseBodyData struct {
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ContentAsyncDetectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ContentAsyncDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ContentAsyncDetectResponseBodyData) SetTaskId(v string) *ContentAsyncDetectResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ContentAsyncDetectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
