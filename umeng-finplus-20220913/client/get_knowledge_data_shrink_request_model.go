// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *GetKnowledgeDataShrinkRequest
	GetBodyShrink() *string
}

type GetKnowledgeDataShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeDataShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *GetKnowledgeDataShrinkRequest) SetBodyShrink(v string) *GetKnowledgeDataShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *GetKnowledgeDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
