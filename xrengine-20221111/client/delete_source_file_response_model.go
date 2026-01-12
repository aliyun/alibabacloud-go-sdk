// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSourceFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSourceFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSourceFileResponseBody) *DeleteSourceFileResponse
	GetBody() *DeleteSourceFileResponseBody
}

type DeleteSourceFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSourceFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSourceFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteSourceFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSourceFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSourceFileResponse) GetBody() *DeleteSourceFileResponseBody {
	return s.Body
}

func (s *DeleteSourceFileResponse) SetHeaders(v map[string]*string) *DeleteSourceFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteSourceFileResponse) SetStatusCode(v int32) *DeleteSourceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSourceFileResponse) SetBody(v *DeleteSourceFileResponseBody) *DeleteSourceFileResponse {
	s.Body = v
	return s
}

func (s *DeleteSourceFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
