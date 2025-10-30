// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyTenantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMyTenantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMyTenantsResponse
	GetStatusCode() *int32
	SetBody(v *GetMyTenantsResponseBody) *GetMyTenantsResponse
	GetBody() *GetMyTenantsResponseBody
}

type GetMyTenantsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMyTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMyTenantsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMyTenantsResponse) GoString() string {
	return s.String()
}

func (s *GetMyTenantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMyTenantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMyTenantsResponse) GetBody() *GetMyTenantsResponseBody {
	return s.Body
}

func (s *GetMyTenantsResponse) SetHeaders(v map[string]*string) *GetMyTenantsResponse {
	s.Headers = v
	return s
}

func (s *GetMyTenantsResponse) SetStatusCode(v int32) *GetMyTenantsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMyTenantsResponse) SetBody(v *GetMyTenantsResponseBody) *GetMyTenantsResponse {
	s.Body = v
	return s
}

func (s *GetMyTenantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
