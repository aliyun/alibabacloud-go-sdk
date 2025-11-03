// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepoTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepoTagResponse
	GetStatusCode() *int32
	SetBody(v *ListRepoTagResponseBody) *ListRepoTagResponse
	GetBody() *ListRepoTagResponseBody
}

type ListRepoTagResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTagResponse) GoString() string {
	return s.String()
}

func (s *ListRepoTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepoTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepoTagResponse) GetBody() *ListRepoTagResponseBody {
	return s.Body
}

func (s *ListRepoTagResponse) SetHeaders(v map[string]*string) *ListRepoTagResponse {
	s.Headers = v
	return s
}

func (s *ListRepoTagResponse) SetStatusCode(v int32) *ListRepoTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoTagResponse) SetBody(v *ListRepoTagResponseBody) *ListRepoTagResponse {
	s.Body = v
	return s
}

func (s *ListRepoTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
