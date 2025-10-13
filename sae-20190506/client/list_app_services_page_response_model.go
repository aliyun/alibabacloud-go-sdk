// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppServicesPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppServicesPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppServicesPageResponse
	GetStatusCode() *int32
	SetBody(v *ListAppServicesPageResponseBody) *ListAppServicesPageResponse
	GetBody() *ListAppServicesPageResponseBody
}

type ListAppServicesPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppServicesPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppServicesPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesPageResponse) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppServicesPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppServicesPageResponse) GetBody() *ListAppServicesPageResponseBody {
	return s.Body
}

func (s *ListAppServicesPageResponse) SetHeaders(v map[string]*string) *ListAppServicesPageResponse {
	s.Headers = v
	return s
}

func (s *ListAppServicesPageResponse) SetStatusCode(v int32) *ListAppServicesPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppServicesPageResponse) SetBody(v *ListAppServicesPageResponseBody) *ListAppServicesPageResponse {
	s.Body = v
	return s
}

func (s *ListAppServicesPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
