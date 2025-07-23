// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBackupsResponse
	GetStatusCode() *int32
	SetBody(v *ListBackupsResponseBody) *ListBackupsResponse
	GetBody() *ListBackupsResponseBody
}

type ListBackupsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBackupsResponse) GoString() string {
	return s.String()
}

func (s *ListBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBackupsResponse) GetBody() *ListBackupsResponseBody {
	return s.Body
}

func (s *ListBackupsResponse) SetHeaders(v map[string]*string) *ListBackupsResponse {
	s.Headers = v
	return s
}

func (s *ListBackupsResponse) SetStatusCode(v int32) *ListBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBackupsResponse) SetBody(v *ListBackupsResponseBody) *ListBackupsResponse {
	s.Body = v
	return s
}

func (s *ListBackupsResponse) Validate() error {
	return dara.Validate(s)
}
