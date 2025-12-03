// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateImageTranslateTaskResponseBody
	GetCode() *int32
	SetData(v *CreateImageTranslateTaskResponseBodyData) *CreateImageTranslateTaskResponseBody
	GetData() *CreateImageTranslateTaskResponseBodyData
	SetMessage(v string) *CreateImageTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateImageTranslateTaskResponseBody
	GetRequestId() *string
}

type CreateImageTranslateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateImageTranslateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A41F6E25-8520-4AF0-90EF-AF7E32840108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateImageTranslateTaskResponseBody) GetData() *CreateImageTranslateTaskResponseBodyData {
	return s.Data
}

func (s *CreateImageTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateImageTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageTranslateTaskResponseBody) SetCode(v int32) *CreateImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetData(v *CreateImageTranslateTaskResponseBodyData) *CreateImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetMessage(v string) *CreateImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetRequestId(v string) *CreateImageTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageTranslateTaskResponseBodyData struct {
	// example:
	//
	// A41F6E25-8520-4AF0-90EF-111111
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageTranslateTaskResponseBodyData) SetTaskId(v string) *CreateImageTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateImageTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
