// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationSSLResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationSSLResponseBody) *ModifyApplicationSSLResponse
	GetBody() *ModifyApplicationSSLResponseBody
}

type ModifyApplicationSSLResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationSSLResponse) GetBody() *ModifyApplicationSSLResponseBody {
	return s.Body
}

func (s *ModifyApplicationSSLResponse) SetHeaders(v map[string]*string) *ModifyApplicationSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationSSLResponse) SetStatusCode(v int32) *ModifyApplicationSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationSSLResponse) SetBody(v *ModifyApplicationSSLResponseBody) *ModifyApplicationSSLResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationSSLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
