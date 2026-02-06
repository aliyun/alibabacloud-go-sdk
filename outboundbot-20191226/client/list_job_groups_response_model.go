// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListJobGroupsResponseBody) *ListJobGroupsResponse
	GetBody() *ListJobGroupsResponseBody
}

type ListJobGroupsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListJobGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobGroupsResponse) GetBody() *ListJobGroupsResponseBody {
	return s.Body
}

func (s *ListJobGroupsResponse) SetHeaders(v map[string]*string) *ListJobGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListJobGroupsResponse) SetStatusCode(v int32) *ListJobGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobGroupsResponse) SetBody(v *ListJobGroupsResponseBody) *ListJobGroupsResponse {
	s.Body = v
	return s
}

func (s *ListJobGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
