// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceUsageResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceUsageResponseBody) *CreateServiceUsageResponse
	GetBody() *CreateServiceUsageResponseBody
}

type CreateServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceUsageResponse) GetBody() *CreateServiceUsageResponseBody {
	return s.Body
}

func (s *CreateServiceUsageResponse) SetHeaders(v map[string]*string) *CreateServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceUsageResponse) SetStatusCode(v int32) *CreateServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceUsageResponse) SetBody(v *CreateServiceUsageResponseBody) *CreateServiceUsageResponse {
	s.Body = v
	return s
}

func (s *CreateServiceUsageResponse) Validate() error {
	return dara.Validate(s)
}
