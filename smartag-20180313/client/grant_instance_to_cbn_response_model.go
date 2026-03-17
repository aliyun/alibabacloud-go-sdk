// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToCbnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantInstanceToCbnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantInstanceToCbnResponse
	GetStatusCode() *int32
	SetBody(v *GrantInstanceToCbnResponseBody) *GrantInstanceToCbnResponse
	GetBody() *GrantInstanceToCbnResponseBody
}

type GrantInstanceToCbnResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantInstanceToCbnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantInstanceToCbnResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToCbnResponse) GoString() string {
	return s.String()
}

func (s *GrantInstanceToCbnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantInstanceToCbnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantInstanceToCbnResponse) GetBody() *GrantInstanceToCbnResponseBody {
	return s.Body
}

func (s *GrantInstanceToCbnResponse) SetHeaders(v map[string]*string) *GrantInstanceToCbnResponse {
	s.Headers = v
	return s
}

func (s *GrantInstanceToCbnResponse) SetStatusCode(v int32) *GrantInstanceToCbnResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantInstanceToCbnResponse) SetBody(v *GrantInstanceToCbnResponseBody) *GrantInstanceToCbnResponse {
	s.Body = v
	return s
}

func (s *GrantInstanceToCbnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
