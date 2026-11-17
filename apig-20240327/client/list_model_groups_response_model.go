// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelGroupsResponseBody) *ListModelGroupsResponse
	GetBody() *ListModelGroupsResponseBody
}

type ListModelGroupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListModelGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelGroupsResponse) GetBody() *ListModelGroupsResponseBody {
	return s.Body
}

func (s *ListModelGroupsResponse) SetHeaders(v map[string]*string) *ListModelGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListModelGroupsResponse) SetStatusCode(v int32) *ListModelGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelGroupsResponse) SetBody(v *ListModelGroupsResponseBody) *ListModelGroupsResponse {
	s.Body = v
	return s
}

func (s *ListModelGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
