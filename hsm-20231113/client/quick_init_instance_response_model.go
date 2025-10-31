// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickInitInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuickInitInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuickInitInstanceResponse
	GetStatusCode() *int32
	SetBody(v *QuickInitInstanceResponseBody) *QuickInitInstanceResponse
	GetBody() *QuickInitInstanceResponseBody
}

type QuickInitInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuickInitInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuickInitInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QuickInitInstanceResponse) GoString() string {
	return s.String()
}

func (s *QuickInitInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuickInitInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuickInitInstanceResponse) GetBody() *QuickInitInstanceResponseBody {
	return s.Body
}

func (s *QuickInitInstanceResponse) SetHeaders(v map[string]*string) *QuickInitInstanceResponse {
	s.Headers = v
	return s
}

func (s *QuickInitInstanceResponse) SetStatusCode(v int32) *QuickInitInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QuickInitInstanceResponse) SetBody(v *QuickInitInstanceResponseBody) *QuickInitInstanceResponse {
	s.Body = v
	return s
}

func (s *QuickInitInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
