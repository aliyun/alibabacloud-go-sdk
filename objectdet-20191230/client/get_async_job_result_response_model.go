// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncJobResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncJobResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncJobResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse
	GetBody() *GetAsyncJobResultResponseBody
}

type GetAsyncJobResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncJobResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncJobResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncJobResultResponse) GetBody() *GetAsyncJobResultResponseBody {
	return s.Body
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

func (s *GetAsyncJobResultResponse) Validate() error {
	return dara.Validate(s)
}
