// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContext0SecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyContext0SecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyContext0SecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyContext0SecurityIpsResponseBody) *ModifyContext0SecurityIpsResponse
	GetBody() *ModifyContext0SecurityIpsResponseBody
}

type ModifyContext0SecurityIpsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyContext0SecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyContext0SecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyContext0SecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyContext0SecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyContext0SecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyContext0SecurityIpsResponse) GetBody() *ModifyContext0SecurityIpsResponseBody {
	return s.Body
}

func (s *ModifyContext0SecurityIpsResponse) SetHeaders(v map[string]*string) *ModifyContext0SecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyContext0SecurityIpsResponse) SetStatusCode(v int32) *ModifyContext0SecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponse) SetBody(v *ModifyContext0SecurityIpsResponseBody) *ModifyContext0SecurityIpsResponse {
	s.Body = v
	return s
}

func (s *ModifyContext0SecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
