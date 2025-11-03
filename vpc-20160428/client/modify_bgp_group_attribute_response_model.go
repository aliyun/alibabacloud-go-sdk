// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBgpGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBgpGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBgpGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBgpGroupAttributeResponseBody) *ModifyBgpGroupAttributeResponse
	GetBody() *ModifyBgpGroupAttributeResponseBody
}

type ModifyBgpGroupAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBgpGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBgpGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBgpGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyBgpGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBgpGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBgpGroupAttributeResponse) GetBody() *ModifyBgpGroupAttributeResponseBody {
	return s.Body
}

func (s *ModifyBgpGroupAttributeResponse) SetHeaders(v map[string]*string) *ModifyBgpGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyBgpGroupAttributeResponse) SetStatusCode(v int32) *ModifyBgpGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBgpGroupAttributeResponse) SetBody(v *ModifyBgpGroupAttributeResponseBody) *ModifyBgpGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyBgpGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
