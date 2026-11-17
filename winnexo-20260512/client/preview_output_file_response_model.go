// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewOutputFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewOutputFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewOutputFileResponse
	GetStatusCode() *int32
	SetBody(v *PreviewOutputFileResponseBody) *PreviewOutputFileResponse
	GetBody() *PreviewOutputFileResponseBody
}

type PreviewOutputFileResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewOutputFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewOutputFileResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewOutputFileResponse) GoString() string {
	return s.String()
}

func (s *PreviewOutputFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewOutputFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewOutputFileResponse) GetBody() *PreviewOutputFileResponseBody {
	return s.Body
}

func (s *PreviewOutputFileResponse) SetHeaders(v map[string]*string) *PreviewOutputFileResponse {
	s.Headers = v
	return s
}

func (s *PreviewOutputFileResponse) SetStatusCode(v int32) *PreviewOutputFileResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewOutputFileResponse) SetBody(v *PreviewOutputFileResponseBody) *PreviewOutputFileResponse {
	s.Body = v
	return s
}

func (s *PreviewOutputFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
