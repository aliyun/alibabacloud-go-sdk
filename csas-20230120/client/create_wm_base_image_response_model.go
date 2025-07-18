// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmBaseImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWmBaseImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWmBaseImageResponse
	GetStatusCode() *int32
	SetBody(v *CreateWmBaseImageResponseBody) *CreateWmBaseImageResponse
	GetBody() *CreateWmBaseImageResponseBody
}

type CreateWmBaseImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWmBaseImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWmBaseImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageResponse) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWmBaseImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWmBaseImageResponse) GetBody() *CreateWmBaseImageResponseBody {
	return s.Body
}

func (s *CreateWmBaseImageResponse) SetHeaders(v map[string]*string) *CreateWmBaseImageResponse {
	s.Headers = v
	return s
}

func (s *CreateWmBaseImageResponse) SetStatusCode(v int32) *CreateWmBaseImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWmBaseImageResponse) SetBody(v *CreateWmBaseImageResponseBody) *CreateWmBaseImageResponse {
	s.Body = v
	return s
}

func (s *CreateWmBaseImageResponse) Validate() error {
	return dara.Validate(s)
}
