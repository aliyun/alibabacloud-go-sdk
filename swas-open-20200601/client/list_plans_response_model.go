// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListPlansResponseBody) *ListPlansResponse
	GetBody() *ListPlansResponseBody
}

type ListPlansResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPlansResponse) GoString() string {
	return s.String()
}

func (s *ListPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPlansResponse) GetBody() *ListPlansResponseBody {
	return s.Body
}

func (s *ListPlansResponse) SetHeaders(v map[string]*string) *ListPlansResponse {
	s.Headers = v
	return s
}

func (s *ListPlansResponse) SetStatusCode(v int32) *ListPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlansResponse) SetBody(v *ListPlansResponseBody) *ListPlansResponse {
	s.Body = v
	return s
}

func (s *ListPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
