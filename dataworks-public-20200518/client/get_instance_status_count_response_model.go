// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStatusCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceStatusCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceStatusCountResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceStatusCountResponseBody) *GetInstanceStatusCountResponse
	GetBody() *GetInstanceStatusCountResponseBody
}

type GetInstanceStatusCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceStatusCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceStatusCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStatusCountResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceStatusCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceStatusCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceStatusCountResponse) GetBody() *GetInstanceStatusCountResponseBody {
	return s.Body
}

func (s *GetInstanceStatusCountResponse) SetHeaders(v map[string]*string) *GetInstanceStatusCountResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceStatusCountResponse) SetStatusCode(v int32) *GetInstanceStatusCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceStatusCountResponse) SetBody(v *GetInstanceStatusCountResponseBody) *GetInstanceStatusCountResponse {
	s.Body = v
	return s
}

func (s *GetInstanceStatusCountResponse) Validate() error {
	return dara.Validate(s)
}
