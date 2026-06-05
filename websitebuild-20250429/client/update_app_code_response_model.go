// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppCodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppCodeResponseBody) *UpdateAppCodeResponse
	GetBody() *UpdateAppCodeResponseBody
}

type UpdateAppCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppCodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppCodeResponse) GetBody() *UpdateAppCodeResponseBody {
	return s.Body
}

func (s *UpdateAppCodeResponse) SetHeaders(v map[string]*string) *UpdateAppCodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppCodeResponse) SetStatusCode(v int32) *UpdateAppCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppCodeResponse) SetBody(v *UpdateAppCodeResponseBody) *UpdateAppCodeResponse {
	s.Body = v
	return s
}

func (s *UpdateAppCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
