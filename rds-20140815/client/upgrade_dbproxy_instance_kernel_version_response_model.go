// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBProxyInstanceKernelVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBProxyInstanceKernelVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBProxyInstanceKernelVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBProxyInstanceKernelVersionResponseBody) *UpgradeDBProxyInstanceKernelVersionResponse
	GetBody() *UpgradeDBProxyInstanceKernelVersionResponseBody
}

type UpgradeDBProxyInstanceKernelVersionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBProxyInstanceKernelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBProxyInstanceKernelVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBProxyInstanceKernelVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBProxyInstanceKernelVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBProxyInstanceKernelVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBProxyInstanceKernelVersionResponse) GetBody() *UpgradeDBProxyInstanceKernelVersionResponseBody {
	return s.Body
}

func (s *UpgradeDBProxyInstanceKernelVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBProxyInstanceKernelVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionResponse) SetStatusCode(v int32) *UpgradeDBProxyInstanceKernelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionResponse) SetBody(v *UpgradeDBProxyInstanceKernelVersionResponseBody) *UpgradeDBProxyInstanceKernelVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
