// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOptimizeRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOptimizeRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOptimizeRunResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOptimizeRunResponseBody) *UpdateOptimizeRunResponse
	GetBody() *UpdateOptimizeRunResponseBody
}

type UpdateOptimizeRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOptimizeRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOptimizeRunResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOptimizeRunResponse) GoString() string {
	return s.String()
}

func (s *UpdateOptimizeRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOptimizeRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOptimizeRunResponse) GetBody() *UpdateOptimizeRunResponseBody {
	return s.Body
}

func (s *UpdateOptimizeRunResponse) SetHeaders(v map[string]*string) *UpdateOptimizeRunResponse {
	s.Headers = v
	return s
}

func (s *UpdateOptimizeRunResponse) SetStatusCode(v int32) *UpdateOptimizeRunResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOptimizeRunResponse) SetBody(v *UpdateOptimizeRunResponseBody) *UpdateOptimizeRunResponse {
	s.Body = v
	return s
}

func (s *UpdateOptimizeRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
