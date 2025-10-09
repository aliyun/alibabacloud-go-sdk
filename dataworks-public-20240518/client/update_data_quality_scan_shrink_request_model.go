// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityScanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResourceShrink(v string) *UpdateDataQualityScanShrinkRequest
	GetComputeResourceShrink() *string
	SetDescription(v string) *UpdateDataQualityScanShrinkRequest
	GetDescription() *string
	SetHooksShrink(v string) *UpdateDataQualityScanShrinkRequest
	GetHooksShrink() *string
	SetId(v int64) *UpdateDataQualityScanShrinkRequest
	GetId() *int64
	SetName(v string) *UpdateDataQualityScanShrinkRequest
	GetName() *string
	SetOwner(v string) *UpdateDataQualityScanShrinkRequest
	GetOwner() *string
	SetParametersShrink(v string) *UpdateDataQualityScanShrinkRequest
	GetParametersShrink() *string
	SetProjectId(v int64) *UpdateDataQualityScanShrinkRequest
	GetProjectId() *int64
	SetRuntimeResourceShrink(v string) *UpdateDataQualityScanShrinkRequest
	GetRuntimeResourceShrink() *string
	SetSpec(v string) *UpdateDataQualityScanShrinkRequest
	GetSpec() *string
	SetTriggerShrink(v string) *UpdateDataQualityScanShrinkRequest
	GetTriggerShrink() *string
}

type UpdateDataQualityScanShrinkRequest struct {
	// The compute engine used during execution. If it\\"s not specified, the data source connection defined in the Spec will be used.
	ComputeResourceShrink *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// Description of the data quality monitor.
	//
	// example:
	//
	// Daily data quality scanning of ods tables.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hook configuration after the data quality monitor stops.
	HooksShrink *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// 231263586109857423
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The definition of execution parameters for the data quality monitor.
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the DataWorks workspace where the data quality monitor resides. You can obtain the workspace ID by calling the [ListProjects](https://help.aliyun.com/document_detail/2852607.html) operation.
	//
	// example:
	//
	// 101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group used during the execution of the data quality monitor.
	RuntimeResourceShrink *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The Spec code of the data quality monitoring content. For more information, see [Data quality Spec configuration description](https://help.aliyun.com/document_detail/2963394.html).
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
	// Trigger settings for the data quality monitor.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s UpdateDataQualityScanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanShrinkRequest) GetComputeResourceShrink() *string {
	return s.ComputeResourceShrink
}

func (s *UpdateDataQualityScanShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataQualityScanShrinkRequest) GetHooksShrink() *string {
	return s.HooksShrink
}

func (s *UpdateDataQualityScanShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityScanShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityScanShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateDataQualityScanShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpdateDataQualityScanShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityScanShrinkRequest) GetRuntimeResourceShrink() *string {
	return s.RuntimeResourceShrink
}

func (s *UpdateDataQualityScanShrinkRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateDataQualityScanShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *UpdateDataQualityScanShrinkRequest) SetComputeResourceShrink(v string) *UpdateDataQualityScanShrinkRequest {
	s.ComputeResourceShrink = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetDescription(v string) *UpdateDataQualityScanShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetHooksShrink(v string) *UpdateDataQualityScanShrinkRequest {
	s.HooksShrink = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetId(v int64) *UpdateDataQualityScanShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetName(v string) *UpdateDataQualityScanShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetOwner(v string) *UpdateDataQualityScanShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetParametersShrink(v string) *UpdateDataQualityScanShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetProjectId(v int64) *UpdateDataQualityScanShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetRuntimeResourceShrink(v string) *UpdateDataQualityScanShrinkRequest {
	s.RuntimeResourceShrink = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetSpec(v string) *UpdateDataQualityScanShrinkRequest {
	s.Spec = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) SetTriggerShrink(v string) *UpdateDataQualityScanShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *UpdateDataQualityScanShrinkRequest) Validate() error {
	return dara.Validate(s)
}
