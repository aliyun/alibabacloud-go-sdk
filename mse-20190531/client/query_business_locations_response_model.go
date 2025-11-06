// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBusinessLocationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBusinessLocationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBusinessLocationsResponse
	GetStatusCode() *int32
	SetBody(v *QueryBusinessLocationsResponseBody) *QueryBusinessLocationsResponse
	GetBody() *QueryBusinessLocationsResponseBody
}

type QueryBusinessLocationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBusinessLocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBusinessLocationsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBusinessLocationsResponse) GoString() string {
	return s.String()
}

func (s *QueryBusinessLocationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBusinessLocationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBusinessLocationsResponse) GetBody() *QueryBusinessLocationsResponseBody {
	return s.Body
}

func (s *QueryBusinessLocationsResponse) SetHeaders(v map[string]*string) *QueryBusinessLocationsResponse {
	s.Headers = v
	return s
}

func (s *QueryBusinessLocationsResponse) SetStatusCode(v int32) *QueryBusinessLocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBusinessLocationsResponse) SetBody(v *QueryBusinessLocationsResponseBody) *QueryBusinessLocationsResponse {
	s.Body = v
	return s
}

func (s *QueryBusinessLocationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
