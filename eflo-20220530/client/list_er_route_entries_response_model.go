// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListErRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListErRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListErRouteEntriesResponseBody) *ListErRouteEntriesResponse
	GetBody() *ListErRouteEntriesResponseBody
}

type ListErRouteEntriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListErRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListErRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListErRouteEntriesResponse) GetBody() *ListErRouteEntriesResponseBody {
	return s.Body
}

func (s *ListErRouteEntriesResponse) SetHeaders(v map[string]*string) *ListErRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListErRouteEntriesResponse) SetStatusCode(v int32) *ListErRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErRouteEntriesResponse) SetBody(v *ListErRouteEntriesResponseBody) *ListErRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *ListErRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
