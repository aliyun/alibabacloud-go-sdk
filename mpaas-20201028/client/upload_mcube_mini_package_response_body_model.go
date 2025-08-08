// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMcubeMiniPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadMcubeMiniPackageResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UploadMcubeMiniPackageResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *UploadMcubeMiniPackageResponseBody
	GetResultMessage() *string
	SetUploadMiniPackageResult(v *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) *UploadMcubeMiniPackageResponseBody
	GetUploadMiniPackageResult() *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult
}

type UploadMcubeMiniPackageResponseBody struct {
	RequestId               *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode              *string                                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage           *string                                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	UploadMiniPackageResult *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult `json:"UploadMiniPackageResult,omitempty" xml:"UploadMiniPackageResult,omitempty" type:"Struct"`
}

func (s UploadMcubeMiniPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeMiniPackageResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMcubeMiniPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadMcubeMiniPackageResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UploadMcubeMiniPackageResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UploadMcubeMiniPackageResponseBody) GetUploadMiniPackageResult() *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult {
	return s.UploadMiniPackageResult
}

func (s *UploadMcubeMiniPackageResponseBody) SetRequestId(v string) *UploadMcubeMiniPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMcubeMiniPackageResponseBody) SetResultCode(v string) *UploadMcubeMiniPackageResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UploadMcubeMiniPackageResponseBody) SetResultMessage(v string) *UploadMcubeMiniPackageResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UploadMcubeMiniPackageResponseBody) SetUploadMiniPackageResult(v *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) *UploadMcubeMiniPackageResponseBody {
	s.UploadMiniPackageResult = v
	return s
}

func (s *UploadMcubeMiniPackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult struct {
	ResultMsg           *string                                                                       `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	ReturnPackageResult *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult `json:"ReturnPackageResult,omitempty" xml:"ReturnPackageResult,omitempty" type:"Struct"`
	Success             *bool                                                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) GoString() string {
	return s.String()
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) GetReturnPackageResult() *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult {
	return s.ReturnPackageResult
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) GetSuccess() *bool {
	return s.Success
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) SetResultMsg(v string) *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult {
	s.ResultMsg = &v
	return s
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) SetReturnPackageResult(v *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult {
	s.ReturnPackageResult = v
	return s
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) SetSuccess(v bool) *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult {
	s.Success = &v
	return s
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResult) Validate() error {
	return dara.Validate(s)
}

type UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult struct {
	DebugUrl  *string `json:"DebugUrl,omitempty" xml:"DebugUrl,omitempty"`
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) GoString() string {
	return s.String()
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) GetDebugUrl() *string {
	return s.DebugUrl
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) GetPackageId() *string {
	return s.PackageId
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) GetUserId() *string {
	return s.UserId
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) SetDebugUrl(v string) *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult {
	s.DebugUrl = &v
	return s
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) SetPackageId(v string) *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult {
	s.PackageId = &v
	return s
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) SetUserId(v string) *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult {
	s.UserId = &v
	return s
}

func (s *UploadMcubeMiniPackageResponseBodyUploadMiniPackageResultReturnPackageResult) Validate() error {
	return dara.Validate(s)
}
