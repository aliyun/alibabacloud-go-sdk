// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeNebulaAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeNebulaAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeNebulaAppResponseBody) *CreateMcubeNebulaAppResponse
	GetBody() *CreateMcubeNebulaAppResponseBody
}

type CreateMcubeNebulaAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeNebulaAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeNebulaAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaAppResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeNebulaAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeNebulaAppResponse) GetBody() *CreateMcubeNebulaAppResponseBody {
	return s.Body
}

func (s *CreateMcubeNebulaAppResponse) SetHeaders(v map[string]*string) *CreateMcubeNebulaAppResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeNebulaAppResponse) SetStatusCode(v int32) *CreateMcubeNebulaAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeNebulaAppResponse) SetBody(v *CreateMcubeNebulaAppResponseBody) *CreateMcubeNebulaAppResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeNebulaAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
