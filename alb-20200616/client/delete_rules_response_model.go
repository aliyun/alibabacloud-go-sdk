// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRulesResponseBody) *DeleteRulesResponse
	GetBody() *DeleteRulesResponseBody
}

type DeleteRulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRulesResponse) GetBody() *DeleteRulesResponseBody {
	return s.Body
}

func (s *DeleteRulesResponse) SetHeaders(v map[string]*string) *DeleteRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRulesResponse) SetStatusCode(v int32) *DeleteRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRulesResponse) SetBody(v *DeleteRulesResponseBody) *DeleteRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
