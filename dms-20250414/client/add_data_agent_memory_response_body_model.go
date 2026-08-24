// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataAgentMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddDataAgentMemoryResponseBodyData) *AddDataAgentMemoryResponseBody
	GetData() *AddDataAgentMemoryResponseBodyData
	SetErrorCode(v string) *AddDataAgentMemoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddDataAgentMemoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddDataAgentMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDataAgentMemoryResponseBody
	GetSuccess() *bool
}

type AddDataAgentMemoryResponseBody struct {
	// The response struct.
	Data *AddDataAgentMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful. Valid values:
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

func (s AddDataAgentMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataAgentMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataAgentMemoryResponseBody) GetData() *AddDataAgentMemoryResponseBodyData {
	return s.Data
}

func (s *AddDataAgentMemoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddDataAgentMemoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddDataAgentMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataAgentMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataAgentMemoryResponseBody) SetData(v *AddDataAgentMemoryResponseBodyData) *AddDataAgentMemoryResponseBody {
	s.Data = v
	return s
}

func (s *AddDataAgentMemoryResponseBody) SetErrorCode(v string) *AddDataAgentMemoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddDataAgentMemoryResponseBody) SetErrorMessage(v string) *AddDataAgentMemoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddDataAgentMemoryResponseBody) SetRequestId(v string) *AddDataAgentMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataAgentMemoryResponseBody) SetSuccess(v bool) *AddDataAgentMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataAgentMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDataAgentMemoryResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// dlc1********63eqm
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Indicates whether the task submission request is successful. Valid values:
	//
	// - True: The request is successful.
	//
	// - False: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDataAgentMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddDataAgentMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDataAgentMemoryResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *AddDataAgentMemoryResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataAgentMemoryResponseBodyData) SetJobId(v string) *AddDataAgentMemoryResponseBodyData {
	s.JobId = &v
	return s
}

func (s *AddDataAgentMemoryResponseBodyData) SetSuccess(v bool) *AddDataAgentMemoryResponseBodyData {
	s.Success = &v
	return s
}

func (s *AddDataAgentMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
