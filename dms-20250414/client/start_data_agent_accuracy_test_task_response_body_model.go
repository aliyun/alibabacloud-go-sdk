// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataAgentAccuracyTestTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartDataAgentAccuracyTestTaskResponseBodyData) *StartDataAgentAccuracyTestTaskResponseBody
	GetData() *StartDataAgentAccuracyTestTaskResponseBodyData
	SetErrorCode(v string) *StartDataAgentAccuracyTestTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartDataAgentAccuracyTestTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *StartDataAgentAccuracyTestTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StartDataAgentAccuracyTestTaskResponseBody
	GetSuccess() *string
}

type StartDataAgentAccuracyTestTaskResponseBody struct {
	// The response struct.
	Data *StartDataAgentAccuracyTestTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D94F5232-xxx-EH0H28FGGI5I
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - True: The request is successful.
	//
	// - False: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartDataAgentAccuracyTestTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDataAgentAccuracyTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) GetData() *StartDataAgentAccuracyTestTaskResponseBodyData {
	return s.Data
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) SetData(v *StartDataAgentAccuracyTestTaskResponseBodyData) *StartDataAgentAccuracyTestTaskResponseBody {
	s.Data = v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) SetErrorCode(v string) *StartDataAgentAccuracyTestTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) SetErrorMessage(v string) *StartDataAgentAccuracyTestTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) SetRequestId(v string) *StartDataAgentAccuracyTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) SetSuccess(v string) *StartDataAgentAccuracyTestTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartDataAgentAccuracyTestTaskResponseBodyData struct {
	// The ID of the accuracy test task.
	//
	// example:
	//
	// 692abb8f-xxx-77fec862db34
	AccuracyTestTaskId *string `json:"AccuracyTestTaskId,omitempty" xml:"AccuracyTestTaskId,omitempty"`
}

func (s StartDataAgentAccuracyTestTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartDataAgentAccuracyTestTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartDataAgentAccuracyTestTaskResponseBodyData) GetAccuracyTestTaskId() *string {
	return s.AccuracyTestTaskId
}

func (s *StartDataAgentAccuracyTestTaskResponseBodyData) SetAccuracyTestTaskId(v string) *StartDataAgentAccuracyTestTaskResponseBodyData {
	s.AccuracyTestTaskId = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
