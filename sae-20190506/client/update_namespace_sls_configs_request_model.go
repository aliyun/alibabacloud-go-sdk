// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceSlsConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNameSpaceShortId(v string) *UpdateNamespaceSlsConfigsRequest
	GetNameSpaceShortId() *string
	SetNamespaceId(v string) *UpdateNamespaceSlsConfigsRequest
	GetNamespaceId() *string
	SetSlsConfigs(v string) *UpdateNamespaceSlsConfigsRequest
	GetSlsConfigs() *string
	SetSlsLogEnvTags(v string) *UpdateNamespaceSlsConfigsRequest
	GetSlsLogEnvTags() *string
}

type UpdateNamespaceSlsConfigsRequest struct {
	// The short ID of the namespace. You do not need to include the region. This parameter is recommended.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configuration for collecting logs to SLS.
	//
	// - To use an SLS resource that is automatically created by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use a custom SLS resource: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The parameters are described as follows:
	//
	// - `projectName`: The name of the SLS project.
	//
	// - `logDir`: The log path.
	//
	// - `logType`: The log type. A value of `stdout` specifies container standard output logs. You can specify only one `stdout` configuration. If you do not set this parameter, file logs are collected.
	//
	// - `logstoreName`: The name of the SLS logstore.
	//
	// - `logtailName`: The name of the Logtail. If you do not specify this parameter, a new Logtail is created.
	//
	// If the SLS configuration remains the same across deployments, you can omit this parameter. To disable log collection to SLS, set the value of `SlsConfigs` to an empty string ("").
	//
	// > SAE automatically deletes a project when you delete the task template used to create it. Therefore, when you select an existing project, do not select a project that was automatically created by SAE.
	//
	// example:
	//
	// [{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The SLS log tags.
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
}

func (s UpdateNamespaceSlsConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceSlsConfigsRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceSlsConfigsRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *UpdateNamespaceSlsConfigsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNamespaceSlsConfigsRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *UpdateNamespaceSlsConfigsRequest) GetSlsLogEnvTags() *string {
	return s.SlsLogEnvTags
}

func (s *UpdateNamespaceSlsConfigsRequest) SetNameSpaceShortId(v string) *UpdateNamespaceSlsConfigsRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsRequest) SetNamespaceId(v string) *UpdateNamespaceSlsConfigsRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsRequest) SetSlsConfigs(v string) *UpdateNamespaceSlsConfigsRequest {
	s.SlsConfigs = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsRequest) SetSlsLogEnvTags(v string) *UpdateNamespaceSlsConfigsRequest {
	s.SlsLogEnvTags = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsRequest) Validate() error {
	return dara.Validate(s)
}
