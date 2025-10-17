// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDynamicConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDynamicConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDynamicConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDynamicConfigResponseBody) *ModifyDynamicConfigResponse
	GetBody() *ModifyDynamicConfigResponseBody
}

type ModifyDynamicConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDynamicConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDynamicConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDynamicConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDynamicConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDynamicConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDynamicConfigResponse) GetBody() *ModifyDynamicConfigResponseBody {
	return s.Body
}

func (s *ModifyDynamicConfigResponse) SetHeaders(v map[string]*string) *ModifyDynamicConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDynamicConfigResponse) SetStatusCode(v int32) *ModifyDynamicConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDynamicConfigResponse) SetBody(v *ModifyDynamicConfigResponseBody) *ModifyDynamicConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDynamicConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
