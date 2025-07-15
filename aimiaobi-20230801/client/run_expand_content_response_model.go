// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunExpandContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunExpandContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunExpandContentResponse
	GetStatusCode() *int32
	SetBody(v *RunExpandContentResponseBody) *RunExpandContentResponse
	GetBody() *RunExpandContentResponseBody
}

type RunExpandContentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunExpandContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunExpandContentResponse) String() string {
	return dara.Prettify(s)
}

func (s RunExpandContentResponse) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunExpandContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunExpandContentResponse) GetBody() *RunExpandContentResponseBody {
	return s.Body
}

func (s *RunExpandContentResponse) SetHeaders(v map[string]*string) *RunExpandContentResponse {
	s.Headers = v
	return s
}

func (s *RunExpandContentResponse) SetStatusCode(v int32) *RunExpandContentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunExpandContentResponse) SetBody(v *RunExpandContentResponseBody) *RunExpandContentResponse {
	s.Body = v
	return s
}

func (s *RunExpandContentResponse) Validate() error {
	return dara.Validate(s)
}
