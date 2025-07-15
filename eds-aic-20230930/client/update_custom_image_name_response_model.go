// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomImageNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomImageNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomImageNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomImageNameResponseBody) *UpdateCustomImageNameResponse
	GetBody() *UpdateCustomImageNameResponseBody
}

type UpdateCustomImageNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomImageNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomImageNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomImageNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomImageNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomImageNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomImageNameResponse) GetBody() *UpdateCustomImageNameResponseBody {
	return s.Body
}

func (s *UpdateCustomImageNameResponse) SetHeaders(v map[string]*string) *UpdateCustomImageNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomImageNameResponse) SetStatusCode(v int32) *UpdateCustomImageNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomImageNameResponse) SetBody(v *UpdateCustomImageNameResponseBody) *UpdateCustomImageNameResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomImageNameResponse) Validate() error {
	return dara.Validate(s)
}
