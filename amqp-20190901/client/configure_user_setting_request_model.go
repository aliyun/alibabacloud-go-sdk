// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ConfigureUserSettingRequest
	GetBucketName() *string
	SetConsoleSessionId(v string) *ConfigureUserSettingRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *ConfigureUserSettingRequest
	GetInstanceId() *string
	SetLogstore(v string) *ConfigureUserSettingRequest
	GetLogstore() *string
	SetProjectName(v string) *ConfigureUserSettingRequest
	GetProjectName() *string
	SetPutType(v string) *ConfigureUserSettingRequest
	GetPutType() *string
}

type ConfigureUserSettingRequest struct {
	BucketName       *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Logstore         *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	ProjectName      *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	PutType *string `json:"PutType,omitempty" xml:"PutType,omitempty"`
}

func (s ConfigureUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureUserSettingRequest) GoString() string {
	return s.String()
}

func (s *ConfigureUserSettingRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ConfigureUserSettingRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ConfigureUserSettingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigureUserSettingRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *ConfigureUserSettingRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ConfigureUserSettingRequest) GetPutType() *string {
	return s.PutType
}

func (s *ConfigureUserSettingRequest) SetBucketName(v string) *ConfigureUserSettingRequest {
	s.BucketName = &v
	return s
}

func (s *ConfigureUserSettingRequest) SetConsoleSessionId(v string) *ConfigureUserSettingRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ConfigureUserSettingRequest) SetInstanceId(v string) *ConfigureUserSettingRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigureUserSettingRequest) SetLogstore(v string) *ConfigureUserSettingRequest {
	s.Logstore = &v
	return s
}

func (s *ConfigureUserSettingRequest) SetProjectName(v string) *ConfigureUserSettingRequest {
	s.ProjectName = &v
	return s
}

func (s *ConfigureUserSettingRequest) SetPutType(v string) *ConfigureUserSettingRequest {
	s.PutType = &v
	return s
}

func (s *ConfigureUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
