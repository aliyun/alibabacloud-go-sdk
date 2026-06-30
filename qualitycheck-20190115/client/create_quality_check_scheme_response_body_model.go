// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityCheckSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateQualityCheckSchemeResponseBody
	GetCode() *string
	SetData(v int64) *CreateQualityCheckSchemeResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateQualityCheckSchemeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateQualityCheckSchemeResponseBody
	GetMessage() *string
	SetMessages(v *CreateQualityCheckSchemeResponseBodyMessages) *CreateQualityCheckSchemeResponseBody
	GetMessages() *CreateQualityCheckSchemeResponseBodyMessages
	SetRequestId(v string) *CreateQualityCheckSchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQualityCheckSchemeResponseBody
	GetSuccess() *bool
}

type CreateQualityCheckSchemeResponseBody struct {
	// Result code. A value of 200 indicates success. Any other value indicates failure. Use this field to identify the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ID of the created quality check plan.
	//
	// example:
	//
	// 12
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error details if the request failed. If the request succeeded, the value is **successful**.
	//
	// example:
	//
	// successful
	Message  *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *CreateQualityCheckSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. If true, the request succeeded. If false or null, the request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQualityCheckSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateQualityCheckSchemeResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateQualityCheckSchemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateQualityCheckSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateQualityCheckSchemeResponseBody) GetMessages() *CreateQualityCheckSchemeResponseBodyMessages {
	return s.Messages
}

func (s *CreateQualityCheckSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityCheckSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQualityCheckSchemeResponseBody) SetCode(v string) *CreateQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetData(v int64) *CreateQualityCheckSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetHttpStatusCode(v int32) *CreateQualityCheckSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetMessage(v string) *CreateQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetMessages(v *CreateQualityCheckSchemeResponseBodyMessages) *CreateQualityCheckSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetRequestId(v string) *CreateQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetSuccess(v bool) *CreateQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateQualityCheckSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s CreateQualityCheckSchemeResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckSchemeResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *CreateQualityCheckSchemeResponseBodyMessages) SetMessage(v []*string) *CreateQualityCheckSchemeResponseBodyMessages {
	s.Message = v
	return s
}

func (s *CreateQualityCheckSchemeResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
