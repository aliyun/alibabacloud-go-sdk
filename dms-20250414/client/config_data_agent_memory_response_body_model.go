// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigDataAgentMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ConfigDataAgentMemoryResponseBodyData) *ConfigDataAgentMemoryResponseBody
	GetData() *ConfigDataAgentMemoryResponseBodyData
	SetErrorCode(v string) *ConfigDataAgentMemoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ConfigDataAgentMemoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ConfigDataAgentMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfigDataAgentMemoryResponseBody
	GetSuccess() *bool
}

type ConfigDataAgentMemoryResponseBody struct {
	Data *ConfigDataAgentMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigDataAgentMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigDataAgentMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigDataAgentMemoryResponseBody) GetData() *ConfigDataAgentMemoryResponseBodyData {
	return s.Data
}

func (s *ConfigDataAgentMemoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ConfigDataAgentMemoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ConfigDataAgentMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigDataAgentMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigDataAgentMemoryResponseBody) SetData(v *ConfigDataAgentMemoryResponseBodyData) *ConfigDataAgentMemoryResponseBody {
	s.Data = v
	return s
}

func (s *ConfigDataAgentMemoryResponseBody) SetErrorCode(v string) *ConfigDataAgentMemoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConfigDataAgentMemoryResponseBody) SetErrorMessage(v string) *ConfigDataAgentMemoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ConfigDataAgentMemoryResponseBody) SetRequestId(v string) *ConfigDataAgentMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigDataAgentMemoryResponseBody) SetSuccess(v bool) *ConfigDataAgentMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigDataAgentMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigDataAgentMemoryResponseBodyData struct {
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// true
	RecallEnabled *bool `json:"RecallEnabled,omitempty" xml:"RecallEnabled,omitempty"`
}

func (s ConfigDataAgentMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfigDataAgentMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfigDataAgentMemoryResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConfigDataAgentMemoryResponseBodyData) GetRecallEnabled() *bool {
	return s.RecallEnabled
}

func (s *ConfigDataAgentMemoryResponseBodyData) SetEnabled(v bool) *ConfigDataAgentMemoryResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *ConfigDataAgentMemoryResponseBodyData) SetRecallEnabled(v bool) *ConfigDataAgentMemoryResponseBodyData {
	s.RecallEnabled = &v
	return s
}

func (s *ConfigDataAgentMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
