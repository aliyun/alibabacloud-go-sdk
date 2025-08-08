// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateToolsetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateToolsetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateToolsetResponse
	GetStatusCode() *int32
	SetBody(v *Toolset) *CreateToolsetResponse
	GetBody() *Toolset
}

type CreateToolsetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Toolset           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateToolsetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateToolsetResponse) GoString() string {
	return s.String()
}

func (s *CreateToolsetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateToolsetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateToolsetResponse) GetBody() *Toolset {
	return s.Body
}

func (s *CreateToolsetResponse) SetHeaders(v map[string]*string) *CreateToolsetResponse {
	s.Headers = v
	return s
}

func (s *CreateToolsetResponse) SetStatusCode(v int32) *CreateToolsetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateToolsetResponse) SetBody(v *Toolset) *CreateToolsetResponse {
	s.Body = v
	return s
}

func (s *CreateToolsetResponse) Validate() error {
	return dara.Validate(s)
}
