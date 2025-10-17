// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActivationCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateActivationCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateActivationCodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateActivationCodeResponseBody) *CreateActivationCodeResponse
	GetBody() *CreateActivationCodeResponseBody
}

type CreateActivationCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateActivationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateActivationCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationCodeResponse) GoString() string {
	return s.String()
}

func (s *CreateActivationCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateActivationCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateActivationCodeResponse) GetBody() *CreateActivationCodeResponseBody {
	return s.Body
}

func (s *CreateActivationCodeResponse) SetHeaders(v map[string]*string) *CreateActivationCodeResponse {
	s.Headers = v
	return s
}

func (s *CreateActivationCodeResponse) SetStatusCode(v int32) *CreateActivationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateActivationCodeResponse) SetBody(v *CreateActivationCodeResponseBody) *CreateActivationCodeResponse {
	s.Body = v
	return s
}

func (s *CreateActivationCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
