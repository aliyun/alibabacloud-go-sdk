// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeNebulaAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcubeNebulaAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcubeNebulaAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcubeNebulaAppResponseBody) *DeleteMcubeNebulaAppResponse
	GetBody() *DeleteMcubeNebulaAppResponseBody
}

type DeleteMcubeNebulaAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcubeNebulaAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcubeNebulaAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeNebulaAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcubeNebulaAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcubeNebulaAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcubeNebulaAppResponse) GetBody() *DeleteMcubeNebulaAppResponseBody {
	return s.Body
}

func (s *DeleteMcubeNebulaAppResponse) SetHeaders(v map[string]*string) *DeleteMcubeNebulaAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcubeNebulaAppResponse) SetStatusCode(v int32) *DeleteMcubeNebulaAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcubeNebulaAppResponse) SetBody(v *DeleteMcubeNebulaAppResponseBody) *DeleteMcubeNebulaAppResponse {
	s.Body = v
	return s
}

func (s *DeleteMcubeNebulaAppResponse) Validate() error {
	return dara.Validate(s)
}
