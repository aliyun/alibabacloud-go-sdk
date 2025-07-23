// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserResourcesResponseBody) *ListUserResourcesResponse
	GetBody() *ListUserResourcesResponseBody
}

type ListUserResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserResourcesResponse) GetBody() *ListUserResourcesResponseBody {
	return s.Body
}

func (s *ListUserResourcesResponse) SetHeaders(v map[string]*string) *ListUserResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListUserResourcesResponse) SetStatusCode(v int32) *ListUserResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserResourcesResponse) SetBody(v *ListUserResourcesResponseBody) *ListUserResourcesResponse {
	s.Body = v
	return s
}

func (s *ListUserResourcesResponse) Validate() error {
	return dara.Validate(s)
}
