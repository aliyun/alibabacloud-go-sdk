// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPrecisionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitPrecisionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitPrecisionTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitPrecisionTaskResponseBody) *SubmitPrecisionTaskResponse
	GetBody() *SubmitPrecisionTaskResponseBody
}

type SubmitPrecisionTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPrecisionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitPrecisionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitPrecisionTaskResponse) GetBody() *SubmitPrecisionTaskResponseBody {
	return s.Body
}

func (s *SubmitPrecisionTaskResponse) SetHeaders(v map[string]*string) *SubmitPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitPrecisionTaskResponse) SetStatusCode(v int32) *SubmitPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPrecisionTaskResponse) SetBody(v *SubmitPrecisionTaskResponseBody) *SubmitPrecisionTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitPrecisionTaskResponse) Validate() error {
	return dara.Validate(s)
}
