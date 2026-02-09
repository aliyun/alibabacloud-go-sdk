// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeHotpatchResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteHotpatchResourceResult(v *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) *DeleteMcubeHotpatchResourceResponseBody
	GetDeleteHotpatchResourceResult() *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult
	SetRequestId(v string) *DeleteMcubeHotpatchResourceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMcubeHotpatchResourceResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *DeleteMcubeHotpatchResourceResponseBody
	GetResultMessage() *string
}

type DeleteMcubeHotpatchResourceResponseBody struct {
	DeleteHotpatchResourceResult *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult `json:"DeleteHotpatchResourceResult,omitempty" xml:"DeleteHotpatchResourceResult,omitempty" type:"Struct"`
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

func (s DeleteMcubeHotpatchResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeHotpatchResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcubeHotpatchResourceResponseBody) GetDeleteHotpatchResourceResult() *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult {
	return s.DeleteHotpatchResourceResult
}

func (s *DeleteMcubeHotpatchResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcubeHotpatchResourceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMcubeHotpatchResourceResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMcubeHotpatchResourceResponseBody) SetDeleteHotpatchResourceResult(v *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) *DeleteMcubeHotpatchResourceResponseBody {
	s.DeleteHotpatchResourceResult = v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBody) SetRequestId(v string) *DeleteMcubeHotpatchResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBody) SetResultCode(v string) *DeleteMcubeHotpatchResourceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBody) SetResultMessage(v string) *DeleteMcubeHotpatchResourceResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBody) Validate() error {
	if s.DeleteHotpatchResourceResult != nil {
		if err := s.DeleteHotpatchResourceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult struct {
	// example:
	//
	// success
	DeleteResult *string `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 61B9F63C-4E6F-5D8E-A694-899450987B48
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

func (s DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) GoString() string {
	return s.String()
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) GetDeleteResult() *string {
	return s.DeleteResult
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) SetDeleteResult(v string) *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult {
	s.DeleteResult = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) SetErrorCode(v string) *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) SetRequestId(v string) *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult {
	s.RequestId = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) SetResultMsg(v string) *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult {
	s.ResultMsg = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) SetSuccess(v bool) *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult {
	s.Success = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponseBodyDeleteHotpatchResourceResult) Validate() error {
	return dara.Validate(s)
}
