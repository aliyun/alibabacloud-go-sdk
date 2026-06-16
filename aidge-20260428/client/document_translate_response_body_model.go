// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DocumentTranslateResponseBody
	GetCode() *string
	SetData(v *DocumentTranslateResponseBodyData) *DocumentTranslateResponseBody
	GetData() *DocumentTranslateResponseBodyData
	SetMessage(v string) *DocumentTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DocumentTranslateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DocumentTranslateResponseBody
	GetSuccess() *bool
}

type DocumentTranslateResponseBody struct {
	// The response code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asynchronous task information.
	Data *DocumentTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 39E8A74B-F99E-1195-A5FF-3ECC5F94F304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DocumentTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocumentTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *DocumentTranslateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DocumentTranslateResponseBody) GetData() *DocumentTranslateResponseBodyData {
	return s.Data
}

func (s *DocumentTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DocumentTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocumentTranslateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DocumentTranslateResponseBody) SetCode(v string) *DocumentTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *DocumentTranslateResponseBody) SetData(v *DocumentTranslateResponseBodyData) *DocumentTranslateResponseBody {
	s.Data = v
	return s
}

func (s *DocumentTranslateResponseBody) SetMessage(v string) *DocumentTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *DocumentTranslateResponseBody) SetRequestId(v string) *DocumentTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocumentTranslateResponseBody) SetSuccess(v bool) *DocumentTranslateResponseBody {
	s.Success = &v
	return s
}

func (s *DocumentTranslateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocumentTranslateResponseBodyData struct {
	// The unique identifier of the asynchronous task. Use this ID to query the task status and result.
	//
	// example:
	//
	// 6071a030-5c92-9df1-96d0-44952343439a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DocumentTranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DocumentTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *DocumentTranslateResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DocumentTranslateResponseBodyData) SetTaskId(v string) *DocumentTranslateResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DocumentTranslateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
