// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIMBotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIMBotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIMBotResponse
	GetStatusCode() *int32
	SetBody(v *IMBotResult) *UpdateIMBotResponse
	GetBody() *IMBotResult
}

type UpdateIMBotResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IMBotResult       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIMBotResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIMBotResponse) GoString() string {
	return s.String()
}

func (s *UpdateIMBotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIMBotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIMBotResponse) GetBody() *IMBotResult {
	return s.Body
}

func (s *UpdateIMBotResponse) SetHeaders(v map[string]*string) *UpdateIMBotResponse {
	s.Headers = v
	return s
}

func (s *UpdateIMBotResponse) SetStatusCode(v int32) *UpdateIMBotResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIMBotResponse) SetBody(v *IMBotResult) *UpdateIMBotResponse {
	s.Body = v
	return s
}

func (s *UpdateIMBotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
