// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAttackCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAttackCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAttackCountResponse
	GetStatusCode() *int32
	SetBody(v *QueryAttackCountResponseBody) *QueryAttackCountResponse
	GetBody() *QueryAttackCountResponseBody
}

type QueryAttackCountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAttackCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAttackCountResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAttackCountResponse) GoString() string {
	return s.String()
}

func (s *QueryAttackCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAttackCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAttackCountResponse) GetBody() *QueryAttackCountResponseBody {
	return s.Body
}

func (s *QueryAttackCountResponse) SetHeaders(v map[string]*string) *QueryAttackCountResponse {
	s.Headers = v
	return s
}

func (s *QueryAttackCountResponse) SetStatusCode(v int32) *QueryAttackCountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAttackCountResponse) SetBody(v *QueryAttackCountResponseBody) *QueryAttackCountResponse {
	s.Body = v
	return s
}

func (s *QueryAttackCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
