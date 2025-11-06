// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashboardListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *DashboardListRequest
	GetConsoleSessionId() *string
	SetDashboardName(v string) *DashboardListRequest
	GetDashboardName() *string
}

type DashboardListRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	DashboardName *string `json:"DashboardName,omitempty" xml:"DashboardName,omitempty"`
}

func (s DashboardListRequest) String() string {
	return dara.Prettify(s)
}

func (s DashboardListRequest) GoString() string {
	return s.String()
}

func (s *DashboardListRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DashboardListRequest) GetDashboardName() *string {
	return s.DashboardName
}

func (s *DashboardListRequest) SetConsoleSessionId(v string) *DashboardListRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DashboardListRequest) SetDashboardName(v string) *DashboardListRequest {
	s.DashboardName = &v
	return s
}

func (s *DashboardListRequest) Validate() error {
	return dara.Validate(s)
}
