// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaEntityDefsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaEntityDefsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaEntityDefsResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaEntityDefsResponseBody) *ListMetaEntityDefsResponse
	GetBody() *ListMetaEntityDefsResponseBody
}

type ListMetaEntityDefsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaEntityDefsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaEntityDefsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntityDefsResponse) GoString() string {
	return s.String()
}

func (s *ListMetaEntityDefsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaEntityDefsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaEntityDefsResponse) GetBody() *ListMetaEntityDefsResponseBody {
	return s.Body
}

func (s *ListMetaEntityDefsResponse) SetHeaders(v map[string]*string) *ListMetaEntityDefsResponse {
	s.Headers = v
	return s
}

func (s *ListMetaEntityDefsResponse) SetStatusCode(v int32) *ListMetaEntityDefsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaEntityDefsResponse) SetBody(v *ListMetaEntityDefsResponseBody) *ListMetaEntityDefsResponse {
	s.Body = v
	return s
}

func (s *ListMetaEntityDefsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
