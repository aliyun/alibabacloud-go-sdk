// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryIndividuationTextShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTextIdListShrink(v string) *BatchQueryIndividuationTextShrinkRequest
	GetTextIdListShrink() *string
}

type BatchQueryIndividuationTextShrinkRequest struct {
	TextIdListShrink *string `json:"textIdList,omitempty" xml:"textIdList,omitempty"`
}

func (s BatchQueryIndividuationTextShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryIndividuationTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextShrinkRequest) GetTextIdListShrink() *string {
	return s.TextIdListShrink
}

func (s *BatchQueryIndividuationTextShrinkRequest) SetTextIdListShrink(v string) *BatchQueryIndividuationTextShrinkRequest {
	s.TextIdListShrink = &v
	return s
}

func (s *BatchQueryIndividuationTextShrinkRequest) Validate() error {
	return dara.Validate(s)
}
