// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConfigRulesResponseBody) *DeleteConfigRulesResponse
	GetBody() *DeleteConfigRulesResponseBody
}

type DeleteConfigRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConfigRulesResponse) GetBody() *DeleteConfigRulesResponseBody {
	return s.Body
}

func (s *DeleteConfigRulesResponse) SetHeaders(v map[string]*string) *DeleteConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigRulesResponse) SetStatusCode(v int32) *DeleteConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigRulesResponse) SetBody(v *DeleteConfigRulesResponseBody) *DeleteConfigRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
