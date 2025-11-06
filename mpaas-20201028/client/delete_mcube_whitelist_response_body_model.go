// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteWhitelistResult(v *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) *DeleteMcubeWhitelistResponseBody
	GetDeleteWhitelistResult() *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult
	SetRequestId(v string) *DeleteMcubeWhitelistResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMcubeWhitelistResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *DeleteMcubeWhitelistResponseBody
	GetResultMessage() *string
}

type DeleteMcubeWhitelistResponseBody struct {
	DeleteWhitelistResult *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult `json:"DeleteWhitelistResult,omitempty" xml:"DeleteWhitelistResult,omitempty" type:"Struct"`
	RequestId             *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode            *string                                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage         *string                                                `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMcubeWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcubeWhitelistResponseBody) GetDeleteWhitelistResult() *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult {
	return s.DeleteWhitelistResult
}

func (s *DeleteMcubeWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcubeWhitelistResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMcubeWhitelistResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMcubeWhitelistResponseBody) SetDeleteWhitelistResult(v *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) *DeleteMcubeWhitelistResponseBody {
	s.DeleteWhitelistResult = v
	return s
}

func (s *DeleteMcubeWhitelistResponseBody) SetRequestId(v string) *DeleteMcubeWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcubeWhitelistResponseBody) SetResultCode(v string) *DeleteMcubeWhitelistResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMcubeWhitelistResponseBody) SetResultMessage(v string) *DeleteMcubeWhitelistResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMcubeWhitelistResponseBody) Validate() error {
	if s.DeleteWhitelistResult != nil {
		if err := s.DeleteWhitelistResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) GoString() string {
	return s.String()
}

func (s *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) GetData() *string {
	return s.Data
}

func (s *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) SetData(v string) *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult {
	s.Data = &v
	return s
}

func (s *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) SetResultMsg(v string) *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult {
	s.ResultMsg = &v
	return s
}

func (s *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) SetSuccess(v bool) *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult {
	s.Success = &v
	return s
}

func (s *DeleteMcubeWhitelistResponseBodyDeleteWhitelistResult) Validate() error {
	return dara.Validate(s)
}
