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
	// The update content. You can modify the information about up to 20 audio and video files at a time. Separate multiple audio and video object information entries with commas (,). If you specify more than 20 objects, the update is failed and the `CountExceededMax` error is returned.
	//
	// The value is a JSON character string. For more details about the parameters, see the **UpdateContent*	- table below.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"VideoId":"f45cf4eba5cb90233389558c39****","Title":"Alibaba Cloud VOD Video Title1"},{"VideoId":"f45cf4eba5c84233389558c36****","Title":"Alibaba Cloud VOD Video Title2"}]
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
