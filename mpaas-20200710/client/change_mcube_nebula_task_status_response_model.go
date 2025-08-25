// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubeNebulaTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeMcubeNebulaTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeMcubeNebulaTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *ChangeMcubeNebulaTaskStatusResponseBody) *ChangeMcubeNebulaTaskStatusResponse
	GetBody() *ChangeMcubeNebulaTaskStatusResponseBody
}

type ChangeMcubeNebulaTaskStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeMcubeNebulaTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeMcubeNebulaTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubeNebulaTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeMcubeNebulaTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeMcubeNebulaTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeMcubeNebulaTaskStatusResponse) GetBody() *ChangeMcubeNebulaTaskStatusResponseBody {
	return s.Body
}

func (s *ChangeMcubeNebulaTaskStatusResponse) SetHeaders(v map[string]*string) *ChangeMcubeNebulaTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponse) SetStatusCode(v int32) *ChangeMcubeNebulaTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponse) SetBody(v *ChangeMcubeNebulaTaskStatusResponseBody) *ChangeMcubeNebulaTaskStatusResponse {
	s.Body = v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
