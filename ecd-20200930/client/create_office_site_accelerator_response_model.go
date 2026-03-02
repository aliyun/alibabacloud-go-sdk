// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfficeSiteAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOfficeSiteAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOfficeSiteAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *CreateOfficeSiteAcceleratorResponseBody) *CreateOfficeSiteAcceleratorResponse
	GetBody() *CreateOfficeSiteAcceleratorResponseBody
}

type CreateOfficeSiteAcceleratorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOfficeSiteAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOfficeSiteAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOfficeSiteAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *CreateOfficeSiteAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOfficeSiteAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOfficeSiteAcceleratorResponse) GetBody() *CreateOfficeSiteAcceleratorResponseBody {
	return s.Body
}

func (s *CreateOfficeSiteAcceleratorResponse) SetHeaders(v map[string]*string) *CreateOfficeSiteAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *CreateOfficeSiteAcceleratorResponse) SetStatusCode(v int32) *CreateOfficeSiteAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorResponse) SetBody(v *CreateOfficeSiteAcceleratorResponseBody) *CreateOfficeSiteAcceleratorResponse {
	s.Body = v
	return s
}

func (s *CreateOfficeSiteAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
