// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIndexJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitIndexJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitIndexJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitIndexJobResponseBody) *SubmitIndexJobResponse
	GetBody() *SubmitIndexJobResponseBody
}

type SubmitIndexJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIndexJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIndexJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitIndexJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitIndexJobResponse) GetBody() *SubmitIndexJobResponseBody {
	return s.Body
}

func (s *SubmitIndexJobResponse) SetHeaders(v map[string]*string) *SubmitIndexJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitIndexJobResponse) SetStatusCode(v int32) *SubmitIndexJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIndexJobResponse) SetBody(v *SubmitIndexJobResponseBody) *SubmitIndexJobResponse {
	s.Body = v
	return s
}

func (s *SubmitIndexJobResponse) Validate() error {
	return dara.Validate(s)
}
