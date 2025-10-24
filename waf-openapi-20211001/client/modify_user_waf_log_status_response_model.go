// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserWafLogStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserWafLogStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserWafLogStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserWafLogStatusResponseBody) *ModifyUserWafLogStatusResponse
	GetBody() *ModifyUserWafLogStatusResponseBody
}

type ModifyUserWafLogStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserWafLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserWafLogStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserWafLogStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserWafLogStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserWafLogStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserWafLogStatusResponse) GetBody() *ModifyUserWafLogStatusResponseBody {
	return s.Body
}

func (s *ModifyUserWafLogStatusResponse) SetHeaders(v map[string]*string) *ModifyUserWafLogStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserWafLogStatusResponse) SetStatusCode(v int32) *ModifyUserWafLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserWafLogStatusResponse) SetBody(v *ModifyUserWafLogStatusResponseBody) *ModifyUserWafLogStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyUserWafLogStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
