// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewLogstashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewLogstashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewLogstashResponse
	GetStatusCode() *int32
	SetBody(v *RenewLogstashResponseBody) *RenewLogstashResponse
	GetBody() *RenewLogstashResponseBody
}

type RenewLogstashResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewLogstashResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewLogstashResponse) GoString() string {
	return s.String()
}

func (s *RenewLogstashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewLogstashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewLogstashResponse) GetBody() *RenewLogstashResponseBody {
	return s.Body
}

func (s *RenewLogstashResponse) SetHeaders(v map[string]*string) *RenewLogstashResponse {
	s.Headers = v
	return s
}

func (s *RenewLogstashResponse) SetStatusCode(v int32) *RenewLogstashResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewLogstashResponse) SetBody(v *RenewLogstashResponseBody) *RenewLogstashResponse {
	s.Body = v
	return s
}

func (s *RenewLogstashResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
