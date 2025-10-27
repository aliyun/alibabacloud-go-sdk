// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackPathWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackPathWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackPathWhitelistResponseBody) *GetAttackPathWhitelistResponse
	GetBody() *GetAttackPathWhitelistResponseBody
}

type GetAttackPathWhitelistResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackPathWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackPathWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathWhitelistResponse) GoString() string {
	return s.String()
}

func (s *GetAttackPathWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackPathWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackPathWhitelistResponse) GetBody() *GetAttackPathWhitelistResponseBody {
	return s.Body
}

func (s *GetAttackPathWhitelistResponse) SetHeaders(v map[string]*string) *GetAttackPathWhitelistResponse {
	s.Headers = v
	return s
}

func (s *GetAttackPathWhitelistResponse) SetStatusCode(v int32) *GetAttackPathWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackPathWhitelistResponse) SetBody(v *GetAttackPathWhitelistResponseBody) *GetAttackPathWhitelistResponse {
	s.Body = v
	return s
}

func (s *GetAttackPathWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
