// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectorKibanaInstance interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *CollectorKibanaInstance
	GetConfigType() *string
	SetHost(v string) *CollectorKibanaInstance
	GetHost() *string
	SetInstanceId(v string) *CollectorKibanaInstance
	GetInstanceId() *string
	SetKibanaHost(v string) *CollectorKibanaInstance
	GetKibanaHost() *string
	SetPassword(v string) *CollectorKibanaInstance
	GetPassword() *string
	SetProtocol(v string) *CollectorKibanaInstance
	GetProtocol() *string
	SetUserName(v string) *CollectorKibanaInstance
	GetUserName() *string
}

type CollectorKibanaInstance struct {
	// This parameter is required.
	//
	// example:
	//
	// collectorElasticsearchForKibana
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// es-cn-*****-kibana.internal.elasticsearch.aliyuncs.com:5601
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// es-cn-45dfy****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://es-cn-****.kibana.elasticsearch.aliyuncs.com:5601
	KibanaHost *string `json:"kibanaHost,omitempty" xml:"kibanaHost,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// username
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s CollectorKibanaInstance) String() string {
	return dara.Prettify(s)
}

func (s CollectorKibanaInstance) GoString() string {
	return s.String()
}

func (s *CollectorKibanaInstance) GetConfigType() *string {
	return s.ConfigType
}

func (s *CollectorKibanaInstance) GetHost() *string {
	return s.Host
}

func (s *CollectorKibanaInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CollectorKibanaInstance) GetKibanaHost() *string {
	return s.KibanaHost
}

func (s *CollectorKibanaInstance) GetPassword() *string {
	return s.Password
}

func (s *CollectorKibanaInstance) GetProtocol() *string {
	return s.Protocol
}

func (s *CollectorKibanaInstance) GetUserName() *string {
	return s.UserName
}

func (s *CollectorKibanaInstance) SetConfigType(v string) *CollectorKibanaInstance {
	s.ConfigType = &v
	return s
}

func (s *CollectorKibanaInstance) SetHost(v string) *CollectorKibanaInstance {
	s.Host = &v
	return s
}

func (s *CollectorKibanaInstance) SetInstanceId(v string) *CollectorKibanaInstance {
	s.InstanceId = &v
	return s
}

func (s *CollectorKibanaInstance) SetKibanaHost(v string) *CollectorKibanaInstance {
	s.KibanaHost = &v
	return s
}

func (s *CollectorKibanaInstance) SetPassword(v string) *CollectorKibanaInstance {
	s.Password = &v
	return s
}

func (s *CollectorKibanaInstance) SetProtocol(v string) *CollectorKibanaInstance {
	s.Protocol = &v
	return s
}

func (s *CollectorKibanaInstance) SetUserName(v string) *CollectorKibanaInstance {
	s.UserName = &v
	return s
}

func (s *CollectorKibanaInstance) Validate() error {
	return dara.Validate(s)
}
