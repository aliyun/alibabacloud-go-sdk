// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetToolsetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetToolsetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetToolsetResponse
	GetStatusCode() *int32
	SetBody(v *Toolset) *GetToolsetResponse
	GetBody() *Toolset
}

type GetToolsetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Toolset           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetToolsetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetToolsetResponse) GoString() string {
	return s.String()
}

func (s *GetToolsetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetToolsetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetToolsetResponse) GetBody() *Toolset {
	return s.Body
}

func (s *GetToolsetResponse) SetHeaders(v map[string]*string) *GetToolsetResponse {
	s.Headers = v
	return s
}

func (s *GetToolsetResponse) SetStatusCode(v int32) *GetToolsetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetToolsetResponse) SetBody(v *Toolset) *GetToolsetResponse {
	s.Body = v
	return s
}

func (s *GetToolsetResponse) Validate() error {
	return dara.Validate(s)
}
