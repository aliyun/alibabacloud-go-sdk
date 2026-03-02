// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewComponentCrdSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PreviewComponentCrdSchemaCmd) *PreviewComponentCrdSchemaRequest
	GetBody() *PreviewComponentCrdSchemaCmd
}

type PreviewComponentCrdSchemaRequest struct {
	Body *PreviewComponentCrdSchemaCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewComponentCrdSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewComponentCrdSchemaRequest) GoString() string {
	return s.String()
}

func (s *PreviewComponentCrdSchemaRequest) GetBody() *PreviewComponentCrdSchemaCmd {
	return s.Body
}

func (s *PreviewComponentCrdSchemaRequest) SetBody(v *PreviewComponentCrdSchemaCmd) *PreviewComponentCrdSchemaRequest {
	s.Body = v
	return s
}

func (s *PreviewComponentCrdSchemaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
