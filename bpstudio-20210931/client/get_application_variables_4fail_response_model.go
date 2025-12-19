// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationVariables4FailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationVariables4FailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationVariables4FailResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationVariables4FailResponseBody) *GetApplicationVariables4FailResponse
	GetBody() *GetApplicationVariables4FailResponseBody
}

type GetApplicationVariables4FailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationVariables4FailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationVariables4FailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariables4FailResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationVariables4FailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationVariables4FailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationVariables4FailResponse) GetBody() *GetApplicationVariables4FailResponseBody {
	return s.Body
}

func (s *GetApplicationVariables4FailResponse) SetHeaders(v map[string]*string) *GetApplicationVariables4FailResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationVariables4FailResponse) SetStatusCode(v int32) *GetApplicationVariables4FailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationVariables4FailResponse) SetBody(v *GetApplicationVariables4FailResponseBody) *GetApplicationVariables4FailResponse {
	s.Body = v
	return s
}

func (s *GetApplicationVariables4FailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
