// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBClusterVersionZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBClusterVersionZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBClusterVersionZonalResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBClusterVersionZonalResponseBody) *UpgradeDBClusterVersionZonalResponse
	GetBody() *UpgradeDBClusterVersionZonalResponseBody
}

type UpgradeDBClusterVersionZonalResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBClusterVersionZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBClusterVersionZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBClusterVersionZonalResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBClusterVersionZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBClusterVersionZonalResponse) GetBody() *UpgradeDBClusterVersionZonalResponseBody {
	return s.Body
}

func (s *UpgradeDBClusterVersionZonalResponse) SetHeaders(v map[string]*string) *UpgradeDBClusterVersionZonalResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBClusterVersionZonalResponse) SetStatusCode(v int32) *UpgradeDBClusterVersionZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalResponse) SetBody(v *UpgradeDBClusterVersionZonalResponseBody) *UpgradeDBClusterVersionZonalResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBClusterVersionZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
