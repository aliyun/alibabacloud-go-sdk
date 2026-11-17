// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAppInstanceCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *RunAppInstanceCommandRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupIds(v []*string) *RunAppInstanceCommandRequest
	GetAppInstanceGroupIds() []*string
	SetAppInstancePersistentId(v string) *RunAppInstanceCommandRequest
	GetAppInstancePersistentId() *string
	SetAppInstancePersistentIds(v []*string) *RunAppInstanceCommandRequest
	GetAppInstancePersistentIds() []*string
	SetCommandContent(v string) *RunAppInstanceCommandRequest
	GetCommandContent() *string
	SetCommandParam(v string) *RunAppInstanceCommandRequest
	GetCommandParam() *string
	SetCommandType(v string) *RunAppInstanceCommandRequest
	GetCommandType() *string
	SetDescription(v string) *RunAppInstanceCommandRequest
	GetDescription() *string
	SetProductType(v string) *RunAppInstanceCommandRequest
	GetProductType() *string
	SetRunScriptType(v string) *RunAppInstanceCommandRequest
	GetRunScriptType() *string
	SetTimeout(v int64) *RunAppInstanceCommandRequest
	GetTimeout() *int64
}

type RunAppInstanceCommandRequest struct {
	// The ID of the delivery group. This parameter is used to execute commands within a single delivery group. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID. If AppInstanceGroupIds is not specified, you must specify this parameter. Otherwise, a parameter error is returned. If AppInstanceGroupIds is specified, this parameter does not take effect.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The list of delivery group IDs. This parameter is used to issue commands in batches to multiple delivery groups in the same region. The command is executed on all persistent sessions in each delivery group in the list, and only one command execution ID is returned. You can specify up to 100 delivery group IDs. Duplicate IDs are automatically deduplicated. All delivery groups must belong to the same region. Otherwise, the InvalidRegion.NotMatch error is returned. If any delivery group does not exist, the InvalidAppInstanceGroup.NotFound error is returned. If this parameter is specified, AppInstanceGroupId and AppInstancePersistentIds do not take effect.
	AppInstanceGroupIds []*string `json:"AppInstanceGroupIds,omitempty" xml:"AppInstanceGroupIds,omitempty" type:"Repeated"`
	// The ID of the persistent session. This parameter is used to select a single persistent session within the delivery group specified by AppInstanceGroupId to execute the command. The session must belong to the specified delivery group. Otherwise, an error is returned.
	//
	// example:
	//
	// p-0cc7s3mw2fg4j****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The list of persistent session IDs. This parameter is used to select multiple persistent sessions within the delivery group specified by AppInstanceGroupId to execute the command. You can specify up to 100 session IDs. If AppInstanceGroupIds is specified, this parameter does not take effect, and the command is executed on all persistent sessions in each delivery group.
	AppInstancePersistentIds []*string `json:"AppInstancePersistentIds,omitempty" xml:"AppInstancePersistentIds,omitempty" type:"Repeated"`
	// The command content, which is the script content executed in the session. The content must match the script type specified by RunScriptType.
	//
	// This parameter is required.
	//
	// example:
	//
	// echo hello
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command parameters, which are passed through to the session execution end along with the command content. Whether to specify this parameter and the specific format depend on the command content. If not required, you can leave this parameter empty.
	//
	// example:
	//
	// {"name":"value"}
	CommandParam *string `json:"CommandParam,omitempty" xml:"CommandParam,omitempty"`
	// The command type. Currently, only CloudAssistantCmd (Cloud Assistant command) is supported. Specifying other values results in a parameter error.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudAssistantCmd
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The description of the command. This parameter is used to describe the purpose of the issued command for subsequent identification.
	//
	// example:
	//
	// Clean up temporary files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product type. Currently, only WuyingServer (WUYING Workspace) is supported. Specifying other values results in a parameter error.
	//
	// This parameter is required.
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The type of the script to execute. Valid values:
	//
	// - RunBatScript: BAT script.
	//
	// - RunPowerShellScript: PowerShell script.
	//
	// - RunShellScript: Shell script.
	//
	// - RunAIAgentScript: AI agent script.
	//
	// example:
	//
	// RunShellScript
	RunScriptType *string `json:"RunScriptType,omitempty" xml:"RunScriptType,omitempty"`
	// The timeout period for command execution, in seconds. If this parameter is not specified, the default value is 60 seconds.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RunAppInstanceCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RunAppInstanceCommandRequest) GoString() string {
	return s.String()
}

func (s *RunAppInstanceCommandRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *RunAppInstanceCommandRequest) GetAppInstanceGroupIds() []*string {
	return s.AppInstanceGroupIds
}

func (s *RunAppInstanceCommandRequest) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *RunAppInstanceCommandRequest) GetAppInstancePersistentIds() []*string {
	return s.AppInstancePersistentIds
}

func (s *RunAppInstanceCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunAppInstanceCommandRequest) GetCommandParam() *string {
	return s.CommandParam
}

func (s *RunAppInstanceCommandRequest) GetCommandType() *string {
	return s.CommandType
}

func (s *RunAppInstanceCommandRequest) GetDescription() *string {
	return s.Description
}

func (s *RunAppInstanceCommandRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RunAppInstanceCommandRequest) GetRunScriptType() *string {
	return s.RunScriptType
}

func (s *RunAppInstanceCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RunAppInstanceCommandRequest) SetAppInstanceGroupId(v string) *RunAppInstanceCommandRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *RunAppInstanceCommandRequest) SetAppInstanceGroupIds(v []*string) *RunAppInstanceCommandRequest {
	s.AppInstanceGroupIds = v
	return s
}

func (s *RunAppInstanceCommandRequest) SetAppInstancePersistentId(v string) *RunAppInstanceCommandRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *RunAppInstanceCommandRequest) SetAppInstancePersistentIds(v []*string) *RunAppInstanceCommandRequest {
	s.AppInstancePersistentIds = v
	return s
}

func (s *RunAppInstanceCommandRequest) SetCommandContent(v string) *RunAppInstanceCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunAppInstanceCommandRequest) SetCommandParam(v string) *RunAppInstanceCommandRequest {
	s.CommandParam = &v
	return s
}

func (s *RunAppInstanceCommandRequest) SetCommandType(v string) *RunAppInstanceCommandRequest {
	s.CommandType = &v
	return s
}

func (s *RunAppInstanceCommandRequest) SetDescription(v string) *RunAppInstanceCommandRequest {
	s.Description = &v
	return s
}

func (s *RunAppInstanceCommandRequest) SetProductType(v string) *RunAppInstanceCommandRequest {
	s.ProductType = &v
	return s
}

func (s *RunAppInstanceCommandRequest) SetRunScriptType(v string) *RunAppInstanceCommandRequest {
	s.RunScriptType = &v
	return s
}

func (s *RunAppInstanceCommandRequest) SetTimeout(v int64) *RunAppInstanceCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunAppInstanceCommandRequest) Validate() error {
	return dara.Validate(s)
}
