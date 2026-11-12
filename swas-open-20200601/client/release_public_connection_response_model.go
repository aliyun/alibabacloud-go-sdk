// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleasePublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleasePublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleasePublicConnectionResponseBody) *ReleasePublicConnectionResponse
	GetBody() *ReleasePublicConnectionResponseBody
}

type ReleasePublicConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleasePublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleasePublicConnectionResponse) GetBody() *ReleasePublicConnectionResponseBody {
	return s.Body
}

func (s *ReleasePublicConnectionResponse) SetHeaders(v map[string]*string) *ReleasePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicConnectionResponse) SetStatusCode(v int32) *ReleasePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePublicConnectionResponse) SetBody(v *ReleasePublicConnectionResponseBody) *ReleasePublicConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleasePublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
