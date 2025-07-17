// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTimingSyntheticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTimingSyntheticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTimingSyntheticTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetTimingSyntheticTaskResponseBody) *GetTimingSyntheticTaskResponse
	GetBody() *GetTimingSyntheticTaskResponseBody
}

type GetTimingSyntheticTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTimingSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTimingSyntheticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTimingSyntheticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTimingSyntheticTaskResponse) GetBody() *GetTimingSyntheticTaskResponseBody {
	return s.Body
}

func (s *GetTimingSyntheticTaskResponse) SetHeaders(v map[string]*string) *GetTimingSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTimingSyntheticTaskResponse) SetStatusCode(v int32) *GetTimingSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTimingSyntheticTaskResponse) SetBody(v *GetTimingSyntheticTaskResponseBody) *GetTimingSyntheticTaskResponse {
	s.Body = v
	return s
}

func (s *GetTimingSyntheticTaskResponse) Validate() error {
	return dara.Validate(s)
}
