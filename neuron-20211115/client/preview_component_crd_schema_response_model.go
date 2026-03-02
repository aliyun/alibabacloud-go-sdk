// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewComponentCrdSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewComponentCrdSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewComponentCrdSchemaResponse
	GetStatusCode() *int32
	SetBody(v *PreviewComponentCrdSchemaResponseBody) *PreviewComponentCrdSchemaResponse
	GetBody() *PreviewComponentCrdSchemaResponseBody
}

type PreviewComponentCrdSchemaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewComponentCrdSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewComponentCrdSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewComponentCrdSchemaResponse) GoString() string {
	return s.String()
}

func (s *PreviewComponentCrdSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewComponentCrdSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewComponentCrdSchemaResponse) GetBody() *PreviewComponentCrdSchemaResponseBody {
	return s.Body
}

func (s *PreviewComponentCrdSchemaResponse) SetHeaders(v map[string]*string) *PreviewComponentCrdSchemaResponse {
	s.Headers = v
	return s
}

func (s *PreviewComponentCrdSchemaResponse) SetStatusCode(v int32) *PreviewComponentCrdSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewComponentCrdSchemaResponse) SetBody(v *PreviewComponentCrdSchemaResponseBody) *PreviewComponentCrdSchemaResponse {
	s.Body = v
	return s
}

func (s *PreviewComponentCrdSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
