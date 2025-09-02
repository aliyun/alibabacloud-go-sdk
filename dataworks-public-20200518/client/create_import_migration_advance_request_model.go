// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateImportMigrationAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalculateEngineMap(v string) *CreateImportMigrationAdvanceRequest
	GetCalculateEngineMap() *string
	SetCommitRule(v string) *CreateImportMigrationAdvanceRequest
	GetCommitRule() *string
	SetDescription(v string) *CreateImportMigrationAdvanceRequest
	GetDescription() *string
	SetName(v string) *CreateImportMigrationAdvanceRequest
	GetName() *string
	SetPackageFileObject(v io.Reader) *CreateImportMigrationAdvanceRequest
	GetPackageFileObject() io.Reader
	SetPackageType(v string) *CreateImportMigrationAdvanceRequest
	GetPackageType() *string
	SetProjectId(v int64) *CreateImportMigrationAdvanceRequest
	GetProjectId() *int64
	SetResourceGroupMap(v string) *CreateImportMigrationAdvanceRequest
	GetResourceGroupMap() *string
	SetWorkspaceMap(v string) *CreateImportMigrationAdvanceRequest
	GetWorkspaceMap() *string
}

type CreateImportMigrationAdvanceRequest struct {
	// The mapping between the source compute engine instance and the destination compute engine instance. The following types of compute engine instances are supported: MaxCompute, E-MapReduce (EMR), Hadoop CDH, and Hologres.
	//
	// example:
	//
	// {     "ODPS": {       "zxy_8221431_engine": "wzp_kaifazheban_engine"     },     "EMR": {         "aaaa": "bbb"     }   }
	CalculateEngineMap *string `json:"CalculateEngineMap,omitempty" xml:"CalculateEngineMap,omitempty"`
	// The rule configured for automatically committing and deploying the import task. The rule contains the following parameters:
	//
	// 	- resourceAutoCommit: specifies whether resources are automatically committed. The value true indicates yes and the value false indicates no.
	//
	// 	- resourceAutoDeploy: specifies whether resources are automatically deployed. The value true indicates yes and the value false indicates no.
	//
	// 	- functionAutoCommit: specifies whether the function is automatically committed. The value true indicates yes and the value false indicates no.
	//
	// 	- functionAutoDeploy: specifies whether the function is automatically deployed. The value true indicates yes and the value false indicates no.
	//
	// 	- tableAutoCommitToDev: specifies whether the table is automatically committed to the development environment. The value true indicates yes and the value false indicates no.
	//
	// 	- tableAutoCommitToProd: specifies whether the table is automatically committed to the production environment. The value true indicates yes and the value false indicates no.
	//
	// 	- ignoreLock: specifies whether the lock is automatically ignored when an import task is locked. The value true indicates yes and the value false indicates no. If you set this parameter to true for an import task, you can forcefully update the task even if the task is locked.
	//
	// 	- fileAutoCommit: specifies whether the file is automatically committed. The value true indicates yes and the value false indicates no.
	//
	// 	- fileAutoDeploy: specifies whether the file is automatically deployed. The value true indicates yes and the value false indicates no.
	//
	// example:
	//
	// {     "resourceAutoCommit": false,     "resourceAutoDeploy": false,     "functionAutoCommit": false,     "functionAutoDeploy": false,     "tableAutoCommitToDev": false,     "tableAutoCommitToProd": false,     "ignoreLock": false,     "fileAutoCommit": false,     "fileAutoDeploy": false   }
	CommitRule *string `json:"CommitRule,omitempty" xml:"CommitRule,omitempty"`
	// The description of the import package.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the import task. The name must be unique within the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_import_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path of the import package. **The import package must be uploaded. Example of the upload method:**
	//
	// ```java
	//
	//         Config config = new Config();
	//
	//         config.setAccessKeyId(accessId);
	//
	//         config.setAccessKeySecret(accessKey);
	//
	//         config.setEndpoint(popEndpoint);
	//
	//         config.setRegionId(regionId);
	//
	//         Client client = new Client(config);
	//
	//         CreateImportMigrationAdvanceRequest request = new CreateImportMigrationAdvanceRequest();
	//
	//         request.setName("test_migration_api_" + System.currentTimeMillis());
	//
	//         request.setProjectId(123456L);
	//
	//         request.setPackageType("DATAWORKS_MODEL");
	//
	//         request.setPackageFileObject(new FileInputStream("/home/admin/Downloads/test.zip"));
	//
	//         RuntimeOptions runtime = new RuntimeOptions();
	//
	//         CreateImportMigrationResponse response = client.createImportMigrationAdvance(request, runtime);
	//
	// ```
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/admin/xxx/import.zip
	PackageFileObject io.Reader `json:"PackageFile,omitempty" xml:"PackageFile,omitempty"`
	// The type of the import package. Valid values:
	//
	// 	- DATAWORKS_MODEL (standard format)
	//
	// 	- DATAWORKS_V2 (Apsara Stack DataWorks V3.6.1 to V3.11)
	//
	// 	- DATAWORKS_V3 (Apsara Stack DataWorks V3.12 and later)
	//
	// This parameter is required.
	//
	// example:
	//
	// DATAWORKS_MODEL
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The mapping between the resource group for scheduling and the resource group for Data Integration. The keys and values in the mapping are the identifiers of the resource groups. Specify the mapping in the following format:
	//
	// ```json
	//
	// {
	//
	//     "SCHEDULER_RESOURCE_GROUP": {
	//
	//         "xxx": "yyy"
	//
	//     },
	//
	//     "DI_RESOURCE_GROUP": {
	//
	//         "ccc": "dfdd"
	//
	//     }
	//
	// }
	//
	// ```
	//
	// example:
	//
	// {"SCHEDULER_RESOURCE_GROUP": {"xxx":"yyy"},"DI_RESOURCE_GROUP":{"ccc":"ddd"}}
	ResourceGroupMap *string `json:"ResourceGroupMap,omitempty" xml:"ResourceGroupMap,omitempty"`
	// The mapping between the prefixes for the names of the source and destination workspaces. When the system performs the import operation, the prefix for the name of the source workspace in the import package is replaced based on the mapping.
	//
	// example:
	//
	// {"test_workspace_src": "test_workspace_target"}
	WorkspaceMap *string `json:"WorkspaceMap,omitempty" xml:"WorkspaceMap,omitempty"`
}

func (s CreateImportMigrationAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImportMigrationAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateImportMigrationAdvanceRequest) GetCalculateEngineMap() *string {
	return s.CalculateEngineMap
}

func (s *CreateImportMigrationAdvanceRequest) GetCommitRule() *string {
	return s.CommitRule
}

func (s *CreateImportMigrationAdvanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImportMigrationAdvanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateImportMigrationAdvanceRequest) GetPackageFileObject() io.Reader {
	return s.PackageFileObject
}

func (s *CreateImportMigrationAdvanceRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *CreateImportMigrationAdvanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateImportMigrationAdvanceRequest) GetResourceGroupMap() *string {
	return s.ResourceGroupMap
}

func (s *CreateImportMigrationAdvanceRequest) GetWorkspaceMap() *string {
	return s.WorkspaceMap
}

func (s *CreateImportMigrationAdvanceRequest) SetCalculateEngineMap(v string) *CreateImportMigrationAdvanceRequest {
	s.CalculateEngineMap = &v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) SetCommitRule(v string) *CreateImportMigrationAdvanceRequest {
	s.CommitRule = &v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) SetDescription(v string) *CreateImportMigrationAdvanceRequest {
	s.Description = &v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) SetName(v string) *CreateImportMigrationAdvanceRequest {
	s.Name = &v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) SetPackageFileObject(v io.Reader) *CreateImportMigrationAdvanceRequest {
	s.PackageFileObject = v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) SetPackageType(v string) *CreateImportMigrationAdvanceRequest {
	s.PackageType = &v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) SetProjectId(v int64) *CreateImportMigrationAdvanceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) SetResourceGroupMap(v string) *CreateImportMigrationAdvanceRequest {
	s.ResourceGroupMap = &v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) SetWorkspaceMap(v string) *CreateImportMigrationAdvanceRequest {
	s.WorkspaceMap = &v
	return s
}

func (s *CreateImportMigrationAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
