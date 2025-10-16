// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateOfficeSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateOfficeSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateOfficeSiteResponse
	GetStatusCode() *int32
	SetBody(v *ActivateOfficeSiteResponseBody) *ActivateOfficeSiteResponse
	GetBody() *ActivateOfficeSiteResponseBody
}

type ActivateOfficeSiteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateOfficeSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *ActivateOfficeSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateOfficeSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateOfficeSiteResponse) GetBody() *ActivateOfficeSiteResponseBody {
	return s.Body
}

func (s *ActivateOfficeSiteResponse) SetHeaders(v map[string]*string) *ActivateOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *ActivateOfficeSiteResponse) SetStatusCode(v int32) *ActivateOfficeSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateOfficeSiteResponse) SetBody(v *ActivateOfficeSiteResponseBody) *ActivateOfficeSiteResponse {
	s.Body = v
	return s
}

func (s *ActivateOfficeSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
