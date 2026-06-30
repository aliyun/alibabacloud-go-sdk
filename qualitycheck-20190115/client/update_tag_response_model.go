// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTagResponseBody) *UpdateTagResponse
	GetBody() *UpdateTagResponseBody
}

type UpdateTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTagResponse) GetBody() *UpdateTagResponseBody {
	return s.Body
}

func (s *UpdateTagResponse) SetHeaders(v map[string]*string) *UpdateTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateTagResponse) SetStatusCode(v int32) *UpdateTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTagResponse) SetBody(v *UpdateTagResponseBody) *UpdateTagResponse {
	s.Body = v
	return s
}

func (s *UpdateTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
