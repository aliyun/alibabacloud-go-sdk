// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotNewsRecommendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *HotNewsRecommendRequest
	GetPrompt() *string
}

type HotNewsRecommendRequest struct {
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
}

func (s HotNewsRecommendRequest) String() string {
	return dara.Prettify(s)
}

func (s HotNewsRecommendRequest) GoString() string {
	return s.String()
}

func (s *HotNewsRecommendRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *HotNewsRecommendRequest) SetPrompt(v string) *HotNewsRecommendRequest {
	s.Prompt = &v
	return s
}

func (s *HotNewsRecommendRequest) Validate() error {
	return dara.Validate(s)
}
