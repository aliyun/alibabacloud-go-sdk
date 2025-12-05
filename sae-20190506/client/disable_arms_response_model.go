// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableArmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableArmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableArmsResponse
	GetStatusCode() *int32
	SetBody(v *DisableArmsResponseBody) *DisableArmsResponse
	GetBody() *DisableArmsResponseBody
}

type DisableArmsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableArmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableArmsResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableArmsResponse) GoString() string {
	return s.String()
}

func (s *DisableArmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableArmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableArmsResponse) GetBody() *DisableArmsResponseBody {
	return s.Body
}

func (s *DisableArmsResponse) SetHeaders(v map[string]*string) *DisableArmsResponse {
	s.Headers = v
	return s
}

func (s *DisableArmsResponse) SetStatusCode(v int32) *DisableArmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableArmsResponse) SetBody(v *DisableArmsResponseBody) *DisableArmsResponse {
	s.Body = v
	return s
}

func (s *DisableArmsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
