// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrustedOriginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrustedOriginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrustedOriginResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrustedOriginResponseBody) *UpdateTrustedOriginResponse
	GetBody() *UpdateTrustedOriginResponseBody
}

type UpdateTrustedOriginResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrustedOriginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrustedOriginResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrustedOriginResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrustedOriginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrustedOriginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrustedOriginResponse) GetBody() *UpdateTrustedOriginResponseBody {
	return s.Body
}

func (s *UpdateTrustedOriginResponse) SetHeaders(v map[string]*string) *UpdateTrustedOriginResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrustedOriginResponse) SetStatusCode(v int32) *UpdateTrustedOriginResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrustedOriginResponse) SetBody(v *UpdateTrustedOriginResponseBody) *UpdateTrustedOriginResponse {
	s.Body = v
	return s
}

func (s *UpdateTrustedOriginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
