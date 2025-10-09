// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComponentResponse
	GetStatusCode() *int32
	SetBody(v *CreateComponentResponseBody) *CreateComponentResponse
	GetBody() *CreateComponentResponseBody
}

type CreateComponentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentResponse) GoString() string {
	return s.String()
}

func (s *CreateComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComponentResponse) GetBody() *CreateComponentResponseBody {
	return s.Body
}

func (s *CreateComponentResponse) SetHeaders(v map[string]*string) *CreateComponentResponse {
	s.Headers = v
	return s
}

func (s *CreateComponentResponse) SetStatusCode(v int32) *CreateComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComponentResponse) SetBody(v *CreateComponentResponseBody) *CreateComponentResponse {
	s.Body = v
	return s
}

func (s *CreateComponentResponse) Validate() error {
	return dara.Validate(s)
}
