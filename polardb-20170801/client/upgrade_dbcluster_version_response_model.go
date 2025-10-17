// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBClusterVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBClusterVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBClusterVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBClusterVersionResponseBody) *UpgradeDBClusterVersionResponse
	GetBody() *UpgradeDBClusterVersionResponseBody
}

type UpgradeDBClusterVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBClusterVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBClusterVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBClusterVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBClusterVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBClusterVersionResponse) GetBody() *UpgradeDBClusterVersionResponseBody {
	return s.Body
}

func (s *UpgradeDBClusterVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBClusterVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBClusterVersionResponse) SetStatusCode(v int32) *UpgradeDBClusterVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBClusterVersionResponse) SetBody(v *UpgradeDBClusterVersionResponseBody) *UpgradeDBClusterVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBClusterVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
