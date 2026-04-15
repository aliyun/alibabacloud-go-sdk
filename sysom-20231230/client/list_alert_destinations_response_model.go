// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertDestinationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertDestinationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertDestinationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertDestinationsResponseBody) *ListAlertDestinationsResponse
	GetBody() *ListAlertDestinationsResponseBody
}

type ListAlertDestinationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertDestinationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertDestinationsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertDestinationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertDestinationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertDestinationsResponse) GetBody() *ListAlertDestinationsResponseBody {
	return s.Body
}

func (s *ListAlertDestinationsResponse) SetHeaders(v map[string]*string) *ListAlertDestinationsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertDestinationsResponse) SetStatusCode(v int32) *ListAlertDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertDestinationsResponse) SetBody(v *ListAlertDestinationsResponseBody) *ListAlertDestinationsResponse {
	s.Body = v
	return s
}

func (s *ListAlertDestinationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
