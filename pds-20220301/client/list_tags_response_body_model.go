// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTags(v []*ImageTag) *ListTagsResponseBody
	GetTags() []*ImageTag
}

type ListTagsResponseBody struct {
	// The information about the tags.
	Tags []*ImageTag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) GetTags() []*ImageTag {
	return s.Tags
}

func (s *ListTagsResponseBody) SetTags(v []*ImageTag) *ListTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListTagsResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
