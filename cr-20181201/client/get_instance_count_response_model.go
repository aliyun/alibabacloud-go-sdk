// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceCountResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceCountResponseBody) *GetInstanceCountResponse
	GetBody() *GetInstanceCountResponseBody
}

type GetInstanceCountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCountResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceCountResponse) GetBody() *GetInstanceCountResponseBody {
	return s.Body
}

func (s *GetInstanceCountResponse) SetHeaders(v map[string]*string) *GetInstanceCountResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceCountResponse) SetStatusCode(v int32) *GetInstanceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceCountResponse) SetBody(v *GetInstanceCountResponseBody) *GetInstanceCountResponse {
	s.Body = v
	return s
}

func (s *GetInstanceCountResponse) Validate() error {
	return dara.Validate(s)
}
