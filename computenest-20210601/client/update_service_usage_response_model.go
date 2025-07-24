// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceUsageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceUsageResponseBody) *UpdateServiceUsageResponse
	GetBody() *UpdateServiceUsageResponseBody
}

type UpdateServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceUsageResponse) GetBody() *UpdateServiceUsageResponseBody {
	return s.Body
}

func (s *UpdateServiceUsageResponse) SetHeaders(v map[string]*string) *UpdateServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceUsageResponse) SetStatusCode(v int32) *UpdateServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceUsageResponse) SetBody(v *UpdateServiceUsageResponseBody) *UpdateServiceUsageResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceUsageResponse) Validate() error {
	return dara.Validate(s)
}
