// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseContextDBPublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseContextDBPublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseContextDBPublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseContextDBPublicConnectionResponseBody) *ReleaseContextDBPublicConnectionResponse
	GetBody() *ReleaseContextDBPublicConnectionResponseBody
}

type ReleaseContextDBPublicConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseContextDBPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseContextDBPublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContextDBPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseContextDBPublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseContextDBPublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseContextDBPublicConnectionResponse) GetBody() *ReleaseContextDBPublicConnectionResponseBody {
	return s.Body
}

func (s *ReleaseContextDBPublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseContextDBPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponse) SetStatusCode(v int32) *ReleaseContextDBPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponse) SetBody(v *ReleaseContextDBPublicConnectionResponseBody) *ReleaseContextDBPublicConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
