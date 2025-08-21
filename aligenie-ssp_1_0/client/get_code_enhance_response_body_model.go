// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeEnhanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetCodeEnhanceResponseBody
	GetCode() *int32
	SetMessage(v string) *GetCodeEnhanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCodeEnhanceResponseBody
	GetRequestId() *string
	SetResult(v string) *GetCodeEnhanceResponseBody
	GetResult() *string
}

type GetCodeEnhanceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Aexfgc
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetCodeEnhanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCodeEnhanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetCodeEnhanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCodeEnhanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCodeEnhanceResponseBody) GetResult() *string {
	return s.Result
}

func (s *GetCodeEnhanceResponseBody) SetCode(v int32) *GetCodeEnhanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetCodeEnhanceResponseBody) SetMessage(v string) *GetCodeEnhanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetCodeEnhanceResponseBody) SetRequestId(v string) *GetCodeEnhanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeEnhanceResponseBody) SetResult(v string) *GetCodeEnhanceResponseBody {
	s.Result = &v
	return s
}

func (s *GetCodeEnhanceResponseBody) Validate() error {
	return dara.Validate(s)
}
