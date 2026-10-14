// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAifgAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartAifgAnalysisResponseBody
	GetCode() *string
	SetData(v *StartAifgAnalysisResponseBodyData) *StartAifgAnalysisResponseBody
	GetData() *StartAifgAnalysisResponseBodyData
	SetMessage(v string) *StartAifgAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartAifgAnalysisResponseBody
	GetRequestId() *string
}

type StartAifgAnalysisResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200, 404, 500
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data details.
	Data *StartAifgAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The message content.
	//
	// example:
	//
	// Processed successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartAifgAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAifgAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *StartAifgAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartAifgAnalysisResponseBody) GetData() *StartAifgAnalysisResponseBodyData {
	return s.Data
}

func (s *StartAifgAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartAifgAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAifgAnalysisResponseBody) SetCode(v string) *StartAifgAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *StartAifgAnalysisResponseBody) SetData(v *StartAifgAnalysisResponseBodyData) *StartAifgAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *StartAifgAnalysisResponseBody) SetMessage(v string) *StartAifgAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *StartAifgAnalysisResponseBody) SetRequestId(v string) *StartAifgAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAifgAnalysisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartAifgAnalysisResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 9cd7aee5-d365-4d4e-8236-58a110a4ea09
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s StartAifgAnalysisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartAifgAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartAifgAnalysisResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *StartAifgAnalysisResponseBodyData) SetTaskId(v string) *StartAifgAnalysisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *StartAifgAnalysisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
