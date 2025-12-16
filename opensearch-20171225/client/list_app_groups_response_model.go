// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppGroupsResponseBody) *ListAppGroupsResponse
	GetBody() *ListAppGroupsResponseBody
}

type ListAppGroupsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppGroupsResponse) GetBody() *ListAppGroupsResponseBody {
	return s.Body
}

func (s *ListAppGroupsResponse) SetHeaders(v map[string]*string) *ListAppGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAppGroupsResponse) SetStatusCode(v int32) *ListAppGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppGroupsResponse) SetBody(v *ListAppGroupsResponseBody) *ListAppGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAppGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
