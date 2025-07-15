// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantInstanceToCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantInstanceToCenResponse
	GetStatusCode() *int32
	SetBody(v *GrantInstanceToCenResponseBody) *GrantInstanceToCenResponse
	GetBody() *GrantInstanceToCenResponseBody
}

type GrantInstanceToCenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantInstanceToCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantInstanceToCenResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToCenResponse) GoString() string {
	return s.String()
}

func (s *GrantInstanceToCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantInstanceToCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantInstanceToCenResponse) GetBody() *GrantInstanceToCenResponseBody {
	return s.Body
}

func (s *GrantInstanceToCenResponse) SetHeaders(v map[string]*string) *GrantInstanceToCenResponse {
	s.Headers = v
	return s
}

func (s *GrantInstanceToCenResponse) SetStatusCode(v int32) *GrantInstanceToCenResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantInstanceToCenResponse) SetBody(v *GrantInstanceToCenResponseBody) *GrantInstanceToCenResponse {
	s.Body = v
	return s
}

func (s *GrantInstanceToCenResponse) Validate() error {
	return dara.Validate(s)
}
