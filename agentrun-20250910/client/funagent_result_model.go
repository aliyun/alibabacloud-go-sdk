// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunagentResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FunagentResult
	GetCode() *string
	SetData(v *Funagent) *FunagentResult
	GetData() *Funagent
	SetRequestId(v string) *FunagentResult
	GetRequestId() *string
}

type FunagentResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string   `json:"code,omitempty" xml:"code,omitempty"`
	Data *Funagent `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s FunagentResult) String() string {
	return dara.Prettify(s)
}

func (s FunagentResult) GoString() string {
	return s.String()
}

func (s *FunagentResult) GetCode() *string {
	return s.Code
}

func (s *FunagentResult) GetData() *Funagent {
	return s.Data
}

func (s *FunagentResult) GetRequestId() *string {
	return s.RequestId
}

func (s *FunagentResult) SetCode(v string) *FunagentResult {
	s.Code = &v
	return s
}

func (s *FunagentResult) SetData(v *Funagent) *FunagentResult {
	s.Data = v
	return s
}

func (s *FunagentResult) SetRequestId(v string) *FunagentResult {
	s.RequestId = &v
	return s
}

func (s *FunagentResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
