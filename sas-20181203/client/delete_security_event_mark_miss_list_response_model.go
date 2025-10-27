// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityEventMarkMissListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityEventMarkMissListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityEventMarkMissListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityEventMarkMissListResponseBody) *DeleteSecurityEventMarkMissListResponse
	GetBody() *DeleteSecurityEventMarkMissListResponseBody
}

type DeleteSecurityEventMarkMissListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityEventMarkMissListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityEventMarkMissListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityEventMarkMissListResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityEventMarkMissListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityEventMarkMissListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityEventMarkMissListResponse) GetBody() *DeleteSecurityEventMarkMissListResponseBody {
	return s.Body
}

func (s *DeleteSecurityEventMarkMissListResponse) SetHeaders(v map[string]*string) *DeleteSecurityEventMarkMissListResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityEventMarkMissListResponse) SetStatusCode(v int32) *DeleteSecurityEventMarkMissListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityEventMarkMissListResponse) SetBody(v *DeleteSecurityEventMarkMissListResponseBody) *DeleteSecurityEventMarkMissListResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityEventMarkMissListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
