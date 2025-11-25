// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpsRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpsRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpsRulesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpsRulesResponseBody) *ModifyIpsRulesResponse
	GetBody() *ModifyIpsRulesResponseBody
}

type ModifyIpsRulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpsRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpsRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpsRulesResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpsRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpsRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpsRulesResponse) GetBody() *ModifyIpsRulesResponseBody {
	return s.Body
}

func (s *ModifyIpsRulesResponse) SetHeaders(v map[string]*string) *ModifyIpsRulesResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpsRulesResponse) SetStatusCode(v int32) *ModifyIpsRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpsRulesResponse) SetBody(v *ModifyIpsRulesResponseBody) *ModifyIpsRulesResponse {
	s.Body = v
	return s
}

func (s *ModifyIpsRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
