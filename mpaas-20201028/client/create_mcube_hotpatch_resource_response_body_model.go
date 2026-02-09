// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateHotpatchResourceResult(v *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) *CreateMcubeHotpatchResourceResponseBody
	GetCreateHotpatchResourceResult() *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult
	SetRequestId(v string) *CreateMcubeHotpatchResourceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeHotpatchResourceResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeHotpatchResourceResponseBody
	GetResultMessage() *string
}

type CreateMcubeHotpatchResourceResponseBody struct {
	CreateHotpatchResourceResult *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult `json:"CreateHotpatchResourceResult,omitempty" xml:"CreateHotpatchResourceResult,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeHotpatchResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchResourceResponseBody) GetCreateHotpatchResourceResult() *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult {
	return s.CreateHotpatchResourceResult
}

func (s *CreateMcubeHotpatchResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeHotpatchResourceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeHotpatchResourceResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeHotpatchResourceResponseBody) SetCreateHotpatchResourceResult(v *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) *CreateMcubeHotpatchResourceResponseBody {
	s.CreateHotpatchResourceResult = v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBody) SetRequestId(v string) *CreateMcubeHotpatchResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBody) SetResultCode(v string) *CreateMcubeHotpatchResourceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBody) SetResultMessage(v string) *CreateMcubeHotpatchResourceResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBody) Validate() error {
	if s.CreateHotpatchResourceResult != nil {
		if err := s.CreateHotpatchResourceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult struct {
	// example:
	//
	// OK
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1768
	HotpatchResourceId *string `json:"HotpatchResourceId,omitempty" xml:"HotpatchResourceId,omitempty"`
	// example:
	//
	// EA606F90-F758-5EDC-A70F-939F089CA496
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) GetHotpatchResourceId() *string {
	return s.HotpatchResourceId
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) SetErrorCode(v string) *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) SetHotpatchResourceId(v string) *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult {
	s.HotpatchResourceId = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) SetRequestId(v string) *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) SetResultMsg(v string) *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) SetSuccess(v bool) *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponseBodyCreateHotpatchResourceResult) Validate() error {
	return dara.Validate(s)
}
