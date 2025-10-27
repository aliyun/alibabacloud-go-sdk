// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttackPathWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAttackPathWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAttackPathWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAttackPathWhitelistResponseBody) *UpdateAttackPathWhitelistResponse
	GetBody() *UpdateAttackPathWhitelistResponseBody
}

type UpdateAttackPathWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAttackPathWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAttackPathWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttackPathWhitelistResponse) GoString() string {
	return s.String()
}

func (s *UpdateAttackPathWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAttackPathWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAttackPathWhitelistResponse) GetBody() *UpdateAttackPathWhitelistResponseBody {
	return s.Body
}

func (s *UpdateAttackPathWhitelistResponse) SetHeaders(v map[string]*string) *UpdateAttackPathWhitelistResponse {
	s.Headers = v
	return s
}

func (s *UpdateAttackPathWhitelistResponse) SetStatusCode(v int32) *UpdateAttackPathWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAttackPathWhitelistResponse) SetBody(v *UpdateAttackPathWhitelistResponseBody) *UpdateAttackPathWhitelistResponse {
	s.Body = v
	return s
}

func (s *UpdateAttackPathWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
