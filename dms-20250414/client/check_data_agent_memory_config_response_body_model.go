// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataAgentMemoryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CheckDataAgentMemoryConfigResponseBodyData) *CheckDataAgentMemoryConfigResponseBody
	GetData() *CheckDataAgentMemoryConfigResponseBodyData
	SetErrorCode(v string) *CheckDataAgentMemoryConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CheckDataAgentMemoryConfigResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CheckDataAgentMemoryConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckDataAgentMemoryConfigResponseBody
	GetSuccess() *bool
}

type CheckDataAgentMemoryConfigResponseBody struct {
	// The response struct.
	Data *CheckDataAgentMemoryConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
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
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDataAgentMemoryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDataAgentMemoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataAgentMemoryConfigResponseBody) GetData() *CheckDataAgentMemoryConfigResponseBodyData {
	return s.Data
}

func (s *CheckDataAgentMemoryConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckDataAgentMemoryConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckDataAgentMemoryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDataAgentMemoryConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckDataAgentMemoryConfigResponseBody) SetData(v *CheckDataAgentMemoryConfigResponseBodyData) *CheckDataAgentMemoryConfigResponseBody {
	s.Data = v
	return s
}

func (s *CheckDataAgentMemoryConfigResponseBody) SetErrorCode(v string) *CheckDataAgentMemoryConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckDataAgentMemoryConfigResponseBody) SetErrorMessage(v string) *CheckDataAgentMemoryConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CheckDataAgentMemoryConfigResponseBody) SetRequestId(v string) *CheckDataAgentMemoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDataAgentMemoryConfigResponseBody) SetSuccess(v bool) *CheckDataAgentMemoryConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CheckDataAgentMemoryConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckDataAgentMemoryConfigResponseBodyData struct {
	// Indicates whether memory generation is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether memory recall is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	RecallEnabled *bool `json:"RecallEnabled,omitempty" xml:"RecallEnabled,omitempty"`
}

func (s CheckDataAgentMemoryConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckDataAgentMemoryConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckDataAgentMemoryConfigResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *CheckDataAgentMemoryConfigResponseBodyData) GetRecallEnabled() *bool {
	return s.RecallEnabled
}

func (s *CheckDataAgentMemoryConfigResponseBodyData) SetEnabled(v bool) *CheckDataAgentMemoryConfigResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *CheckDataAgentMemoryConfigResponseBodyData) SetRecallEnabled(v bool) *CheckDataAgentMemoryConfigResponseBodyData {
	s.RecallEnabled = &v
	return s
}

func (s *CheckDataAgentMemoryConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
