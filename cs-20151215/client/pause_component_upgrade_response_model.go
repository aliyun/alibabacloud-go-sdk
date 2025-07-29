// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseComponentUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseComponentUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseComponentUpgradeResponse
	GetStatusCode() *int32
}

type PauseComponentUpgradeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PauseComponentUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *PauseComponentUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseComponentUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseComponentUpgradeResponse) SetHeaders(v map[string]*string) *PauseComponentUpgradeResponse {
	s.Headers = v
	return s
}

func (s *PauseComponentUpgradeResponse) SetStatusCode(v int32) *PauseComponentUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseComponentUpgradeResponse) Validate() error {
	return dara.Validate(s)
}
