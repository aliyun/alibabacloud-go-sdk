// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateHotpatchTaskResult(v *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) *CreateMcubeHotpatchTaskResponseBody
	GetCreateHotpatchTaskResult() *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult
	SetRequestId(v string) *CreateMcubeHotpatchTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeHotpatchTaskResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeHotpatchTaskResponseBody
	GetResultMessage() *string
}

type CreateMcubeHotpatchTaskResponseBody struct {
	CreateHotpatchTaskResult *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult `json:"CreateHotpatchTaskResult,omitempty" xml:"CreateHotpatchTaskResult,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeHotpatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchTaskResponseBody) GetCreateHotpatchTaskResult() *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult {
	return s.CreateHotpatchTaskResult
}

func (s *CreateMcubeHotpatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeHotpatchTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeHotpatchTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeHotpatchTaskResponseBody) SetCreateHotpatchTaskResult(v *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) *CreateMcubeHotpatchTaskResponseBody {
	s.CreateHotpatchTaskResult = v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBody) SetRequestId(v string) *CreateMcubeHotpatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBody) SetResultCode(v string) *CreateMcubeHotpatchTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBody) SetResultMessage(v string) *CreateMcubeHotpatchTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBody) Validate() error {
	if s.CreateHotpatchTaskResult != nil {
		if err := s.CreateHotpatchTaskResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult struct {
	// example:
	//
	// OK
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1490
	HotpatchTaskId *string `json:"HotpatchTaskId,omitempty" xml:"HotpatchTaskId,omitempty"`
	// example:
	//
	// FD352990-0D38-5580-B6E1-EBE64B234059
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) GetHotpatchTaskId() *string {
	return s.HotpatchTaskId
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) SetErrorCode(v string) *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult {
	s.ErrorCode = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) SetHotpatchTaskId(v string) *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult {
	s.HotpatchTaskId = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) SetRequestId(v string) *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) SetResultMsg(v string) *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) SetSuccess(v bool) *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponseBodyCreateHotpatchTaskResult) Validate() error {
	return dara.Validate(s)
}
