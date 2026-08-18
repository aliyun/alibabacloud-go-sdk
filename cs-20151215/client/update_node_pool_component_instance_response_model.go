// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodePoolComponentInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodePoolComponentInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodePoolComponentInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodePoolComponentInstanceResponseBody) *UpdateNodePoolComponentInstanceResponse
	GetBody() *UpdateNodePoolComponentInstanceResponseBody
}

type UpdateNodePoolComponentInstanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodePoolComponentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodePoolComponentInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodePoolComponentInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodePoolComponentInstanceResponse) GetBody() *UpdateNodePoolComponentInstanceResponseBody {
	return s.Body
}

func (s *UpdateNodePoolComponentInstanceResponse) SetHeaders(v map[string]*string) *UpdateNodePoolComponentInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodePoolComponentInstanceResponse) SetStatusCode(v int32) *UpdateNodePoolComponentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceResponse) SetBody(v *UpdateNodePoolComponentInstanceResponseBody) *UpdateNodePoolComponentInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateNodePoolComponentInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
