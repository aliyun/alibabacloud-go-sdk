// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityClassifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityClassifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityClassifyResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityClassifyResponseBody) *CreateSecurityClassifyResponse
	GetBody() *CreateSecurityClassifyResponseBody
}

type CreateSecurityClassifyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityClassifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityClassifyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityClassifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityClassifyResponse) GetBody() *CreateSecurityClassifyResponseBody {
	return s.Body
}

func (s *CreateSecurityClassifyResponse) SetHeaders(v map[string]*string) *CreateSecurityClassifyResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityClassifyResponse) SetStatusCode(v int32) *CreateSecurityClassifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityClassifyResponse) SetBody(v *CreateSecurityClassifyResponseBody) *CreateSecurityClassifyResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityClassifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
