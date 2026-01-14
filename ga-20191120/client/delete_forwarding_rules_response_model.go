// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteForwardingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteForwardingRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteForwardingRulesResponseBody) *DeleteForwardingRulesResponse
	GetBody() *DeleteForwardingRulesResponseBody
}

type DeleteForwardingRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteForwardingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardingRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteForwardingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteForwardingRulesResponse) GetBody() *DeleteForwardingRulesResponseBody {
	return s.Body
}

func (s *DeleteForwardingRulesResponse) SetHeaders(v map[string]*string) *DeleteForwardingRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteForwardingRulesResponse) SetStatusCode(v int32) *DeleteForwardingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteForwardingRulesResponse) SetBody(v *DeleteForwardingRulesResponseBody) *DeleteForwardingRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteForwardingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
