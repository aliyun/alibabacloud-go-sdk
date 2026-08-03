// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDataAgentMemoryResponseBodyData) *DeleteDataAgentMemoryResponseBody
	GetData() *DeleteDataAgentMemoryResponseBodyData
	SetErrorCode(v string) *DeleteDataAgentMemoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataAgentMemoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataAgentMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataAgentMemoryResponseBody
	GetSuccess() *bool
}

type DeleteDataAgentMemoryResponseBody struct {
	// The response struct.
	Data *DeleteDataAgentMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request is abnormal.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAgentMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentMemoryResponseBody) GetData() *DeleteDataAgentMemoryResponseBodyData {
	return s.Data
}

func (s *DeleteDataAgentMemoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataAgentMemoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataAgentMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataAgentMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAgentMemoryResponseBody) SetData(v *DeleteDataAgentMemoryResponseBodyData) *DeleteDataAgentMemoryResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataAgentMemoryResponseBody) SetErrorCode(v string) *DeleteDataAgentMemoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataAgentMemoryResponseBody) SetErrorMessage(v string) *DeleteDataAgentMemoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataAgentMemoryResponseBody) SetRequestId(v string) *DeleteDataAgentMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataAgentMemoryResponseBody) SetSuccess(v bool) *DeleteDataAgentMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDataAgentMemoryResponseBodyData struct {
	// The call ID.
	//
	// example:
	//
	// dlc7c***********2r8v7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Indicates whether the task submission request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAgentMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentMemoryResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *DeleteDataAgentMemoryResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAgentMemoryResponseBodyData) SetJobId(v string) *DeleteDataAgentMemoryResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DeleteDataAgentMemoryResponseBodyData) SetSuccess(v bool) *DeleteDataAgentMemoryResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
