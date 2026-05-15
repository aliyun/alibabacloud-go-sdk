// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelConnectionsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelConnectionsResult
	GetCode() *string
	SetData(v *ListModelConnectionsOutput) *ListModelConnectionsResult
	GetData() *ListModelConnectionsOutput
	SetRequestId(v string) *ListModelConnectionsResult
	GetRequestId() *string
}

type ListModelConnectionsResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 模型连接列表的详细信息
	//
	// example:
	//
	// {}
	Data *ListModelConnectionsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListModelConnectionsResult) String() string {
	return dara.Prettify(s)
}

func (s ListModelConnectionsResult) GoString() string {
	return s.String()
}

func (s *ListModelConnectionsResult) GetCode() *string {
	return s.Code
}

func (s *ListModelConnectionsResult) GetData() *ListModelConnectionsOutput {
	return s.Data
}

func (s *ListModelConnectionsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelConnectionsResult) SetCode(v string) *ListModelConnectionsResult {
	s.Code = &v
	return s
}

func (s *ListModelConnectionsResult) SetData(v *ListModelConnectionsOutput) *ListModelConnectionsResult {
	s.Data = v
	return s
}

func (s *ListModelConnectionsResult) SetRequestId(v string) *ListModelConnectionsResult {
	s.RequestId = &v
	return s
}

func (s *ListModelConnectionsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
