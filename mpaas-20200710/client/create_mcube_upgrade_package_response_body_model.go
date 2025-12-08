// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeUpgradePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateUpgradePackageResult(v *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) *CreateMcubeUpgradePackageResponseBody
	GetCreateUpgradePackageResult() *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult
	SetRequestId(v string) *CreateMcubeUpgradePackageResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeUpgradePackageResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeUpgradePackageResponseBody
	GetResultMessage() *string
}

type CreateMcubeUpgradePackageResponseBody struct {
	CreateUpgradePackageResult *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult `json:"CreateUpgradePackageResult,omitempty" xml:"CreateUpgradePackageResult,omitempty" type:"Struct"`
	RequestId                  *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                 *string                                                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage              *string                                                          `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeUpgradePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradePackageResponseBody) GetCreateUpgradePackageResult() *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult {
	return s.CreateUpgradePackageResult
}

func (s *CreateMcubeUpgradePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeUpgradePackageResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeUpgradePackageResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeUpgradePackageResponseBody) SetCreateUpgradePackageResult(v *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) *CreateMcubeUpgradePackageResponseBody {
	s.CreateUpgradePackageResult = v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBody) SetRequestId(v string) *CreateMcubeUpgradePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBody) SetResultCode(v string) *CreateMcubeUpgradePackageResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBody) SetResultMessage(v string) *CreateMcubeUpgradePackageResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBody) Validate() error {
	if s.CreateUpgradePackageResult != nil {
		if err := s.CreateUpgradePackageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult struct {
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg        *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	UpgradePackageId *string `json:"UpgradePackageId,omitempty" xml:"UpgradePackageId,omitempty"`
}

func (s CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) GetUpgradePackageId() *string {
	return s.UpgradePackageId
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) SetErrorCode(v string) *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult {
	s.ErrorCode = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) SetRequestId(v string) *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) SetResultMsg(v string) *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) SetSuccess(v bool) *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) SetUpgradePackageId(v string) *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult {
	s.UpgradePackageId = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyCreateUpgradePackageResult) Validate() error {
	return dara.Validate(s)
}
