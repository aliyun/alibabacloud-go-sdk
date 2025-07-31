// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceConfResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceConfResponseBody) *GetServiceConfResponse
	GetBody() *GetServiceConfResponseBody
}

type GetServiceConfResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceConfResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfResponse) GoString() string {
	return s.String()
}

func (s *GetServiceConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceConfResponse) GetBody() *GetServiceConfResponseBody {
	return s.Body
}

func (s *GetServiceConfResponse) SetHeaders(v map[string]*string) *GetServiceConfResponse {
	s.Headers = v
	return s
}

func (s *GetServiceConfResponse) SetStatusCode(v int32) *GetServiceConfResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceConfResponse) SetBody(v *GetServiceConfResponseBody) *GetServiceConfResponse {
	s.Body = v
	return s
}

func (s *GetServiceConfResponse) Validate() error {
	return dara.Validate(s)
}
