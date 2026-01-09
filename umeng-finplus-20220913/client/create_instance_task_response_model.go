// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceTaskResponseBody) *CreateInstanceTaskResponse
	GetBody() *CreateInstanceTaskResponseBody
}

type CreateInstanceTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceTaskResponse) GetBody() *CreateInstanceTaskResponseBody {
	return s.Body
}

func (s *CreateInstanceTaskResponse) SetHeaders(v map[string]*string) *CreateInstanceTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceTaskResponse) SetStatusCode(v int32) *CreateInstanceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceTaskResponse) SetBody(v *CreateInstanceTaskResponseBody) *CreateInstanceTaskResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
