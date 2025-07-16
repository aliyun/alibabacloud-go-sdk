// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaId(v string) *UpdateAppServiceRequest
	GetQuotaId() *string
	SetWorkspaceId(v string) *UpdateAppServiceRequest
	GetWorkspaceId() *string
	SetAppType(v string) *UpdateAppServiceRequest
	GetAppType() *string
	SetAppVersion(v string) *UpdateAppServiceRequest
	GetAppVersion() *string
	SetConfig(v map[string]interface{}) *UpdateAppServiceRequest
	GetConfig() map[string]interface{}
	SetReplicas(v int32) *UpdateAppServiceRequest
	GetReplicas() *int32
	SetServiceSpec(v string) *UpdateAppServiceRequest
	GetServiceSpec() *string
}

type UpdateAppServiceRequest struct {
	// The quota ID.
	//
	// example:
	//
	// abcdef
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The application type.
	//
	// Valid values:
	//
	// 	- LLM: the large language model (LLM) application
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// LLM
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The application version.
	//
	// example:
	//
	// v1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The additional configurations that are required for service deployment.
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The number of instances. This value must be greater than 0.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The service specifications. Valid values:
	//
	// 	- llama_7b_fp16
	//
	// 	- llama_7b_int8
	//
	// 	- llama_13b_fp16
	//
	// 	- llama_7b_int8
	//
	// 	- chatglm_6b_fp16
	//
	// 	- chatglm_6b_int8
	//
	// 	- chatglm2_6b_fp16
	//
	// 	- baichuan_7b_int8
	//
	// 	- baichuan_13b_fp16
	//
	// 	- baichuan_7b_fp16
	//
	// example:
	//
	// llama_7b_fp16
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
}

func (s UpdateAppServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppServiceRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *UpdateAppServiceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateAppServiceRequest) GetAppType() *string {
	return s.AppType
}

func (s *UpdateAppServiceRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *UpdateAppServiceRequest) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *UpdateAppServiceRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *UpdateAppServiceRequest) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *UpdateAppServiceRequest) SetQuotaId(v string) *UpdateAppServiceRequest {
	s.QuotaId = &v
	return s
}

func (s *UpdateAppServiceRequest) SetWorkspaceId(v string) *UpdateAppServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateAppServiceRequest) SetAppType(v string) *UpdateAppServiceRequest {
	s.AppType = &v
	return s
}

func (s *UpdateAppServiceRequest) SetAppVersion(v string) *UpdateAppServiceRequest {
	s.AppVersion = &v
	return s
}

func (s *UpdateAppServiceRequest) SetConfig(v map[string]interface{}) *UpdateAppServiceRequest {
	s.Config = v
	return s
}

func (s *UpdateAppServiceRequest) SetReplicas(v int32) *UpdateAppServiceRequest {
	s.Replicas = &v
	return s
}

func (s *UpdateAppServiceRequest) SetServiceSpec(v string) *UpdateAppServiceRequest {
	s.ServiceSpec = &v
	return s
}

func (s *UpdateAppServiceRequest) Validate() error {
	return dara.Validate(s)
}
