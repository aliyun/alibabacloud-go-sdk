// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncModifyAgLoginEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsyncModifyAgLoginEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsyncModifyAgLoginEmailResponse
	GetStatusCode() *int32
	SetBody(v *AsyncModifyAgLoginEmailResponseBody) *AsyncModifyAgLoginEmailResponse
	GetBody() *AsyncModifyAgLoginEmailResponseBody
}

type AsyncModifyAgLoginEmailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncModifyAgLoginEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsyncModifyAgLoginEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s AsyncModifyAgLoginEmailResponse) GoString() string {
	return s.String()
}

func (s *AsyncModifyAgLoginEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsyncModifyAgLoginEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsyncModifyAgLoginEmailResponse) GetBody() *AsyncModifyAgLoginEmailResponseBody {
	return s.Body
}

func (s *AsyncModifyAgLoginEmailResponse) SetHeaders(v map[string]*string) *AsyncModifyAgLoginEmailResponse {
	s.Headers = v
	return s
}

func (s *AsyncModifyAgLoginEmailResponse) SetStatusCode(v int32) *AsyncModifyAgLoginEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *AsyncModifyAgLoginEmailResponse) SetBody(v *AsyncModifyAgLoginEmailResponseBody) *AsyncModifyAgLoginEmailResponse {
	s.Body = v
	return s
}

func (s *AsyncModifyAgLoginEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
