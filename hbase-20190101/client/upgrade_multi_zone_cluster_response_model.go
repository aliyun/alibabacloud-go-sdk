// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMultiZoneClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeMultiZoneClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeMultiZoneClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeMultiZoneClusterResponseBody) *UpgradeMultiZoneClusterResponse
	GetBody() *UpgradeMultiZoneClusterResponseBody
}

type UpgradeMultiZoneClusterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeMultiZoneClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeMultiZoneClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeMultiZoneClusterResponse) GetBody() *UpgradeMultiZoneClusterResponseBody {
	return s.Body
}

func (s *UpgradeMultiZoneClusterResponse) SetHeaders(v map[string]*string) *UpgradeMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMultiZoneClusterResponse) SetStatusCode(v int32) *UpgradeMultiZoneClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMultiZoneClusterResponse) SetBody(v *UpgradeMultiZoneClusterResponseBody) *UpgradeMultiZoneClusterResponse {
	s.Body = v
	return s
}

func (s *UpgradeMultiZoneClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
