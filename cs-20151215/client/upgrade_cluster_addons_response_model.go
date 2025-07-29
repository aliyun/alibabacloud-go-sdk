// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterAddonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeClusterAddonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeClusterAddonsResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeClusterAddonsResponseBody) *UpgradeClusterAddonsResponse
	GetBody() *UpgradeClusterAddonsResponseBody
}

type UpgradeClusterAddonsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeClusterAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeClusterAddonsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeClusterAddonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeClusterAddonsResponse) GetBody() *UpgradeClusterAddonsResponseBody {
	return s.Body
}

func (s *UpgradeClusterAddonsResponse) SetHeaders(v map[string]*string) *UpgradeClusterAddonsResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClusterAddonsResponse) SetStatusCode(v int32) *UpgradeClusterAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeClusterAddonsResponse) SetBody(v *UpgradeClusterAddonsResponseBody) *UpgradeClusterAddonsResponse {
	s.Body = v
	return s
}

func (s *UpgradeClusterAddonsResponse) Validate() error {
	return dara.Validate(s)
}
