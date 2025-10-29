// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIDEEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventDetail(v *GetIDEEventDetailResponseBodyEventDetail) *GetIDEEventDetailResponseBody
	GetEventDetail() *GetIDEEventDetailResponseBodyEventDetail
	SetRequestId(v string) *GetIDEEventDetailResponseBody
	GetRequestId() *string
}

type GetIDEEventDetailResponseBody struct {
	// The data snapshot that is generated when an extension point event is triggered.
	//
	// The fields contained in data snapshots vary based on the types of the triggered extension point events. For more information, see the description of the fields.
	EventDetail *GetIDEEventDetailResponseBodyEventDetail `json:"EventDetail,omitempty" xml:"EventDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIDEEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBody) GetEventDetail() *GetIDEEventDetailResponseBodyEventDetail {
	return s.EventDetail
}

func (s *GetIDEEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIDEEventDetailResponseBody) SetEventDetail(v *GetIDEEventDetailResponseBodyEventDetail) *GetIDEEventDetailResponseBody {
	s.EventDetail = v
	return s
}

func (s *GetIDEEventDetailResponseBody) SetRequestId(v string) *GetIDEEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIDEEventDetailResponseBody) Validate() error {
	if s.EventDetail != nil {
		if err := s.EventDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIDEEventDetailResponseBodyEventDetail struct {
	// The data snapshot when the file is committed and deployed.
	//
	// This parameter is valid only if the message type is IDE_FILE_SUBMIT_BEFORE or IDE_FILE_DEPLOY_BEFORE.
	CommittedFile *GetIDEEventDetailResponseBodyEventDetailCommittedFile `json:"CommittedFile,omitempty" xml:"CommittedFile,omitempty" type:"Struct"`
	// The data snapshot when the file is deleted. This parameter is valid only if the message type is IDE_FILE_DELETE_BEFORE.
	DeletedFile *GetIDEEventDetailResponseBodyEventDetailDeletedFile `json:"DeletedFile,omitempty" xml:"DeletedFile,omitempty" type:"Struct"`
	// The data snapshot when the code in the file is run. This parameter is valid only if the message type is IDE_FILE_EXECUTE_BEFORE.
	FileExecutionCommand *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand `json:"FileExecutionCommand,omitempty" xml:"FileExecutionCommand,omitempty" type:"Struct"`
	// The data snapshot when the table is committed and deployed. This parameter is valid only if the message type is IDE_TABLE_SUBMIT_BEFORE or IDE_TABLE_DEPLOY_BEFORE.
	TableModel *GetIDEEventDetailResponseBodyEventDetailTableModel `json:"TableModel,omitempty" xml:"TableModel,omitempty" type:"Struct"`
}

func (s GetIDEEventDetailResponseBodyEventDetail) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetail) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetail) GetCommittedFile() *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	return s.CommittedFile
}

func (s *GetIDEEventDetailResponseBodyEventDetail) GetDeletedFile() *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	return s.DeletedFile
}

func (s *GetIDEEventDetailResponseBodyEventDetail) GetFileExecutionCommand() *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand {
	return s.FileExecutionCommand
}

func (s *GetIDEEventDetailResponseBodyEventDetail) GetTableModel() *GetIDEEventDetailResponseBodyEventDetailTableModel {
	return s.TableModel
}

func (s *GetIDEEventDetailResponseBodyEventDetail) SetCommittedFile(v *GetIDEEventDetailResponseBodyEventDetailCommittedFile) *GetIDEEventDetailResponseBodyEventDetail {
	s.CommittedFile = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetail) SetDeletedFile(v *GetIDEEventDetailResponseBodyEventDetailDeletedFile) *GetIDEEventDetailResponseBodyEventDetail {
	s.DeletedFile = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetail) SetFileExecutionCommand(v *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) *GetIDEEventDetailResponseBodyEventDetail {
	s.FileExecutionCommand = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetail) SetTableModel(v *GetIDEEventDetailResponseBodyEventDetailTableModel) *GetIDEEventDetailResponseBodyEventDetail {
	s.TableModel = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetail) Validate() error {
	if s.CommittedFile != nil {
		if err := s.CommittedFile.Validate(); err != nil {
			return err
		}
	}
	if s.DeletedFile != nil {
		if err := s.DeletedFile.Validate(); err != nil {
			return err
		}
	}
	if s.FileExecutionCommand != nil {
		if err := s.FileExecutionCommand.Validate(); err != nil {
			return err
		}
	}
	if s.TableModel != nil {
		if err := s.TableModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIDEEventDetailResponseBodyEventDetailCommittedFile struct {
	// The type of the change to the file of the current version. Valid values: CREATE, UPDATE, and DELETE.
	//
	// example:
	//
	// UPDATE
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The description of the file version.
	//
	// example:
	//
	// Second version submission
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the Alibaba Cloud account that is used to generate the file of the current version.
	//
	// example:
	//
	// 7384234****
	Committor *string `json:"Committor,omitempty" xml:"Committor,omitempty"`
	// The code in the file of the current version.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 1234123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// hello_dataworks.sql
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The details of the file.
	FilePropertyContent *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent `json:"FilePropertyContent,omitempty" xml:"FilePropertyContent,omitempty" type:"Struct"`
	// The file type. The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 10
	FileType *int64 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The scheduling properties of the node that corresponds to the file.
	NodeConfiguration *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration `json:"NodeConfiguration,omitempty" xml:"NodeConfiguration,omitempty" type:"Struct"`
	// The ID of the node that is scheduled.
	//
	// example:
	//
	// 421429
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The module to which the file belongs. Valid values:
	//
	// 	- NORMAL: The file is used for DataStudio.
	//
	// 	- MANUAL: The file is used for a manually triggered node.
	//
	// 	- MANUAL_BIZ: The file is used for a manually triggered workflow.
	//
	// 	- SKIP: The file is used for a dry-run node in DataStudio.
	//
	// 	- ADHOCQUERY: The file is used for an ad hoc query.
	//
	// 	- COMPONENT: The file is used for a script template.
	//
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFile) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFile) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetChangeType() *string {
	return s.ChangeType
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetComment() *string {
	return s.Comment
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetCommittor() *string {
	return s.Committor
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetContent() *string {
	return s.Content
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetFileId() *int64 {
	return s.FileId
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetFileName() *string {
	return s.FileName
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetFilePropertyContent() *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent {
	return s.FilePropertyContent
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetFileType() *int64 {
	return s.FileType
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetNodeConfiguration() *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	return s.NodeConfiguration
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) GetUseType() *string {
	return s.UseType
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetChangeType(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.ChangeType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetComment(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.Comment = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetCommittor(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.Committor = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetContent(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.Content = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetFileId(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.FileId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetFileName(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.FileName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetFilePropertyContent(v *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.FilePropertyContent = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetFileType(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.FileType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetNodeConfiguration(v *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.NodeConfiguration = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetNodeId(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.NodeId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) SetUseType(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFile {
	s.UseType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFile) Validate() error {
	if s.FilePropertyContent != nil {
		if err := s.FilePropertyContent.Validate(); err != nil {
			return err
		}
	}
	if s.NodeConfiguration != nil {
		if err := s.NodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent struct {
	// The ID of the workflow to which the file belongs.
	//
	// example:
	//
	// 74328
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The latest version number of the file.
	//
	// example:
	//
	// 1
	CurrentVersion *int64 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The name of the data source with which the file is associated.
	//
	// example:
	//
	// odps_source
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The ID of the folder to which the file belongs. You can call the [GetFolder](https://help.aliyun.com/document_detail/173952.html) operation to query the details of the file based on the folder ID.
	//
	// example:
	//
	// aldurie78l2falure
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The file owner.
	//
	// example:
	//
	// 7384234****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the do-while node or for-each node that corresponds to the file.
	//
	// example:
	//
	// 1234122
	ParentFileId *int64 `json:"ParentFileId,omitempty" xml:"ParentFileId,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) GetCurrentVersion() *int64 {
	return s.CurrentVersion
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) GetFolderId() *string {
	return s.FolderId
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) GetOwner() *string {
	return s.Owner
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) GetParentFileId() *int64 {
	return s.ParentFileId
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) SetBusinessId(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent {
	s.BusinessId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) SetCurrentVersion(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent {
	s.CurrentVersion = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) SetDataSourceName(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent {
	s.DataSourceName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) SetFolderId(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent {
	s.FolderId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) SetOwner(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent {
	s.Owner = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) SetParentFileId(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent {
	s.ParentFileId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileFilePropertyContent) Validate() error {
	return dara.Validate(s)
}

type GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration struct {
	// The interval at which the node corresponding to the file is rerun. Unit: milliseconds.
	//
	// example:
	//
	// 120000
	AutoRerunIntervalMillis *int64 `json:"AutoRerunIntervalMillis,omitempty" xml:"AutoRerunIntervalMillis,omitempty"`
	// The number of times that the node corresponding to the file can be rerun.
	//
	// example:
	//
	// 3
	AutoRerunTimes *int64 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// The cron expression that is used to schedule the node corresponding to the file.
	//
	// example:
	//
	// 00 05 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of the scheduling cycle of the node that corresponds to the file. Valid values: NOT_DAY and DAY. The value NOT_DAY indicates that the node is scheduled to run by minute or hour. The value DAY indicates that the node is scheduled to run by day, week, or month.
	//
	// This parameter corresponds to the Scheduling Cycle parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The ID of the node on which the node that corresponds to the file depends when the DependentType parameter is set to USER_DEFINE. Multiple IDs are separated by commas (,).
	//
	// The value of this parameter is equivalent to the ID of the node that you specified after you select Other Nodes for Cross-Cycle Dependency (Original Previous-Cycle Dependency) in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 5,10,15,20
	DependentNodeIdList *string `json:"DependentNodeIdList,omitempty" xml:"DependentNodeIdList,omitempty"`
	// The type of the cross-cycle scheduling dependency of the node. Valid values:
	//
	// 	- SELF: The instance generated for the node in the current cycle depends on the instance generated for the node in the previous cycle.
	//
	// 	- CHILD: The instance generated for the node in the current cycle depends on the instances generated for the descendant nodes at the nearest level of the node in the previous cycle.
	//
	// 	- USER_DEFINE: The instance generated for the node in the current cycle depends on the instances generated for one or more specified nodes in the previous cycle.
	//
	// 	- NONE: No cross-cycle scheduling dependency type is selected for the node.
	//
	// example:
	//
	// USER_DEFINE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// The output information about the parent files on which the current file depends.
	InputList []*GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList `json:"InputList,omitempty" xml:"InputList,omitempty" type:"Repeated"`
	// The output information about the current file.
	OutputList []*GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList `json:"OutputList,omitempty" xml:"OutputList,omitempty" type:"Repeated"`
	// The scheduling parameters of the node.
	//
	// This parameter corresponds to the Scheduling Parameter section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console). For more information about the configurations of scheduling parameters, see [Configure scheduling parameters](https://help.aliyun.com/document_detail/137548.html).
	//
	// example:
	//
	// a=x b=y
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// Indicates whether the node that corresponds to the file can be rerun. Valid values:
	//
	// 	- ALL_ALLOWED: The node can be rerun regardless of whether it is successfully run or fails to run.
	//
	// 	- FAILURE_ALLOWED: The node can be rerun only after it fails to run.
	//
	// 	- ALL_DENIED: The node cannot be rerun regardless of whether it is successfully run or fails to run.
	//
	// This parameter corresponds to the Rerun parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// ALL_ALLOWED
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The ID of the resource group that is used to run the node that corresponds to the file. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the available resource groups in the workspace.
	//
	// example:
	//
	// 375827434852437
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling type of the node. Valid values:
	//
	// 	- NORMAL: The node is an auto triggered node.
	//
	// 	- MANUAL: The node is a manually triggered node. Manually triggered nodes cannot be automatically triggered. They correspond to the nodes in the Manually Triggered Workflows pane.
	//
	// 	- PAUSE: The node is a paused node.
	//
	// 	- SKIP: The node is a dry-run node. Dry-run nodes are started as scheduled, but the system sets the status of the nodes to successful when it starts to run them.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetAutoRerunIntervalMillis() *int64 {
	return s.AutoRerunIntervalMillis
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetAutoRerunTimes() *int64 {
	return s.AutoRerunTimes
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetCronExpress() *string {
	return s.CronExpress
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetCycleType() *string {
	return s.CycleType
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetDependentNodeIdList() *string {
	return s.DependentNodeIdList
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetDependentType() *string {
	return s.DependentType
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetInputList() []*GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList {
	return s.InputList
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetOutputList() []*GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList {
	return s.OutputList
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetParaValue() *string {
	return s.ParaValue
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetRerunMode() *string {
	return s.RerunMode
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetResourceGroupId() *int64 {
	return s.ResourceGroupId
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetAutoRerunIntervalMillis(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.AutoRerunIntervalMillis = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetAutoRerunTimes(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.AutoRerunTimes = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetCronExpress(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.CronExpress = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetCycleType(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.CycleType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetDependentNodeIdList(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.DependentNodeIdList = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetDependentType(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.DependentType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetInputList(v []*GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.InputList = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetOutputList(v []*GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.OutputList = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetParaValue(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.ParaValue = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetRerunMode(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.RerunMode = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetResourceGroupId(v int64) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) SetSchedulerType(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration {
	s.SchedulerType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfiguration) Validate() error {
	if s.InputList != nil {
		for _, item := range s.InputList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputList != nil {
		for _, item := range s.OutputList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList struct {
	// The output name of the parent file on which the current file depends.
	//
	// This parameter corresponds to the Output Name of Ancestor Node parameter under Parent Nodes in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// dw_project_root
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The mode of the configuration file dependency. Valid values:
	//
	// 	- MANUAL: Scheduling dependencies are manually configured.
	//
	// 	- AUTO: Scheduling dependencies are automatically parsed.
	//
	// example:
	//
	// MANUAL
	ParseType *string `json:"ParseType,omitempty" xml:"ParseType,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList) GetInput() *string {
	return s.Input
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList) GetParseType() *string {
	return s.ParseType
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList) SetInput(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList {
	s.Input = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList) SetParseType(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList {
	s.ParseType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationInputList) Validate() error {
	return dara.Validate(s)
}

type GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList struct {
	// The output name of the current file.
	//
	// This parameter corresponds to the Output Name parameter under Output Name of Current Node in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// dw_project.002_out
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The output table name of the current file.
	//
	// This parameter corresponds to the Output Table Name parameter under Output Name of Current Node in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// ods_user_info_d
	RefTableName *string `json:"RefTableName,omitempty" xml:"RefTableName,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList) GetOutput() *string {
	return s.Output
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList) GetRefTableName() *string {
	return s.RefTableName
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList) SetOutput(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList {
	s.Output = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList) SetRefTableName(v string) *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList {
	s.RefTableName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailCommittedFileNodeConfigurationOutputList) Validate() error {
	return dara.Validate(s)
}

type GetIDEEventDetailResponseBodyEventDetailDeletedFile struct {
	// The ID of the workflow to which the file belongs.
	//
	// example:
	//
	// 74328
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The code in the file of the current version.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The latest version number of the file.
	//
	// example:
	//
	// 1
	CurrentVersion *int64 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The name of the data source with which the file is associated.
	//
	// example:
	//
	// odps_source
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 1234123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// hello_dataworks.sql
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type. The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 10
	FileType *int64 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The ID of the folder to which the file belongs. You can call the [GetFolder](https://help.aliyun.com/document_detail/173952.html) operation to query the details of the file based on the folder ID.
	//
	// example:
	//
	// aldurie78l2falure
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the node that is scheduled.
	//
	// example:
	//
	// 421429
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The file owner.
	//
	// example:
	//
	// 7384234****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the do-while node or for-each node that corresponds to the file.
	//
	// example:
	//
	// 1234122
	ParentFileId *int64 `json:"ParentFileId,omitempty" xml:"ParentFileId,omitempty"`
	// The module to which the file belongs. Valid values:
	//
	// 	- NORMAL: The file is used for DataStudio.
	//
	// 	- MANUAL: The file is used for a manually triggered node.
	//
	// 	- MANUAL_BIZ: The file is used for a manually triggered workflow.
	//
	// 	- SKIP: The file is used for a dry-run node in DataStudio.
	//
	// 	- ADHOCQUERY: The file is used for an ad hoc query.
	//
	// 	- COMPONENT: The file is used for a script template.
	//
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailDeletedFile) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailDeletedFile) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetContent() *string {
	return s.Content
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetCurrentVersion() *int64 {
	return s.CurrentVersion
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetFileId() *int64 {
	return s.FileId
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetFileName() *string {
	return s.FileName
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetFileType() *int64 {
	return s.FileType
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetFolderId() *string {
	return s.FolderId
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetOwner() *string {
	return s.Owner
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetParentFileId() *int64 {
	return s.ParentFileId
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) GetUseType() *string {
	return s.UseType
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetBusinessId(v int64) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.BusinessId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetContent(v string) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.Content = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetCurrentVersion(v int64) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.CurrentVersion = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetDataSourceName(v string) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.DataSourceName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetFileId(v int64) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.FileId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetFileName(v string) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.FileName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetFileType(v int64) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.FileType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetFolderId(v string) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.FolderId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetNodeId(v int64) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.NodeId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetOwner(v string) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.Owner = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetParentFileId(v int64) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.ParentFileId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) SetUseType(v string) *GetIDEEventDetailResponseBodyEventDetailDeletedFile {
	s.UseType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailDeletedFile) Validate() error {
	return dara.Validate(s)
}

type GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand struct {
	// The code in the file of the current version.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The name of the data source with which the file is associated.
	//
	// example:
	//
	// odps_source
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 1234123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file type. The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 10
	FileType *int64 `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) GetContent() *string {
	return s.Content
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) GetFileId() *int64 {
	return s.FileId
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) GetFileType() *int64 {
	return s.FileType
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) SetContent(v string) *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand {
	s.Content = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) SetDataSourceName(v string) *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand {
	s.DataSourceName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) SetFileId(v int64) *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand {
	s.FileId = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) SetFileType(v int64) *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand {
	s.FileType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailFileExecutionCommand) Validate() error {
	return dara.Validate(s)
}

type GetIDEEventDetailResponseBodyEventDetailTableModel struct {
	// The columns in the table.
	Columns []*GetIDEEventDetailResponseBodyEventDetailTableModelColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The remarks of the table.
	//
	// example:
	//
	// A new table
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the data source to which the table belongs.
	//
	// example:
	//
	// odps_source
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The environment in which the table is used. Valid values:
	//
	// 	- DEV
	//
	// 	- PROD
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The lifecycle of the table. Unit: day.
	//
	// example:
	//
	// 7
	LifeCycle *int64 `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// The path of the table.
	//
	// example:
	//
	// hdfs://path/to/object
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// tb_hello
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailTableModel) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailTableModel) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) GetColumns() []*GetIDEEventDetailResponseBodyEventDetailTableModelColumns {
	return s.Columns
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) GetComment() *string {
	return s.Comment
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) GetEnv() *string {
	return s.Env
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) GetLifeCycle() *int64 {
	return s.LifeCycle
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) GetLocation() *string {
	return s.Location
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) GetTableName() *string {
	return s.TableName
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) SetColumns(v []*GetIDEEventDetailResponseBodyEventDetailTableModelColumns) *GetIDEEventDetailResponseBodyEventDetailTableModel {
	s.Columns = v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) SetComment(v string) *GetIDEEventDetailResponseBodyEventDetailTableModel {
	s.Comment = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) SetDataSourceName(v string) *GetIDEEventDetailResponseBodyEventDetailTableModel {
	s.DataSourceName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) SetEnv(v string) *GetIDEEventDetailResponseBodyEventDetailTableModel {
	s.Env = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) SetLifeCycle(v int64) *GetIDEEventDetailResponseBodyEventDetailTableModel {
	s.LifeCycle = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) SetLocation(v string) *GetIDEEventDetailResponseBodyEventDetailTableModel {
	s.Location = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) SetTableName(v string) *GetIDEEventDetailResponseBodyEventDetailTableModel {
	s.TableName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModel) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetIDEEventDetailResponseBodyEventDetailTableModelColumns struct {
	// The name of the column.
	//
	// example:
	//
	// ID
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// BIGINT
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The remarks of the column.
	//
	// example:
	//
	// ID
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Indicates whether the column is a partition key column. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsPartitionColumn *bool `json:"IsPartitionColumn,omitempty" xml:"IsPartitionColumn,omitempty"`
}

func (s GetIDEEventDetailResponseBodyEventDetailTableModelColumns) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailResponseBodyEventDetailTableModelColumns) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) GetColumnType() *string {
	return s.ColumnType
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) GetComment() *string {
	return s.Comment
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) GetIsPartitionColumn() *bool {
	return s.IsPartitionColumn
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) SetColumnName(v string) *GetIDEEventDetailResponseBodyEventDetailTableModelColumns {
	s.ColumnName = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) SetColumnType(v string) *GetIDEEventDetailResponseBodyEventDetailTableModelColumns {
	s.ColumnType = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) SetComment(v string) *GetIDEEventDetailResponseBodyEventDetailTableModelColumns {
	s.Comment = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) SetIsPartitionColumn(v bool) *GetIDEEventDetailResponseBodyEventDetailTableModelColumns {
	s.IsPartitionColumn = &v
	return s
}

func (s *GetIDEEventDetailResponseBodyEventDetailTableModelColumns) Validate() error {
	return dara.Validate(s)
}
