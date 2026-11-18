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
	// The update content. You can modify object information of up to 20 audio or video files at a time. Separate object information of multiple audio or video files with commas (,). If you attempt to update more than 20 files at a time, the update fails and the `CountExceededMax` error is returned.
	//
	// The value must be a JSON string. For more information about how to construct the parameters, see the **UpdateContent*	- table below.
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
