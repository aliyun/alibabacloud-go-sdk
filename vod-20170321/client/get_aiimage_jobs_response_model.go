// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIImageJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIImageJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIImageJobsResponse
	GetStatusCode() *int32
	SetBody(v *GetAIImageJobsResponseBody) *GetAIImageJobsResponse
	GetBody() *GetAIImageJobsResponseBody
}

type GetAIImageJobsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIImageJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIImageJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIImageJobsResponse) GoString() string {
	return s.String()
}

func (s *GetAIImageJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIImageJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIImageJobsResponse) GetBody() *GetAIImageJobsResponseBody {
	return s.Body
}

func (s *GetAIImageJobsResponse) SetHeaders(v map[string]*string) *GetAIImageJobsResponse {
	s.Headers = v
	return s
}

func (s *GetAIImageJobsResponse) SetStatusCode(v int32) *GetAIImageJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIImageJobsResponse) SetBody(v *GetAIImageJobsResponseBody) *GetAIImageJobsResponse {
	s.Body = v
	return s
}

func (s *GetAIImageJobsResponse) Validate() error {
	return dara.Validate(s)
}
