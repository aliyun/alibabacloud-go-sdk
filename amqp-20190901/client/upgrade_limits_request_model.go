// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLimitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *UpgradeLimitsRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *UpgradeLimitsRequest
	GetInstanceId() *string
}

type UpgradeLimitsRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpgradeLimitsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLimitsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeLimitsRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *UpgradeLimitsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeLimitsRequest) SetConsoleSessionId(v string) *UpgradeLimitsRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *UpgradeLimitsRequest) SetInstanceId(v string) *UpgradeLimitsRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeLimitsRequest) Validate() error {
	return dara.Validate(s)
}
