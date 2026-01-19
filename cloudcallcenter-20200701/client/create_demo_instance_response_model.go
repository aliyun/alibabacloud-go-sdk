// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemoInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDemoInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDemoInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDemoInstanceResponseBody) *CreateDemoInstanceResponse
	GetBody() *CreateDemoInstanceResponseBody
}

type CreateDemoInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDemoInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDemoInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDemoInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDemoInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDemoInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDemoInstanceResponse) GetBody() *CreateDemoInstanceResponseBody {
	return s.Body
}

func (s *CreateDemoInstanceResponse) SetHeaders(v map[string]*string) *CreateDemoInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDemoInstanceResponse) SetStatusCode(v int32) *CreateDemoInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDemoInstanceResponse) SetBody(v *CreateDemoInstanceResponseBody) *CreateDemoInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateDemoInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
