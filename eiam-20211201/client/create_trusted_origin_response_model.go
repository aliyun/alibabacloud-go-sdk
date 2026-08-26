// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrustedOriginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrustedOriginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrustedOriginResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrustedOriginResponseBody) *CreateTrustedOriginResponse
	GetBody() *CreateTrustedOriginResponseBody
}

type CreateTrustedOriginResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrustedOriginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrustedOriginResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrustedOriginResponse) GoString() string {
	return s.String()
}

func (s *CreateTrustedOriginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrustedOriginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrustedOriginResponse) GetBody() *CreateTrustedOriginResponseBody {
	return s.Body
}

func (s *CreateTrustedOriginResponse) SetHeaders(v map[string]*string) *CreateTrustedOriginResponse {
	s.Headers = v
	return s
}

func (s *CreateTrustedOriginResponse) SetStatusCode(v int32) *CreateTrustedOriginResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrustedOriginResponse) SetBody(v *CreateTrustedOriginResponseBody) *CreateTrustedOriginResponse {
	s.Body = v
	return s
}

func (s *CreateTrustedOriginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
