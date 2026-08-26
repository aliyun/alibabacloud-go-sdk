// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrustedOriginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrustedOriginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrustedOriginResponse
	GetStatusCode() *int32
	SetBody(v *GetTrustedOriginResponseBody) *GetTrustedOriginResponse
	GetBody() *GetTrustedOriginResponseBody
}

type GetTrustedOriginResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrustedOriginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrustedOriginResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrustedOriginResponse) GoString() string {
	return s.String()
}

func (s *GetTrustedOriginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrustedOriginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrustedOriginResponse) GetBody() *GetTrustedOriginResponseBody {
	return s.Body
}

func (s *GetTrustedOriginResponse) SetHeaders(v map[string]*string) *GetTrustedOriginResponse {
	s.Headers = v
	return s
}

func (s *GetTrustedOriginResponse) SetStatusCode(v int32) *GetTrustedOriginResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrustedOriginResponse) SetBody(v *GetTrustedOriginResponseBody) *GetTrustedOriginResponse {
	s.Body = v
	return s
}

func (s *GetTrustedOriginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
