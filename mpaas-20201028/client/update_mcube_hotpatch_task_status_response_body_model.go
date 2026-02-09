// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcubeHotpatchTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMcubeHotpatchTaskStatusResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UpdateMcubeHotpatchTaskStatusResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *UpdateMcubeHotpatchTaskStatusResponseBody
	GetResultMessage() *string
	SetUpdateHotpatchTaskStatusResult(v *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) *UpdateMcubeHotpatchTaskStatusResponseBody
	GetUpdateHotpatchTaskStatusResult() *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult
}

type UpdateMcubeHotpatchTaskStatusResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// success
	ResultMessage                  *string                                                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	UpdateHotpatchTaskStatusResult *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult `json:"UpdateHotpatchTaskStatusResult,omitempty" xml:"UpdateHotpatchTaskStatusResult,omitempty" type:"Struct"`
}

func (s UpdateMcubeHotpatchTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeHotpatchTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) GetUpdateHotpatchTaskStatusResult() *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult {
	return s.UpdateHotpatchTaskStatusResult
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) SetRequestId(v string) *UpdateMcubeHotpatchTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) SetResultCode(v string) *UpdateMcubeHotpatchTaskStatusResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) SetResultMessage(v string) *UpdateMcubeHotpatchTaskStatusResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) SetUpdateHotpatchTaskStatusResult(v *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) *UpdateMcubeHotpatchTaskStatusResponseBody {
	s.UpdateHotpatchTaskStatusResult = v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBody) Validate() error {
	if s.UpdateHotpatchTaskStatusResult != nil {
		if err := s.UpdateHotpatchTaskStatusResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6844B5-279D-5FFD-BD5A-2E1F9BEC39EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) GoString() string {
	return s.String()
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) GetResult() *string {
	return s.Result
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) SetErrorCode(v string) *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) SetRequestId(v string) *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult {
	s.RequestId = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) SetResult(v string) *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult {
	s.Result = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) SetResultMsg(v string) *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult {
	s.ResultMsg = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) SetSuccess(v bool) *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult {
	s.Success = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponseBodyUpdateHotpatchTaskStatusResult) Validate() error {
	return dara.Validate(s)
}
