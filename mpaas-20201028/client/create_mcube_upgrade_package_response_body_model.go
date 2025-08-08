// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeUpgradePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMcubeUpgradePackageResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeUpgradePackageResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateMcubeUpgradePackageResponseBodyResultContent) *CreateMcubeUpgradePackageResponseBody
	GetResultContent() *CreateMcubeUpgradePackageResponseBodyResultContent
	SetResultMessage(v string) *CreateMcubeUpgradePackageResponseBody
	GetResultMessage() *string
}

type CreateMcubeUpgradePackageResponseBody struct {
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateMcubeUpgradePackageResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeUpgradePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeUpgradePackageResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeUpgradePackageResponseBody) GetResultContent() *CreateMcubeUpgradePackageResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateMcubeUpgradePackageResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeUpgradePackageResponseBody) SetRequestId(v string) *CreateMcubeUpgradePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBody) SetResultCode(v string) *CreateMcubeUpgradePackageResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBody) SetResultContent(v *CreateMcubeUpgradePackageResponseBodyResultContent) *CreateMcubeUpgradePackageResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBody) SetResultMessage(v string) *CreateMcubeUpgradePackageResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcubeUpgradePackageResponseBodyResultContent struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeUpgradePackageResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradePackageResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) SetData(v string) *CreateMcubeUpgradePackageResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) SetErrorCode(v string) *CreateMcubeUpgradePackageResponseBodyResultContent {
	s.ErrorCode = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) SetRequestId(v string) *CreateMcubeUpgradePackageResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) SetResultMsg(v string) *CreateMcubeUpgradePackageResponseBodyResultContent {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) SetSuccess(v bool) *CreateMcubeUpgradePackageResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
