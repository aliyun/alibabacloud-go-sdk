// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScgSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScgFilterShrink(v string) *ScgSearchShrinkRequest
	GetScgFilterShrink() *string
	SetTopicId(v string) *ScgSearchShrinkRequest
	GetTopicId() *string
}

type ScgSearchShrinkRequest struct {
	// This parameter is required.
	ScgFilterShrink *string `json:"ScgFilter,omitempty" xml:"ScgFilter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MC201132
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ScgSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScgSearchShrinkRequest) GetScgFilterShrink() *string {
	return s.ScgFilterShrink
}

func (s *ScgSearchShrinkRequest) GetTopicId() *string {
	return s.TopicId
}

func (s *ScgSearchShrinkRequest) SetScgFilterShrink(v string) *ScgSearchShrinkRequest {
	s.ScgFilterShrink = &v
	return s
}

func (s *ScgSearchShrinkRequest) SetTopicId(v string) *ScgSearchShrinkRequest {
	s.TopicId = &v
	return s
}

func (s *ScgSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
