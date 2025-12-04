// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesByNcdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancesByNcdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancesByNcdResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancesByNcdResponseBody) *ListInstancesByNcdResponse
	GetBody() *ListInstancesByNcdResponseBody
}

type ListInstancesByNcdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesByNcdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesByNcdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesByNcdResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancesByNcdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancesByNcdResponse) GetBody() *ListInstancesByNcdResponseBody {
	return s.Body
}

func (s *ListInstancesByNcdResponse) SetHeaders(v map[string]*string) *ListInstancesByNcdResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesByNcdResponse) SetStatusCode(v int32) *ListInstancesByNcdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesByNcdResponse) SetBody(v *ListInstancesByNcdResponseBody) *ListInstancesByNcdResponse {
	s.Body = v
	return s
}

func (s *ListInstancesByNcdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
