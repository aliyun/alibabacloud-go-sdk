// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeKernelVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeKernelVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeKernelVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeKernelVersionResponseBody) *UpgradeKernelVersionResponse
	GetBody() *UpgradeKernelVersionResponseBody
}

type UpgradeKernelVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeKernelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeKernelVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeKernelVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeKernelVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeKernelVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeKernelVersionResponse) GetBody() *UpgradeKernelVersionResponseBody {
	return s.Body
}

func (s *UpgradeKernelVersionResponse) SetHeaders(v map[string]*string) *UpgradeKernelVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeKernelVersionResponse) SetStatusCode(v int32) *UpgradeKernelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeKernelVersionResponse) SetBody(v *UpgradeKernelVersionResponseBody) *UpgradeKernelVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeKernelVersionResponse) Validate() error {
	return dara.Validate(s)
}
