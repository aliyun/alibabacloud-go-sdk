// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFilterRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebFilterRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebFilterRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebFilterRulesResponseBody) *DeleteWebFilterRulesResponse
	GetBody() *DeleteWebFilterRulesResponseBody
}

type DeleteWebFilterRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebFilterRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebFilterRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFilterRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebFilterRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebFilterRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebFilterRulesResponse) GetBody() *DeleteWebFilterRulesResponseBody {
	return s.Body
}

func (s *DeleteWebFilterRulesResponse) SetHeaders(v map[string]*string) *DeleteWebFilterRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebFilterRulesResponse) SetStatusCode(v int32) *DeleteWebFilterRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebFilterRulesResponse) SetBody(v *DeleteWebFilterRulesResponseBody) *DeleteWebFilterRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteWebFilterRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
