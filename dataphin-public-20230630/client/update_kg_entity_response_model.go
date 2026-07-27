// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKgEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKgEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKgEntityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKgEntityResponseBody) *UpdateKgEntityResponse
	GetBody() *UpdateKgEntityResponseBody
}

type UpdateKgEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKgEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKgEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateKgEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKgEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKgEntityResponse) GetBody() *UpdateKgEntityResponseBody {
	return s.Body
}

func (s *UpdateKgEntityResponse) SetHeaders(v map[string]*string) *UpdateKgEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateKgEntityResponse) SetStatusCode(v int32) *UpdateKgEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKgEntityResponse) SetBody(v *UpdateKgEntityResponseBody) *UpdateKgEntityResponse {
	s.Body = v
	return s
}

func (s *UpdateKgEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
