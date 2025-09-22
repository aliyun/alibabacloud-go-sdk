// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmissionSourceConstituteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmissionSourceConstituteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmissionSourceConstituteResponse
	GetStatusCode() *int32
	SetBody(v *GetEmissionSourceConstituteResponseBody) *GetEmissionSourceConstituteResponse
	GetBody() *GetEmissionSourceConstituteResponseBody
}

type GetEmissionSourceConstituteResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmissionSourceConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmissionSourceConstituteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmissionSourceConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetEmissionSourceConstituteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmissionSourceConstituteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmissionSourceConstituteResponse) GetBody() *GetEmissionSourceConstituteResponseBody {
	return s.Body
}

func (s *GetEmissionSourceConstituteResponse) SetHeaders(v map[string]*string) *GetEmissionSourceConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetEmissionSourceConstituteResponse) SetStatusCode(v int32) *GetEmissionSourceConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmissionSourceConstituteResponse) SetBody(v *GetEmissionSourceConstituteResponseBody) *GetEmissionSourceConstituteResponse {
	s.Body = v
	return s
}

func (s *GetEmissionSourceConstituteResponse) Validate() error {
	return dara.Validate(s)
}
