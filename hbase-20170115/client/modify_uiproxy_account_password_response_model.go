// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUIProxyAccountPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUIProxyAccountPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUIProxyAccountPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUIProxyAccountPasswordResponseBody) *ModifyUIProxyAccountPasswordResponse
	GetBody() *ModifyUIProxyAccountPasswordResponseBody
}

type ModifyUIProxyAccountPasswordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUIProxyAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUIProxyAccountPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUIProxyAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyUIProxyAccountPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUIProxyAccountPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUIProxyAccountPasswordResponse) GetBody() *ModifyUIProxyAccountPasswordResponseBody {
	return s.Body
}

func (s *ModifyUIProxyAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyUIProxyAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyUIProxyAccountPasswordResponse) SetStatusCode(v int32) *ModifyUIProxyAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordResponse) SetBody(v *ModifyUIProxyAccountPasswordResponseBody) *ModifyUIProxyAccountPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyUIProxyAccountPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
