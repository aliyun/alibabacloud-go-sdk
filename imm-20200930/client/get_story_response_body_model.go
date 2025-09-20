// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStoryResponseBody
	GetRequestId() *string
	SetStory(v *Story) *GetStoryResponseBody
	GetStory() *Story
}

type GetStoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B3D5E0A-D8B8-4DA0-8127-ED32C851****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the story.
	Story *Story `json:"Story,omitempty" xml:"Story,omitempty"`
}

func (s GetStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetStoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStoryResponseBody) GetStory() *Story {
	return s.Story
}

func (s *GetStoryResponseBody) SetRequestId(v string) *GetStoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStoryResponseBody) SetStory(v *Story) *GetStoryResponseBody {
	s.Story = v
	return s
}

func (s *GetStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
