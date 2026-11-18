// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaFoldersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaFoldersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaFoldersResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaFoldersResponseBody) *ListMediaFoldersResponse
	GetBody() *ListMediaFoldersResponseBody
}

type ListMediaFoldersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaFoldersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaFoldersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaFoldersResponse) GoString() string {
	return s.String()
}

func (s *ListMediaFoldersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaFoldersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaFoldersResponse) GetBody() *ListMediaFoldersResponseBody {
	return s.Body
}

func (s *ListMediaFoldersResponse) SetHeaders(v map[string]*string) *ListMediaFoldersResponse {
	s.Headers = v
	return s
}

func (s *ListMediaFoldersResponse) SetStatusCode(v int32) *ListMediaFoldersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaFoldersResponse) SetBody(v *ListMediaFoldersResponseBody) *ListMediaFoldersResponse {
	s.Body = v
	return s
}

func (s *ListMediaFoldersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
