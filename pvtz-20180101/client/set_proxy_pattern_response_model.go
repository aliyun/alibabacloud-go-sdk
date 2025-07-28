// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetProxyPatternResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetProxyPatternResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetProxyPatternResponse
	GetStatusCode() *int32
	SetBody(v *SetProxyPatternResponseBody) *SetProxyPatternResponse
	GetBody() *SetProxyPatternResponseBody
}

type SetProxyPatternResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetProxyPatternResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetProxyPatternResponse) String() string {
	return dara.Prettify(s)
}

func (s SetProxyPatternResponse) GoString() string {
	return s.String()
}

func (s *SetProxyPatternResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetProxyPatternResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetProxyPatternResponse) GetBody() *SetProxyPatternResponseBody {
	return s.Body
}

func (s *SetProxyPatternResponse) SetHeaders(v map[string]*string) *SetProxyPatternResponse {
	s.Headers = v
	return s
}

func (s *SetProxyPatternResponse) SetStatusCode(v int32) *SetProxyPatternResponse {
	s.StatusCode = &v
	return s
}

func (s *SetProxyPatternResponse) SetBody(v *SetProxyPatternResponseBody) *SetProxyPatternResponse {
	s.Body = v
	return s
}

func (s *SetProxyPatternResponse) Validate() error {
	return dara.Validate(s)
}
