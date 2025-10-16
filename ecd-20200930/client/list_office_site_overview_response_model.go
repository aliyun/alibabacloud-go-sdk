// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfficeSiteOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOfficeSiteOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOfficeSiteOverviewResponse
	GetStatusCode() *int32
	SetBody(v *ListOfficeSiteOverviewResponseBody) *ListOfficeSiteOverviewResponse
	GetBody() *ListOfficeSiteOverviewResponseBody
}

type ListOfficeSiteOverviewResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOfficeSiteOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOfficeSiteOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOfficeSiteOverviewResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOfficeSiteOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOfficeSiteOverviewResponse) GetBody() *ListOfficeSiteOverviewResponseBody {
	return s.Body
}

func (s *ListOfficeSiteOverviewResponse) SetHeaders(v map[string]*string) *ListOfficeSiteOverviewResponse {
	s.Headers = v
	return s
}

func (s *ListOfficeSiteOverviewResponse) SetStatusCode(v int32) *ListOfficeSiteOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOfficeSiteOverviewResponse) SetBody(v *ListOfficeSiteOverviewResponseBody) *ListOfficeSiteOverviewResponse {
	s.Body = v
	return s
}

func (s *ListOfficeSiteOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
