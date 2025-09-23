// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterOrcaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDBClusterOrcaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDBClusterOrcaResponse
	GetStatusCode() *int32
	SetBody(v *DisableDBClusterOrcaResponseBody) *DisableDBClusterOrcaResponse
	GetBody() *DisableDBClusterOrcaResponseBody
}

type DisableDBClusterOrcaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDBClusterOrcaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDBClusterOrcaResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterOrcaResponse) GoString() string {
	return s.String()
}

func (s *DisableDBClusterOrcaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDBClusterOrcaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDBClusterOrcaResponse) GetBody() *DisableDBClusterOrcaResponseBody {
	return s.Body
}

func (s *DisableDBClusterOrcaResponse) SetHeaders(v map[string]*string) *DisableDBClusterOrcaResponse {
	s.Headers = v
	return s
}

func (s *DisableDBClusterOrcaResponse) SetStatusCode(v int32) *DisableDBClusterOrcaResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDBClusterOrcaResponse) SetBody(v *DisableDBClusterOrcaResponseBody) *DisableDBClusterOrcaResponse {
	s.Body = v
	return s
}

func (s *DisableDBClusterOrcaResponse) Validate() error {
	return dara.Validate(s)
}
