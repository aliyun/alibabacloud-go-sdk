// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileTagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileTagResponseBody) *UpdateFileTagResponse
	GetBody() *UpdateFileTagResponseBody
}

type UpdateFileTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileTagResponse) GetBody() *UpdateFileTagResponseBody {
	return s.Body
}

func (s *UpdateFileTagResponse) SetHeaders(v map[string]*string) *UpdateFileTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileTagResponse) SetStatusCode(v int32) *UpdateFileTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileTagResponse) SetBody(v *UpdateFileTagResponseBody) *UpdateFileTagResponse {
	s.Body = v
	return s
}

func (s *UpdateFileTagResponse) Validate() error {
	return dara.Validate(s)
}
