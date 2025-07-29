// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetJobInfoResponseBody) *GetJobInfoResponse
	GetBody() *GetJobInfoResponseBody
}

type GetJobInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponse) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobInfoResponse) GetBody() *GetJobInfoResponseBody {
	return s.Body
}

func (s *GetJobInfoResponse) SetHeaders(v map[string]*string) *GetJobInfoResponse {
	s.Headers = v
	return s
}

func (s *GetJobInfoResponse) SetStatusCode(v int32) *GetJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobInfoResponse) SetBody(v *GetJobInfoResponseBody) *GetJobInfoResponse {
	s.Body = v
	return s
}

func (s *GetJobInfoResponse) Validate() error {
	return dara.Validate(s)
}
