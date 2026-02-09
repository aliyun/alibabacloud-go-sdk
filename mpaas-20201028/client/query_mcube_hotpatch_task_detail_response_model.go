// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeHotpatchTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMcubeHotpatchTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMcubeHotpatchTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryMcubeHotpatchTaskDetailResponseBody) *QueryMcubeHotpatchTaskDetailResponse
	GetBody() *QueryMcubeHotpatchTaskDetailResponseBody
}

type QueryMcubeHotpatchTaskDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMcubeHotpatchTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMcubeHotpatchTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeHotpatchTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryMcubeHotpatchTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMcubeHotpatchTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMcubeHotpatchTaskDetailResponse) GetBody() *QueryMcubeHotpatchTaskDetailResponseBody {
	return s.Body
}

func (s *QueryMcubeHotpatchTaskDetailResponse) SetHeaders(v map[string]*string) *QueryMcubeHotpatchTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponse) SetStatusCode(v int32) *QueryMcubeHotpatchTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponse) SetBody(v *QueryMcubeHotpatchTaskDetailResponseBody) *QueryMcubeHotpatchTaskDetailResponse {
	s.Body = v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
