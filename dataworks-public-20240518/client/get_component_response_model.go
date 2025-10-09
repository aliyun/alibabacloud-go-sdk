// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComponentResponse
	GetStatusCode() *int32
	SetBody(v *GetComponentResponseBody) *GetComponentResponse
	GetBody() *GetComponentResponseBody
}

type GetComponentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComponentResponse) GoString() string {
	return s.String()
}

func (s *GetComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComponentResponse) GetBody() *GetComponentResponseBody {
	return s.Body
}

func (s *GetComponentResponse) SetHeaders(v map[string]*string) *GetComponentResponse {
	s.Headers = v
	return s
}

func (s *GetComponentResponse) SetStatusCode(v int32) *GetComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComponentResponse) SetBody(v *GetComponentResponseBody) *GetComponentResponse {
	s.Body = v
	return s
}

func (s *GetComponentResponse) Validate() error {
	return dara.Validate(s)
}
