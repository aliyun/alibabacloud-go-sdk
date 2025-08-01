// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIsolationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIsolationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIsolationRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIsolationRulesResponseBody) *DeleteIsolationRulesResponse
	GetBody() *DeleteIsolationRulesResponseBody
}

type DeleteIsolationRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIsolationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIsolationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIsolationRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteIsolationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIsolationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIsolationRulesResponse) GetBody() *DeleteIsolationRulesResponseBody {
	return s.Body
}

func (s *DeleteIsolationRulesResponse) SetHeaders(v map[string]*string) *DeleteIsolationRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteIsolationRulesResponse) SetStatusCode(v int32) *DeleteIsolationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIsolationRulesResponse) SetBody(v *DeleteIsolationRulesResponseBody) *DeleteIsolationRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteIsolationRulesResponse) Validate() error {
	return dara.Validate(s)
}
