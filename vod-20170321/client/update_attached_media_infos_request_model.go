// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttachedMediaInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateContent(v string) *UpdateAttachedMediaInfosRequest
	GetUpdateContent() *string
}

type UpdateAttachedMediaInfosRequest struct {
	// The new information about the one or more images. You can modify the information about up to 20 auxiliary media assets at a time. For more information, see the **UpdateContent*	- section of this topic.
	//
	// > 	- You cannot specify emojis for `Title`, `Description`, or `Tags`.
	//
	// > 	- The specific parameter of a video is updated only when a new value is passed in the parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"MediaId ":"bbc65bba53f6ed90de118a7849****","Title":" test title1","Description":"test description1","Tags":"tag1,tag2"},{"MediaId ":"f45cf4eba5cb90233389558c39****","Title2":"test title2","Description2":"test description2","Tags":"tag3,tag4"}]
	UpdateContent *string `json:"UpdateContent,omitempty" xml:"UpdateContent,omitempty"`
}

func (s UpdateAttachedMediaInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttachedMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *UpdateAttachedMediaInfosRequest) GetUpdateContent() *string {
	return s.UpdateContent
}

func (s *UpdateAttachedMediaInfosRequest) SetUpdateContent(v string) *UpdateAttachedMediaInfosRequest {
	s.UpdateContent = &v
	return s
}

func (s *UpdateAttachedMediaInfosRequest) Validate() error {
	return dara.Validate(s)
}
