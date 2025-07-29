// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelClusterUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelClusterUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelClusterUpgradeResponse
	GetStatusCode() *int32
}

type CancelClusterUpgradeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CancelClusterUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelClusterUpgradeResponse) GoString() string {
	return s.String()
}

func (s *CancelClusterUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelClusterUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelClusterUpgradeResponse) SetHeaders(v map[string]*string) *CancelClusterUpgradeResponse {
	s.Headers = v
	return s
}

func (s *CancelClusterUpgradeResponse) SetStatusCode(v int32) *CancelClusterUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelClusterUpgradeResponse) Validate() error {
	return dara.Validate(s)
}
