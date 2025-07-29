// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolNodeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNodePoolNodeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNodePoolNodeConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNodePoolNodeConfigResponseBody) *ModifyNodePoolNodeConfigResponse
	GetBody() *ModifyNodePoolNodeConfigResponseBody
}

type ModifyNodePoolNodeConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodePoolNodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodePoolNodeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolNodeConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolNodeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNodePoolNodeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNodePoolNodeConfigResponse) GetBody() *ModifyNodePoolNodeConfigResponseBody {
	return s.Body
}

func (s *ModifyNodePoolNodeConfigResponse) SetHeaders(v map[string]*string) *ModifyNodePoolNodeConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodePoolNodeConfigResponse) SetStatusCode(v int32) *ModifyNodePoolNodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodePoolNodeConfigResponse) SetBody(v *ModifyNodePoolNodeConfigResponseBody) *ModifyNodePoolNodeConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyNodePoolNodeConfigResponse) Validate() error {
	return dara.Validate(s)
}
