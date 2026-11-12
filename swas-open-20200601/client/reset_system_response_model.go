// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetSystemResponse
	GetStatusCode() *int32
	SetBody(v *ResetSystemResponseBody) *ResetSystemResponse
	GetBody() *ResetSystemResponseBody
}

type ResetSystemResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetSystemResponse) GoString() string {
	return s.String()
}

func (s *ResetSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetSystemResponse) GetBody() *ResetSystemResponseBody {
	return s.Body
}

func (s *ResetSystemResponse) SetHeaders(v map[string]*string) *ResetSystemResponse {
	s.Headers = v
	return s
}

func (s *ResetSystemResponse) SetStatusCode(v int32) *ResetSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSystemResponse) SetBody(v *ResetSystemResponseBody) *ResetSystemResponse {
	s.Body = v
	return s
}

func (s *ResetSystemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
