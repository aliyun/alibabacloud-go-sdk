// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOriginProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOriginProtectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOriginProtectionResponseBody) *DeleteOriginProtectionResponse
	GetBody() *DeleteOriginProtectionResponseBody
}

type DeleteOriginProtectionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOriginProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOriginProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginProtectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteOriginProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOriginProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOriginProtectionResponse) GetBody() *DeleteOriginProtectionResponseBody {
	return s.Body
}

func (s *DeleteOriginProtectionResponse) SetHeaders(v map[string]*string) *DeleteOriginProtectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteOriginProtectionResponse) SetStatusCode(v int32) *DeleteOriginProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOriginProtectionResponse) SetBody(v *DeleteOriginProtectionResponseBody) *DeleteOriginProtectionResponse {
	s.Body = v
	return s
}

func (s *DeleteOriginProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
