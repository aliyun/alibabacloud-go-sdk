// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewMCTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewMCTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewMCTableResponse
	GetStatusCode() *int32
	SetBody(v *PreviewMCTableResponseBody) *PreviewMCTableResponse
	GetBody() *PreviewMCTableResponseBody
}

type PreviewMCTableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewMCTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewMCTableResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewMCTableResponse) GoString() string {
	return s.String()
}

func (s *PreviewMCTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewMCTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewMCTableResponse) GetBody() *PreviewMCTableResponseBody {
	return s.Body
}

func (s *PreviewMCTableResponse) SetHeaders(v map[string]*string) *PreviewMCTableResponse {
	s.Headers = v
	return s
}

func (s *PreviewMCTableResponse) SetStatusCode(v int32) *PreviewMCTableResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewMCTableResponse) SetBody(v *PreviewMCTableResponseBody) *PreviewMCTableResponse {
	s.Body = v
	return s
}

func (s *PreviewMCTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
