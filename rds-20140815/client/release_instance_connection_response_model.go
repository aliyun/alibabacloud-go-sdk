// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseInstanceConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseInstanceConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseInstanceConnectionResponseBody) *ReleaseInstanceConnectionResponse
	GetBody() *ReleaseInstanceConnectionResponseBody
}

type ReleaseInstanceConnectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstanceConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseInstanceConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseInstanceConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseInstanceConnectionResponse) GetBody() *ReleaseInstanceConnectionResponseBody {
	return s.Body
}

func (s *ReleaseInstanceConnectionResponse) SetHeaders(v map[string]*string) *ReleaseInstanceConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceConnectionResponse) SetStatusCode(v int32) *ReleaseInstanceConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceConnectionResponse) SetBody(v *ReleaseInstanceConnectionResponseBody) *ReleaseInstanceConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleaseInstanceConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
