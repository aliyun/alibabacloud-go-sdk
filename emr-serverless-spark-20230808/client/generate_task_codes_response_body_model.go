// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTaskCodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GenerateTaskCodesResponseBody
	GetCode() *int32
	SetData(v []*int64) *GenerateTaskCodesResponseBody
	GetData() []*int64
	SetFailed(v bool) *GenerateTaskCodesResponseBody
	GetFailed() *bool
	SetHttpStatusCode(v int32) *GenerateTaskCodesResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *GenerateTaskCodesResponseBody
	GetMsg() *string
	SetRequestId(v string) *GenerateTaskCodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateTaskCodesResponseBody
	GetSuccess() *bool
}

type GenerateTaskCodesResponseBody struct {
	// example:
	//
	// 1000000
	Code *int32   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*int64 `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Failed *bool `json:"failed,omitempty" xml:"failed,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// No permission for resource action
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GenerateTaskCodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateTaskCodesResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTaskCodesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GenerateTaskCodesResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *GenerateTaskCodesResponseBody) GetFailed() *bool {
	return s.Failed
}

func (s *GenerateTaskCodesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateTaskCodesResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GenerateTaskCodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateTaskCodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateTaskCodesResponseBody) SetCode(v int32) *GenerateTaskCodesResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateTaskCodesResponseBody) SetData(v []*int64) *GenerateTaskCodesResponseBody {
	s.Data = v
	return s
}

func (s *GenerateTaskCodesResponseBody) SetFailed(v bool) *GenerateTaskCodesResponseBody {
	s.Failed = &v
	return s
}

func (s *GenerateTaskCodesResponseBody) SetHttpStatusCode(v int32) *GenerateTaskCodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateTaskCodesResponseBody) SetMsg(v string) *GenerateTaskCodesResponseBody {
	s.Msg = &v
	return s
}

func (s *GenerateTaskCodesResponseBody) SetRequestId(v string) *GenerateTaskCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateTaskCodesResponseBody) SetSuccess(v bool) *GenerateTaskCodesResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateTaskCodesResponseBody) Validate() error {
	return dara.Validate(s)
}
