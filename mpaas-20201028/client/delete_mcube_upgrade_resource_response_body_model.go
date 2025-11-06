// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeUpgradeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteResult(v *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) *DeleteMcubeUpgradeResourceResponseBody
	GetDeleteResult() *DeleteMcubeUpgradeResourceResponseBodyDeleteResult
	SetRequestId(v string) *DeleteMcubeUpgradeResourceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMcubeUpgradeResourceResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *DeleteMcubeUpgradeResourceResponseBody
	GetResultMessage() *string
}

type DeleteMcubeUpgradeResourceResponseBody struct {
	DeleteResult  *DeleteMcubeUpgradeResourceResponseBodyDeleteResult `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty" type:"Struct"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMcubeUpgradeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeUpgradeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcubeUpgradeResourceResponseBody) GetDeleteResult() *DeleteMcubeUpgradeResourceResponseBodyDeleteResult {
	return s.DeleteResult
}

func (s *DeleteMcubeUpgradeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcubeUpgradeResourceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMcubeUpgradeResourceResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMcubeUpgradeResourceResponseBody) SetDeleteResult(v *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) *DeleteMcubeUpgradeResourceResponseBody {
	s.DeleteResult = v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponseBody) SetRequestId(v string) *DeleteMcubeUpgradeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponseBody) SetResultCode(v string) *DeleteMcubeUpgradeResourceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponseBody) SetResultMessage(v string) *DeleteMcubeUpgradeResourceResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponseBody) Validate() error {
	if s.DeleteResult != nil {
		if err := s.DeleteResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMcubeUpgradeResourceResponseBodyDeleteResult struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcubeUpgradeResourceResponseBodyDeleteResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeUpgradeResourceResponseBodyDeleteResult) GoString() string {
	return s.String()
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) SetErrorCode(v string) *DeleteMcubeUpgradeResourceResponseBodyDeleteResult {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) SetRequestId(v string) *DeleteMcubeUpgradeResourceResponseBodyDeleteResult {
	s.RequestId = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) SetResultMsg(v string) *DeleteMcubeUpgradeResourceResponseBodyDeleteResult {
	s.ResultMsg = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) SetSuccess(v bool) *DeleteMcubeUpgradeResourceResponseBodyDeleteResult {
	s.Success = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponseBodyDeleteResult) Validate() error {
	return dara.Validate(s)
}
