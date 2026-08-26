// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrustedOriginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrustedOriginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrustedOriginResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrustedOriginResponseBody) *DeleteTrustedOriginResponse
	GetBody() *DeleteTrustedOriginResponseBody
}

type DeleteTrustedOriginResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrustedOriginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrustedOriginResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrustedOriginResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrustedOriginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrustedOriginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrustedOriginResponse) GetBody() *DeleteTrustedOriginResponseBody {
	return s.Body
}

func (s *DeleteTrustedOriginResponse) SetHeaders(v map[string]*string) *DeleteTrustedOriginResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrustedOriginResponse) SetStatusCode(v int32) *DeleteTrustedOriginResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrustedOriginResponse) SetBody(v *DeleteTrustedOriginResponseBody) *DeleteTrustedOriginResponse {
	s.Body = v
	return s
}

func (s *DeleteTrustedOriginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
