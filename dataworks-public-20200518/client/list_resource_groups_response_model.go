// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceGroupsResponseBody) *ListResourceGroupsResponse
	GetBody() *ListResourceGroupsResponseBody
}

type ListResourceGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceGroupsResponse) GetBody() *ListResourceGroupsResponseBody {
	return s.Body
}

func (s *ListResourceGroupsResponse) SetHeaders(v map[string]*string) *ListResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupsResponse) SetStatusCode(v int32) *ListResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupsResponse) SetBody(v *ListResourceGroupsResponseBody) *ListResourceGroupsResponse {
	s.Body = v
	return s
}

func (s *ListResourceGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
