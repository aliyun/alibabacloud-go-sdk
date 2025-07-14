// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationApmServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeApplicationApmServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeApplicationApmServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeApplicationApmServiceResponseBody) *UpgradeApplicationApmServiceResponse
	GetBody() *UpgradeApplicationApmServiceResponseBody
}

type UpgradeApplicationApmServiceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeApplicationApmServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeApplicationApmServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationApmServiceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationApmServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeApplicationApmServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeApplicationApmServiceResponse) GetBody() *UpgradeApplicationApmServiceResponseBody {
	return s.Body
}

func (s *UpgradeApplicationApmServiceResponse) SetHeaders(v map[string]*string) *UpgradeApplicationApmServiceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeApplicationApmServiceResponse) SetStatusCode(v int32) *UpgradeApplicationApmServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeApplicationApmServiceResponse) SetBody(v *UpgradeApplicationApmServiceResponseBody) *UpgradeApplicationApmServiceResponse {
	s.Body = v
	return s
}

func (s *UpgradeApplicationApmServiceResponse) Validate() error {
	return dara.Validate(s)
}
