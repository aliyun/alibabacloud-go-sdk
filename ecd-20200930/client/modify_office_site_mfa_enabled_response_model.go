// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteMfaEnabledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOfficeSiteMfaEnabledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOfficeSiteMfaEnabledResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOfficeSiteMfaEnabledResponseBody) *ModifyOfficeSiteMfaEnabledResponse
	GetBody() *ModifyOfficeSiteMfaEnabledResponseBody
}

type ModifyOfficeSiteMfaEnabledResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOfficeSiteMfaEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOfficeSiteMfaEnabledResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOfficeSiteMfaEnabledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOfficeSiteMfaEnabledResponse) GetBody() *ModifyOfficeSiteMfaEnabledResponseBody {
	return s.Body
}

func (s *ModifyOfficeSiteMfaEnabledResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteMfaEnabledResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledResponse) SetStatusCode(v int32) *ModifyOfficeSiteMfaEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledResponse) SetBody(v *ModifyOfficeSiteMfaEnabledResponseBody) *ModifyOfficeSiteMfaEnabledResponse {
	s.Body = v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
