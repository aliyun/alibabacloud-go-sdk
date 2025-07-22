// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceResponseBody) *GetInstanceResponse
	GetBody() *GetInstanceResponseBody
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceResponse) GetBody() *GetInstanceResponseBody {
	return s.Body
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

func (s *GetInstanceResponse) Validate() error {
	return dara.Validate(s)
}
