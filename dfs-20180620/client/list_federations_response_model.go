// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFederationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFederationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFederationsResponse
	GetStatusCode() *int32
	SetBody(v *ListFederationsResponseBody) *ListFederationsResponse
	GetBody() *ListFederationsResponseBody
}

type ListFederationsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFederationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFederationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFederationsResponse) GoString() string {
	return s.String()
}

func (s *ListFederationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFederationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFederationsResponse) GetBody() *ListFederationsResponseBody {
	return s.Body
}

func (s *ListFederationsResponse) SetHeaders(v map[string]*string) *ListFederationsResponse {
	s.Headers = v
	return s
}

func (s *ListFederationsResponse) SetStatusCode(v int32) *ListFederationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFederationsResponse) SetBody(v *ListFederationsResponseBody) *ListFederationsResponse {
	s.Body = v
	return s
}

func (s *ListFederationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
