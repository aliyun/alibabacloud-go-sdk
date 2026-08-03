// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateDataAgentMemoryResponseBodyData) *UpdateDataAgentMemoryResponseBody
	GetData() *UpdateDataAgentMemoryResponseBodyData
	SetErrorCode(v string) *UpdateDataAgentMemoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataAgentMemoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateDataAgentMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataAgentMemoryResponseBody
	GetSuccess() *bool
}

type UpdateDataAgentMemoryResponseBody struct {
	// The response struct.
	Data *UpdateDataAgentMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
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
	// Indicates whether the query is successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataAgentMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentMemoryResponseBody) GetData() *UpdateDataAgentMemoryResponseBodyData {
	return s.Data
}

func (s *UpdateDataAgentMemoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataAgentMemoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataAgentMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataAgentMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataAgentMemoryResponseBody) SetData(v *UpdateDataAgentMemoryResponseBodyData) *UpdateDataAgentMemoryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDataAgentMemoryResponseBody) SetErrorCode(v string) *UpdateDataAgentMemoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataAgentMemoryResponseBody) SetErrorMessage(v string) *UpdateDataAgentMemoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataAgentMemoryResponseBody) SetRequestId(v string) *UpdateDataAgentMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataAgentMemoryResponseBody) SetSuccess(v bool) *UpdateDataAgentMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataAgentMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataAgentMemoryResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// dlc1********63eqm
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

func (s UpdateDataAgentMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentMemoryResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *UpdateDataAgentMemoryResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataAgentMemoryResponseBodyData) SetJobId(v string) *UpdateDataAgentMemoryResponseBodyData {
	s.JobId = &v
	return s
}

func (s *UpdateDataAgentMemoryResponseBodyData) SetSuccess(v bool) *UpdateDataAgentMemoryResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateDataAgentMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
