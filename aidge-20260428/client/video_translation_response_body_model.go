// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoTranslationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VideoTranslationResponseBody
	GetCode() *string
	SetData(v *VideoTranslationResponseBodyData) *VideoTranslationResponseBody
	GetData() *VideoTranslationResponseBodyData
	SetMessage(v string) *VideoTranslationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VideoTranslationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VideoTranslationResponseBody
	GetSuccess() *bool
}

type VideoTranslationResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Processing
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asynchronous task submit status.
	Data *VideoTranslationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description message.
	//
	// example:
	//
	// Translation processing
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// req_20260608_jkl012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true: The call is successful. false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VideoTranslationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VideoTranslationResponseBody) GoString() string {
	return s.String()
}

func (s *VideoTranslationResponseBody) GetCode() *string {
	return s.Code
}

func (s *VideoTranslationResponseBody) GetData() *VideoTranslationResponseBodyData {
	return s.Data
}

func (s *VideoTranslationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VideoTranslationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VideoTranslationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VideoTranslationResponseBody) SetCode(v string) *VideoTranslationResponseBody {
	s.Code = &v
	return s
}

func (s *VideoTranslationResponseBody) SetData(v *VideoTranslationResponseBodyData) *VideoTranslationResponseBody {
	s.Data = v
	return s
}

func (s *VideoTranslationResponseBody) SetMessage(v string) *VideoTranslationResponseBody {
	s.Message = &v
	return s
}

func (s *VideoTranslationResponseBody) SetRequestId(v string) *VideoTranslationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VideoTranslationResponseBody) SetSuccess(v bool) *VideoTranslationResponseBody {
	s.Success = &v
	return s
}

func (s *VideoTranslationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoTranslationResponseBodyData struct {
	// The asynchronous task ID, used for subsequent queries.
	//
	// example:
	//
	// 0ea3b66e88a543658520c994f08896a0
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VideoTranslationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VideoTranslationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VideoTranslationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VideoTranslationResponseBodyData) SetTaskId(v string) *VideoTranslationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VideoTranslationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
