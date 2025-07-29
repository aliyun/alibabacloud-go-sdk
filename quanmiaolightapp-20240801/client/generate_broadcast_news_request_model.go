// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateBroadcastNewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *GenerateBroadcastNewsRequest
	GetPrompt() *string
}

type GenerateBroadcastNewsRequest struct {
	// This parameter is required.
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
}

func (s GenerateBroadcastNewsRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateBroadcastNewsRequest) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *GenerateBroadcastNewsRequest) SetPrompt(v string) *GenerateBroadcastNewsRequest {
	s.Prompt = &v
	return s
}

func (s *GenerateBroadcastNewsRequest) Validate() error {
	return dara.Validate(s)
}
