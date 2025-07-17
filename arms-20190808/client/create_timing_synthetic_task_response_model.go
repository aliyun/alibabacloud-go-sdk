// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTimingSyntheticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTimingSyntheticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTimingSyntheticTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTimingSyntheticTaskResponseBody) *CreateTimingSyntheticTaskResponse
	GetBody() *CreateTimingSyntheticTaskResponseBody
}

type CreateTimingSyntheticTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTimingSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTimingSyntheticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTimingSyntheticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTimingSyntheticTaskResponse) GetBody() *CreateTimingSyntheticTaskResponseBody {
	return s.Body
}

func (s *CreateTimingSyntheticTaskResponse) SetHeaders(v map[string]*string) *CreateTimingSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTimingSyntheticTaskResponse) SetStatusCode(v int32) *CreateTimingSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTimingSyntheticTaskResponse) SetBody(v *CreateTimingSyntheticTaskResponseBody) *CreateTimingSyntheticTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTimingSyntheticTaskResponse) Validate() error {
	return dara.Validate(s)
}
