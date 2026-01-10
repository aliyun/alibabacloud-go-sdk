// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomDomainResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CustomDomainResult
	GetCode() *string
	SetData(v *CustomDomain) *CustomDomainResult
	GetData() *CustomDomain
	SetRequestId(v string) *CustomDomainResult
	GetRequestId() *string
}

type CustomDomainResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 自定义域名的详细信息
	//
	// example:
	//
	// {}
	Data *CustomDomain `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CustomDomainResult) String() string {
	return dara.Prettify(s)
}

func (s CustomDomainResult) GoString() string {
	return s.String()
}

func (s *CustomDomainResult) GetCode() *string {
	return s.Code
}

func (s *CustomDomainResult) GetData() *CustomDomain {
	return s.Data
}

func (s *CustomDomainResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomDomainResult) SetCode(v string) *CustomDomainResult {
	s.Code = &v
	return s
}

func (s *CustomDomainResult) SetData(v *CustomDomain) *CustomDomainResult {
	s.Data = v
	return s
}

func (s *CustomDomainResult) SetRequestId(v string) *CustomDomainResult {
	s.RequestId = &v
	return s
}

func (s *CustomDomainResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
