// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeNebulaResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcubeNebulaResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcubeNebulaResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetMcubeNebulaResourceResponseBody) *GetMcubeNebulaResourceResponse
	GetBody() *GetMcubeNebulaResourceResponseBody
}

type GetMcubeNebulaResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcubeNebulaResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcubeNebulaResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaResourceResponse) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcubeNebulaResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcubeNebulaResourceResponse) GetBody() *GetMcubeNebulaResourceResponseBody {
	return s.Body
}

func (s *GetMcubeNebulaResourceResponse) SetHeaders(v map[string]*string) *GetMcubeNebulaResourceResponse {
	s.Headers = v
	return s
}

func (s *GetMcubeNebulaResourceResponse) SetStatusCode(v int32) *GetMcubeNebulaResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcubeNebulaResourceResponse) SetBody(v *GetMcubeNebulaResourceResponseBody) *GetMcubeNebulaResourceResponse {
	s.Body = v
	return s
}

func (s *GetMcubeNebulaResourceResponse) Validate() error {
	return dara.Validate(s)
}
