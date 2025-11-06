// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetUserSettingRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *GetUserSettingRequest
	GetInstanceId() *string
}

type GetUserSettingRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserSettingRequest) GoString() string {
	return s.String()
}

func (s *GetUserSettingRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetUserSettingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserSettingRequest) SetConsoleSessionId(v string) *GetUserSettingRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetUserSettingRequest) SetInstanceId(v string) *GetUserSettingRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
