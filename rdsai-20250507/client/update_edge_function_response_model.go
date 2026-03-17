// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEdgeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEdgeFunctionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEdgeFunctionResponseBody) *UpdateEdgeFunctionResponse
	GetBody() *UpdateEdgeFunctionResponseBody
}

type UpdateEdgeFunctionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEdgeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEdgeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateEdgeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEdgeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEdgeFunctionResponse) GetBody() *UpdateEdgeFunctionResponseBody {
	return s.Body
}

func (s *UpdateEdgeFunctionResponse) SetHeaders(v map[string]*string) *UpdateEdgeFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateEdgeFunctionResponse) SetStatusCode(v int32) *UpdateEdgeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEdgeFunctionResponse) SetBody(v *UpdateEdgeFunctionResponseBody) *UpdateEdgeFunctionResponse {
	s.Body = v
	return s
}

func (s *UpdateEdgeFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
