// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCmsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCmsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCmsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListCmsInstancesResponseBody) *ListCmsInstancesResponse
	GetBody() *ListCmsInstancesResponseBody
}

type ListCmsInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCmsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCmsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCmsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCmsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCmsInstancesResponse) GetBody() *ListCmsInstancesResponseBody {
	return s.Body
}

func (s *ListCmsInstancesResponse) SetHeaders(v map[string]*string) *ListCmsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListCmsInstancesResponse) SetStatusCode(v int32) *ListCmsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCmsInstancesResponse) SetBody(v *ListCmsInstancesResponseBody) *ListCmsInstancesResponse {
	s.Body = v
	return s
}

func (s *ListCmsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
