// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeniedMemberAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeniedMemberAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeniedMemberAccessResponse
	GetStatusCode() *int32
	SetBody(v *DeniedMemberAccessResponseBody) *DeniedMemberAccessResponse
	GetBody() *DeniedMemberAccessResponseBody
}

type DeniedMemberAccessResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeniedMemberAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeniedMemberAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeniedMemberAccessResponse) GoString() string {
	return s.String()
}

func (s *DeniedMemberAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeniedMemberAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeniedMemberAccessResponse) GetBody() *DeniedMemberAccessResponseBody {
	return s.Body
}

func (s *DeniedMemberAccessResponse) SetHeaders(v map[string]*string) *DeniedMemberAccessResponse {
	s.Headers = v
	return s
}

func (s *DeniedMemberAccessResponse) SetStatusCode(v int32) *DeniedMemberAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeniedMemberAccessResponse) SetBody(v *DeniedMemberAccessResponseBody) *DeniedMemberAccessResponse {
	s.Body = v
	return s
}

func (s *DeniedMemberAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
