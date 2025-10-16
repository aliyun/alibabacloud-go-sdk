// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpsRulesToDefaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpsRulesToDefaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpsRulesToDefaultResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpsRulesToDefaultResponseBody) *ModifyIpsRulesToDefaultResponse
	GetBody() *ModifyIpsRulesToDefaultResponseBody
}

type ModifyIpsRulesToDefaultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpsRulesToDefaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpsRulesToDefaultResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpsRulesToDefaultResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpsRulesToDefaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpsRulesToDefaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpsRulesToDefaultResponse) GetBody() *ModifyIpsRulesToDefaultResponseBody {
	return s.Body
}

func (s *ModifyIpsRulesToDefaultResponse) SetHeaders(v map[string]*string) *ModifyIpsRulesToDefaultResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpsRulesToDefaultResponse) SetStatusCode(v int32) *ModifyIpsRulesToDefaultResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpsRulesToDefaultResponse) SetBody(v *ModifyIpsRulesToDefaultResponseBody) *ModifyIpsRulesToDefaultResponse {
	s.Body = v
	return s
}

func (s *ModifyIpsRulesToDefaultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
