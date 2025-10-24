// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScreenshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScreenshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScreenshotResponse
	GetStatusCode() *int32
	SetBody(v *CreateScreenshotResponseBody) *CreateScreenshotResponse
	GetBody() *CreateScreenshotResponseBody
}

type CreateScreenshotResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScreenshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScreenshotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScreenshotResponse) GoString() string {
	return s.String()
}

func (s *CreateScreenshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScreenshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScreenshotResponse) GetBody() *CreateScreenshotResponseBody {
	return s.Body
}

func (s *CreateScreenshotResponse) SetHeaders(v map[string]*string) *CreateScreenshotResponse {
	s.Headers = v
	return s
}

func (s *CreateScreenshotResponse) SetStatusCode(v int32) *CreateScreenshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScreenshotResponse) SetBody(v *CreateScreenshotResponseBody) *CreateScreenshotResponse {
	s.Body = v
	return s
}

func (s *CreateScreenshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
