// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReopenInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReopenInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReopenInstanceResponseBody) *ReopenInstanceResponse
	GetBody() *ReopenInstanceResponseBody
}

type ReopenInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReopenInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReopenInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReopenInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReopenInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReopenInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReopenInstanceResponse) GetBody() *ReopenInstanceResponseBody {
	return s.Body
}

func (s *ReopenInstanceResponse) SetHeaders(v map[string]*string) *ReopenInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReopenInstanceResponse) SetStatusCode(v int32) *ReopenInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReopenInstanceResponse) SetBody(v *ReopenInstanceResponseBody) *ReopenInstanceResponse {
	s.Body = v
	return s
}

func (s *ReopenInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
