// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodePoolComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodePoolComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodePoolComponentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodePoolComponentResponseBody) *UpdateNodePoolComponentResponse
	GetBody() *UpdateNodePoolComponentResponseBody
}

type UpdateNodePoolComponentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodePoolComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodePoolComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodePoolComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodePoolComponentResponse) GetBody() *UpdateNodePoolComponentResponseBody {
	return s.Body
}

func (s *UpdateNodePoolComponentResponse) SetHeaders(v map[string]*string) *UpdateNodePoolComponentResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodePoolComponentResponse) SetStatusCode(v int32) *UpdateNodePoolComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodePoolComponentResponse) SetBody(v *UpdateNodePoolComponentResponseBody) *UpdateNodePoolComponentResponse {
	s.Body = v
	return s
}

func (s *UpdateNodePoolComponentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
