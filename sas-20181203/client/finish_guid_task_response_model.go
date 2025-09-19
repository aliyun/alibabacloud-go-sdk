// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishGuidTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FinishGuidTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FinishGuidTaskResponse
	GetStatusCode() *int32
	SetBody(v *FinishGuidTaskResponseBody) *FinishGuidTaskResponse
	GetBody() *FinishGuidTaskResponseBody
}

type FinishGuidTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishGuidTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishGuidTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s FinishGuidTaskResponse) GoString() string {
	return s.String()
}

func (s *FinishGuidTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FinishGuidTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FinishGuidTaskResponse) GetBody() *FinishGuidTaskResponseBody {
	return s.Body
}

func (s *FinishGuidTaskResponse) SetHeaders(v map[string]*string) *FinishGuidTaskResponse {
	s.Headers = v
	return s
}

func (s *FinishGuidTaskResponse) SetStatusCode(v int32) *FinishGuidTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishGuidTaskResponse) SetBody(v *FinishGuidTaskResponseBody) *FinishGuidTaskResponse {
	s.Body = v
	return s
}

func (s *FinishGuidTaskResponse) Validate() error {
	return dara.Validate(s)
}
