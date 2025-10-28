// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBuildPackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBuildPackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBuildPackResponse
	GetStatusCode() *int32
	SetBody(v *ListBuildPackResponseBody) *ListBuildPackResponse
	GetBody() *ListBuildPackResponseBody
}

type ListBuildPackResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBuildPackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBuildPackResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBuildPackResponse) GoString() string {
	return s.String()
}

func (s *ListBuildPackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBuildPackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBuildPackResponse) GetBody() *ListBuildPackResponseBody {
	return s.Body
}

func (s *ListBuildPackResponse) SetHeaders(v map[string]*string) *ListBuildPackResponse {
	s.Headers = v
	return s
}

func (s *ListBuildPackResponse) SetStatusCode(v int32) *ListBuildPackResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBuildPackResponse) SetBody(v *ListBuildPackResponseBody) *ListBuildPackResponse {
	s.Body = v
	return s
}

func (s *ListBuildPackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
