// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeNebulaTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcubeNebulaTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcubeNebulaTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetMcubeNebulaTaskDetailResponseBody) *GetMcubeNebulaTaskDetailResponse
	GetBody() *GetMcubeNebulaTaskDetailResponseBody
}

type GetMcubeNebulaTaskDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcubeNebulaTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcubeNebulaTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcubeNebulaTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcubeNebulaTaskDetailResponse) GetBody() *GetMcubeNebulaTaskDetailResponseBody {
	return s.Body
}

func (s *GetMcubeNebulaTaskDetailResponse) SetHeaders(v map[string]*string) *GetMcubeNebulaTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponse) SetStatusCode(v int32) *GetMcubeNebulaTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponse) SetBody(v *GetMcubeNebulaTaskDetailResponseBody) *GetMcubeNebulaTaskDetailResponse {
	s.Body = v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
