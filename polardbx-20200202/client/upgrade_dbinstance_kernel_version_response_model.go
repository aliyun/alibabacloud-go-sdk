// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceKernelVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBInstanceKernelVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBInstanceKernelVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBInstanceKernelVersionResponseBody) *UpgradeDBInstanceKernelVersionResponse
	GetBody() *UpgradeDBInstanceKernelVersionResponseBody
}

type UpgradeDBInstanceKernelVersionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBInstanceKernelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBInstanceKernelVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBInstanceKernelVersionResponse) GetBody() *UpgradeDBInstanceKernelVersionResponseBody {
	return s.Body
}

func (s *UpgradeDBInstanceKernelVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceKernelVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponse) SetStatusCode(v int32) *UpgradeDBInstanceKernelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponse) SetBody(v *UpgradeDBInstanceKernelVersionResponseBody) *UpgradeDBInstanceKernelVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponse) Validate() error {
	return dara.Validate(s)
}
