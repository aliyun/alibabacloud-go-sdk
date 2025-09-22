// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgConstituteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrgConstituteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrgConstituteResponse
	GetStatusCode() *int32
	SetBody(v *GetOrgConstituteResponseBody) *GetOrgConstituteResponse
	GetBody() *GetOrgConstituteResponseBody
}

type GetOrgConstituteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgConstituteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrgConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetOrgConstituteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrgConstituteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrgConstituteResponse) GetBody() *GetOrgConstituteResponseBody {
	return s.Body
}

func (s *GetOrgConstituteResponse) SetHeaders(v map[string]*string) *GetOrgConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetOrgConstituteResponse) SetStatusCode(v int32) *GetOrgConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgConstituteResponse) SetBody(v *GetOrgConstituteResponseBody) *GetOrgConstituteResponse {
	s.Body = v
	return s
}

func (s *GetOrgConstituteResponse) Validate() error {
	return dara.Validate(s)
}
