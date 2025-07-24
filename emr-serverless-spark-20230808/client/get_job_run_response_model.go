// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobRunResponse
	GetStatusCode() *int32
	SetBody(v *GetJobRunResponseBody) *GetJobRunResponse
	GetBody() *GetJobRunResponseBody
}

type GetJobRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobRunResponse) GoString() string {
	return s.String()
}

func (s *GetJobRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobRunResponse) GetBody() *GetJobRunResponseBody {
	return s.Body
}

func (s *GetJobRunResponse) SetHeaders(v map[string]*string) *GetJobRunResponse {
	s.Headers = v
	return s
}

func (s *GetJobRunResponse) SetStatusCode(v int32) *GetJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobRunResponse) SetBody(v *GetJobRunResponseBody) *GetJobRunResponse {
	s.Body = v
	return s
}

func (s *GetJobRunResponse) Validate() error {
	return dara.Validate(s)
}
