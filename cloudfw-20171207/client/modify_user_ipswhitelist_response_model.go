// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserIPSWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserIPSWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserIPSWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserIPSWhitelistResponseBody) *ModifyUserIPSWhitelistResponse
	GetBody() *ModifyUserIPSWhitelistResponseBody
}

type ModifyUserIPSWhitelistResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserIPSWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserIPSWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserIPSWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserIPSWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserIPSWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserIPSWhitelistResponse) GetBody() *ModifyUserIPSWhitelistResponseBody {
	return s.Body
}

func (s *ModifyUserIPSWhitelistResponse) SetHeaders(v map[string]*string) *ModifyUserIPSWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserIPSWhitelistResponse) SetStatusCode(v int32) *ModifyUserIPSWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserIPSWhitelistResponse) SetBody(v *ModifyUserIPSWhitelistResponseBody) *ModifyUserIPSWhitelistResponse {
	s.Body = v
	return s
}

func (s *ModifyUserIPSWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
