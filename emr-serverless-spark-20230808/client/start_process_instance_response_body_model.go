// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartProcessInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *StartProcessInstanceResponseBody
	GetCode() *int32
	SetData(v interface{}) *StartProcessInstanceResponseBody
	GetData() interface{}
	SetFailed(v bool) *StartProcessInstanceResponseBody
	GetFailed() *bool
	SetHttpStatusCode(v int32) *StartProcessInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *StartProcessInstanceResponseBody
	GetMsg() *string
	SetRequestId(v string) *StartProcessInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartProcessInstanceResponseBody
	GetSuccess() *bool
}

type StartProcessInstanceResponseBody struct {
	// The code that is returned by the backend server.
	//
	// example:
	//
	// 1400009
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {\\"sessionBizId\\": \\"sc-dc85644dba1c8c63\\", \\"bizId\\": \\"st-aeed3b0d4f87418a9a9dcbd757477658\\", \\"gmtCreated\\": \\"Thu Sep 12 02:28:45 UTC 2024\\"}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Indicates whether the workflow fails to be run manually.
	//
	// example:
	//
	// false
	Failed *bool `json:"failed,omitempty" xml:"failed,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the returned code.
	//
	// example:
	//
	// No permission for resource action
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartProcessInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *StartProcessInstanceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *StartProcessInstanceResponseBody) GetFailed() *bool {
	return s.Failed
}

func (s *StartProcessInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartProcessInstanceResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *StartProcessInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartProcessInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartProcessInstanceResponseBody) SetCode(v int32) *StartProcessInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetData(v interface{}) *StartProcessInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StartProcessInstanceResponseBody) SetFailed(v bool) *StartProcessInstanceResponseBody {
	s.Failed = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetHttpStatusCode(v int32) *StartProcessInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetMsg(v string) *StartProcessInstanceResponseBody {
	s.Msg = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetRequestId(v string) *StartProcessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetSuccess(v bool) *StartProcessInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartProcessInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
