// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFilterResponse
	GetStatusCode() *int32
	SetBody(v *CreateFilterResponseBody) *CreateFilterResponse
	GetBody() *CreateFilterResponseBody
}

type CreateFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFilterResponse) GoString() string {
	return s.String()
}

func (s *CreateFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFilterResponse) GetBody() *CreateFilterResponseBody {
	return s.Body
}

func (s *CreateFilterResponse) SetHeaders(v map[string]*string) *CreateFilterResponse {
	s.Headers = v
	return s
}

func (s *CreateFilterResponse) SetStatusCode(v int32) *CreateFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFilterResponse) SetBody(v *CreateFilterResponseBody) *CreateFilterResponse {
	s.Body = v
	return s
}

func (s *CreateFilterResponse) Validate() error {
	return dara.Validate(s)
}
