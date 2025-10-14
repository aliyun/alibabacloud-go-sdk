// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventRulesResponseBody) *DeleteEventRulesResponse
	GetBody() *DeleteEventRulesResponseBody
}

type DeleteEventRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventRulesResponse) GetBody() *DeleteEventRulesResponseBody {
	return s.Body
}

func (s *DeleteEventRulesResponse) SetHeaders(v map[string]*string) *DeleteEventRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRulesResponse) SetStatusCode(v int32) *DeleteEventRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventRulesResponse) SetBody(v *DeleteEventRulesResponseBody) *DeleteEventRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteEventRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
