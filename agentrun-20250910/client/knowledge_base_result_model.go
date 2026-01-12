// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *KnowledgeBaseResult
	GetCode() *string
	SetData(v *KnowledgeBase) *KnowledgeBaseResult
	GetData() *KnowledgeBase
	SetRequestId(v string) *KnowledgeBaseResult
	GetRequestId() *string
}

type KnowledgeBaseResult struct {
	Code      *string        `json:"code,omitempty" xml:"code,omitempty"`
	Data      *KnowledgeBase `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s KnowledgeBaseResult) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseResult) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseResult) GetCode() *string {
	return s.Code
}

func (s *KnowledgeBaseResult) GetData() *KnowledgeBase {
	return s.Data
}

func (s *KnowledgeBaseResult) GetRequestId() *string {
	return s.RequestId
}

func (s *KnowledgeBaseResult) SetCode(v string) *KnowledgeBaseResult {
	s.Code = &v
	return s
}

func (s *KnowledgeBaseResult) SetData(v *KnowledgeBase) *KnowledgeBaseResult {
	s.Data = v
	return s
}

func (s *KnowledgeBaseResult) SetRequestId(v string) *KnowledgeBaseResult {
	s.RequestId = &v
	return s
}

func (s *KnowledgeBaseResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
