// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHivesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHivesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHivesResponse
	GetStatusCode() *int32
	SetBody(v *ListHivesResponseBody) *ListHivesResponse
	GetBody() *ListHivesResponseBody
}

type ListHivesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHivesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHivesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHivesResponse) GoString() string {
	return s.String()
}

func (s *ListHivesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHivesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHivesResponse) GetBody() *ListHivesResponseBody {
	return s.Body
}

func (s *ListHivesResponse) SetHeaders(v map[string]*string) *ListHivesResponse {
	s.Headers = v
	return s
}

func (s *ListHivesResponse) SetStatusCode(v int32) *ListHivesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHivesResponse) SetBody(v *ListHivesResponseBody) *ListHivesResponse {
	s.Body = v
	return s
}

func (s *ListHivesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
