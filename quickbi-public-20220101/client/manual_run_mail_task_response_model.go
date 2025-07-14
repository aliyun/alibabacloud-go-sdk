// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualRunMailTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManualRunMailTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManualRunMailTaskResponse
	GetStatusCode() *int32
	SetBody(v *ManualRunMailTaskResponseBody) *ManualRunMailTaskResponse
	GetBody() *ManualRunMailTaskResponseBody
}

type ManualRunMailTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualRunMailTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManualRunMailTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ManualRunMailTaskResponse) GoString() string {
	return s.String()
}

func (s *ManualRunMailTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManualRunMailTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManualRunMailTaskResponse) GetBody() *ManualRunMailTaskResponseBody {
	return s.Body
}

func (s *ManualRunMailTaskResponse) SetHeaders(v map[string]*string) *ManualRunMailTaskResponse {
	s.Headers = v
	return s
}

func (s *ManualRunMailTaskResponse) SetStatusCode(v int32) *ManualRunMailTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualRunMailTaskResponse) SetBody(v *ManualRunMailTaskResponseBody) *ManualRunMailTaskResponse {
	s.Body = v
	return s
}

func (s *ManualRunMailTaskResponse) Validate() error {
	return dara.Validate(s)
}
