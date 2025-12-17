// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileUploadResponse
	GetStatusCode() *int32
	SetBody(v *ListFileUploadResponseBody) *ListFileUploadResponse
	GetBody() *ListFileUploadResponseBody
}

type ListFileUploadResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileUploadResponse) GoString() string {
	return s.String()
}

func (s *ListFileUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileUploadResponse) GetBody() *ListFileUploadResponseBody {
	return s.Body
}

func (s *ListFileUploadResponse) SetHeaders(v map[string]*string) *ListFileUploadResponse {
	s.Headers = v
	return s
}

func (s *ListFileUploadResponse) SetStatusCode(v int32) *ListFileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileUploadResponse) SetBody(v *ListFileUploadResponseBody) *ListFileUploadResponse {
	s.Body = v
	return s
}

func (s *ListFileUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
