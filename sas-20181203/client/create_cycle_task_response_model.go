// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCycleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCycleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCycleTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateCycleTaskResponseBody) *CreateCycleTaskResponse
	GetBody() *CreateCycleTaskResponseBody
}

type CreateCycleTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCycleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCycleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCycleTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCycleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCycleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCycleTaskResponse) GetBody() *CreateCycleTaskResponseBody {
	return s.Body
}

func (s *CreateCycleTaskResponse) SetHeaders(v map[string]*string) *CreateCycleTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCycleTaskResponse) SetStatusCode(v int32) *CreateCycleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCycleTaskResponse) SetBody(v *CreateCycleTaskResponseBody) *CreateCycleTaskResponse {
	s.Body = v
	return s
}

func (s *CreateCycleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
