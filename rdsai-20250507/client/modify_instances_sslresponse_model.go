// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancesSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstancesSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstancesSSLResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstancesSSLResponseBody) *ModifyInstancesSSLResponse
	GetBody() *ModifyInstancesSSLResponseBody
}

type ModifyInstancesSSLResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstancesSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstancesSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancesSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstancesSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstancesSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstancesSSLResponse) GetBody() *ModifyInstancesSSLResponseBody {
	return s.Body
}

func (s *ModifyInstancesSSLResponse) SetHeaders(v map[string]*string) *ModifyInstancesSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstancesSSLResponse) SetStatusCode(v int32) *ModifyInstancesSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstancesSSLResponse) SetBody(v *ModifyInstancesSSLResponseBody) *ModifyInstancesSSLResponse {
	s.Body = v
	return s
}

func (s *ModifyInstancesSSLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
