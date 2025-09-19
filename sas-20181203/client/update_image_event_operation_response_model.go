// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageEventOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImageEventOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImageEventOperationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImageEventOperationResponseBody) *UpdateImageEventOperationResponse
	GetBody() *UpdateImageEventOperationResponseBody
}

type UpdateImageEventOperationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageEventOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageEventOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageEventOperationResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageEventOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImageEventOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImageEventOperationResponse) GetBody() *UpdateImageEventOperationResponseBody {
	return s.Body
}

func (s *UpdateImageEventOperationResponse) SetHeaders(v map[string]*string) *UpdateImageEventOperationResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageEventOperationResponse) SetStatusCode(v int32) *UpdateImageEventOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageEventOperationResponse) SetBody(v *UpdateImageEventOperationResponseBody) *UpdateImageEventOperationResponse {
	s.Body = v
	return s
}

func (s *UpdateImageEventOperationResponse) Validate() error {
	return dara.Validate(s)
}
