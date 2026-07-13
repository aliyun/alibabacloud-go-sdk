// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CheckRulesResponseBodyData) *CheckRulesResponseBody
	GetData() *CheckRulesResponseBodyData
	SetRequestId(v string) *CheckRulesResponseBody
	GetRequestId() *string
}

type CheckRulesResponseBody struct {
	// The returned data.
	Data *CheckRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// 700683DE-0154-56D4-8D76-3B7A2C2C7DF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRulesResponseBody) GetData() *CheckRulesResponseBodyData {
	return s.Data
}

func (s *CheckRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckRulesResponseBody) SetData(v *CheckRulesResponseBodyData) *CheckRulesResponseBody {
	s.Data = v
	return s
}

func (s *CheckRulesResponseBody) SetRequestId(v string) *CheckRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckRulesResponseBodyData struct {
	// The unique ID of the asynchronous task.
	//
	// example:
	//
	// t-0000e4w0u1v592zdf6s7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CheckRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckRulesResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CheckRulesResponseBodyData) SetTaskId(v string) *CheckRulesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CheckRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
