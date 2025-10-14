// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipCurrentStepResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SkipCurrentStepResponseBodyData) *SkipCurrentStepResponseBody
	GetData() *SkipCurrentStepResponseBodyData
	SetMessage(v string) *SkipCurrentStepResponseBody
	GetMessage() *string
	SetRequestId(v string) *SkipCurrentStepResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SkipCurrentStepResponseBody
	GetSuccess() *bool
}

type SkipCurrentStepResponseBody struct {
	Data *SkipCurrentStepResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SkipCurrentStepResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SkipCurrentStepResponseBody) GoString() string {
	return s.String()
}

func (s *SkipCurrentStepResponseBody) GetData() *SkipCurrentStepResponseBodyData {
	return s.Data
}

func (s *SkipCurrentStepResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SkipCurrentStepResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SkipCurrentStepResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SkipCurrentStepResponseBody) SetData(v *SkipCurrentStepResponseBodyData) *SkipCurrentStepResponseBody {
	s.Data = v
	return s
}

func (s *SkipCurrentStepResponseBody) SetMessage(v string) *SkipCurrentStepResponseBody {
	s.Message = &v
	return s
}

func (s *SkipCurrentStepResponseBody) SetRequestId(v string) *SkipCurrentStepResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipCurrentStepResponseBody) SetSuccess(v bool) *SkipCurrentStepResponseBody {
	s.Success = &v
	return s
}

func (s *SkipCurrentStepResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SkipCurrentStepResponseBodyData struct {
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s SkipCurrentStepResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SkipCurrentStepResponseBodyData) GoString() string {
	return s.String()
}

func (s *SkipCurrentStepResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *SkipCurrentStepResponseBodyData) SetSlinkTaskId(v string) *SkipCurrentStepResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *SkipCurrentStepResponseBodyData) Validate() error {
	return dara.Validate(s)
}
