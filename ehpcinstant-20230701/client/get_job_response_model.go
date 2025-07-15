// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobResponse
	GetStatusCode() *int32
	SetBody(v *GetJobResponseBody) *GetJobResponse
	GetBody() *GetJobResponseBody
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobResponse) GetBody() *GetJobResponseBody {
	return s.Body
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

func (s *GetJobResponse) Validate() error {
	return dara.Validate(s)
}
