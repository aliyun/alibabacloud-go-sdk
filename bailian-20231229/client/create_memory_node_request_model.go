// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateMemoryNodeRequest
	GetContent() *string
}

type CreateMemoryNodeRequest struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CreateMemoryNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryNodeRequest) GetContent() *string {
	return s.Content
}

func (s *CreateMemoryNodeRequest) SetContent(v string) *CreateMemoryNodeRequest {
	s.Content = &v
	return s
}

func (s *CreateMemoryNodeRequest) Validate() error {
	return dara.Validate(s)
}
