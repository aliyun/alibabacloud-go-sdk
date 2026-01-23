// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityClassifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityClassifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityClassifyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityClassifyResponseBody) *DeleteSecurityClassifyResponse
	GetBody() *DeleteSecurityClassifyResponseBody
}

type DeleteSecurityClassifyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityClassifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityClassifyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityClassifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityClassifyResponse) GetBody() *DeleteSecurityClassifyResponseBody {
	return s.Body
}

func (s *DeleteSecurityClassifyResponse) SetHeaders(v map[string]*string) *DeleteSecurityClassifyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityClassifyResponse) SetStatusCode(v int32) *DeleteSecurityClassifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityClassifyResponse) SetBody(v *DeleteSecurityClassifyResponseBody) *DeleteSecurityClassifyResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityClassifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
