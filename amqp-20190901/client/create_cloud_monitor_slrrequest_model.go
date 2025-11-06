// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMonitorSLRRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *CreateCloudMonitorSLRRequest
	GetConsoleSessionId() *string
}

type CreateCloudMonitorSLRRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s CreateCloudMonitorSLRRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMonitorSLRRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudMonitorSLRRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *CreateCloudMonitorSLRRequest) SetConsoleSessionId(v string) *CreateCloudMonitorSLRRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *CreateCloudMonitorSLRRequest) Validate() error {
	return dara.Validate(s)
}
