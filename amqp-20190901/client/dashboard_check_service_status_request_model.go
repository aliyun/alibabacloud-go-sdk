// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashboardCheckServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *DashboardCheckServiceStatusRequest
	GetConsoleSessionId() *string
}

type DashboardCheckServiceStatusRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s DashboardCheckServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DashboardCheckServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *DashboardCheckServiceStatusRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DashboardCheckServiceStatusRequest) SetConsoleSessionId(v string) *DashboardCheckServiceStatusRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DashboardCheckServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
