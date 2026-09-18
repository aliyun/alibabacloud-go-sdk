// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewGroupSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewGroupSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewGroupSourceResponse
	GetStatusCode() *int32
	SetBody(v *PreviewGroupSourceResponseBody) *PreviewGroupSourceResponse
	GetBody() *PreviewGroupSourceResponseBody
}

type PreviewGroupSourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewGroupSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewGroupSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewGroupSourceResponse) GoString() string {
	return s.String()
}

func (s *PreviewGroupSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewGroupSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewGroupSourceResponse) GetBody() *PreviewGroupSourceResponseBody {
	return s.Body
}

func (s *PreviewGroupSourceResponse) SetHeaders(v map[string]*string) *PreviewGroupSourceResponse {
	s.Headers = v
	return s
}

func (s *PreviewGroupSourceResponse) SetStatusCode(v int32) *PreviewGroupSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewGroupSourceResponse) SetBody(v *PreviewGroupSourceResponseBody) *PreviewGroupSourceResponse {
	s.Body = v
	return s
}

func (s *PreviewGroupSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
