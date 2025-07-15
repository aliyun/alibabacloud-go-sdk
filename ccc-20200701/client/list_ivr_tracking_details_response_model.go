// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIvrTrackingDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIvrTrackingDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIvrTrackingDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListIvrTrackingDetailsResponseBody) *ListIvrTrackingDetailsResponse
	GetBody() *ListIvrTrackingDetailsResponseBody
}

type ListIvrTrackingDetailsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIvrTrackingDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIvrTrackingDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIvrTrackingDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIvrTrackingDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIvrTrackingDetailsResponse) GetBody() *ListIvrTrackingDetailsResponseBody {
	return s.Body
}

func (s *ListIvrTrackingDetailsResponse) SetHeaders(v map[string]*string) *ListIvrTrackingDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListIvrTrackingDetailsResponse) SetStatusCode(v int32) *ListIvrTrackingDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIvrTrackingDetailsResponse) SetBody(v *ListIvrTrackingDetailsResponseBody) *ListIvrTrackingDetailsResponse {
	s.Body = v
	return s
}

func (s *ListIvrTrackingDetailsResponse) Validate() error {
	return dara.Validate(s)
}
