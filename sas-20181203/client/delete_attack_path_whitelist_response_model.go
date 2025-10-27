// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackPathWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAttackPathWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAttackPathWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAttackPathWhitelistResponseBody) *DeleteAttackPathWhitelistResponse
	GetBody() *DeleteAttackPathWhitelistResponseBody
}

type DeleteAttackPathWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAttackPathWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAttackPathWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackPathWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteAttackPathWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAttackPathWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAttackPathWhitelistResponse) GetBody() *DeleteAttackPathWhitelistResponseBody {
	return s.Body
}

func (s *DeleteAttackPathWhitelistResponse) SetHeaders(v map[string]*string) *DeleteAttackPathWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteAttackPathWhitelistResponse) SetStatusCode(v int32) *DeleteAttackPathWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAttackPathWhitelistResponse) SetBody(v *DeleteAttackPathWhitelistResponseBody) *DeleteAttackPathWhitelistResponse {
	s.Body = v
	return s
}

func (s *DeleteAttackPathWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
