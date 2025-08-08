// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeNebulaResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeNebulaResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeNebulaResourceResponseBody) *CreateMcubeNebulaResourceResponse
	GetBody() *CreateMcubeNebulaResourceResponseBody
}

type CreateMcubeNebulaResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeNebulaResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeNebulaResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeNebulaResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeNebulaResourceResponse) GetBody() *CreateMcubeNebulaResourceResponseBody {
	return s.Body
}

func (s *CreateMcubeNebulaResourceResponse) SetHeaders(v map[string]*string) *CreateMcubeNebulaResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeNebulaResourceResponse) SetStatusCode(v int32) *CreateMcubeNebulaResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponse) SetBody(v *CreateMcubeNebulaResourceResponseBody) *CreateMcubeNebulaResourceResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeNebulaResourceResponse) Validate() error {
	return dara.Validate(s)
}
