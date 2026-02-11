// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePocTaskListResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribePocTaskListResponseBody
	GetHttpStatusCode() *string
	SetMessage(v bool) *DescribePocTaskListResponseBody
	GetMessage() *bool
	SetRequestId(v string) *DescribePocTaskListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribePocTaskListResponseBodyResultObject) *DescribePocTaskListResponseBody
	GetResultObject() []*DescribePocTaskListResponseBodyResultObject
}

type DescribePocTaskListResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *bool `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	ResultObject []*DescribePocTaskListResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
}

func (s DescribePocTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePocTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePocTaskListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePocTaskListResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribePocTaskListResponseBody) GetMessage() *bool {
	return s.Message
}

func (s *DescribePocTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePocTaskListResponseBody) GetResultObject() []*DescribePocTaskListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribePocTaskListResponseBody) SetCode(v string) *DescribePocTaskListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePocTaskListResponseBody) SetHttpStatusCode(v string) *DescribePocTaskListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribePocTaskListResponseBody) SetMessage(v bool) *DescribePocTaskListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePocTaskListResponseBody) SetRequestId(v string) *DescribePocTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePocTaskListResponseBody) SetResultObject(v []*DescribePocTaskListResponseBodyResultObject) *DescribePocTaskListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribePocTaskListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePocTaskListResponseBodyResultObject struct {
	// Retro task ID.
	//
	// example:
	//
	// 01
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Retro task name.
	//
	// example:
	//
	// xxx
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribePocTaskListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribePocTaskListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribePocTaskListResponseBodyResultObject) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePocTaskListResponseBodyResultObject) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribePocTaskListResponseBodyResultObject) SetTaskId(v string) *DescribePocTaskListResponseBodyResultObject {
	s.TaskId = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) SetTaskName(v string) *DescribePocTaskListResponseBodyResultObject {
	s.TaskName = &v
	return s
}

func (s *DescribePocTaskListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
