// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMcubeRsaKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadMcubeRsaKeyResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UploadMcubeRsaKeyResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *UploadMcubeRsaKeyResponseBody
	GetResultMessage() *string
	SetUploadRsaResult(v *UploadMcubeRsaKeyResponseBodyUploadRsaResult) *UploadMcubeRsaKeyResponseBody
	GetUploadRsaResult() *UploadMcubeRsaKeyResponseBodyUploadRsaResult
}

type UploadMcubeRsaKeyResponseBody struct {
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode      *string                                       `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage   *string                                       `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	UploadRsaResult *UploadMcubeRsaKeyResponseBodyUploadRsaResult `json:"UploadRsaResult,omitempty" xml:"UploadRsaResult,omitempty" type:"Struct"`
}

func (s UploadMcubeRsaKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeRsaKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMcubeRsaKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadMcubeRsaKeyResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UploadMcubeRsaKeyResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UploadMcubeRsaKeyResponseBody) GetUploadRsaResult() *UploadMcubeRsaKeyResponseBodyUploadRsaResult {
	return s.UploadRsaResult
}

func (s *UploadMcubeRsaKeyResponseBody) SetRequestId(v string) *UploadMcubeRsaKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMcubeRsaKeyResponseBody) SetResultCode(v string) *UploadMcubeRsaKeyResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UploadMcubeRsaKeyResponseBody) SetResultMessage(v string) *UploadMcubeRsaKeyResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UploadMcubeRsaKeyResponseBody) SetUploadRsaResult(v *UploadMcubeRsaKeyResponseBodyUploadRsaResult) *UploadMcubeRsaKeyResponseBody {
	s.UploadRsaResult = v
	return s
}

func (s *UploadMcubeRsaKeyResponseBody) Validate() error {
	if s.UploadRsaResult != nil {
		if err := s.UploadRsaResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadMcubeRsaKeyResponseBodyUploadRsaResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadMcubeRsaKeyResponseBodyUploadRsaResult) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeRsaKeyResponseBodyUploadRsaResult) GoString() string {
	return s.String()
}

func (s *UploadMcubeRsaKeyResponseBodyUploadRsaResult) GetData() *string {
	return s.Data
}

func (s *UploadMcubeRsaKeyResponseBodyUploadRsaResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *UploadMcubeRsaKeyResponseBodyUploadRsaResult) GetSuccess() *bool {
	return s.Success
}

func (s *UploadMcubeRsaKeyResponseBodyUploadRsaResult) SetData(v string) *UploadMcubeRsaKeyResponseBodyUploadRsaResult {
	s.Data = &v
	return s
}

func (s *UploadMcubeRsaKeyResponseBodyUploadRsaResult) SetResultMsg(v string) *UploadMcubeRsaKeyResponseBodyUploadRsaResult {
	s.ResultMsg = &v
	return s
}

func (s *UploadMcubeRsaKeyResponseBodyUploadRsaResult) SetSuccess(v bool) *UploadMcubeRsaKeyResponseBodyUploadRsaResult {
	s.Success = &v
	return s
}

func (s *UploadMcubeRsaKeyResponseBodyUploadRsaResult) Validate() error {
	return dara.Validate(s)
}
