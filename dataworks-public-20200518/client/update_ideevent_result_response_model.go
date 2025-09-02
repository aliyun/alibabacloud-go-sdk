// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIDEEventResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIDEEventResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIDEEventResultResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIDEEventResultResponseBody) *UpdateIDEEventResultResponse
	GetBody() *UpdateIDEEventResultResponseBody
}

type UpdateIDEEventResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIDEEventResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIDEEventResultResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIDEEventResultResponse) GoString() string {
	return s.String()
}

func (s *UpdateIDEEventResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIDEEventResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIDEEventResultResponse) GetBody() *UpdateIDEEventResultResponseBody {
	return s.Body
}

func (s *UpdateIDEEventResultResponse) SetHeaders(v map[string]*string) *UpdateIDEEventResultResponse {
	s.Headers = v
	return s
}

func (s *UpdateIDEEventResultResponse) SetStatusCode(v int32) *UpdateIDEEventResultResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIDEEventResultResponse) SetBody(v *UpdateIDEEventResultResponseBody) *UpdateIDEEventResultResponse {
	s.Body = v
	return s
}

func (s *UpdateIDEEventResultResponse) Validate() error {
	return dara.Validate(s)
}
