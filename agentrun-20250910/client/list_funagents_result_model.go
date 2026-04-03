// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunagentsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFunagentsResult
	GetCode() *string
	SetData(v *ListFunagentsOutput) *ListFunagentsResult
	GetData() *ListFunagentsOutput
	SetRequestId(v string) *ListFunagentsResult
	GetRequestId() *string
}

type ListFunagentsResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string              `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListFunagentsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListFunagentsResult) String() string {
	return dara.Prettify(s)
}

func (s ListFunagentsResult) GoString() string {
	return s.String()
}

func (s *ListFunagentsResult) GetCode() *string {
	return s.Code
}

func (s *ListFunagentsResult) GetData() *ListFunagentsOutput {
	return s.Data
}

func (s *ListFunagentsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFunagentsResult) SetCode(v string) *ListFunagentsResult {
	s.Code = &v
	return s
}

func (s *ListFunagentsResult) SetData(v *ListFunagentsOutput) *ListFunagentsResult {
	s.Data = v
	return s
}

func (s *ListFunagentsResult) SetRequestId(v string) *ListFunagentsResult {
	s.RequestId = &v
	return s
}

func (s *ListFunagentsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
