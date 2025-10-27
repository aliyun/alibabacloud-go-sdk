// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackPathWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAttackPathWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAttackPathWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ListAttackPathWhitelistResponseBody) *ListAttackPathWhitelistResponse
	GetBody() *ListAttackPathWhitelistResponseBody
}

type ListAttackPathWhitelistResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAttackPathWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAttackPathWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAttackPathWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ListAttackPathWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAttackPathWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAttackPathWhitelistResponse) GetBody() *ListAttackPathWhitelistResponseBody {
	return s.Body
}

func (s *ListAttackPathWhitelistResponse) SetHeaders(v map[string]*string) *ListAttackPathWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ListAttackPathWhitelistResponse) SetStatusCode(v int32) *ListAttackPathWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAttackPathWhitelistResponse) SetBody(v *ListAttackPathWhitelistResponseBody) *ListAttackPathWhitelistResponse {
	s.Body = v
	return s
}

func (s *ListAttackPathWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
