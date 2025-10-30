// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAxgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAxgResponse
	GetStatusCode() *int32
	SetBody(v *BindAxgResponseBody) *BindAxgResponse
	GetBody() *BindAxgResponseBody
}

type BindAxgResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxgResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAxgResponse) GoString() string {
	return s.String()
}

func (s *BindAxgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAxgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAxgResponse) GetBody() *BindAxgResponseBody {
	return s.Body
}

func (s *BindAxgResponse) SetHeaders(v map[string]*string) *BindAxgResponse {
	s.Headers = v
	return s
}

func (s *BindAxgResponse) SetStatusCode(v int32) *BindAxgResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxgResponse) SetBody(v *BindAxgResponseBody) *BindAxgResponse {
	s.Body = v
	return s
}

func (s *BindAxgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
