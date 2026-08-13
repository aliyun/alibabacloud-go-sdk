// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutputFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOutputFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOutputFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListOutputFilesResponseBody) *ListOutputFilesResponse
	GetBody() *ListOutputFilesResponseBody
}

type ListOutputFilesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOutputFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOutputFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesResponse) GoString() string {
	return s.String()
}

func (s *ListOutputFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOutputFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOutputFilesResponse) GetBody() *ListOutputFilesResponseBody {
	return s.Body
}

func (s *ListOutputFilesResponse) SetHeaders(v map[string]*string) *ListOutputFilesResponse {
	s.Headers = v
	return s
}

func (s *ListOutputFilesResponse) SetStatusCode(v int32) *ListOutputFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOutputFilesResponse) SetBody(v *ListOutputFilesResponseBody) *ListOutputFilesResponse {
	s.Body = v
	return s
}

func (s *ListOutputFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
