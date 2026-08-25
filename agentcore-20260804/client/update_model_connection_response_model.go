// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelConnectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelConnectionResponseBody) *UpdateModelConnectionResponse
	GetBody() *UpdateModelConnectionResponseBody
}

type UpdateModelConnectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelConnectionResponse) GetBody() *UpdateModelConnectionResponseBody {
	return s.Body
}

func (s *UpdateModelConnectionResponse) SetHeaders(v map[string]*string) *UpdateModelConnectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelConnectionResponse) SetStatusCode(v int32) *UpdateModelConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelConnectionResponse) SetBody(v *UpdateModelConnectionResponseBody) *UpdateModelConnectionResponse {
	s.Body = v
	return s
}

func (s *UpdateModelConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
