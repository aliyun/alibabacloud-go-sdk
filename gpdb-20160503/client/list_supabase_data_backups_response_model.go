// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseDataBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupabaseDataBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupabaseDataBackupsResponse
	GetStatusCode() *int32
	SetBody(v *ListSupabaseDataBackupsResponseBody) *ListSupabaseDataBackupsResponse
	GetBody() *ListSupabaseDataBackupsResponseBody
}

type ListSupabaseDataBackupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupabaseDataBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupabaseDataBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseDataBackupsResponse) GoString() string {
	return s.String()
}

func (s *ListSupabaseDataBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupabaseDataBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupabaseDataBackupsResponse) GetBody() *ListSupabaseDataBackupsResponseBody {
	return s.Body
}

func (s *ListSupabaseDataBackupsResponse) SetHeaders(v map[string]*string) *ListSupabaseDataBackupsResponse {
	s.Headers = v
	return s
}

func (s *ListSupabaseDataBackupsResponse) SetStatusCode(v int32) *ListSupabaseDataBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupabaseDataBackupsResponse) SetBody(v *ListSupabaseDataBackupsResponseBody) *ListSupabaseDataBackupsResponse {
	s.Body = v
	return s
}

func (s *ListSupabaseDataBackupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
