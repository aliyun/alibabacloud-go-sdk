// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCodeInterpreterResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCodeInterpreterResult
	GetCode() *string
	SetData(v *CodeInterpreter) *DeleteCodeInterpreterResult
	GetData() *CodeInterpreter
	SetRequestId(v string) *DeleteCodeInterpreterResult
	GetRequestId() *string
}

type DeleteCodeInterpreterResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 被删除的代码解释器详细信息
	//
	// example:
	//
	// {}
	Data *CodeInterpreter `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteCodeInterpreterResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteCodeInterpreterResult) GoString() string {
	return s.String()
}

func (s *DeleteCodeInterpreterResult) GetCode() *string {
	return s.Code
}

func (s *DeleteCodeInterpreterResult) GetData() *CodeInterpreter {
	return s.Data
}

func (s *DeleteCodeInterpreterResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCodeInterpreterResult) SetCode(v string) *DeleteCodeInterpreterResult {
	s.Code = &v
	return s
}

func (s *DeleteCodeInterpreterResult) SetData(v *CodeInterpreter) *DeleteCodeInterpreterResult {
	s.Data = v
	return s
}

func (s *DeleteCodeInterpreterResult) SetRequestId(v string) *DeleteCodeInterpreterResult {
	s.RequestId = &v
	return s
}

func (s *DeleteCodeInterpreterResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
