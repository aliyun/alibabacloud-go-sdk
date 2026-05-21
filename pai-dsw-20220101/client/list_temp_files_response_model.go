// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTempFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTempFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTempFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListTempFilesResponseBody) *ListTempFilesResponse
	GetBody() *ListTempFilesResponseBody
}

type ListTempFilesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTempFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTempFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTempFilesResponse) GoString() string {
	return s.String()
}

func (s *ListTempFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTempFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTempFilesResponse) GetBody() *ListTempFilesResponseBody {
	return s.Body
}

func (s *ListTempFilesResponse) SetHeaders(v map[string]*string) *ListTempFilesResponse {
	s.Headers = v
	return s
}

func (s *ListTempFilesResponse) SetStatusCode(v int32) *ListTempFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTempFilesResponse) SetBody(v *ListTempFilesResponseBody) *ListTempFilesResponse {
	s.Body = v
	return s
}

func (s *ListTempFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
