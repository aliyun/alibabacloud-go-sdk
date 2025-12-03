// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateAsyncTranslateResponseBody
	GetCode() *int32
	SetData(v *CreateAsyncTranslateResponseBodyData) *CreateAsyncTranslateResponseBody
	GetData() *CreateAsyncTranslateResponseBodyData
	SetMessage(v string) *CreateAsyncTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAsyncTranslateResponseBody
	GetRequestId() *string
}

type CreateAsyncTranslateResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateAsyncTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAsyncTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsyncTranslateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateAsyncTranslateResponseBody) GetData() *CreateAsyncTranslateResponseBodyData {
	return s.Data
}

func (s *CreateAsyncTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAsyncTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAsyncTranslateResponseBody) SetCode(v int32) *CreateAsyncTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAsyncTranslateResponseBody) SetData(v *CreateAsyncTranslateResponseBodyData) *CreateAsyncTranslateResponseBody {
	s.Data = v
	return s
}

func (s *CreateAsyncTranslateResponseBody) SetMessage(v string) *CreateAsyncTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAsyncTranslateResponseBody) SetRequestId(v string) *CreateAsyncTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAsyncTranslateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAsyncTranslateResponseBodyData struct {
	// This parameter is required.
	//
	// example:
	//
	// 98bbb007-71bb-448b-bab0-2695ce8f8599
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAsyncTranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAsyncTranslateResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreateAsyncTranslateResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateAsyncTranslateResponseBodyData) SetJobId(v string) *CreateAsyncTranslateResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateAsyncTranslateResponseBodyData) SetStatus(v string) *CreateAsyncTranslateResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateAsyncTranslateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
