// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOtaTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOtaTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOtaTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListOtaTaskResponseBody) *ListOtaTaskResponse
	GetBody() *ListOtaTaskResponseBody
}

type ListOtaTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOtaTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOtaTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOtaTaskResponse) GetBody() *ListOtaTaskResponseBody {
	return s.Body
}

func (s *ListOtaTaskResponse) SetHeaders(v map[string]*string) *ListOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *ListOtaTaskResponse) SetStatusCode(v int32) *ListOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOtaTaskResponse) SetBody(v *ListOtaTaskResponseBody) *ListOtaTaskResponse {
	s.Body = v
	return s
}

func (s *ListOtaTaskResponse) Validate() error {
	return dara.Validate(s)
}
