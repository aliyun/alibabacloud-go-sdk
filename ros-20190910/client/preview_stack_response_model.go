// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewStackResponse
	GetStatusCode() *int32
	SetBody(v *PreviewStackResponseBody) *PreviewStackResponse
	GetBody() *PreviewStackResponseBody
}

type PreviewStackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewStackResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackResponse) GoString() string {
	return s.String()
}

func (s *PreviewStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewStackResponse) GetBody() *PreviewStackResponseBody {
	return s.Body
}

func (s *PreviewStackResponse) SetHeaders(v map[string]*string) *PreviewStackResponse {
	s.Headers = v
	return s
}

func (s *PreviewStackResponse) SetStatusCode(v int32) *PreviewStackResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewStackResponse) SetBody(v *PreviewStackResponseBody) *PreviewStackResponse {
	s.Body = v
	return s
}

func (s *PreviewStackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
