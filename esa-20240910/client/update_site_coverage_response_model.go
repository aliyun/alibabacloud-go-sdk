// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteCoverageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSiteCoverageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSiteCoverageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSiteCoverageResponseBody) *UpdateSiteCoverageResponse
	GetBody() *UpdateSiteCoverageResponseBody
}

type UpdateSiteCoverageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteCoverageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteCoverageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteCoverageResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteCoverageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSiteCoverageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSiteCoverageResponse) GetBody() *UpdateSiteCoverageResponseBody {
	return s.Body
}

func (s *UpdateSiteCoverageResponse) SetHeaders(v map[string]*string) *UpdateSiteCoverageResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteCoverageResponse) SetStatusCode(v int32) *UpdateSiteCoverageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteCoverageResponse) SetBody(v *UpdateSiteCoverageResponseBody) *UpdateSiteCoverageResponse {
	s.Body = v
	return s
}

func (s *UpdateSiteCoverageResponse) Validate() error {
	return dara.Validate(s)
}
