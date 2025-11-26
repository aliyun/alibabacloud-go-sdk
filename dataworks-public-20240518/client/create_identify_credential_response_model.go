// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentifyCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIdentifyCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIdentifyCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CreateIdentifyCredentialResponseBody) *CreateIdentifyCredentialResponse
	GetBody() *CreateIdentifyCredentialResponseBody
}

type CreateIdentifyCredentialResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIdentifyCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdentifyCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentifyCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateIdentifyCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIdentifyCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIdentifyCredentialResponse) GetBody() *CreateIdentifyCredentialResponseBody {
	return s.Body
}

func (s *CreateIdentifyCredentialResponse) SetHeaders(v map[string]*string) *CreateIdentifyCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateIdentifyCredentialResponse) SetStatusCode(v int32) *CreateIdentifyCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdentifyCredentialResponse) SetBody(v *CreateIdentifyCredentialResponseBody) *CreateIdentifyCredentialResponse {
	s.Body = v
	return s
}

func (s *CreateIdentifyCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
