// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerFromClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeResourceServerFromClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeResourceServerFromClientResponse
	GetStatusCode() *int32
	SetBody(v *RevokeResourceServerFromClientResponseBody) *RevokeResourceServerFromClientResponse
	GetBody() *RevokeResourceServerFromClientResponseBody
}

type RevokeResourceServerFromClientResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResourceServerFromClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResourceServerFromClientResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerFromClientResponse) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerFromClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeResourceServerFromClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeResourceServerFromClientResponse) GetBody() *RevokeResourceServerFromClientResponseBody {
	return s.Body
}

func (s *RevokeResourceServerFromClientResponse) SetHeaders(v map[string]*string) *RevokeResourceServerFromClientResponse {
	s.Headers = v
	return s
}

func (s *RevokeResourceServerFromClientResponse) SetStatusCode(v int32) *RevokeResourceServerFromClientResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResourceServerFromClientResponse) SetBody(v *RevokeResourceServerFromClientResponseBody) *RevokeResourceServerFromClientResponse {
	s.Body = v
	return s
}

func (s *RevokeResourceServerFromClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
