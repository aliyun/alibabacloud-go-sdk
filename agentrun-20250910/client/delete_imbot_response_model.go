// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIMBotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIMBotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIMBotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIMBotNoContent) *DeleteIMBotResponse
	GetBody() *DeleteIMBotNoContent
}

type DeleteIMBotResponse struct {
	Headers    map[string]*string    `json:"headers" xml:"headers"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIMBotNoContent `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIMBotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIMBotResponse) GoString() string {
	return s.String()
}

func (s *DeleteIMBotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIMBotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIMBotResponse) GetBody() *DeleteIMBotNoContent {
	return s.Body
}

func (s *DeleteIMBotResponse) SetHeaders(v map[string]*string) *DeleteIMBotResponse {
	s.Headers = v
	return s
}

func (s *DeleteIMBotResponse) SetStatusCode(v int32) *DeleteIMBotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIMBotResponse) SetBody(v *DeleteIMBotNoContent) *DeleteIMBotResponse {
	s.Body = v
	return s
}

func (s *DeleteIMBotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
