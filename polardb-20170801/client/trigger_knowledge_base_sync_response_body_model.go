// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerKnowledgeBaseSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerKnowledgeBaseSyncResponseBody
	GetRequestId() *string
}

type TriggerKnowledgeBaseSyncResponseBody struct {
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerKnowledgeBaseSyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerKnowledgeBaseSyncResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerKnowledgeBaseSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerKnowledgeBaseSyncResponseBody) SetRequestId(v string) *TriggerKnowledgeBaseSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerKnowledgeBaseSyncResponseBody) Validate() error {
	return dara.Validate(s)
}
