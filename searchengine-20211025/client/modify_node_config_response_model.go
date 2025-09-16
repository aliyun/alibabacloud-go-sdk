// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNodeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNodeConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNodeConfigResponseBody) *ModifyNodeConfigResponse
	GetBody() *ModifyNodeConfigResponseBody
}

type ModifyNodeConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNodeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNodeConfigResponse) GetBody() *ModifyNodeConfigResponseBody {
	return s.Body
}

func (s *ModifyNodeConfigResponse) SetHeaders(v map[string]*string) *ModifyNodeConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeConfigResponse) SetStatusCode(v int32) *ModifyNodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeConfigResponse) SetBody(v *ModifyNodeConfigResponseBody) *ModifyNodeConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyNodeConfigResponse) Validate() error {
	return dara.Validate(s)
}
