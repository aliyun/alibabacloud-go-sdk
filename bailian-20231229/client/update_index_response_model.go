// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIndexResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIndexResponseBody) *UpdateIndexResponse
	GetBody() *UpdateIndexResponseBody
}

type UpdateIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIndexResponse) GoString() string {
	return s.String()
}

func (s *UpdateIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIndexResponse) GetBody() *UpdateIndexResponseBody {
	return s.Body
}

func (s *UpdateIndexResponse) SetHeaders(v map[string]*string) *UpdateIndexResponse {
	s.Headers = v
	return s
}

func (s *UpdateIndexResponse) SetStatusCode(v int32) *UpdateIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIndexResponse) SetBody(v *UpdateIndexResponseBody) *UpdateIndexResponse {
	s.Body = v
	return s
}

func (s *UpdateIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
