// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVmcoreDiagnosisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVmcoreDiagnosisTaskResponseBody
	GetCode() *string
	SetData(v *CreateVmcoreDiagnosisTaskResponseBodyData) *CreateVmcoreDiagnosisTaskResponseBody
	GetData() *CreateVmcoreDiagnosisTaskResponseBodyData
	SetMessage(v string) *CreateVmcoreDiagnosisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVmcoreDiagnosisTaskResponseBody
	GetRequestId() *string
}

type CreateVmcoreDiagnosisTaskResponseBody struct {
	// example:
	//
	// Success
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateVmcoreDiagnosisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateVmcoreDiagnosisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVmcoreDiagnosisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) GetData() *CreateVmcoreDiagnosisTaskResponseBodyData {
	return s.Data
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) SetCode(v string) *CreateVmcoreDiagnosisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) SetData(v *CreateVmcoreDiagnosisTaskResponseBodyData) *CreateVmcoreDiagnosisTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) SetMessage(v string) *CreateVmcoreDiagnosisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) SetRequestId(v string) *CreateVmcoreDiagnosisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVmcoreDiagnosisTaskResponseBodyData struct {
	// example:
	//
	// bbe94a98-4192-4172-b856-95777e0a55d7
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateVmcoreDiagnosisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateVmcoreDiagnosisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateVmcoreDiagnosisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVmcoreDiagnosisTaskResponseBodyData) SetTaskId(v string) *CreateVmcoreDiagnosisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
