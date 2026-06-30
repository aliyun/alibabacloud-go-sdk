// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAsrVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAsrVocabResponseBody
	GetCode() *string
	SetData(v string) *UpdateAsrVocabResponseBody
	GetData() *string
	SetMessage(v string) *UpdateAsrVocabResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAsrVocabResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAsrVocabResponseBody
	GetSuccess() *bool
}

type UpdateAsrVocabResponseBody struct {
	// The status code. A value of **200*	- indicates that the request was successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the hotword vocabulary. This is the value of the **vocabId*	- that you specified in the request.
	//
	// example:
	//
	// 71b1795ac8634bd8bdf4d3878480c7c2
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned. If the request is successful, **successful*	- is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates success. A value of **false*	- or **null*	- indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAsrVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAsrVocabResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateAsrVocabResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAsrVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAsrVocabResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAsrVocabResponseBody) SetCode(v string) *UpdateAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetData(v string) *UpdateAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetMessage(v string) *UpdateAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetRequestId(v string) *UpdateAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetSuccess(v bool) *UpdateAsrVocabResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) Validate() error {
	return dara.Validate(s)
}
