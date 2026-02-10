// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTemplatesResult
	GetCode() *string
	SetData(v *ListTemplatesOutput) *ListTemplatesResult
	GetData() *ListTemplatesOutput
	SetRequestId(v string) *ListTemplatesResult
	GetRequestId() *string
}

type ListTemplatesResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 模板列表的详细信息
	Data *ListTemplatesOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// C0595DB0-D1EE-55C3-8DDD-790872C7EC2F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListTemplatesResult) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResult) GoString() string {
	return s.String()
}

func (s *ListTemplatesResult) GetCode() *string {
	return s.Code
}

func (s *ListTemplatesResult) GetData() *ListTemplatesOutput {
	return s.Data
}

func (s *ListTemplatesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatesResult) SetCode(v string) *ListTemplatesResult {
	s.Code = &v
	return s
}

func (s *ListTemplatesResult) SetData(v *ListTemplatesOutput) *ListTemplatesResult {
	s.Data = v
	return s
}

func (s *ListTemplatesResult) SetRequestId(v string) *ListTemplatesResult {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
