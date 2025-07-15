// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveMpuTaskSeiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveMpuTaskSeiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveMpuTaskSeiResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveMpuTaskSeiResponseBody) *SetLiveMpuTaskSeiResponse
	GetBody() *SetLiveMpuTaskSeiResponseBody
}

type SetLiveMpuTaskSeiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveMpuTaskSeiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveMpuTaskSeiResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveMpuTaskSeiResponse) GoString() string {
	return s.String()
}

func (s *SetLiveMpuTaskSeiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveMpuTaskSeiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveMpuTaskSeiResponse) GetBody() *SetLiveMpuTaskSeiResponseBody {
	return s.Body
}

func (s *SetLiveMpuTaskSeiResponse) SetHeaders(v map[string]*string) *SetLiveMpuTaskSeiResponse {
	s.Headers = v
	return s
}

func (s *SetLiveMpuTaskSeiResponse) SetStatusCode(v int32) *SetLiveMpuTaskSeiResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveMpuTaskSeiResponse) SetBody(v *SetLiveMpuTaskSeiResponseBody) *SetLiveMpuTaskSeiResponse {
	s.Body = v
	return s
}

func (s *SetLiveMpuTaskSeiResponse) Validate() error {
	return dara.Validate(s)
}
