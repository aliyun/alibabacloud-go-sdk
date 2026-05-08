// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseListResult interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBases(v []*KnowledgeBaseInfo) *KnowledgeBaseListResult
	GetKnowledgeBases() []*KnowledgeBaseInfo
	SetRequestId(v string) *KnowledgeBaseListResult
	GetRequestId() *string
	SetTotal(v int32) *KnowledgeBaseListResult
	GetTotal() *int32
}

type KnowledgeBaseListResult struct {
	KnowledgeBases []*KnowledgeBaseInfo `json:"knowledgeBases,omitempty" xml:"knowledgeBases,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s KnowledgeBaseListResult) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseListResult) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseListResult) GetKnowledgeBases() []*KnowledgeBaseInfo {
	return s.KnowledgeBases
}

func (s *KnowledgeBaseListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *KnowledgeBaseListResult) GetTotal() *int32 {
	return s.Total
}

func (s *KnowledgeBaseListResult) SetKnowledgeBases(v []*KnowledgeBaseInfo) *KnowledgeBaseListResult {
	s.KnowledgeBases = v
	return s
}

func (s *KnowledgeBaseListResult) SetRequestId(v string) *KnowledgeBaseListResult {
	s.RequestId = &v
	return s
}

func (s *KnowledgeBaseListResult) SetTotal(v int32) *KnowledgeBaseListResult {
	s.Total = &v
	return s
}

func (s *KnowledgeBaseListResult) Validate() error {
	if s.KnowledgeBases != nil {
		for _, item := range s.KnowledgeBases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
