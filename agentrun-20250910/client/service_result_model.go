// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ServiceResult
	GetCode() *string
	SetData(v interface{}) *ServiceResult
	GetData() interface{}
	SetRequestId(v string) *ServiceResult
	GetRequestId() *string
}

type ServiceResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 服务的详细信息
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ServiceResult) String() string {
	return dara.Prettify(s)
}

func (s ServiceResult) GoString() string {
	return s.String()
}

func (s *ServiceResult) GetCode() *string {
	return s.Code
}

func (s *ServiceResult) GetData() interface{} {
	return s.Data
}

func (s *ServiceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ServiceResult) SetCode(v string) *ServiceResult {
	s.Code = &v
	return s
}

func (s *ServiceResult) SetData(v interface{}) *ServiceResult {
	s.Data = v
	return s
}

func (s *ServiceResult) SetRequestId(v string) *ServiceResult {
	s.RequestId = &v
	return s
}

func (s *ServiceResult) Validate() error {
	return dara.Validate(s)
}
