// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncWritingBiddingDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AsyncWritingBiddingDocResponseBody
	GetCode() *string
	SetData(v *AsyncWritingBiddingDocResponseBodyData) *AsyncWritingBiddingDocResponseBody
	GetData() *AsyncWritingBiddingDocResponseBodyData
	SetHttpStatusCode(v int32) *AsyncWritingBiddingDocResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AsyncWritingBiddingDocResponseBody
	GetMessage() *string
	SetRequestId(v string) *AsyncWritingBiddingDocResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AsyncWritingBiddingDocResponseBody
	GetSuccess() *bool
}

type AsyncWritingBiddingDocResponseBody struct {
	// example:
	//
	// successful
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AsyncWritingBiddingDocResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AsyncWritingBiddingDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsyncWritingBiddingDocResponseBody) GoString() string {
	return s.String()
}

func (s *AsyncWritingBiddingDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *AsyncWritingBiddingDocResponseBody) GetData() *AsyncWritingBiddingDocResponseBodyData {
	return s.Data
}

func (s *AsyncWritingBiddingDocResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AsyncWritingBiddingDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AsyncWritingBiddingDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncWritingBiddingDocResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncWritingBiddingDocResponseBody) SetCode(v string) *AsyncWritingBiddingDocResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncWritingBiddingDocResponseBody) SetData(v *AsyncWritingBiddingDocResponseBodyData) *AsyncWritingBiddingDocResponseBody {
	s.Data = v
	return s
}

func (s *AsyncWritingBiddingDocResponseBody) SetHttpStatusCode(v int32) *AsyncWritingBiddingDocResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AsyncWritingBiddingDocResponseBody) SetMessage(v string) *AsyncWritingBiddingDocResponseBody {
	s.Message = &v
	return s
}

func (s *AsyncWritingBiddingDocResponseBody) SetRequestId(v string) *AsyncWritingBiddingDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncWritingBiddingDocResponseBody) SetSuccess(v bool) *AsyncWritingBiddingDocResponseBody {
	s.Success = &v
	return s
}

func (s *AsyncWritingBiddingDocResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AsyncWritingBiddingDocResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AsyncWritingBiddingDocResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AsyncWritingBiddingDocResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsyncWritingBiddingDocResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncWritingBiddingDocResponseBodyData) SetTaskId(v string) *AsyncWritingBiddingDocResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AsyncWritingBiddingDocResponseBodyData) Validate() error {
	return dara.Validate(s)
}
