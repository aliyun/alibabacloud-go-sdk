// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAScriptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAScriptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAScriptsResponse
	GetStatusCode() *int32
	SetBody(v *CreateAScriptsResponseBody) *CreateAScriptsResponse
	GetBody() *CreateAScriptsResponseBody
}

type CreateAScriptsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAScriptsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAScriptsResponse) GoString() string {
	return s.String()
}

func (s *CreateAScriptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAScriptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAScriptsResponse) GetBody() *CreateAScriptsResponseBody {
	return s.Body
}

func (s *CreateAScriptsResponse) SetHeaders(v map[string]*string) *CreateAScriptsResponse {
	s.Headers = v
	return s
}

func (s *CreateAScriptsResponse) SetStatusCode(v int32) *CreateAScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAScriptsResponse) SetBody(v *CreateAScriptsResponseBody) *CreateAScriptsResponse {
	s.Body = v
	return s
}

func (s *CreateAScriptsResponse) Validate() error {
	return dara.Validate(s)
}
