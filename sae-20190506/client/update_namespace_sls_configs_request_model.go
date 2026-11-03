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
	// The short ID of the namespace. The region is not required. This parameter is recommended.
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
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - To use SLS resources automatically created by Serverless App Engine (SAE): `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The parameter settings are as follows:
	//
	// - **projectName**: the name of the SLS project.
	//
	// - **logDir**: the log path.
	//
	// - **logType**: the log type. **stdout*	- indicates container standard output logs. Only one stdout log can be set. If it is not set, file logs are collected.
	//
	// - **logstoreName**: the name of the SLS Logstore.
	//
	// - **logtailName**: the name of the SLS Logtail. If this parameter is not specified, a new Logtail is created.
	//
	// If the SLS collection configuration remains unchanged during multiple deployments, you do not need to set this parameter (that is, the **SlsConfigs*	- field is not required in the request). If you no longer need the SLS collection feature, you must set the value of this field to an empty string (that is, set the **SlsConfigs*	- field to "") in the request.
	//
	// > A project automatically created with a job template is deleted when the job template is deleted. Therefore, when you select an existing project, you cannot select a project automatically created by SAE.
	//
	// example:
	//
	// [{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// sls log tags
	//
	// example:
	//
	// []
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
