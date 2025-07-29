// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterNodepoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeClusterNodepoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeClusterNodepoolResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeClusterNodepoolResponseBody) *UpgradeClusterNodepoolResponse
	GetBody() *UpgradeClusterNodepoolResponseBody
}

type UpgradeClusterNodepoolResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeClusterNodepoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeClusterNodepoolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterNodepoolResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterNodepoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeClusterNodepoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeClusterNodepoolResponse) GetBody() *UpgradeClusterNodepoolResponseBody {
	return s.Body
}

func (s *UpgradeClusterNodepoolResponse) SetHeaders(v map[string]*string) *UpgradeClusterNodepoolResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClusterNodepoolResponse) SetStatusCode(v int32) *UpgradeClusterNodepoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeClusterNodepoolResponse) SetBody(v *UpgradeClusterNodepoolResponseBody) *UpgradeClusterNodepoolResponse {
	s.Body = v
	return s
}

func (s *UpgradeClusterNodepoolResponse) Validate() error {
	return dara.Validate(s)
}
