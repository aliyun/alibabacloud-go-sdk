// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRiskCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRiskCheckTaskResponseBody
	GetCode() *string
	SetData(v *CreateRiskCheckTaskResponseBodyData) *CreateRiskCheckTaskResponseBody
	GetData() *CreateRiskCheckTaskResponseBodyData
	SetMessage(v string) *CreateRiskCheckTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRiskCheckTaskResponseBody
	GetRequestId() *string
}

type CreateRiskCheckTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateRiskCheckTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01A0220E-1F41-5260-A418-68286DF6B53D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRiskCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRiskCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRiskCheckTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRiskCheckTaskResponseBody) GetData() *CreateRiskCheckTaskResponseBodyData {
	return s.Data
}

func (s *CreateRiskCheckTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRiskCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRiskCheckTaskResponseBody) SetCode(v string) *CreateRiskCheckTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRiskCheckTaskResponseBody) SetData(v *CreateRiskCheckTaskResponseBodyData) *CreateRiskCheckTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateRiskCheckTaskResponseBody) SetMessage(v string) *CreateRiskCheckTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRiskCheckTaskResponseBody) SetRequestId(v string) *CreateRiskCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRiskCheckTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRiskCheckTaskResponseBodyData struct {
	// example:
	//
	// rct-xxxxxxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateRiskCheckTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateRiskCheckTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRiskCheckTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateRiskCheckTaskResponseBodyData) SetTaskId(v string) *CreateRiskCheckTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateRiskCheckTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
