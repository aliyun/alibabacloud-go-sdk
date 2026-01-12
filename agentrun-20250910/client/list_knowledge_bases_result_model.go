// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBasesResult
	GetCode() *string
	SetData(v *ListKnowledgeBasesOutput) *ListKnowledgeBasesResult
	GetData() *ListKnowledgeBasesOutput
	SetRequestId(v string) *ListKnowledgeBasesResult
	GetRequestId() *string
}

type ListKnowledgeBasesResult struct {
	Code      *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListKnowledgeBasesOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListKnowledgeBasesResult) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResult) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResult) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBasesResult) GetData() *ListKnowledgeBasesOutput {
	return s.Data
}

func (s *ListKnowledgeBasesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBasesResult) SetCode(v string) *ListKnowledgeBasesResult {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBasesResult) SetData(v *ListKnowledgeBasesOutput) *ListKnowledgeBasesResult {
	s.Data = v
	return s
}

func (s *ListKnowledgeBasesResult) SetRequestId(v string) *ListKnowledgeBasesResult {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBasesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
