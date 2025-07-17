// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeDependenciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeDependenciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeDependenciesResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeDependenciesResponseBody) *ListNodeDependenciesResponse
	GetBody() *ListNodeDependenciesResponseBody
}

type ListNodeDependenciesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeDependenciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeDependenciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponse) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeDependenciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeDependenciesResponse) GetBody() *ListNodeDependenciesResponseBody {
	return s.Body
}

func (s *ListNodeDependenciesResponse) SetHeaders(v map[string]*string) *ListNodeDependenciesResponse {
	s.Headers = v
	return s
}

func (s *ListNodeDependenciesResponse) SetStatusCode(v int32) *ListNodeDependenciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeDependenciesResponse) SetBody(v *ListNodeDependenciesResponseBody) *ListNodeDependenciesResponse {
	s.Body = v
	return s
}

func (s *ListNodeDependenciesResponse) Validate() error {
	return dara.Validate(s)
}
