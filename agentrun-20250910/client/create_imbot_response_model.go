// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIMBotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIMBotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIMBotResponse
	GetStatusCode() *int32
	SetBody(v *IMBotResult) *CreateIMBotResponse
	GetBody() *IMBotResult
}

type CreateIMBotResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IMBotResult       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIMBotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIMBotResponse) GoString() string {
	return s.String()
}

func (s *CreateIMBotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIMBotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIMBotResponse) GetBody() *IMBotResult {
	return s.Body
}

func (s *CreateIMBotResponse) SetHeaders(v map[string]*string) *CreateIMBotResponse {
	s.Headers = v
	return s
}

func (s *CreateIMBotResponse) SetStatusCode(v int32) *CreateIMBotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIMBotResponse) SetBody(v *IMBotResult) *CreateIMBotResponse {
	s.Body = v
	return s
}

func (s *CreateIMBotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
