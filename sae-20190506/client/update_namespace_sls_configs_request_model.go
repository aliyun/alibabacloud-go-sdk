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
	// The short ID of the namespace. No need to specify a region ID. We recommend configuring this parameter.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The logging configurations of Simple Log Service.
	//
	// 	- `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`: Simple Log Service resources automatically created by Serverless App Engine (SAE) are used.
	//
	// 	- To use custom Simple Log Service resources, set this parameter to `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Take note of the following subparameters:
	//
	// 	- **projectName**: the name of the Simple Log Service project.
	//
	// 	- **logDir**: the path in which logs are stored.
	//
	// 	- **logType**: the log type. **stdout*	- indicates the standard output (stdout) logs of the container. You can specify only one stdout value for this parameter. If not specified, file logs are collected.
	//
	// 	- **logstoreName**: the name of the Logstore in Simple Log Service.
	//
	// 	- **logtailName**: the name of the Logtail in Simple Log Service. If not specified, a new Logtail is created.
	//
	// If logging configuration changes are not needed during job template deployment, specify **SlsConfigs*	- only in the first request. Omit this parameter in later requests. To stop using Simple Log Service, leave **SlsConfigs*	- empty.
	//
	// > Projects automatically created by SAE for job templates are deleted when their corresponding job templates are deleted. Therefore, you should not select an existing SAE-created project for log collection.
	//
	// example:
	//
	// [{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]
	SlsConfigs    *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
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
