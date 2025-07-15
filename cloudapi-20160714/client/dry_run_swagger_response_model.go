// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunSwaggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DryRunSwaggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DryRunSwaggerResponse
	GetStatusCode() *int32
	SetBody(v *DryRunSwaggerResponseBody) *DryRunSwaggerResponse
	GetBody() *DryRunSwaggerResponseBody
}

type DryRunSwaggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DryRunSwaggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DryRunSwaggerResponse) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponse) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DryRunSwaggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DryRunSwaggerResponse) GetBody() *DryRunSwaggerResponseBody {
	return s.Body
}

func (s *DryRunSwaggerResponse) SetHeaders(v map[string]*string) *DryRunSwaggerResponse {
	s.Headers = v
	return s
}

func (s *DryRunSwaggerResponse) SetStatusCode(v int32) *DryRunSwaggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DryRunSwaggerResponse) SetBody(v *DryRunSwaggerResponseBody) *DryRunSwaggerResponse {
	s.Body = v
	return s
}

func (s *DryRunSwaggerResponse) Validate() error {
	return dara.Validate(s)
}
