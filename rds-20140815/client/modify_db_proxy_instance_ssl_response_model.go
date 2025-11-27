// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDbProxyInstanceSslResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDbProxyInstanceSslResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDbProxyInstanceSslResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDbProxyInstanceSslResponseBody) *ModifyDbProxyInstanceSslResponse
	GetBody() *ModifyDbProxyInstanceSslResponseBody
}

type ModifyDbProxyInstanceSslResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDbProxyInstanceSslResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDbProxyInstanceSslResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDbProxyInstanceSslResponse) GoString() string {
	return s.String()
}

func (s *ModifyDbProxyInstanceSslResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDbProxyInstanceSslResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDbProxyInstanceSslResponse) GetBody() *ModifyDbProxyInstanceSslResponseBody {
	return s.Body
}

func (s *ModifyDbProxyInstanceSslResponse) SetHeaders(v map[string]*string) *ModifyDbProxyInstanceSslResponse {
	s.Headers = v
	return s
}

func (s *ModifyDbProxyInstanceSslResponse) SetStatusCode(v int32) *ModifyDbProxyInstanceSslResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDbProxyInstanceSslResponse) SetBody(v *ModifyDbProxyInstanceSslResponseBody) *ModifyDbProxyInstanceSslResponse {
	s.Body = v
	return s
}

func (s *ModifyDbProxyInstanceSslResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
