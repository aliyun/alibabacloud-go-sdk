// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIMBotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIMBotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIMBotResponse
	GetStatusCode() *int32
	SetBody(v *IMBotResult) *GetIMBotResponse
	GetBody() *IMBotResult
}

type GetIMBotResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IMBotResult       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIMBotResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIMBotResponse) GoString() string {
	return s.String()
}

func (s *GetIMBotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIMBotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIMBotResponse) GetBody() *IMBotResult {
	return s.Body
}

func (s *GetIMBotResponse) SetHeaders(v map[string]*string) *GetIMBotResponse {
	s.Headers = v
	return s
}

func (s *GetIMBotResponse) SetStatusCode(v int32) *GetIMBotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIMBotResponse) SetBody(v *IMBotResult) *GetIMBotResponse {
	s.Body = v
	return s
}

func (s *GetIMBotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
