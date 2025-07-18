// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalSchemasForApprovalProcessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcesses(v []*ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) *ListApprovalSchemasForApprovalProcessesResponseBody
	GetProcesses() []*ListApprovalSchemasForApprovalProcessesResponseBodyProcesses
	SetRequestId(v string) *ListApprovalSchemasForApprovalProcessesResponseBody
	GetRequestId() *string
}

type ListApprovalSchemasForApprovalProcessesResponseBody struct {
	Processes []*ListApprovalSchemasForApprovalProcessesResponseBodyProcesses `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApprovalSchemasForApprovalProcessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasForApprovalProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBody) GetProcesses() []*ListApprovalSchemasForApprovalProcessesResponseBodyProcesses {
	return s.Processes
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBody) SetProcesses(v []*ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) *ListApprovalSchemasForApprovalProcessesResponseBody {
	s.Processes = v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBody) SetRequestId(v string) *ListApprovalSchemasForApprovalProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApprovalSchemasForApprovalProcessesResponseBodyProcesses struct {
	// example:
	//
	// approval-process-35ee09077ee9****
	ProcessId *string                                                                `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	Schemas   []*ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
}

func (s ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) GetSchemas() []*ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	return s.Schemas
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) SetProcessId(v string) *ListApprovalSchemasForApprovalProcessesResponseBodyProcesses {
	s.ProcessId = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) SetSchemas(v []*ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) *ListApprovalSchemasForApprovalProcessesResponseBodyProcesses {
	s.Schemas = v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcesses) Validate() error {
	return dara.Validate(s)
}

type ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas struct {
	// example:
	//
	// 2024-03-15 14:44:07
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// example:
	//
	// {"form": {"labelCol": 6,"wrapperCol": 12}}
	SchemaContent *string `json:"SchemaContent,omitempty" xml:"SchemaContent,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId   *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// example:
	//
	// 1715680940
	SchemaVersion *int64 `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GetDescription() *string {
	return s.Description
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GetSchemaContent() *string {
	return s.SchemaContent
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) GetSchemaVersion() *int64 {
	return s.SchemaVersion
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) SetCreateTime(v string) *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	s.CreateTime = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) SetDescription(v string) *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	s.Description = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) SetIsDefault(v bool) *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	s.IsDefault = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) SetPolicyType(v string) *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	s.PolicyType = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) SetSchemaContent(v string) *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	s.SchemaContent = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) SetSchemaId(v string) *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) SetSchemaName(v string) *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	s.SchemaName = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) SetSchemaVersion(v int64) *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas {
	s.SchemaVersion = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponseBodyProcessesSchemas) Validate() error {
	return dara.Validate(s)
}
