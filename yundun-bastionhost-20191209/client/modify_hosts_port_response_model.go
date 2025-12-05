// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsPortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostsPortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostsPortResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostsPortResponseBody) *ModifyHostsPortResponse
	GetBody() *ModifyHostsPortResponseBody
}

type ModifyHostsPortResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostsPortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostsPortResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsPortResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostsPortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostsPortResponse) GetBody() *ModifyHostsPortResponseBody {
	return s.Body
}

func (s *ModifyHostsPortResponse) SetHeaders(v map[string]*string) *ModifyHostsPortResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostsPortResponse) SetStatusCode(v int32) *ModifyHostsPortResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostsPortResponse) SetBody(v *ModifyHostsPortResponseBody) *ModifyHostsPortResponse {
	s.Body = v
	return s
}

func (s *ModifyHostsPortResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
