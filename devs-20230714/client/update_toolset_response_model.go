// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateToolsetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateToolsetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateToolsetResponse
	GetStatusCode() *int32
	SetBody(v *Toolset) *UpdateToolsetResponse
	GetBody() *Toolset
}

type UpdateToolsetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Toolset           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateToolsetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateToolsetResponse) GoString() string {
	return s.String()
}

func (s *UpdateToolsetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateToolsetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateToolsetResponse) GetBody() *Toolset {
	return s.Body
}

func (s *UpdateToolsetResponse) SetHeaders(v map[string]*string) *UpdateToolsetResponse {
	s.Headers = v
	return s
}

func (s *UpdateToolsetResponse) SetStatusCode(v int32) *UpdateToolsetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateToolsetResponse) SetBody(v *Toolset) *UpdateToolsetResponse {
	s.Body = v
	return s
}

func (s *UpdateToolsetResponse) Validate() error {
	return dara.Validate(s)
}
