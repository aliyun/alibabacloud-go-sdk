// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLogConfigResponseBody) *ModifyLogConfigResponse
	GetBody() *ModifyLogConfigResponseBody
}

type ModifyLogConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLogConfigResponse) GetBody() *ModifyLogConfigResponseBody {
	return s.Body
}

func (s *ModifyLogConfigResponse) SetHeaders(v map[string]*string) *ModifyLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogConfigResponse) SetStatusCode(v int32) *ModifyLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogConfigResponse) SetBody(v *ModifyLogConfigResponseBody) *ModifyLogConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
