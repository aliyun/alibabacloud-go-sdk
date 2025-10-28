// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceGroupsResponseBody) *ListServiceGroupsResponse
	GetBody() *ListServiceGroupsResponseBody
}

type ListServiceGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceGroupsResponse) GetBody() *ListServiceGroupsResponseBody {
	return s.Body
}

func (s *ListServiceGroupsResponse) SetHeaders(v map[string]*string) *ListServiceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceGroupsResponse) SetStatusCode(v int32) *ListServiceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceGroupsResponse) SetBody(v *ListServiceGroupsResponseBody) *ListServiceGroupsResponse {
	s.Body = v
	return s
}

func (s *ListServiceGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
