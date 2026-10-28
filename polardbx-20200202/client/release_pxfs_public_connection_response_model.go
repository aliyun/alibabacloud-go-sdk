// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePxfsPublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleasePxfsPublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleasePxfsPublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleasePxfsPublicConnectionResponseBody) *ReleasePxfsPublicConnectionResponse
	GetBody() *ReleasePxfsPublicConnectionResponseBody
}

type ReleasePxfsPublicConnectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePxfsPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePxfsPublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleasePxfsPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleasePxfsPublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleasePxfsPublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleasePxfsPublicConnectionResponse) GetBody() *ReleasePxfsPublicConnectionResponseBody {
	return s.Body
}

func (s *ReleasePxfsPublicConnectionResponse) SetHeaders(v map[string]*string) *ReleasePxfsPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleasePxfsPublicConnectionResponse) SetStatusCode(v int32) *ReleasePxfsPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponse) SetBody(v *ReleasePxfsPublicConnectionResponseBody) *ReleasePxfsPublicConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleasePxfsPublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
