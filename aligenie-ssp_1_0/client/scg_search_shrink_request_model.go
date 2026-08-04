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
	// Query filter
	//
	// This parameter is required.
	ScgFilterShrink *string `json:"ScgFilter,omitempty" xml:"ScgFilter,omitempty"`
	// Selection pool ID. Optional values: MC201132 (Ethnic Chinese Style), MC201136 (Pop Music), MC201139 (Sweet Love), MC201133 (Folk), MC201137 (Relaxing Reading), MC201138 (Happiness), PA202029 (Stories), PA202030 (Children\\"s Songs), PA202028 (Chinese Classics and History), PA202032 (Encyclopedia), PA202031 (English Children\\"s Songs)
	//
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
