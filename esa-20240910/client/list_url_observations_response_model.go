// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUrlObservationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUrlObservationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUrlObservationsResponse
	GetStatusCode() *int32
	SetBody(v *ListUrlObservationsResponseBody) *ListUrlObservationsResponse
	GetBody() *ListUrlObservationsResponseBody
}

type ListUrlObservationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUrlObservationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUrlObservationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUrlObservationsResponse) GoString() string {
	return s.String()
}

func (s *ListUrlObservationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUrlObservationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUrlObservationsResponse) GetBody() *ListUrlObservationsResponseBody {
	return s.Body
}

func (s *ListUrlObservationsResponse) SetHeaders(v map[string]*string) *ListUrlObservationsResponse {
	s.Headers = v
	return s
}

func (s *ListUrlObservationsResponse) SetStatusCode(v int32) *ListUrlObservationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUrlObservationsResponse) SetBody(v *ListUrlObservationsResponseBody) *ListUrlObservationsResponse {
	s.Body = v
	return s
}

func (s *ListUrlObservationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
