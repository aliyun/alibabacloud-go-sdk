// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNodePoolAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNodePoolAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNodePoolAttributeResponseBody) *ModifyNodePoolAttributeResponse
	GetBody() *ModifyNodePoolAttributeResponseBody
}

type ModifyNodePoolAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodePoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodePoolAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNodePoolAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNodePoolAttributeResponse) GetBody() *ModifyNodePoolAttributeResponseBody {
	return s.Body
}

func (s *ModifyNodePoolAttributeResponse) SetHeaders(v map[string]*string) *ModifyNodePoolAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodePoolAttributeResponse) SetStatusCode(v int32) *ModifyNodePoolAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodePoolAttributeResponse) SetBody(v *ModifyNodePoolAttributeResponseBody) *ModifyNodePoolAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyNodePoolAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
