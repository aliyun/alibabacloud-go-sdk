// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBuildConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBuildConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBuildConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListBuildConfigsResponseBody) *ListBuildConfigsResponse
	GetBody() *ListBuildConfigsResponseBody
}

type ListBuildConfigsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBuildConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBuildConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBuildConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBuildConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBuildConfigsResponse) GetBody() *ListBuildConfigsResponseBody {
	return s.Body
}

func (s *ListBuildConfigsResponse) SetHeaders(v map[string]*string) *ListBuildConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListBuildConfigsResponse) SetStatusCode(v int32) *ListBuildConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBuildConfigsResponse) SetBody(v *ListBuildConfigsResponseBody) *ListBuildConfigsResponse {
	s.Body = v
	return s
}

func (s *ListBuildConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
