// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListFilesResponseBody) *ListFilesResponse
	GetBody() *ListFilesResponseBody
}

type ListFilesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFilesResponse) GoString() string {
	return s.String()
}

func (s *ListFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFilesResponse) GetBody() *ListFilesResponseBody {
	return s.Body
}

func (s *ListFilesResponse) SetHeaders(v map[string]*string) *ListFilesResponse {
	s.Headers = v
	return s
}

func (s *ListFilesResponse) SetStatusCode(v int32) *ListFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFilesResponse) SetBody(v *ListFilesResponseBody) *ListFilesResponse {
	s.Body = v
	return s
}

func (s *ListFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
