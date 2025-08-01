// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMseServiceApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMseServiceApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMseServiceApplicationResponse
	GetStatusCode() *int32
	SetBody(v *CreateMseServiceApplicationResponseBody) *CreateMseServiceApplicationResponse
	GetBody() *CreateMseServiceApplicationResponseBody
}

type CreateMseServiceApplicationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMseServiceApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMseServiceApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMseServiceApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateMseServiceApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMseServiceApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMseServiceApplicationResponse) GetBody() *CreateMseServiceApplicationResponseBody {
	return s.Body
}

func (s *CreateMseServiceApplicationResponse) SetHeaders(v map[string]*string) *CreateMseServiceApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateMseServiceApplicationResponse) SetStatusCode(v int32) *CreateMseServiceApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMseServiceApplicationResponse) SetBody(v *CreateMseServiceApplicationResponseBody) *CreateMseServiceApplicationResponse {
	s.Body = v
	return s
}

func (s *CreateMseServiceApplicationResponse) Validate() error {
	return dara.Validate(s)
}
