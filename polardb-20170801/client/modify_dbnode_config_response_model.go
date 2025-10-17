// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodeConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodeConfigResponseBody) *ModifyDBNodeConfigResponse
	GetBody() *ModifyDBNodeConfigResponseBody
}

type ModifyDBNodeConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodeConfigResponse) GetBody() *ModifyDBNodeConfigResponseBody {
	return s.Body
}

func (s *ModifyDBNodeConfigResponse) SetHeaders(v map[string]*string) *ModifyDBNodeConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodeConfigResponse) SetStatusCode(v int32) *ModifyDBNodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodeConfigResponse) SetBody(v *ModifyDBNodeConfigResponseBody) *ModifyDBNodeConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
