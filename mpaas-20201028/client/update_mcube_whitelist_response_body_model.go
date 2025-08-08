// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcubeWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddWhitelistResult(v *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) *UpdateMcubeWhitelistResponseBody
	GetAddWhitelistResult() *UpdateMcubeWhitelistResponseBodyAddWhitelistResult
	SetRequestId(v string) *UpdateMcubeWhitelistResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UpdateMcubeWhitelistResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *UpdateMcubeWhitelistResponseBody
	GetResultMessage() *string
}

type UpdateMcubeWhitelistResponseBody struct {
	AddWhitelistResult *UpdateMcubeWhitelistResponseBodyAddWhitelistResult `json:"AddWhitelistResult,omitempty" xml:"AddWhitelistResult,omitempty" type:"Struct"`
	RequestId          *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMcubeWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcubeWhitelistResponseBody) GetAddWhitelistResult() *UpdateMcubeWhitelistResponseBodyAddWhitelistResult {
	return s.AddWhitelistResult
}

func (s *UpdateMcubeWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcubeWhitelistResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UpdateMcubeWhitelistResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UpdateMcubeWhitelistResponseBody) SetAddWhitelistResult(v *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) *UpdateMcubeWhitelistResponseBody {
	s.AddWhitelistResult = v
	return s
}

func (s *UpdateMcubeWhitelistResponseBody) SetRequestId(v string) *UpdateMcubeWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcubeWhitelistResponseBody) SetResultCode(v string) *UpdateMcubeWhitelistResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMcubeWhitelistResponseBody) SetResultMessage(v string) *UpdateMcubeWhitelistResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UpdateMcubeWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateMcubeWhitelistResponseBodyAddWhitelistResult struct {
	AddWhitelistInfo *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo `json:"AddWhitelistInfo,omitempty" xml:"AddWhitelistInfo,omitempty" type:"Struct"`
	ResultMsg        *string                                                             `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success          *bool                                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMcubeWhitelistResponseBodyAddWhitelistResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeWhitelistResponseBodyAddWhitelistResult) GoString() string {
	return s.String()
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) GetAddWhitelistInfo() *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo {
	return s.AddWhitelistInfo
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) SetAddWhitelistInfo(v *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) *UpdateMcubeWhitelistResponseBodyAddWhitelistResult {
	s.AddWhitelistInfo = v
	return s
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) SetResultMsg(v string) *UpdateMcubeWhitelistResponseBodyAddWhitelistResult {
	s.ResultMsg = &v
	return s
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) SetSuccess(v bool) *UpdateMcubeWhitelistResponseBodyAddWhitelistResult {
	s.Success = &v
	return s
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResult) Validate() error {
	return dara.Validate(s)
}

type UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo struct {
	FailNum     *int64  `json:"FailNum,omitempty" xml:"FailNum,omitempty"`
	FailUserIds *string `json:"FailUserIds,omitempty" xml:"FailUserIds,omitempty"`
	SuccessNum  *int64  `json:"SuccessNum,omitempty" xml:"SuccessNum,omitempty"`
}

func (s UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) GoString() string {
	return s.String()
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) GetFailNum() *int64 {
	return s.FailNum
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) GetFailUserIds() *string {
	return s.FailUserIds
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) GetSuccessNum() *int64 {
	return s.SuccessNum
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) SetFailNum(v int64) *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo {
	s.FailNum = &v
	return s
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) SetFailUserIds(v string) *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo {
	s.FailUserIds = &v
	return s
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) SetSuccessNum(v int64) *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo {
	s.SuccessNum = &v
	return s
}

func (s *UpdateMcubeWhitelistResponseBodyAddWhitelistResultAddWhitelistInfo) Validate() error {
	return dara.Validate(s)
}
