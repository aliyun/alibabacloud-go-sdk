// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeServiceInstanceResponseBody) *UpgradeServiceInstanceResponse
	GetBody() *UpgradeServiceInstanceResponseBody
}

type UpgradeServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeServiceInstanceResponse) GetBody() *UpgradeServiceInstanceResponseBody {
	return s.Body
}

func (s *UpgradeServiceInstanceResponse) SetHeaders(v map[string]*string) *UpgradeServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeServiceInstanceResponse) SetStatusCode(v int32) *UpgradeServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeServiceInstanceResponse) SetBody(v *UpgradeServiceInstanceResponseBody) *UpgradeServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *UpgradeServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
