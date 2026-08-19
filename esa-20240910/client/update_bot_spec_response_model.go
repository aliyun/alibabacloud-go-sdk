// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBotSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBotSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBotSpecResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBotSpecResponseBody) *UpdateBotSpecResponse
	GetBody() *UpdateBotSpecResponseBody
}

type UpdateBotSpecResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBotSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBotSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBotSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateBotSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBotSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBotSpecResponse) GetBody() *UpdateBotSpecResponseBody {
	return s.Body
}

func (s *UpdateBotSpecResponse) SetHeaders(v map[string]*string) *UpdateBotSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateBotSpecResponse) SetStatusCode(v int32) *UpdateBotSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBotSpecResponse) SetBody(v *UpdateBotSpecResponseBody) *UpdateBotSpecResponse {
	s.Body = v
	return s
}

func (s *UpdateBotSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
