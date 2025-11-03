// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentForClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeAgentForClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeAgentForClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeAgentForClusterResponseBody) *UpgradeAgentForClusterResponse
	GetBody() *UpgradeAgentForClusterResponseBody
}

type UpgradeAgentForClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAgentForClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAgentForClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentForClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeAgentForClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeAgentForClusterResponse) GetBody() *UpgradeAgentForClusterResponseBody {
	return s.Body
}

func (s *UpgradeAgentForClusterResponse) SetHeaders(v map[string]*string) *UpgradeAgentForClusterResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAgentForClusterResponse) SetStatusCode(v int32) *UpgradeAgentForClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAgentForClusterResponse) SetBody(v *UpgradeAgentForClusterResponseBody) *UpgradeAgentForClusterResponse {
	s.Body = v
	return s
}

func (s *UpgradeAgentForClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
