// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteKnowledgeBaseResult
	GetCode() *string
	SetData(v *KnowledgeBase) *DeleteKnowledgeBaseResult
	GetData() *KnowledgeBase
	SetRequestId(v string) *DeleteKnowledgeBaseResult
	GetRequestId() *string
}

type DeleteKnowledgeBaseResult struct {
	Code      *string        `json:"code,omitempty" xml:"code,omitempty"`
	Data      *KnowledgeBase `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteKnowledgeBaseResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseResult) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseResult) GetCode() *string {
	return s.Code
}

func (s *DeleteKnowledgeBaseResult) GetData() *KnowledgeBase {
	return s.Data
}

func (s *DeleteKnowledgeBaseResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeBaseResult) SetCode(v string) *DeleteKnowledgeBaseResult {
	s.Code = &v
	return s
}

func (s *DeleteKnowledgeBaseResult) SetData(v *KnowledgeBase) *DeleteKnowledgeBaseResult {
	s.Data = v
	return s
}

func (s *DeleteKnowledgeBaseResult) SetRequestId(v string) *DeleteKnowledgeBaseResult {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeBaseResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
