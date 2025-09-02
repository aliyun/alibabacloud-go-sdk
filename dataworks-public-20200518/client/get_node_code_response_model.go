// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeCodeResponseBody) *GetNodeCodeResponse
	GetBody() *GetNodeCodeResponseBody
}

type GetNodeCodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeCodeResponse) GoString() string {
	return s.String()
}

func (s *GetNodeCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeCodeResponse) GetBody() *GetNodeCodeResponseBody {
	return s.Body
}

func (s *GetNodeCodeResponse) SetHeaders(v map[string]*string) *GetNodeCodeResponse {
	s.Headers = v
	return s
}

func (s *GetNodeCodeResponse) SetStatusCode(v int32) *GetNodeCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeCodeResponse) SetBody(v *GetNodeCodeResponseBody) *GetNodeCodeResponse {
	s.Body = v
	return s
}

func (s *GetNodeCodeResponse) Validate() error {
	return dara.Validate(s)
}
