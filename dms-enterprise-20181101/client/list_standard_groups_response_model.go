// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStandardGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStandardGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStandardGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListStandardGroupsResponseBody) *ListStandardGroupsResponse
	GetBody() *ListStandardGroupsResponseBody
}

type ListStandardGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStandardGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStandardGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStandardGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListStandardGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStandardGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStandardGroupsResponse) GetBody() *ListStandardGroupsResponseBody {
	return s.Body
}

func (s *ListStandardGroupsResponse) SetHeaders(v map[string]*string) *ListStandardGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListStandardGroupsResponse) SetStatusCode(v int32) *ListStandardGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStandardGroupsResponse) SetBody(v *ListStandardGroupsResponseBody) *ListStandardGroupsResponse {
	s.Body = v
	return s
}

func (s *ListStandardGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
