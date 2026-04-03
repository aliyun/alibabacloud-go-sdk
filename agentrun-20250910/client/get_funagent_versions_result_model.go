// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunagentVersionsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFunagentVersionsResult
	GetCode() *string
	SetData(v *GetFunagentVersionsOutput) *GetFunagentVersionsResult
	GetData() *GetFunagentVersionsOutput
	SetRequestId(v string) *GetFunagentVersionsResult
	GetRequestId() *string
}

type GetFunagentVersionsResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetFunagentVersionsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetFunagentVersionsResult) String() string {
	return dara.Prettify(s)
}

func (s GetFunagentVersionsResult) GoString() string {
	return s.String()
}

func (s *GetFunagentVersionsResult) GetCode() *string {
	return s.Code
}

func (s *GetFunagentVersionsResult) GetData() *GetFunagentVersionsOutput {
	return s.Data
}

func (s *GetFunagentVersionsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunagentVersionsResult) SetCode(v string) *GetFunagentVersionsResult {
	s.Code = &v
	return s
}

func (s *GetFunagentVersionsResult) SetData(v *GetFunagentVersionsOutput) *GetFunagentVersionsResult {
	s.Data = v
	return s
}

func (s *GetFunagentVersionsResult) SetRequestId(v string) *GetFunagentVersionsResult {
	s.RequestId = &v
	return s
}

func (s *GetFunagentVersionsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
