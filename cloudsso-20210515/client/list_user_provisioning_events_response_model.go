// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProvisioningEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserProvisioningEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserProvisioningEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserProvisioningEventsResponseBody) *ListUserProvisioningEventsResponse
	GetBody() *ListUserProvisioningEventsResponseBody
}

type ListUserProvisioningEventsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserProvisioningEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserProvisioningEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserProvisioningEventsResponse) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserProvisioningEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserProvisioningEventsResponse) GetBody() *ListUserProvisioningEventsResponseBody {
	return s.Body
}

func (s *ListUserProvisioningEventsResponse) SetHeaders(v map[string]*string) *ListUserProvisioningEventsResponse {
	s.Headers = v
	return s
}

func (s *ListUserProvisioningEventsResponse) SetStatusCode(v int32) *ListUserProvisioningEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserProvisioningEventsResponse) SetBody(v *ListUserProvisioningEventsResponseBody) *ListUserProvisioningEventsResponse {
	s.Body = v
	return s
}

func (s *ListUserProvisioningEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
