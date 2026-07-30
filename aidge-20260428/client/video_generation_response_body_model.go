// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VideoGenerationResponseBody
	GetCode() *string
	SetData(v *VideoGenerationResponseBodyData) *VideoGenerationResponseBody
	GetData() *VideoGenerationResponseBodyData
	SetMessage(v string) *VideoGenerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VideoGenerationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VideoGenerationResponseBody
	GetSuccess() *bool
}

type VideoGenerationResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *VideoGenerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VideoGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *VideoGenerationResponseBody) GetCode() *string {
	return s.Code
}

func (s *VideoGenerationResponseBody) GetData() *VideoGenerationResponseBodyData {
	return s.Data
}

func (s *VideoGenerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VideoGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VideoGenerationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VideoGenerationResponseBody) SetCode(v string) *VideoGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *VideoGenerationResponseBody) SetData(v *VideoGenerationResponseBodyData) *VideoGenerationResponseBody {
	s.Data = v
	return s
}

func (s *VideoGenerationResponseBody) SetMessage(v string) *VideoGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *VideoGenerationResponseBody) SetRequestId(v string) *VideoGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VideoGenerationResponseBody) SetSuccess(v bool) *VideoGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *VideoGenerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoGenerationResponseBodyData struct {
	TaskId   *string           `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s VideoGenerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VideoGenerationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VideoGenerationResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *VideoGenerationResponseBodyData) SetTaskId(v string) *VideoGenerationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VideoGenerationResponseBodyData) SetUsageMap(v map[string]*int64) *VideoGenerationResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *VideoGenerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
