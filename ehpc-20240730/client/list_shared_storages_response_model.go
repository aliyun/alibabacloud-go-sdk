// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedStoragesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSharedStoragesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSharedStoragesResponse
	GetStatusCode() *int32
	SetBody(v *ListSharedStoragesResponseBody) *ListSharedStoragesResponse
	GetBody() *ListSharedStoragesResponseBody
}

type ListSharedStoragesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSharedStoragesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSharedStoragesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSharedStoragesResponse) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSharedStoragesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSharedStoragesResponse) GetBody() *ListSharedStoragesResponseBody {
	return s.Body
}

func (s *ListSharedStoragesResponse) SetHeaders(v map[string]*string) *ListSharedStoragesResponse {
	s.Headers = v
	return s
}

func (s *ListSharedStoragesResponse) SetStatusCode(v int32) *ListSharedStoragesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSharedStoragesResponse) SetBody(v *ListSharedStoragesResponseBody) *ListSharedStoragesResponse {
	s.Body = v
	return s
}

func (s *ListSharedStoragesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
