// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeEcsSpecsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServerIdeEcsSpecsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServerIdeEcsSpecsResponse
	GetStatusCode() *int32
	SetBody(v *ListServerIdeEcsSpecsResponseBody) *ListServerIdeEcsSpecsResponse
	GetBody() *ListServerIdeEcsSpecsResponseBody
}

type ListServerIdeEcsSpecsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServerIdeEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServerIdeEcsSpecsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeEcsSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListServerIdeEcsSpecsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServerIdeEcsSpecsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServerIdeEcsSpecsResponse) GetBody() *ListServerIdeEcsSpecsResponseBody {
	return s.Body
}

func (s *ListServerIdeEcsSpecsResponse) SetHeaders(v map[string]*string) *ListServerIdeEcsSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListServerIdeEcsSpecsResponse) SetStatusCode(v int32) *ListServerIdeEcsSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponse) SetBody(v *ListServerIdeEcsSpecsResponseBody) *ListServerIdeEcsSpecsResponse {
	s.Body = v
	return s
}

func (s *ListServerIdeEcsSpecsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
