// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSiteInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSiteInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSiteInstanceResponseBody) *UpdateSiteInstanceResponse
	GetBody() *UpdateSiteInstanceResponseBody
}

type UpdateSiteInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSiteInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSiteInstanceResponse) GetBody() *UpdateSiteInstanceResponseBody {
	return s.Body
}

func (s *UpdateSiteInstanceResponse) SetHeaders(v map[string]*string) *UpdateSiteInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteInstanceResponse) SetStatusCode(v int32) *UpdateSiteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteInstanceResponse) SetBody(v *UpdateSiteInstanceResponseBody) *UpdateSiteInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateSiteInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
