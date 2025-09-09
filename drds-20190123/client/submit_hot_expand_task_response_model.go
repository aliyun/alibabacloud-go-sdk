// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotExpandTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitHotExpandTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitHotExpandTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitHotExpandTaskResponseBody) *SubmitHotExpandTaskResponse
	GetBody() *SubmitHotExpandTaskResponseBody
}

type SubmitHotExpandTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHotExpandTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHotExpandTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitHotExpandTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitHotExpandTaskResponse) GetBody() *SubmitHotExpandTaskResponseBody {
	return s.Body
}

func (s *SubmitHotExpandTaskResponse) SetHeaders(v map[string]*string) *SubmitHotExpandTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotExpandTaskResponse) SetStatusCode(v int32) *SubmitHotExpandTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotExpandTaskResponse) SetBody(v *SubmitHotExpandTaskResponseBody) *SubmitHotExpandTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitHotExpandTaskResponse) Validate() error {
	return dara.Validate(s)
}
