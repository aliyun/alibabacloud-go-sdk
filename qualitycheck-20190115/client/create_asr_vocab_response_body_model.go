// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsrVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAsrVocabResponseBody
	GetCode() *string
	SetData(v string) *CreateAsrVocabResponseBody
	GetData() *string
	SetMessage(v string) *CreateAsrVocabResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAsrVocabResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAsrVocabResponseBody
	GetSuccess() *bool
}

type CreateAsrVocabResponseBody struct {
	// Result code. **200*	- means success. Any other value means failure. Use this field to diagnose errors.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Hotword ID
	//
	// example:
	//
	// 71b1795ac8634bd8bdf4d3878480c7c2
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error details if the call fails. Returns **successful*	- on success.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. true means success. **false*	- or **null*	- means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAsrVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAsrVocabResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateAsrVocabResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAsrVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAsrVocabResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAsrVocabResponseBody) SetCode(v string) *CreateAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetData(v string) *CreateAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetMessage(v string) *CreateAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetRequestId(v string) *CreateAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetSuccess(v bool) *CreateAsrVocabResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAsrVocabResponseBody) Validate() error {
	return dara.Validate(s)
}
