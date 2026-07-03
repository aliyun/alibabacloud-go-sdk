// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentAccuracyTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDataAgentAccuracyTestResponseBodyData) *CreateDataAgentAccuracyTestResponseBody
	GetData() *CreateDataAgentAccuracyTestResponseBodyData
	SetErrorCode(v string) *CreateDataAgentAccuracyTestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataAgentAccuracyTestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataAgentAccuracyTestResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateDataAgentAccuracyTestResponseBody
	GetSuccess() *string
}

type CreateDataAgentAccuracyTestResponseBody struct {
	// The returned result.
	Data *CreateDataAgentAccuracyTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A61C2009-xxx-BE7E95CEDF2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataAgentAccuracyTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentAccuracyTestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataAgentAccuracyTestResponseBody) GetData() *CreateDataAgentAccuracyTestResponseBodyData {
	return s.Data
}

func (s *CreateDataAgentAccuracyTestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataAgentAccuracyTestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataAgentAccuracyTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataAgentAccuracyTestResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateDataAgentAccuracyTestResponseBody) SetData(v *CreateDataAgentAccuracyTestResponseBodyData) *CreateDataAgentAccuracyTestResponseBody {
	s.Data = v
	return s
}

func (s *CreateDataAgentAccuracyTestResponseBody) SetErrorCode(v string) *CreateDataAgentAccuracyTestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataAgentAccuracyTestResponseBody) SetErrorMessage(v string) *CreateDataAgentAccuracyTestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataAgentAccuracyTestResponseBody) SetRequestId(v string) *CreateDataAgentAccuracyTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataAgentAccuracyTestResponseBody) SetSuccess(v string) *CreateDataAgentAccuracyTestResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataAgentAccuracyTestResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentAccuracyTestResponseBodyData struct {
	// The instance ID of the accuracy test.
	//
	// example:
	//
	// at-xxxxxxxxxxxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
}

func (s CreateDataAgentAccuracyTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentAccuracyTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDataAgentAccuracyTestResponseBodyData) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *CreateDataAgentAccuracyTestResponseBodyData) SetAccuracyTestInsId(v string) *CreateDataAgentAccuracyTestResponseBodyData {
	s.AccuracyTestInsId = &v
	return s
}

func (s *CreateDataAgentAccuracyTestResponseBodyData) Validate() error {
	return dara.Validate(s)
}
