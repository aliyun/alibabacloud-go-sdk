// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLgfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLgfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLgfResponse
	GetStatusCode() *int32
	SetBody(v *ListLgfResponseBody) *ListLgfResponse
	GetBody() *ListLgfResponseBody
}

type ListLgfResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLgfResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLgfResponse) GoString() string {
	return s.String()
}

func (s *ListLgfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLgfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLgfResponse) GetBody() *ListLgfResponseBody {
	return s.Body
}

func (s *ListLgfResponse) SetHeaders(v map[string]*string) *ListLgfResponse {
	s.Headers = v
	return s
}

func (s *ListLgfResponse) SetStatusCode(v int32) *ListLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLgfResponse) SetBody(v *ListLgfResponseBody) *ListLgfResponse {
	s.Body = v
	return s
}

func (s *ListLgfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
