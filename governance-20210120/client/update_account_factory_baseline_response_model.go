// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountFactoryBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccountFactoryBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccountFactoryBaselineResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccountFactoryBaselineResponseBody) *UpdateAccountFactoryBaselineResponse
	GetBody() *UpdateAccountFactoryBaselineResponseBody
}

type UpdateAccountFactoryBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountFactoryBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountFactoryBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountFactoryBaselineResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountFactoryBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccountFactoryBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccountFactoryBaselineResponse) GetBody() *UpdateAccountFactoryBaselineResponseBody {
	return s.Body
}

func (s *UpdateAccountFactoryBaselineResponse) SetHeaders(v map[string]*string) *UpdateAccountFactoryBaselineResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountFactoryBaselineResponse) SetStatusCode(v int32) *UpdateAccountFactoryBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountFactoryBaselineResponse) SetBody(v *UpdateAccountFactoryBaselineResponseBody) *UpdateAccountFactoryBaselineResponse {
	s.Body = v
	return s
}

func (s *UpdateAccountFactoryBaselineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
