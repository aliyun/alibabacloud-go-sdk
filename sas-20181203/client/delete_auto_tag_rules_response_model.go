// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoTagRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoTagRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoTagRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoTagRulesResponseBody) *DeleteAutoTagRulesResponse
	GetBody() *DeleteAutoTagRulesResponseBody
}

type DeleteAutoTagRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoTagRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoTagRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoTagRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoTagRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoTagRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoTagRulesResponse) GetBody() *DeleteAutoTagRulesResponseBody {
	return s.Body
}

func (s *DeleteAutoTagRulesResponse) SetHeaders(v map[string]*string) *DeleteAutoTagRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoTagRulesResponse) SetStatusCode(v int32) *DeleteAutoTagRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoTagRulesResponse) SetBody(v *DeleteAutoTagRulesResponseBody) *DeleteAutoTagRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoTagRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
