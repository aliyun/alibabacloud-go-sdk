// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteCustomLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSiteCustomLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSiteCustomLogResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSiteCustomLogResponseBody) *UpdateSiteCustomLogResponse
	GetBody() *UpdateSiteCustomLogResponseBody
}

type UpdateSiteCustomLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteCustomLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteCustomLogResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteCustomLogResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteCustomLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSiteCustomLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSiteCustomLogResponse) GetBody() *UpdateSiteCustomLogResponseBody {
	return s.Body
}

func (s *UpdateSiteCustomLogResponse) SetHeaders(v map[string]*string) *UpdateSiteCustomLogResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteCustomLogResponse) SetStatusCode(v int32) *UpdateSiteCustomLogResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteCustomLogResponse) SetBody(v *UpdateSiteCustomLogResponseBody) *UpdateSiteCustomLogResponse {
	s.Body = v
	return s
}

func (s *UpdateSiteCustomLogResponse) Validate() error {
	return dara.Validate(s)
}
