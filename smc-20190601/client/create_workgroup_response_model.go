// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkgroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkgroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkgroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkgroupResponseBody) *CreateWorkgroupResponse
	GetBody() *CreateWorkgroupResponseBody
}

type CreateWorkgroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkgroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkgroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkgroupResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkgroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkgroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkgroupResponse) GetBody() *CreateWorkgroupResponseBody {
	return s.Body
}

func (s *CreateWorkgroupResponse) SetHeaders(v map[string]*string) *CreateWorkgroupResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkgroupResponse) SetStatusCode(v int32) *CreateWorkgroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkgroupResponse) SetBody(v *CreateWorkgroupResponseBody) *CreateWorkgroupResponse {
	s.Body = v
	return s
}

func (s *CreateWorkgroupResponse) Validate() error {
	return dara.Validate(s)
}
