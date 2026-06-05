// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpecificationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSpecificationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSpecificationsResponse
	GetStatusCode() *int32
	SetBody(v *ListSpecificationsResponseBody) *ListSpecificationsResponse
	GetBody() *ListSpecificationsResponseBody
}

type ListSpecificationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSpecificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSpecificationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSpecificationsResponse) GoString() string {
	return s.String()
}

func (s *ListSpecificationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSpecificationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSpecificationsResponse) GetBody() *ListSpecificationsResponseBody {
	return s.Body
}

func (s *ListSpecificationsResponse) SetHeaders(v map[string]*string) *ListSpecificationsResponse {
	s.Headers = v
	return s
}

func (s *ListSpecificationsResponse) SetStatusCode(v int32) *ListSpecificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSpecificationsResponse) SetBody(v *ListSpecificationsResponseBody) *ListSpecificationsResponse {
	s.Body = v
	return s
}

func (s *ListSpecificationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
