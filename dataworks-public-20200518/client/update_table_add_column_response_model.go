// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableAddColumnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTableAddColumnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTableAddColumnResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTableAddColumnResponseBody) *UpdateTableAddColumnResponse
	GetBody() *UpdateTableAddColumnResponseBody
}

type UpdateTableAddColumnResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableAddColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableAddColumnResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableAddColumnResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableAddColumnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTableAddColumnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTableAddColumnResponse) GetBody() *UpdateTableAddColumnResponseBody {
	return s.Body
}

func (s *UpdateTableAddColumnResponse) SetHeaders(v map[string]*string) *UpdateTableAddColumnResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableAddColumnResponse) SetStatusCode(v int32) *UpdateTableAddColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableAddColumnResponse) SetBody(v *UpdateTableAddColumnResponseBody) *UpdateTableAddColumnResponse {
	s.Body = v
	return s
}

func (s *UpdateTableAddColumnResponse) Validate() error {
	return dara.Validate(s)
}
