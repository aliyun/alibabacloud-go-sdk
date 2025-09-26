// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CommonResult
	GetCode() *string
	SetData(v interface{}) *CommonResult
	GetData() interface{}
	SetRequestId(v string) *CommonResult
	GetRequestId() *string
}

type CommonResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 实际的业务数据内容
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CommonResult) String() string {
	return dara.Prettify(s)
}

func (s CommonResult) GoString() string {
	return s.String()
}

func (s *CommonResult) GetCode() *string {
	return s.Code
}

func (s *CommonResult) GetData() interface{} {
	return s.Data
}

func (s *CommonResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CommonResult) SetCode(v string) *CommonResult {
	s.Code = &v
	return s
}

func (s *CommonResult) SetData(v interface{}) *CommonResult {
	s.Data = v
	return s
}

func (s *CommonResult) SetRequestId(v string) *CommonResult {
	s.RequestId = &v
	return s
}

func (s *CommonResult) Validate() error {
	return dara.Validate(s)
}
