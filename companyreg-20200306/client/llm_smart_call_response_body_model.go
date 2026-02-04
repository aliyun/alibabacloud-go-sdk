// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallRecordId(v string) *LlmSmartCallResponseBody
	GetCallRecordId() *string
	SetErrorCode(v string) *LlmSmartCallResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *LlmSmartCallResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *LlmSmartCallResponseBody
	GetRequestId() *string
	SetSuccess(v string) *LlmSmartCallResponseBody
	GetSuccess() *string
}

type LlmSmartCallResponseBody struct {
	// example:
	//
	// UUID
	CallRecordId *string `json:"CallRecordId,omitempty" xml:"CallRecordId,omitempty"`
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5D8BD6E8-28D9-5451-BBA1-3D3DCA6971F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LlmSmartCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *LlmSmartCallResponseBody) GetCallRecordId() *string {
	return s.CallRecordId
}

func (s *LlmSmartCallResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *LlmSmartCallResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *LlmSmartCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LlmSmartCallResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *LlmSmartCallResponseBody) SetCallRecordId(v string) *LlmSmartCallResponseBody {
	s.CallRecordId = &v
	return s
}

func (s *LlmSmartCallResponseBody) SetErrorCode(v string) *LlmSmartCallResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LlmSmartCallResponseBody) SetErrorMsg(v string) *LlmSmartCallResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *LlmSmartCallResponseBody) SetRequestId(v string) *LlmSmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *LlmSmartCallResponseBody) SetSuccess(v string) *LlmSmartCallResponseBody {
	s.Success = &v
	return s
}

func (s *LlmSmartCallResponseBody) Validate() error {
	return dara.Validate(s)
}
