// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityScanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDataQualityScanShrinkRequest
	GetClientToken() *string
	SetComputeResourceShrink(v string) *CreateDataQualityScanShrinkRequest
	GetComputeResourceShrink() *string
	SetDescription(v string) *CreateDataQualityScanShrinkRequest
	GetDescription() *string
	SetHooksShrink(v string) *CreateDataQualityScanShrinkRequest
	GetHooksShrink() *string
	SetName(v string) *CreateDataQualityScanShrinkRequest
	GetName() *string
	SetOwner(v string) *CreateDataQualityScanShrinkRequest
	GetOwner() *string
	SetParametersShrink(v string) *CreateDataQualityScanShrinkRequest
	GetParametersShrink() *string
	SetProjectId(v int64) *CreateDataQualityScanShrinkRequest
	GetProjectId() *int64
	SetRuntimeResourceShrink(v string) *CreateDataQualityScanShrinkRequest
	GetRuntimeResourceShrink() *string
	SetSpec(v string) *CreateDataQualityScanShrinkRequest
	GetSpec() *string
	SetTriggerShrink(v string) *CreateDataQualityScanShrinkRequest
	GetTriggerShrink() *string
}

type CreateDataQualityScanShrinkRequest struct {
	// The idempotency token.
	//
	// This parameter is required.
	//
	// example:
	//
	// a-customized-uuid
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The compute engine used when the data quality monitor runs. This parameter is optional. If it is not specified, the data source connection specified in the Spec is used.
	ComputeResourceShrink *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The description of the data quality monitor.
	//
	// example:
	//
	// Daily data quality scanning of ods tables.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configuration for hooks that run after the data quality monitor finishes.
	HooksShrink *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// The name of the data quality monitor.
	//
	// example:
	//
	// data_quality_scan_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the owner of the data quality monitor.
	//
	// example:
	//
	// 95279527****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The execution parameter definitions of the data quality monitor.
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/overview) and go to the Workspace Management page to obtain the workspace ID. This parameter specifies the DataWorks workspace for the API call.
	//
	// example:
	//
	// 101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group used when the data quality monitor runs.
	RuntimeResourceShrink *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The Spec code that defines the data quality monitor. For more information, see [Data quality spec configuration guide](https://help.aliyun.com/document_detail/2963394.html).
	//
	// example:
	//
	// {
	//
	//     "datasets": [
	//
	//         {
	//
	//             "type": "Table",
	//
	//             "dataSource": {
	//
	//                 "name": "odps_first",
	//
	//                 "envType": "Prod"
	//
	//             },
	//
	//             "tables": [
	//
	//                 "ods_d_user_info"
	//
	//             ],
	//
	//             "filter": "pt = $[yyyymmdd-1]"
	//
	//         }
	//
	//     ],
	//
	//     "rules": [
	//
	//         {
	//
	//             "assertion": "row_count > 0"
	//
	//         }, {
	//
	//             "templateId": "SYSTEM:field:null_value:fixed",
	//
	//             "pass": "when = 0",
	//
	//             "name": "The id cannot be empty.",
	//
	//             "severity": "High",
	//
	//              "identity": "a-customized-data-quality-rule-uuid"
	//
	//         }
	//
	//     ]
	//
	// }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The trigger settings of the data quality monitor.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s CreateDataQualityScanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDataQualityScanShrinkRequest) GetComputeResourceShrink() *string {
	return s.ComputeResourceShrink
}

func (s *CreateDataQualityScanShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataQualityScanShrinkRequest) GetHooksShrink() *string {
	return s.HooksShrink
}

func (s *CreateDataQualityScanShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityScanShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateDataQualityScanShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *CreateDataQualityScanShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityScanShrinkRequest) GetRuntimeResourceShrink() *string {
	return s.RuntimeResourceShrink
}

func (s *CreateDataQualityScanShrinkRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateDataQualityScanShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *CreateDataQualityScanShrinkRequest) SetClientToken(v string) *CreateDataQualityScanShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetComputeResourceShrink(v string) *CreateDataQualityScanShrinkRequest {
	s.ComputeResourceShrink = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetDescription(v string) *CreateDataQualityScanShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetHooksShrink(v string) *CreateDataQualityScanShrinkRequest {
	s.HooksShrink = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetName(v string) *CreateDataQualityScanShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetOwner(v string) *CreateDataQualityScanShrinkRequest {
	s.Owner = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetParametersShrink(v string) *CreateDataQualityScanShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetProjectId(v int64) *CreateDataQualityScanShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetRuntimeResourceShrink(v string) *CreateDataQualityScanShrinkRequest {
	s.RuntimeResourceShrink = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetSpec(v string) *CreateDataQualityScanShrinkRequest {
	s.Spec = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) SetTriggerShrink(v string) *CreateDataQualityScanShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *CreateDataQualityScanShrinkRequest) Validate() error {
	return dara.Validate(s)
}
