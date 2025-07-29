// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseClusterUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseClusterUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseClusterUpgradeResponse
	GetStatusCode() *int32
}

type PauseClusterUpgradeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PauseClusterUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseClusterUpgradeResponse) GoString() string {
	return s.String()
}

func (s *PauseClusterUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseClusterUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseClusterUpgradeResponse) SetHeaders(v map[string]*string) *PauseClusterUpgradeResponse {
	s.Headers = v
	return s
}

func (s *PauseClusterUpgradeResponse) SetStatusCode(v int32) *PauseClusterUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseClusterUpgradeResponse) Validate() error {
	return dara.Validate(s)
}
