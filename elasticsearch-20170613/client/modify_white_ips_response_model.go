// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhiteIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWhiteIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWhiteIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWhiteIpsResponseBody) *ModifyWhiteIpsResponse
	GetBody() *ModifyWhiteIpsResponseBody
}

type ModifyWhiteIpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWhiteIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyWhiteIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWhiteIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWhiteIpsResponse) GetBody() *ModifyWhiteIpsResponseBody {
	return s.Body
}

func (s *ModifyWhiteIpsResponse) SetHeaders(v map[string]*string) *ModifyWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyWhiteIpsResponse) SetStatusCode(v int32) *ModifyWhiteIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWhiteIpsResponse) SetBody(v *ModifyWhiteIpsResponseBody) *ModifyWhiteIpsResponse {
	s.Body = v
	return s
}

func (s *ModifyWhiteIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
