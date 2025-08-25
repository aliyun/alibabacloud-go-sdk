// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeNebulaAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteMcubeNebulaAppResult(v *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) *DeleteMcubeNebulaAppResponseBody
	GetDeleteMcubeNebulaAppResult() *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult
	SetRequestId(v string) *DeleteMcubeNebulaAppResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMcubeNebulaAppResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *DeleteMcubeNebulaAppResponseBody
	GetResultMessage() *string
}

type DeleteMcubeNebulaAppResponseBody struct {
	DeleteMcubeNebulaAppResult *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult `json:"DeleteMcubeNebulaAppResult,omitempty" xml:"DeleteMcubeNebulaAppResult,omitempty" type:"Struct"`
	RequestId                  *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                 *string                                                     `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage              *string                                                     `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMcubeNebulaAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeNebulaAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcubeNebulaAppResponseBody) GetDeleteMcubeNebulaAppResult() *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult {
	return s.DeleteMcubeNebulaAppResult
}

func (s *DeleteMcubeNebulaAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcubeNebulaAppResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMcubeNebulaAppResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMcubeNebulaAppResponseBody) SetDeleteMcubeNebulaAppResult(v *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) *DeleteMcubeNebulaAppResponseBody {
	s.DeleteMcubeNebulaAppResult = v
	return s
}

func (s *DeleteMcubeNebulaAppResponseBody) SetRequestId(v string) *DeleteMcubeNebulaAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcubeNebulaAppResponseBody) SetResultCode(v string) *DeleteMcubeNebulaAppResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMcubeNebulaAppResponseBody) SetResultMessage(v string) *DeleteMcubeNebulaAppResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMcubeNebulaAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) GoString() string {
	return s.String()
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) SetErrorCode(v string) *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) SetRequestId(v string) *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult {
	s.RequestId = &v
	return s
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) SetResultMsg(v string) *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult {
	s.ResultMsg = &v
	return s
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) SetSuccess(v bool) *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult {
	s.Success = &v
	return s
}

func (s *DeleteMcubeNebulaAppResponseBodyDeleteMcubeNebulaAppResult) Validate() error {
	return dara.Validate(s)
}
