// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalProcessesForApprovalSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListApprovalProcessesForApprovalSchemasResponseBody
	GetRequestId() *string
	SetSchemas(v []*ListApprovalProcessesForApprovalSchemasResponseBodySchemas) *ListApprovalProcessesForApprovalSchemasResponseBody
	GetSchemas() []*ListApprovalProcessesForApprovalSchemasResponseBodySchemas
}

type ListApprovalProcessesForApprovalSchemasResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schemas   []*ListApprovalProcessesForApprovalSchemasResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
}

func (s ListApprovalProcessesForApprovalSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesForApprovalSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBody) GetSchemas() []*ListApprovalProcessesForApprovalSchemasResponseBodySchemas {
	return s.Schemas
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBody) SetRequestId(v string) *ListApprovalProcessesForApprovalSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBody) SetSchemas(v []*ListApprovalProcessesForApprovalSchemasResponseBodySchemas) *ListApprovalProcessesForApprovalSchemasResponseBody {
	s.Schemas = v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesForApprovalSchemasResponseBodySchemas struct {
	Processes []*ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesForApprovalSchemasResponseBodySchemas) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesForApprovalSchemasResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemas) GetProcesses() []*ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses {
	return s.Processes
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemas) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemas) SetProcesses(v []*ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) *ListApprovalProcessesForApprovalSchemasResponseBodySchemas {
	s.Processes = v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemas) SetSchemaId(v string) *ListApprovalProcessesForApprovalSchemasResponseBodySchemas {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemas) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses struct {
	// example:
	//
	// 2022-10-25 10:44:09
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// approval-process-dc61e92ba5c5****
	ProcessId   *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
}

func (s ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) GetDescription() *string {
	return s.Description
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) SetCreateTime(v string) *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses {
	s.CreateTime = &v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) SetDescription(v string) *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses {
	s.Description = &v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) SetProcessId(v string) *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses {
	s.ProcessId = &v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) SetProcessName(v string) *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses {
	s.ProcessName = &v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponseBodySchemasProcesses) Validate() error {
	return dara.Validate(s)
}
