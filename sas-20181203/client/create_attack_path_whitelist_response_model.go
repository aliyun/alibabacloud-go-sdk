// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttackPathWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAttackPathWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAttackPathWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *CreateAttackPathWhitelistResponseBody) *CreateAttackPathWhitelistResponse
	GetBody() *CreateAttackPathWhitelistResponseBody
}

type CreateAttackPathWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAttackPathWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAttackPathWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackPathWhitelistResponse) GoString() string {
	return s.String()
}

func (s *CreateAttackPathWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAttackPathWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAttackPathWhitelistResponse) GetBody() *CreateAttackPathWhitelistResponseBody {
	return s.Body
}

func (s *CreateAttackPathWhitelistResponse) SetHeaders(v map[string]*string) *CreateAttackPathWhitelistResponse {
	s.Headers = v
	return s
}

func (s *CreateAttackPathWhitelistResponse) SetStatusCode(v int32) *CreateAttackPathWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAttackPathWhitelistResponse) SetBody(v *CreateAttackPathWhitelistResponseBody) *CreateAttackPathWhitelistResponse {
	s.Body = v
	return s
}

func (s *CreateAttackPathWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
