// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *CancelUserSettingRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *CancelUserSettingRequest
	GetInstanceId() *string
}

type CancelUserSettingRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CancelUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelUserSettingRequest) GoString() string {
	return s.String()
}

func (s *CancelUserSettingRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *CancelUserSettingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelUserSettingRequest) SetConsoleSessionId(v string) *CancelUserSettingRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *CancelUserSettingRequest) SetInstanceId(v string) *CancelUserSettingRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
