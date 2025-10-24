// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainPunishStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDomainPunishStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDomainPunishStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDomainPunishStatusResponseBody) *ModifyDomainPunishStatusResponse
	GetBody() *ModifyDomainPunishStatusResponseBody
}

type ModifyDomainPunishStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDomainPunishStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDomainPunishStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainPunishStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainPunishStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDomainPunishStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDomainPunishStatusResponse) GetBody() *ModifyDomainPunishStatusResponseBody {
	return s.Body
}

func (s *ModifyDomainPunishStatusResponse) SetHeaders(v map[string]*string) *ModifyDomainPunishStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainPunishStatusResponse) SetStatusCode(v int32) *ModifyDomainPunishStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainPunishStatusResponse) SetBody(v *ModifyDomainPunishStatusResponseBody) *ModifyDomainPunishStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyDomainPunishStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
