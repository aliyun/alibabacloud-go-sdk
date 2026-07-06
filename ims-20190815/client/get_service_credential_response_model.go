// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceCredentialResponseBody) *GetServiceCredentialResponse
	GetBody() *GetServiceCredentialResponseBody
}

type GetServiceCredentialResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetServiceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceCredentialResponse) GetBody() *GetServiceCredentialResponseBody {
	return s.Body
}

func (s *GetServiceCredentialResponse) SetHeaders(v map[string]*string) *GetServiceCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetServiceCredentialResponse) SetStatusCode(v int32) *GetServiceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceCredentialResponse) SetBody(v *GetServiceCredentialResponseBody) *GetServiceCredentialResponse {
	s.Body = v
	return s
}

func (s *GetServiceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
