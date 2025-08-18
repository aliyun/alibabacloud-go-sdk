// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateManagedTransformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateManagedTransformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateManagedTransformResponse
	GetStatusCode() *int32
	SetBody(v *UpdateManagedTransformResponseBody) *UpdateManagedTransformResponse
	GetBody() *UpdateManagedTransformResponseBody
}

type UpdateManagedTransformResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateManagedTransformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateManagedTransformResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedTransformResponse) GoString() string {
	return s.String()
}

func (s *UpdateManagedTransformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateManagedTransformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateManagedTransformResponse) GetBody() *UpdateManagedTransformResponseBody {
	return s.Body
}

func (s *UpdateManagedTransformResponse) SetHeaders(v map[string]*string) *UpdateManagedTransformResponse {
	s.Headers = v
	return s
}

func (s *UpdateManagedTransformResponse) SetStatusCode(v int32) *UpdateManagedTransformResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateManagedTransformResponse) SetBody(v *UpdateManagedTransformResponseBody) *UpdateManagedTransformResponse {
	s.Body = v
	return s
}

func (s *UpdateManagedTransformResponse) Validate() error {
	return dara.Validate(s)
}
