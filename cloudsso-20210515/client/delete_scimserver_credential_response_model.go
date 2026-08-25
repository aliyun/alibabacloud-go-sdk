// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSCIMServerCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSCIMServerCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSCIMServerCredentialResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSCIMServerCredentialResponseBody) *DeleteSCIMServerCredentialResponse
	GetBody() *DeleteSCIMServerCredentialResponseBody
}

type DeleteSCIMServerCredentialResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSCIMServerCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSCIMServerCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSCIMServerCredentialResponse) GoString() string {
	return s.String()
}

func (s *DeleteSCIMServerCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSCIMServerCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSCIMServerCredentialResponse) GetBody() *DeleteSCIMServerCredentialResponseBody {
	return s.Body
}

func (s *DeleteSCIMServerCredentialResponse) SetHeaders(v map[string]*string) *DeleteSCIMServerCredentialResponse {
	s.Headers = v
	return s
}

func (s *DeleteSCIMServerCredentialResponse) SetStatusCode(v int32) *DeleteSCIMServerCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSCIMServerCredentialResponse) SetBody(v *DeleteSCIMServerCredentialResponseBody) *DeleteSCIMServerCredentialResponse {
	s.Body = v
	return s
}

func (s *DeleteSCIMServerCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
