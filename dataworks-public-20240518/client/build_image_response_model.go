// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuildImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuildImageResponse
	GetStatusCode() *int32
	SetBody(v *BuildImageResponseBody) *BuildImageResponse
	GetBody() *BuildImageResponseBody
}

type BuildImageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuildImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildImageResponse) String() string {
	return dara.Prettify(s)
}

func (s BuildImageResponse) GoString() string {
	return s.String()
}

func (s *BuildImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuildImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuildImageResponse) GetBody() *BuildImageResponseBody {
	return s.Body
}

func (s *BuildImageResponse) SetHeaders(v map[string]*string) *BuildImageResponse {
	s.Headers = v
	return s
}

func (s *BuildImageResponse) SetStatusCode(v int32) *BuildImageResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildImageResponse) SetBody(v *BuildImageResponseBody) *BuildImageResponse {
	s.Body = v
	return s
}

func (s *BuildImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
