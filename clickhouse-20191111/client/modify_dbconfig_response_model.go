// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBConfigResponseBody) *ModifyDBConfigResponse
	GetBody() *ModifyDBConfigResponseBody
}

type ModifyDBConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBConfigResponse) GetBody() *ModifyDBConfigResponseBody {
	return s.Body
}

func (s *ModifyDBConfigResponse) SetHeaders(v map[string]*string) *ModifyDBConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBConfigResponse) SetStatusCode(v int32) *ModifyDBConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBConfigResponse) SetBody(v *ModifyDBConfigResponseBody) *ModifyDBConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDBConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
