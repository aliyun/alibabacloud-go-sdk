// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeClusterResponseBody) *UpgradeClusterResponse
	GetBody() *UpgradeClusterResponseBody
}

type UpgradeClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeClusterResponse) GetBody() *UpgradeClusterResponseBody {
	return s.Body
}

func (s *UpgradeClusterResponse) SetHeaders(v map[string]*string) *UpgradeClusterResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClusterResponse) SetStatusCode(v int32) *UpgradeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeClusterResponse) SetBody(v *UpgradeClusterResponseBody) *UpgradeClusterResponse {
	s.Body = v
	return s
}

func (s *UpgradeClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
