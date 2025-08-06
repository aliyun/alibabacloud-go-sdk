// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateContent(v string) *UpdateVideoInfosRequest
	GetUpdateContent() *string
}

type UpdateVideoInfosRequest struct {
	// The new information about audios or videos. You can modify the information about up to 20 audios or videos at a time. Separate multiple audios or videos with commas (,). When you modify the information exceed 20 audios or videos at a time, the update will fail with an error code **CountExceededMax**.
	//
	// The value is a JSON string. For more information, see the **UpdateContent*	- section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"VideoId":"f45cf4eba5cb90233389558c39****","Title":"test title1"},{"VideoId":"f45cf4eba5c84233389558c36****","Title":"test title2"}]
	UpdateContent *string `json:"UpdateContent,omitempty" xml:"UpdateContent,omitempty"`
}

func (s UpdateVideoInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoInfosRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfosRequest) GetUpdateContent() *string {
	return s.UpdateContent
}

func (s *UpdateVideoInfosRequest) SetUpdateContent(v string) *UpdateVideoInfosRequest {
	s.UpdateContent = &v
	return s
}

func (s *UpdateVideoInfosRequest) Validate() error {
	return dara.Validate(s)
}
