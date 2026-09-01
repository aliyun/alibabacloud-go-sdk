// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseFileShardingStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKnowledgeBaseFileShardingStrategyResponseBody
	GetRequestId() *string
}

type UpdateKnowledgeBaseFileShardingStrategyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 019F7F29-BF69-1734-AE5A-02D391385BAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseFileShardingStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
