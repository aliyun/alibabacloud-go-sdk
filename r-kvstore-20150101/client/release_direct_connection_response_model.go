// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDirectConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseDirectConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseDirectConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseDirectConnectionResponseBody) *ReleaseDirectConnectionResponse
	GetBody() *ReleaseDirectConnectionResponseBody
}

type ReleaseDirectConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseDirectConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseDirectConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDirectConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseDirectConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseDirectConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseDirectConnectionResponse) GetBody() *ReleaseDirectConnectionResponseBody {
	return s.Body
}

func (s *ReleaseDirectConnectionResponse) SetHeaders(v map[string]*string) *ReleaseDirectConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseDirectConnectionResponse) SetStatusCode(v int32) *ReleaseDirectConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseDirectConnectionResponse) SetBody(v *ReleaseDirectConnectionResponseBody) *ReleaseDirectConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleaseDirectConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
