// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ResetApplicationResponseBody) *ResetApplicationResponse
	GetBody() *ResetApplicationResponseBody
}

type ResetApplicationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetApplicationResponse) GoString() string {
	return s.String()
}

func (s *ResetApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetApplicationResponse) GetBody() *ResetApplicationResponseBody {
	return s.Body
}

func (s *ResetApplicationResponse) SetHeaders(v map[string]*string) *ResetApplicationResponse {
	s.Headers = v
	return s
}

func (s *ResetApplicationResponse) SetStatusCode(v int32) *ResetApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetApplicationResponse) SetBody(v *ResetApplicationResponseBody) *ResetApplicationResponse {
	s.Body = v
	return s
}

func (s *ResetApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
