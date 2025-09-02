// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectDetailResponseBody) *GetProjectDetailResponse
	GetBody() *GetProjectDetailResponseBody
}

type GetProjectDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *GetProjectDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectDetailResponse) GetBody() *GetProjectDetailResponseBody {
	return s.Body
}

func (s *GetProjectDetailResponse) SetHeaders(v map[string]*string) *GetProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *GetProjectDetailResponse) SetStatusCode(v int32) *GetProjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectDetailResponse) SetBody(v *GetProjectDetailResponseBody) *GetProjectDetailResponse {
	s.Body = v
	return s
}

func (s *GetProjectDetailResponse) Validate() error {
	return dara.Validate(s)
}
