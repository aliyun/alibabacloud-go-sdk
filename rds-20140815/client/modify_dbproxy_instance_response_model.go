// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBProxyInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBProxyInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBProxyInstanceResponseBody) *ModifyDBProxyInstanceResponse
	GetBody() *ModifyDBProxyInstanceResponseBody
}

type ModifyDBProxyInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBProxyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBProxyInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBProxyInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBProxyInstanceResponse) GetBody() *ModifyDBProxyInstanceResponseBody {
	return s.Body
}

func (s *ModifyDBProxyInstanceResponse) SetHeaders(v map[string]*string) *ModifyDBProxyInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBProxyInstanceResponse) SetStatusCode(v int32) *ModifyDBProxyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBProxyInstanceResponse) SetBody(v *ModifyDBProxyInstanceResponseBody) *ModifyDBProxyInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyDBProxyInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
