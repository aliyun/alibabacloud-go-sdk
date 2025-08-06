// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetJobDetailResponseBody) *GetJobDetailResponse
	GetBody() *GetJobDetailResponseBody
}

type GetJobDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobDetailResponse) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobDetailResponse) GetBody() *GetJobDetailResponseBody {
	return s.Body
}

func (s *GetJobDetailResponse) SetHeaders(v map[string]*string) *GetJobDetailResponse {
	s.Headers = v
	return s
}

func (s *GetJobDetailResponse) SetStatusCode(v int32) *GetJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobDetailResponse) SetBody(v *GetJobDetailResponseBody) *GetJobDetailResponse {
	s.Body = v
	return s
}

func (s *GetJobDetailResponse) Validate() error {
	return dara.Validate(s)
}
