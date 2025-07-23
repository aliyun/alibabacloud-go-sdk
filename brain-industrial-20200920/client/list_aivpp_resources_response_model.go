// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAivppResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAivppResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAivppResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListAivppResourcesResponseBody) *ListAivppResourcesResponse
	GetBody() *ListAivppResourcesResponseBody
}

type ListAivppResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAivppResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAivppResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAivppResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAivppResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAivppResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAivppResourcesResponse) GetBody() *ListAivppResourcesResponseBody {
	return s.Body
}

func (s *ListAivppResourcesResponse) SetHeaders(v map[string]*string) *ListAivppResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAivppResourcesResponse) SetStatusCode(v int32) *ListAivppResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAivppResourcesResponse) SetBody(v *ListAivppResourcesResponseBody) *ListAivppResourcesResponse {
	s.Body = v
	return s
}

func (s *ListAivppResourcesResponse) Validate() error {
	return dara.Validate(s)
}
