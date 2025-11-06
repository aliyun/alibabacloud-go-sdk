// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateWhitelistResult(v *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) *CreateMcubeWhitelistResponseBody
	GetCreateWhitelistResult() *CreateMcubeWhitelistResponseBodyCreateWhitelistResult
	SetRequestId(v string) *CreateMcubeWhitelistResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeWhitelistResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeWhitelistResponseBody
	GetResultMessage() *string
}

type CreateMcubeWhitelistResponseBody struct {
	CreateWhitelistResult *CreateMcubeWhitelistResponseBodyCreateWhitelistResult `json:"CreateWhitelistResult,omitempty" xml:"CreateWhitelistResult,omitempty" type:"Struct"`
	RequestId             *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode            *string                                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage         *string                                                `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeWhitelistResponseBody) GetCreateWhitelistResult() *CreateMcubeWhitelistResponseBodyCreateWhitelistResult {
	return s.CreateWhitelistResult
}

func (s *CreateMcubeWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeWhitelistResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeWhitelistResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeWhitelistResponseBody) SetCreateWhitelistResult(v *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) *CreateMcubeWhitelistResponseBody {
	s.CreateWhitelistResult = v
	return s
}

func (s *CreateMcubeWhitelistResponseBody) SetRequestId(v string) *CreateMcubeWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeWhitelistResponseBody) SetResultCode(v string) *CreateMcubeWhitelistResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeWhitelistResponseBody) SetResultMessage(v string) *CreateMcubeWhitelistResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeWhitelistResponseBody) Validate() error {
	if s.CreateWhitelistResult != nil {
		if err := s.CreateWhitelistResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcubeWhitelistResponseBodyCreateWhitelistResult struct {
	ResultMsg   *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	WhitelistId *string `json:"WhitelistId,omitempty" xml:"WhitelistId,omitempty"`
}

func (s CreateMcubeWhitelistResponseBodyCreateWhitelistResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeWhitelistResponseBodyCreateWhitelistResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) GetWhitelistId() *string {
	return s.WhitelistId
}

func (s *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) SetResultMsg(v string) *CreateMcubeWhitelistResponseBodyCreateWhitelistResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) SetSuccess(v bool) *CreateMcubeWhitelistResponseBodyCreateWhitelistResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) SetWhitelistId(v string) *CreateMcubeWhitelistResponseBodyCreateWhitelistResult {
	s.WhitelistId = &v
	return s
}

func (s *CreateMcubeWhitelistResponseBodyCreateWhitelistResult) Validate() error {
	return dara.Validate(s)
}
