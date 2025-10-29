// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessageAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessageAppResponse
	GetStatusCode() *int32
	SetBody(v *ListMessageAppResponseBody) *ListMessageAppResponse
	GetBody() *ListMessageAppResponseBody
}

type ListMessageAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessageAppResponse) GoString() string {
	return s.String()
}

func (s *ListMessageAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessageAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessageAppResponse) GetBody() *ListMessageAppResponseBody {
	return s.Body
}

func (s *ListMessageAppResponse) SetHeaders(v map[string]*string) *ListMessageAppResponse {
	s.Headers = v
	return s
}

func (s *ListMessageAppResponse) SetStatusCode(v int32) *ListMessageAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageAppResponse) SetBody(v *ListMessageAppResponseBody) *ListMessageAppResponse {
	s.Body = v
	return s
}

func (s *ListMessageAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
