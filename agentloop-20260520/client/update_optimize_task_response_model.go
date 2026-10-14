// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOptimizeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOptimizeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOptimizeTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOptimizeTaskResponseBody) *UpdateOptimizeTaskResponse
	GetBody() *UpdateOptimizeTaskResponseBody
}

type UpdateOptimizeTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOptimizeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOptimizeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOptimizeTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateOptimizeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOptimizeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOptimizeTaskResponse) GetBody() *UpdateOptimizeTaskResponseBody {
	return s.Body
}

func (s *UpdateOptimizeTaskResponse) SetHeaders(v map[string]*string) *UpdateOptimizeTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateOptimizeTaskResponse) SetStatusCode(v int32) *UpdateOptimizeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOptimizeTaskResponse) SetBody(v *UpdateOptimizeTaskResponseBody) *UpdateOptimizeTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateOptimizeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
