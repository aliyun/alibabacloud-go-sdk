// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageToVideoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *SubmitImageToVideoTaskRequest
	GetImageUrl() *string
	SetPosPrompt(v string) *SubmitImageToVideoTaskRequest
	GetPosPrompt() *string
}

type SubmitImageToVideoTaskRequest struct {
	// example:
	//
	// http://xxx/image.png
	ImageUrl  *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	PosPrompt *string `json:"posPrompt,omitempty" xml:"posPrompt,omitempty"`
}

func (s SubmitImageToVideoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageToVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageToVideoTaskRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SubmitImageToVideoTaskRequest) GetPosPrompt() *string {
	return s.PosPrompt
}

func (s *SubmitImageToVideoTaskRequest) SetImageUrl(v string) *SubmitImageToVideoTaskRequest {
	s.ImageUrl = &v
	return s
}

func (s *SubmitImageToVideoTaskRequest) SetPosPrompt(v string) *SubmitImageToVideoTaskRequest {
	s.PosPrompt = &v
	return s
}

func (s *SubmitImageToVideoTaskRequest) Validate() error {
	return dara.Validate(s)
}
