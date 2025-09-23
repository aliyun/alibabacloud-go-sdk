// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceShareResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceShareResponseBody) *CreateResourceShareResponse
	GetBody() *CreateResourceShareResponseBody
}

type CreateResourceShareResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceShareResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceShareResponse) GetBody() *CreateResourceShareResponseBody {
	return s.Body
}

func (s *CreateResourceShareResponse) SetHeaders(v map[string]*string) *CreateResourceShareResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceShareResponse) SetStatusCode(v int32) *CreateResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceShareResponse) SetBody(v *CreateResourceShareResponseBody) *CreateResourceShareResponse {
	s.Body = v
	return s
}

func (s *CreateResourceShareResponse) Validate() error {
	return dara.Validate(s)
}
