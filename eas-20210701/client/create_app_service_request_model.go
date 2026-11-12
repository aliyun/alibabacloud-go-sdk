// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateAppServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateAppServiceRequest
	GetDryRun() *bool
	SetQuotaId(v string) *CreateAppServiceRequest
	GetQuotaId() *string
	SetWorkspaceId(v string) *CreateAppServiceRequest
	GetWorkspaceId() *string
	SetAppType(v string) *CreateAppServiceRequest
	GetAppType() *string
	SetAppVersion(v string) *CreateAppServiceRequest
	GetAppVersion() *string
	SetConfig(v map[string]interface{}) *CreateAppServiceRequest
	GetConfig() map[string]interface{}
	SetReplicas(v int32) *CreateAppServiceRequest
	GetReplicas() *int32
	SetServiceName(v string) *CreateAppServiceRequest
	GetServiceName() *string
	SetServiceSpec(v string) *CreateAppServiceRequest
	GetServiceSpec() *string
}

type CreateAppServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 5db780a0-9d3e-4f35-a84a-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: performs a dry run and does not create the instance. The system checks whether the required parameters are specified, the request format is valid, the service limits are met, and the ECS inventory is sufficient. If the request fails the check, an error is returned. If the request passes the check, the DryRunOperation error code is returned.
	//
	// - false (default): performs a normal request and creates the instance if the request passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The type of the application service.
	//
	// This parameter is required.
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
	// The additional configuration information required for deployment.
	//
	// example:
	//
	// { "ModelStorage": "oss", "OssPath": "oss://examplebucket/xxxx/" }
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The number of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service specifications. Valid values:
	//
	// - llama_7b_fp16
	//
	// - llama_7b_int8
	//
	// - llama_13b_fp16
	//
	// - llama_7b_int8
	//
	// - chatglm_6b_fp16
	//
	// - chatglm_6b_int8
	//
	// - chatglm2_6b_fp16
	//
	// - baichuan_7b_int8
	//
	// - baichuan_13b_fp16
	//
	// - baichuan_7b_fp16
	//
	// This parameter is required.
	//
	// example:
	//
	// llama_7b_fp16
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
}

func (s CreateAppServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateAppServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAppServiceRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *CreateAppServiceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAppServiceRequest) GetAppType() *string {
	return s.AppType
}

func (s *CreateAppServiceRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *CreateAppServiceRequest) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *CreateAppServiceRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateAppServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateAppServiceRequest) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *CreateAppServiceRequest) SetClientToken(v string) *CreateAppServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppServiceRequest) SetDryRun(v bool) *CreateAppServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAppServiceRequest) SetQuotaId(v string) *CreateAppServiceRequest {
	s.QuotaId = &v
	return s
}

func (s *CreateAppServiceRequest) SetWorkspaceId(v string) *CreateAppServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAppServiceRequest) SetAppType(v string) *CreateAppServiceRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppServiceRequest) SetAppVersion(v string) *CreateAppServiceRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppServiceRequest) SetConfig(v map[string]interface{}) *CreateAppServiceRequest {
	s.Config = v
	return s
}

func (s *CreateAppServiceRequest) SetReplicas(v int32) *CreateAppServiceRequest {
	s.Replicas = &v
	return s
}

func (s *CreateAppServiceRequest) SetServiceName(v string) *CreateAppServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateAppServiceRequest) SetServiceSpec(v string) *CreateAppServiceRequest {
	s.ServiceSpec = &v
	return s
}

func (s *CreateAppServiceRequest) Validate() error {
	return dara.Validate(s)
}
