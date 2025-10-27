// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterSSLResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterSSLResponseBody) *ModifyDBClusterSSLResponse
	GetBody() *ModifyDBClusterSSLResponseBody
}

type ModifyDBClusterSSLResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterSSLResponse) GetBody() *ModifyDBClusterSSLResponseBody {
	return s.Body
}

func (s *ModifyDBClusterSSLResponse) SetHeaders(v map[string]*string) *ModifyDBClusterSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterSSLResponse) SetStatusCode(v int32) *ModifyDBClusterSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterSSLResponse) SetBody(v *ModifyDBClusterSSLResponseBody) *ModifyDBClusterSSLResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterSSLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
