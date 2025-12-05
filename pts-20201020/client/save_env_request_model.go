// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveEnvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v *SaveEnvRequestEnv) *SaveEnvRequest
	GetEnv() *SaveEnvRequestEnv
}

type SaveEnvRequest struct {
	// The JMeter environment.
	//
	// This parameter is required.
	Env *SaveEnvRequestEnv `json:"Env,omitempty" xml:"Env,omitempty" type:"Struct"`
}

func (s SaveEnvRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveEnvRequest) GoString() string {
	return s.String()
}

func (s *SaveEnvRequest) GetEnv() *SaveEnvRequestEnv {
	return s.Env
}

func (s *SaveEnvRequest) SetEnv(v *SaveEnvRequestEnv) *SaveEnvRequest {
	s.Env = v
	return s
}

func (s *SaveEnvRequest) Validate() error {
	if s.Env != nil {
		if err := s.Env.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveEnvRequestEnv struct {
	// The ID of the JMeter environment. To create a JMeter environment, leave this parameter empty. To update a JMeter environment, specify the ID of the environment.
	//
	// example:
	//
	// 10YPA8H
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// The name of the environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-create
	EnvName *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	// The files on which the environment depends.
	//
	// This parameter is required.
	Files []*SaveEnvRequestEnvFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The extension label.
	//
	// example:
	//
	// test
	JmeterPluginLabel *string `json:"JmeterPluginLabel,omitempty" xml:"JmeterPluginLabel,omitempty"`
	// The JMeter attributes.
	Properties []*SaveEnvRequestEnvProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s SaveEnvRequestEnv) String() string {
	return dara.Prettify(s)
}

func (s SaveEnvRequestEnv) GoString() string {
	return s.String()
}

func (s *SaveEnvRequestEnv) GetEnvId() *string {
	return s.EnvId
}

func (s *SaveEnvRequestEnv) GetEnvName() *string {
	return s.EnvName
}

func (s *SaveEnvRequestEnv) GetFiles() []*SaveEnvRequestEnvFiles {
	return s.Files
}

func (s *SaveEnvRequestEnv) GetJmeterPluginLabel() *string {
	return s.JmeterPluginLabel
}

func (s *SaveEnvRequestEnv) GetProperties() []*SaveEnvRequestEnvProperties {
	return s.Properties
}

func (s *SaveEnvRequestEnv) SetEnvId(v string) *SaveEnvRequestEnv {
	s.EnvId = &v
	return s
}

func (s *SaveEnvRequestEnv) SetEnvName(v string) *SaveEnvRequestEnv {
	s.EnvName = &v
	return s
}

func (s *SaveEnvRequestEnv) SetFiles(v []*SaveEnvRequestEnvFiles) *SaveEnvRequestEnv {
	s.Files = v
	return s
}

func (s *SaveEnvRequestEnv) SetJmeterPluginLabel(v string) *SaveEnvRequestEnv {
	s.JmeterPluginLabel = &v
	return s
}

func (s *SaveEnvRequestEnv) SetProperties(v []*SaveEnvRequestEnvProperties) *SaveEnvRequestEnv {
	s.Properties = v
	return s
}

func (s *SaveEnvRequestEnv) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveEnvRequestEnvFiles struct {
	// The name of the file. Make sure that the file name is the same as the file name in the value of **FileOssAddress**.
	//
	// This parameter is required.
	//
	// example:
	//
	// json.jar
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The Object Storage Service (OSS) URL of the file. Make sure that the file is accessible from the Internet.
	//
	// >  Only OSS URLs in the China (Shanghai) region are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/json.jar
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
}

func (s SaveEnvRequestEnvFiles) String() string {
	return dara.Prettify(s)
}

func (s SaveEnvRequestEnvFiles) GoString() string {
	return s.String()
}

func (s *SaveEnvRequestEnvFiles) GetFileName() *string {
	return s.FileName
}

func (s *SaveEnvRequestEnvFiles) GetFileOssAddress() *string {
	return s.FileOssAddress
}

func (s *SaveEnvRequestEnvFiles) SetFileName(v string) *SaveEnvRequestEnvFiles {
	s.FileName = &v
	return s
}

func (s *SaveEnvRequestEnvFiles) SetFileOssAddress(v string) *SaveEnvRequestEnvFiles {
	s.FileOssAddress = &v
	return s
}

func (s *SaveEnvRequestEnvFiles) Validate() error {
	return dara.Validate(s)
}

type SaveEnvRequestEnvProperties struct {
	// The description of the attribute.
	//
	// example:
	//
	// 远程主机
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the attribute.
	//
	// example:
	//
	// remote_hosts
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// 127.0.0.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SaveEnvRequestEnvProperties) String() string {
	return dara.Prettify(s)
}

func (s SaveEnvRequestEnvProperties) GoString() string {
	return s.String()
}

func (s *SaveEnvRequestEnvProperties) GetDescription() *string {
	return s.Description
}

func (s *SaveEnvRequestEnvProperties) GetName() *string {
	return s.Name
}

func (s *SaveEnvRequestEnvProperties) GetValue() *string {
	return s.Value
}

func (s *SaveEnvRequestEnvProperties) SetDescription(v string) *SaveEnvRequestEnvProperties {
	s.Description = &v
	return s
}

func (s *SaveEnvRequestEnvProperties) SetName(v string) *SaveEnvRequestEnvProperties {
	s.Name = &v
	return s
}

func (s *SaveEnvRequestEnvProperties) SetValue(v string) *SaveEnvRequestEnvProperties {
	s.Value = &v
	return s
}

func (s *SaveEnvRequestEnvProperties) Validate() error {
	return dara.Validate(s)
}
