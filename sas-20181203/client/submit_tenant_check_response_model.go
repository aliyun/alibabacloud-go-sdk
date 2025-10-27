// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTenantCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTenantCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTenantCheckResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTenantCheckResponseBody) *SubmitTenantCheckResponse
	GetBody() *SubmitTenantCheckResponseBody
}

type SubmitTenantCheckResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTenantCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTenantCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTenantCheckResponse) GoString() string {
	return s.String()
}

func (s *SubmitTenantCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTenantCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTenantCheckResponse) GetBody() *SubmitTenantCheckResponseBody {
	return s.Body
}

func (s *SubmitTenantCheckResponse) SetHeaders(v map[string]*string) *SubmitTenantCheckResponse {
	s.Headers = v
	return s
}

func (s *SubmitTenantCheckResponse) SetStatusCode(v int32) *SubmitTenantCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTenantCheckResponse) SetBody(v *SubmitTenantCheckResponseBody) *SubmitTenantCheckResponse {
	s.Body = v
	return s
}

func (s *SubmitTenantCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
