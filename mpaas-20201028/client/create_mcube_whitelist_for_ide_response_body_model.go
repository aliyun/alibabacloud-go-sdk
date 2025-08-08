// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeWhitelistForIdeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateWhitelistForIdeResult(v *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) *CreateMcubeWhitelistForIdeResponseBody
	GetCreateWhitelistForIdeResult() *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult
	SetRequestId(v string) *CreateMcubeWhitelistForIdeResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeWhitelistForIdeResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeWhitelistForIdeResponseBody
	GetResultMessage() *string
}

type CreateMcubeWhitelistForIdeResponseBody struct {
	CreateWhitelistForIdeResult *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult `json:"CreateWhitelistForIdeResult,omitempty" xml:"CreateWhitelistForIdeResult,omitempty" type:"Struct"`
	RequestId                   *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                  *string                                                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage               *string                                                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeWhitelistForIdeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeWhitelistForIdeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeWhitelistForIdeResponseBody) GetCreateWhitelistForIdeResult() *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult {
	return s.CreateWhitelistForIdeResult
}

func (s *CreateMcubeWhitelistForIdeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeWhitelistForIdeResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeWhitelistForIdeResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeWhitelistForIdeResponseBody) SetCreateWhitelistForIdeResult(v *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) *CreateMcubeWhitelistForIdeResponseBody {
	s.CreateWhitelistForIdeResult = v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponseBody) SetRequestId(v string) *CreateMcubeWhitelistForIdeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponseBody) SetResultCode(v string) *CreateMcubeWhitelistForIdeResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponseBody) SetResultMessage(v string) *CreateMcubeWhitelistForIdeResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult struct {
	ResultMsg   *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	WhitelistId *string `json:"WhitelistId,omitempty" xml:"WhitelistId,omitempty"`
}

func (s CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) GetWhitelistId() *string {
	return s.WhitelistId
}

func (s *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) SetResultMsg(v string) *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) SetSuccess(v bool) *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) SetWhitelistId(v string) *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult {
	s.WhitelistId = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponseBodyCreateWhitelistForIdeResult) Validate() error {
	return dara.Validate(s)
}
