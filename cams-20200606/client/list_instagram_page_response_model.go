// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstagramPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstagramPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstagramPageResponse
	GetStatusCode() *int32
	SetBody(v *ListInstagramPageResponseBody) *ListInstagramPageResponse
	GetBody() *ListInstagramPageResponseBody
}

type ListInstagramPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstagramPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstagramPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPageResponse) GoString() string {
	return s.String()
}

func (s *ListInstagramPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstagramPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstagramPageResponse) GetBody() *ListInstagramPageResponseBody {
	return s.Body
}

func (s *ListInstagramPageResponse) SetHeaders(v map[string]*string) *ListInstagramPageResponse {
	s.Headers = v
	return s
}

func (s *ListInstagramPageResponse) SetStatusCode(v int32) *ListInstagramPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstagramPageResponse) SetBody(v *ListInstagramPageResponseBody) *ListInstagramPageResponse {
	s.Body = v
	return s
}

func (s *ListInstagramPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
