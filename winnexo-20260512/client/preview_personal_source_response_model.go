// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPersonalSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewPersonalSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewPersonalSourceResponse
	GetStatusCode() *int32
	SetBody(v *PreviewPersonalSourceResponseBody) *PreviewPersonalSourceResponse
	GetBody() *PreviewPersonalSourceResponseBody
}

type PreviewPersonalSourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewPersonalSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewPersonalSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewPersonalSourceResponse) GoString() string {
	return s.String()
}

func (s *PreviewPersonalSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewPersonalSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewPersonalSourceResponse) GetBody() *PreviewPersonalSourceResponseBody {
	return s.Body
}

func (s *PreviewPersonalSourceResponse) SetHeaders(v map[string]*string) *PreviewPersonalSourceResponse {
	s.Headers = v
	return s
}

func (s *PreviewPersonalSourceResponse) SetStatusCode(v int32) *PreviewPersonalSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewPersonalSourceResponse) SetBody(v *PreviewPersonalSourceResponseBody) *PreviewPersonalSourceResponse {
	s.Body = v
	return s
}

func (s *PreviewPersonalSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
