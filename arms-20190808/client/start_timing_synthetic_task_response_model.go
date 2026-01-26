// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTimingSyntheticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTimingSyntheticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTimingSyntheticTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartTimingSyntheticTaskResponseBody) *StartTimingSyntheticTaskResponse
	GetBody() *StartTimingSyntheticTaskResponseBody
}

type StartTimingSyntheticTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTimingSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTimingSyntheticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTimingSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *StartTimingSyntheticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTimingSyntheticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTimingSyntheticTaskResponse) GetBody() *StartTimingSyntheticTaskResponseBody {
	return s.Body
}

func (s *StartTimingSyntheticTaskResponse) SetHeaders(v map[string]*string) *StartTimingSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *StartTimingSyntheticTaskResponse) SetStatusCode(v int32) *StartTimingSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTimingSyntheticTaskResponse) SetBody(v *StartTimingSyntheticTaskResponseBody) *StartTimingSyntheticTaskResponse {
	s.Body = v
	return s
}

func (s *StartTimingSyntheticTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
