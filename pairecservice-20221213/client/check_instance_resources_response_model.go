// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInstanceResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInstanceResourcesResponse
	GetStatusCode() *int32
	SetBody(v *CheckInstanceResourcesResponseBody) *CheckInstanceResourcesResponse
	GetBody() *CheckInstanceResourcesResponseBody
}

type CheckInstanceResourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInstanceResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInstanceResourcesResponse) GetBody() *CheckInstanceResourcesResponseBody {
	return s.Body
}

func (s *CheckInstanceResourcesResponse) SetHeaders(v map[string]*string) *CheckInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceResourcesResponse) SetStatusCode(v int32) *CheckInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceResourcesResponse) SetBody(v *CheckInstanceResourcesResponseBody) *CheckInstanceResourcesResponse {
	s.Body = v
	return s
}

func (s *CheckInstanceResourcesResponse) Validate() error {
	return dara.Validate(s)
}
