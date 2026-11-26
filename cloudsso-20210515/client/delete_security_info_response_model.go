// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityInfoResponseBody) *DeleteSecurityInfoResponse
	GetBody() *DeleteSecurityInfoResponseBody
}

type DeleteSecurityInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityInfoResponse) GetBody() *DeleteSecurityInfoResponseBody {
	return s.Body
}

func (s *DeleteSecurityInfoResponse) SetHeaders(v map[string]*string) *DeleteSecurityInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityInfoResponse) SetStatusCode(v int32) *DeleteSecurityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityInfoResponse) SetBody(v *DeleteSecurityInfoResponseBody) *DeleteSecurityInfoResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
