// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStatisticsByVhostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetStatisticsByVhostRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *GetStatisticsByVhostRequest
	GetInstanceId() *string
	SetVhostName(v string) *GetStatisticsByVhostRequest
	GetVhostName() *string
}

type GetStatisticsByVhostRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetStatisticsByVhostRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStatisticsByVhostRequest) GoString() string {
	return s.String()
}

func (s *GetStatisticsByVhostRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetStatisticsByVhostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetStatisticsByVhostRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetStatisticsByVhostRequest) SetConsoleSessionId(v string) *GetStatisticsByVhostRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetStatisticsByVhostRequest) SetInstanceId(v string) *GetStatisticsByVhostRequest {
	s.InstanceId = &v
	return s
}

func (s *GetStatisticsByVhostRequest) SetVhostName(v string) *GetStatisticsByVhostRequest {
	s.VhostName = &v
	return s
}

func (s *GetStatisticsByVhostRequest) Validate() error {
	return dara.Validate(s)
}
