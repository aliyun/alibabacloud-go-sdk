// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountFactoryBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccountFactoryBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccountFactoryBaselineResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccountFactoryBaselineResponseBody) *CreateAccountFactoryBaselineResponse
	GetBody() *CreateAccountFactoryBaselineResponseBody
}

type CreateAccountFactoryBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountFactoryBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccountFactoryBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountFactoryBaselineResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountFactoryBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccountFactoryBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccountFactoryBaselineResponse) GetBody() *CreateAccountFactoryBaselineResponseBody {
	return s.Body
}

func (s *CreateAccountFactoryBaselineResponse) SetHeaders(v map[string]*string) *CreateAccountFactoryBaselineResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountFactoryBaselineResponse) SetStatusCode(v int32) *CreateAccountFactoryBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountFactoryBaselineResponse) SetBody(v *CreateAccountFactoryBaselineResponseBody) *CreateAccountFactoryBaselineResponse {
	s.Body = v
	return s
}

func (s *CreateAccountFactoryBaselineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
