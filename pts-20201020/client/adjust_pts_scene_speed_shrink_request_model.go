// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustPtsSceneSpeedShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiSpeedListShrink(v string) *AdjustPtsSceneSpeedShrinkRequest
	GetApiSpeedListShrink() *string
	SetSceneId(v string) *AdjustPtsSceneSpeedShrinkRequest
	GetSceneId() *string
}

type AdjustPtsSceneSpeedShrinkRequest struct {
	// The stress testing speed in the PTS scenario.
	ApiSpeedListShrink *string `json:"ApiSpeedList,omitempty" xml:"ApiSpeedList,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYXXX12H
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s AdjustPtsSceneSpeedShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AdjustPtsSceneSpeedShrinkRequest) GoString() string {
	return s.String()
}

func (s *AdjustPtsSceneSpeedShrinkRequest) GetApiSpeedListShrink() *string {
	return s.ApiSpeedListShrink
}

func (s *AdjustPtsSceneSpeedShrinkRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *AdjustPtsSceneSpeedShrinkRequest) SetApiSpeedListShrink(v string) *AdjustPtsSceneSpeedShrinkRequest {
	s.ApiSpeedListShrink = &v
	return s
}

func (s *AdjustPtsSceneSpeedShrinkRequest) SetSceneId(v string) *AdjustPtsSceneSpeedShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *AdjustPtsSceneSpeedShrinkRequest) Validate() error {
	return dara.Validate(s)
}
