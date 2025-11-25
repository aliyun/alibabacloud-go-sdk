// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTlsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTlsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTlsConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTlsConfigResponseBody) *ModifyTlsConfigResponse
	GetBody() *ModifyTlsConfigResponseBody
}

type ModifyTlsConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTlsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTlsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTlsConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyTlsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTlsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTlsConfigResponse) GetBody() *ModifyTlsConfigResponseBody {
	return s.Body
}

func (s *ModifyTlsConfigResponse) SetHeaders(v map[string]*string) *ModifyTlsConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyTlsConfigResponse) SetStatusCode(v int32) *ModifyTlsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTlsConfigResponse) SetBody(v *ModifyTlsConfigResponseBody) *ModifyTlsConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyTlsConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
