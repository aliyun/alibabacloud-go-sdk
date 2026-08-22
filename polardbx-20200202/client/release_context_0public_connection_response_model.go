// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseContext0PublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseContext0PublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseContext0PublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseContext0PublicConnectionResponseBody) *ReleaseContext0PublicConnectionResponse
	GetBody() *ReleaseContext0PublicConnectionResponseBody
}

type ReleaseContext0PublicConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseContext0PublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseContext0PublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContext0PublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseContext0PublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseContext0PublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseContext0PublicConnectionResponse) GetBody() *ReleaseContext0PublicConnectionResponseBody {
	return s.Body
}

func (s *ReleaseContext0PublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseContext0PublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseContext0PublicConnectionResponse) SetStatusCode(v int32) *ReleaseContext0PublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponse) SetBody(v *ReleaseContext0PublicConnectionResponseBody) *ReleaseContext0PublicConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleaseContext0PublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
