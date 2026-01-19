// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeResourceServerScopesFromClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeResourceServerScopesFromClientResponse
	GetStatusCode() *int32
	SetBody(v *RevokeResourceServerScopesFromClientResponseBody) *RevokeResourceServerScopesFromClientResponse
	GetBody() *RevokeResourceServerScopesFromClientResponseBody
}

type RevokeResourceServerScopesFromClientResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResourceServerScopesFromClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResourceServerScopesFromClientResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromClientResponse) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeResourceServerScopesFromClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeResourceServerScopesFromClientResponse) GetBody() *RevokeResourceServerScopesFromClientResponseBody {
	return s.Body
}

func (s *RevokeResourceServerScopesFromClientResponse) SetHeaders(v map[string]*string) *RevokeResourceServerScopesFromClientResponse {
	s.Headers = v
	return s
}

func (s *RevokeResourceServerScopesFromClientResponse) SetStatusCode(v int32) *RevokeResourceServerScopesFromClientResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResourceServerScopesFromClientResponse) SetBody(v *RevokeResourceServerScopesFromClientResponseBody) *RevokeResourceServerScopesFromClientResponse {
	s.Body = v
	return s
}

func (s *RevokeResourceServerScopesFromClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
